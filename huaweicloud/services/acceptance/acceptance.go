//nolint:revive
package acceptance

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud"
)

var (
	HW_REGION_NAME        = os.Getenv("HW_REGION_NAME")
	HW_REGION_NAME_1      = os.Getenv("HW_REGION_NAME_1")
	HW_REGION_NAME_2      = os.Getenv("HW_REGION_NAME_2")
	HW_REGION_NAME_3      = os.Getenv("HW_REGION_NAME_3")
	HW_CUSTOM_REGION_NAME = os.Getenv("HW_CUSTOM_REGION_NAME")
	HW_AVAILABILITY_ZONE  = os.Getenv("HW_AVAILABILITY_ZONE")
	HW_ACCESS_KEY         = os.Getenv("HW_ACCESS_KEY")
	HW_SECRET_KEY         = os.Getenv("HW_SECRET_KEY")
	HW_USER_ID            = os.Getenv("HW_USER_ID")
	HW_USER_NAME          = os.Getenv("HW_USER_NAME")
	HW_PROJECT_ID         = os.Getenv("HW_PROJECT_ID")
	HW_PROJECT_ID_1       = os.Getenv("HW_PROJECT_ID_1")
	HW_PROJECT_ID_2       = os.Getenv("HW_PROJECT_ID_2")
	HW_PROJECT_ID_3       = os.Getenv("HW_PROJECT_ID_3")

	HW_DEDICATED_HOST_ID = os.Getenv("HW_DEDICATED_HOST_ID")

	HW_DOMAIN_ID                          = os.Getenv("HW_DOMAIN_ID")
	HW_DOMAIN_NAME                        = os.Getenv("HW_DOMAIN_NAME")
	HW_ENTERPRISE_PROJECT_ID_TEST         = os.Getenv("HW_ENTERPRISE_PROJECT_ID_TEST")
	HW_ENTERPRISE_MIGRATE_PROJECT_ID_TEST = os.Getenv("HW_ENTERPRISE_MIGRATE_PROJECT_ID_TEST")

	HW_FLAVOR_ID                     = os.Getenv("HW_FLAVOR_ID")
	HW_FLAVOR_NAME                   = os.Getenv("HW_FLAVOR_NAME")
	HW_IMAGE_ID                      = os.Getenv("HW_IMAGE_ID")
	HW_IMAGE_NAME                    = os.Getenv("HW_IMAGE_NAME")
	HW_IMS_DATA_DISK_IMAGE_ID        = os.Getenv("HW_IMS_DATA_DISK_IMAGE_ID")
	HW_VPC_ID                        = os.Getenv("HW_VPC_ID")
	HW_VPC_ENHANCED_LOCAL_ROUTE      = os.Getenv("HW_VPC_ENHANCED_LOCAL_ROUTE")
	HW_VPN_P2C_GATEWAY_ID            = os.Getenv("HW_VPN_P2C_GATEWAY_ID")
	HW_VPN_P2C_SERVER                = os.Getenv("HW_VPN_P2C_SERVER")
	HW_VPN_P2C_CLIENT_CA_CERTIFICATE = os.Getenv("HW_VPN_P2C_CLIENT_CA_CERTIFICATE")
	HW_NETWORK_ID                    = os.Getenv("HW_NETWORK_ID")
	HW_SUBNET_ID                     = os.Getenv("HW_SUBNET_ID")
	HW_SECURITY_GROUP_ID             = os.Getenv("HW_SECURITY_GROUP_ID")
	HW_ENTERPRISE_PROJECT_ID         = os.Getenv("HW_ENTERPRISE_PROJECT_ID")
	HW_ADMIN                         = os.Getenv("HW_ADMIN")
	HW_IAM_V5                        = os.Getenv("HW_IAM_V5")
	HW_RUNNER_PUBLIC_IP              = os.Getenv("HW_RUNNER_PUBLIC_IP")

	HW_APIG_DEDICATED_INSTANCE_ID             = os.Getenv("HW_APIG_DEDICATED_INSTANCE_ID")
	HW_APIG_DEDICATED_INSTANCE_USED_SUBNET_ID = os.Getenv("HW_APIG_DEDICATED_INSTANCE_USED_SUBNET_ID")

	HW_CAE_ENVIRONMENT_ID     = os.Getenv("HW_CAE_ENVIRONMENT_ID")
	HW_CAE_APPLICATION_ID     = os.Getenv("HW_CAE_APPLICATION_ID")
	HW_CAE_CODE_URL           = os.Getenv("HW_CAE_CODE_URL")
	HW_CAE_CODE_BRANCH        = os.Getenv("HW_CAE_CODE_BRANCH")
	HW_CAE_CODE_AUTH_NAME     = os.Getenv("HW_CAE_CODE_AUTH_NAME")
	HW_CAE_CODE_NAMESPACE     = os.Getenv("HW_CAE_CODE_NAMESPACE")
	HW_CAE_ARTIFACT_NAMESPACE = os.Getenv("HW_CAE_ARTIFACT_NAMESPACE")
	HW_CAE_BUILD_BASE_IMAGE   = os.Getenv("HW_CAE_BUILD_BASE_IMAGE")
	HW_CAE_IMAGE_URL          = os.Getenv("HW_CAE_IMAGE_URL")

	HW_MAPREDUCE_CUSTOM           = os.Getenv("HW_MAPREDUCE_CUSTOM")
	HW_MAPREDUCE_BOOTSTRAP_SCRIPT = os.Getenv("HW_MAPREDUCE_BOOTSTRAP_SCRIPT")

	HW_CNAD_ENABLE_FLAG       = os.Getenv("HW_CNAD_ENABLE_FLAG")
	HW_CNAD_PROJECT_OBJECT_ID = os.Getenv("HW_CNAD_PROJECT_OBJECT_ID")

	HW_OBS_DESTINATION_BUCKET = os.Getenv("HW_OBS_DESTINATION_BUCKET")
	HW_OBS_USER_DOMAIN_NAME1  = os.Getenv("HW_OBS_USER_DOMAIN_NAME1")
	HW_OBS_USER_DOMAIN_NAME2  = os.Getenv("HW_OBS_USER_DOMAIN_NAME2")
	HW_OBS_ENDPOINT           = os.Getenv("HW_OBS_ENDPOINT")

	HW_OMS_ENABLE_FLAG = os.Getenv("HW_OMS_ENABLE_FLAG")

	HW_DEPRECATED_ENVIRONMENT = os.Getenv("HW_DEPRECATED_ENVIRONMENT")
	HW_INTERNAL_USED          = os.Getenv("HW_INTERNAL_USED")

	HW_WAF_ENABLE_FLAG         = os.Getenv("HW_WAF_ENABLE_FLAG")
	HW_WAF_CERTIFICATE_ID      = os.Getenv("HW_WAF_CERTIFICATE_ID")
	HW_WAF_TYPE                = os.Getenv("HW_WAF_TYPE")
	HW_WAF_CLOUD_INSTANCE_FLAG = os.Getenv("HW_WAF_CLOUD_INSTANCE_FLAG")
	HW_WAF_INTERNATIONAL_FLAG  = os.Getenv("HW_WAF_INTERNATIONAL_FLAG")
	HW_WAF_GROUP_FLAG          = os.Getenv("HW_WAF_GROUP_FLAG")

	HW_ELB_CERT_ID = os.Getenv("HW_ELB_CERT_ID")

	HW_DBSS_INSATNCE_ID = os.Getenv("HW_DBSS_INSATNCE_ID")

	HW_DEW_ENABLE_FLAG   = os.Getenv("HW_DEW_ENABLE_FLAG")
	HW_KPS_KEY_FILE_PATH = os.Getenv("HW_KPS_KEY_FILE_PATH")

	HW_DEST_REGION          = os.Getenv("HW_DEST_REGION")
	HW_DEST_PROJECT_ID      = os.Getenv("HW_DEST_PROJECT_ID")
	HW_DEST_PROJECT_ID_TEST = os.Getenv("HW_DEST_PROJECT_ID_TEST")
	HW_CHARGING_MODE        = os.Getenv("HW_CHARGING_MODE")
	HW_HIGH_COST_ALLOW      = os.Getenv("HW_HIGH_COST_ALLOW")
	HW_SWR_SHARING_ACCOUNT  = os.Getenv("HW_SWR_SHARING_ACCOUNT")

	HW_RAM_SHARE_ACCOUNT_ID          = os.Getenv("HW_RAM_SHARE_ACCOUNT_ID")
	HW_RAM_SHARE_RESOURCE_URN        = os.Getenv("HW_RAM_SHARE_RESOURCE_URN")
	HW_RAM_SHARE_UPDATE_ACCOUNT_ID   = os.Getenv("HW_RAM_SHARE_UPDATE_ACCOUNT_ID")
	HW_RAM_SHARE_UPDATE_RESOURCE_URN = os.Getenv("HW_RAM_SHARE_UPDATE_RESOURCE_URN")
	HW_RAM_ENABLE_FLAG               = os.Getenv("HW_RAM_ENABLE_FLAG")
	HW_RAM_SHARE_INVITATION_ID       = os.Getenv("HW_RAM_SHARE_INVITATION_ID")
	HW_RAM_SHARE_ID                  = os.Getenv("HW_RAM_SHARE_ID")

	HW_RMS_TARGET_ID_FOR_FGS        = os.Getenv("HW_RMS_TARGET_ID_FOR_FGS")
	HW_RMS_TARGET_ID_FOR_RFS        = os.Getenv("HW_RMS_TARGET_ID_FOR_RFS")
	HW_RMS_EXCLUDED_ACCOUNT_1       = os.Getenv("HW_RMS_EXCLUDED_ACCOUNT_1")
	HW_RMS_EXCLUDED_ACCOUNT_2       = os.Getenv("HW_RMS_EXCLUDED_ACCOUNT_2")
	HW_RMS_RESOURCE_RECORDER_CLOSED = os.Getenv("HW_RMS_RESOURCE_RECORDER_CLOSED")

	HW_CDN_DOMAIN_NAME = os.Getenv("HW_CDN_DOMAIN_NAME")
	// `HW_CDN_CERT_DOMAIN_NAME` Configure the domain name environment variable of the certificate type.
	HW_CDN_CERT_DOMAIN_NAME = os.Getenv("HW_CDN_CERT_DOMAIN_NAME")
	HW_CDN_DOMAIN_URL       = os.Getenv("HW_CDN_DOMAIN_URL")
	HW_CDN_CERT_PATH        = os.Getenv("HW_CDN_CERT_PATH")
	HW_CDN_PRIVATE_KEY_PATH = os.Getenv("HW_CDN_PRIVATE_KEY_PATH")
	HW_CDN_ENABLE_FLAG      = os.Getenv("HW_CDN_ENABLE_FLAG")
	HW_CDN_TIMESTAMP        = os.Getenv("HW_CDN_TIMESTAMP")
	HW_CDN_START_TIME       = os.Getenv("HW_CDN_START_TIME")
	HW_CDN_END_TIME         = os.Getenv("HW_CDN_END_TIME")
	HW_CDN_STAT_TYPE        = os.Getenv("HW_CDN_STAT_TYPE")

	// CCM environment
	HW_CCM_CERTIFICATE_CONTENT_PATH    = os.Getenv("HW_CCM_CERTIFICATE_CONTENT_PATH")
	HW_CCM_CERTIFICATE_CHAIN_PATH      = os.Getenv("HW_CCM_CERTIFICATE_CHAIN_PATH")
	HW_CCM_PRIVATE_KEY_PATH            = os.Getenv("HW_CCM_PRIVATE_KEY_PATH")
	HW_CCM_ENC_CERTIFICATE_PATH        = os.Getenv("HW_CCM_ENC_CERTIFICATE_PATH")
	HW_CCM_ENC_PRIVATE_KEY_PATH        = os.Getenv("HW_CCM_ENC_PRIVATE_KEY_PATH")
	HW_CCM_CERTIFICATE_PROJECT         = os.Getenv("HW_CCM_CERTIFICATE_PROJECT")
	HW_CCM_CERTIFICATE_PROJECT_UPDATED = os.Getenv("HW_CCM_CERTIFICATE_PROJECT_UPDATED")
	HW_CCM_CERTIFICATE_NAME            = os.Getenv("HW_CCM_CERTIFICATE_NAME")
	HW_CCM_PRIVATE_CA_ID               = os.Getenv("HW_CCM_PRIVATE_CA_ID")
	HW_CCM_SSL_CERTIFICATE_ID          = os.Getenv("HW_CCM_SSL_CERTIFICATE_ID")
	HW_CCM_ENABLE_FLAG                 = os.Getenv("HW_CCM_ENABLE_FLAG")

	HW_CPH_OBS_OBJECT_PATH = os.Getenv("HW_CPH_OBS_OBJECT_PATH")
	HW_CPH_OBS_BUCKET_NAME = os.Getenv("HW_CPH_OBS_BUCKET_NAME")

	HW_DMS_ENVIRONMENT   = os.Getenv("HW_DMS_ENVIRONMENT")
	HW_SMS_SOURCE_SERVER = os.Getenv("HW_SMS_SOURCE_SERVER")

	HW_DLI_AUTHORIZED_USER_NAME         = os.Getenv("HW_DLI_AUTHORIZED_USER_NAME")
	HW_DLI_FLINK_JAR_OBS_PATH           = os.Getenv("HW_DLI_FLINK_JAR_OBS_PATH")
	HW_DLI_DS_AUTH_CSS_OBS_PATH         = os.Getenv("HW_DLI_DS_AUTH_CSS_OBS_PATH")
	HW_DLI_DS_AUTH_KAFKA_TRUST_OBS_PATH = os.Getenv("HW_DLI_DS_AUTH_KAFKA_TRUST_OBS_PATH")
	HW_DLI_DS_AUTH_KAFKA_KEY_OBS_PATH   = os.Getenv("HW_DLI_DS_AUTH_KAFKA_KEY_OBS_PATH")
	HW_DLI_DS_AUTH_KRB_CONF_OBS_PATH    = os.Getenv("HW_DLI_DS_AUTH_KRB_CONF_OBS_PATH")
	HW_DLI_DS_AUTH_KRB_TAB_OBS_PATH     = os.Getenv("HW_DLI_DS_AUTH_KRB_TAB_OBS_PATH")
	HW_DLI_AGENCY_FLAG                  = os.Getenv("HW_DLI_AGENCY_FLAG")
	HW_DLI_OWNER                        = os.Getenv("HW_DLI_OWNER")
	HW_DLI_ELASTIC_RESOURCE_POOL_NAMES  = os.Getenv("HW_DLI_ELASTIC_RESOURCE_POOL_NAMES")
	HW_DLI_SQL_QUEUE_NAME               = os.Getenv("HW_DLI_SQL_QUEUE_NAME")
	HW_DLI_GENERAL_QUEUE_NAME           = os.Getenv("HW_DLI_GENERAL_QUEUE_NAME")
	HW_DLI_UPDATED_OWNER                = os.Getenv("HW_DLI_UPDATED_OWNER")
	HW_DLI_FLINK_VERSION                = os.Getenv("HW_DLI_FLINK_VERSION")
	HW_DLI_FLINK_STREAM_GRAPH           = os.Getenv("HW_DLI_FLINK_STREAM_GRAPH")
	HW_DLI_ELASTIC_RESOURCE_POOL        = os.Getenv("HW_DLI_ELASTIC_RESOURCE_POOL")

	HW_GITHUB_PERSONAL_TOKEN   = os.Getenv("HW_GITHUB_PERSONAL_TOKEN")   // Personal access token (Github, Gitlab, Gitee)
	HW_GITHUB_REPO_PWD         = os.Getenv("HW_GITHUB_REPO_PWD")         // Repository password (DevCloud, BitBucket)
	HW_GITHUB_REPO_URL         = os.Getenv("HW_GITHUB_REPO_URL")         // Repository URL (Github, Gitlab, Gitee)
	HW_OBS_STORAGE_URL         = os.Getenv("HW_OBS_STORAGE_URL")         // OBS storage URL where ZIP file is located
	HW_BUILD_IMAGE_URL         = os.Getenv("HW_BUILD_IMAGE_URL")         // SWR Image URL for component deployment
	HW_BUILD_IMAGE_URL_UPDATED = os.Getenv("HW_BUILD_IMAGE_URL_UPDATED") // SWR Image URL for component deployment update

	HW_GAUSSDB_MYSQL_INSTANCE_ID               = os.Getenv("HW_GAUSSDB_MYSQL_INSTANCE_ID")
	HW_GAUSSDB_MYSQL_NODE_ID                   = os.Getenv("HW_GAUSSDB_MYSQL_NODE_ID")
	HW_GAUSSDB_MYSQL_DATABASE_NAME             = os.Getenv("HW_GAUSSDB_MYSQL_DATABASE_NAME")
	HW_GAUSSDB_MYSQL_TABLE_NAME                = os.Getenv("HW_GAUSSDB_MYSQL_TABLE_NAME")
	HW_GAUSSDB_MYSQL_INSTANCE_CONFIGURATION_ID = os.Getenv("HW_GAUSSDB_MYSQL_INSTANCE_CONFIGURATION_ID")
	HW_GAUSSDB_MYSQL_BACKUP_BEGIN_TIME         = os.Getenv("HW_GAUSSDB_MYSQL_BACKUP_BEGIN_TIME")
	HW_GAUSSDB_MYSQL_BACKUP_END_TIME           = os.Getenv("HW_GAUSSDB_MYSQL_BACKUP_END_TIME")
	HW_GAUSSDB_MYSQL_JOB_ID                    = os.Getenv("HW_GAUSSDB_MYSQL_JOB_ID")
	HW_GAUSSDB_MYSQL_START_TIME                = os.Getenv("HW_GAUSSDB_MYSQL_START_TIME")
	HW_GAUSSDB_MYSQL_END_TIME                  = os.Getenv("HW_GAUSSDB_MYSQL_END_TIME")

	HW_GAUSSDB_OPENGAUSS_PARAMETER_TEMPLATE_ID = os.Getenv("HW_GAUSSDB_OPENGAUSS_PARAMETER_TEMPLATE_ID")
	HW_GAUSSDB_OPENGAUSS_JOB_ID                = os.Getenv("HW_GAUSSDB_OPENGAUSS_JOB_ID")
	HW_GAUSSDB_OPENGAUSS_INSTANCE_ID           = os.Getenv("HW_GAUSSDB_OPENGAUSS_INSTANCE_ID")
	HW_GAUSSDB_OPENGAUSS_START_TIME            = os.Getenv("HW_GAUSSDB_OPENGAUSS_START_TIME")
	HW_GAUSSDB_OPENGAUSS_END_TIME              = os.Getenv("HW_GAUSSDB_OPENGAUSS_END_TIME")

	HW_VOD_WATERMARK_FILE   = os.Getenv("HW_VOD_WATERMARK_FILE")
	HW_VOD_MEDIA_ASSET_FILE = os.Getenv("HW_VOD_MEDIA_ASSET_FILE")

	HW_LTS_ENABLE_FLAG                 = os.Getenv("HW_LTS_ENABLE_FLAG")
	HW_LTS_STRUCT_CONFIG_TEMPLATE_ID   = os.Getenv("HW_LTS_STRUCT_CONFIG_TEMPLATE_ID")
	HW_LTS_STRUCT_CONFIG_TEMPLATE_NAME = os.Getenv("HW_LTS_STRUCT_CONFIG_TEMPLATE_NAME")

	HW_LIVE_STREAMING_DOMAIN_NAME          = os.Getenv("HW_LIVE_STREAMING_DOMAIN_NAME")
	HW_LIVE_INGEST_RTMP_DOMAIN_NAME        = os.Getenv("HW_LIVE_INGEST_RTMP_DOMAIN_NAME")
	HW_LIVE_INGEST_SRT_DOMAIN_NAME         = os.Getenv("HW_LIVE_INGEST_SRT_DOMAIN_NAME")
	HW_LIVE_INGEST_DOMAIN_NAME             = os.Getenv("HW_LIVE_INGEST_DOMAIN_NAME")
	HW_LIVE_TRANSCODING_TEPLATE_ID         = os.Getenv("HW_LIVE_TRANSCODING_TEPLATE_ID")
	HW_LIVE_TRANSCODING_TEPLATE_ANOTHER_ID = os.Getenv("HW_LIVE_TRANSCODING_TEPLATE_ANOTHER_ID")
	HW_LIVE_HTTPS_TLS_CERT_BODY_PATH       = os.Getenv("HW_LIVE_HTTPS_TLS_CERT_BODY_PATH")
	HW_LIVE_HTTPS_TLS_CERT_KEY_PATH        = os.Getenv("HW_LIVE_HTTPS_TLS_CERT_KEY_PATH")

	HW_CHAIR_EMAIL              = os.Getenv("HW_CHAIR_EMAIL")
	HW_GUEST_EMAIL              = os.Getenv("HW_GUEST_EMAIL")
	HW_MEETING_ACCOUNT_NAME     = os.Getenv("HW_MEETING_ACCOUNT_NAME")
	HW_MEETING_ACCOUNT_PASSWORD = os.Getenv("HW_MEETING_ACCOUNT_PASSWORD")
	HW_MEETING_APP_ID           = os.Getenv("HW_MEETING_APP_ID")
	HW_MEETING_APP_KEY          = os.Getenv("HW_MEETING_APP_KEY")
	HW_MEETING_USER_ID          = os.Getenv("HW_MEETING_USER_ID")
	HW_MEETING_ROOM_ID          = os.Getenv("HW_MEETING_ROOM_ID")

	HW_AAD_INSTANCE_ID = os.Getenv("HW_AAD_INSTANCE_ID")
	HW_AAD_IP_ADDRESS  = os.Getenv("HW_AAD_IP_ADDRESS")

	HW_WORKSPACE_AD_DOMAIN_NAMES = os.Getenv("HW_WORKSPACE_AD_DOMAIN_NAMES") // Domain name, e.g. "example.com".
	// Please make sure all AD servers (master and standby) have the same account configuration (with same name and password).
	HW_WORKSPACE_AD_SERVER_ACCOUNT = os.Getenv("HW_WORKSPACE_AD_SERVER_ACCOUNT") // The admin user name of the AD servers.
	HW_WORKSPACE_AD_SERVER_PWD     = os.Getenv("HW_WORKSPACE_AD_SERVER_PWD")     // The password of the admin user for the AD servers.
	// The IP addresses of the AD servers.
	// The format is '{master IP address},{standby IP address}' (with the standby AD server) or '{master IP address}'.
	HW_WORKSPACE_AD_DOMAIN_IPS = os.Getenv("HW_WORKSPACE_AD_DOMAIN_IPS")
	HW_WORKSPACE_AD_VPC_ID     = os.Getenv("HW_WORKSPACE_AD_VPC_ID")     // The VPC ID to which the AD servers and desktops belong.
	HW_WORKSPACE_AD_NETWORK_ID = os.Getenv("HW_WORKSPACE_AD_NETWORK_ID") // The network ID to which the AD servers belong.
	// The internet access port to which the Workspace service.
	HW_WORKSPACE_INTERNET_ACCESS_PORT              = os.Getenv("HW_WORKSPACE_INTERNET_ACCESS_PORT")
	HW_WORKSPACE_APP_SERVER_GROUP_FLAVOR_ID        = os.Getenv("HW_WORKSPACE_APP_SERVER_GROUP_FLAVOR_ID")
	HW_WORKSPACE_APP_SERVER_GROUP_IMAGE_ID         = os.Getenv("HW_WORKSPACE_APP_SERVER_GROUP_IMAGE_ID")
	HW_WORKSPACE_APP_SERVER_GROUP_IMAGE_PRODUCT_ID = os.Getenv("HW_WORKSPACE_APP_SERVER_GROUP_IMAGE_PRODUCT_ID")
	HW_WORKSPACE_APP_SERVER_GROUP_IMAGE_SPEC_CODE  = os.Getenv("HW_WORKSPACE_APP_SERVER_GROUP_IMAGE_SPEC_CODE")
	HW_WORKSPACE_OU_NAME                           = os.Getenv("HW_WORKSPACE_OU_NAME")
	HW_WORKSPACE_APP_FILE_NAME                     = os.Getenv("HW_WORKSPACE_APP_FILE_NAME")
	HW_WORKSPACE_USER_NAMES                        = os.Getenv("HW_WORKSPACE_USER_NAMES")

	HW_FGS_AGENCY_NAME         = os.Getenv("HW_FGS_AGENCY_NAME")
	HW_FGS_APP_AGENCY_NAME     = os.Getenv("HW_FGS_APP_AGENCY_NAME")
	HW_FGS_GPU_TYPE            = os.Getenv("HW_FGS_GPU_TYPE")
	HW_FGS_DEPENDENCY_OBS_LINK = os.Getenv("HW_FGS_DEPENDENCY_OBS_LINK")

	HW_KMS_ENVIRONMENT     = os.Getenv("HW_KMS_ENVIRONMENT")
	HW_KMS_HSM_CLUSTER_ID  = os.Getenv("HW_KMS_HSM_CLUSTER_ID")
	HW_KMS_KEY_ID          = os.Getenv("HW_KMS_KEY_ID")
	HW_KMS_IMPORT_TOKEN    = os.Getenv("HW_KMS_IMPORT_TOKEN")
	HW_KMS_KEY_MATERIAL    = os.Getenv("HW_KMS_KEY_MATERIAL")
	HW_KMS_KEY_PRIVATE_KEY = os.Getenv("HW_KMS_KEY_PRIVATE_KEY")

	HW_MULTI_ACCOUNT_ENVIRONMENT            = os.Getenv("HW_MULTI_ACCOUNT_ENVIRONMENT")
	HW_ORGANIZATIONS_OPEN                   = os.Getenv("HW_ORGANIZATIONS_OPEN")
	HW_ORGANIZATIONS_ACCOUNT_ID             = os.Getenv("HW_ORGANIZATIONS_ACCOUNT_ID")
	HW_ORGANIZATIONS_ACCOUNT_NAME           = os.Getenv("HW_ORGANIZATIONS_ACCOUNT_NAME")
	HW_ORGANIZATIONS_INVITE_ACCOUNT_ID      = os.Getenv("HW_ORGANIZATIONS_INVITE_ACCOUNT_ID")
	HW_ORGANIZATIONS_ORGANIZATIONAL_UNIT_ID = os.Getenv("HW_ORGANIZATIONS_ORGANIZATIONAL_UNIT_ID")
	HW_ORGANIZATIONS_INVITATION_ID          = os.Getenv("HW_ORGANIZATIONS_INVITATION_ID")

	HW_RGC_ORGANIZATIONAL_UNIT_ID    = os.Getenv("HW_RGC_ORGANIZATIONAL_UNIT_ID")
	HW_RGC_ORGANIZATIONAL_UNIT_NAME  = os.Getenv("HW_RGC_ORGANIZATIONAL_UNIT_NAME")
	HW_RGC_BLUEPRINT_PRODUCT_ID      = os.Getenv("HW_RGC_BLUEPRINT_PRODUCT_ID")
	HW_RGC_BLUEPRINT_PRODUCT_VERSION = os.Getenv("HW_RGC_BLUEPRINT_PRODUCT_VERSION")

	HW_IDENTITY_CENTER_ACCOUNT_ID                = os.Getenv("HW_IDENTITY_CENTER_ACCOUNT_ID")
	HW_IDENTITY_CENTER_IDENTITY_POLICY_ID        = os.Getenv("HW_IDENTITY_CENTER_IDENTITY_POLICY_ID")
	HW_IDENTITY_CENTER_IDENTITY_POLICY_ID_UPDATE = os.Getenv("HW_IDENTITY_CENTER_IDENTITY_POLICY_ID_UPDATE")

	HW_ER_TEST_ON = os.Getenv("HW_ER_TEST_ON") // Whether to run the ER related tests.

	// The OBS address where the HCL/JSON template archive (No variables) is located.
	HW_RF_TEMPLATE_ARCHIVE_NO_VARS_URI = os.Getenv("HW_RF_TEMPLATE_ARCHIVE_NO_VARS_URI")
	// The OBS address where the HCL/JSON template archive is located.
	HW_RF_TEMPLATE_ARCHIVE_URI = os.Getenv("HW_RF_TEMPLATE_ARCHIVE_URI")
	// The OBS address where the variable archive corresponding to the HCL/JSON template is located.
	HW_RF_VARIABLES_ARCHIVE_URI = os.Getenv("HW_RF_VARIABLES_ARCHIVE_URI")

	// The direct connection ID (provider does not support direct connection resource).
	HW_DC_DIRECT_CONNECT_ID    = os.Getenv("HW_DC_DIRECT_CONNECT_ID")
	HW_DC_RESOURCE_TENANT_ID   = os.Getenv("HW_DC_RESOURCE_TENANT_ID")
	HW_DC_HOSTTING_ID          = os.Getenv("HW_DC_HOSTTING_ID")
	HW_DC_TARGET_TENANT_VGW_ID = os.Getenv("HW_DC_TARGET_TENANT_VGW_ID")
	HW_DC_VIRTUAL_INTERFACE_ID = os.Getenv("HW_DC_VIRTUAL_INTERFACE_ID")
	HW_DC_ENABLE_FLAG          = os.Getenv("HW_DC_ENABLE_FLAG")

	HW_DSC_INSTANCE_ID    = os.Getenv("HW_DSC_INSTANCE_ID")
	HW_DSC_ALARM_TOPIC_ID = os.Getenv("HW_DSC_ALARM_TOPIC_ID")

	HW_CES_START_TIME = os.Getenv("HW_CES_START_TIME")
	HW_CES_END_TIME   = os.Getenv("HW_CES_END_TIME")

	// The CFW instance ID
	HW_CFW_INSTANCE_ID               = os.Getenv("HW_CFW_INSTANCE_ID")
	HW_CFW_EAST_WEST_FIREWALL        = os.Getenv("HW_CFW_EAST_WEST_FIREWALL")
	HW_CFW_START_TIME                = os.Getenv("HW_CFW_START_TIME")
	HW_CFW_END_TIME                  = os.Getenv("HW_CFW_END_TIME")
	HW_CFW_PREDEFINED_SERVICE_GROUP1 = os.Getenv("HW_CFW_PREDEFINED_SERVICE_GROUP1")
	HW_CFW_PREDEFINED_SERVICE_GROUP2 = os.Getenv("HW_CFW_PREDEFINED_SERVICE_GROUP2")
	HW_CFW_PREDEFINED_ADDRESS_GROUP1 = os.Getenv("HW_CFW_PREDEFINED_ADDRESS_GROUP1")
	HW_CFW_PREDEFINED_ADDRESS_GROUP2 = os.Getenv("HW_CFW_PREDEFINED_ADDRESS_GROUP2")
	HW_CFW_IPS_CUSTOM_RULE           = os.Getenv("HW_CFW_IPS_CUSTOM_RULE")

	HW_CTS_START_TIME = os.Getenv("HW_CTS_START_TIME")
	HW_CTS_END_TIME   = os.Getenv("HW_CTS_END_TIME")

	// The cluster ID of the CCE
	HW_CCE_CLUSTER_ID = os.Getenv("HW_CCE_CLUSTER_ID")
	// The absolute chart path of the CCE
	HW_CCE_CHART_PATH = os.Getenv("HW_CCE_CHART_PATH")
	// The cluster name of the CCE
	HW_CCE_CLUSTER_NAME = os.Getenv("HW_CCE_CLUSTER_NAME")
	// The partition az of the CCE
	HW_CCE_PARTITION_AZ = os.Getenv("HW_CCE_PARTITION_AZ")
	// The namespace of the workload is located
	HW_WORKLOAD_NAMESPACE = os.Getenv("HW_WORKLOAD_NAMESPACE")
	// The workload type deployed in CCE/CCI
	HW_WORKLOAD_TYPE = os.Getenv("HW_WORKLOAD_TYPE")
	// The workload name deployed in CCE/CCI
	HW_WORKLOAD_NAME = os.Getenv("HW_WORKLOAD_NAME")
	// The target region of SWR image auto sync
	HW_SWR_TARGET_REGION = os.Getenv("HW_SWR_TARGET_REGION")
	// The target organization of SWR image auto sync
	HW_SWR_TARGET_ORGANIZATION = os.Getenv("HW_SWR_TARGET_ORGANIZATION")
	// The organization of SWR image tag
	HW_SWR_ORGANIZATION = os.Getenv("HW_SWR_ORGANIZATION")
	// The repository of SWR image tag
	HW_SWR_REPOSITORY = os.Getenv("HW_SWR_REPOSITORY")

	// The ID of the CBR vault
	HW_IMS_VAULT_ID = os.Getenv("HW_IMS_VAULT_ID")
	// The ID of the CBR backup
	HW_IMS_BACKUP_ID = os.Getenv("HW_IMS_BACKUP_ID")
	// The image file url in the OBS bucket
	HW_IMS_IMAGE_URL = os.Getenv("HW_IMS_IMAGE_URL")
	// The shared backup ID wants to accept.
	HW_SHARED_BACKUP_ID = os.Getenv("HW_SHARED_BACKUP_ID")

	// The SecMaster workspace ID
	HW_SECMASTER_WORKSPACE_ID = os.Getenv("HW_SECMASTER_WORKSPACE_ID")
	// The SecMaster indicator ID
	HW_SECMASTER_INDICATOR_TYPE_ID        = os.Getenv("HW_SECMASTER_INDICATOR_TYPE_ID")
	HW_SECMASTER_INDICATOR_TYPE_ID_UPDATE = os.Getenv("HW_SECMASTER_INDICATOR_TYPE_ID_UPDATE")
	HW_SECMASTER_METRIC_ID                = os.Getenv("HW_SECMASTER_METRIC_ID")

	// The SecMaster pipeline ID
	HW_SECMASTER_PIPELINE_ID = os.Getenv("HW_SECMASTER_PIPELINE_ID")

	// The SecMaster playbook instance ID
	HW_SECMASTER_INSTANCE_ID = os.Getenv("HW_SECMASTER_INSTANCE_ID")

	HW_MODELARTS_HAS_SUBSCRIBE_MODEL = os.Getenv("HW_MODELARTS_HAS_SUBSCRIBE_MODEL")
	HW_MODELARTS_USER_LOGIN_PASSWORD = os.Getenv("HW_MODELARTS_USER_LOGIN_PASSWORD")
	HW_MODELARTS_DEVSERVER_FLAVOR    = os.Getenv("HW_MODELARTS_DEVSERVER_FLAVOR")
	HW_MODELARTS_DEVSERVER_IMAGE_ID  = os.Getenv("HW_MODELARTS_DEVSERVER_IMAGE_ID")

	// The CMDB sub-application ID of AOM service
	HW_AOM_SUB_APPLICATION_ID                    = os.Getenv("HW_AOM_SUB_APPLICATION_ID")
	HW_AOM_MULTI_ACCOUNT_AGGREGATION_RULE_ENABLE = os.Getenv("HW_AOM_MULTI_ACCOUNT_AGGREGATION_RULE_ENABLE")

	// the ID of ECS instance which has installed uniagent
	HW_COC_INSTANCE_ID = os.Getenv("HW_COC_INSTANCE_ID")

	// Deprecated
	HW_SRC_ACCESS_KEY = os.Getenv("HW_SRC_ACCESS_KEY")
	HW_SRC_SECRET_KEY = os.Getenv("HW_SRC_SECRET_KEY")
	HW_EXTGW_ID       = os.Getenv("HW_EXTGW_ID")
	HW_POOL_NAME      = os.Getenv("HW_POOL_NAME")

	HW_IMAGE_SHARE_SOURCE_IMAGE_ID = os.Getenv("HW_IMAGE_SHARE_SOURCE_IMAGE_ID")

	HW_CERTIFICATE_CONTENT         = os.Getenv("HW_CERTIFICATE_CONTENT")
	HW_CERTIFICATE_CONTENT_UPDATE  = os.Getenv("HW_CERTIFICATE_CONTENT_UPDATE")
	HW_CERTIFICATE_PRIVATE_KEY     = os.Getenv("HW_CERTIFICATE_PRIVATE_KEY")
	HW_CERTIFICATE_ROOT_CA         = os.Getenv("HW_CERTIFICATE_ROOT_CA")
	HW_NEW_CERTIFICATE_CONTENT     = os.Getenv("HW_NEW_CERTIFICATE_CONTENT")
	HW_NEW_CERTIFICATE_PRIVATE_KEY = os.Getenv("HW_NEW_CERTIFICATE_PRIVATE_KEY")
	HW_NEW_CERTIFICATE_ROOT_CA     = os.Getenv("HW_NEW_CERTIFICATE_ROOT_CA")

	HW_GM_CERTIFICATE_CONTENT             = os.Getenv("HW_GM_CERTIFICATE_CONTENT")
	HW_GM_CERTIFICATE_PRIVATE_KEY         = os.Getenv("HW_GM_CERTIFICATE_PRIVATE_KEY")
	HW_GM_ENC_CERTIFICATE_CONTENT         = os.Getenv("HW_GM_ENC_CERTIFICATE_CONTENT")
	HW_GM_ENC_CERTIFICATE_PRIVATE_KEY     = os.Getenv("HW_GM_ENC_CERTIFICATE_PRIVATE_KEY")
	HW_GM_CERTIFICATE_CHAIN               = os.Getenv("HW_GM_CERTIFICATE_CHAIN")
	HW_NEW_GM_CERTIFICATE_CONTENT         = os.Getenv("HW_NEW_GM_CERTIFICATE_CONTENT")
	HW_NEW_GM_CERTIFICATE_PRIVATE_KEY     = os.Getenv("HW_NEW_GM_CERTIFICATE_PRIVATE_KEY")
	HW_NEW_GM_ENC_CERTIFICATE_CONTENT     = os.Getenv("HW_NEW_GM_ENC_CERTIFICATE_CONTENT")
	HW_NEW_GM_ENC_CERTIFICATE_PRIVATE_KEY = os.Getenv("HW_NEW_GM_ENC_CERTIFICATE_PRIVATE_KEY")
	HW_NEW_GM_CERTIFICATE_CHAIN           = os.Getenv("HW_NEW_GM_CERTIFICATE_CHAIN")

	HW_CODEARTS_RESOURCE_POOL_ID   = os.Getenv("HW_CODEARTS_RESOURCE_POOL_ID")
	HW_CODEARTS_DEPLOYMENT_TASK_ID = os.Getenv("HW_CODEARTS_DEPLOYMENT_TASK_ID")

	HW_CODEARTS_ENABLE_FLAG       = os.Getenv("HW_CODEARTS_ENABLE_FLAG")
	HW_CODEARTS_PUBLIC_IP_ADDRESS = os.Getenv("HW_CODEARTS_PUBLIC_IP_ADDRESS")
	HW_CODEARTS_SSH_CREDENTIAL_ID = os.Getenv("HW_CODEARTS_SSH_CREDENTIAL_ID")

	HW_EG_TEST_ON     = os.Getenv("HW_EG_TEST_ON") // Whether to run the EG related tests.
	HW_EG_CHANNEL_ID  = os.Getenv("HW_EG_CHANNEL_ID")
	HW_EG_AGENCY_NAME = os.Getenv("HW_EG_AGENCY_NAME")

	HW_KOOGALLERY_ASSET = os.Getenv("HW_KOOGALLERY_ASSET")

	HW_CCI_NAMESPACE = os.Getenv("HW_CCI_NAMESPACE")

	HW_CC_GLOBAL_GATEWAY_ID  = os.Getenv("HW_CC_GLOBAL_GATEWAY_ID")
	HW_CC_PEER_DOMAIN_ID     = os.Getenv("HW_CC_PEER_DOMAIN_ID")
	HW_CC_PEER_CONNECTION_ID = os.Getenv("HW_CC_PEER_CONNECTION_ID")

	HW_CC_PERMISSION_ID = os.Getenv("HW_CC_PERMISSION_ID")

	HW_CSE_MICROSERVICE_ENGINE_ID             = os.Getenv("HW_CSE_MICROSERVICE_ENGINE_ID")
	HW_CSE_MICROSERVICE_ENGINE_ADMIN_PASSWORD = os.Getenv("HW_CSE_MICROSERVICE_ENGINE_ADMIN_PASSWORD")
	HW_CSE_NACOS_MICROSERVICE_ENGINE_ID       = os.Getenv("HW_CSE_NACOS_MICROSERVICE_ENGINE_ID")

	HW_CSS_LOCAL_DISK_FLAVOR  = os.Getenv("HW_CSS_LOCAL_DISK_FLAVOR")
	HW_CSS_ELB_AGENCY         = os.Getenv("HW_CSS_ELB_AGENCY")
	HW_CSS_UPGRADE_AGENCY     = os.Getenv("HW_CSS_UPGRADE_AGENCY")
	HW_CSS_LOW_ENGINE_VERSION = os.Getenv("HW_CSS_LOW_ENGINE_VERSION")
	HW_CSS_TARGET_IMAGE_ID    = os.Getenv("HW_CSS_TARGET_IMAGE_ID")
	HW_CSS_REPLACE_AGENCY     = os.Getenv("HW_CSS_REPLACE_AGENCY")
	HW_CSS_AZ_MIGRATE_AGENCY  = os.Getenv("HW_CSS_AZ_MIGRATE_AGENCY")

	HW_CERT_BATCH_PUSH_ID     = os.Getenv("HW_CERT_BATCH_PUSH_ID")
	HW_CERT_BATCH_PUSH_WAF_ID = os.Getenv("HW_CERT_BATCH_PUSH_WAF_ID")

	HW_AS_SCALING_GROUP_ID         = os.Getenv("HW_AS_SCALING_GROUP_ID")
	HW_AS_SCALING_CONFIGURATION_ID = os.Getenv("HW_AS_SCALING_CONFIGURATION_ID")
	HW_AS_SCALING_POLICY_ID        = os.Getenv("HW_AS_SCALING_POLICY_ID")
	HW_AS_LIFECYCLE_ACTION_KEY     = os.Getenv("HW_AS_LIFECYCLE_ACTION_KEY")
	HW_AS_INSTANCE_ID              = os.Getenv("HW_AS_INSTANCE_ID")
	HW_AS_LIFECYCLE_HOOK_NAME      = os.Getenv("HW_AS_LIFECYCLE_HOOK_NAME")

	// Common
	HW_DATAARTS_WORKSPACE_ID                               = os.Getenv("HW_DATAARTS_WORKSPACE_ID")
	HW_DATAARTS_CDM_NAME                                   = os.Getenv("HW_DATAARTS_CDM_NAME")
	HW_DATAARTS_MANAGER_ID                                 = os.Getenv("HW_DATAARTS_MANAGER_ID")
	HW_DATAARTS_BIZ_CATALOG_ID                             = os.Getenv("HW_DATAARTS_BIZ_CATALOG_ID")
	HW_DATAARTS_SECRECY_LEVEL_ID                           = os.Getenv("HW_DATAARTS_SECRECY_LEVEL_ID")
	HW_DATAARTS_SECRECY_LEVEL_ID_UPDATE                    = os.Getenv("HW_DATAARTS_SECRECY_LEVEL_ID_UPDATE")
	HW_DATAARTS_CATEGORY_ID                                = os.Getenv("HW_DATAARTS_CATEGORY_ID")
	HW_DATAARTS_CATEGORY_ID_UPDATE                         = os.Getenv("HW_DATAARTS_CATEGORY_ID_UPDATE")
	HW_DATAARTS_BUILTIN_RULE_ID                            = os.Getenv("HW_DATAARTS_BUILTIN_RULE_ID")
	HW_DATAARTS_BUILTIN_RULE_NAME                          = os.Getenv("HW_DATAARTS_BUILTIN_RULE_NAME")
	HW_DATAARTS_SUBJECT_ID                                 = os.Getenv("HW_DATAARTS_SUBJECT_ID")
	HW_DATAARTS_ARCHITECTURE_USER_ID                       = os.Getenv("HW_DATAARTS_ARCHITECTURE_USER_ID")
	HW_DATAARTS_ARCHITECTURE_USER_NAME                     = os.Getenv("HW_DATAARTS_ARCHITECTURE_USER_NAME")
	HW_DATAARTS_SECURITY_PERMISSSIONSET_MEMBER_OBJECT_ID   = os.Getenv("HW_DATAARTS_SECURITY_PERMISSSIONSET_MEMBER_OBJECT_ID")
	HW_DATAARTS_SECURITY_PERMISSSIONSET_MEMBER_OBJECT_NAME = os.Getenv("HW_DATAARTS_SECURITY_PERMISSSIONSET_MEMBER_OBJECT_NAME")
	// Management Center
	HW_DATAARTS_CONNECTION_ID   = os.Getenv("HW_DATAARTS_CONNECTION_ID")
	HW_DATAARTS_CONNECTION_NAME = os.Getenv("HW_DATAARTS_CONNECTION_NAME")
	// Data Service
	HW_DATAARTS_REVIEWER_NAME       = os.Getenv("HW_DATAARTS_REVIEWER_NAME")
	HW_DATAARTS_DLI_QUEUE_NAME      = os.Getenv("HW_DATAARTS_DLI_QUEUE_NAME")
	HW_DATAARTS_INSTANCE_ID_IN_APIG = os.Getenv("HW_DATAARTS_INSTANCE_ID_IN_APIG")

	HW_EVS_AVAILABILITY_ZONE_GPSSD2 = os.Getenv("HW_EVS_AVAILABILITY_ZONE_GPSSD2")
	HW_EVS_AVAILABILITY_ZONE_ESSD2  = os.Getenv("HW_EVS_AVAILABILITY_ZONE_ESSD2")
	HW_EVS_TRANSFER_ID              = os.Getenv("HW_EVS_TRANSFER_ID")
	HW_EVS_TRANSFER_AUTH_KEY        = os.Getenv("HW_EVS_TRANSFER_AUTH_KEY")
	HW_EVS_ENABLE_FLAG              = os.Getenv("HW_EVS_ENABLE_FLAG")

	HW_ECS_LAUNCH_TEMPLATE_ID = os.Getenv("HW_ECS_LAUNCH_TEMPLATE_ID")

	HW_IOTDA_ACCESS_ADDRESS      = os.Getenv("HW_IOTDA_ACCESS_ADDRESS")
	HW_IOTDA_BATCHTASK_FILE_PATH = os.Getenv("HW_IOTDA_BATCHTASK_FILE_PATH")

	HW_DWS_MUTIL_AZS               = os.Getenv("HW_DWS_MUTIL_AZS")
	HW_DWS_CLUSTER_ID              = os.Getenv("HW_DWS_CLUSTER_ID")
	HW_DWS_LOGICAL_MODE_CLUSTER_ID = os.Getenv("HW_DWS_LOGICAL_MODE_CLUSTER_ID")
	HW_DWS_LOGICAL_CLUSTER_NAME    = os.Getenv("HW_DWS_LOGICAL_CLUSTER_NAME")
	HW_DWS_SNAPSHOT_POLICY_NAME    = os.Getenv("HW_DWS_SNAPSHOT_POLICY_NAME")
	// The list of the user names under specified DWS cluster. Using commas (,) to separate multiple names.
	HW_DWS_ASSOCIATE_USER_NAMES  = os.Getenv("HW_DWS_ASSOCIATE_USER_NAMES")
	HW_DWS_AUTOMATED_SNAPSHOT_ID = os.Getenv("HW_DWS_AUTOMATED_SNAPSHOT_ID")
	// The OBS agency name list of the DWS data source. Using commas (,) to separate multiple names.
	HW_DWS_OBS_AGENCY_NAMES = os.Getenv("HW_DWS_OBS_AGENCY_NAMES")

	HW_DCS_ACCOUNT_WHITELIST = os.Getenv("HW_DCS_ACCOUNT_WHITELIST")

	HW_DCS_INSTANCE_ID = os.Getenv("HW_DCS_INSTANCE_ID")
	HW_DCS_BEGIN_TIME  = os.Getenv("HW_DCS_BEGIN_TIME")
	HW_DCS_END_TIME    = os.Getenv("HW_DCS_END_TIME")

	HW_ELB_GATEWAY_TYPE = os.Getenv("HW_ELB_GATEWAY_TYPE")

	HW_LTS_AGENCY_STREAM_NAME = os.Getenv("HW_LTS_AGENCY_STREAM_NAME")
	HW_LTS_AGENCY_STREAM_ID   = os.Getenv("HW_LTS_AGENCY_STREAM_ID")
	HW_LTS_AGENCY_GROUP_NAME  = os.Getenv("HW_LTS_AGENCY_GROUP_NAME")
	HW_LTS_AGENCY_GROUP_ID    = os.Getenv("HW_LTS_AGENCY_GROUP_ID")
	HW_LTS_LOG_STREAM_NAME    = os.Getenv("HW_LTS_LOG_STREAM_NAME")
	HW_LTS_LOG_STREAM_ID      = os.Getenv("HW_LTS_LOG_STREAM_ID")
	HW_LTS_LOG_GROUP_NAME     = os.Getenv("HW_LTS_LOG_GROUP_NAME")
	HW_LTS_LOG_GROUP_ID       = os.Getenv("HW_LTS_LOG_GROUP_ID")
	HW_LTS_AGENCY_PROJECT_ID  = os.Getenv("HW_LTS_AGENCY_PROJECT_ID")
	HW_LTS_AGENCY_DOMAIN_NAME = os.Getenv("HW_LTS_AGENCY_DOMAIN_NAME")
	HW_LTS_AGENCY_NAME        = os.Getenv("HW_LTS_AGENCY_NAME")
	// The ID list of the LTS hosts. Using commas (,) to separate multiple IDs (At least two UUIDs we need).
	HW_LTS_HOST_IDS       = os.Getenv("HW_LTS_HOST_IDS")
	HW_LTS_CCE_CLUSTER_ID = os.Getenv("HW_LTS_CCE_CLUSTER_ID")

	HW_LTS_LOG_CONVERGE_ORGANIZATION_ID       = os.Getenv("HW_LTS_LOG_CONVERGE_ORGANIZATION_ID")
	HW_LTS_LOG_CONVERGE_MANAGEMENT_ACCOUNT_ID = os.Getenv("HW_LTS_LOG_CONVERGE_MANAGEMENT_ACCOUNT_ID")
	HW_LTS_LOG_CONVERGE_MEMBER_ACCOUNT_ID     = os.Getenv("HW_LTS_LOG_CONVERGE_MEMBER_ACCOUNT_ID")

	HW_LTS_LOG_CONVERGE_SOURCE_LOG_GROUP_ID  = os.Getenv("HW_LTS_LOG_CONVERGE_SOURCE_LOG_GROUP_ID")
	HW_LTS_LOG_CONVERGE_SOURCE_LOG_STREAM_ID = os.Getenv("HW_LTS_LOG_CONVERGE_SOURCE_LOG_STREAM_ID")

	HW_LTS_CLUSTER_ID           = os.Getenv("HW_LTS_CLUSTER_ID")
	HW_LTS_CLUSTER_NAME         = os.Getenv("HW_LTS_CLUSTER_NAME")
	HW_LTS_CLUSTER_ID_ANOTHER   = os.Getenv("HW_LTS_CLUSTER_ID_ANOTHER")
	HW_LTS_CLUSTER_NAME_ANOTHER = os.Getenv("HW_LTS_CLUSTER_NAME_ANOTHER")

	HW_VPCEP_SERVICE_ID = os.Getenv("HW_VPCEP_SERVICE_ID")

	HW_HSS_HOST_PROTECTION_HOST_ID  = os.Getenv("HW_HSS_HOST_PROTECTION_HOST_ID")
	HW_HSS_HOST_PROTECTION_QUOTA_ID = os.Getenv("HW_HSS_HOST_PROTECTION_QUOTA_ID")

	HW_DDM_INSTANCE_ID = os.Getenv("HW_DDM_INSTANCE_ID")
	HW_DDM_PROCESS_ID  = os.Getenv("HW_DDM_PROCESS_ID")
	HW_DDM_START_TIME  = os.Getenv("HW_DDM_START_TIME")
	HW_DDM_END_TIME    = os.Getenv("HW_DDM_END_TIME")

	HW_DDS_SECOND_LEVEL_MONITORING_ENABLED = os.Getenv("HW_DDS_SECOND_LEVEL_MONITORING_ENABLED")
	HW_DDS_INSTANCE_ID                     = os.Getenv("HW_DDS_INSTANCE_ID")
	HW_DDS_START_TIME                      = os.Getenv("HW_DDS_START_TIME")
	HW_DDS_END_TIME                        = os.Getenv("HW_DDS_END_TIME")
	HW_DDS_RECYCLE_INSTANCES_ENABLED       = os.Getenv("HW_DDS_RECYCLE_INSTANCES_ENABLED")
	HW_DDS_SCHEDULED_TASKS_ENABLED         = os.Getenv("HW_DDS_SCHEDULED_TASKS_ENABLED")

	HW_RDS_CROSS_REGION_BACKUP_INSTANCE_ID = os.Getenv("HW_RDS_CROSS_REGION_BACKUP_INSTANCE_ID")
	HW_RDS_INSTANCE_ID                     = os.Getenv("HW_RDS_INSTANCE_ID")
	HW_RDS_INSTANCE_NAME                   = os.Getenv("HW_RDS_INSTANCE_NAME")
	HW_RDS_DATABASE_NAME                   = os.Getenv("HW_RDS_DATABASE_NAME")
	HW_RDS_TABLE_NAME                      = os.Getenv("HW_RDS_TABLE_NAME")
	HW_RDS_BACKUP_ID                       = os.Getenv("HW_RDS_BACKUP_ID")
	HW_RDS_START_TIME                      = os.Getenv("HW_RDS_START_TIME")
	HW_RDS_END_TIME                        = os.Getenv("HW_RDS_END_TIME")
	HW_RDS_FIXED_IP                        = os.Getenv("HW_RDS_FIXED_IP")

	HW_DMS_KAFKA_INSTANCE_ID         = os.Getenv("HW_DMS_KAFKA_INSTANCE_ID")
	HW_DMS_KAFKA_TOPIC_NAME          = os.Getenv("HW_DMS_KAFKA_TOPIC_NAME")
	HW_DMS_KAFKA_CONSUMER_GROUP_NAME = os.Getenv("HW_DMS_KAFKA_CONSUMER_GROUP_NAME")

	HW_DMS_ROCKETMQ_INSTANCE_ID = os.Getenv("HW_DMS_ROCKETMQ_INSTANCE_ID")
	HW_DMS_ROCKETMQ_TOPIC_NAME  = os.Getenv("HW_DMS_ROCKETMQ_TOPIC_NAME")
	HW_DMS_ROCKETMQ_GROUP_NAME  = os.Getenv("HW_DMS_ROCKETMQ_GROUP_NAME")

	HW_SFS_TURBO_BACKUP_ID   = os.Getenv("HW_SFS_TURBO_BACKUP_ID")
	HW_SFS_FILE_SYSTEM_NAMES = os.Getenv("HW_SFS_FILE_SYSTEM_NAMES")

	HW_SMN_SUBSCRIBED_TOPIC_URN = os.Getenv("HW_SMN_SUBSCRIBED_TOPIC_URN")
)

