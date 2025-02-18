package huaweicloud

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/go-homedir"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/aad"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/accessanalyzer"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/antiddos"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/aom"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/apig"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/apigateway"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/apm"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/as"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/asm"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/bcs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/bms"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cae"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cbh"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cbr"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cc"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cce"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cceautopilot"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cci"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/ccm"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cdm"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cdn"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/ces"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cfw"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cloudtable"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cmdb"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cnad"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/coc"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/codearts"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/codeartsdeploy"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/codeartsinspector"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cph"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cpts"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cse"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/css"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/cts"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dataarts"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dbss"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dc"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dcs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/ddm"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dds"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/deprecated"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dew"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dis"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dli"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dms"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dns"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/drs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dsc"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/dws"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/ecs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/eg"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/eip"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/elb"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/eps"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/er"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/evs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/fgs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/ga"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/gaussdb"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/geminidb"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/ges"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/hss"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/iam"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/identitycenter"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/iec"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/ims"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/iotda"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/koogallery"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/lb"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/live"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/lts"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/meeting"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/modelarts"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/mpc"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/mrs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/nat"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/obs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/oms"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/organizations"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/ram"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/rds"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/rfs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/rgc"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/rms"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/sdrs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/secmaster"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/servicestage"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/sfsturbo"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/smn"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/sms"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/swr"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/taurusdb"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/tms"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/ucs"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/vod"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/vpc"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/vpcep"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/vpn"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/waf"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/workspace"
)

const (
	defaultCloud       string = "myhuaweicloud.com"
	defaultEuropeCloud string = "myhuaweicloud.eu"
	prefixEuropeRegion string = "eu-west-1"
)

var Version string

