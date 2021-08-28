---
subcategory: "MapReduce Service (MRS)"
---

# huaweicloud_mapreduce_cluster

Manages a cluster resource within HuaweiCloud MRS.

## Example Usage

### Create an analysis cluster

```hcl
data "huaweicloud_availability_zones" "test" {}

variable "cluster_name" {}
variable "password" {}
variable "vpc_id" {}
variable "subnet_id" {}

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = var.cluster_name
  version            = "MRS 1.9.2"
  type               = "ANALYSIS"
  component_list     = [
    "Hadoop",
    "Hive",
    "Tez"]
  manager_admin_pass = var.password
  node_admin_pass    = var.password
  vpc_id             = var.vpc_id
  subnet_id          = var.subnet_id

  master_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  analysis_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  analysis_task_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 1
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }

  tags = {
    foo = "bar"
    key = "value"
  }
}
```

### Create a stream cluster

```hcl
data "huaweicloud_availability_zones" "test" {}

variable "cluster_name" {}
variable "password" {}
variable "vpc_id" {}
variable "subnet_id" {}

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = var.cluster_name
  type               = "STREAMING"
  version            = "MRS 1.9.2"
  manager_admin_pass = var.password
  node_admin_pass    = var.password
  vpc_id             = var.vpc_id
  subnet_id          = var.subnet_id
  component_list     = [
    "Storm"]

  master_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }

  tags = {
    foo = "bar"
    key = "value"
  }
}
```

### Create a hybrid cluster

