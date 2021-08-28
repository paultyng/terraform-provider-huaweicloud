---
subcategory: "Identity and Access Management (IAM)"
---

# huaweicloud_identity_user

Manages a User resource within HuaweiCloud IAM service. This is an alternative to `huaweicloud_identity_user_v3`

Note: You _must_ have admin privileges in your HuaweiCloud cloud to use this resource.

## Example Usage

```hcl
resource "huaweicloud_identity_user" "user_1" {
  name        = "user_1"
  description = "A user"
  password    = "password123!"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Specifies the name of the user. The user name consists of 5 to 32 characters. It can
  contain only uppercase letters, lowercase letters, digits, spaces, and special characters (-_) and cannot start with a
  digit.

* `description` - (Optional, String) Specifies the description of the user.

* `email` - (Optional, String) Specifies the email address with a maximum of 255 characters.

* `phone` - (Optional, String) Specifies the mobile number with a maximum of 32 digits. This parameter must be used
  together with `country_code`.

* `country_code` - (Optional, String) Specifies the country code. The country code of the Chinese mainland is 0086. This
  parameter must be used together with `phone`.

* `password` - (Optional, String) Specifies the password for the user with 6 to 32 characters. It must contain at least
  two of the following character types: uppercase letters, lowercase letters, digits, and special characters.

* `enabled` - (Optional, Bool) Specifies whether the user is enabled or disabled. Valid values are `true` and `false`.

* `access_type` - (Optional, String) Specifies the access type of the user. Available values are:
  + default: support both programmatic and management console access.
  + programmatic: only support programmatic access.
  + console: only support management console access.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID in UUID format.

* `password_strength` - Indicates the password strength.

* `create_time` - The time when the IAM user was created.

* `last_login` - The tiem when the IAM user last login.

## Import

Users can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_identity_user.user_1 89c60255-9bd6-460c-822a-e2b959ede9d2
```

But due to the security reason, `password` can not be imported, you can ignore it as below.

```
resource "huaweicloud_identity_user" "user_1" {
  ...

  lifecycle {
    ignore_changes = [
      "password",
    ]
  }
}
```
