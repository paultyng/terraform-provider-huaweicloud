---
subcategory: "Key Management Service (KMS)"
---

# huaweicloud_kms_key

Use this data source to get the ID of an available HuaweiCloud KMS key. This is an alternative
to `huaweicloud_kms_key_v1`

## Example Usage

```hcl
data "huaweicloud_kms_key" "key_1" {
  key_alias        = "test_key"
  key_description  = "test key description"
  key_state        = "2"
  key_id           = "af650527-a0ff-4527-aef3-c493df1f3012"
  default_key_flag = "0"
  domain_id        = "b168fe00ff56492495a7d22974df2d0b"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the keys. If omitted, the provider-level region will be
  used.

* `key_alias` - (Optional, String) The alias in which to create the key. It is required when we create a new key.
  Changing this gets the new key.

* `key_description` - (Optional, String) The description of the key as viewed in Huawei console. Changing this gets a
  new key.

* `key_id` - (Optional, String) The globally unique identifier for the key. Changing this gets the new key.

* `default_key_flag` - (Optional, String) Identification of a Master Key. The value "1" indicates a Default Master Key,
  and the value "0" indicates a key. Changing this gets a new key.

* `key_state` - (Optional, String) The state of a key. "1" indicates that the key is waiting to be activated.
  "2" indicates that the key is enabled. "3" indicates that the key is disabled. "4" indicates that the key is scheduled
  for deletion. Changing this gets a new key.

* `domain_id` - (Optional, String)  - ID of a user domain for the key. Changing this gets a new key.

* `enterprise_project_id` - (Optional, String) The enterprise project id of the kms key.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.
* `scheduled_deletion_date` - Scheduled deletion time (time stamp) of a key.
* `expiration_time` - Expiration time.
* `creation_date` - Creation time (time stamp) of a key.
* `tags` - The key/value pairs to associate with the kms key.
