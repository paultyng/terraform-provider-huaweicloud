---
subcategory: "Distributed Message Service (DMS)"
---

# huaweicloud_dms_rabbitmq_instance

## Example Usage

### Basic Instance

```hcl
data "huaweicloud_dms_az" "test" {}

data "huaweicloud_vpc" "test" {
  name = "vpc-default"
}

data "huaweicloud_vpc_subnet" "test" {
  name = "subnet-default"
}

data "huaweicloud_dms_product" "test" {
  engine        = "rabbitmq"
  instance_type = "cluster"
  version       = "3.7.17"
}

resource "huaweicloud_networking_secgroup" "test" {
  name        = "secgroup_1"
  description = "secgroup for rabbitmq"
}

resource "huaweicloud_dms_rabbitmq_instance" "test" {
  name              = "instance_1"
  description       = "rabbitmq test"
  access_user       = "user"
  password          = "Rabbitmqtest@123"
  vpc_id            = data.huaweicloud_vpc.test.id
  network_id        = data.huaweicloud_vpc_subnet.test.id
  security_group_id = huaweicloud_networking_secgroup.test.id
  available_zones   = [data.huaweicloud_dms_az.test.id]
  product_id        = data.huaweicloud_dms_product.test.id
  engine_version    = data.huaweicloud_dms_product.test.version
  storage_space     = data.huaweicloud_dms_product.test.storage
  storage_spec_code = data.huaweicloud_dms_product.test.storage_spec_code

  tags = {
    key   = "value"
    owner = "terraform"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the DMS rabbitmq instance resource.
  If omitted, the provider-level region will be used. Changing this creates a new instance resource.

* `name` - (Required, String) Specifies the name of the DMS rabbitmq instance. An instance name starts with a letter,
  consists of 4 to 64 characters, and supports only letters, digits, hyphens (-) and underscores (_).

* `description` - (Optional, String) Specifies the description of the DMS rabbitmq instance. It is a character
  string containing not more than 1024 characters.

* `engine_version` - (Optional, String, ForceNew) Specifies the version of the rabbitmq engine. Default to "3.7.17".
  Changing this creates a new instance resource.

* `storage_space` - (Required, Int, ForceNew) Specifies the message storage space. Value range:
  + Single-node RabbitMQ instance: 100–90000 GB
  + Cluster RabbitMQ instance: 100 GB x Number of nodes to 90000 GB, 200 GB x Number of nodes to 90000 GB, and 300 GB x Number of nodes to 90000 GB

  Changing this creates a new instance resource.

* `storage_spec_code` - (Required, String, ForceNew) Specifies the storage I/O specification. Value range:
  + dms.physical.storage.high
  + dms.physical.storage.ultra

  Changing this creates a new instance resource.

* `vpc_id` - (Required, String, ForceNew) Specifies the ID of a VPC. Changing this creates a new instance resource.

* `network_id` - (Required, String, ForceNew) Specifies the ID of a subnet.
  Changing this creates a new instance resource.

* `security_group_id` - (Required, String) Specifies the ID of a security group.

* `available_zones` - (Required, List, ForceNew) Specifies the ID of an AZ. The parameter value can not be
  left blank or an empty array. Changing this creates a new instance resource.

* `product_id` - (Required, String, ForceNew) Specifies a product ID. Changing this creates a new instance resource.

* `access_user` - (Required, String, ForceNew) Specifies a username.
  A username consists of 4 to 64 characters and supports only letters, digits, and hyphens (-).
  Changing this creates a new instance resource.

* `password` - (Required, String, ForceNew) Specifies the password of the DMS rabbitmq instance. A password
  must meet the following complexity requirements: Must be 8 to 32 characters long.
  Must contain at least 2 of the following character types: lowercase letters, uppercase
  letters, digits, and special characters (`~!@#$%^&*()-_=+\\|[{}]:'",<.>/?).
  Changing this creates a new instance resource.

* `maintain_begin` - (Optional, String) Specifies the time at which a maintenance time window starts.
  Format: HH:mm.
  The start time and end time of a maintenance time window must indicate the time segment of
  a supported maintenance time window. The start time must be set to 22:00, 02:00, 06:00, 10:00,
  14:00, or 18:00. Parameters `maintain_begin` and `maintain_end` must be set in pairs. If
  parameter `maintain_begin` is left blank, parameter `maintain_end` is also blank. In this case,
  the system automatically allocates the default start time 02:00.

* `maintain_end` - (Optional, String) Specifies the time at which a maintenance time window ends.
  Format: HH:mm.
  The start time and end time of a maintenance time window must indicate the time segment of
  a supported maintenance time window. The end time is four hours later than the start time.
  For example, if the start time is 22:00, the end time is 02:00. Parameters `maintain_begin`
  and `maintain_end` must be set in pairs. If parameter `maintain_end` is left blank, parameter
  `maintain_begin` is also blank. In this case, the system automatically allocates the default
  end time 06:00.

* `ssl_enable` - (Optional, String, ForceNew) Specifies whether to enable public access for the DMS rabbitmq instance.
  Changing this creates a new instance resource.

* `public_ip_id` - (Optional, String) Specifies the ID of the elastic IP address (EIP)
  bound to the DMS rabbitmq instance.

* `enterprise_project_id` - (Optional, String) Specifies the enterprise project ID of the rabbitmq instance.

* `tags` - (Optional, Map) The key/value pairs to associate with the DMS rabbitmq instance.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.
* `engine` - Indicates the message engine.
* `specification` - Indicates the instance specification. For a single-node DMS rabbitmq instance,
  VM specifications are returned. For a cluster DMS rabbitmq instance, VM specifications and the number of
  nodes are returned.
* `used_storage_space` - Indicates the used message storage space. Unit: GB
* `port` - Indicates the port number of the DMS rabbitmq instance.
* `status` - Indicates the status of the DMS rabbitmq instance.
* `enable_public_ip` - Indicates whether public access to the DMS rabbitmq instance is enabled.
* `resource_spec_code` - Indicates a resource specifications identifier.
* `type` - Indicates the DMS rabbitmq instance type.
* `user_id` - Indicates the ID of the user who created the DMS rabbitmq instance
* `user_name` - Indicates the name of the user who created the DMS rabbitmq instance
* `connect_address` - Indicates the IP address of the DMS rabbitmq instance.
* `manegement_connect_address` - Indicates the management address of the DMS rabbitmq instance.

## Import

DMS rabbitmq instance can be imported using the instance id, e.g.

```
 $ terraform import huaweicloud_dms_rabbitmq_instance.instance_1 8d3c7938-dc47-4937-a30f-c80de381c5e3
```

Note that the imported state may not be identical to your resource definition, due to some attrubutes missing from
the API response, security or some other reason. The missing attributes include:
`password`.
It is generally recommended running `terraform plan` after importing a DMS rabbitmq instance. You can then decide if
changes should be applied to the instance, or the resource definition should be updated to align with the instance.
Also you can ignore changes as below.

```
resource "huaweicloud_dms_rabbitmq_instance" "instance_1" {
    ...

  lifecycle {
    ignore_changes = [
      password,
    ]
  }
}
```
