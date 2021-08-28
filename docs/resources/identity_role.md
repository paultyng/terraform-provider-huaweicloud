---
subcategory: "Identity and Access Management (IAM)"
---

# huaweicloud_identity_role

Manages a **Custom Policy** resource within HuaweiCloud IAM service.

->**Note** You _must_ have admin privileges in your HuaweiCloud cloud to use this resource.

## Example Usage

```hcl
resource "huaweicloud_identity_role" role1 {
  name        = "test"
  description = "created by terraform"
  type        = "AX"
  policy      = <<EOF
{
  "Version": "1.1",
  "Statement": [
    {
      "Action": [
        "obs:bucket:GetBucketAcl"
      ],
      "Effect": "Allow",
      "Resource": [
        "obs:*:*:bucket:*"
      ],
      "Condition": {
        "StringStartWith": {
          "g:ProjectName": [
            "cn-north-4"
          ]
        }
      }
    }
  ]
}
EOF
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Name of the custom policy.

* `description` - (Required, String) Description of the custom policy.

* `type` - (Required, String) Display mode. Valid options are _AX_: Account level and _XA_: Project level.

* `policy` - (Required, String) Document of the custom policy in JSON format. For more details, please refer to the
  [offical document](https://support.huaweicloud.com/intl/en-us/usermanual-iam/iam_01_0017.html).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The role id.

* `references` - The number of references.

## Import

Roles can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_identity_role.role_1 89c60255-9bd6-460c-822a-e2b959ede9d2
```
