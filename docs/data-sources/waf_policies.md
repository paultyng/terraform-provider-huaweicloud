---
subcategory: "Web Application Firewall (WAF)"
---

# huaweicloud_waf_policies

Use this data source to get a list of WAF policies.

## Example Usage

```hcl
variable "policy_name" {}
data "huaweicloud_waf_policies" "policies" {
  name = var.policy_name
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) The region in which to obtain the WAF policies. If omitted, the provider-level region
  will be used.

* `name` - (Optional, String) Policy name used for matching. The value is case sensitive and supports fuzzy matching.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `policies` - A list of WAF policies.

The `policies` block supports:

* `id` - The WAF Policy ID.

* `name` - The WAF policy name.

* `protection_mode` - Specifies the protective action after a rule is matched. Valid values are:
  + `block`: WAF blocks and logs detected attacks.
  + `log`: WAF logs detected attacks only.

* `level` - Specifies the protection level. Valid values are:
  + `1`: low
  + `2`: medium
  + `3`: high

* `full_detection` - The detection mode in Precise Protection.
  + `true`: full detection. Full detection finishes all threat detections before blocking requests that meet Precise
      Protection specified conditions.
  + `false`: instant detection. Instant detection immediately ends threat detection after blocking a request that
      meets Precise Protection specified conditions.

* `options` - The protection switches. The options object structure is documented below.

The `options` block supports:

* `basic_web_protection` - Indicates whether Basic Web Protection is enabled.

* `general_check` - Indicates whether General Check in Basic Web Protection is enabled.

* `crawler` - Indicates whether the master crawler detection switch in Basic Web Protection is enabled.

* `crawler_engine` - Indicates whether the Search Engine switch in Basic Web Protection is enabled.

* `crawler_scanner` - Indicates whether the Scanner switch in Basic Web Protection is enabled.

* `crawler_script` - Indicates whether the Script Tool switch in Basic Web Protection is enabled.

* `crawler_other` - Indicates whether detection of other crawlers in Basic Web Protection is enabled.

* `webshell` - Indicates whether webshell detection in Basic Web Protection is enabled.

* `cc_attack_protection` - Indicates whether CC Attack Protection is enabled.

* `precise_protection` - Indicates whether Precise Protection is enabled.

* `blacklist` - Indicates whether Blacklist and Whitelist is enabled.

* `data_masking` - Indicates whether Data Masking is enabled.

* `false_alarm_masking` - Indicates whether False Alarm Masking is enabled.

* `web_tamper_protection` - Indicates whether Web Tamper Protection is enabled.
