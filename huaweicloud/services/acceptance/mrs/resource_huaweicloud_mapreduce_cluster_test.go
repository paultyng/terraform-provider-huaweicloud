package mrs

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/mrs/v1/cluster"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils/fmtp"
)

type GroupNodeNum struct {
	AnalysisCoreNum int
	StreamCoreNum   int
	AnalysisTaskNum int
	StreamTaskNum   int
}

func TestAccMrsMapReduceCluster_basic(t *testing.T) {
	var clusterGet cluster.Cluster
	resourceName := "huaweicloud_mapreduce_cluster.test"
	rName := fmt.Sprintf("tf-acc-test-%s", acctest.RandString(5))
	password := fmt.Sprintf("TF%s%s%d", acctest.RandString(10), acctest.RandStringFromCharSet(1, "-_"),
		acctest.RandIntRange(0, 99))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAccPreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckMRSV2ClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMrsMapReduceClusterConfig_basic(rName, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "STREAMING"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "tags.foo", "bar"),
					resource.TestCheckResourceAttr(resourceName, "tags.key", "value"),
				),
			},
			{
				Config: testAccMrsMapReduceClusterConfig_update(rName, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "STREAMING"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "tags.foo1", "bar"),
					resource.TestCheckResourceAttr(resourceName, "tags.key", "update_value"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"manager_admin_pass",
					"node_admin_pass",
				},
			},
		},
	})
}

func TestAccMrsMapReduceCluster_keypair(t *testing.T) {
	var clusterGet cluster.Cluster
	resourceName := "huaweicloud_mapreduce_cluster.test"
	rName := fmt.Sprintf("tf-acc-test-%s", acctest.RandString(5))
	password := fmt.Sprintf("TF%s%s%d", acctest.RandString(10), acctest.RandStringFromCharSet(1, "-_"),
		acctest.RandIntRange(0, 99))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAccPreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckMRSV2ClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMrsMapReduceClusterConfig_keypair(rName, password),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "STREAMING"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"manager_admin_pass",
				},
			},
		},
	})
}

func TestAccMrsMapReduceCluster_analysis(t *testing.T) {
	var clusterGet cluster.Cluster
	resourceName := "huaweicloud_mapreduce_cluster.test"
	rName := fmt.Sprintf("tf-acc-test-%s", acctest.RandString(5))
	password := fmt.Sprintf("TF%s%s%d", acctest.RandString(10), acctest.RandStringFromCharSet(1, "-_"),
		acctest.RandIntRange(0, 99))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAccPreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckMRSV2ClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMrsMapReduceClusterConfig_analysis(rName, password, buildGroupNodeNumbers(2, 0, 1, 0)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "ANALYSIS"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "analysis_core_nodes.0.node_number", "2"),
					resource.TestCheckResourceAttr(resourceName, "analysis_task_nodes.0.node_number", "1"),
				),
			},
			{
				Config: testAccMrsMapReduceClusterConfig_analysis(rName, password, buildGroupNodeNumbers(3, 0, 2, 0)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "ANALYSIS"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "analysis_core_nodes.0.node_number", "3"),
					resource.TestCheckResourceAttr(resourceName, "analysis_task_nodes.0.node_number", "2"),
				),
			},
			{
				Config: testAccMrsMapReduceClusterConfig_analysis(rName, password, buildGroupNodeNumbers(2, 0, 1, 0)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "ANALYSIS"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "analysis_core_nodes.0.node_number", "2"),
					resource.TestCheckResourceAttr(resourceName, "analysis_task_nodes.0.node_number", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"manager_admin_pass",
					"node_admin_pass",
				},
			},
		},
	})
}

