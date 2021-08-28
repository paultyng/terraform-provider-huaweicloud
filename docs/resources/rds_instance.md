---
subcategory: "Relational Database Service (RDS)"
---

# huaweicloud_rds_instance

Manage RDS instance resource within HuaweiCloud. This is an alternative to `huaweicloud_rds_instance_v3`

## Example Usage

### create a single db instance

```hcl
resource "huaweicloud_networking_secgroup" "secgroup" {
  name        = "terraform_test_security_group"
  description = "terraform security group acceptance test"
}

resource "huaweicloud_rds_instance" "instance" {
  name              = "terraform_test_rds_instance"
  flavor            = "rds.pg.n1.large.2"
  vpc_id            = "{{ vpc_id }}"
  subnet_id         = "{{ subnet_id }}"
  security_group_id = huaweicloud_networking_secgroup.secgroup.id
  availability_zone = [
    "{{ availability_zone }}"]

  db {
    type     = "PostgreSQL"
    version  = "12"
    password = "Huangwei!120521"
  }

  volume {
    type = "ULTRAHIGH"
    size = 100
  }

  backup_strategy {
    start_time = "08:00-09:00"
    keep_days  = 1
  }
}
```

### create a primary/standby db instance

```hcl
resource "huaweicloud_networking_secgroup" "secgroup" {
  name        = "terraform_test_security_group"
  description = "terraform security group acceptance test"
}

resource "huaweicloud_rds_instance" "instance" {
  name                = "terraform_test_rds_instance"
  flavor              = "rds.pg.n1.large.2.ha"
  ha_replication_mode = "async"
  vpc_id              = "{{ vpc_id }}"
  subnet_id           = "{{ subnet_id }}"
  security_group_id   = huaweicloud_networking_secgroup.secgroup.id
  availability_zone   = [
    "{{ availability_zone_1 }}",
    "{{ availability_zone_2 }}"]

  db {
    type     = "PostgreSQL"
    version  = "12"
    password = "Huangwei!120521"
  }
  volume {
    type = "ULTRAHIGH"
    size = 100
  }
  backup_strategy {
    start_time = "08:00-09:00"
    keep_days  = 1
  }
}
```

### create a single db instance with encrypted volume

