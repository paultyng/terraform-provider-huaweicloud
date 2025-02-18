---
subcategory: "Cloud Container Engine (CCE)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_cce_node_pool"
description: ""
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

* `region` - (Optional, String) Specifies the region in which to obtain the CCE node pools.
  If omitted, the provider-level region will be used.

* `cluster_id` - (Required, String) Specifies the ID of container cluster.

* `name` - (Optional, String) Specifies the name of the node pool.

* `node_pool_id` - (Optional, String) Specifies the ID of the node pool.

* `status` - (Optional, String) Specifies the state of the node pool.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID in UUID format.

* `initial_node_count` - Initial number of nodes in the node pool.

* `current_node_count` - Current number of nodes in the node pool.

* `flavor_id` - The flavor ID.

* `type` - Node Pool type.

* `availability_zone` - The name of the available partition (AZ).

* `os` - Operating System of the node.

* `key_pair` - Key pair name when logging in to select the key pair mode.

* `subnet_id` - The ID of the subnet to which the NIC belongs.

* `max_pods` - The maximum number of instances a node is allowed to create.

* `extend_param` - Extended parameter.

* `scale_enable` - Whether auto scaling is enabled.

* `min_node_count` - Minimum number of nodes allowed if auto scaling is enabled.

* `max_node_count` - Maximum number of nodes allowed if auto scaling is enabled.

* `scale_down_cooldown_time` - Interval between two scaling operations, in minutes.

* `priority` - Weight of a node pool. A node pool with a higher weight has a higher priority during scaling.

* `labels` - Tags of a Kubernetes node, key/value pair format.

* `tags` - Tags of a VM node, key/value pair format.

* `root_volume` - It corresponds to the system disk related configuration. Structure is documented below.

* `data_volumes` - Represents the data disk to be created. Structure is documented below.

* `enterprise_project_id` - The enterprise project ID of the node pool.

* `hostname_config` - The hostname config of the kubernetes node.
  The [object](#hostname_config) structure is documented below.

The `root_volume` and `data_volumes` blocks support:

* `size` - Disk size in GB.

* `volumetype` - Disk type.

* `extend_params` - Disk expansion parameters.

<a name="hostname_config"></a>
The `hostname_config` block supports:

* `type` - The hostname type of the kubernetes node.
