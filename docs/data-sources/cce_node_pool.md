---
subcategory: "Cloud Container Engine (CCE)"
---

# huaweicloud_cce_node_pool

To get the specified node pool in a cluster.

## Example Usage

```hcl
variable "cluster_id" {}
variable "node_pool_name" {}

data "huaweicloud_cce_node_pool" "node_pool" {
  cluster_id = var.cluster_id
  name       = var.node_pool_name
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) The region in which to obtain the cce node pools. If omitted, the provider-level region
  will be used.

* `cluster_id` - (Required, String) Specifies the id of container cluster.

* `name` - (Optional, String) Specifies the name of the node pool.

* `node_pool_id` - (Optional, String) Specifies the id of the node pool.

* `status` - (Optional, String) Specifies the state of the node pool.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID in UUID format.

* `initial_node_count` - Initial number of nodes in the node pool.

* `current_node_count` - Current number of nodes in the node pool.

* `flavor_id` - The flavor id.

* `type` - Node Pool type.

* `availability_zone` - The name of the available partition (AZ).

* `os` - Operating System of the node.

* `key_pair` - Key pair name when logging in to select the key pair mode.

* `subnet_id` - The ID of the subnet to which the NIC belongs.

* `max_pods` - The maximum number of instances a node is allowed to create.

* `extend_param` - Extended parameter.

* `scall_enable` - Whether auto scaling is enabled.

* `min_node_count` - Minimum number of nodes allowed if auto scaling is enabled.

* `max_node_count` - Maximum number of nodes allowed if auto scaling is enabled.

* `scale_down_cooldown_time` - Interval between two scaling operations, in minutes.

* `priority` - Weight of a node pool. A node pool with a higher weight has a higher priority during scaling.

* `labels` - Tags of a Kubernetes node, key/value pair format.

* `tags` - Tags of a VM node, key/value pair format.

* `root_volume` - It corresponds to the system disk related configuration.

* `data_volumes` - Represents the data disk to be created.

The `root_volume` block supports:

* `size` - Disk size in GB.

* `volumetype` - Disk type.

* `extend_params` - Disk expansion parameters.

The `data_volumes` block supports:

* `size` - Disk size in GB.

* `volumetype` - Disk type.

* `extend_params` - Disk expansion parameters.
