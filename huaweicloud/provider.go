package huaweicloud

import (
	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// This is a global MutexKV for use within this plugin.
var osMutexKV = mutexkv.NewMutexKV()

// Provider returns a schema.Provider for HuaweiCloud.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_ACCESS_KEY", ""),
				Description: descriptions["access_key"],
			},

			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_SECRET_KEY", ""),
				Description: descriptions["secret_key"],
			},

			"auth_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_AUTH_URL", ""),
				Description: descriptions["auth_url"],
			},

			"region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["region"],
				DefaultFunc: schema.EnvDefaultFunc("OS_REGION_NAME", ""),
			},

			"user_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USERNAME", ""),
				Description: descriptions["user_name"],
			},

			"user_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_ID", ""),
				Description: descriptions["user_name"],
			},

			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_ID",
					"OS_PROJECT_ID",
				}, ""),
				Description: descriptions["tenant_id"],
			},

			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_NAME",
					"OS_PROJECT_NAME",
				}, ""),
				Description: descriptions["tenant_name"],
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PASSWORD", ""),
				Description: descriptions["password"],
			},

			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_AUTH_TOKEN", ""),
				Description: descriptions["token"],
			},

			"domain_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_USER_DOMAIN_ID",
					"OS_PROJECT_DOMAIN_ID",
					"OS_DOMAIN_ID",
				}, ""),
				Description: descriptions["domain_id"],
			},

			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_USER_DOMAIN_NAME",
					"OS_PROJECT_DOMAIN_NAME",
					"OS_DOMAIN_NAME",
					"OS_DEFAULT_DOMAIN",
				}, ""),
				Description: descriptions["domain_name"],
			},

			"insecure": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_INSECURE", ""),
				Description: descriptions["insecure"],
			},

			"endpoint_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_ENDPOINT_TYPE", ""),
			},

			"cacert_file": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CACERT", ""),
				Description: descriptions["cacert_file"],
			},

			"cert": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CERT", ""),
				Description: descriptions["cert"],
			},

			"key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_KEY", ""),
				Description: descriptions["key"],
			},

			"swauth": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_SWAUTH", ""),
				Description: descriptions["swauth"],
			},

			"use_octavia": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USE_OCTAVIA", ""),
				Description: descriptions["use_octavia"],
			},

			"cloud": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CLOUD", ""),
				Description: descriptions["cloud"],
			},

			"agency_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_AGENCY_NAME", ""),
				Description: descriptions["agency_name"],
			},

			"agency_domain_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_AGENCY_DOMAIN_NAME", ""),
				Description: descriptions["agency_domain_name"],
			},
			"delegated_project": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DELEGATED_PROJECT", ""),
				Description: descriptions["delegated_project"],
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"huaweicloud_networking_network_v2":  dataSourceNetworkingNetworkV2(),
			"huaweicloud_networking_subnet_v2":   dataSourceNetworkingSubnetV2(),
			"huaweicloud_networking_secgroup_v2": dataSourceNetworkingSecGroupV2(),
			"huaweicloud_s3_bucket_object":       dataSourceS3BucketObject(),
			"huaweicloud_kms_key_v1":             dataSourceKmsKeyV1(),
			"huaweicloud_kms_data_key_v1":        dataSourceKmsDataKeyV1(),
			"huaweicloud_rds_flavors_v1":         dataSourceRdsFlavorV1(),
			"huaweicloud_rts_stack_v1":           dataSourceRTSStackV1(),
			"huaweicloud_rts_stack_resource_v1":  dataSourceRTSStackResourcesV1(),
			"huaweicloud_rts_software_config_v1": dataSourceRtsSoftwareConfigV1(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"huaweicloud_blockstorage_volume_v2":          resourceBlockStorageVolumeV2(),
			"huaweicloud_compute_instance_v2":             resourceComputeInstanceV2(),
			"huaweicloud_compute_keypair_v2":              resourceComputeKeypairV2(),
			"huaweicloud_compute_secgroup_v2":             resourceComputeSecGroupV2(),
			"huaweicloud_compute_servergroup_v2":          resourceComputeServerGroupV2(),
			"huaweicloud_compute_floatingip_v2":           resourceComputeFloatingIPV2(),
			"huaweicloud_compute_floatingip_associate_v2": resourceComputeFloatingIPAssociateV2(),
			"huaweicloud_compute_volume_attach_v2":        resourceComputeVolumeAttachV2(),
			"huaweicloud_dns_recordset_v2":                resourceDNSRecordSetV2(),
			"huaweicloud_dns_zone_v2":                     resourceDNSZoneV2(),
			"huaweicloud_fw_firewall_group_v2":            resourceFWFirewallGroupV2(),
			"huaweicloud_fw_policy_v2":                    resourceFWPolicyV2(),
			"huaweicloud_fw_rule_v2":                      resourceFWRuleV2(),
			"huaweicloud_kms_key_v1":                      resourceKmsKeyV1(),
			"huaweicloud_elb_loadbalancer":                resourceELBLoadBalancer(),
			"huaweicloud_elb_listener":                    resourceELBListener(),
			"huaweicloud_elb_healthcheck":                 resourceELBHealthCheck(),
			"huaweicloud_elb_backendecs":                  resourceELBBackendECS(),
			"huaweicloud_lb_loadbalancer_v2":              resourceLoadBalancerV2(),
			"huaweicloud_lb_listener_v2":                  resourceListenerV2(),
			"huaweicloud_lb_pool_v2":                      resourcePoolV2(),
			"huaweicloud_lb_member_v2":                    resourceMemberV2(),
			"huaweicloud_lb_monitor_v2":                   resourceMonitorV2(),
			"huaweicloud_networking_network_v2":           resourceNetworkingNetworkV2(),
			"huaweicloud_networking_subnet_v2":            resourceNetworkingSubnetV2(),
			"huaweicloud_networking_floatingip_v2":        resourceNetworkingFloatingIPV2(),
			"huaweicloud_networking_port_v2":              resourceNetworkingPortV2(),
			"huaweicloud_networking_router_v2":            resourceNetworkingRouterV2(),
			"huaweicloud_networking_router_interface_v2":  resourceNetworkingRouterInterfaceV2(),
			"huaweicloud_networking_router_route_v2":      resourceNetworkingRouterRouteV2(),
			"huaweicloud_networking_secgroup_v2":          resourceNetworkingSecGroupV2(),
			"huaweicloud_networking_secgroup_rule_v2":     resourceNetworkingSecGroupRuleV2(),
			"huaweicloud_s3_bucket":                       resourceS3Bucket(),
			"huaweicloud_s3_bucket_policy":                resourceS3BucketPolicy(),
			"huaweicloud_s3_bucket_object":                resourceS3BucketObject(),
			"huaweicloud_smn_topic_v2":                    resourceTopic(),
			"huaweicloud_smn_subscription_v2":             resourceSubscription(),
			"huaweicloud_rds_instance_v1":                 resourceRdsInstance(),
			"huaweicloud_nat_gateway_v2":                  resourceNatGatewayV2(),
			"huaweicloud_nat_snat_rule_v2":                resourceNatSnatRuleV2(),
			"huaweicloud_vpc_eip_v1":                      resourceVpcEIPV1(),
			"huaweicloud_rts_stack_v1":                    resourceRTSStackV1(),
			"huaweicloud_rts_software_config_v1":          resourceSoftwareConfigV1(),
		},

		ConfigureFunc: configureProvider,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"auth_url": "The Identity authentication URL.",

		"region": "The HuaweiCloud region to connect to.",

		"user_name": "Username to login with.",

		"user_id": "User ID to login with.",

		"tenant_id": "The ID of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"tenant_name": "The name of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"password": "Password to login with.",

		"token": "Authentication token to use as an alternative to username/password.",

		"domain_id": "The ID of the Domain to scope to (Identity v3).",

		"domain_name": "The name of the Domain to scope to (Identity v3).",

		"insecure": "Trust self-signed certificates.",

		"cacert_file": "A Custom CA certificate.",

		"endpoint_type": "The catalog endpoint type to use.",

		"cert": "A client certificate to authenticate with.",

		"key": "A client private key to authenticate with.",

		"swauth": "Use Swift's authentication system instead of Keystone. Only used for\n" +
			"interaction with Swift.",

		"use_octavia": "If set to `true`, API requests will go the Load Balancer\n" +
			"service (Octavia) instead of the Networking service (Neutron).",

		"cloud": "An entry in a `clouds.yaml` file to use.",

		"agency_name": "The name of agency",

		"agency_domain_name": "The name of domain who created the agency (Identity v3).",

		"delegated_project": "The name of delegated project (Identity v3).",
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccessKey:        d.Get("access_key").(string),
		SecretKey:        d.Get("secret_key").(string),
		CACertFile:       d.Get("cacert_file").(string),
		ClientCertFile:   d.Get("cert").(string),
		ClientKeyFile:    d.Get("key").(string),
		Cloud:            d.Get("cloud").(string),
		DomainID:         d.Get("domain_id").(string),
		DomainName:       d.Get("domain_name").(string),
		EndpointType:     d.Get("endpoint_type").(string),
		IdentityEndpoint: d.Get("auth_url").(string),
		Insecure:         d.Get("insecure").(bool),
		Password:         d.Get("password").(string),
		Region:           d.Get("region").(string),
		Swauth:           d.Get("swauth").(bool),
		Token:            d.Get("token").(string),
		TenantID:         d.Get("tenant_id").(string),
		TenantName:       d.Get("tenant_name").(string),
		Username:         d.Get("user_name").(string),
		UserID:           d.Get("user_id").(string),
		useOctavia:       d.Get("use_octavia").(bool),
		AgencyName:       d.Get("agency_name").(string),
		AgencyDomainName: d.Get("agency_domain_name").(string),
		DelegatedProject: d.Get("delegated_project").(string),
	}

	if err := config.LoadAndValidate(); err != nil {
		return nil, err
	}

	return &config, nil
}
