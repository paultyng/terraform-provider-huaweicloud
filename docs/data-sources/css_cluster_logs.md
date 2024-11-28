---
subcategory: "Cloud Search Service (CSS)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_css_cluster_logs"
description: |-
  Use this data source to get the list of CSS cluster logs.
---

# huaweicloud_css_cluster_logs

Use this data source to get the list of CSS cluster logs.

-> **NOTE:** Up to 100 logs can be retrieved.

## Example Usage

```hcl
variable "cluster_id" {}
variable "instance_name" {}

data "huaweicloud_css_cluster_logs" "test" {
  cluster_id    = var.cluster_id
  instance_name = var.instance_name
  log_type      = "instance"
  level         = "WARN"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the resource.
  If omitted, the provider-level region will be used.

* `cluster_id` - (Required, String) Specifies the ID of the cluster.

* `instance_name` - (Required, String) Specifies the node name.

* `log_type` - (Required, String) Specifies the log type.
  The types of logs that can be queried are **deprecation**, **indexingSlow**, **searchSlow**, and **instance**.

* `level` - (Required, String) Specifies the log level.
  The levels of logs that can be queried are **INFO**, **ERROR**, **DEBUG**, and **WARN**.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `logs` - The log list.

  The [logs](#logs_struct) structure is documented below.

<a name="logs_struct"></a>
The `logs` block supports:

* `level` - The log level.

* `date` - The log date.

* `content` - The log content.
