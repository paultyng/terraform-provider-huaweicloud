// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://www.github.com/huaweicloud/magic-modules
//
// ----------------------------------------------------------------------------

package huaweicloud

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/huaweicloud/golangsdk"
)

func TestAccRdsInstanceV3_basic(t *testing.T) {
	name := acctest.RandString(10)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRdsInstanceV3Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRdsInstanceV3_basic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdsInstanceV3Exists(),
				),
			},
			{
				ResourceName:      "huaweicloud_rds_instance_v3.instance",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"db",
				},
			},
			{
				Config: testAccRdsInstanceV3_update(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdsInstanceV3Exists(),
				),
			},
		},
	})
}

func testAccRdsInstanceV3_basic(val string) string {
	return fmt.Sprintf(`
resource "huaweicloud_rds_instance_v3" "instance" {
  availability_zone = ["%s"]
  db {
    password = "Huangwei!120521"
    type = "PostgreSQL"
    version = "10"
    port = "8635"
  }
  name = "terraform_test_rds_instance%s"
  security_group_id = "3b5ceb06-3b8d-43ee-866a-dc0443b85de8"
  subnet_id = "%s"
  vpc_id = "%s"
  volume {
    type = "ULTRAHIGH"
    size = 50
  }
  flavor = "rds.pg.c2.medium"
  backup_strategy {
    start_time = "08:00-09:00"
    keep_days = 1
  }
}
	`, OS_AVAILABILITY_ZONE, val, OS_NETWORK_ID, OS_VPC_ID)
}

func testAccRdsInstanceV3_update(val string) string {
	return fmt.Sprintf(`
resource "huaweicloud_rds_instance_v3" "instance" {
  availability_zone = ["%s"]
  db {
    password = "Huangwei!120521"
    type = "PostgreSQL"
    version = "10"
    port = "8635"
  }
  name = "terraform_test_rds_instance%s"
  security_group_id = "3b5ceb06-3b8d-43ee-866a-dc0443b85de8"
  subnet_id = "%s"
  vpc_id = "%s"
  volume {
    type = "ULTRAHIGH"
    size = 100
  }
  flavor = "rds.pg.c2.medium"
  backup_strategy {
    start_time = "09:00-10:00"
    keep_days = 2
  }
}
	`, OS_AVAILABILITY_ZONE, val, OS_NETWORK_ID, OS_VPC_ID)
}

func testAccCheckRdsInstanceV3Destroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.sdkClient(OS_REGION_NAME, "rdsv3", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "huaweicloud_rds_instance_v3" {
			continue
		}

		_, err = fetchRdsInstanceV3ByListOnTest(rs, client)
		if err != nil {
			if strings.Index(err.Error(), "Error finding the resource by list api") != -1 {
				return nil
			}
			return err
		}
		return fmt.Errorf("huaweicloud_rds_instance_v3 still exists")
	}

	return nil
}

func testAccCheckRdsInstanceV3Exists() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)
		client, err := config.sdkClient(OS_REGION_NAME, "rdsv3", serviceProjectLevel)
		if err != nil {
			return fmt.Errorf("Error creating sdk client, err=%s", err)
		}

		rs, ok := s.RootModule().Resources["huaweicloud_rds_instance_v3.instance"]
		if !ok {
			return fmt.Errorf("Error checking huaweicloud_rds_instance_v3.instance exist, err=not found this resource")
		}

		_, err = fetchRdsInstanceV3ByListOnTest(rs, client)
		if err != nil {
			if strings.Index(err.Error(), "Error finding the resource by list api") != -1 {
				return fmt.Errorf("huaweicloud_rds_instance_v3 is not exist")
			}
			return fmt.Errorf("Error checking huaweicloud_rds_instance_v3.instance exist, err=%s", err)
		}
		return nil
	}
}

func fetchRdsInstanceV3ByListOnTest(rs *terraform.ResourceState,
	client *golangsdk.ServiceClient) (interface{}, error) {

	identity := map[string]interface{}{"id": rs.Primary.ID}

	queryLink := "?id=" + identity["id"].(string)

	link := client.ServiceURL("instances") + queryLink

	return findRdsInstanceV3ByList(client, link, identity)
}
