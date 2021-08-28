---
subcategory: "Content Delivery Network (CDN)"
---

# huaweicloud_cdn_domain

CDN domain management This is an alternative to `huaweicloud_cdn_domain_v1`

## Example Usage

### create a cdn domain

```hcl
resource "huaweicloud_cdn_domain" "domain_1" {
  name = var.domain_name
  type = "web"

  sources {
    origin      = var.origin_server
    origin_type = "ipaddr"
    active      = 1
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String, ForceNew) The acceleration domain name. Changing this parameter will create a new
  resource.

* `type` - (Required, String, ForceNew) The service type. The valid values are  'web', 'download' and 'video'. Changing
  this parameter will create a new resource.

* `sources` - (Required, List, ForceNew) An array of one or more objects specifies the domain name of the origin server.
  The sources object structure is documented below.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id. Changing this parameter will create
  a new resource.

The `sources` block supports:

* `origin` - (Required, String) The domain name or IP address of the origin server.

* `origin_type` - (Required, String) The origin server type. The valid values are 'ipaddr', 'domain', and 'obs_bucket'.

* `active` - (Optional, Int) Whether an origin server is active or standby (1: active; 0: standby). The default value is
  1.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The acceleration domain name ID.

* `cname` - The CNAME of the acceleration domain name.

* `domain_status` - The status of the acceleration domain name. The available values are
  'online', 'offline', 'configuring', 'configure_failed', 'checking', 'check_failed' and 'deleting.'

* `service_area` - The area covered by the acceleration service.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 20 minute.
* `delete` - Default is 20 minute.

## Import

Domains can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_cdn_domain.domain_1 fe2462fac09a4a42a76ecc4a1ef542f1
```
