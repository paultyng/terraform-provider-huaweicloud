package cph

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccCphPhoneDataExport_basic(t *testing.T) {
	name := acceptance.RandomAccResourceName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPrecheckCphObsBucketName(t)
			acceptance.TestAccPrecheckCphAdbObjectPath(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testCphPhoneDataExport_basic(name),
			},
			{
				Config: testCphServerBase(name),
				Check: resource.ComposeTestCheckFunc(
					waitForDeletionCooldownComplete(),
				),
			},
		},
	})
}

func testCphPhoneDataExport_basic(name string) string {
	return fmt.Sprintf(`
%[1]s

data "huaweicloud_cph_phones" "test" {
  server_id = huaweicloud_cph_server.test.id
}

resource "huaweicloud_cph_phone_data_export" "test" {
  phone_id      = data.huaweicloud_cph_phones.test.phones[0].phone_id
  bucket_name   = "%[2]s"
  object_path   = "%[3]s"
  include_files = ["/data/app", "/data/local", "/data/media"]
}
`, testCphServer_basic(name), acceptance.HW_CPH_OBS_BUCKET_NAME, acceptance.HW_CPH_OBS_OBJECT_PATH)
}
