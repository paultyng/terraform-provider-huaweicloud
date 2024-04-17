---
subcategory: "Cloud Connect (CC)"
---

# huaweicloud_cc_connection_routes

Use this data source to get the list of cloud connection routes.

## Example Usage

```hcl
variable "cloud_connection_route_id" {}

data "huaweicloud_cc_connection_routes" "test" {
  cloud_connection_route_id = var.cloud_connection_route_id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the resource.
  If omitted, the provider-level region will be used.

* `cloud_connection_route_id` - (Optional, String) Specifies cloud connection route ID.

* `cloud_connection_id` - (Optional, String) Specifies cloud connection ID.

* `instance_id` - (Optional, String) Specifies network instance ID of cloud connection route.

* `region_id` - (Optional, String) Specifies region ID of cloud connection route.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `cloud_connection_routes` - The list of cloud connection routes.

  The [cloud_connection_routes](#cloud_connection_routes_struct) structure is documented below.

<a name="cloud_connection_routes_struct"></a>
The `cloud_connection_routes` block supports:

* `id` - The cloud connection route ID.

* `cloud_connection_id` - The cloud connection ID.

* `instance_id` - The network instance ID of cloud connection route.

* `region_id` - The region ID of cloud connection route.

* `project_id` - The project ID of cloud connection route.

* `type` - The type of the network instance that the next hop of a route points to.

* `destination` - The destination address.
