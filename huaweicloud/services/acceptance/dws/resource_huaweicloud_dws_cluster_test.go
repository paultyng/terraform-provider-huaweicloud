package dws

import (
	"fmt"
	"testing"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"

	"github.com/chnsz/golangsdk/openstack/dws/v1/cluster"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
)

func getDwsResourceFunc(config *config.Config, state *terraform.ResourceState) (interface{}, error) {
	client, err := config.DwsV1Client(acceptance.HW_REGION_NAME)
	if err != nil {
		return nil, fmt.Errorf("error creating DWS v1 client, err=%s", err)
	}
	return cluster.Get(client, state.Primary.ID)
}

func TestAccResourceDWS_basic(t *testing.T) {
	var clusterInstance cluster.CreateOpts
	resourceName := "huaweicloud_dws_cluster.test"
	name := acceptance.RandomAccResourceName()

	rc := acceptance.InitResourceCheck(
		resourceName,
		&clusterInstance,
		getDwsResourceFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccDwsCluster_basic(name, 3, cluster.PublicBindTypeAuto, "cluster123@!", "bar"),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "number_of_node", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags.key", "val"),
					resource.TestCheckResourceAttr(resourceName, "tags.foo", "bar"),
				),
			},
			{
				Config: testAccDwsCluster_basic(name, 6, cluster.PublicBindTypeAuto, "cluster123@!u", "cat"),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "number_of_node", "6"),
					resource.TestCheckResourceAttr(resourceName, "tags.key", "val"),
					resource.TestCheckResourceAttr(resourceName, "tags.foo", "cat"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"user_pwd", "number_of_cn"},
			},
		},
	})
}

func testAccBaseResource(rName string) string {
	return fmt.Sprintf(`
resource "huaweicloud_vpc" "test" {
  name = "%s"
  cidr = "192.168.0.0/16"
}

resource "huaweicloud_vpc_subnet" "test" {
  name       = "%s"
  cidr       = "192.168.0.0/24"
  vpc_id     = huaweicloud_vpc.test.id
  gateway_ip = "192.168.0.1"
}

resource "huaweicloud_networking_secgroup" "test" {
  name        = "%s"
  description = "terraform security group acceptance test"
}

data "huaweicloud_availability_zones" "test" {}
`, rName, rName, rName)
}

func testAccDwsCluster_basic(rName string, numberOfNode int, publicIpBindType, password, tag string) string {
	baseResource := testAccBaseResource(rName)

	return fmt.Sprintf(`
%s

data "huaweicloud_dws_flavors" "test" {
  vcpus             = 8
  memory            = 64
  availability_zone = data.huaweicloud_availability_zones.test.names[0]
}

resource "huaweicloud_dws_cluster" "test" {
  name              = "%s"
  node_type         = data.huaweicloud_dws_flavors.test.flavors.0.flavor_id
  number_of_node    = %d
  vpc_id            = huaweicloud_vpc.test.id
  network_id        = huaweicloud_vpc_subnet.test.id
  security_group_id = huaweicloud_networking_secgroup.test.id
  availability_zone = data.huaweicloud_availability_zones.test.names[0]
  user_name         = "test_cluster_admin"
  user_pwd          = "%s"

  public_ip {
    public_bind_type = "%s"
  }

  tags = {
    key = "val"
    foo = "%s"
  }
}
`, baseResource, rName, numberOfNode, password, publicIpBindType, tag)
}
