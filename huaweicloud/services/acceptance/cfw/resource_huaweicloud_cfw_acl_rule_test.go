package cfw

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cfw"
)

func getACLRuleResourceFunc(conf *config.Config, state *terraform.ResourceState) (interface{}, error) {
	region := acceptance.HW_REGION_NAME
	client, err := conf.NewServiceClient("cfw", region)
	if err != nil {
		return nil, fmt.Errorf("error creating CFW client: %s", err)
	}

	return cfw.GetACLRule(client, state.Primary.ID, state.Primary.Attributes["object_id"])
}

func TestAccACLRule_basic(t *testing.T) {
	var obj interface{}

	name := acceptance.RandomAccResourceName()
	rName := "huaweicloud_cfw_acl_rule.r1"

	rc := acceptance.InitResourceCheck(
		rName,
		&obj,
		getACLRuleResourceFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckCfw(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccACLRule_basic(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "description", "terraform test"),
					resource.TestCheckResourceAttr(rName, "type", "0"),
					resource.TestCheckResourceAttr(rName, "address_type", "0"),
					resource.TestCheckResourceAttr(rName, "action_type", "0"),
					resource.TestCheckResourceAttr(rName, "applications.#", "1"),
					resource.TestCheckResourceAttr(rName, "long_connect_enable", "0"),
					resource.TestCheckResourceAttr(rName, "status", "1"),
					resource.TestCheckResourceAttr(rName, "direction", "0"),
					resource.TestCheckResourceAttr(rName, "source_addresses.0", "1.1.1.1"),
					resource.TestCheckResourceAttr(rName, "destination_addresses.0", "1.1.1.2"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.protocol", "6"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.source_port", "81"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.dest_port", "82"),
					resource.TestCheckResourceAttr(rName, "tags.key", "value"),
					resource.TestCheckResourceAttrSet(rName, "rule_hit_count"),
				),
			},
			{
				Config: testAccACLRule_basic_update(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "applications.#", "2"),
					resource.TestCheckResourceAttr(rName, "source_addresses.#", "2"),
					resource.TestCheckResourceAttr(rName, "destination_addresses.#", "2"),
					resource.TestCheckResourceAttr(rName, "source_addresses.0", "1.1.1.2"),
					resource.TestCheckResourceAttr(rName, "source_addresses.1", "3.3.3.3"),
					resource.TestCheckResourceAttr(rName, "destination_addresses.0", "2.2.2.2"),
					resource.TestCheckResourceAttr(rName, "destination_addresses.1", "4.4.4.4"),
					resource.TestCheckResourceAttr(rName, "custom_services.#", "2"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.protocol", "6"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.dest_port", "84"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.source_port", "83"),
					resource.TestCheckResourceAttr(rName, "custom_services.1.protocol", "6"),
					resource.TestCheckResourceAttr(rName, "custom_services.1.source_port", "85"),
					resource.TestCheckResourceAttr(rName, "custom_services.1.dest_port", "86"),
					resource.TestCheckResourceAttr(rName, "tags.k1", "v1"),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testACLRuleImportState(rName),
				ImportStateVerifyIgnore: []string{
					"sequence", "type",
				},
			},
		},
	})
}

