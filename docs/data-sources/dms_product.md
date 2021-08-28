---
subcategory: "Distributed Message Service (DMS)"
---

# huaweicloud_dms_product

Use this data source to get the ID of an available HuaweiCloud dms product. This is an alternative
to `huaweicloud_dms_product_v1`

## Example Usage

```hcl
data "huaweicloud_dms_product" "product1" {
  engine            = "kafka"
  version           = "1.1.0"
  instance_type     = "cluster"
  partition_num     = 300
  storage           = 600
  storage_spec_code = "dms.physical.storage.high"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the dms products. If omitted, the provider-level region
  will be used.

* `engine` - (Required, String) Indicates the name of a message engine.

* `version` - (Optional, String) Indicates the version of a message engine.

* `instance_type` - (Required, String) Indicates an instance type. Options: "single" and "cluster"

* `vm_specification` - (Optional, String) Indicates VM specifications.

* `storage` - (Optional, String) Indicates the message storage space.

* `bandwidth` - (Optional, String) Indicates the baseline bandwidth of a Kafka instance.

* `partition_num` - (Optional, String) Indicates the maximum number of topics that can be created for a Kafka instance.

* `storage_spec_code` - (Optional, String) Indicates an I/O specification.

* `io_type` - (Optional, String) Indicates an I/O type.

* `node_num` - (Optional, String) Indicates the number of nodes in a cluster.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.

