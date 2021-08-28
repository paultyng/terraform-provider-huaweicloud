---
subcategory: "Distributed Cache Service"
---

# huaweicloud_dcs_az

Use this data source to get the ID of an available Huaweicloud dcs az.

## Example Usage

```hcl
data "huaweicloud_dcs_az" "az1" {
  code = "cn-north-1a"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the dcs az. If omitted, the provider-level region will be
  used.

* `code` - (Required, String) Specifies the code of an AZ, e.g. "cn-north-1a".

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Indicates a data source ID in UUID format.
* `name` - Indicates the name of an AZ.
* `port` - Indicates the port number of an AZ.
