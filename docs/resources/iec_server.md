---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_server

Manages a IEC server resource within HuaweiCloud.

## Example Usage

### Basic Server Instance

```hcl
variable "iec_server_name" {}
variable "iec_iamge_id" {}
variable "iec_flavor_id" {}
variable "iec_site_id" {}
variable "iec_site_operator" {}
variable "iec_vpc_id" {}
variable "iec_subnet_id" {}
variable "iec_secgroup_id" {}
variable "iec_server_password" {}

resource "huaweicloud_iec_server" "server_test" {
  name            = var.iec_server_name
  image_id        = var.iec_iamge_id
  flavor_id       = var.iec_flavor_id
  vpc_id          = var.iec_vpc_id
  subnet_ids      = [
    var.iec_subnet_id]
  security_groups = [
    var.iec_secgroup_id]

  admin_pass       = var.iec_server_password
  bind_eip         = true
  system_disk_type = "SAS"
  system_disk_size = 40

  coverage_sites {
    site_id  = var.iec_site_id
    operator = var.iec_site_operator
  }
}
```

### Server Instance With Multiple Data Disks

```hcl
variable "iec_server_name" {}
variable "iec_iamge_id" {}
variable "iec_flavor_id" {}
variable "iec_site_id" {}
variable "iec_site_operator" {}
variable "iec_vpc_id" {}
variable "iec_subnet_id" {}
variable "iec_secgroup_id" {}
variable "iec_server_password" {}

resource "huaweicloud_iec_server" "server_test" {
  name            = var.iec_server_name
  image_id        = var.iec_iamge_id
  flavor_id       = var.iec_flavor_id
  vpc_id          = var.iec_vpc_id
  subnet_ids      = [
    var.iec_subnet_id]
  security_groups = [
    var.iec_secgroup_id]

  admin_pass       = var.iec_server_password
  bind_eip         = true
  system_disk_type = "SAS"
  system_disk_size = 40

  data_disks {
    type = "SAS"
    size = "20"
  }
  data_disks {
    type = "SAS"
    size = "40"
  }

  coverage_sites {
    site_id  = var.iec_site_id
    operator = var.iec_site_operator
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String, ForceNew) Specifies the IEC server name. This parameter can contain a maximum of 64
  characters, which may consist of letters, digits, dot(.), underscores (_), and hyphens (-). Changing this parameter
  creates a new IEC server resource.

* `flavor_id` - (Required, String, ForceNew) Specifies the flavor ID of the desired flavor for the IEC server. Changing
  this parameter creates a new IEC server resource.

* `image_id` - (Required, String, ForceNew) Specifies the image ID of the desired image for the IEC server. Changing
  this parameter creates a new IEC server resource.

* `vpc_id` - (Required, String, ForceNew) Specifies the ID of vpc for the IEC server. VPC mode only *CUSTOMER* can be
  used to create IEC server. Changing this parameter creates a new IEC server resource.

* `subnet_ids` - (Required, List, ForceNew) Specifies an array of one or more subnet ID of Network for the IEC server
  binding. Changing this parameter creates a new IEC server resource.

* `security_groups` - (Required, List, ForceNew) Specifies an array of one or more security group IDs to associate with
  the IEC server. Changing this parameter creates a new IEC server resource.

* `system_disk_type` - (Required, String, ForceNew) Specifies the type of system disk for the IEC server binding. Valid
  value is *SAS*(high I/O disk type). Changing this parameter creates a new IEC server resource.

* `system_disk_size` - (Required, String, ForceNew) Specifies the size of system disk for the IEC server binding. The
  value range is 40 to 100 in GB. Changing this parameter creates a new IEC server resource.

* `coverage_sites` - (Required, List, ForceNew) Specifies an array of site ID and operator for the IEC server. The
  object structure is documented below. Changing this parameter creates a new IEC server resource.

* `admin_pass` - (Optional, String, ForceNew) Specifies the administrative password to assign to the IEC server. This
  parameter can contain a maximum of 26 characters, which may consist of letters, digits and Special characters(~!?,.:
  ;-_'"(){}[]/<>@#$%^&*+|\\=) and space. This parameter and `key_pair` are alternative. Changing this changes the root
  password on the existing server.

* `key_pair` - (Optional, String, ForceNew) Specifies the name of a key pair to put on the IEC server. The key pair must
  already be created and associated with the tenant's account. This parameter and `admin_pass` are alternative. Changing
  this parameter creates a new IEC server resource.

* `bind_eip` - (Optional, String, ForceNew) Specifies whether the IEC server is bound to EIP. Changing this parameter
  creates a new IEC server resource.

* `coverage_level` - (Optional, String, ForceNew) Specifies the coverage level of IEC sites. Valid value is *SITE*.
  Changing this parameter creates a new IEC server resource.

* `coverage_policy` - (Optional, String, ForceNew) Specifies the policy of IEC sites. Valid values are *centralize*
  and *discrete*, *centralize* is default. Changing this parameter creates a new IEC server resource.

* `data_disks` - (Optional, List, ForceNew) Specifies the array of data disks to attach to the IEC server. Up to two
  data disks can be specified. The object structure is documented below. Changing this parameter creates a new IEC
  server resource.

* `user_data` - (Optional, String, ForceNew) Specifies the user data (information after encoding) configured during IEC
  server creation. The value can come from a variety of sources: inline, read in from the *file* function. Changing this
  parameter creates a new IEC server resource.

The `coverage_sites` block supports:

* `site_id` - (Required, String, ForceNew) Specifies the ID of IEC site.
* `operator` - (Required, String, ForceNew) Specifies the operator of the IEC site.

The `data_disks` block supports:

* `type` - (Required, String, ForceNew) Specifies the type of data disk for the IEC server binding. Valid value is *
  SAS*(high I/O disk type). Changing this parameter creates a new IEC server resource.
* `size` - (Required, String, ForceNew) Specifies the size of data disk for the IEC server binding. The value range is
  10 to 500 in GB. Changing this parameter creates a new IEC server resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `edgecloud_id` - The ID of the edgecloud service.
* `edgecloud_name` - The Name of the edgecloud service.
* `image_name` - The image name of the IEC server.
* `flavor_name` - The flavor name of the IEC server.
* `nics` - An array of one or more networks to attach to the IEC server. The object structure is documented below.
* `volume_attached` - An array of one or more disks to attach to the IEC server. The object structure is documented
  below.
* `public_ip` - The EIP address that is associted to the IEC server.
* `system_disk_id` - The system disk voume ID.
* `origin_server_id` - The ID of origin server.
* `status` - The status of IEC server.

The `nics` block supports:

* `port` - The port ID corresponding to the IP address on that network.
* `mac` - The MAC address of the NIC on that network.
* `address` - The IPv4 address of the server on that network.

The `volume_attached` block supports:

* `volume_id` - The volume ID on that attachment.
* `boot_index` - The volume boot index on that attachment.
* `size` - The volume size on that attachment.
* `type` - The volume type on that attachment.
* `device` - The device name in the IEC server.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 30 minute.
* `update` - Default is 30 minute.
* `delete` - Default is 30 minute.
