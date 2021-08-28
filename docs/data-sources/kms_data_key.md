---
subcategory: "Key Management Service (KMS)"
---

# huaweicloud_kms_data_key

Use this data source to get the plaintext and the ciphertext of an available HuaweiCloud KMS DEK (data encryption key).
This is an alternative to `huaweicloud_kms_key_v1`

## Example Usage

```hcl
resource "huaweicloud_kms_key" "key1" {
  key_alias       = "key_1"
  pending_days    = "7"
  key_description = "first test key"
}

data "huaweicloud_kms_data_key" "kms_datakey1" {
  key_id         = huaweicloud_kms_key.key1.id
  datakey_length = "512"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the keys. If omitted, the provider-level region will be
  used.

* `key_id` - (Required, String) The globally unique identifier for the key. Changing this gets the new data encryption
  key.

* `encryption_context` - (Optional, String) The value of this parameter must be a series of
  "key:value" pairs used to record resource context information. The value of this parameter must not contain sensitive
  information and must be within 8192 characters in length. Example: {"Key1":"Value1","Key2":"Value2"}

* `datakey_length` - (Required, String) Number of bits in the length of a DEK (data encryption keys). The maximum number
  is 512. Changing this gets the new data encryption key.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.
* `plain_text` - The plaintext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.
* `cipher_text` - The ciphertext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.