// Provider returns a schema.Provider for HuaweiCloud.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"region": {
				Type:         schema.TypeString,
				Optional:     true,
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

			"assume_role": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agency_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: descriptions["assume_role_agency_name"],
							DefaultFunc: schema.EnvDefaultFunc("HW_ASSUME_ROLE_AGENCY_NAME", nil),
						},
						"domain_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: descriptions["assume_role_domain_name"],
							DefaultFunc: schema.EnvDefaultFunc("HW_ASSUME_ROLE_DOMAIN_NAME", nil),
						},
						"domain_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: descriptions["assume_role_domain_id"],
							DefaultFunc: schema.EnvDefaultFunc("HW_ASSUME_ROLE_DOMAIN_ID", nil),
						},
					},
				},
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
				DefaultFunc: schema.EnvDefaultFunc("HW_CLOUD", ""),
			},

			"endpoints": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: descriptions["endpoints"],
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"regional": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions["regional"],
			},

			"shared_config_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["shared_config_file"],
				DefaultFunc: schema.EnvDefaultFunc("HW_SHARED_CONFIG_FILE", ""),
			},

			"profile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["profile"],
				DefaultFunc: schema.EnvDefaultFunc("HW_PROFILE", ""),
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

			"enable_force_new": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: descriptions["enable_force_new"],
				DefaultFunc: schema.EnvDefaultFunc("HW_ENABLE_FORCE_NEW", false),
			},
			"signing_algorithm": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["signing_algorithm"],
				DefaultFunc: schema.EnvDefaultFunc("HW_SIGNING_ALGORITHM", ""),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"huaweicloud_access_analyzers":              accessanalyzer.DataSourceAccessAnalyzers(),
			"huaweicloud_access_analyzer_archive_rules": accessanalyzer.DataSourceAccessAnalyzerArchiveRules(),

			"huaweicloud_antiddos_config_ranges":                antiddos.DataSourceConfigRanges(),
			"huaweicloud_antiddos_weekly_protection_statistics": antiddos.DataSourceWeeklyProtectionStatistics(),
			"huaweicloud_antiddos_eip_defense_statuses":         antiddos.DataSourceEipDefenseStatuses(),

			"huaweicloud_aom_alarm_action_rules":              aom.DataSourceAomAlarmActionRules(),
			"huaweicloud_aom_alarm_group_rules":               aom.DataSourceAlarmGroupRules(),
			"huaweicloud_aom_prom_instances":                  aom.DataSourceAomPromInstances(),
			"huaweicloud_aom_multi_account_aggregation_rules": aom.DataSourceMultiAccountAggregationRules(),
			"huaweicloud_aom_aggregation_metrics":             aom.DataSourceAggregationMetrics(),
			"huaweicloud_aom_organization_accounts":           aom.DataSourceOrganizationAccounts(),
			"huaweicloud_aom_message_templates":               aom.DataSourceMessageTemplates(),
			"huaweicloud_aom_dashboards_folders":              aom.DataSourceDashboardsFolders(),
			"huaweicloud_aom_alarm_rules":                     aom.DataSourceAlarmRules(),
			"huaweicloud_aom_dashboards":                      aom.DataSourceDashboards(),
			"huaweicloud_aom_alarm_rules_templates":           aom.DataSourceAlarmRulesTemplates(),
			"huaweicloud_aom_alarm_silence_rules":             aom.DataSourceAlarmSilenceRules(),
			"huaweicloud_aom_service_discovery_rules":         aom.DataSourceServiceDiscoveryRules(),

			"huaweicloud_apig_acl_policies":                       apig.DataSourceAclPolicies(),
			"huaweicloud_apig_api_associated_acl_policies":        apig.DataSourceApiAssociatedAclPolicies(),
			"huaweicloud_apig_api_associated_applications":        apig.DataSourceApiAssociatedApplications(),
			"huaweicloud_apig_api_associated_plugins":             apig.DataSourceApiAssociatedPlugins(),
			"huaweicloud_apig_api_associated_signatures":          apig.DataSourceApiAssociatedSignatures(),
			"huaweicloud_apig_api_associated_throttling_policies": apig.DataSourceApiAssociatedThrottlingPolicies(),
			"huaweicloud_apig_api_basic_configurations":           apig.DataSourceApiBasicConfigurations(),
			"huaweicloud_apig_api":                                apig.DataSourceApi(),
			"huaweicloud_apig_appcodes":                           apig.DataSourceAppcodes(),
			"huaweicloud_apig_applications":                       apig.DataSourceApplications(),
			"huaweicloud_apig_application_acl":                    apig.DataSourceApplicationAcl(),
			"huaweicloud_apig_application_quotas":                 apig.DataSourceApigApplicationQuotas(),
			"huaweicloud_apig_channels":                           apig.DataSourceChannels(),
			"huaweicloud_apig_custom_authorizers":                 apig.DataSourceCustomAuthorizers(),
			"huaweicloud_apig_endpoint_connections":               apig.DataSourceApigEndpointConnections(),
			"huaweicloud_apig_environment_variables":              apig.DataSourceApigEnvironmentVariables(),
			"huaweicloud_apig_environments":                       apig.DataSourceEnvironments(),
			"huaweicloud_apig_groups":                             apig.DataSourceGroups(),
			"huaweicloud_apig_instance_features":                  apig.DataSourceInstanceFeatures(),
			"huaweicloud_apig_instance_supported_features":        apig.DataSourceInstanceSupportedFeatures(),
			"huaweicloud_apig_instances":                          apig.DataSourceInstances(),
			"huaweicloud_apig_orchestration_rules":                apig.DataSourceOrchestrationRules(),
			"huaweicloud_apig_signatures":                         apig.DataSourceSignatures(),
			"huaweicloud_apig_throttling_policies":                apig.DataSourceThrottlingPolicies(),

			"huaweicloud_as_activity_logs":       as.DataSourceActivityLogs(),
			"huaweicloud_as_configurations":      as.DataSourceASConfigurations(),
			"huaweicloud_as_groups":              as.DataSourceASGroups(),
			"huaweicloud_as_hook_instances":      as.DataSourceAsHookInstances(),
			"huaweicloud_as_instances":           as.DataSourceASInstances(),
			"huaweicloud_as_lifecycle_hooks":     as.DataSourceLifeCycleHooks(),
			"huaweicloud_as_notifications":       as.DataSourceAsNotifications(),
			"huaweicloud_as_planned_tasks":       as.DataSourceAsPlannedTasks(),
			"huaweicloud_as_policies":            as.DataSourceASPolicies(),
			"huaweicloud_as_policy_execute_logs": as.DataSourcePolicyExecuteLogs(),
			"huaweicloud_as_quotas":              as.DataSourceAsQuotas(),
			"huaweicloud_as_group_quotas":        as.DataSourceAsGroupQuotas(),

			"huaweicloud_asm_meshes": asm.DataSourceAsmMeshes(),

			"huaweicloud_account":            DataSourceAccount(),
			"huaweicloud_availability_zones": DataSourceAvailabilityZones(),

			"huaweicloud_bms_flavors":   bms.DataSourceBmsFlavors(),
			"huaweicloud_bms_instances": bms.DataSourceBmsInstances(),

			"huaweicloud_cae_applications":       cae.DataSourceApplications(),
			"huaweicloud_cae_components":         cae.DataSourceComponents(),
			"huaweicloud_cae_environments":       cae.DataSourceEnvironments(),
			"huaweicloud_cae_notification_rules": cae.DataSourceCaeNotificationRules(),

			"huaweicloud_cbr_backup":   cbr.DataSourceBackup(),
			"huaweicloud_cbr_vaults":   cbr.DataSourceVaults(),
			"huaweicloud_cbr_policies": cbr.DataSourcePolicies(),

			"huaweicloud_cbh_instances":          cbh.DataSourceCbhInstances(),
			"huaweicloud_cbh_flavors":            cbh.DataSourceCbhFlavors(),
			"huaweicloud_cbh_availability_zones": cbh.DataSourceAvailabilityZones(),
			"huaweicloud_cbh_instance_quota":     cbh.DataSourceInstanceQuota(),

			"huaweicloud_cc_authorizations":                               cc.DataSourceCcAuthorizations(),
			"huaweicloud_cc_bandwidth_packages":                           cc.DataSourceCcBandwidthPackages(),
			"huaweicloud_cc_central_networks":                             cc.DataSourceCcCentralNetworks(),
			"huaweicloud_cc_central_network_capabilities":                 cc.DataSourceCcCentralNetworkCapabilities(),
			"huaweicloud_cc_central_network_connections":                  cc.DataSourceCcCentralNetworkConnections(),
			"huaweicloud_cc_central_network_policies":                     cc.DataSourceCcCentralNetworkPolicies(),
			"huaweicloud_cc_central_network_attachments":                  cc.DataSourceCcCentralNetworkAttachments(),
			"huaweicloud_cc_central_network_policies_change_set":          cc.DataSourceCcCentralNetworkPoliciesChangeSet(),
			"huaweicloud_cc_central_network_quotas":                       cc.DataSourceCcCentralNetworkQuotas(),
			"huaweicloud_cc_connections":                                  cc.DataSourceCloudConnections(),
			"huaweicloud_cc_connection_routes":                            cc.DataSourceCcConnectionRoutes(),
			"huaweicloud_cc_connection_tags":                              cc.DataSourceCcConnectionTags(),
			"huaweicloud_cc_inter_region_bandwidths":                      cc.DataSourceCcInterRegionBandwidths(),
			"huaweicloud_cc_global_connection_bandwidths":                 cc.DataSourceCcGlobalConnectionBandwidths(),
			"huaweicloud_cc_global_connection_bandwidth_line_levels":      cc.DataSourceCcGlobalConnectionBandwidthLineLevels(),
			"huaweicloud_cc_global_connection_bandwidth_spec_codes":       cc.DataSourceCcGlobalConnectionBandwidthSpecCodes(),
			"huaweicloud_cc_global_connection_bandwidth_sites":            cc.DataSourceCcGlobalConnectionBandwidthSites(),
			"huaweicloud_cc_support_binding_global_connection_bandwidths": cc.DataSourceCcSupportBindingGlobalConnectionBandwidths(),
			"huaweicloud_cc_network_instances":                            cc.DataSourceCcNetworkInstances(),
			"huaweicloud_cc_permissions":                                  cc.DataSourceCcPermissions(),

			"huaweicloud_cce_addon_template":         cce.DataSourceAddonTemplate(),
			"huaweicloud_cce_cluster":                cce.DataSourceCCEClusterV3(),
			"huaweicloud_cce_clusters":               cce.DataSourceCCEClusters(),
			"huaweicloud_cce_cluster_certificate":    cce.DataSourceCCEClusterCertificate(),
			"huaweicloud_cce_cluster_openid_jwks":    cce.DataSourceCCEClusterOpenIDJWKS(),
			"huaweicloud_cce_node":                   cce.DataSourceNode(),
			"huaweicloud_cce_nodes":                  cce.DataSourceNodes(),
			"huaweicloud_cce_node_pool":              cce.DataSourceCCENodePoolV3(),
			"huaweicloud_cce_charts":                 cce.DataSourceCCECharts(),
			"huaweicloud_cce_cluster_configurations": cce.DataSourceClusterConfigurations(),

			"huaweicloud_cci_namespaces": cci.DataSourceCciNamespaces(),

			"huaweicloud_ccm_certificates":               ccm.DataSourceCertificates(),
			"huaweicloud_ccm_certificate_export":         ccm.DataSourceCertificateExport(),
			"huaweicloud_ccm_private_cas":                ccm.DataSourcePrivateCas(),
			"huaweicloud_ccm_private_ca_export":          ccm.DataSourcePrivateCaExport(),
			"huaweicloud_ccm_private_certificates":       ccm.DataSourcePrivateCertificates(),
			"huaweicloud_ccm_private_certificate_export": ccm.DataSourcePrivateCertificateExport(),

			"huaweicloud_cdn_domain_statistics":   cdn.DataSourceStatistics(),
			"huaweicloud_cdn_domains":             cdn.DataSourceCdnDomains(),
			"huaweicloud_cdn_domain_certificates": cdn.DataSourceDomainCertificates(),
			"huaweicloud_cdn_cache_url_tasks":     cdn.DataSourceCacheUrlTasks(),
			"huaweicloud_cdn_cache_history_tasks": cdn.DataSourceCacheHistoryTasks(),
			"huaweicloud_cdn_billing_option":      cdn.DataSourceBillingOption(),
			"huaweicloud_cdn_logs":                cdn.DataSourceCdnLogs(),
			"huaweicloud_cdn_analytics":           cdn.DataSourceCdnAnalytics(),

			"huaweicloud_ces_agent_dimensions":                 ces.DataSourceCesAgentDimensions(),
			"huaweicloud_ces_alarm_templates":                  ces.DataSourceCesAlarmTemplates(),
			"huaweicloud_ces_alarmrules":                       ces.DataSourceCesAlarmRules(),
			"huaweicloud_ces_dashboards":                       ces.DataSourceCesDashboards(),
			"huaweicloud_ces_dashboard_widgets":                ces.DataSourceCesDashboardWidgets(),
			"huaweicloud_ces_events":                           ces.DataSourceCesEvents(),
			"huaweicloud_ces_event_details":                    ces.DataSourceCesEventDetails(),
			"huaweicloud_ces_metrics":                          ces.DataSourceCesMetrics(),
			"huaweicloud_ces_metric_data":                      ces.DataSourceCesMetricData(),
			"huaweicloud_ces_multiple_metrics_data":            ces.DataSourceMultipleMetricsData(),
			"huaweicloud_ces_one_click_alarms":                 ces.DataSourceCesOneClickAlarms(),
			"huaweicloud_ces_one_click_alarm_rules":            ces.DataSourceCesOneClickAlarmRules(),
			"huaweicloud_ces_resource_groups":                  ces.DataSourceCesGroups(),
			"huaweicloud_ces_resource_group_service_resources": ces.DataSourceCesGroupServiceResources(),

			"huaweicloud_cfw_firewalls":             cfw.DataSourceFirewalls(),
			"huaweicloud_cfw_address_groups":        cfw.DataSourceCfwAddressGroups(),
			"huaweicloud_cfw_address_group_members": cfw.DataSourceCfwAddressGroupMembers(),
			"huaweicloud_cfw_black_white_lists":     cfw.DataSourceCfwBlackWhiteLists(),
			"huaweicloud_cfw_capture_tasks":         cfw.DataSourceCfwCaptureTasks(),
			"huaweicloud_cfw_capture_task_results":  cfw.DataSourceCfwCaptureTaskResults(),
			"huaweicloud_cfw_domain_name_groups":    cfw.DataSourceCfwDomainNameGroups(),
			"huaweicloud_cfw_protection_rules":      cfw.DataSourceCfwProtectionRules(),
			"huaweicloud_cfw_service_groups":        cfw.DataSourceCfwServiceGroups(),
			"huaweicloud_cfw_service_group_members": cfw.DataSourceCfwServiceGroupMembers(),
			"huaweicloud_cfw_access_control_logs":   cfw.DataSourceCfwAccessControlLogs(),
			"huaweicloud_cfw_attack_logs":           cfw.DataSourceCfwAttackLogs(),
			"huaweicloud_cfw_flow_logs":             cfw.DataSourceCfwFlowLogs(),
			"huaweicloud_cfw_regions":               cfw.DataSourceCfwRegions(),
			"huaweicloud_cfw_ips_rules":             cfw.DataSourceCfwIpsRules(),
			"huaweicloud_cfw_ips_custom_rules":      cfw.DataSourceCfwIpsCustomRules(),
			"huaweicloud_cfw_ips_rule_details":      cfw.DataSourceCfwIpsRuleDetails(),
			"huaweicloud_cfw_resource_tags":         cfw.DataSourceCfwResourceTags(),
			"huaweicloud_cfw_tags":                  cfw.DataSourceCfwTags(),

			"huaweicloud_cnad_advanced_instances":         cnad.DataSourceInstances(),
			"huaweicloud_cnad_advanced_available_objects": cnad.DataSourceAvailableProtectedObjects(),
			"huaweicloud_cnad_advanced_protected_objects": cnad.DataSourceProtectedObjects(),

			"huaweicloud_compute_flavors":                 ecs.DataSourceEcsFlavors(),
			"huaweicloud_compute_instance":                ecs.DataSourceComputeInstance(),
			"huaweicloud_compute_instances":               ecs.DataSourceComputeInstances(),
			"huaweicloud_compute_servergroups":            ecs.DataSourceComputeServerGroups(),
			"huaweicloud_compute_instance_remote_console": ecs.DataSourceComputeInstanceRemoteConsole(),

			// CodeArts
			"huaweicloud_codearts_deploy_groups":                         codeartsdeploy.DataSourceCodeartsDeployGroups(),
			"huaweicloud_codearts_deploy_hosts":                          codeartsdeploy.DataSourceCodeartsDeployHosts(),
			"huaweicloud_codearts_deploy_application_groups":             codeartsdeploy.DataSourceCodeartsDeployApplicationGroups(),
			"huaweicloud_codearts_deploy_applications":                   codeartsdeploy.DataSourceCodeartsDeployApplications(),
			"huaweicloud_codearts_deploy_application_deployment_records": codeartsdeploy.DataSourceCodeartsDeployApplicationDeploymentRecords(),
			"huaweicloud_codearts_deploy_environments":                   codeartsdeploy.DataSourceCodeartsDeployEnvironments(),

			"huaweicloud_codearts_inspector_host_groups": codeartsinspector.DataSourceCodeartsInspectorHostGroups(),

			"huaweicloud_cts_notifications": cts.DataSourceNotifications(),
			"huaweicloud_cts_traces":        cts.DataSourceCtsTraces(),
			"huaweicloud_cts_trackers":      cts.DataSourceCtsTrackers(),
			"huaweicloud_cts_operations":    cts.DataSourceCtsOperations(),
			"huaweicloud_cts_quotas":        cts.DataSourceCtsQuotas(),
			"huaweicloud_cts_resources":     cts.DataSourceCtsResources(),
			"huaweicloud_cts_users":         cts.DataSourceCtsUsers(),

			"huaweicloud_cdm_clusters":              cdm.DataSourceCdmClusters(),
			"huaweicloud_cdm_flavors":               cdm.DataSourceCdmFlavors(),
			"huaweicloud_cdm_job_execution_records": cdm.DataSourceCdmJobExecutionRecords(),

			"huaweicloud_cph_server_flavors":      cph.DataSourceServerFlavors(),
			"huaweicloud_cph_phone_flavors":       cph.DataSourcePhoneFlavors(),
			"huaweicloud_cph_phone_images":        cph.DataSourcePhoneImages(),
			"huaweicloud_cph_servers":             cph.DataSourceCphServers(),
			"huaweicloud_cph_phones":              cph.DataSourceCphPhones(),
			"huaweicloud_cph_phone_custom_images": cph.DataSourceCphPhoneCustomImages(),
			"huaweicloud_cph_server_bandwidths":   cph.DataSourceCphServerBandwidths(),
			"huaweicloud_cph_encode_servers":      cph.DataSourceCphEncodeServers(),
			"huaweicloud_cph_phone_connections":   cph.DataSourceCphPhoneConnections(),

			"huaweicloud_cse_microservice_engine_configurations": cse.DataSourceMicroserviceEngineConfigurations(),
			"huaweicloud_cse_microservice_engine_flavors":        cse.DataSourceMicroserviceEngineFlavors(),
			"huaweicloud_cse_microservice_engines":               cse.DataSourceMicroserviceEngines(),
			"huaweicloud_cse_microservice_instances":             cse.DataSourceMicroserviceInstances(),
			"huaweicloud_cse_nacos_namespaces":                   cse.DataSourceNacosNamespaces(),

			"huaweicloud_csms_events":                   dew.DataSourceDewCsmsEvents(),
			"huaweicloud_csms_secrets":                  dew.DataSourceDewCsmsSecrets(),
			"huaweicloud_csms_secret_version":           dew.DataSourceDewCsmsSecret(),
			"huaweicloud_csms_secret_versions":          dew.DataSourceDewCsmsSecretVersions(),
			"huaweicloud_css_flavors":                   css.DataSourceCssFlavors(),
			"huaweicloud_css_clusters":                  css.DataSourceCssClusters(),
			"huaweicloud_css_logstash_pipelines":        css.DataSourceCssLogstashPipelines(),
			"huaweicloud_css_logstash_configurations":   css.DataSourceCssLogstashConfigurations(),
			"huaweicloud_css_elb_loadbalancers":         css.DataSourceCssElbLoadbalancers(),
			"huaweicloud_css_logstash_certificates":     css.DataSourceCssLogstashCertificates(),
			"huaweicloud_css_logstash_pipeline_actions": css.DataSourceCssLogstashPipelineActions(),
			"huaweicloud_css_upgrade_target_images":     css.DataSourceCssUpgradeTargetImages(),
			"huaweicloud_css_logstash_templates":        css.DataSourceCssLogstashTemplates(),
			"huaweicloud_css_cluster_tags":              css.DataSourceCssClusterTags(),
			"huaweicloud_css_log_backup_records":        css.DataSourceCssLogBackupRecords(),
			"huaweicloud_css_scan_tasks":                css.DataSourceCssScanTasks(),
			"huaweicloud_css_snapshots":                 css.DataSourceCssSnapshots(),

			"huaweicloud_dataarts_studio_data_connections": dataarts.DataSourceDataConnections(),
			"huaweicloud_dataarts_studio_workspaces":       dataarts.DataSourceDataArtsStudioWorkspaces(),
			// DataArts Architecture
			"huaweicloud_dataarts_architecture_ds_template_optionals": dataarts.DataSourceTemplateOptionalFields(),
			"huaweicloud_dataarts_architecture_model_statistic":       dataarts.DataSourceArchitectureModelStatistic(),
			"huaweicloud_dataarts_architecture_table_models":          dataarts.DataSourceArchitectureTableModels(),
			// DataArts DataService
			"huaweicloud_dataarts_dataservice_apis":            dataarts.DataSourceDataServiceApis(),
			"huaweicloud_dataarts_dataservice_apps":            dataarts.DataSourceDataServiceApps(),
			"huaweicloud_dataarts_dataservice_authorized_apps": dataarts.DataSourceDataServiceAuthorizedApps(),
			"huaweicloud_dataarts_dataservice_instances":       dataarts.DataSourceDataServiceInstances(),
			"huaweicloud_dataarts_dataservice_messages":        dataarts.DataSourceDataServiceMessages(),
			// DataArts Quality
			"huaweicloud_dataarts_quality_tasks": dataarts.DataSourceQualityTasks(),
			// DataArts Factory
			"huaweicloud_dataarts_factory_jobs": dataarts.DataSourceFactoryJobs(),

			"huaweicloud_dbss_audit_data_masking_rules":  dbss.DataSourceDbssAuditDataMaskingRules(),
			"huaweicloud_dbss_audit_risk_rules":          dbss.DataSourceDbssAuditRiskRules(),
			"huaweicloud_dbss_audit_rule_scopes":         dbss.DataSourceDbssAuditRuleScopes(),
			"huaweicloud_dbss_audit_sql_injection_rules": dbss.DataSourceDbssAuditSqlInjectionRules(),
			"huaweicloud_dbss_availability_zones":        dbss.DataSourceDbssAvailabilityZones(),
			"huaweicloud_dbss_databases":                 dbss.DataSourceDbssDatabases(),
			"huaweicloud_dbss_flavors":                   dbss.DataSourceDbssFlavors(),
			"huaweicloud_dbss_instances":                 dbss.DataSourceDbssInstances(),
			"huaweicloud_dbss_operation_logs":            dbss.DataSourceOperationLogs(),
			"huaweicloud_dbss_rds_databases":             dbss.DataSourceDbssRdsDatabases(),

			"huaweicloud_dc_connections":        dc.DataSourceDcConnections(),
			"huaweicloud_dc_quotas":             dc.DataSourceDcQuotas(),
			"huaweicloud_dc_virtual_gateways":   dc.DataSourceDCVirtualGateways(),
			"huaweicloud_dc_virtual_interfaces": dc.DataSourceDCVirtualInterfaces(),

			"huaweicloud_dcs_flavors":         dcs.DataSourceDcsFlavorsV2(),
			"huaweicloud_dcs_maintainwindow":  dcs.DataSourceDcsMaintainWindow(),
			"huaweicloud_dcs_instances":       dcs.DataSourceDcsInstance(),
			"huaweicloud_dcs_templates":       dcs.DataSourceTemplates(),
			"huaweicloud_dcs_template_detail": dcs.DataSourceTemplateDetail(),
			"huaweicloud_dcs_backups":         dcs.DataSourceBackups(),
			"huaweicloud_dcs_hotkey_analyses": dcs.DataSourceDcsHotkeyAnalyses(),
			"huaweicloud_dcs_bigkey_analyses": dcs.DataSourceDcsBigkeyAnalyses(),
			"huaweicloud_dcs_accounts":        dcs.DataSourceDcsAccounts(),
			"huaweicloud_dcs_diagnosis_tasks": dcs.DataSourceDcsDiagnosisTasks(),

			"huaweicloud_dds_quotas":                                  dds.DataSourceDdsQuotas(),
			"huaweicloud_dds_audit_logs":                              dds.DataSourceDdsAuditLogs(),
			"huaweicloud_dds_audit_log_links":                         dds.DataSourceDdsAuditLogLinks(),
			"huaweicloud_dds_database_versions":                       dds.DataSourceDdsDatabaseVersions(),
			"huaweicloud_dds_flavors":                                 dds.DataSourceDDSFlavorV3(),
			"huaweicloud_dds_migrate_availability_zones":              dds.DataSourceDDSMigrateAvailabilityZones(),
			"huaweicloud_dds_instances":                               dds.DataSourceDdsInstance(),
			"huaweicloud_dds_parameter_templates":                     dds.DataSourceDdsParameterTemplates(),
			"huaweicloud_dds_pt_applicable_instances":                 dds.DataSourceDdsPtApplicableInstances(),
			"huaweicloud_dds_pt_application_records":                  dds.DataSourceDdsPtApplicationRecords(),
			"huaweicloud_dds_pt_modification_records":                 dds.DataSourceDdsPtModificationRecords(),
			"huaweicloud_dds_instance_parameter_modification_records": dds.DataSourceDdsInstanceParameterModificationRecords(),
			"huaweicloud_dds_databases":                               dds.DataSourceDdsDatabases(),
			"huaweicloud_dds_database_users":                          dds.DateSourceDDSDatabaseUser(),
			"huaweicloud_dds_storage_types":                           dds.DataSourceDdsStorageTypes(),
			"huaweicloud_dds_restore_databases":                       dds.DataSourceDdsRestoreDatabases(),
			"huaweicloud_dds_restore_collections":                     dds.DataSourceDdsRestoreCollections(),
			"huaweicloud_dds_restore_time_ranges":                     dds.DataSourceDdsRestoreTimeRanges(),
			"huaweicloud_dds_recycle_instances":                       dds.DataSourceDdsRecycleInstances(),
			"huaweicloud_dds_backups":                                 dds.DataSourceDDSBackups(),
			"huaweicloud_dds_database_roles":                          dds.DateSourceDDSDatabaseRoles(),
			"huaweicloud_dds_error_logs":                              dds.DataSourceDDSErrorLogs(),
			"huaweicloud_dds_error_log_links":                         dds.DataSourceDDSErrorLogLinks(),
			"huaweicloud_dds_slow_log_links":                          dds.DataSourceDDSSlowLogLinks(),
			"huaweicloud_dds_slow_logs":                               dds.DataSourceDDSSlowLogs(),
			"huaweicloud_dds_backup_download_links":                   dds.DataSourceDdsBackupDownloadLinks(),
			"huaweicloud_dds_ssl_cert_download_links":                 dds.DataSourceDdsSslCertDownloadLinks(),
			"huaweicloud_dds_instant_tasks":                           dds.DataSourceDdsInstantTasks(),
			"huaweicloud_dds_scheduled_tasks":                         dds.DataSourceDdsScheduledTasks(),

			"huaweicloud_dli_datasource_auths":       dli.DataSourceAuths(),
			"huaweicloud_dli_datasource_connections": dli.DataSourceConnections(),
			"huaweicloud_dli_elastic_resource_pools": dli.DataSourceDliElasticPools(),
			"huaweicloud_dli_flink_templates":        dli.DataSourceDliFlinkTemplates(),
			"huaweicloud_dli_flinkjar_jobs":          dli.DataSourceDliFlinkjarJobs(),
			"huaweicloud_dli_flinksql_jobs":          dli.DataSourceDliFlinkSQLJobs(),
			"huaweicloud_dli_quotas":                 dli.DataSourceDliQuotas(),
			"huaweicloud_dli_spark_templates":        dli.DataSourceDliSparkTemplates(),
			"huaweicloud_dli_sql_jobs":               dli.DataSourceDliSqlJobs(),
			"huaweicloud_dli_sql_templates":          dli.DataSourceDliSqlTemplates(),

			"huaweicloud_dms_kafka_background_tasks":        dms.DataSourceDmsKafkaBackgroundTasks(),
			"huaweicloud_dms_kafka_flavors":                 dms.DataSourceKafkaFlavors(),
			"huaweicloud_dms_kafka_extend_flavors":          dms.DataSourceDmsKafkaExtendFlavors(),
			"huaweicloud_dms_kafka_instances":               dms.DataSourceDmsKafkaInstances(),
			"huaweicloud_dms_kafka_consumer_groups":         dms.DataSourceDmsKafkaConsumerGroups(),
			"huaweicloud_dms_product":                       dms.DataSourceDmsProduct(),
			"huaweicloud_dms_maintainwindow":                dms.DataSourceDmsMaintainWindow(),
			"huaweicloud_dms_kafka_smart_connect_tasks":     dms.DataSourceDmsKafkaSmartConnectTasks(),
			"huaweicloud_dms_kafkav2_smart_connect_tasks":   dms.DataSourceDmsKafkav2SmartConnectTasks(),
			"huaweicloud_dms_kafka_user_client_quotas":      dms.DataSourceDmsKafkaUserClientQuotas(),
			"huaweicloud_dms_kafka_topic_producers":         dms.DataSourceDmsKafkaTopicProducers(),
			"huaweicloud_dms_kafka_topics":                  dms.DataSourceDmsKafkaTopics(),
			"huaweicloud_dms_kafka_topic_partitions":        dms.DataSourceDmsKafkaTopicPartitions(),
			"huaweicloud_dms_kafka_users":                   dms.DataSourceDmsKafkaUsers(),
			"huaweicloud_dms_kafka_message_diagnosis_tasks": dms.DataSourceDmsKafkaMessageDiagnosisTasks(),
			"huaweicloud_dms_kafka_messages":                dms.DataSourceDmsKafkaMessages(),

			"huaweicloud_dms_rabbitmq_background_tasks": dms.DataSourceDmsRabbitMQBackgroundTasks(),
			"huaweicloud_dms_rabbitmq_flavors":          dms.DataSourceRabbitMQFlavors(),
			"huaweicloud_dms_rabbitmq_plugins":          dms.DataSourceDmsRabbitmqPlugins(),
			"huaweicloud_dms_rabbitmq_instances":        dms.DataSourceDmsRabbitMQInstances(),
			"huaweicloud_dms_rabbitmq_extend_flavors":   dms.DataSourceDmsRabbitmqExtendFlavors(),
			"huaweicloud_dms_rabbitmq_vhosts":           dms.DataSourceDmsRabbitmqVhosts(),
			"huaweicloud_dms_rabbitmq_exchanges":        dms.DataSourceDmsRabbitmqExchanges(),
			"huaweicloud_dms_rabbitmq_queues":           dms.DataSourceDmsRabbitmqQueues(),
			"huaweicloud_dms_rabbitmq_users":            dms.DataSourceDmsRabbitmqUsers(),

			"huaweicloud_dms_rocketmq_broker":                      dms.DataSourceDmsRocketMQBroker(),
			"huaweicloud_dms_rocketmq_instances":                   dms.DataSourceDmsRocketMQInstances(),
			"huaweicloud_dms_rocketmq_topics":                      dms.DataSourceDmsRocketMQTopics(),
			"huaweicloud_dms_rocketmq_topic_access_users":          dms.DataSourceDmsRocketmqTopicAccessUsers(),
			"huaweicloud_dms_rocketmq_users":                       dms.DataSourceDmsRocketMQUsers(),
			"huaweicloud_dms_rocketmq_consumer_groups":             dms.DataSourceDmsRocketMQConsumerGroups(),
			"huaweicloud_dms_rocketmq_consumers":                   dms.DataSourceDmsRocketmqConsumers(),
			"huaweicloud_dms_rocketmq_consumer_group_access_users": dms.DataSourceDmsRocketmqConsumerGroupAccessUsers(),
			"huaweicloud_dms_rocketmq_flavors":                     dms.DataSourceRocketMQFlavors(),
			"huaweicloud_dms_rocketmq_migration_tasks":             dms.DataSourceDmsRocketmqMigrationTasks(),
			"huaweicloud_dms_rocketmq_topic_consumer_groups":       dms.DataSourceDmsRocketmqTopicConsumerGroups(),
			"huaweicloud_dms_rocketmq_message_traces":              dms.DataSourceDmsRocketmqMessageTraces(),
			"huaweicloud_dms_rocketmq_extend_flavors":              dms.DataSourceDmsRocketmqExtendFlavors(),
			"huaweicloud_dms_rocketmq_messages":                    dms.DataSourceDmsRocketMQMessages(),

			"huaweicloud_dns_custom_lines":        dns.DataSourceDNSCustomLines(),
			"huaweicloud_dns_floating_ptrrecords": dns.DataSourceFloatingPtrrecords(),
			"huaweicloud_dns_line_groups":         dns.DataSourceLineGroups(),
			"huaweicloud_dns_nameservers":         dns.DataSourceNameservers(),
			"huaweicloud_dns_quotas":              dns.DataSourceDNSQuotas(),
			"huaweicloud_dns_recordsets":          dns.DataSourceRecordsets(),
			"huaweicloud_dns_zones":               dns.DataSourceZones(),

			"huaweicloud_drs_availability_zones": drs.DataSourceAvailabilityZones(),
			"huaweicloud_drs_node_types":         drs.DataSourceNodeTypes(),

			"huaweicloud_eg_custom_event_channels": eg.DataSourceCustomEventChannels(),
			"huaweicloud_eg_custom_event_sources":  eg.DataSourceCustomEventSources(),
			"huaweicloud_eg_event_channels":        eg.DataSourceEventChannels(),
			"huaweicloud_eg_event_sources":         eg.DataSourceEventSources(),

			"huaweicloud_enterprise_project":  eps.DataSourceEnterpriseProject(),
			"huaweicloud_enterprise_projects": eps.DataSourceEnterpriseProjects(),

			"huaweicloud_er_associations":       er.DataSourceAssociations(),
			"huaweicloud_er_attachments":        er.DataSourceAttachments(),
			"huaweicloud_er_available_routes":   er.DataSourceErAvailableRoutes(),
			"huaweicloud_er_flow_logs":          er.DataSourceFlowLogs(),
			"huaweicloud_er_instances":          er.DataSourceInstances(),
			"huaweicloud_er_propagations":       er.DataSourcePropagations(),
			"huaweicloud_er_quotas":             er.DataSourceErQuotas(),
			"huaweicloud_er_route_tables":       er.DataSourceRouteTables(),
			"huaweicloud_er_availability_zones": er.DataSourceAvailabilityZones(),

			"huaweicloud_evs_volumes":            evs.DataSourceEvsVolumesV2(),
			"huaweicloud_evs_snapshots":          evs.DataSourceEvsSnapshots(),
			"huaweicloud_evs_availability_zones": evs.DataSourceEvsAvailabilityZones(),
			"huaweicloud_evs_volume_types":       evs.DataSourceEvsVolumeTypes(),
			"huaweicloud_evs_volume_transfers":   evs.DataSourceEvsVolumeTransfers(),

			"huaweicloud_fgs_applications":          fgs.DataSourceApplications(),
			"huaweicloud_fgs_application_templates": fgs.DataSourceApplicationTemplates(),
			"huaweicloud_fgs_dependencies":          fgs.DataSourceFunctionGraphDependencies(),
			"huaweicloud_fgs_dependency_versions":   fgs.DataSourceDependencieVersions(),
			"huaweicloud_fgs_function_events":       fgs.DataSourceFunctionEvents(),
			"huaweicloud_fgs_function_triggers":     fgs.DataSourceFunctionTriggers(),
			"huaweicloud_fgs_functions":             fgs.DataSourceFunctions(),
			"huaweicloud_fgs_quotas":                fgs.DataSourceQuotas(),

			"huaweicloud_ga_accelerators":       ga.DataSourceAccelerators(),
			"huaweicloud_ga_access_logs":        ga.DataSourceGaAccessLogs(),
			"huaweicloud_ga_address_groups":     ga.DataSourceAddressGroups(),
			"huaweicloud_ga_availability_zones": ga.DataSourceAvailabilityZones(),
			"huaweicloud_ga_endpoint_groups":    ga.DataSourceEndpointGroups(),
			"huaweicloud_ga_endpoints":          ga.DataSourceEndpoints(),
			"huaweicloud_ga_health_checks":      ga.DataSourceHealthChecks(),
			"huaweicloud_ga_listeners":          ga.DataSourceListeners(),

			"huaweicloud_gaussdb_nosql_flavors":                geminidb.DataSourceGaussDBNoSQLFlavors(),
			"huaweicloud_gaussdb_cassandra_dedicated_resource": geminidb.DataSourceGeminiDBDehResource(),
			"huaweicloud_gaussdb_cassandra_flavors":            geminidb.DataSourceCassandraFlavors(),
			"huaweicloud_gaussdb_cassandra_instance":           geminidb.DataSourceGeminiDBInstance(),
			"huaweicloud_gaussdb_cassandra_instances":          geminidb.DataSourceGeminiDBInstances(),
			"huaweicloud_gaussdb_redis_instance":               geminidb.DataSourceGaussRedisInstance(),
			"huaweicloud_gaussdb_redis_flavors":                geminidb.DataSourceGaussDBRedisFlavors(),
			"huaweicloud_gaussdb_influx_instances":             geminidb.DataSourceGaussDBInfluxInstances(),

			"huaweicloud_gaussdb_opengauss_storage_types":             gaussdb.DataSourceGaussdbOpengaussStorageTypes(),
			"huaweicloud_gaussdb_opengauss_datastores":                gaussdb.DataSourceGaussdbOpengaussDatastores(),
			"huaweicloud_gaussdb_opengauss_flavors":                   gaussdb.DataSourceGaussdbOpengaussFlavors(),
			"huaweicloud_gaussdb_opengauss_available_flavors":         gaussdb.DataSourceGaussdbOpengaussAvailableFlavors(),
			"huaweicloud_gaussdb_opengauss_instance":                  gaussdb.DataSourceOpenGaussInstance(),
			"huaweicloud_gaussdb_opengauss_instances":                 gaussdb.DataSourceOpenGaussInstances(),
			"huaweicloud_gaussdb_opengauss_instance_nodes":            gaussdb.DataSourceGaussdbOpengaussInstanceNodes(),
			"huaweicloud_gaussdb_opengauss_instance_coordinators":     gaussdb.DataSourceGaussdbOpengaussInstanceCoordinators(),
			"huaweicloud_gaussdb_opengauss_instance_features":         gaussdb.DataSourceGaussdbOpengaussInstanceFeatures(),
			"huaweicloud_gaussdb_opengauss_instance_snapshot":         gaussdb.DataSourceGaussdbOpengaussInstanceSnapshot(),
			"huaweicloud_gaussdb_opengauss_databases":                 gaussdb.DataSourceOpenGaussDatabases(),
			"huaweicloud_gaussdb_opengauss_schemas":                   gaussdb.DataSourceOpenGaussSchemas(),
			"huaweicloud_gaussdb_opengauss_parameter_templates":       gaussdb.DataSourceGaussdbOpengaussParameterTemplates(),
			"huaweicloud_gaussdb_opengauss_pt_modify_records":         gaussdb.DataSourceGaussdbOpengaussPtModifyRecords(),
			"huaweicloud_gaussdb_opengauss_pt_apply_records":          gaussdb.DataSourceGaussdbOpengaussPtApplyRecords(),
			"huaweicloud_gaussdb_opengauss_pt_applicable_instances":   gaussdb.DataSourceGaussdbOpengaussPtApplicableInstances(),
			"huaweicloud_gaussdb_opengauss_restore_time_ranges":       gaussdb.DataSourceGaussdbOpengaussRestoreTimeRanges(),
			"huaweicloud_gaussdb_opengauss_restorable_instances":      gaussdb.DataSourceGaussdbOpengaussRestorableInstances(),
			"huaweicloud_gaussdb_opengauss_backups":                   gaussdb.DataSourceGaussdbOpengaussBackups(),
			"huaweicloud_gaussdb_opengauss_backup_files":              gaussdb.DataSourceGaussdbOpengaussBackupFiles(),
			"huaweicloud_gaussdb_opengauss_recycling_instances":       gaussdb.DataSourceGaussdbOpengaussRecyclingInstances(),
			"huaweicloud_gaussdb_opengauss_ssl_cert_download_link":    gaussdb.DataSourceGaussdbOpengaussSslCertDownloadLink(),
			"huaweicloud_gaussdb_opengauss_solution_template_setting": gaussdb.DataSourceGaussdbOpengaussSolutionTemplateSetting(),
			"huaweicloud_gaussdb_opengauss_top_io_traffics":           gaussdb.DataSourceGaussdbOpengaussTopIoTraffics(),
			"huaweicloud_gaussdb_opengauss_project_quotas":            gaussdb.DataSourceGaussdbOpengaussProjectQuotas(),
			"huaweicloud_gaussdb_opengauss_tasks":                     gaussdb.DataSourceOpenGaussTasks(),
			"huaweicloud_gaussdb_opengauss_quotas":                    gaussdb.DataSourceGaussdbOpengaussQuotas(),
			"huaweicloud_gaussdb_opengauss_upgrade_versions":          gaussdb.DataSourceGaussdbOpengaussUpgradeVersions(),
			"huaweicloud_gaussdb_opengauss_tags":                      gaussdb.DataSourceOpenGaussTags(),
			"huaweicloud_gaussdb_opengauss_predefined_tags":           gaussdb.DataSourceOpenGaussPredefinedTags(),
			"huaweicloud_gaussdb_opengauss_slow_logs":                 gaussdb.DataSourceOpenGaussSlowLogs(),
			"huaweicloud_gaussdb_opengauss_error_logs":                gaussdb.DataSourceGaussdbOpengaussErrorLogs(),
			"huaweicloud_gaussdb_opengauss_sql_templates":             gaussdb.DataSourceGaussdbOpengaussSqlTemplates(),
			"huaweicloud_gaussdb_opengauss_sql_throttling_tasks":      gaussdb.DataSourceGaussdbOpengaussSqlThrottlingTasks(),

			"huaweicloud_gaussdb_mysql_engine_versions":          taurusdb.DataSourceGaussdbMysqlEngineVersions(),
			"huaweicloud_gaussdb_mysql_configuration":            taurusdb.DataSourceGaussdbMysqlConfiguration(),
			"huaweicloud_gaussdb_mysql_configurations":           taurusdb.DataSourceGaussdbMysqlConfigurations(),
			"huaweicloud_gaussdb_mysql_dedicated_resource":       taurusdb.DataSourceGaussDBMysqlDehResource(),
			"huaweicloud_gaussdb_mysql_flavors":                  taurusdb.DataSourceGaussdbMysqlFlavors(),
			"huaweicloud_gaussdb_mysql_instance":                 taurusdb.DataSourceGaussDBMysqlInstance(),
			"huaweicloud_gaussdb_mysql_instances":                taurusdb.DataSourceGaussDBMysqlInstances(),
			"huaweicloud_gaussdb_mysql_backups":                  taurusdb.DataSourceGaussdbMysqlBackups(),
			"huaweicloud_gaussdb_mysql_restore_time_ranges":      taurusdb.DataSourceGaussdbMysqlRestoreTimeRanges(),
			"huaweicloud_gaussdb_mysql_database_character_set":   taurusdb.DataSourceGaussdbMysqlDatabaseCharacterSet(),
			"huaweicloud_gaussdb_mysql_databases":                taurusdb.DataSourceGaussdbMysqlDatabases(),
			"huaweicloud_gaussdb_mysql_proxy_flavors":            taurusdb.DataSourceGaussdbMysqlProxyFlavors(),
			"huaweicloud_gaussdb_mysql_proxies":                  taurusdb.DataSourceGaussdbMysqlProxies(),
			"huaweicloud_gaussdb_mysql_pt_modify_records":        taurusdb.DataSourceGaussdbMysqlPtModifyRecords(),
			"huaweicloud_gaussdb_mysql_pt_apply_records":         taurusdb.DataSourceGaussdbMysqlPtApplyRecords(),
			"huaweicloud_gaussdb_mysql_pt_applicable_instances":  taurusdb.DataSourceGaussdbMysqlPtApplicableInstances(),
			"huaweicloud_gaussdb_mysql_recycling_instances":      taurusdb.DataSourceGaussdbMysqlRecyclingInstances(),
			"huaweicloud_gaussdb_mysql_auto_scaling_records":     taurusdb.DataSourceGaussdbMysqlAutoScalingRecords(),
			"huaweicloud_gaussdb_mysql_incremental_backups":      taurusdb.DataSourceGaussdbMysqlIncrementalBackups(),
			"huaweicloud_gaussdb_mysql_restored_tables":          taurusdb.DataSourceGaussdbMysqlRestoredTables(),
			"huaweicloud_gaussdb_mysql_project_quotas":           taurusdb.DataSourceGaussdbMysqlProjectQuotas(),
			"huaweicloud_gaussdb_mysql_instant_tasks":            taurusdb.DataSourceGaussDBMysqlInstantTasks(),
			"huaweicloud_gaussdb_mysql_scheduled_tasks":          taurusdb.DataSourceGaussDBMysqlScheduledTasks(),
			"huaweicloud_gaussdb_mysql_slow_logs":                taurusdb.DataSourceGaussDBMysqlSlowLogs(),
			"huaweicloud_gaussdb_mysql_error_logs":               taurusdb.DataSourceGaussDBMysqlErrorLogs(),
			"huaweicloud_gaussdb_mysql_diagnosis_statistics":     taurusdb.DataSourceGaussdbMysqlDiagnosisStatistics(),
			"huaweicloud_gaussdb_mysql_diagnosis_instances":      taurusdb.DataSourceGaussDBMysqlDiagnosisInstances(),
			"huaweicloud_gaussdb_mysql_audit_log_download_links": taurusdb.DataSourceGaussDBMysqlAuditLogDownloadLinks(),

			"huaweicloud_hss_ransomware_protection_policies": hss.DataSourceRansomwareProtectionPolicies(),
			"huaweicloud_hss_host_groups":                    hss.DataSourceHostGroups(),
			"huaweicloud_hss_hosts":                          hss.DataSourceHosts(),
			"huaweicloud_hss_webtamper_hosts":                hss.DataSourceWebTamperHosts(),
			"huaweicloud_hss_quotas":                         hss.DataSourceQuotas(),

			"huaweicloud_identity_permissions": iam.DataSourceIdentityPermissions(),
			"huaweicloud_identity_role":        iam.DataSourceIdentityRole(),
			"huaweicloud_identity_custom_role": iam.DataSourceIdentityCustomRole(),
			"huaweicloud_identity_group":       iam.DataSourceIdentityGroup(),
			"huaweicloud_identity_projects":    iam.DataSourceIdentityProjects(),
			"huaweicloud_identity_users":       iam.DataSourceIdentityUsers(),
			"huaweicloud_identity_agencies":    iam.DataSourceIdentityAgencies(),
			"huaweicloud_identity_providers":   iam.DataSourceIamIdentityProviders(),

			"huaweicloud_identitycenter_instance":                                identitycenter.DataSourceIdentityCenter(),
			"huaweicloud_identitycenter_groups":                                  identitycenter.DataSourceIdentityCenterGroups(),
			"huaweicloud_identitycenter_users":                                   identitycenter.DataSourceIdentityCenterUsers(),
			"huaweicloud_identitycenter_access_control_attribute_configurations": identitycenter.DataSourceAccessControlAttributeConfigurations(),
			"huaweicloud_identitycenter_permission_sets":                         identitycenter.DataSourceIdentitycenterPermissionSets(),
			"huaweicloud_identitycenter_account_provisioning_permission_sets":    identitycenter.DataSourceAccountProvisioningPermissionSets(),
			"huaweicloud_identitycenter_permission_set_provisioning_accounts":    identitycenter.DataSourcePermissionSetProvisioningAccounts(),
			"huaweicloud_identitycenter_permission_set_provisionings":            identitycenter.DataSourceIdentitycenterPermissionSetProvisionings(),
			"huaweicloud_identitycenter_system_policy_attachments":               identitycenter.DataSourceIdentitycenterSystemPolicyAttachments(),
			"huaweicloud_identitycenter_system_identity_policy_attachments":      identitycenter.DataSourceSystemIdentityPolicyAttachments(),

			"huaweicloud_iec_bandwidths":     iec.DataSourceBandWidths(),
			"huaweicloud_iec_eips":           iec.DataSourceEips(),
			"huaweicloud_iec_flavors":        iec.DataSourceFlavors(),
			"huaweicloud_iec_images":         iec.DataSourceImages(),
			"huaweicloud_iec_keypair":        iec.DataSourceKeypair(),
			"huaweicloud_iec_network_acl":    iec.DataSourceNetworkACL(),
			"huaweicloud_iec_port":           iec.DataSourcePort(),
			"huaweicloud_iec_security_group": iec.DataSourceSecurityGroup(),
			"huaweicloud_iec_server":         iec.DataSourceServer(),
			"huaweicloud_iec_sites":          iec.DataSourceSites(),
			"huaweicloud_iec_vpc":            iec.DataSourceVpc(),
			"huaweicloud_iec_vpc_subnets":    iec.DataSourceVpcSubnets(),

			"huaweicloud_images_image":    ims.DataSourceImagesImageV2(),
			"huaweicloud_images_images":   ims.DataSourceImagesImages(),
			"huaweicloud_ims_os_versions": ims.DataSourceOsVersions(),
			"huaweicloud_ims_quotas":      ims.DataSourceImsQuotas(),

			"huaweicloud_kms_data_key":      dew.DataSourceKmsDataKeyV1(),
			"huaweicloud_kms_grants":        dew.DataSourceKmsGrants(),
			"huaweicloud_kms_key":           dew.DataSourceKmsKey(),
			"huaweicloud_kms_keys":          dew.DataSourceKmsKeys(),
			"huaweicloud_kps_failed_tasks":  dew.DataSourceDewKpsFailedTasks(),
			"huaweicloud_kps_running_tasks": dew.DataSourceDewKpsRunningTasks(),
			"huaweicloud_kps_keypairs":      dew.DataSourceKeypairs(),

			"huaweicloud_iotda_device_messages":            iotda.DataSourceDeviceMessages(),
			"huaweicloud_iotda_device_proxies":             iotda.DataSourceDeviceProxies(),
			"huaweicloud_iotda_device_binding_groups":      iotda.DataSourceDeviceBindingGroups(),
			"huaweicloud_iotda_amqps":                      iotda.DataSourceAMQPQueues(),
			"huaweicloud_iotda_batchtasks":                 iotda.DataSourceBatchTasks(),
			"huaweicloud_iotda_dataforwarding_rules":       iotda.DataSourceDataForwardingRules(),
			"huaweicloud_iotda_data_flow_control_policies": iotda.DataSourceDataFlowControlPolicies(),
			"huaweicloud_iotda_data_backlog_policies":      iotda.DataSourceDataBacklogPolicies(),
			"huaweicloud_iotda_devices":                    iotda.DataSourceDevices(),
			"huaweicloud_iotda_device_certificates":        iotda.DataSourceDeviceCertificates(),
			"huaweicloud_iotda_device_groups":              iotda.DataSourceDeviceGroups(),
			"huaweicloud_iotda_device_linkage_rules":       iotda.DataSourceDeviceLinkageRules(),
			"huaweicloud_iotda_products":                   iotda.DataSourceProducts(),
			"huaweicloud_iotda_spaces":                     iotda.DataSourceSpaces(),
			"huaweicloud_iotda_upgrade_packages":           iotda.DataSourceUpgradePackages(),

			"huaweicloud_koogallery_assets": koogallery.DataSourceKooGalleryAssets(),

			"huaweicloud_lb_listeners":    lb.DataSourceListeners(),
			"huaweicloud_lb_loadbalancer": lb.DataSourceELBV2Loadbalancer(),
			"huaweicloud_lb_certificate":  lb.DataSourceLBCertificateV2(),
			"huaweicloud_lb_pools":        lb.DataSourcePools(),

			"huaweicloud_live_disable_push_streams": live.DataSourceDisablePushStreams(),
			"huaweicloud_live_domains":              live.DataSourceLiveDomains(),
			"huaweicloud_live_recordings":           live.DataSourceLiveRecordings(),
			"huaweicloud_live_transcodings":         live.DataSourceLiveTranscodings(),
			"huaweicloud_live_snapshots":            live.DataSourceLiveSnapshots(),
			"huaweicloud_live_geo_blockings":        live.DataSourceGeoBlockings(),
			"huaweicloud_live_record_callbacks":     live.DataSourceLiveRecordCallbacks(),
			"huaweicloud_live_channels":             live.DataSourceLiveChannels(),
			"huaweicloud_live_cdn_ips":              live.DataSourceLiveCdnIps(),

			"huaweicloud_lts_aom_accesses":                 lts.DataSourceAOMAccesses(),
			"huaweicloud_lts_cce_accesses":                 lts.DataSourceCceAccesses(),
			"huaweicloud_lts_groups":                       lts.DataSourceLtsGroups(),
			"huaweicloud_lts_host_groups":                  lts.DataSourceLtsHostGroups(),
			"huaweicloud_lts_notification_templates":       lts.DataSourceLtsNotificationTemplates(),
			"huaweicloud_lts_search_criteria":              lts.DataSourceLtsSearchCriteria(),
			"huaweicloud_lts_streams":                      lts.DataSourceLtsStreams(),
			"huaweicloud_lts_structuring_custom_templates": lts.DataSourceCustomTemplates(),
			"huaweicloud_lts_transfers":                    lts.DataSourceLtsTransfers(),

			"huaweicloud_elb_certificate":          elb.DataSourceELBCertificateV3(),
			"huaweicloud_elb_flavors":              elb.DataSourceElbFlavorsV3(),
			"huaweicloud_elb_pools":                elb.DataSourcePools(),
			"huaweicloud_elb_active_standby_pools": elb.DataSourceActiveStandbyPools(),
			"huaweicloud_elb_loadbalancers":        elb.DataSourceElbLoadbalances(),
			"huaweicloud_elb_listeners":            elb.DataSourceElbListeners(),
			"huaweicloud_elb_members":              elb.DataSourceElbMembers(),
			"huaweicloud_elb_ipgroups":             elb.DataSourceElbIpGroups(),
			"huaweicloud_elb_logtanks":             elb.DataSourceElbLogtanks(),
			"huaweicloud_elb_l7rules":              elb.DataSourceElbL7rules(),
			"huaweicloud_elb_l7policies":           elb.DataSourceElbL7policies(),
			"huaweicloud_elb_security_policies":    elb.DataSourceElbSecurityPolicies(),
			"huaweicloud_elb_monitors":             elb.DataSourceElbMonitors(),

			"huaweicloud_nat_gateway":             nat.DataSourcePublicGateway(),
			"huaweicloud_nat_gateways":            nat.DataSourcePublicGateways(),
			"huaweicloud_nat_dnat_rules":          nat.DataSourceDnatRules(),
			"huaweicloud_nat_private_dnat_rules":  nat.DataSourcePrivateDnatRules(),
			"huaweicloud_nat_private_gateways":    nat.DataSourcePrivateGateways(),
			"huaweicloud_nat_private_snat_rules":  nat.DataSourcePrivateSnatRules(),
			"huaweicloud_nat_private_transit_ips": nat.DataSourcePrivateTransitIps(),
			"huaweicloud_nat_snat_rules":          nat.DataSourceSnatRules(),

			"huaweicloud_networking_port":           vpc.DataSourceNetworkingPortV2(),
			"huaweicloud_networking_secgroup":       vpc.DataSourceNetworkingSecGroup(),
			"huaweicloud_networking_secgroups":      vpc.DataSourceNetworkingSecGroups(),
			"huaweicloud_networking_secgroup_rules": vpc.DataSourceNetworkingSecGroupRules(),

			"huaweicloud_mapreduce_versions": mrs.DataSourceMrsVersions(),

			"huaweicloud_modelarts_datasets":         modelarts.DataSourceDatasets(),
			"huaweicloud_modelarts_dataset_versions": modelarts.DataSourceDatasetVerions(),
			"huaweicloud_modelarts_notebook_images":  modelarts.DataSourceNotebookImages(),
			"huaweicloud_modelarts_notebook_flavors": modelarts.DataSourceNotebookFlavors(),
			"huaweicloud_modelarts_service_flavors":  modelarts.DataSourceServiceFlavors(),
			"huaweicloud_modelarts_models":           modelarts.DataSourceModels(),
			"huaweicloud_modelarts_model_templates":  modelarts.DataSourceModelTemplates(),
			"huaweicloud_modelarts_workspaces":       modelarts.DataSourceWorkspaces(),
			"huaweicloud_modelarts_services":         modelarts.DataSourceServices(),
			"huaweicloud_modelarts_resource_flavors": modelarts.DataSourceResourceFlavors(),

			"huaweicloud_mapreduce_clusters": mrs.DataSourceMrsClusters(),

			"huaweicloud_obs_buckets":       obs.DataSourceObsBuckets(),
			"huaweicloud_obs_bucket_object": obs.DataSourceObsBucketObject(),

			"huaweicloud_ram_resource_permissions":                  ram.DataSourceRAMPermissions(),
			"huaweicloud_ram_resource_share_invitations":            ram.DataSourceResourceShareInvitations(),
			"huaweicloud_ram_shared_resources":                      ram.DataSourceRAMSharedResources(),
			"huaweicloud_ram_shared_principals":                     ram.DataSourceRAMSharedPrincipals(),
			"huaweicloud_ram_resource_share_associations":           ram.DataSourceShareAssociations(),
			"huaweicloud_ram_resource_share_associated_permissions": ram.DataSourceAssociatedPermissions(),
			"huaweicloud_ram_resource_shares":                       ram.DataSourceRAMShares(),

			"huaweicloud_rds_flavors":                         rds.DataSourceRdsFlavor(),
			"huaweicloud_rds_engine_versions":                 rds.DataSourceRdsEngineVersionsV3(),
			"huaweicloud_rds_instances":                       rds.DataSourceRdsInstances(),
			"huaweicloud_rds_backups":                         rds.DataSourceRdsBackups(),
			"huaweicloud_rds_storage_types":                   rds.DataSourceStoragetype(),
			"huaweicloud_rds_sqlserver_collations":            rds.DataSourceSQLServerCollations(),
			"huaweicloud_rds_sqlserver_databases":             rds.DataSourceSQLServerDatabases(),
			"huaweicloud_rds_sqlserver_accounts":              rds.DataSourceRdsSQLServerAccounts(),
			"huaweicloud_rds_sqlserver_database_privileges":   rds.DataSourceSQLServerDatabasePrivileges(),
			"huaweicloud_rds_pg_plugins":                      rds.DataSourcePgPlugins(),
			"huaweicloud_rds_pg_accounts":                     rds.DataSourcePgAccounts(),
			"huaweicloud_rds_pg_roles":                        rds.DataSourceRdsPgRoles(),
			"huaweicloud_rds_pg_databases":                    rds.DataSourcePgDatabases(),
			"huaweicloud_rds_mysql_databases":                 rds.DataSourceRdsMysqlDatabases(),
			"huaweicloud_rds_mysql_database_privileges":       rds.DataSourceRdsMysqlDatabasePrivileges(),
			"huaweicloud_rds_mysql_accounts":                  rds.DataSourceRdsMysqlAccounts(),
			"huaweicloud_rds_mysql_binlog":                    rds.DataSourceRdsMysqlBinlog(),
			"huaweicloud_rds_mysql_proxy_flavors":             rds.DataSourceRdsMysqlProxyFlavors(),
			"huaweicloud_rds_mysql_proxies":                   rds.DataSourceRdsMysqlProxies(),
			"huaweicloud_rds_parametergroups":                 rds.DataSourceParametergroups(),
			"huaweicloud_rds_sql_audit_operations":            rds.DataSourceRdsSqlAuditTypes(),
			"huaweicloud_rds_cross_region_backups":            rds.DataSourceRdsCrossRegionBackups(),
			"huaweicloud_rds_cross_region_backup_instances":   rds.DataSourceRdsCrossRegionBackupInstances(),
			"huaweicloud_rds_sql_audit_logs":                  rds.DataSourceRdsSqlAuditLogs(),
			"huaweicloud_rds_sql_audit_log_links":             rds.DataSourceRdsSqlAuditLogLinks(),
			"huaweicloud_rds_error_logs":                      rds.DataSourceRdsErrorLogs(),
			"huaweicloud_rds_error_log_link":                  rds.DataSourceRdsErrorLogLink(),
			"huaweicloud_rds_slow_logs":                       rds.DataSourceRdsSlowLogs(),
			"huaweicloud_rds_slow_log_link":                   rds.DataSourceRdsSlowLogLink(),
			"huaweicloud_rds_pg_sql_limits":                   rds.DataSourceRdsPgSqlLimits(),
			"huaweicloud_rds_recycling_instances":             rds.DataSourceRdsRecyclingInstances(),
			"huaweicloud_rds_pg_plugin_parameter_value_range": rds.DataSourceRdsPgPluginParameterValueRange(),
			"huaweicloud_rds_pg_plugin_parameter_values":      rds.DataSourceRdsPgPluginParameterValues(),
			"huaweicloud_rds_restore_time_ranges":             rds.DataSourceRdsRestoreTimeRanges(),
			"huaweicloud_rds_restored_databases":              rds.DataSourceRdsRestoredDatabases(),
			"huaweicloud_rds_restored_tables":                 rds.DataSourceRdsRestoredTables(),
			"huaweicloud_rds_extend_log_files":                rds.DataSourceRdsExtendLogFiles(),
			"huaweicloud_rds_extend_log_links":                rds.DataSourceRdsExtendLogLinks(),
			"huaweicloud_rds_slow_log_files":                  rds.DataSourceRdsSlowLogFiles(),

			"huaweicloud_rms_policy_definitions":                       rms.DataSourcePolicyDefinitions(),
			"huaweicloud_rms_assignment_package_templates":             rms.DataSourceTemplates(),
			"huaweicloud_rms_regions":                                  rms.DataSourceRmsRegions(),
			"huaweicloud_rms_services":                                 rms.DataSourceRmsServices(),
			"huaweicloud_rms_policy_assignments":                       rms.DataSourceRmsPolicyAssignments(),
			"huaweicloud_rms_advanced_query_schemas":                   rms.DataSourceRmsAdvancedQuerySchemas(),
			"huaweicloud_rms_assignment_packages":                      rms.DataSourceRmsAssignmentPackages(),
			"huaweicloud_rms_organizational_policy_assignments":        rms.DataSourceRmsOrganizationalPolicyAssignments(),
			"huaweicloud_rms_organizational_assignment_packages":       rms.DataSourceRmsOrganizationalAssignmentPackages(),
			"huaweicloud_rms_advanced_query":                           rms.DataSourceAdvancedQuery(),
			"huaweicloud_rms_advanced_queries":                         rms.DataSourceRmsAdvancedQueries(),
			"huaweicloud_rms_resource_aggregators":                     rms.DataSourceRmsAggregators(),
			"huaweicloud_rms_resources":                                rms.DataSourceResources(),
			"huaweicloud_rms_resources_summary":                        rms.DataSourceResourcesSummary(),
			"huaweicloud_rms_resource_aggregation_pending_requests":    rms.DataSourceRmsAggregationPendingRequests(),
			"huaweicloud_rms_resource_aggregator_source_statuses":      rms.DataSourceRmsAggregatorSourceStatuses(),
			"huaweicloud_rms_policy_states":                            rms.DataSourcePolicyStates(),
			"huaweicloud_rms_assignment_package_scores":                rms.DataSourceRmsAssignmentPackageScores(),
			"huaweicloud_rms_assignment_package_results":               rms.DataSourceRmsAssignmentPackageResults(),
			"huaweicloud_rms_resource_aggregator_discovered_resources": rms.DataSourceAggregatorDiscoveredResources(),
			"huaweicloud_rms_resource_aggregator_advanced_query":       rms.DataSourceAggregatorAdvancedQuery(),
			"huaweicloud_rms_resource_aggregator_policy_states":        rms.DataSourceAggregatorPolicyStates(),
			"huaweicloud_rms_resource_aggregator_policy_assignments":   rms.DataSourceAggregatorPolicyAssignments(),
			"huaweicloud_rms_resource_histories":                       rms.DataSourceRmsHistories(),
			"huaweicloud_rms_resource_relations_details":               rms.DataSourceRmsRelationsDetails(),
			"huaweicloud_rms_remediation_execution_statuses":           rms.DataSourceRemediationExecutionStatuses(),
			"huaweicloud_rms_resource_tags":                            rms.DataSourceResourceTags(),
			"huaweicloud_rms_resource_instances":                       rms.DataSourceResourceInstances(),

			"huaweicloud_sdrs_domain": sdrs.DataSourceSDRSDomain(),

			"huaweicloud_secmaster_workflows":                 secmaster.DataSourceSecmasterWorkflows(),
			"huaweicloud_secmaster_workspaces":                secmaster.DataSourceSecmasterWorkspaces(),
			"huaweicloud_secmaster_incidents":                 secmaster.DataSourceIncidents(),
			"huaweicloud_secmaster_alerts":                    secmaster.DataSourceAlerts(),
			"huaweicloud_secmaster_indicators":                secmaster.DataSourceIndicators(),
			"huaweicloud_secmaster_metric_results":            secmaster.DataSourceMetricResults(),
			"huaweicloud_secmaster_baseline_check_results":    secmaster.DataSourceSecmasterBaselineCheckResults(),
			"huaweicloud_secmaster_playbooks":                 secmaster.DataSourceSecmasterPlaybooks(),
			"huaweicloud_secmaster_alert_rules":               secmaster.DataSourceSecmasterAlertRules(),
			"huaweicloud_secmaster_alert_rule_templates":      secmaster.DataSourceSecmasterAlertRuleTemplates(),
			"huaweicloud_secmaster_playbook_versions":         secmaster.DataSourceSecmasterPlaybookVersions(),
			"huaweicloud_secmaster_playbook_instances":        secmaster.DataSourceSecmasterPlaybookInstances(),
			"huaweicloud_secmaster_data_classes":              secmaster.DataSourceSecmasterDataClasses(),
			"huaweicloud_secmaster_data_class_fields":         secmaster.DataSourceSecmasterDataClassFields(),
			"huaweicloud_secmaster_playbook_action_instances": secmaster.DataSourceSecmasterPlaybookActionInstances(),
			"huaweicloud_secmaster_playbook_statistics":       secmaster.DataSourceSecmasterPlaybookStatistics(),
			"huaweicloud_secmaster_playbook_audit_logs":       secmaster.DataSourceSecmasterPlaybookAuditLogs(),
			"huaweicloud_secmaster_playbook_monitors":         secmaster.DataSourceSecmasterPlaybookMonitors(),
			"huaweicloud_secmaster_alert_rule_metrics":        secmaster.DataSourceSecmasterAlertRuleMetrics(),

			// Querying by Ver.2 APIs
			"huaweicloud_servicestage_component_runtimes": servicestage.DataSourceComponentRuntimes(),
			// Querying by Ver.3 APIs
			"huaweicloud_servicestagev3_applications":   servicestage.DataSourceV3Applications(),
			"huaweicloud_servicestagev3_components":     servicestage.DataSourceV3Components(),
			"huaweicloud_servicestagev3_environments":   servicestage.DataSourceV3Environments(),
			"huaweicloud_servicestagev3_runtime_stacks": servicestage.DataSourceV3RuntimeStacks(),

			"huaweicloud_smn_topics":              smn.DataSourceTopics(),
			"huaweicloud_smn_message_templates":   smn.DataSourceSmnMessageTemplates(),
			"huaweicloud_smn_subscriptions":       smn.DataSourceSmnSubscriptions(),
			"huaweicloud_smn_logtanks":            smn.DataSourceSmnLogtanks(),
			"huaweicloud_smn_topic_subscriptions": smn.DataSourceSmnTopicSubscriptions(),

			"huaweicloud_sms_source_servers": sms.DataSourceServers(),

			"huaweicloud_sfs_turbos":            sfsturbo.DataSourceTurbos(),
			"huaweicloud_sfs_turbo_data_tasks":  sfsturbo.DataSourceSfsTurboDataTasks(),
			"huaweicloud_sfs_turbo_du_tasks":    sfsturbo.DataSourceSfsTurboDuTasks(),
			"huaweicloud_sfs_turbo_obs_targets": sfsturbo.DataSourceSfsTurboObsTargets(),
			"huaweicloud_sfs_turbo_perm_rules":  sfsturbo.DataSourceSfsTurboPermRules(),

			"huaweicloud_swr_organizations":             swr.DataSourceOrganizations(),
			"huaweicloud_swr_repositories":              swr.DataSourceRepositories(),
			"huaweicloud_swr_shared_repositories":       swr.DataSourceSharedRepositories(),
			"huaweicloud_swr_image_triggers":            swr.DataSourceImageTriggers(),
			"huaweicloud_swr_image_tags":                swr.DataSourceImageTags(),
			"huaweicloud_swr_shared_accounts":           swr.DataSourceSharedAccounts(),
			"huaweicloud_swr_image_retention_policies":  swr.DataSourceImageRetentionPolicies(),
			"huaweicloud_swr_image_retention_histories": swr.DataSourceSwrImageRetentionHistories(),
			"huaweicloud_swr_quotas":                    swr.DataSourceSwrQuotas(),
			"huaweicloud_swr_feature_gates":             swr.DataSourceSwrFeatureGates(),
			"huaweicloud_swr_domain_overviews":          swr.DataSourceSwrDomainOverviews(),
			"huaweicloud_swr_domain_resource_reports":   swr.DataSourceSwrDomainReports(),

			"huaweicloud_tms_resource_types": tms.DataSourceResourceTypes(),

			"huaweicloud_vpc_bandwidth":              eip.DataSourceBandWidth(),
			"huaweicloud_vpc_bandwidths":             eip.DataSourceBandWidths(),
			"huaweicloud_vpc_eip":                    eip.DataSourceVpcEip(),
			"huaweicloud_vpc_eips":                   eip.DataSourceVpcEips(),
			"huaweicloud_vpc_internet_gateways":      eip.DataSourceVPCInternetGateways(),
			"huaweicloud_global_eip_pools":           eip.DataSourceGlobalEIPPools(),
			"huaweicloud_global_eip_access_sites":    eip.DataSourceGlobalEIPAccessSites(),
			"huaweicloud_global_internet_bandwidths": eip.DataSourceGlobalInternetBandwidths(),
			"huaweicloud_global_eips":                eip.DataSourceGlobalEIPs(),

			"huaweicloud_vpc":                             vpc.DataSourceVpcV1(),
			"huaweicloud_vpc_address_groups":              vpc.DataSourceVpcAddressGroups(),
			"huaweicloud_vpc_flow_logs":                   vpc.DataSourceVpcFlowLogs(),
			"huaweicloud_vpc_network_acls":                vpc.DataSourceNetworkAcls(),
			"huaweicloud_vpc_peering_connection":          vpc.DataSourceVpcPeeringConnectionV2(),
			"huaweicloud_vpc_route_table":                 vpc.DataSourceVPCRouteTable(),
			"huaweicloud_vpc_routes":                      vpc.DataSourceVpcRoutes(),
			"huaweicloud_vpc_sub_network_interfaces":      vpc.DataSourceVpcSubNetworkInterfaces(),
			"huaweicloud_vpc_subnet":                      vpc.DataSourceVpcSubnetV1(),
			"huaweicloud_vpc_subnet_ids":                  vpc.DataSourceVpcSubnetIdsV1(),
			"huaweicloud_vpc_subnet_private_ips":          vpc.DataSourceVpcSubnetPrivateIps(),
			"huaweicloud_vpc_subnets":                     vpc.DataSourceVpcSubnets(),
			"huaweicloud_vpc_traffic_mirror_filter_rules": vpc.DataSourceVpcTrafficMirrorFilterRules(),
			"huaweicloud_vpc_traffic_mirror_filters":      vpc.DataSourceVpcTrafficMirrorFilters(),
			"huaweicloud_vpc_traffic_mirror_sessions":     vpc.DataSourceVpcTrafficMirrorSessions(),
			"huaweicloud_vpcs":                            vpc.DataSourceVpcs(),

			"huaweicloud_vpcep_endpoints":           vpcep.DataSourceVPCEPEndpoints(),
			"huaweicloud_vpcep_public_services":     vpcep.DataSourceVPCEPPublicServices(),
			"huaweicloud_vpcep_quotas":              vpcep.DataSourceVpcepQuotas(),
			"huaweicloud_vpcep_services":            vpcep.DataSourceVPCEPServices(),
			"huaweicloud_vpcep_service_connections": vpcep.DataSourceVPCEPServiceConnections(),
			"huaweicloud_vpcep_service_permissions": vpcep.DataSourceVPCEPServicePermissions(),
			"huaweicloud_vpcep_service_summary":     vpcep.DataSourceVpcepServiceSummary(),

			"huaweicloud_vpn_access_policies":                vpn.DataSourceVpnAccessPolicies(),
			"huaweicloud_vpn_gateway_availability_zones":     vpn.DataSourceVpnGatewayAZs(),
			"huaweicloud_vpn_gateways":                       vpn.DataSourceGateways(),
			"huaweicloud_vpn_customer_gateways":              vpn.DataSourceVpnCustomerGateways(),
			"huaweicloud_vpn_connections":                    vpn.DataSourceVpnConnections(),
			"huaweicloud_vpn_connection_health_checks":       vpn.DataSourceVpnConnectionHealthChecks(),
			"huaweicloud_vpn_p2c_gateways":                   vpn.DataSourceVpnP2cGateways(),
			"huaweicloud_vpn_p2c_gateway_availability_zones": vpn.DataSourceVpnP2cGatewayAvailabilityZones(),
			"huaweicloud_vpn_p2c_gateway_connections":        vpn.DataSourceVpnP2cGatewayConnections(),
			"huaweicloud_vpn_servers":                        vpn.DataSourceVpnServers(),
			"huaweicloud_vpn_users":                          vpn.DataSourceVpnUsers(),
			"huaweicloud_vpn_user_groups":                    vpn.DataSourceVpnUserGroups(),
			"huaweicloud_vpn_quotas":                         vpn.DataSourceVpnQuotas(),

			"huaweicloud_waf_address_groups":                       waf.DataSourceWafAddressGroups(),
			"huaweicloud_waf_certificate":                          waf.DataSourceWafCertificate(),
			"huaweicloud_waf_dedicated_domains":                    waf.DataSourceWafDedicatedDomains(),
			"huaweicloud_waf_dedicated_instances":                  waf.DataSourceWafDedicatedInstances(),
			"huaweicloud_waf_domains":                              waf.DataSourceWafDomains(),
			"huaweicloud_waf_instance_groups":                      waf.DataSourceWafInstanceGroups(),
			"huaweicloud_waf_policies":                             waf.DataSourceWafPolicies(),
			"huaweicloud_waf_reference_tables":                     waf.DataSourceWafReferenceTables(),
			"huaweicloud_waf_rules_anti_crawler":                   waf.DataSourceWafRulesAntiCrawler(),
			"huaweicloud_waf_rules_blacklist":                      waf.DataSourceWafRulesBlacklist(),
			"huaweicloud_waf_rules_cc_protection":                  waf.DataSourceWafRulesCcProtection(),
			"huaweicloud_waf_rules_data_masking":                   waf.DataSourceWafRulesDataMasking(),
			"huaweicloud_waf_rules_geolocation_access_control":     waf.DataSourceWafRulesGeolocationAccessControl(),
			"huaweicloud_waf_rules_global_protection_whitelist":    waf.DataSourceWafRulesGlobalProtectionWhitelist(),
			"huaweicloud_waf_rules_information_leakage_prevention": waf.DataSourceWafRulesInformationLeakagePrevention(),
			"huaweicloud_waf_rules_known_attack_source":            waf.DataSourceWafRulesKnownAttackSource(),
			"huaweicloud_waf_rules_precise_protection":             waf.DataSourceWafRulesPreciseProtection(),
			"huaweicloud_waf_rules_web_tamper_protection":          waf.DataSourceWafRulesWebTamperProtection(),
			"huaweicloud_waf_source_ips":                           waf.DataSourceWafSourceIps(),
			"huaweicloud_waf_certificates":                         waf.DataSourceWafCertificates(),

			"huaweicloud_dws_alarm_subscriptions":             dws.DataSourceAlarmSubscriptions(),
			"huaweicloud_dws_availability_zones":              dws.DataSourceDwsAvailabilityZones(),
			"huaweicloud_dws_cluster_cns":                     dws.DataSourceDwsClusterCns(),
			"huaweicloud_dws_cluster_logs":                    dws.DataSourceDwsClusterLogs(),
			"huaweicloud_dws_cluster_nodes":                   dws.DataSourceDwsClusterNodes(),
			"huaweicloud_dws_cluster_parameters":              dws.DataSourceClusterParameters(),
			"huaweicloud_dws_cluster_snapshot_statistics":     dws.DataSourceDwsClusterSnapshotStatistics(),
			"huaweicloud_dws_cluster_topo_rings":              dws.DataSourceDwsClusterTopoRings(),
			"huaweicloud_dws_clusters":                        dws.DataSourceDwsClusters(),
			"huaweicloud_dws_disaster_recovery_tasks":         dws.DataSourceDisasterRecoveryTasks(),
			"huaweicloud_dws_event_subscriptions":             dws.DataSourceEventSubscriptions(),
			"huaweicloud_dws_flavors":                         dws.DataSourceDwsFlavors(),
			"huaweicloud_dws_logical_cluster_rings":           dws.DataSourceLogicalClusterRings(),
			"huaweicloud_dws_logical_cluster_volumes":         dws.DataSourceDwsLogicalClusterVolumes(),
			"huaweicloud_dws_logical_clusters":                dws.DataSourceDwsLogicalClusters(),
			"huaweicloud_dws_om_account_configuration":        dws.DataSourceOmAccountConfiguration(),
			"huaweicloud_dws_quotas":                          dws.DataSourceDwsQuotas(),
			"huaweicloud_dws_snapshot_policies":               dws.DataSourceDwsSnapshotPolicies(),
			"huaweicloud_dws_snapshots":                       dws.DataSourceDwsSnapshots(),
			"huaweicloud_dws_statistics":                      dws.DataSourceDwsStatistics(),
			"huaweicloud_dws_workload_plans":                  dws.DataSourceDwsWorkloadPlans(),
			"huaweicloud_dws_workload_queue_associated_users": dws.DataSourceDwsWorkloadQueueAssociatedUsers(),
			"huaweicloud_dws_workload_queues":                 dws.DataSourceWorkloadQueues(),

			"huaweicloud_workspace_app_group_authorizations": workspace.DataSourceWorkspaceAppGroupAuthorizations(),
			"huaweicloud_workspace_app_groups":               workspace.DataSourceWorkspaceAppGroups(),
			"huaweicloud_workspace_app_image_servers":        workspace.DataSourceWorkspaceAppImageServers(),
			"huaweicloud_workspace_app_nas_storages":         workspace.DataSourceAppNasStorages(),
			"huaweicloud_workspace_app_publishable_apps":     workspace.DataSourceWorkspaceAppPublishableApps(),
			"huaweicloud_workspace_app_storage_policies":     workspace.DataSourceAppStoragePolicies(),
			"huaweicloud_workspace_desktops":                 workspace.DataSourceDesktops(),
			"huaweicloud_workspace_flavors":                  workspace.DataSourceWorkspaceFlavors(),
			"huaweicloud_workspace_service":                  workspace.DataSourceService(),

			"huaweicloud_cpts_projects": cpts.DataSourceCptsProjects(),

			// Legacy
			"huaweicloud_images_image_v2":        ims.DataSourceImagesImageV2(),
			"huaweicloud_networking_port_v2":     vpc.DataSourceNetworkingPortV2(),
			"huaweicloud_networking_secgroup_v2": vpc.DataSourceNetworkingSecGroup(),

			"huaweicloud_kms_key_v1":      dew.DataSourceKmsKey(),
			"huaweicloud_kms_data_key_v1": dew.DataSourceKmsDataKeyV1(),

			"huaweicloud_rds_flavors_v3": rds.DataSourceRdsFlavor(),

			"huaweicloud_vpc_v1":                    vpc.DataSourceVpcV1(),
			"huaweicloud_vpc_ids_v1":                vpc.DataSourceVpcIdsV1(),
			"huaweicloud_vpc_peering_connection_v2": vpc.DataSourceVpcPeeringConnectionV2(),
			"huaweicloud_vpc_subnet_v1":             vpc.DataSourceVpcSubnetV1(),
			"huaweicloud_vpc_subnet_ids_v1":         vpc.DataSourceVpcSubnetIdsV1(),

			"huaweicloud_cce_cluster_v3":                    cce.DataSourceCCEClusterV3(),
			"huaweicloud_cce_node_v3":                       cce.DataSourceNode(),
			"huaweicloud_cce_autopilot_clusters":            cceautopilot.DataSourceCceAutopilotClusters(),
			"huaweicloud_cce_autopilot_addon_templates":     cceautopilot.DataSourceCceAutopilotAddonTemplates(),
			"huaweicloud_cce_autopilot_cluster_certificate": cceautopilot.DataSourceCceAutopilotClusterCertificate(),

			"huaweicloud_dms_product_v1":        dms.DataSourceDmsProduct(),
			"huaweicloud_dms_maintainwindow_v1": dms.DataSourceDmsMaintainWindow(),

			"huaweicloud_dcs_maintainwindow_v1": dcs.DataSourceDcsMaintainWindow(),

			"huaweicloud_dds_flavors_v3":               dds.DataSourceDDSFlavorV3(),
			"huaweicloud_identity_role_v3":             iam.DataSourceIdentityRole(),
			"huaweicloud_identity_virtual_mfa_devices": iam.DataSourceIamIdentityVirtualMfaDevices(),
			"huaweicloud_cdm_flavors_v1":               cdm.DataSourceCdmFlavors(),

			"huaweicloud_ddm_engines":                     ddm.DataSourceDdmEngines(),
			"huaweicloud_ddm_flavors":                     ddm.DataSourceDdmFlavors(),
			"huaweicloud_ddm_instances":                   ddm.DataSourceDdmInstances(),
			"huaweicloud_ddm_instance_nodes":              ddm.DataSourceDdmInstanceNodes(),
			"huaweicloud_ddm_instance_groups":             ddm.DataSourceDdmInstanceGroups(),
			"huaweicloud_ddm_instance_available_versions": ddm.DataSourceDdmInstanceAvailableVersions(),
			"huaweicloud_ddm_schemas":                     ddm.DataSourceDdmSchemas(),
			"huaweicloud_ddm_accounts":                    ddm.DataSourceDdmAccounts(),
			"huaweicloud_ddm_physical_sessions":           ddm.DataSourceDdmPhysicalSessions(),
			"huaweicloud_ddm_logical_sessions":            ddm.DataSourceDdmLogicalSessions(),
			"huaweicloud_ddm_killing_sessions_audit_logs": ddm.DataSourceDdmKillingSessionsAuditLogs(),

			"huaweicloud_organizations_organization":         organizations.DataSourceOrganization(),
			"huaweicloud_organizations_organizational_units": organizations.DataSourceOrganizationalUnits(),
			"huaweicloud_organizations_accounts":             organizations.DataSourceAccounts(),
			"huaweicloud_organizations_policies":             organizations.DataSourcePolicies(),
			"huaweicloud_organizations_sent_invitations":     organizations.DataSourceOrganizationsSentInvitations(),
			"huaweicloud_organizations_received_invitations": organizations.DataSourceOrganizationsReceivedInvitations(),
			"huaweicloud_organizations_services":             organizations.DataSourceOrganizationsServices(),
			"huaweicloud_organizations_trusted_services":     organizations.DataSourceOrganizationsTrustedServices(),
			"huaweicloud_organizations_effective_policies":   organizations.DataSourceOrganizationsEffectivePolicies(),
			"huaweicloud_organizations_tag_policy_services":  organizations.DataSourceOrganizationsTagPolicyServices(),
			"huaweicloud_organizations_resource_tags":        organizations.DataSourceOrganizationsTags(),
			"huaweicloud_organizations_resource_instances":   organizations.DataSourceOrganizationsResourceInstances(),

			// Deprecated ongoing (without DeprecationMessage), used by other providers
			"huaweicloud_vpc_route":        vpc.DataSourceVpcRouteV2(),
			"huaweicloud_vpc_route_ids":    vpc.DataSourceVpcRouteIdsV2(),
			"huaweicloud_vpc_route_v2":     vpc.DataSourceVpcRouteV2(),
			"huaweicloud_vpc_route_ids_v2": vpc.DataSourceVpcRouteIdsV2(),
			"huaweicloud_vpc_ids":          vpc.DataSourceVpcIdsV1(),

			// Deprecated Just discard the resource name, use `huaweicloud_ccm_certificates` instead
			"huaweicloud_scm_certificates": ccm.DataSourceCertificates(),

			// Deprecated
			"huaweicloud_antiddos":                      deprecated.DataSourceAntiDdosV1(),
			"huaweicloud_antiddos_v1":                   deprecated.DataSourceAntiDdosV1(),
			"huaweicloud_compute_availability_zones_v2": deprecated.DataSourceComputeAvailabilityZonesV2(),
			"huaweicloud_csbs_backup":                   deprecated.DataSourceCSBSBackupV1(),
			"huaweicloud_csbs_backup_policy":            deprecated.DataSourceCSBSBackupPolicyV1(),
			"huaweicloud_csbs_backup_policy_v1":         deprecated.DataSourceCSBSBackupPolicyV1(),
			"huaweicloud_csbs_backup_v1":                deprecated.DataSourceCSBSBackupV1(),
			"huaweicloud_networking_network_v2":         deprecated.DataSourceNetworkingNetworkV2(),
			"huaweicloud_networking_subnet_v2":          deprecated.DataSourceNetworkingSubnetV2(),
			"huaweicloud_cts_tracker":                   deprecated.DataSourceCTSTrackerV1(),
			"huaweicloud_dcs_az":                        deprecated.DataSourceDcsAZV1(),
			"huaweicloud_dcs_az_v1":                     deprecated.DataSourceDcsAZV1(),
			"huaweicloud_dcs_product":                   deprecated.DataSourceDcsProductV1(),
			"huaweicloud_dcs_product_v1":                deprecated.DataSourceDcsProductV1(),
			"huaweicloud_dms_az":                        deprecated.DataSourceDmsAZ(),
			"huaweicloud_dms_az_v1":                     deprecated.DataSourceDmsAZ(),
			"huaweicloud_sfs_file_system":               deprecated.DataSourceSFSFileSystemV2(),
			"huaweicloud_sfs_file_system_v2":            deprecated.DataSourceSFSFileSystemV2(),
			"huaweicloud_vbs_backup_policy":             deprecated.DataSourceVBSBackupPolicyV2(),
			"huaweicloud_vbs_backup":                    deprecated.DataSourceVBSBackupV2(),
			"huaweicloud_vbs_backup_policy_v2":          deprecated.DataSourceVBSBackupPolicyV2(),
			"huaweicloud_vbs_backup_v2":                 deprecated.DataSourceVBSBackupV2(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"huaweicloud_aad_forward_rule": aad.ResourceForwardRule(),

			"huaweicloud_antiddos_basic":                     antiddos.ResourceCloudNativeAntiDdos(),
			"huaweicloud_antiddos_default_protection_policy": antiddos.ResourceDefaultProtectionPolicy(),

			"huaweicloud_access_analyzer":                    accessanalyzer.ResourceAccessAnalyzer(),
			"huaweicloud_access_analyzer_archive_rule":       accessanalyzer.ResourceArchiveRule(),
			"huaweicloud_access_analyzer_archive_rule_apply": accessanalyzer.ResourceArchiveRuleApply(),

			"huaweicloud_aom_alarm_rule":                     aom.ResourceAlarmRule(),
			"huaweicloud_aomv4_alarm_rule":                   aom.ResourceAlarmRuleV4(),
			"huaweicloud_aom_event_alarm_rule":               aom.ResourceEventAlarmRule(),
			"huaweicloud_aom_service_discovery_rule":         aom.ResourceServiceDiscoveryRule(),
			"huaweicloud_aom_alarm_action_rule":              aom.ResourceAlarmActionRule(),
			"huaweicloud_aom_alarm_silence_rule":             aom.ResourceAlarmSilenceRule(),
			"huaweicloud_aom_cmdb_application":               aom.ResourceCmdbApplication(),
			"huaweicloud_aom_cmdb_component":                 aom.ResourceCmdbComponent(),
			"huaweicloud_aom_cmdb_environment":               aom.ResourceCmdbEnvironment(),
			"huaweicloud_aom_prom_instance":                  aom.ResourcePromInstance(),
			"huaweicloud_aom_multi_account_aggregation_rule": aom.ResourceMultiAccountAggregationRule(),
			"huaweicloud_aom_dashboards_folder":              aom.ResourceDashboardsFolder(),
			"huaweicloud_aom_message_template":               aom.ResourceMessageTemplate(),
			"huaweicloud_aom_cloud_service_access":           aom.ResourceCloudServiceAccess(),
			"huaweicloud_aom_dashboard":                      aom.ResourceDashboard(),
			"huaweicloud_aom_alarm_rules_template":           aom.ResourceAlarmRulesTemplate(),
			"huaweicloud_aom_alarm_group_rule":               aom.ResourceAlarmGroupRule(),

			"huaweicloud_rfs_execution_plan": rfs.ResourceExecutionPlan(),
			"huaweicloud_rfs_private_hook":   rfs.ResourcePrivateHook(),
			"huaweicloud_rfs_stack":          rfs.ResourceStack(),

			"huaweicloud_api_gateway_api":         apigateway.ResourceAPI(),
			"huaweicloud_api_gateway_environment": apigateway.ResourceEnvironment(),
			"huaweicloud_api_gateway_group":       apigateway.ResourceGroup(),

			"huaweicloud_apig_acl_policy":                     apig.ResourceAclPolicy(),
			"huaweicloud_apig_acl_policy_associate":           apig.ResourceAclPolicyAssociate(),
			"huaweicloud_apig_api":                            apig.ResourceApigAPIV2(),
			"huaweicloud_apig_api_publishment":                apig.ResourceApigApiPublishment(),
			"huaweicloud_apig_appcode":                        apig.ResourceAppcode(),
			"huaweicloud_apig_application":                    apig.ResourceApigApplicationV2(),
			"huaweicloud_apig_application_acl":                apig.ResourceApplicationAcl(),
			"huaweicloud_apig_application_authorization":      apig.ResourceAppAuth(),
			"huaweicloud_apig_application_quota":              apig.ResourceApplicationQuota(),
			"huaweicloud_apig_application_quota_associate":    apig.ResourceApplicationQuotaAssociate(),
			"huaweicloud_apig_certificate":                    apig.ResourceCertificate(),
			"huaweicloud_apig_channel":                        apig.ResourceChannel(),
			"huaweicloud_apig_custom_authorizer":              apig.ResourceApigCustomAuthorizerV2(),
			"huaweicloud_apig_endpoint_connection_management": apig.ResourceEndpointConnectionManagement(),
			"huaweicloud_apig_environment":                    apig.ResourceApigEnvironmentV2(),
			"huaweicloud_apig_environment_variable":           apig.ResourceEnvironmentVariable(),
			"huaweicloud_apig_group":                          apig.ResourceApigGroupV2(),
			"huaweicloud_apig_instance_feature":               apig.ResourceInstanceFeature(),
			"huaweicloud_apig_instance_routes":                apig.ResourceInstanceRoutes(),
			"huaweicloud_apig_instance":                       apig.ResourceApigInstanceV2(),
			"huaweicloud_apig_orchestration_rule":             apig.ResourceOrchestrationRule(),
			"huaweicloud_apig_plugin_associate":               apig.ResourcePluginAssociate(),
			"huaweicloud_apig_plugin":                         apig.ResourcePlugin(),
			"huaweicloud_apig_response":                       apig.ResourceApigResponseV2(),
			"huaweicloud_apig_signature_associate":            apig.ResourceSignatureAssociate(),
			"huaweicloud_apig_signature":                      apig.ResourceSignature(),
			"huaweicloud_apig_throttling_policy_associate":    apig.ResourceThrottlingPolicyAssociate(),
			"huaweicloud_apig_throttling_policy":              apig.ResourceApigThrottlingPolicyV2(),
			"huaweicloud_apig_endpoint_whitelist":             apig.ResourceEndpointWhiteList(),

			"huaweicloud_as_configuration":           as.ResourceASConfiguration(),
			"huaweicloud_as_group":                   as.ResourceASGroup(),
			"huaweicloud_as_lifecycle_hook":          as.ResourceASLifecycleHook(),
			"huaweicloud_as_instance_attach":         as.ResourceASInstanceAttach(),
			"huaweicloud_as_notification":            as.ResourceAsNotification(),
			"huaweicloud_as_policy":                  as.ResourceASPolicy(),
			"huaweicloud_as_bandwidth_policy":        as.ResourceASBandWidthPolicy(),
			"huaweicloud_as_planned_task":            as.ResourcePlannedTask(),
			"huaweicloud_as_lifecycle_hook_callback": as.ResourceLifecycleHookCallBack(),

			"huaweicloud_asm_mesh": asm.ResourceAsmMesh(),

			"huaweicloud_bms_instance": bms.ResourceBmsInstance(),
			"huaweicloud_bcs_instance": bcs.ResourceInstance(),

			"huaweicloud_cae_component_action":         cae.ResourceComponentAction(),
			"huaweicloud_cae_application":              cae.ResourceApplication(),
			"huaweicloud_cae_certificate":              cae.ResourceCertificate(),
			"huaweicloud_cae_component":                cae.ResourceComponent(),
			"huaweicloud_cae_component_configurations": cae.ResourceComponentConfigurations(),
			"huaweicloud_cae_domain":                   cae.ResourceDomain(),
			"huaweicloud_cae_environment":              cae.ResourceEnvironment(),
			"huaweicloud_cae_notification_rule":        cae.ResourceNotificationRule(),
			"huaweicloud_cae_timer_rule":               cae.ResourceTimerRule(),
			"huaweicloud_cae_vpc_egress":               cae.ResourceVpcEgress(),

			"huaweicloud_cbr_backup_share_accepter": cbr.ResourceBackupShareAccepter(),
			"huaweicloud_cbr_backup_share":          cbr.ResourceBackupShare(),
			"huaweicloud_cbr_checkpoint":            cbr.ResourceCheckpoint(),
			"huaweicloud_cbr_policy":                cbr.ResourcePolicy(),
			"huaweicloud_cbr_vault":                 cbr.ResourceVault(),

			"huaweicloud_cbh_instance":                   cbh.ResourceCBHInstance(),
			"huaweicloud_cbh_ha_instance":                cbh.ResourceCBHHAInstance(),
			"huaweicloud_cbh_asset_agency_authorization": cbh.ResourceAssetAgencyAuthorization(),

			"huaweicloud_cc_connection":                                     cc.ResourceCloudConnection(),
			"huaweicloud_cc_network_instance":                               cc.ResourceNetworkInstance(),
			"huaweicloud_cc_authorization":                                  cc.ResourceAuthorization(),
			"huaweicloud_cc_bandwidth_package":                              cc.ResourceBandwidthPackage(),
			"huaweicloud_cc_inter_region_bandwidth":                         cc.ResourceInterRegionBandwidth(),
			"huaweicloud_cc_central_network":                                cc.ResourceCentralNetwork(),
			"huaweicloud_cc_central_network_policy":                         cc.ResourceCentralNetworkPolicy(),
			"huaweicloud_cc_central_network_policy_apply":                   cc.ResourceCentralNetworkPolicyApply(),
			"huaweicloud_cc_central_network_attachment":                     cc.ResourceCentralNetworkAttachment(),
			"huaweicloud_cc_central_network_connection_bandwidth_associate": cc.ResourceCentralNetworkConnectionBandwidthAssociate(),
			"huaweicloud_cc_global_connection_bandwidth":                    cc.ResourceGlobalConnectionBandwidth(),
			"huaweicloud_cc_global_connection_bandwidth_associate":          cc.ResourceGlobalConnectionBandwidthAssociate(),

			"huaweicloud_cce_autopilot_cluster": cceautopilot.ResourceAutopilotCluster(),
			"huaweicloud_cce_autopilot_addon":   cceautopilot.ResourceAutopilotAddon(),

			"huaweicloud_cce_cluster":                    cce.ResourceCluster(),
			"huaweicloud_cce_cluster_log_config":         cce.ResourceClusterLogConfig(),
			"huaweicloud_cce_cluster_upgrade":            cce.ResourceClusterUpgrade(),
			"huaweicloud_cce_node":                       cce.ResourceNode(),
			"huaweicloud_cce_node_attach":                cce.ResourceNodeAttach(),
			"huaweicloud_cce_node_sync":                  cce.ResourceNodeSync(),
			"huaweicloud_cce_addon":                      cce.ResourceAddon(),
			"huaweicloud_cce_node_pool":                  cce.ResourceNodePool(),
			"huaweicloud_cce_node_pool_nodes_add":        cce.ResourcePoolNodesAdd(),
			"huaweicloud_cce_namespace":                  cce.ResourceCCENamespaceV1(),
			"huaweicloud_cce_pvc":                        cce.ResourceCcePersistentVolumeClaimsV1(),
			"huaweicloud_cce_partition":                  cce.ResourcePartition(),
			"huaweicloud_cce_chart":                      cce.ResourceChart(),
			"huaweicloud_cce_cluster_certificate_revoke": cce.ResourceCertificateRevoke(),
			"huaweicloud_cce_node_pool_scale":            cce.ResourceNodePoolScale(),

			"huaweicloud_cts_tracker":      cts.ResourceCTSTracker(),
			"huaweicloud_cts_data_tracker": cts.ResourceCTSDataTracker(),
			"huaweicloud_cts_notification": cts.ResourceCTSNotification(),
			"huaweicloud_cci_namespace":    cci.ResourceCciNamespace(),
			"huaweicloud_cci_network":      cci.ResourceCciNetworkV1(),
			"huaweicloud_cci_pvc":          cci.ResourcePersistentVolumeClaimV1(),

			"huaweicloud_ccm_certificate":                ccm.ResourceCCMCertificate(),
			"huaweicloud_ccm_certificate_apply":          ccm.ResourceCertificateApply(),
			"huaweicloud_ccm_certificate_deploy":         ccm.ResourceCertificateDeploy(),
			"huaweicloud_ccm_certificate_import":         ccm.ResourceCertificateImport(),
			"huaweicloud_ccm_certificate_push":           ccm.ResourceCertificatePush(),
			"huaweicloud_ccm_private_ca":                 ccm.ResourcePrivateCertificateAuthority(),
			"huaweicloud_ccm_private_ca_revoke":          ccm.ResourcePrivateCaRevoke(),
			"huaweicloud_ccm_private_certificate":        ccm.ResourcePrivateCertificate(),
			"huaweicloud_ccm_private_certificate_revoke": ccm.ResourcePrivateCertificateRevoke(),
			"huaweicloud_ccm_private_ca_restore":         ccm.ResourcePrivateCaRestore(),

			"huaweicloud_cdm_cluster":        cdm.ResourceCdmCluster(),
			"huaweicloud_cdm_cluster_action": cdm.ResourceClusterAction(),
			"huaweicloud_cdm_job":            cdm.ResourceCdmJob(),
			"huaweicloud_cdm_link":           cdm.ResourceCdmLink(),

			"huaweicloud_cdn_domain":         cdn.ResourceCdnDomain(),
			"huaweicloud_cdn_domain_rule":    cdn.ResourceCdnDomainRule(),
			"huaweicloud_cdn_billing_option": cdn.ResourceBillingOption(),
			"huaweicloud_cdn_cache_preheat":  cdn.ResourceCachePreheat(),
			"huaweicloud_cdn_cache_refresh":  cdn.ResourceCacheRefresh(),

			"huaweicloud_ces_alarmrule":        ces.ResourceAlarmRule(),
			"huaweicloud_ces_alarm_template":   ces.ResourceCesAlarmTemplate(),
			"huaweicloud_ces_dashboard":        ces.ResourceDashboard(),
			"huaweicloud_ces_dashboard_widget": ces.ResourceDashboardWidget(),
			"huaweicloud_ces_event_report":     ces.ResourceCesEventReport(),
			"huaweicloud_ces_metric_data_add":  ces.ResourceMetricDataAdd(),
			"huaweicloud_ces_one_click_alarm":  ces.ResourceOneClickAlarm(),
			"huaweicloud_ces_resource_group":   ces.ResourceResourceGroup(),

			"huaweicloud_cfw_acl_rule":             cfw.ResourceAclRule(),
			"huaweicloud_cfw_address_group":        cfw.ResourceAddressGroup(),
			"huaweicloud_cfw_address_group_member": cfw.ResourceAddressGroupMember(),
			"huaweicloud_cfw_alarm_config":         cfw.ResourceAlarmConfig(),
			"huaweicloud_cfw_anti_virus":           cfw.ResourceAntiVirus(),
			"huaweicloud_cfw_black_white_list":     cfw.ResourceBlackWhiteList(),
			"huaweicloud_cfw_eip_protection":       cfw.ResourceEipProtection(),
			"huaweicloud_cfw_service_group":        cfw.ResourceServiceGroup(),
			"huaweicloud_cfw_service_group_member": cfw.ResourceServiceGroupMember(),
			"huaweicloud_cfw_firewall":             cfw.ResourceFirewall(),
			"huaweicloud_cfw_domain_name_group":    cfw.ResourceDomainNameGroup(),
			"huaweicloud_cfw_lts_log":              cfw.ResourceLtsLog(),
			"huaweicloud_cfw_dns_resolution":       cfw.ResourceDNSResolution(),
			"huaweicloud_cfw_capture_task":         cfw.ResourceCaptureTask(),
			"huaweicloud_cfw_ips_rule_mode_change": cfw.ResourceCfwIpsRuleModeChange(),

			"huaweicloud_cloudtable_cluster": cloudtable.ResourceCloudTableCluster(),

			"huaweicloud_cnad_advanced_black_white_list": cnad.ResourceBlackWhiteList(),
			"huaweicloud_cnad_advanced_policy":           cnad.ResourceCNADAdvancedPolicy(),
			"huaweicloud_cnad_advanced_policy_associate": cnad.ResourcePolicyAssociate(),
			"huaweicloud_cnad_advanced_protected_object": cnad.ResourceProtectedObject(),

			"huaweicloud_compute_instance":          ecs.ResourceComputeInstance(),
			"huaweicloud_compute_interface_attach":  ecs.ResourceComputeInterfaceAttach(),
			"huaweicloud_compute_keypair":           ResourceComputeKeypairV2(),
			"huaweicloud_compute_servergroup":       ecs.ResourceComputeServerGroup(),
			"huaweicloud_compute_eip_associate":     ecs.ResourceComputeEIPAssociate(),
			"huaweicloud_compute_volume_attach":     ecs.ResourceComputeVolumeAttach(),
			"huaweicloud_compute_auto_launch_group": ecs.ResourceComputeAutoLaunchGroup(),

			"huaweicloud_coc_script":         coc.ResourceScript(),
			"huaweicloud_coc_script_execute": coc.ResourceScriptExecute(),

			"huaweicloud_cph_server":         cph.ResourceCphServer(),
			"huaweicloud_cph_adb_command":    cph.ResourceAdbCommand(),
			"huaweicloud_cph_phone_stop":     cph.ResourcePhoneStop(),
			"huaweicloud_cph_server_restart": cph.ResourceServerRestart(),
			"huaweicloud_cph_phone_reset":    cph.ResourcePhoneReset(),
			"huaweicloud_cph_share_app":      cph.ResourceShareApp(),
			"huaweicloud_cph_phone_property": cph.ResourcePhoneProperty(),
			"huaweicloud_cph_phone_restart":  cph.ResourcePhoneRestart(),

			"huaweicloud_cse_microservice":                      cse.ResourceMicroservice(),
			"huaweicloud_cse_microservice_engine":               cse.ResourceMicroserviceEngine(),
			"huaweicloud_cse_microservice_engine_configuration": cse.ResourceMicroserviceEngineConfiguration(),
			"huaweicloud_cse_microservice_instance":             cse.ResourceMicroserviceInstance(),
			"huaweicloud_cse_nacos_namespace":                   cse.ResourceNacosNamespace(),

			"huaweicloud_csms_event":                dew.ResourceCsmsEvent(),
			"huaweicloud_csms_secret":               dew.ResourceSecret(),
			"huaweicloud_csms_secret_version_state": dew.ResourceSecretVersionState(),

			"huaweicloud_css_cluster":                     css.ResourceCssCluster(),
			"huaweicloud_css_cluster_restart":             css.ResourceCssClusterRestart(),
			"huaweicloud_css_cluster_node_replace":        css.ResourceCssClusterNodeReplace(),
			"huaweicloud_css_snapshot":                    css.ResourceCssSnapshot(),
			"huaweicloud_css_thesaurus":                   css.ResourceCssthesaurus(),
			"huaweicloud_css_configuration":               css.ResourceCssConfiguration(),
			"huaweicloud_css_scan_task":                   css.ResourceScanTask(),
			"huaweicloud_css_es_loadbalancer_config":      css.ResourceEsLoadbalancerConfig(),
			"huaweicloud_css_es_core_upgrade":             css.ResourceCssEsCoreUpgrade(),
			"huaweicloud_css_cluster_az_migrate":          css.ResourceCssClusterAzMigrate(),
			"huaweicloud_css_log_setting":                 css.ResourceLogSetting(),
			"huaweicloud_css_manual_log_backup":           css.ResourceManualLogBackup(),
			"huaweicloud_css_logstash_cluster":            css.ResourceLogstashCluster(),
			"huaweicloud_css_logstash_cluster_restart":    css.ResourceLogstashClusterRestart(),
			"huaweicloud_css_logstash_configuration":      css.ResourceLogstashConfiguration(),
			"huaweicloud_css_logstash_connectivity":       css.ResourceLogstashConnectivity(),
			"huaweicloud_css_logstash_pipeline":           css.ResourceLogstashPipeline(),
			"huaweicloud_css_logstash_custom_certificate": css.ResourceLogstashCertificate(),
			"huaweicloud_css_logstash_custom_template":    css.ResourceLogstashCustomTemplate(),

			"huaweicloud_dbss_audit_risk_rule_action": dbss.ResourceRiskRuleAction(),
			"huaweicloud_dbss_ecs_database":           dbss.ResourceAddEcsDatabase(),
			"huaweicloud_dbss_instance":               dbss.ResourceInstance(),
			"huaweicloud_dbss_rds_database":           dbss.ResourceAddRdsDatabase(),

			"huaweicloud_dc_virtual_gateway":            dc.ResourceVirtualGateway(),
			"huaweicloud_dc_virtual_interface":          dc.ResourceVirtualInterface(),
			"huaweicloud_dc_virtual_interface_accepter": dc.ResourceInterfaceAccepter(),
			"huaweicloud_dc_hosted_connect":             dc.ResourceHostedConnect(),

			"huaweicloud_dcs_instance":         dcs.ResourceDcsInstance(),
			"huaweicloud_dcs_backup":           dcs.ResourceDcsBackup(),
			"huaweicloud_dcs_custom_template":  dcs.ResourceCustomTemplate(),
			"huaweicloud_dcs_hotkey_analysis":  dcs.ResourceHotKeyAnalysis(),
			"huaweicloud_dcs_bigkey_analysis":  dcs.ResourceBigKeyAnalysis(),
			"huaweicloud_dcs_account":          dcs.ResourceDcsAccount(),
			"huaweicloud_dcs_instance_restore": dcs.ResourceDcsRestore(),
			"huaweicloud_dcs_diagnosis_task":   dcs.ResourceDiagnosisTask(),

			"huaweicloud_dds_database_role":               dds.ResourceDatabaseRole(),
			"huaweicloud_dds_database_user":               dds.ResourceDatabaseUser(),
			"huaweicloud_dds_instance":                    dds.ResourceDdsInstanceV3(),
			"huaweicloud_dds_backup":                      dds.ResourceDdsBackup(),
			"huaweicloud_dds_parameter_template":          dds.ResourceDdsParameterTemplate(),
			"huaweicloud_dds_audit_log_policy":            dds.ResourceDdsAuditLogPolicy(),
			"huaweicloud_dds_audit_log_delete":            dds.ResourceDDSAuditLogDelete(),
			"huaweicloud_dds_lts_log":                     dds.ResourceDdsLtsLog(),
			"huaweicloud_dds_instance_restart":            dds.ResourceDDSInstanceRestart(),
			"huaweicloud_dds_instance_internal_ip_modify": dds.ResourceDDSInstanceModifyIP(),
			"huaweicloud_dds_instance_eip_associate":      dds.ResourceDDSInstanceBindEIP(),
			"huaweicloud_dds_instance_restore":            dds.ResourceDDSInstanceRestore(),
			"huaweicloud_dds_collection_restore":          dds.ResourceDDSCollectionRestore(),
			"huaweicloud_dds_instance_parameters_modify":  dds.ResourceDDSInstanceParametersModify(),
			"huaweicloud_dds_primary_standby_switch":      dds.ResourceDDSPrimaryStandbySwitch(),
			"huaweicloud_dds_recycle_policy":              dds.ResourceDDSRecyclePolicy(),
			"huaweicloud_dds_parameter_template_reset":    dds.ResourceDDSParameterTemplateReset(),
			"huaweicloud_dds_parameter_template_copy":     dds.ResourceDDSParameterTemplateCopy(),
			"huaweicloud_dds_parameter_template_compare":  dds.ResourceDDSParameterTemplateCompare(),
			"huaweicloud_dds_parameter_template_apply":    dds.ResourceDDSParameterTemplateApply(),
			"huaweicloud_dds_scheduled_task_cancel":       dds.ResourceDDSScheduledTaskCancel(),

			"huaweicloud_ddm_instance":               ddm.ResourceDdmInstance(),
			"huaweicloud_ddm_instance_restart":       ddm.ResourceDdmInstanceRestart(),
			"huaweicloud_ddm_instance_upgrade":       ddm.ResourceDdmInstanceUpgrade(),
			"huaweicloud_ddm_instance_rollback":      ddm.ResourceDdmInstanceRollback(),
			"huaweicloud_ddm_schema":                 ddm.ResourceDdmSchema(),
			"huaweicloud_ddm_account":                ddm.ResourceDdmAccount(),
			"huaweicloud_ddm_instance_read_strategy": ddm.ResourceDdmInstanceReadStrategy(),
			"huaweicloud_ddm_physical_sessions_kill": ddm.ResourceDdmPhysicalSessionsKill(),
			"huaweicloud_ddm_logical_sessions_kill":  ddm.ResourceDdmLogicalSessionsKill(),

			"huaweicloud_dis_stream": dis.ResourceDisStream(),

			"huaweicloud_dli_database":                        dli.ResourceDliSqlDatabaseV1(),
			"huaweicloud_dli_database_privilege":              dli.ResourceDatabasePrivilege(),
			"huaweicloud_dli_elastic_resource_pool":           dli.ResourceElasticResourcePool(),
			"huaweicloud_dli_package":                         dli.ResourceDliPackageV2(),
			"huaweicloud_dli_queue":                           dli.ResourceDliQueue(),
			"huaweicloud_dli_spark_job":                       dli.ResourceDliSparkJobV2(),
			"huaweicloud_dli_sql_job":                         dli.ResourceSqlJob(),
			"huaweicloud_dli_table":                           dli.ResourceDliTable(),
			"huaweicloud_dli_flinksql_job":                    dli.ResourceFlinkSqlJob(),
			"huaweicloud_dli_flinkjar_job":                    dli.ResourceFlinkJarJob(),
			"huaweicloud_dli_permission":                      dli.ResourceDliPermission(),
			"huaweicloud_dli_datasource_connection":           dli.ResourceDatasourceConnection(),
			"huaweicloud_dli_datasource_connection_associate": dli.ResourceDatasourceConnectionAssociate(),
			"huaweicloud_dli_datasource_connection_privilege": dli.ResourceDatasourceConnectionPrivilege(),
			"huaweicloud_dli_datasource_auth":                 dli.ResourceDatasourceAuth(),
			"huaweicloud_dli_sql_template":                    dli.ResourceSQLTemplate(),
			"huaweicloud_dli_flink_template":                  dli.ResourceFlinkTemplate(),
			"huaweicloud_dli_global_variable":                 dli.ResourceGlobalVariable(),
			"huaweicloud_dli_spark_template":                  dli.ResourceSparkTemplate(),
			"huaweicloud_dli_agency":                          dli.ResourceDliAgency(),

			"huaweicloud_dms_kafka_background_task_delete":    dms.ResourceDmsKafkaBackgroundTaskDelete(),
			"huaweicloud_dms_kafka_user":                      dms.ResourceDmsKafkaUser(),
			"huaweicloud_dms_kafka_permissions":               dms.ResourceDmsKafkaPermissions(),
			"huaweicloud_dms_kafka_instance":                  dms.ResourceDmsKafkaInstance(),
			"huaweicloud_dms_kafka_instance_restart":          dms.ResourceDmsKafkaInstanceRestart(),
			"huaweicloud_dms_kafka_topic":                     dms.ResourceDmsKafkaTopic(),
			"huaweicloud_dms_kafka_message_produce":           dms.ResourceDmsKafkaMessageProduce(),
			"huaweicloud_dms_kafka_partition_reassign":        dms.ResourceDmsKafkaPartitionReassign(),
			"huaweicloud_dms_kafka_consumer_group":            dms.ResourceDmsKafkaConsumerGroup(),
			"huaweicloud_dms_kafka_message_offset_reset":      dms.ResourceDmsKafkaMessageOffsetReset(),
			"huaweicloud_dms_kafka_smart_connect":             dms.ResourceDmsKafkaSmartConnect(),
			"huaweicloud_dms_kafka_smart_connect_task":        dms.ResourceDmsKafkaSmartConnectTask(),
			"huaweicloud_dms_kafkav2_smart_connect_task":      dms.ResourceDmsKafkav2SmartConnectTask(),
			"huaweicloud_dms_kafka_smart_connect_task_action": dms.ResourceDmsKafkaSmartConnectTaskAction(),
			"huaweicloud_dms_kafka_user_client_quota":         dms.ResourceDmsKafkaUserClientQuota(),
			"huaweicloud_dms_kafka_message_diagnosis_task":    dms.ResourceDmsKafkaMessageDiagnosisTask(),

			"huaweicloud_dms_rabbitmq_background_task_delete": dms.ResourceDmsRabbitMQBackgroundTaskDelete(),
			"huaweicloud_dms_rabbitmq_instance":               dms.ResourceDmsRabbitmqInstance(),
			"huaweicloud_dms_rabbitmq_plugin":                 dms.ResourceDmsRabbitmqPlugin(),
			"huaweicloud_dms_rabbitmq_vhost":                  dms.ResourceDmsRabbitmqVhost(),
			"huaweicloud_dms_rabbitmq_exchange":               dms.ResourceDmsRabbitmqExchange(),
			"huaweicloud_dms_rabbitmq_queue":                  dms.ResourceDmsRabbitmqQueue(),
			"huaweicloud_dms_rabbitmq_queue_message_clear":    dms.ResourceDmsRabbitmqQueueMessageClear(),
			"huaweicloud_dms_rabbitmq_exchange_associate":     dms.ResourceDmsRabbitmqExchangeAssociate(),
			"huaweicloud_dms_rabbitmq_user":                   dms.ResourceDmsRabbitmqUser(),

			"huaweicloud_dms_rocketmq_instance":             dms.ResourceDmsRocketMQInstance(),
			"huaweicloud_dms_rocketmq_consumer_group":       dms.ResourceDmsRocketMQConsumerGroup(),
			"huaweicloud_dms_rocketmq_consumption_verify":   dms.ResourceDmsRocketMQConsumptionVerify(),
			"huaweicloud_dms_rocketmq_message_offset_reset": dms.ResourceDmsRocketMQMessageOffsetReset(),
			"huaweicloud_dms_rocketmq_dead_letter_resend":   dms.ResourceDmsRocketMQDeadLetterResend(),
			"huaweicloud_dms_rocketmq_topic":                dms.ResourceDmsRocketMQTopic(),
			"huaweicloud_dms_rocketmq_user":                 dms.ResourceDmsRocketMQUser(),
			"huaweicloud_dms_rocketmq_migration_task":       dms.ResourceDmsRocketmqMigrationTask(),

			"huaweicloud_dns_custom_line":             dns.ResourceDNSCustomLine(),
			"huaweicloud_dns_ptrrecord":               dns.ResourceDNSPtrRecord(),
			"huaweicloud_dns_recordset":               dns.ResourceDNSRecordset(),
			"huaweicloud_dns_zone":                    dns.ResourceDNSZone(),
			"huaweicloud_dns_endpoint_assignment":     dns.ResourceEndpointAssignment(),
			"huaweicloud_dns_endpoint":                dns.ResourceDNSEndpoint(),
			"huaweicloud_dns_resolver_rule":           dns.ResourceDNSResolverRule(),
			"huaweicloud_dns_resolver_rule_associate": dns.ResourceDNSResolverRuleAssociate(),
			"huaweicloud_dns_line_group":              dns.ResourceDNSLineGroup(),

			"huaweicloud_drs_job":                        drs.ResourceDrsJob(),
			"huaweicloud_drs_job_primary_standby_switch": drs.ResourceDRSPrimaryStandbySwitch(),

			"huaweicloud_dws_alarm_subscription":            dws.ResourceDwsAlarmSubs(),
			"huaweicloud_dws_cluster_restart":               dws.ResourceClusterRestart(),
			"huaweicloud_dws_cluster":                       dws.ResourceDwsCluster(),
			"huaweicloud_dws_disaster_recovery_task":        dws.ResourceDwsDisasterRecoveryTask(),
			"huaweicloud_dws_event_subscription":            dws.ResourceDwsEventSubs(),
			"huaweicloud_dws_ext_data_source":               dws.ResourceDwsExtDataSource(),
			"huaweicloud_dws_logical_cluster_restart":       dws.ResourceLogicalClusterRestart(),
			"huaweicloud_dws_logical_cluster":               dws.ResourceLogicalCluster(),
			"huaweicloud_dws_om_account_action":             dws.ResourceOmAccountAction(),
			"huaweicloud_dws_parameter_configurations":      dws.ResourceParameterConfigurations(),
			"huaweicloud_dws_public_domain_associate":       dws.ResourcePublicDomainAssociate(),
			"huaweicloud_dws_snapshot_copy":                 dws.ResourceSnapshotCopy(),
			"huaweicloud_dws_snapshot_policy":               dws.ResourceDwsSnapshotPolicy(),
			"huaweicloud_dws_snapshot":                      dws.ResourceDwsSnapshot(),
			"huaweicloud_dws_workload_configuration":        dws.ResourceWorkLoadConfiguration(),
			"huaweicloud_dws_workload_plan_execution":       dws.ResourceWorkLoadPlanExecution(),
			"huaweicloud_dws_workload_plan_stage":           dws.ResourceWorkLoadPlanStage(),
			"huaweicloud_dws_workload_plan":                 dws.ResourceWorkLoadPlan(),
			"huaweicloud_dws_workload_queue_user_associate": dws.ResourceWorkloadQueueUserAssociate(),
			"huaweicloud_dws_workload_queue":                dws.ResourceWorkLoadQueue(),

			"huaweicloud_eg_connection":           eg.ResourceConnection(),
			"huaweicloud_eg_custom_event_channel": eg.ResourceCustomEventChannel(),
			"huaweicloud_eg_custom_event_source":  eg.ResourceCustomEventSource(),
			"huaweicloud_eg_endpoint":             eg.ResourceEndpoint(),
			"huaweicloud_eg_event_stream":         eg.ResourceEventStream(),
			"huaweicloud_eg_event_subscription":   eg.ResourceEventSubscription(),

			"huaweicloud_elb_certificate":                  elb.ResourceCertificateV3(),
			"huaweicloud_elb_certificate_private_key_echo": elb.ResourceCertificatePrivateKeyEcho(),
			"huaweicloud_elb_l7policy":                     elb.ResourceL7PolicyV3(),
			"huaweicloud_elb_l7rule":                       elb.ResourceL7RuleV3(),
			"huaweicloud_elb_listener":                     elb.ResourceListenerV3(),
			"huaweicloud_elb_loadbalancer":                 elb.ResourceLoadBalancerV3(),
			"huaweicloud_elb_monitor":                      elb.ResourceMonitorV3(),
			"huaweicloud_elb_ipgroup":                      elb.ResourceIpGroupV3(),
			"huaweicloud_elb_pool":                         elb.ResourcePoolV3(),
			"huaweicloud_elb_active_standby_pool":          elb.ResourceActiveStandbyPool(),
			"huaweicloud_elb_member":                       elb.ResourceMemberV3(),
			"huaweicloud_elb_logtank":                      elb.ResourceLogTank(),
			"huaweicloud_elb_security_policy":              elb.ResourceSecurityPolicy(),

			"huaweicloud_enterprise_project":           eps.ResourceEnterpriseProject(),
			"huaweicloud_enterprise_project_authority": eps.ResourceAuthority(),

			"huaweicloud_er_association":    er.ResourceAssociation(),
			"huaweicloud_er_instance":       er.ResourceInstance(),
			"huaweicloud_er_propagation":    er.ResourcePropagation(),
			"huaweicloud_er_route_table":    er.ResourceRouteTable(),
			"huaweicloud_er_static_route":   er.ResourceStaticRoute(),
			"huaweicloud_er_vpc_attachment": er.ResourceVpcAttachment(),
			"huaweicloud_er_flow_log":       er.ResourceFlowLog(),

			"huaweicloud_evs_snapshot":                 evs.ResourceEvsSnapshotV2(),
			"huaweicloud_evs_volume":                   evs.ResourceEvsVolume(),
			"huaweicloud_evs_snapshot_rollback":        evs.ResourceSnapshotRollBack(),
			"huaweicloud_evs_volume_transfer":          evs.ResourceVolumeTransfer(),
			"huaweicloud_evs_volume_transfer_accepter": evs.ResourceVolumeTransferAccepter(),

			"huaweicloud_fgs_application":                fgs.ResourceApplication(),
			"huaweicloud_fgs_async_invoke_configuration": fgs.ResourceAsyncInvokeConfiguration(),
			"huaweicloud_fgs_dependency":                 fgs.ResourceDependency(),
			"huaweicloud_fgs_dependency_version":         fgs.ResourceDependencyVersion(),
			"huaweicloud_fgs_function":                   fgs.ResourceFgsFunctionV2(),
			"huaweicloud_fgs_function_event":             fgs.ResourceFunctionEvent(),
			"huaweicloud_fgs_function_topping":           fgs.ResourceFunctionTopping(),
			"huaweicloud_fgs_function_trigger":           fgs.ResourceFunctionTrigger(),

			"huaweicloud_ga_accelerator":    ga.ResourceAccelerator(),
			"huaweicloud_ga_access_log":     ga.ResourceAccessLog(),
			"huaweicloud_ga_address_group":  ga.ResourceIpAddressGroup(),
			"huaweicloud_ga_listener":       ga.ResourceListener(),
			"huaweicloud_ga_endpoint_group": ga.ResourceEndpointGroup(),
			"huaweicloud_ga_endpoint":       ga.ResourceEndpoint(),
			"huaweicloud_ga_health_check":   ga.ResourceHealthCheck(),

			"huaweicloud_gaussdb_cassandra_instance":  geminidb.ResourceGeminiDBInstanceV3(),
			"huaweicloud_gaussdb_redis_instance":      geminidb.ResourceGaussRedisInstanceV3(),
			"huaweicloud_gaussdb_redis_eip_associate": geminidb.ResourceGaussRedisEipAssociate(),
			"huaweicloud_gaussdb_influx_instance":     geminidb.ResourceGaussDBInfluxInstanceV3(),
			"huaweicloud_gaussdb_mongo_instance":      geminidb.ResourceGaussDBMongoInstanceV3(),

			"huaweicloud_gaussdb_mysql_instance":                   taurusdb.ResourceGaussDBInstance(),
			"huaweicloud_gaussdb_mysql_instance_node_config":       taurusdb.ResourceGaussDBMysqlNodeConfig(),
			"huaweicloud_gaussdb_mysql_instance_restart":           taurusdb.ResourceGaussDBMysqlRestart(),
			"huaweicloud_gaussdb_mysql_instance_upgrade":           taurusdb.ResourceGaussDBMysqlUpgrade(),
			"huaweicloud_gaussdb_mysql_proxy":                      taurusdb.ResourceGaussDBProxy(),
			"huaweicloud_gaussdb_mysql_proxy_restart":              taurusdb.ResourceGaussDBProxyRestart(),
			"huaweicloud_gaussdb_mysql_database":                   taurusdb.ResourceGaussDBDatabase(),
			"huaweicloud_gaussdb_mysql_account":                    taurusdb.ResourceGaussDBAccount(),
			"huaweicloud_gaussdb_mysql_account_privilege":          taurusdb.ResourceGaussDBAccountPrivilege(),
			"huaweicloud_gaussdb_mysql_sql_control_rule":           taurusdb.ResourceGaussDBSqlControlRule(),
			"huaweicloud_gaussdb_mysql_parameter_template":         taurusdb.ResourceGaussDBMysqlTemplate(),
			"huaweicloud_gaussdb_mysql_parameter_template_apply":   taurusdb.ResourceGaussDBMysqlTemplateApply(),
			"huaweicloud_gaussdb_mysql_parameter_template_compare": taurusdb.ResourceGaussDBMysqlTemplateCompare(),
			"huaweicloud_gaussdb_mysql_backup":                     taurusdb.ResourceGaussDBMysqlBackup(),
			"huaweicloud_gaussdb_mysql_lts_log":                    taurusdb.ResourceGaussDBMysqlLtsLog(),
			"huaweicloud_gaussdb_mysql_restore":                    taurusdb.ResourceGaussDBMysqlRestore(),
			"huaweicloud_gaussdb_mysql_table_restore":              taurusdb.ResourceGaussDBMysqlTableRestore(),
			"huaweicloud_gaussdb_mysql_eip_associate":              taurusdb.ResourceGaussMysqlEipAssociate(),
			"huaweicloud_gaussdb_mysql_recycling_policy":           taurusdb.ResourceGaussDBRecyclingPolicy(),
			"huaweicloud_gaussdb_mysql_quota":                      taurusdb.ResourceGaussDBMysqlQuota(),
			"huaweicloud_gaussdb_mysql_scheduled_task_cancel":      taurusdb.ResourceGaussDBScheduledTaskCancel(),
			"huaweicloud_gaussdb_mysql_scheduled_task_delete":      taurusdb.ResourceGaussDBScheduledTaskDelete(),
			"huaweicloud_gaussdb_mysql_instant_task_delete":        taurusdb.ResourceGaussDBInstantTaskDelete(),

			"huaweicloud_gaussdb_opengauss_instance":                   gaussdb.ResourceOpenGaussInstance(),
			"huaweicloud_gaussdb_opengauss_instance_restart":           gaussdb.ResourceOpenGaussInstanceRestart(),
			"huaweicloud_gaussdb_opengauss_instance_upgrade":           gaussdb.ResourceOpenGaussInstanceUpgrade(),
			"huaweicloud_gaussdb_opengauss_instance_node_startup":      gaussdb.ResourceOpenGaussInstanceNodeStartup(),
			"huaweicloud_gaussdb_opengauss_instance_node_stop":         gaussdb.ResourceOpenGaussInstanceNodeStop(),
			"huaweicloud_gaussdb_opengauss_database":                   gaussdb.ResourceOpenGaussDatabase(),
			"huaweicloud_gaussdb_opengauss_schema":                     gaussdb.ResourceOpenGaussSchema(),
			"huaweicloud_gaussdb_opengauss_backup":                     gaussdb.ResourceGaussDBOpenGaussBackup(),
			"huaweicloud_gaussdb_opengauss_backup_stop":                gaussdb.ResourceOpenGaussBackupStop(),
			"huaweicloud_gaussdb_opengauss_restore":                    gaussdb.ResourceOpenGaussRestore(),
			"huaweicloud_gaussdb_opengauss_eip_associate":              gaussdb.ResourceOpenGaussEipAssociate(),
			"huaweicloud_gaussdb_opengauss_primary_standby_switch":     gaussdb.ResourceOpenGaussPrimaryStandbySwitch(),
			"huaweicloud_gaussdb_opengauss_parameter_template":         gaussdb.ResourceOpenGaussParameterTemplate(),
			"huaweicloud_gaussdb_opengauss_parameter_template_apply":   gaussdb.ResourceOpenGaussParameterTemplateApply(),
			"huaweicloud_gaussdb_opengauss_parameter_template_compare": gaussdb.ResourceOpenGaussParameterTemplateCompare(),
			"huaweicloud_gaussdb_opengauss_parameter_template_reset":   gaussdb.ResourceOpenGaussParameterTemplateReset(),
			"huaweicloud_gaussdb_opengauss_recycling_policy":           gaussdb.ResourceOpenGaussRecyclingPolicy(),
			"huaweicloud_gaussdb_opengauss_task_delete":                gaussdb.ResourceOpenGaussTaskDelete(),
			"huaweicloud_gaussdb_opengauss_sync_sql_throttling_task":   gaussdb.ResourceOpenGaussSyncSqlThrottlingTask(),
			"huaweicloud_gaussdb_opengauss_quota":                      gaussdb.ResourceOpenGaussQuota(),
			"huaweicloud_gaussdb_opengauss_sql_throttling_task":        gaussdb.ResourceOpenGaussSqlThrottlingTask(),

			"huaweicloud_ges_graph":    ges.ResourceGesGraph(),
			"huaweicloud_ges_metadata": ges.ResourceGesMetadata(),
			"huaweicloud_ges_backup":   ges.ResourceGesBackup(),

			"huaweicloud_hss_host_group":           hss.ResourceHostGroup(),
			"huaweicloud_hss_cce_protection":       hss.ResourceCCEProtection(),
			"huaweicloud_hss_host_protection":      hss.ResourceHostProtection(),
			"huaweicloud_hss_webtamper_protection": hss.ResourceWebTamperProtection(),
			"huaweicloud_hss_quota":                hss.ResourceQuota(),

			"huaweicloud_identity_access_key":            iam.ResourceIdentityKey(),
			"huaweicloud_identity_acl":                   iam.ResourceIdentityACL(),
			"huaweicloud_identity_agency":                iam.ResourceIAMAgencyV3(),
			"huaweicloud_identity_service_agency":        iam.ResourceIAMServiceAgency(),
			"huaweicloud_identity_trust_agency":          iam.ResourceIAMTrustAgency(),
			"huaweicloud_identity_group":                 iam.ResourceIdentityGroup(),
			"huaweicloud_identity_group_membership":      iam.ResourceIdentityGroupMembership(),
			"huaweicloud_identity_group_role_assignment": iam.ResourceIdentityGroupRoleAssignment(),
			"huaweicloud_identity_project":               iam.ResourceIdentityProject(),
			"huaweicloud_identity_role":                  iam.ResourceIdentityRole(),
			"huaweicloud_identity_role_assignment":       iam.ResourceIdentityGroupRoleAssignment(),
			"huaweicloud_identity_user":                  iam.ResourceIdentityUser(),
			"huaweicloud_identity_user_role_assignment":  iam.ResourceIdentityUserRoleAssignment(),
			"huaweicloud_identity_provider":              iam.ResourceIdentityProvider(),
			"huaweicloud_identity_password_policy":       iam.ResourceIdentityPasswordPolicy(),
			"huaweicloud_identity_protection_policy":     iam.ResourceIdentityProtectionPolicy(),
			"huaweicloud_identity_login_policy":          iam.ResourceIdentityLoginPolicy(),
			"huaweicloud_identity_virtual_mfa_device":    iam.ResourceIdentityVirtualMFADevice(),
			"huaweicloud_identity_user_token":            iam.ResourceIdentityUserToken(),
			"huaweicloud_identity_policy":                iam.ResourceIdentityPolicy(),

			"huaweicloud_identitycenter_user":                                   identitycenter.ResourceIdentityCenterUser(),
			"huaweicloud_identitycenter_group":                                  identitycenter.ResourceIdentityCenterGroup(),
			"huaweicloud_identitycenter_group_membership":                       identitycenter.ResourceGroupMembership(),
			"huaweicloud_identitycenter_permission_set":                         identitycenter.ResourcePermissionSet(),
			"huaweicloud_identitycenter_system_policy_attachment":               identitycenter.ResourceSystemPolicyAttachment(),
			"huaweicloud_identitycenter_system_identity_policy_attachment":      identitycenter.ResourceSystemIdentityPolicyAttachment(),
			"huaweicloud_identitycenter_account_assignment":                     identitycenter.ResourceIdentityCenterAccountAssignment(),
			"huaweicloud_identitycenter_custom_policy_attachment":               identitycenter.ResourceCustomPolicyAttachment(),
			"huaweicloud_identitycenter_custom_role_attachment":                 identitycenter.ResourceCustomRoleAttachment(),
			"huaweicloud_identitycenter_access_control_attribute_configuration": identitycenter.ResourceAccessControlAttributeConfiguration(),
			"huaweicloud_identitycenter_provision_permission_set":               identitycenter.ResourceProvisionPermissionSet(),

			"huaweicloud_iec_eip":                 iec.ResourceEip(),
			"huaweicloud_iec_keypair":             iec.ResourceKeypair(),
			"huaweicloud_iec_network_acl":         iec.ResourceNetworkACL(),
			"huaweicloud_iec_network_acl_rule":    iec.ResourceNetworkACLRule(),
			"huaweicloud_iec_security_group_rule": iec.ResourceSecurityGroupRule(),
			"huaweicloud_iec_security_group":      iec.ResourceSecurityGroup(),
			"huaweicloud_iec_server":              iec.ResourceServer(),
			"huaweicloud_iec_vip":                 iec.ResourceVip(),
			"huaweicloud_iec_vpc":                 iec.ResourceVpc(),
			"huaweicloud_iec_vpc_subnet":          iec.ResourceSubnet(),

			"huaweicloud_ims_ecs_system_image":        ims.ResourceEcsSystemImage(),
			"huaweicloud_ims_ecs_whole_image":         ims.ResourceEcsWholeImage(),
			"huaweicloud_ims_cbr_whole_image":         ims.ResourceCbrWholeImage(),
			"huaweicloud_ims_evs_data_image":          ims.ResourceEvsDataImage(),
			"huaweicloud_ims_evs_system_image":        ims.ResourceEvsSystemImage(),
			"huaweicloud_ims_obs_data_image":          ims.ResourceObsDataImage(),
			"huaweicloud_ims_obs_system_image":        ims.ResourceObsSystemImage(),
			"huaweicloud_ims_obs_iso_image":           ims.ResourceObsIsoImage(),
			"huaweicloud_ims_image_export":            ims.ResourceImageExport(),
			"huaweicloud_images_image_copy":           ims.ResourceImsImageCopy(),
			"huaweicloud_images_image_share":          ims.ResourceImsImageShare(),
			"huaweicloud_images_image_share_accepter": ims.ResourceImsImageShareAccepter(),

			"huaweicloud_iotda_access_credential":        iotda.ResourceAccessCredential(),
			"huaweicloud_iotda_amqp":                     iotda.ResourceAmqp(),
			"huaweicloud_iotda_batchtask":                iotda.ResourceBatchTask(),
			"huaweicloud_iotda_dataforwarding_rule":      iotda.ResourceDataForwardingRule(),
			"huaweicloud_iotda_data_flow_control_policy": iotda.ResourceDataFlowControlPolicy(),
			"huaweicloud_iotda_data_backlog_policy":      iotda.ResourceDataBacklogPolicy(),
			"huaweicloud_iotda_device_message":           iotda.ResourceDeviceMessage(),
			"huaweicloud_iotda_device":                   iotda.ResourceDevice(),
			"huaweicloud_iotda_device_async_command":     iotda.ResourceDeviceAsyncCommand(),
			"huaweicloud_iotda_device_certificate":       iotda.ResourceDeviceCertificate(),
			"huaweicloud_iotda_device_group":             iotda.ResourceDeviceGroup(),
			"huaweicloud_iotda_device_linkage_rule":      iotda.ResourceDeviceLinkageRule(),
			"huaweicloud_iotda_device_proxy":             iotda.ResourceDeviceProxy(),
			"huaweicloud_iotda_product":                  iotda.ResourceProduct(),
			"huaweicloud_iotda_space":                    iotda.ResourceSpace(),
			"huaweicloud_iotda_upgrade_package":          iotda.ResourceUpgradePackage(),

			"huaweicloud_kms_data_encrypt_decrypt": dew.ResourceKmsDataEncryptDecrypt(),
			"huaweicloud_kms_key":                  dew.ResourceKmsKey(),
			"huaweicloud_kps_keypair":              dew.ResourceKeypair(),
			"huaweicloud_kms_grant":                dew.ResourceKmsGrant(),
			"huaweicloud_kms_dedicated_keystore":   dew.ResourceKmsDedicatedKeystore(),
			"huaweicloud_kms_key_material":         dew.ResourceKmsKeyMaterial(),

			"huaweicloud_lb_certificate":  lb.ResourceCertificateV2(),
			"huaweicloud_lb_l7policy":     lb.ResourceL7PolicyV2(),
			"huaweicloud_lb_l7rule":       lb.ResourceL7RuleV2(),
			"huaweicloud_lb_loadbalancer": lb.ResourceLoadBalancer(),
			"huaweicloud_lb_listener":     lb.ResourceListener(),
			"huaweicloud_lb_member":       lb.ResourceMemberV2(),
			"huaweicloud_lb_monitor":      lb.ResourceMonitorV2(),
			"huaweicloud_lb_pool":         lb.ResourcePoolV2(),
			"huaweicloud_lb_whitelist":    lb.ResourceWhitelistV2(),

			"huaweicloud_live_bucket_authorization":       live.ResourceBucketAuthorization(),
			"huaweicloud_live_channel":                    live.ResourceChannel(),
			"huaweicloud_live_disable_push_stream":        live.ResourceDisablePushStream(),
			"huaweicloud_live_domain":                     live.ResourceDomain(),
			"huaweicloud_live_geo_blocking":               live.ResourceGeoBlocking(),
			"huaweicloud_live_hls_configuration":          live.ResourceHlsConfiguration(),
			"huaweicloud_live_https_certificate":          live.ResourceHTTPSCertificate(),
			"huaweicloud_live_ip_acl":                     live.ResourceIpAcl(),
			"huaweicloud_live_notification_configuration": live.ResourceNotificationConfiguration(),
			"huaweicloud_live_record_callback":            live.ResourceRecordCallback(),
			"huaweicloud_live_recording":                  live.ResourceRecording(),
			"huaweicloud_live_referer_validation":         live.ResourceRefererValidation(),
			"huaweicloud_live_snapshot":                   live.ResourceSnapshot(),
			"huaweicloud_live_transcoding":                live.ResourceTranscoding(),
			"huaweicloud_live_url_authentication":         live.ResourceUrlAuthentication(),
			"huaweicloud_live_url_validation":             live.ResourceUrlValidation(),
			"huaweicloud_live_origin_pull_configuration":  live.ResourceOriginPullConfiguration(),
			"huaweicloud_live_stream_delay":               live.ResourceStreamDelay(),

			"huaweicloud_lts_aom_access":                       lts.ResourceAOMAccess(),
			"huaweicloud_lts_group":                            lts.ResourceLTSGroup(),
			"huaweicloud_lts_host_group":                       lts.ResourceHostGroup(),
			"huaweicloud_lts_host_access":                      lts.ResourceHostAccessConfig(),
			"huaweicloud_lts_log_converge":                     lts.ResourceLogConverge(),
			"huaweicloud_lts_log_converge_switch":              lts.ResourceLogConvergeSwitch(),
			"huaweicloud_lts_metric_rule":                      lts.ResourceMetricRule(),
			"huaweicloud_lts_stream":                           lts.ResourceLTSStream(),
			"huaweicloud_lts_structing_template":               lts.ResourceStructConfig(),
			"huaweicloud_lts_structuring_custom_configuration": lts.ResourceStructCustomConfig(),
			"huaweicloud_lts_transfer":                         lts.ResourceLtsTransfer(),
			"huaweicloud_lts_keywords_alarm_rule":              lts.ResourceKeywordsAlarmRule(),
			"huaweicloud_lts_sql_alarm_rule":                   lts.ResourceSQLAlarmRule(),
			"huaweicloud_lts_notification_template":            lts.ResourceNotificationTemplate(),
			"huaweicloud_lts_search_criteria":                  lts.ResourceSearchCriteria(),
			"huaweicloud_lts_waf_access":                       lts.ResourceWAFAccess(),
			"huaweicloud_lts_cross_account_access":             lts.ResourceCrossAccountAccess(),
			"huaweicloud_lts_cce_access":                       lts.ResourceCceAccessConfig(),

			"huaweicloud_mapreduce_cluster":         mrs.ResourceMRSClusterV2(),
			"huaweicloud_mapreduce_job":             mrs.ResourceMRSJobV2(),
			"huaweicloud_mapreduce_data_connection": mrs.ResourceDataConnection(),
			"huaweicloud_mapreduce_scaling_policy":  mrs.ResourceScalingPolicy(),

			"huaweicloud_meeting_admin_assignment": meeting.ResourceAdminAssignment(),
			"huaweicloud_meeting_conference":       meeting.ResourceConference(),
			"huaweicloud_meeting_user":             meeting.ResourceUser(),

			"huaweicloud_modelarts_dataset":                modelarts.ResourceDataset(),
			"huaweicloud_modelarts_dataset_version":        modelarts.ResourceDatasetVersion(),
			"huaweicloud_modelarts_devserver_action":       modelarts.ResourceDevServerAction(),
			"huaweicloud_modelarts_devserver":              modelarts.ResourceDevServer(),
			"huaweicloud_modelarts_notebook":               modelarts.ResourceNotebook(),
			"huaweicloud_modelarts_notebook_mount_storage": modelarts.ResourceNotebookMountStorage(),
			"huaweicloud_modelarts_model":                  modelarts.ResourceModelartsModel(),
			"huaweicloud_modelarts_service":                modelarts.ResourceModelartsService(),
			"huaweicloud_modelarts_workspace":              modelarts.ResourceModelartsWorkspace(),
			"huaweicloud_modelarts_authorization":          modelarts.ResourceModelArtsAuthorization(),
			"huaweicloud_modelarts_network":                modelarts.ResourceModelartsNetwork(),
			"huaweicloud_modelarts_resource_pool":          modelarts.ResourceModelartsResourcePool(),

			// DataArts Studio - Management Center
			"huaweicloud_dataarts_studio_data_connection": dataarts.ResourceDataConnection(),
			"huaweicloud_dataarts_studio_instance":        dataarts.ResourceStudioInstance(),
			// DataArts Architecture
			"huaweicloud_dataarts_architecture_directory":              dataarts.ResourceArchitectureDirectory(),
			"huaweicloud_dataarts_architecture_model":                  dataarts.ResourceArchitectureModel(),
			"huaweicloud_dataarts_architecture_subject":                dataarts.ResourceArchitectureSubject(),
			"huaweicloud_dataarts_architecture_table_model":            dataarts.ResourceArchitectureTableModel(),
			"huaweicloud_dataarts_architecture_batch_publish":          dataarts.ResourceArchitectureBatchPublish(),
			"huaweicloud_dataarts_architecture_batch_publishment":      dataarts.ResourceArchitectureBatchPublishment(),
			"huaweicloud_dataarts_architecture_batch_unpublish":        dataarts.ResourceArchitectureBatchUnpublish(),
			"huaweicloud_dataarts_architecture_business_metric":        dataarts.ResourceBusinessMetric(),
			"huaweicloud_dataarts_architecture_process":                dataarts.ResourceArchitectureProcess(),
			"huaweicloud_dataarts_architecture_code_table":             dataarts.ResourceArchitectureCodeTable(),
			"huaweicloud_dataarts_architecture_code_table_values":      dataarts.ResourceArchitectureCodeTableValues(),
			"huaweicloud_dataarts_architecture_data_standard":          dataarts.ResourceDataStandard(),
			"huaweicloud_dataarts_architecture_data_standard_template": dataarts.ResourceDataStandardTemplate(),
			"huaweicloud_dataarts_architecture_reviewer":               dataarts.ResourceDataArtsArchitectureReviewer(),
			// DataArts Factory
			"huaweicloud_dataarts_factory_resource":   dataarts.ResourceFactoryResource(),
			"huaweicloud_dataarts_factory_job_action": dataarts.ResourceFactoryJobAction(),
			"huaweicloud_dataarts_factory_job":        dataarts.ResourceFactoryJob(),
			"huaweicloud_dataarts_factory_script":     dataarts.ResourceDataArtsFactoryScript(),
			// DataArts Security
			"huaweicloud_dataarts_security_data_recognition_rule":    dataarts.ResourceSecurityRule(),
			"huaweicloud_dataarts_security_data_secrecy_level":       dataarts.ResourceSecurityDataSecrecyLevel(),
			"huaweicloud_dataarts_security_permission_set":           dataarts.ResourceSecurityPermissionSet(),
			"huaweicloud_dataarts_security_permission_set_member":    dataarts.ResourceSecurityPermissionSetMember(),
			"huaweicloud_dataarts_security_permission_set_privilege": dataarts.ResourceSecurityPermissionSetPrivilege(),
			// DataArts DataService
			"huaweicloud_dataarts_dataservice_api":             dataarts.ResourceDataServiceApi(),
			"huaweicloud_dataarts_dataservice_api_action":      dataarts.ResourceDataServiceApiAction(),
			"huaweicloud_dataarts_dataservice_api_auth":        dataarts.ResourceDataServiceApiAuth(),
			"huaweicloud_dataarts_dataservice_api_auth_action": dataarts.ResourceDataServiceApiAuthAction(),
			"huaweicloud_dataarts_dataservice_api_debug":       dataarts.ResourceDataServiceApiDebug(),
			"huaweicloud_dataarts_dataservice_api_publish":     dataarts.ResourceDataServiceApiPublish(),
			"huaweicloud_dataarts_dataservice_api_publishment": dataarts.ResourceDataServiceApiPublishment(),
			"huaweicloud_dataarts_dataservice_app":             dataarts.ResourceDataServiceApp(),
			"huaweicloud_dataarts_dataservice_catalog":         dataarts.ResourceDatatServiceCatalog(),
			"huaweicloud_dataarts_dataservice_message_approve": dataarts.ResourceDataServiceMessageApprove(),

			"huaweicloud_mpc_transcoding_template":       mpc.ResourceTranscodingTemplate(),
			"huaweicloud_mpc_transcoding_template_group": mpc.ResourceTranscodingTemplateGroup(),

			"huaweicloud_mrs_cluster": ResourceMRSClusterV1(),
			"huaweicloud_mrs_job":     ResourceMRSJobV1(),

			"huaweicloud_nat_dnat_rule": nat.ResourcePublicDnatRule(),
			"huaweicloud_nat_gateway":   nat.ResourcePublicGateway(),
			"huaweicloud_nat_snat_rule": nat.ResourcePublicSnatRule(),

			"huaweicloud_nat_private_dnat_rule":  nat.ResourcePrivateDnatRule(),
			"huaweicloud_nat_private_gateway":    nat.ResourcePrivateGateway(),
			"huaweicloud_nat_private_snat_rule":  nat.ResourcePrivateSnatRule(),
			"huaweicloud_nat_private_transit_ip": nat.ResourcePrivateTransitIp(),

			"huaweicloud_networking_secgroup":      vpc.ResourceNetworkingSecGroup(),
			"huaweicloud_networking_secgroup_rule": vpc.ResourceNetworkingSecGroupRule(),
			"huaweicloud_networking_vip":           vpc.ResourceNetworkingVip(),
			"huaweicloud_networking_vip_associate": vpc.ResourceNetworkingVIPAssociateV2(),

			"huaweicloud_obs_bucket":             obs.ResourceObsBucket(),
			"huaweicloud_obs_bucket_acl":         obs.ResourceOBSBucketAcl(),
			"huaweicloud_obs_bucket_object":      obs.ResourceObsBucketObject(),
			"huaweicloud_obs_bucket_object_acl":  obs.ResourceOBSBucketObjectAcl(),
			"huaweicloud_obs_bucket_policy":      obs.ResourceObsBucketPolicy(),
			"huaweicloud_obs_bucket_replication": obs.ResourceObsBucketReplication(),

			"huaweicloud_oms_migration_sync_task":  oms.ResourceMigrationSyncTask(),
			"huaweicloud_oms_migration_task":       oms.ResourceMigrationTask(),
			"huaweicloud_oms_migration_task_group": oms.ResourceMigrationTaskGroup(),

			"huaweicloud_ram_organization":            ram.ResourceRAMOrganization(),
			"huaweicloud_ram_resource_share":          ram.ResourceRAMShare(),
			"huaweicloud_ram_resource_share_accepter": ram.ResourceShareAccepter(),

			"huaweicloud_rds_mysql_account":                rds.ResourceMysqlAccount(),
			"huaweicloud_rds_mysql_binlog":                 rds.ResourceMysqlBinlog(),
			"huaweicloud_rds_mysql_database":               rds.ResourceMysqlDatabase(),
			"huaweicloud_rds_mysql_database_privilege":     rds.ResourceMysqlDatabasePrivilege(),
			"huaweicloud_rds_mysql_database_table_restore": rds.ResourceMysqlDatabaseTableRestore(),
			"huaweicloud_rds_mysql_proxy":                  rds.ResourceMysqlProxy(),
			"huaweicloud_rds_mysql_proxy_restart":          rds.ResourceMysqlProxyRestart(),
			"huaweicloud_rds_pg_account":                   rds.ResourcePgAccount(),
			"huaweicloud_rds_pg_account_roles":             rds.ResourcePgAccountRoles(),
			"huaweicloud_rds_pg_account_privileges":        rds.ResourcePgAccountPrivileges(),
			"huaweicloud_rds_pg_database":                  rds.ResourcePgDatabase(),
			"huaweicloud_rds_pg_database_privilege":        rds.ResourcePgDatabasePrivilege(),
			"huaweicloud_rds_sqlserver_account":            rds.ResourceSQLServerAccount(),
			"huaweicloud_rds_sqlserver_database":           rds.ResourceSQLServerDatabase(),
			"huaweicloud_rds_sqlserver_database_privilege": rds.ResourceSQLServerDatabasePrivilege(),
			"huaweicloud_rds_instance":                     rds.ResourceRdsInstance(),
			"huaweicloud_rds_instance_eip_associate":       rds.ResourceRdsInstanceEipAssociate(),
			"huaweicloud_rds_parametergroup":               rds.ResourceRdsConfiguration(),
			"huaweicloud_rds_read_replica_instance":        rds.ResourceRdsReadReplicaInstance(),
			"huaweicloud_rds_backup":                       rds.ResourceBackup(),
			"huaweicloud_rds_restore":                      rds.ResourceRdsRestore(),
			"huaweicloud_rds_cross_region_backup_strategy": rds.ResourceBackupStrategy(),
			"huaweicloud_rds_sql_audit":                    rds.ResourceSQLAudit(),
			"huaweicloud_rds_pg_plugin":                    rds.ResourceRdsPgPlugin(),
			"huaweicloud_rds_pg_plugin_update":             rds.ResourceRdsPgPluginUpdate(),
			"huaweicloud_rds_pg_hba":                       rds.ResourcePgHba(),
			"huaweicloud_rds_pg_sql_limit":                 rds.ResourcePgSqlLimit(),
			"huaweicloud_rds_pg_plugin_parameter":          rds.ResourcePgPluginParameter(),
			"huaweicloud_rds_lts_log":                      rds.ResourceRdsLtsLog(),
			"huaweicloud_rds_recycling_policy":             rds.ResourceRecyclingPolicy(),
			"huaweicloud_rds_primary_standby_switch":       rds.ResourceRdsInstanceSwitch(),
			"huaweicloud_rds_database_logs_shrinking":      rds.ResourceRdsDbLogsShrinking(),
			"huaweicloud_rds_extend_log_link":              rds.ResourceRdsExtendLogLink(),

			"huaweicloud_rgc_account": rgc.ResourceAccount(),

			"huaweicloud_rms_policy_assignment":                  rms.ResourcePolicyAssignment(),
			"huaweicloud_rms_policy_assignment_evaluate":         rms.ResourcePolicyAssignmentEvaluate(),
			"huaweicloud_rms_resource_aggregator":                rms.ResourceAggregator(),
			"huaweicloud_rms_resource_aggregation_authorization": rms.ResourceAggregationAuthorization(),
			"huaweicloud_rms_resource_recorder":                  rms.ResourceRecorder(),
			"huaweicloud_rms_advanced_query":                     rms.ResourceAdvancedQuery(),
			"huaweicloud_rms_assignment_package":                 rms.ResourceAssignmentPackage(),
			"huaweicloud_rms_organizational_assignment_package":  rms.ResourceOrgAssignmentPackage(),
			"huaweicloud_rms_organizational_policy_assignment":   rms.ResourceOrganizationalPolicyAssignment(),
			"huaweicloud_rms_remediation_configuration":          rms.ResourceRemediationConfiguration(),
			"huaweicloud_rms_remediation_exception":              rms.ResourceRemediationException(),
			"huaweicloud_rms_remediation_execution":              rms.ResourceRemediationExecution(),

			"huaweicloud_sdrs_drill":              sdrs.ResourceDrill(),
			"huaweicloud_sdrs_replication_pair":   sdrs.ResourceReplicationPair(),
			"huaweicloud_sdrs_protection_group":   sdrs.ResourceProtectionGroup(),
			"huaweicloud_sdrs_protected_instance": sdrs.ResourceProtectedInstance(),
			"huaweicloud_sdrs_replication_attach": sdrs.ResourceReplicationAttach(),

			"huaweicloud_secmaster_incident":                    secmaster.ResourceIncident(),
			"huaweicloud_secmaster_indicator":                   secmaster.ResourceIndicator(),
			"huaweicloud_secmaster_alert":                       secmaster.ResourceAlert(),
			"huaweicloud_secmaster_alert_rule":                  secmaster.ResourceAlertRule(),
			"huaweicloud_secmaster_data_object_relations":       secmaster.ResourceDataObjectRelations(),
			"huaweicloud_secmaster_playbook":                    secmaster.ResourcePlaybook(),
			"huaweicloud_secmaster_playbook_enable":             secmaster.ResourcePlaybookEnable(),
			"huaweicloud_secmaster_playbook_version":            secmaster.ResourcePlaybookVersion(),
			"huaweicloud_secmaster_playbook_version_action":     secmaster.ResourcePlaybookVersionAction(),
			"huaweicloud_secmaster_playbook_rule":               secmaster.ResourcePlaybookRule(),
			"huaweicloud_secmaster_playbook_action":             secmaster.ResourcePlaybookAction(),
			"huaweicloud_secmaster_playbook_approval":           secmaster.ResourcePlaybookApproval(),
			"huaweicloud_secmaster_playbook_instance_operation": secmaster.ResourcePlaybookInstanceOperation(),
			"huaweicloud_secmaster_alert_rule_simulation":       secmaster.ResourceAlertRuleSimulation(),

			"huaweicloud_servicestage_application":                 servicestage.ResourceApplication(),
			"huaweicloud_servicestage_component_instance":          servicestage.ResourceComponentInstance(),
			"huaweicloud_servicestage_component":                   servicestage.ResourceComponent(),
			"huaweicloud_servicestage_environment":                 servicestage.ResourceEnvironment(),
			"huaweicloud_servicestage_repo_token_authorization":    servicestage.ResourceRepoTokenAuth(),
			"huaweicloud_servicestage_repo_password_authorization": servicestage.ResourceRepoPwdAuth(),
			// v3 managements
			"huaweicloud_servicestagev3_application":           servicestage.ResourceV3Application(),
			"huaweicloud_servicestagev3_component":             servicestage.ResourceV3Component(),
			"huaweicloud_servicestagev3_environment":           servicestage.ResourceV3Environment(),
			"huaweicloud_servicestagev3_environment_associate": servicestage.ResourceV3EnvironmentAssociate(),

			"huaweicloud_sfs_turbo":            sfsturbo.ResourceSFSTurbo(),
			"huaweicloud_sfs_turbo_dir":        sfsturbo.ResourceSfsTurboDir(),
			"huaweicloud_sfs_turbo_dir_quota":  sfsturbo.ResourceSfsTurboDirQuota(),
			"huaweicloud_sfs_turbo_data_task":  sfsturbo.ResourceDataTask(),
			"huaweicloud_sfs_turbo_du_task":    sfsturbo.ResourceDuTask(),
			"huaweicloud_sfs_turbo_obs_target": sfsturbo.ResourceOBSTarget(),
			"huaweicloud_sfs_turbo_perm_rule":  sfsturbo.ResourceSFSTurboPermRule(),

			"huaweicloud_smn_topic":                      smn.ResourceTopic(),
			"huaweicloud_smn_subscription":               smn.ResourceSubscription(),
			"huaweicloud_smn_message_template":           smn.ResourceSmnMessageTemplate(),
			"huaweicloud_smn_logtank":                    smn.ResourceSmnLogtank(),
			"huaweicloud_smn_message_detection":          smn.ResourceMessageDetection(),
			"huaweicloud_smn_message_publish":            smn.ResourceMessagePublish(),
			"huaweicloud_smn_subscription_filter_policy": smn.ResourceSubscriptionFilterPolicy(),

			"huaweicloud_sms_server_template": sms.ResourceServerTemplate(),
			"huaweicloud_sms_task":            sms.ResourceMigrateTask(),

			"huaweicloud_swr_organization":             swr.ResourceSWROrganization(),
			"huaweicloud_swr_organization_permissions": swr.ResourceSWROrganizationPermissions(),
			"huaweicloud_swr_repository":               swr.ResourceSWRRepository(),
			"huaweicloud_swr_repository_sharing":       swr.ResourceSWRRepositorySharing(),
			"huaweicloud_swr_image_permissions":        swr.ResourceSwrImagePermissions(),
			"huaweicloud_swr_image_trigger":            swr.ResourceSwrImageTrigger(),
			"huaweicloud_swr_image_retention_policy":   swr.ResourceSwrImageRetentionPolicy(),
			"huaweicloud_swr_image_auto_sync":          swr.ResourceSwrImageAutoSync(),

			"huaweicloud_tms_resource_tags": tms.ResourceResourceTags(),
			"huaweicloud_tms_tags":          tms.ResourceTmsTag(),

			"huaweicloud_ucs_fleet":   ucs.ResourceFleet(),
			"huaweicloud_ucs_cluster": ucs.ResourceCluster(),
			"huaweicloud_ucs_policy":  ucs.ResourcePolicy(),

			"huaweicloud_vod_media_asset":                vod.ResourceMediaAsset(),
			"huaweicloud_vod_media_category":             vod.ResourceMediaCategory(),
			"huaweicloud_vod_transcoding_template_group": vod.ResourceTranscodingTemplateGroup(),
			"huaweicloud_vod_watermark_template":         vod.ResourceWatermarkTemplate(),

			"huaweicloud_vpc_bandwidth":           eip.ResourceVpcBandWidthV2(),
			"huaweicloud_vpc_bandwidth_associate": eip.ResourceBandWidthAssociate(),
			"huaweicloud_vpc_eip":                 eip.ResourceVpcEIPV1(),
			"huaweicloud_vpc_eip_associate":       eip.ResourceEIPAssociate(),
			"huaweicloud_vpc_eipv3_associate":     eip.ResourceEipv3Associate(),
			"huaweicloud_vpc_internet_gateway":    eip.ResourceVPCInternetGateway(),

			"huaweicloud_global_internet_bandwidth": eip.ResourceGlobalInternetBandwidth(),
			"huaweicloud_global_eip":                eip.ResourceGlobalEIP(),
			"huaweicloud_global_eip_associate":      eip.ResourceGlobalEIPAssociate(),

			"huaweicloud_vpc_peering_connection":          vpc.ResourceVpcPeeringConnectionV2(),
			"huaweicloud_vpc_peering_connection_accepter": vpc.ResourceVpcPeeringConnectionAccepterV2(),
			"huaweicloud_vpc_route_table":                 vpc.ResourceVPCRouteTable(),
			"huaweicloud_vpc_route":                       vpc.ResourceVPCRouteTableRoute(),
			"huaweicloud_vpc":                             vpc.ResourceVirtualPrivateCloudV1(),
			"huaweicloud_vpc_subnet":                      vpc.ResourceVpcSubnetV1(),
			"huaweicloud_vpc_subnet_private_ip":           vpc.ResourceSubnetPrivateIP(),
			"huaweicloud_vpc_address_group":               vpc.ResourceVpcAddressGroup(),
			"huaweicloud_vpc_flow_log":                    vpc.ResourceVpcFlowLog(),
			"huaweicloud_vpc_network_interface":           vpc.ResourceNetworkInterface(),
			"huaweicloud_vpc_network_acl":                 vpc.ResourceNetworkAcl(),
			"huaweicloud_vpc_sub_network_interface":       vpc.ResourceSubNetworkInterface(),
			"huaweicloud_vpc_traffic_mirror_filter":       vpc.ResourceTrafficMirrorFilter(),
			"huaweicloud_vpc_traffic_mirror_filter_rule":  vpc.ResourceTrafficMirrorFilterRule(),
			"huaweicloud_vpc_traffic_mirror_session":      vpc.ResourceTrafficMirrorSession(),

			"huaweicloud_vpcep_approval": vpcep.ResourceVPCEndpointApproval(),
			"huaweicloud_vpcep_endpoint": vpcep.ResourceVPCEndpoint(),
			"huaweicloud_vpcep_service":  vpcep.ResourceVPCEndpointService(),

			"huaweicloud_vpn_access_policy":           vpn.ResourceAccessPolicy(),
			"huaweicloud_vpn_gateway":                 vpn.ResourceGateway(),
			"huaweicloud_vpn_customer_gateway":        vpn.ResourceCustomerGateway(),
			"huaweicloud_vpn_connection":              vpn.ResourceConnection(),
			"huaweicloud_vpn_connection_health_check": vpn.ResourceConnectionHealthCheck(),
			"huaweicloud_vpn_user":                    vpn.ResourceUser(),
			"huaweicloud_vpn_user_group":              vpn.ResourceUserGroup(),
			"huaweicloud_vpn_client_ca_certificate":   vpn.ResourceClientCACertificate(),

			"huaweicloud_waf_address_group":                       waf.ResourceWafAddressGroup(),
			"huaweicloud_waf_certificate":                         waf.ResourceWafCertificate(),
			"huaweicloud_waf_cloud_instance":                      waf.ResourceCloudInstance(),
			"huaweicloud_waf_domain":                              waf.ResourceWafDomain(),
			"huaweicloud_waf_policy":                              waf.ResourceWafPolicy(),
			"huaweicloud_waf_rule_anti_crawler":                   waf.ResourceRuleAntiCrawler(),
			"huaweicloud_waf_rule_blacklist":                      waf.ResourceWafRuleBlackList(),
			"huaweicloud_waf_rule_cc_protection":                  waf.ResourceRuleCCProtection(),
			"huaweicloud_waf_rule_data_masking":                   waf.ResourceWafRuleDataMasking(),
			"huaweicloud_waf_rule_global_protection_whitelist":    waf.ResourceRuleGlobalProtectionWhitelist(),
			"huaweicloud_waf_rule_known_attack_source":            waf.ResourceRuleKnownAttack(),
			"huaweicloud_waf_rule_precise_protection":             waf.ResourceRulePreciseProtection(),
			"huaweicloud_waf_rule_web_tamper_protection":          waf.ResourceWafRuleWebTamperProtection(),
			"huaweicloud_waf_rule_geolocation_access_control":     waf.ResourceRuleGeolocation(),
			"huaweicloud_waf_rule_information_leakage_prevention": waf.ResourceRuleLeakagePrevention(),
			"huaweicloud_waf_dedicated_instance":                  waf.ResourceWafDedicatedInstance(),
			"huaweicloud_waf_dedicated_domain":                    waf.ResourceWafDedicatedDomain(),
			"huaweicloud_waf_instance_group":                      waf.ResourceWafInstanceGroup(),
			"huaweicloud_waf_instance_group_associate":            waf.ResourceWafInstGroupAssociate(),
			"huaweicloud_waf_reference_table":                     waf.ResourceWafReferenceTable(),

			"huaweicloud_workspace_app_group_authorization": workspace.ResourceAppGroupAuthorization(),
			"huaweicloud_workspace_app_group":               workspace.ResourceWorkspaceAppGroup(),
			"huaweicloud_workspace_app_image_server":        workspace.ResourceAppImageServer(),
			"huaweicloud_workspace_app_image":               workspace.ResourceAppImage(),
			"huaweicloud_workspace_app_nas_storage":         workspace.ResourceAppNasStorage(),
			"huaweicloud_workspace_app_personal_folders":    workspace.ResourceAppPersonalFolders(),
			"huaweicloud_workspace_app_policy_group":        workspace.ResourceAppPolicyGroup(),
			"huaweicloud_workspace_app_publishment":         workspace.ResourceAppPublishment(),
			"huaweicloud_workspace_app_server_group":        workspace.ResourceAppServerGroup(),
			"huaweicloud_workspace_app_server":              workspace.ResourceAppServer(),
			"huaweicloud_workspace_app_shared_folder":       workspace.ResourceAppSharedFolder(),
			"huaweicloud_workspace_app_storage_policy":      workspace.ResourceAppStoragePolicy(),
			"huaweicloud_workspace_app_warehouse_app":       workspace.ResourceWarehouseApplication(),
			"huaweicloud_workspace_user_group":              workspace.ResourceUserGroup(),
			"huaweicloud_workspace_access_policy":           workspace.ResourceAccessPolicy(),
			"huaweicloud_workspace_desktop_name_rule":       workspace.ResourceDesktopNameRule(),
			"huaweicloud_workspace_desktop":                 workspace.ResourceDesktop(),
			"huaweicloud_workspace_policy_group":            workspace.ResourcePolicyGroup(),
			"huaweicloud_workspace_service":                 workspace.ResourceService(),
			"huaweicloud_workspace_terminal_binding":        workspace.ResourceTerminalBinding(),
			"huaweicloud_workspace_user":                    workspace.ResourceUser(),
			"huaweicloud_workspace_eip_associate":           workspace.ResourceEipAssociate(),

			"huaweicloud_cpts_project": cpts.ResourceProject(),
			"huaweicloud_cpts_task":    cpts.ResourceTask(),

			// CodeArts
			"huaweicloud_codearts_project":    codearts.ResourceProject(),
			"huaweicloud_codearts_repository": codearts.ResourceRepository(),

			"huaweicloud_codearts_deploy_application":            codeartsdeploy.ResourceDeployApplication(),
			"huaweicloud_codearts_deploy_application_copy":       codeartsdeploy.ResourceDeployApplicationCopy(),
			"huaweicloud_codearts_deploy_application_permission": codeartsdeploy.ResourceDeployApplicationPermission(),
			"huaweicloud_codearts_deploy_application_deploy":     codeartsdeploy.ResourceDeployApplicationDeploy(),
			"huaweicloud_codearts_deploy_application_group":      codeartsdeploy.ResourceDeployApplicationGroup(),
			"huaweicloud_codearts_deploy_application_group_move": codeartsdeploy.ResourceDeployApplicationGroupMove(),
			"huaweicloud_codearts_deploy_environment":            codeartsdeploy.ResourceDeployEnvironment(),
			"huaweicloud_codearts_deploy_environment_permission": codeartsdeploy.ResourceDeployEnvironmentPermission(),
			"huaweicloud_codearts_deploy_group":                  codeartsdeploy.ResourceDeployGroup(),
			"huaweicloud_codearts_deploy_group_permission":       codeartsdeploy.ResourceDeployGroupPermission(),
			"huaweicloud_codearts_deploy_host":                   codeartsdeploy.ResourceDeployHost(),
			"huaweicloud_codearts_deploy_hosts_copy":             codeartsdeploy.ResourceCodeArtsDeployHostsCopy(),

			"huaweicloud_codearts_inspector_website":      codeartsinspector.ResourceInspectorWebsite(),
			"huaweicloud_codearts_inspector_website_scan": codeartsinspector.ResourceInspectorWebsiteScan(),
			"huaweicloud_codearts_inspector_host_group":   codeartsinspector.ResourceInspectorHostGroup(),

			"huaweicloud_dsc_instance":           dsc.ResourceDscInstance(),
			"huaweicloud_dsc_asset_obs":          dsc.ResourceAssetObs(),
			"huaweicloud_dsc_alarm_notification": dsc.ResourceAlarmNotification(),

			// internal only
			"huaweicloud_apm_aksk":                apm.ResourceApmAkSk(),
			"huaweicloud_aom_alarm_policy":        aom.ResourceAlarmPolicy(),
			"huaweicloud_aom_prometheus_instance": aom.ResourcePrometheusInstance(),

			"huaweicloud_aom_application":                 cmdb.ResourceAomApplication(),
			"huaweicloud_aom_component":                   cmdb.ResourceAomComponent(),
			"huaweicloud_aom_cmdb_resource_relationships": cmdb.ResourceCiRelationships(),
			"huaweicloud_aom_environment":                 cmdb.ResourceAomEnvironment(),

			"huaweicloud_lts_access_rule":     lts.ResourceAomMappingRule(),
			"huaweicloud_lts_dashboard":       lts.ResourceLtsDashboard(),
			"huaweicloud_elb_log":             lts.ResourceLtsElb(),
			"huaweicloud_lts_struct_template": lts.ResourceLtsStruct(),

			// Legacy
			"huaweicloud_networking_eip_associate": eip.ResourceEIPAssociate(),
			"huaweicloud_dds_instance_recovery":    dds.ResourceDDSInstanceRestore(),

			"huaweicloud_projectman_project": codearts.ResourceProject(),
			"huaweicloud_codehub_repository": codearts.ResourceRepository(),

			"huaweicloud_compute_instance_v2":             ecs.ResourceComputeInstance(),
			"huaweicloud_compute_interface_attach_v2":     ecs.ResourceComputeInterfaceAttach(),
			"huaweicloud_compute_keypair_v2":              ResourceComputeKeypairV2(),
			"huaweicloud_compute_servergroup_v2":          ecs.ResourceComputeServerGroup(),
			"huaweicloud_compute_volume_attach_v2":        ecs.ResourceComputeVolumeAttach(),
			"huaweicloud_compute_floatingip_associate_v2": ecs.ResourceComputeEIPAssociate(),

			"huaweicloud_dns_ptrrecord_v2": dns.ResourceDNSPtrRecord(),
			"huaweicloud_dns_recordset_v2": dns.ResourceDNSRecordSetV2(),
			"huaweicloud_dns_zone_v2":      dns.ResourceDNSZone(),

			"huaweicloud_dcs_instance_v1":      dcs.ResourceDcsInstance(),
			"huaweicloud_dds_instance_v3":      dds.ResourceDdsInstanceV3(),
			"huaweicloud_dcs_isntance_restore": dcs.ResourceDcsRestore(),

			"huaweicloud_kms_key_v1": dew.ResourceKmsKey(),

			"huaweicloud_lb_certificate_v2":  lb.ResourceCertificateV2(),
			"huaweicloud_lb_loadbalancer_v2": lb.ResourceLoadBalancer(),
			"huaweicloud_lb_listener_v2":     lb.ResourceListener(),
			"huaweicloud_lb_pool_v2":         lb.ResourcePoolV2(),
			"huaweicloud_lb_member_v2":       lb.ResourceMemberV2(),
			"huaweicloud_lb_monitor_v2":      lb.ResourceMonitorV2(),
			"huaweicloud_lb_l7policy_v2":     lb.ResourceL7PolicyV2(),
			"huaweicloud_lb_l7rule_v2":       lb.ResourceL7RuleV2(),
			"huaweicloud_lb_whitelist_v2":    lb.ResourceWhitelistV2(),

			"huaweicloud_mrs_cluster_v1": ResourceMRSClusterV1(),
			"huaweicloud_mrs_job_v1":     ResourceMRSJobV1(),

			"huaweicloud_networking_secgroup_v2":      vpc.ResourceNetworkingSecGroup(),
			"huaweicloud_networking_secgroup_rule_v2": vpc.ResourceNetworkingSecGroupRule(),

			"huaweicloud_smn_topic_v2":        smn.ResourceTopic(),
			"huaweicloud_smn_subscription_v2": smn.ResourceSubscription(),

			"huaweicloud_rds_account":            rds.ResourceMysqlAccount(),
			"huaweicloud_rds_database":           rds.ResourceMysqlDatabase(),
			"huaweicloud_rds_database_privilege": rds.ResourceMysqlDatabasePrivilege(),
			"huaweicloud_rds_instance_v3":        rds.ResourceRdsInstance(),
			"huaweicloud_rds_parametergroup_v3":  rds.ResourceRdsConfiguration(),

			"huaweicloud_rf_stack": rfs.ResourceStack(),

			"huaweicloud_nat_dnat_rule_v2": nat.ResourcePublicDnatRule(),
			"huaweicloud_nat_gateway_v2":   nat.ResourcePublicGateway(),
			"huaweicloud_nat_snat_rule_v2": nat.ResourcePublicSnatRule(),

			"huaweicloud_iam_agency":    iam.ResourceIAMAgencyV3(),
			"huaweicloud_iam_agency_v3": iam.ResourceIAMAgencyV3(),

			"huaweicloud_vpc_bandwidth_v2":                   eip.ResourceVpcBandWidthV2(),
			"huaweicloud_vpc_eip_v1":                         eip.ResourceVpcEIPV1(),
			"huaweicloud_vpc_peering_connection_v2":          vpc.ResourceVpcPeeringConnectionV2(),
			"huaweicloud_vpc_peering_connection_accepter_v2": vpc.ResourceVpcPeeringConnectionAccepterV2(),
			"huaweicloud_vpc_v1":                             vpc.ResourceVirtualPrivateCloudV1(),
			"huaweicloud_vpc_subnet_v1":                      vpc.ResourceVpcSubnetV1(),

			"huaweicloud_cce_cluster_v3": cce.ResourceCCEClusterV3(),
			"huaweicloud_cce_node_v3":    cce.ResourceNode(),

			"huaweicloud_as_configuration_v1": as.ResourceASConfiguration(),
			"huaweicloud_as_group_v1":         as.ResourceASGroup(),
			"huaweicloud_as_policy_v1":        as.ResourceASPolicy(),

			"huaweicloud_identity_project_v3":          iam.ResourceIdentityProject(),
			"huaweicloud_identity_role_assignment_v3":  iam.ResourceIdentityGroupRoleAssignment(),
			"huaweicloud_identity_user_v3":             iam.ResourceIdentityUser(),
			"huaweicloud_identity_group_v3":            iam.ResourceIdentityGroup(),
			"huaweicloud_identity_group_membership_v3": iam.ResourceIdentityGroupMembership(),
			"huaweicloud_identity_provider_conversion": iam.ResourceIAMProviderConversion(),

			"huaweicloud_cdm_cluster_v1": cdm.ResourceCdmCluster(),
			"huaweicloud_css_cluster_v1": css.ResourceCssCluster(),
			"huaweicloud_dis_stream_v2":  dis.ResourceDisStream(),

			"huaweicloud_organizations_organization":            organizations.ResourceOrganization(),
			"huaweicloud_organizations_organizational_unit":     organizations.ResourceOrganizationalUnit(),
			"huaweicloud_organizations_account":                 organizations.ResourceAccount(),
			"huaweicloud_organizations_account_associate":       organizations.ResourceAccountAssociate(),
			"huaweicloud_organizations_account_invite":          organizations.ResourceAccountInvite(),
			"huaweicloud_organizations_account_invite_accepter": organizations.ResourceAccountInviteAccepter(),
			"huaweicloud_organizations_account_invite_decliner": organizations.ResourceAccountInviteDecliner(),
			"huaweicloud_organizations_trusted_service":         organizations.ResourceTrustedService(),
			"huaweicloud_organizations_policy":                  organizations.ResourcePolicy(),
			"huaweicloud_organizations_policy_attach":           organizations.ResourcePolicyAttach(),
			"huaweicloud_organizations_delegated_administrator": organizations.ResourceDelegatedAdministrator(),

			"huaweicloud_dli_queue_v1":                dli.ResourceDliQueue(),
			"huaweicloud_networking_vip_v2":           vpc.ResourceNetworkingVip(),
			"huaweicloud_networking_vip_associate_v2": vpc.ResourceNetworkingVIPAssociateV2(),
			"huaweicloud_fgs_function_v2":             fgs.ResourceFgsFunctionV2(),
			"huaweicloud_cdn_domain_v1":               cdn.ResourceCdnDomain(),
			"huaweicloud_scm_certificate":             ccm.ResourceCertificateImport(),

			// Deprecated
			"huaweicloud_apig_vpc_channel":               deprecated.ResourceApigVpcChannelV2(),
			"huaweicloud_blockstorage_volume_v2":         deprecated.ResourceBlockStorageVolumeV2(),
			"huaweicloud_cae_component_deployment":       cae.ResourceComponentAction(),
			"huaweicloud_cfw_protection_rule":            deprecated.ResourceProtectionRule(),
			"huaweicloud_csbs_backup":                    deprecated.ResourceCSBSBackupV1(),
			"huaweicloud_csbs_backup_policy":             deprecated.ResourceCSBSBackupPolicyV1(),
			"huaweicloud_csbs_backup_policy_v1":          deprecated.ResourceCSBSBackupPolicyV1(),
			"huaweicloud_csbs_backup_v1":                 deprecated.ResourceCSBSBackupV1(),
			"huaweicloud_fgs_trigger":                    deprecated.ResourceFunctionGraphTrigger(),
			"huaweicloud_network_acl":                    deprecated.ResourceNetworkACL(),
			"huaweicloud_network_acl_rule":               deprecated.ResourceNetworkACLRule(),
			"huaweicloud_networking_network_v2":          deprecated.ResourceNetworkingNetworkV2(),
			"huaweicloud_networking_subnet_v2":           deprecated.ResourceNetworkingSubnetV2(),
			"huaweicloud_networking_floatingip_v2":       deprecated.ResourceNetworkingFloatingIPV2(),
			"huaweicloud_networking_router_v2":           deprecated.ResourceNetworkingRouterV2(),
			"huaweicloud_networking_router_interface_v2": deprecated.ResourceNetworkingRouterInterfaceV2(),
			"huaweicloud_networking_router_route_v2":     deprecated.ResourceNetworkingRouterRouteV2(),
			"huaweicloud_networking_port":                deprecated.ResourceNetworkingPortV2(),
			"huaweicloud_networking_port_v2":             deprecated.ResourceNetworkingPortV2(),
			"huaweicloud_vpc_route_v2":                   deprecated.ResourceVPCRouteV2(),
			"huaweicloud_ecs_instance_v1":                deprecated.ResourceEcsInstanceV1(),
			"huaweicloud_compute_secgroup_v2":            deprecated.ResourceComputeSecGroupV2(),
			"huaweicloud_compute_floatingip_v2":          deprecated.ResourceComputeFloatingIPV2(),
			"huaweicloud_oms_task":                       deprecated.ResourceMaasTaskV1(),

			"huaweicloud_fw_firewall_group_v2": deprecated.ResourceFWFirewallGroupV2(),
			"huaweicloud_fw_policy_v2":         deprecated.ResourceFWPolicyV2(),
			"huaweicloud_fw_rule_v2":           deprecated.ResourceFWRuleV2(),

			"huaweicloud_images_image_v2": deprecated.ResourceImagesImageV2(),

			"huaweicloud_dms_instance":    deprecated.ResourceDmsInstancesV1(),
			"huaweicloud_dms_instance_v1": deprecated.ResourceDmsInstancesV1(),
			"huaweicloud_dms_group":       deprecated.ResourceDmsGroups(),
			"huaweicloud_dms_group_v1":    deprecated.ResourceDmsGroups(),
			"huaweicloud_dms_queue":       deprecated.ResourceDmsQueues(),
			"huaweicloud_dms_queue_v1":    deprecated.ResourceDmsQueues(),

			"huaweicloud_dli_template_sql":   dli.ResourceSQLTemplate(),
			"huaweicloud_dli_template_flink": dli.ResourceFlinkTemplate(),
			"huaweicloud_dli_template_spark": dli.ResourceSparkTemplate(),

			"huaweicloud_cs_cluster":            deprecated.ResourceCsClusterV1(),
			"huaweicloud_cs_cluster_v1":         deprecated.ResourceCsClusterV1(),
			"huaweicloud_cs_route":              deprecated.ResourceCsRouteV1(),
			"huaweicloud_cs_route_v1":           deprecated.ResourceCsRouteV1(),
			"huaweicloud_cs_peering_connect":    deprecated.ResourceCsPeeringConnectV1(),
			"huaweicloud_cs_peering_connect_v1": deprecated.ResourceCsPeeringConnectV1(),

			"huaweicloud_lts_structuring_configuration": lts.ResourceStructConfig(),

			"huaweicloud_sfs_access_rule":    deprecated.ResourceSFSAccessRuleV2(),
			"huaweicloud_sfs_file_system":    deprecated.ResourceSFSFileSystemV2(),
			"huaweicloud_sfs_access_rule_v2": deprecated.ResourceSFSAccessRuleV2(),
			"huaweicloud_sfs_file_system_v2": deprecated.ResourceSFSFileSystemV2(),

			"huaweicloud_vbs_backup":           deprecated.ResourceVBSBackupV2(),
			"huaweicloud_vbs_backup_policy":    deprecated.ResourceVBSBackupPolicyV2(),
			"huaweicloud_vbs_backup_policy_v2": deprecated.ResourceVBSBackupPolicyV2(),
			"huaweicloud_vbs_backup_v2":        deprecated.ResourceVBSBackupV2(),

			"huaweicloud_vpnaas_ipsec_policy_v2":    deprecated.ResourceVpnIPSecPolicyV2(),
			"huaweicloud_vpnaas_service_v2":         deprecated.ResourceVpnServiceV2(),
			"huaweicloud_vpnaas_ike_policy_v2":      deprecated.ResourceVpnIKEPolicyV2(),
			"huaweicloud_vpnaas_endpoint_group_v2":  deprecated.ResourceVpnEndpointGroupV2(),
			"huaweicloud_vpnaas_site_connection_v2": deprecated.ResourceVpnSiteConnectionV2(),

			"huaweicloud_vpnaas_endpoint_group":  deprecated.ResourceVpnEndpointGroupV2(),
			"huaweicloud_vpnaas_ike_policy":      deprecated.ResourceVpnIKEPolicyV2(),
			"huaweicloud_vpnaas_ipsec_policy":    deprecated.ResourceVpnIPSecPolicyV2(),
			"huaweicloud_vpnaas_service":         deprecated.ResourceVpnServiceV2(),
			"huaweicloud_vpnaas_site_connection": deprecated.ResourceVpnSiteConnectionV2(),

			"huaweicloud_iotda_batchtask_file": deprecated.ResourceBatchTaskFile(),

			"huaweicloud_images_image": deprecated.ResourceImsImage(),
		},
	}

	provider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11 cc
			terraformVersion = "0.11+compatible"
		}

		return configureProvider(ctx, d, terraformVersion)
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

		"assume_role_agency_name": "The name of agency for assume role.",

		"assume_role_domain_name": "The name of domain for assume role.",

		"assume_role_domain_id": "The id of domain for v5 assume role.",

		"cloud": "The endpoint of cloud provider, defaults to myhuaweicloud.com",

		"endpoints": "The custom endpoints used to override the default endpoint URL.",

		"regional": "Whether the service endpoints are regional",

		"shared_config_file": "The path to the shared config file. If not set, the default is ~/.hcloud/config.json.",

		"profile": "The profile name as set in the shared config file.",

		"max_retries": "How many times HTTP connection should be retried until giving up.",

		"enterprise_project_id": "enterprise project id",

		"enable_force_new": "Whether to enable ForceNew",

		"signing_algorithm": "The signing algorithm for authentication",
	}
}