func TestAccMrsMapReduceCluster_stream(t *testing.T) {
	var clusterGet cluster.Cluster
	resourceName := "huaweicloud_mapreduce_cluster.test"
	rName := fmt.Sprintf("tf-acc-test-%s", acctest.RandString(5))
	password := fmt.Sprintf("TF%s%s%d", acctest.RandString(10), acctest.RandStringFromCharSet(1, "-_"),
		acctest.RandIntRange(0, 99))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAccPreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckMRSV2ClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMrsMapReduceClusterConfig_stream(rName, password, buildGroupNodeNumbers(0, 2, 0, 1)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "STREAMING"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "streaming_core_nodes.0.node_number", "2"),
					resource.TestCheckResourceAttr(resourceName, "streaming_task_nodes.0.node_number", "1"),
				),
			},
			{
				Config: testAccMrsMapReduceClusterConfig_stream(rName, password, buildGroupNodeNumbers(0, 3, 0, 2)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "STREAMING"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "streaming_core_nodes.0.node_number", "3"),
					resource.TestCheckResourceAttr(resourceName, "streaming_task_nodes.0.node_number", "2"),
				),
			},
			{
				Config: testAccMrsMapReduceClusterConfig_stream(rName, password, buildGroupNodeNumbers(0, 2, 0, 1)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "STREAMING"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "streaming_core_nodes.0.node_number", "2"),
					resource.TestCheckResourceAttr(resourceName, "streaming_task_nodes.0.node_number", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"manager_admin_pass",
					"node_admin_pass",
				},
			},
		},
	})
}

func TestAccMrsMapReduceCluster_hybrid(t *testing.T) {
	var clusterGet cluster.Cluster
	resourceName := "huaweicloud_mapreduce_cluster.test"
	rName := fmt.Sprintf("tf-acc-test-%s", acctest.RandString(5))
	password := fmt.Sprintf("TF%s%s%d", acctest.RandString(10), acctest.RandStringFromCharSet(1, "-_"),
		acctest.RandIntRange(0, 99))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAccPreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckMRSV2ClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMrsMapReduceClusterConfig_hybrid(rName, password, buildGroupNodeNumbers(2, 2, 1, 1)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "MIXED"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "analysis_core_nodes.0.node_number", "2"),
					resource.TestCheckResourceAttr(resourceName, "streaming_core_nodes.0.node_number", "2"),
					resource.TestCheckResourceAttr(resourceName, "analysis_task_nodes.0.node_number", "1"),
					resource.TestCheckResourceAttr(resourceName, "streaming_task_nodes.0.node_number", "1"),
					resource.TestCheckResourceAttr(resourceName, "analysis_core_nodes.0.host_ips.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "streaming_core_nodes.0.host_ips.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "analysis_task_nodes.0.host_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "streaming_task_nodes.0.host_ips.#", "1"),
				),
			},
			{
				Config: testAccMrsMapReduceClusterConfig_hybrid(rName, password, buildGroupNodeNumbers(3, 3, 2, 2)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "MIXED"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "analysis_core_nodes.0.node_number", "3"),
					resource.TestCheckResourceAttr(resourceName, "streaming_core_nodes.0.node_number", "3"),
					resource.TestCheckResourceAttr(resourceName, "analysis_task_nodes.0.node_number", "2"),
					resource.TestCheckResourceAttr(resourceName, "streaming_task_nodes.0.node_number", "2"),
					resource.TestCheckResourceAttr(resourceName, "analysis_core_nodes.0.host_ips.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "streaming_core_nodes.0.host_ips.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "analysis_task_nodes.0.host_ips.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "streaming_task_nodes.0.host_ips.#", "2"),
				),
			},
			{
				Config: testAccMrsMapReduceClusterConfig_hybrid(rName, password, buildGroupNodeNumbers(2, 2, 1, 1)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "MIXED"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "analysis_core_nodes.0.node_number", "2"),
					resource.TestCheckResourceAttr(resourceName, "streaming_core_nodes.0.node_number", "2"),
					resource.TestCheckResourceAttr(resourceName, "analysis_task_nodes.0.node_number", "1"),
					resource.TestCheckResourceAttr(resourceName, "streaming_task_nodes.0.node_number", "1"),
					resource.TestCheckResourceAttr(resourceName, "analysis_core_nodes.0.host_ips.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "streaming_core_nodes.0.host_ips.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "analysis_task_nodes.0.host_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "streaming_task_nodes.0.host_ips.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"manager_admin_pass",
					"node_admin_pass",
				},
			},
		},
	})
}

