---
subcategory: "Deprecated"
---

# huaweicloud\_dms\_instance

!> **WARNING:** It has been deprecated, use `huaweicloud_dms_kafka_instance` or
`huaweicloud_dms_rabbitmq_instance` instead.

Manages a DMS instance in the huaweicloud DMS Service. This is an alternative to `huaweicloud_dms_instance_v1`

## Example Usage

### Automatically detect the correct network

```hcl
data "huaweicloud_dms_az" "az_1" {
}
data "huaweicloud_dms_product" "product_1" {
  engine        = "rabbitmq"
  instance_type = "single"
  version       = "3.7.17"
}

resource "huaweicloud_networking_secgroup" "secgroup_1" {
  name        = "secgroup_1"
  description = "secgroup_1"
}
resource "huaweicloud_dms_instance" "instance_1" {
  name              = var.instance_name
  engine            = "rabbitmq"
  access_user       = "user"
  password          = "Dmstest@123"
  vpc_id            = var.vpc_id
  subnet_id         = var.subnet_id
  security_group_id = huaweicloud_networking_secgroup.secgroup_1.id
  available_zones   = [
    data.huaweicloud_dms_az.az_1.id]
  product_id        = data.huaweicloud_dms_product.product_1.id
  engine_version    = data.huaweicloud_dms_product.product_1.version
  storage_space     = data.huaweicloud_dms_product.product_1.storage
  storage_spec_code = data.huaweicloud_dms_product.product_1.storage_spec_code
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the DMS instance resource. If omitted, the
  provider-level region will be used. Changing this creates a new DMS instance resource.

* `name` - (Required, String) Indicates the name of an instance. An instance name starts with a letter, consists of 4 to
  64 characters, and supports only letters, digits, and hyphens (-).

* `description` - (Optional, String) Indicates the description of an instance. It is a character string containing not
  more than 1024 characters.

* `engine` - (Optional, String, ForceNew) Indicates a message engine. Options: rabbitmq and kafka.

* `engine_version` - (Optional, String, ForceNew) Indicates the version of a message engine.

* `specification` - (Optional, String) This parameter is mandatory if the engine is kafka. Indicates the baseline
  bandwidth of a Kafka instance, that is, the maximum amount of data transferred per unit time. Unit: byte/s. Options:
  300 MB, 600 MB, 1200 MB.

* `storage_space` - (Required, Int) Indicates the message storage space. Value range:
  + Single-node RabbitMQ instance: 100–90000 GB
  + Cluster RabbitMQ instance: 100 GB x Number of nodes to 90000 GB, 200 GB x Number of nodes to 90000 GB, 300 GB x
      Number of nodes to 90000 GB
  + Kafka instance with specification being 300 MB: 1200–90000 GB
  + Kafka instance with specification being 600 MB: 2400–90000 GB
  + Kafka instance with specification being 1200 MB: 4800–90000 GB

* `storage_spec_code` - (Required, String) Indicates the storage I/O specification. Value range:

  Options for a RabbitMQ instance:
  + dms.physical.storage.normal
  + dms.physical.storage.high
  + dms.physical.storage.ultra

      Options for a Kafka instance:
  + When specification is 300 MB: dms.physical.storage.high or dms.physical.storage.ultra
  + When specification is 600 MB: dms.physical.storage.ultra
  + When specification is 1200 MB: dms.physical.storage.ultra

* `partition_num` - (Optional, Int) This parameter is mandatory when a Kafka instance is created. Indicates the maximum
  number of topics in a Kafka instance.
  + When specification is 300 MB: 900
  + When specification is 600 MB: 1800
  + When specification is 1200 MB: 1800

* `access_user` - (Optional, String) Indicates a username. If the engine is rabbitmq, this parameter is mandatory. If
  the engine is kafka, this parameter is optional. A username consists of 4 to 64 characters and supports only letters,
  digits, and hyphens (-).

* `password` - (Optional, String) If the engine is rabbitmq, this parameter is mandatory. If the engine is kafka, this
  parameter is mandatory when ssl_enable is true and is invalid when ssl_enable is false. Indicates the password of an
  instance. An instance password must meet the following complexity requirements: Must be 8 to 32 characters long. Must
  contain at least 2 of the following character types: lowercase letters, uppercase letters, digits, and special
  characters (`~!@#$%^&*()-_=+\|[{}]:'",<.>/?).

* `vpc_id` - (Required, String) Indicates the ID of a VPC.

* `subnet_id` - (Required, String) Indicates the ID of a subnet.

* `security_group_id` - (Required, String) Indicates the ID of a security group.

* `available_zones` - (Required, List) Indicates the ID of an AZ. The parameter value can not be left blank or an empty
  array. For details, see section Querying AZ Information.

* `product_id` - (Required, String) Indicates a product ID.

* `maintain_begin` - (Optional, String) Indicates the time at which a maintenance time window starts. Format: HH:mm:ss.
  The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance
  time window. For details, see section Querying Maintenance Time Windows. The start time must be set to 22:00, 02:00,
  06:00, 10:00, 14:00, or 18:00. Parameters maintain_begin and maintain_end must be set in pairs. If parameter
  maintain_begin is left blank, parameter maintain_end is also blank. In this case, the system automatically allocates
  the default start time 02:00.

* `maintain_end` - (Optional, String) Indicates the time at which a maintenance time window ends. Format: HH:mm:ss. The
  start time and end time of a maintenance time window must indicate the time segment of a supported maintenance time
  window. For details, see section Querying Maintenance Time Windows. The end time is four hours later than the start
  time. For example, if the start time is 22:00, the end time is 02:00. Parameters maintain_begin and maintain_end must
  be set in pairs. If parameter maintain_end is left blank, parameter maintain_begin is also blank. In this case, the
  system automatically allocates the default end time 06:00.

* `enable_publicip` - (Optional, Bool) Indicates whether to enable public access to a RabbitMQ instance. true: enable,
  false: disable

* `publicip_id` - (Optional, String) Indicates the ID of the elastic IP address (EIP) bound to a RabbitMQ instance. This
  parameter is mandatory if public access is enabled (that is, enable_publicip is set to true).

* `tags` - (Optional, Map) The key/value pairs to associate with the instance.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.
* `storage_space` - Indicates the time when a instance is created.
* `security_group_name` - Indicates the name of a security group.
* `subnet_name` - Indicates the name of a subnet.
* `subnet_cidr` - Indicates a subnet segment.
* `used_storage_space` - Indicates the used message storage space. Unit: GB
* `connect_address` - Indicates the IP address of an instance.
* `port` - Indicates the port number of an instance.
* `status` - Indicates the status of an instance. For details, see section Instance Status.
* `instance_id` - Indicates the ID of an instance.
* `resource_spec_code` - Indicates a resource specifications identifier.
* `type` - Indicates an instance type. Options: "single" and "cluster"
* `created_at` - Indicates the time when an instance is created. The time is in the format of timestamp, that is, the
  offset milliseconds from 1970-01-01 00:00:00 UTC to the specified time.
* `user_id` - Indicates a user ID.
* `user_name` - Indicates a username.
