// ---------------------------------------------------------------
// *** AUTO GENERATED CODE ***
// @Product ModelArts
// ---------------------------------------------------------------

package modelarts

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jmespath/go-jmespath"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API ModelArts POST /v2/{project_id}/workflows
// @API ModelArts DELETE /v2/{project_id}/workflows/{id}
// @API ModelArts GET /v2/{project_id}/workflows/{id}
// @API ModelArts PUT /v2/{project_id}/workflows/{id}
func ResourceModelartsWorkflow() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceModelartsWorkflowCreate,
		UpdateContext: resourceModelartsWorkflowUpdate,
		ReadContext:   resourceModelartsWorkflowRead,
		DeleteContext: resourceModelartsWorkflowDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Name of the workflow, which consists of 1 to 64 characters.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The description of the workflow.`,
			},
			"steps": {
				Type:        schema.TypeList,
				Elem:        modelartsWorkflowWorkflowStepSchema(),
				Optional:    true,
				Computed:    true,
				Description: `List of workflow steps.`,
			},
			"workspace_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Workspace ID, which defaults to 0.`,
			},
			"data_requirements": {
				Type:        schema.TypeList,
				Elem:        modelartsWorkflowDataRequirementSchema(),
				Optional:    true,
				Computed:    true,
				Description: `List of data requirements.`,
			},
			"data": {
				Type:        schema.TypeList,
				Elem:        modelartsWorkflowDataSchema(),
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `List of data included in workflow.`,
			},
			"parameters": {
				Type:        schema.TypeList,
				Elem:        modelartsWorkflowWorkflowParameterSchema(),
				Optional:    true,
				Computed:    true,
				Description: `List of workflow parameters.`,
			},
			"source_workflow_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Workflow ID to be copied.`,
			},
			"gallery_subscription": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem:     modelartsWorkflowWorkflowGallerySubscriptionSchema(),
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"source": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Workflow source.`,
			},
			"storages": {
				Type:        schema.TypeList,
				Elem:        modelartsWorkflowWorkflowStorageSchema(),
				Optional:    true,
				Computed:    true,
				Description: `List of workflow storage.`,
			},
			"labels": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `List of workflow labels.`,
			},
			"assets": {
				Type:        schema.TypeList,
				Elem:        modelartsWorkflowWorkflowAssetSchema(),
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `List of workflow assets.`,
			},
			"sub_graphs": {
				Type:        schema.TypeList,
				Elem:        modelartsWorkflowWorkflowSubgraphSchema(),
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `List of workflow subgraphs.`,
			},
			"policy": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem:     modelartsWorkflowWorkflowPolicySchema(),
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"smn_switch": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: `Whether the SMN message subscription is enabled.`,
			},
			"subscription_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `SMN message subscription ID.`,
			},
			"exeml_template_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Auto-learning template ID.`,
			},
			"run_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Number of times the workflow has been run.`,
			},
			"param_ready": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `Whether all the required parameters of the workflow have been filled in.`,
			},
		},
	}
}

func modelartsWorkflowWorkflowStepSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Name of the workflow step.`,
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Type of the workflow step.`,
			},
			"inputs": {
				Type:        schema.TypeList,
				Elem:        modelartsWorkflowWorkflowStepJobInputSchema(),
				Optional:    true,
				Computed:    true,
				Description: `List of workflow step input items.`,
			},
			"outputs": {
				Type:        schema.TypeList,
				Elem:        modelartsWorkflowWorkflowStepJobOutputSchema(),
				Optional:    true,
				Computed:    true,
				Description: `List of workflow step output items.`,
			},
			"title": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Title of the workflow step.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Description of the workflow step.`,
			},
			"properties": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(_, old, new string, _ *schema.ResourceData) bool {
					return utils.JSONStringsEqual(old, new)
				},
				Description: `Properties of the workflow step.`,
			},
			"depend_steps": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `List of dependent workflow steps.`,
			},
			"conditions": {
				Type:        schema.TypeList,
				Elem:        modelartsWorkflowWorkflowStepStepConditionSchema(),
				Optional:    true,
				Computed:    true,
				Description: `List of workflow step execution conditions.`,
			},
			"if_then_steps": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `List of branch steps that meet the conditions.`,
			},
			"else_then_steps": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `List of branch steps that do not meet the conditions.`,
			},
			"policy": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem:     modelartsWorkflowWorkflowStepWorkflowStepPolicySchema(),
				Optional: true,
				Computed: true,
			},
		},
	}
	return &sc
}

func modelartsWorkflowWorkflowStepJobInputSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Name of the input item.`,
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Type of the input item.`,
			},
			"data": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Data of the Input item.`,
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Value of the input item.`,
			},
		},
	}
	return &sc
}

func modelartsWorkflowWorkflowStepJobOutputSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Name of the output item.`,
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Type of the output item.`,
			},
			"config": {
				Type:        schema.TypeMap,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `The configuration of the output item.`,
			},
		},
	}
	return &sc
}

func modelartsWorkflowWorkflowStepStepConditionSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Type of the condition.`,
			},
			"left": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Branch when the condition is true.`,
			},
			"right": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Branch when the condition is false.`,
			},
		},
	}
	return &sc
}