func TestAccMrsMapReduceCluster_custom_compact(t *testing.T) {
	var clusterGet cluster.Cluster
	resourceName := "huaweicloud_mapreduce_cluster.test"
	rName := acceptance.RandomAccResourceName()
	password := fmt.Sprintf("TF%s%s%d", acctest.RandString(10), acctest.RandStringFromCharSet(1, "-_"),
		acctest.RandIntRange(0, 99))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckMrsCustom(t)
		},
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckMRSV2ClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMrsMapReduceClusterConfig_customCompact(rName, password, 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.node_number", "3"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.host_ips.#", "3"),
				),
			},
			{
				Config: testAccMrsMapReduceClusterConfig_customCompact(rName, password, 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.node_number", "4"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.host_ips.#", "4"),
				),
			},
			{
				Config: testAccMrsMapReduceClusterConfig_customCompact(rName, password, 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.node_number", "3"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.host_ips.#", "3"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"manager_admin_pass",
					"node_admin_pass",
				},
			},
		},
	})
}

func TestAccMrsMapReduceCluster_custom_seperate(t *testing.T) {
	var clusterGet cluster.Cluster
	resourceName := "huaweicloud_mapreduce_cluster.test"
	rName := acceptance.RandomAccResourceName()
	password := fmt.Sprintf("TF%s%s%d", acctest.RandString(10), acctest.RandStringFromCharSet(1, "-_"),
		acctest.RandIntRange(0, 99))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckMrsCustom(t)
		},
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckMRSV2ClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMrsMapReduceClusterConfig_customSeperate(rName, password, 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.node_number", "3"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.host_ips.#", "3"),
				),
			},
			{
				Config: testAccMrsMapReduceClusterConfig_customSeperate(rName, password, 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.node_number", "4"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.host_ips.#", "4"),
				),
			},
			{
				Config: testAccMrsMapReduceClusterConfig_customSeperate(rName, password, 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.node_number", "3"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.host_ips.#", "3"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"manager_admin_pass",
					"node_admin_pass",
				},
			},
		},
	})
}

func TestAccMrsMapReduceCluster_custom_fullsize(t *testing.T) {
	var clusterGet cluster.Cluster
	resourceName := "huaweicloud_mapreduce_cluster.test"
	rName := acceptance.RandomAccResourceName()
	password := fmt.Sprintf("TF%s%s%d", acctest.RandString(10), acctest.RandStringFromCharSet(1, "-_"),
		acctest.RandIntRange(0, 99))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckMrsCustom(t)
		},
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckMRSV2ClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMrsMapReduceClusterConfig_customFullsize(rName, password, 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.node_number", "3"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.host_ips.#", "3"),
				),
			},
			{
				Config: testAccMrsMapReduceClusterConfig_customFullsize(rName, password, 4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.node_number", "4"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.host_ips.#", "4"),
				),
			},
			{
				Config: testAccMrsMapReduceClusterConfig_customFullsize(rName, password, 3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMRSV2ClusterExists(resourceName, &clusterGet),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "CUSTOM"),
					resource.TestCheckResourceAttr(resourceName, "safe_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "status", "running"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.node_number", "3"),
					resource.TestCheckResourceAttr(resourceName, "custom_nodes.0.host_ips.#", "3"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"manager_admin_pass",
					"node_admin_pass",
				},
			},
		},
	})
}

func buildGroupNodeNumbers(analysisCoreNum, streamCoreNum, analysisTaskNum, streamTaskNum int) GroupNodeNum {
	return GroupNodeNum{
		AnalysisCoreNum: analysisCoreNum,
		StreamCoreNum:   streamCoreNum,
		AnalysisTaskNum: analysisTaskNum,
		StreamTaskNum:   streamTaskNum,
	}
}

func testAccCheckMRSV2ClusterDestroy(s *terraform.State) error {
	config := acceptance.TestAccProvider.Meta().(*config.Config)
	client, err := config.MrsV1Client(acceptance.HW_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating huaweicloud mrs: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "huaweicloud_mapreduce_cluster" {
			continue
		}

		clusterGet, err := cluster.Get(client, rs.Primary.ID).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return nil
			}
			return fmt.Errorf("MRS cluster (%s) is still exists", rs.Primary.ID)
		}
		if clusterGet.Clusterstate == "terminated" {
			return nil
		}
	}

	return nil
}

