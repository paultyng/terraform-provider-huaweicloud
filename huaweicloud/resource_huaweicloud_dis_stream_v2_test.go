// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://www.github.com/huaweicloud/magic-modules
//
// ----------------------------------------------------------------------------

package huaweicloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/huaweicloud/golangsdk"
)

func TestAccDisStreamV2_basic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDisStreamV2Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDisStreamV2_basic(acctest.RandString(10)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDisStreamV2Exists(),
				),
			},
		},
	})
}

func testAccDisStreamV2_basic(val string) string {
	return fmt.Sprintf(`
resource "huaweicloud_dis_stream_v2" "stream" {
  stream_name = "terraform_test_dis_stream%s"
  partition_count = 1
}
	`, val)
}

func testAccCheckDisStreamV2Destroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.disV2Client(HW_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "huaweicloud_dis_stream_v2" {
			continue
		}

		url, err := replaceVarsForTest(rs, "streams/{id}")
		if err != nil {
			return err
		}
		url = client.ServiceURL(url)

		_, err = client.Get(url, nil, &golangsdk.RequestOpts{
			MoreHeaders: map[string]string{"Content-Type": "application/json"}})
		if err == nil {
			return fmt.Errorf("huaweicloud_dis_stream_v2 still exists at %s", url)
		}
	}

	return nil
}

func testAccCheckDisStreamV2Exists() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)
		client, err := config.disV2Client(HW_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating sdk client, err=%s", err)
		}

		rs, ok := s.RootModule().Resources["huaweicloud_dis_stream_v2.stream"]
		if !ok {
			return fmt.Errorf("Error checking huaweicloud_dis_stream_v2.stream exist, err=not found this resource")
		}

		url, err := replaceVarsForTest(rs, "streams/{id}")
		if err != nil {
			return fmt.Errorf("Error checking huaweicloud_dis_stream_v2.stream exist, err=building url failed: %s", err)
		}
		url = client.ServiceURL(url)

		_, err = client.Get(url, nil, &golangsdk.RequestOpts{
			MoreHeaders: map[string]string{"Content-Type": "application/json"}})
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return fmt.Errorf("huaweicloud_dis_stream_v2.stream is not exist")
			}
			return fmt.Errorf("Error checking huaweicloud_dis_stream_v2.stream exist, err=send request failed: %s", err)
		}
		return nil
	}
}
