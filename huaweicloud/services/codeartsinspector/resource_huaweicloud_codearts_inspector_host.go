package codeartsinspector

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API VSS POST /v3/{project_id}/hostscan/hosts
// @API VSS GET /v3/{project_id}/hostscan/hosts
// @API VSS DELETE /v3/{project_id}/hostscan/hosts/delete/{host_id}
func ResourceInspectorHost() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInspectorHostCreate,
		ReadContext:   resourceInspectorHostRead,
		DeleteContext: resourceInspectorHostDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the host name.`,
			},
			"ip": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the host IP.`,
			},
			"os_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the host os type.`,
			},
			"group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Specifies the host group ID.`,
			},
			"ssh_credential_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  `Specifies the host ssh credential ID for linux host.`,
				ExactlyOneOf: []string{"ssh_credential_id", "smb_credential_id"},
			},
			"jumper_server_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Specifies the jumper server ID. Only available for linux host.`,
			},
			"smb_credential_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Specifies the smb credential ID for windows host.`,
			},
			"auth_status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Indicates the auth status.`,
			},
			"last_scan_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Indicates the last scan ID.`,
			},
			"last_scan_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Indicates the last scan informations.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_weak_passwd": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: `Indicates whether weak password check enabled.`,
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `Indicates the task status.`,
						},
						"reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the task description.`,
						},
						"progress": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `Indicates the task progress.`,
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the scan task create time.`,
						},
						"end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the scan task end time.`,
						},
					},
				},
			},
		},
	}
}

func resourceInspectorHostCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)
	client, err := cfg.NewServiceClient("vss", region)
	if err != nil {
		return diag.Errorf("error creating CodeArts inspector client: %s", err)
	}

	createHttpUrl := "v3/{project_id}/hostscan/hosts"
	createPath := client.Endpoint + createHttpUrl
	createPath = strings.ReplaceAll(createPath, "{project_id}", client.ProjectID)
	createOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		JSONBody:         utils.RemoveNil(buildCreateInspectorHostBodyParams(d)),
	}

	createResp, err := client.Request("POST", createPath, &createOpt)
	if err != nil {
		return diag.Errorf("error creating host: %s", err)
	}

	createRespBody, err := utils.FlattenResponse(createResp)
	if err != nil {
		return diag.FromErr(err)
	}

	hostId := utils.PathSearch("items|[0].id", createRespBody, "").(string)
	if hostId == "" {
		return diag.Errorf("unable find host ID from the API response")
	}
	d.SetId(hostId)

	return resourceInspectorHostRead(ctx, d, meta)
}

func buildCreateInspectorHostBodyParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"hosts": []map[string]interface{}{
			{
				"name":              d.Get("name"),
				"ip":                d.Get("ip"),
				"os_type":           d.Get("os_type"),
				"group_id":          utils.ValueIgnoreEmpty(d.Get("group_id")),
				"ssh_credential_id": utils.ValueIgnoreEmpty(d.Get("ssh_credential_id")),
				"jumper_server_id":  utils.ValueIgnoreEmpty(d.Get("jumper_server_id")),
				"smb_credential_id": utils.ValueIgnoreEmpty(d.Get("smb_credential_id")),
			},
		},
	}
}

func resourceInspectorHostRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)
	client, err := cfg.NewServiceClient("vss", region)
	if err != nil {
		return diag.Errorf("error creating CodeArts inspector client: %s", err)
	}

	host, err := GetInspectorHost(client, d.Id())
	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving host")
	}

	mErr := multierror.Append(nil,
		d.Set("name", utils.PathSearch("name", host, nil)),
		d.Set("ip", utils.PathSearch("ip", host, nil)),
		d.Set("os_type", utils.PathSearch("os_type", host, nil)),
		d.Set("group_id", utils.PathSearch("group_id", host, nil)),
		d.Set("ssh_credential_id", utils.PathSearch("ssh_credential_id", host, nil)),
		d.Set("jumper_server_id", utils.PathSearch("jumper_server_id", host, nil)),
		d.Set("smb_credential_id", utils.PathSearch("smb_credential_id", host, nil)),
		d.Set("auth_status", utils.PathSearch("auth_status", host, nil)),
		d.Set("last_scan_id", utils.PathSearch("last_scan_id", host, nil)),
		d.Set("last_scan_info", flattenInspectorHostLastScanInfo(utils.PathSearch("last_scan_info", host, nil))),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func flattenInspectorHostLastScanInfo(params interface{}) []interface{} {
	if params == nil {
		return nil
	}

	rst := map[string]interface{}{
		"enable_weak_passwd": utils.PathSearch("enable_weak_passwd", params, nil),
		"status":             utils.PathSearch("status", params, nil),
		"reason":             utils.PathSearch("reason", params, nil),
		"progress":           utils.PathSearch("progress", params, nil),
		"create_time": utils.FormatTimeStampRFC3339(
			int64(utils.PathSearch("create_time", params, float64(0)).(float64))/1000, false),
		"end_time": utils.FormatTimeStampRFC3339(
			int64(utils.PathSearch("end_time", params, float64(0)).(float64))/1000, false),
	}

	return []interface{}{rst}
}

func GetInspectorHost(client *golangsdk.ServiceClient, id string) (interface{}, error) {
	listGroupsHttpUrl := "v3/{project_id}/hostscan/hosts"
	listGroupsPath := client.Endpoint + listGroupsHttpUrl
	listGroupsPath = strings.ReplaceAll(listGroupsPath, "{project_id}", client.ProjectID)
	listGroupsOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	// pagelimit is `10`
	listGroupsPath += fmt.Sprintf("?limit=%v", pageLimit)

	currentTotal := 0
	for {
		currentPath := listGroupsPath + fmt.Sprintf("&offset=%d", currentTotal)
		listGroupsResp, err := client.Request("GET", currentPath, &listGroupsOpt)
		if err != nil {
			return nil, err
		}
		listGroupsRespBody, err := utils.FlattenResponse(listGroupsResp)
		if err != nil {
			return nil, err
		}

		searchPath := fmt.Sprintf("items[?id=='%s']|[0]", id)
		host := utils.PathSearch(searchPath, listGroupsRespBody, nil)
		if host != nil {
			return host, nil
		}

		hosts := utils.PathSearch("items", listGroupsRespBody, make([]interface{}, 0)).([]interface{})
		currentTotal += len(hosts)
		totalCount := utils.PathSearch("total", listGroupsRespBody, float64(0))
		if int(totalCount.(float64)) <= currentTotal {
			break
		}
	}

	return nil, golangsdk.ErrDefault404{}
}

func resourceInspectorHostDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)
	client, err := cfg.NewServiceClient("vss", region)
	if err != nil {
		return diag.Errorf("error creating CodeArts inspector client: %s", err)
	}

	deleteHttpUrl := "v3/{project_id}/hostscan/hosts/delete/{host_id}"
	deletePath := client.Endpoint + deleteHttpUrl
	deletePath = strings.ReplaceAll(deletePath, "{project_id}", client.ProjectID)
	deletePath = strings.ReplaceAll(deletePath, "{host_id}", d.Id())

	deleteOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	_, err = client.Request("DELETE", deletePath, &deleteOpt)
	if err != nil {
		return common.CheckDeletedDiag(d,
			common.ConvertExpected401ErrInto404Err(err, "error_code", "Host is invalid"),
			"error deleting host",
		)
	}

	return nil
}
