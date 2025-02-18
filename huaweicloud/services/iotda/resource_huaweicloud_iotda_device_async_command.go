package iotda

import (
	"context"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API IoTDA POST /v5/iot/{project_id}/devices/{device_id}/async-commands
// @API IoTDA GET /v5/iot/{project_id}/devices/{device_id}/async-commands/{command_id}
func ResourceDeviceAsyncCommand() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDeviceAsyncCommandCreate,
		ReadContext:   resourceDeviceAsyncCommandRead,
		DeleteContext: resourceDeviceAsyncCommandDelete,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"device_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"send_strategy": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"service_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"paras": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"expire_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"result": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sent_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delivered_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"response_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func buildDeviceAsyncCommandBodyParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"service_id":    utils.ValueIgnoreEmpty(d.Get("service_id")),
		"command_name":  utils.ValueIgnoreEmpty(d.Get("name")),
		"paras":         utils.ValueIgnoreEmpty(d.Get("paras")),
		"expire_time":   utils.ValueIgnoreEmpty(d.Get("expire_time")),
		"send_strategy": d.Get("send_strategy"),
	}
}

func resourceDeviceAsyncCommandCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg       = meta.(*config.Config)
		region    = cfg.GetRegion(d)
		isDerived = WithDerivedAuth(cfg, region)
		deviceId  = d.Get("device_id").(string)
		httpUrl   = "v5/iot/{project_id}/devices/{device_id}/async-commands"
	)

	client, err := cfg.NewServiceClientWithDerivedAuth("iotda", region, isDerived)
	if err != nil {
		return diag.Errorf("error creating IoTDA client: %s", err)
	}

	createPath := client.Endpoint + httpUrl
	createPath = strings.ReplaceAll(createPath, "{project_id}", client.ProjectID)
	createPath = strings.ReplaceAll(createPath, "{device_id}", deviceId)

	createOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		JSONBody:         utils.RemoveNil(buildDeviceAsyncCommandBodyParams(d)),
	}

	resp, err := client.Request("POST", createPath, &createOpt)
	if err != nil {
		return diag.Errorf("error creating device asynchronous command: %s", err)
	}

	respBody, err := utils.FlattenResponse(resp)
	if err != nil {
		return diag.FromErr(err)
	}

	commandId := utils.PathSearch("command_id", respBody, "").(string)
	if commandId == "" {
		return diag.Errorf("error creating device asynchronous command: ID is not found in API response")
	}

	d.SetId(commandId)

	return resourceDeviceAsyncCommandRead(ctx, d, meta)
}

func resourceDeviceAsyncCommandRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg       = meta.(*config.Config)
		region    = cfg.GetRegion(d)
		isDerived = WithDerivedAuth(cfg, region)
		deviceId  = d.Get("device_id").(string)
		httpUrl   = "v5/iot/{project_id}/devices/{device_id}/async-commands/{command_id}"
	)

	client, err := cfg.NewServiceClientWithDerivedAuth("iotda", region, isDerived)
	if err != nil {
		return diag.Errorf("error creating IoTDA client: %s", err)
	}

	getPath := client.Endpoint + httpUrl
	getPath = strings.ReplaceAll(getPath, "{project_id}", client.ProjectID)
	getPath = strings.ReplaceAll(getPath, "{device_id}", deviceId)
	getPath = strings.ReplaceAll(getPath, "{command_id}", d.Id())
	getOpts := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	getResp, err := client.Request("GET", getPath, &getOpts)
	if err != nil {
		// The `device_id` or `command_id` does not exist, query API will return `404`.
		return common.CheckDeletedDiag(d, err, "error retrieving device asynchronous command")
	}

	getRespBody, err := utils.FlattenResponse(getResp)
	if err != nil {
		return diag.FromErr(err)
	}

	mErr := multierror.Append(
		d.Set("region", region),
		d.Set("device_id", utils.PathSearch("device_id", getRespBody, nil)),
		d.Set("service_id", utils.PathSearch("service_id", getRespBody, nil)),
		d.Set("name", utils.PathSearch("command_name", getRespBody, nil)),
		d.Set("status", utils.PathSearch("status", getRespBody, nil)),
		d.Set("result", utils.PathSearch("result", getRespBody, nil)),
		d.Set("send_strategy", utils.PathSearch("send_strategy", getRespBody, nil)),
		d.Set("expire_time", utils.PathSearch("expire_time", getRespBody, nil)),
		d.Set("sent_time", utils.PathSearch("sent_time", getRespBody, nil)),
		d.Set("delivered_time", utils.PathSearch("delivered_time", getRespBody, nil)),
		d.Set("response_time", utils.PathSearch("response_time", getRespBody, nil)),
		d.Set("created_at", utils.PathSearch("created_time", getRespBody, nil)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func resourceDeviceAsyncCommandDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	// No processing is performed in the 'Delete()' method because the resource is a action resource.
	return nil
}
