---
subcategory: "GaussDB"
---

# huaweicloud_gaussdb_mysql_instances

Use this data source to list all available HuaweiCloud gaussdb mysql instances.

## Example Usage

```hcl
data "huaweicloud_gaussdb_mysql_instances" "this" {
  name = "gaussdb-instance"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the instances. If omitted, the provider-level region will
  be used.

* `name` - (Optional, String) Specifies the name of the instance.

* `vpc_id` - (Optional, String) Specifies the VPC ID.

* `subnet_id` - (Optional, String) Specifies the network ID of a subnet.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.

* `instances` - An array of available instances.

The `instances` block supports:

* `region` - The region of the instance.

* `name` - Indicates the name of the instance.

* `vpc_id` - Indicates the VPC ID.

* `subnet_id` - Indicates the network ID of a subnet.

* `id` - Indicates the ID of the instance.

* `flavor` - Indicates the instance specifications.

* `security_group_id` - Indicates the security group ID.

* `configuration_id` - Indicates the configuration ID.

* `enterprise_project_id` - Indicates the enterprise project id.

* `read_replicas` - Indicates the count of read replicas.

* `time_zone` - Indicates the time zone.

* `availability_zone_mode` - Indicates the availability zone mode: "single" or "multi".

* `master_availability_zone` - Indicates the availability zone where the master node resides.

* `datastore` - Indicates the database information. Structure is documented below.

* `backup_strategy` - Indicates the advanced backup policy. Structure is documented below.

* `status` - Indicates the DB instance status.

* `port` - Indicates the database port.

* `mode` - Indicates the instance mode.

* `db_user_name` - Indicates the default username.

* `private_write_ip` - Indicates the private IP address of the DB instance.

* `nodes` - Indicates the instance nodes information. Structure is documented below.

The `datastore` block supports:

* `engine` - Indicates the database engine.
* `version` - Indicates the database version.

The `backup_strategy` block supports:

* `start_time` - Indicates the backup time window.
* `keep_days` - Indicates the number of days to retain the generated

The `nodes` block contains:

* `id` - Indicates the node ID.
* `name` - Indicates the node name.
* `type` - Indicates the node type: master or slave.
* `status` - Indicates the node status.
* `private_read_ip` - Indicates the private IP address of a node.
* `availability_zone` - Indicates the availability zone where the node resides.
