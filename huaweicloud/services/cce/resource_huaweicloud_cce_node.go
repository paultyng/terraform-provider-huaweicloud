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
	"github.com/chnsz/golangsdk/openstack/cce/v3/nodes"
	"github.com/chnsz/golangsdk/openstack/common/tags"
	"github.com/chnsz/golangsdk/openstack/ecs/v1/cloudservers"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func ResourceNode() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNodeCreate,
		ReadContext:   resourceNodeRead,
		UpdateContext: resourceNodeUpdate,
		DeleteContext: resourceNodeDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceNodeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"availability_zone": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"os": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"key_pair": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"password", "key_pair"},
			},
			"password": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				ExactlyOneOf: []string{"password", "key_pair"},
			},
			"private_key": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"root_volume":  resourceNodeRootVolume(),
			"data_volumes": resourceNodeDataVolume(),
			"storage": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"selectors": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Default:  "evs",
									},
									"match_label_size": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"match_label_volume_type": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"match_label_metadata_encrypted": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"match_label_metadata_cmkid": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"match_label_count": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
						"groups": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"cce_managed": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"selector_names": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"virtual_spaces": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"size": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"lvm_lv_type": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"lvm_path": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"runtime_lv_type": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"taints": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"effect": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"eip_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ConflictsWith: []string{
					"eip_ids", "iptype", "bandwidth_charge_mode", "bandwidth_size", "sharetype",
				},
			},
			"iptype": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				RequiredWith: []string{
					"iptype", "bandwidth_size", "sharetype",
				},
			},
			"bandwidth_charge_mode": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"eip_ids", "eip_id"},
			},
			"sharetype": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				RequiredWith: []string{
					"iptype", "bandwidth_size", "sharetype",
				},
			},
			"bandwidth_size": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				RequiredWith: []string{
					"iptype", "bandwidth_size", "sharetype",
				},
			},
			"runtime": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"docker", "containerd",
				}, false),
			},
			"ecs_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ecs_performance_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "schema: Deprecated",
			},
			"product_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "schema: Deprecated",
			},
			"max_pods": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "schema: Deprecated",
			},
			"public_key": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "schema: Deprecated",
			},
			"preinstall": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "schema: Deprecated",
				StateFunc:   utils.DecodeHashAndHexEncode,
			},
			"postinstall": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "schema: Deprecated",
				StateFunc:   utils.DecodeHashAndHexEncode,
			},
			// (k8s_tags)
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			// (node/ecs_tags)
			"tags": common.TagsSchema(),
			"annotations": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "schema: Internal",
			},

			// charge info: charging_mode, period_unit, period, auto_renew, auto_pay
			"charging_mode": common.SchemaChargingMode(nil),
			"period_unit":   common.SchemaPeriodUnit(nil),
			"period":        common.SchemaPeriod(nil),
			"auto_renew":    common.SchemaAutoRenewUpdatable(nil),
			"auto_pay":      common.SchemaAutoPay(nil),

			"extend_param": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "schema: Deprecated",
			},
			"extend_params": resourceNodeExtendParamsSchema([]string{
				"max_pods", "public_key", "preinstall", "postinstall", "extend_param",
				"billing_mode", "order_id", "product_id", "ecs_performance_type",
			}),
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"fixed_ip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"extension_nics": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"dedicated_host_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"initialized_conditions": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"keep_ecs": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "schema: Internal",
			},
			"private_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			// Deprecated
			"eip_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				ConflictsWith: []string{
					"eip_id", "iptype", "bandwidth_charge_mode", "bandwidth_size", "sharetype",
				},
				Deprecated: "use eip_id instead",
			},
			"billing_mode": {
				Type:       schema.TypeInt,
				Optional:   true,
				ForceNew:   true,
				Computed:   true,
				Deprecated: "use charging_mode instead",
			},
			"extend_param_charging_mode": {
				Type:       schema.TypeInt,
				Optional:   true,
				ForceNew:   true,
				Deprecated: "use charging_mode instead",
			},
			"order_id": {
				Type:       schema.TypeString,
				Optional:   true,
				ForceNew:   true,
				Deprecated: "will be removed after v1.26.0",
			},
			"partition": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "schema: Internal",
			},
		},
	}
}

func buildResourceNodeAnnotations(d *schema.ResourceData) map[string]string {
	m := make(map[string]string)
	for key, val := range d.Get("annotations").(map[string]interface{}) {
		m[key] = val.(string)
	}
	return m
}

