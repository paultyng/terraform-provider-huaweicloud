---
subcategory: "Document Database Service (DDS)"
---

# huaweicloud_dds_database_versions

Use this data source to get the versions of DDS instances.

## Example Usage

```hcl
data "huaweicloud_dds_database_versions" "test1" {
  name = "DDS-Community"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the resources.
  If omitted, the provider-level region will be used.

* `datastore_name` - (Required, String) Specifies the database name.
  The valid values are **DDS-Community** and **DDS-Enhanced**.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `versions` - Indicates the database version.
