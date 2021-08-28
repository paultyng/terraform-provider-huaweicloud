---
subcategory: "Deprecated"
---

# huaweicloud_networking_network_v2

Manages a V2 Neutron network resource within HuaweiCloud.

!> **Warning:** It has been deprecated, use `huaweicloud_vpc_subnet` instead.

## Example Usage

```hcl
resource "huaweicloud_networking_network_v2" "network_1" {
  name           = "network_1"
  admin_state_up = "true"
}

resource "huaweicloud_networking_subnet_v2" "subnet_1" {
  name       = "subnet_1"
  network_id = huaweicloud_networking_network_v2.network_1.id
  cidr       = "192.168.199.0/24"
  ip_version = 4
}

resource "huaweicloud_networking_secgroup" "secgroup_1" {
  name        = "secgroup_1"
  description = "a security group"
}

resource "huaweicloud_networking_secgroup_rule" "secgroup_rule_1" {
  direction         = "ingress"
  ethertype         = "IPv4"
  port_range_max    = 22
  port_range_min    = 22
  protocol          = "tcp"
  remote_ip_prefix  = "0.0.0.0/0"
  security_group_id = huaweicloud_networking_secgroup.secgroup_1.id
}

resource "huaweicloud_networking_port_v2" "port_1" {
  name               = "port_1"
  network_id         = huaweicloud_networking_network_v2.network_1.id
  admin_state_up     = "true"
  security_group_ids = [huaweicloud_networking_secgroup.secgroup_1.id]

  fixed_ip {
    subnet_id  = huaweicloud_networking_subnet_v2.subnet_1.id
    ip_address = "192.168.199.10"
  }
}

resource "huaweicloud_compute_instance_v2" "instance_1" {
  name            = "instance_1"
  security_groups = [huaweicloud_networking_secgroup.secgroup_1.name]

  network {
    port = huaweicloud_networking_port_v2.port_1.id
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the V2 Networking client.
    A Networking client is needed to create a Neutron network. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    network.

* `name` - (Optional, String) The name of the network. Changing this updates the name of
    the existing network.

* `shared` - (Optional, String)  Specifies whether the network resource can be accessed
    by any tenant or not. Changing this updates the sharing capabalities of the
    existing network.

* `admin_state_up` - (Optional, String) The administrative state of the network.
    Acceptable values are "true" and "false". Changing this value updates the
    state of the existing network.

* `segments` - (Optional, List, ForceNew) An array of one or more provider segment objects.

* `value_specs` - (Optional, Map, ForceNew) Map of additional options.

The `segments` block supports:

* `physical_network` - (Optional, String, ForceNew) The phisical network where this network is implemented.
* `segmentation_id` - (Optional, Int, ForceNew) An isolated segment on the physical network.
* `network_type` - (Optional, String, ForceNew) The type of physical network.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

## Timeouts
This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

Networks can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_networking_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9
```
