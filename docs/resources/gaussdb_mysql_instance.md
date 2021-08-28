---
subcategory: "GaussDB"
---

# huaweicloud_gaussdb_mysql_instance

GaussDB mysql instance management within HuaweiCoud.

## Example Usage

### create a basic instance

```hcl
resource "huaweicloud_gaussdb_mysql_instance" "instance_1" {
  name              = "gaussdb_instance_1"
  password          = var.password
  flavor            = "gaussdb.mysql.4xlarge.x86.4"
  vpc_id            = var.vpc_id
  subnet_id         = var.subnet_id
  security_group_id = var.secgroup_id
}
```

### create a gaussdb mysql instance with backup strategy

```hcl
resource "huaweicloud_gaussdb_mysql_instance" "instance_1" {
  name              = "gaussdb_instance_1"
  password          = var.password
  flavor            = "gaussdb.mysql.4xlarge.x86.4"
  vpc_id            = var.vpc_id
  subnet_id         = var.subnet_id
  security_group_id = var.secgroup_id

  backup_strategy {
    start_time = "03:00-04:00"
    keep_days  = 7
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the GaussDB mysql instance resource. If omitted,
  the provider-level region will be used. Changing this creates a new instance resource.

* `name` - (Required, String) Specifies the instance name, which can be the same as an existing instance name. The value
  must be 4 to 64 characters in length and start with a letter. It is case-sensitive and can contain only letters,
  digits, hyphens (-), and underscores (_).

* `flavor` - (Required, String) Specifies the instance specifications. Please use
  `gaussdb_mysql_flavors` data source to fetch the available flavors.

* `password` - (Required, String) Specifies the database password. The value must be 8 to 32 characters in length,
  including uppercase and lowercase letters, digits, and special characters, such as ~!@#%^*-_=+? You are advised to
  enter a strong password to improve security, preventing security risks such as brute force cracking.

* `vpc_id` - (Required, String, ForceNew) Specifies the VPC ID. Changing this parameter will create a new resource.

* `subnet_id` - (Required, String, ForceNew) Specifies the network ID of a subnet. Changing this parameter will create a
  new resource.

* `security_group_id` - (Optional, String, ForceNew) Specifies the security group ID. Required if the selected subnet
  doesn't enable network ACL. Changing this parameter will create a new resource.

* `configuration_id` - (Optional, String, ForceNew) Specifies the configuration ID. Changing this parameter will create
  a new resource.

* `dedicated_resource_id` - (Optional, String, ForceNew) Specifies the dedicated resource ID. Changing this parameter
  will create a new resource.

* `enterprise_project_id` - (Optional, String, ForceNew) Specifies the enterprise project id. Required if EPS enabled.
  Changing this parameter will create a new resource.

* `table_name_case_sensitivity` - (Optional, Bool) Whether the kernel table name is case sensitive. The value can
  be `true` (case sensitive) and `false` (case insensitive). Defaults to `false`. This parameter only works during
  creation.

* `read_replicas` - (Optional, Int) Specifies the count of read replicas. Defaults to 1.

* `time_zone` - (Optional, String, ForceNew) Specifies the time zone. Defaults to "UTC+08:00". Changing this parameter
  will create a new resource.

* `availability_zone_mode` - (Optional, String, ForceNew) Specifies the availability zone mode: "single" or "multi".
  Defaults to "single". Changing this parameter will create a new resource.

* `master_availability_zone` - (Optional, String, ForceNew) Specifies the availability zone where the master node
  resides. The parameter is required in multi availability zone mode. Changing this parameter will create a new
  resource.

* `datastore` - (Optional, List, ForceNew) Specifies the database information. Structure is documented below. Changing
  this parameter will create a new resource.

* `backup_strategy` - (Optional, List) Specifies the advanced backup policy. Structure is documented below.

* `proxy_flavor` - (Optional, String) Specifies the flavor of the proxy.

* `proxy_node_num` - (Optional, Int) Specifies the node count of the proxy.

* `force_import` - (Optional, Bool) If specified, try to import the instance instead of creating if the name already
  existed.

* `volume_size` - (Optional, Int) Specifies the volume size of the instance. The new storage space must be greater than
  the current storage and must be a multiple of 10 GB. Only valid when in prePaid mode.

The `datastore` block supports:

* `engine` - (Optional, String, ForceNew) Specifies the database engine. Only "gauss-mysql" is supported now.

* `version` - (Optional, String, ForceNew) Specifies the database version. Only "8.0" is supported now.

The `backup_strategy` block supports:

* `start_time` - (Required, String) Specifies the backup time window. Automated backups will be triggered during the
  backup time window. It must be a valid value in the "hh:mm-HH:MM" format. The current time is in the UTC format. The
  HH value must be 1 greater than the hh value. The values of mm and MM must be the same and must be set to 00. Example
  value: 08:00-09:00, 03:00-04:00.

* `keep_days` - (Optional, Int) Specifies the number of days to retain the generated backup files. The value ranges from
  0 to 35. If this parameter is set to 0, the automated backup policy is not set. If this parameter is not transferred,
  the automated backup policy is enabled by default. Backup files are stored for seven days by default.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.
* `status` - Indicates the DB instance status.
* `port` - Indicates the database port.
* `mode` - Indicates the instance mode.
* `db_user_name` - Indicates the default username.
* `private_write_ip` - Indicates the private IP address of the DB instance.
* `nodes` - Indicates the instance nodes information. Structure is documented below.
* `proxy_address` - Indicates the address of the proxy.
* `proxy_port` - Indicates the port of the proxy.

The `nodes` block contains:

* `id` - Indicates the node ID.
* `name` - Indicates the node name.
* `type` - Indicates the node type: master or slave.
* `status` - Indicates the node status.
* `private_read_ip` - Indicates the private IP address of a node.
* `availability_zone` - Indicates the availability zone where the node resides.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 60 minute.
* `update` - Default is 60 minute.
* `delete` - Default is 30 minute.

## Import

GaussDB instance can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_gaussdb_mysql_instance.instance_1 ee678f40-ce8e-4d0c-8221-38dead426f06
```
