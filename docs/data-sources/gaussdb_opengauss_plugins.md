---
subcategory: "GaussDB"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_gaussdb_opengauss_plugins"
description: |-
  Use this data source to get list of supported plug-ins.
---

# huaweicloud_gaussdb_opengauss_plugins

Use this data source to get the list of supported plug-ins.

## Example Usage

```hcl
data "huaweicloud_gaussdb_opengauss_plugins" "test" {}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the resource.
  If omitted, the provider-level region will be used.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `plugins` - Indicates the list plug-ins.
