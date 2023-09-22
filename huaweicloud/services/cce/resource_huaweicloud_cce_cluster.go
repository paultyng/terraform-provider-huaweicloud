package cce

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/openstack/aom/v1/icagents"
	"github.com/chnsz/golangsdk/openstack/cce/v3/clusters"
	"github.com/chnsz/golangsdk/openstack/cce/v3/nodes"
	"github.com/chnsz/golangsdk/openstack/common/tags"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// ResourceCCEClusterV3 defines the CCE cluster resource schema and functions.
// Deprecated: It's a deprecated function, please refer to the function 'ResourceCluster'.
func ResourceCCEClusterV3() *schema.Resource {
	return ResourceCluster()
}

var associateDeleteSchema *schema.Schema = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
	ValidateFunc: validation.StringInSlice([]string{
		"true", "try", "false",
	}, true),
	ConflictsWith: []string{"delete_all"},
}

var associateDeleteSchemaInternal *schema.Schema = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
	ValidateFunc: validation.StringInSlice([]string{
		"true", "try", "false",
	}, true),
	ConflictsWith: []string{"delete_all"},
	Description:   "schema: Internal",
}

// ResourceCluster defines the CCE cluster resource schema and functions.
func ResourceCluster() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceClusterCreate,
		ReadContext:   resourceClusterRead,
		UpdateContext: resourceClusterUpdate,
		DeleteContext: resourceClusterDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		//request and response parameters
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
				ForceNew: true,
			},
			"flavor_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_version": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: utils.SuppressVersionDiffs,
			},
			"cluster_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "VirtualMachine",
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "schema: Internal",
			},
			"annotations": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "schema: Internal",
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"security_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"highway_subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "schema: Internal",
			},
			"container_network_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"container_network_cidr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"eni_subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "the IPv4 subnet ID of the subnet where the ENI resides",
			},
			"eni_subnet_cidr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "schema: Computed",
			},
			"enable_distribute_management": {
				Type:         schema.TypeBool,
				Optional:     true,
				ForceNew:     true,
				RequiredWith: []string{"eni_subnet_id", "eni_subnet_cidr"},
				Description:  "schema: Internal",
			},
			"authentication_mode": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "rbac",
			},
			"authenticating_proxy_ca": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				RequiredWith: []string{"authenticating_proxy_cert", "authenticating_proxy_private_key"},
			},
			"authenticating_proxy_cert": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				RequiredWith: []string{"authenticating_proxy_ca", "authenticating_proxy_private_key"},
			},
			"authenticating_proxy_private_key": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				RequiredWith: []string{"authenticating_proxy_ca", "authenticating_proxy_cert"},
			},
			"multi_az": {
				Type:          schema.TypeBool,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"masters"},
			},
			"masters": {
				Type:          schema.TypeList,
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				MaxItems:      3,
				ConflictsWith: []string{"multi_az"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"availability_zone": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
					},
				},
			},
			"eip": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.ValidateIP,
			},
			"service_network_cidr": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"kube_proxy_mode": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"enterprise_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"extend_param": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hibernate": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tags": common.TagsForceNewSchema(),

			// charge info: charging_mode, period_unit, period, auto_renew, auto_pay
			"charging_mode": common.SchemaChargingMode(nil),
			"period_unit":   common.SchemaPeriodUnit(nil),
			"period":        common.SchemaPeriod(nil),
			"auto_renew":    common.SchemaAutoRenewUpdatable(nil),
			"auto_pay":      common.SchemaAutoPay(nil),

			"delete_efs": associateDeleteSchema,
			"delete_eni": associateDeleteSchemaInternal,
			"delete_evs": associateDeleteSchema,
			"delete_net": associateDeleteSchemaInternal,
			"delete_obs": associateDeleteSchema,
			"delete_sfs": associateDeleteSchema,
			"delete_all": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"true", "try", "false",
				}, true),
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kube_config_raw": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"server": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificate_authority_data": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"certificate_users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"client_certificate_data": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"client_key_data": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Deprecated
			"billing_mode": {
				Type:       schema.TypeInt,
				Optional:   true,
				Computed:   true,
				ForceNew:   true,
				Deprecated: "use charging_mode instead",
			},
		},
	}
}

func resourceClusterLabels(d *schema.ResourceData) map[string]string {
	m := make(map[string]string)
	for key, val := range d.Get("labels").(map[string]interface{}) {
		m[key] = val.(string)
	}
	return m
}

