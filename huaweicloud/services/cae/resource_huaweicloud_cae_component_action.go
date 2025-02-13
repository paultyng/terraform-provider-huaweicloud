package cae

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API CAE POST /v1/{project_id}/cae/applications/{application_id}/components/{component_id}/action
// @API CAE GET /v1/{project_id}/cae/jobs/{job_id}
// ResourceComponentAction is a definition of the one-time action resource that used to operate component.
func ResourceComponentAction() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComponentActionCreate,
		ReadContext:   resourceComponentActionRead,
		UpdateContext: resourceComponentActionUpdate,
		DeleteContext: resourceComponentActionDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `The region in which to create the resource.`,
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the environment where the application is located.`,
			},
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the application where the component is located.`,
			},
			"component_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the component being operated.`,
			},
			"metadata": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The action name.`,
						},
						"annotations": {
							Type:        schema.TypeMap,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: `The key/value pairs parameters related to the component being operated.`,
						},
					},
				},
				Description: `The metadata of this action request.`,
			},
			"spec": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsJSON,
				Description:  `The specification detail of the action, in JSON format.`,
			},
		},
	}
}

func buildCreateComponentActionBodyParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"api_version": "v1",
		"kind":        "Action",
		"metadata": map[string]interface{}{
			"annotations": d.Get("metadata.0.annotations"),
			"name":        d.Get("metadata.0.name"),
		},
		"spec": utils.ValueIgnoreEmpty(utils.StringToJson(d.Get("spec").(string))),
	}
}

func getActionJobDetail(client *golangsdk.ServiceClient, environmentId, jobId string) (interface{}, error) {
	httpUrl := "v1/{project_id}/cae/jobs/{job_id}"
	getPath := client.Endpoint + httpUrl
	getPath = strings.ReplaceAll(getPath, "{project_id}", client.ProjectID)
	getPath = strings.ReplaceAll(getPath, "{job_id}", jobId)

	getOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"X-Environment-ID": environmentId,
		},
	}
	requestResp, err := client.Request("GET", getPath, &getOpt)
	if err != nil {
		return nil, fmt.Errorf("error querying the operation component job detail by its ID (%s): %s", jobId, err)
	}
	respBody, err := utils.FlattenResponse(requestResp)
	if err != nil {
		return nil, fmt.Errorf("error retrieving the operation component job (%s) detail: %s", jobId, err)
	}
	return respBody, nil
}

func deployJobRefreshFunc(client *golangsdk.ServiceClient, environmentId, jobId string, targets []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, err := getActionJobDetail(client, environmentId, jobId)
		if err != nil {
			return resp, "ERROR", err
		}

		status := utils.PathSearch("spec.status", resp, "null").(string)
		if utils.StrSliceContains(targets, status) {
			return resp, "COMPLETED", nil
		}

		if status == "failed" {
			task := utils.PathSearch("spec.tasks[?status=='failed']|[0]", resp, nil)

			return resp, "ERROR",
				fmt.Errorf("the job (%s) execution failed: (%s)", utils.PathSearch("name", task, "").(string),
					utils.PathSearch("detail", task, "").(string))
		}

		return resp, "PENDING", nil
	}
}

func resourceComponentActionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg         = meta.(*config.Config)
		region      = cfg.GetRegion(d)
		componentId = d.Get("component_id").(string)
	)
	client, err := cfg.NewServiceClient("cae", region)
	if err != nil {
		return diag.Errorf("error creating CAE client: %s", err)
	}

	err = doActionComponent(ctx, client, d, componentId, buildCreateComponentActionBodyParams(d), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.Errorf("error operating (%s) the component (%s): %s", d.Get("action").(string), componentId, err)
	}
	d.SetId(componentId)

	return resourceComponentActionRead(ctx, d, meta)
}

func resourceComponentActionRead(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	// No processing is performed in the 'Read()' method because the resource is a one-time action resource.
	return nil
}

func resourceComponentActionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg         = meta.(*config.Config)
		region      = cfg.GetRegion(d)
		componentId = d.Get("component_id").(string)
	)
	client, err := cfg.NewServiceClient("cae", region)
	if err != nil {
		return diag.Errorf("error creating CAE client: %s", err)
	}

	err = doActionComponent(ctx, client, d, componentId, buildCreateComponentActionBodyParams(d), d.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return diag.Errorf("unable to operate (%s) the component (%s): %s", d.Get("action").(string), componentId, err)
	}

	return resourceComponentActionRead(ctx, d, meta)
}

func resourceComponentActionDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	// No processing is performed in the 'Delete()' method because the resource is a one-time action resource.
	errorMsg := `This resource is only a one-time action resource for operating the component. Deleting this resource will
not clear the corresponding request record, but will only remove the resource information from the tfstate file.`
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  errorMsg,
		},
	}
}