func buildResourceNodeK8sTags(d *schema.ResourceData) map[string]string {
	m := make(map[string]string)
	for key, val := range d.Get("labels").(map[string]interface{}) {
		m[key] = val.(string)
	}
	return m
}

func buildResourceNodeTags(d *schema.ResourceData) []tags.ResourceTag {
	tagRaw := d.Get("tags").(map[string]interface{})
	return utils.ExpandResourceTags(tagRaw)
}

func buildResourceNodeTaint(d *schema.ResourceData) []nodes.TaintSpec {
	taintRaw := d.Get("taints").([]interface{})
	taints := make([]nodes.TaintSpec, len(taintRaw))
	for i, raw := range taintRaw {
		rawMap := raw.(map[string]interface{})
		taints[i] = nodes.TaintSpec{
			Key:    rawMap["key"].(string),
			Value:  rawMap["value"].(string),
			Effect: rawMap["effect"].(string),
		}
	}
	return taints
}

func buildResourceNodeEipIDs(d *schema.ResourceData) []string {
	if v, ok := d.GetOk("eip_id"); ok {
		return []string{v.(string)}
	}
	rawID := d.Get("eip_ids").(*schema.Set)
	id := make([]string, rawID.Len())
	for i, raw := range rawID.List() {
		id[i] = raw.(string)
	}
	return id
}

func buildResourceNodeStorage(d *schema.ResourceData) *nodes.StorageSpec {
	v, ok := d.GetOk("storage")
	if !ok {
		return nil
	}

	var storageSpec nodes.StorageSpec
	storageSpecRaw := v.([]interface{})
	storageSpecRawMap := storageSpecRaw[0].(map[string]interface{})
	storageSelectorSpecRaw := storageSpecRawMap["selectors"].([]interface{})
	storageGroupSpecRaw := storageSpecRawMap["groups"].([]interface{})

	var selectors []nodes.StorageSelectorsSpec
	for _, s := range storageSelectorSpecRaw {
		sMap := s.(map[string]interface{})
		selector := nodes.StorageSelectorsSpec{
			Name:        sMap["name"].(string),
			StorageType: sMap["type"].(string),
			MatchLabels: nodes.MatchLabelsSpec{
				Size:              sMap["match_label_size"].(string),
				VolumeType:        sMap["match_label_volume_type"].(string),
				MetadataEncrypted: sMap["match_label_metadata_encrypted"].(string),
				MetadataCmkid:     sMap["match_label_metadata_cmkid"].(string),
				Count:             sMap["match_label_count"].(string),
			},
		}
		selectors = append(selectors, selector)
	}
	storageSpec.StorageSelectors = selectors

	var groups []nodes.StorageGroupsSpec
	for _, g := range storageGroupSpecRaw {
		gMap := g.(map[string]interface{})
		group := nodes.StorageGroupsSpec{
			Name:          gMap["name"].(string),
			CceManaged:    gMap["cce_managed"].(bool),
			SelectorNames: utils.ExpandToStringList(gMap["selector_names"].([]interface{})),
		}

		virtualSpacesRaw := gMap["virtual_spaces"].([]interface{})
		virtualSpaces := make([]nodes.VirtualSpacesSpec, 0, len(virtualSpacesRaw))
		for _, v := range virtualSpacesRaw {
			virtualSpaceMap := v.(map[string]interface{})
			virtualSpace := nodes.VirtualSpacesSpec{
				Name: virtualSpaceMap["name"].(string),
				Size: virtualSpaceMap["size"].(string),
			}

			if virtualSpaceMap["lvm_lv_type"].(string) != "" {
				lvmConfig := nodes.LVMConfigSpec{
					LvType: virtualSpaceMap["lvm_lv_type"].(string),
					Path:   virtualSpaceMap["lvm_path"].(string),
				}
				virtualSpace.LVMConfig = &lvmConfig
			}

			if virtualSpaceMap["runtime_lv_type"].(string) != "" {
				runtimeConfig := nodes.RuntimeConfigSpec{
					LvType: virtualSpaceMap["runtime_lv_type"].(string),
				}
				virtualSpace.RuntimeConfig = &runtimeConfig
			}

			virtualSpaces = append(virtualSpaces, virtualSpace)
		}
		group.VirtualSpaces = virtualSpaces

		groups = append(groups, group)
	}

	storageSpec.StorageGroups = groups
	return &storageSpec
}

