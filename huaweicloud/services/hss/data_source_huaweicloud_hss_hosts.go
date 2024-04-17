package hss

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	hssv5model "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/hss/v5/model"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API HSS GET /v5/{project_id}/host-management/hosts
func DataSourceHosts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHostsRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"agent_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protect_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protect_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protect_charging_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"detect_result": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"asset_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enterprise_project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hosts": {
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
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"agent_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"agent_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protect_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protect_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protect_charging_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quota_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"detect_result": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"policy_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"asset_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"open_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"private_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"asset_risk_num": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vulnerability_risk_num": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"baseline_risk_num": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"intrusion_risk_num": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"enterprise_project_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceHostsRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		cfg      = meta.(*config.Config)
		region   = cfg.GetRegion(d)
		epsId    = cfg.DataGetEnterpriseProjectID(d)
		limit    = int32(20)
		offset   int32
		allHosts []hssv5model.Host
	)

	client, err := cfg.HcHssV5Client(region)
	if err != nil {
		return diag.Errorf("error creating HSS v5 client: %s", err)
	}

	for {
		request := hssv5model.ListHostStatusRequest{
			Region:              &region,
			Limit:               utils.Int32(limit),
			Offset:              utils.Int32(offset),
			HostId:              utils.StringIgnoreEmpty(d.Get("host_id").(string)),
			HostName:            utils.StringIgnoreEmpty(d.Get("name").(string)),
			HostStatus:          utils.StringIgnoreEmpty(d.Get("status").(string)),
			OsType:              utils.StringIgnoreEmpty(d.Get("os_type").(string)),
			AgentStatus:         utils.StringIgnoreEmpty(d.Get("agent_status").(string)),
			ProtectStatus:       utils.StringIgnoreEmpty(d.Get("protect_status").(string)),
			Version:             utils.StringIgnoreEmpty(d.Get("protect_version").(string)),
			ChargingMode:        utils.StringIgnoreEmpty(convertChargingModeRequest(d.Get("protect_charging_mode").(string))),
			DetectResult:        utils.StringIgnoreEmpty(d.Get("detect_result").(string)),
			GroupId:             utils.StringIgnoreEmpty(d.Get("group_id").(string)),
			PolicyGroupId:       utils.StringIgnoreEmpty(d.Get("policy_group_id").(string)),
			AssetValue:          utils.StringIgnoreEmpty(d.Get("asset_value").(string)),
			EnterpriseProjectId: utils.String(epsId),
		}

		listResp, listErr := client.ListHostStatus(&request)
		if listErr != nil {
			return diag.Errorf("error querying HSS hosts: %s", listErr)
		}

		if listResp == nil || listResp.DataList == nil {
			break
		}
		if len(*listResp.DataList) == 0 {
			break
		}

		allHosts = append(allHosts, *listResp.DataList...)
		offset += limit
	}

	uuId, err := uuid.GenerateUUID()
	if err != nil {
		return diag.Errorf("unable to generate ID: %s", err)
	}

	d.SetId(uuId)

	mErr := multierror.Append(nil,
		d.Set("region", region),
		d.Set("hosts", flattenHosts(allHosts)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func flattenHosts(hosts []hssv5model.Host) []interface{} {
	if len(hosts) == 0 {
		return nil
	}

	rst := make([]interface{}, 0, len(hosts))
	for _, v := range hosts {
		rst = append(rst, map[string]interface{}{
			"id":                     v.HostId,
			"name":                   v.HostName,
			"status":                 v.HostStatus,
			"os_type":                v.OsType,
			"agent_id":               v.AgentId,
			"agent_status":           v.AgentStatus,
			"protect_status":         v.ProtectStatus,
			"protect_version":        v.Version,
			"protect_charging_mode":  convertChargingMode(v.ChargingMode),
			"quota_id":               v.ResourceId,
			"detect_result":          v.DetectResult,
			"group_id":               v.GroupId,
			"policy_group_id":        v.PolicyGroupId,
			"asset_value":            v.AssetValue,
			"open_time":              convertOpenTime(v.OpenTime),
			"private_ip":             v.PrivateIp,
			"public_ip":              v.PublicIp,
			"asset_risk_num":         v.Asset,
			"vulnerability_risk_num": v.Vulnerability,
			"baseline_risk_num":      v.Baseline,
			"intrusion_risk_num":     v.Intrusion,
			"enterprise_project_id":  v.EnterpriseProjectId,
		})
	}

	return rst
}