func modelartsWorkflowWorkflowStepWorkflowStepPolicySchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"poll_interval_seconds": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Execution interval.`,
			},
			"max_execution_minutes": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Maximum execution time.`,
			},
		},
	}
	return &sc
}

func modelartsWorkflowDataRequirementSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Name of the data.`,
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Type of the data.`,
			},
			"conditions": {
				Type:        schema.TypeList,
				Elem:        modelartsWorkflowDataRequirementConstraintSchema(),
				Optional:    true,
				Computed:    true,
				Description: `Data constraint conditions.`,
			},
			"value": {
				Type:        schema.TypeMap,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `Data value.`,
			},
			"used_steps": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `Workflow steps that use the data.`,
			},
			"delay": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: `Delay parameter flag.`,
			},
		},
	}
	return &sc
}

func modelartsWorkflowDataRequirementConstraintSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Attribute.`,
			},
			"operator": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Operation.`,
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Value.`,
			},
		},
	}
	return &sc
}

func modelartsWorkflowDataSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Name of the data.`,
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Type of the data.`,
			},
			"value": {
				Type:        schema.TypeMap,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `Data value.`,
			},
			"used_steps": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `Workflow steps that use the data.`,
			},
		},
	}
	return &sc
}

func modelartsWorkflowWorkflowParameterSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Name of the parameter.`,
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Type of the parameter.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Description of the parameter.`,
			},
			"example": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Example of the parameter.`,
			},
			"delay": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: `Whether it is a delayed input parameters.`,
			},
			"default": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Default value of the parameter.`,
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Value of the parameter.`,
			},
			"enum": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `Enumeration value of the parameters.`,
			},
			"used_steps": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `Workflow steps that use the parameter.`,
			},
			"format": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The format of the parameter.`,
			},
			"constraint": {
				Type:        schema.TypeMap,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `Parameter constraint conditions.`,
			},
		},
	}
	return &sc
}

func modelartsWorkflowWorkflowGallerySubscriptionSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"content_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `ID of the content to be subscribed.`,
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Version of the content to be subscribed.`,
			},
			"expired_at": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Subscription expiration time.`,
			},
		},
	}
	return &sc
}

func modelartsWorkflowWorkflowStorageSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Name of the storage.`,
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Type of the storage.`,
			},
			"path": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Storage path.`,
			},
		},
	}
	return &sc
}

func modelartsWorkflowWorkflowAssetSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Name of the asset.`,
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Type of the asset.`,
			},
			"content_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `ID of the asset to be subscribed.`,
			},
			"subscription_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `ID of the subscription.`,
			},
		},
	}
	return &sc
}

func modelartsWorkflowWorkflowSubgraphSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Name of the subgraph.`,
			},
			"steps": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `List of subgraph steps.`,
			},
		},
	}
	return &sc
}

func modelartsWorkflowWorkflowPolicySchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"use_scene": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Use scene.`,
			},
			"scene_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Scene ID.`,
			},
			"scenes": {
				Type:        schema.TypeList,
				Elem:        modelartsWorkflowWorkflowPolicySceneSchema(),
				Optional:    true,
				Computed:    true,
				Description: `List of scenes.`,
			},
		},
	}
	return &sc
}

func modelartsWorkflowWorkflowPolicySceneSchema() *schema.Resource {
	sc := schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Scene ID.`,
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Scene name.`,
			},
			"steps": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Computed:    true,
				Description: `List of steps.`,
			},
		},
	}
	return &sc
}

func resourceModelartsWorkflowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	// createWorkflow: create a Modelarts workflow.
	var (
		createWorkflowHttpUrl = "v2/{project_id}/workflows"
		createWorkflowProduct = "modelarts"
	)
	createWorkflowClient, err := cfg.NewServiceClient(createWorkflowProduct, region)
	if err != nil {
		return diag.Errorf("error creating ModelArts client: %s", err)
	}

	createWorkflowPath := createWorkflowClient.Endpoint + createWorkflowHttpUrl
	createWorkflowPath = strings.ReplaceAll(createWorkflowPath, "{project_id}", createWorkflowClient.ProjectID)

	createWorkflowOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			201,
		},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	}

	createWorkflowOpt.JSONBody = utils.RemoveNil(buildCreateWorkflowBodyParams(d))
	createWorkflowResp, err := createWorkflowClient.Request("POST", createWorkflowPath, &createWorkflowOpt)
	if err != nil {
		return diag.Errorf("error creating Modelarts workflow: %s", err)
	}

	createWorkflowRespBody, err := utils.FlattenResponse(createWorkflowResp)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := jmespath.Search("workflow_id", createWorkflowRespBody)
	if err != nil {
		return diag.Errorf("error creating Modelarts workflow: ID is not found in API response")
	}
	d.SetId(id.(string))

	return resourceModelartsWorkflowRead(ctx, d, meta)
}

func buildCreateWorkflowBodyParams(d *schema.ResourceData) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"name":                 d.Get("name"),
		"description":          utils.ValueIngoreEmpty(d.Get("description")),
		"steps":                buildWorkflowReqBodyWorkflowStep(d.Get("steps")),
		"workspace_id":         utils.ValueIngoreEmpty(d.Get("workspace_id")),
		"data_requirements":    buildWorkflowReqBodyDataRequirement(d.Get("data_requirements")),
		"data":                 buildCreateWorkflowReqBodyData(d.Get("data")),
		"parameters":           buildWorkflowReqBodyWorkflowParameter(d.Get("parameters")),
		"source_workflow_id":   utils.ValueIngoreEmpty(d.Get("source_workflow_id")),
		"gallery_subscription": buildCreateWorkflowReqBodyWorkflowGallerySubscription(d.Get("gallery_subscription")),
		"source":               utils.ValueIngoreEmpty(d.Get("source")),
		"storages":             buildWorkflowReqBodyWorkflowStorage(d.Get("storages")),
		"labels":               utils.ValueIngoreEmpty(d.Get("labels")),
		"assets":               buildCreateWorkflowReqBodyWorkflowAsset(d.Get("assets")),
		"sub_graphs":           buildCreateWorkflowReqBodyWorkflowSubgraph(d.Get("sub_graphs")),
		"policy":               buildCreateWorkflowReqBodyWorkflowPolicy(d.Get("policy")),
		"smn_switch":           utils.ValueIngoreEmpty(d.Get("smn_switch")),
		"subscription_id":      utils.ValueIngoreEmpty(d.Get("subscription_id")),
		"exeml_template_id":    utils.ValueIngoreEmpty(d.Get("exeml_template_id")),
	}
	return bodyParams
}

func buildWorkflowReqBodyWorkflowStep(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})

			properties := make(map[string]interface{})
			if v, ok := raw["properties"].(string); ok && len(v) > 0 {
				err := json.Unmarshal([]byte(v), &properties)
				if err != nil {
					log.Printf("[ERROR] Invalid type of the properties, it is not JSON format, %s, %s", err, v)
					return nil
				}
			}

			rst[i] = map[string]interface{}{
				"name":            utils.ValueIngoreEmpty(raw["name"]),
				"type":            utils.ValueIngoreEmpty(raw["type"]),
				"inputs":          buildWorkflowStepJobInput(raw["inputs"]),
				"outputs":         buildWorkflowStepJobOutput(raw["outputs"]),
				"title":           utils.ValueIngoreEmpty(raw["title"]),
				"description":     utils.ValueIngoreEmpty(raw["description"]),
				"properties":      properties,
				"depend_steps":    utils.ValueIngoreEmpty(raw["depend_steps"]),
				"conditions":      buildWorkflowStepStepCondition(raw["conditions"]),
				"if_then_steps":   utils.ValueIngoreEmpty(raw["if_then_steps"]),
				"else_then_steps": utils.ValueIngoreEmpty(raw["else_then_steps"]),
				"policy":          buildWorkflowStepWorkflowStepPolicy(raw["policy"]),
			}
		}
		return rst
	}
	return nil
}

