package workspace

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

// @API Workspace POST /v2/{project_id}/ous
// @API Workspace GET /v2/{project_id}/ous
// @API Workspace PUT /v2/{project_id}/ous/{ou_id}
// @API Workspace DELETE /v2/{project_id}/ous/{ou_id}
func ResourceOu() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOuCreate,
		ReadContext:   resourceOuRead,
		UpdateContext: resourceOuUpdate,
		DeleteContext: resourceOuDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceOuImportState,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The name of the OU.`,
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The domain name to which the OU belongs.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The description of the OU.`,
			},
		},
	}
}

func buildCreateOuBodyParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"ou_name":     d.Get("name"),
		"domain":      d.Get("domain"),
		"description": utils.ValueIgnoreEmpty(d.Get("description")),
	}
}

func resourceOuCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg     = meta.(*config.Config)
		httpUrl = "v2/{project_id}/ous"
	)

	client, err := cfg.NewServiceClient("workspace", cfg.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating Workspace client: %s", err)
	}

	createPath := client.Endpoint + httpUrl
	createPath = strings.ReplaceAll(createPath, "{project_id}", client.ProjectID)
	createOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		JSONBody:         utils.RemoveNil(buildCreateOuBodyParams(d)),
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}
	resp, err := client.Request("POST", createPath, &createOpt)
	if err != nil {
		return diag.Errorf("error creating OU: %s", err)
	}

	respBody, err := utils.FlattenResponse(resp)
	if err != nil {
		return diag.FromErr(err)
	}

	ouId := utils.PathSearch("id", respBody, "").(string)
	if ouId == "" {
		return diag.Errorf("unable to find OU ID from API response")
	}

	d.SetId(ouId)
	return resourceOuRead(ctx, d, meta)
}

func GetOuByName(client *golangsdk.ServiceClient, ouName string) (interface{}, error) {
	httpUrl := "v2/{project_id}/ous"
	listPath := client.Endpoint + httpUrl
	listPath = strings.ReplaceAll(listPath, "{project_id}", client.ProjectID)
	listPath = fmt.Sprintf("%s?ou_name=%s", listPath, ouName)
	listOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}
	resp, err := client.Request("GET", listPath, &listOpt)
	if err != nil {
		// WKS.0202: The resource not found.
		// WKS.00010099: The Workspace service is disabled.
		return nil, common.ConvertExpected500ErrInto404Err(
			common.ConvertExpected403ErrInto404Err(err, "error_code", "WKS.00010099"),
			"error_code", "WKS.0202")
	}

	respBody, err := utils.FlattenResponse(resp)
	if err != nil {
		return nil, err
	}

	ous := utils.PathSearch("ous", respBody, make([]interface{}, 0)).([]interface{})
	if len(ous) == 0 {
		return nil, golangsdk.ErrDefault404{}
	}

	return ous[0], nil
}

func resourceOuRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg    = meta.(*config.Config)
		region = cfg.GetRegion(d)
	)

	client, err := cfg.NewServiceClient("workspace", region)
	if err != nil {
		return diag.Errorf("error creating Workspace client: %s", err)
	}

	respBody, err := GetOuByName(client, d.Get("name").(string))
	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving Workspace OU")
	}

	mErr := multierror.Append(nil,
		d.Set("region", region),
		d.Set("name", utils.PathSearch("ou_name", respBody, nil)),
		d.Set("domain", utils.PathSearch("domain", respBody, nil)),
		d.Set("description", utils.PathSearch("description", respBody, nil)),
	)
	return diag.FromErr(mErr.ErrorOrNil())
}

func buildUpdateOuBodyParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"ou_name":     d.Get("name"),
		"description": d.Get("description"),
	}
}

func resourceOuUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg     = meta.(*config.Config)
		httpUrl = "v2/{project_id}/ous/{ou_id}"
	)

	client, err := cfg.NewServiceClient("workspace", cfg.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating Workspace client: %s", err)
	}

	updatePath := client.Endpoint + httpUrl
	updatePath = strings.ReplaceAll(updatePath, "{project_id}", client.ProjectID)
	updatePath = strings.ReplaceAll(updatePath, "{ou_id}", d.Id())
	updateOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		JSONBody:         utils.RemoveNil(buildUpdateOuBodyParams(d)),
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}
	_, err = client.Request("PUT", updatePath, &updateOpt)
	if err != nil {
		return diag.Errorf("error updating OU (%s): %s", d.Get("name").(string), err)
	}

	return resourceOuRead(ctx, d, meta)
}

func resourceOuDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg     = meta.(*config.Config)
		httpUrl = "v2/{project_id}/ous/{ou_id}"
	)

	client, err := cfg.NewServiceClient("workspace", cfg.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating Workspace client: %s", err)
	}

	deletePath := client.Endpoint + httpUrl
	deletePath = strings.ReplaceAll(deletePath, "{project_id}", client.ProjectID)
	deletePath = strings.ReplaceAll(deletePath, "{ou_id}", d.Id())
	deleteOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	}
	_, err = client.Request("DELETE", deletePath, &deleteOpt)
	if err != nil {
		// WKS.0915: The resource not found.
		// WKS.00010099: The Workspace service is disabled.
		return common.CheckDeletedDiag(d, common.ConvertExpected500ErrInto404Err(
			common.ConvertExpected403ErrInto404Err(err, "error_code", "WKS.00010099"),
			"error_code", "WKS.0915"),
			"error deleting Workspace OU")
	}

	return nil
}

func resourceOuImportState(_ context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	importId := d.Id()
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	client, err := cfg.NewServiceClient("workspace", region)
	if err != nil {
		return nil, fmt.Errorf("error creating Workspace client: %s", err)
	}

	ouInfo, err := GetOuByName(client, importId)
	if err != nil {
		return nil, err
	}

	ouId := utils.PathSearch("id", ouInfo, "").(string)
	if ouId == "" {
		return nil, fmt.Errorf("unable to find the OU ID using its name (%s): %s", importId, err)
	}

	d.SetId(ouId)
	mErr := multierror.Append(nil,
		d.Set("name", importId),
	)
	return []*schema.ResourceData{d}, mErr.ErrorOrNil()
}
