---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_vpc_subnets

Use this data source to get a list of subnets belong to a specific IEC VPC.

## Example Usage

```hcl
variable "vpc_id" {}

data "huaweicloud_iec_vpc_subnets" "all_subnets" {
  vpc_id = var.vpc_id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) The region in which to obtain the subnets. If omitted, the provider-level region will be
  used.

* `vpc_id` - (Required, String) Specifies the ID of the IEC VPC.

* `site_id` - (Optional, String) Specifies the ID of the IEC site.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `subnets` - A list of all the subnets found. The object is documented below.

The `subnets` block supports:

* `id` - Indicates the ID of the subnet.
* `name` - Indicates the name of the subnet.
* `cidr` - Indicates the cidr block of the subnet.
* `gateway_ip` - Indicates the gateway of the subnet.
* `dns_list` - Indicates the DNS server address list of the subnet.
* `site_id` - Indicates the ID of the IEC site.
* `site_info` - Indicates the located information of the iec site. It contains area, province and city.
* `status` - Indicates the status of the subnet.