func buildWorkflowStepJobInput(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})
			rst[i] = map[string]interface{}{
				"name":  utils.ValueIngoreEmpty(raw["name"]),
				"type":  utils.ValueIngoreEmpty(raw["type"]),
				"data":  utils.ValueIngoreEmpty(raw["data"]),
				"value": utils.ValueIngoreEmpty(raw["value"]),
			}
		}
		return rst
	}
	return nil
}

func buildWorkflowStepJobOutput(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})
			rst[i] = map[string]interface{}{
				"name":   utils.ValueIngoreEmpty(raw["name"]),
				"type":   utils.ValueIngoreEmpty(raw["type"]),
				"config": utils.ValueIngoreEmpty(raw["config"]),
			}
		}
		return rst
	}
	return nil
}

func buildWorkflowStepStepCondition(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})
			rst[i] = map[string]interface{}{
				"type":  utils.ValueIngoreEmpty(raw["type"]),
				"left":  utils.ValueIngoreEmpty(raw["left"]),
				"right": utils.ValueIngoreEmpty(raw["right"]),
			}
		}
		return rst
	}
	return nil
}

func buildWorkflowStepWorkflowStepPolicy(rawParams interface{}) map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}
		raw, ok := rawArray[0].(map[string]interface{})
		if !ok {
			return nil
		}

		params := map[string]interface{}{
			"poll_interval_seconds": utils.ValueIngoreEmpty(raw["poll_interval_seconds"]),
			"max_execution_minutes": utils.ValueIngoreEmpty(raw["max_execution_minutes"]),
		}
		return params
	}
	return nil
}

func buildWorkflowReqBodyDataRequirement(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})
			rst[i] = map[string]interface{}{
				"name":       utils.ValueIngoreEmpty(raw["name"]),
				"type":       utils.ValueIngoreEmpty(raw["type"]),
				"conditions": buildDataRequirementConstraint(raw["conditions"]),
				"value":      utils.ValueIngoreEmpty(raw["value"]),
				"used_steps": utils.ValueIngoreEmpty(raw["used_steps"]),
				"delay":      utils.ValueIngoreEmpty(raw["delay"]),
			}
		}
		return rst
	}
	return nil
}

func buildDataRequirementConstraint(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})
			rst[i] = map[string]interface{}{
				"attribute": utils.ValueIngoreEmpty(raw["attribute"]),
				"operator":  utils.ValueIngoreEmpty(raw["operator"]),
				"value":     utils.ValueIngoreEmpty(raw["value"]),
			}
		}
		return rst
	}
	return nil
}

func buildCreateWorkflowReqBodyData(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})
			rst[i] = map[string]interface{}{
				"name":       utils.ValueIngoreEmpty(raw["name"]),
				"type":       utils.ValueIngoreEmpty(raw["type"]),
				"value":      utils.ValueIngoreEmpty(raw["value"]),
				"used_steps": utils.ValueIngoreEmpty(raw["used_steps"]),
			}
		}
		return rst
	}
	return nil
}

func buildWorkflowReqBodyWorkflowParameter(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})
			rst[i] = map[string]interface{}{
				"name":        utils.ValueIngoreEmpty(raw["name"]),
				"type":        utils.ValueIngoreEmpty(raw["type"]),
				"description": utils.ValueIngoreEmpty(raw["description"]),
				"example":     utils.ValueIngoreEmpty(raw["example"]),
				"delay":       utils.ValueIngoreEmpty(raw["delay"]),
				"default":     utils.ValueIngoreEmpty(raw["default"]),
				"value":       utils.ValueIngoreEmpty(raw["value"]),
				"enum":        utils.ValueIngoreEmpty(raw["enum"]),
				"used_steps":  utils.ValueIngoreEmpty(raw["used_steps"]),
				"format":      utils.ValueIngoreEmpty(raw["format"]),
				"constraint":  utils.ValueIngoreEmpty(raw["constraint"]),
			}

			// convert the value to the correct type
			var s interface{}
			var err error
			switch typeValue := raw["type"].(string); typeValue {
			case "int":
				s, err = strconv.Atoi(raw["value"].(string))
			case "bool":
				s, err = strconv.ParseBool(raw["value"].(string))
			case "float":
				s, err = strconv.ParseFloat(raw["value"].(string), 64)
			default:
				s = raw["value"].(string)
				err = nil
			}

			if err != nil {
				log.Printf("[ERROR] the type of the parameter is %s, but value is %s", raw["type"], raw["value"])
			}

			rst[i]["value"] = s
		}
		return rst
	}
	return nil
}