func TestAccACLRule_serviceGroups_regionList(t *testing.T) {
	var obj interface{}

	name := acceptance.RandomAccResourceName()
	rName := "huaweicloud_cfw_acl_rule.test"

	rc := acceptance.InitResourceCheck(
		rName,
		&obj,
		getACLRuleResourceFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckCfw(t)
			acceptance.TestAccPreCheckCfwPredefinedServiceGroup(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccACLRule_serviceGroups_regionList(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "description", "terraform test"),
					resource.TestCheckResourceAttr(rName, "type", "0"),
					resource.TestCheckResourceAttr(rName, "address_type", "0"),
					resource.TestCheckResourceAttr(rName, "action_type", "0"),
					resource.TestCheckResourceAttr(rName, "long_connect_enable", "0"),
					resource.TestCheckResourceAttr(rName, "status", "1"),
					resource.TestCheckResourceAttr(rName, "direction", "0"),
					resource.TestCheckResourceAttr(rName, "source_region_list.#", "1"),
					resource.TestCheckResourceAttr(rName, "source_region_list.0.description_cn", "中国"),
					resource.TestCheckResourceAttr(rName, "source_region_list.0.description_en", "Chinese Mainland"),
					resource.TestCheckResourceAttr(rName, "source_region_list.0.region_id", "CN"),
					resource.TestCheckResourceAttr(rName, "source_region_list.0.region_type", "0"),
					resource.TestCheckResourceAttr(rName, "destination_addresses.0", "1.1.1.2"),
					resource.TestCheckResourceAttrPair(rName, "custom_service_groups.0.group_ids.0", "huaweicloud_cfw_service_group.s1", "id"),
					resource.TestCheckResourceAttr(rName, "custom_service_groups.0.protocols.0", "6"),
					resource.TestCheckResourceAttr(rName, "tags.key", "value"),
					resource.TestCheckResourceAttrSet(rName, "rule_hit_count"),
				),
			},
			{
				Config: testAccACLRule_serviceGroups_regionList_update(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "direction", "1"),
					resource.TestCheckResourceAttr(rName, "source_addresses.0", "1.1.1.2"),
					resource.TestCheckResourceAttr(rName, "destination_region_list.#", "2"),
					resource.TestCheckResourceAttr(rName, "destination_region_list.0.description_cn", "中国"),
					resource.TestCheckResourceAttr(rName, "destination_region_list.0.description_en", "Chinese Mainland"),
					resource.TestCheckResourceAttr(rName, "destination_region_list.0.region_id", "CN"),
					resource.TestCheckResourceAttr(rName, "destination_region_list.0.region_type", "0"),
					resource.TestCheckResourceAttr(rName, "destination_region_list.1.description_cn", "希腊"),
					resource.TestCheckResourceAttr(rName, "destination_region_list.1.description_en", "Greece"),
					resource.TestCheckResourceAttr(rName, "destination_region_list.1.region_id", "GR"),
					resource.TestCheckResourceAttr(rName, "destination_region_list.1.region_type", "0"),
					resource.TestCheckResourceAttr(rName, "custom_service_groups.#", "1"),
					resource.TestCheckResourceAttr(rName, "custom_service_groups.0.group_ids.#", "2"),
					resource.TestCheckResourceAttrPair(rName, "custom_service_groups.0.group_ids.0", "huaweicloud_cfw_service_group.s2", "id"),
					resource.TestCheckResourceAttrPair(rName, "custom_service_groups.0.group_ids.1", "huaweicloud_cfw_service_group.s3", "id"),
					resource.TestCheckResourceAttr(rName, "custom_service_groups.0.protocols.0", "6"),
					resource.TestCheckResourceAttr(rName, "tags.k1", "v1"),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testACLRuleImportState(rName),
				ImportStateVerifyIgnore: []string{
					"sequence", "type", "predefined_service_groups",
				},
			},
		},
	})
}

