---
subcategory: "Enterprise Project Management Service (EPS)"
---

# huaweicloud_enterprise_project

Use this data source to get an enterprise project from HuaweiCloud

## Example Usage

```hcl
data "huaweicloud_enterprise_project" "test" {
  name = "test"
}
```

## Resources Supported Currently:

Service Name | Resource Name | Sub Resource Name
---- | --- | ---
AS  | huaweicloud_as_group |
BCS | huaweicloud_bcs_instance |
BMS | huaweicloud_bms_instance |
CBR | huaweicloud_cbr_vault |
CCE | huaweicloud_cce_cluster | huaweicloud_cce_node<br>huaweicloud_cce_node_pool<br>huaweicloud_cce_addon
CDM | huaweicloud_cdm_cluster |
CDN | huaweicloud_cdn_domain |
CES | huaweicloud_ces_alarmrule |
DCS | huaweicloud_dcs_instance |
DDS | huaweicloud_dds_instance |
DMS | huaweicloud_dms_kafka_instance<br>huaweicloud_dms_rabbitmq_instance |
DNS | huaweicloud_dns_ptrrecord<br>huaweicloud_dns_zone |
ECS | huaweicloud_compute_instance |
EIP | huaweicloud_vpc_eip<br>huaweicloud_vpc_bandwidth |
ELB | huaweicloud_lb_loadbalancer |
Dedicated ELB | huaweicloud_elb_certificate<br>huaweicloud_elb_ipgroup<br>huaweicloud_elb_loadbalancer |
EVS | huaweicloud_evs_volume |
FGS | huaweicloud_fgs_function |
GaussDB | huaweicloud_gaussdb_cassandra_instance<br>huaweicloud_gaussdb_mysql_instance<br>huaweicloud_gaussdb_opengauss_instance |
IMS | huaweicloud_images_image |
KMS | huaweicloud_kms_key |
NAT | huaweicloud_nat_gateway | huaweicloud_nat_snat_rule<br>huaweicloud_nat_dnat_rule
OBS | huaweicloud_obs_bucket | huaweicloud_obs_bucket_object<br>huaweicloud_obs_bucket_policy
RDS | huaweicloud_rds_instance<br>huaweicloud_rds_read_replica_instance |
SFS | huaweicloud_sfs_file_system<br>huaweicloud_sfs_turbo | huaweicloud_sfs_access_rule
VPC | huaweicloud_vpc<br>huaweicloud_networking_secgroup | huaweicloud_vpc_subnet<br>huaweicloud_vpc_route<br>huaweicloud_networking_secgroup_rule

## Argument Reference

* `name` - (Optional, String) Specifies the enterprise project name. Fuzzy search is supported.

* `id` - (Optional, String) Specifies the ID of an enterprise project. The value 0 indicates enterprise project default.

* `status` - (Optional, Int) Specifies the status of an enterprise project.
  + 1 indicates Enabled.
  + 2 indicates Disabled.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - Provides supplementary information about the enterprise project.

* `created_at` - Specifies the time (UTC) when the enterprise project was created. Example: 2018-05-18T06:49:06Z

* `updated_at` - Specifies the time (UTC) when the enterprise project was modified. Example: 2018-05-28T02:21:36Z

