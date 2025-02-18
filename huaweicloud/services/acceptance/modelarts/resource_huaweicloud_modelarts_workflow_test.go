package modelarts

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func getModelartsWorkflowResourceFunc(cfg *config.Config, state *terraform.ResourceState) (interface{}, error) {
	region := acceptance.HW_REGION_NAME
	var (
		getWorkflowHttpUrl = "v2/{project_id}/workflows/{id}"
		getWorkflowProduct = "modelarts"
	)
	getWorkflowClient, err := cfg.NewServiceClient(getWorkflowProduct, region)
	if err != nil {
		return nil, fmt.Errorf("error creating ModelArts client: %s", err)
	}

	getWorkflowPath := getWorkflowClient.Endpoint + getWorkflowHttpUrl
	getWorkflowPath = strings.ReplaceAll(getWorkflowPath, "{project_id}", getWorkflowClient.ProjectID)
	getWorkflowPath = strings.ReplaceAll(getWorkflowPath, "{id}", state.Primary.ID)

	getWorkflowOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	}

	getWorkflowResp, err := getWorkflowClient.Request("GET", getWorkflowPath, &getWorkflowOpt)
	if err != nil {
		return nil, fmt.Errorf("error retrieving ModelArts workflow: %s", err)
	}

	getWorkflowRespBody, err := utils.FlattenResponse(getWorkflowResp)
	if err != nil {
		return nil, fmt.Errorf("error retrieving ModelArts workflow: %s", err)
	}

	return getWorkflowRespBody, nil
}

func TestAccModelartsWorkflow_basic(t *testing.T) {
	var obj interface{}

	name := acceptance.RandomAccResourceName()
	rName := "huaweicloud_modelarts_workflow.test"

	rc := acceptance.InitResourceCheck(
		rName,
		&obj,
		getModelartsWorkflowResourceFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testModelartsWorkflow_basic(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name),
					resource.TestCheckResourceAttr(rName, "description", "This is a demo"),
					resource.TestCheckResourceAttr(rName, "workspace_id", "0"),
					resource.TestCheckResourceAttr(rName, "steps.#", "4"),
					resource.TestCheckResourceAttr(rName, "steps.0.name", "condition_step_test"),
					resource.TestCheckResourceAttr(rName, "labels.#", "1"),
					resource.TestCheckResourceAttr(rName, "labels.0", "subgraph"),
					resource.TestCheckResourceAttr(rName, "data.#", "3"),
					resource.TestCheckResourceAttr(rName, "data.0.name", "a2ff296da618452daa8243399f06db8e"),
					resource.TestCheckResourceAttr(rName, "parameters.#", "5"),
					resource.TestCheckResourceAttr(rName, "parameters.0.name", "is_true"),
					resource.TestCheckResourceAttrSet(rName, "param_ready"),
				),
			},
			{
				Config: testModelartsWorkflow_basic_update(name + "_update"),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName, "name", name+"_update"),
					resource.TestCheckResourceAttr(rName, "description", "This is a demo update"),
					resource.TestCheckResourceAttr(rName, "workspace_id", "0"),
					resource.TestCheckResourceAttr(rName, "steps.#", "4"),
					resource.TestCheckResourceAttr(rName, "steps.0.name", "condition_step_test"),
					resource.TestCheckResourceAttr(rName, "labels.#", "2"),
					resource.TestCheckResourceAttr(rName, "labels.0", "subgraph"),
					resource.TestCheckResourceAttr(rName, "labels.1", "test"),
					resource.TestCheckResourceAttr(rName, "data.#", "3"),
					resource.TestCheckResourceAttr(rName, "data.0.name", "a2ff296da618452daa8243399f06db8e"),
					resource.TestCheckResourceAttr(rName, "parameters.#", "6"),
					resource.TestCheckResourceAttr(rName, "parameters.0.name", "is_true"),
					resource.TestCheckResourceAttr(rName, "parameters.5.name", "service_config4"),
					resource.TestCheckResourceAttrSet(rName, "param_ready"),
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

