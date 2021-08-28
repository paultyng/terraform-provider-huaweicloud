---
subcategory: "GaussDB"
---

# huaweicloud_gaussdb_mysql_dedicated_resource

Use this data source to get available HuaweiCloud gaussdb mysql dedicated resource.

## Example Usage

```hcl
data "huaweicloud_gaussdb_mysql_dedicated_resource" "this" {
  resource_name = "test"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the dedicated resource. If omitted, the provider-level
  region will be used.

* `resource_name` - (Optional, String) Specifies the dedicated resource name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Indicates the ID of the instance.

* `status` - Indicates the status of the dedicated resource.

* `availability_zone` - Indicates the availability zone of the dedicated resource.

* `architecture` - Indicates the architecture of the dedicated resource.

* `vcpus` - Indicates the vcpus count of the dedicated resource.

* `ram` - Indicates the ram size of the dedicated resource.

* `volume` - Indicates the volume size of the dedicated resource.
