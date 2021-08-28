---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_flavors

Use this data source to get the available of HuaweiCloud IEC flavors.

## Example Usage

```hcl
variable "flavor_name" {
  default = "c6.large.2"
}

data "huaweicloud_iec_flavors" "iec_flavor_test" {
  name = var.flavor_name
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) The region in which to obtain the flavors. If omitted, the provider-level region will be
  used.

* `name` - (Optional, String) Specifies the flavor name, which can be queried with a regular expression.

* `site_ids` - (Optional, String) Specifies the list of edge service site.

* `area` - (Optional, String) Specifies the province of the iec instance located.

* `province` - (Optional, String) Specifies the province of the iec instance located.

* `city` - (Optional, String) Specifies the province of the iec instance located.

* `operator` - (Optional, String) Specifies the operator supported of the iec instance.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `flavors` - An array of one or more flavors. The flavors object structure is documented below.

The `flavors` block supports:

* `id` - The id of the iec flavor.
* `name` - The name of the iec flavor.
* `vcpus` - The vcpus of the iec flavor.
* `memory` - The memory of the iec flavor.
