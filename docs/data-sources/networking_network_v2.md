---
subcategory: "Deprecated"
---

# huaweicloud\_networking\_network\_v2

Use this data source to get the ID of an available HuaweiCloud network.

!> **WARNING:** It has been deprecated, use `huaweicloud_vpc_subnet` instead.

## Example Usage

```hcl
data "huaweicloud_networking_network_v2" "network" {
  name = "tf_test_network"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the V2 Neutron client. A Neutron client is needed to
  retrieve networks ids. If omitted, the
  `region` argument of the provider is used.

* `network_id` - (Optional, String) The ID of the network.

* `name` - (Optional, String) The name of the network.

* `status` - (Optional, String) The status of the network.

* `matching_subnet_cidr` - (Optional, String) The CIDR of a subnet within the network.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.

* `admin_state_up` - The administrative state of the network.

* `shared` - Specifies whether the network resource can be accessed by any tenant or not.