// TestAccProviders is a static map containing only the main provider instance.
//
// Deprecated: Terraform Plugin SDK version 2 uses TestCase.ProviderFactories
// but supports this value in TestCase.Providers for backwards compatibility.
// In the future Providers: TestAccProviders will be changed to
// ProviderFactories: TestAccProviderFactories
var TestAccProviders map[string]*schema.Provider

// TestAccProviderFactories is a static map containing only the main provider instance
var TestAccProviderFactories map[string]func() (*schema.Provider, error)

// TestAccProvider is the "main" provider instance
var TestAccProvider *schema.Provider

func init() {
	TestAccProvider = huaweicloud.Provider()

	TestAccProviders = map[string]*schema.Provider{
		"huaweicloud": TestAccProvider,
	}

	TestAccProviderFactories = map[string]func() (*schema.Provider, error){
		"huaweicloud": func() (*schema.Provider, error) {
			return TestAccProvider, nil
		},
	}
}

func preCheckRequiredEnvVars(t *testing.T) {
	if HW_REGION_NAME == "" {
		t.Fatal("HW_REGION_NAME must be set for acceptance tests")
	}
}

// use this function to precheck langding zone services, such as Organizations and Identity Center
// lintignore:AT003
func TestAccPreCheckMultiAccount(t *testing.T) {
	if HW_MULTI_ACCOUNT_ENVIRONMENT == "" {
		t.Skip("This environment does not support multi-account tests")
	}
}