func testAccCheckMRSV2ClusterExists(n string, clusterGet *cluster.Cluster) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource %s not found", n)
		}

		if rs.Primary.ID == "" {
			return fmtp.Errorf("No MRS cluster ID")
		}

		config := acceptance.TestAccProvider.Meta().(*config.Config)
		mrsClient, err := config.MrsV1Client(acceptance.HW_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating huaweicloud MRS client: %s ", err)
		}

		found, err := cluster.Get(mrsClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}
		*clusterGet = *found
		return nil
	}
}

func testAccMrsMapReduceClusterConfig_base(rName string) string {
	return fmt.Sprintf(`
data "huaweicloud_availability_zones" "test" {}

resource "huaweicloud_vpc" "test" {
  name = "%s"
  cidr = "192.168.0.0/16"
}

resource "huaweicloud_vpc_subnet" "test" {
  name       = "%s"
  cidr       = "192.168.0.0/20"
  vpc_id     = huaweicloud_vpc.test.id
  gateway_ip = "192.168.0.1"
}
`, rName, rName)
}

// The task node has not contain data disks.
func testAccMrsMapReduceClusterConfig_basic(rName, pwd string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = "%s"
  type               = "STREAMING"
  version            = "MRS 1.9.2"
  manager_admin_pass = "%s"
  node_admin_pass    = "%s"
  subnet_id          = huaweicloud_vpc_subnet.test.id
  vpc_id             = huaweicloud_vpc.test.id
  component_list     = ["Storm"]

  master_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_task_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 1
    root_volume_type  = "SAS"
    root_volume_size  = 300
  }

  tags = {
    foo = "bar"
    key = "value"
  }
}`, testAccMrsMapReduceClusterConfig_base(rName), rName, pwd, pwd)
}

func testAccMrsMapReduceClusterConfig_update(rName, pwd string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = "%s"
  type               = "STREAMING"
  version            = "MRS 1.9.2"
  manager_admin_pass = "%s"
  node_admin_pass    = "%s"
  subnet_id          = huaweicloud_vpc_subnet.test.id
  vpc_id             = huaweicloud_vpc.test.id
  component_list     = ["Storm"]

  master_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_task_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 1
    root_volume_type  = "SAS"
    root_volume_size  = 300
  }

  tags = {
    foo1 = "bar"
    key  = "update_value"
  }
}`, testAccMrsMapReduceClusterConfig_base(rName), rName, pwd, pwd)
}

func testAccMrsMapReduceClusterConfig_keypair(rName, pwd string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_compute_keypair" "test" {
  name = "%s"

  lifecycle {
    ignore_changes = [
      public_key,
    ]
  }
}

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = "%s"
  type               = "STREAMING"
  version            = "MRS 1.9.2"
  manager_admin_pass = "%s"
  node_key_pair      = huaweicloud_compute_keypair.test.name
  subnet_id          = huaweicloud_vpc_subnet.test.id
  vpc_id             = huaweicloud_vpc.test.id
  component_list     = ["Storm"]

  master_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
}`, testAccMrsMapReduceClusterConfig_base(rName), rName, rName, pwd)
}

func testAccMrsMapReduceClusterConfig_analysis(rName, pwd string, nodeNums GroupNodeNum) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = "%s"
  type               = "ANALYSIS"
  version            = "MRS 1.9.2"
  manager_admin_pass = "%s"
  node_admin_pass    = "%s"
  subnet_id          = huaweicloud_vpc_subnet.test.id
  vpc_id             = huaweicloud_vpc.test.id
  component_list     = ["Hadoop", "Hive", "Tez"]

  master_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 600
    data_volume_type  = "SAS"
    data_volume_size  = 600
    data_volume_count = 1
  }
  analysis_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = %d
    root_volume_type  = "SAS"
    root_volume_size  = 600
    data_volume_type  = "SAS"
    data_volume_size  = 600
    data_volume_count = 1
  }
  analysis_task_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = %d
    root_volume_type  = "SAS"
    root_volume_size  = 600
    data_volume_type  = "SAS"
    data_volume_size  = 600
    data_volume_count = 1
  }
}`, testAccMrsMapReduceClusterConfig_base(rName), rName, pwd, pwd,
		nodeNums.AnalysisCoreNum, nodeNums.AnalysisTaskNum)
}

