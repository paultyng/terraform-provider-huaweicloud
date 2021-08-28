---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_vpc

Manages an IEC VPC resource within HuaweiCloud.

## Example Usage

```hcl
variable "vpc_name" {}
variable "vpc_cidr" {}

resource "huaweicloud_iec_vpc" "vpc" {
  name = var.vpc_name
  cidr = var.vpc_cidr
}

resource "huaweicloud_iec_vpc" "vpc_by_customer" {
  name = var.vpc_name
  cidr = var.vpc_cidr
  mode = "CUSTOMER"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the IEC VPC. If omitted, the provider-level
  region will be used. Changing this creates a new resource.

* `name` - (Required, String) Specifies the name of the IEC VPC. The name can contain a maximum of 64 characters. Only
  letters, digits, underscores (_), hyphens (-), and periods (.) are allowed.

* `cidr` - (Required, String) Specifies the IP address range for the VPC. The subnet IP address in the VPC must be
  within the IP address range of the VPC. The following CIDR blocks are supported:
  *10.0.0.0/8-16*, *172.16.0.0/12-16*, *192.168.0.0/16*.

* `mode` - (Optional, String, ForceNew) Specifies the mode of the IEC VPC. Possible values are "SYSTEM" and "CUSTOMER",
  defaults to "SYSTEM". Changing this creates a new IEC VPC.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the IEC VPC.
* `subent_num` - Indicates the number of subents.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 3 minute.

## Import

VPCs can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_iec_vpc.myvpc 7117d38e-4c8f-4624-a505-bd96b97d024c
```
