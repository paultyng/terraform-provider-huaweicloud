# huaweicloud_availability_zones

Use this data source to get a list of availability zones from HuaweiCloud

## Example Usage

```hcl
data "huaweicloud_availability_zones" "zones" {}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the available zones. If omitted, the provider-level region
  will be used.

* `state` - (Optional, String) The `state` of the availability zones to match, default ("available").

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.

* `names` - The names of the availability zones, ordered alphanumerically, that match the queried `state`
