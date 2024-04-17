package lts

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jmespath/go-jmespath"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/eps"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

const EPSTagKey string = "_sys_enterprise_project_id"

// @API LTS POST /v2/{project_id}/groups/{log_group_id}/streams
// @API LTS GET /v2/{project_id}/groups/{log_group_id}/streams
// @API LTS DELETE /v2/{project_id}/groups/{log_group_id}/streams/{log_stream_id}
func ResourceLTSStream() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceStreamCreate,
		ReadContext:   resourceStreamRead,
		DeleteContext: resourceStreamDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceStreamImportState,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"stream_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ttl_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"enterprise_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Attributes
			// tags of stream will be changed when the tags of group has been changed
			// and the API cannot support updating tags, so we should mark tags as computed.
			"tags": common.TagsComputedSchema(),
			"filter_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceStreamCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg     = meta.(*config.Config)
		region  = cfg.GetRegion(d)
		httpUrl = "v2/{project_id}/groups/{log_group_id}/streams"
	)

	client, err := cfg.NewServiceClient("lts", region)
	if err != nil {
		return diag.Errorf("error creating LTS client: %s", err)
	}
	createPath := client.Endpoint + httpUrl
	createPath = strings.ReplaceAll(createPath, "{project_id}", client.ProjectID)
	createPath = strings.ReplaceAll(createPath, "{log_group_id}", d.Get("group_id").(string))

	createOpts := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders:      map[string]string{"Content-Type": "application/json"},
		JSONBody:         utils.RemoveNil(buildCreateStreamBodyParams(cfg, d)),
	}

	requestResp, err := client.Request("POST", createPath, &createOpts)
	if err != nil {
		return diag.Errorf("error creating LTS stream: %s", err)
	}

	respBody, err := utils.FlattenResponse(requestResp)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := jmespath.Search("log_stream_id", respBody)
	if err != nil {
		return diag.Errorf("error creating flow log: ID is not found in API response")
	}

	d.SetId(id.(string))

	return resourceStreamRead(ctx, d, meta)
}

func buildCreateStreamBodyParams(cfg *config.Config, d *schema.ResourceData) map[string]interface{} {
	bodyParams := map[string]interface{}{
		"log_stream_name": d.Get("stream_name"),
		"ttl_in_days":     utils.ValueIngoreEmpty(d.Get("ttl_in_days")),
	}

	userNoPermission := []string{"EPS.0004"}
	epsId := cfg.GetEnterpriseProjectID(d)
	if epsId == "" {
		epsId = "0"
	}
	epsInfo, err := getEnterpriseProjectById(cfg, cfg.GetRegion(d), epsId)
	// If we catch error 403, it means that the user does not have EPS permissions, return immediately.
	if parsedErr := eps.ParseQueryError403(err, userNoPermission, "No permission, skip the enterprise project query"); parsedErr == nil {
		// Unable to set enterprise project ID for log stream via parameter 'enterprise_project_id' and
		// 'tags._sys_enterprise_project_id'. Currently, only parameter 'enterprise_project_name' is available.
		bodyParams["enterprise_project_name"] = utils.PathSearch("enterprise_project.name", epsInfo, nil)
	}
	// If not, insert enterprise project name into bodyParams.
	return bodyParams
}

func getEnterpriseProjectById(cfg *config.Config, region, epsId string) (interface{}, error) {
	client, err := cfg.NewServiceClient("eps", region)
	if err != nil {
		return nil, fmt.Errorf("error creating EPS client: %s", err)
	}

	httpUrl := "v1.0/enterprise-projects/{enterprise_project_id}"
	getPath := client.Endpoint + httpUrl
	getPath = strings.ReplaceAll(getPath, "{enterprise_project_id}", epsId)

	getOpts := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders:      map[string]string{"Content-Type": "application/json"},
	}

	requestResp, err := client.Request("GET", getPath, &getOpts)
	if err != nil {
		return nil, err
	}

	respBody, err := utils.FlattenResponse(requestResp)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func resourceStreamRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg      = meta.(*config.Config)
		region   = cfg.GetRegion(d)
		httpUrl  = "v2/{project_id}/groups/{log_group_id}/streams"
		streamId = d.Id()
	)

	client, err := cfg.NewServiceClient("lts", region)
	if err != nil {
		return diag.Errorf("error creating LTS client: %s", err)
	}
	getPath := client.Endpoint + httpUrl
	getPath = strings.ReplaceAll(getPath, "{project_id}", client.ProjectID)
	getPath = strings.ReplaceAll(getPath, "{log_group_id}", d.Get("group_id").(string))

	getOpts := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders:      map[string]string{"Content-Type": "application/json"},
	}

	requestResp, err := client.Request("GET", getPath, &getOpts)
	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving log stream")
	}

	respBody, err := utils.FlattenResponse(requestResp)
	if err != nil {
		return diag.Errorf("error parsing the log stream: %s", err)
	}

	streamResult := utils.PathSearch(fmt.Sprintf("log_streams|[?log_stream_id=='%s']|[0]", streamId), respBody, nil)
	if streamResult == nil {
		return common.CheckDeletedDiag(d, err, fmt.Sprintf("unable to find log stream by its ID (%s)", streamId))
	}

	mErr := multierror.Append(nil,
		d.Set("region", region),
		d.Set("stream_name", utils.PathSearch("log_stream_name", streamResult, nil)),
		d.Set("enterprise_project_id", utils.PathSearch("tag._sys_enterprise_project_id", streamResult, nil)),
		d.Set("tags", ignoreSysEpsTag(utils.PathSearch("tag", streamResult, make(map[string]interface{})).(map[string]interface{}))),
		d.Set("filter_count", utils.PathSearch("filter_count", streamResult, nil)),
		d.Set("created_at", utils.FormatTimeStampRFC3339(int64(utils.PathSearch("creation_time", streamResult, 0).(float64))/1000, false)),
	)
	return diag.FromErr(mErr.ErrorOrNil())
}

func resourceStreamDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg      = meta.(*config.Config)
		region   = cfg.GetRegion(d)
		httpUrl  = "v2/{project_id}/groups/{log_group_id}/streams/{log_stream_id}"
		streamId = d.Id()
	)

	client, err := cfg.NewServiceClient("lts", region)
	if err != nil {
		return diag.Errorf("error creating LTS client: %s", err)
	}
	deletePath := client.Endpoint + httpUrl
	deletePath = strings.ReplaceAll(deletePath, "{project_id}", client.ProjectID)
	deletePath = strings.ReplaceAll(deletePath, "{log_group_id}", d.Get("group_id").(string))
	deletePath = strings.ReplaceAll(deletePath, "{log_stream_id}", streamId)

	deleteOpts := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders:      map[string]string{"Content-Type": "application/json"},
	}
	_, err = client.Request("DELETE", deletePath, &deleteOpts)
	if err != nil {
		return common.CheckDeletedDiag(d, err, "error deleting log stream")
	}
	return nil
}

func resourceStreamImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid ID format, want '<group_id>/<stream_id>', but '%s'", d.Id())
	}

	groupID := parts[0]
	streamID := parts[1]

	d.SetId(streamID)
	mErr := multierror.Append(nil,
		d.Set("group_id", groupID),
	)
	return []*schema.ResourceData{d}, mErr.ErrorOrNil()
}
