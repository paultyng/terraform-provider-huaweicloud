---
subcategory: "Cloud Phone (CPH)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_cph_phone_data_export"
description: |-
  Manages a CPH phone data export resource within HuaweiCloud.
---

# huaweicloud_cph_phone_data_export

Manages a CPH phone data export resource within HuaweiCloud.

-> The current resource is a one-time resource, and destroying this resource is only removed from the state.

## Example Usage

```hcl
variable "phone_id" {}
variable "bucket_name" {}
variable "object_path" {}
variable "include_files" {}
variable "exclude_files" {}

resource "huaweicloud_cph_phone_data_export" "test" {
  phone_id      = var.phone_id
  bucket_name   = var.bucket_name
  object_path   = var.object_path
  include_files = var.include_files
  exclude_files = var.exclude_files
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the resource.
  If omitted, the provider-level region will be used.
  Changing this creates a new resource.

* `phone_id` - (Required, String) Specifies the phone ID.

* `bucket_name` - (Required, String) Specifies the bucket name of OBS.

* `object_path` - (Required, String) Specifies the object path of OBS.

* `include_files` - (Required, List) Specifies the include files.

* `exclude_files` - (Optional, List) Specifies the exclude files.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 30 minutes.
