package cce

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/chnsz/golangsdk/openstack/cce/v3/addons"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func getAddonFunc(cfg *config.Config, state *terraform.ResourceState) (interface{}, error) {
	client, err := cfg.CceAddonV3Client(acceptance.HW_REGION_NAME)
	if err != nil {
		return nil, fmt.Errorf("error creating CCE v3 client: %s", err)
	}
	return addons.Get(client, state.Primary.ID, state.Primary.Attributes["cluster_id"]).Extract()
}

func TestAccAddon_basic(t *testing.T) {
	var (
		addon addons.Addon

		name         = acceptance.RandomAccResourceNameWithDash()
		resourceName = "huaweicloud_cce_addon.test"

		rc = acceptance.InitResourceCheck(
			resourceName,
			&addon,
			getAddonFunc,
		)
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccAddon_basic(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(resourceName, "cluster_id",
						"huaweicloud_cce_cluster.test", "id"),
					resource.TestCheckResourceAttrSet(resourceName, "version"),
					resource.TestCheckResourceAttr(resourceName, "template_name", "metrics-server"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccAddonImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccAddonImportStateIdFunc(resName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		var clusterId, addonId string
		rs, ok := s.RootModule().Resources[resName]
		if !ok {
			return "", fmt.Errorf("the resource (%s) of CCE add-on is not found in the tfstate", resName)
		}
		clusterId = rs.Primary.Attributes["cluster_id"]
		addonId = rs.Primary.ID
		if clusterId == "" || addonId == "" {
			return "", fmt.Errorf("the CCE add-on ID is not exist or related CCE cluster ID is missing")
		}
		return fmt.Sprintf("%s/%s", clusterId, addonId), nil
	}
}

func testAccAddon_Base(rName string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_cce_node" "test" {
  cluster_id        = huaweicloud_cce_cluster.test.id
  name              = "%[2]s"
  flavor_id         = data.huaweicloud_compute_flavors.test.ids[0]
  availability_zone = data.huaweicloud_availability_zones.test.names[0]
  key_pair          = huaweicloud_kps_keypair.test.name

  root_volume {
    size       = 40
    volumetype = "SSD"
  }
  data_volumes {
    size       = 100
    volumetype = "SSD"
  }
}
`, testAccNode_Base(rName), rName)
}

func testAccAddon_basic(rName string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_cce_addon" "test" {
  cluster_id    = huaweicloud_cce_cluster.test.id
  template_name = "metrics-server"
  depends_on    = [huaweicloud_cce_node.test]
}
`, testAccAddon_Base(rName))
}

func TestAccAddon_values(t *testing.T) {
	var (
		addon addons.Addon

		name         = acceptance.RandomAccResourceNameWithDash()
		resourceName = "huaweicloud_cce_addon.test"

		rc = acceptance.InitResourceCheck(
			resourceName,
			&addon,
			getAddonFunc,
		)
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckProjectID(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccAddon_values_step1(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(resourceName, "cluster_id",
						"huaweicloud_cce_cluster.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "version", "1.25.21"),
					resource.TestCheckResourceAttr(resourceName, "template_name", "autoscaler"),
					resource.TestCheckResourceAttrSet(resourceName, "status"),
				),
			},
			{
				Config: testAccAddon_values_step2(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					// the values not set, only check if the updating request is successful
					resource.TestCheckResourceAttrPair(resourceName, "cluster_id",
						"huaweicloud_cce_cluster.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "version", "1.25.21"),
					resource.TestCheckResourceAttr(resourceName, "template_name", "autoscaler"),
					resource.TestCheckResourceAttrSet(resourceName, "status"),
				),
			},
		},
	})
}

func testAccAddon_values_base(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_cce_node_pool" "test" {
  cluster_id         = huaweicloud_cce_cluster.test.id
  name               = "%[2]s"
  os                 = "EulerOS 2.9"
  flavor_id          = data.huaweicloud_compute_flavors.test.ids[0]
  initial_node_count = 4
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  key_pair           = huaweicloud_kps_keypair.test.name
  scale_enable       = true
  min_node_count     = 2
  max_node_count     = 10
  priority           = 1
  type               = "vm"

  root_volume {
    size       = 40
    volumetype = "SSD"
  }
  data_volumes {
    size       = 100
    volumetype = "SSD"
  }
}

data "huaweicloud_cce_addon_template" "test" {
  cluster_id = huaweicloud_cce_cluster.test.id
  name       = "autoscaler"
  version    = "1.25.21"
}
`, testAccNodePool_base(name), name)
}

func testAccAddon_values_step1(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_cce_addon" "test" {
  depends_on = [
    huaweicloud_cce_node_pool.test,
  ]

  cluster_id    = huaweicloud_cce_cluster.test.id
  template_name = "autoscaler"
  version       = "1.25.21"

  values {
    basic       = jsondecode(data.huaweicloud_cce_addon_template.test.spec).basic
    custom_json = jsonencode(merge(
      jsondecode(data.huaweicloud_cce_addon_template.test.spec).parameters.custom,
      {
        cluster_id = huaweicloud_cce_cluster.test.id
        tenant_id  = "%[2]s"
        logLevel   = 3
      }
    ))
    flavor_json = jsonencode(jsondecode(data.huaweicloud_cce_addon_template.test.spec).parameters.flavor1)
  }
}
`, testAccAddon_values_base(name), acceptance.HW_PROJECT_ID)
}

func testAccAddon_values_step2(rName string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_cce_addon" "test" {
  depends_on = [
    huaweicloud_cce_node_pool.test,
  ]

  cluster_id    = huaweicloud_cce_cluster.test.id
  template_name = "autoscaler"
  version       = "1.25.21"

  values {
    basic       = jsondecode(data.huaweicloud_cce_addon_template.test.spec).basic
    custom_json = jsonencode(merge(
      jsondecode(data.huaweicloud_cce_addon_template.test.spec).parameters.custom,
      {
        cluster_id = huaweicloud_cce_cluster.test.id
        tenant_id  = "%[2]s"
        logLevel   = 4
      }
    ))
    flavor_json = jsonencode(jsondecode(data.huaweicloud_cce_addon_template.test.spec).parameters.flavor2)
  }
}
`, testAccAddon_values_base(rName), acceptance.HW_PROJECT_ID)
}
