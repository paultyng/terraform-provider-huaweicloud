package huaweicloud

// ServiceCatalog defines a struct which was used to generate a service client for huaweicloud.
// the endpoint likes https://{Name}.{Region}.myhuaweicloud.com/{Version}/{project_id}/{ResourceBase}
// For more information, please refer to Config.NewServiceClient
type ServiceCatalog struct {
	Name             string
	Version          string
	Scope            string
	Admin            bool
	ResourceBase     string
	WithOutProjectID bool
}

var allServiceCatalog = map[string]ServiceCatalog{
	// catalog for global service
	// identity is used for openstack keystone APIs
	"identity": ServiceCatalog{
		Name:             "iam",
		Version:          "v3",
		Scope:            "global",
		Admin:            true,
		WithOutProjectID: true,
	},
	// iam is used for huaweicloud IAM APIs
	"iam": ServiceCatalog{
		Name:             "iam",
		Version:          "v3.0",
		Scope:            "global",
		Admin:            true,
		WithOutProjectID: true,
	},
	"cdn": ServiceCatalog{
		Name:             "cdn",
		Version:          "v1.0",
		Scope:            "global",
		WithOutProjectID: true,
	},
	"dns": ServiceCatalog{
		Name:             "dns",
		Version:          "v2",
		Scope:            "global",
		WithOutProjectID: true,
	},
	// ******* client for Compute start *******
	"ecs": ServiceCatalog{
		Name:    "ecs",
		Version: "v1",
	},
	"ecsv11": ServiceCatalog{
		Name:    "ecs",
		Version: "v1.1",
	},
	"ecsv21": ServiceCatalog{
		Name:    "ecs",
		Version: "v2.1",
	},
	"autoscalingv1": ServiceCatalog{
		Name:    "as",
		Version: "autoscaling-api/v1",
	},
	"imagev2": ServiceCatalog{
		Name:             "ims",
		Version:          "v2",
		WithOutProjectID: true,
	},
	"ccev3": ServiceCatalog{
		Name:    "cce",
		Version: "api/v3/projects",
	},
	"cceaddonv3": ServiceCatalog{
		Name:             "cce",
		Version:          "api/v3",
		WithOutProjectID: true,
	},
	"cciv1": ServiceCatalog{
		Name:             "cci",
		Version:          "apis/networking.cci.io/v1beta1",
		WithOutProjectID: true,
	},
	"fgsv2": ServiceCatalog{
		Name:    "functiongraph",
		Version: "v2",
	},
	// ******* client for Compute end  *******//

	// ******* client for storage start ******//
	"blockstoragev2": ServiceCatalog{
		Name:    "evs",
		Version: "v2",
	},
	"blockstoragev3": ServiceCatalog{
		Name:    "evs",
		Version: "v3",
	},
	"sfsv2": ServiceCatalog{
		Name:    "sfs",
		Version: "v2",
	},
	"csbsv1": ServiceCatalog{
		Name:    "csbs",
		Version: "v1",
	},
	"vbsv2": ServiceCatalog{
		Name:    "vbs",
		Version: "v2",
	},
	// ******* client for storage end   ******//

	// ******* client for network start ******//
	"network": ServiceCatalog{
		Name:             "vpc",
		Version:          "v1",
		WithOutProjectID: true,
	},
	"networkv2": ServiceCatalog{
		Name:             "vpc",
		Version:          "v2.0",
		WithOutProjectID: true,
	},
	"natv2": ServiceCatalog{
		Name:             "nat",
		Version:          "v2.0",
		WithOutProjectID: true,
	},
	"loadelasticloadbalancer": ServiceCatalog{
		Name:             "elb",
		Version:          "v1.0",
		WithOutProjectID: true,
	},
	"fwv2": ServiceCatalog{
		Name:             "vpc",
		Version:          "v2.0",
		WithOutProjectID: true,
	},
	// ******* client for network end   ******//

	"vpc": ServiceCatalog{
		Name:             "vpc",
		Version:          "v1",
		WithOutProjectID: true,
	},
	"volumev2": ServiceCatalog{
		Name:    "evs",
		Version: "v2",
	},
	"volumev3": ServiceCatalog{
		Name:    "evs",
		Version: "v3",
	},
	"sfs-turbo": ServiceCatalog{
		Name:    "sfs-turbo",
		Version: "v1",
	},
	"dcsv2": ServiceCatalog{
		Name:    "dcs",
		Version: "v2",
	},

	// catalog for database
	"rdsv1": ServiceCatalog{
		Name:    "rds",
		Version: "rds/v1",
	},
	"rdsv3": ServiceCatalog{
		Name:    "rds",
		Version: "v3",
	},
	"ddsv3": ServiceCatalog{
		Name:    "dds",
		Version: "v3",
	},
	"cassandra": ServiceCatalog{
		Name:    "gaussdb-nosql",
		Version: "v3",
	},
	"gaussdb": ServiceCatalog{
		Name:    "gaussdb",
		Version: "mysql/v3",
	},
	"opengauss": ServiceCatalog{
		Name:    "gaussdb",
		Version: "opengauss/v3",
	},
	"bss": ServiceCatalog{
		Name:             "bss",
		Version:          "v1.0",
		WithOutProjectID: true,
	},
	// catalog for management service
	"ces": ServiceCatalog{
		Name:    "ces",
		Version: "V1.0",
	},
	"cts": ServiceCatalog{
		Name:    "cts",
		Version: "v1.0",
	},
	"lts": ServiceCatalog{
		Name:    "lts",
		Version: "v2",
	},

	// catalog for Security service
	"anti-ddos": ServiceCatalog{
		Name:    "antiddos",
		Version: "v1",
	},
	"kms": ServiceCatalog{
		Name:             "kms",
		Version:          "v1.0",
		WithOutProjectID: true,
	},

	// catalog for Enterprise Intelligence
	"mrs": ServiceCatalog{
		Name:    "mrs",
		Version: "v1.1",
	},
	"smn": ServiceCatalog{
		Name:         "smn",
		Version:      "v2",
		ResourceBase: "notifications",
	},

	// catalog for Application
	"apig": ServiceCatalog{
		Name:             "apig",
		Version:          "v1.0",
		ResourceBase:     "apigw",
		WithOutProjectID: true,
	},
	"dcsv1": ServiceCatalog{
		Name:    "dcs",
		Version: "v1.0",
	},
	"dms": ServiceCatalog{
		Name:    "dms",
		Version: "v1.0",
	},
}