func resourceClusterTags(d *schema.ResourceData) []tags.ResourceTag {
	tagRaw := d.Get("tags").(map[string]interface{})
	return utils.ExpandResourceTags(tagRaw)
}

func resourceClusterAnnotations(d *schema.ResourceData) map[string]string {
	m := make(map[string]string)
	for key, val := range d.Get("annotations").(map[string]interface{}) {
		m[key] = val.(string)
	}
	return m
}

func resourceClusterExtendParam(d *schema.ResourceData, config *config.Config) map[string]interface{} {
	extendParam := make(map[string]interface{})
	if v, ok := d.GetOk("extend_param"); ok {
		for key, val := range v.(map[string]interface{}) {
			extendParam[key] = val.(string)
		}
	}

	// assemble the charge info
	var isPrePaid bool
	var billingMode int
	if v, ok := d.GetOk("charging_mode"); ok && v.(string) == "prePaid" {
		isPrePaid = true
	}
	if v, ok := d.GetOk("billing_mode"); ok {
		billingMode = v.(int)
	}
	if isPrePaid || billingMode == 1 {
		extendParam["isAutoRenew"] = "false"
		extendParam["isAutoPay"] = common.GetAutoPay(d)
	}

	if v, ok := d.GetOk("period_unit"); ok {
		extendParam["periodType"] = v.(string)
	}
	if v, ok := d.GetOk("period"); ok {
		extendParam["periodNum"] = v.(int)
	}
	if v, ok := d.GetOk("auto_renew"); ok {
		extendParam["isAutoRenew"] = v.(string)
	}

	if multi_az, ok := d.GetOk("multi_az"); ok && multi_az == true {
		extendParam["clusterAZ"] = "multi_az"
	}
	if kube_proxy_mode, ok := d.GetOk("kube_proxy_mode"); ok {
		extendParam["kubeProxyMode"] = kube_proxy_mode.(string)
	}
	if eip, ok := d.GetOk("eip"); ok {
		extendParam["clusterExternalIP"] = eip.(string)
	}

	epsID := config.GetEnterpriseProjectID(d)
	if epsID != "" {
		extendParam["enterpriseProjectId"] = epsID
	}

	return extendParam
}

func resourceClusterMasters(d *schema.ResourceData) ([]clusters.MasterSpec, error) {
	if v, ok := d.GetOk("masters"); ok {
		flavorId := d.Get("flavor_id").(string)
		mastersRaw := v.([]interface{})
		if strings.Contains(flavorId, "s1") && len(mastersRaw) != 1 {
			return nil, fmt.Errorf("error creating CCE cluster: "+
				"single-master cluster need 1 az for master node, but got %d", len(mastersRaw))
		}
		if strings.Contains(flavorId, "s2") && len(mastersRaw) != 3 {
			return nil, fmt.Errorf("error creating CCE cluster: "+
				"high-availability cluster need 3 az for master nodes, but got %d", len(mastersRaw))
		}
		masters := make([]clusters.MasterSpec, len(mastersRaw))
		for i, raw := range mastersRaw {
			rawMap := raw.(map[string]interface{})
			masters[i] = clusters.MasterSpec{
				MasterAZ: rawMap["availability_zone"].(string),
			}
		}
		return masters, nil
	}

	return nil, nil
}

func buildContainerNetworkCidrsOpts(cidrs string) []clusters.CidrSpec {
	if cidrs == "" {
		return nil
	}

	cidrList := strings.Split(cidrs, ",")

	res := make([]clusters.CidrSpec, len(cidrList))
	for i, cidr := range cidrList {
		res[i] = clusters.CidrSpec{
			Cidr: cidr,
		}
	}

	return res
}

func buildEniNetworkOpts(eniSubnetID string) *clusters.EniNetworkSpec {
	if eniSubnetID == "" {
		return nil
	}

	subnetIDs := strings.Split(eniSubnetID, ",")
	subnets := make([]clusters.EniSubnetSpec, len(subnetIDs))
	for i, subnetID := range subnetIDs {
		subnets[i] = clusters.EniSubnetSpec{
			SubnetID: subnetID,
		}
	}

	eniNetwork := clusters.EniNetworkSpec{
		Subnets: subnets,
	}

	return &eniNetwork
}

func resourceClusterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	cceClient, err := config.CceV3Client(config.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating CCE v3 client: %s", err)
	}
	icAgentClient, err := config.AomV1Client(config.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating AOM v1 client: %s", err)
	}

	authenticating_proxy := make(map[string]string)
	if common.HasFilledOpt(d, "authenticating_proxy_ca") {
		authenticating_proxy["ca"] = utils.TryBase64EncodeString(d.Get("authenticating_proxy_ca").(string))
		authenticating_proxy["cert"] = utils.TryBase64EncodeString(d.Get("authenticating_proxy_cert").(string))
		authenticating_proxy["privateKey"] = utils.TryBase64EncodeString(d.Get("authenticating_proxy_private_key").(string))
	}

	billingMode := 0
	if d.Get("charging_mode").(string) == "prePaid" || d.Get("billing_mode").(int) == 1 {
		billingMode = 1
		if err := common.ValidatePrePaidChargeInfo(d); err != nil {
			return diag.FromErr(err)
		}
	}

	clusterName := d.Get("name").(string)
	createOpts := clusters.CreateOpts{
		Kind:       "Cluster",
		ApiVersion: "v3",
		Metadata: clusters.CreateMetaData{
			Name:        clusterName,
			Labels:      resourceClusterLabels(d),
			Annotations: resourceClusterAnnotations(d)},
		Spec: clusters.Spec{
			Type:        d.Get("cluster_type").(string),
			Flavor:      d.Get("flavor_id").(string),
			Version:     d.Get("cluster_version").(string),
			Description: d.Get("description").(string),
			HostNetwork: clusters.HostNetworkSpec{
				VpcId:         d.Get("vpc_id").(string),
				SubnetId:      d.Get("subnet_id").(string),
				HighwaySubnet: d.Get("highway_subnet_id").(string),
				SecurityGroup: d.Get("security_group_id").(string),
			},
			ContainerNetwork: clusters.ContainerNetworkSpec{
				Mode:  d.Get("container_network_type").(string),
				Cidrs: buildContainerNetworkCidrsOpts(d.Get("container_network_cidr").(string)),
			},
			EniNetwork: buildEniNetworkOpts(d.Get("eni_subnet_id").(string)),
			Authentication: clusters.AuthenticationSpec{
				Mode:                d.Get("authentication_mode").(string),
				AuthenticatingProxy: authenticating_proxy,
			},
			BillingMode:          billingMode,
			ExtendParam:          resourceClusterExtendParam(d, config),
			KubernetesSvcIPRange: d.Get("service_network_cidr").(string),
			ClusterTags:          resourceClusterTags(d),
		},
	}

	if _, ok := d.GetOk("enable_distribute_management"); ok {
		createOpts.Spec.EnableDistMgt = d.Get("enable_distribute_management").(bool)
	}

	masters, err := resourceClusterMasters(d)
	if err != nil {
		return diag.FromErr(err)
	}
	createOpts.Spec.Masters = masters

	s, err := clusters.Create(cceClient, createOpts).Extract()
	if err != nil {
		return diag.Errorf("error creating CCE cluster: %s", err)
	}

	if orderId, ok := s.Spec.ExtendParam["orderID"]; ok && orderId != "" {
		bssClient, err := config.BssV2Client(config.GetRegion(d))
		if err != nil {
			return diag.Errorf("error creating BSS v2 client: %s", err)
		}
		err = common.WaitOrderComplete(ctx, bssClient, orderId.(string), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return diag.FromErr(err)
		}
		resourceId, err := common.WaitOrderResourceComplete(ctx, bssClient, orderId.(string), d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId(resourceId)
	} else {
		jobID := s.Status.JobID
		if jobID == "" {
			return diag.Errorf("error fetching job ID after creating CCE cluster: %s", clusterName)
		}

		clusterID, err := getClusterIDFromJob(ctx, cceClient, jobID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(clusterID)
	}

	log.Printf("[DEBUG] Waiting for CCE cluster (%s) to become available", d.Id())
	stateConf := &resource.StateChangeConf{
		// The statuses of pending phase include "Creating".
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      clusterStateRefreshFunc(cceClient, d.Id(), []string{"Available"}),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        20 * time.Second,
		PollInterval: 20 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("error creating CCE cluster: %s", err)
	}

	log.Printf("[DEBUG] Installing ICAgent for CCE cluster (%s)", d.Id())
	installParam := icagents.InstallParam{
		ClusterId: d.Id(),
		NameSpace: "default",
	}
	result := icagents.Create(icAgentClient, installParam)
	var diags diag.Diagnostics
	if result.Err != nil {
		diagIcagent := diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Error installing ICAgent",
			Detail:   fmt.Sprintf("error installing ICAgent in CCE cluster: %s", result.Err),
		}
		diags = append(diags, diagIcagent)
	}

	// create a hibernating cluster
	if d.Get("hibernate").(bool) {
		err = resourceClusterHibernate(ctx, d, cceClient)
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	diags = append(diags, resourceClusterRead(ctx, d, meta)...)

	return diags

}

func resourceClusterRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	cceClient, err := config.CceV3Client(config.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating CCE v3 client: %s", err)
	}

	clusterId := d.Id()
	n, err := clusters.Get(cceClient, clusterId).Extract()
	if err != nil {
		return common.CheckDeletedDiag(d, err, "CCE cluster")
	}

	// Special blacklist of CCE service.
	sysTagsBL := []string{
		"CCE-Cluster-ID",
		// HCS provider used
		fmt.Sprintf("CCE-Cluster-ID.%s", clusterId),
		clusterId,
	}
	mErr := multierror.Append(nil,
		d.Set("region", config.GetRegion(d)),
		d.Set("name", n.Metadata.Name),
		d.Set("status", n.Status.Phase),
		d.Set("flavor_id", n.Spec.Flavor),
		d.Set("cluster_version", n.Spec.Version),
		d.Set("cluster_type", n.Spec.Type),
		d.Set("description", n.Spec.Description),
		d.Set("vpc_id", n.Spec.HostNetwork.VpcId),
		d.Set("subnet_id", n.Spec.HostNetwork.SubnetId),
		d.Set("highway_subnet_id", n.Spec.HostNetwork.HighwaySubnet),
		d.Set("container_network_type", n.Spec.ContainerNetwork.Mode),
		d.Set("container_network_cidr", flattenContainerNetworkCidrs(n.Spec.ContainerNetwork)),
		d.Set("eni_subnet_id", flattenEniSubnetID(n.Spec.EniNetwork)),
		d.Set("eni_subnet_cidr", n.Spec.EniNetwork.Cidr),
		d.Set("authentication_mode", n.Spec.Authentication.Mode),
		d.Set("security_group_id", n.Spec.HostNetwork.SecurityGroup),
		d.Set("enterprise_project_id", n.Spec.ExtendParam["enterpriseProjectId"]),
		d.Set("service_network_cidr", n.Spec.KubernetesSvcIPRange),
		d.Set("billing_mode", n.Spec.BillingMode),
		d.Set("tags", utils.TagsToMapWithIndicator(n.Spec.ClusterTags, sysTagsBL, d.Get("tags").(map[string]interface{}))),
	)

	if n.Spec.BillingMode != 0 {
		mErr = multierror.Append(mErr, d.Set("charging_mode", "prePaid"))
	}

	r := clusters.GetCert(cceClient, d.Id())

	kubeConfigRaw, err := utils.JsonMarshal(r.Body)

	if err != nil {
		log.Printf("error marshaling r.Body: %s", err)
	}

	mErr = multierror.Append(mErr, d.Set("kube_config_raw", string(kubeConfigRaw)))

	cert, err := r.Extract()

	if err != nil {
		log.Printf("error retrieving CCE cluster certificate: %s", err)
	}

	//Set Certificate Clusters
	var clusterList []map[string]interface{}
	for _, clusterObj := range cert.Clusters {
		clusterCert := make(map[string]interface{})
		clusterCert["name"] = clusterObj.Name
		clusterCert["server"] = clusterObj.Cluster.Server
		clusterCert["certificate_authority_data"] = clusterObj.Cluster.CertAuthorityData
		clusterList = append(clusterList, clusterCert)
	}
	mErr = multierror.Append(mErr, d.Set("certificate_clusters", clusterList))

	//Set Certificate Users
	var userList []map[string]interface{}
	for _, userObj := range cert.Users {
		userCert := make(map[string]interface{})
		userCert["name"] = userObj.Name
		userCert["client_certificate_data"] = userObj.User.ClientCertData
		userCert["client_key_data"] = userObj.User.ClientKeyData
		userList = append(userList, userCert)
	}
	mErr = multierror.Append(mErr, d.Set("certificate_users", userList))

	// Set masters
	var masterList []map[string]interface{}
	for _, masterObj := range n.Spec.Masters {
		master := make(map[string]interface{})
		master["availability_zone"] = masterObj.MasterAZ
		masterList = append(masterList, master)
	}
	mErr = multierror.Append(mErr, d.Set("masters", masterList))

	if err = mErr.ErrorOrNil(); err != nil {
		return diag.Errorf("error setting CCE cluster fields: %s", err)
	}

	return nil
}

