---
subcategory: "Relational Database Service (RDS)"
---

# huaweicloud_rds_read_replica_instance

Manage RDS Read Replica Instance resource.

## Example Usage

### Create a Rds read replica instance

```hcl
resource "huaweicloud_networking_secgroup" "secgroup" {
  name          = "test_sg_for_rds"
  description   = "security group for rds read replica instance"
}

resource "huaweicloud_rds_instance" "instance" {
  name                  = "terraform_test_rds_instance"
  flavor                = "rds.pg.n1.large.2"
  availability_zone     = ["{{ availability_zone }}"]
  vpc_id                = "{{ vpc_id }}"
  subnet_id             = "{{ subnet_id }}"
  security_group_id     = huaweicloud_networking_secgroup.secgroup.id
  enterprise_project_id = "{{ enterprise_project_id }}"

  db {
    type        = "PostgreSQL"
    version     = "12"
    password    = "Huangwei!120521"
    port        = "8635"
  }
  volume {
    type = "ULTRAHIGH"
    size = 50
  }
  backup_strategy {
    start_time  = "08:00-09:00"
    keep_days   = 1
  }
}

resource "huaweicloud_rds_read_replica_instance" "replica_instance" {
  name                  = "test_rds_readonly_instance"
  flavor                = "rds.pg.n1.large.2.rr"
  primary_instance_id   = huaweicloud_rds_instance.instance.id
  availability_zone     = "{{ availability_zone }}"
  enterprise_project_id = "{{ enterprise_project_id }}"
  volume {
    type = "ULTRAHIGH"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the rds read replica instance resource. If
  omitted, the provider-level region will be used.

  Currently, read replicas can be created *only* in the same region as that of the primary DB instance.

* `availability_zone` - (Required, String, ForceNew) Specifies the AZ name. Changing this parameter will create a new
  resource.

* `name` - (Required, String, ForceNew) Specifies the DB instance name. The DB instance name of the same type must be
  unique for the same tenant. The value must be 4 to 64 characters in length and start with a letter. It is
  case-sensitive and can contain only letters, digits, hyphens (-), and underscores (_). Changing this parameter will
  create a new resource.

* `flavor` - (Required, String) Specifies the specification code.

* `primary_instance_id` - (Required, String, ForceNew) Specifies the DB instance ID, which is used to create a read
  replica. Changing this parameter will create a new resource.

* `volume` - (Required, List, ForceNew) Specifies the volume information. Structure is documented below. Changing this
  parameter will create a new resource.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id of the read replica instance.
  Changing this parameter will create a new resource.

* `tags` - (Optional, Map) A mapping of tags to assign to the RDS read replica instance. Each tag is represented by one
  key-value pair.

The `volume` block supports:

* `type` - (Required, String, ForceNew) Specifies the volume type. Its value can be any of the following and is
  case-sensitive:
  + *ULTRAHIGH*: SSD storage.
  + *LOCALSSD*: local SSD storage.
  + *CLOUDSSD*: cloud SSD storage. This storage type is supported only with general-purpose and dedicated DB
      instances.
  + *ESSD*: extreme SSD storage.

  Changing this parameter will create a new resource.

* `disk_encryption_id` - (Optional, String, ForceNew) Specifies the key ID for disk encryption. Changing this parameter
  will create a new resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Indicates the instance ID.

* `status` - Indicates the instance status.

* `db` - Indicates the database information. Structure is documented below.

* `private_ips` - Indicates the private IP address list.

* `public_ips` - Indicates the public IP address list.

* `security_group_id` - Indicates the security group which the RDS DB instance belongs to.

* `subnet_id` - Indicates the subnet id.

* `vpc_id` - Indicates the VPC ID.

The `db` block supports:

* `port` - Indicates the database port information.

* `type` - Indicates the DB engine. Value: MySQL, PostgreSQL, SQLServer.

* `user_name` - Indicates the default user name of database.

* `version` - Indicates the database version.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 30 minute.
* `delete` - Default is 30 minute.

## Import

RDS read replica instance can be imported by `id`, e.g.

```shell
$ terraform import huaweicloud_rds_read_replica_instance.replica_instance 92302c133d13424cbe357506ce057ea5in03
```
