---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_security_group

Manages a IEC security group resource within HuaweiCloud.

## Example Usage

```hcl
variable "iec_secgroup_name" {}

resource "huaweicloud_iec_security_group" "secgroup_test" {
  name = var.iec_secgroup_name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String, ForceNew) Specifies the name for the security group. This parameter can contain a maximum
  of 64 characters, which may consist of letters, digits, dot (.), underscores (_), and hyphens (-). The iec security
  group allowed to have the same name. Changing this parameter will creates a new iec security group resource.

* `description` - (Optional, String, ForceNew) Specifies the description of the iec security group. description must be
  0 to 64 characters in length, and does not contain angle brackets (<) and (>). Changing this parameter will creates a
  new iec security group resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - A resource ID in UUID format.

* `security_group_rules` - An Array of one or more security group rules. The security_group_rules object structure is
  documented below.

The `security_group_rules` block supports:

* `id` - The id of the iec security group rules.
* `security_group_id` - The id of the iec security group rules.
* `description` - The description for the iec security group rules.
* `direction` - The direction of the iec security group rules.
* `ethertype` - The layer 3 protocol type.
* `port_range_max` - The higher part of the allowed port range.
* `port_range_min` - The lower part of the allowed port range.
* `protocol` - The layer 4 protocol type.
* `remote_ip_prefix` - The remote CIDR of the iec security group rules.
* `remote_group_id` - The remote group id of the iec security group rules.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

IEC Security Groups can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_iec_security_group.secgroup_test 2a02d1d3-437c-11eb-b721-fa163e8ac569
```
