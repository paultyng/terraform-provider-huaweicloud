---
subcategory: "Cloud Container Engine (CCE)"
---

# huaweicloud_cce_node_pool
Add a node pool to a container cluster.

## Example Usage

```hcl
variable "cluster_id" { }
variable "key_pair" { }
variable "availability_zone" { }

resource "huaweicloud_cce_node_pool" "node_pool" {
  cluster_id               = var.cluster_id
  name                     = "testpool"
  os                       = "EulerOS 2.5"
  initial_node_count       = 2
  flavor_id                = "s3.large.4"
  availability_zone        = var.availability_zone
  key_pair                 = var.keypair
  scall_enable             = true
  min_node_count           = 1
  max_node_count           = 10
  scale_down_cooldown_time = 100
  priority                 = 1
  type                     = "vm"

  root_volume {
    size       = 40
    volumetype = "SAS"
  }
  data_volumes {
    size       = 100
    volumetype = "SAS"
  }
}
```

## Argument Reference
The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the cce pool resource. If omitted, the provider-level region will be used. Changing this creates a new cce node pool resource.

* `cluster_id` - (Required, String, ForceNew) ID of the cluster. Changing this parameter will create a new resource.

* `name` - (Required, String) Node Pool Name.

* `initial_node_count` - (Required, Int) Initial number of expected nodes in the node pool.

* `flavor_id` - (Required, String, ForceNew) Specifies the flavor id. Changing this parameter will create a new resource.

*  `type` - (Optional, String, ForceNew) Node Pool type. Possible values are: "vm" and "ElasticBMS".
 
* `availability_zone` - (Optional, String, ForceNew) specify the name of the available partition (AZ). Default value is random
    to create nodes in a random AZ in the node pool.
    Changing this parameter will create a new resource.

* `os` - (Optional, String) Operating System of the node. The value can be EulerOS 2.5 and CentOS 7.6.
    Changing this parameter will create a new resource.

* `key_pair` - (Optional, String, ForceNew) Key pair name when logging in to select the key pair mode. This parameter and `password` are alternative.
    Changing this parameter will create a new resource.

* `password` - (Optional, String, ForceNew) root password when logging in to select the password mode. This parameter must be salted and alternative to `key_pair`.
    Changing this parameter will create a new resource.

* `subnet_id` - (Optional, String, ForceNew) The ID of the subnet to which the NIC belongs. Changing this parameter will create a new resource.

* `max_pods` - (Optional, Int, ForceNew) The maximum number of instances a node is allowed to create.
    Changing this parameter will create a new resource.

* `preinstall` - (Optional, String, ForceNew) Script required before installation. The input value can be a Base64 encoded string or not.
    Changing this parameter will create a new resource.

* `postinstall` - (Optional, String, ForceNew) Script required after the installation. The input value can be a Base64 encoded string or not.
    Changing this parameter will create a new resource.

* `extend_param` - (Optional, Map, ForceNew) Extended parameter. Changing this parameter will create a new resource. Availiable keys :

  + `agency_name` - Specifies the agency name to provide temporary credentials for CCE node to access other cloud services.
  + `alpha.cce/NodeImageID` - This parameter is required when a custom image is used to create a BMS node.
  + `dockerBaseSize` - The available disk space of a single docker container on the node in device mapper mode.
  + `DockerLVMConfigOverride` - Docker data disk configurations. The following is an example default configuration:

```hcl
  extend_param = {
    DockerLVMConfigOverride = "dockerThinpool=vgpaas/90%VG;kubernetesLV=vgpaas/10%VG;diskType=evs;lvType=linear"
  }
```

* `scall_enable` - (Optional, Bool) Whether to enable auto scaling. If Autoscaler is enabled, install the autoscaler add-on to use the auto scaling feature.

* `min_node_count` - (Optional, Int) Minimum number of nodes allowed if auto scaling is enabled.

* `max_node_count` - (Optional, Int) Maximum number of nodes allowed if auto scaling is enabled.

* `scale_down_cooldown_time` - (Optional, Int) Interval between two scaling operations, in minutes.

* `priority` - (Optional, Int) Weight of a node pool. A node pool with a higher weight has a higher priority during scaling.

* `labels` - (Optional, Map) Tags of a Kubernetes node, key/value pair format.

* `tags` - (Optional, Map) Tags of a VM node, key/value pair format.

* `root_volume` - (Required, List, ForceNew) It corresponds to the system disk related configuration.
  The structure is described below. Changing this parameter will create a new resource.

* `data_volumes` - (Required, List, ForceNew) Represents the data disk to be created.
  The structure is described below. Changing this parameter will create a new resource.

* `taints` - (Optional, List) You can add taints to created nodes to configure anti-affinity.
  The structure is described below.

The `root_volume` block supports:

* `size` - (Required, Int) Disk size in GB.

* `volumetype` - (Required, String) Disk type.

* `extend_params` - (Optional, Map) Disk expansion parameters.

The `data_volumes` block supports:

* `size` - (Required, Int) Disk size in GB.

* `volumetype` - (Required, String) Disk type.

* `extend_params` - (Optional, Map) Disk expansion parameters.

The `taints` block supports:

* `key` - (Required, String, ForceNew) A key must contain 1 to 63 characters starting with a letter or digit. Only letters, digits, hyphens (-),
  underscores (_), and periods (.) are allowed. A DNS subdomain name can be used as the prefix of a key.

* `value` - (Required, String) A value must start with a letter or digit and can contain a maximum of 63 characters, including letters,
  digits, hyphens (-), underscores (_), and periods (.).

* `effect` - (Required, String) Available options are NoSchedule, PreferNoSchedule, and NoExecute.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

* `status` - Node status information.

* `billing_mode` - Billing mode of a node.

## Timeouts
This resource provides the following timeouts configuration options:

* `create` - Default is 20 minute.
* `delete` - Default is 20 minute.

## Import

CCE node pool can be imported using the cluster ID and node pool ID
separated by a slash, e.g.:

```
$ terraform import huaweicloud_cce_node_pool.my_node_pool 5c20fdad-7288-11eb-b817-0255ac10158b/e9287dff-7288-11eb-b817-0255ac10158b
```

Note that the imported state may not be identical to your resource definition, due to some attrubutes missing from the
API response, security or some other reason. The missing attributes include:
`password`, `subnet_id`, `preinstall`, `posteinstall`, `taints`.
It is generally recommended running `terraform plan` after importing a node pool. 
You can then decide if changes should be applied to the node pool, or the resource definition should be updated to align
with the node pool. Also you can ignore changes as below.

```
resource "huaweicloud_cce_node_pool" "my_node_pool" {
    ...

  lifecycle {
    ignore_changes = [
      password, subnet_id,
    ]
  }
}
```
