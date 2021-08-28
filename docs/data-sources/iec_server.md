---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_server

Use this data source to get the details of a specified IEC server.

## Example Usage

```hcl
variable "server_name" {}

data "huaweicloud_iec_server" "demo" {
  name = var.server_name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Specifies the IEC server name, which can be queried with a regular expression.

* `status` - (Optional, String) Specifies the status of IEC server.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The IEC server ID in UUID format.
* `edgecloud_id` - The ID of the edgecloud service.
* `edgecloud_name` - The Name of the edgecloud service.
* `coverage_sites` - An array of site ID and operator for the IEC server. The object structure is documented below.
* `flavor_id` - The flavor ID of the IEC server.
* `flavor_name` - The flavor name of the IEC server.
* `image_name` - The image name of the IEC server.
* `vpc_id` - The ID of vpc for the IEC server.
* `security_groups` - An array of one or more security group IDs to associate with the IEC server.
* `nics` - An array of one or more networks to attach to the IEC server. The object structure is documented below.
* `volume_attached` - An array of one or more disks to attach to the IEC server. The object structure is documented
  below.
* `public_ip` - The EIP address that is associted to the IEC server.
* `system_disk_id` - The system disk voume ID.

The `coverage_sites` block supports:

* `site_id` - The ID of IEC site.
* `site_info` - The located information of the IEC site. It contains area, province and city.
* `operator` - The operator of the IEC site.

The `nics` block supports:

* `port` - The port ID corresponding to the IP address on that network.
* `mac` - The MAC address of the NIC on that network.
* `address` - The IPv4 address of the server on that network.

The `volume_attached` block supports:

* `volume_id` - The volume ID on that attachment.
* `boot_index` - The volume boot index on that attachment.
* `size` - The volume size on that attachment.
* `type` - The volume type on that attachment.
* `device` - The device name in the IEC server.
