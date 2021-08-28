---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_network_acl_rule

Manages a network ACL rule resource within HuaweiCloud IEC.

## Example Usage

```hcl
resource "huaweicloud_iec_network_acl" "acl_test" {
  name = "iec-network-acl-demo"
}

resource "huaweicloud_iec_network_acl_rule" "rule_test" {
  network_acl_id         = huaweicloud_iec_network_acl.acl_test.id
  direction              = "ingress"
  protocol               = "tcp"
  action                 = "allow"
  source_ip_address      = "152.16.30.0/24"
  destination_ip_address = "192.168.128.0/18"
  destination_port       = "445"
  enabled                = true
}
```

## Argument Reference

The following arguments are supported:

* `network_acl_id` - (Required, String) Specifies a unique id for the iec network ACL.

* `direction` - (Required, String, ForceNew) Specifies the direction of the rule, valid values are *ingress* or *egress*
  . Changing this parameter creates a new iec network ACL rule resource.

* `description` - (Optional, String) Specifies the description for the iec network ACL rule.

* `protocol` - (Optional, String) Specifies the protocol supported by the iec network ACL rule. Valid values are: *tcp*
  , *udp*, *icmp* and *any*.

* `action` - (Optional, String) Specifies the action in the iec network ACL rule. Currently, the value can be *allow*
  or *deny*.

* `source_ip_address` - (Optional, String) Specifies the source IP address that the traffic is allowed from. The default
  value is *0.0.0.0/0*. For example:
  xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).

* `destination_ip_address` - (Optional, String) Specifies the destination IP address to which the traffic is allowed.
  The default value is *0.0.0.0/0*. For example: xxx.xxx.xxx.xxx (IP address), xxx.xxx.xxx.0/24 (CIDR block).

* `source_port` - (Optional, String) Specifies the source port number or port number range. The value ranges from 1 to
  65535. For a port number range, enter two port numbers connected by a hyphen (-). For example, 1-100.

* `destination_port` - (Optional, String) Specifies the destination port number or port number range. The value ranges
  from 1 to 65535. For a port number range, enter two port numbers connected by a hyphen (-). For example, 1-100.

* `enabled` - (Optional, Bool) Specifies the Enabled status for the iec network ACL rule. The default value is true.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID in UUID format.

* `ip_version` - The version of elastic IP address. IEC services only support IPv4(4) now.

* `poilicy_id` - The ID of the firewall policy for the iec network ACL.
