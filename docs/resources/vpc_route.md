---
subcategory: "Virtual Private Cloud (VPC)"
---

# huaweicloud_vpc_route

Provides a resource to create a route. This is an alternative to `huaweicloud_vpc_route_v2`

## Example Usage

```hcl
resource "huaweicloud_vpc_route" "vpc_route" {
  type        = "peering"
  nexthop     = var.nexthop
  destination = "192.168.0.0/16"
  vpc_id      = var.vpc_id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the vpc route. If omitted, the provider-level
  region will be used. Changing this creates a new Route resource.

* `destination` (Required, String, ForceNew) - Specifies the destination IP address or CIDR block. Changing this creates
  a new Route.

* `nexthop` (Required, String, ForceNew) - Specifies the next hop. If the route type is peering, enter the VPC peering
  connection ID. Changing this creates a new Route.

* `type` (Required, String, ForceNew) - Specifies the route type. Currently, the value can only be **peering**. Changing
  this creates a new Route.

* `vpc_id` (Required, String, ForceNew) - Specifies the VPC for which a route is to be added. Changing this creates a
  new Route.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The route ID.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 10 minute.

