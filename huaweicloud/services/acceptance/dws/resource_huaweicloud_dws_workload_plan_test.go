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

func getWorkLoadPlanResourceFunc(cfg *config.Config, state *terraform.ResourceState) (interface{}, error) {
	var (
		region  = acceptance.HW_REGION_NAME
		httpUrl = "v2/{project_id}/clusters/{cluster_id}/workload/plans"
		product = "dws"
	)

	client, err := cfg.NewServiceClient(product, region)
	if err != nil {
		return nil, fmt.Errorf("error creating DWS client: %s", err)
	}

	listPath := client.Endpoint + httpUrl
	listPath = strings.ReplaceAll(listPath, "{project_id}", client.ProjectID)
	listPath = strings.ReplaceAll(listPath, "{cluster_id}", state.Primary.Attributes["cluster_id"])
	listOpt := golangsdk.RequestOpts{
		MoreHeaders:      map[string]string{"Content-Type": "application/json;charset=UTF-8"},
		KeepResponseBody: true,
	}

	listResp, err := client.Request("GET", listPath, &listOpt)
	if err != nil {
		return nil, fmt.Errorf("error querying DWS workload plans: %s", err)
	}

	listRespBody, err := utils.FlattenResponse(listResp)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	jsonPaths := fmt.Sprintf("plan_list[?plan_name=='%s']", state.Primary.Attributes["name"])
	plans := utils.PathSearch(jsonPaths, listRespBody, make([]interface{}, 0)).([]interface{})
	if len(plans) == 0 {
		return nil, golangsdk.ErrDefault404{}
	}

	return plans[0], nil
}

func TestAccResourceWorkLoadPlan_basic(t *testing.T) {
	var (
		obj          interface{}
		resourceName = "huaweicloud_dws_workload_plan.test"
		name         = acceptance.RandomAccResourceName()
	)

	rc := acceptance.InitResourceCheck(
		resourceName,
		&obj,
		getWorkLoadPlanResourceFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkLoadPlan_basic(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "status", "disabled"),
					resource.TestCheckResourceAttrPair(resourceName, "cluster_id", "huaweicloud_dws_cluster.test", "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testWorkLoadPlanImportState(resourceName),
			},
		},
	})
}

func testAccWorkLoadPlan_base(name, password string) string {
	return fmt.Sprintf(`
%[1]s

data "huaweicloud_availability_zones" "test" {}

resource "huaweicloud_dws_cluster" "test" {
  name              = "%[2]s"
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

func testAccWorkLoadPlan_basic(name string) string {
	password := acceptance.RandomPassword("@#%&_=!")

	return fmt.Sprintf(`
%s

resource "huaweicloud_dws_workload_plan" "test" {
  cluster_id = huaweicloud_dws_cluster.test.id
  name       = "%s"
}
`, testAccWorkLoadPlan_base(name, password), name)
}

func testWorkLoadPlanImportState(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("resource (%s) not found: %s", name, rs)
		}
		clusterId := rs.Primary.Attributes["cluster_id"]
		planName := rs.Primary.Attributes["name"]
		if clusterId == "" || planName == "" {
			return "", fmt.Errorf("the workload plan name or related cluster ID is missing")
		}

		return fmt.Sprintf("%s/%s", clusterId, planName), nil
	}
}
