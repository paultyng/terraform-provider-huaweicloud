package huaweicloud

import (
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/aom/v1/icagents"
	"github.com/huaweicloud/golangsdk/openstack/cce/v3/clusters"
	"github.com/huaweicloud/golangsdk/openstack/cce/v3/nodes"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils/fmtp"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils/logp"
)

var associateDeleteSchema *schema.Schema = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
	ValidateFunc: validation.StringInSlice([]string{
		"true", "try", "false",
	}, true),
	ConflictsWith: []string{"delete_all"},
}

func ResourceCCEClusterV3() *schema.Resource {
	return &schema.Resource{
		Create: resourceCCEClusterV3Create,
		Read:   resourceCCEClusterV3Read,
		Update: resourceCCEClusterV3Update,
		Delete: resourceCCEClusterV3Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"highway_subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				RequiredWith: []string{"eni_subnet_cidr"},
			},
			"eni_subnet_cidr": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				RequiredWith: []string{"eni_subnet_id"},
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
				ForceNew:     true,
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

			// charge info: charging_mode, period_unit, period, auto_renew
			"charging_mode": schemeChargingMode(nil),
			"period_unit":   schemaPeriodUnit(nil),
			"period":        schemaPeriod(nil),
			"auto_renew":    schemaAutoRenew(nil),

			"delete_efs": associateDeleteSchema,
			"delete_eni": associateDeleteSchema,
			"delete_evs": associateDeleteSchema,
			"delete_net": associateDeleteSchema,
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
			"security_group_id": {
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

func resourceClusterLabelsV3(d *schema.ResourceData) map[string]string {
	m := make(map[string]string)
	for key, val := range d.Get("labels").(map[string]interface{}) {
		m[key] = val.(string)
	}
	return m
}
func resourceClusterAnnotationsV3(d *schema.ResourceData) map[string]string {
	m := make(map[string]string)
	for key, val := range d.Get("annotations").(map[string]interface{}) {
		m[key] = val.(string)
	}
	return m
}

func resourceClusterExtendParamV3(d *schema.ResourceData, config *config.Config) map[string]interface{} {
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
		extendParam["isAutoPay"] = "true"
		extendParam["isAutoRenew"] = "false"
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

	epsID := GetEnterpriseProjectID(d, config)
	if epsID != "" {
		extendParam["enterpriseProjectId"] = epsID
	}

	return extendParam
}

func resourceClusterMastersV3(d *schema.ResourceData) ([]clusters.MasterSpec, error) {
	if v, ok := d.GetOk("masters"); ok {
		flavorId := d.Get("flavor_id").(string)
		mastersRaw := v.([]interface{})
		if strings.Contains(flavorId, "s1") && len(mastersRaw) != 1 {
			return nil, fmtp.Errorf("Error creating HuaweiCloud Cluster: "+
				"single-master cluster need 1 az for master node, but got %d", len(mastersRaw))
		}
		if strings.Contains(flavorId, "s2") && len(mastersRaw) != 3 {
			return nil, fmtp.Errorf("Error creating HuaweiCloud Cluster: "+
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

func resourceCCEClusterV3Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*config.Config)
	cceClient, err := config.CceV3Client(GetRegion(d, config))
	if err != nil {
		return fmtp.Errorf("Unable to create HuaweiCloud CCE client : %s", err)
	}
	icAgentClient, err := config.AomV1Client(GetRegion(d, config))
	if err != nil {
		return fmtp.Errorf("Unable to create HuaweiCloud AOM client : %s", err)
	}

	authenticating_proxy := make(map[string]string)
	if hasFilledOpt(d, "authenticating_proxy_ca") {
		authenticating_proxy["ca"] = utils.EncodeBase64IfNot(d.Get("authenticating_proxy_ca").(string))
		authenticating_proxy["cert"] = utils.EncodeBase64IfNot(d.Get("authenticating_proxy_cert").(string))
		authenticating_proxy["privateKey"] = utils.EncodeBase64IfNot(d.Get("authenticating_proxy_private_key").(string))
	}

	billingMode := 0
	if d.Get("charging_mode").(string) == "prePaid" || d.Get("billing_mode").(int) == 1 {
		billingMode = 1
		if err := validatePrePaidChargeInfo(d); err != nil {
			return err
		}
	}

	clusterName := d.Get("name").(string)
	createOpts := clusters.CreateOpts{
		Kind:       "Cluster",
		ApiVersion: "v3",
		Metadata: clusters.CreateMetaData{
			Name:        clusterName,
			Labels:      resourceClusterLabelsV3(d),
			Annotations: resourceClusterAnnotationsV3(d)},
		Spec: clusters.Spec{
			Type:        d.Get("cluster_type").(string),
			Flavor:      d.Get("flavor_id").(string),
			Version:     d.Get("cluster_version").(string),
			Description: d.Get("description").(string),
			HostNetwork: clusters.HostNetworkSpec{
				VpcId:         d.Get("vpc_id").(string),
				SubnetId:      d.Get("subnet_id").(string),
				HighwaySubnet: d.Get("highway_subnet_id").(string),
			},
			ContainerNetwork: clusters.ContainerNetworkSpec{
				Mode: d.Get("container_network_type").(string),
				Cidr: d.Get("container_network_cidr").(string),
			},
			Authentication: clusters.AuthenticationSpec{
				Mode:                d.Get("authentication_mode").(string),
				AuthenticatingProxy: authenticating_proxy,
			},
			BillingMode:          billingMode,
			ExtendParam:          resourceClusterExtendParamV3(d, config),
			KubernetesSvcIPRange: d.Get("service_network_cidr").(string),
		},
	}

	if _, ok := d.GetOk("eni_subnet_id"); ok {
		eniNetwork := clusters.EniNetworkSpec{
			SubnetId: d.Get("eni_subnet_id").(string),
			Cidr:     d.Get("eni_subnet_cidr").(string),
		}
		createOpts.Spec.EniNetwork = &eniNetwork
	}

	masters, err := resourceClusterMastersV3(d)
	if err != nil {
		return err
	}
	createOpts.Spec.Masters = masters

	s, err := clusters.Create(cceClient, createOpts).Extract()
	if err != nil {
		return fmtp.Errorf("Error creating HuaweiCloud Cluster: %s", err)
	}

	jobID := s.Status.JobID
	if jobID == "" {
		return fmtp.Errorf("Error fetching job id after creating cce cluster: %s", clusterName)
	}

	clusterID, err := getCCEClusterIDFromJob(cceClient, jobID)
	if err != nil {
		return err
	}
	d.SetId(clusterID)

	logp.Printf("[DEBUG] Waiting for HuaweiCloud CCE cluster (%s) to become available", clusterID)
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"Creating"},
		Target:       []string{"Available"},
		Refresh:      waitForCCEClusterActive(cceClient, clusterID),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        150 * time.Second,
		PollInterval: 20 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmtp.Errorf("Error creating HuaweiCloud CCE cluster: %s", err)
	}

	logp.Printf("[DEBUG] installing ICAgent for CCE cluster (%s)", d.Id())
	installParam := icagents.InstallParam{
		ClusterId: d.Id(),
		NameSpace: "default",
	}
	result := icagents.Create(icAgentClient, installParam)
	if result.Err != nil {
		logp.Printf("Error installing ICAgent in CCE cluster: %s", result.Err)
	}

	// create a hibernating cluster
	if d.Get("hibernate").(bool) {
		err = resourceCCEClusterV3Hibernate(d, cceClient)
		if err != nil {
			return err
		}
	}

	return resourceCCEClusterV3Read(d, meta)
}

func resourceCCEClusterV3Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*config.Config)
	cceClient, err := config.CceV3Client(GetRegion(d, config))
	if err != nil {
		return fmtp.Errorf("Error creating HuaweiCloud CCE client: %s", err)
	}

	n, err := clusters.Get(cceClient, d.Id()).Extract()
	if err != nil {
		if _, ok := err.(golangsdk.ErrDefault404); ok {
			d.SetId("")
			return nil
		}

		return fmtp.Errorf("Error retrieving HuaweiCloud CCE: %s", err)
	}

	d.Set("name", n.Metadata.Name)
	d.Set("status", n.Status.Phase)
	d.Set("flavor_id", n.Spec.Flavor)
	d.Set("cluster_version", n.Spec.Version)
	d.Set("cluster_type", n.Spec.Type)
	d.Set("description", n.Spec.Description)
	d.Set("vpc_id", n.Spec.HostNetwork.VpcId)
	d.Set("subnet_id", n.Spec.HostNetwork.SubnetId)
	d.Set("highway_subnet_id", n.Spec.HostNetwork.HighwaySubnet)
	d.Set("container_network_type", n.Spec.ContainerNetwork.Mode)
	d.Set("container_network_cidr", n.Spec.ContainerNetwork.Cidr)
	d.Set("eni_subnet_id", n.Spec.EniNetwork.SubnetId)
	d.Set("eni_subnet_cidr", n.Spec.EniNetwork.Cidr)
	d.Set("authentication_mode", n.Spec.Authentication.Mode)
	d.Set("security_group_id", n.Spec.HostNetwork.SecurityGroup)
	d.Set("region", GetRegion(d, config))
	d.Set("enterprise_project_id", n.Spec.ExtendParam["enterpriseProjectId"])
	d.Set("service_network_cidr", n.Spec.KubernetesSvcIPRange)
	d.Set("billing_mode", n.Spec.BillingMode)
	if n.Spec.BillingMode != 0 {
		d.Set("charging_mode", "prePaid")
	}

	r := clusters.GetCert(cceClient, d.Id())

	kubeConfigRaw, err := utils.JsonMarshal(r.Body)

	if err != nil {
		logp.Printf("Error marshaling r.Body: %s", err)
	}

	d.Set("kube_config_raw", string(kubeConfigRaw))

	cert, err := r.Extract()

	if err != nil {
		logp.Printf("Error retrieving HuaweiCloud CCE cluster cert: %s", err)
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
	d.Set("certificate_clusters", clusterList)

	//Set Certificate Users
	var userList []map[string]interface{}
	for _, userObj := range cert.Users {
		userCert := make(map[string]interface{})
		userCert["name"] = userObj.Name
		userCert["client_certificate_data"] = userObj.User.ClientCertData
		userCert["client_key_data"] = userObj.User.ClientKeyData
		userList = append(userList, userCert)
	}
	d.Set("certificate_users", userList)

	// Set masters
	var masterList []map[string]interface{}
	for _, masterObj := range n.Spec.Masters {
		master := make(map[string]interface{})
		master["availability_zone"] = masterObj.MasterAZ
		masterList = append(masterList, master)
	}
	d.Set("masters", masterList)

	return nil
}

func resourceCCEClusterV3Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*config.Config)
	cceClient, err := config.CceV3Client(GetRegion(d, config))
	if err != nil {
		return fmtp.Errorf("Error creating HuaweiCloud CCE Client: %s", err)
	}

	if d.HasChange("description") {
		var updateOpts clusters.UpdateOpts
		updateOpts.Spec.Description = d.Get("description").(string)
		_, err = clusters.Update(cceClient, d.Id(), updateOpts).Extract()

		if err != nil {
			return fmtp.Errorf("Error updating HuaweiCloud CCE: %s", err)
		}
	}

	if d.HasChange("hibernate") {
		if d.Get("hibernate").(bool) {
			err = resourceCCEClusterV3Hibernate(d, cceClient)
			if err != nil {
				return err
			}
		} else {
			err = resourceCCEClusterV3Awake(d, cceClient)
			if err != nil {
				return err
			}
		}
	}

	return resourceCCEClusterV3Read(d, meta)
}

