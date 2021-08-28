---
subcategory: "Identity and Access Management (IAM)"
---

# huaweicloud_identity_acl

Manages a ACL resource within HuaweiCloud IAM service. The ACL allowing user access only from specified IP address
ranges and IPv4 CIDR blocks. The ACL take effect for IAM users under the Domain account rather than the account itself.

Note: You _must_ have admin privileges in your HuaweiCloud cloud to use this resource.

## Example Usage

```hcl
resource "huaweicloud_identity_acl" "acl" {
  type = "console"

  ip_cidrs {
    cidr        = "159.138.39.192/32"
    description = "This is a test ip address"
  }
  ip_ranges {
    range       = "0.0.0.0-255.255.255.0"
    description = "This is a test ip range"
  }
}
```

## Argument Reference

The following arguments are supported:

* `type` - (Required, String, ForceNew) Specifies the ACL is created through the Console or API. valid value are '
  console' and 'api'. Changing this parameter will create a new ACL.

* `ip_cidrs` - (Optional, List) Specifies the IPv4 CIDR blocks from which console access or api access is allowed.
  The `ip_cidrs` cannot repeat. The structure is documented below.

* `ip_ranges` - (Optional, List) Specifies the IP address ranges from which console access or api access is allowed.
  The `ip_ranges` cannot repeat. The structure is documented below.

The `ip_cidrs` block supports:

* `cidr` - (Required, String) Specifies the IPv4 CIDR block, for example, __192.168.0.0/24__.

* `description` - (Optional, String) Specifies a description about an IPv4 CIDR block. This parameter can contain a
  maximum of 255 characters and the following charactors are not allowed:__@#%^&*<>\\__.

The `ip_ranges` block supports:

* `range` - (Required, String) Specifies the Ip address range, for example, __0.0.0.0-255.255.255.0__.

* `description` - (Optional, String) Specifies a description about an IP address range. This parameter can contain a
  maximum of 255 characters and the following charactors are not allowed:__@#%^&*<>\\__.

**note**: Up to 200 `ip_cidrs` and `ip_ranges` can be created in total for each access method.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of identity acl.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 5 minute.
* `update` - Default is 5 minute.
* `delete` - Default is 3 minute.
