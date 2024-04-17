---
subcategory: "Virtual Private Cloud (VPC)"
---

# huaweicloud_vpc_address_groups

Use this data source to get a list of VPC IP address groups.

## Example Usage

```hcl
data "huaweicloud_vpc_address_groups" "demo1" {
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the resource.
  If omitted, the provider-level region will be used.

* `group_id` - (Optional, String) Unique ID of an IP address group, which can be used to filter the IP address group.

* `name` - (Optional, String) Name of an IP address group, which can be used to filter the IP address group.

* `ip_version` - (Optional, Int) Version of IP addresses in an IP address group,
  which can be used to filter the IP address group.

* `description` - (Optional, String) Provides supplementary information about an IP address group,
  which can be used to filter the IP address group.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `address_groups` - The IP address groups.

  The [address_groups](#address_groups_struct) structure is documented below.

<a name="address_groups_struct"></a>
The `address_groups` block supports:

* `id` - IP address group ID, which uniquely identifies the IP address group.

* `name` - IP address group name.

* `ip_version` - Whether it is an IPv4 or IPv6 address group.

* `status` - The status of IP address group.
  Valid values are:
    + `NORMAL`: normal status.
    + `UPDATING`: updating.
    + `UPDATE_FAILED`: update failed.
  When the status of IP address group is `UPDATING`, the IP address group cannot be updated again.

* `max_capacity` - Maximum number of entries in an address group,
  which limits the number of addresses that can be contained in an address group.

* `description` - The supplementary information about the IP address group.

* `status_message` - The status details of IP address group.

* `addresses` - IP address sets in an IP address group.
  Value range: a single IP address, IP address range, or CIDR block.

* `ip_extra_set` - IP addresses and their remarks in an IP address group.

  The [ip_extra_set](#address_groups_ip_extra_set_struct) structure is documented below.

* `enterprise_project_id` - Enterprise project ID.

* `updated_at` - Time when the IP address group was last updated.

* `created_at` - Time when the IP address group is created.

<a name="address_groups_ip_extra_set_struct"></a>
The `ip_extra_set` block supports:

* `ip` - An IP address, IP address range, or CIDR block.

* `remarks` - Provides supplementary information about the IP address, IP address range, or CIDR block.