func TestAccACLRule_domainAddressName_domainGroup(t *testing.T) {
	var obj interface{}

	name := acceptance.RandomAccResourceName()
	rName := "huaweicloud_cfw_acl_rule.test"

	rc := acceptance.InitResourceCheck(
		rName,
		&obj,
		getACLRuleResourceFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckCfw(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccACLRule_domainAddressName_domainGroup_step1(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "description", "terraform test"),
					resource.TestCheckResourceAttr(rName, "type", "0"),
					resource.TestCheckResourceAttr(rName, "address_type", "0"),
					resource.TestCheckResourceAttr(rName, "action_type", "0"),
					resource.TestCheckResourceAttr(rName, "long_connect_enable", "0"),
					resource.TestCheckResourceAttr(rName, "status", "1"),
					resource.TestCheckResourceAttr(rName, "direction", "1"),
					resource.TestCheckResourceAttr(rName, "source_addresses.#", "1"),
					resource.TestCheckResourceAttr(rName, "source_addresses.0", "1.1.1.1"),
					resource.TestCheckResourceAttr(rName, "destination_domain_address_name", "*.baidu.com"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.protocol", "6"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.source_port", "83"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.dest_port", "84"),
					resource.TestCheckResourceAttrSet(rName, "rule_hit_count"),
				),
			},
			{
				Config: testAccACLRule_domainAddressName_domainGroup_step2(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "direction", "1"),
					resource.TestCheckResourceAttr(rName, "source_addresses.#", "1"),
					resource.TestCheckResourceAttr(rName, "source_addresses.0", "1.1.1.2"),
					resource.TestCheckResourceAttr(rName, "destination_domain_address_name", "www.baidu.com"),
				),
			},
			{
				Config: testAccACLRule_domainAddressName_domainGroup_step3(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "direction", "1"),
					resource.TestCheckResourceAttr(rName, "source_addresses.#", "1"),
					resource.TestCheckResourceAttr(rName, "source_addresses.0", "1.1.1.2"),
					resource.TestCheckResourceAttrPair(rName, "destination_domain_group_id", "huaweicloud_cfw_domain_name_group.dg1", "id"),
					resource.TestCheckResourceAttrPair(rName, "destination_domain_group_name", "huaweicloud_cfw_domain_name_group.dg1", "name"),
					resource.TestCheckResourceAttr(rName, "destination_domain_group_type", "4"),
				),
			},
			{
				Config: testAccACLRule_domainAddressName_domainGroup_step4(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "direction", "1"),
					resource.TestCheckResourceAttr(rName, "source_addresses.#", "1"),
					resource.TestCheckResourceAttr(rName, "source_addresses.0", "1.1.1.2"),
					resource.TestCheckResourceAttrPair(rName, "destination_domain_group_id", "huaweicloud_cfw_domain_name_group.dg2", "id"),
					resource.TestCheckResourceAttrPair(rName, "destination_domain_group_name", "huaweicloud_cfw_domain_name_group.dg2", "name"),
					resource.TestCheckResourceAttr(rName, "destination_domain_group_type", "4"),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testACLRuleImportState(rName),
				ImportStateVerifyIgnore: []string{
					"sequence", "type",
				},
			},
		},
	})
}

func TestAccACLRule_addressGroups(t *testing.T) {
	var obj interface{}

	name := acceptance.RandomAccResourceName()
	rName := "huaweicloud_cfw_acl_rule.test"

	rc := acceptance.InitResourceCheck(
		rName,
		&obj,
		getACLRuleResourceFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckCfw(t)
			acceptance.TestAccPreCheckCfwPredefinedAddressGroup(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccACLRule_addressGroups_step1(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "description", "terraform test"),
					resource.TestCheckResourceAttr(rName, "type", "0"),
					resource.TestCheckResourceAttr(rName, "address_type", "0"),
					resource.TestCheckResourceAttr(rName, "action_type", "0"),
					resource.TestCheckResourceAttr(rName, "long_connect_enable", "0"),
					resource.TestCheckResourceAttr(rName, "status", "1"),
					resource.TestCheckResourceAttr(rName, "source_address_groups.#", "2"),
					resource.TestCheckResourceAttrPair(rName, "source_address_groups.0", "huaweicloud_cfw_address_group.g1", "id"),
					resource.TestCheckResourceAttrPair(rName, "source_address_groups.1", "huaweicloud_cfw_address_group.g2", "id"),
					resource.TestCheckResourceAttr(rName, "destination_addresses.0", "1.1.1.1"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.protocol", "6"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.source_port", "80"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.dest_port", "80"),
					resource.TestCheckResourceAttrSet(rName, "rule_hit_count"),
				),
			},
			{
				Config: testAccACLRule_addressGroups_step2(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "destination_addresses.0", "1.1.1.1"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.protocol", "6"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.source_port", "80"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.dest_port", "80"),
				),
			},
			{
				Config: testAccACLRule_addressGroups_step3(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "source_addresses.0", "1.1.1.1"),
					resource.TestCheckResourceAttr(rName, "destination_address_groups.#", "2"),
					resource.TestCheckResourceAttrPair(rName, "destination_address_groups.0", "huaweicloud_cfw_address_group.g1", "id"),
					resource.TestCheckResourceAttrPair(rName, "destination_address_groups.1", "huaweicloud_cfw_address_group.g2", "id"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.protocol", "6"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.source_port", "80"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.dest_port", "80"),
				),
			},
			{
				Config: testAccACLRule_addressGroups_step4(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "source_addresses.0", "1.1.1.1"),
					resource.TestCheckResourceAttr(rName, "destination_address_groups.#", "1"),
					resource.TestCheckResourceAttrPair(rName, "destination_address_groups.0", "huaweicloud_cfw_address_group.g1", "id"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.protocol", "6"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.source_port", "80"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.dest_port", "80"),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testACLRuleImportState(rName),
				ImportStateVerifyIgnore: []string{
					"sequence", "type", "source_predefined_groups",
				},
			},
		},
	})
}

