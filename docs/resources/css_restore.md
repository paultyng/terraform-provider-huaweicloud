---
subcategory: "Cloud Search Service (CSS)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_css_restore"
description: |-
  Manages CSS cluster restore resource within HuaweiCloud
---

# huaweicloud_css_restore

Manages CSS cluster restore resource within HuaweiCloud

## Example Usage

### restore by snapshot_id

```hcl
variable "target_cluster_id" {}
variable "source_cluster_id" {}
variable "snapshot_id" {}

resource "huaweicloud_css_restore" "test" {
  source_cluster_id     = var.target_cluster_id
  target_cluster_id     = var.source_cluster_id
  snapshot_id           = var.snapshot_id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the resource. If omitted, the
  provider-level region will be used. Changing this creates a new resource.

* `target_cluster_id` - (Required, String, ForceNew) Specifies the target cluster ID.

  Changing this creates a new resource.

* `source_cluster_id` - (Required, String, ForceNew) Specifies the source cluster ID.

  Changing this creates a new resource.

* `snapshot_id` - (Required, String, ForceNew) Specifies the ID of the snapshot to be restored.

  Changing this creates a new resource.

* `indices` - (Optional, String, ForceNew) Name of an index to be restored. Multiple indexes are separated by commas (,).
  By default, all indexes are restored.You can use \ * to match multiple indexes. For example, if you specify 2018-06*,
  then the data of the indexes with the prefix 2018-06 will be restored.The value can contain 0 to 1,024 characters.
  Uppercase letters, spaces, and the following special characters are not allowed: "\<|>/?.

  Changing this creates a new resource.

* `rename_pattern` - (Optional, String, ForceNew) Rule for defining the indexes to be restored.The value can contain 0 to
  1,024 characters. Uppercase letters, spaces, and the following special characters are not allowed: "\<|>/?. Indexes
  that match this rule will be restored. The filtering condition must be a regular expression.

  Changing this creates a new resource.

* `rename_replacement` - (Optional, String, ForceNew) Rule for renaming an index. The value can contain 0 to 1,024
  characters. Uppercase letters, spaces, and the following special characters are not allowed: "\<|>/? For example,
  restored_index_$1 indicates adding the restored_ prefix to the names of all the restored indexes.The
  rename_replacement
  parameter takes effect only if rename_pattern has been enabled.

  Changing this creates a new resource.

## Attribute Reference

In addition to all arguments above, the following attribute is exported:

* `id` - The resource ID. The value is the restore job ID.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 60 minutes.
