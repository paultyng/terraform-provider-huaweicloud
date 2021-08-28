---
subcategory: "Deprecated"
---

# huaweicloud\_ecs\_instance\_v1

!> **Warning:** It has been deprecated, use `huaweicloud_compute_instance` instead.

Manages a ECS instance resource within HuaweiCloud.

## Example Usage

### Basic Instance

```hcl
resource "huaweicloud_ecs_instance_v1" "basic" {
  name     = "server_1"
  image_id = "ad091b52-742f-469e-8f3c-fd81cadf0743"
  flavor   = "s1.medium"
  vpc_id   = "8eed4fc7-e5e5-44a2-b5f2-23b3e5d46235"

  nics {
    network_id = "55534eaa-533a-419d-9b40-ec427ea7195a"
  }

  availability_zone = "cn-north-1a"
  key_name          = "KeyPair-test"
  security_groups   = ["default"]
}
```

### Instance with Data Disks

```hcl
resource "huaweicloud_ecs_instance_v1" "basic" {
  name     = "server_1"
  image_id = "ad091b52-742f-469e-8f3c-fd81cadf0743"
  flavor   = "s1.medium"
  vpc_id   = "8eed4fc7-e5e5-44a2-b5f2-23b3e5d46235"

  nics {
    network_id = "55534eaa-533a-419d-9b40-ec427ea7195a"
  }

  system_disk_type = "SAS"
  system_disk_size = 40

  data_disks {
    type = "SATA"
    size = "10"
  }
  data_disks {
    type = "SAS"
    size = "20"
  }

  delete_disks_on_termination = true
  availability_zone           = "cn-north-1a"
  key_name                    = "KeyPair-test"
  security_groups             = ["default"]
}
```

### Instance With Attached Volume

```hcl
resource "huaweicloud_blockstorage_volume_v2" "myvol" {
  name = "myvol"
  size = 1
}

resource "huaweicloud_ecs_instance_v1" "basic" {
  name     = "server_1"
  image_id = "ad091b52-742f-469e-8f3c-fd81cadf0743"
  flavor   = "s1.medium"
  vpc_id   = "8eed4fc7-e5e5-44a2-b5f2-23b3e5d46235"

  nics {
    network_id = "55534eaa-533a-419d-9b40-ec427ea7195a"
  }

  availability_zone = "cn-north-1a"
  key_name          = "KeyPair-test"
  security_groups   = ["default"]
}

resource "huaweicloud_compute_volume_attach" "attached" {
  instance_id = huaweicloud_ecs_instance_v1.basic.id
  volume_id   = huaweicloud_blockstorage_volume_v2.myvol.id
}
```

### Instance With Multiple Networks

```hcl
resource "huaweicloud_networking_floatingip_v2" "myip" {
}

resource "huaweicloud_ecs_instance_v1" "multi-net" {
  name     = "server_1"
  image_id = "ad091b52-742f-469e-8f3c-fd81cadf0743"
  flavor   = "s1.medium"
  vpc_id   = "8eed4fc7-e5e5-44a2-b5f2-23b3e5d46235"

  nics {
    network_id = "55534eaa-533a-419d-9b40-ec427ea7195a"
  }

  nics {
    network_id = "2c0a74a9-4395-4e62-a17b-e3e86fbf66b7"
  }

  availability_zone = "cn-north-1a"
  key_name          = "KeyPair-test"
  security_groups   = ["default"]
}

resource "huaweicloud_compute_eip_associate" "myip" {
  floating_ip = huaweicloud_networking_floatingip_v2.myip.address
  instance_id = huaweicloud_ecs_instance_v1.multi-net.id
  fixed_ip    = huaweicloud_ecs_instance_v1.multi-net.nics.0.ip_address
}
```

### Instance with User Data (cloud-init)

```hcl
resource "huaweicloud_ecs_instance_v1" "basic" {
  name     = "server_1"
  image_id = "ad091b52-742f-469e-8f3c-fd81cadf0743"
  flavor   = "s1.medium"
  vpc_id   = "8eed4fc7-e5e5-44a2-b5f2-23b3e5d46235"

  nics {
    network_id = "55534eaa-533a-419d-9b40-ec427ea7195a"
  }

  user_data         = "#cloud-config\nhostname: server_1.example.com\nfqdn: server_1.example.com"
  key_name          = "KeyPair-test"
  security_groups   = ["default"]
}
```

`user_data` can come from a variety of sources: inline, read in from the `file`
function, or the `template_cloudinit_config` resource.

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) A unique name for the instance.

