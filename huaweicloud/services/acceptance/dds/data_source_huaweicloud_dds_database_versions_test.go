package dds

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceDdsDatabaseVersions_basic(t *testing.T) {
	dataSource := "data.huaweicloud_dds_database_versions.test1"
	dc := acceptance.InitDataSourceCheck(dataSource)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceDataSourceDdsDatabaseVersions_basic(),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(dataSource, "versions.0"),
				),
			},
		},
	})
}

func testDataSourceDataSourceDdsDatabaseVersions_basic() string {
	return `
data "huaweicloud_dds_database_versions" "test1" {
  datastore_name = "DDS-Community"
}
`
}
