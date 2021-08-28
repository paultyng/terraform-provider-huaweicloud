---
subcategory: "GaussDB"
---

# huaweicloud_gaussdb_opengauss_instance

GaussDB OpenGauss instance management within HuaweiCoud.

## Example Usage

### create a basic instance

```hcl
resource "huaweicloud_gaussdb_opengauss_instance" "instance_acc" {
  name              = "opengaussdb_instance_1"
  password          = "Test@123"
  flavor            = "gaussdb.opengauss.ee.dn.m6.2xlarge.8.in"
  vpc_id            = var.vpc_id
  subnet_id         = var.subnet_id
  availability_zone = "cn-north-4a,cn-north-4a,cn-north-4a"
  security_group_id = var.secgroup.id
  sharding_num      = 1
  coordinator_num   = 1

  ha {
    mode             = "enterprise"
    replication_mode = "sync"
    consistency      = "strong"
  }

  volume {
    type = "ULTRAHIGH"
    size = 40
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the instance. If omitted, the provider-level
  region will be used. Changing this creates a new resource.

* `name` - (Required, String) Specifies the instance name, which can be the same as an existing instance name. The value
  must be 4 to 64 characters in length and start with a letter. It is case-sensitive and can contain only letters,
  digits, hyphens (-), and underscores (_).

* `flavor` - (Required, String, ForceNew) Specifies the instance specifications. Please reference the API docs for valid
  options. Changing this parameter will create a new resource.

* `password` - (Required, String) Specifies the database password. The value must be 8 to 32 characters in length,
  including uppercase and lowercase letters, digits, and special characters, such as ~!@#%^*-_=+? You are advised to
  enter a strong password to improve security, preventing security risks such as brute force cracking.

* `availability_zone` - (Required, String, ForceNew) Specifies the Availability Zone information, can be three same or
  different az like "cn-north-4a,cn-north-4a,cn-north-4a". Changing this parameter will create a new resource.

* `vpc_id` - (Required, String, ForceNew) Specifies the VPC ID. Changing this parameter will create a new resource.

* `subnet_id` - (Required, String, ForceNew) Specifies the network ID of a subnet. Changing this parameter will create a
  new resource.

* `security_group_id` - (Optional, String, ForceNew) Specifies the security group ID. Changing this parameter will
  create a new resource.

* `volume` - (Required, List) Specifies the volume storage information. Structure is documented below.

* `port` - (Optional, String) Specifies the port information. Defaults to "8000". Changing this parameter will create a
  new resource.

* `configuration_id` - (Optional, String, ForceNew) The parameter template id. Changing this parameter will create a new
  resource.

* `sharding_num` - (Optional, Int) The Sharding num. Values: 1~9. The default value is 3.

* `coordinator_num` - (Optional, Int) The Coordinator num. Values: 1~9. The default value is 3. The value must not be
  greater than twice value of `sharding_num`.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id. Changing this parameter will create
  a new resource.

* `time_zone` - (Optional, String, ForceNew) Specifies the time zone. Defaults to "UTC+08:00". Changing this parameter
  will create a new resource.

* `force_import` - (Optional, Bool) If specified, try to import the instance instead of creating if the name already
  existed.

* `datastore` - (Optional, List, ForceNew) Specifies the datastore information. Structure is documented below. Changing
  this parameter will create a new resource.

* `backup_strategy` - (Optional, List) Specifies the advanced backup policy. Structure is documented below.

* `ha` - (Optional, List, ForceNew) Specifies the HA information. Structure is documented below. Changing this parameter
  will create a new resource.

The `datastore` block supports:

* `engine` - (Required, String, ForceNew) Specifies the database engine. Only "GaussDB(for openGauss)" is supported now.

* `version` - (Optional, String, ForceNew) Specifies the database version. Defaults to the latest version. Please
  reference to the API docs for valid options.

The `volume` block supports:

* `type` - (Required, String, ForceNew) Specifies the volume type. Only "ULTRAHIGH" is supported now.

* `size` - (Required, Int) Specifies the volume size (in gigabytes) for a Sharding. The value should between 40G ~ 5TB.

The `ha` block supports:

* `mode` - (Required, String, ForceNew) Specifies the database mode. Only "enterprise" is supported now.

* `replication_mode` - (Required, String, ForceNew) Specifies the database replication mode. Only "sync" is supported
  now.

* `consistency` - (Optional, String, ForceNew) Specifies the database consistency mode. Valid options are "strong" and "
  eventual".

The `backup_strategy` block supports:

* `start_time` - (Required, String) Specifies the backup time window. Automated backups will be triggered during the
  backup time window. It must be a valid value in the "hh:mm-HH:MM" format. The current time is in the UTC format. The
  HH value must be 1 greater than the hh value. The values of mm and MM must be the same and must be set to 00, 15, 30
  or 45. Example value: 08:15-09:15, 23:00-00:00.

* `keep_days` - (Optional, Int) Specifies the number of days to retain the generated backup files. The value ranges from
  0 to 732. If this parameter is set to 0, the automated backup policy is not set. If this parameter is not transferred,
  the automated backup policy is enabled by default.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Indicates the DB instance ID.
* `status` - Indicates the DB instance status.
* `type` - Indicates the database type.
* `port` - Indicates the database port.
* `private_ips` - Indicates the private IP address of the DB instance.
* `public_ips` - Indicates the public IP address of the DB instance.
* `endpoints` - Indicates the connection endpoints list of the DB instance. Example: [127.0.0.1:8000].
* `db_user_name` - Indicates the default username.
* `switch_strategy` - Indicates the switch strategy.
* `maintenance_window` - Indicates the maintenance window.
* `nodes` - Indicates the instance nodes information. Structure is documented below.

The `nodes` block contains:

* `id` - Indicates the node ID.
* `name` - Indicates the node name.
* `role` - Indicates the node role: master or slave.
* `status` - Indicates the node status.
* `availability_zone` - Indicates the availability zone of the node.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 120 minute.
* `update` - Default is 60 minute.
* `delete` - Default is 30 minute.

## Import

OpenGaussDB instance can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_gaussdb_opengauss_instance.instance_1 ee678f40-ce8e-4d0c-8221-38dead426f06
```
