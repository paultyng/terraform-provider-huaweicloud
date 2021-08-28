---
subcategory: "Distributed Message Service (DMS)"
---

# huaweicloud_dms_az

Use this data source to get the ID of an available HuaweiCloud dms az. This is an alternative to `huaweicloud_dms_az_v1`

## Example Usage

```hcl
data "huaweicloud_dms_az" "az1" {
  code = "cn-north-4a"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the dms az. If omitted, the provider-level region will be
  used.

* `code` - (Optional, String) Specifies the code of an AZ.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Indicates a data source ID in UUID format.

* `name` - Indicates the name of an AZ.

* `port` - Indicates the port number of an AZ.
