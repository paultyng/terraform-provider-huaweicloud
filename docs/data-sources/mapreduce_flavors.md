---
subcategory: "MapReduce Service (MRS)"
---

# huaweicloud_mapreduce_flavors

Use this data source to get available cluster flavors of MapReduce.

## Example Usage

```hcl
data "huaweicloud_mapreduce_flavors" "test" {
  version_name = 'MRS 3.2.0-LTS.1'
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the data source.
  If omitted, the provider-level region will be used.

* `version_name` - (Required, String) The version of cluster.

* `availability_zone` - (Optional, String) The AZ name.

* `node_type` - (Optional, String) The node type supported by this flavor.  
  The options are as follows:
    - **master**: Indicates the master node.
    - **core**: Indicates the core node.
    - **task**: Indicates the task node.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.

* `flavors` - List of available cluster flavors.
  The [flavors](#MrsFlavors_Flavor) structure is documented below.

<a name="MrsFlavors_Flavor"></a>
The `flavors` block supports:

* `flavor_id` - The flavor ID.

* `version_name` - The version of cluster.

* `availability_zone` - The availability zone.

* `node_type` - The node type supported by this flavor.  
  The options are as follows:
    - **master**: Indicates the master node.
    - **core**: Indicates the core node.
    - **task**: Indicates the task node.
