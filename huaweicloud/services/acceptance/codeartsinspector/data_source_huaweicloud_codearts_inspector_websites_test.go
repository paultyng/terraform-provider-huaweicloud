package codeartsinspector

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceCodeartsInspectorWebsites_basic(t *testing.T) {
	dataSource := "data.huaweicloud_codearts_inspector_websites.test"
	name := acceptance.RandomAccResourceName()
	dc := acceptance.InitDataSourceCheck(dataSource)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceCodeartsInspectorWebsites_basic(name),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(dataSource, "websites.#"),
					resource.TestCheckResourceAttrSet(dataSource, "websites.0.id"),
					resource.TestCheckResourceAttrSet(dataSource, "websites.0.website_name"),
					resource.TestCheckResourceAttrSet(dataSource, "websites.0.website_address"),
					resource.TestCheckResourceAttrSet(dataSource, "websites.0.high"),
					resource.TestCheckResourceAttrSet(dataSource, "websites.0.middle"),
					resource.TestCheckResourceAttrSet(dataSource, "websites.0.low"),
					resource.TestCheckResourceAttrSet(dataSource, "websites.0.hint"),
					resource.TestCheckResourceAttrSet(dataSource, "websites.0.created_at"),
					resource.TestCheckResourceAttrSet(dataSource, "websites.0.auth_status"),

					resource.TestCheckOutput("is_id_filter_useful", "true"),
					resource.TestCheckOutput("is_auth_status_filter_useful", "true"),
				),
			},
		},
	})
}

func testDataSourceCodeartsInspectorWebsites_basic(name string) string {
	return fmt.Sprintf(`
%[1]s

data "huaweicloud_codearts_inspector_websites" "test" {
  depends_on = [huaweicloud_codearts_inspector_website.test]
}

// filter by id
data "huaweicloud_codearts_inspector_websites" "filter_by_id" {
  domain_id = huaweicloud_codearts_inspector_website.test.id
}

locals {
  filter_result_by_id = [for v in data.huaweicloud_codearts_inspector_websites.filter_by_id.websites[*].id : 
    v == huaweicloud_codearts_inspector_website.test.id]
}

output "is_id_filter_useful" {
  value = length(local.filter_result_by_id) == 1
}

// filter by auth status
data "huaweicloud_codearts_inspector_websites" "filter_by_auth_status" {
  auth_status = huaweicloud_codearts_inspector_website.test.auth_status
}

locals {
  filter_result_by_auth_status = [for v in data.huaweicloud_codearts_inspector_websites.filter_by_auth_status.websites[*].auth_status : 
    v == huaweicloud_codearts_inspector_website.test.auth_status]
}

output "is_auth_status_filter_useful" {
  value = length(local.filter_result_by_auth_status) > 0 && alltrue(local.filter_result_by_auth_status) 
}
`, testInspectorWebsite_basic(name))
}