func TestAccACLRule_anyAddress_anyService(t *testing.T) {
	var obj interface{}

	name := acceptance.RandomAccResourceName()
	rName := "huaweicloud_cfw_acl_rule.test"

	rc := acceptance.InitResourceCheck(
		rName,
		&obj,
		getACLRuleResourceFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckCfw(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccACLRule_anyAddress_anyService(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "description", "terraform test"),
					resource.TestCheckResourceAttr(rName, "applications.#", "1"),
					resource.TestCheckResourceAttr(rName, "applications.0", "HTTPS"),
					resource.TestCheckResourceAttr(rName, "destination_addresses.0", "1.1.1.2"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.protocol", "6"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.source_port", "80"),
					resource.TestCheckResourceAttr(rName, "custom_services.0.dest_port", "80"),
				),
			},
			{
				Config: testAccACLRule_anyAddress_anyService_update(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "applications.#", "1"),
					resource.TestCheckResourceAttr(rName, "applications.0", "ANY"),
					resource.TestCheckResourceAttr(rName, "source_addresses.0", "1.1.1.1"),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testACLRuleImportState(rName),
				ImportStateVerifyIgnore: []string{
					"sequence", "type",
				},
			},
		},
	})
}

func testACLRuleImportState(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Resource (%s) not found: %s", name, rs)
		}
		if rs.Primary.Attributes["object_id"] == "" {
			return "", fmt.Errorf("Attribute (object_id) of Resource (%s) not found: %s", name, rs)
		}
		if rs.Primary.ID == "" {
			return "", fmt.Errorf("Attribute (ID) of Resource (%s) not found: %s", name, rs)
		}

		return rs.Primary.Attributes["object_id"] + "/" +
			rs.Primary.ID, nil
	}
}

func testAccACLRule_basic(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_cfw_acl_rule" "r1" {
  name                = "%[2]s"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  applications        = ["HTTPS"]
  long_connect_enable = 0
  status              = 1

  source_addresses      = ["1.1.1.1"]
  destination_addresses = ["1.1.1.2"]

  custom_services {
    protocol    = 6
    source_port = 81
    dest_port   = 82
  }

  sequence {
    top = 1
  }

  tags = {
    key = "value"
  }

  depends_on = [
    huaweicloud_cfw_acl_rule.r2,
  ]
}

