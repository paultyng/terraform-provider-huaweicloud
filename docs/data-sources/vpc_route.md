---
subcategory: "Virtual Private Cloud (VPC)"
---

# huaweicloud_vpc_route

Provides details about a specific VPC route. This is an alternative to `huaweicloud_vpc_route_v2`

## Example Usage

```hcl
data "huaweicloud_vpc_route" "vpc_route" {
  vpc_id = var.vpc_id
}
```

## Argument Reference

The arguments of this data source act as filters for querying the available routes in the current tenant. The given
filters must match exactly one route whose data will be exported as attributes.

* `region` - (Optional, String) The region in which to obtain the vpc route. If omitted, the provider-level region will
  be used.

* `id` - (Optional, String) The id of the specific route to retrieve.

* `vpc_id` - (Optional, String) The id of the VPC that the desired route belongs to.

* `destination` - (Optional, String) The route destination address (CIDR).

* `tenant_id` - (Optional, String) Only the administrator can specify the tenant ID of other tenants.

* `type` - (Optional, String) Route type for filtering.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `nexthop` - The next hop of the route. If the route type is peering, it will provide VPC peering connection ID.
