package dws

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDisasterRecoveryTasksDataSource_basic(t *testing.T) {
	resourceName := "data.huaweicloud_dws_disaster_recovery_tasks.name_filter"
	dc := acceptance.InitDataSourceCheck(resourceName)
	name := acceptance.RandomAccResourceName()
	password := acceptance.RandomPassword()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDisasterRecoveryTasksDataSource_basic(name, password),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(resourceName, "tasks.#"),
					resource.TestCheckOutput("name_filter_is_useful", "true"),
					resource.TestCheckOutput("drType_filter_is_useful", "true"),
					resource.TestCheckOutput("status_filter_is_useful", "true"),
					resource.TestCheckOutput("primary_cluster_name_result_is_useful", "true"),
					resource.TestCheckOutput("standby_cluster_name_result_is_useful", "true"),
					resource.TestCheckResourceAttrPair("data.huaweicloud_dws_disaster_recovery_tasks.name_filter", "tasks.0.name",
						"huaweicloud_dws_disaster_recovery_task.test", "name"),
					resource.TestCheckResourceAttrPair("data.huaweicloud_dws_disaster_recovery_tasks.drType_filter", "tasks.0.dr_type",
						"huaweicloud_dws_disaster_recovery_task.test", "dr_type"),
					resource.TestCheckResourceAttrPair("data.huaweicloud_dws_disaster_recovery_tasks.status_filter", "tasks.0.status",
						"huaweicloud_dws_disaster_recovery_task.test", "status"),
					resource.TestCheckResourceAttrPair("data.huaweicloud_dws_disaster_recovery_tasks.p_name_filter", "tasks.0.primary_cluster_name",
						"huaweicloud_dws_disaster_recovery_task.test", "primary_cluster.0.name"),
					resource.TestCheckResourceAttrPair("data.huaweicloud_dws_disaster_recovery_tasks.s_name_filter", "tasks.0.standby_cluster_name",
						"huaweicloud_dws_disaster_recovery_task.test", "standby_cluster.0.name"),
				),
			},
		},
	})
}

func testAccDisasterRecoveryTasksDataSource_basic(name, password string) string {
	return fmt.Sprintf(`
%s

data "huaweicloud_dws_disaster_recovery_tasks" "name_filter" {
  name = huaweicloud_dws_disaster_recovery_task.test.name
}

data "huaweicloud_dws_disaster_recovery_tasks" "drType_filter" {
  dr_type = huaweicloud_dws_disaster_recovery_task.test.dr_type
}

data "huaweicloud_dws_disaster_recovery_tasks" "status_filter" {
  status = huaweicloud_dws_disaster_recovery_task.test.status
}

data "huaweicloud_dws_disaster_recovery_tasks" "p_name_filter" {
  primary_cluster_name = huaweicloud_dws_disaster_recovery_task.test.primary_cluster[0].name
}

data "huaweicloud_dws_disaster_recovery_tasks" "s_name_filter" {
  standby_cluster_name = huaweicloud_dws_disaster_recovery_task.test.standby_cluster[0].name
}

locals {
  name_filter = [for v in data.huaweicloud_dws_disaster_recovery_tasks.name_filter.tasks[*].name :
  v == huaweicloud_dws_disaster_recovery_task.test.name]
  drType_filter = [for v in data.huaweicloud_dws_disaster_recovery_tasks.drType_filter.tasks[*].dr_type :
  v == huaweicloud_dws_disaster_recovery_task.test.dr_type]
  status_filter = [for v in data.huaweicloud_dws_disaster_recovery_tasks.status_filter.tasks[*].status :
  v == huaweicloud_dws_disaster_recovery_task.test.status]
  primary_cluster_name_filter = [for v in data.huaweicloud_dws_disaster_recovery_tasks.p_name_filter.tasks[*].primary_cluster_name :
  v == huaweicloud_dws_disaster_recovery_task.test.primary_cluster[0].name]
  standby_cluster_name_filter = [for v in data.huaweicloud_dws_disaster_recovery_tasks.s_name_filter.tasks[*].standby_cluster_name :
  v == huaweicloud_dws_disaster_recovery_task.test.standby_cluster[0].name]
}

output "name_filter_is_useful" {
  value = alltrue(local.name_filter) && length(local.name_filter) > 0
}
output "drType_filter_is_useful" {
  value = alltrue(local.drType_filter) && length(local.drType_filter) > 0
}
output "status_filter_is_useful" {
  value = alltrue(local.status_filter) && length(local.status_filter) > 0
}
output "primary_cluster_name_result_is_useful" {
  value = alltrue(local.primary_cluster_name_filter) && length(local.primary_cluster_name_filter) > 0
}
output "standby_cluster_name_result_is_useful" {
  value = alltrue(local.standby_cluster_name_filter) && length(local.standby_cluster_name_filter) > 0
}
`, testAcDisasterRecoveryTask_basic(name, password))
}
