---
subcategory: "Deprecated"
---

# huaweicloud_vpnaas_service_v2

Manages a V2 VPN service resource within HuaweiCloud.

## Example Usage

```hcl
resource "huaweicloud_vpnaas_service_v2" "service_1" {
  name           = "my_service"
  router_id      = "14a75700-fc03-4602-9294-26ee44f366b3"
  admin_state_up = "true"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional) The region in which to create the V2 Networking client. A Networking client is needed to create
  a VPN service. If omitted, the
  `region` argument of the provider is used. Changing this creates a new service.

* `name` - (Optional) The name of the service. Changing this updates the name of the existing service.

* `description` - (Optional) The human-readable description for the service. Changing this updates the description of
  the existing service.

* `admin_state_up` - (Optional) The administrative state of the resource. Can either be up(true) or down (false).
  Defaults to `true`. Changing this updates the administrative state of the existing service.

* `subnet_id` - (Optional) SubnetID is the ID of the subnet. Default is null.

* `router_id` - (Required) The ID of the router. Changing this creates a new service.

* `value_specs` - (Optional) Map of additional options.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.
* `status` - Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR,
  PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
* `external_v6_ip` - The read-only external (public) IPv6 address that is used for the VPN service.
* `external_v4_ip` - The read-only external (public) IPv4 address that is used for the VPN service.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `update` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

Services can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_vpnaas_service_v2.service_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
```
