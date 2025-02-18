package codeartsinspector

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

const (
	taskNotExistCode = "CodeArtsInspector.00002008" // The task does not exist.
	taskEndedCode    = "CodeArtsInspector.00002023" // The task has ended, the task cannot be canceled.
)

// @API VSS POST /v3/{project_id}/webscan/tasks
// @API VSS PUT /v3/{project_id}/webscan/tasks
// @API VSS GET /v3/{project_id}/webscan/tasks
func ResourceInspectorWebsiteScan() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInspectorWebsiteScanCreate,
		ReadContext:   resourceInspectorWebsiteScanRead,
		DeleteContext: resourceInspectorWebsiteScanDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"task_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the task name.`,
			},
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the URL.`,
			},
			"task_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Specifies the task type.`,
			},
			"timer": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Specifies the scheduled trigger time of the normal task.`,
			},
			"trigger_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Specifies the scheduled trigger time of the monitor task.`,
			},
			"task_period": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Specifies the scheduled trigger period of the monitor task.`,
			},
			"scan_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Specifies the task scan mode.`,
			},
			"port_scan": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Specifies whether to perform port scanning.`,
			},
			"weak_pwd_scan": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Specifies whether to scan for weak passwords.`,
			},
			"cve_check": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Specifies whether to perform CVE vulnerability scanning.`,
			},
			"text_check": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Specifies whether to conduct website content compliance text detection.`,
			},
			"picture_check": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Specifies whether to conduct website content compliance image detection.`,
			},
			"malicious_code": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Specifies whether to perform malicious code scanning.`,
			},
			"malicious_link": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: `Specifies whether to perform link health detection.`,
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The create time of the task.`,
			},
			"task_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The task status.`,
			},
			"schedule_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The monitor task status.`,
			},
			"progress": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The task progress.`,
			},
			"reason": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Task status description.`,
			},
			"pack_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The total number of packages.`,
			},
			"score": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The safety score.`,
			},
			"safe_level": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The security level.`,
			},
			"high": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The number of high-risk vulnerabilities.`,
			},
			"middle": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The number of medium-risk vulnerabilities.`,
			},
			"low": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The number of low-severity vulnerabilities.`,
			},
			"hint": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The number of hint-risk vulnerabilities.`,
			},
		},
	}
}

func resourceInspectorWebsiteScanCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg     = meta.(*config.Config)
		region  = cfg.GetRegion(d)
		httpUrl = "v3/{project_id}/webscan/tasks"
		product = "vss"
	)
	client, err := cfg.NewServiceClient(product, region)
	if err != nil {
		return diag.Errorf("error creating CodeArts inspector client: %s", err)
	}

	createPath := client.Endpoint + httpUrl
	createPath = strings.ReplaceAll(createPath, "{project_id}", client.ProjectID)
	createOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		JSONBody:         utils.RemoveNil(buildCreateInspectorWebsiteScanBodyParams(d)),
	}

	createResp, err := client.Request("POST", createPath, &createOpt)
	if err != nil {
		return diag.Errorf("error creating CodeArts inspector website scan: %s", err)
	}

	createRespBody, err := utils.FlattenResponse(createResp)
	if err != nil {
		return diag.FromErr(err)
	}

	taskId := utils.PathSearch("task_id", createRespBody, "").(string)
	if taskId == "" {
		return diag.Errorf("unable find the CodeArts inspector website scan ID from the API response")
	}
	d.SetId(taskId)

	return resourceInspectorWebsiteScanRead(ctx, d, meta)
}

func buildCreateInspectorWebsiteScanBodyParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"task_name":    d.Get("task_name"),
		"url":          d.Get("url"),
		"task_type":    utils.ValueIgnoreEmpty(d.Get("task_type")),
		"timer":        utils.ValueIgnoreEmpty(d.Get("timer")),
		"trigger_time": utils.ValueIgnoreEmpty(d.Get("trigger_time")),
		"task_period":  utils.ValueIgnoreEmpty(d.Get("task_period")),
		"task_config":  buildCreateInspectorWebsiteScanTaskConfig(d),
	}
}

func buildCreateInspectorWebsiteScanTaskConfig(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"scan_mode":      utils.ValueIgnoreEmpty(d.Get("scan_mode")),
		"port_scan":      d.Get("port_scan"),
		"weak_pwd_scan":  d.Get("weak_pwd_scan"),
		"cve_check":      d.Get("cve_check"),
		"text_check":     d.Get("text_check"),
		"picture_check":  d.Get("picture_check"),
		"malicious_code": d.Get("malicious_code"),
		"malicious_link": d.Get("malicious_link"),
	}
}

func resourceInspectorWebsiteScanRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg     = meta.(*config.Config)
		region  = cfg.GetRegion(d)
		mErr    *multierror.Error
		httpUrl = "v3/{project_id}/webscan/tasks"
		product = "vss"
	)
	client, err := cfg.NewServiceClient(product, region)
	if err != nil {
		return diag.Errorf("error creating CodeArts inspector client: %s", err)
	}

	getPath := client.Endpoint + httpUrl
	getPath = strings.ReplaceAll(getPath, "{project_id}", client.ProjectID)
	getPath += buildGetInspectorWebsiteScanQueryParams(d)
	getOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	getResp, err := client.Request("GET", getPath, &getOpt)
	if err != nil {
		return common.CheckDeletedDiag(d,
			common.ConvertUndefinedErrInto404Err(err, 418, "error_code", taskNotExistCode),
			"error retrieving CodeArts inspector website scan",
		)
	}

	getRespBody, err := utils.FlattenResponse(getResp)
	if err != nil {
		return diag.FromErr(err)
	}

	taskStatus := utils.PathSearch("task_status", getRespBody, "")
	if taskStatus == "" {
		return diag.Errorf("error retrieving CodeArts inspector website scan: field `task_status` is not found" +
			" in detail API response")
	}

	if taskStatus == "canceled" {
		return common.CheckDeletedDiag(d, golangsdk.ErrDefault404{}, "")
	}

	mErr = multierror.Append(
		mErr,
		d.Set("task_name", utils.PathSearch("task_name", getRespBody, nil)),
		d.Set("url", utils.PathSearch("url", getRespBody, nil)),
		d.Set("task_type", utils.PathSearch("task_type", getRespBody, nil)),
		d.Set("timer", utils.PathSearch("task_settings.timer", getRespBody, nil)),
		d.Set("trigger_time", utils.PathSearch("task_settings.trigger_time", getRespBody, nil)),
		d.Set("task_period", utils.PathSearch("task_settings.task_period", getRespBody, nil)),
		d.Set("scan_mode", utils.PathSearch("task_settings.task_config.scan_mode", getRespBody, nil)),
		d.Set("port_scan", utils.PathSearch("task_settings.task_config.port_scan", getRespBody, nil)),
		d.Set("weak_pwd_scan", utils.PathSearch("task_settings.task_config.weak_pwd_scan", getRespBody, nil)),
		d.Set("cve_check", utils.PathSearch("task_settings.task_config.cve_check", getRespBody, nil)),
		d.Set("text_check", utils.PathSearch("task_settings.task_config.text_check", getRespBody, nil)),
		d.Set("picture_check", utils.PathSearch("task_settings.task_config.picture_check", getRespBody, nil)),
		d.Set("malicious_code", utils.PathSearch("task_settings.task_config.malicious_code", getRespBody, nil)),
		d.Set("malicious_link", utils.PathSearch("task_settings.task_config.malicious_link", getRespBody, nil)),
		d.Set("created_at", utils.PathSearch("create_time", getRespBody, nil)),
		d.Set("task_status", taskStatus),
		d.Set("schedule_status", utils.PathSearch("schedule_status", getRespBody, nil)),
		d.Set("progress", utils.PathSearch("progress", getRespBody, nil)),
		d.Set("reason", utils.PathSearch("reason", getRespBody, nil)),
		d.Set("pack_num", utils.PathSearch("pack_num", getRespBody, nil)),
		d.Set("score", utils.PathSearch("score", getRespBody, nil)),
		d.Set("safe_level", utils.PathSearch("safe_level", getRespBody, nil)),
		d.Set("high", utils.PathSearch("statistics.high", getRespBody, nil)),
		d.Set("middle", utils.PathSearch("statistics.middle", getRespBody, nil)),
		d.Set("low", utils.PathSearch("statistics.low", getRespBody, nil)),
		d.Set("hint", utils.PathSearch("statistics.hint", getRespBody, nil)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func buildGetInspectorWebsiteScanQueryParams(d *schema.ResourceData) string {
	return fmt.Sprintf("?task_id=%s", d.Id())
}

func resourceInspectorWebsiteScanDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	taskStatus := d.Get("task_status").(string)
	ignoreStatuses := []string{"success", "canceled", "failure"}
	// When the task status is `success`, `canceled`, `failure`, the task cancellation operation is not performed.
	if utils.StrSliceContains(ignoreStatuses, taskStatus) {
		log.Printf("[DEBUG] The current task status is `%s`, the task (%s) cancellation operation is not performed.",
			taskStatus, d.Id())
		return nil
	}

	var (
		cfg     = meta.(*config.Config)
		region  = cfg.GetRegion(d)
		httpUrl = "v3/{project_id}/webscan/tasks"
		product = "vss"
	)
	client, err := cfg.NewServiceClient(product, region)
	if err != nil {
		return diag.Errorf("error creating CodeArts inspector client: %s", err)
	}

	cancelPath := client.Endpoint + httpUrl
	cancelPath = strings.ReplaceAll(cancelPath, "{project_id}", client.ProjectID)
	cancelOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		JSONBody:         buildDeleteInspectorWebsiteScanBodyParams(d),
	}

	_, err = client.Request("PUT", cancelPath, &cancelOpt)
	if err != nil {
		return common.CheckDeletedDiag(d,
			common.ConvertUndefinedErrInto404Err(err, 418, "error_code", []string{taskNotExistCode, taskEndedCode}...),
			"error canceling CodeArts inspector website scan",
		)
	}

	return nil
}

func buildDeleteInspectorWebsiteScanBodyParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"task_id": d.Id(),
		"action":  "cancel",
	}
}