func resourceCCEClusterV3Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*config.Config)
	cceClient, err := config.CceV3Client(GetRegion(d, config))
	if err != nil {
		return fmtp.Errorf("Error creating HuaweiCloud CCE Client: %s", err)
	}

	// for prePaid mode, we should unsubscribe the resource
	if d.Get("charging_mode").(string) == "prePaid" || d.Get("billing_mode").(int) == 1 {
		if err := UnsubscribePrePaidResource(d, config, []string{d.Id()}); err != nil {
			return fmtp.Errorf("Error unsubscribing HuaweiCloud CCE cluster: %s", err)
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
			return fmtp.Errorf("Error deleting HuaweiCloud CCE Cluster: %s", err)
		}
	}

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"Deleting", "Available", "Unavailable"},
		Target:       []string{"Deleted"},
		Refresh:      waitForCCEClusterDelete(cceClient, d.Id()),
		Timeout:      d.Timeout(schema.TimeoutDelete),
		Delay:        60 * time.Second,
		PollInterval: 20 * time.Second,
	}

	_, err = stateConf.WaitForState()

	if err != nil {
		return fmtp.Errorf("Error deleting HuaweiCloud CCE cluster: %s", err)
	}

	d.SetId("")
	return nil
}

func waitForCCEClusterActive(cceClient *golangsdk.ServiceClient, clusterId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		n, err := clusters.Get(cceClient, clusterId).Extract()
		if err != nil {
			return nil, "", err
		}

		return n, n.Status.Phase, nil
	}
}

