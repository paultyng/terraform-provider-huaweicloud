package codeartsinspector

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/codeartsinspector"
)

func getInspectorHostResourceFunc(cfg *config.Config, state *terraform.ResourceState) (interface{}, error) {
	client, err := cfg.NewServiceClient("vss", acceptance.HW_REGION_NAME)
	if err != nil {
		return nil, fmt.Errorf("error creating CodeArts inspector client: %s", err)
	}

	return codeartsinspector.GetInspectorHost(client, state.Primary.ID)
}

func TestAccInspectorHost_basic(t *testing.T) {
	var obj interface{}

	name := acceptance.RandomAccResourceName()
	rName := "huaweicloud_codearts_inspector_host.test"

	rc := acceptance.InitResourceCheck(
		rName,
		&obj,
		getInspectorHostResourceFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckCodeArtsSshCredentialID(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testInspectorHost_basic(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttrPair(rName, "ip", "huaweicloud_compute_instance.test", "public_ip"),
					resource.TestCheckResourceAttr(rName, "os_type", "linux"),
					resource.TestCheckResourceAttr(rName, "ssh_credential_id", acceptance.HW_CODEARTS_SSH_CREDENTIAL_ID),
					resource.TestCheckResourceAttrSet(rName, "auth_status"),
				),
			},
			{
				ResourceName:      rName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testInspectorHost_ECS_Base(name string) string {
	return fmt.Sprintf(`
resource "huaweicloud_compute_instance" "test" {
  depends_on = [huaweicloud_networking_secgroup_rule.out_v4_all]

  name                        = "%[1]s"
  image_id                    = data.huaweicloud_images_image.test.id
  flavor_id                   = data.huaweicloud_compute_flavors.test.ids[0]
  availability_zone           = data.huaweicloud_availability_zones.test.names[0]
  admin_pass                  = "Terraform@123"
  delete_disks_on_termination = true
  security_group_ids          = [huaweicloud_networking_secgroup.test.id]
  
  eip_type = "5_bgp"
  bandwidth {
    share_type  = "PER"
    size        = 1
    charge_mode = ""
  }

  network {
    uuid = huaweicloud_vpc_subnet.test.id
  }
}
`, name)
}

func testInspectorHost_basic(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_networking_secgroup_rule" "out_v4_all" {
  security_group_id = huaweicloud_networking_secgroup.test.id
  ethertype         = "IPv4"
  direction         = "egress"
  remote_ip_prefix  = "0.0.0.0/0"
}

%[2]s

resource "huaweicloud_codearts_inspector_host" "test" {
  name              = "%[3]s"
  ip                = huaweicloud_compute_instance.test.public_ip
  os_type           = "linux"
  ssh_credential_id = "%[4]s"
}
`, common.TestBaseComputeResources(name), testInspectorHost_ECS_Base(name), name, acceptance.HW_CODEARTS_SSH_CREDENTIAL_ID)
}
