---
subcategory: "Deprecated"
---

# huaweicloud_networking_router_v2

Manages a V2 router resource within HuaweiCloud.

!> **WARNING:** It has been deprecated, use `huaweicloud_vpc` instead.

## Example Usage

```hcl
resource "huaweicloud_networking_router_v2" "router_1" {
  name                = "my_router"
  admin_state_up      = true
  external_network_id = "f67f0d72-0ddf-11e4-9d95-e1f29f417e2f"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the V2 networking client. A networking client is
  needed to create a router. If omitted, the
  `region` argument of the provider is used. Changing this creates a new router.

* `name` - (Optional, String) A unique name for the router. Changing this updates the `name` of an existing router.

* `admin_state_up` - (Optional, Bool) Administrative up/down status for the router
  (must be "true" or "false" if provided). Changing this updates the
  `admin_state_up` of an existing router.

* `distributed` - (Optional, Bool, ForceNew) Indicates whether or not to create a distributed router. The default policy
  setting in Neutron restricts usage of this property to administrative users only.

* `external_network_id` - (Optional, String) The network UUID of an external gateway for the router. A router with an
  external gateway is required if any compute instances or load balancers will be using floating IPs. Changing this
  updates the external gateway of the router.

* `enable_snat` - (Optional, Bool) Enable Source NAT for the router. Valid values are
  "true" or "false". An `external_network_id` has to be set in order to set this property. Changing this updates
  the `enable_snat` of the router.

* `external_fixed_ip` - (Optional, List) An external fixed IP for the router. This can be repeated. The structure is
  described below. An `external_network_id`
  has to be set in order to set this property. Changing this updates the external fixed IPs of the router.

* `value_specs` - (Optional, Map, ForceNew) Map of additional driver-specific options.

The `external_fixed_ip` block supports:

* `subnet_id` - (Optional, String) Subnet in which the fixed IP belongs to.

* `ip_address` - (Optional, String) The IP address to set on the router.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the router.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

Routers can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_networking_router_v2.router_1 014395cd-89fc-4c9b-96b7-13d1ee79dad2
```
