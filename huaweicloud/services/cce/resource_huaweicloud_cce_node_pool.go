package cce

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/openstack/cce/v3/nodepools"
	"github.com/chnsz/golangsdk/openstack/cce/v3/nodes"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API CCE GET /api/v3/projects/{project_id}/clusters/{cluster_id}
// @API CCE POST /api/v3/projects/{project_id}/clusters/{cluster_id}/nodepools
// @API CCE GET /api/v3/projects/{project_id}/clusters/{cluster_id}/nodepools/{nodepool_id}
// @API CCE PUT /api/v3/projects/{project_id}/clusters/{cluster_id}/nodepools/{nodepool_id}
// @API CCE DELETE /api/v3/projects/{project_id}/clusters/{cluster_id}/nodepools/{nodepool_id}

func ResourceNodePool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNodePoolCreate,
		ReadContext:   resourceNodePoolRead,
		UpdateContext: resourceNodePoolUpdate,
		DeleteContext: resourceNodePoolDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceNodePoolImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"initial_node_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"flavor_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"labels": { // (k8s_tags)
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"root_volume":  resourceNodeRootVolume(),
			"data_volumes": resourceNodeDataVolume(),
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "random",
			},
			"os": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_pair": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"password", "key_pair"},
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"storage": resourceNodeStorageSchema(),
			"taints": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"effect": {
							Type:     schema.TypeString,
							Required: true,
						},
					}},
			},
			"tags": common.TagsSchema(),
			// charge info: charging_mode, period_unit, period, auto_renew
			"charging_mode": common.SchemaChargingMode(nil),
			"period_unit":   common.SchemaPeriodUnit(nil),
			"period":        common.SchemaPeriod(nil),
			"auto_renew":    common.SchemaAutoRenew(nil),

			"runtime": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"docker", "containerd",
				}, false),
			},
			"extend_params": resourceNodeExtendParamsSchema([]string{
				"max_pods", "preinstall", "postinstall", "extend_param",
			}),
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"scale_enable": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"min_node_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_node_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scale_down_cooldown_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"security_groups": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pod_security_groups": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ecs_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"initialized_conditions": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"label_policy_on_existing_nodes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag_policy_on_existing_nodes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"taint_policy_on_existing_nodes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"enterprise_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extension_scale_groups": resourceExtensionScaleGroupsSchema(),
			"current_node_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"billing_mode": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			// Deprecated parameters
			"max_pods": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "schema: Deprecated; This parameter can be configured in the 'extend_params' parameter.",
			},
			"preinstall": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				StateFunc:   utils.DecodeHashAndHexEncode,
				Description: "schema: Deprecated; This parameter can be configured in the 'extend_params' parameter.",
			},
			"postinstall": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				StateFunc:   utils.DecodeHashAndHexEncode,
				Description: "schema: Deprecated; This parameter can be configured in the 'extend_params' parameter.",
			},
			"extend_param": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "schema: Deprecated; This parameter has been replaced by the 'extend_params' parameter.",
			},
		},
	}
}

func resourceExtensionScaleGroupsSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"metadata": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"name": {
								Type:     schema.TypeString,
								Optional: true,
							},
						},
					},
				},
				"spec": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"flavor": {
								Type:     schema.TypeString,
								Optional: true,
							},
							"az": {
								Type:     schema.TypeString,
								Optional: true,
							},
							"capacity_reservation_specification": {
								Type:     schema.TypeList,
								Optional: true,
								MaxItems: 1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"id": {
											Type:     schema.TypeString,
											Optional: true,
										},
										"preference": {
											Type:     schema.TypeString,
											Optional: true,
										},
									},
								},
							},
							"autoscaling": {
								Type:     schema.TypeList,
								Optional: true,
								MaxItems: 1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"enable": {
											Type:     schema.TypeBool,
											Optional: true,
										},
										"extension_priority": {
											Type:     schema.TypeInt,
											Optional: true,
										},
										"min_node_count": {
											Type:     schema.TypeInt,
											Optional: true,
										},
										"max_node_count": {
											Type:     schema.TypeInt,
											Optional: true,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func buildPodSecurityGroups(ids []interface{}) []nodepools.PodSecurityGroupSpec {
	if len(ids) == 0 {
		return nil
	}

	groups := make([]nodepools.PodSecurityGroupSpec, len(ids))
	for i, id := range ids {
		groups[i] = nodepools.PodSecurityGroupSpec{Id: id.(string)}
	}

	return groups
}

func buildExtensionScaleGroups(d *schema.ResourceData) []nodepools.ExtensionScaleGroups {
	if v, ok := d.GetOk("extension_scale_groups"); ok {
		groups := v.([]interface{})
		res := make([]nodepools.ExtensionScaleGroups, len(groups))
		for i, v := range groups {
			if group, ok := v.(map[string]interface{}); ok {
				res[i] = nodepools.ExtensionScaleGroups{
					Metadata: buildExtensionScaleGroupsMetadata(utils.PathSearch("metadata", group, nil)),
					Spec:     buildExtensionScaleGroupsSpec(utils.PathSearch("spec", group, nil)),
				}
			}
		}

		return res
	}

	return nil
}

func buildExtensionScaleGroupsMetadata(metadata interface{}) *nodepools.ExtensionScaleGroupsMetadata {
	if len(metadata.([]interface{})) == 0 {
		return nil
	}

	res := nodepools.ExtensionScaleGroupsMetadata{
		Uid:  utils.PathSearch("[0].uid", metadata, "").(string),
		Name: utils.PathSearch("[0].name", metadata, "").(string),
	}

	return &res
}

func buildExtensionScaleGroupsSpec(spec interface{}) *nodepools.ExtensionScaleGroupsSpec {
	if len(spec.([]interface{})) == 0 {
		return nil
	}

	res := nodepools.ExtensionScaleGroupsSpec{
		Flavor:      utils.PathSearch("[0].flavor", spec, "").(string),
		Az:          utils.PathSearch("[0].az", spec, "").(string),
		Autoscaling: buildAutoscaling(utils.PathSearch("[0].autoscaling", spec, nil)),
		CapacityReservationSpecification: buildCapacityReservationSpecification(
			utils.PathSearch("[0].capacity_reservation_specification", spec, nil)),
	}

	return &res
}

func buildCapacityReservationSpecification(capacityReservationSpecification interface{}) *nodepools.CapacityReservationSpecification {
	if len(capacityReservationSpecification.([]interface{})) == 0 {
		return nil
	}

	res := nodepools.CapacityReservationSpecification{
		ID:         utils.PathSearch("[0].id", capacityReservationSpecification, "").(string),
		Preference: utils.PathSearch("[0].preference", capacityReservationSpecification, "").(string),
	}

	return &res
}

func buildAutoscaling(autoscaling interface{}) *nodepools.Autoscaling {
	if len(autoscaling.([]interface{})) == 0 {
		return nil
	}

	res := nodepools.Autoscaling{
		Enable:            utils.PathSearch("[0].enable", autoscaling, false).(bool),
		ExtensionPriority: utils.PathSearch("[0].extension_priority", autoscaling, 0).(int),
		MinNodeCount:      utils.PathSearch("[0].min_node_count", autoscaling, 0).(int),
		MaxNodeCount:      utils.PathSearch("[0].max_node_count", autoscaling, 0).(int),
	}

	return &res
}

func buildResourceNodePoolNicSpec(d *schema.ResourceData) nodes.NodeNicSpec {
	res := nodes.NodeNicSpec{
		PrimaryNic: nodes.PrimaryNic{
			SubnetId:   d.Get("subnet_id").(string),
			SubnetList: utils.ExpandToStringList(d.Get("subnet_list").([]interface{})),
		},
	}

	return res
}

func buildNodePoolCreateOpts(d *schema.ResourceData, cfg *config.Config) (*nodepools.CreateOpts, error) {
	// Validate whether prepaid parameters are configured.
	billingMode := 0
	if d.Get("charging_mode").(string) == "prePaid" {
		if err := common.ValidatePrePaidChargeInfo(d); err != nil {
			return nil, err
		}
		billingMode = 1
	}

	createOpts := nodepools.CreateOpts{
		Kind:       "NodePool",
		ApiVersion: "v3",
		Metadata: nodepools.CreateMetaData{
			Name: d.Get("name").(string),
		},
		Spec: nodepools.CreateSpec{
			Type: d.Get("type").(string),
			NodeTemplate: nodes.Spec{
				Flavor:                    d.Get("flavor_id").(string),
				Az:                        d.Get("availability_zone").(string),
				Os:                        d.Get("os").(string),
				RootVolume:                buildResourceNodeRootVolume(d),
				DataVolumes:               buildResourceNodeDataVolume(d),
				Storage:                   buildResourceNodeStorage(d),
				K8sTags:                   buildResourceNodeK8sTags(d),
				BillingMode:               billingMode,
				Count:                     1,
				NodeNicSpec:               buildResourceNodePoolNicSpec(d),
				ExtendParam:               buildExtendParams(d),
				Taints:                    buildResourceNodeTaint(d),
				UserTags:                  utils.ExpandResourceTags(d.Get("tags").(map[string]interface{})),
				InitializedConditions:     utils.ExpandToStringList(d.Get("initialized_conditions").([]interface{})),
				HostnameConfig:            buildResourceNodeHostnameConfig(d),
				ServerEnterpriseProjectID: cfg.GetEnterpriseProjectID(d),
			},
			Autoscaling: nodepools.AutoscalingSpec{
				Enable:                d.Get("scale_enable").(bool),
				MinNodeCount:          d.Get("min_node_count").(int),
				MaxNodeCount:          d.Get("max_node_count").(int),
				ScaleDownCooldownTime: d.Get("scale_down_cooldown_time").(int),
				Priority:              d.Get("priority").(int),
			},
			InitialNodeCount:     utils.Int(d.Get("initial_node_count").(int)),
			PodSecurityGroups:    buildPodSecurityGroups(d.Get("pod_security_groups").([]interface{})),
			CustomSecurityGroups: utils.ExpandToStringList(d.Get("security_groups").([]interface{})),
			NodeManagement: nodepools.NodeManagementSpec{
				ServerGroupReference: d.Get("ecs_group_id").(string),
			},
			LabelPolicyOnExistingNodes:   d.Get("label_policy_on_existing_nodes").(string),
			UserTagPolicyOnExistingNodes: d.Get("tag_policy_on_existing_nodes").(string),
			TaintPolicyOnExistingNodes:   d.Get("taint_policy_on_existing_nodes").(string),
			ExtensionScaleGroups:         buildExtensionScaleGroups(d),
		},
	}

	if v, ok := d.GetOk("runtime"); ok {
		createOpts.Spec.NodeTemplate.RunTime = &nodes.RunTimeSpec{
			Name: v.(string),
		}
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	// Add loginSpec here, so it wouldn't go in the above log entry
	loginSpec, err := buildResourceNodeLoginSpec(d)
	if err != nil {
		return nil, err
	}
	createOpts.Spec.NodeTemplate.Login = loginSpec
	return &createOpts, nil
}

func resourceNodePoolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	cceClient, err := cfg.CceV3Client(cfg.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating CCE v3 client: %s", err)
	}

	// Wait for the cce cluster to become available.
	clusterId := d.Get("cluster_id").(string)
	stateCluster := &resource.StateChangeConf{
		Pending:    []string{"PENDING"},
		Target:     []string{"COMPLETED"},
		Refresh:    clusterStateRefreshFunc(cceClient, clusterId, []string{"Available"}),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 5 * time.Second,
	}
	_, err = stateCluster.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("error waiting for CCE cluster to be available: %s", err)
	}

	createOpts, err := buildNodePoolCreateOpts(d, cfg)
	if err != nil {
		return diag.Errorf("error creating CreateOpts structure of 'Create' method for CCE node pool: %s", err)
	}
	resp, err := nodepools.Create(cceClient, clusterId, createOpts).Extract()
	if err != nil {
		return diag.Errorf("error creating CCE node pool: %s", err)
	}
	log.Printf("[DEBUG] The response of the 'Create' method for CCE node pool is: %v", resp)
	if resp.Metadata.Id == "" {
		return diag.Errorf("error fetching resource ID from the API response of CCE node pool")
	}
	d.SetId(resp.Metadata.Id)

	stateConf := &resource.StateChangeConf{
		// The statuses of pending phase includes "Synchronizing" and "Synchronized".
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      nodePoolStateRefreshFunc(cceClient, clusterId, d.Id(), []string{""}),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        120 * time.Second,
		PollInterval: 20 * time.Second,
	}
	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("error waiting for CCE node pool to become available: %s", err)
	}

	return resourceNodePoolRead(ctx, d, meta)
}

func resourceNodePoolRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)
	cceClient, err := cfg.CceV3Client(region)
	if err != nil {
		return diag.Errorf("error creating CCE v3 client: %s", err)
	}
	clusterId := d.Get("cluster_id").(string)
	s, err := nodepools.Get(cceClient, clusterId, d.Id()).Extract()

	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving CCE node pool")
	}

	// The following parameters are not returned:
	// password, extend_params, taints, initial_node_count, pod_security_groups
	// extension_scale_groups not save, because the order of groups will change and computed not working in TypeSet
	mErr := multierror.Append(nil,
		d.Set("region", region),
		d.Set("name", s.Metadata.Name),
		d.Set("flavor_id", s.Spec.NodeTemplate.Flavor),
		d.Set("availability_zone", s.Spec.NodeTemplate.Az),
		d.Set("os", s.Spec.NodeTemplate.Os),
		d.Set("billing_mode", s.Spec.NodeTemplate.BillingMode),
		d.Set("key_pair", s.Spec.NodeTemplate.Login.SshKey),
		d.Set("scale_enable", s.Spec.Autoscaling.Enable),
		d.Set("min_node_count", s.Spec.Autoscaling.MinNodeCount),
		d.Set("max_node_count", s.Spec.Autoscaling.MaxNodeCount),
		d.Set("current_node_count", s.Status.CurrentNode),
		d.Set("scale_down_cooldown_time", s.Spec.Autoscaling.ScaleDownCooldownTime),
		d.Set("priority", s.Spec.Autoscaling.Priority),
		d.Set("type", s.Spec.Type),
		d.Set("ecs_group_id", s.Spec.NodeManagement.ServerGroupReference),
		d.Set("storage", flattenResourceNodeStorage(s.Spec.NodeTemplate.Storage)),
		d.Set("security_groups", s.Spec.CustomSecurityGroups),
		d.Set("tags", utils.TagsToMap(s.Spec.NodeTemplate.UserTags)),
		d.Set("status", s.Status.Phase),
		d.Set("data_volumes", flattenResourceNodeDataVolume(d, s.Spec.NodeTemplate.DataVolumes)),
		d.Set("root_volume", flattenResourceNodeRootVolume(d, s.Spec.NodeTemplate.RootVolume)),
		d.Set("initialized_conditions", s.Spec.NodeTemplate.InitializedConditions),
		d.Set("label_policy_on_existing_nodes", s.Spec.LabelPolicyOnExistingNodes),
		d.Set("tag_policy_on_existing_nodes", s.Spec.UserTagPolicyOnExistingNodes),
		d.Set("taint_policy_on_existing_nodes", s.Spec.TaintPolicyOnExistingNodes),
		d.Set("hostname_config", flattenResourceNodeHostnameConfig(s.Spec.NodeTemplate.HostnameConfig)),
		d.Set("enterprise_project_id", s.Spec.NodeTemplate.ServerEnterpriseProjectID),
		d.Set("subnet_id", s.Spec.NodeTemplate.NodeNicSpec.PrimaryNic.SubnetId),
		d.Set("subnet_list", s.Spec.NodeTemplate.NodeNicSpec.PrimaryNic.SubnetList),
	)

	if s.Spec.NodeTemplate.BillingMode != 0 {
		mErr = multierror.Append(mErr, d.Set("charging_mode", "prePaid"))
	}

	if s.Spec.NodeTemplate.RunTime != nil {
		mErr = multierror.Append(mErr, d.Set("runtime", s.Spec.NodeTemplate.RunTime.Name))
	}

	labels := make(map[string]interface{})
	for key, val := range s.Spec.NodeTemplate.K8sTags {
		if strings.Contains(key, "cce.cloud.com") {
			continue
		}
		labels[key] = val
	}
	mErr = multierror.Append(mErr, d.Set("labels", labels))
	if err = mErr.ErrorOrNil(); err != nil {
		return diag.Errorf("error setting CCE node pool fields: %s", err)
	}
	return nil
}

