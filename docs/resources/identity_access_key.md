---
subcategory: "Identity and Access Management (IAM)"
---

# huaweicloud_identity_access_key

Manages a permanent Access Key resource within HuaweiCloud IAM service.

-> **NOTE:** You _must_ have admin privileges in your HuaweiCloud cloud to use this resource.

## Example Usage

```hcl
resource "huaweicloud_identity_user" "user_1" {
  name        = "user_1"
  description = "A user"
  password    = "password123!"
}

resource "huaweicloud_identity_access_key" "key_1" {
  user_id = huaweicloud_identity_user.user_1.id
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required, String, ForceNew) Specifies the ID of the user who is requesting to create an access key.
  Changing this creates a new resource.

* `description` - (Optional, String) Specifies the description of the access key.

* `status` - (Optional, String) Specifies the status of the access key. It must be *active* or *inactive*. Default value
  is *active*.

* `secret_file` - (Optional, String, ForceNew) Specifies the file name that can save access key and access secret key.
  Defaults to *./credentials-{{user name}}.csv*. Changing this creates a new resource.

* `pgp_key` - (Optional, String, ForceNew) Either a base-64 encoded PGP public key, or a keybase username in the form
  `keybase:some_person_that_exists`. Changing this creates a new resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The access key ID.
* `secret` - The access secret key. Setting the value only when writing to `secret_file` failed.
* `key_fingerprint` - The fingerprint of the PGP key used to encrypt the secret
* `encrypted_secret` - The encrypted secret, base64 encoded. The encrypted secret may be decrypted using the command
  line, for example: `terraform output encrypted_secret | base64 --decode | keybase pgp decrypt`.
* `user_name` - The name of IAM user.
* `create_time` - The time when the access key was created.
