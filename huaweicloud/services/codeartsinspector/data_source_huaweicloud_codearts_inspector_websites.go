package codeartsinspector

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API VSS GET /v3/{project_id}/webscan/domains
func DataSourceCodeartsInspectorWebsites() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCodeartsInspectorWebsitesRead,

		Schema: map[string]*schema.Schema{
			"domain_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the domain ID.`,
			},
			"auth_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the auth status of website.`,
			},
			"top_level_domain_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Indicates the number of top level domain`,
			},
			"websites": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Indicates the websites list.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the domain ID.`,
						},
						"website_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the website name.`,
						},
						"website_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the website address.`,
						},
						"high": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `Indicates the number of high-risk vulnerabilities.`,
						},
						"middle": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `Indicates the number of medium-risk vulnerabilities.`,
						},
						"low": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `Indicates the number of low-severity vulnerabilities.`,
						},
						"hint": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `Indicates the number of hint-risk vulnerabilities.`,
						},
						"top_level_domain_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the top level domain ID.`,
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the time to create website.`,
						},
						"auth_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the auth status of website.`,
						},
					},
				},
			},
		},
	}
}

func dataSourceCodeartsInspectorWebsitesRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)
	client, err := cfg.NewServiceClient("vss", region)
	if err != nil {
		return diag.Errorf("error creating CodeArts inspector client: %s", err)
	}

	getHttpUrl := "v3/{project_id}/webscan/domains"
	getPath := client.Endpoint + getHttpUrl
	getPath = strings.ReplaceAll(getPath, "{project_id}", client.ProjectID)
	getOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	// pageLimit is `10`
	getPath += fmt.Sprintf("?limit=%v", pageLimit)
	getPath += buildCodeartsInspectorWebsitesQueryParams(d)

	currentTotal := 0

	rst := make([]map[string]interface{}, 0)
	topLevelDomainNum := 0
	for {
		currentPath := getPath + fmt.Sprintf("&offset=%d", currentTotal)
		getResp, err := client.Request("GET", currentPath, &getOpt)
		if err != nil {
			return diag.Errorf("error retrieving websites: %s", err)
		}
		getRespBody, err := utils.FlattenResponse(getResp)
		if err != nil {
			return diag.Errorf("error flatten response: %s", err)
		}

		websites := utils.PathSearch("domains", getRespBody, make([]interface{}, 0)).([]interface{})
		topLevelDomainNum = int(utils.PathSearch("top_level_domain_num", getRespBody, float64(0)).(float64))
		for _, website := range websites {
			rst = append(rst, map[string]interface{}{
				"id":                  utils.PathSearch("domain_id", website, nil),
				"high":                utils.PathSearch("high", website, nil),
				"middle":              utils.PathSearch("middle", website, nil),
				"low":                 utils.PathSearch("low", website, nil),
				"hint":                utils.PathSearch("hint", website, nil),
				"top_level_domain_id": utils.PathSearch("top_level_domain_id", website, nil),
				"created_at":          utils.PathSearch("create_time", website, nil),
				"auth_status":         utils.PathSearch("auth_status", website, nil),
				"website_name":        utils.PathSearch("alias", website, nil),
				"website_address":     flattenWebsiteAddress(website),
			})
		}

		currentTotal += len(websites)
		totalCount := utils.PathSearch("total", getRespBody, float64(0))
		if int(totalCount.(float64)) <= currentTotal {
			break
		}
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	mErr := multierror.Append(nil,
		d.Set("websites", rst),
		d.Set("top_level_domain_num", topLevelDomainNum),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func buildCodeartsInspectorWebsitesQueryParams(d *schema.ResourceData) string {
	res := ""

	if v, ok := d.GetOk("domain_id"); ok {
		res = fmt.Sprintf("%s&domain_id=%v", res, v)
	}
	if v, ok := d.GetOk("auth_status"); ok {
		res = fmt.Sprintf("%s&auth_status=%v", res, v)
	}

	return res
}
