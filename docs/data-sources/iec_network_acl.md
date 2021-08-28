---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_network_acl

Use this data source to get the details of a specific IEC network ACL.

## Example Usage

```hcl
variable "acl_name" {}

data "huaweicloud_iec_network_acl" "firewall" {
  name = var.acl_name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, String) Specifies the name if the IEC network ACL with a maximum of 64 characters.

* `id` - (Optional, String) Specifies the ID of the IEC network ACL to retrieve.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - The description of the IEC network ACL.
* `status` - The status of the IEC network ACL.
