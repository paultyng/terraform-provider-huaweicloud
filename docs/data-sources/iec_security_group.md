---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_security_group

Use this data source to get the details of a specific IEC security group.

## Example Usage

```hcl
variable "sg_name" {}

data "huaweicloud_iec_security_group" "my_sg" {
  name = var.sg_name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Specifies the name of the security group with a maximum of 64 characters.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - A data source ID in UUID format.
* `description` - The description of the IEC security group.
* `security_group_rules` - An Array of one or more security group rules. The object is documented below.

The `security_group_rules` block supports:

* `id` - The ID of the IEC security group rules.
* `security_group_id` - The id of the IEC security group rules.
* `description` - The description for the IEC security group rules.
* `direction` - The direction of the IEC security group rules.
* `ethertype` - The layer 3 protocol type.
* `port_range_max` - The higher part of the allowed port range.
* `port_range_min` - The lower part of the allowed port range.
* `protocol` - The layer 4 protocol type.
* `remote_ip_prefix` - The remote CIDR of the IEC security group rules.
* `remote_group_id` - The remote group id of the IEC security group rules.
