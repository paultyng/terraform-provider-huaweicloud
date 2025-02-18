package mrs

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDatasourceMrsFlavors_basic(t *testing.T) {
	rName := "data.huaweicloud_mapreduce_flavors.test"
	dc := acceptance.InitDataSourceCheck(rName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDatasourceMrsFlavors_basic(),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "flavors.0.version_name", "MRS 3.2.0-LTS.1"),
					resource.TestCheckResourceAttrSet(rName, "flavors.0.flavor_id"),
					resource.TestCheckResourceAttrSet(rName, "flavors.0.version_name"),
					resource.TestCheckResourceAttrSet(rName, "flavors.0.availability_zone"),
					resource.TestCheckResourceAttrSet(rName, "flavors.0.node_type"),

					resource.TestCheckOutput("availability_zone_filter_is_useful", "true"),

					resource.TestCheckOutput("node_type_filter_is_useful", "true"),
				),
			},
		},
	})
}

func testAccDatasourceMrsFlavors_basic() string {
	return `
data "huaweicloud_mapreduce_flavors" "test" {
  version_name = "MRS 3.2.0-LTS.1"
}

data "huaweicloud_availability_zones" "test" {}

data "huaweicloud_mapreduce_flavors" "availability_zone_filter" {
  version_name      = "MRS 3.2.0-LTS.1"
  availability_zone = data.huaweicloud_availability_zones.test.names[0]
}
output "availability_zone_filter_is_useful" {
  value = length(data.huaweicloud_mapreduce_flavors.availability_zone_filter.flavors) > 0 && alltrue(
    [for v in data.huaweicloud_mapreduce_flavors.availability_zone_filter.flavors[*].availability_zone :
    v == data.huaweicloud_availability_zones.test.names[0]]
  )
}

data "huaweicloud_mapreduce_flavors" "node_type_filter" {
  version_name = "MRS 3.2.0-LTS.1"
  node_type    = "master"
}

output "node_type_filter_is_useful" {
  value = length(data.huaweicloud_mapreduce_flavors.node_type_filter.flavors) > 0 && alltrue(
    [for v in data.huaweicloud_mapreduce_flavors.node_type_filter.flavors[*].node_type : v == "master"]
  )
}
`
}
