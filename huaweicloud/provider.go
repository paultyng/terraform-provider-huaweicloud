package huaweicloud

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/mutexkv"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/apig"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/deprecated"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dli"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/elb"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/fgs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/lb"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/mrs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/vpc"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/waf"
)

const defaultCloud string = "myhuaweicloud.com"

// This is a global MutexKV for use within this plugin.
var osMutexKV = mutexkv.NewMutexKV()

// Provider returns a schema.Provider for HuaweiCloud.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"region": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  descriptions["region"],
				InputDefault: "cn-north-1",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_REGION_NAME",
					"OS_REGION_NAME",
				}, nil),
			},

			"access_key": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  descriptions["access_key"],
				RequiredWith: []string{"secret_key"},
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_ACCESS_KEY",
					"OS_ACCESS_KEY",
				}, nil),
			},

			"secret_key": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  descriptions["secret_key"],
				RequiredWith: []string{"access_key"},
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_SECRET_KEY",
					"OS_SECRET_KEY",
				}, nil),
			},

			"security_token": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  descriptions["security_token"],
				RequiredWith: []string{"access_key"},
				DefaultFunc:  schema.EnvDefaultFunc("HW_SECURITY_TOKEN", nil),
			},

			"domain_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["domain_id"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_DOMAIN_ID",
					"OS_DOMAIN_ID",
					"OS_USER_DOMAIN_ID",
					"OS_PROJECT_DOMAIN_ID",
				}, ""),
			},

			"domain_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["domain_name"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_DOMAIN_NAME",
					"OS_DOMAIN_NAME",
					"OS_USER_DOMAIN_NAME",
					"OS_PROJECT_DOMAIN_NAME",
				}, ""),
			},

			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["user_name"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_USER_NAME",
					"OS_USERNAME",
				}, ""),
			},

			"user_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["user_id"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_USER_ID",
					"OS_USER_ID",
				}, ""),
			},

			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: descriptions["password"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_USER_PASSWORD",
					"OS_PASSWORD",
				}, ""),
			},

			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["project_id"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_PROJECT_ID",
					"OS_PROJECT_ID",
				}, nil),
			},

			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["project_name"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_PROJECT_NAME",
					"OS_PROJECT_NAME",
				}, nil),
			},

			"tenant_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["tenant_id"],
				DefaultFunc: schema.EnvDefaultFunc("OS_TENANT_ID", ""),
			},

			"tenant_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["tenant_name"],
				DefaultFunc: schema.EnvDefaultFunc("OS_TENANT_NAME", ""),
			},

			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["token"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_AUTH_TOKEN",
					"OS_AUTH_TOKEN",
				}, ""),
			},

			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions["insecure"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_INSECURE",
					"OS_INSECURE",
				}, false),
			},

			"cacert_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CACERT", ""),
				Description: descriptions["cacert_file"],
			},

			"cert": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CERT", ""),
				Description: descriptions["cert"],
			},

			"key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_KEY", ""),
				Description: descriptions["key"],
			},

			"agency_name": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("OS_AGENCY_NAME", nil),
				Description:  descriptions["agency_name"],
				RequiredWith: []string{"agency_domain_name"},
			},

			"agency_domain_name": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("OS_AGENCY_DOMAIN_NAME", nil),
				Description:  descriptions["agency_domain_name"],
				RequiredWith: []string{"agency_name"},
			},

			"delegated_project": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DELEGATED_PROJECT", ""),
				Description: descriptions["delegated_project"],
			},

			"auth_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["auth_url"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HW_AUTH_URL",
					"OS_AUTH_URL",
				}, nil),
			},

			"cloud": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["cloud"],
				DefaultFunc: schema.EnvDefaultFunc("HW_CLOUD", defaultCloud),
			},

			"endpoints": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: descriptions["endpoints"],
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"enterprise_project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["enterprise_project_id"],
				DefaultFunc: schema.EnvDefaultFunc("HW_ENTERPRISE_PROJECT_ID", ""),
			},

			"max_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: descriptions["max_retries"],
				DefaultFunc: schema.EnvDefaultFunc("HW_MAX_RETRIES", 5),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"huaweicloud_antiddos":                             dataSourceAntiDdosV1(),
			"huaweicloud_availability_zones":                   DataSourceAvailabilityZones(),
			"huaweicloud_cce_addon_template":                   DataSourceCCEAddonTemplateV3(),
			"huaweicloud_cce_cluster":                          DataSourceCCEClusterV3(),
			"huaweicloud_cce_node":                             DataSourceCCENodeV3(),
			"huaweicloud_cce_node_pool":                        DataSourceCCENodePoolV3(),
			"huaweicloud_cdm_flavors":                          DataSourceCdmFlavorV1(),
			"huaweicloud_compute_flavors":                      DataSourceEcsFlavors(),
			"huaweicloud_compute_instance":                     DataSourceComputeInstance(),
			"huaweicloud_csbs_backup":                          dataSourceCSBSBackupV1(),
			"huaweicloud_csbs_backup_policy":                   dataSourceCSBSBackupPolicyV1(),
			"huaweicloud_dcs_az":                               DataSourceDcsAZV1(),
			"huaweicloud_dcs_maintainwindow":                   DataSourceDcsMaintainWindowV1(),
			"huaweicloud_dcs_product":                          DataSourceDcsProductV1(),
			"huaweicloud_dds_flavors":                          DataSourceDDSFlavorV3(),
			"huaweicloud_dis_partition":                        DataSourceDisPartitionV2(),
			"huaweicloud_dms_az":                               DataSourceDmsAZV1(),
			"huaweicloud_dms_product":                          DataSourceDmsProductV1(),
			"huaweicloud_dms_maintainwindow":                   DataSourceDmsMaintainWindowV1(),
			"huaweicloud_elb_flavors":                          dataSourceElbFlavorsV3(),
			"huaweicloud_enterprise_project":                   DataSourceEnterpriseProject(),
			"huaweicloud_fgs_dependencies":                     fgs.DataSourceFunctionGraphDependencies(),
			"huaweicloud_gaussdb_cassandra_dedicated_resource": dataSourceGeminiDBDehResource(),
			"huaweicloud_gaussdb_cassandra_instance":           dataSourceGeminiDBInstance(),
			"huaweicloud_gaussdb_cassandra_instances":          dataSourceGeminiDBInstances(),
			"huaweicloud_gaussdb_opengauss_instance":           dataSourceOpenGaussInstance(),
			"huaweicloud_gaussdb_mysql_configuration":          dataSourceGaussdbMysqlConfigurations(),
			"huaweicloud_gaussdb_mysql_dedicated_resource":     dataSourceGaussDBMysqlDehResource(),
			"huaweicloud_gaussdb_mysql_flavors":                dataSourceGaussdbMysqlFlavors(),
			"huaweicloud_gaussdb_mysql_instance":               dataSourceGaussDBMysqlInstance(),
			"huaweicloud_gaussdb_mysql_instances":              dataSourceGaussDBMysqlInstances(),
			"huaweicloud_gaussdb_redis_instance":               dataSourceGaussRedisInstance(),
			"huaweicloud_identity_role":                        DataSourceIdentityRoleV3(),
			"huaweicloud_identity_custom_role":                 DataSourceIdentityCustomRole(),
			"huaweicloud_iec_eips":                             dataSourceIECNetworkEips(),
			"huaweicloud_iec_flavors":                          dataSourceIecFlavors(),
			"huaweicloud_iec_images":                           dataSourceIecImages(),
			"huaweicloud_iec_keypair":                          dataSourceIECKeypair(),
			"huaweicloud_iec_network_acl":                      dataSourceIECNetworkACL(),
			"huaweicloud_iec_port":                             DataSourceIECPort(),
			"huaweicloud_iec_security_group":                   dataSourceIECSecurityGroup(),
			"huaweicloud_iec_server":                           dataSourceIECServer(),
			"huaweicloud_iec_sites":                            dataSourceIecSites(),
			"huaweicloud_iec_vpc":                              DataSourceIECVpc(),
			"huaweicloud_iec_vpc_subnets":                      DataSourceIECVpcSubnets(),
			"huaweicloud_images_image":                         DataSourceImagesImageV2(),
			"huaweicloud_kms_key":                              DataSourceKmsKeyV1(),
			"huaweicloud_kms_data_key":                         DataSourceKmsDataKeyV1(),
			"huaweicloud_lb_loadbalancer":                      DataSourceELBV2Loadbalancer(),
			"huaweicloud_lb_certificate":                       lb.DataSourceLBCertificateV2(),
			"huaweicloud_elb_certificate":                      elb.DataSourceELBCertificateV3(),
			"huaweicloud_nat_gateway":                          DataSourceNatGatewayV2(),
			"huaweicloud_networking_port":                      DataSourceNetworkingPortV2(),
			"huaweicloud_networking_secgroup":                  DataSourceNetworkingSecGroupV2(),
			"huaweicloud_obs_bucket_object":                    DataSourceObsBucketObject(),
			"huaweicloud_rds_flavors":                          DataSourceRdsFlavorV3(),
			"huaweicloud_sfs_file_system":                      DataSourceSFSFileSystemV2(),
			"huaweicloud_vbs_backup_policy":                    dataSourceVBSBackupPolicyV2(),
			"huaweicloud_vbs_backup":                           dataSourceVBSBackupV2(),
			"huaweicloud_vpc":                                  DataSourceVirtualPrivateCloudVpcV1(),
			"huaweicloud_vpc_bandwidth":                        DataSourceBandWidth(),
			"huaweicloud_vpc_eip":                              DataSourceVpcEip(),
			"huaweicloud_vpc_ids":                              dataSourceVirtualPrivateCloudVpcIdsV1(),
			"huaweicloud_vpc_peering_connection":               dataSourceVpcPeeringConnectionV2(),
			"huaweicloud_vpc_route":                            DataSourceVPCRouteV2(),
			"huaweicloud_vpc_route_ids":                        dataSourceVPCRouteIdsV2(),
			"huaweicloud_vpc_subnet":                           DataSourceVpcSubnetV1(),
			"huaweicloud_vpc_subnet_ids":                       DataSourceVpcSubnetIdsV1(),
			"huaweicloud_vpcep_public_services":                DataSourceVPCEPPublicServices(),
			"huaweicloud_waf_certificate":                      waf.DataSourceWafCertificateV1(),
			"huaweicloud_waf_policies":                         waf.DataSourceWafPoliciesV1(),
			"huaweicloud_waf_dedicated_instances":              waf.DataSourceWafDedicatedInstancesV1(),

			// Legacy
			"huaweicloud_images_image_v2":           DataSourceImagesImageV2(),
			"huaweicloud_networking_port_v2":        DataSourceNetworkingPortV2(),
			"huaweicloud_networking_secgroup_v2":    DataSourceNetworkingSecGroupV2(),
			"huaweicloud_kms_key_v1":                DataSourceKmsKeyV1(),
			"huaweicloud_kms_data_key_v1":           DataSourceKmsDataKeyV1(),
			"huaweicloud_rds_flavors_v3":            DataSourceRdsFlavorV3(),
			"huaweicloud_sfs_file_system_v2":        DataSourceSFSFileSystemV2(),
			"huaweicloud_vpc_v1":                    DataSourceVirtualPrivateCloudVpcV1(),
			"huaweicloud_vpc_ids_v1":                dataSourceVirtualPrivateCloudVpcIdsV1(),
			"huaweicloud_vpc_peering_connection_v2": dataSourceVpcPeeringConnectionV2(),
			"huaweicloud_vpc_route_v2":              DataSourceVPCRouteV2(),
			"huaweicloud_vpc_route_ids_v2":          dataSourceVPCRouteIdsV2(),
			"huaweicloud_vpc_subnet_v1":             DataSourceVpcSubnetV1(),
			"huaweicloud_vpc_subnet_ids_v1":         DataSourceVpcSubnetIdsV1(),
			"huaweicloud_cce_cluster_v3":            DataSourceCCEClusterV3(),
			"huaweicloud_cce_node_v3":               DataSourceCCENodeV3(),
			"huaweicloud_csbs_backup_v1":            dataSourceCSBSBackupV1(),
			"huaweicloud_csbs_backup_policy_v1":     dataSourceCSBSBackupPolicyV1(),
			"huaweicloud_dms_az_v1":                 DataSourceDmsAZV1(),
			"huaweicloud_dms_product_v1":            DataSourceDmsProductV1(),
			"huaweicloud_dms_maintainwindow_v1":     DataSourceDmsMaintainWindowV1(),
			"huaweicloud_vbs_backup_policy_v2":      dataSourceVBSBackupPolicyV2(),
			"huaweicloud_vbs_backup_v2":             dataSourceVBSBackupV2(),
			"huaweicloud_antiddos_v1":               dataSourceAntiDdosV1(),
			"huaweicloud_dcs_az_v1":                 DataSourceDcsAZV1(),
			"huaweicloud_dcs_maintainwindow_v1":     DataSourceDcsMaintainWindowV1(),
			"huaweicloud_dcs_product_v1":            DataSourceDcsProductV1(),
			"huaweicloud_dds_flavors_v3":            DataSourceDDSFlavorV3(),
			"huaweicloud_identity_role_v3":          DataSourceIdentityRoleV3(),
			"huaweicloud_cdm_flavors_v1":            DataSourceCdmFlavorV1(),
			"huaweicloud_dis_partition_v2":          DataSourceDisPartitionV2(),
			// Deprecated
			"huaweicloud_compute_availability_zones_v2": dataSourceComputeAvailabilityZonesV2(),
			"huaweicloud_networking_network_v2":         dataSourceNetworkingNetworkV2(),
			"huaweicloud_networking_subnet_v2":          dataSourceNetworkingSubnetV2(),
			"huaweicloud_rts_stack_v1":                  dataSourceRTSStackV1(),
			"huaweicloud_rts_stack_resource_v1":         dataSourceRTSStackResourcesV1(),
			"huaweicloud_rts_software_config_v1":        dataSourceRtsSoftwareConfigV1(),
			"huaweicloud_cts_tracker":                   deprecated.DataSourceCTSTrackerV1(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"huaweicloud_api_gateway_api":                 ResourceAPIGatewayAPI(),
			"huaweicloud_api_gateway_group":               ResourceAPIGatewayGroup(),
			"huaweicloud_apig_api":                        apig.ResourceApigAPIV2(),
			"huaweicloud_apig_instance":                   apig.ResourceApigInstanceV2(),
			"huaweicloud_apig_application":                apig.ResourceApigApplicationV2(),
			"huaweicloud_apig_custom_authorizer":          apig.ResourceApigCustomAuthorizerV2(),
			"huaweicloud_apig_environment":                apig.ResourceApigEnvironmentV2(),
			"huaweicloud_apig_group":                      apig.ResourceApigGroupV2(),
			"huaweicloud_apig_response":                   apig.ResourceApigResponseV2(),
			"huaweicloud_apig_throttling_policy":          apig.ResourceApigThrottlingPolicyV2(),
			"huaweicloud_apig_vpc_channel":                apig.ResourceApigVpcChannelV2(),
			"huaweicloud_as_configuration":                ResourceASConfiguration(),
			"huaweicloud_as_group":                        ResourceASGroup(),
			"huaweicloud_as_lifecycle_hook":               ResourceASLifecycleHook(),
			"huaweicloud_as_policy":                       ResourceASPolicy(),
			"huaweicloud_bms_instance":                    ResourceBmsInstance(),
			"huaweicloud_bcs_instance":                    resourceBCSInstanceV2(),
			"huaweicloud_cbr_policy":                      resourceCBRPolicyV3(),
			"huaweicloud_cbr_vault":                       resourceCBRVaultV3(),
			"huaweicloud_cce_cluster":                     ResourceCCEClusterV3(),
			"huaweicloud_cce_node":                        ResourceCCENodeV3(),
			"huaweicloud_cce_node_attach":                 ResourceCCENodeAttachV3(),
			"huaweicloud_cce_addon":                       ResourceCCEAddonV3(),
			"huaweicloud_cce_node_pool":                   ResourceCCENodePool(),
			"huaweicloud_cci_network":                     resourceCCINetworkV1(),
			"huaweicloud_cci_pvc":                         ResourceCCIPersistentVolumeClaimV1(),
			"huaweicloud_cdm_cluster":                     ResourceCdmClusterV1(),
			"huaweicloud_cdn_domain":                      resourceCdnDomainV1(),
			"huaweicloud_ces_alarmrule":                   ResourceAlarmRule(),
			"huaweicloud_cloudtable_cluster":              resourceCloudtableClusterV2(),
			"huaweicloud_compute_instance":                ResourceComputeInstanceV2(),
			"huaweicloud_compute_interface_attach":        ResourceComputeInterfaceAttachV2(),
			"huaweicloud_compute_keypair":                 ResourceComputeKeypairV2(),
			"huaweicloud_compute_servergroup":             ResourceComputeServerGroupV2(),
			"huaweicloud_compute_eip_associate":           ResourceComputeFloatingIPAssociateV2(),
			"huaweicloud_compute_volume_attach":           ResourceComputeVolumeAttachV2(),
			"huaweicloud_cs_cluster":                      deprecated.ResourceCsClusterV1(),
			"huaweicloud_cs_peering_connect":              deprecated.ResourceCsPeeringConnectV1(),
			"huaweicloud_cs_route":                        deprecated.ResourceCsRouteV1(),
			"huaweicloud_csbs_backup":                     resourceCSBSBackupV1(),
			"huaweicloud_csbs_backup_policy":              resourceCSBSBackupPolicyV1(),
			"huaweicloud_css_cluster":                     ResourceCssClusterV1(),
			"huaweicloud_css_snapshot":                    ResourceCssSnapshot(),
			"huaweicloud_dcs_instance":                    ResourceDcsInstanceV1(),
			"huaweicloud_dds_instance":                    ResourceDdsInstanceV3(),
			"huaweicloud_dis_stream":                      ResourceDisStreamV2(),
			"huaweicloud_dli_queue":                       dli.ResourceDliQueue(),
			"huaweicloud_dms_group":                       ResourceDmsGroupsV1(),
			"huaweicloud_dms_instance":                    ResourceDmsInstancesV1(),
			"huaweicloud_dms_queue":                       ResourceDmsQueuesV1(),
			"huaweicloud_dms_kafka_instance":              ResourceDmsKafkaInstance(),
			"huaweicloud_dms_kafka_topic":                 ResourceDmsKafkaTopic(),
			"huaweicloud_dms_rabbitmq_instance":           ResourceDmsRabbitmqInstance(),
			"huaweicloud_dns_ptrrecord":                   ResourceDNSPtrRecordV2(),
			"huaweicloud_dns_recordset":                   ResourceDNSRecordSetV2(),
			"huaweicloud_dns_zone":                        ResourceDNSZoneV2(),
			"huaweicloud_dws_cluster":                     ResourceDwsCluster(),
			"huaweicloud_elb_certificate":                 ResourceCertificateV3(),
			"huaweicloud_elb_l7policy":                    ResourceL7PolicyV3(),
			"huaweicloud_elb_l7rule":                      ResourceL7RuleV3(),
			"huaweicloud_elb_listener":                    ResourceListenerV3(),
			"huaweicloud_elb_loadbalancer":                ResourceLoadBalancerV3(),
			"huaweicloud_elb_monitor":                     ResourceMonitorV3(),
			"huaweicloud_elb_ipgroup":                     ResourceIpGroupV3(),
			"huaweicloud_elb_pool":                        ResourcePoolV3(),
			"huaweicloud_elb_member":                      ResourceMemberV3(),
			"huaweicloud_evs_snapshot":                    ResourceEvsSnapshotV2(),
			"huaweicloud_evs_volume":                      ResourceEvsStorageVolumeV3(),
			"huaweicloud_fgs_function":                    ResourceFgsFunctionV2(),
			"huaweicloud_gaussdb_cassandra_instance":      resourceGeminiDBInstanceV3(),
			"huaweicloud_gaussdb_mysql_instance":          resourceGaussDBInstance(),
			"huaweicloud_gaussdb_opengauss_instance":      resourceOpenGaussInstance(),
			"huaweicloud_gaussdb_redis_instance":          resourceGaussRedisInstanceV3(),
			"huaweicloud_ges_graph":                       ResourceGesGraphV1(),
			"huaweicloud_identity_access_key":             resourceIdentityKey(),
			"huaweicloud_identity_acl":                    ResourceIdentityACL(),
			"huaweicloud_identity_agency":                 ResourceIAMAgencyV3(),
			"huaweicloud_identity_group":                  ResourceIdentityGroupV3(),
			"huaweicloud_identity_group_membership":       ResourceIdentityGroupMembershipV3(),
			"huaweicloud_identity_project":                ResourceIdentityProjectV3(),
			"huaweicloud_identity_role":                   ResourceIdentityRole(),
			"huaweicloud_identity_role_assignment":        ResourceIdentityRoleAssignmentV3(),
			"huaweicloud_identity_user":                   ResourceIdentityUserV3(),
			"huaweicloud_iec_eip":                         resourceIecNetworkEip(),
			"huaweicloud_iec_keypair":                     resourceIecKeypair(),
			"huaweicloud_iec_network_acl":                 resourceIecNetworkACL(),
			"huaweicloud_iec_network_acl_rule":            resourceIecNetworkACLRule(),
			"huaweicloud_iec_security_group":              resourceIecSecurityGroup(),
			"huaweicloud_iec_security_group_rule":         resourceIecSecurityGroupRule(),
			"huaweicloud_iec_server":                      resourceIecServer(),
			"huaweicloud_iec_vip":                         resourceIecVipV1(),
			"huaweicloud_iec_vpc":                         ResourceIecVpc(),
			"huaweicloud_iec_vpc_subnet":                  resourceIecSubnet(),
			"huaweicloud_images_image":                    ResourceImsImage(),
			"huaweicloud_kms_key":                         ResourceKmsKeyV1(),
			"huaweicloud_lb_certificate":                  ResourceCertificateV2(),
			"huaweicloud_lb_l7policy":                     ResourceL7PolicyV2(),
			"huaweicloud_lb_l7rule":                       ResourceL7RuleV2(),
			"huaweicloud_lb_listener":                     ResourceListenerV2(),
			"huaweicloud_lb_loadbalancer":                 ResourceLoadBalancerV2(),
			"huaweicloud_lb_member":                       ResourceMemberV2(),
			"huaweicloud_lb_monitor":                      ResourceMonitorV2(),
			"huaweicloud_lb_pool":                         ResourcePoolV2(),
			"huaweicloud_lb_whitelist":                    ResourceWhitelistV2(),
			"huaweicloud_lts_group":                       resourceLTSGroupV2(),
			"huaweicloud_lts_stream":                      resourceLTSStreamV2(),
			"huaweicloud_oms_task":                        resourceMaasTaskV1(),
			"huaweicloud_mls_instance":                    resourceMlsInstance(),
			"huaweicloud_mapreduce_cluster":               mrs.ResourceMRSClusterV2(),
			"huaweicloud_mapreduce_job":                   mrs.ResourceMRSJobV2(),
			"huaweicloud_mrs_cluster":                     ResourceMRSClusterV1(),
			"huaweicloud_mrs_job":                         ResourceMRSJobV1(),
			"huaweicloud_nat_dnat_rule":                   ResourceNatDnatRuleV2(),
			"huaweicloud_nat_gateway":                     ResourceNatGatewayV2(),
			"huaweicloud_nat_snat_rule":                   ResourceNatSnatRuleV2(),
			"huaweicloud_network_acl":                     ResourceNetworkACL(),
			"huaweicloud_network_acl_rule":                ResourceNetworkACLRule(),
			"huaweicloud_networking_eip_associate":        ResourceNetworkingFloatingIPAssociateV2(),
			"huaweicloud_networking_port":                 ResourceNetworkingPortV2(),
			"huaweicloud_networking_secgroup":             ResourceNetworkingSecGroupV2(),
			"huaweicloud_networking_secgroup_rule":        ResourceNetworkingSecGroupRuleV2(),
			"huaweicloud_networking_vip":                  resourceNetworkingVIPV2(),
			"huaweicloud_networking_vip_associate":        resourceNetworkingVIPAssociateV2(),
			"huaweicloud_obs_bucket":                      ResourceObsBucket(),
			"huaweicloud_obs_bucket_object":               ResourceObsBucketObject(),
			"huaweicloud_obs_bucket_policy":               ResourceObsBucketPolicy(),
			"huaweicloud_rds_instance":                    ResourceRdsInstanceV3(),
			"huaweicloud_rds_parametergroup":              ResourceRdsConfigurationV3(),
			"huaweicloud_rds_read_replica_instance":       ResourceRdsReadReplicaInstance(),
			"huaweicloud_sfs_access_rule":                 ResourceSFSAccessRuleV2(),
			"huaweicloud_sfs_file_system":                 ResourceSFSFileSystemV2(),
			"huaweicloud_sfs_turbo":                       ResourceSFSTurbo(),
			"huaweicloud_smn_topic":                       ResourceTopic(),
			"huaweicloud_smn_subscription":                ResourceSubscription(),
			"huaweicloud_swr_organization":                resourceSWROrganization(),
			"huaweicloud_vbs_backup":                      resourceVBSBackupV2(),
			"huaweicloud_vbs_backup_policy":               resourceVBSBackupPolicyV2(),
			"huaweicloud_vpc":                             vpc.ResourceVirtualPrivateCloudV1(),
			"huaweicloud_vpc_bandwidth":                   ResourceVpcBandWidthV2(),
			"huaweicloud_vpc_eip":                         ResourceVpcEIPV1(),
			"huaweicloud_vpc_peering_connection":          ResourceVpcPeeringConnectionV2(),
			"huaweicloud_vpc_peering_connection_accepter": resourceVpcPeeringConnectionAccepterV2(),
			"huaweicloud_vpc_route":                       ResourceVPCRouteV2(),
			"huaweicloud_vpc_subnet":                      vpc.ResourceVpcSubnetV1(),
			"huaweicloud_vpcep_approval":                  ResourceVPCEndpointApproval(),
			"huaweicloud_vpcep_endpoint":                  ResourceVPCEndpoint(),
			"huaweicloud_vpcep_service":                   ResourceVPCEndpointService(),
			"huaweicloud_vpnaas_endpoint_group":           ResourceVpnEndpointGroupV2(),
			"huaweicloud_vpnaas_ike_policy":               ResourceVpnIKEPolicyV2(),
			"huaweicloud_vpnaas_ipsec_policy":             ResourceVpnIPSecPolicyV2(),
			"huaweicloud_vpnaas_service":                  ResourceVpnServiceV2(),
			"huaweicloud_vpnaas_site_connection":          ResourceVpnSiteConnectionV2(),
			"huaweicloud_scm_certificate":                 resourceScmCertificateV3(),
			"huaweicloud_waf_certificate":                 waf.ResourceWafCertificateV1(),
			"huaweicloud_waf_domain":                      waf.ResourceWafDomainV1(),
			"huaweicloud_waf_policy":                      waf.ResourceWafPolicyV1(),
			"huaweicloud_waf_rule_blacklist":              waf.ResourceWafRuleBlackListV1(),
			"huaweicloud_waf_rule_data_masking":           waf.ResourceWafRuleDataMaskingV1(),
			"huaweicloud_waf_rule_web_tamper_protection":  waf.ResourceWafRuleWebTamperProtectionV1(),
			"huaweicloud_waf_dedicated_instance":          waf.ResourceWafDedicatedInstanceV1(),
			"huaweicloud_waf_dedicated_domain":            waf.ResourceWafDedicatedDomainV1(),
			"huaweicloud_waf_reference_table":             waf.ResourceWafReferenceTableV1(),

			// Legacy
			"huaweicloud_compute_instance_v2":                ResourceComputeInstanceV2(),
			"huaweicloud_compute_interface_attach_v2":        ResourceComputeInterfaceAttachV2(),
			"huaweicloud_compute_keypair_v2":                 ResourceComputeKeypairV2(),
			"huaweicloud_compute_servergroup_v2":             ResourceComputeServerGroupV2(),
			"huaweicloud_compute_volume_attach_v2":           ResourceComputeVolumeAttachV2(),
			"huaweicloud_dns_ptrrecord_v2":                   ResourceDNSPtrRecordV2(),
			"huaweicloud_dns_recordset_v2":                   ResourceDNSRecordSetV2(),
			"huaweicloud_dns_zone_v2":                        ResourceDNSZoneV2(),
			"huaweicloud_dcs_instance_v1":                    ResourceDcsInstanceV1(),
			"huaweicloud_dds_instance_v3":                    ResourceDdsInstanceV3(),
			"huaweicloud_fw_firewall_group_v2":               resourceFWFirewallGroupV2(),
			"huaweicloud_fw_policy_v2":                       resourceFWPolicyV2(),
			"huaweicloud_fw_rule_v2":                         resourceFWRuleV2(),
			"huaweicloud_kms_key_v1":                         ResourceKmsKeyV1(),
			"huaweicloud_dms_queue_v1":                       ResourceDmsQueuesV1(),
			"huaweicloud_dms_group_v1":                       ResourceDmsGroupsV1(),
			"huaweicloud_dms_instance_v1":                    ResourceDmsInstancesV1(),
			"huaweicloud_images_image_v2":                    resourceImagesImageV2(),
			"huaweicloud_lb_certificate_v2":                  ResourceCertificateV2(),
			"huaweicloud_lb_loadbalancer_v2":                 ResourceLoadBalancerV2(),
			"huaweicloud_lb_listener_v2":                     ResourceListenerV2(),
			"huaweicloud_lb_pool_v2":                         ResourcePoolV2(),
			"huaweicloud_lb_member_v2":                       ResourceMemberV2(),
			"huaweicloud_lb_monitor_v2":                      ResourceMonitorV2(),
			"huaweicloud_lb_l7policy_v2":                     ResourceL7PolicyV2(),
			"huaweicloud_lb_l7rule_v2":                       ResourceL7RuleV2(),
			"huaweicloud_lb_whitelist_v2":                    ResourceWhitelistV2(),
			"huaweicloud_mrs_cluster_v1":                     ResourceMRSClusterV1(),
			"huaweicloud_mrs_job_v1":                         ResourceMRSJobV1(),
			"huaweicloud_networking_port_v2":                 ResourceNetworkingPortV2(),
			"huaweicloud_networking_secgroup_v2":             ResourceNetworkingSecGroupV2(),
			"huaweicloud_networking_secgroup_rule_v2":        ResourceNetworkingSecGroupRuleV2(),
			"huaweicloud_smn_topic_v2":                       ResourceTopic(),
			"huaweicloud_smn_subscription_v2":                ResourceSubscription(),
			"huaweicloud_rds_instance_v3":                    ResourceRdsInstanceV3(),
			"huaweicloud_rds_parametergroup_v3":              ResourceRdsConfigurationV3(),
			"huaweicloud_nat_gateway_v2":                     ResourceNatGatewayV2(),
			"huaweicloud_nat_snat_rule_v2":                   ResourceNatSnatRuleV2(),
			"huaweicloud_nat_dnat_rule_v2":                   ResourceNatDnatRuleV2(),
			"huaweicloud_sfs_access_rule_v2":                 ResourceSFSAccessRuleV2(),
			"huaweicloud_sfs_file_system_v2":                 ResourceSFSFileSystemV2(),
			"huaweicloud_iam_agency":                         ResourceIAMAgencyV3(),
			"huaweicloud_iam_agency_v3":                      ResourceIAMAgencyV3(),
			"huaweicloud_vpc_v1":                             vpc.ResourceVirtualPrivateCloudV1(),
			"huaweicloud_vpc_bandwidth_v2":                   ResourceVpcBandWidthV2(),
			"huaweicloud_vpc_eip_v1":                         ResourceVpcEIPV1(),
			"huaweicloud_vpc_peering_connection_v2":          ResourceVpcPeeringConnectionV2(),
			"huaweicloud_vpc_peering_connection_accepter_v2": resourceVpcPeeringConnectionAccepterV2(),
			"huaweicloud_vpc_route_v2":                       ResourceVPCRouteV2(),
			"huaweicloud_vpc_subnet_v1":                      vpc.ResourceVpcSubnetV1(),
			"huaweicloud_cce_cluster_v3":                     ResourceCCEClusterV3(),
			"huaweicloud_cce_node_v3":                        ResourceCCENodeV3(),
			"huaweicloud_cci_network_v1":                     resourceCCINetworkV1(),
			"huaweicloud_as_configuration_v1":                ResourceASConfiguration(),
			"huaweicloud_as_group_v1":                        ResourceASGroup(),
			"huaweicloud_as_policy_v1":                       ResourceASPolicy(),
			"huaweicloud_csbs_backup_v1":                     resourceCSBSBackupV1(),
			"huaweicloud_csbs_backup_policy_v1":              resourceCSBSBackupPolicyV1(),
			"huaweicloud_vbs_backup_policy_v2":               resourceVBSBackupPolicyV2(),
			"huaweicloud_vbs_backup_v2":                      resourceVBSBackupV2(),
			"huaweicloud_maas_task":                          resourceMaasTaskV1(),
			"huaweicloud_maas_task_v1":                       resourceMaasTaskV1(),
			"huaweicloud_identity_project_v3":                ResourceIdentityProjectV3(),
			"huaweicloud_identity_role_assignment_v3":        ResourceIdentityRoleAssignmentV3(),
			"huaweicloud_identity_user_v3":                   ResourceIdentityUserV3(),
			"huaweicloud_identity_group_v3":                  ResourceIdentityGroupV3(),
			"huaweicloud_identity_group_membership_v3":       ResourceIdentityGroupMembershipV3(),
			"huaweicloud_cdm_cluster_v1":                     ResourceCdmClusterV1(),
			"huaweicloud_ges_graph_v1":                       ResourceGesGraphV1(),
			"huaweicloud_cloudtable_cluster_v2":              resourceCloudtableClusterV2(),
			"huaweicloud_css_cluster_v1":                     ResourceCssClusterV1(),
			"huaweicloud_dis_stream_v2":                      ResourceDisStreamV2(),
			"huaweicloud_cs_cluster_v1":                      deprecated.ResourceCsClusterV1(),
			"huaweicloud_cs_peering_connect_v1":              deprecated.ResourceCsPeeringConnectV1(),
			"huaweicloud_vpnaas_ipsec_policy_v2":             ResourceVpnIPSecPolicyV2(),
			"huaweicloud_vpnaas_service_v2":                  ResourceVpnServiceV2(),
			"huaweicloud_vpnaas_ike_policy_v2":               ResourceVpnIKEPolicyV2(),
			"huaweicloud_vpnaas_endpoint_group_v2":           ResourceVpnEndpointGroupV2(),
			"huaweicloud_vpnaas_site_connection_v2":          ResourceVpnSiteConnectionV2(),
			"huaweicloud_dli_queue_v1":                       dli.ResourceDliQueue(),
			"huaweicloud_cs_route_v1":                        deprecated.ResourceCsRouteV1(),
			"huaweicloud_networking_vip_v2":                  resourceNetworkingVIPV2(),
			"huaweicloud_networking_vip_associate_v2":        resourceNetworkingVIPAssociateV2(),
			"huaweicloud_fgs_function_v2":                    ResourceFgsFunctionV2(),
			"huaweicloud_cdn_domain_v1":                      resourceCdnDomainV1(),
			// Deprecated
			"huaweicloud_blockstorage_volume_v2":             resourceBlockStorageVolumeV2(),
			"huaweicloud_networking_network_v2":              resourceNetworkingNetworkV2(),
			"huaweicloud_networking_subnet_v2":               resourceNetworkingSubnetV2(),
			"huaweicloud_networking_floatingip_v2":           resourceNetworkingFloatingIPV2(),
			"huaweicloud_networking_router_v2":               resourceNetworkingRouterV2(),
			"huaweicloud_networking_router_interface_v2":     resourceNetworkingRouterInterfaceV2(),
			"huaweicloud_networking_router_route_v2":         resourceNetworkingRouterRouteV2(),
			"huaweicloud_networking_floatingip_associate_v2": ResourceNetworkingFloatingIPAssociateV2(),
			"huaweicloud_ecs_instance_v1":                    resourceEcsInstanceV1(),
			"huaweicloud_compute_secgroup_v2":                ResourceComputeSecGroupV2(),
			"huaweicloud_compute_floatingip_v2":              ResourceComputeFloatingIPV2(),
			"huaweicloud_compute_floatingip_associate_v2":    ResourceComputeFloatingIPAssociateV2(),
			"huaweicloud_rts_stack_v1":                       resourceRTSStackV1(),
			"huaweicloud_rts_software_config_v1":             resourceSoftwareConfigV1(),
			"huaweicloud_cts_tracker":                        deprecated.ResourceCTSTrackerV1(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return configureProvider(d, terraformVersion)
	}

	return provider
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"auth_url": "The Identity authentication URL.",

		"region": "The HuaweiCloud region to connect to.",

		"user_name": "Username to login with.",

		"user_id": "User ID to login with.",

		"project_id": "The ID of the project to login with.",

		"project_name": "The name of the project to login with.",

		"tenant_id": "The ID of the Tenant (Identity v2) to login with.",

		"tenant_name": "The name of the Tenant (Identity v2) to login with.",

		"password": "Password to login with.",

		"token": "Authentication token to use as an alternative to username/password.",

		"domain_id": "The ID of the Domain to scope to.",

		"domain_name": "The name of the Domain to scope to.",

		"access_key":     "The access key of the HuaweiCloud to use.",
		"secret_key":     "The secret key of the HuaweiCloud to use.",
		"security_token": "The security token to authenticate with a temporary security credential.",

		"insecure": "Trust self-signed certificates.",

		"cacert_file": "A Custom CA certificate.",

		"cert": "A client certificate to authenticate with.",

		"key": "A client private key to authenticate with.",

		"agency_name": "The name of agency",

		"agency_domain_name": "The name of domain who created the agency (Identity v3).",

		"delegated_project": "The name of delegated project (Identity v3).",

		"cloud": "The endpoint of cloud provider, defaults to myhuaweicloud.com",

		"endpoints": "The custom endpoints used to override the default endpoint URL.",

		"max_retries": "How many times HTTP connection should be retried until giving up.",

		"enterprise_project_id": "enterprise project id",
	}
}

