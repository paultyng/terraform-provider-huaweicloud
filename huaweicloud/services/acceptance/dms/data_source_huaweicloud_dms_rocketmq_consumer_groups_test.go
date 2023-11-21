package dms

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDmsRocketMQConsumerGroups_basic(t *testing.T) {
	name := acceptance.RandomAccResourceName()
	dataSourceName := "data.huaweicloud_dms_rocketmq_consumer_groups.test"
	dc := acceptance.InitDataSourceCheck(dataSourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDmsRocketMQSearchConsumerGroups(name),
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(dataSourceName, "groups.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "groups.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "groups.0.enabled"),
					resource.TestCheckResourceAttrSet(dataSourceName, "groups.0.broadcast"),
					resource.TestCheckResourceAttrSet(dataSourceName, "groups.0.retry_max_time"),
					resource.TestCheckOutput("name_filter_is_useful", "true"),
					resource.TestCheckOutput("enabled_filter_is_useful", "true"),
					resource.TestCheckOutput("broadcast_filter_is_useful", "true"),
					resource.TestCheckOutput("retry_max_time_filter_is_useful", "true"),
				),
			},
		},
	})
}

func testAccDmsRocketMQSearchConsumerGroups(name string) string {
	return fmt.Sprintf(`
%[1]s

data "huaweicloud_dms_rocketmq_consumer_groups" "test" {
  depends_on  = [huaweicloud_dms_rocketmq_consumer_group.test]
  instance_id = huaweicloud_dms_rocketmq_instance.test.id
}

data "huaweicloud_dms_rocketmq_consumer_groups" "name_filter" {
  depends_on  = [huaweicloud_dms_rocketmq_consumer_group.test]
  instance_id = huaweicloud_dms_rocketmq_instance.test.id
  name        = "%[2]s"
}
  
output "name_filter_is_useful" {
  value = length(data.huaweicloud_dms_rocketmq_consumer_groups.name_filter.groups) > 0 && alltrue(
  [for v in data.huaweicloud_dms_rocketmq_consumer_groups.name_filter.groups[*].name : v == "%[2]s"]
  )  
}

data "huaweicloud_dms_rocketmq_consumer_groups" "enabled_filter" {
  depends_on  = [huaweicloud_dms_rocketmq_consumer_group.test]
  instance_id = huaweicloud_dms_rocketmq_instance.test.id
  enabled     = huaweicloud_dms_rocketmq_consumer_group.test.enabled
}

locals {
  enabled = huaweicloud_dms_rocketmq_consumer_group.test.enabled
}
    
output "enabled_filter_is_useful" {
  value = length(data.huaweicloud_dms_rocketmq_consumer_groups.enabled_filter.groups) > 0 && alltrue(
  [for v in data.huaweicloud_dms_rocketmq_consumer_groups.enabled_filter.groups[*].enabled : v == local.enabled]
  )  
}

data "huaweicloud_dms_rocketmq_consumer_groups" "broadcast_filter" {
  depends_on  = [huaweicloud_dms_rocketmq_consumer_group.test]
  instance_id = huaweicloud_dms_rocketmq_instance.test.id
  broadcast   = huaweicloud_dms_rocketmq_consumer_group.test.broadcast
}

locals {
  broadcast = huaweicloud_dms_rocketmq_consumer_group.test.broadcast
}
    
output "broadcast_filter_is_useful" {
  value = length(data.huaweicloud_dms_rocketmq_consumer_groups.broadcast_filter.groups) > 0 && alltrue(
  [for v in data.huaweicloud_dms_rocketmq_consumer_groups.broadcast_filter.groups[*].broadcast : v == local.broadcast]
  )  
}

data "huaweicloud_dms_rocketmq_consumer_groups" "retry_max_time_filter" {
  depends_on     = [huaweicloud_dms_rocketmq_consumer_group.test]
  instance_id    = huaweicloud_dms_rocketmq_instance.test.id
  retry_max_time = huaweicloud_dms_rocketmq_consumer_group.test.retry_max_time
}

locals {
  retry_max_time = huaweicloud_dms_rocketmq_consumer_group.test.retry_max_time
}
    
output "retry_max_time_filter_is_useful" {
  value = length(data.huaweicloud_dms_rocketmq_consumer_groups.retry_max_time_filter.groups) > 0 && alltrue(
  [for v in data.huaweicloud_dms_rocketmq_consumer_groups.retry_max_time_filter.groups[*].retry_max_time : v == local.retry_max_time]
  )
}

`, testDmsRocketMQConsumerGroup_basic(name), name)
}
