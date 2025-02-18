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

// @API VSS POST /v3/{project_id}/webscan/domains
// @API VSS POST /v3/{project_id}/webscan/domains/authenticate
// @API VSS POST /v3/{project_id}/webscan/domains/settings
// @API VSS GET /v3/{project_id}/webscan/domains
// @API VSS GET /v3/{project_id}/webscan/domains/settings
// @API VSS DELETE /v3/{project_id}/webscan/domains
func ResourceInspectorWebsite() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInspectorWebsiteCreate,
		UpdateContext: resourceInspectorWebsiteUpdate,
		ReadContext:   resourceInspectorWebsiteRead,
		DeleteContext: resourceInspectorWebsiteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"website_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the website name.`,
			},
			"website_address": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the website address.`,
			},
			"auth_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Specifies the authentication type.`,
			},
			"login_url": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"login_username", "login_password"},
				Description:  `Specifies the login URL.`,
			},
			"login_username": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"login_url", "login_password"},
				Description:  `Specifies the login username.`,
			},
			"login_password": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				RequiredWith: []string{"login_url", "login_username"},
				Description:  `Specifies the login password.`,
			},
			"login_cookie": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: `Specifies the login cookies.`,
			},
			"http_headers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Specifies the HTTP headers.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"verify_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the verify URL.`,
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
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time to create website.`,
			},
			"auth_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The auth status.`,
			},
		},
	}
}

func resourceInspectorWebsiteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg     = meta.(*config.Config)
		region  = cfg.GetRegion(d)
		product = "vss"
	)
	client, err := cfg.NewServiceClient(product, region)
	if err != nil {
		return diag.Errorf("error creating CodeArts inspector client: %s", err)
	}

	if err := createInspectorWebsite(client, d); err != nil {
		return diag.Errorf("error creating CodeArts inspector website: %s", err)
	}

	if err := authorizeInspectorWebsite(client, d); err != nil {
		return diag.Errorf("error authorizing CodeArts inspector website: %s", err)
	}

	if err := updateInspectorWebsiteSettings(client, d); err != nil {
		return diag.Errorf("error updating CodeArts inspector website settings: %s", err)
	}

	return resourceInspectorWebsiteRead(ctx, d, meta)
}

func createInspectorWebsite(client *golangsdk.ServiceClient, d *schema.ResourceData) error {
	createPath := client.Endpoint + "v3/{project_id}/webscan/domains"
	createPath = strings.ReplaceAll(createPath, "{project_id}", client.ProjectID)
	createOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		JSONBody:         buildCreateInspectorWebsiteBodyParams(d),
	}

	createResp, err := client.Request("POST", createPath, &createOpt)
	if err != nil {
		return err
	}

	createRespBody, err := utils.FlattenResponse(createResp)
	if err != nil {
		return err
	}

	domainId := utils.PathSearch("domain_id", createRespBody, "").(string)
	if domainId == "" {
		return fmt.Errorf("unable to find the CodeArts website domain ID from the API response")
	}
	d.SetId(domainId)

	return nil
}

func buildCreateInspectorWebsiteBodyParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"alias":       d.Get("website_name"),
		"domain_name": d.Get("website_address"),
	}
}

func authorizeInspectorWebsite(client *golangsdk.ServiceClient, d *schema.ResourceData) error {
	authorizePath := client.Endpoint + "v3/{project_id}/webscan/domains/authenticate"
	authorizePath = strings.ReplaceAll(authorizePath, "{project_id}", client.ProjectID)
	authorizeOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
		JSONBody:         buildAuthenticateInspectorWebsiteBodyParams(d),
	}

	_, err := client.Request("POST", authorizePath, &authorizeOpt)
	return err
}

func buildAuthenticateInspectorWebsiteBodyParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"domain_name": d.Get("website_address"),
		"auth_mode":   d.Get("auth_type"),
	}
}

func updateInspectorWebsiteSettings(client *golangsdk.ServiceClient, d *schema.ResourceData) error {
	updateSettingsChanges := []string{
		"login_url",
		"login_username",
		"login_password",
		"login_cookie",
		"verify_url",
		"http_headers",
	}

	if d.HasChanges(updateSettingsChanges...) {
		updatePath := client.Endpoint + "v3/{project_id}/webscan/domains/settings"
		updatePath = strings.ReplaceAll(updatePath, "{project_id}", client.ProjectID)
		updateOpt := golangsdk.RequestOpts{
			KeepResponseBody: true,
			JSONBody:         utils.RemoveNil(buildUpdateInspectorWebsiteBodyParams(d)),
		}

		_, err := client.Request("POST", updatePath, &updateOpt)
		return err
	}
	return nil
}

func buildUpdateInspectorWebsiteBodyParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"domain_id":      d.Id(),
		"login_url":      utils.ValueIgnoreEmpty(d.Get("login_url")),
		"login_username": utils.ValueIgnoreEmpty(d.Get("login_username")),
		"login_password": utils.ValueIgnoreEmpty(d.Get("login_password")),
		"login_cookies":  utils.ValueIgnoreEmpty(d.Get("login_cookie")),
		"verify_url":     utils.ValueIgnoreEmpty(d.Get("verify_url")),
		"http_headers":   utils.ValueIgnoreEmpty(d.Get("http_headers")),
	}
}

func resourceInspectorWebsiteRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		mErr    *multierror.Error
		cfg     = meta.(*config.Config)
		region  = cfg.GetRegion(d)
		product = "vss"
	)
	client, err := cfg.NewServiceClient(product, region)
	if err != nil {
		return diag.Errorf("error creating CodeArts inspector client: %s", err)
	}

	getResp, err := readInspectorWebsite(client, d)
	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving CodeArts inspector website")
	}

	getSettingsResp, err := readInspectorWebsiteSettings(client, d)
	if err != nil {
		return diag.Errorf("error retrieving CodeArts inspector website settings: %s", err)
	}

	mErr = multierror.Append(
		mErr,
		d.Set("website_name", utils.PathSearch("alias", getResp, nil)),
		d.Set("website_address", flattenWebsiteAddress(getResp)),
		d.Set("high", utils.PathSearch("high", getResp, nil)),
		d.Set("middle", utils.PathSearch("middle", getResp, nil)),
		d.Set("low", utils.PathSearch("low", getResp, nil)),
		d.Set("hint", utils.PathSearch("hint", getResp, nil)),
		d.Set("created_at", utils.PathSearch("create_time", getResp, nil)),
		d.Set("auth_status", utils.PathSearch("auth_status", getResp, nil)),
		d.Set("login_url", utils.PathSearch("login_url", getSettingsResp, nil)),
		d.Set("login_username", utils.PathSearch("login_username", getSettingsResp, nil)),
		d.Set("verify_url", utils.PathSearch("verify_url", getSettingsResp, nil)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func flattenWebsiteAddress(resp interface{}) interface{} {
	protocolType := utils.PathSearch("protocol_type", resp, "").(string)
	domainName := utils.PathSearch("domain_name", resp, "").(string)
	if protocolType == "" || domainName == "" {
		return nil
	}
	return protocolType + domainName
}

// An example of response body is: {"total": 1, "domains":[{"alias":"XXX", "auth_status":"XXX", ...}]}.
func readInspectorWebsite(client *golangsdk.ServiceClient, d *schema.ResourceData) (interface{}, error) {
	getPath := client.Endpoint + "v3/{project_id}/webscan/domains"
	getPath = strings.ReplaceAll(getPath, "{project_id}", client.ProjectID)
	getPath += buildGetInspectorWebsiteQueryParams(d)
	getOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	getResp, err := client.Request("GET", getPath, &getOpt)
	if err != nil {
		return nil, err
	}

	getRespBody, err := utils.FlattenResponse(getResp)
	if err != nil {
		return nil, err
	}

	domainInfo := utils.PathSearch("domains|[0]", getRespBody, nil)
	if domainInfo == nil {
		return nil, golangsdk.ErrDefault404{}
	}
	return domainInfo, nil
}

func readInspectorWebsiteSettings(client *golangsdk.ServiceClient, d *schema.ResourceData) (interface{}, error) {
	getSettingsPath := client.Endpoint + "v3/{project_id}/webscan/domains/settings"
	getSettingsPath = strings.ReplaceAll(getSettingsPath, "{project_id}", client.ProjectID)
	getSettingsPath += buildGetInspectorWebsiteQueryParams(d)
	getSettingsOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	getSettingsResp, err := client.Request("GET", getSettingsPath, &getSettingsOpt)
	if err != nil {
		return nil, err
	}
	return utils.FlattenResponse(getSettingsResp)
}

func buildGetInspectorWebsiteQueryParams(d *schema.ResourceData) string {
	return fmt.Sprintf("?domain_id=%s", d.Id())
}

func resourceInspectorWebsiteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg     = meta.(*config.Config)
		region  = cfg.GetRegion(d)
		product = "vss"
	)
	client, err := cfg.NewServiceClient(product, region)
	if err != nil {
		return diag.Errorf("error creating CodeArts inspector client: %s", err)
	}

	if err := updateInspectorWebsiteSettings(client, d); err != nil {
		return diag.Errorf("error updating CodeArts inspector website settings: %s", err)
	}

	return resourceInspectorWebsiteRead(ctx, d, meta)
}

func resourceInspectorWebsiteDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg           = meta.(*config.Config)
		product       = "vss"
		region        = cfg.GetRegion(d)
		deleteHttpUrl = "v3/{project_id}/webscan/domains"
	)
	client, err := cfg.NewServiceClient(product, region)
	if err != nil {
		return diag.Errorf("error creating CodeArts inspector client: %s", err)
	}

	deletePath := client.Endpoint + deleteHttpUrl
	deletePath = strings.ReplaceAll(deletePath, "{project_id}", client.ProjectID)
	deletePath += buildDeleteInspectorWebsiteQueryParams(d)
	deleteOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	_, err = client.Request("DELETE", deletePath, &deleteOpt)
	if err != nil {
		// CodeArtsInspector.00006022: The domain name does not exist.
		return common.CheckDeletedDiag(d,
			common.ConvertUndefinedErrInto404Err(err, 418, "error_code", "CodeArtsInspector.00006022"),
			"error deleting CodeArts inspector website",
		)
	}

	return nil
}

func buildDeleteInspectorWebsiteQueryParams(d *schema.ResourceData) string {
	return fmt.Sprintf("?domain_name=%s", d.Get("website_address").(string))
}