func configureProvider(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	var tenantName, tenantID, delegated_project, identityEndpoint string
	region := d.Get("region").(string)
	cloud := d.Get("cloud").(string)

	// project_name is prior to tenant_name
	// if neither of them was set, use region as the default project
	if v, ok := d.GetOk("project_name"); ok && v.(string) != "" {
		tenantName = v.(string)
	} else if v, ok := d.GetOk("tenant_name"); ok && v.(string) != "" {
		tenantName = v.(string)
	} else {
		tenantName = region
	}

	// project_id is prior to tenant_id
	if v, ok := d.GetOk("project_id"); ok && v.(string) != "" {
		tenantID = v.(string)
	} else {
		tenantID = d.Get("tenant_id").(string)
	}

	// Use region as delegated_project if it's not set
	if v, ok := d.GetOk("delegated_project"); ok && v.(string) != "" {
		delegated_project = v.(string)
	} else {
		delegated_project = region
	}

	// use auth_url as identityEndpoint if provided
	if v, ok := d.GetOk("auth_url"); ok {
		identityEndpoint = v.(string)
	} else {
		// use cloud as basis for identityEndpoint
		if cloud == defaultCloud {
			identityEndpoint = fmt.Sprintf("https://iam.%s:443/v3", cloud)
		} else {
			identityEndpoint = fmt.Sprintf("https://iam.%s.%s:443/v3", region, cloud)
		}
	}

	config := config.Config{
		AccessKey:           d.Get("access_key").(string),
		SecretKey:           d.Get("secret_key").(string),
		CACertFile:          d.Get("cacert_file").(string),
		ClientCertFile:      d.Get("cert").(string),
		ClientKeyFile:       d.Get("key").(string),
		DomainID:            d.Get("domain_id").(string),
		DomainName:          d.Get("domain_name").(string),
		IdentityEndpoint:    identityEndpoint,
		Insecure:            d.Get("insecure").(bool),
		Password:            d.Get("password").(string),
		Token:               d.Get("token").(string),
		SecurityToken:       d.Get("security_token").(string),
		Region:              region,
		TenantID:            tenantID,
		TenantName:          tenantName,
		Username:            d.Get("user_name").(string),
		UserID:              d.Get("user_id").(string),
		AgencyName:          d.Get("agency_name").(string),
		AgencyDomainName:    d.Get("agency_domain_name").(string),
		DelegatedProject:    delegated_project,
		Cloud:               cloud,
		MaxRetries:          d.Get("max_retries").(int),
		EnterpriseProjectID: d.Get("enterprise_project_id").(string),
		TerraformVersion:    terraformVersion,
		RegionProjectIDMap:  make(map[string]string),
		RPLock:              new(sync.Mutex),
	}

	// get custom endpoints
	endpoints, err := flattenProviderEndpoints(d)
	if err != nil {
		return nil, err
	}
	config.Endpoints = endpoints

	if err := config.LoadAndValidate(); err != nil {
		return nil, err
	}

	return &config, nil
}