func testModelartsWorkflow_basic(name string) string {
	return fmt.Sprintf(`
resource "huaweicloud_modelarts_workflow" "test" {
  name         = "%s"
  description  = "This is a demo"
  workspace_id = "0"

  steps {
    name  = "condition_step_test"
    title = "condition_step_test"
    type  = "condition"
    conditions {
      type  = "=="
      left  = "$ref/parameters/is_true"
      right = true
    }
    if_then_steps   = ["training_job1"]
    else_then_steps = ["training_job2"]
  }

  steps {
    name  = "training_job1"
    title = "labeling"
    type  = "job"
    inputs {
      name = "data_url"
      type = "obs"
      data = "$ref/data/a2ff296da618452daa8243399f06db8e"
    }
    outputs {
      name = "train_url"
      type = "obs"
      config = {
        obs_url = "/test-lh/test-metrics/"
      }
    }
    outputs {
      name = "service-link"
      type = "service_content"
      config = {
        config_file = "$ref/parameters/service_config1"
      }
    }

    properties = jsonencode({
      "algorithm" : {
        "id" : "21ef85a8-5e40-4618-95ee-aa48ec224b43",
        "parameters" : []
      },
      "kind" : "job",
      "metadata" : {
        "name" : "workflow-b0b9fa4c06254b2ebb0e48ba1f7a916c"
      },
      "spec" : {
        "resource" : {
          "flavor_id" : "$ref/parameters/train_spec",
          "node_count" : 1,
          "policy" : "regular"
        }
      }
    })

    depend_steps = ["condition_step_test"]
  }

  steps {
    name  = "training_job2"
    title = "labeling"
    type  = "job"
    inputs {
      name = "data_url"
      type = "obs"
      data = "$ref/data/f78e46676a454ccdacb9907f589f8d67"
    }
    outputs {
      name = "train_url"
      type = "obs"
      config = {
        obs_url = "/test-lh/test-metrics/"
      }
    }
    outputs {
      name = "service-link"
      type = "service_content"
      config = {
        config_file = "$ref/parameters/service_config2"
      }
    }

    properties = jsonencode({
      "algorithm" : {
        "id" : "21ef85a8-5e40-4618-95ee-aa48ec224b43",
        "parameters" : []
      },
      "kind" : "job",
      "metadata" : {
        "name" : "workflow-4a4317eb49ad4370bd087e6b726d84cf"
      },
      "spec" : {
        "resource" : {
          "flavor_id" : "$ref/parameters/train_spec",
          "node_count" : 1,
          "policy" : "regular"
        }
      }
    })

    depend_steps = ["condition_step_test"]
  }

  steps {
    name  = "training_job3"
    title = "labeling"
    type  = "job"
    inputs {
      name = "data_url"
      type = "obs"
      data = "$ref/data/f78e46676a454ccdacb9907f589f8d67"
    }
    outputs {
      name = "train_url"
      type = "obs"
      config = {
        obs_url = "/test-lh/test-metrics/"
      }
    }
    outputs {
      name = "service-link"
      type = "service_content"
      config = {
        config_file = "$ref/parameters/service_config3"
      }
    }

    properties = jsonencode({
      "algorithm" : {
        "id" : "21ef85a8-5e40-4618-95ee-aa48ec224b43",
        "parameters" : []
      },
      "kind" : "job",
      "metadata" : {
        "name" : "workflow-4a4317eb49ad4370bd087e6b726d84cf"
      },
      "spec" : {
        "resource" : {
          "flavor_id" : "$ref/parameters/train_spec",
          "node_count" : 1,
          "policy" : "regular"
        }
      }
    })

    depend_steps = ["condition_step_test"]
  }

  labels = ["subgraph"]

  data {
    name = "a2ff296da618452daa8243399f06db8e"
    type = "obs"
    value = {
      obs_url = "/test-lh/test-metrics/"
    }
    used_steps = ["training_job1"]
  }

  data {
    name = "f78e46676a454ccdacb9907f589f8d67"
    type = "obs"
    value = {
      obs_url = "/test-lh/test-metrics/"
    }
    used_steps = ["training_job2"]
  }

  data {
    name = "dee65054c96b4bf3b7ac98c0709f9ae0"
    type = "obs"
    value = {
      obs_url = "/test-lh/test-metrics/"
    }
    used_steps = ["training_job3"]
  }

  parameters {
    name       = "is_true"
    type       = "bool"
    delay      = true
    value      = true
    used_steps = ["condition_step_test"]
  }

  parameters {
    name        = "train_spec"
    type        = "str"
    format      = "flavor"
    description = "training specification"
    default     = "modelarts.vm.cpu.8u"
    used_steps  = ["training_job1", "training_job2", "training_job3"]
  }

  parameters {
    name       = "service_config1"
    type       = "str"
    default    = "/test-lh/test-metrics/metrics.json"
    used_steps = ["training_job1"]
  }

  parameters {
    name       = "service_config2"
    type       = "str"
    default    = "/test-lh/test-metrics/metrics.json"
    used_steps = ["training_job2"]
  }

  parameters {
    name       = "service_config3"
    type       = "str"
    default    = "/test-lh/test-metrics/metrics.json"
    used_steps = ["training_job3"]
  }
}
`, name)
}

