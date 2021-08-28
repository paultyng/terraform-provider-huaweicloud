---
subcategory: "VPC Endpoint (VPCEP)"
---

# huaweicloud_vpcep_approval

Provides a resource to manage the VPC endpoint connections.

## Example Usage

```hcl
variable "service_vpc_id" {}
variable "vm_port" {}
variable "vpc_id" {}
variable "network_id" {}

resource "huaweicloud_vpcep_service" "demo" {
  name        = "demo-service"
  server_type = "VM"
  vpc_id      = var.service_vpc_id
  port_id     = var.vm_port
  approval    = true

  port_mapping {
    service_port  = 8080
    terminal_port = 80
  }
}

resource "huaweicloud_vpcep_endpoint" "demo" {
  service_id = huaweicloud_vpcep_service.demo.id
  vpc_id     = var.vpc_id
  network_id = var.network_id
  enable_dns = true

  lifecycle {
    # enable_dns and ip_address are not assigned until connecting to the service
    ignore_changes = [
      enable_dns,
      ip_address]
  }
}

resource "huaweicloud_vpcep_approval" "approval" {
  service_id = huaweicloud_vpcep_service.demo.id
  endpoints  = [
    huaweicloud_vpcep_endpoint.demo.id]
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to obtain the VPC endpoint service. If omitted, the
  provider-level region will be used. Changing this creates a new resource.

* `service_id` (Required, String, ForceNew) - Specifies the ID of the VPC endpoint service. Changing this creates a new
  resource.

* `endpoints` (Required, List) - Specifies the list of VPC endpoint IDs which accepted to connect to VPC endpoint
  service. The VPC endpoints will be rejected when the resource was destroyed.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique ID in UUID format which equals to the ID of the VPC endpoint service.

* `connections` - An array of VPC endpoints connect to the VPC endpoint service. Structure is documented below.
  + `endpoint_id` - The unique ID of the VPC endpoint.
  + `packet_id` - The packet ID of the VPC endpoint.
  + `domain_id` - The user's domain ID.
  + `status` - The connection status of the VPC endpoint.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 3 minute.