func testAccMrsMapReduceClusterConfig_stream(rName, pwd string, nodeNums GroupNodeNum) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = "%s"
  type               = "STREAMING"
  version            = "MRS 1.9.2"
  manager_admin_pass = "%s"
  node_admin_pass    = "%s"
  subnet_id          = huaweicloud_vpc_subnet.test.id
  vpc_id             = huaweicloud_vpc.test.id
  component_list     = ["Storm"]

  master_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = %d
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_task_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = %d
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
}`, testAccMrsMapReduceClusterConfig_base(rName), rName, pwd, pwd,
		nodeNums.StreamCoreNum, nodeNums.StreamTaskNum)
}

func testAccMrsMapReduceClusterConfig_hybrid(rName, pwd string, nodeNums GroupNodeNum) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = "%s"
  type               = "MIXED"
  version            = "MRS 1.9.2"
  safe_mode          = true
  manager_admin_pass = "%s"
  node_admin_pass    = "%s"
  subnet_id          = huaweicloud_vpc_subnet.test.id
  vpc_id             = huaweicloud_vpc.test.id
  component_list     = ["Hadoop", "Spark", "Hive", "Tez", "Storm"]

  master_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  analysis_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = %d
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = %d
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  analysis_task_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = %d
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_task_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = %d
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
}`, testAccMrsMapReduceClusterConfig_base(rName), rName, pwd, pwd,
		nodeNums.AnalysisCoreNum, nodeNums.StreamCoreNum, nodeNums.AnalysisTaskNum, nodeNums.StreamTaskNum)
}

func testAccMrsMapReduceClusterConfig_customCompact(rName, pwd string, nodeNum1 int) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = "%s"
  type               = "CUSTOM"
  version            = "MRS 3.1.0"
  safe_mode          = true
  manager_admin_pass = "%s"
  node_admin_pass    = "%s"
  subnet_id          = huaweicloud_vpc_subnet.test.id
  vpc_id             = huaweicloud_vpc.test.id
  template_id        = "mgmt_control_combined_v4"
  component_list     = ["Hadoop", "ZooKeeper", "Ranger"]