func buildCreateWorkflowReqBodyWorkflowGallerySubscription(rawParams interface{}) map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}
		raw, ok := rawArray[0].(map[string]interface{})
		if !ok {
			return nil
		}

		params := map[string]interface{}{
			"content_id": utils.ValueIngoreEmpty(raw["content_id"]),
			"version":    utils.ValueIngoreEmpty(raw["version"]),
			"expired_at": utils.ValueIngoreEmpty(raw["expired_at"]),
		}
		return params
	}
	return nil
}

func buildWorkflowReqBodyWorkflowStorage(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})
			rst[i] = map[string]interface{}{
				"name": utils.ValueIngoreEmpty(raw["name"]),
				"type": utils.ValueIngoreEmpty(raw["type"]),
				"path": utils.ValueIngoreEmpty(raw["path"]),
			}
		}
		return rst
	}
	return nil
}

func buildCreateWorkflowReqBodyWorkflowAsset(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})
			rst[i] = map[string]interface{}{
				"name":            utils.ValueIngoreEmpty(raw["name"]),
				"type":            utils.ValueIngoreEmpty(raw["type"]),
				"content_id":      utils.ValueIngoreEmpty(raw["content_id"]),
				"subscription_id": utils.ValueIngoreEmpty(raw["subscription_id"]),
			}
		}
		return rst
	}
	return nil
}

func buildCreateWorkflowReqBodyWorkflowSubgraph(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})
			rst[i] = map[string]interface{}{
				"name":  utils.ValueIngoreEmpty(raw["name"]),
				"steps": utils.ValueIngoreEmpty(raw["steps"]),
			}
		}
		return rst
	}
	return nil
}

func buildCreateWorkflowReqBodyWorkflowPolicy(rawParams interface{}) map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}
		raw, ok := rawArray[0].(map[string]interface{})
		if !ok {
			return nil
		}

		params := map[string]interface{}{
			"use_scene": utils.ValueIngoreEmpty(raw["use_scene"]),
			"scene_id":  utils.ValueIngoreEmpty(raw["scene_id"]),
			"scenes":    buildWorkflowPolicyScene(raw["scenes"]),
		}
		return params
	}
	return nil
}

func buildWorkflowPolicyScene(rawParams interface{}) []map[string]interface{} {
	if rawArray, ok := rawParams.([]interface{}); ok {
		if len(rawArray) == 0 {
			return nil
		}

		rst := make([]map[string]interface{}, len(rawArray))
		for i, v := range rawArray {
			raw := v.(map[string]interface{})
			rst[i] = map[string]interface{}{
				"id":    utils.ValueIngoreEmpty(raw["id"]),
				"name":  utils.ValueIngoreEmpty(raw["name"]),
				"steps": utils.ValueIngoreEmpty(raw["steps"]),
			}
		}
		return rst
	}
	return nil
}

func resourceModelartsWorkflowRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	var mErr *multierror.Error

	// getWorkflow: Query the Modelarts workflow.
	var (
		getWorkflowHttpUrl = "v2/{project_id}/workflows/{id}"
		getWorkflowProduct = "modelarts"
	)
	getWorkflowClient, err := cfg.NewServiceClient(getWorkflowProduct, region)
	if err != nil {
		return diag.Errorf("error creating ModelArts client: %s", err)
	}

	getWorkflowPath := getWorkflowClient.Endpoint + getWorkflowHttpUrl
	getWorkflowPath = strings.ReplaceAll(getWorkflowPath, "{project_id}", getWorkflowClient.ProjectID)
	getWorkflowPath = strings.ReplaceAll(getWorkflowPath, "{id}", d.Id())

	getWorkflowOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			200,
		},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	}

	getWorkflowResp, err := getWorkflowClient.Request("GET", getWorkflowPath, &getWorkflowOpt)

	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving Modelarts workflow")
	}

	getWorkflowRespBody, err := utils.FlattenResponse(getWorkflowResp)
	if err != nil {
		return diag.FromErr(err)
	}

	mErr = multierror.Append(
		mErr,
		d.Set("region", region),
		d.Set("name", utils.PathSearch("name", getWorkflowRespBody, nil)),
		d.Set("description", utils.PathSearch("description", getWorkflowRespBody, nil)),
		d.Set("steps", flattenGetWorkflowResponseBodyWorkflowStep(getWorkflowRespBody)),
		d.Set("workspace_id", utils.PathSearch("workspace_id", getWorkflowRespBody, nil)),
		d.Set("data_requirements", flattenGetWorkflowResponseBodyDataRequirement(getWorkflowRespBody)),
		d.Set("data", flattenGetWorkflowResponseBodyData(getWorkflowRespBody)),
		d.Set("parameters", flattenGetWorkflowResponseBodyWorkflowParameter(getWorkflowRespBody)),
		d.Set("source_workflow_id", utils.PathSearch("source_workflow_id", getWorkflowRespBody, nil)),
		d.Set("gallery_subscription", flattenGetWorkflowResponseBodyWorkflowGallerySubscription(getWorkflowRespBody)),
		d.Set("run_count", utils.PathSearch("run_count", getWorkflowRespBody, nil)),
		d.Set("param_ready", utils.PathSearch("param_ready", getWorkflowRespBody, nil)),
		d.Set("source", utils.PathSearch("source", getWorkflowRespBody, nil)),
		d.Set("storages", flattenGetWorkflowResponseBodyWorkflowStorage(getWorkflowRespBody)),
		d.Set("labels", utils.PathSearch("labels", getWorkflowRespBody, nil)),
		d.Set("assets", flattenGetWorkflowResponseBodyWorkflowAsset(getWorkflowRespBody)),
		d.Set("sub_graphs", flattenGetWorkflowResponseBodyWorkflowSubgraph(getWorkflowRespBody)),
		d.Set("policy", flattenGetWorkflowResponseBodyWorkflowPolicy(getWorkflowRespBody)),
		d.Set("smn_switch", utils.PathSearch("smn_switch", getWorkflowRespBody, nil)),
		d.Set("subscription_id", utils.PathSearch("subscription_id", getWorkflowRespBody, nil)),
		d.Set("exeml_template_id", utils.PathSearch("exeml_template_id", getWorkflowRespBody, nil)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func flattenGetWorkflowResponseBodyWorkflowStep(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("steps", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		properties, err := json.Marshal(utils.PathSearch("properties", v, nil))
		if err != nil {
			log.Printf("[ERROR] error parsing properties from response= %#v", v)
		}

		rst = append(rst, map[string]interface{}{
			"name":            utils.PathSearch("name", v, nil),
			"type":            utils.PathSearch("type", v, nil),
			"inputs":          flattenWorkflowStepJobInput(v),
			"outputs":         flattenWorkflowStepJobOutput(v),
			"title":           utils.PathSearch("title", v, nil),
			"description":     utils.PathSearch("description", v, nil),
			"properties":      string(properties),
			"depend_steps":    utils.PathSearch("depend_steps", v, nil),
			"conditions":      flattenWorkflowStepStepCondition(v),
			"if_then_steps":   utils.PathSearch("if_then_steps", v, nil),
			"else_then_steps": utils.PathSearch("else_then_steps", v, nil),
			"policy":          flattenWorkflowStepWorkflowStepPolicy(v),
		})
	}
	return rst
}

func flattenWorkflowStepJobInput(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("inputs", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"name":  utils.PathSearch("name", v, nil),
			"type":  utils.PathSearch("type", v, nil),
			"data":  utils.PathSearch("data", v, nil),
			"value": utils.PathSearch("value", v, nil),
		})
	}
	return rst
}

func flattenWorkflowStepJobOutput(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("outputs", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"name":   utils.PathSearch("name", v, nil),
			"type":   utils.PathSearch("type", v, nil),
			"config": utils.PathSearch("config", v, nil),
		})
	}
	return rst
}

func flattenWorkflowStepStepCondition(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("conditions", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"type":  utils.PathSearch("type", v, nil),
			"left":  utils.PathSearch("left", v, nil),
			"right": utils.PathSearch("right", v, nil),
		})
	}
	return rst
}

