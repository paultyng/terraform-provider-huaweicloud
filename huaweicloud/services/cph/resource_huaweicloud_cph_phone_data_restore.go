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

var phoneDataRestoreNonUpdatableParams = []string{
	"phone_id",
	"bucket_name",
	"object_path",
}

// @API CPH POST /v1/{project_id}/cloud-phone/phones/batch-restore
// @API CPH GET /v1/{project_id}/cloud-phone/phones/{phone_id}
func ResourcePhoneDataRestore() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePhoneDataRestoreCreate,
		UpdateContext: resourcePhoneDataRestoreUpdate,
		ReadContext:   resourcePhoneDataRestoreRead,
		DeleteContext: resourcePhoneDataRestoreDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: config.FlexibleForceNew(phoneDataRestoreNonUpdatableParams),

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
			"enable_force_new": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"true", "false"}, false),
				Description:  utils.SchemaDesc("", utils.SchemaDescInput{Internal: true}),
			},
		},
	}
}

func resourcePhoneDataRestoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	client, err := cfg.NewServiceClient("cph", region)
	if err != nil {
		return diag.Errorf("error creating CPH client: %s", err)
	}

	// createPhoneDataRestore: create CPH phone data restore
	createPhoneDataRestoreHttpUrl := "v1/{project_id}/cloud-phone/phones/batch-restore"
	createPhoneDataRestorePath := client.Endpoint + createPhoneDataRestoreHttpUrl
	createPhoneDataRestorePath = strings.ReplaceAll(createPhoneDataRestorePath, "{project_id}", client.ProjectID)

	createPhoneDataRestoreOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	createPhoneDataRestoreOpt.JSONBody = map[string]interface{}{
		"restore_infos": []map[string]interface{}{
			{
				"phone_id":    d.Get("phone_id"),
				"bucket_name": d.Get("bucket_name"),
				"object_path": d.Get("object_path"),
			},
		},
	}
	createPhoneDataRestoreResp, err := client.Request("POST", createPhoneDataRestorePath, &createPhoneDataRestoreOpt)
	if err != nil {
		return diag.Errorf("error creating CPH phone data restore: %s", err)
	}

	resp, err := utils.FlattenResponse(createPhoneDataRestoreResp)
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
		return diag.Errorf("failed to restore CPH phone (phone_id: %s) data, error_code: %s, error_msg: %s", id, errorCode, errorMsg)
	}

	err = checkPhoneDataRestoreJobStatus(ctx, client, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourcePhoneDataRestoreRead(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePhoneDataRestoreUpdate(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePhoneDataRestoreDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	errorMsg := "Deleting CPH phone data restore resource is not supported. The resource is only removed from the state."
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  errorMsg,
		},
	}
}

func checkPhoneDataRestoreJobStatus(ctx context.Context, client *golangsdk.ServiceClient, id string, timeout time.Duration) error {
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
		return fmt.Errorf("error waiting for CPH phone data restore to be completed: %s", err)
	}
	return nil
}
