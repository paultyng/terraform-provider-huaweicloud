---
subcategory: "Deprecated"
---

# huaweicloud\_networking\_subnet\_v2

Manages a V2 Neutron subnet resource within HuaweiCloud.

!> **WARNING:** It has been deprecated, use `huaweicloud_vpc_subnet` instead.

## Example Usage

```hcl
resource "huaweicloud_networking_network_v2" "network_1" {
  name           = "tf_test_network"
  admin_state_up = "true"
}

resource "huaweicloud_networking_subnet_v2" "subnet_1" {
  network_id = huaweicloud_networking_network_v2.network_1.id
  cidr       = "192.168.199.0/24"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the V2 Networking client. A Networking client is
  needed to create a Neutron subnet. If omitted, the
  `region` argument of the provider is used. Changing this creates a new subnet.

* `network_id` - (Required, String, ForceNew) The UUID of the parent network. Changing this creates a new subnet.

* `cidr` - (Required, String, ForceNew) CIDR representing IP range for this subnet, based on IP version. Changing this
  creates a new subnet.

* `ip_version` - (Optional, Int, ForceNew) IP version, either 4 (default) or 6. Changing this creates a new subnet.

* `name` - (Optional, String) The name of the subnet. Changing this updates the name of the existing subnet.

* `allocation_pools` - (Optional, List) An array of sub-ranges of CIDR available for dynamic allocation to ports. The
  allocation_pool object structure is documented below. Changing this creates a new subnet.

* `gateway_ip` - (Optional, String)  Default gateway used by devices in this subnet. Leaving this blank and not
  setting `no_gateway` will cause a default gateway of `.1` to be used. Changing this updates the gateway IP of the
  existing subnet.

* `no_gateway` - (Optional, Bool) Do not set a gateway IP on this subnet. Changing this removes or adds a default
  gateway IP of the existing subnet.

* `enable_dhcp` - (Optional, Bool) The administrative state of the network. The value must be "true".

* `dns_nameservers` - (Optional, String) An array of DNS name server names used by hosts in this subnet. Changing this
  updates the DNS name servers for the existing subnet.

* `host_routes` - (Optional, List) An array of routes that should be used by devices with IPs from this subnet (not
  including local subnet route). The host_route object structure is documented below. Changing this updates the host
  routes for the existing subnet.

* `value_specs` - (Optional, Map, ForceNew) Map of additional options.

The `allocation_pools` block supports:

* `start` - (Required, String) The starting address.

* `end` - (Required, String) The ending address.

The `host_routes` block supports:

* `destination_cidr` - (Required, String) The destination CIDR.

* `next_hop` - (Required, String) The next hop in the route.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

Subnets can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_networking_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048
```
