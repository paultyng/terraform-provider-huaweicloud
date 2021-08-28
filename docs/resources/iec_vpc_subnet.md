---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_vpc_subnet

Manages a VPC subnet resource within HuaweiCloud IEC.

## Example Usage

```hcl
data "huaweicloud_iec_sites" "sites_test" {}

resource "huaweicloud_iec_vpc" "vpc_test" {
  name = "vpc_demo"
  cidr = "192.168.0.0/16"
  mode = "CUSTOMER"
}

resource "huaweicloud_iec_vpc_subnet" "subnet_test" {
  name       = "subnet_demo"
  cidr       = "192.168.128.0/18"
  vpc_id     = huaweicloud_iec_vpc.vpc_test.id
  site_id    = data.huaweicloud_iec_sites.sites_test.sites[0].id
  gateway_ip = "192.168.128.1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Specifies the name of the iec vpc subnet. The value is a string of 1 to 64 characters that
  can contain letters, digits, underscores(_), and hyphens(-).

* `cidr` - (Required, String, ForceNew) Specifies the network segment on which the subnet resides. The value must be in
  CIDR format and within the CIDR block of the iec vpc. Changing this parameter creates a new subnet resource.

* `vpc_id` - (Required, String, ForceNew) Specifies the ID of the iec __CUSTOMER__
  vpc to which the subnet belongs. Changing this parameter creates a new subnet resource.

* `site_id` - (Required, String, ForceNew) Specifies the ID of the iec site. Changing this parameter creates a new
  subnet resource.

* `gateway_ip` - (Required, String, ForceNew)  Specifies the gateway of the subnet. The value must be a valid IP address
  and in the subnet segment. Changing this parameter creates a new subnet resource.

* `dns_list` - (Optional, List) Specifies the DNS server address list of a subnet. These DNS server address must be
  valid IP addresses.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID in UUID format.

* `dhcp_enable` - The status of subnet DHCP is enabled or not.

* `site_info` - The located information of the iec site. It contains area, province and city.

* `status` - The status of the subnet.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 3 minute.

## Import

IEC vpc subnet can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_iec_vpc_subnet.subnet_demo 51be9f2b-5a3b-406a-9271-36f0c929fbcc
```
