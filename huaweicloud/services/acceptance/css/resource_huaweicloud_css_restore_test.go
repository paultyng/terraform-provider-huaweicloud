package css

import (
	"fmt"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance/common"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccCssInstanceRestore_basic(t *testing.T) {
	name := acceptance.RandomAccResourceName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCssInstanceRestoreConfig_basic(name),
			},
		},
	})
}

func testAccCssInstanceRestoreConfig_basic(name string) string {
	return fmt.Sprintf(`
%[1]s

data "huaweicloud_availability_zones" "test" {}

resource "huaweicloud_obs_bucket" "cssObs" {
  bucket        = "%s"
  acl           = "private"
  force_destroy = true
}

resource "huaweicloud_css_cluster" "test" {
  count			 = 2
  name           = "%[2]s__${count.index}"
  engine_version = "7.10.2"
  security_mode  = true
  password       = "Test@passw0rd"

  ess_node_config {
    flavor          = "ess.spec-4u8g"
    instance_number = 1
    volume {
      volume_type = "HIGH"
      size        = 40
    }
  }

  availability_zone = data.huaweicloud_availability_zones.test.names[0]
  security_group_id = huaweicloud_networking_secgroup.test.id
  subnet_id         = huaweicloud_vpc_subnet.test.id
  vpc_id            = huaweicloud_vpc.test.id

  backup_strategy {
    keep_days   = 1
    start_time  = "00:00 GMT+08:00"
    prefix      = "snapshot"
    bucket      = huaweicloud_obs_bucket.cssObs.bucket
    agency      = "css_obs_agency"
    backup_path = "css_repository/acctest"
  }
}
resource "huaweicloud_css_snapshot" "snapshot" {
  name        = "snapshot-%[2]s"
  description = "a snapshot created by terraform acctest"
  cluster_id  = huaweicloud_css_cluster.test[0].id
}

resource "huaweicloud_css_restore" "test" {
  source_cluster_id 	= huaweicloud_css_cluster.test[0].id
  target_cluster_id 	= huaweicloud_css_cluster.test[1].id
  snapshot_id        	= huaweicloud_css_snapshot.snapshot.id
}
`, common.TestBaseNetwork(name), name)
}