* `image_id` - (Required, String, ForceNew) The ID of the desired image for the server. Changing this creates a new server.

* `flavor` - (Required, String) The name of the desired flavor for the server. Changing this resizes the existing server.

* `user_data` - (Optional, String, ForceNew) The user data to provide when launching the instance.
    Changing this creates a new server.

* `password` - (Optional, String, ForceNew) The administrative password to assign to the server.
    Changing this creates a new server.

* `key_name` - (Optional, String, ForceNew) The name of a key pair to put on the server. The key
    pair must already be created and associated with the tenant's account.
    Changing this creates a new server.

* `vpc_id` - (Required, String, ForceNew) The ID of the desired VPC for the server. Changing this creates a new server.

* `nics` - (Optional, List, ForceNew) An array of one or more networks to attach to the
    instance. The nics object structure is documented below. Changing this
    creates a new server.

* `system_disk_type` - (Optional, String, ForceNew) The system disk type of the server. For HANA, HL1, and HL2 ECSs use co-p1 and uh-l1 disks.
    Changing this creates a new server. Available options are:
	+ `SATA`: common I/O disk type.
	+ `SAS`: high I/O disk type.
	+ `SSD`: ultra-high I/O disk type.
	+ `co-p1`: high I/O(performance-optimized) disk type.
	+ `uh-l1`: ultra-high I/O(latency-optimized) disk type.

* `system_disk_size` - (Optional, Int, ForceNew) The system disk size in GB, The value range is 1 to 1024. Changing this creates a new server.

* `data_disks` - (Optional, List, ForceNew) An array of one or more data disks to attach to the
    instance. The data_disks object structure is documented below. Changing this
    creates a new server.

* `security_groups` - (Optional, String) An array of one or more security group names
    to associate with the server. Changing this results in adding/removing
    security groups from the existing server.

* `availability_zone` - (Required, String, ForceNew) The availability zone in which to create
    the server. Please refer to [endpoint reference](https://developer.huaweicloud.com/endpoint)
    for the values. Changing this creates a new server.

* `charging_mode` - (Optional, String, ForceNew) The charging mode of the instance. Valid options are: prePaid and postPaid, defaults to postPaid. Changing this creates a new server.

* `period_unit` - (Optional, String, ForceNew) The charging period unit of the instance. Valid options are: month and year, defaults to month. Changing this creates a new server.

* `period` - (Optional, Int, ForceNew) The charging period of the instance. Changing this creates a new server.

* `auto_renew` - (Optional, String, ForceNew) Specifies whether auto renew is enabled. Changing this creates a new server.

* `auto_recovery` - (Optional, Bool) Whether configure automatic recovery of an instance.

* `delete_disks_on_termination` - (Optional, Bool) Delete the data disks upon termination of the instance. Defaults to false. Changing this creates a new server.

* `enterprise_project_id` - (Optional, String) The enterprise project id. Changing this creates a new server.

* `tags` - (Optional, Map) Tags key/value pairs to associate with the instance.

* `op_svc_userid` - (Optional, String, ForceNew) User ID, required when using key_name. Changing this creates a new server.


The `nics` block supports:

* `network_id` - (Required, String, ForceNew) The network UUID to attach to the server. Changing this creates a new server.

* `ip_address` - (Optional, String, ForceNew) Specifies a fixed IPv4 address to be used on this
    network. Changing this creates a new server.

The `data_disks` block supports:

* `type` - (Required, String, ForceNew) The data disk type of the server. For HANA, HL1, and HL2 ECSs use co-p1 and uh-l1 disks.
    Changing this creates a new server. Available options are:
	+ `SATA`: common I/O disk type.
	+ `SAS`: high I/O disk type.
	+ `SSD`: ultra-high I/O disk type.
	+ `co-p1`: high I/O(performance-optimized) disk type.
	+ `uh-l1`: ultra-high I/O(latency-optimized) disk type.

* `size` - (Required, Int, ForceNew) The size of the data disk in GB. The value range is 10 to 32768.
    Changing this creates a new server.

* `snapshot_id` - (Optional, String, ForceNew) Specifies the snapshot ID or ID of the original data disk contained in the full-ECS image.
    Changing this creates a new server.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the server.
* `nics/mac_address` - The MAC address of the NIC on that network.
* `nics/port_id` - The port ID of the NIC on that network.

## Timeouts
This resource provides the following timeouts configuration options:
* `create` - Default is 30 minute.
* `update` - Default is 30 minute.
* `delete` - Default is 30 minute.

## Import

Instances can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_ecs_instance_v1.instance_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9
