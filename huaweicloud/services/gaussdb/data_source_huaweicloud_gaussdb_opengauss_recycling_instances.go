// Generated by PMS #513
package gaussdb

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tidwall/gjson"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/httphelper"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/schemas"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func DataSourceGaussdbOpengaussRecyclingInstances() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGaussdbOpengaussRecyclingInstancesRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the GaussDB OpenGauss instance name.`,
			},
			"instances": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Indicates the information about all instances in the recycle bin.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the instance ID.`,
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the instance name.`,
						},
						"mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the product type.`,
						},
						"ha_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the deployment model.`,
						},
						"engine_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the engine name.`,
						},
						"engine_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the engine version.`,
						},
						"pay_model": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the billing mode.`,
						},
						"volume_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the disk type.`,
						},
						"volume_size": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the disk size.`,
						},
						"enterprise_project_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the enterprise project ID.`,
						},
						"enterprise_project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the enterprise project name.`,
						},
						"recycle_backup_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the backup ID.`,
						},
						"backup_level": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the backup level.`,
						},
						"data_vip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the private IP address.`,
						},
						"recycle_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the backup status in the recycle bin.`,
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the creation time in the **yyyy-mm-ddThh:mm:ssZ** format.`,
						},
						"deleted_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the deletion time in the **yyyy-mm-ddThh:mm:ssZ** format.`,
						},
					},
				},
			},
		},
	}
}

type OpengaussRecyclingInstancesDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newOpengaussRecyclingInstancesDSWrapper(d *schema.ResourceData, meta interface{}) *OpengaussRecyclingInstancesDSWrapper {
	return &OpengaussRecyclingInstancesDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceGaussdbOpengaussRecyclingInstancesRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newOpengaussRecyclingInstancesDSWrapper(d, meta)
	lisRecInsDetRst, err := wrapper.ListRecycleInstancesDetails()
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	err = wrapper.listRecycleInstancesDetailsToSchema(lisRecInsDetRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API GaussDB GET /v3.1/{project_id}/recycle-instances
func (w *OpengaussRecyclingInstancesDSWrapper) ListRecycleInstancesDetails() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "opengauss")
	if err != nil {
		return nil, err
	}

	uri := "/v3.1/{project_id}/recycle-instances"
	params := map[string]any{
		"instance_name": w.Get("instance_name"),
	}
	params = utils.RemoveNil(params)
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		Query(params).
		Request().
		Result()
}

func (w *OpengaussRecyclingInstancesDSWrapper) listRecycleInstancesDetailsToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("instances", schemas.SliceToList(body.Get("instances"),
			func(instances gjson.Result) any {
				return map[string]any{
					"id":                      instances.Get("id").Value(),
					"name":                    instances.Get("name").Value(),
					"mode":                    instances.Get("mode").Value(),
					"ha_mode":                 instances.Get("ha_mode").Value(),
					"engine_name":             instances.Get("engine_name").Value(),
					"engine_version":          instances.Get("engine_version").Value(),
					"pay_model":               instances.Get("pay_model").Value(),
					"volume_type":             instances.Get("volume_type").Value(),
					"volume_size":             instances.Get("volume_size").Value(),
					"enterprise_project_id":   instances.Get("enterprise_project_id").Value(),
					"enterprise_project_name": instances.Get("enterprise_project_name").Value(),
					"recycle_backup_id":       instances.Get("recycle_backup_id").Value(),
					"backup_level":            instances.Get("backup_level").Value(),
					"data_vip":                instances.Get("data_vip").Value(),
					"recycle_status":          instances.Get("recycle_status").Value(),
					"created_at":              instances.Get("created_at").Value(),
					"deleted_at":              instances.Get("deleted_at").Value(),
				}
			},
		)),
	)
	return mErr.ErrorOrNil()
}
