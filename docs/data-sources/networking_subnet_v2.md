---
subcategory: "Deprecated"
---

# huaweicloud\_networking\_subnet\_v2

Use this data source to get the ID of an available HuaweiCloud subnet.

!> **WARNING:** It has been deprecated, use `huaweicloud_vpc_subnet` instead.

## Example Usage

```hcl
data "huaweicloud_networking_subnet_v2" "subnet_1" {
  name = "subnet_1"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the V2 Neutron client. A Neutron client is needed to
  retrieve subnet ids. If omitted, the
  `region` argument of the provider is used.

* `name` - (Optional, String) The name of the subnet.

* `dhcp_enabled` - (Optional, Bool) If the subnet has DHCP enabled.

* `dhcp_disabled` - (Optional, Bool) If the subnet has DHCP disabled.

* `ip_version` - (Optional, Int) The IP version of the subnet (either 4 or 6).

* `gateway_ip` - (Optional, String) The IP of the subnet's gateway.

* `cidr` - (Optional, String) The CIDR of the subnet.

* `subnet_id` - (Optional, String) The ID of the subnet.

* `network_id` - (Optional, String) The ID of the network the subnet belongs to.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.
* `allocation_pools` - Allocation pools of the subnet.
* `enable_dhcp` - Whether the subnet has DHCP enabled or not.
* `dns_nameservers` - DNS Nameservers of the subnet.
* `host_routes` - Host Routes of the subnet.

