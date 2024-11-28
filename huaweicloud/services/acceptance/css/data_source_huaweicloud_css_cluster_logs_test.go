package css

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceCssClusterLogs_basic(t *testing.T) {
	dataSource := "data.huaweicloud_css_cluster_logs.test"
	rName := acceptance.RandomAccResourceName()
	dc := acceptance.InitDataSourceCheck(dataSource)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceCssClusterLogs_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(dataSource, "logs.0.content"),
					resource.TestCheckResourceAttrSet(dataSource, "logs.0.date"),
					resource.TestCheckResourceAttrSet(dataSource, "logs.0.level"),
				),
			},
		},
	})
}

func testDataSourceCssClusterLogs_basic(name string) string {
	return fmt.Sprintf(`
%s

data "huaweicloud_css_cluster_logs" "test" {
  cluster_id    = huaweicloud_css_cluster.test.id
  instance_name = huaweicloud_css_cluster.test.nodes[0].name
  log_type      = "instance"
  level         = "INFO"
}
`, testAccCssCluster_basic(name, "Test@passw0rd", 7, "bar"))
}