func configureProvider(_ context.Context, d *schema.ResourceData, terraformVersion string) (interface{},
	diag.Diagnostics) {
	var tenantName, tenantID, delegatedProject, identityEndpoint string

	conf := config.Config{
		AccessKey:           d.Get("access_key").(string),
		SecretKey:           d.Get("secret_key").(string),
		CACertFile:          d.Get("cacert_file").(string),
		ClientCertFile:      d.Get("cert").(string),
		ClientKeyFile:       d.Get("key").(string),
		DomainID:            d.Get("domain_id").(string),
		DomainName:          d.Get("domain_name").(string),
		Insecure:            d.Get("insecure").(bool),
		Password:            d.Get("password").(string),
		Token:               d.Get("token").(string),
		SecurityToken:       d.Get("security_token").(string),
		Username:            d.Get("user_name").(string),
		UserID:              d.Get("user_id").(string),
		AgencyName:          d.Get("agency_name").(string),
		AgencyDomainName:    d.Get("agency_domain_name").(string),
		MaxRetries:          d.Get("max_retries").(int),
		EnterpriseProjectID: d.Get("enterprise_project_id").(string),
		SharedConfigFile:    d.Get("shared_config_file").(string),
		Profile:             d.Get("profile").(string),
		TerraformVersion:    terraformVersion,
		RegionProjectIDMap:  make(map[string]string),
		RPLock:              new(sync.Mutex),
		SecurityKeyLock:     new(sync.Mutex),
		EnableForceNew:      d.Get("enable_force_new").(bool),
		SigningAlgorithm:    d.Get("signing_algorithm").(string),
	}

	// get assume role
	assumeRoleList := d.Get("assume_role").([]interface{})
	if len(assumeRoleList) == 0 {
		// without assume_role block in provider
		delegatedAgencyName := os.Getenv("HW_ASSUME_ROLE_AGENCY_NAME")
		delegatedDomianName := os.Getenv("HW_ASSUME_ROLE_DOMAIN_NAME")
		delegatedDomianID := os.Getenv("HW_ASSUME_ROLE_DOMAIN_ID")
		if delegatedAgencyName != "" && delegatedDomianName != "" {
			conf.AssumeRoleAgency = delegatedAgencyName
			conf.AssumeRoleDomain = delegatedDomianName
			conf.AssumeRoleDomainID = delegatedDomianID
		}
	} else {
		assumeRole := assumeRoleList[0].(map[string]interface{})
		conf.AssumeRoleAgency = assumeRole["agency_name"].(string)
		conf.AssumeRoleDomain = assumeRole["domain_name"].(string)
		conf.AssumeRoleDomainID = assumeRole["domain_id"].(string)
	}

	conf.Region = d.Get("region").(string)

	if conf.SharedConfigFile != "" || conf.Profile != "" {
		err := readConfig(&conf)
		if err != nil {
			return nil, diag.FromErr(err)
		}
	}

	if conf.Region == "" {
		return nil, diag.Errorf("region should be provided")
	}

	cloud := getCloudDomain(d.Get("cloud").(string), conf.Region)
	conf.Cloud = cloud

	isRegional := d.Get("regional").(bool)
	if strings.HasPrefix(conf.Region, prefixEuropeRegion) {
		// the default format of endpoints in Europe site is xxx.{{region}}.{{cloud}}
		isRegional = true
	}
	conf.RegionClient = isRegional

	// if can't read from shared config, keep the original way
	if conf.TenantID == "" {
		// project_id is prior to tenant_id
		if v, ok := d.GetOk("project_id"); ok && v.(string) != "" {
			tenantID = v.(string)
		} else {
			tenantID = d.Get("tenant_id").(string)
		}
		conf.TenantID = tenantID

		// project_name is prior to tenant_name
		// if neither of them was set, use region as the default project
		if v, ok := d.GetOk("project_name"); ok && v.(string) != "" {
			tenantName = v.(string)
		} else if v, ok := d.GetOk("tenant_name"); ok && v.(string) != "" {
			tenantName = v.(string)
		} else {
			tenantName = conf.Region
		}
		conf.TenantName = tenantName
	}

	// Use region as delegated_project if it's not set
	if v, ok := d.GetOk("delegated_project"); ok && v.(string) != "" {
		delegatedProject = v.(string)
	} else {
		delegatedProject = conf.Region
	}
	conf.DelegatedProject = delegatedProject

	// use auth_url as identityEndpoint if specified
	if v, ok := d.GetOk("auth_url"); ok {
		identityEndpoint = v.(string)
	} else {
		// use cloud as basis for identityEndpoint
		identityEndpoint = fmt.Sprintf("https://iam.%s.%s/v3", conf.Region, cloud)
	}
	conf.IdentityEndpoint = identityEndpoint

	// get custom endpoints
	endpoints, err := flattenProviderEndpoints(d)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	conf.Endpoints = endpoints

	if err := conf.LoadAndValidate(); err != nil {
		return nil, diag.FromErr(err)
	}

	if conf.Cloud == defaultCloud {
		if os.Getenv("SKIP_CHECK_WEBSITE_TYPE") != "true" {
			if err := conf.SetWebsiteType(); err != nil {
				log.Printf("[WARN] failed to get the website type: %s", err)
			}

			if conf.GetWebsiteType() == config.InternationalSite {
				// refer to https://developer.huaweicloud.com/intl/en-us/endpoint
				bssIntlEndpoint := fmt.Sprintf("https://bss-intl.%s/", conf.Cloud)
				tmsIntlEndpoint := fmt.Sprintf("https://tms.ap-southeast-1.%s/", conf.Cloud)

				conf.SetServiceEndpoint("bss", bssIntlEndpoint)
				conf.SetServiceEndpoint("tms", tmsIntlEndpoint)
			}
		} else {
			log.Printf("[WARN] check website type skipped")
		}
	}

	if conf.Cloud == defaultEuropeCloud {
		cdnEndpoint := fmt.Sprintf("https://cdn.%s/", conf.Cloud)
		conf.SetServiceEndpoint("cdn", cdnEndpoint)
	}

	return &conf, config.CheckUpgrade(Version)
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

	// unify the endpoint which has multiple versions
	for key := range endpoints {
		ep, ok := epMap[key]
		if !ok {
			continue
		}

		multiKeys := config.GetServiceDerivedCatalogKeys(key)
		for _, k := range multiKeys {
			epMap[k] = ep
		}
	}

	log.Printf("[DEBUG] customer endpoints: %+v", epMap)
	return epMap, nil
}

