---
subcategory: "Cloud Container Engine (CCE)"
---

# huaweicloud_cce_node

To get the specified node in a cluster. This is an alternative to `huaweicloud_cce_node_v3`

## Example Usage

```hcl
variable "cluster_id" {}
variable "node_name" {}

data "huaweicloud_cce_node" "node" {
  cluster_id = var.cluster_id
  name       = var.node_name
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) The region in which to obtain the cce nodes. If omitted, the provider-level region will
  be used.

* `Cluster_id` - (Required, String) The id of container cluster.

* `name` - (Optional, String) Name of the node.

* `node_id` - (Optional, String) The id of the node.

* `status` - (Optional, String) The state of the node.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `flavor_id` - The flavor id to be used.

* `availability_zone` - Available partitions where the node is located.

* `os` - Operating System of the node.

* `subnet_id` - The ID of the subnet which the NIC belongs to.

* `esc_group_id` - The ID of Ecs group which the node belongs to.

* `tags` - Tags of a VM node, key/value pair format.

* `key_pair` - Key pair name when logging in to select the key pair mode.

* `billing_mode` - Node's billing mode: The value is 0 (on demand).

* `server_id` - The node's virtual machine ID in ECS.

* `public_ip` - Elastic IP parameters of the node.

* `private_ip` - Private IP of the node

* `root_volume` - It corresponds to the system disk related configuration.

  + `size` - Disk size in GB.
  + `volumetype` - Disk type.
  + `extend_params` - Disk expansion parameters.

* `data_volumes` - Represents the data disk to be created.

  + `size` - Disk size in GB.
  + `volumetype` - Disk type.
  + `extend_params` - Disk expansion parameters.

