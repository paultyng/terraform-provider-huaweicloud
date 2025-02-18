package css

import (
	"context"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API CSS POST /v1.0/{project_id}/clusters/{cluster_id}/index_snapshot/{snapshot_id}/restore
// @API CSS GET /v1.0/{project_id}/clusters
func ResourceCssRestore() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCssRestoreCreate,
		ReadContext:   resourceCssRestoreRead,
		DeleteContext: resourceCssRestoreDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the source cluster ID.`,
			},
			"snapshot_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the ID of the snapshot to be restored.`,
			},
			"target_cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the target cluster ID.`,
			},
			"indices": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Name of an index to be restored.`,
			},
			"rename_pattern": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Rule for defining the indexes to be restored.`,
			},
			"rename_replacement": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Rule for renaming an index.`,
			},
		},
	}
}

func resourceCssRestoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)
	cssV1Client, err := cfg.CssV1Client(region)
	if err != nil {
		return diag.Errorf("error creating CSS V1 client: %s", err)
	}
	hcCssV1Client, err := cfg.HcCssV1Client(region)
	if err != nil {
		return diag.Errorf("error creating CSS V1 client: %s", err)
	}

	restoreCssCreateHttpUrl := "v1.0/{project_id}/clusters/{cluster_id}/index_snapshot/{snapshot_id}/restore"

	targetClusterID := d.Get("target_cluster_id").(string)
	backupId := d.Get("snapshot_id").(string)
	clusterId := d.Get("source_cluster_id").(string)
	restoreCssCreatePath := cssV1Client.Endpoint + restoreCssCreateHttpUrl
	restoreCssCreatePath = strings.ReplaceAll(restoreCssCreatePath, "{project_id}", cssV1Client.ProjectID)
	restoreCssCreatePath = strings.ReplaceAll(restoreCssCreatePath, "{cluster_id}", clusterId)
	restoreCssCreatePath = strings.ReplaceAll(restoreCssCreatePath, "{snapshot_id}", backupId)

	restoreCssCreateOpt := golangsdk.RequestOpts{KeepResponseBody: true}
	restoreCssCreateOpt.JSONBody = utils.RemoveNil(buildCreateRestoreBodyParams(d))

	_, err = cssV1Client.Request("POST", restoreCssCreatePath, &restoreCssCreateOpt)
	if err != nil {
		return diag.Errorf("error restore CSS cluster snapshot extend, cluster_id: %s, error: %s", d.Id(), err)
	}

	err = checkClusterOperationCompleted(ctx, hcCssV1Client, targetClusterID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(backupId)

	return resourceCssRestoreRead(ctx, d, meta)
}

func buildCreateRestoreBodyParams(d *schema.ResourceData) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"targetCluster":     d.Get("target_cluster_id"),
		"indices":           utils.ValueIgnoreEmpty(d.Get("indices")),
		"renamePattern":     utils.ValueIgnoreEmpty(d.Get("rename_pattern")),
		"renameReplacement": utils.ValueIgnoreEmpty(d.Get("rename_replacement")),
	}
	return bodyParams
}

func resourceCssRestoreRead(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourceCssRestoreDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	errorMsg := "Deleting restoration record is not supported. The restoration record is only removed from the state," +
		" but it remains in the cloud. And the instance doesn't return to the state before restoration."
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  errorMsg,
		},
	}
}
