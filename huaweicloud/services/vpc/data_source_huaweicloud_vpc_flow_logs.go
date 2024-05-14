// Generated by PMS #150
package vpc

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

func DataSourceVpcFlowLogs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVpcFlowLogsRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the VPC flow log name.`,
			},
			"flow_log_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the VPC flow log ID.`,
			},
			"resource_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the resource type for which that the logs to be collected.`,
			},
			"resource_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the resource ID for which that the logs to be collected.`,
			},
			"log_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the LTS log group ID.`,
			},
			"log_stream_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the LTS log stream ID.`,
			},
			"traffic_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the type of traffic to log.`,
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the status of the flow log.`,
			},
			"flow_logs": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The list of VPC flow logs.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The VPC flow log name.`,
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The ID of a VPC flow log`,
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The VPC flow log description.`,
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The resource type for which that the logs to be collected.`,
						},
						"resource_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The resource ID for which that the logs to be collected.`,
						},
						"log_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The LTS log group ID.`,
						},
						"log_stream_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The LTS log stream ID.`,
						},
						"traffic_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The type of traffic to log.`,
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: `Whether to enable the VPC flow log.`,
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The VPC flow log status.`,
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The time when the resource is created.`,
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The time when the resource is last updated.`,
						},
					},
				},
			},
		},
	}
}

type FlowLogsDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newFlowLogsDSWrapper(d *schema.ResourceData, meta interface{}) *FlowLogsDSWrapper {
	return &FlowLogsDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceVpcFlowLogsRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newFlowLogsDSWrapper(d, meta)
	listFlowLogsRst, err := wrapper.ListFlowLogs()
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	err = wrapper.listFlowLogsToSchema(listFlowLogsRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API VPC GET /v1/{project_id}/fl/flow_logs
func (w *FlowLogsDSWrapper) ListFlowLogs() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "vpc")
	if err != nil {
		return nil, err
	}

	uri := "/v1/{project_id}/fl/flow_logs"
	params := map[string]any{
		"id":            w.Get("flow_log_id"),
		"name":          w.Get("name"),
		"resource_type": w.Get("resource_type"),
		"resource_id":   w.Get("resource_id"),
		"traffic_type":  w.Get("traffic_type"),
		"log_group_id":  w.Get("log_group_id"),
		"log_topic_id":  w.Get("log_stream_id"),
		"status":        w.Get("status"),
	}
	params = utils.RemoveNil(params)
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		Query(params).
		MarkerPager("flow_logs", "flow_logs[*].id | [-1]", "marker").
		Request().
		Result()
}

func (w *FlowLogsDSWrapper) listFlowLogsToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("flow_logs", schemas.SliceToList(body.Get("flow_logs"),
			func(flowLog gjson.Result) any {
				return map[string]any{
					"name":          flowLog.Get("name").Value(),
					"id":            flowLog.Get("id").Value(),
					"description":   flowLog.Get("description").Value(),
					"resource_type": flowLog.Get("resource_type").Value(),
					"resource_id":   flowLog.Get("resource_id").Value(),
					"log_group_id":  flowLog.Get("log_group_id").Value(),
					"log_stream_id": flowLog.Get("log_topic_id").Value(),
					"traffic_type":  flowLog.Get("traffic_type").Value(),
					"enabled":       flowLog.Get("admin_state").Value(),
					"status":        flowLog.Get("status").Value(),
					"created_at":    flowLog.Get("created_at").Value(),
					"updated_at":    flowLog.Get("updated_at").Value(),
				}
			},
		)),
	)
	return mErr.ErrorOrNil()
}
