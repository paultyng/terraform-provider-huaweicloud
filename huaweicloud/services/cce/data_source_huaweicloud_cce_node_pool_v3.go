package cce

import (
	"context"
	"strconv"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk/openstack/cce/v3/nodepools"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils/fmtp"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils/logp"
)

// @API CCE GET /api/v3/projects/{project_id}/clusters/{clusterid}/nodepools
func DataSourceCCENodePoolV3() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCceNodePoolsV3Read,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"node_pool_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"initial_node_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"current_node_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"flavor_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"labels": { //(k8s_tags)
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"root_volume": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"volumetype": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"extend_params": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					}},
			},
			"data_volumes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"volumetype": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"extend_params": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					}},
			},
			"max_pods": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"extend_param": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"availability_zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_pair": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": common.TagsComputedSchema(),
			"subnet_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scale_enable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"min_node_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_node_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"scale_down_cooldown_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"hostname_config": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"enterprise_project_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCceNodePoolsV3Read(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	cceClient, err := config.CceV3Client(config.GetRegion(d))
	if err != nil {
		return fmtp.DiagErrorf("unable to create HuaweiCloud CCE client : %s", err)
	}

	listOpts := nodepools.ListOpts{
		Uid:   d.Get("node_pool_id").(string),
		Name:  d.Get("name").(string),
		Phase: d.Get("status").(string),
	}

	refinedNodePools, err := nodepools.List(cceClient, d.Get("cluster_id").(string), listOpts)

	if err != nil {
		return fmtp.DiagErrorf("unable to retrieve Node Pools: %s", err)
	}

	if len(refinedNodePools) < 1 {
		return fmtp.DiagErrorf("your query returned no results, please change your search criteria and try again")
	}

	if len(refinedNodePools) > 1 {
		return fmtp.DiagErrorf("your query returned more than one result, please try a more specific search criteria")
	}

	NodePool := refinedNodePools[0]

	logp.Printf("[DEBUG] Retrieved Node Pools using given filter %s: %+v", NodePool.Metadata.Id, NodePool)

	d.SetId(NodePool.Metadata.Id)
	mErr := multierror.Append(nil,
		d.Set("region", config.GetRegion(d)),
		d.Set("node_pool_id", NodePool.Metadata.Id),
		d.Set("name", NodePool.Metadata.Name),
		d.Set("type", NodePool.Spec.Type),
		d.Set("flavor_id", NodePool.Spec.NodeTemplate.Flavor),
		d.Set("availability_zone", NodePool.Spec.NodeTemplate.Az),
		d.Set("os", NodePool.Spec.NodeTemplate.Os),
		d.Set("key_pair", NodePool.Spec.NodeTemplate.Login.SshKey),
		d.Set("scale_enable", NodePool.Spec.Autoscaling.Enable),
		d.Set("initial_node_count", NodePool.Spec.InitialNodeCount),
		d.Set("current_node_count", NodePool.Status.CurrentNode),
		d.Set("min_node_count", NodePool.Spec.Autoscaling.MinNodeCount),
		d.Set("max_node_count", NodePool.Spec.Autoscaling.MaxNodeCount),
		d.Set("scale_down_cooldown_time", NodePool.Spec.Autoscaling.ScaleDownCooldownTime),
		d.Set("priority", NodePool.Spec.Autoscaling.Priority),
		d.Set("subnet_id", NodePool.Spec.NodeTemplate.NodeNicSpec.PrimaryNic.SubnetId),
		d.Set("status", NodePool.Status.Phase),
		d.Set("hostname_config", flattenResourceNodeHostnameConfig(NodePool.Spec.NodeTemplate.HostnameConfig)),
		d.Set("enterprise_project_id", NodePool.Spec.NodeTemplate.ServerEnterpriseProjectID),
	)

	// set extend_param
	var extendParam = NodePool.Spec.NodeTemplate.ExtendParam
	mErr = multierror.Append(mErr, d.Set("max_pods", extendParam["maxPods"]))
	delete(extendParam, "maxPods")

	extendParamToSet := map[string]string{}
	for k, v := range extendParam {
		switch v := v.(type) {
		case string:
			extendParamToSet[k] = v
		case int:
			extendParamToSet[k] = strconv.Itoa(v)
		case int32:
			extendParamToSet[k] = strconv.FormatInt(int64(v), 10)
		case float64:
			extendParamToSet[k] = strconv.FormatFloat(v, 'f', -1, 64)
		case bool:
			extendParamToSet[k] = strconv.FormatBool(v)
		default:
			logp.Printf("[WARN] can not set %s to extend_param, the value is %v", k, v)
		}
	}

	mErr = multierror.Append(mErr, d.Set("extend_param", extendParamToSet))

	// set labels
	labels := map[string]string{}
	for key, val := range NodePool.Spec.NodeTemplate.K8sTags {
		if strings.Contains(key, "cce.cloud.com") {
			continue
		}
		labels[key] = val
	}
	mErr = multierror.Append(mErr, d.Set("labels", labels))

	// set data volumes
	volumes := make([]map[string]interface{}, 0, len(NodePool.Spec.NodeTemplate.DataVolumes))
	for _, pairObject := range NodePool.Spec.NodeTemplate.DataVolumes {
		volume := make(map[string]interface{})
		volume["size"] = pairObject.Size
		volume["volumetype"] = pairObject.VolumeType
		volume["extend_params"] = pairObject.ExtendParam
		volumes = append(volumes, volume)
	}
	mErr = multierror.Append(mErr, d.Set("data_volumes", volumes))

	// set root volume
	rootVolume := []map[string]interface{}{
		{
			"size":          NodePool.Spec.NodeTemplate.RootVolume.Size,
			"volumetype":    NodePool.Spec.NodeTemplate.RootVolume.VolumeType,
			"extend_params": NodePool.Spec.NodeTemplate.RootVolume.ExtendParam,
		},
	}
	mErr = multierror.Append(mErr, d.Set("root_volume", rootVolume))

	// set tags
	tagmap := utils.TagsToMap(NodePool.Spec.NodeTemplate.UserTags)
	mErr = multierror.Append(mErr, d.Set("tags", tagmap))

	if err = mErr.ErrorOrNil(); err != nil {
		return fmtp.DiagErrorf("Error setting node pool fields: %s", err)
	}

	return nil
}
