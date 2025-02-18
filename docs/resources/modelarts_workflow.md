---
subcategory: "AI Development Platform (ModelArts)"
---

# huaweicloud_modelarts_workflow

Manages a Modelarts workflow resource within HuaweiCloud.

## Example Usage

```hcl
resource "huaweicloud_modelarts_workflow" "test" {
  name         = "graph-test-has-condition-step"
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
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the resource.
  If omitted, the provider-level region will be used. Changing this parameter will create a new resource.

* `name` - (Required, String) Name of the workflow, which consists of 1 to 64 characters.  
  Only letters, digits, hyphens (-), and underscores (_) are allowed.

* `description` - (Optional, String) The description of the workflow.

* `steps` - (Optional, List) List of workflow steps.
  The [steps](#ModelartsWorkflow_WorkflowStep) structure is documented below.

* `workspace_id` - (Optional, String, ForceNew) Workspace ID, which defaults to 0.

  Changing this parameter will create a new resource.

* `data_requirements` - (Optional, List) List of data requirements.
  The [data_requirements](#ModelartsWorkflow_DataRequirement) structure is documented below.

* `data` - (Optional, List, ForceNew) List of data included in workflow.

  Changing this parameter will create a new resource.
  The [data](#ModelartsWorkflow_Data) structure is documented below.

* `parameters` - (Optional, List) List of workflow parameters.
  The [parameters](#ModelartsWorkflow_WorkflowParameter) structure is documented below.

* `source_workflow_id` - (Optional, String, ForceNew) Workflow ID to be copied.  
  When creating a workflow by copying, this parameter is mandatory.

  Changing this parameter will create a new resource.

* `gallery_subscription` - (Optional, List, ForceNew) When creating a workflow by subscribing to the gallery,
  this parameter is mandatory.

  Changing this parameter will create a new resource.
  The [gallery_subscription](#ModelartsWorkflow_WorkflowGallerySubscription) structure is documented below.

* `source` - (Optional, String, ForceNew) Workflow source.  
  Value options are as follows:
    + **ai_gallery**: Imported from AI Gallery.

  Changing this parameter will create a new resource.

* `storages` - (Optional, List) List of workflow storage.
  The [storages](#ModelartsWorkflow_WorkflowStorage) structure is documented below.

* `labels` - (Optional, List) List of workflow labels.

* `assets` - (Optional, List, ForceNew) List of workflow assets.

  Changing this parameter will create a new resource.
  The [assets](#ModelartsWorkflow_WorkflowAsset) structure is documented below.

* `sub_graphs` - (Optional, List, ForceNew) List of workflow subgraphs.

  Changing this parameter will create a new resource.
  The [sub_graphs](#ModelartsWorkflow_WorkflowSubgraph) structure is documented below.

* `policy` - (Optional, List, ForceNew) Workflow execution policy.

  Changing this parameter will create a new resource.
  The [policy](#ModelartsWorkflow_WorkflowPolicy) structure is documented below.

* `smn_switch` - (Optional, Bool) Whether the SMN message subscription is enabled.  
  Value options are as follows:
    + **true**: SMN message subscription is enabled.
    + **false**: SMN message subscription is disabled.

* `subscription_id` - (Optional, String, ForceNew) SMN message subscription ID.

  Changing this parameter will create a new resource.

* `exeml_template_id` - (Optional, String, ForceNew) Auto-learning template ID.

  Changing this parameter will create a new resource.

<a name="ModelartsWorkflow_WorkflowStep"></a>
The `steps` block supports:

* `name` - (Optional, String) Name of the workflow step.

* `type` - (Optional, String) Type of the workflow step.  
  Value options are as follows:
    + **job**: Training
    + **labeling**: Labeling
    + **release_dataset**: Dataset release
    + **model**: Model release
    + **service**: Service deployment
    + **mrs_job**: MRS job
    + **dataset_import**: Dataset import
    + **create_dataset**: Dataset creation

* `inputs` - (Optional, List) List of workflow step input items.
  The [inputs](#ModelartsWorkflow_JobInput) structure is documented below.

* `outputs` - (Optional, List) List of workflow step output items.
  The [outputs](#ModelartsWorkflow_JobOutput) structure is documented below.

* `title` - (Optional, String) Title of the workflow step.

* `description` - (Optional, String) Description of the workflow step.

* `properties` - (Optional, String) Properties of the workflow step.  
  The properties are described in JSON format.

* `depend_steps` - (Optional, List) List of dependent workflow steps.

* `conditions` - (Optional, List) List of workflow step execution conditions.
  The [conditions](#ModelartsWorkflow_StepCondition) structure is documented below.

* `if_then_steps` - (Optional, List) List of branch steps that meet the conditions.

* `else_then_steps` - (Optional, List) List of branch steps that do not meet the conditions.

* `policy` - (Optional, List) Workflow step execution policy.
  The [policy](#ModelartsWorkflow_WorkflowStepPolicy) structure is documented below.

<a name="ModelartsWorkflow_JobInput"></a>
The `inputs` block supports:

* `name` - (Optional, String) Name of the input item.

* `type` - (Optional, String) Type of the input item.  
  Value options are as follows:
    + **dataset**: Dataset
    + **obs**: OBS
    + **data_selector**: Data selection

* `data` - (Optional, String) Data of the Input item.

* `value` - (Optional, String) Value of the input item.

<a name="ModelartsWorkflow_JobOutput"></a>
The `outputs` block supports:

* `name` - (Optional, String) Name of the output item.

* `type` - (Optional, String) Type of the output item.  
  Value options are as follows:
    + **obs**: OBS
    + **model**: AI application metamodel

* `config` - (Optional, Map) The configuration of the output item.

<a name="ModelartsWorkflow_StepCondition"></a>
The `conditions` block supports:

* `type` - (Optional, String) Type of the condition.  
  Value options are as follows:
    + **==**: Equal to
    + **!=**: Not equal to
    + **>**: Greater than
    + **>=**: Greater than or equal to
    + **<**: Less than
    + **<=**: Less than or equal to
    + **in**: In
    + **or**: Or

* `left` - (Optional, String) Branch when the condition is `true`.

* `right` - (Optional, String) Branch when the condition is `false`.

<a name="ModelartsWorkflow_WorkflowStepPolicy"></a>
The `policy` block supports:

* `poll_interval_seconds` - (Optional, String) Execution interval.

* `max_execution_minutes` - (Optional, String) Maximum execution time.

<a name="ModelartsWorkflow_DataRequirement"></a>
The `data_requirements` block supports:

* `name` - (Optional, String) Name of the data.

* `type` - (Optional, String) Type of the data.  
  Value options are as follows:
    + **dataset**: Dataset
    + **obs**: OBS
    + **swr**: SWR
    + **model_list**: AI application list
    + **label_task**: Labeling task
    + **service**: Online service

* `conditions` - (Optional, List) Data constraint conditions.
  The [conditions](#ModelartsWorkflow_Constraint) structure is documented below.

* `value` - (Optional, Map) Data value.

* `used_steps` - (Optional, List) Workflow steps that use the data.

* `delay` - (Optional, Bool) Delay parameter flag.

<a name="ModelartsWorkflow_Constraint"></a>
The `conditions` block supports:

* `attribute` - (Optional, String) Attribute.

* `operator` - (Optional, String) Operation.  
  Currently, only the **equal** operation is supported.

* `value` - (Optional, String) Value.

<a name="ModelartsWorkflow_Data"></a>
The `data` block supports:

* `name` - (Optional, String) Name of the data.

* `type` - (Optional, String) Type of the data.  
  Value options are as follows:
    + **dataset**: Dataset
    + **obs**: OBS
    + **swr**: SWR
    + **model**: AI application
    + **label_task**: Labeling task
    + **service**: Online service
    + **image**: Image

* `value` - (Optional, Map) Data value.

* `used_steps` - (Optional, List) Workflow steps that use the data.

<a name="ModelartsWorkflow_WorkflowParameter"></a>
The `parameters` block supports:

* `name` - (Optional, String) Name of the parameter.

* `type` - (Optional, String) Type of the parameter.  
  Value options are as follows:
    + **str**: String
    + **int**: Integer
    + **bool**: Boolean
    + **float**: Floating point number

* `description` - (Optional, String) Description of the parameter.

* `example` - (Optional, String) Example of the parameter.

* `delay` - (Optional, Bool) Whether it is a delayed input parameters.

* `default` - (Optional, String) Default value of the parameter.

* `value` - (Optional, String) Value of the parameter.

* `enum` - (Optional, List) Enumeration value of the parameters.

* `used_steps` - (Optional, List) Workflow steps that use the parameter.

* `format` - (Optional, String) The format of the parameter.

* `constraint` - (Optional, Map) Parameter constraint conditions.

<a name="ModelartsWorkflow_WorkflowGallerySubscription"></a>
The `gallery_subscription` block supports:

* `content_id` - (Optional, String) ID of the content to be subscribed.

* `version` - (Optional, String) Version of the content to be subscribed.

* `expired_at` - (Optional, String) Subscription expiration time.

<a name="ModelartsWorkflow_WorkflowStorage"></a>
The `storages` block supports:

* `name` - (Optional, String) Name of the storage.

* `type` - (Optional, String) Type of the storage.  
  Value options are as follows:
    + **obs**: OBS

* `path` - (Optional, String) Storage path.

<a name="ModelartsWorkflow_WorkflowAsset"></a>
The `assets` block supports:

* `name` - (Optional, String) Name of the asset.

* `type` - (Optional, String) Type of the asset.  
  Value options are as follows:
    + **algorithm**: Algorithm
    + **algorithm2**: New algorithm
    + **model**: Model

* `content_id` - (Optional, String) ID of the asset to be subscribed.

* `subscription_id` - (Optional, String) ID of the subscription.

<a name="ModelartsWorkflow_WorkflowSubgraph"></a>
The `sub_graphs` block supports:

* `name` - (Optional, String) Name of the subgraph.

* `steps` - (Optional, List) List of subgraph steps.

<a name="ModelartsWorkflow_WorkflowPolicy"></a>
The `policy` block supports:

* `use_scene` - (Optional, String) Use scene.

* `scene_id` - (Optional, String) Scene ID.

* `scenes` - (Optional, List) List of scenes.
  The [scenes](#ModelartsWorkflow_Scene) structure is documented below.

<a name="ModelartsWorkflow_Scene"></a>
The `scenes` block supports:

* `id` - (Optional, String) Scene ID.

* `name` - (Optional, String) Scene name.

* `steps` - (Optional, List) List of steps.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.

* `run_count` - Number of times the workflow has been run.  
  The default value is 0.

* `param_ready` - Whether all the required parameters of the workflow have been filled in.  
  The default value is `false`.

## Import

The modelarts workflow can be imported using the `id`, e.g.

```bash
$ terraform import huaweicloud_modelarts_workflow.test 0ce123456a00f2591fabc00385ff1234
```
