---
subcategory: "Deprecated"
---

# huaweicloud\_fw\_policy_v2

Manages a v2 firewall policy resource within HuaweiCloud.

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
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the v2 networking client. A networking client is
  needed to create a firewall policy. If omitted, the
  `region` argument of the provider is used. Changing this creates a new firewall policy.

* `name` - (Optional, String) A name for the firewall policy. Changing this updates the `name` of an existing firewall
  policy.

* `description` - (Optional, String) A description for the firewall policy. Changing this updates the `description` of
  an existing firewall policy.

* `rules` - (Optional, List, String) An array of one or more firewall rules that comprise the policy. Changing this
  results in adding/removing rules from the existing firewall policy.

* `audited` - (Optional, Bool) Audit status of the firewall policy
  (must be "true" or "false" if provided - defaults to "false"). This status is set to "false" whenever the firewall
  policy or any of its rules are changed. Changing this updates the `audited` status of an existing firewall policy.

* `shared` - (Optional, Bool) Sharing status of the firewall policy (must be "true"
  or "false" if provided). If this is "true" the policy is visible to, and can be used in, firewalls in other tenants.
  Changing this updates the
  `shared` status of an existing firewall policy. Only administrative users can specify if the policy should be shared.

* `value_specs` - (Optional, Map, ForceNew) Map of additional options.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.

## Import

Firewall Policies can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_fw_policy_v2.policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0
```
