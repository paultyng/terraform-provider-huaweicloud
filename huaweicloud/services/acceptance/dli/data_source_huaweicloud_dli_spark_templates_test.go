package dli

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceDliSparkTemplates_basic(t *testing.T) {
	dataSource := "data.huaweicloud_dli_spark_templates.test"
	rName := acceptance.RandomAccResourceName()
	dc := acceptance.InitDataSourceCheck(dataSource)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceDataSourceDliSparkTemplates_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckOutput("template_id_filter_is_useful", "true"),
					resource.TestCheckOutput("name_filter_is_useful", "true"),
					resource.TestCheckOutput("group_filter_is_useful", "true"),
				),
			},
		},
	})
}

func testDataSourceDataSourceDliSparkTemplates_basic(name string) string {
	return fmt.Sprintf(`
%s

data "huaweicloud_dli_spark_templates" "test" {
  depends_on = [
    huaweicloud_dli_spark_template.test
  ]
}

data "huaweicloud_dli_spark_templates" "template_id_filter" {
  template_id = huaweicloud_dli_spark_template.test.id
}
  
locals {
  template_id = huaweicloud_dli_spark_template.test.id
}
  
output "template_id_filter_is_useful" {
  value = length(data.huaweicloud_dli_spark_templates.template_id_filter.templates) > 0 && alltrue(
    [for v in data.huaweicloud_dli_spark_templates.template_id_filter.templates[*].id : v == local.template_id]
  )
}

data "huaweicloud_dli_spark_templates" "name_filter" {
  name = huaweicloud_dli_spark_template.test.name
}
  
locals {
  name = huaweicloud_dli_spark_template.test.name
}
  
output "name_filter_is_useful" {
  value = length(data.huaweicloud_dli_spark_templates.name_filter.templates) > 0 && alltrue(
    [for v in data.huaweicloud_dli_spark_templates.name_filter.templates[*].name : v == local.name]
  )
}

data "huaweicloud_dli_spark_templates" "group_filter" {
  group = huaweicloud_dli_spark_template.test.group
}
  
locals {
  group = huaweicloud_dli_spark_template.test.group
}
  
output "group_filter_is_useful" {
  value = length(data.huaweicloud_dli_spark_templates.group_filter.templates) > 0 && alltrue(
    [for v in data.huaweicloud_dli_spark_templates.group_filter.templates[*].group : v == local.group]
  )
}
`, testSparkTemplate_basic(name))
}
