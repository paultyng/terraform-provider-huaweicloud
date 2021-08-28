---
subcategory: "Deprecated"
---

# huaweicloud\_networking\_floatingip\_v2

!> **WARNING:** It has been deprecated, use `huaweicloud_vpc_eip` instead.

Manages a V2 floating IP resource within HuaweiCloud Neutron (networking)

## Example Usage

```hcl
resource "huaweicloud_networking_floatingip_v2" "floatip_1" {
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the V2 Networking client. A Networking client is
  needed to create a floating IP that can be used with another networking resource, such as a load balancer. If omitted,
  the
  `region` argument of the provider is used. Changing this creates a new floating IP (which may or may not have a
  different address).

* `pool` - (Optional, String, ForceNew) The name of the pool from which to create the floating IP. Only
  admin_external_net is valid. Changing this creates a new floating IP.

* `port_id` - (Optional, String) ID of an existing port with at least one IP address to associate with this floating IP.

* `fixed_ip` - (Optional, String) Fixed IP of the port to associate with this floating IP. Required if the port has
  multiple fixed IPs.

* `value_specs` - (Optional, Map, ForceNew) Map of additional options.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.
* `address` - The actual floating IP address itself.
* `port_id` - ID of associated port.
* `tenant_id` - the ID of the tenant in which to create the floating IP.
* `fixed_ip` - The fixed IP which the floating IP maps to.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

Floating IPs can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_networking_floatingip_v2.floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1
```
