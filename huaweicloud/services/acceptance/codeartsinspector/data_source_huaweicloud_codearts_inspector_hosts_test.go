package codeartsinspector

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceCodeartsInspectorHosts_basic(t *testing.T) {
	dataSource := "data.huaweicloud_codearts_inspector_hosts.test"
	rName := acceptance.RandomAccResourceName()
	dc := acceptance.InitDataSourceCheck(dataSource)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckCodeArtsSshCredentialID(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceCodeartsInspectorHosts_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.#"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.id"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.name"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.ip"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.os_type"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.ssh_credential_id"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.auth_status"),
				),
			},
		},
	})
}

func testDataSourceCodeartsInspectorHosts_basic(name string) string {
	return fmt.Sprintf(`
%[1]s

data "huaweicloud_codearts_inspector_hosts" "test" {
  depends_on = [huaweicloud_codearts_inspector_host.test]
}
`, testInspectorHost_basic(name))
}
