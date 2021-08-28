---
subcategory: "Virtual Private Cloud (VPC)"
---

# huaweicloud_networking_vip_associate

Manages a Vip associate resource within HuaweiCloud. This is an alternative to `huaweicloud_networking_vip_associate_v2`

## Example Usage

```hcl
data "huaweicloud_vpc_subnet" "mynet" {
  name = "subnet-default"
}

resource "huaweicloud_networking_vip" "myvip" {
  network_id = data.huaweicloud_vpc_subnet.mynet.id
}

resource "huaweicloud_networking_vip_associate" "vip_associated" {
  vip_id   = huaweicloud_networking_vip.myvip.id
  port_ids = [
    var.port_1,
    var.port_2]
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, ForceNew) The region in which to create the vip associate resource. If omitted, the
  provider-level region will be used.

* `vip_id` - (Required, ForceNew) The ID of vip to attach the ports to.

* `port_ids` - (Required, List) An array of one or more IDs of the ports to attach the vip to.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.
* `vip_subnet_id` - The ID of the subnet this vip connects to.
* `vip_ip_address` - The IP address in the subnet for this vip.
* `ip_addresses` - The IP addresses of ports to attach the vip to.
