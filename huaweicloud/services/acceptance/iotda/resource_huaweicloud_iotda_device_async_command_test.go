package iotda

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/iotda"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func getDeviceAsyncCommandFunc(cfg *config.Config, state *terraform.ResourceState) (interface{}, error) {
	var (
		region    = acceptance.HW_REGION_NAME
		isDerived = iotda.WithDerivedAuth(cfg, region)
		httpUrl   = "v5/iot/{project_id}/devices/{device_id}/async-commands/{command_id}"
	)

	client, err := cfg.NewServiceClientWithDerivedAuth("iotda", region, isDerived)
	if err != nil {
		return nil, fmt.Errorf("error creating IoTDA client: %s", err)
	}

	getPath := client.Endpoint + httpUrl
	getPath = strings.ReplaceAll(getPath, "{project_id}", client.ProjectID)
	getPath = strings.ReplaceAll(getPath, "{device_id}", state.Primary.Attributes["device_id"])
	getPath = strings.ReplaceAll(getPath, "{command_id}", state.Primary.ID)
	getOpts := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	getResp, err := client.Request("GET", getPath, &getOpts)
	if err != nil {
		return nil, fmt.Errorf("error retrieving device asynchronous command: %s", err)
	}

	return utils.FlattenResponse(getResp)
}

func TestAccDeviceAsyncCommand_basic(t *testing.T) {
	var (
		commondObj interface{}
		name       = acceptance.RandomAccResourceName()
		rName      = "huaweicloud_iotda_device_async_command.test"
	)

	rc := acceptance.InitResourceCheck(
		rName,
		&commondObj,
		getDeviceAsyncCommandFunc,
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckHWIOTDAAccessAddress(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceAsyncCommand_basic(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttrPair(rName, "device_id", "huaweicloud_iotda_device.test", "id"),
					resource.TestCheckResourceAttrPair(rName, "service_id", "huaweicloud_iotda_product.test", "services.0.id"),
					resource.TestCheckResourceAttrPair(rName, "name", "huaweicloud_iotda_product.test", "services.0.commands.0.name"),
					resource.TestCheckResourceAttr(rName, "send_strategy", "delay"),
					resource.TestCheckResourceAttr(rName, "expire_time", "80000"),
					resource.TestCheckResourceAttrSet(rName, "status"),
					resource.TestCheckResourceAttrSet(rName, "created_at"),
				),
			},
		},
	})
}

func testAccDeviceAsyncCommand_base(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_iotda_space" "test" {
  name = "%[2]s"
}

resource "huaweicloud_iotda_product" "test" {
  name        = "%[2]s"
  device_type = "AI"
  protocol    = "CoAP"
  space_id    = huaweicloud_iotda_space.test.id
  data_type   = "json"
  industry    = "smart-home"

  services {
    id   = "001211985996"
    type = "001002"

    commands {
      name = "cmd-test"

      paras {
        name       = "cmd-req"
        type       = "string"
        max_length = 20
      }

      responses {
        name       = "cmd-resp"
        type       = "string"
        max_length = 20
      }
    }
  }
}

resource "huaweicloud_iotda_device" "test" {
  node_id     = "101112026"
  name        = "%[2]s"
  space_id    = huaweicloud_iotda_space.test.id
  product_id  = huaweicloud_iotda_product.test.id
  description = "demo"
}
`, buildIoTDAEndpoint(), name)
}

func testAccDeviceAsyncCommand_basic(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_iotda_device_async_command" "test" {
  device_id     = huaweicloud_iotda_device.test.id
  service_id    = huaweicloud_iotda_product.test.services.0.id
  name          = huaweicloud_iotda_product.test.services.0.commands.0.name
  send_strategy = "delay"
  expire_time   = "80000"

  paras = {
    (huaweicloud_iotda_product.test.services.0.commands.0.paras.0.name) = "tf-acc"
  }
}
`, testAccDeviceAsyncCommand_base(name))
}
