package hss

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceHosts_basic(t *testing.T) {
	dataSource := "data.huaweicloud_hss_hosts.test"
	dc := acceptance.InitDataSourceCheck(dataSource)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckHSSHostProtectionHostId(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceHosts_basic(),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.#"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.id"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.name"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.status"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.os_type"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.agent_id"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.agent_status"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.protect_status"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.protect_version"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.protect_charging_mode"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.quota_id"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.detect_result"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.policy_group_id"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.asset_value"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.open_time"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.private_ip"),
					resource.TestCheckResourceAttrSet(dataSource, "hosts.0.enterprise_project_id"),

					resource.TestCheckOutput("is_host_id_filter_useful", "true"),
					resource.TestCheckOutput("is_agent_status_filter_useful", "true"),
					resource.TestCheckOutput("is_protect_version_filter_useful", "true"),
					resource.TestCheckOutput("is_protect_charging_mode_filter_useful", "true"),
					resource.TestCheckOutput("not_found_validation_pass", "true"),
				),
			},
		},
	})
}

func testDataSourceHosts_basic() string {
	return `

data "huaweicloud_hss_hosts" "test" {}

# Filter using host ID.
locals {
  host_id = data.huaweicloud_hss_hosts.test.hosts[0].id
}

data "huaweicloud_hss_hosts" "host_id_filter" {
  host_id = local.host_id
}

output "is_host_id_filter_useful" {
  value = length(data.huaweicloud_hss_hosts.host_id_filter.hosts) > 0 && alltrue(
    [for v in data.huaweicloud_hss_hosts.host_id_filter.hosts[*].id : v == local.host_id]
  )
}

# Filter using agent_status
locals {
  agent_status = data.huaweicloud_hss_hosts.test.hosts[0].agent_status
}

data "huaweicloud_hss_hosts" "agent_status_filter" {
  agent_status = local.agent_status
}

output "is_agent_status_filter_useful" {
  value = length(data.huaweicloud_hss_hosts.agent_status_filter.hosts) > 0 && alltrue(
    [for v in data.huaweicloud_hss_hosts.agent_status_filter.hosts[*].agent_status : v == local.agent_status]
  )
}

# Filter using protect_version
locals {
  protect_version = data.huaweicloud_hss_hosts.test.hosts[0].protect_version
}

data "huaweicloud_hss_hosts" "protect_version_filter" {
  protect_version = local.protect_version
}

output "is_protect_version_filter_useful" {
  value = length(data.huaweicloud_hss_hosts.protect_version_filter.hosts) > 0 && alltrue(
    [for v in data.huaweicloud_hss_hosts.protect_version_filter.hosts[*].protect_version : v == local.protect_version]
  )
}

# Filter using protect_charging_mode
locals {
  protect_charging_mode = data.huaweicloud_hss_hosts.test.hosts[0].protect_charging_mode
}

data "huaweicloud_hss_hosts" "protect_charging_mode_filter" {
  protect_charging_mode = local.protect_charging_mode
}

output "is_protect_charging_mode_filter_useful" {
  value = length(data.huaweicloud_hss_hosts.protect_charging_mode_filter.hosts) > 0 && alltrue(
    [for v in data.huaweicloud_hss_hosts.protect_charging_mode_filter.hosts[*].protect_charging_mode : v == local.protect_charging_mode]
  )
}

# Filter using non existent name.
data "huaweicloud_hss_hosts" "not_found" {
  name = "resource_not_found"
}

output "not_found_validation_pass" {
  value = length(data.huaweicloud_hss_hosts.not_found.hosts) == 0
}
`
}
