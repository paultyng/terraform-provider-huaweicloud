---
subcategory: "Distributed Cache Service"
---

# huaweicloud_dcs_instance

Manages a DCS instance in the huaweicloud DCS Service.

## Example Usage

### DCS Single Instance

```hcl
data "huaweicloud_dcs_az" "az_1" {
  code = "cn-north-1a"
}

resource "huaweicloud_networking_secgroup" "secgroup_1" {
  name        = "secgroup_1"
  description = "secgroup_1"
}
resource "huaweicloud_vpc" "vpc_1" {
  name = "test_vpc1"
  cidr = "192.168.0.0/16"
}
resource "huaweicloud_vpc_subnet" "subnet_1" {
  name       = "test_subnet1"
  cidr       = "192.168.0.0/24"
  gateway_ip = "192.168.0.1"
  vpc_id     = huaweicloud_vpc.vpc_1.id
}

resource "huaweicloud_dcs_instance" "instance_1" {
  name              = "test_dcs_instance"
  engine            = "Redis"
  engine_version    = "5.0"
  password          = "Huawei_test"
  capacity          = 2
  vpc_id            = huaweicloud_vpc.vpc_1.id
  subnet_id         = huaweicloud_vpc_subnet.subnet_1.id
  security_group_id = huaweicloud_networking_secgroup.secgroup_1.id
  available_zones   = [
    data.huaweicloud_dcs_az.az_1.id]
  product_id        = "redis.single.xu1.large.2-h"
}
```

### DCS HA Instance

```hcl
resource "huaweicloud_dcs_instance" "instance_1" {
  name            = "test_dcs_instance"
  engine          = "Redis"
  engine_version  = "5.0"
  password        = "Huawei_test"
  capacity        = 2
  vpc_id          = huaweicloud_vpc.vpc_1.id
  subnet_id       = huaweicloud_vpc_subnet.subnet_1.id
  available_zones = [
    data.huaweicloud_dcs_az.az_1.id]
  product_id      = "redis.ha.au1.large.r2.2-h"

  backup_policy {
    save_days   = 1
    backup_type = "manual"
    begin_at    = "00:00-01:00"
    period_type = "weekly"
    backup_at   = [
      1,
      2,
      4,
      6]
  }

  whitelists {
    group_name = "test-group1"
    ip_address = [
      "192.168.10.100",
      "192.168.0.0/24"]
  }
  whitelists {
    group_name = "test-group2"
    ip_address = [
      "172.16.10.100",
      "172.16.0.0/24"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the DCS instance resource. If omitted,
  the provider-level region will be used. Changing this creates a new DCS instance resource.

* `name` - (Required, String) Specifies the name of an instance. It is a string with 4 to 64 characters that contains
  letters, digits, underscores (_), and hyphens (-) and starts with a letter.

* `description` - (Optional, String) Specifies the description of an instance. It is a string that contains a maximum of
  1024 characters.

* `engine` - (Required, String, ForceNew) Specifies a cache engine. Options: *Redis* and *Memcached*. Changing this
  creates a new instance.

* `engine_version` - (Optional, String, ForceNew) Specifies the version of a cache engine. It is mandatory when the
  engine is *Redis*, the value can be 3.0, 4.0, or 5.0. Changing this creates a new instance.

* `capacity` - (Required, Float, ForceNew) Specifies the cache capacity. Unit: GB.
  + Redis4.0 and Redis5.0: Stand-alone and active/standby type instance values:
      0.125, 0.25, 0.5, 1, 2, 4, 8, 16, 32, 64. Cluster instance specifications support 24, 32, 48, 64, 96, 128, 192,
      256, 384, 512, 768, 1024.

  + Redis3.0: Stand-alone and active/standby type instance values: 2, 4, 8, 16, 32, 64. Proxy cluster instance
      specifications support 64, 128, 256, 512, and 1024.

  + Memcached: Stand-alone and active/standby type instance values: 2, 4, 8, 16, 32, 64.

* `product_id` - (Required, String, ForceNew) Specifies the product ID with various cache modes:
  *single*, *ha*, *cluster*, *proxy* and *ha_rw_split*. The value format is *spec_code* + "-h", indicates that the
  charging mode is pay-per-use. Changing this creates a new instance. You can query the *spec_code* as follows:
  + Query the specifications
      in [DCS Instance Specifications](https://support.huaweicloud.com/intl/en-us/productdesc-dcs/dcs-pd-200713003.html)
  + Log in to the DCS console, click *Buy DCS Instance*, and find the corresponding instance specification.

* `vpc_id` - (Required, String, ForceNew) Specifies the id of the VPC. Changing this creates a new instance.

* `subnet_id` - (Required, String, ForceNew) Specifies the id of the subnet. Changing this creates a new instance.

* `security_group_id` - (Optional, String) Specifies the id of the security group which the instance belongs to. This
  parameter is mandatory for Memcached and Redis 3.0 version.

* `whitelists` - (Optional, List) Specifies the IP addresses which can access the instance. This parameter is valid for
  Redis 4.0 and 5.0 versions. The structure is described below.

* `whitelist_enable` - (Optional, Bool) Enable or disable the IP addresse whitelists. Default to true. If the whitelist
  is disabled, all IP addresses connected to the VPC can access the instance.

* `available_zones` - (Required, List, ForceNew) Specifies IDs of the AZs where cache nodes reside. If you are creating
  active/standby, Proxy cluster, and Cluster cluster instances to support cross-zone deployment, you can specify the
  standby zone for the standby node. Changing this creates a new instance.

* `access_user` - (Optional, String, ForceNew) Specifies the username used for accessing a DCS instance. If the cache
  engine is *Redis*, you do not need to set this parameter. A username starts with a letter, consists of 1 to 64
  characters, and supports only letters, digits, and hyphens (-). Changing this creates a new instance.

* `password` - (Optional, String, ForceNew) Specifies the password of a DCS instance. Changing this creates a new
  instance. The password of a DCS Redis instance must meet the following complexity requirements:
  + Enter a string of 8 to 32 bits in length.
  + The new password cannot be the same as the old password.
  + Must contain three combinations of the following four characters: Lower case letters, uppercase letter, digital,
      Special characters include (`~!@#$%^&*()-_=+|[{}]:'",<.>/?).

