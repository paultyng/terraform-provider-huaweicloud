package elb

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDatasourceListeners_basic(t *testing.T) {
	rName := "data.huaweicloud_elb_listeners.test"
	dc := acceptance.InitDataSourceCheck(rName)
	name := acceptance.RandomAccResourceName()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDatasourceListeners_basic(name),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(rName, "listeners.#"),
					resource.TestCheckResourceAttrSet(rName, "listeners.0.name"),
					resource.TestCheckResourceAttrSet(rName, "listeners.0.id"),
					resource.TestCheckResourceAttrSet(rName, "listeners.0.loadbalancer_id"),
					resource.TestCheckResourceAttrSet(rName, "listeners.0.description"),
					resource.TestCheckResourceAttrSet(rName, "listeners.0.protocol"),
					resource.TestCheckResourceAttrSet(rName, "listeners.0.protocol_port"),
					resource.TestCheckOutput("name_filter_is_useful", "true"),
					resource.TestCheckOutput("protocol_filter_is_useful", "true"),
					resource.TestCheckOutput("protocol_port_filter_is_useful", "true"),
					resource.TestCheckOutput("description_filter_is_useful", "true"),
					resource.TestCheckOutput("loadbalancer_id_filter_is_useful", "true"),
				),
			},
		},
	})
}

func testAccElbListenerConfig_basic(name string) string {
	return fmt.Sprintf(`
resource "huaweicloud_vpc" "test" {
  name = "%[1]s"
  cidr = "192.168.0.0/16"
}

resource "huaweicloud_vpc_subnet" "test" {
  name       = "%[1]s"
  vpc_id     = huaweicloud_vpc.test.id
  cidr       = "192.168.0.0/24"
  gateway_ip = "192.168.0.1"
}

data "huaweicloud_availability_zones" "test" {}

resource "huaweicloud_elb_loadbalancer" "test" {
  name           = "%[2]s"
  vpc_id         = huaweicloud_vpc.test.id
  ipv4_subnet_id = huaweicloud_vpc_subnet.test.ipv4_subnet_id
  
  availability_zone = [
    data.huaweicloud_availability_zones.test.names[0]
  ]
  backend_subnets = [
    huaweicloud_vpc_subnet.test.id
  ]
}

resource "huaweicloud_elb_listener" "test" {
 name                        = "%[1]s"
 description                 = "test description"
 protocol                    = "HTTP"
 protocol_port               = 8080
 loadbalancer_id             = huaweicloud_elb_loadbalancer.test.id
 advanced_forwarding_enabled = false

}
`, name, name)
}

func testAccDatasourceListeners_basic(name string) string {
	return fmt.Sprintf(`
%[1]s

data "huaweicloud_elb_listeners" "test" {
  depends_on = [huaweicloud_elb_listener.test]
}

data "huaweicloud_elb_listeners" "name_filter" {
  name       = "%[2]s"
  depends_on = [huaweicloud_elb_listener.test]
}
output "name_filter_is_useful" {
  value = length(data.huaweicloud_elb_listeners.name_filter.listeners) > 0 && alltrue(
  [for v in data.huaweicloud_elb_listeners.name_filter.listeners[*].name :v == "%[2]s"]
  )  
}

data "huaweicloud_elb_listeners" "description_filter" {
  description = huaweicloud_elb_listener.test.description
  depends_on  = [huaweicloud_elb_listener.test]
}
locals {
  description = huaweicloud_elb_listener.test.description
}
output "description_filter_is_useful" {
  value = length(data.huaweicloud_elb_listeners.description_filter.listeners) > 0 && alltrue(
  [for v in data.huaweicloud_elb_listeners.description_filter.listeners[*].description :v == local.description]
  )  
}

data "huaweicloud_elb_listeners" "protocol_filter" {
  protocol = huaweicloud_elb_listener.test.protocol
  depends_on  = [huaweicloud_elb_listener.test]
}
locals {
  protocol = huaweicloud_elb_listener.test.protocol
}
output "protocol_filter_is_useful" {
  value = length(data.huaweicloud_elb_listeners.protocol_filter.listeners) > 0 && alltrue(
  [for v in data.huaweicloud_elb_listeners.protocol_filter.listeners[*].protocol :v == local.protocol]
  )  
}

data "huaweicloud_elb_listeners" "protocol_port_filter" {
  protocol_port = huaweicloud_elb_listener.test.protocol_port
  depends_on  = [huaweicloud_elb_listener.test]
}
locals {
  protocol_port = huaweicloud_elb_listener.test.protocol_port
}
output "protocol_port_filter_is_useful" {
  value = length(data.huaweicloud_elb_listeners.protocol_port_filter.listeners) > 0 && alltrue(
  [for v in data.huaweicloud_elb_listeners.protocol_port_filter.listeners[*].protocol_port :v == local.protocol_port]
  )  
}

data "huaweicloud_elb_listeners" "loadbalancer_id_filter" {
   depends_on  = [huaweicloud_elb_listener.test]
   loadbalancer_id = huaweicloud_elb_listener.test.loadbalancer_id
}
locals {
  loadbalancer_id = huaweicloud_elb_listener.test.loadbalancer_id
}
output "loadbalancer_id_filter_is_useful" {
  value = length(data.huaweicloud_elb_listeners.loadbalancer_id_filter.listeners) > 0 && alltrue(
  [for v in data.huaweicloud_elb_listeners.loadbalancer_id_filter.listeners[*].loadbalancer_id :v == local.loadbalancer_id]
  )  
}

`, testAccElbListenerConfig_basic(name), name)
}
