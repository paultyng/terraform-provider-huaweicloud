---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_port

Use this data source to get the details of a specific IEC subnet port.

## Example Usage

```hcl
variable "subent_id" {}

data "huaweicloud_iec_port" "port_1" {
  subnet_id = var.subent_id
  fixed_ip  = "192.168.1.123"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the port. If omitted, the provider-level region will be
  used.

* `id` - (Optional, String) The ID of the port.

* `subnet_id` - (Optional, String) The ID of the subnet which the port belongs to.

* `fixed_ip` - (Optional, String) The IP address of the port.

* `mac_address` - (Optional, String) The MAC address of the port.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.
* `status` - Indicates the status of the port.
* `site_id` - Indicates the ID of the IEC site.
* `security_groups` - Indicates the list of security group IDs applied on the port.