resource "huaweicloud_cfw_acl_rule" "r2" {
  name                = "%[2]s2"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1
  
  source_addresses      = ["2.2.2.2"]
  destination_addresses = ["3.3.3.3"]
  
  custom_services {
    protocol    = 6
    source_port = 81
    dest_port   = 82
  }
  
  sequence {
    top = 1
  }
  
  tags = {
    key = "value"
  }
}
`, testAccDatasourceFirewalls_basic(), name)
}

func testAccACLRule_basic_update(name string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_cfw_acl_rule" "r1" {
  name                = "%s"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  applications        = ["HTTPS", "HTTP"]
  long_connect_enable = 0
  status              = 1

  source_addresses      = ["1.1.1.2", "3.3.3.3"]
  destination_addresses = ["2.2.2.2", "4.4.4.4"]

  custom_services {
    protocol    = 6
    source_port = 83
    dest_port   = 84
  }

  custom_services {
    protocol    = 6
    source_port = 85
    dest_port   = 86
  }

  sequence {
    top          = 0
    dest_rule_id = huaweicloud_cfw_acl_rule.r2.id
  }

  tags = {
    k1 = "v1"
  }

  depends_on = [
    huaweicloud_cfw_acl_rule.r2,
  ]
}

resource "huaweicloud_cfw_acl_rule" "r2" {
  name                = "%[2]s2"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1
  
  source_addresses      = ["2.2.2.2"]
  destination_addresses = ["3.3.3.3"]
  
  custom_services {
    protocol    = 6
    source_port = 81
    dest_port   = 82
  }
  
  sequence {
    top = 1
  }
  
  tags = {
    key = "value"
  }
}
`, testAccDatasourceFirewalls_basic(), name)
}

func testAccACLRule_serviceGroups_regionList(name string) string {
	return fmt.Sprintf(`
%[1]s

%[2]s

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = "%[3]s"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1

  source_region_list {
    description_cn = "中国"
    description_en = "Chinese Mainland"
    region_id      = "CN"
    region_type    = 0
  }

  destination_addresses = ["1.1.1.2"]

  custom_service_groups {
    protocols = [6]
    group_ids = [
      huaweicloud_cfw_service_group.s1.id,	
    ]
  }

  predefined_service_groups {
    protocols = [6]
    group_ids = [
      "%[4]s",
    ]
  }

  sequence {
    top = 1
  }

  tags = {
    key = "value"
  }
}
`, testAccDatasourceFirewalls_basic(), testAccACLRule_advanced_base(name), name, acceptance.HW_CFW_PREDEFINED_SERVICE_GROUP1)
}

func testAccACLRule_serviceGroups_regionList_update(name string) string {
	return fmt.Sprintf(`
%[1]s

%[2]s

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = "%[3]s"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1
  direction           = 1

  source_addresses = ["1.1.1.2"]

  destination_region_list {
    description_cn = "中国"
    description_en = "Chinese Mainland"
    region_id      = "CN"
    region_type    = 0
  }

  destination_region_list {
    description_cn = "希腊"
    description_en = "Greece"
    region_id      = "GR"
    region_type    = 0
  }

  custom_service_groups {
    protocols = [6]
    group_ids = [
      huaweicloud_cfw_service_group.s2.id,
      huaweicloud_cfw_service_group.s3.id,
    ]
  }

  predefined_service_groups {
    protocols = [6]
    group_ids = [
      "%[4]s",
    ]
  }

  sequence {
    top = 1
  }

  tags = {
    k1 = "v1"
  }
}
`, testAccDatasourceFirewalls_basic(), testAccACLRule_advanced_base(name), name, acceptance.HW_CFW_PREDEFINED_SERVICE_GROUP2)
}

func testAccACLRule_domainAddressName_domainGroup_step1(name string) string {
	return fmt.Sprintf(`
%[1]s

%[2]s

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = "%[3]s"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1
  direction           = 1

  source_addresses                = ["1.1.1.1"]
  destination_domain_address_name = "*.baidu.com"

  custom_services {
    protocol    = 6
    source_port = 83
    dest_port   = 84
  }

  sequence {
    top = 1
  }
}
`, testAccDatasourceFirewalls_basic(), testAccACLRule_advanced_base(name), name)
}

