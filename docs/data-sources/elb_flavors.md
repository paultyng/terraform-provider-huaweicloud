---
subcategory: "Dedicated Load Balance (Dedicated ELB)"
---

# huaweicloud_elb_flavors

Use this data source to get the available ELB Flavors.

## Example Usage

```hcl
data "huaweicloud_elb_flavors" "flavors" {
  type            = "L7"
  max_connections = 200000
  cps             = 2000
  bandwidth       = 50
}

# Create Dedicated Load Balancer with the first matched flavor

resource "huaweicloud_elb_loadbalancer" "lb" {
  l7_flavor_id = data.huaweicloud_elb_flavors.flavors.ids[0]

  # Other properties...
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the flavors. If omitted, the provider-level region will be
  used.

* `type` - (Optional, String) Specifies the flavor type. Valid values are L4 and L7.

* `max_connections` - (Optional, Int) Specifies the maximum connections in the flavor.

* `bandwidth` - (Optional, Int) Specifies the bandwidth size(Mbit/s) in the flavor.

* `cps` - (Optional, Int) Specifies the cps in the flavor.

* `qps` - (Optional, Int) Specifies the qps in the L7 flavor.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.

* `ids` - A list of flavor IDs.

* `flavors` - A list of flavors. Each element contains the following attributes:
  + `id` - ID of the flavor.
  + `name` - Name of the flavor.
  + `type` - Type of the flavor.
  + `max_connections` - Maximum connections of the flavor.
  + `cps` - Cps of the flavor.
  + `qps` - Qps of the L7 flavor.
  + `bandwidth` - Bandwidth size(Mbit/s) of the flavor.
