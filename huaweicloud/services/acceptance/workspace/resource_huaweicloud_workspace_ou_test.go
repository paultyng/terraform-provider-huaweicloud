package workspace

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/workspace"
)

func getOuFun(conf *config.Config, state *terraform.ResourceState) (interface{}, error) {
	client, err := conf.NewServiceClient("workspace", acceptance.HW_REGION_NAME)
	if err != nil {
		return nil, fmt.Errorf("error creating Workspace client: %s", err)
	}

	return workspace.GetOuByName(client, state.Primary.Attributes["name"])
}

func TestAccOu_basic(t *testing.T) {
	var (
		ouObj        interface{}
		resourceName = "huaweicloud_workspace_ou.test"
		rc           = acceptance.InitResourceCheck(
			resourceName,
			&ouObj,
			getOuFun,
		)
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckWorkspaceAD(t)
			acceptance.TestAccPreCheckWorkspaceOUName(t)
			acceptance.TestAccPreCheckWorkspaceOUUpdateName(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccOu_step1(),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_WORKSPACE_OU_NAME),
					resource.TestCheckResourceAttrPair(resourceName, "domain",
						"huaweicloud_workspace_service.test", "ad_domain.0.name"),
					resource.TestCheckResourceAttr(resourceName, "description", "Created by terraform script"),
				),
			},
			{
				Config: testAccOu_step2(),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_WORKSPACE_OU_UPDATE_NAME),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccOuImportStateFunc(resourceName),
			},
		},
	})
}

func testAccOu_base() string {
	return fmt.Sprintf(`
resource "huaweicloud_workspace_service" "test" {
  ad_domain {
    name                   = try(element(regexall("\\w+\\.(.*)", element(split(",", "%[1]s"), 0))[0], 0), "")
    active_domain_name     = element(split(",", "%[1]s"), 0)
    active_domain_ip       = element(split(",", "%[2]s"), 0)
    active_dns_ip          = element(split(",", "%[2]s"), 0)
    admin_account          = "%[3]s"
    password               = "%[4]s"
    delete_computer_object = true
  }

  auth_type   = "LOCAL_AD"
  access_mode = "INTERNET"
  vpc_id      = "%[5]s"
  network_ids = ["%[6]s"]
}
`, acceptance.HW_WORKSPACE_AD_DOMAIN_NAMES,
		acceptance.HW_WORKSPACE_AD_DOMAIN_IPS,
		acceptance.HW_WORKSPACE_AD_SERVER_ACCOUNT,
		acceptance.HW_WORKSPACE_AD_SERVER_PWD,
		acceptance.HW_WORKSPACE_AD_VPC_ID,
		acceptance.HW_WORKSPACE_AD_NETWORK_ID)
}

func testAccOu_step1() string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_workspace_ou" "test" {
  depends_on = [huaweicloud_workspace_service.test]

  name        = "%[2]s"
  domain      = huaweicloud_workspace_service.test.ad_domain[0].name
  description = "Created by terraform script"
}
`, testAccOu_base(), acceptance.HW_WORKSPACE_OU_NAME)
}

func testAccOu_step2() string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_workspace_ou" "test" {
  depends_on = [huaweicloud_workspace_service.test]

  name   = "%[2]s"
  domain = huaweicloud_workspace_service.test.ad_domain[0].name
}
`, testAccOu_base(), acceptance.HW_WORKSPACE_OU_UPDATE_NAME)
}

func testAccOuImportStateFunc(rName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[rName]
		if !ok {
			return "", fmt.Errorf("resource (%s) not found: %s", rName, rs)
		}
		ouName := rs.Primary.Attributes["name"]
		if ouName == "" {
			return "", fmt.Errorf("invalid format specified for import ID, want '<name>', but got '%s'", ouName)
		}
		return ouName, nil
	}
}