func buildResourceNodePublicIP(d *schema.ResourceData) nodes.PublicIPSpec {
	// eipCount must be specified when bandwidth_size parameters was set
	eipCount := 0
	if _, ok := d.GetOk("bandwidth_size"); ok {
		eipCount = 1
	}

	res := nodes.PublicIPSpec{
		Ids:   buildResourceNodeEipIDs(d),
		Count: eipCount,
		Eip: nodes.EipSpec{
			IpType: d.Get("iptype").(string),
			Bandwidth: nodes.BandwidthOpts{
				ChargeMode: d.Get("bandwidth_charge_mode").(string),
				Size:       d.Get("bandwidth_size").(int),
				ShareType:  d.Get("sharetype").(string),
			},
		},
	}

	return res
}

func buildResourceNodeNicSpec(d *schema.ResourceData) nodes.NodeNicSpec {
	res := nodes.NodeNicSpec{
		PrimaryNic: nodes.PrimaryNic{
			SubnetId: d.Get("subnet_id").(string),
		},
	}

	if v, ok := d.GetOk("fixed_ip"); ok {
		res.PrimaryNic.FixedIps = []string{v.(string)}
	}

	if v, ok := d.GetOk("extension_nics"); ok {
		nicsRaw := v.([]interface{})
		extNics := make([]nodes.ExtNic, len(nicsRaw))
		for i, v := range nicsRaw {
			nic := v.(map[string]interface{})
			extNics[i] = nodes.ExtNic{
				SubnetId: nic["subnet_id"].(string),
			}
		}

		res.ExtNics = extNics
	}

	return res
}

func buildResourceNodeLoginSpec(d *schema.ResourceData) (nodes.LoginSpec, error) {
	var loginSpec nodes.LoginSpec
	if v, ok := d.GetOk("key_pair"); ok {
		loginSpec = nodes.LoginSpec{
			SshKey: v.(string),
		}
	} else {
		password, err := utils.TryPasswordEncrypt(d.Get("password").(string))
		if err != nil {
			return loginSpec, err
		}
		loginSpec = nodes.LoginSpec{
			UserPassword: nodes.UserPassword{
				Username: "root",
				Password: password,
			},
		}
	}

	return loginSpec, nil
}

func resourceNodeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)
	nodeClient, err := cfg.CceV3Client(region)
	if err != nil {
		return diag.Errorf("error creating CCE Node client: %s", err)
	}

	// validation
	billingMode := 0
	if d.Get("charging_mode").(string) == "prePaid" || d.Get("billing_mode").(int) == 1 {
		billingMode = 1
		if err := common.ValidatePrePaidChargeInfo(d); err != nil {
			return diag.FromErr(err)
		}
	}

	// wait for the cce cluster to become available
	clusterid := d.Get("cluster_id").(string)
	stateCluster := &resource.StateChangeConf{
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      clusterStateRefreshFunc(nodeClient, clusterid, []string{"Available"}),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        5 * time.Second,
		PollInterval: 5 * time.Second,
	}
	_, err = stateCluster.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("error waiting for CCE cluster to be Available: %s", err)
	}

	createOpts := nodes.CreateOpts{
		Kind:       "Node",
		ApiVersion: "v3",
		Metadata: nodes.CreateMetaData{
			Name:        d.Get("name").(string),
			Annotations: buildResourceNodeAnnotations(d),
		},
		Spec: nodes.Spec{
			Flavor:                d.Get("flavor_id").(string),
			Az:                    d.Get("availability_zone").(string),
			Os:                    d.Get("os").(string),
			RootVolume:            buildResourceNodeRootVolume(d),
			DataVolumes:           buildResourceNodeDataVolume(d),
			Storage:               buildResourceNodeStorage(d),
			PublicIP:              buildResourceNodePublicIP(d),
			BillingMode:           billingMode,
			Count:                 1,
			NodeNicSpec:           buildResourceNodeNicSpec(d),
			EcsGroupID:            d.Get("ecs_group_id").(string),
			ExtendParam:           buildExtendParams(d),
			Taints:                buildResourceNodeTaint(d),
			K8sTags:               buildResourceNodeK8sTags(d),
			UserTags:              buildResourceNodeTags(d),
			DedicatedHostID:       d.Get("dedicated_host_id").(string),
			InitializedConditions: utils.ExpandToStringList(d.Get("initialized_conditions").([]interface{})),
		},
	}

	if v, ok := d.GetOk("runtime"); ok {
		createOpts.Spec.RunTime = &nodes.RunTimeSpec{
			Name: v.(string),
		}
	}

	// Create a node in the specified partition
	if v, ok := d.GetOk("partition"); ok {
		createOpts.Spec.Partition = v.(string)
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)

	// Add loginSpec here so it wouldn't go in the above log entry
	loginSpec, err := buildResourceNodeLoginSpec(d)
	if err != nil {
		diag.FromErr(err)
	}
	createOpts.Spec.Login = loginSpec

	s, err := nodes.Create(nodeClient, clusterid, createOpts).Extract()
	if err != nil {
		return diag.Errorf("error creating Node: %s", err)
	}

	if orderId, ok := s.Spec.ExtendParam["orderID"]; ok && orderId != "" {
		bssClient, err := cfg.BssV2Client(region)
		if err != nil {
			return diag.Errorf("error creating BSS v2 client: %s", err)
		}
		err = common.WaitOrderComplete(ctx, bssClient, orderId.(string), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return diag.FromErr(err)
		}
		// The resource ID generated by the CBC service only means that the underlying ECS of CCE node is created.
		_, err = common.WaitOrderResourceComplete(ctx, bssClient, orderId.(string), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return diag.FromErr(err)
		}
	}
	// The completion of the creation of the underlying resource (ECS) corresponding to the CCE node does not mean that
	// the creation of the CCE node is completed.
	nodeID, err := getResourceIDFromJob(ctx, nodeClient, s.Status.JobID, "CreateNode", "CreateNodeVM",
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(nodeID)

	log.Printf("[DEBUG] Waiting for CCE Node (%s) to become available", s.Metadata.Name)
	stateConf := &resource.StateChangeConf{
		// The statuses of pending phase includes "Build" and "Installing".
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      nodeStateRefreshFunc(nodeClient, clusterid, nodeID, []string{"Active"}),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        20 * time.Second,
		PollInterval: 20 * time.Second,
	}
	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("error creating CCE Node: %s", err)
	}

	return resourceNodeRead(ctx, d, meta)
}

func resourceNodeRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)
	nodeClient, err := cfg.CceV3Client(region)
	if err != nil {
		return diag.Errorf("error creating CCE Node client: %s", err)
	}
	clusterid := d.Get("cluster_id").(string)
	nodeId := d.Id()
	s, err := nodes.Get(nodeClient, clusterid, nodeId).Extract()

	if err != nil {
		return common.CheckDeletedDiag(d, err, "error retrieving CCE Node")
	}

	mErr := multierror.Append(nil,
		d.Set("region", region),
		d.Set("name", s.Metadata.Name),
		d.Set("flavor_id", s.Spec.Flavor),
		d.Set("availability_zone", s.Spec.Az),
		d.Set("os", s.Spec.Os),
		d.Set("key_pair", s.Spec.Login.SshKey),
		d.Set("subnet_id", s.Spec.NodeNicSpec.PrimaryNic.SubnetId),
		d.Set("ecs_group_id", s.Spec.EcsGroupID),
		d.Set("server_id", s.Status.ServerID),
		d.Set("private_ip", s.Status.PrivateIP),
		d.Set("public_ip", s.Status.PublicIP),
		d.Set("status", s.Status.Phase),
		d.Set("root_volume", flattenResourceNodeRootVolume(s.Spec.RootVolume)),
		d.Set("data_volumes", flattenResourceNodeDataVolume(s.Spec.DataVolumes)),
		d.Set("initialized_conditions", s.Spec.InitializedConditions),
	)

	if s.Spec.BillingMode != 0 {
		mErr = multierror.Append(mErr, d.Set("charging_mode", "prePaid"))
	}
	if s.Spec.RunTime != nil {
		mErr = multierror.Append(mErr, d.Set("runtime", s.Spec.RunTime.Name))
	}

	computeClient, err := cfg.ComputeV1Client(region)
	if err != nil {
		return diag.Errorf("error creating compute client: %s", err)
	}

	serverId := s.Status.ServerID
	// fetch key_pair from ECS instance
	if server, err := cloudservers.Get(computeClient, serverId).Extract(); err == nil {
		mErr = multierror.Append(mErr, d.Set("key_pair", server.KeyName))
	} else {
		log.Printf("[WARN] Error fetching ECS instance (%s): %s", serverId, err)
	}

	// fetch tags from ECS instance
	if resourceTags, err := tags.Get(computeClient, "cloudservers", serverId).Extract(); err == nil {
		// Special blacklist of CCE service.
		sysTagsBL := []string{
			"CCE-Cluster-ID",
			"CCE-Dynamic-Provisioning-Node",
			// HCS provider used
			fmt.Sprintf("CCE-Cluster-ID.%s", clusterid),
			fmt.Sprintf("CCE-Dynamic-Provisioning-Node.%s", nodeId),
			clusterid,
			nodeId,
		}
		tagmap := utils.TagsToMapWithIndicator(resourceTags.Tags, sysTagsBL, d.Get("tags").(map[string]interface{}))
		mErr = multierror.Append(mErr, d.Set("tags", tagmap))
	} else {
		log.Printf("[WARN] Error fetching tags of ECS instance (%s): %s", serverId, err)
	}

	if err = mErr.ErrorOrNil(); err != nil {
		return diag.Errorf("error setting CCE Node fields: %s", err)
	}
	return nil
}

func resourceNodeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	nodeClient, err := cfg.CceV3Client(region)
	if err != nil {
		return diag.Errorf("error creating CCE client: %s", err)
	}
	computeClient, err := cfg.ComputeV1Client(region)
	if err != nil {
		return diag.Errorf("error creating compute client: %s", err)
	}

	if d.HasChange("name") {
		var updateOpts nodes.UpdateOpts
		updateOpts.Metadata.Name = d.Get("name").(string)

		clusterid := d.Get("cluster_id").(string)
		_, err = nodes.Update(nodeClient, clusterid, d.Id(), updateOpts).Extract()
		if err != nil {
			return diag.Errorf("error updating cce node: %s", err)
		}
	}

	serverId := d.Get("server_id").(string)

	// update node tags with ECS API
	if d.HasChange("tags") {
		tagErr := utils.UpdateResourceTags(computeClient, d, "cloudservers", serverId)
		if tagErr != nil {
			return diag.Errorf("error updating tags of cce node %s: %s", d.Id(), tagErr)
		}
	}

	// update node key_pair with DEW API
	if d.HasChange("key_pair") {
		kmsClient, err := cfg.KmsV3Client(region)
		if err != nil {
			return diag.Errorf("error creating KMS v3 client: %s", err)
		}

		currentPwd, _ := d.GetChange("password")
		o, n := d.GetChange("key_pair")
		keyPairOpts := &common.KeypairAuthOpts{
			InstanceID:       serverId,
			InUsedKeyPair:    o.(string),
			NewKeyPair:       n.(string),
			InUsedPrivateKey: d.Get("private_key").(string),
			Password:         currentPwd.(string),
			DisablePassword:  true,
			Timeout:          d.Timeout(schema.TimeoutUpdate),
		}
		if err := common.UpdateEcsInstanceKeyPair(ctx, computeClient, kmsClient, keyPairOpts); err != nil {
			return diag.FromErr(err)
		}
	}

	// update node password with ECS API
	// A new password takes effect after the ECS is started or restarted.
	if d.HasChange("password") {
		// if the password is empty, it means that the ECS instance will bind a new keypair
		if newPwd, ok := d.GetOk("password"); ok {
			err := cloudservers.ChangeAdminPassword(computeClient, serverId, newPwd.(string)).ExtractErr()
			if err != nil {
				return diag.Errorf("error changing password of cce node %s: %s", d.Id(), err)
			}
		}
	}

	if d.HasChange("auto_renew") {
		bssClient, err := cfg.BssV2Client(region)
		if err != nil {
			return diag.Errorf("error creating BSS V2 client: %s", err)
		}
		if err = common.UpdateAutoRenew(bssClient, d.Get("auto_renew").(string), d.Get("server_id").(string)); err != nil {
			// Do not output the underlying ECS instance ID externally.
			return diag.Errorf("error updating the auto-renew of the node (%s): %s", d.Id(), err)
		}
	}

	return resourceNodeRead(ctx, d, meta)
}

func resourceNodeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)
	nodeClient, err := cfg.CceV3Client(region)
	if err != nil {
		return diag.Errorf("error creating CCE client: %s", err)
	}

	clusterid := d.Get("cluster_id").(string)
	// remove node without deleting ecs
	if d.Get("keep_ecs").(bool) {
		err := removeNode(nodeClient, d, clusterid)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		err := deleteNode(ctx, cfg, nodeClient, d, clusterid)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	stateConf := &resource.StateChangeConf{
		// The statuses of pending phase include "Deleting".
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      nodeStateRefreshFunc(nodeClient, clusterid, d.Id(), nil),
		Timeout:      d.Timeout(schema.TimeoutDelete),
		Delay:        60 * time.Second,
		PollInterval: 20 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("error deleting CCE Node: %s", err)
	}

	d.SetId("")
	return nil
}

func removeNode(nodeClient *golangsdk.ServiceClient, d *schema.ResourceData, clusterID string) error {
	loginSpec, err := buildResourceNodeLoginSpec(d)
	if err != nil {
		return err
	}

	removeOpts := nodes.RemoveOpts{
		Spec: nodes.RemoveNodeSpec{
			Login: loginSpec,
			Nodes: []nodes.NodeItem{
				{
					Uid: d.Id(),
				},
			},
		},
	}

	err = nodes.Remove(nodeClient, clusterID, removeOpts).ExtractErr()
	if err != nil {
		return fmt.Errorf("error removing CCE node: %s", err)
	}

	return nil
}

func deleteNode(ctx context.Context, cfg *config.Config, nodeClient *golangsdk.ServiceClient,
	d *schema.ResourceData, clusterID string) error {
	// for prePaid node, firstly, we should unsubscribe the ecs server, and then delete it
	if d.Get("charging_mode").(string) == "prePaid" || d.Get("billing_mode").(int) == 1 {
		region := cfg.GetRegion(d)
		serverID := d.Get("server_id").(string)
		publicIP := d.Get("public_ip").(string)

		resourceIDs, err := getResourceIDsToUnsubscribe(cfg, d, serverID, publicIP)
		if err != nil {
			return err
		}

		if len(resourceIDs) > 0 {
			if err := common.UnsubscribePrePaidResource(d, cfg, resourceIDs); err != nil {
				return fmt.Errorf("error unsubscribing CCE node: %s", err)
			}
		}

		// wait for the ecs server of the prePaid node to be deleted
		computeClient, err := cfg.ComputeV1Client(region)
		if err != nil {
			return fmt.Errorf("error creating compute client: %s", err)
		}

		pending := []string{"ACTIVE", "SHUTOFF"}
		target := []string{"DELETED", "SOFT_DELETED"}
		deleteTimeout := d.Timeout(schema.TimeoutDelete)
		if err := waitForServerTargetState(ctx, computeClient, serverID, pending, target, deleteTimeout); err != nil {
			return fmt.Errorf("state waiting timeout: %s", err)
		}
	}

	err := nodes.Delete(nodeClient, clusterID, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("error deleting CCE node: %s", err)
	}

	return nil
}

func getResourceIDsToUnsubscribe(cfg *config.Config, d *schema.ResourceData, serverID, publicIP string) ([]string, error) {
	resourceIDs := make([]string, 0, 2)
	region := cfg.GetRegion(d)

	// check whether the ecs server of the perPaid exists before unsubscribe it
	// because resource could not be found cannot be unsubscribed
	if serverID != "" {
		computeClient, err := cfg.ComputeV1Client(region)
		if err != nil {
			return nil, fmt.Errorf("error creating compute client: %s", err)
		}

		server, err := cloudservers.Get(computeClient, serverID).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); !ok {
				return nil, fmt.Errorf("error retrieving ECS intance: %s", err)
			}
		} else if server.Status != "DELETED" && server.Status != "SOFT_DELETED" {
			resourceIDs = append(resourceIDs, serverID)
		}
	}

	// unsubscribe the eip if necessary
	if _, ok := d.GetOk("iptype"); ok && publicIP != "" {
		eipClient, err := cfg.NetworkingV1Client(region)
		if err != nil {
			return nil, fmt.Errorf("error creating networking client: %s", err)
		}

		epsID := "all_granted_eps"
		if eipID, err := common.GetEipIDbyAddress(eipClient, publicIP, epsID); err == nil {
			resourceIDs = append(resourceIDs, eipID)
		} else {
			log.Printf("[WARN] error fetching EIP ID of CCE Node (%s): %s", d.Id(), err)
		}
	}

	return resourceIDs, nil
}

