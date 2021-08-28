---
subcategory: "Elastic IP (EIP)"
---

# huaweicloud_vpc_bandwidth

Manages a Shared Bandwidth resource within HuaweiCloud. This is an alternative to `huaweicloud_vpc_bandwidth_v2`

## Example Usage

```hcl
resource "huaweicloud_vpc_bandwidth" "bandwidth_1" {
  name = "bandwidth_1"
  size = 5
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the Shared Bandwidth. If omitted, the
  provider-level region will be used. Changing this creates a new Shared Bandwidth resource.

* `name` - (Required, String) The name of the Shared Bandwidth.

* `size` - (Required, Int) The size of the Shared Bandwidth. The value ranges from 5 to 2000 G.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id of the Shared Bandwidth. Changing
  this creates a new bandwidth.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the Shared Bandwidth.

* `share_type` - Indicates whether the bandwidth is shared or dedicated.

* `bandwidth_type` - Indicates the bandwidth type.

* `charge_mode` - Indicates whether the billing is based on traffic, bandwidth, or 95th percentile bandwidth (enhanced).

* `status` - Indicates the bandwidth status.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `update` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

Shared Bandwidths can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_vpc_bandwidth.bandwidth_1 7117d38e-4c8f-4624-a505-bd96b97d024c
```
