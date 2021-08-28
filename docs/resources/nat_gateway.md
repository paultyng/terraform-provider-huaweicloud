---
subcategory: "NAT Gateway (NAT)"
---

# huaweicloud_nat_gateway

Manages a Nat gateway resource within HuaweiCloud Nat This is an alternative to `huaweicloud_nat_gateway_v2`

## Example Usage

```hcl
resource "huaweicloud_nat_gateway" "nat_1" {
  name        = "test"
  description = "test for terraform"
  spec        = "3"
  vpc_id      = "2c1fe4bd-ebad-44ca-ae9d-e94e63847b75"
  subnet_id   = "dc8632e2-d9ff-41b1-aa0c-d455557314a0"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the Nat gateway resource. If omitted,
  the provider-level region will be used. Changing this creates a new nat gateway.

* `name` - (Required, String) Specifies the nat gateway name. The name can contain only digits, letters, underscores (_)
  , and hyphens(-).

* `spec` - (Required, String) Specifies the nat gateway type. The value can be:
  + `1`: small type, which supports up to 10,000 SNAT connections.
  + `2`: medium type, which supports up to 50,000 SNAT connections.
  + `3`: large type, which supports up to 200,000 SNAT connections.
  + `4`: extra-large type, which supports up to 1,000,000 SNAT connections.

* `vpc_id` - (Required, String, ForceNew) Specifies the ID of the VPC this nat gateway belongs to. Changing this creates
  a new nat gateway.

* `subnet_id` - (Required, String, ForceNew) Specifies the subnet ID of the downstream interface (the next hop of the
  DVR) of the NAT gateway. Changing this creates a new nat gateway.

* `description` - (Optional, String) Specifies the description of the nat gateway. The value contains 0 to 255
  characters, and angle brackets (<)
  and (>) are not allowed.

* `enterprise_project_id` - (Optional, String, ForceNew) Specifies the enterprise project id of the nat gateway. The
  value can contains maximum of 36 characters which it is string "0" or in UUID format with hyphens (-). Changing this
  creates a new nat gateway.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID in UUID format.

* `status` - The status of the nat gateway.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

Nat gateway can be imported using the following format:

```
$ terraform import huaweicloud_nat_gateway.nat_1 d126fb87-43ce-4867-a2ff-cf34af3765d9
```
