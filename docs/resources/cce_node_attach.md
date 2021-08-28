---
subcategory: "Cloud Container Engine (CCE)"
---

# huaweicloud_cce_node_attach

Add a node from an existing ecs server to a CCE cluster.

## Basic Usage

```hcl
resource "huaweicloud_cce_node_attach" "test" {
  cluster_id = var.cluster_id
  server_id  = var.server_id
  key_pair   = var.keypair_name
  os         = "EulerOS 2.5"

  tags = {
    foo = "bar"
    key = "value"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the cce node attach resource. If omitted, the
  provider-level region will be used. Changing this creates a new cce node attach resource.

* `cluster_id` - (Required, String, ForceNew) Specifies the ID of the cluster. Changing this parameter will create a new
  resource.

* `name` - (Optional, String) Specifies the Node Name.

* `server_id` - (Required, String, ForceNew) Specifies the ecs server id. Changing this parameter will create a new
  resource.

* `os` - (Required, String, ForceNew) Specifies the operating System of the node. Changing this parameter will create a
  new resource.
  + For VM nodes, clusters of v1.13 and later support *EulerOS 2.5* and *CentOS 7.6*.

* `key_pair` - (Optional, String, ForceNew) Specifies the key pair name when logging in to select the key pair mode.
  This parameter and `password` are alternative. Changing this parameter will create a new resource.

* `password` - (Optional, String, ForceNew) Specifies the root password when logging in to select the password mode.
  This parameter must be salted and alternative to `key_pair`. Changing this parameter will create a new resource.

* `max_pods` - (Optional, Int, ForceNew) Specifies the the maximum number of instances a node is allowed to create.
  Changing this parameter will create a new resource.

* `docker_base_size` - (Optional, Int, ForceNew) Specifies the available disk space of a single docker container on the
  node in device mapper mode. Changing this parameter will create a new resource.

* `lvm_config` - (Optional, String, ForceNew) Specifies the docker data disk configurations. The following is an
  example:

```hcl
  lvm_config = "dockerThinpool=vgpaas/90%VG;kubernetesLV=vgpaas/10%VG"
```

Changing this parameter will create a new resource.

* `preinstall` - (Optional, String, ForceNew) Specifies the script required before installation. The input value can be
  a Base64 encoded string or not. Changing this parameter will create a new resource.

* `postinstall` - (Optional, String, ForceNew) Specifies the script required after installation. The input value can be
  a Base64 encoded string or not. Changing this parameter will create a new resource.

* `tags` - (Optional, Map) Specifies the tags of a VM node, key/value pair format.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.
* `status` - Node status information.
* `private_ip` - Private IP of the CCE node.
* `public_ip` - Public IP of the CCE node.
* `flavor_id` - The flavor id of the CCE node.
* `availability_zone` - The name of the available partition (AZ).
* `root_volume` - The system disk related configuration.
* `data_volumes` - The data disks related configuration.
* `runtime` - The runtime of the CCE node.
* `ecs_group_id` - The Ecs group id.
* `subnet_id` - The ID of the subnet to which the NIC belongs.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 20 minute.
* `delete` - Default is 20 minute.
