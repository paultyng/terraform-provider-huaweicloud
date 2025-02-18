---
subcategory: "GaussDB"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_gaussdb_opengauss_available_flavors"
description: |-
  Use this data source to get the specifications that a GaussDB OpenGauss instance can be changed to.
---

# huaweicloud_gaussdb_opengauss_available_flavors

Use this data source to get the specifications that a GaussDB OpenGauss instance can be changed to.

## Example Usage

```hcl
variable "instance_id" {}

data "huaweicloud_gaussdb_opengauss_available_flavors" "test" {
  instance_id = var.instance_id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the resource.
  If omitted, the provider-level region will be used.

* `instance_id` - (Required, String) Specifies the ID of the GaussDB OpenGauss instance.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `flavors` - Indicates the list of available flavors.

  The [flavors](#flavors_struct) structure is documented below.

<a name="flavors_struct"></a>
The `flavors` block supports:

* `spec_code` - Indicates the resource specification code.

* `vcpus` - Indicates the number of vCPUs.

* `ram` - Indicates the memory size in GB.

* `az_status` - Indicates the key/value pairs of the availability zone status.
  **key** indicates the AZ ID, and **value** indicates the specification status in the AZ.
  The **value** can be any of the following:
  + **normal**: available.
  + **unsupported**: not supported.
  + **sellout**: sold out.