func waitForCCEClusterDelete(cceClient *golangsdk.ServiceClient, clusterId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		logp.Printf("[DEBUG] Attempting to delete HuaweiCloud CCE cluster %s.\n", clusterId)

		r, err := clusters.Get(cceClient, clusterId).Extract()

		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				logp.Printf("[DEBUG] Successfully deleted HuaweiCloud CCE cluster %s", clusterId)
				return r, "Deleted", nil
			}
		}
		if r.Status.Phase == "Deleting" {
			return r, "Deleting", nil
		}
		logp.Printf("[DEBUG] HuaweiCloud CCE cluster %s still available.\n", clusterId)
		return r, "Available", nil
	}
}

func getCCEClusterIDFromJob(client *golangsdk.ServiceClient, jobID string) (string, error) {
	stateJob := &resource.StateChangeConf{
		Pending: []string{"Initializing"},
		Target:  []string{"Running", "Success"},
		Refresh: waitForJobStatus(client, jobID),
		Timeout: 5 * time.Minute,
		// waiting for 35 seconds to avoid 401 response code
		Delay:        35 * time.Second,
		PollInterval: 10 * time.Second,
	}

	v, err := stateJob.WaitForState()
	if err != nil {
		return "", fmtp.Errorf("Error waiting for job (%s) to become running: %s", jobID, err)
	}

	job := v.(*nodes.Job)
	clusterID := job.Spec.ClusterID
	if clusterID == "" {
		return "", fmtp.Errorf("Error fetching CCE cluster id")
	}
	return clusterID, nil
}

