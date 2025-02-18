package dws

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func getDisasterRecoveryTaskResourceFunc(cfg *config.Config, state *terraform.ResourceState) (interface{}, error) {
	var (
		region  = acceptance.HW_REGION_NAME
		httpUrl = "v2/{project_id}/disaster-recovery/{disaster_recovery_id}"
		product = "dws"
	)
	client, err := cfg.NewServiceClient(product, region)
	if err != nil {
		return nil, fmt.Errorf("error creating DWS client: %s", err)
	}
	getPath := client.Endpoint + httpUrl
	getPath = strings.ReplaceAll(getPath, "{project_id}", client.ProjectID)
	getPath = strings.ReplaceAll(getPath, "{disaster_recovery_id}", state.Primary.ID)
	getOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json;charset=UTF-8",
		},
	}
	resp, err := client.Request("GET", getPath, &getOpt)
	if err != nil {
		return nil, fmt.Errorf("error retrieving DWS disaster recovery: %s", err)
	}
	respBody, err := utils.FlattenResponse(resp)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
func TestAccResourceDisasterRecoveryTask_basic(t *testing.T) {
	var (
		obj          interface{}
		resourceName = "huaweicloud_dws_disaster_recovery_task.test"
		name         = acceptance.RandomAccResourceName()
		password     = acceptance.RandomPassword()
	)
	rc := acceptance.InitResourceCheck(
		resourceName,
		&obj,
		getDisasterRecoveryTaskResourceFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAcDisasterRecoveryTask_basic(name, password),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "dr_sync_period", "2H"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "primary_cluster.0.name"),
					resource.TestCheckResourceAttrSet(resourceName, "standby_cluster.0.name"),
				),
			},
			{
				Config: testAcDisasterRecoveryTask_update(name, password),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "dr_sync_period", "3H"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
				),
			},
			{
				Config: testAcDisasterRecoveryTask_switch(name, password),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttrPair(resourceName, "primary_cluster.0.id", resourceName, "standby_cluster_id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"action", "primary_cluster_id", "standby_cluster_id"},
			},
		},
	})
}

func testAcDisasterRecoveryTask_base(name, password string) string {
	return fmt.Sprintf(`
%[1]s

data "huaweicloud_availability_zones" "test" {}

resource "huaweicloud_dws_cluster" "test" {
  name              = "%[2]s-1"
  node_type         = "dwsk2.xlarge"
  number_of_node    = 3
  vpc_id            = huaweicloud_vpc.test.id
  network_id        = huaweicloud_vpc_subnet.test.id
  security_group_id = huaweicloud_networking_secgroup.test.id
  availability_zone = data.huaweicloud_availability_zones.test.names[0]
  user_name         = "admin_user"
  user_pwd          = "%[3]s"
}

resource "huaweicloud_dws_cluster" "test1" {
  name              = "%[2]s-2"
  node_type         = "dwsk2.xlarge"
  number_of_node    = 3
  vpc_id            = huaweicloud_vpc.test.id
  network_id        = huaweicloud_vpc_subnet.test.id
  security_group_id = huaweicloud_networking_secgroup.test.id
  availability_zone = data.huaweicloud_availability_zones.test.names[0]
  user_name         = "admin_user"
  user_pwd          = "%[3]s"
}
`, common.TestBaseNetwork(name), name, password)
}

func testAcDisasterRecoveryTask_basic(name, password string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_dws_disaster_recovery_task" "test" {
  name               = "%s"
  dr_type            = "az"
  primary_cluster_id = huaweicloud_dws_cluster.test.id
  standby_cluster_id = huaweicloud_dws_cluster.test1.id
  dr_sync_period     = "2H"
}
`, testAcDisasterRecoveryTask_base(name, password), name)
}

func testAcDisasterRecoveryTask_update(name, password string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_dws_disaster_recovery_task" "test" {
  name               = "%s"
  dr_type            = "az"
  primary_cluster_id = huaweicloud_dws_cluster.test.id
  standby_cluster_id = huaweicloud_dws_cluster.test1.id
  dr_sync_period     = "3H"
  action             = "start"
}
`, testAcDisasterRecoveryTask_base(name, password), name)
}

func testAcDisasterRecoveryTask_switch(name, password string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_dws_disaster_recovery_task" "test" {
  name               = "%s"
  dr_type            = "az"
  primary_cluster_id = huaweicloud_dws_cluster.test.id
  standby_cluster_id = huaweicloud_dws_cluster.test1.id
  dr_sync_period     = "3H"
  action             = "switchover"
}
`, testAcDisasterRecoveryTask_base(name, password), name)
}