func buildNodePoolUpdateOpts(d *schema.ResourceData, cfg *config.Config) (*nodepools.UpdateOpts, error) {
	updateOpts := nodepools.UpdateOpts{
		Metadata: nodepools.UpdateMetaData{
			Name: d.Get("name").(string),
		},
		Spec: nodepools.UpdateSpec{
			InitialNodeCount: utils.Int(d.Get("initial_node_count").(int)),
			Autoscaling: nodepools.AutoscalingSpec{
				Enable:                d.Get("scale_enable").(bool),
				MinNodeCount:          d.Get("min_node_count").(int),
				MaxNodeCount:          d.Get("max_node_count").(int),
				ScaleDownCooldownTime: d.Get("scale_down_cooldown_time").(int),
				Priority:              d.Get("priority").(int),
			},
			NodeTemplate: nodepools.UpdateNodeTemplate{
				UserTags:                  utils.ExpandResourceTags(d.Get("tags").(map[string]interface{})),
				K8sTags:                   buildResourceNodeK8sTags(d),
				Taints:                    buildResourceNodeTaint(d),
				InitializedConditions:     utils.ExpandToStringList(d.Get("initialized_conditions").([]interface{})),
				ServerEnterpriseProjectID: cfg.GetEnterpriseProjectID(d),
				Os:                        d.Get("os").(string),
				NodeNicSpecUpdate:         buildResourceNodePoolNicSpec(d),
			},
			LabelPolicyOnExistingNodes:   d.Get("label_policy_on_existing_nodes").(string),
			UserTagPolicyOnExistingNodes: d.Get("tag_policy_on_existing_nodes").(string),
			TaintPolicyOnExistingNodes:   d.Get("taint_policy_on_existing_nodes").(string),
			ExtensionScaleGroups:         buildExtensionScaleGroups(d),
		},
	}
	return &updateOpts, nil
}

func resourceNodePoolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	cceClient, err := cfg.CceV3Client(cfg.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating CCE v3 client: %s", err)
	}

	updateOpts, err := buildNodePoolUpdateOpts(d, cfg)
	if err != nil {
		return diag.FromErr(err)
	}
	clusterId := d.Get("cluster_id").(string)
	nodePoolId := d.Id()
	_, err = nodepools.Update(cceClient, clusterId, nodePoolId, updateOpts).Extract()
	if err != nil {
		return diag.Errorf("error updating CCE node pool (%s): %s", nodePoolId, err)
	}

	stateConf := &resource.StateChangeConf{
		// The statuses of pending phase includes "Synchronizing" and "Synchronized".
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      nodePoolStateRefreshFunc(cceClient, clusterId, nodePoolId, []string{""}),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        60 * time.Second,
		PollInterval: 10 * time.Second,
	}
	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("error waiting for CCE node pool (%s) to become available: %s", nodePoolId, err)
	}

	return resourceNodePoolRead(ctx, d, meta)
}

func resourceNodePoolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	cceClient, err := cfg.CceV3Client(cfg.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating CCE v3 client: %s", err)
	}

	clusterId := d.Get("cluster_id").(string)
	nodePoolId := d.Id()
	err = nodepools.Delete(cceClient, clusterId, nodePoolId).ExtractErr()
	if err != nil {
		return diag.Errorf("error deleting CCE node pool (%s): %s", nodePoolId, err)
	}

	stateConf := &resource.StateChangeConf{
		// The statuses of pending phase include "Deleting".
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      nodePoolStateRefreshFunc(cceClient, clusterId, nodePoolId, nil),
		Timeout:      d.Timeout(schema.TimeoutDelete),
		Delay:        60 * time.Second,
		PollInterval: 20 * time.Second,
	}
	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("error waiting for CCE node pool (%s) to become deleted: %s", nodePoolId, err)
	}
	return nil
}

func nodePoolStateRefreshFunc(cceClient *golangsdk.ServiceClient, clusterId, nodePoolId string,
	targets []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Expect the status of CCE add-on to be any one of the status list: %v.", targets)
		resp, err := nodepools.Get(cceClient, clusterId, nodePoolId).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				log.Printf("[DEBUG] The node pool (%s) has been deleted", nodePoolId)
				return resp, "COMPLETED", nil
			}
			return nil, "ERROR", err
		}

		invalidStatuses := []string{"Error", "Shelved", "Unknow"}
		if utils.IsStrContainsSliceElement(resp.Status.Phase, invalidStatuses, true, true) {
			return resp, "", fmt.Errorf("unexpect status (%s)", resp.Status.Phase)
		}

		if utils.StrSliceContains(targets, resp.Status.Phase) {
			return resp, "COMPLETED", nil
		}
		return resp, "PENDING", nil
	}
}

func resourceNodePoolImport(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	importId := d.Id()
	parts := strings.Split(importId, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid format specified for CCE node pool, want '<cluster_id>/<id>', but got '%s'",
			importId)
	}

	clusterID := parts[0]
	nodePoolID := parts[1]

	d.SetId(nodePoolID)
	err := d.Set("cluster_id", clusterID)
	return []*schema.ResourceData{d}, err
}
