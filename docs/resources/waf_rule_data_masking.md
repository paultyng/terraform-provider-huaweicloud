---
subcategory: "Web Application Firewall (WAF)"
---

# huaweicloud_waf_rule_data_masking

Manages a WAF Data Masking Rule resource within HuaweiCloud.

## Example Usage

```hcl
resource "huaweicloud_waf_policy" "policy_1" {
  name = "policy_1"
}

resource "huaweicloud_waf_rule_data_masking" "rule_1" {
  policy_id = huaweicloud_waf_policy.policy_1.id
  path      = "/login"
  field     = "params"
  subfield  = "password"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the WAF Data Masking rule resource. If omitted,
  the provider-level region will be used. Changing this setting will create a new rule.

* `policy_id` - (Required, String, ForceNew) Specifies the WAF policy ID. Changing this creates a new rule.

* `path` - (Required, String) Specifies the URL to which the data masking rule applies (exact match by default).

* `field` - (Required, String) The position where the masked field stored. Valid values are:
  + `params`: The field in the parameter.
  + `header`: The field in the header.
  + `form`: The field in the form.
  + `cookie`: The field in the cookie.

* `subfield` - (Required, String) Specifies the name of the masked field, e.g.: password.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The rule ID in UUID format.

## Import

Data Masking Rules can be imported using the policy ID and rule ID separated by a slash, e.g.:

```sh
terraform import huaweicloud_waf_rule_data_masking.rule_1 d78b439fd5e54ea08886e5f63ee7b3f5/ac01a092d50e4e6ba3cd622c1128ba2c
```