func flattenContainerNetworkCidrs(containerNetwork clusters.ContainerNetworkSpec) string {
	cidrs := containerNetwork.Cidrs
	if len(cidrs) != 0 {
		cidrList := make([]string, len(cidrs))
		for i, v := range cidrs {
			cidrList[i] = v.Cidr
		}

		return strings.Join(cidrList, ",")
	}

	return containerNetwork.Cidr
}

func flattenEniSubnetID(eniNetwork *clusters.EniNetworkSpec) string {
	if eniNetwork == nil {
		return ""
	}

	subnets := eniNetwork.Subnets
	subnetIDs := make([]string, len(subnets))
	for i, v := range subnets {
		subnetIDs[i] = v.SubnetID
	}

	return strings.Join(subnetIDs, ",")
}

func resourceClusterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	cceClient, err := config.CceV3Client(config.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating CCE v3 client: %s", err)
	}

	var updateOpts = clusters.UpdateOpts{}

	if d.HasChanges("description") {
		updateOpts.Spec.Description = d.Get("description").(string)
	}

	if d.HasChanges("eni_subnet_id") {
		updateOpts.Spec.EniNetwork = buildEniNetworkOpts(d.Get("eni_subnet_id").(string))
	}

	if d.HasChange("security_group_id") {
		updateOpts.Spec.HostNetwork = &clusters.UpdateHostNetworkSpec{
			SecurityGroup: d.Get("security_group_id").(string),
		}
	}

	if !reflect.DeepEqual(updateOpts, clusters.UpdateOpts{}) {
		_, err = clusters.Update(cceClient, d.Id(), updateOpts).Extract()
		if err != nil {
			return diag.Errorf("error updating CCE cluster: %s", err)
		}
	}

	if d.HasChange("hibernate") {
		if d.Get("hibernate").(bool) {
			err = resourceClusterHibernate(ctx, d, cceClient)
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			err = resourceClusterAwake(ctx, d, cceClient)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if d.HasChange("eip") {
		eipClient, err := config.NetworkingV1Client(config.GetRegion(d))
		if err != nil {
			return diag.Errorf("error creating VPC v1 client: %s", err)
		}

		oldEip, newEip := d.GetChange("eip")
		if oldEip.(string) != "" {
			err = resourceClusterEipAction(cceClient, eipClient, d.Id(), oldEip.(string), "unbind")
			if err != nil {
				return diag.FromErr(err)
			}
		}
		if newEip.(string) != "" {
			err = resourceClusterEipAction(cceClient, eipClient, d.Id(), newEip.(string), "bind")
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if d.HasChange("auto_renew") {
		bssClient, err := config.BssV2Client(config.GetRegion(d))
		if err != nil {
			return diag.Errorf("error creating BSS v2 client: %s", err)
		}
		if err = common.UpdateAutoRenew(bssClient, d.Get("auto_renew").(string), d.Id()); err != nil {
			return diag.Errorf("error updating the auto-renew of the CCE cluster (%s): %s", d.Id(), err)
		}
	}

	return resourceClusterRead(ctx, d, meta)
}

func resourceClusterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	cceClient, err := config.CceV3Client(config.GetRegion(d))
	if err != nil {
		return diag.Errorf("error creating CCE v3 client: %s", err)
	}

	// for prePaid mode, we should unsubscribe the resource
	if d.Get("charging_mode").(string) == "prePaid" || d.Get("billing_mode").(int) == 1 {
		if err := common.UnsubscribePrePaidResource(d, config, []string{d.Id()}); err != nil {
			return diag.Errorf("error unsubscribing CCE cluster: %s", err)
		}
	} else {
		deleteOpts := clusters.DeleteOpts{}
		if v, ok := d.GetOk("delete_all"); ok && v.(string) != "false" {
			deleteOpt := d.Get("delete_all").(string)
			deleteOpts.DeleteEfs = deleteOpt
			deleteOpts.DeleteEvs = deleteOpt
			deleteOpts.DeleteObs = deleteOpt
			deleteOpts.DeleteSfs = deleteOpt
		} else {
			deleteOpts.DeleteEfs = d.Get("delete_efs").(string)
			deleteOpts.DeleteENI = d.Get("delete_eni").(string)
			deleteOpts.DeleteEvs = d.Get("delete_evs").(string)
			deleteOpts.DeleteNet = d.Get("delete_net").(string)
			deleteOpts.DeleteObs = d.Get("delete_obs").(string)
			deleteOpts.DeleteSfs = d.Get("delete_sfs").(string)
		}
		err = clusters.DeleteWithOpts(cceClient, d.Id(), deleteOpts).ExtractErr()
		if err != nil {
			return diag.Errorf("error deleting CCE cluster: %s", err)
		}
	}

	stateConf := &resource.StateChangeConf{
		// The statuses of pending phase includes "Deleting", "Available" and "Unavailable".
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      clusterStateRefreshFunc(cceClient, d.Id(), nil),
		Timeout:      d.Timeout(schema.TimeoutDelete),
		Delay:        60 * time.Second,
		PollInterval: 20 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(ctx)

	if err != nil {
		return diag.Errorf("error deleting CCE cluster: %s", err)
	}

	d.SetId("")
	return nil
}

func clusterStateRefreshFunc(cceClient *golangsdk.ServiceClient, clusterId string,
	targets []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Expect the status of CCE cluster to be any one of the status list: %v", targets)
		resp, err := clusters.Get(cceClient, clusterId).Extract()
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				log.Printf("[DEBUG] The cluster (%s) has been deleted", clusterId)
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

func getClusterIDFromJob(ctx context.Context, client *golangsdk.ServiceClient, jobID string, timeout time.Duration) (string, error) {
	stateJob := &resource.StateChangeConf{
		Pending:      []string{"Initializing", "Running"},
		Target:       []string{"Success"},
		Refresh:      waitForJobStatus(client, jobID),
		Timeout:      timeout,
		Delay:        150 * time.Second,
		PollInterval: 20 * time.Second,
	}

	v, err := stateJob.WaitForStateContext(ctx)
	if err != nil {
		if job, ok := v.(*nodes.Job); ok {
			return "", fmt.Errorf("error waiting for job (%s) to become success: %s, reason: %s",
				jobID, err, job.Status.Reason)
		} else {
			return "", fmt.Errorf("error waiting for job (%s) to become success: %s", jobID, err)
		}

	}

	job := v.(*nodes.Job)
	clusterID := job.Spec.ClusterID
	if clusterID == "" {
		return "", fmt.Errorf("error fetching CCE cluster ID")
	}
	return clusterID, nil
}

func resourceClusterHibernate(ctx context.Context, d *schema.ResourceData, cceClient *golangsdk.ServiceClient) error {
	clusterID := d.Id()
	err := clusters.Operation(cceClient, clusterID, "hibernate").ExtractErr()
	if err != nil {
		return fmt.Errorf("error hibernating CCE cluster: %s", err)
	}

	log.Printf("[DEBUG] Waiting for CCE cluster (%s) to become hibernate", clusterID)
	stateConf := &resource.StateChangeConf{
		// The statuses of pending phase includes "Available" and "Hibernating".
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      clusterStateRefreshFunc(cceClient, clusterID, []string{"Hibernation"}),
		Timeout:      d.Timeout(schema.TimeoutUpdate),
		Delay:        20 * time.Second,
		PollInterval: 20 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return fmt.Errorf("error hibernating CCE cluster: %s", err)
	}
	return nil
}

func resourceClusterAwake(ctx context.Context, d *schema.ResourceData, cceClient *golangsdk.ServiceClient) error {
	clusterID := d.Id()
	err := clusters.Operation(cceClient, clusterID, "awake").ExtractErr()
	if err != nil {
		return fmt.Errorf("error awaking CCE cluster: %s", err)
	}

	log.Printf("[DEBUG] Waiting for CCE cluster (%s) to become available", clusterID)
	stateConf := &resource.StateChangeConf{
		// The statuses of pending phase include "Awaking".
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      clusterStateRefreshFunc(cceClient, clusterID, []string{"Available"}),
		Timeout:      d.Timeout(schema.TimeoutUpdate),
		Delay:        100 * time.Second,
		PollInterval: 20 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return fmt.Errorf("error awaking CCE cluster: %s", err)
	}
	return nil
}

func resourceClusterEipAction(cceClient, eipClient *golangsdk.ServiceClient,
	clusterID, eip, action string) error {
	eipID, err := common.GetEipIDbyAddress(eipClient, eip, "all_granted_eps")
	if err != nil {
		return fmt.Errorf("error fetching EIP ID: %s", err)
	}

	opts := clusters.UpdateIpOpts{
		Action: action,
		Spec: clusters.IpSpec{
			ID: eipID,
		},
	}

	err = clusters.UpdateMasterIp(cceClient, clusterID, opts).ExtractErr()
	if err != nil {
		return fmt.Errorf("error %sing the public IP of CCE cluster: %s", action, err)
	}
	return nil
}