```hcl
resource "huaweicloud_kms_key" "key" {
  key_alias       = "key_1"
  key_description = "first test key"
  is_enabled      = true
}

resource "huaweicloud_networking_secgroup" "secgroup" {
  name        = "test_security_group"
  description = "security group acceptance test"
}

resource "huaweicloud_rds_instance" "instance" {
  name              = "terraform_test_rds_instance"
  flavor            = "rds.pg.n1.large.2"
  vpc_id            = "{{ vpc_id }}"
  subnet_id         = "{{ subnet_id }}"
  security_group_id = huaweicloud_networking_secgroup.secgroup.id
  availability_zone = [
    "{{ availability_zone }}"]

  db {
    type     = "PostgreSQL"
    version  = "12"
    password = "Huangwei!120521"
  }
  volume {
    type               = "ULTRAHIGH"
    size               = 100
    disk_encryption_id = huaweicloud_kms_key.key.id
  }
  backup_strategy {
    start_time = "08:00-09:00"
    keep_days  = 1
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the rds instance resource. If omitted, the
  provider-level region will be used. Changing this creates a new rds instance resource.

* `availability_zone` - (Required, List, ForceNew) Specifies the list of AZ name. Changing this parameter will create a
  new resource.

* `name` - (Required, String) Specifies the DB instance name. The DB instance name of the same type must be unique for
  the same tenant. The value must be 4 to 64 characters in length and start with a letter. It is case-sensitive and can
  contain only letters, digits, hyphens (-), and underscores (_).

* `flavor` - (Required, String) Specifies the specification code.

  -> **NOTE:** Services will be interrupted for 5 to 10 minutes when you change RDS instance flavor.

* `db` - (Required, String, ForceNew) Specifies the database information. Structure is documented below. Changing this
  parameter will create a new resource.

* `vpc_id` - (Required, String, ForceNew) Specifies the VPC ID. Changing this parameter will create a new resource.

* `subnet_id` - (Required, String, ForceNew) Specifies the network id of a subnet. Changing this parameter will create a
  new resource.

* `security_group_id` - (Required, String) Specifies the security group which the RDS DB instance belongs to.

* `volume` - (Required, List) Specifies the volume information. Structure is documented below.

* `fixed_ip` - (Optional, String, ForceNew) Specifies an intranet floating IP address of RDS DB instance. Changing this
  parameter will create a new resource.

* `backup_strategy` - (Optional, List) Specifies the advanced backup policy. Structure is documented below.

* `ha_replication_mode` - (Optional, String, ForceNew) Specifies the replication mode for the standby DB instance.
  Changing this parameter will create a new resource.
  + For MySQL, the value is *async* or *semisync*.
  + For PostgreSQL, the value is *async* or *sync*.
  + For Microsoft SQL Server, the value is *sync*.

  -> **NOTE:** async indicates the asynchronous replication mode. semisync indicates the semi-synchronous replication
  mode. sync indicates the synchronous replication mode.

* `param_group_id` - (Optional, String, ForceNew) Specifies the parameter group ID. Changing this parameter will create
  a new resource.

* `time_zone` - (Optional, String, ForceNew) Specifies the UTC time zone. For MySQL and PostgreSQL Chinese mainland site
  and international site use UTC by default. The value ranges from UTC-12:00 to UTC+12:00 at the full hour. For
  Microsoft SQL Server international site use UTC by default and Chinese mainland site use China Standard Time. The time
  zone is expressed as a character string, refer to
  [HuaweiCloud Document](https://support.huaweicloud.com/intl/en-us/api-rds/rds_01_0002.html#rds_01_0002__table613473883617)
  .

* `charging_mode` - (Optional, String, ForceNew) Specifies the charging mode of the RDS DB instance. Valid values are *
  prePaid* and *postPaid*, defaults to *postPaid*. Changing this creates a new resource.

* `period_unit` - (Optional, String, ForceNew) Specifies the charging period unit of the RDS DB instance. Valid values
  are *month* and *year*. This parameter is mandatory if `charging_mode` is set to *prePaid*. Changing this creates a
  new resource.

* `period` - (Optional, Int, ForceNew) Specifies the charging period of the RDS DB instance. If `period_unit` is set
  to *month*, the value ranges from 1 to 9. If `period_unit` is set to *year*, the value ranges from 1 to 3. This
  parameter is mandatory if `charging_mode` is set to *prePaid*. Changing this creates a new resource.

* `auto_renew` - (Optional, String, ForceNew) Specifies whether auto renew is enabled. Valid values are "true" and "
  false". Changing this creates a new resource.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id of the RDS instance. Changing this
  parameter creates a new RDS instance.

* `ssl_enable` - (Optional, Bool) Specifies whether to enable the SSL for MySQL database.

* `tags` - (Optional, Map) A mapping of tags to assign to the RDS instance. Each tag is represented by one key-value
  pair.

The `db` block supports:

* `type` - (Required, String, ForceNew) Specifies the DB engine. Available value are *MySQL*, *PostgreSQL* and *
  SQLServer*. Changing this parameter will create a new resource.

* `version` - (Required, String, ForceNew) Specifies the database version. Changing this parameter will create a new
  resource. Available values detailed in
  [DB Engines and Versions](https://support.huaweicloud.com/intl/en-us/productdesc-rds/en-us_topic_0043898356.html).

* `password` - (Required, String, ForceNew) Specifies the database password. The value cannot be empty and should
  contain 8 to 32 characters, including uppercase and lowercase letters, digits, and the following special
  characters: ~!@#%^*-_=+? You are advised to enter a strong password to improve security, preventing security risks
  such as brute force cracking. Changing this parameter will create a new resource.

* `port` - (Optional, Int) Specifies the database port.
  + The MySQL database port ranges from 1024 to 65535 (excluding 12017 and 33071, which are occupied by the RDS system
      and cannot be used). The default value is 3306.
  + The PostgreSQL database port ranges from 2100 to 9500. The default value is 5432.
  + The Microsoft SQL Server database port can be 1433 or ranges from 2100 to 9500, excluding 5355 and 5985. The
      default value is 1433.

The `volume` block supports:

* `size` - (Required, Int) Specifies the volume size. Its value range is from 40 GB to 4000 GB. The value must be a
  multiple of 10 and greater than the original size.

* `type` - (Required, String, ForceNew) Specifies the volume type. Its value can be any of the following and is
  case-sensitive:
  + *ULTRAHIGH*: SSD storage.
  + *LOCALSSD*: local SSD storage.
  + *CLOUDSSD*: cloud SSD storage. This storage type is supported only with general-purpose and dedicated DB
      instances.
  + *ESSD*: extreme SSD storage.

  Changing this parameter will create a new resource. For details about volume types, see
  [DB Instance Storage Types](https://support.huaweicloud.com/intl/en-us/productdesc-rds/rds_01_0020.html).

* `disk_encryption_id` - (Optional) Specifies the key ID for disk encryption. Changing this parameter will create a new
  resource.

The `backup_strategy` block supports:

* `keep_days` - (Optional, Int) Specifies the retention days for specific backup files. The value range is from 0 to
  732. If this parameter is not specified or set to 0, the automated backup policy is disabled.

  -> **NOTE:** Primary/standby DB instances of Microsoft SQL Server do not support disabling the automated backup
  policy.

* `start_time` - (Required, String) Specifies the backup time window. Automated backups will be triggered during the
  backup time window. It must be a valid value in the **hh:mm-HH:MM**
  format. The current time is in the UTC format. The HH value must be 1 greater than the hh value. The values of mm and
  MM must be the same and must be set to any of the following: 00, 15, 30, or 45. Example value: 08:15-09:15 23:00-00:
  00.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

* `status` - Indicates the DB instance status.

* `created` - Indicates the creation time.

* `nodes` - Indicates the instance nodes information. Structure is documented below.

* `private_ips` - Indicates the private IP address list. It is a blank string until an ECS is created.

* `public_ips` - Indicates the public IP address list.

The `nodes` block contains:

* `availability_zone` - Indicates the AZ.

* `id` - Indicates the node ID.

* `name` - Indicates the node name.

* `role` - Indicates the node type. The value can be master or slave, indicating the primary node or standby node
  respectively.

* `status` - Indicates the node status.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 30 minute.
* `update` - Default is 30 minute.

## Import

RDS instance can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_rds_instance.instance_1 7117d38e-4c8f-4624-a505-bd96b97d024c
```

But due to some attrubutes missing from the API response, it's required to ignore changes as below.

```
resource "huaweicloud_rds_instance" "instance_1" {
  ...

  lifecycle {
    ignore_changes = [
      "db",
    ]
  }
}
```
