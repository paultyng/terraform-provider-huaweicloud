package cae

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccComponentAction_basic(t *testing.T) {
	baseConfig := testAccComponentAction_base()
	// Avoid CheckDestroy because this resource is a one-time action resource and there is nothing in the destroy
	// method.
	// lintignore:AT001
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckCaeEnvironment(t)
			acceptance.TestAccPreCheckCaeApplication(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				// One-time action resource do not need to be checked and no processing is performed in the Read method.
				Config: testAccComponentAction_basic_step1(baseConfig, "deploy"),
			},
			{
				Config: testAccComponentAction_basic_step2(baseConfig),
			},
			{
				Config: testAccComponentAction_basic_step3(baseConfig),
			},
			{
				Config: testAccComponentAction_basic_step4(baseConfig),
			},
			{
				Config: testAccComponentAction_basic_step5(baseConfig),
			},
			{
				Config: testAccComponentAction_basic_step6(baseConfig),
			},

			{
				Config: testAccComponentAction_basic_step7(baseConfig),
			},
		},
	})
}

func testAccComponentAction_base() string {
	name := acceptance.RandomAccResourceNameWithDash()

	return fmt.Sprintf(`
data "huaweicloud_swr_repositories" "test" {}

locals {
  swr_repositories = [for v in data.huaweicloud_swr_repositories.test.repositories : v if length(v.tags) > 0]
}

resource "huaweicloud_cae_component" "test" {
  environment_id = "%[1]s"
  application_id = "%[2]s"

  metadata {
    name = "%[3]s"

    annotations = {
      version = "1.0.0"
    }
  }

  spec {
    replica = 1
    runtime = "Docker"

    source {
      type = "image"
      url  = try(format("%%s:%%s", local.swr_repositories[0].path, local.swr_repositories[0].tags[0]), "")
    }

    resource_limit {
      cpu    = "500m"
      memory = "1Gi"
    }
  }

  lifecycle {
    ignore_changes = [
      metadata.0.annotations.version, spec.0.source, spec.0.resource_limit
    ]
  }
}
`, acceptance.HW_CAE_ENVIRONMENT_ID, acceptance.HW_CAE_APPLICATION_ID, name)
}

func testAccComponentAction_basic_step1(baseConfig, action string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_cae_component_action" "test" {
  environment_id = "%[2]s"
  application_id = "%[3]s"
  component_id   = huaweicloud_cae_component.test.id

  metadata {
    name = "%[4]s"
  }
}
`, baseConfig, acceptance.HW_CAE_ENVIRONMENT_ID, acceptance.HW_CAE_APPLICATION_ID, action)
}

func testAccComponentAction_basic_step2(baseConfig string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_cae_component_action" "test" {
  environment_id = "%[2]s"
  application_id = "%[3]s"
  component_id   = huaweicloud_cae_component.test.id

  metadata {
    name = "upgrade"

    annotations = {
      version = "2.0.0"
    }
  }

  spec = jsonencode({
    "source" : {
      "type" : "image",
      "url" : "nginx:alpine-perl"
    },
    "resource_limit" : {
      "cpu_limit" : "1000m",
      "memory_limit" : "2Gi"
    }
  })
}
`, baseConfig, acceptance.HW_CAE_ENVIRONMENT_ID, acceptance.HW_CAE_APPLICATION_ID)
}

func testAccComponentAction_basic_step3(baseConfig string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_cae_component_action" "test" {
  environment_id = "%[2]s"
  application_id = "%[3]s"
  component_id   = huaweicloud_cae_component.test.id

  metadata {
    name = "rollback"

    annotations = {
      version = "1.0.0"
    }
  }

  spec = jsonencode({
    "snapshot_index" : 1
  })
}
`, baseConfig, acceptance.HW_CAE_ENVIRONMENT_ID, acceptance.HW_CAE_APPLICATION_ID)
}

func testAccComponentAction_basic_step4(baseConfig string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_cae_component_configurations" "test" {
  environment_id = "%[2]s"
  application_id = "%[3]s"
  component_id   = huaweicloud_cae_component.test.id

  items {
    type = "env"
    data = jsonencode({
      "spec": {
        "envs": {
            "key": "value",
            "foo": "baar"
        }
      }
    })
  }
}

resource "huaweicloud_cae_component_action" "test" {
  environment_id = "%[2]s"
  application_id = "%[3]s"
  component_id   = huaweicloud_cae_component.test.id

  metadata {
    name = "configure"
  }

  lifecycle {
    replace_triggered_by = [
      huaweicloud_cae_component_configurations.test.items
    ]
  }
}
`, baseConfig, acceptance.HW_CAE_ENVIRONMENT_ID, acceptance.HW_CAE_APPLICATION_ID)
}

func testAccComponentAction_basic_step5(baseConfig string) string {
	return testAccComponentAction_basic_step1(baseConfig, "stop")
}

func testAccComponentAction_basic_step6(baseConfig string) string {
	return testAccComponentAction_basic_step1(baseConfig, "start")
}

func testAccComponentAction_basic_step7(baseConfig string) string {
	return testAccComponentAction_basic_step1(baseConfig, "restart")
}
