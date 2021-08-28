---
subcategory: "Deprecated"
---

# huaweicloud\_fw\_firewall\_group\_v2

Manages a v2 firewall group resource within HuaweiCloud.

!> **WARNING:** It has been deprecated, use `huaweicloud_network_acl` instead.

## Example Usage

```hcl
resource "huaweicloud_fw_rule_v2" "rule_1" {
  name             = "my-rule-1"
  description      = "drop TELNET traffic"
  action           = "deny"
  protocol         = "tcp"
  destination_port = "23"
  enabled          = "true"
}

resource "huaweicloud_fw_rule_v2" "rule_2" {
  name             = "my-rule-2"
  description      = "drop NTP traffic"
  action           = "deny"
  protocol         = "udp"
  destination_port = "123"
  enabled          = "false"
}

resource "huaweicloud_fw_policy_v2" "policy_1" {
  name = "my-policy"

  rules = [
    huaweicloud_fw_rule_v2.rule_1.id,
    huaweicloud_fw_rule_v2.rule_2.id,
  ]
}

resource "huaweicloud_fw_firewall_group_v2" "firewall_group_1" {
  name              = "my-firewall-group"
  ingress_policy_id = huaweicloud_fw_policy_v2.policy_1.id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the v2 networking client. A networking client is
  needed to create a firewall group. If omitted, the
  `region` argument of the provider is used. Changing this creates a new firewall group.

* `ingress_policy_id` - (Optional, String) The ingress policy resource id for the firewall group. Changing this updates
  the `ingress_policy_id` of an existing firewall group.

* `egress_policy_id` - (Optional, String) The egress policy resource id for the firewall group. Changing this updates
  the `egress_policy_id` of an existing firewall group.

* `name` - (Optional, String) A name for the firewall group. Changing this updates the `name` of an existing firewall
  group.

* `description` - (Required, String) A description for the firewall group. Changing this updates the `description` of an
  existing firewall group.

* `admin_state_up` - (Optional, Bool) Administrative up/down status for the firewall group
  (must be "true" or "false" if provided - defaults to "true"). Changing this updates the `admin_state_up` of an
  existing firewall group.

* `ports` - (Optional, String) Port(s) to associate this firewall group instance with. Must be a list of strings.
  Changing this updates the associated routers of an existing firewall group.

* `value_specs` - (Optional, Map, ForceNew) Map of additional options.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `update` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

Firewall Groups can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_fw_firewall_group_v2.firewall_group_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97
```
