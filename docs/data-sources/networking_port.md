---
subcategory: "Virtual Private Cloud (VPC)"
---

# huaweicloud_networking_port

Use this data source to get the ID of an available HuaweiCloud port. This is an alternative
to `huaweicloud_networking_port_v2`

## Example Usage

```hcl
data "huaweicloud_networking_port" "port_1" {
  network_id = var.network_id
  fixed_ip   = "192.168.0.100"
}
```

## Argument Reference

* `region` - (Optional, String) Specifies the region in which to obtain the port. If omitted, the provider-level region
  will be used.

* `port_id` - (Optional, String) Specifies the ID of the port.

* `network_id` - (Optional, String) Specifies the ID of the network the port belongs to.

* `fixed_ip` - (Optional, String) Specifies the port IP address filter.

* `mac_address` - (Optional, String) Specifies the MAC address of the port.

* `status` - (Optional, String) Specifies the status of the port.

* `security_group_ids` - (Optional, String) The list of port security group IDs to filter.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID in UUID format.

* `name` - The name of the port.

* `admin_state_up` - The administrative state of the port.

* `device_owner` - The device owner of the port.

* `device_id` - The ID of the device the port belongs to.

* `all_fixed_ips` - The collection of Fixed IP addresses on the port.

* `all_security_group_ids` - The collection of security group IDs applied on the port.