func flattenWorkflowStepWorkflowStepPolicy(resp interface{}) []interface{} {
	var rst []interface{}
	curJson, err := jmespath.Search("policy", resp)
	if err != nil {
		log.Printf("[ERROR] error parsing policy from response= %#v", resp)
		return rst
	}

	rst = []interface{}{
		map[string]interface{}{
			"poll_interval_seconds": utils.PathSearch("poll_interval_seconds", curJson, nil),
			"max_execution_minutes": utils.PathSearch("max_execution_minutes", curJson, nil),
		},
	}
	return rst
}

func flattenGetWorkflowResponseBodyDataRequirement(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("data_requirements", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"name":       utils.PathSearch("name", v, nil),
			"type":       utils.PathSearch("type", v, nil),
			"conditions": flattenDataRequirementConstraint(v),
			"value":      utils.PathSearch("value", v, nil),
			"used_steps": utils.PathSearch("used_steps", v, nil),
			"delay":      utils.PathSearch("delay", v, nil),
		})
	}
	return rst
}

func flattenDataRequirementConstraint(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("conditions", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"attribute": utils.PathSearch("attribute", v, nil),
			"operator":  utils.PathSearch("operator", v, nil),
			"value":     utils.PathSearch("value", v, nil),
		})
	}
	return rst
}

func flattenGetWorkflowResponseBodyData(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("data", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"name":       utils.PathSearch("name", v, nil),
			"type":       utils.PathSearch("type", v, nil),
			"value":      utils.PathSearch("value", v, nil),
			"used_steps": utils.PathSearch("used_steps", v, nil),
		})
	}
	return rst
}

func flattenGetWorkflowResponseBodyWorkflowParameter(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("parameters", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"name":        utils.PathSearch("name", v, nil),
			"type":        utils.PathSearch("type", v, nil),
			"description": utils.PathSearch("description", v, nil),
			"example":     utils.PathSearch("example", v, nil),
			"delay":       utils.PathSearch("delay", v, nil),
			"default":     utils.PathSearch("default", v, nil),
			"value":       fmt.Sprint(utils.PathSearch("value", v, "")),
			"enum":        utils.PathSearch("enum", v, nil),
			"used_steps":  utils.PathSearch("used_steps", v, nil),
			"format":      utils.PathSearch("format", v, nil),
			"constraint":  utils.PathSearch("constraint", v, nil),
		})
	}
	return rst
}

func flattenGetWorkflowResponseBodyWorkflowGallerySubscription(resp interface{}) []interface{} {
	var rst []interface{}
	curJson, err := jmespath.Search("gallery_subscription", resp)
	if err != nil {
		log.Printf("[ERROR] error parsing gallery_subscription from response= %#v", resp)
		return rst
	}

	rst = []interface{}{
		map[string]interface{}{
			"content_id": utils.PathSearch("content_id", curJson, nil),
			"version":    utils.PathSearch("version", curJson, nil),
			"expired_at": utils.PathSearch("expired_at", curJson, nil),
		},
	}
	return rst
}

func flattenGetWorkflowResponseBodyWorkflowStorage(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("storages", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"name": utils.PathSearch("name", v, nil),
			"type": utils.PathSearch("type", v, nil),
			"path": utils.PathSearch("path", v, nil),
		})
	}
	return rst
}

func flattenGetWorkflowResponseBodyWorkflowAsset(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("assets", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"name":            utils.PathSearch("name", v, nil),
			"type":            utils.PathSearch("type", v, nil),
			"content_id":      utils.PathSearch("content_id", v, nil),
			"subscription_id": utils.PathSearch("subscription_id", v, nil),
		})
	}
	return rst
}

func flattenGetWorkflowResponseBodyWorkflowSubgraph(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("sub_graphs", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"name":  utils.PathSearch("name", v, nil),
			"steps": utils.PathSearch("steps", v, nil),
		})
	}
	return rst
}

func flattenGetWorkflowResponseBodyWorkflowPolicy(resp interface{}) []interface{} {
	var rst []interface{}
	curJson, err := jmespath.Search("policy", resp)
	if err != nil {
		log.Printf("[ERROR] error parsing policy from response= %#v", resp)
		return rst
	}

	rst = []interface{}{
		map[string]interface{}{
			"use_scene": utils.PathSearch("use_scene", curJson, nil),
			"scene_id":  utils.PathSearch("scene_id", curJson, nil),
			"scenes":    flattenWorkflowPolicyScenes(curJson),
		},
	}
	return rst
}