func testAccACLRule_domainAddressName_domainGroup_step2(name string) string {
	return fmt.Sprintf(`
%[1]s

%[2]s

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = "%[3]s"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1
  direction           = 1

  source_addresses                = ["1.1.1.2"]
  destination_domain_address_name = "www.baidu.com"

  custom_services {
    protocol    = 6
    source_port = 83
    dest_port   = 84
  }

  sequence {
    top = 1
  }
}
`, testAccDatasourceFirewalls_basic(), testAccACLRule_advanced_base(name), name)
}

func testAccACLRule_domainAddressName_domainGroup_step3(name string) string {
	return fmt.Sprintf(`
%[1]s

%[2]s

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = "%[3]s"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1
  direction           = 1

  source_addresses              = ["1.1.1.2"]
  destination_domain_group_id   = huaweicloud_cfw_domain_name_group.dg1.id
  destination_domain_group_name = huaweicloud_cfw_domain_name_group.dg1.name
  destination_domain_group_type = 4

  custom_services {
    protocol    = 6
    source_port = 83
    dest_port   = 84
  }

  sequence {
    top = 1
  }
}
`, testAccDatasourceFirewalls_basic(), testAccACLRule_advanced_base(name), name)
}

func testAccACLRule_domainAddressName_domainGroup_step4(name string) string {
	return fmt.Sprintf(`
%[1]s

%[2]s

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = "%[3]s"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1
  direction           = 1

  source_addresses              = ["1.1.1.2"]
  destination_domain_group_id   = huaweicloud_cfw_domain_name_group.dg2.id
  destination_domain_group_name = huaweicloud_cfw_domain_name_group.dg2.name
  destination_domain_group_type = 4

  custom_services {
    protocol    = 6
    source_port = 83
    dest_port   = 84
  }

  sequence {
    top = 1
  }
}
`, testAccDatasourceFirewalls_basic(), testAccACLRule_advanced_base(name), name)
}

func testAccACLRule_addressGroups_step1(name string) string {
	return fmt.Sprintf(`
%[1]s

%[2]s

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = "%[3]s"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1
  
  source_address_groups = [
    huaweicloud_cfw_address_group.g1.id,
    huaweicloud_cfw_address_group.g2.id,
  ]

  destination_addresses = ["1.1.1.1"]

  custom_services {
    protocol    = 6
    source_port = 80
    dest_port   = 80
  }

  sequence {
    bottom = 1
  }
}
	`, testAccDatasourceFirewalls_basic(), testAccACLRule_advanced_base(name), name)
}

func testAccACLRule_addressGroups_step2(name string) string {
	return fmt.Sprintf(`
%[1]s

%[2]s

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = "%[3]s"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1
  
  source_predefined_groups = [
    "%[4]s",
    "%[5]s",
  ]

  destination_addresses = ["1.1.1.1"]

  custom_services {
    protocol    = 6
    source_port = 80
    dest_port   = 80
  }

  sequence {
    bottom = 1
  }
}
`, testAccDatasourceFirewalls_basic(), testAccACLRule_advanced_base(name), name,
		acceptance.HW_CFW_PREDEFINED_ADDRESS_GROUP1, acceptance.HW_CFW_PREDEFINED_ADDRESS_GROUP2)
}

func testAccACLRule_addressGroups_step3(name string) string {
	return fmt.Sprintf(`
%[1]s

%[2]s

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = "%[3]s"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1
  
  source_addresses = ["1.1.1.1"]
  
  destination_address_groups = [
    huaweicloud_cfw_address_group.g1.id,
    huaweicloud_cfw_address_group.g2.id,
  ]

  custom_services {
    protocol    = 6
    source_port = 80
    dest_port   = 80
  }

  sequence {
    bottom = 1
  }
}
`, testAccDatasourceFirewalls_basic(), testAccACLRule_advanced_base(name), name,
		acceptance.HW_CFW_PREDEFINED_ADDRESS_GROUP1, acceptance.HW_CFW_PREDEFINED_ADDRESS_GROUP2)
}

