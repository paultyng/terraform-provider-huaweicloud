package gaussdb

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceGaussDBOpenGaussPlugins_basic(t *testing.T) {
	dataSource := "data.huaweicloud_gaussdb_opengauss_plugins.test"
	dc := acceptance.InitDataSourceCheck(dataSource)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceGaussDBOpenGaussPlugins_basic(),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(dataSource, "plugins.#"),
				),
			},
		},
	})
}

func testDataSourceGaussDBOpenGaussPlugins_basic() string {
	return `
data "huaweicloud_gaussdb_opengauss_plugins" "test" {}
`
}
