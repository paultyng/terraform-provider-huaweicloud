---
subcategory: "Key Management Service (KMS)"
---

# huaweicloud_kms_key

Manages a KMS key resource. This is an alernative to `huaweicloud_kms_key_v1`

## Example Usage

```hcl
resource "huaweicloud_kms_key" "key_1" {
  key_alias       = "key_1"
  pending_days    = "7"
  key_description = "first test key"
  is_enabled      = true
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the KMS key resource. If omitted, the
  provider-level region will be used. Changing this creates a new KMS key resource.

* `key_alias` - (Required, String) The alias in which to create the key. It is required when we create a new key.
  Changing this updates the alias of key.

* `key_description` - (Optional, String) The description of the key as viewed in Huawei console. Changing this updates
  the description of key.

* `pending_days` - (Optional, String) Duration in days after which the key is deleted after destruction of the resource,
  must be between 7 and 1096 days. It doesn't have default value. It only be used when delete a key.

* `is_enabled` - (Optional, Bool) Specifies whether the key is enabled. Defaults to true. Changing this updates the
  state of existing key.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id of the kms key. Changing this creates
  a new key.

* `tags` - (Optional, Map) Specifies the key/value pairs to associate with the kms key.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.
* `key_id` - The globally unique identifier for the key.
* `default_key_flag` - Identification of a Master Key. The value 1 indicates a Default Master Key, and the value 0
  indicates a key.
* `scheduled_deletion_date` - Scheduled deletion time (time stamp) of a key.
* `domain_id` - ID of a user domain for the key.
* `expiration_time` - Expiration time.
* `creation_date` - Creation time (time stamp) of a key.

## Import

KMS Keys can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_kms_key.key_1 7056d636-ac60-4663-8a6c-82d3c32c1c64
```
