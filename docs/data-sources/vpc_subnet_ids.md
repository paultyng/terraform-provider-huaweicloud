---
subcategory: "Virtual Private Cloud (VPC)"
---

# huaweicloud_vpc_subnet_ids

Provides a list of subnet ids for a vpc_id This is an alternative to `huaweicloud_vpc_subnet_ids_v1`

This resource can be useful for getting back a list of subnet ids for a vpc.

## Example Usage

The following example shows outputing all cidr blocks for every subnet id in a vpc.

```hcl
data "huaweicloud_vpc_subnet_ids" "subnet_ids" {
  vpc_id = var.vpc_id
}

data "huaweicloud_vpc_subnet" "subnet" {
  count = length(data.huaweicloud_vpc_subnet_ids.subnet_ids.ids)
  id    = data.huaweicloud_vpc_subnet_ids.subnet_ids.ids[count.index]
}

output "subnet_cidr_blocks" {
  value = [for s in data.huaweicloud_vpc_subnet.subnet: "${s.name}: ${s.id}: ${s.cidr}"]
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) The region in which to obtain the subnet ids. If omitted, the provider-level region will
  be used.

* `vpc_id` - (Required, String) Specifies the VPC ID used as the query filter.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.

* `ids` - A set of all the subnet ids found. This data source will fail if none are found.