func getCloudDomain(cloud, region string) string {
	// first, use the specified value
	if cloud != "" {
		return cloud
	}

	// then check whether the region(eu-west-1xx) is located in Europe
	if strings.HasPrefix(region, prefixEuropeRegion) {
		return defaultEuropeCloud
	}
	return defaultCloud
}

func readConfig(c *config.Config) error {
	if c.SharedConfigFile == "" {
		c.SharedConfigFile = fmt.Sprintf("%s/.hcloud/config.json", os.Getenv("HOME"))
		if runtime.GOOS == "windows" {
			c.SharedConfigFile = fmt.Sprintf("%s/.hcloud/config.json", os.Getenv("USERPROFILE"))
		}
	}

	profilePath, err := homedir.Expand(c.SharedConfigFile)
	if err != nil {
		return err
	}

	current := c.Profile
	var providerConfig config.Profile
	_, err = os.Stat(profilePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("The specified shared config file %s does not exist", profilePath)
	}

	data, err := os.ReadFile(profilePath)
	if err != nil {
		return fmt.Errorf("Err reading from shared config file: %s", err)
	}
	sharedConfig := config.SharedConfig{}
	err = json.Unmarshal(data, &sharedConfig)
	if err != nil {
		return err
	}

	// fetch current from shared config if not specified with provider
	if current == "" {
		current = sharedConfig.Current
	}

	// fetch the current profile config
	for _, v := range sharedConfig.Profiles {
		if current == v.Name {
			providerConfig = v
			break
		}
	}
	if (providerConfig == config.Profile{}) {
		return fmt.Errorf("Error finding profile %s from shared config file", current)
	}

	if providerConfig.Mode == "SSO" {
		ssoAuth := providerConfig.SsoAuth
		if (ssoAuth == config.SsoAuth{}) {
			return fmt.Errorf("Error finding ssoAuth config when auth mode is SSO")
		}
		stsToken := ssoAuth.StsToken
		if (stsToken == config.StsToken{}) {
			return fmt.Errorf("Error finding ssoAuth.stsToken config when auth mode is SSO")
		}
		c.AccessKey = stsToken.AccessKeyId
		c.SecretKey = stsToken.SecretAccessKey
		c.SecurityToken = stsToken.SecurityToken
	} else {
		c.AccessKey = providerConfig.AccessKeyId
		c.SecretKey = providerConfig.SecretAccessKey
		c.SecurityToken = providerConfig.SecurityToken
	}

	// non required fields
	if providerConfig.Region != "" {
		c.Region = providerConfig.Region
	}
	if providerConfig.DomainId != "" {
		c.DomainID = providerConfig.DomainId
	}
	if providerConfig.ProjectId != "" {
		c.TenantID = providerConfig.ProjectId
	}
	// assume role
	if providerConfig.AgencyName != "" {
		c.AssumeRoleAgency = providerConfig.AgencyName
	}
	if providerConfig.AgencyDomainName != "" {
		c.AssumeRoleDomain = providerConfig.AgencyDomainName
	}

	return nil
}
