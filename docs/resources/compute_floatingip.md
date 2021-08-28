---
subcategory: "Deprecated"
---

# huaweicloud_compute_floatingip_v2

!> **WARNING:** It has been deprecated, use `huaweicloud_vpc_eip` instead.

Manages a V2 floating IP resource within HuaweiCloud Nova (compute)
that can be used for compute instances.

## Example Usage

```hcl
resource "huaweicloud_compute_floatingip_v2" "floatip_1" {
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the V2 Compute client. A Compute client is
  needed to create a floating IP that can be used with a compute instance. If omitted, the `region` argument of the
  provider is used. Changing this creates a new floating IP (which may or may not have a different address).

* `pool` - (Optional, String, ForceNew) The name of the pool from which to create the floating IP. Only
  admin_external_net is valid. Changing this creates a new floating IP.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.
* `address` - The actual floating IP address itself.
* `fixed_ip` - The fixed IP address corresponding to the floating IP.
* `instance_id` - UUID of the compute instance associated with the floating IP.

## Import

Floating IPs can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_compute_floatingip_v2.floatip_1 89c60255-9bd6-460c-822a-e2b959ede9d2
```
