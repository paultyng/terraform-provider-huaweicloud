package iotda

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

// @API IoTDA GET /v5/iot/{project_id}/device-group
func DataSourceDeviceGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeviceGroupsRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"space_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func buildDeviceGroupsQueryParams(d *schema.ResourceData) string {
	req := ""

	if v, ok := d.GetOk("name"); ok {
		req = fmt.Sprintf("%s&name=%v", req, v)
	}

	if v, ok := d.GetOk("type"); ok {
		req = fmt.Sprintf("%s&group_type=%v", req, v)
	}

	if v, ok := d.GetOk("space_id"); ok {
		return fmt.Sprintf("%s&app_id=%v", req, v)
	}

	return req
}

func dataSourceDeviceGroupsRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg       = meta.(*config.Config)
		region    = cfg.GetRegion(d)
		isDerived = WithDerivedAuth(cfg, region)
		httpUrl   = "v5/iot/{project_id}/device-group?limit=50"
		allGroups = make([]interface{}, 0)
		offset    = 0
	)

	client, err := cfg.NewServiceClientWithDerivedAuth("iotda", region, isDerived)
	if err != nil {
		return diag.Errorf("error creating IoTDA client: %s", err)
	}

	listPath := client.Endpoint + httpUrl
	listPath = strings.ReplaceAll(listPath, "{project_id}", client.ProjectID)
	listPath += buildDeviceGroupsQueryParams(d)
	listOpts := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	for {
		listPathWithOffset := fmt.Sprintf("%s&offset=%d", listPath, offset)
		resp, err := client.Request("GET", listPathWithOffset, &listOpts)
		if err != nil {
			return diag.Errorf("error querying IoTDA device groups: %s", err)
		}

		respBody, err := utils.FlattenResponse(resp)
		if err != nil {
			return diag.FromErr(err)
		}

		groups := utils.PathSearch("device_groups", respBody, make([]interface{}, 0)).([]interface{})
		if len(groups) == 0 {
			break
		}

		allGroups = append(allGroups, groups...)
		offset += len(groups)
	}

	uuId, err := uuid.GenerateUUID()
	if err != nil {
		return diag.Errorf("unable to generate ID: %s", err)
	}

	d.SetId(uuId)

	mErr := multierror.Append(nil,
		d.Set("region", region),
		d.Set("groups", flattenDeviceGroups(filterListDeviceGroups(allGroups, d))),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func filterListDeviceGroups(deviceGroups []interface{}, d *schema.ResourceData) []interface{} {
	if len(deviceGroups) == 0 {
		return nil
	}

	rst := make([]interface{}, 0, len(deviceGroups))
	for _, v := range deviceGroups {
		if deviceGroupId, ok := d.GetOk("group_id"); ok &&
			fmt.Sprint(deviceGroupId) != utils.PathSearch("group_id", v, "").(string) {
			continue
		}

		if parentGroupId, ok := d.GetOk("parent_group_id"); ok &&
			fmt.Sprint(parentGroupId) != utils.PathSearch("super_group_id", v, "").(string) {
			continue
		}

		rst = append(rst, v)
	}

	return rst
}

func flattenDeviceGroups(deviceGroups []interface{}) []interface{} {
	if len(deviceGroups) == 0 {
		return nil
	}

	rst := make([]interface{}, 0, len(deviceGroups))
	for _, v := range deviceGroups {
		rst = append(rst, map[string]interface{}{
			"id":              utils.PathSearch("group_id", v, nil),
			"name":            utils.PathSearch("name", v, nil),
			"description":     utils.PathSearch("description", v, nil),
			"parent_group_id": utils.PathSearch("super_group_id", v, nil),
			"type":            utils.PathSearch("group_type", v, nil),
		})
	}

	return rst
}
