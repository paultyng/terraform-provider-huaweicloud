package fgs

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/huaweicloud/golangsdk/openstack/fgs/v2/trigger"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccFunctionGraphTriggerV2_basic(t *testing.T) {
	var (
		rName        = fmt.Sprintf("tf-acc-test-%s", acctest.RandString(5))
		resourceName = "huaweicloud_fgs_trigger.test"
		timeTrigger  trigger.Trigger
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAccPreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckFunctionGraphTriggerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFunctionGraphTrigger_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFunctionGraphTriggerExists(resourceName, &timeTrigger),
					resource.TestCheckResourceAttr(resourceName, "timer.0.name", rName),
					resource.TestCheckResourceAttr(resourceName, "timer.0.schedule_type", "Rate"),
					resource.TestCheckResourceAttr(resourceName, "timer.0.schedule", "3d"),
					resource.TestCheckResourceAttr(resourceName, "type", "TIMER"),
					resource.TestCheckResourceAttr(resourceName, "status", "ACTIVE"),
				),
			},
			{
				Config: testAccFunctionGraphTrigger_update(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFunctionGraphTriggerExists(resourceName, &timeTrigger),
					resource.TestCheckResourceAttr(resourceName, "timer.0.name", rName),
					resource.TestCheckResourceAttr(resourceName, "timer.0.schedule_type", "Rate"),
					resource.TestCheckResourceAttr(resourceName, "timer.0.schedule", "3d"),
					resource.TestCheckResourceAttr(resourceName, "type", "TIMER"),
					resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccFunctionGraphSubResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccFunctionGraphTriggerV2_cronTimer(t *testing.T) {
	var (
		rName        = fmt.Sprintf("tf-acc-test-%s", acctest.RandString(5))
		resourceName = "huaweicloud_fgs_trigger.test"
		timeTrigger  trigger.Trigger
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAccPreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckFunctionGraphTriggerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFunctionGraphTrigger_cron(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFunctionGraphTriggerExists(resourceName, &timeTrigger),
					resource.TestCheckResourceAttr(resourceName, "timer.0.name", rName),
					resource.TestCheckResourceAttr(resourceName, "timer.0.schedule_type", "Cron"),
					resource.TestCheckResourceAttr(resourceName, "timer.0.schedule", "@every 1h30m"),
					resource.TestCheckResourceAttr(resourceName, "type", "TIMER"),
					resource.TestCheckResourceAttr(resourceName, "status", "ACTIVE"),
				),
			},
			{
				Config: testAccFunctionGraphTrigger_cronUpdate(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFunctionGraphTriggerExists(resourceName, &timeTrigger),
					resource.TestCheckResourceAttr(resourceName, "timer.0.name", rName),
					resource.TestCheckResourceAttr(resourceName, "timer.0.schedule_type", "Cron"),
					resource.TestCheckResourceAttr(resourceName, "timer.0.schedule", "@every 1h30m"),
					resource.TestCheckResourceAttr(resourceName, "type", "TIMER"),
					resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccFunctionGraphSubResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccCheckFunctionGraphTriggerDestroy(s *terraform.State) error {
	config := acceptance.TestAccProvider.Meta().(*config.Config)
	client, err := config.ApigV2Client(acceptance.HW_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating HuaweiCloud APIG v2 client: %s", err)
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "huaweicloud_fgs_trigger" {
			continue
		}
		_, err := trigger.Get(client, rs.Primary.Attributes["function_urn"], rs.Primary.Attributes["type"],
			rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("FunctionGraph v2 trigger (%s) is still exists", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckFunctionGraphTriggerExists(appName string, t *trigger.Trigger) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[appName]
		if !ok {
			return fmt.Errorf("Resource %s not found", appName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No FunctionGraph V2 Trigger ID")
		}

		config := acceptance.TestAccProvider.Meta().(*config.Config)
		client, err := config.FgsV2Client(acceptance.HW_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating HuaweiCloud APIG v2 client: %s", err)
		}
		found, err := trigger.Get(client, rs.Primary.Attributes["function_urn"], rs.Primary.Attributes["type"],
			rs.Primary.ID).Extract()
		if err != nil {
			return err
		}
		*t = *found
		return nil
	}
}

func testAccFunctionGraphSubResourceImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Resource (%s) not found: %s", name, rs)
		}
		if rs.Primary.ID == "" || rs.Primary.Attributes["function_urn"] == "" {
			return "", fmt.Errorf("resource not found: %s/%s", rs.Primary.Attributes["function_urn"], rs.Primary.ID)
		}
		return fmt.Sprintf("%s/%s", rs.Primary.Attributes["function_urn"], rs.Primary.ID), nil
	}
}

func testAccFunctionGraphTrigger_base(rName string) string {
	return fmt.Sprintf(`
resource "huaweicloud_fgs_function" "test" {
  name        = "%s"
  app         = "default"
  handler     = "index.handler"
  memory_size = 128
  timeout     = 3
  runtime     = "Python2.7"
  code_type   = "inline"
  func_code   = "aW1wb3J0IGpzb24KZGVmIGhhbmRsZXIgKGV2ZW50LCBjb250ZXh0KToKICAgIG91dHB1dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganNvbi5kdW1wcyhldmVudCkKICAgIHJldHVybiBvdXRwdXQ="
}
`, rName)
}

func testAccFunctionGraphTrigger_basic(rName string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_fgs_trigger" "test" {
  function_urn = huaweicloud_fgs_function.test.urn
  type         = "TIMER"

  timer {
    name          = "%s"
    schedule_type = "Rate"
    schedule      = "3d"
  }
}
`, testAccFunctionGraphTrigger_base(rName), rName)
}

func testAccFunctionGraphTrigger_update(rName string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_fgs_trigger" "test" {
  function_urn = huaweicloud_fgs_function.test.urn
  type         = "TIMER"
  status       = "DISABLED"

  timer {
	name          = "%s"
	schedule_type = "Rate"
	schedule      = "3d"
  }
}
`, testAccFunctionGraphTrigger_base(rName), rName)
}

func testAccFunctionGraphTrigger_cron(rName string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_fgs_trigger" "test" {
  function_urn = huaweicloud_fgs_function.test.urn
  type         = "TIMER"

  timer {
    name          = "%s"
    schedule_type = "Cron"
    schedule      = "@every 1h30m"
  }
}
`, testAccFunctionGraphTrigger_base(rName), rName)
}

func testAccFunctionGraphTrigger_cronUpdate(rName string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_fgs_trigger" "test" {
  function_urn = huaweicloud_fgs_function.test.urn
  type         = "TIMER"
  status       = "DISABLED"

  timer {
	name          = "%s"
	schedule_type = "Cron"
	schedule      = "@every 1h30m"
  }
}
`, testAccFunctionGraphTrigger_base(rName), rName)
}
