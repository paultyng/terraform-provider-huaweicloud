---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_security_group_rule

Manages a IEC security group rule resource within HuaweiCloud.

## Example Usage

```hcl
var "iec_security_group_id" {}

resource "huaweicloud_iec_security_group_rule" "secgroup_rule_test" {
  direction         = "ingress"
  port_range_min    = 22
  port_range_max    = 22
  ethertype         = "IPv4"
  protocol          = "tcp"
  security_group_id = var.iec_security_group_id
  remote_ip_prefix  = "0.0.0.0/0"
}
```

## Argument Reference

The following arguments are supported:

* `direction` - (Required, String, ForceNew) Specifies the direction of the rule, valid values are __ingress__ or
  __egress__. Changing this parameter creates a new security group rule resource.

* `ethertype` - (Required, String, ForceNew) Specifies the layer 3 protocol type, valid values are __IPv4__(IPv4 is
  default) or __IPv6__. Changing this parameter creates a new security group rule resource.

* `protocol` - (Required, String, ForceNew) Specifies the layer 4 protocol type, valid values are following. The valid
  values are: __tcp__, __udp__, __icmp__ and __gre__. Changing this parameter creates a new security group rule
  resource.

* `security_group_id` - (Required, String, ForceNew) Specifies the security group id the rule should belong to. Changing
  this parameter creates a new security group rule resource.

* `remote_ip_prefix` - (Optional, String, ForceNew) Specifies the remote CIDR, the value to be a valid CIDR (i.e.
  192.168.0.0/16). This parameter and remote_group_id are alternative. Changing this parameter creates a new security
  group rule resource.

* `remote_group_id` - (Optional, String, ForceNew) Specifies the remote group id, the value needs to be an ID of a
  security group. This parameter and remote_ip_prefix are alternative. Changing this parameter creates a new security
  group rule resource.

* `description` - (Optional, String, ForceNew) Specifies a description of the security group rule. Changing this
  parameter creates a new security group rule resource.

* `port_range_min` - (Optional, Int, ForceNew) Specifies the lower part of the allowed port range, valid integer value
  needs to be between 1 and 65535. Changing this parameter creates a new security group rule resource.

* `port_range_max` - (Optional, Int, ForceNew) Specifies the higher part of the allowed port range, valid integer value
  needs to be between 1 and 65535. Changing this parameter creates a new security group rule resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - A resource ID in UUID format.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 10 minute.
