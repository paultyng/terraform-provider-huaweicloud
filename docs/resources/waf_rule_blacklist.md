---
subcategory: "Web Application Firewall (WAF)"
---

# huaweicloud_waf_rule_blacklist

Manages a WAF blacklist and whitelist rule resource within HuaweiCloud.

## Example Usage

```hcl
resource "huaweicloud_waf_policy" "policy_1" {
  name = "policy_1"
}

resource "huaweicloud_waf_rule_blacklist" "rule_1" {
  policy_id  = huaweicloud_waf_policy.policy_1.id
  ip_address = "192.168.0.0/24"
  action     = 0
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the WAF blacklist and whitelist rule resource.
  If omitted, the provider-level region will be used. Changing this setting will push a new certificate.

* `policy_id` - (Required, String, ForceNew) Specifies the WAF policy ID. Changing this creates a new rule. Please make
  sure that the region which the policy belongs to be consistent with the `region`.

* `ip_address` - (Required, String) Specifies the IP address or range. For example, 192.168.0.125 or 192.168.0.0/24.

* `action` - (Optional, Int) Specifies the protective action. Defaults is `0`. The value can be:
  + `0`: block the request.
  + `1`: allow the request.
  + `2`: log the request only.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The rule ID in UUID format.

## Import

Blacklist and Whiltelist Rules can be imported using the policy ID and rule ID separated by a slash, e.g.:

```sh
terraform import huaweicloud_waf_rule_blacklist.rule_1 d78b439fd5e54ea08886e5f63ee7b3f5/ac01a092d50e4e6ba3cd622c1128ba2c
```
