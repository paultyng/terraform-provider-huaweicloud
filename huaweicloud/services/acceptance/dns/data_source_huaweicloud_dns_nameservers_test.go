package dns

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceNameservers_basic(t *testing.T) {
	byType := "data.huaweicloud_dns_nameservers.filter_by_type"
	dcByType := acceptance.InitDataSourceCheck(byType)

	byServerRegion := "data.huaweicloud_dns_nameservers.filter_by_server_region"
	dcByServerRegion := acceptance.InitDataSourceCheck(byServerRegion)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceNameservers_basic(),
				Check: resource.ComposeTestCheckFunc(
					dcByType.CheckResourceExists(),
					resource.TestCheckResourceAttr(byType, "nameservers.0.type", "public"),
					resource.TestCheckResourceAttrSet(byType, "nameservers.0.ns_records.0.priority"),
					resource.TestCheckResourceAttrSet(byType, "nameservers.0.ns_records.0.hostname"),
					resource.TestCheckResourceAttr(byType, "nameservers.0.ns_records.0.address", ""),
					resource.TestCheckOutput("public_filter_is_useful", "true"),

					dcByServerRegion.CheckResourceExists(),
					resource.TestCheckResourceAttr(byServerRegion, "nameservers.0.type", "private"),
					resource.TestCheckResourceAttrSet(byServerRegion, "nameservers.0.ns_records.0.priority"),
					resource.TestCheckResourceAttrSet(byServerRegion, "nameservers.0.ns_records.0.address"),
					resource.TestCheckResourceAttr(byServerRegion, "nameservers.0.ns_records.0.hostname", ""),
					resource.TestCheckOutput("server_region_filter_is_useful", "true"),
					resource.TestCheckOutput("not_found_region", "true"),
				),
			},
		},
	})
}

func testDataSourceNameservers_basic() string {
	return fmt.Sprintf(`
data "huaweicloud_dns_nameservers" "filter_by_type" {
  type = "public"
}

locals {
  type_filter_result = [for v in data.huaweicloud_dns_nameservers.filter_by_type.nameservers[*].type : v == "public"]
}

output "public_filter_is_useful" {
  value = alltrue(local.type_filter_result) && length(local.type_filter_result) > 0
}

data "huaweicloud_dns_nameservers" "filter_by_server_region" {
  server_region = "%[1]s"
}

locals {
  server_region_filter_result = [for v in data.huaweicloud_dns_nameservers.filter_by_server_region.nameservers[*].region : v == "%[1]s"]
}

output "server_region_filter_is_useful" {
  value = alltrue(local.server_region_filter_result) && length(local.server_region_filter_result) == 1
}

data "huaweicloud_dns_nameservers" "not_found_region" {
  server_region = "not_fount"
}
  
output "not_found_region" {
  value = length(data.huaweicloud_dns_nameservers.not_found_region.nameservers) == 0
}
`, acceptance.HW_REGION_NAME)
}