// Before the CAE environment resource is released, temporarily use this environment variable for acceptance tests.
// lintignore:AT003
func TestAccPreCheckCaeEnvironment(t *testing.T) {
	if HW_CAE_ENVIRONMENT_ID == "" {
		t.Skip("HW_CAE_ENVIRONMENT_ID must be set for the CAE acceptance test")
	}
}

// Before the CAE application resource is released, temporarily use this environment variable for acceptance tests.
// lintignore:AT003
func TestAccPreCheckCaeApplication(t *testing.T) {
	if HW_CAE_APPLICATION_ID == "" {
		t.Skip("HW_CAE_APPLICATION_ID must be set for the CAE acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCaeComponent(t *testing.T) {
	if HW_CAE_CODE_URL == "" || HW_CAE_CODE_AUTH_NAME == "" || HW_CAE_CODE_BRANCH == "" || HW_CAE_CODE_NAMESPACE == "" ||
		HW_CAE_ARTIFACT_NAMESPACE == "" || HW_CAE_BUILD_BASE_IMAGE == "" || HW_CAE_IMAGE_URL == "" {
		t.Skip("Skip the CAE acceptance tests.")
	}
}

// when this variable is set, the Organizations service should be enabled, and the organization info
// can be get by the API
// lintignore:AT003
func TestAccPreCheckOrganizationsOpen(t *testing.T) {
	if HW_ORGANIZATIONS_OPEN == "" {
		t.Skip("HW_ORGANIZATIONS_OPEN must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckOrganizationsAccountId(t *testing.T) {
	if HW_ORGANIZATIONS_ACCOUNT_ID == "" {
		t.Skip("HW_ORGANIZATIONS_ACCOUNT_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckOrganizationsAccountName(t *testing.T) {
	if HW_ORGANIZATIONS_ACCOUNT_NAME == "" {
		t.Skip("HW_ORGANIZATIONS_ACCOUNT_NAME must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckOrganizationsInviteAccountId(t *testing.T) {
	if HW_ORGANIZATIONS_INVITE_ACCOUNT_ID == "" {
		t.Skip("HW_ORGANIZATIONS_INVITE_ACCOUNT_ID must be set for acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckOrganizationsInvitationId(t *testing.T) {
	if HW_ORGANIZATIONS_INVITATION_ID == "" {
		t.Skip("HW_ORGANIZATIONS_INVITATION_ID must be set for acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckOrganizationsOrganizationalUnitId(t *testing.T) {
	if HW_ORGANIZATIONS_ORGANIZATIONAL_UNIT_ID == "" {
		t.Skip("HW_ORGANIZATIONS_ORGANIZATIONAL_UNIT_ID must be set for acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckRGCOrganization(t *testing.T) {
	if HW_RGC_ORGANIZATIONAL_UNIT_ID == "" || HW_RGC_ORGANIZATIONAL_UNIT_NAME == "" {
		t.Skip("HW_RGC_ORGANIZATIONAL_UNIT_ID and HW_RGC_ORGANIZATIONAL_UNIT_NAME must be set for acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckRGCBlueprint(t *testing.T) {
	if HW_RGC_BLUEPRINT_PRODUCT_ID == "" || HW_RGC_BLUEPRINT_PRODUCT_VERSION == "" {
		t.Skip("HW_RGC_BLUEPRINT_PRODUCT_ID and HW_RGC_BLUEPRINT_PRODUCT_VERSION must be set for acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckIdentityCenterAccountId(t *testing.T) {
	if HW_IDENTITY_CENTER_ACCOUNT_ID == "" {
		t.Skip("HW_IDENTITY_CENTER_ACCOUNT_ID must be set for acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckIdentityCenterIdentiyPolicyId(t *testing.T) {
	if HW_IDENTITY_CENTER_IDENTITY_POLICY_ID == "" || HW_IDENTITY_CENTER_IDENTITY_POLICY_ID_UPDATE == "" {
		t.Skip("HW_IDENTITY_CENTER_IDENTITY_POLICY_ID, HW_IDENTITY_CENTER_IDENTITY_POLICY_ID_UPDATE must be set for acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheck(t *testing.T) {
	// Do not run the test if this is a deprecated testing environment.
	if HW_DEPRECATED_ENVIRONMENT != "" {
		t.Skip("This environment only runs deprecated tests")
	}

	preCheckRequiredEnvVars(t)
}

// lintignore:AT003
func TestAccPrecheckDomainId(t *testing.T) {
	if HW_DOMAIN_ID == "" {
		t.Skip("HW_DOMAIN_ID must be set for acceptance tests")
	}
}

// lintignore:AT003
func TestAccPrecheckDomainName(t *testing.T) {
	if HW_DOMAIN_NAME == "" {
		t.Skip("HW_DOMAIN_NAME must be set for acceptance tests")
	}
}

// lintignore:AT003
func TestAccPrecheckCustomRegion(t *testing.T) {
	if HW_CUSTOM_REGION_NAME == "" {
		t.Skip("HW_CUSTOM_REGION_NAME must be set for acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckDeprecated(t *testing.T) {
	if HW_DEPRECATED_ENVIRONMENT == "" {
		t.Skip("This environment does not support deprecated tests")
	}

	preCheckRequiredEnvVars(t)
}

// lintignore:AT003
func TestAccPreCheckInternal(t *testing.T) {
	if HW_INTERNAL_USED == "" {
		t.Skip("HW_INTERNAL_USED must be set for internal acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckEpsID(t *testing.T) {
	// The environment variables in tests take HW_ENTERPRISE_PROJECT_ID_TEST instead of HW_ENTERPRISE_PROJECT_ID to
	// ensure that other data-resources that support enterprise projects query the default project without being
	// affected by this variable.
	if HW_ENTERPRISE_PROJECT_ID_TEST == "" {
		t.Skip("The environment variables does not support Enterprise Project ID for acc tests")
	}
}

// lintignore:AT003
func TestAccPreCheckMigrateEpsID(t *testing.T) {
	if HW_ENTERPRISE_PROJECT_ID_TEST == "" || HW_ENTERPRISE_MIGRATE_PROJECT_ID_TEST == "" {
		t.Skip("The environment variables does not support Migrate Enterprise Project ID for acc tests")
	}
}

// lintignore:AT003
func TestAccPreCheckUserId(t *testing.T) {
	if HW_USER_ID == "" {
		t.Skip("The environment variables does not support the user ID (HW_USER_ID) for acc tests")
	}
}

// lintignore:AT003
func TestAccPreCheckApigSubResourcesRelatedInfo(t *testing.T) {
	if HW_APIG_DEDICATED_INSTANCE_ID == "" {
		t.Skip("Before running APIG acceptance tests, please ensure the env 'HW_APIG_DEDICATED_INSTANCE_ID' has been configured")
	}
}

// lintignore:AT003
func TestAccPreCheckApigChannelRelatedInfo(t *testing.T) {
	if HW_APIG_DEDICATED_INSTANCE_USED_SUBNET_ID == "" {
		t.Skip("Before running APIG acceptance tests, please ensure the env 'HW_APIG_DEDICATED_INSTANCE_USED_SUBNET_ID' has been configured")
	}
}

// lintignore:AT003
func TestAccPreCheckSms(t *testing.T) {
	if HW_SMS_SOURCE_SERVER == "" {
		t.Skip("HW_SMS_SOURCE_SERVER must be set for SMS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckMrsCustom(t *testing.T) {
	if HW_MAPREDUCE_CUSTOM == "" {
		t.Skip("HW_MAPREDUCE_CUSTOM must be set for acceptance tests:custom type cluster of map reduce")
	}
}

// lintignore:AT003
func TestAccPreCheckMrsBootstrapScript(t *testing.T) {
	if HW_MAPREDUCE_BOOTSTRAP_SCRIPT == "" {
		t.Skip("HW_MAPREDUCE_BOOTSTRAP_SCRIPT must be set for acceptance tests: cluster of map reduce with bootstrap")
	}
}

// lintignore:AT003
func TestAccPreCheckFgsAgency(t *testing.T) {
	// The agency should be FunctionGraph and authorize these roles:
	// For the acceptance tests of the async invoke configuration:
	// + FunctionGraph FullAccess
	// + DIS Operator
	// + OBS Administrator
	// + SMN Administrator
	// For the acceptance tests of the function trigger and the application:
	// + LTS Administrator
	// + SWR Administrator
	if HW_FGS_AGENCY_NAME == "" {
		t.Skip("HW_FGS_AGENCY_NAME must be set for FGS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckFgsAppAgency(t *testing.T) {
	if HW_FGS_APP_AGENCY_NAME == "" {
		t.Skip("HW_FGS_APP_AGENCY_NAME must be set for FGS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckFgsGpuType(t *testing.T) {
	if HW_FGS_GPU_TYPE == "" {
		t.Skip("HW_FGS_GPU_TYPE must be set for FGS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckFgsDependencyLink(t *testing.T) {
	if HW_FGS_DEPENDENCY_OBS_LINK == "" {
		t.Skip("HW_FGS_DEPENDENCY_OBS_LINK must be set for FGS acceptance tests")
	}
}

// Deprecated
// lintignore:AT003
func TestAccPreCheckMaas(t *testing.T) {
	if HW_ACCESS_KEY == "" || HW_SECRET_KEY == "" || HW_SRC_ACCESS_KEY == "" || HW_SRC_SECRET_KEY == "" {
		t.Skip("HW_ACCESS_KEY, HW_SECRET_KEY, HW_SRC_ACCESS_KEY, and HW_SRC_SECRET_KEY  must be set for MAAS acceptance tests")
	}
}

func RandomAccResourceName() string {
	return fmt.Sprintf("tf_test_%s", acctest.RandString(5))
}

func RandomAccResourceNameWithDash() string {
	return fmt.Sprintf("tf-test-%s", acctest.RandString(5))
}

func RandomCidr() string {
	return fmt.Sprintf("172.16.%d.0/24", acctest.RandIntRange(0, 255))
}

func RandomCidrAndGatewayIp() (string, string) {
	seed := acctest.RandIntRange(0, 255)
	return fmt.Sprintf("172.16.%d.0/24", seed), fmt.Sprintf("172.16.%d.1", seed)
}

func RandomPassword(customChars ...string) string {
	var specialChars string
	if len(customChars) < 1 {
		specialChars = "~!@#%^*-_=+?"
	} else {
		specialChars = customChars[0]
	}
	return fmt.Sprintf("%s%s%s%d",
		acctest.RandStringFromCharSet(2, "ABCDEFGHIJKLMNOPQRSTUVWXZY"),
		acctest.RandStringFromCharSet(3, acctest.CharSetAlpha),
		acctest.RandStringFromCharSet(2, specialChars),
		acctest.RandIntRange(1000, 9999))
}

// lintignore:AT003
func TestAccPrecheckWafInstance(t *testing.T) {
	if HW_WAF_ENABLE_FLAG == "" {
		t.Skip("Skip the WAF acceptance tests.")
	}
}

// lintignore:AT003
// Used as a switch to enable the creation of cloud mode instance test cases.
func TestAccPreCheckWafCloudInstance(t *testing.T) {
	if HW_WAF_CLOUD_INSTANCE_FLAG == "" {
		t.Skip("Skip the WAF cloud acceptance tests.")
	}
}

// lintignore:AT003
// Used as a switch to enable international station test cases.
func TestAccPreCheckWafInternationalInstance(t *testing.T) {
	if HW_WAF_INTERNATIONAL_FLAG == "" {
		t.Skip("Skip the WAF international acceptance tests.")
	}
}

// lintignore:AT003
// Used as a switch to enable the WAF group test cases.
func TestAccPreCheckWafGroup(t *testing.T) {
	if HW_WAF_GROUP_FLAG == "" {
		t.Skip("Skip the WAF group acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckWafCertID(t *testing.T) {
	if HW_WAF_CERTIFICATE_ID == "" {
		t.Skip("HW_WAF_CERTIFICATE_ID must be set for this acceptance test.")
	}
}

// lintignore:AT003
func TestAccPreCheckWafType(t *testing.T) {
	if HW_WAF_TYPE == "" {
		t.Skip("HW_WAF_TYPE must be set for this acceptance test.")
	}
}

// lintignore:AT003
func TestAccPreCheckElbCertID(t *testing.T) {
	if HW_ELB_CERT_ID == "" {
		t.Skip("HW_ELB_CERT_ID must be set for this acceptance test.")
	}
}

// lintignore:AT003
func TestAccPreCheckCNADInstance(t *testing.T) {
	if HW_CNAD_ENABLE_FLAG == "" {
		t.Skip("Skip the CNAD acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckCNADProtectedObject(t *testing.T) {
	if HW_CNAD_PROJECT_OBJECT_ID == "" {
		t.Skip("Skipping test because HW_CNAD_PROJECT_OBJECT_ID is required for this acceptance test.")
	}
}

// lintignore:AT003
func TestAccPreCheckOmsInstance(t *testing.T) {
	if HW_OMS_ENABLE_FLAG == "" {
		t.Skip("Skip the OMS acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckAdminOnly(t *testing.T) {
	if HW_ADMIN == "" {
		t.Skip("Skipping test because it requires the admin privileges")
	}
}

// lintignore:AT003
func TestAccPreCheckIAMV5(t *testing.T) {
	if HW_IAM_V5 == "" {
		t.Skip("This environment does not support IAM v5 tests")
	}
}

// lintignore:AT003
func TestAccPreCheckRunnerPublicIP(t *testing.T) {
	if HW_RUNNER_PUBLIC_IP == "" {
		t.Skip("HW_RUNNER_PUBLIC_IP must be set for this acceptance test.")
	}
}

// lintignore:AT003
func TestAccPreCheckReplication(t *testing.T) {
	if HW_DEST_REGION == "" || HW_DEST_PROJECT_ID == "" {
		t.Skip("Skip the replication policy acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckDestRegion(t *testing.T) {
	if HW_DEST_REGION == "" {
		t.Skip("HW_DEST_REGION must be set for the acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckProjectId(t *testing.T) {
	if HW_DEST_PROJECT_ID_TEST == "" {
		t.Skip("Skipping test because it requires the test project id.")
	}
}

// lintignore:AT003
func TestAccPreCheckDestProjectIds(t *testing.T) {
	if HW_DEST_PROJECT_ID == "" || HW_DEST_PROJECT_ID_TEST == "" {
		t.Skip("HW_DEST_PROJECT_ID and HW_DEST_PROJECT_ID_TEST must be set for acceptance test.")
	}
}

// lintignore:AT003
func TestAccPreCheckProject(t *testing.T) {
	if HW_ENTERPRISE_PROJECT_ID_TEST != "" {
		t.Skip("This environment does not support project tests")
	}
}

// lintignore:AT003
func TestAccPreCheckOBS(t *testing.T) {
	if HW_ACCESS_KEY == "" || HW_SECRET_KEY == "" {
		t.Skip("HW_ACCESS_KEY and HW_SECRET_KEY must be set for OBS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckOBSDestinationBucket(t *testing.T) {
	if HW_OBS_DESTINATION_BUCKET == "" {
		t.Skip("HW_OBS_DESTINATION_BUCKET must be set for OBS destination tests")
	}
}

// lintignore:AT003
func TestAccPreCheckOBSUserDomainNames(t *testing.T) {
	if HW_OBS_USER_DOMAIN_NAME1 == "" || HW_OBS_USER_DOMAIN_NAME2 == "" {
		t.Skip("HW_OBS_USER_DOMAIN_NAME1 and HW_OBS_USER_DOMAIN_NAME2 must be set for OBS user domain name tests")
	}
}

// lintignore:AT003
func TestAccPreCheckOBSEndpoint(t *testing.T) {
	if HW_OBS_ENDPOINT == "" {
		t.Skip("HW_OBS_ENDPOINT must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckChargingMode(t *testing.T) {
	if HW_CHARGING_MODE != "prePaid" {
		t.Skip("This environment does not support prepaid tests")
	}
}

// lintignore:AT003
func TestAccPreCheckAvailabilityZoneGPSSD2(t *testing.T) {
	if HW_EVS_AVAILABILITY_ZONE_GPSSD2 == "" {
		t.Skip("If you want to change the QoS of a GPSSD2 type cloudvolume, you must specify an availability zone that supports GPSSD2 type under the current region")
	}
}

// lintignore:AT003
func TestAccPreCheckAvailabilityZoneESSD2(t *testing.T) {
	if HW_EVS_AVAILABILITY_ZONE_ESSD2 == "" {
		t.Skip("If you want to change the QoS of a ESSD2 type cloudvolume, you must specify an availability zone that supports ESSD2 type under the current region")
	}
}

// lintignore:AT003
func TestAccPreCheckHighCostAllow(t *testing.T) {
	if HW_HIGH_COST_ALLOW == "" {
		t.Skip("Do not allow expensive testing")
	}
}

// lintignore:AT003
func TestAccPreCheckCCMBaseCertificateImport(t *testing.T) {
	if HW_CCM_CERTIFICATE_CONTENT_PATH == "" || HW_CCM_CERTIFICATE_CHAIN_PATH == "" || HW_CCM_PRIVATE_KEY_PATH == "" {
		t.Skip("HW_CCM_CERTIFICATE_CONTENT_PATH, HW_CCM_CERTIFICATE_CHAIN_PATH and HW_CCM_PRIVATE_KEY_PATH " +
			"must be set for CCM certificate import tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckCCMEncCertificateImport(t *testing.T) {
	if HW_CCM_ENC_CERTIFICATE_PATH == "" || HW_CCM_ENC_PRIVATE_KEY_PATH == "" {
		t.Skip("HW_CCM_ENC_CERTIFICATE_PATH and HW_CCM_ENC_PRIVATE_KEY_PATH " +
			"must be set for CCM certificate enc import tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckCCMCertificatePush(t *testing.T) {
	if HW_CCM_CERTIFICATE_PROJECT == "" || HW_CCM_CERTIFICATE_PROJECT_UPDATED == "" {
		t.Skip("HW_CCM_CERTIFICATE_PROJECT and HW_CCM_CERTIFICATE_PROJECT_UPDATED must be set for CCM push certificate tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckSWRDomian(t *testing.T) {
	if HW_SWR_SHARING_ACCOUNT == "" {
		t.Skip("HW_SWR_SHARING_ACCOUNT must be set for swr domian tests, " +
			"the value of HW_SWR_SHARING_ACCOUNT should be another IAM user name")
	}
}

// lintignore:AT003
func TestAccPreCheckRAM(t *testing.T) {
	if HW_RAM_SHARE_ACCOUNT_ID == "" || HW_RAM_SHARE_RESOURCE_URN == "" {
		t.Skip("HW_RAM_SHARE_ACCOUNT_ID and HW_RAM_SHARE_RESOURCE_URN " +
			"must be set for create ram resource tests.")
	}

	if HW_RAM_SHARE_UPDATE_ACCOUNT_ID == "" || HW_RAM_SHARE_UPDATE_RESOURCE_URN == "" {
		t.Skip("HW_RAM_SHARE_UPDATE_ACCOUNT_ID and HW_RAM_SHARE_UPDATE_RESOURCE_URN" +
			" must be set for update ram resource tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckRAMResourceShare(t *testing.T) {
	if HW_RAM_SHARE_ACCOUNT_ID == "" || HW_RAM_SHARE_RESOURCE_URN == "" {
		t.Skip("HW_RAM_SHARE_ACCOUNT_ID and HW_RAM_SHARE_RESOURCE_URN " +
			"must be set for create ram resource tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckRAMEnableFlag(t *testing.T) {
	if HW_RAM_ENABLE_FLAG == "" {
		t.Skip("Skip the RAM acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckRAMShareInvitationId(t *testing.T) {
	if HW_RAM_SHARE_INVITATION_ID == "" {
		t.Skip("HW_RAM_SHARE_INVITATION_ID must be set for the acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckRAMShareId(t *testing.T) {
	if HW_RAM_SHARE_ID == "" {
		t.Skip("HW_RAM_SHARE_ID must be set for the acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckRAMSharedPrincipalsQueryFields(t *testing.T) {
	if HW_RAM_SHARE_ACCOUNT_ID == "" || HW_RAM_SHARE_RESOURCE_URN == "" {
		t.Skip("HW_RAM_SHARE_ACCOUNT_ID and HW_RAM_SHARE_RESOURCE_URN " +
			"must be set for RAM shared principals tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckRMSTargetIDForFGS(t *testing.T) {
	if HW_RMS_TARGET_ID_FOR_FGS == "" {
		t.Skip("HW_RMS_TARGET_ID_FOR_FGS must be set for the acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckRMSTargetIDForRFS(t *testing.T) {
	if HW_RMS_TARGET_ID_FOR_RFS == "" {
		t.Skip("HW_RMS_TARGET_ID_FOR_RFS must be set for the acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckRMSExcludedAccounts(t *testing.T) {
	if HW_RMS_EXCLUDED_ACCOUNT_1 == "" || HW_RMS_EXCLUDED_ACCOUNT_2 == "" {
		t.Skip("HW_RMS_EXCLUDED_ACCOUNT_1 and HW_RMS_EXCLUDED_ACCOUNT_2 must be set for the acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckRMSResourceRecorder(t *testing.T) {
	if HW_RMS_RESOURCE_RECORDER_CLOSED == "" {
		t.Skip("HW_RMS_RESOURCE_RECORDER_CLOSED must be set for the acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckDms(t *testing.T) {
	if HW_DMS_ENVIRONMENT == "" {
		t.Skip("This environment does not support DMS tests")
	}
}

// DLI authorization operations require that the object has administrative rights to the DLI service and has logged in
// to the DLI console.
// lintignore:AT003
func TestAccPreCheckDliAuthorizedUserConfigured(t *testing.T) {
	if HW_DLI_AUTHORIZED_USER_NAME == "" {
		t.Skip("HW_DLI_AUTHORIZED_USER_NAME must be set for DLI privilege acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckDliJarPath(t *testing.T) {
	if HW_DLI_FLINK_JAR_OBS_PATH == "" {
		t.Skip("HW_DLI_FLINK_JAR_OBS_PATH must be set for DLI Flink Jar job acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckDliDsAuthCss(t *testing.T) {
	if HW_DLI_DS_AUTH_CSS_OBS_PATH == "" {
		t.Skip("HW_DLI_DS_AUTH_CSS_OBS_PATH must be set for DLI datasource CSS Auth acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckDliDsAuthKafka(t *testing.T) {
	if HW_DLI_DS_AUTH_KAFKA_TRUST_OBS_PATH == "" || HW_DLI_DS_AUTH_KAFKA_KEY_OBS_PATH == "" {
		t.Skip("HW_DLI_DS_AUTH_KAFKA_TRUST_OBS_PATH,HW_DLI_DS_AUTH_KAFKA_KEY_OBS_PATH must be set for DLI datasource Kafka Auth acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckDliDsAuthKrb(t *testing.T) {
	if HW_DLI_DS_AUTH_KRB_CONF_OBS_PATH == "" || HW_DLI_DS_AUTH_KRB_TAB_OBS_PATH == "" {
		t.Skip("HW_DLI_DS_AUTH_KRB_CONF_OBS_PATH,HW_DLI_DS_AUTH_KRB_TAB_OBS_PATH must be set for DLI datasource Kafka Auth acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckDliAgency(t *testing.T) {
	if HW_DLI_AGENCY_FLAG == "" {
		t.Skip("HW_DLI_AGENCY_FLAG must be set for DLI datasource DLI agency acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckDliOwner(t *testing.T) {
	if HW_DLI_OWNER == "" {
		t.Skip("HW_DLI_OWNER must be set for DLI datasource DLI agency acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckDliElasticResourcePoolName(t *testing.T) {
	// Elastic resource pools for associating DLI datasource enhanced connection.
	// Therefore, two elastic resource pools are provided, one for initial binding and the other for updating binding.
	// Using commas (,) to separate two elastic resource pools.
	// The CU of the latter must be large and can be associated with multiple queues.
	// In the test case, the HW_DLI_SQL_QUEUE_NAME and HW_DLI_GENERAL_QUEUE_NAME belong to the latter resource pool.
	names := strings.Split(HW_DLI_ELASTIC_RESOURCE_POOL_NAMES, ",")
	if len(names) < 2 {
		t.Skip("Before running acceptance tests related to elastic resource pool, " +
			"please ensure +the 'HW_DLI_ELASTIC_RESOURCE_POOL_NAMES' has been configured")
	}

}

// lintignore:AT003
func TestAccPreCheckDliSQLQueueName(t *testing.T) {
	if HW_DLI_SQL_QUEUE_NAME == "" {
		t.Skip("HW_DLI_SQL_QUEUE_NAME must be set for DLI acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckDliGenaralQueueName(t *testing.T) {
	if HW_DLI_GENERAL_QUEUE_NAME == "" {
		t.Skip("HW_DLI_GENERAL_QUEUE_NAME must be set for DLI acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckDliUpdatedOwner(t *testing.T) {
	if HW_DLI_UPDATED_OWNER == "" {
		t.Skip("HW_DLI_UPDATED_OWNER must be set for DLI acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckDliFlinkVersion(t *testing.T) {
	if HW_DLI_FLINK_VERSION == "" {
		t.Skip("HW_DLI_FLINK_VERSION must be set for DLI acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckDliFlinkStreamGraph(t *testing.T) {
	if HW_DLI_FLINK_STREAM_GRAPH == "" {
		t.Skip("HW_DLI_FLINK_STREAM_GRAPH must be set for DLI acceptance tests.")
	}
}

// Since it takes at least one hour to execute the elastic resource pool resource test case,
// this variable is provided to distinguish the full test cases.
// lintignore:AT003
func TestAccPreCheckDliElasticResourcePool(t *testing.T) {
	if HW_DLI_ELASTIC_RESOURCE_POOL == "" {
		t.Skip("HW_DLI_ELASTIC_RESOURCE_POOL must be set for DLI acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckRepoTokenAuth(t *testing.T) {
	if HW_GITHUB_PERSONAL_TOKEN == "" {
		t.Skip("HW_GITHUB_PERSONAL_TOKEN must be set for token authorization test within ServiceStage.")
	}
}

// lintignore:AT003
func TestAccPreCheckRepoPwdAuth(t *testing.T) {
	if HW_DOMAIN_NAME == "" || HW_USER_NAME == "" || HW_GITHUB_REPO_PWD == "" {
		t.Skip("Repository configurations are not completed for acceptance test of password authorization.")
	}
}

// lintignore:AT003
func TestAccPreCheckComponent(t *testing.T) {
	if HW_DOMAIN_NAME == "" || HW_GITHUB_REPO_URL == "" || HW_OBS_STORAGE_URL == "" {
		t.Skip("Repository (package) configurations are not completed for acceptance test of component.")
	}
}

// lintignore:AT003
func TestAccPreCheckImageUrl(t *testing.T) {
	if HW_BUILD_IMAGE_URL == "" {
		t.Skip("HW_BUILD_IMAGE_URL must be set for acceptance test.")
	}
}

// lintignore:AT003
func TestAccPreCheckImageUrlUpdated(t *testing.T) {
	if HW_BUILD_IMAGE_URL_UPDATED == "" {
		t.Skip("HW_BUILD_IMAGE_URL_UPDATED must be set for acceptance test.")
	}
}

// lintignore:AT003
func TestAccPreCheckGaussDBMysqlInstanceId(t *testing.T) {
	if HW_GAUSSDB_MYSQL_INSTANCE_ID == "" {
		t.Skip("HW_GAUSSDB_MYSQL_INSTANCE_ID must be set for GaussDB MySQL acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckGaussDBMysqlNodeId(t *testing.T) {
	if HW_GAUSSDB_MYSQL_NODE_ID == "" {
		t.Skip("HW_GAUSSDB_MYSQL_NODE_ID must be set for GaussDB MySQL acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckGaussDBMysqlDatabaseName(t *testing.T) {
	if HW_GAUSSDB_MYSQL_DATABASE_NAME == "" {
		t.Skip("HW_GAUSSDB_MYSQL_DATABASE_NAME must be set for GaussDB MySQL acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckGaussDBMysqlTableName(t *testing.T) {
	if HW_GAUSSDB_MYSQL_TABLE_NAME == "" {
		t.Skip("HW_GAUSSDB_MYSQL_TABLE_NAME must be set for GaussDB MySQL acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckGaussDBMysqlInstanceConfigurationId(t *testing.T) {
	if HW_GAUSSDB_MYSQL_INSTANCE_CONFIGURATION_ID == "" {
		t.Skip("HW_GAUSSDB_MYSQL_INSTANCE_CONFIGURATION_ID must be set for GaussDB MySQL acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckGaussDBMysqlBackupBeginTime(t *testing.T) {
	if HW_GAUSSDB_MYSQL_BACKUP_BEGIN_TIME == "" {
		t.Skip("HW_GAUSSDB_MYSQL_BACKUP_BEGIN_TIME must be set for GaussDB MySQL acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckGaussDBMysqlBackupEndTime(t *testing.T) {
	if HW_GAUSSDB_MYSQL_BACKUP_END_TIME == "" {
		t.Skip("HW_GAUSSDB_MYSQL_BACKUP_END_TIME must be set for GaussDB MySQL acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckGaussDBMysqlJobId(t *testing.T) {
	if HW_GAUSSDB_MYSQL_JOB_ID == "" {
		t.Skip("HW_GAUSSDB_MYSQL_JOB_ID must be set for GaussDB MySQL acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckGaussDBMysqlTimeRange(t *testing.T) {
	if HW_GAUSSDB_MYSQL_START_TIME == "" || HW_GAUSSDB_MYSQL_END_TIME == "" {
		t.Skip("HW_GAUSSDB_MYSQL_START_TIME and HW_GAUSSDB_MYSQL_END_TIME must be set for GaussDB MySQL acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckGaussDBOpenGaussParameterTemplateId(t *testing.T) {
	if HW_GAUSSDB_OPENGAUSS_PARAMETER_TEMPLATE_ID == "" {
		t.Skip("HW_GAUSSDB_OPENGAUSS_PARAMETER_TEMPLATE_ID must be set for GaussDB OpenGauss acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckGaussDBOpenGaussJobId(t *testing.T) {
	if HW_GAUSSDB_OPENGAUSS_JOB_ID == "" {
		t.Skip("HW_GAUSSDB_OPENGAUSS_JOB_ID must be set for GaussDB OpenGauss acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckGaussDBOpenGaussInstanceId(t *testing.T) {
	if HW_GAUSSDB_OPENGAUSS_INSTANCE_ID == "" {
		t.Skip("HW_GAUSSDB_OPENGAUSS_INSTANCE_ID must be set for GaussDB OpenGauss acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckGaussDBOpenGaussTimeRange(t *testing.T) {
	if HW_GAUSSDB_OPENGAUSS_START_TIME == "" || HW_GAUSSDB_OPENGAUSS_END_TIME == "" {
		t.Skip("HW_GAUSSDB_OPENGAUSS_START_TIME and HW_GAUSSDB_OPENGAUSS_END_TIME must be set for GaussDB " +
			"OpenGauss acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckVODWatermark(t *testing.T) {
	if HW_VOD_WATERMARK_FILE == "" {
		t.Skip("HW_VOD_WATERMARK_FILE must be set for VOD watermark template acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckVODMediaAsset(t *testing.T) {
	if HW_VOD_MEDIA_ASSET_FILE == "" {
		t.Skip("HW_VOD_MEDIA_ASSET_FILE must be set for VOD media asset acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckPwdAuth(t *testing.T) {
	if HW_MEETING_ACCOUNT_NAME == "" || HW_MEETING_ACCOUNT_PASSWORD == "" {
		t.Skip("The account name (HW_MEETING_ACCOUNT_NAME) or password (HW_MEETING_ACCOUNT_PASSWORD) is not " +
			"completed for acceptance test of conference.")
	}
}

// lintignore:AT003
func TestAccPreCheckAppAuth(t *testing.T) {
	if HW_MEETING_APP_ID == "" || HW_MEETING_APP_KEY == "" || HW_MEETING_USER_ID == "" {
		t.Skip("The app ID (HW_MEETING_APP_ID), app KEY (HW_MEETING_APP_KEY) or user ID (HW_MEETING_USER_ID) is not " +
			"completed for acceptance test of conference.")
	}
}

// lintignore:AT003
func TestAccPreCheckMeetingRoom(t *testing.T) {
	if HW_MEETING_ROOM_ID == "" {
		t.Skip("The vmr ID (HW_MEETING_ROOM_ID) is not completed for acceptance test of conference.")
	}
}

// lintignore:AT003
func TestAccPreCheckParticipants(t *testing.T) {
	if HW_CHAIR_EMAIL == "" || HW_GUEST_EMAIL == "" {
		t.Skip("The chair (HW_CHAIR_EMAIL) or guest (HW_GUEST_EMAIL) mailbox is not completed for acceptance test of " +
			"conference.")
	}
}

// lintignore:AT003
func TestAccPreCheckAadForwardRule(t *testing.T) {
	if HW_AAD_INSTANCE_ID == "" || HW_AAD_IP_ADDRESS == "" {
		t.Skip("The instance information is not completed for AAD rule acceptance test.")
	}
}

// lintignore:AT003
func TestAccPreCheckCCMCertificateName(t *testing.T) {
	if HW_CCM_CERTIFICATE_NAME == "" {
		t.Skip("HW_CCM_CERTIFICATE_NAME must be set for CCM SSL acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckCCMSSLCertificateId(t *testing.T) {
	if HW_CCM_SSL_CERTIFICATE_ID == "" {
		t.Skip("HW_CCM_SSL_CERTIFICATE_ID must be set for CCM SSL acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckCCMPrivateCaID(t *testing.T) {
	if HW_CCM_PRIVATE_CA_ID == "" {
		t.Skip("HW_CCM_PRIVATE_CA_ID must be set for CCM acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckCCMEnableFlag(t *testing.T) {
	if HW_CCM_ENABLE_FLAG == "" {
		t.Skip("Skip the CCM acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckKms(t *testing.T) {
	if HW_KMS_ENVIRONMENT == "" {
		t.Skip("This environment does not support KMS tests")
	}
}

// lintignore:AT003
func TestAccPreCheckKmsHsmClusterId(t *testing.T) {
	if HW_KMS_HSM_CLUSTER_ID == "" {
		t.Skip("HW_KMS_HSM_CLUSTER_ID must be set for KMS dedicated keystore acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckKmsImportToken(t *testing.T) {
	if HW_KMS_IMPORT_TOKEN == "" {
		t.Skip("HW_KMS_IMPORT_TOKEN must be set for KMS key material acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckKmsKeyID(t *testing.T) {
	if HW_KMS_KEY_ID == "" {
		t.Skip("HW_KMS_KEY_ID must be set for KMS key material acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckKmsKeyMaterial(t *testing.T) {
	if HW_KMS_KEY_MATERIAL == "" {
		t.Skip("HW_KMS_KEY_MATERIAL must be set for KMS key material acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckKmsKeyPrivateKey(t *testing.T) {
	if HW_KMS_KEY_PRIVATE_KEY == "" {
		t.Skip("HW_KMS_KEY_PRIVATE_KEY must be set for KMS key material acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckProjectID(t *testing.T) {
	if HW_PROJECT_ID == "" {
		t.Skip("HW_PROJECT_ID must be set for acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCCProjectID(t *testing.T) {
	if HW_PROJECT_ID_1 == "" || HW_PROJECT_ID_2 == "" || HW_PROJECT_ID_3 == "" {
		t.Skip("HW_PROJECT_ID_1, HW_PROJECT_ID_2, HW_PROJECT_ID_3 must be set for this acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCCRegionName(t *testing.T) {
	if HW_REGION_NAME_1 == "" || HW_REGION_NAME_2 == "" || HW_REGION_NAME_3 == "" {
		t.Skip("HW_REGION_NAME_1, HW_REGION_NAME_2, HW_REGION_NAME_3 must be set for this acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCCConnectionRouteProjectID(t *testing.T) {
	if HW_PROJECT_ID_1 == "" || HW_PROJECT_ID_2 == "" {
		t.Skip("HW_PROJECT_ID_1, HW_PROJECT_ID_2 must be set for this acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCCConnectionRouteRegionName(t *testing.T) {
	if HW_REGION_NAME_1 == "" || HW_REGION_NAME_2 == "" {
		t.Skip("HW_REGION_NAME_1, HW_REGION_NAME_2 must be set for this acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCtsTimeRange(t *testing.T) {
	if HW_CTS_START_TIME == "" || HW_CTS_END_TIME == "" {
		t.Skip("HW_CTS_START_TIME and HW_CTS_END_TIME must be set for CTS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckWorkspaceADDomainNames(t *testing.T) {
	if len(strings.Split(HW_WORKSPACE_AD_DOMAIN_NAMES, ",")) != 2 {
		t.Skip(`The Workspace AD service need domain name configurations for both master and standby servers, plesse config them in the
HW_WORKSPACE_AD_DOMAIN_NAMES environment variable, separated by a comma (,).`)
	}
}

// lintignore:AT003
func TestAccPreCheckWorkspaceAD(t *testing.T) {
	TestAccPreCheckWorkspaceADDomainNames(t)

	if len(strings.Split(HW_WORKSPACE_AD_DOMAIN_IPS, ",")) != 2 {
		t.Skip(`The Workspace AD service need IP address configurations for both master and standby servers, plesse config them in the
HW_WORKSPACE_AD_DOMAIN_IPS environment variable, separated by a comma (,).`)
	}
	if HW_WORKSPACE_AD_SERVER_ACCOUNT == "" || HW_WORKSPACE_AD_SERVER_PWD == "" || HW_WORKSPACE_AD_VPC_ID == "" || HW_WORKSPACE_AD_NETWORK_ID == "" {
		t.Skip(`The configuration of AD server is not completed for Workspace service acceptance test, please check your inputs for
HW_WORKSPACE_AD_SERVER_ACCOUNT, HW_WORKSPACE_AD_SERVER_PWD, HW_WORKSPACE_AD_VPC_ID and HW_WORKSPACE_AD_NETWORK_ID.`)
	}
}

// lintignore:AT003
func TestAccPreCheckWorkspaceInternetAccessPort(t *testing.T) {
	if HW_WORKSPACE_INTERNET_ACCESS_PORT == "" {
		t.Skip("HW_WORKSPACE_INTERNET_ACCESS_PORT must be set for Workspace service acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckWorkspaceAppServerGroup(t *testing.T) {
	if HW_WORKSPACE_APP_SERVER_GROUP_FLAVOR_ID == "" || HW_WORKSPACE_APP_SERVER_GROUP_IMAGE_ID == "" ||
		HW_WORKSPACE_APP_SERVER_GROUP_IMAGE_PRODUCT_ID == "" {
		t.Skip("Workspace APP server group acceptance test missing configuration parameters.")
	}
}

// lintignore:AT003
func TestAccPreCheckWorkspaceAppImageSpecCode(t *testing.T) {
	if HW_WORKSPACE_APP_SERVER_GROUP_IMAGE_SPEC_CODE == "" {
		t.Skip("HW_WORKSPACE_APP_SERVER_GROUP_IMAGE_SPEC_CODE must be set for Workspace APP acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPrecheckWorkspaceUserNames(t *testing.T) {
	if len(strings.Split(HW_WORKSPACE_USER_NAMES, ",")) < 1 {
		t.Skip("At least one of user must be configured in the HW_WORKSPACE_USER_NAMES, and separated by commas")
	}
}

// lintignore:AT003
func TestAccPreCheckWorkspaceOUName(t *testing.T) {
	if HW_WORKSPACE_OU_NAME == "" {
		t.Skip("HW_WORKSPACE_OU_NAME must be set for Workspace service acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckWorkspaceAppFileName(t *testing.T) {
	if HW_WORKSPACE_APP_FILE_NAME == "" {
		t.Skip("HW_WORKSPACE_APP_FILE_NAME must be set for Workspace service acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckER(t *testing.T) {
	if HW_ER_TEST_ON == "" {
		t.Skip("Skip all ER acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckRfArchives(t *testing.T) {
	if HW_RF_TEMPLATE_ARCHIVE_NO_VARS_URI == "" || HW_RF_TEMPLATE_ARCHIVE_URI == "" ||
		HW_RF_VARIABLES_ARCHIVE_URI == "" {
		t.Skip("Skip the archive URI parameters acceptance test for RF resource stack.")
	}
}

// lintignore:AT003
func TestAccPreCheckDcDirectConnection(t *testing.T) {
	if HW_DC_DIRECT_CONNECT_ID == "" {
		t.Skip("Skip the interface acceptance test because of the direct connection ID is missing.")
	}
}

// lintignore:AT003
func TestAccPreCheckDcHostedConnection(t *testing.T) {
	if HW_DC_RESOURCE_TENANT_ID == "" || HW_DC_HOSTTING_ID == "" {
		t.Skip("HW_DC_RESOURCE_TENANT_ID, HW_DC_HOSTTING_ID must be set for this acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDcResourceTenant(t *testing.T) {
	if HW_DC_RESOURCE_TENANT_ID == "" {
		t.Skip("HW_DC_RESOURCE_TENANT_ID must be set for this acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckTargetTenantDcVGW(t *testing.T) {
	if HW_DC_TARGET_TENANT_VGW_ID == "" {
		t.Skip("HW_DC_TARGET_TENANT_VGW_ID must be set for this acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDCVirtualInterfaceID(t *testing.T) {
	if HW_DC_VIRTUAL_INTERFACE_ID == "" {
		t.Skip("HW_DC_VIRTUAL_INTERFACE_ID must be set for this acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCesTimeRange(t *testing.T) {
	if HW_CES_START_TIME == "" || HW_CES_END_TIME == "" {
		t.Skip("HW_CES_START_TIME and HW_CES_END_TIME must be set for CES acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCfw(t *testing.T) {
	if HW_CFW_INSTANCE_ID == "" {
		t.Skip("HW_CFW_INSTANCE_ID must be set for CFW acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCfwTimeRange(t *testing.T) {
	if HW_CFW_START_TIME == "" || HW_CFW_END_TIME == "" {
		t.Skip("HW_CFW_START_TIME and HW_CFW_END_TIME must be set for CFW acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCfwEastWestFirewall(t *testing.T) {
	if HW_CFW_EAST_WEST_FIREWALL == "" {
		t.Skip("HW_CFW_EAST_WEST_FIREWALL must be set for CFW east-west firewall acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCfwPredefinedServiceGroup(t *testing.T) {
	if HW_CFW_PREDEFINED_SERVICE_GROUP1 == "" || HW_CFW_PREDEFINED_SERVICE_GROUP2 == "" {
		t.Skip("HW_CFW_PREDEFINED_SERVICE_GROUP1 and HW_CFW_PREDEFINED_SERVICE_GROUP2 must be set for CFW ACL rule acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCfwPredefinedAddressGroup(t *testing.T) {
	if HW_CFW_PREDEFINED_ADDRESS_GROUP1 == "" || HW_CFW_PREDEFINED_ADDRESS_GROUP2 == "" {
		t.Skip("HW_CFW_PREDEFINED_ADDRESS_GROUP1 and HW_CFW_PREDEFINED_ADDRESS_GROUP2 must be set for CFW ACL rule acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCfwIpsCustomRule(t *testing.T) {
	if HW_CFW_IPS_CUSTOM_RULE == "" {
		t.Skip("HW_CFW_IPS_CUSTOM_RULE must be set for CFW IPS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckWorkloadType(t *testing.T) {
	if HW_WORKLOAD_TYPE == "" {
		t.Skip("HW_WORKLOAD_TYPE must be set for SWR image trigger acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckWorkloadName(t *testing.T) {
	if HW_WORKLOAD_NAME == "" {
		t.Skip("HW_WORKLOAD_NAME must be set for SWR image trigger acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCceClusterId(t *testing.T) {
	if HW_CCE_CLUSTER_ID == "" {
		t.Skip("HW_CCE_CLUSTER_ID must be set for this acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCceChartPath(t *testing.T) {
	// HW_CCE_CHART_PATH is the absolute path of the chart package
	if HW_CCE_CHART_PATH == "" {
		t.Skip("HW_CCE_CHART_PATH must be set for CCE chart acceptance tests, " +
			"the value is the absolute path of the chart package")
	}
}

// lintignore:AT003
func TestAccPreCheckWorkloadNameSpace(t *testing.T) {
	if HW_WORKLOAD_NAMESPACE == "" {
		t.Skip("HW_WORKLOAD_NAMESPACE must be set for SWR image trigger acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckSwrTargetRegion(t *testing.T) {
	if HW_SWR_TARGET_REGION == "" {
		t.Skip("HW_SWR_TARGET_REGION must be set for SWR image auto sync tests")
	}
}

// lintignore:AT003
func TestAccPreCheckSwrTargetOrigination(t *testing.T) {
	if HW_SWR_TARGET_ORGANIZATION == "" {
		t.Skip("HW_SWR_TARGET_ORGANIZATION must be set for SWR image auto sync tests")
	}
}

// lintignore:AT003
func TestAccPreCheckSwrOrigination(t *testing.T) {
	if HW_SWR_ORGANIZATION == "" {
		t.Skip("HW_SWR_ORGANIZATION must be set for SWR image tags tests")
	}
}

// lintignore:AT003
func TestAccPreCheckSwrRepository(t *testing.T) {
	if HW_SWR_REPOSITORY == "" {
		t.Skip("HW_SWR_REPOSITORY must be set for SWR image tags tests")
	}
}

// lintignore:AT003
func TestAccPreCheckImsVaultId(t *testing.T) {
	if HW_IMS_VAULT_ID == "" {
		t.Skip("HW_IMS_VAULT_ID must be set for IMS whole image tests")
	}
}

// lintignore:AT003
func TestAccPreCheckImsBackupId(t *testing.T) {
	if HW_IMS_BACKUP_ID == "" {
		t.Skip("HW_IMS_BACKUP_ID must be set for IMS whole image with CBR backup id")
	}
}

// lintignore:AT003
func TestAccPreCheckImsImageUrl(t *testing.T) {
	if HW_IMS_IMAGE_URL == "" {
		t.Skip("HW_IMS_IMAGE_URL must be set for IMS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckAcceptBackup(t *testing.T) {
	if HW_SHARED_BACKUP_ID == "" {
		t.Skip("HW_SHARED_BACKUP_ID must be set for CBR backup share acceptance")
	}
}

// lintignore:AT003
func TestAccPreCheckSourceImage(t *testing.T) {
	if HW_IMAGE_SHARE_SOURCE_IMAGE_ID == "" {
		t.Skip("Skip the interface acceptance test because of the source image ID is missing.")
	}
}

// lintignore:AT003
func TestAccPreCheckSecMasterWorkspaceID(t *testing.T) {
	if HW_SECMASTER_WORKSPACE_ID == "" {
		t.Skip("HW_SECMASTER_WORKSPACE_ID must be set for SecMaster acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckSecMasterMetricID(t *testing.T) {
	if HW_SECMASTER_METRIC_ID == "" {
		t.Skip("HW_SECMASTER_METRIC_ID must be set for SecMaster acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckSecMasterPipelineID(t *testing.T) {
	if HW_SECMASTER_PIPELINE_ID == "" {
		t.Skip("HW_SECMASTER_PIPELINE_ID must be set for SecMaster acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckSecMasterInstanceID(t *testing.T) {
	if HW_SECMASTER_INSTANCE_ID == "" {
		t.Skip("HW_SECMASTER_INSTANCE_ID must be set for SecMaster acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckSecMasterIndicatorTypeID(t *testing.T) {
	if HW_SECMASTER_INDICATOR_TYPE_ID == "" {
		t.Skip("HW_SECMASTER_INDICATOR_TYPE_ID must be set for SecMaster acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckSecMaster(t *testing.T) {
	if HW_SECMASTER_WORKSPACE_ID == "" || HW_SECMASTER_INDICATOR_TYPE_ID == "" ||
		HW_SECMASTER_INDICATOR_TYPE_ID_UPDATE == "" {
		t.Skip("HW_SECMASTER_WORKSPACE_ID, HW_SECMASTER_INDICATOR_TYPE_ID and HW_SECMASTER_INDICATOR_TYPE_ID_UPDATE" +
			" must be set for SecMaster acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCcePartitionAz(t *testing.T) {
	if HW_CCE_PARTITION_AZ == "" {
		t.Skip("Skip the interface acceptance test because of the cce partition az is missing.")
	}
}

// lintignore:AT003
func TestAccPreCheckUpdateCertificateContent(t *testing.T) {
	if HW_CERTIFICATE_CONTENT == "" || HW_CERTIFICATE_CONTENT_UPDATE == "" {
		t.Skip("HW_CERTIFICATE_CONTENT, HW_CERTIFICATE_CONTENT_UPDATE must be set for this test")
	}
}

// lintignore:AT003
func TestAccPreCheckCertificateWithoutRootCA(t *testing.T) {
	if HW_CERTIFICATE_CONTENT == "" || HW_CERTIFICATE_PRIVATE_KEY == "" ||
		HW_NEW_CERTIFICATE_CONTENT == "" || HW_NEW_CERTIFICATE_PRIVATE_KEY == "" {
		t.Skip("HW_CERTIFICATE_CONTENT, HW_CERTIFICATE_PRIVATE_KEY, HW_NEW_CERTIFICATE_CONTENT and " +
			"HW_NEW_CERTIFICATE_PRIVATE_KEY must be set for simple acceptance tests of SSL certificate resource")
	}
}

// lintignore:AT003
func TestAccPreCheckCertificateFull(t *testing.T) {
	TestAccPreCheckCertificateWithoutRootCA(t)
	if HW_CERTIFICATE_ROOT_CA == "" || HW_NEW_CERTIFICATE_ROOT_CA == "" {
		t.Skip("HW_CERTIFICATE_ROOT_CA and HW_NEW_CERTIFICATE_ROOT_CA must be set for root CA validation")
	}
}

// lintignore:AT003
func TestAccPreCheckGMCertificate(t *testing.T) {
	if HW_GM_CERTIFICATE_CONTENT == "" || HW_GM_CERTIFICATE_PRIVATE_KEY == "" ||
		HW_GM_ENC_CERTIFICATE_CONTENT == "" || HW_GM_ENC_CERTIFICATE_PRIVATE_KEY == "" ||
		HW_GM_CERTIFICATE_CHAIN == "" ||
		HW_NEW_GM_CERTIFICATE_CONTENT == "" || HW_NEW_GM_CERTIFICATE_PRIVATE_KEY == "" ||
		HW_NEW_GM_ENC_CERTIFICATE_CONTENT == "" || HW_NEW_GM_ENC_CERTIFICATE_PRIVATE_KEY == "" ||
		HW_NEW_GM_CERTIFICATE_CHAIN == "" {
		t.Skip("HW_GM_CERTIFICATE_CONTENT, HW_GM_CERTIFICATE_PRIVATE_KEY, HW_GM_ENC_CERTIFICATE_CONTENT," +
			" HW_GM_ENC_CERTIFICATE_PRIVATE_KEY, HW_GM_CERTIFICATE_CHAIN, HW_NEW_GM_CERTIFICATE_CONTENT," +
			" HW_NEW_GM_CERTIFICATE_PRIVATE_KEY, HW_NEW_GM_ENC_CERTIFICATE_CONTENT," +
			" HW_NEW_GM_ENC_CERTIFICATE_PRIVATE_KEY, HW_NEW_GM_CERTIFICATE_CHAIN must be set for GM certificate")
	}
}

// lintignore:AT003
func TestAccPreCheckCodeArtsDeployResourcePoolID(t *testing.T) {
	if HW_CODEARTS_RESOURCE_POOL_ID == "" {
		t.Skip("HW_CODEARTS_RESOURCE_POOL_ID must be set for this acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCodeArtsDeploymentTaskID(t *testing.T) {
	if HW_CODEARTS_DEPLOYMENT_TASK_ID == "" {
		t.Skip("HW_CODEARTS_DEPLOYMENT_TASK_ID must be set for this acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCodeArtsEnableFlag(t *testing.T) {
	if HW_CODEARTS_ENABLE_FLAG == "" {
		t.Skip("Skip the CodeArts acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckCodeArtsPublicIPAddress(t *testing.T) {
	if HW_CODEARTS_PUBLIC_IP_ADDRESS == "" {
		t.Skip("HW_CODEARTS_PUBLIC_IP_ADDRESS must be set for this acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCodeArtsSshCredentialID(t *testing.T) {
	if HW_CODEARTS_SSH_CREDENTIAL_ID == "" {
		t.Skip("HW_CODEARTS_SSH_CREDENTIAL_ID must be set for this acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckModelArtsHasSubscribeModel(t *testing.T) {
	if HW_MODELARTS_HAS_SUBSCRIBE_MODEL == "" {
		t.Skip("Subscribe two free models from market and set HW_MODELARTS_HAS_SUBSCRIBE_MODEL" +
			" for modelarts service acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckModelartsUserLoginPassword(t *testing.T) {
	if HW_MODELARTS_USER_LOGIN_PASSWORD == "" {
		t.Skip("HW_MODELARTS_USER_LOGIN_PASSWORD must be set for modelarts privilege resource pool acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckModelartsDevServer(t *testing.T) {
	if HW_MODELARTS_DEVSERVER_FLAVOR == "" || HW_MODELARTS_DEVSERVER_IMAGE_ID == "" {
		t.Skip("HW_MODELARTS_DEVSERVER_FLAVOR and HW_MODELARTS_DEVSERVER_IMAGE_ID must be set for ModelArts DevServer acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckEG(t *testing.T) {
	if HW_EG_TEST_ON == "" {
		t.Skip("Skip all sub tests of the EG service.")
	}
}

// lintignore:AT003
func TestAccPreCheckEgChannelId(t *testing.T) {
	if HW_EG_CHANNEL_ID == "" {
		t.Skip("The sub-resource acceptance test of the EG channel must set 'HW_EG_CHANNEL_ID'")
	}
}

// lintignore:AT003
func TestAccPreCheckEgAgencyName(t *testing.T) {
	if HW_EG_AGENCY_NAME == "" {
		t.Skip("HW_EG_AGENCY_NAME must be set for resource creation of the EG connection")
	}
}

// lintignore:AT003
func TestAccPreCheckLtsAomAccess(t *testing.T) {
	if HW_LTS_CLUSTER_ID == "" || HW_LTS_CLUSTER_NAME == "" {
		t.Skip("HW_LTS_CLUSTER_ID and HW_LTS_CLUSTER_NAME must be set for LTS AOM access acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckLtsAomAccessUpdate(t *testing.T) {
	if HW_LTS_CLUSTER_ID_ANOTHER == "" || HW_LTS_CLUSTER_NAME_ANOTHER == "" {
		t.Skip("HW_LTS_CLUSTER_ID_ANOTHER and HW_LTS_CLUSTER_NAME_ANOTHER must be set for LTS AOM access" +
			" acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckLtsStructConfigCustom(t *testing.T) {
	if HW_LTS_STRUCT_CONFIG_TEMPLATE_ID == "" || HW_LTS_STRUCT_CONFIG_TEMPLATE_NAME == "" {
		t.Skip("HW_LTS_STRUCT_CONFIG_TEMPLATE_ID and HW_LTS_STRUCT_CONFIG_TEMPLATE_NAME must be" +
			" set for LTS struct config custom acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckLiveStreamingDomainName(t *testing.T) {
	if HW_LIVE_STREAMING_DOMAIN_NAME == "" {
		t.Skip("HW_LIVE_STREAMING_DOMAIN_NAME must be set for Live acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckLiveIngestRTMPDomainName(t *testing.T) {
	if HW_LIVE_INGEST_RTMP_DOMAIN_NAME == "" {
		t.Skip("HW_LIVE_INGEST_RTMP_DOMAIN_NAME must be set for Live acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckLiveIngestSRTDomainName(t *testing.T) {
	if HW_LIVE_INGEST_SRT_DOMAIN_NAME == "" {
		t.Skip("HW_LIVE_INGEST_SRT_DOMAIN_NAME must be set for Live acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckLiveIngestDomainName(t *testing.T) {
	if HW_LIVE_INGEST_DOMAIN_NAME == "" {
		t.Skip("HW_LIVE_INGEST_DOMAIN_NAME must be set for Live acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckLiveTranscodingTemplateID(t *testing.T) {
	if HW_LIVE_TRANSCODING_TEPLATE_ID == "" {
		t.Skip("HW_LIVE_TRANSCODING_TEPLATE_ID must be set for Live acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckLiveTranscodingTemplateAnotherID(t *testing.T) {
	if HW_LIVE_TRANSCODING_TEPLATE_ANOTHER_ID == "" {
		t.Skip("HW_LIVE_TRANSCODING_TEPLATE_ANOTHER_ID must be set for Live acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckLiveTLSCert(t *testing.T) {
	if HW_LIVE_HTTPS_TLS_CERT_BODY_PATH == "" || HW_LIVE_HTTPS_TLS_CERT_KEY_PATH == "" {
		t.Skip("HW_LIVE_HTTPS_TLS_CERT_BODY_PATH and HW_LIVE_HTTPS_TLS_CERT_KEY_PATH must be set for Live acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckLtsEnableFlag(t *testing.T) {
	if HW_LTS_ENABLE_FLAG == "" {
		t.Skip("Skip the LTS acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckLTSCrossAccountAccess(t *testing.T) {
	if HW_LTS_AGENCY_STREAM_NAME == "" || HW_LTS_AGENCY_STREAM_ID == "" || HW_LTS_AGENCY_GROUP_NAME == "" ||
		HW_LTS_AGENCY_GROUP_ID == "" || HW_LTS_AGENCY_PROJECT_ID == "" ||
		HW_LTS_AGENCY_DOMAIN_NAME == "" || HW_LTS_AGENCY_NAME == "" {
		t.Skip("The delegator account config of HW_LTS_AGENCY_STREAM_NAME, HW_LTS_AGENCY_STREAM_ID, HW_LTS_AGENCY_GROUP_NAME," +
			" HW_LTS_AGENCY_GROUP_ID, HW_LTS_AGENCY_PROJECT_ID, HW_LTS_AGENCY_DOMAIN_NAME and HW_LTS_AGENCY_NAME " +
			"must be set for the acceptance test")
	}

	if HW_LTS_LOG_STREAM_NAME == "" || HW_LTS_LOG_STREAM_ID == "" ||
		HW_LTS_LOG_GROUP_NAME == "" || HW_LTS_LOG_GROUP_ID == "" {
		t.Skip("The delegatee account config of HW_LTS_LOG_STREAM_NAME, HW_LTS_LOG_STREAM_ID, HW_LTS_LOG_GROUP_NAME" +
			" and HW_LTS_LOG_GROUP_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckLTSCCEAccess(t *testing.T) {
	if HW_LTS_CCE_CLUSTER_ID == "" {
		t.Skip("The cce access config of HW_LTS_CCE_CLUSTER_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckLTSHostGroup(t *testing.T) {
	// LTS host groups support updating associated hosts, so at least one host is required to create and update each host group.
	hostIds := strings.Split(HW_LTS_HOST_IDS, ",")
	if len(hostIds) < 2 {
		t.Skip("The length of HW_LTS_HOST_IDS must be at least 2 for the host group acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckLTSLogConvergeBaseConfig(t *testing.T) {
	if HW_LTS_LOG_CONVERGE_ORGANIZATION_ID == "" || HW_LTS_LOG_CONVERGE_MANAGEMENT_ACCOUNT_ID == "" ||
		HW_LTS_LOG_CONVERGE_MEMBER_ACCOUNT_ID == "" {
		t.Skip("The cce access config of HW_LTS_LOG_CONVERGE_ORGANIZATION_ID, HW_LTS_LOG_CONVERGE_MANAGEMENT_ACCOUNT_ID, " +
			"HW_LTS_LOG_CONVERGE_MEMBER_ACCOUNT_ID must be set for the log converge configuration acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckLTSLogConvergeMappingConfig(t *testing.T) {
	if HW_LTS_LOG_CONVERGE_SOURCE_LOG_GROUP_ID == "" || HW_LTS_LOG_CONVERGE_SOURCE_LOG_STREAM_ID == "" {
		t.Skip("The environment variables of HW_LTS_LOG_CONVERGE_SOURCE_LOG_GROUP_ID and HW_LTS_LOG_CONVERGE_SOURCE_LOG_STREAM_ID " +
			"must be set for the log converge configuration acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckAomSubApplicationId(t *testing.T) {
	if HW_AOM_SUB_APPLICATION_ID == "" {
		t.Skip("HW_AOM_SUB_APPLICATION_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckMultiAccountAggregationRuleEnable(t *testing.T) {
	if HW_AOM_MULTI_ACCOUNT_AGGREGATION_RULE_ENABLE == "" {
		t.Skip("HW_AOM_MULTI_ACCOUNT_AGGREGATION_RULE_ENABLE must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCocInstanceID(t *testing.T) {
	if HW_COC_INSTANCE_ID == "" {
		t.Skip("HW_COC_INSTANCE_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPrecheckKooGallery(t *testing.T) {
	if HW_KOOGALLERY_ASSET == "" {
		t.Skip("Skip the KooGallery acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckCCINamespace(t *testing.T) {
	if HW_CCI_NAMESPACE == "" {
		t.Skip("This environment does not support CCI Namespace tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCDN(t *testing.T) {
	if HW_CDN_DOMAIN_NAME == "" {
		t.Skip("This environment does not support CDN tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCertCDN(t *testing.T) {
	if HW_CDN_CERT_DOMAIN_NAME == "" {
		t.Skip("HW_CDN_CERT_DOMAIN_NAME must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCDNURL(t *testing.T) {
	if HW_CDN_DOMAIN_URL == "" {
		t.Skip("HW_CDN_DOMAIN_URL must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCERT(t *testing.T) {
	if HW_CDN_CERT_PATH == "" || HW_CDN_PRIVATE_KEY_PATH == "" {
		t.Skip("This environment does not support CDN certificate tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCDNDomainCertificates(t *testing.T) {
	if HW_CDN_ENABLE_FLAG == "" {
		t.Skip("Skip the CDN acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPreCheckCCGlobalGateway(t *testing.T) {
	if HW_CC_GLOBAL_GATEWAY_ID == "" {
		t.Skip("HW_CC_GLOBAL_GATEWAY_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCCAuth(t *testing.T) {
	if HW_CC_PEER_DOMAIN_ID == "" || HW_CC_PEER_CONNECTION_ID == "" {
		t.Skip("HW_CC_PEER_DOMAIN_ID, HW_CC_PEER_CONNECTION_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCCPermission(t *testing.T) {
	if HW_CC_PERMISSION_ID == "" {
		t.Skip("HW_CC_PERMISSION_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCSEMicroserviceEngineID(t *testing.T) {
	if HW_CSE_MICROSERVICE_ENGINE_ID == "" {
		t.Skip("HW_CSE_MICROSERVICE_ENGINE_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCSEMicroserviceEngineAdminPassword(t *testing.T) {
	if HW_CSE_MICROSERVICE_ENGINE_ADMIN_PASSWORD == "" {
		t.Skip("HW_CSE_MICROSERVICE_ENGINE_ADMIN_PASSWORD must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCSENacosMicroserviceEngineID(t *testing.T) {
	if HW_CSE_NACOS_MICROSERVICE_ENGINE_ID == "" {
		t.Skip("HW_CSE_NACOS_MICROSERVICE_ENGINE_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCSSLocalDiskFlavor(t *testing.T) {
	if HW_CSS_LOCAL_DISK_FLAVOR == "" {
		t.Skip("HW_CSS_LOCAL_DISK_FLAVOR must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCSSElbAgency(t *testing.T) {
	if HW_CSS_ELB_AGENCY == "" {
		t.Skip("HW_CSS_ELB_AGENCY must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCSSUpgradeAgency(t *testing.T) {
	if HW_CSS_UPGRADE_AGENCY == "" {
		t.Skip("HW_CSS_UPGRADE_AGENCY must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckASScalingGroupID(t *testing.T) {
	if HW_AS_SCALING_GROUP_ID == "" {
		t.Skip("HW_AS_SCALING_GROUP_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckASScalingConfigurationID(t *testing.T) {
	if HW_AS_SCALING_CONFIGURATION_ID == "" {
		t.Skip("HW_AS_SCALING_CONFIGURATION_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckASScalingPolicyID(t *testing.T) {
	if HW_AS_SCALING_POLICY_ID == "" {
		t.Skip("HW_AS_SCALING_POLICY_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckASLifecycleActionKey(t *testing.T) {
	if HW_AS_LIFECYCLE_ACTION_KEY == "" {
		t.Skip("HW_AS_LIFECYCLE_ACTION_KEY must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckASINSTANCEID(t *testing.T) {
	if HW_AS_INSTANCE_ID == "" {
		t.Skip("HW_AS_INSTANCE_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckASLifecycleHookName(t *testing.T) {
	if HW_AS_LIFECYCLE_HOOK_NAME == "" {
		t.Skip("HW_AS_LIFECYCLE_HOOK_NAME must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDataArtsWorkSpaceID(t *testing.T) {
	if HW_DATAARTS_WORKSPACE_ID == "" {
		t.Skip("This environment does not support DataArts Studio tests")
	}
}

// lintignore:AT003
func TestAccPreCheckDataArtsReviewerName(t *testing.T) {
	if HW_DATAARTS_REVIEWER_NAME == "" {
		t.Skip("HW_DATAARTS_REVIEWER_NAME must be set for DataService tests")
	}
}

// lintignore:AT003
func TestAccPreCheckDataArtsRelatedDliQueueName(t *testing.T) {
	if HW_DATAARTS_DLI_QUEUE_NAME == "" {
		t.Skip("HW_DATAARTS_DLI_QUEUE_NAME must be set for the DataService tests")
	}
}

// lintignore:AT003
func TestAccPreCheckDataArtsDataServiceApigInstanceId(t *testing.T) {
	if HW_DATAARTS_INSTANCE_ID_IN_APIG == "" {
		t.Skip("HW_DATAARTS_INSTANCE_ID_IN_APIG must be set for the DataService tests")
	}
}

// lintignore:AT003
func TestAccPreCheckDataArtsManagerID(t *testing.T) {
	if HW_DATAARTS_MANAGER_ID == "" {
		t.Skip("This environment does not support DataArts Studio permission set tests")
	}
}

// lintignore:AT003
func TestAccPreCheckDataArtsBizCatalogID(t *testing.T) {
	if HW_DATAARTS_BIZ_CATALOG_ID == "" {
		t.Skip("HW_DATAARTS_BIZ_CATALOG_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDataArtsCdmName(t *testing.T) {
	if HW_DATAARTS_CDM_NAME == "" {
		t.Skip("HW_DATAARTS_CDM_NAME must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDataArtsDataClassificationID(t *testing.T) {
	if HW_DATAARTS_SECRECY_LEVEL_ID == "" || HW_DATAARTS_SECRECY_LEVEL_ID_UPDATE == "" {
		t.Skip("HW_DATAARTS_SECRECY_LEVEL_ID and HW_DATAARTS_SECRECY_LEVEL_ID_UPDATE must be set for the acceptance test")
	}

	if HW_DATAARTS_CATEGORY_ID == "" || HW_DATAARTS_CATEGORY_ID_UPDATE == "" {
		t.Skip("HW_DATAARTS_CATEGORY_ID and HW_DATAARTS_CATEGORY_ID_UPDATE must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDataArtsBuiltinRule(t *testing.T) {
	if HW_DATAARTS_BUILTIN_RULE_ID == "" || HW_DATAARTS_BUILTIN_RULE_NAME == "" {
		t.Skip("HW_DATAARTS_BUILTIN_RULE_ID and HW_DATAARTS_BUILTIN_RULE_NAME must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDataArtsSubjectID(t *testing.T) {
	if HW_DATAARTS_SUBJECT_ID == "" {
		t.Skip("HW_DATAARTS_SUBJECT_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDataArtsConnectionName(t *testing.T) {
	if HW_DATAARTS_CONNECTION_NAME == "" {
		t.Skip("HW_DATAARTS_CONNECTION_NAME must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDataArtsConnectionID(t *testing.T) {
	if HW_DATAARTS_CONNECTION_ID == "" {
		t.Skip("HW_DATAARTS_CONNECTION_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDataArtsArchitectureReviewer(t *testing.T) {
	if HW_DATAARTS_ARCHITECTURE_USER_ID == "" || HW_DATAARTS_ARCHITECTURE_USER_NAME == "" {
		t.Skip("HW_DATAARTS_ARCHITECTURE_USER_ID and HW_DATAARTS_ARCHITECTURE_USER_NAME must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDataArtsSecurityPermissionSetMember(t *testing.T) {
	if HW_DATAARTS_SECURITY_PERMISSSIONSET_MEMBER_OBJECT_ID == "" || HW_DATAARTS_SECURITY_PERMISSSIONSET_MEMBER_OBJECT_NAME == "" {
		t.Skip("HW_DATAARTS_SECURITY_PERMISSSIONSET_MEMBER_OBJECT_ID and HW_DATAARTS_SECURITY_PERMISSSIONSET_MEMBER_OBJECT_NAME " +
			"must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckAKAndSK(t *testing.T) {
	if HW_ACCESS_KEY == "" || HW_SECRET_KEY == "" {
		t.Skip("HW_ACCESS_KEY and HW_SECRET_KEY must be set for acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckECSLaunchTemplateID(t *testing.T) {
	if HW_ECS_LAUNCH_TEMPLATE_ID == "" {
		t.Skip("HW_ECS_LAUNCH_TEMPLATE_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckHWIOTDAAccessAddress(t *testing.T) {
	if HW_IOTDA_ACCESS_ADDRESS == "" {
		t.Skip("HW_IOTDA_ACCESS_ADDRESS must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckIOTDABatchTaskFilePath(t *testing.T) {
	if HW_IOTDA_BATCHTASK_FILE_PATH == "" {
		t.Skip("HW_IOTDA_BATCHTASK_FILE_PATH must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckMutilAZ(t *testing.T) {
	if HW_DWS_MUTIL_AZS == "" {
		t.Skip("HW_DWS_MUTIL_AZS must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDwsClusterId(t *testing.T) {
	if HW_DWS_CLUSTER_ID == "" {
		t.Skip("HW_DWS_CLUSTER_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDwsLogicalModeClusterId(t *testing.T) {
	if HW_DWS_LOGICAL_MODE_CLUSTER_ID == "" {
		t.Skip("HW_DWS_LOGICAL_MODE_CLUSTER_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDwsLogicalClusterName(t *testing.T) {
	if HW_DWS_LOGICAL_CLUSTER_NAME == "" {
		t.Skip("HW_DWS_LOGICAL_CLUSTER_NAME must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDwsSnapshotPolicyName(t *testing.T) {
	if HW_DWS_SNAPSHOT_POLICY_NAME == "" {
		t.Skip("HW_DWS_SNAPSHOT_POLICY_NAME must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDwsClusterUserNames(t *testing.T) {
	userNames := strings.Split(HW_DWS_ASSOCIATE_USER_NAMES, ",")
	// One is used to associate to the queue, and the other is used to update the user associated with the queue.
	if len(userNames) < 2 {
		t.Skip("The length of HW_DWS_ASSOCIATE_USER_NAMES must be 2 for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDwsAutomatedSnapshot(t *testing.T) {
	if HW_DWS_AUTOMATED_SNAPSHOT_ID == "" {
		t.Skip("HW_DWS_AUTOMATED_SNAPSHOT_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDwsExtDataSourceAgencyNames(t *testing.T) {
	// Control OBS type data source acceptance test, one for creating resource and the other for updating resource.
	agencyNames := strings.Split(HW_DWS_OBS_AGENCY_NAMES, ",")
	if len(agencyNames) < 2 {
		t.Skip("The length of HW_DWS_OBS_AGENCY_NAMES must be 2 for the OBS data source acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDCSAccountWhitelist(t *testing.T) {
	if HW_DCS_ACCOUNT_WHITELIST == "" {
		t.Skip("HW_DCS_ACCOUNT_WHITELIST must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDCSInstanceID(t *testing.T) {
	if HW_DCS_INSTANCE_ID == "" {
		t.Skip("HW_DCS_INSTANCE_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDcsTimeRange(t *testing.T) {
	if HW_DCS_BEGIN_TIME == "" || HW_DCS_END_TIME == "" {
		t.Skip("HW_DCS_BEGIN_TIME and HW_DCS_END_TIME must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckElbGatewayType(t *testing.T) {
	if HW_ELB_GATEWAY_TYPE == "" {
		t.Skip("HW_ELB_GATEWAY_TYPE must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCCMPushCertificateID(t *testing.T) {
	if HW_CERT_BATCH_PUSH_ID == "" {
		t.Skip("HW_CERT_BATCH_PUSH_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckCCMPushWAFInstance(t *testing.T) {
	if HW_CERT_BATCH_PUSH_WAF_ID == "" {
		t.Skip("HW_CERT_BATCH_PUSH_WAF_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckVPCEPServiceId(t *testing.T) {
	if HW_VPCEP_SERVICE_ID == "" {
		t.Skip("HW_VPCEP_SERVICE_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckHSSHostProtectionHostId(t *testing.T) {
	if HW_HSS_HOST_PROTECTION_HOST_ID == "" {
		t.Skip("HW_HSS_HOST_PROTECTION_HOST_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckHSSHostProtectionQuotaId(t *testing.T) {
	if HW_HSS_HOST_PROTECTION_QUOTA_ID == "" {
		t.Skip("HW_HSS_HOST_PROTECTION_QUOTA_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckHSSCCEProtection(t *testing.T) {
	if HW_CCE_CLUSTER_ID == "" || HW_CCE_CLUSTER_NAME == "" {
		t.Skip("HW_CCE_CLUSTER_ID and HW_CCE_CLUSTER_NAME must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDDMInstanceID(t *testing.T) {
	if HW_DDM_INSTANCE_ID == "" {
		t.Skip("HW_DDM_INSTANCE_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDDMProcessId(t *testing.T) {
	if HW_DDM_PROCESS_ID == "" {
		t.Skip("HW_DDM_PROCESS_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDDMTimeRange(t *testing.T) {
	if HW_DDM_START_TIME == "" || HW_DDM_END_TIME == "" {
		t.Skip("HW_DDM_START_TIME and HW_DDM_END_TIME must be set for acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDDSInstanceID(t *testing.T) {
	if HW_DDS_INSTANCE_ID == "" {
		t.Skip("HW_DDS_INSTANCE_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDDSSecondLevelMonitoringEnabled(t *testing.T) {
	if HW_DDS_SECOND_LEVEL_MONITORING_ENABLED == "" {
		t.Skip("HW_DDS_SECOND_LEVEL_MONITORING_ENABLED must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDDSScheduledTasksEnabled(t *testing.T) {
	if HW_DDS_SCHEDULED_TASKS_ENABLED == "" {
		t.Skip("HW_DDS_SCHEDULED_TASKS_ENABLED must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDDSTimeRange(t *testing.T) {
	if HW_DDS_START_TIME == "" || HW_DDS_END_TIME == "" {
		t.Skip("HW_DDS_START_TIME and HW_DDS_END_TIME must be set for acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckDDSRecycleInstancesEnabled(t *testing.T) {
	if HW_DDS_RECYCLE_INSTANCES_ENABLED == "" {
		t.Skip("HW_DDS_RECYCLE_INSTANCES_ENABLED must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckRdsCrossRegionBackupInstanceId(t *testing.T) {
	if HW_RDS_CROSS_REGION_BACKUP_INSTANCE_ID == "" {
		t.Skip("HW_RDS_CROSS_REGION_BACKUP_INSTANCE_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckRdsInstanceId(t *testing.T) {
	if HW_RDS_INSTANCE_ID == "" {
		t.Skip("HW_RDS_INSTANCE_ID must be set for RDS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckRdsInstanceName(t *testing.T) {
	if HW_RDS_INSTANCE_NAME == "" {
		t.Skip("HW_RDS_INSTANCE_NAME must be set for RDS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckRdsDatabaseName(t *testing.T) {
	if HW_RDS_DATABASE_NAME == "" {
		t.Skip("HW_RDS_DATABASE_NAME must be set for RDS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckRdsTableName(t *testing.T) {
	if HW_RDS_TABLE_NAME == "" {
		t.Skip("HW_RDS_TABLE_NAME must be set for RDS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckRdsBackupId(t *testing.T) {
	if HW_RDS_BACKUP_ID == "" {
		t.Skip("HW_RDS_BACKUP_ID must be set for RDS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckRdsTimeRange(t *testing.T) {
	if HW_RDS_START_TIME == "" || HW_RDS_END_TIME == "" {
		t.Skip("HW_RDS_START_TIME and HW_RDS_END_TIME must be set for RDS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckRdsFixedIp(t *testing.T) {
	if HW_RDS_FIXED_IP == "" {
		t.Skip("HW_RDS_FIXED_IP must be set for RDS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCssLowEngineVersion(t *testing.T) {
	if HW_CSS_LOW_ENGINE_VERSION == "" {
		t.Skip("HW_CSS_LOW_ENGINE_VERSION must be set for CSS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCssTargetImageId(t *testing.T) {
	if HW_CSS_TARGET_IMAGE_ID == "" {
		t.Skip("HW_CSS_TARGET_IMAGE_ID must be set for CSS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCssReplaceAgency(t *testing.T) {
	if HW_CSS_REPLACE_AGENCY == "" {
		t.Skip("HW_CSS_REPLACE_AGENCY must be set for CSS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckCssAzMigrateAgency(t *testing.T) {
	if HW_CSS_AZ_MIGRATE_AGENCY == "" {
		t.Skip("HW_CSS_AZ_MIGRATE_AGENCY must be set for CSS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckDMSKafkaInstanceID(t *testing.T) {
	if HW_DMS_KAFKA_INSTANCE_ID == "" {
		t.Skip("HW_DMS_KAFKA_INSTANCE_ID must be set for DMS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckDMSKafkaTopicName(t *testing.T) {
	if HW_DMS_KAFKA_TOPIC_NAME == "" {
		t.Skip("HW_DMS_KAFKA_TOPIC_NAME must be set for DMS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckDMSKafkaConsumerGroupName(t *testing.T) {
	if HW_DMS_KAFKA_CONSUMER_GROUP_NAME == "" {
		t.Skip("HW_DMS_KAFKA_CONSUMER_GROUP_NAME must be set for DMS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckDMSRocketMQInstanceID(t *testing.T) {
	if HW_DMS_ROCKETMQ_INSTANCE_ID == "" {
		t.Skip("HW_DMS_ROCKETMQ_INSTANCE_ID must be set for DMS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckDMSRocketMQGroupName(t *testing.T) {
	if HW_DMS_ROCKETMQ_GROUP_NAME == "" {
		t.Skip("HW_DMS_ROCKETMQ_GROUP_NAME must be set for DMS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckDMSRocketMQTopicName(t *testing.T) {
	if HW_DMS_ROCKETMQ_TOPIC_NAME == "" {
		t.Skip("HW_DMS_ROCKETMQ_TOPIC_NAME must be set for DMS acceptance tests")
	}
}

// lintignore:AT003
func TestAccPreCheckAsDedicatedHostId(t *testing.T) {
	if HW_DEDICATED_HOST_ID == "" {
		t.Skip("HW_DEDICATED_HOST_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckAsDataDiskImageId(t *testing.T) {
	if HW_IMS_DATA_DISK_IMAGE_ID == "" {
		t.Skip("HW_IMS_DATA_DISK_IMAGE_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckVpcId(t *testing.T) {
	if HW_VPC_ID == "" {
		t.Skip("HW_VPC_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckVpcEnhancedLocalRoute(t *testing.T) {
	if HW_VPC_ENHANCED_LOCAL_ROUTE == "" {
		t.Skip("HW_VPC_ENHANCED_LOCAL_ROUTE must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckVPNP2cGatewayId(t *testing.T) {
	if HW_VPN_P2C_GATEWAY_ID == "" {
		t.Skip("HW_VPN_P2C_GATEWAY_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckVPNP2cServer(t *testing.T) {
	if HW_VPN_P2C_SERVER == "" {
		t.Skip("HW_VPN_P2C_SERVER must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckVPNP2cClientCACertificate(t *testing.T) {
	if HW_VPN_P2C_CLIENT_CA_CERTIFICATE == "" {
		t.Skip("HW_VPN_P2C_CLIENT_CA_CERTIFICATE must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckSubnetId(t *testing.T) {
	if HW_SUBNET_ID == "" {
		t.Skip("HW_SUBNET_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckSecurityGroupId(t *testing.T) {
	if HW_SECURITY_GROUP_ID == "" {
		t.Skip("HW_SECURITY_GROUP_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPrecheckDcFlag(t *testing.T) {
	if HW_DC_ENABLE_FLAG == "" {
		t.Skip("HW_DC_ENABLE_FLAG must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPrecheckDscInstance(t *testing.T) {
	if HW_DSC_INSTANCE_ID == "" {
		t.Skip("HW_DSC_INSTANCE_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPrecheckDscAlarmTopicID(t *testing.T) {
	if HW_DSC_ALARM_TOPIC_ID == "" {
		t.Skip("HW_DSC_ALARM_TOPIC_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPrecheckTimeStamp(t *testing.T) {
	if HW_CDN_TIMESTAMP == "" {
		t.Skip("HW_CDN_TIMESTAMP must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPrecheckCDNAnalytics(t *testing.T) {
	if HW_CDN_START_TIME == "" || HW_CDN_END_TIME == "" || HW_CDN_STAT_TYPE == "" {
		t.Skip("HW_CDN_START_TIME, HW_CDN_END_TIME, and HW_CDN_STAT_TYPE must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPrecheckSFSTurboBackupId(t *testing.T) {
	if HW_SFS_TURBO_BACKUP_ID == "" {
		t.Skip("HW_SFS_TURBO_BACKUP_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPrecheckSfsFileSystemNames(t *testing.T, min int) {
	// For this acceptance test, you should prepare three SFS file systems.
	if HW_SFS_FILE_SYSTEM_NAMES == "" || len(strings.Split(HW_SFS_FILE_SYSTEM_NAMES, ",")) < min {
		t.Skip("At least three file system name must be supported during the HW_SFS_FILE_SYSTEM_NAMES, and separated by commas")
	}
}

// lintignore:AT003
func TestAccPrecheckEVSTransferAccepter(t *testing.T) {
	if HW_EVS_TRANSFER_ID == "" || HW_EVS_TRANSFER_AUTH_KEY == "" {
		t.Skip("HW_EVS_TRANSFER_ID and HW_EVS_TRANSFER_AUTH_KEY must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckEVSFlag(t *testing.T) {
	if HW_EVS_ENABLE_FLAG == "" {
		t.Skip("Skip the EVS acceptance tests.")
	}
}

// lintignore:AT003
func TestAccPrecheckDewFlag(t *testing.T) {
	// The key pair operation task, such as key pair bind or unbind task
	// Query the task execution status(running or failed)
	if HW_DEW_ENABLE_FLAG == "" {
		t.Skip("HW_DEW_ENABLE_FLAG must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPreCheckKPSKeyFilePath(t *testing.T) {
	if HW_KPS_KEY_FILE_PATH == "" {
		t.Skip("HW_KPS_KEY_FILE_PATH must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPrecheckDbssInstanceId(t *testing.T) {
	if HW_DBSS_INSATNCE_ID == "" {
		t.Skip("HW_DBSS_INSATNCE_ID must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPrecheckCphAdbObjectPath(t *testing.T) {
	if HW_CPH_OBS_OBJECT_PATH == "" {
		t.Skip("HW_CPH_OBS_OBJECT_PATH must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPrecheckCphObsBucketName(t *testing.T) {
	if HW_CPH_OBS_BUCKET_NAME == "" {
		t.Skip("HW_CPH_OBS_BUCKET_NAME must be set for the acceptance test")
	}
}

// lintignore:AT003
func TestAccPrecheckSmnSubscribedTopicUrn(t *testing.T) {
	if HW_SMN_SUBSCRIBED_TOPIC_URN == "" {
		t.Skip("HW_SMN_SUBSCRIBED_TOPIC_URN must be set for the acceptance test")
	}
}