* `maintain_begin` - (Optional, String) Specifies the time at which a maintenance time window starts. Format: HH:mm:ss.
  + The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance
      time window.
  + The start time must be set to 22:00:00, 02:00:00, 06:00:00, 10:00:00, 14:00:00, or 18:00:00.
  + Parameters `maintain_begin` and `maintain_end` must be set in pairs.
  + If parameter maintain_begin is left blank, parameter maintain_end is also blank. In this case, the system
      automatically allocates the default start time 02:00:00.

* `maintain_end` - (Optional, String) Specifies the time at which a maintenance time window ends. Format: HH:mm:ss.
  + The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance
      time window.
  + The end time is four hours later than the start time. For example, if the start time is 22:00:00, the end time is
      02:00:00.
  + Parameters `maintain_begin` and `maintain_end` must be set in pairs.
  + If parameter maintain_end is left blank, parameter maintain_begin is also blank. In this case, the system
      automatically allocates the default end time 06:00:00.

* `backup_policy` - (Optional, List) Specifies the backup configuration to be used with the instance. The structure is
  described below.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id of the dcs instance. Changing this
  creates a new instance.

* `tags` - (Optional, Map) The key/value pairs to associate with the dcs instance.

The `whitelists` block supports:

* `group_name` - (Required, String) Specifies the name of IP address group.

* `ip_address` - (Required, List) Specifies the list of IP address or CIDR which can be whitelisted for an instance.

The `backup_policy` block supports:

* `save_days` - (Optional, Int) Retention time. Unit: day. Range: 1–7.

* `backup_type` - (Optional, String) Backup type. Options:
  + *auto*: automatic backup
  + *manual*: manual backup (default)

* `begin_at` - (Required, String) Time at which backup starts. "00:00-01:00" indicates that backup starts at 00:00:00.

* `period_type` - (Required, String) Interval at which backup is performed. Currently, only weekly backup is supported.

* `backup_at` - (Required, List) Day in a week on which backup starts. Range: 1–7. Where: 1 indicates Monday; 7
  indicates Sunday.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - A resource ID in UUID format.
* `vpc_name` - Indicates the name of a vpc.
* `subnet_name` - Indicates the name of a subnet.
* `security_group_name` - Indicates the name of a security group.
* `resource_spec_code` - Resource specification code.
* `used_memory` - Size of the used memory. Unit: MB.
* `internal_version` - Internal DCS version.
* `max_memory` - Overall memory size. Unit: MB.
* `user_id` - Indicates a user ID.
* `user_name` - Username.
* `ip` - Cache node's IP address in tenant's VPC.
* `port` - Port of the cache node.
