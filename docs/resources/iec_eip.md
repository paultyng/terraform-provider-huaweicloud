---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_eip

Manages a eip resource within HuaweiCloud IEC.

## Example Usage

```hcl
data "huaweicloud_iec_sites" "iec_sites" {}

resource "huaweicloud_iec_eip" "eip_test" {
  site_id = data.huaweicloud_iec_sites.iec_sites.sites[0].id
}
```

## Argument Reference

The following arguments are supported:

* `site_id` - (Required, String, ForceNew) Specifies the ID of IEC sevice site. Changing this parameter creates a new
  resource.

* `port_id` - (Optional, String) Specifies the port ID which this eip will associate with.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID in UUID format.
* `ip_version` - The version of elastic IP address.
* `status` - The status of elastic IP.
* `public_ip` - The address of elastic IP.
* `private_ip` - The address of private IP.
* `bandwitch_name` - The name of bandwidth.
* `bandwitch_size` - The size of bandwidth.
* `bandwitch_share_type` - Whether the bandwidth is shared or exclusive.
* `site_info` - The located information of the IEC site. It contains area, province and city.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 3 minute.

## Import

IEC EIPs can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_iec_eip.eip_test b5ad19d1-57d1-48fd-aab7-1378f9bee169
```
