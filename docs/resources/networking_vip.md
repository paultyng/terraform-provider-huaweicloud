---
subcategory: "Virtual Private Cloud (VPC)"
---

# huaweicloud_networking_vip

Manages a Vip resource within HuaweiCloud. This is an alternative to `huaweicloud_networking_vip_v2`

## Example Usage

```hcl
data "huaweicloud_vpc_subnet" "mynet" {
  name = "subnet-default"
}

resource "huaweicloud_networking_vip" "myvip" {
  network_id = data.huaweicloud_vpc_subnet.mynet.id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, ForceNew) The region in which to create the vip resource. If omitted, the provider-level region
  will be used.

* `network_id` - (Required, ForceNew) Specifies the ID of the network to which the vip belongs.

* `subnet_id` - (Optional, ForceNew) Specifies the subnet in which to allocate IP address for this vip.

* `ip_address` - (Optional, ForceNew) Specifies the IP address desired in the subnet for this vip. If you don't
  specify `ip_address`, an available IP address from the specified subnet will be allocated to this vip.

* `name` - (Optional, String) Specifies a unique name for the vip.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the vip.
* `mac_address` - The MAC address of the vip.
* `status` - The status of vip.
* `tenant_id` - The tenant ID of the vip.
* `device_owner` - The device owner of the vip.

## Import

Networking VIP can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_networking_vip.myvip ce595799-da26-4015-8db5-7733c6db292e
```