```hcl
data "huaweicloud_availability_zones" "test" {}

variable "cluster_name" {}
variable "password" {}
variable "vpc_id" {}
variable "subnet_id" {}

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = var.cluster_name
  version            = "MRS 1.9.2"
  type               = "MIXED"
  component_list     = [
    "Hadoop",
    "Spark",
    "Hive",
    "Tez",
    "Storm"]
  manager_admin_pass = var.password
  node_admin_pass    = var.password
  vpc_id             = var.vpc_id
  subnet_id          = var.subnet_id

  master_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  analysis_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  analysis_task_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 1
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_task_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 1
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }

  tags = {
    foo = "bar"
    key = "value"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the MapReduce cluster resource. If omitted, the
  provider-level region will be used. Changing this will create a new MapReduce cluster resource.

* `availability_zone` - (Required, String, ForceNew) Specifies the availability zone in which to create the cluster.
  Please following [reference](https://developer.huaweicloud.com/intl/en-us/endpoint?all)
  Changing this will create a new MapReduce cluster resource.

* `name` - (Required, String, ForceNew) Specifies the name of the MapReduce cluster. The name can contain 2 to 64
  characters, which may consist of letters, digits, underscores (_) and hyphens (-). Changing this will create a new
  MapReduce cluster resource.

* `version` - (Required, String, ForceNew) Specifies the MapReduce cluster version. The valid values are `MRS 1.9.2`
  , `MRS 3.0.5` and `MRS 3.1.0`. Changing this will create a new MapReduce cluster resource.

* `component_list` - (Required, List, ForceNew) Specifies the list of component names. For the components supported by
  the cluster, please following [reference](https://support.huaweicloud.com/intl/en-us/productdesc-mrs/mrs_08_0005.html)
  Changing this will create a new MapReduce cluster resource.

* `master_nodes` - (Required, List, ForceNew) Specifies a list of the informations about the master nodes in the
  MapReduce cluster. The `nodes` object structure of the `master_nodes` is documented below. Changing this will create a
  new MapReduce cluster resource.

* `vpc_id` - (Required, String, ForceNew) Specifies the ID of the VPC which bound to the MapReduce cluster. Changing
  this will create a new MapReduce cluster resource.

* `subnet_id` - (Required, String, ForceNew) Specifies the network ID of a subnet which bound to the MapReduce cluster.
  Changing this will create a new MapReduce cluster resource.

* `type` - (Optional, String, ForceNew) Specifies the type of the MapReduce cluster. The valid values are *ANALYSIS*, *
  STREAMING* and *MIXED*, default to *ANALYSIS*. Changing this will create a new MapReduce cluster resource.

* `enterprise_project_id` - (Optional, String, ForceNew) Specifies a unique ID in UUID format of enterprise project.
  Changing this will create a new MapReduce cluster resource.

* `eip_id` - (Optional, String, ForceNew) Specifies the EIP ID which bound to the MapReduce cluster. Changing this will
  create a new MapReduce cluster resource.

* `log_collection` - (Optional, Bool, ForceNew) Specifies whether logs are collected when cluster installation fails.
  Default to true. If `log_collection` set true, the OBS buckets will be created and only used to collect logs that
  record MapReduce cluster creation failures. Changing this will create a new MapReduce cluster resource.

* `manager_admin_pass` - (Optional, String, ForceNew) Specifies the administrator password, which is used to log in to
  the cluster management page. The password can contain 8 to 26 charactors and cannot be the username or the username
  spelled backwards. The password must contain lowercase letters, uppercase letters, digits, spaces and the special
  characters: `!?,.:-_{}[]@$^+=/`. Changing this will create a new MapReduce cluster resource.

* `node_admin_pass` - (Optional, String, ForceNew) Specifies the administrator password, which is used to log in to the
  each nodes(/ECSs). The password can contain 8 to 26 charactors and cannot be the username or the username spelled
  backwards. The password must contain lowercase letters, uppercase letters, digits, spaces and the special
  characters: `!?,.:-_{}[]@$^+=/`. Changing this will create a new MapReduce cluster resource. This parameter
  and `node_key_pair` are alternative.

* `node_key_pair` - (Optional, String, ForceNew) Specifies the name of a key pair, which is used to log in to the each
  nodes(/ECSs). Changing this will create a new MapReduce cluster resource.

* `safe_mode` - (Optional, Bool, ForceNew) Specifies whether the running mode of the MapReduce cluster is secure,
  default to true.
  + true: enable Kerberos authentication.
  + false: disable Kerberos authentication. Changing this will create a new MapReduce cluster resource.

* `security_group_ids` - (Optional, List, ForceNew) Specifies an array of one or more security group ID to attach to the
  MapReduce cluster. If using the specified security group, the group need to open the specified port (9022) rules.

* `analysis_core_nodes` - (Optional, List) Specifies a list of the informations about the analysis core nodes in the
  MapReduce cluster. The `nodes` object structure of the `analysis_core_nodes` is documented below.

* `streaming_core_nodes` - (Optional, List) Specifies a list of the informations about the streaming core nodes in the
  MapReduce cluster. The `nodes` object structure of the `streaming_core_nodes` is documented below.

* `analysis_task_nodes` - (Optional, List) Specifies a list of the informations about the analysis task nodes in the
  MapReduce cluster. The `nodes` object structure of the `analysis_task_nodes` is documented below.

* `streaming_task_nodes` - (Optional, List) Specifies a list of the informations about the streaming task nodes in the
  MapReduce cluster. The `nodes` object structure of the `streaming_task_nodes` is documented below.

* `tags` - (Optional, Map, ForceNew) Specifies the key/value pairs to associate with the cluster.

The `nodes` block supports:

* `flavor` - (Required, String, ForceNew) Specifies the instance specifications for each nodes in node group. Changing
  this will create a new MapReduce cluster resource.

* `node_number` - (Required, Int) Specifies the number of nodes for the node group.

  -> **NOTE:** Only the core group and task group updations are allowed. The number of nodes after scaling cannot be
  less than the number of nodes originally created.

* `root_volume_type` - (Required, String, ForceNew) Specifies the system disk flavor of the nodes. Changing this will
  create a new MapReduce cluster resource.

* `root_volume_size` - (Required, Int, ForceNew) Specifies the system disk size of the nodes. Changing this will create
  a new MapReduce cluster resource.

* `data_volume_type` - (Optional, String, ForceNew) Specifies the data disk flavor of the nodes. Required
  if `data_volume_count` is not empty. Changing this will create a new MapReduce cluster resource.

* `data_volume_size` - (Optional, Int, ForceNew) Specifies the data disk size of the nodes. Required
  if `data_volume_count` is not empty. Changing this will create a new MapReduce cluster resource.

* `data_volume_count` - (Optional, Int, ForceNew) Specifies the data disk number of the nodes. The number configuration
  of each node are as follows:
  + master_nodes: 1.
  + analysis_core_nodes: minimum is one and the maximum is subject to the configuration of the corresponding flavor.
  + streaming_core_nodes: minimum is one and the maximum is subject to the configuration of the corresponding flavor.
  + analysis_task_nodes: minimum is zero and the maximum is subject to the configuration of the corresponding flavor.
  + streaming_task_nodes: minimum is zero and the maximum is subject to the configuration of the corresponding flavor.

  Changing this will create a new MapReduce cluster resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The cluster ID in UUID format.
* `total_node_number` - The total number of nodes deployed in the cluster.
* `master_node_ip` - The IP address of the master node.
* `private_ip` - The preferred private IP address of the master node.
* `status` - The cluster state, which include: running, frozen, abnormal and failed.
* `create_time` - The cluster creation time, in RFC-3339 format.
* `update_time` - The cluster update time, in RFC-3339 format.
* `charging_start_time` - The charging start time which is the start time of billing, in RFC-3339 format.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 60 minute.
* `update` - Default is 180 minute.
* `delete` - Default is 40 minute.

## Import

Clusters can be imported by their `id`. For example,

```
terraform import huaweicloud_mapreduce_cluster.test b11b407c-e604-4e8d-8bc4-92398320b847
```

Note that the imported state may not be identical to your resource definition, due to some attrubutes missing from the
API response, security or some other reason. The missing attributes include:
`manager_admin_pass`, `node_admin_pass` and `eip_id`. It is generally recommended running `terraform plan` after
importing a cluster. You can then decide if changes should be applied to the cluster, or the resource definition should
be updated to align with the cluster. Also you can ignore changes as below.

```
resource "huaweicloud_mapreduce_cluster" "test" {
    ...

  lifecycle {
    ignore_changes = [
      manager_admin_pass, node_admin_pass, eip_id,
    ]
  }
}
```
