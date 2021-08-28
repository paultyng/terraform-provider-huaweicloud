---
subcategory: "Scalable File Service (SFS)"
---

# huaweicloud_sfs_turbo

Provides an Shared File System (SFS) Turbo resource.

## Example Usage

```hcl
variable "vpc_id" {}
variable "subnet_id" {}
variable "secgroup_id" {}
variable "test_az" {}

resource "huaweicloud_sfs_turbo" "sfs-turbo-1" {
  name              = "sfs-turbo-1"
  size              = 500
  share_proto       = "NFS"
  vpc_id            = var.vpc_id
  subnet_id         = var.subnet_id
  security_group_id = var.secgroup_id
  availability_zone = var.test_az
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the SFS Turbo resource. If omitted, the
  provider-level region will be used. Changing this creates a new SFS Turbo resource.

* `name` - (Required, String, ForceNew) Specifies the name of an SFS Turbo file system. The value contains 4 to 64
  characters and must start with a letter. Changing this will create a new resource.

* `size` - (Required, Int) Specifies the capacity of a common file system, in GB. The value ranges from 500 to 32768,
  and must be large than 10240 for an enhanced file system.

* `share_proto` - (Optional, String, ForceNew) Specifies the protocol for sharing file systems. The valid value is NFS.
  Changing this will create a new resource.

* `share_type` - (Optional, String, ForceNew) Specifies the file system type. The valid values are STANDARD and
  PERFORMANCE Changing this will create a new resource.

* `availability_zone` - (Required, String, ForceNew) Specifies the availability zone where the file system is located.
  Changing this will create a new resource.

* `vpc_id` - (Required, String, ForceNew) Specifies the VPC ID. Changing this will create a new resource.

* `subnet_id` - (Required, String, ForceNew) Specifies the network ID of the subnet. Changing this will create a new
  resource.

* `security_group_id` - (Required, String, ForceNew) Specifies the security group ID. Changing this will create a new
  resource.

* `enhanced` - (Optional, Bool, ForceNew) Specifies whether the file system is enhanced or not. Changing this will
  create a new resource.

* `crypt_key_id` - (Optional, String, ForceNew) Specifies the ID of a KMS key to encrypt the file system. Changing this
  will create a new resource.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id of the file system. Changing this
  will create a new resource.

-> **NOTE:**
SFS Turbo will create two private IP addresses and one virtual IP address under the subnet you specified. To ensure
normal use, SFS Turbo will enable the inbound rules for ports *111*, *445*, *2049*, *2051*, *2052*, and *20048* in the
security group you specified.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The UUID of the SFS Turbo file system.

* `region` - The region of the SFS Turbo file system.

* `status` - The status of the SFS Turbo file system.

* `version` - The version ID of the SFS Turbo file system.

* `export_location` - Tthe mount point of the SFS Turbo file system.

* `available_capacity` - The available capacity of the SFS Turbo file system in the unit of GB.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

SFS Turbo can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_sfs_turbo 1e3d5306-24c9-4316-9185-70e9787d71ab
```
