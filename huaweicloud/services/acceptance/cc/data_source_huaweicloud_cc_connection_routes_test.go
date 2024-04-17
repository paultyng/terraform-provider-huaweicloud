package cc

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceCcConnectionRoutes_basic(t *testing.T) {
	dataSource := "data.huaweicloud_cc_connection_routes.test"
	rName := acceptance.RandomAccResourceName()
	dc := acceptance.InitDataSourceCheck(dataSource)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckCCConnectionRouteProjectID(t)
			acceptance.TestAccPreCheckCCConnectionRouteRegionName(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceDataSourceCcConnectionRoutes_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(dataSource, "cloud_connection_routes.0.id"),
					resource.TestCheckResourceAttrSet(dataSource, "cloud_connection_routes.0.cloud_connection_id"),
					resource.TestCheckResourceAttrSet(dataSource, "cloud_connection_routes.0.instance_id"),
					resource.TestCheckResourceAttrSet(dataSource, "cloud_connection_routes.0.region_id"),

					resource.TestCheckOutput("id_filter_is_useful", "true"),
					resource.TestCheckOutput("conn_id_filter_is_useful", "true"),
					resource.TestCheckOutput("instance_id_filter_is_useful", "true"),
					resource.TestCheckOutput("region_id_filter_is_useful", "true"),
				),
			},
		},
	})
}

func testDataSourceDataSourceCcConnectionRoutes_basic(name string) string {
	return fmt.Sprintf(`
%s

data "huaweicloud_cc_connection_routes" "test" {
  depends_on = [
    huaweicloud_cc_network_instance.test1,
    huaweicloud_cc_network_instance.test2,
  ]
}

locals {
  cloud_connection_routes = data.huaweicloud_cc_connection_routes.test.cloud_connection_routes
  id                      = local.cloud_connection_routes[0].id
  cloud_connection_id     = local.cloud_connection_routes[0].cloud_connection_id
  instance_id             = local.cloud_connection_routes[0].instance_id
  region_id               = local.cloud_connection_routes[0].region_id
}

data "huaweicloud_cc_connection_routes" "filter_by_id" {
  cloud_connection_route_id = local.id
}

data "huaweicloud_cc_connection_routes" "filter_by_conn_id" {
  cloud_connection_id = local.cloud_connection_id
}

data "huaweicloud_cc_connection_routes" "filter_by_instance_id" {
  instance_id = local.instance_id
}

data "huaweicloud_cc_connection_routes" "filter_by_region_id" {
  region_id = local.region_id
}

locals {
  listById         = data.huaweicloud_cc_connection_routes.filter_by_id.cloud_connection_routes
  listByConnId     = data.huaweicloud_cc_connection_routes.filter_by_conn_id.cloud_connection_routes
  listByInstanceId = data.huaweicloud_cc_connection_routes.filter_by_instance_id.cloud_connection_routes
  listByRegionId   = data.huaweicloud_cc_connection_routes.filter_by_region_id.cloud_connection_routes
}

output "id_filter_is_useful" {
  value = length(local.listById) > 0 && alltrue(
    [for v in local.listById[*].id : v == local.id]
  )
}

output "conn_id_filter_is_useful" {
  value = length(local.listByConnId) > 0 && alltrue(
    [for v in local.listByConnId[*].cloud_connection_id : v == local.cloud_connection_id]
  )
}

output "instance_id_filter_is_useful" {
  value = length(local.listByInstanceId) > 0 && alltrue(
    [for v in local.listByInstanceId[*].instance_id : v == local.instance_id]
  )
}

output "region_id_filter_is_useful" {
  value = length(local.listByRegionId) > 0 && alltrue(
    [for v in local.listByRegionId[*].region_id : v == local.region_id]
  )
}
`, testDataSourceCcConnectionRoutes_dataBasic(name))
}

func testDataSourceCcConnectionRoutes_dataBasic(name string) string {
	return fmt.Sprintf(`
resource "huaweicloud_vpc" "test1" {
  region = "%[2]s"
  name   = "%[1]s_1"
  cidr   = cidrsubnet("10.12.0.0/16", 4, 1)
}

resource "huaweicloud_vpc" "test2" {
  region = "%[3]s"
  name   = "%[1]s_2"
  cidr   = "192.168.0.0/24"
}
	
resource "huaweicloud_vpc_subnet" "test1" {
  region     = "%[2]s"
  name       = "%[1]s_1"
  vpc_id     = huaweicloud_vpc.test1.id
  cidr       = cidrsubnet(huaweicloud_vpc.test1.cidr, 4, 1)
  gateway_ip = cidrhost(cidrsubnet(huaweicloud_vpc.test1.cidr, 4, 1), 1)
}

resource "huaweicloud_vpc_subnet" "test2" {
  region     = "%[3]s"
  name       = "%[1]s_2"
  vpc_id     = huaweicloud_vpc.test2.id
  cidr       = cidrsubnet(huaweicloud_vpc.test2.cidr, 4, 1)
  gateway_ip = cidrhost(cidrsubnet(huaweicloud_vpc.test2.cidr, 4, 1), 1)
}

resource "huaweicloud_cc_connection" "test" {
  name                  = "%[1]s"
  enterprise_project_id = "0"
  description           = "accDemo"
}

resource "huaweicloud_cc_network_instance" "test1" {
  type                = "vpc"
  cloud_connection_id = huaweicloud_cc_connection.test.id
  instance_id         = huaweicloud_vpc.test1.id
  project_id          = "%[4]s"
  region_id           = huaweicloud_vpc.test1.region
  description         = "desc 1"

  cidrs = [
    huaweicloud_vpc_subnet.test1.cidr,
  ]
}

resource "huaweicloud_cc_network_instance" "test2" {
  type                = "vpc"
  cloud_connection_id = huaweicloud_cc_connection.test.id
  instance_id         = huaweicloud_vpc.test2.id
  project_id          = "%[5]s"
  region_id           = huaweicloud_vpc.test2.region
  description         = "desc 2"
	
  cidrs = [
    huaweicloud_vpc_subnet.test2.cidr,
  ]

  depends_on = [
    huaweicloud_cc_network_instance.test1,
  ]
}
`, name, acceptance.HW_REGION_NAME_1, acceptance.HW_REGION_NAME_2,
		acceptance.HW_PROJECT_ID_1, acceptance.HW_PROJECT_ID_2)
}