func flattenWorkflowPolicyScenes(resp interface{}) []interface{} {
	if resp == nil {
		return nil
	}
	curJson := utils.PathSearch("scenes", resp, make([]interface{}, 0))
	curArray := curJson.([]interface{})
	rst := make([]interface{}, 0, len(curArray))
	for _, v := range curArray {
		rst = append(rst, map[string]interface{}{
			"id":    utils.PathSearch("id", v, nil),
			"name":  utils.PathSearch("name", v, nil),
			"steps": utils.PathSearch("steps", v, nil),
		})
	}
	return rst
}

func resourceModelartsWorkflowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	updateWorkflowChanges := []string{
		"name",
		"description",
		"steps",
		"data_requirements",
		"parameters",
		"storages",
		"labels",
		"smn_switch",
	}

	if d.HasChanges(updateWorkflowChanges...) {
		var (
			updateWorkflowHttpUrl = "v2/{project_id}/workflows/{id}"
			updateWorkflowProduct = "modelarts"
		)
		updateWorkflowClient, err := cfg.NewServiceClient(updateWorkflowProduct, region)
		if err != nil {
			return diag.Errorf("error creating ModelArts client: %s", err)
		}

		updateWorkflowPath := updateWorkflowClient.Endpoint + updateWorkflowHttpUrl
		updateWorkflowPath = strings.ReplaceAll(updateWorkflowPath, "{project_id}", updateWorkflowClient.ProjectID)
		updateWorkflowPath = strings.ReplaceAll(updateWorkflowPath, "{id}", d.Id())

		updateWorkflowOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
			OkCodes: []int{
				200,
			},
			MoreHeaders: map[string]string{"Content-Type": "application/json"},
		}

		updateWorkflowOpt.JSONBody = utils.RemoveNil(buildUpdateWorkflowBodyParams(d))
		_, err = updateWorkflowClient.Request("PUT", updateWorkflowPath, &updateWorkflowOpt)
		if err != nil {
			return diag.Errorf("error updating Modelarts workflow: %s", err)
		}
	}
	return resourceModelartsWorkflowRead(ctx, d, meta)
}

func buildUpdateWorkflowBodyParams(d *schema.ResourceData) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"name":              utils.ValueIngoreEmpty(d.Get("name")),
		"description":       utils.ValueIngoreEmpty(d.Get("description")),
		"steps":             buildWorkflowReqBodyWorkflowStep(d.Get("steps")),
		"data_requirements": buildWorkflowReqBodyDataRequirement(d.Get("data_requirements")),
		"parameters":        buildWorkflowReqBodyWorkflowParameter(d.Get("parameters")),
		"storages":          buildWorkflowReqBodyWorkflowStorage(d.Get("storages")),
		"labels":            utils.ValueIngoreEmpty(d.Get("labels")),
		"smn_switch":        utils.ValueIngoreEmpty(d.Get("smn_switch")),
	}
	return bodyParams
}

func resourceModelartsWorkflowDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	// deleteWorkflow: delete Modelarts workflow
	var (
		deleteWorkflowHttpUrl = "v2/{project_id}/workflows/{id}"
		deleteWorkflowProduct = "modelarts"
	)
	deleteWorkflowClient, err := cfg.NewServiceClient(deleteWorkflowProduct, region)
	if err != nil {
		return diag.Errorf("error creating ModelArts client: %s", err)
	}

	deleteWorkflowPath := deleteWorkflowClient.Endpoint + deleteWorkflowHttpUrl
	deleteWorkflowPath = strings.ReplaceAll(deleteWorkflowPath, "{project_id}", deleteWorkflowClient.ProjectID)
	deleteWorkflowPath = strings.ReplaceAll(deleteWorkflowPath, "{id}", d.Id())

	deleteWorkflowOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		OkCodes: []int{
			204,
		},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	}

	_, err = deleteWorkflowClient.Request("DELETE", deleteWorkflowPath, &deleteWorkflowOpt)
	if err != nil {
		return diag.Errorf("error deleting Modelarts workflow: %s", err)
	}

	return nil
}
