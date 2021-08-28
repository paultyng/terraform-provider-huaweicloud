---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_vpc

Use this data source to get the details of a specific IEC VPC.

## Example Usage

```hcl
variable "vpc_name" {}

data "huaweicloud_iec_vpc" "my_vpc" {
  name = var.vpc_name
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to obtain the vpc. If omitted, the provider-level region
  will be used.

* `name` - (Optional, String) Specifies the name of the IEC VPC. The name can contain a maximum of 64 characters. Only
  letters, digits, underscores (_), hyphens (-), and periods (.) are allowed.

* `id` - (Optional, String) Specifies the ID of the IEC VPC to retrieve.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cidr` - Indicates the IP address range for the VPC.
* `mode` - Indicates the mode of the IEC VPC. Possible values are *SYSTEM* and *CUSTOMER*.
* `subent_num` - Indicates the number of subents.
