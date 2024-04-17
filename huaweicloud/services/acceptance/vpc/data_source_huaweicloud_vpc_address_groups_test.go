package vpc

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceVpcAddressGroups_basic(t *testing.T) {
	dataSource := "data.huaweicloud_vpc_address_groups.test1"
	rName := acceptance.RandomAccResourceName()
	dc := acceptance.InitDataSourceCheck(dataSource)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceVpcAddressGroups_noArgs(rName),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(dataSource, "address_groups.0.name"),
					resource.TestCheckResourceAttrSet(dataSource, "address_groups.0.max_capacity"),
					resource.TestCheckResourceAttrSet(dataSource, "address_groups.0.ip_extra_set.#"),
					resource.TestCheckResourceAttrSet(dataSource, "address_groups.0.addresses.#"),
					resource.TestCheckResourceAttrSet(dataSource, "address_groups.0.id"),
					resource.TestCheckResourceAttrSet(dataSource, "address_groups.0.ip_version"),
				),
			},
			{
				Config: testDataSourceVpcAddressGroups_args(rName),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttr(dataSource, "address_groups.#", "1"),
					resource.TestCheckResourceAttr(dataSource, "address_groups.0.name", rName),
					resource.TestCheckResourceAttr(dataSource, "address_groups.0.max_capacity", "10"),
					resource.TestCheckResourceAttr(dataSource, "address_groups.0.ip_extra_set.#", "3"),
					resource.TestCheckResourceAttr(dataSource, "address_groups.0.addresses.#", "3"),
					resource.TestCheckResourceAttrSet(dataSource, "address_groups.0.id"),
					resource.TestCheckResourceAttrSet(dataSource, "address_groups.0.ip_version"),
				),
			},
		},
	})
}

func testDataSourceVpcAddressGroups_noArgs(name string) string {
	return fmt.Sprintf(`
resource "huaweicloud_vpc_address_group" "test" {
  name        = "%s"
  description = "created by acc test"
  addresses   = [
    "192.168.5.0/24",
    "192.168.3.2",
    "192.168.3.20-192.168.3.100"
  ]
  max_capacity = 10
}

data "huaweicloud_vpc_address_groups" "test1" {

  depends_on = [huaweicloud_vpc_address_group.test]
}
`, name)
}

func testDataSourceVpcAddressGroups_args(name string) string {
	return fmt.Sprintf(`
resource "huaweicloud_vpc_address_group" "test" {
  name        = "%s"
  description = "created by acc test"
  addresses   = [
    "192.168.5.0/24",
    "192.168.3.2",
    "192.168.3.20-192.168.3.100"
  ]
  max_capacity = 10
}

data "huaweicloud_vpc_address_groups" "test1" {
  description = "created by acc test"
  name        = "%s"

  depends_on = [huaweicloud_vpc_address_group.test]
}
`, name, name)
}