master_nodes {
    flavor            = "c6.4xlarge.4.linux.bigdata"
    node_number       = 3
    root_volume_type  = "SAS"
    root_volume_size  = 480
    data_volume_type  = "SAS"
    data_volume_size  = 600
    data_volume_count = 1
    assigned_roles = [
      "OMSServer:1,2",
      "SlapdServer:1,2",
      "KerberosServer:1,2",
      "KerberosAdmin:1,2",
      "quorumpeer:1,2,3",
      "NameNode:2,3",
      "Zkfc:2,3",
      "JournalNode:1,2,3",
      "ResourceManager:2,3",
      "JobHistoryServer:3",
      "DBServer:1,3",
      "HttpFS:1,3",
      "TimelineServer:3",
      "RangerAdmin:1,2",
      "UserSync:2",
      "TagSync:2",
      "KerberosClient",
      "SlapdClient",
      "meta"
    ]
  }

  custom_nodes {
    group_name        = "node_group_1"
    flavor            = "c6.4xlarge.4.linux.bigdata"
    node_number       = %d
    root_volume_type  = "SAS"
    root_volume_size  = 480
    data_volume_type  = "SAS"
    data_volume_size  = 600
    data_volume_count = 1
    assigned_roles = [
      "DataNode",
      "NodeManager",
      "RegionServer",
      "Flume",
      "KerberosClient",
      "SlapdClient",
      "meta"
    ]

  }
  
}`, testAccMrsMapReduceClusterConfig_base(rName), rName, pwd, pwd, nodeNum1)
}

func testAccMrsMapReduceClusterConfig_customSeperate(rName, pwd string, nodeNum1 int) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = "%s"
  type               = "CUSTOM"
  version            = "MRS 3.1.0"
  safe_mode          = true
  manager_admin_pass = "%s"
  node_admin_pass    = "%s"
  subnet_id          = huaweicloud_vpc_subnet.test.id
  vpc_id             = huaweicloud_vpc.test.id
  template_id        = "mgmt_control_separated_v4"
  component_list     = ["DBService", "Hadoop", "ZooKeeper", "Ranger"]

master_nodes {
    flavor            = "c6.4xlarge.4.linux.bigdata"
    node_number       = 5
    root_volume_type  = "SAS"
    root_volume_size  = 480
    data_volume_type  = "SAS"
    data_volume_size  = 600
    data_volume_count = 1
    assigned_roles = [
      "OMSServer:1,2",
      "SlapdServer:3,4",
      "KerberosServer:3,4",
      "KerberosAdmin:3,4",
      "quorumpeer:3,4,5",
      "NameNode:4,5",
      "Zkfc:4,5",
      "JournalNode:3,4,5",
      "ResourceManager:4,5",
      "JobHistoryServer:5",
      "DBServer:3,5",
      "HttpFS:3,5",
      "TimelineServer:5",
      "RangerAdmin:3,4",
      "UserSync:4",
      "TagSync:4",
      "KerberosClient",
      "SlapdClient",
      "meta"
    ]
  }

  custom_nodes {
    group_name        = "node_group_1"
    flavor            = "c6.4xlarge.4.linux.bigdata"
    node_number       = %d
    root_volume_type  = "SAS"
    root_volume_size  = 480
    data_volume_type  = "SAS"
    data_volume_size  = 600
    data_volume_count = 1
    assigned_roles = [
      "DataNode",
      "NodeManager",
      "KerberosClient",
      "SlapdClient",
      "meta"
    ]
  }

}`, testAccMrsMapReduceClusterConfig_base(rName), rName, pwd, pwd, nodeNum1)
}

func testAccMrsMapReduceClusterConfig_customFullsize(rName, pwd string, nodeNum1 int) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = "%s"
  type               = "CUSTOM"
  version            = "MRS 3.1.0"
  safe_mode          = true
  manager_admin_pass = "%s"
  node_admin_pass    = "%s"
  subnet_id          = huaweicloud_vpc_subnet.test.id
  vpc_id             = huaweicloud_vpc.test.id
  template_id        = "mgmt_control_data_separated_v4"
  component_list     = ["Hadoop", "Ranger", "ZooKeeper","DBServer"]

  master_nodes {
    flavor            = "c6.4xlarge.4.linux.bigdata"
    node_number       = 9
    root_volume_type  = "SAS"
    root_volume_size  = 480
    data_volume_type  = "SAS"
    data_volume_size  = 600
    data_volume_count = 1
    assigned_roles = [
      "OMSServer:1,2",
      "SlapdServer:5,6",
      "KerberosServer:5,6",
      "KerberosAdmin:5,6",
      "quorumpeer:5,6,7,8,9",
      "NameNode:3,4",
      "Zkfc:3,4",
      "JournalNode:5,6,7",
      "ResourceManager:8,9",
      "JobHistoryServer:8",
      "DBServer:8,9",
      "HttpFS:8,9",
      "TimelineServer:5",
      "RangerAdmin:4,5",
      "UserSync:5",
      "TagSync:5",
      "KerberosClient",
      "SlapdClient",
      "meta"
    ]
  }

  custom_nodes {
    group_name        = "node_group_1"
    flavor            = "c6.4xlarge.4.linux.bigdata"
    node_number       = %d
    root_volume_type  = "SAS"
    root_volume_size  = 480
    data_volume_type  = "SAS"
    data_volume_size  = 600
    data_volume_count = 1
    assigned_roles = [
      "DataNode",
      "NodeManager",
      "KerberosClient",
      "SlapdClient",
      "meta"
    ]
  }

}`, testAccMrsMapReduceClusterConfig_base(rName), rName, pwd, pwd, nodeNum1)
}
