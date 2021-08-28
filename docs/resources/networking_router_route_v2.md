---
subcategory: "Deprecated"
---

# huaweicloud_networking_router_route_v2

Creates a routing entry on a HuaweiCloud V2 router.

!> **WARNING:** It has been deprecated, use `huaweicloud_vpc_route` instead.

## Example Usage

```hcl
resource "huaweicloud_networking_router_v2" "router_1" {
  name           = "router_1"
  admin_state_up = "true"
}

resource "huaweicloud_networking_network_v2" "network_1" {
  name           = "network_1"
  admin_state_up = "true"
}

resource "huaweicloud_networking_subnet_v2" "subnet_1" {
  network_id = huaweicloud_networking_network_v2.network_1.id
  cidr       = "192.168.199.0/24"
  ip_version = 4
}

resource "huaweicloud_networking_router_interface_v2" "int_1" {
  router_id = huaweicloud_networking_router_v2.router_1.id
  subnet_id = huaweicloud_networking_subnet_v2.subnet_1.id
}

resource "huaweicloud_networking_router_route_v2" "router_route_1" {
  depends_on       = [
    "huaweicloud_networking_router_interface_v2.int_1"]
  router_id        = huaweicloud_networking_router_v2.router_1.id
  destination_cidr = "10.0.1.0/24"
  next_hop         = "192.168.199.254"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the V2 networking client. A networking client is
  needed to configure a routing entry on a router. If omitted, the
  `region` argument of the provider is used. Changing this creates a new routing entry.

* `router_id` - (Required, String, ForceNew) ID of the router this routing entry belongs to. Changing this creates a new
  routing entry.

* `destination_cidr` - (Required, String, ForceNew) CIDR block to match on the packet’s destination IP. Changing this
  creates a new routing entry.

* `next_hop` - (Required, String, ForceNew) IP address of the next hop gateway. Changing this creates a new routing
  entry.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

## Notes

The `next_hop` IP address must be directly reachable from the router at the ``huaweicloud_networking_router_route_v2``
resource creation time. You can ensure that by explicitly specifying a dependency on
the ``huaweicloud_networking_router_interface_v2``
resource that connects the next hop to the router, as in the example above.

## Import

Routing entries can be imported using a combined ID using the following
format: ``<router_id>-route-<destination_cidr>-<next_hop>``

```
$ terraform import huaweicloud_networking_router_route_v2.router_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25
```