func resourceCCEClusterV3Hibernate(d *schema.ResourceData, cceClient *golangsdk.ServiceClient) error {
	clusterID := d.Id()
	err := clusters.Operation(cceClient, clusterID, "hibernate").ExtractErr()
	if err != nil {
		return fmtp.Errorf("Error hibernating HuaweiCloud CCE cluster: %s", err)
	}

	logp.Printf("[DEBUG] Waiting for HuaweiCloud CCE cluster (%s) to become hibernating", clusterID)
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"Available"},
		Target:       []string{"Hibernation"},
		Refresh:      waitForCCEClusterActive(cceClient, clusterID),
		Timeout:      d.Timeout(schema.TimeoutUpdate),
		Delay:        100 * time.Second,
		PollInterval: 20 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmtp.Errorf("Error hibernating HuaweiCloud CCE cluster: %s", err)
	}
	return nil
}

func resourceCCEClusterV3Awake(d *schema.ResourceData, cceClient *golangsdk.ServiceClient) error {
	clusterID := d.Id()
	err := clusters.Operation(cceClient, clusterID, "awake").ExtractErr()
	if err != nil {
		return fmtp.Errorf("Error awakting HuaweiCloud CCE cluster: %s", err)
	}

	logp.Printf("[DEBUG] Waiting for HuaweiCloud CCE cluster (%s) to become available", clusterID)
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"Awaking"},
		Target:       []string{"Available"},
		Refresh:      waitForCCEClusterActive(cceClient, clusterID),
		Timeout:      d.Timeout(schema.TimeoutUpdate),
		Delay:        100 * time.Second,
		PollInterval: 20 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmtp.Errorf("Error awaking HuaweiCloud CCE cluster: %s", err)
	}
	return nil
}
