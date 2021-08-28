---
subcategory: "Web Application Firewall (WAF)"
---

# huaweicloud_waf_rule_web_tamper_protection

Manages a WAF web tamper protection rule resource within HuaweiCloud.

## Example Usage

```hcl
resource "huaweicloud_waf_policy" "policy_1" {
  name = "policy_1"
}

resource "huaweicloud_waf_rule_web_tamper_protection" "rule_1" {
  policy_id = huaweicloud_waf_policy.policy_1.id
  domain    = "www.your-domain.com"
  path      = "/payment"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the WAF web tamper protection rules resource. If
  omitted, the provider-level region will be used. Changing this creates a new rule.

* `policy_id` - (Required, String, ForceNew) Specifies the WAF policy ID. Changing this creates a new rule.

* `domain` - (Required, String, ForceNew) Specifies the domain name. Changing this creates a new rule.

* `path` - (Required, String, ForceNew) Specifies the URL protected by the web tamper protection rule, excluding a
  domain name. Changing this creates a new rule.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The rule ID in UUID format.

## Import

Web Tamper Protection Rules can be imported using the policy ID and rule ID separated by a slash, e.g.:

```sh
terraform import huaweicloud_waf_rule_web_tamper_protection.rule_1 840c6dfdd5604c1781eea033eae2004f/c6dbc13bb7e74788ae53ecc9254b3ea8
```