func testAccACLRule_addressGroups_step4(name string) string {
	return fmt.Sprintf(`
%[1]s

%[2]s

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = "%[3]s"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1

  source_addresses = ["1.1.1.1"]
  
  destination_address_groups = [
    huaweicloud_cfw_address_group.g1.id,
  ]

  custom_services {
    protocol    = 6
    source_port = 80
    dest_port   = 80
  }

  sequence {
    bottom = 1
  }
}
`, testAccDatasourceFirewalls_basic(), testAccACLRule_advanced_base(name), name,
		acceptance.HW_CFW_PREDEFINED_ADDRESS_GROUP1, acceptance.HW_CFW_PREDEFINED_ADDRESS_GROUP2)
}

func testAccACLRule_anyAddress_anyService(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = "%[2]s"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  applications        = ["HTTPS"]
  long_connect_enable = 0
  status              = 1
  
  destination_addresses = ["1.1.1.2"]

  custom_services {
    protocol    = 6
    source_port = 80
    dest_port   = 80
  }
  
  sequence {
    top = 1
  }
  
  tags = {
    key = "value"
  }
}
`, testAccDatasourceFirewalls_basic(), name)
}

func testAccACLRule_anyAddress_anyService_update(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = "%[2]s"
  object_id           = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  description         = "terraform test"
  type                = 0
  address_type        = 0
  action_type         = 0
  applications        = ["ANY"]
  long_connect_enable = 0
  status              = 1
  
  source_addresses    = ["1.1.1.1"]
  
  sequence {
    top = 1
  }
  
  tags = {
    key = "value"
  }
}
`, testAccDatasourceFirewalls_basic(), name)
}

func testAccACLRule_advanced_base(name string) string {
	return fmt.Sprintf(`
resource "huaweicloud_cfw_address_group" "g1" {
  object_id   = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  name        = "%[1]s_ag1"
  description = "address group 1"
}

resource "huaweicloud_cfw_address_group" "g2" {
  object_id   = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  name        = "%[1]s_ag2"
  description = "address group 2"
}

resource "huaweicloud_cfw_address_group" "g3" {
  object_id   = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  name        = "%[1]s_ag3"
  description = "address group 3"
}

resource "huaweicloud_cfw_service_group" "s1" {
  object_id   = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  name        = "%[1]s_sg1"
  description = "service group 1"
}

resource "huaweicloud_cfw_service_group" "s2" {
  object_id   = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  name        = "%[1]s_sg2"
  description = "service group 2"
}

resource "huaweicloud_cfw_service_group" "s3" {
  object_id   = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  name        = "%[1]s_sg3"
  description = "service group 3"
}

resource "huaweicloud_cfw_service_group_member" "m1" {
  group_id    = huaweicloud_cfw_service_group.s1.id
  protocol    = 6
  source_port = "80"
  dest_port   = "22"
}

resource "huaweicloud_cfw_service_group_member" "m2" {
  group_id    = huaweicloud_cfw_service_group.s2.id
  protocol    = 6
  source_port = "81"
  dest_port   = "23"
}

resource "huaweicloud_cfw_service_group_member" "m3" {
  group_id    = huaweicloud_cfw_service_group.s3.id
  protocol    = 6
  source_port = "82"
  dest_port   = "24"
}

resource "huaweicloud_cfw_domain_name_group" "dg1" {
  fw_instance_id = data.huaweicloud_cfw_firewalls.test.records[0].fw_instance_id
  object_id      = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  name           = "%[1]s_dg1"
  type           = 0
  description    = "created by terraform"
  
  domain_names {
    domain_name = "www.cfw-test1.com"
    description = "test domain 1"
  }
}

resource "huaweicloud_cfw_domain_name_group" "dg2" {
  fw_instance_id = data.huaweicloud_cfw_firewalls.test.records[0].fw_instance_id
  object_id      = data.huaweicloud_cfw_firewalls.test.records[0].protect_objects[0].object_id
  name           = "%[1]s_dg2"
  type           = 0
  description    = "created by terraform"
  
  domain_names {
    domain_name = "www.cfw-test2.com"
    description = "test domain 2"
  }
}
`, name)
}