func testModelartsWorkflow_basic_update(name string) string {
	return fmt.Sprintf(`
resource "huaweicloud_modelarts_workflow" "test" {
  name         = "%s"
  description  = "This is a demo update"
  workspace_id = "0"

  steps {
    name  = "condition_step_test"
    title = "condition_step_test"
    type  = "condition"
    conditions {
      type  = "=="
      left  = "$ref/parameters/is_true"
      right = true
    }
    if_then_steps   = ["training_job1"]
    else_then_steps = ["training_job2"]
  }

  steps {
    name  = "training_job1"
    title = "labeling"
    type  = "job"
    inputs {
      name = "data_url"
      type = "obs"
      data = "$ref/data/a2ff296da618452daa8243399f06db8e"
    }
    outputs {
      name = "train_url"
      type = "obs"
      config = {
        obs_url = "/test-lh/test-metrics/"
      }
    }
    outputs {
      name = "service-link"
      type = "service_content"
      config = {
        config_file = "$ref/parameters/service_config1"
      }
    }

    properties = jsonencode({
      "algorithm" : {
        "id" : "21ef85a8-5e40-4618-95ee-aa48ec224b43",
        "parameters" : []
      },
      "kind" : "job",
      "metadata" : {
        "name" : "workflow-b0b9fa4c06254b2ebb0e48ba1f7a916c"
      },
      "spec" : {
        "resource" : {
          "flavor_id" : "$ref/parameters/train_spec",
          "node_count" : 1,
          "policy" : "regular"
        }
      }
    })

    depend_steps = ["condition_step_test"]
  }

  steps {
    name  = "training_job2"
    title = "labeling"
    type  = "job"
    inputs {
      name = "data_url"
      type = "obs"
      data = "$ref/data/f78e46676a454ccdacb9907f589f8d67"
    }
    outputs {
      name = "train_url"
      type = "obs"
      config = {
        obs_url = "/test-lh/test-metrics/"
      }
    }
    outputs {
      name = "service-link"
      type = "service_content"
      config = {
        config_file = "$ref/parameters/service_config2"
      }
    }

    properties = jsonencode({
      "algorithm" : {
        "id" : "21ef85a8-5e40-4618-95ee-aa48ec224b43",
        "parameters" : []
      },
      "kind" : "job",
      "metadata" : {
        "name" : "workflow-4a4317eb49ad4370bd087e6b726d84cf"
      },
      "spec" : {
        "resource" : {
          "flavor_id" : "$ref/parameters/train_spec",
          "node_count" : 1,
          "policy" : "regular"
        }
      }
    })

    depend_steps = ["condition_step_test"]
  }

  steps {
    name  = "training_job3"
    title = "labeling"
    type  = "job"
    inputs {
      name = "data_url"
      type = "obs"
      data = "$ref/data/f78e46676a454ccdacb9907f589f8d67"
    }
    outputs {
      name = "train_url"
      type = "obs"
      config = {
        obs_url = "/test-lh/test-metrics/"
      }
    }
    outputs {
      name = "service-link"
      type = "service_content"
      config = {
        config_file = "$ref/parameters/service_config3"
      }
    }

    properties = jsonencode({
      "algorithm" : {
        "id" : "21ef85a8-5e40-4618-95ee-aa48ec224b43",
        "parameters" : []
      },
      "kind" : "job",
      "metadata" : {
        "name" : "workflow-4a4317eb49ad4370bd087e6b726d84cf"
      },
      "spec" : {
        "resource" : {
          "flavor_id" : "$ref/parameters/train_spec",
          "node_count" : 1,
          "policy" : "regular"
        }
      }
    })

    depend_steps = ["condition_step_test"]
  }

  labels = ["subgraph", "test"]

  data {
    name = "a2ff296da618452daa8243399f06db8e"
    type = "obs"
    value = {
      obs_url = "/test-lh/test-metrics/"
    }
    used_steps = ["training_job1"]
  }

  data {
    name = "f78e46676a454ccdacb9907f589f8d67"
    type = "obs"
    value = {
      obs_url = "/test-lh/test-metrics/"
    }
    used_steps = ["training_job2"]
  }

  data {
    name = "dee65054c96b4bf3b7ac98c0709f9ae0"
    type = "obs"
    value = {
      obs_url = "/test-lh/test-metrics/"
    }
    used_steps = ["training_job3"]
  }

  parameters {
    name       = "is_true"
    type       = "bool"
    delay      = true
    value      = true
    used_steps = ["condition_step_test"]
  }

  parameters {
    name        = "train_spec"
    type        = "str"
    format      = "flavor"
    description = "training specification"
    default     = "modelarts.vm.cpu.8u"
    used_steps  = ["training_job1", "training_job2", "training_job3"]
  }

  parameters {
    name       = "service_config1"
    type       = "str"
    default    = "/test-lh/test-metrics/metrics.json"
    used_steps = ["training_job1"]
  }

  parameters {
    name       = "service_config2"
    type       = "str"
    default    = "/test-lh/test-metrics/metrics.json"
    used_steps = ["training_job2"]
  }

  parameters {
    name       = "service_config3"
    type       = "str"
    default    = "/test-lh/test-metrics/metrics.json"
    used_steps = ["training_job3"]
  }

  parameters {
    name       = "service_config4"
    type       = "str"
    default    = "/test-lh/test-metrics/metrics.json"
  }
}
`, name)
}
