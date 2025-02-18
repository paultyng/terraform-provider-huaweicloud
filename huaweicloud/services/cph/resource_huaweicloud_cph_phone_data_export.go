package cph

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

var phoneDataExportNonUpdatableParams = []string{
	"phone_id",
	"bucket_name",
	"object_path",
	"include_files",
	"exclude_files",
}

// @API CPH POST /v1/{project_id}/cloud-phone/phones/batch-storage
// @API CPH GET /v1/{project_id}/cloud-phone/phones/{phone_id}
func ResourcePhoneDataExport() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePhoneDataExportCreate,
		UpdateContext: resourcePhoneDataExportUpdate,
		ReadContext:   resourcePhoneDataExportRead,
		DeleteContext: resourcePhoneDataExportDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: config.FlexibleForceNew(phoneDataExportNonUpdatableParams),

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"phone_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the phone ID.`,
			},
			"bucket_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the bucket name of OBS.`,
			},
			"object_path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the object path of OBS.`,
			},
			"include_files": {
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: `Specifies the include files.`,
			},
			"exclude_files": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: `Specifies the exclude files.`,
			},
			"enable_force_new": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"true", "false"}, false),
				Description:  utils.SchemaDesc("", utils.SchemaDescInput{Internal: true}),
			},
		},
	}
}

func resourcePhoneDataExportCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	client, err := cfg.NewServiceClient("cph", region)
	if err != nil {
		return diag.Errorf("error creating CPH client: %s", err)
	}

	// createPhoneDataExport: create CPH phone data export
	createPhoneDataExportHttpUrl := "v1/{project_id}/cloud-phone/phones/batch-storage"
	createPhoneDataExportPath := client.Endpoint + createPhoneDataExportHttpUrl
	createPhoneDataExportPath = strings.ReplaceAll(createPhoneDataExportPath, "{project_id}", client.ProjectID)

	createPhoneDataExportOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	createPhoneDataExportOpt.JSONBody = utils.RemoveNil(map[string]interface{}{
		"storage_infos": []map[string]interface{}{
			{
				"phone_id":      d.Get("phone_id"),
				"bucket_name":   d.Get("bucket_name"),
				"object_path":   d.Get("object_path"),
				"include_files": d.Get("include_files"),
				"exclude_files": utils.ValueIgnoreEmpty(d.Get("exclude_files")),
			},
		},
	})
	createPhoneDataExportResp, err := client.Request("POST", createPhoneDataExportPath, &createPhoneDataExportOpt)
	if err != nil {
		return diag.Errorf("error creating CPH phone data export: %s", err)
	}

	resp, err := utils.FlattenResponse(createPhoneDataExportResp)
	if err != nil {
		return diag.FromErr(err)
	}
	id := utils.PathSearch("jobs|[0].phone_id", resp, "").(string)
	if id == "" {
		return diag.Errorf("unable to find the phone ID from the API response")
	}
	d.SetId(id)

	errorCode := utils.PathSearch("jobs|[0].error_code", resp, "").(string)
	if errorCode != "" {
		errorMsg := utils.PathSearch("jobs|[0].error_msg", resp, "").(string)
		return diag.Errorf("failed to export CPH phone (phone_id: %s) data, error_code: %s, error_msg: %s", id, errorCode, errorMsg)
	}

	err = checkPhoneDataExportJobStatus(ctx, client, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourcePhoneDataExportRead(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePhoneDataExportUpdate(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePhoneDataExportDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	errorMsg := "Deleting CPH phone data export resource is not supported. The resource is only removed from the state."
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  errorMsg,
		},
	}
}

func checkPhoneDataExportJobStatus(ctx context.Context, client *golangsdk.ServiceClient, id string, timeout time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      jobStatusRefreshFunc(client, id),
		Timeout:      timeout,
		PollInterval: 10 * timeout,
		Delay:        10 * time.Second,
	}
	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return fmt.Errorf("error waiting for CPH phone data export to be completed: %s", err)
	}
	return nil
}
