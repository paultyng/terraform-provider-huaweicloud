---
subcategory: "Elastic Cloud Server (ECS)"
---

# huaweicloud_compute_instance

Use this data source to get the details of a specified compute instance.

## Example Usage

```hcl
variable "ecs_name" {}

data "huaweicloud_compute_instance" "demo" {
  name = var.ecs_name
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) The region in which to obtain the instance. If omitted, the provider-level region will
  be used.

* `name` - (Optional, String) Specifies the ECS name, which can be queried with a regular expression.

* `fixed_ip_v4` - (Optional, String)  Specifies the IPv4 addresses of the ECS.

* `flavor_id` - (Optional, String) Specifies the flavor ID.

* `enterprise_project_id` - (Optional, String) Specifies the enterprise project id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The instance ID in UUID format.
* `availability_zone` - The availability zone where the instance is located.
* `image_id` - The image ID of the instance.
* `image_name` - The image name of the instance.
* `flavor_name` - The flavor name of the instance.
* `key_pair` - The key pair that is used to authenticate the instance.
* `public_ip` - The EIP address that is associted to the instance.
* `system_disk_id` - The system disk voume ID.
* `user_data` - The user data (information after encoding) configured during instance creation.
* `security_group_ids` - An array of one or more security group IDs to associate with the instance.
* `network` - An array of one or more networks to attach to the instance. The network object structure is documented
  below.
* `volume_attached` - An array of one or more disks to attach to the instance. The volume_attached object structure is
  documented below.
* `scheduler_hints` - The scheduler with hints on how the instance should be launched. The available hints are described
  below.
* `tags` - The key/value pairs to associate with the instance.
* `status` - The status of the instance.

The `network` block supports:

* `uuid` - The network UUID to attach to the server.
* `port` - The port ID corresponding to the IP address on that network.
* `mac` - The MAC address of the NIC on that network.
* `fixed_ip_v4` - The fixed IPv4 address of the instance on this network.
* `fixed_ip_v6` - The Fixed IPv6 address of the instance on that network.

The `volume_attached` block supports:

* `volume_id` - The volume id on that attachment.
* `boot_index` - The volume boot index on that attachment.
* `size` - The volume size on that attachment.
* `type` - The volume type on that attachment.
* `pci_address` - The volume pci address on that attachment.

The `scheduler_hints` block supports:

* `group` - The UUID of a Server Group where the instance will be placed into.
