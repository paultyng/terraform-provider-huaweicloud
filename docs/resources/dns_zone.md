---
subcategory: "Domain Name Service (DNS)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_dns_zone"
description: ""
---

# huaweicloud_dns_zone

Manages a DNS zone in the HuaweiCloud DNS Service.

## Example Usage

### Create a public DNS zone

```hcl
resource "huaweicloud_dns_zone" "my_public_zone" {
  name        = "example.com."
  email       = "jdoe@example.com"
  description = "An example zone"
  ttl         = 3000
  zone_type   = "public"
}
```

### Create a private DNS zone

```hcl
resource "huaweicloud_dns_zone" "my_private_zone" {
  name        = "1.example.com."
  email       = "jdoe@example.com"
  description = "An example zone"
  ttl         = 3000
  zone_type   = "private"

  router {
    router_id = "2c1fe4bd-ebad-44ca-ae9d-e94e63847b75"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the DNS zone. If omitted, the `region` argument
  of the provider will be used. Changing this creates a new DNS zone.

* `name` - (Required, String, ForceNew) The name of the zone. Note the `.` at the end of the name. Changing this creates
  a new DNS zone.

* `email` - (Optional, String) The email address of the administrator managing the zone.

* `zone_type` - (Optional, String, ForceNew) The type of zone. Can either be `public` or `private`. Changing this
  creates a new DNS zone.

* `router` - (Optional, List) Router configuration block which is required if zone_type is private. The router
  structure is documented below.

  -> Before changing this parameter, make sure the zone status is enabled.

* `ttl` - (Optional, Int) The time to live (TTL) of the zone.  
  The valid value is range from `1` to `2,147,483,647`.

* `description` - (Optional, String) A description of the zone.  
  A maximum of `255` characters are allowed.

* `tags` - (Optional, Map) The key/value pairs to associate with the zone.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id of the zone. Changing this creates a
  new zone.

* `status` - (Optional, String) Specifies the status of the zone.  
  The valid values are as follows:
  + **ENABLE**
  + **DISABLE**

  -> This is a one-time action.

* `proxy_pattern` - (Optional, String) Specifies the recursive resolution proxy mode for subdomains of the private zone.  
Defaults to **AUTHORITY**.
  The valid values are as follows:
  + **AUTHORITY**: The recursive resolution proxy is disabled for the private zone.
  + **RECURSIVE**: The recursive resolution proxy is enabled for the private zone.

The `router` block supports:

* `router_id` - (Required, String) ID of the associated VPC.

* `router_region` - (Optional, String) The region of the VPC.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

* `masters` - An array of master DNS servers.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minutes.
* `update` - Default is 10 minutes.
* `delete` - Default is 10 minutes.

## Import

This resource can be imported by specifying the zone ID:

```bash
$ terraform import huaweicloud_dns_zone.zone_1 <zone_id>
```
