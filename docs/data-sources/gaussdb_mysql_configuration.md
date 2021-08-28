---
subcategory: "GaussDB"
---

# huaweicloud_gaussdb_mysql_configuration

Use this data source to get available HuaweiCloud gaussdb mysql configuration.

## Example Usage

```hcl
data "huaweicloud_gaussdb_mysql_configuration" "this" {
  name = "Default-GaussDB-for-MySQL 8.0"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the configurations. If omitted, the provider-level region
  will be used.

* `name` - (Optional, String) Specifies the name of the parameter template.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Indicates the ID of the configuration.
* `description` - Indicates the description of the configuration.
* `datastore_name` - Indicates the datastore name of the configuration.
* `datastore_version` - Indicates the datastore version of the configuration.
