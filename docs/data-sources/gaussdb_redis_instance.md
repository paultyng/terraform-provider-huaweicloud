---
subcategory: "GaussDB"
---

# huaweicloud_gaussdb_redis_instance

Use this data source to get available HuaweiCloud gaussdb redis instance.

## Example Usage

```hcl
data "huaweicloud_gaussdb_redis_instance" "test" {
  name = "gaussdb-instance"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the instance. If omitted, the provider-level region will
  be used.

* `name` - (Optional, String) Specifies the name of the instance.

* `vpc_id` - (Optional, String) Specifies the VPC ID.

* `subnet_id` - (Optional, String) Specifies the network ID of a subnet.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Indicates the ID of the instance.

* `status` - Indicates the DB instance status.

* `mode` - Indicates the instance mode.

* `flavor` - Indicates the instance specifications.

* `security_group_id` - Indicates the security group ID.

* `enterprise_project_id` - Indicates the enterprise project id.

* `db_user_name` - Indicates the default username.

* `port` - Indicates the database port.

* `availability_zone` - Indicates the instance availability zone.

* `node_num` - Indicates the count of the nodes.

* `volume_size` - Indicates the size of the volume.

* `private_ips` - Indicates the list of private IP address of the nodes.

* `datastore` - Indicates the database information. Structure is documented below.

* `backup_strategy` - Indicates the advanced backup policy. Structure is documented below.

* `nodes` - Indicates the instance nodes information. Structure is documented below.

* `tags` - Indicates the key/value tags of the instance.

The `datastore` block supports:

* `engine` - Indicates the database engine.
* `storage_engine` - Indicates the database storage engine.
* `version` - Indicates the database version.

The `backup_strategy` block supports:

* `start_time` - Indicates the backup time window.
* `keep_days` - Indicates the number of days to retain the generated

The `nodes` block contains:

* `id` - Indicates the node ID.
* `name` - Indicates the node name.
* `private_ip` - Indicates the private IP address of a node.
* `status` - Indicates the node status.
* `support_reduce` - Indicates whether the node support reduce.
* `availability_zone` - Indicates the availability zone where the node resides.
