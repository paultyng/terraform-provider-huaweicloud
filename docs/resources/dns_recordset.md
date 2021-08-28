---
subcategory: "Domain Name Service (DNS)"
---

# huaweicloud_dns_recordset

Manages a DNS record set in the HuaweiCloud DNS Service. This is an alternative to `huaweicloud_dns_recordset_v2`

## Example Usage

### Automatically detect the correct network

```hcl
resource "huaweicloud_dns_zone" "example_zone" {
  name        = "example.com."
  email       = "email2@example.com"
  description = "a zone"
  ttl         = 6000
  zone_type   = "public"
}

resource "huaweicloud_dns_recordset" "rs_example_com" {
  zone_id     = huaweicloud_dns_zone.example_zone.id
  name        = "rs.example.com."
  description = "An example record set"
  ttl         = 3000
  type        = "A"
  records     = [
    "10.0.0.1"]
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the DNS record set. If omitted, the `region`
  argument of the provider will be used. Changing this creates a new DNS record set.

* `zone_id` - (Required, String, ForceNew) The ID of the zone in which to create the record set. Changing this creates a
  new DNS record set.

* `name` - (Required, String, ForceNew) The name of the record set. Note the `.` at the end of the name. Changing this
  creates a new DNS record set.

* `type` - (Required, String, ForceNew) The type of record set. The options include `A`, `AAAA`, `MX`,
  `CNAME`, `TXT`, `NS`, `SRV`, `CAA`, and `PTR`. Changing this creates a new DNS record set.

* `records` - (Required, List) An array of DNS records.

* `ttl` - (Optional, Int) The time to live (TTL) of the record set (in seconds). The value range is 300–2147483647. The
  default value is 300.

* `description` - (Optional, String) A description of the record set.

* `tags` - (Optional, Map) The key/value pairs to associate with the record set.

* `value_specs` - (Optional, Map, ForceNew) Map of additional options. Changing this creates a new record set.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `update` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash.

```
$ terraform import huaweicloud_dns_recordset.recordset_1 < zone_id >/< recordset_id >
```