func flattenProviderEndpoints(d *schema.ResourceData) (map[string]string, error) {
	endpoints := d.Get("endpoints").(map[string]interface{})
	epMap := make(map[string]string)

	for key, val := range endpoints {
		endpoint := strings.TrimSpace(val.(string))
		// check empty string
		if endpoint == "" {
			return nil, fmt.Errorf("the value of customer endpoint %s must be specified", key)
		}

		// add prefix "https://" and suffix "/"
		if !strings.HasPrefix(endpoint, "http") {
			endpoint = fmt.Sprintf("https://%s", endpoint)
		}
		if !strings.HasSuffix(endpoint, "/") {
			endpoint = fmt.Sprintf("%s/", endpoint)
		}
		epMap[key] = endpoint
	}

	// unify the endpoint which has multi types
	if endpoint, ok := epMap["iam"]; ok {
		epMap["identity"] = endpoint
	}
	if endpoint, ok := epMap["ecs"]; ok {
		epMap["ecsv11"] = endpoint
		epMap["ecsv21"] = endpoint
	}
	if endpoint, ok := epMap["cce"]; ok {
		epMap["cce_addon"] = endpoint
	}
	if endpoint, ok := epMap["evs"]; ok {
		epMap["volumev2"] = endpoint
	}
	if endpoint, ok := epMap["vpc"]; ok {
		epMap["networkv2"] = endpoint
		epMap["security_group"] = endpoint
	}

	log.Printf("[DEBUG] customer endpoints: %+v", epMap)
	return epMap, nil
}
