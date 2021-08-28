---
subcategory: "Deprecated"
---

# huaweicloud_vpnaas_ike_policy_v2

Manages a V2 IKE policy resource within HuaweiCloud.

## Example Usage

```hcl
resource "huaweicloud_vpnaas_ike_policy_v2" "policy_1" {
  name = "my_policy"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional) The region in which to create the V2 Networking client. A Networking client is needed to create
  a VPN service. If omitted, the
  `region` argument of the provider is used. Changing this creates a new service.

* `name` - (Optional) The name of the policy. Changing this updates the name of the existing policy.

* `description` - (Optional) The human-readable description for the policy. Changing this updates the description of the
  existing policy.

* `auth_algorithm` - (Optional) The authentication hash algorithm. Valid values are md5, sha1, sha2-256, sha2-384,
  sha2-512. Default is sha1. Changing this updates the algorithm of the existing policy.

* `encryption_algorithm` - (Optional) The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on. The
  default value is aes-128. Changing this updates the existing policy.

* `pfs` - (Optional) The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
  Changing this updates the existing policy.

* `phase1_negotiation_mode` - (Optional) The IKE mode. A valid value is main, which is the default. Changing this
  updates the existing policy.

* `ike_version` - (Optional) The IKE mode. A valid value is v1 or v2. Default is v1. Changing this updates the existing
  policy.

* `lifetime` - (Optional) The lifetime of the security association. Consists of Unit and Value.
  + `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
      Default is seconds.
  + `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer. Default
      is 3600.

* `value_specs` - (Optional) Map of additional options.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.

## Import

Services can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_vpnaas_ike_policy_v2.policy_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
```
