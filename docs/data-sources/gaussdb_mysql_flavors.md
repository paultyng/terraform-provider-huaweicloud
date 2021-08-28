---
subcategory: "GaussDB"
---

# huaweicloud_gaussdb_mysql_flavors

Use this data source to get available HuaweiCloud gaussdb mysql flavors.

## Example Usage

```hcl
data "huaweicloud_gaussdb_mysql_flavors" "flavors" {
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the flavors. If omitted, the provider-level region will be
  used.

* `engine` - (Optional, String) Specifies the database engine. Only "gaussdb-mysql" is supported now.

* `version` - (Optional, String) Specifies the database version. Only "8.0" is supported now.

* `availability_zone_mode` - (Optional, String) Specifies the availability zone mode. Currently support `single` and '
  multi'. Defaults to `single`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.

* `flavors` - Indicates the flavors information. Structure is documented below.

The `flavors` block contains:

* `name` - The name of the gaussdb mysql flavor.
* `vcpus` - Indicates the CPU size.
* `memory` - Indicates the memory size in GB.
* `mode` - Indicates the database mode.
* `version` - Indicates the database version.
