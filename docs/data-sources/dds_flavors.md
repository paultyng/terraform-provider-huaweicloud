---
subcategory: "Document Database Service (DDS)"
---

# huaweicloud_dds_flavors

Use this data source to get the ID of an available HuaweiCloud dds flavor. This is an alternative
to `huaweicloud_dds_flavors_v3`

## Example Usage

```hcl
data "huaweicloud_dds_flavors" "flavor" {
  engine_name = "DDS-Community"
  vcpus       = 8
}
```

## Argument Reference

* `region` - (Optional, String) Specifies the region in which to obtain the V3 dds client. If omitted, the
  provider-level region will be used.

* `engine_name` - (Required, String) Specifies the engine name of the dds, "DDS-Community" and "DDS-Enhanced" are
  supported.

* `type` - (Optional, String) Specifies the type of the dds falvor. "mongos", "shard", "config", "replica" and "single"
  are supported.

* `vcpus` - (Optional, String) Specifies the vcpus of the dds flavor.

* `memory` - (Optional, String) Specifies the ram of the dds flavor in GB.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.

* `flavors` - Indicates the flavors information. Structure is documented below.

The `flavors` block contains:

* `spec_code - The name of the rds flavor.
* `type` - See `type` above.
* `vcpus` - See `vcpus` above.
* `memory` - See 'memory' above.
