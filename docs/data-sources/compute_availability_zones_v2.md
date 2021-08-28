---
subcategory: "Deprecated"
---

# huaweicloud\_compute\_availability\_zones\_v2

Use this data source to get a list of availability zones from HuaweiCloud

!> **WARNING:** It has been deprecated, use `huaweicloud_availability_zones` instead.

## Example Usage

```hcl
data "huaweicloud_compute_availability_zones_v2" "zones" {}
```

## Argument Reference

* `state` - (Optional, String) The `state` of the availability zones to match, default ("available").

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.

* `names` - The names of the availability zones, ordered alphanumerically, that match the queried `state`