func getResourceIDFromJob(ctx context.Context, client *golangsdk.ServiceClient, jobID, jobType, subJobType string,
	timeout time.Duration) (string, error) {
	stateJob := &resource.StateChangeConf{
		Pending:      []string{"Initializing", "Running"},
		Target:       []string{"Success"},
		Refresh:      waitForJobStatus(client, jobID),
		Timeout:      timeout,
		Delay:        120 * time.Second,
		PollInterval: 20 * time.Second,
	}

	v, err := stateJob.WaitForStateContext(ctx)
	if err != nil {
		if job, ok := v.(*nodes.Job); ok {
			return "", fmt.Errorf("error waiting for job (%s) to become success: %s, reason: %s",
				jobID, err, job.Status.Reason)
		}

		return "", fmt.Errorf("error waiting for job (%s) to become success: %s", jobID, err)
	}

	job := v.(*nodes.Job)
	if len(job.Spec.SubJobs) == 0 {
		return "", fmt.Errorf("error fetching sub jobs from %s", jobID)
	}

	var subJobID string
	var refreshJob bool
	for _, s := range job.Spec.SubJobs {
		// postPaid: should get details of sub job ID
		if s.Spec.Type == jobType {
			subJobID = s.Metadata.ID
			refreshJob = true
			break
		}
	}

	if refreshJob {
		job, err = nodes.GetJobDetails(client, subJobID).ExtractJob()
		if err != nil {
			return "", fmt.Errorf("error fetching sub Job %s: %s", subJobID, err)
		}
	}

	for _, s := range job.Spec.SubJobs {
		if s.Spec.Type == subJobType {
			return s.Spec.ResourceID, nil
		}
	}

	return "", fmt.Errorf("error fetching the resource ID from the specified job (type: %s)", subJobType)
}

func nodeStateRefreshFunc(cceClient *golangsdk.ServiceClient, clusterId, nodeId string,
	targets []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Expect the status of CCE node to be any one of the status list: %v", targets)
		resp, err := nodes.Get(cceClient, clusterId, nodeId).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				log.Printf("[DEBUG] The node (%s) has been deleted", clusterId)
				return resp, "COMPLETED", nil
			}
			return nil, "ERROR", err
		}

		invalidStatuses := []string{"Error", "Shelved", "Unknow"}
		if utils.IsStrContainsSliceElement(resp.Status.Phase, invalidStatuses, true, true) {
			return resp, "ERROR", fmt.Errorf("unexpected status: %s", resp.Status.Phase)
		}

		if utils.StrSliceContains(targets, resp.Status.Phase) {
			return resp, "COMPLETED", nil
		}
		return resp, "PENDING", nil
	}
}

func waitForJobStatus(cceClient *golangsdk.ServiceClient, jobID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		job, err := nodes.GetJobDetails(cceClient, jobID).ExtractJob()
		if err != nil {
			return nil, "", err
		}

		return job, job.Status.Phase, nil
	}
}

func resourceNodeImport(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		err := fmt.Errorf("invalid format specified for CCE Node. Format must be <cluster id>/<node id>")
		return nil, err
	}

	clusterID := parts[0]
	nodeID := parts[1]

	d.SetId(nodeID)
	d.Set("cluster_id", clusterID)

	return []*schema.ResourceData{d}, nil
}

func waitForServerTargetState(ctx context.Context, client *golangsdk.ServiceClient, id string,
	pending, target []string, timeout time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending:      pending,
		Target:       target,
		Refresh:      ServerV1StateRefreshFunc(client, id),
		Timeout:      timeout,
		Delay:        5 * time.Second,
		PollInterval: 5 * time.Second,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return fmt.Errorf("error waiting for instance (%s) to become target state (%v): %s", id, target, err)
	}
	return nil
}

// ServerV1StateRefreshFunc returns a resource.StateRefreshFunc that is used to watch an instance.
func ServerV1StateRefreshFunc(client *golangsdk.ServiceClient, instanceID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		s, err := cloudservers.Get(client, instanceID).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return s, "DELETED", nil
			}
			return nil, "", err
		}

		// get fault message when status is ERROR
		if s.Status == "ERROR" {
			fault := fmt.Errorf("[error code: %d, message: %s]", s.Fault.Code, s.Fault.Message)
			return s, "ERROR", fault
		}
		return s, s.Status, nil
	}
}
