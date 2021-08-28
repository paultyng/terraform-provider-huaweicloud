---
subcategory: "Cloud Container Engine (CCE)"
---

# huaweicloud_cce_cluster

Provides a CCE cluster resource. This is an alternative to `huaweicloud_cce_cluster_v3`

## Basic Usage

```hcl
resource "huaweicloud_vpc" "myvpc" {
  name = "vpc"
  cidr = "192.168.0.0/16"
}

resource "huaweicloud_vpc_subnet" "mysubnet" {
  name       = "subnet"
  cidr       = "192.168.0.0/16"
  gateway_ip = "192.168.0.1"

  //dns is required for cce node installing
  primary_dns   = "100.125.1.250"
  secondary_dns = "100.125.21.250"
  vpc_id        = huaweicloud_vpc.myvpc.id
}

resource "huaweicloud_cce_cluster" "cluster" {
  name                   = "cluster"
  flavor_id              = "cce.s1.small"
  vpc_id                 = huaweicloud_vpc.myvpc.id
  subnet_id              = huaweicloud_vpc_subnet.mysubnet.id
  container_network_type = "overlay_l2"
}
```

## Cluster With Eip

```hcl
resource "huaweicloud_vpc" "myvpc" {
  name = "vpc"
  cidr = "192.168.0.0/16"
}

resource "huaweicloud_vpc_subnet" "mysubnet" {
  name       = "subnet"
  cidr       = "192.168.0.0/16"
  gateway_ip = "192.168.0.1"

  //dns is required for cce node installing
  primary_dns   = "100.125.1.250"
  secondary_dns = "100.125.21.250"
  vpc_id        = huaweicloud_vpc.myvpc.id
}

resource "huaweicloud_vpc_eip" "myeip" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name        = "test"
    size        = 8
    share_type  = "PER"
    charge_mode = "traffic"
  }
}

resource "huaweicloud_cce_cluster" "cluster" {
  name                   = "cluster"
  cluster_type           = "VirtualMachine"
  flavor_id              = "cce.s1.small"
  vpc_id                 = huaweicloud_vpc.myvpc.id
  subnet_id              = huaweicloud_vpc_subnet.mysubnet.id
  container_network_type = "overlay_l2"
  authentication_mode    = "rbac"
  eip                    = huaweicloud_vpc_eip.myeip.address
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the cce cluster resource. If omitted, the
  provider-level region will be used. Changing this creates a new cce cluster resource.

* `name` - (Required, String, ForceNew) Cluster name. Changing this parameter will create a new cluster resource.

* `flavor_id` - (Required, String, ForceNew) Cluster specifications. Changing this parameter will create a new cluster
  resource. Possible values:

  + `cce.s1.small` - small-scale single cluster (up to 50 nodes).
  + `cce.s1.medium` - medium-scale single cluster (up to 200 nodes).
  + `cce.s2.small` - small-scale HA cluster (up to 50 nodes).
  + `cce.s2.medium` - medium-scale HA cluster (up to 200 nodes).
  + `cce.s2.large` - large-scale HA cluster (up to 1000 nodes).
  + `cce.s2.xlarge` - large-scale HA cluster (up to 2000 nodes).

* `cluster_version` - (Optional, String, ForceNew) For the cluster version, defaults to the latest supported version.
  Changing this parameter will create a new cluster resource.

* `cluster_type` - (Optional, String, ForceNew) Cluster Type, possible values are VirtualMachine and ARM64. Defaults
  to *VirtualMachine*. Changing this parameter will create a new cluster resource.

* `description` - (Optional, String) The Cluster description.

* `vpc_id` - (Required, String, ForceNew) The ID of the VPC used to create the node. Changing this parameter will create
  a new cluster resource.

* `subnet_id` - (Required, String, ForceNew) The ID of the subnet used to create the node which should be configured
  with a *DNS address*. Changing this parameter will create a new cluster resource.

* `container_network_type` - (Required, String, ForceNew) Container network parameters. Possible values:

  + `overlay_l2` - An overlay_l2 network built for containers by using Open vSwitch(OVS).
  + `vpc-router` - An vpc-router network built for containers by using ipvlan and custom VPC routes.
  + `eni` - A Yangtse network built for cce turbo cluster. The container network deeply integrates the native ENI
      capability of VPC, uses the VPC CIDR block to allocate container addresses, and supports direct connections
      between ELB and containers to provide high performance.

* `container_network_cidr` - (Optional, String, ForceNew) Container network segment. Changing this parameter will create
  a new cluster resource.

* `service_network_cidr` - (Optional, String, ForceNew) Service network segment. Changing this parameter will create a
  new cluster resource.

* `eni_subnet_id` - (Optional, String, ForceNew) ENI subnet id. Specified when creating a CCE Turbo cluster. Changing
  this parameter will create a new cluster resource.

* `eni_subnet_cidr` - (Optional, String, ForceNew) ENI network segment. Specified when creating a CCE Turbo cluster.
  Changing this parameter will create a new cluster resource.

* `authentication_mode` - (Optional, String, ForceNew) Authentication mode of the cluster, possible values are
  authenticating_proxy and rbac. Defaults to *rbac*. Changing this parameter will create a new cluster resource.

* `authenticating_proxy_ca` - (Optional, String, ForceNew) CA root certificate provided in the authenticating_proxy
  mode. The CA root certificate is encoded to the Base64 format. Changing this parameter will create a new cluster
  resource.

* `multi_az` - (Optional, Bool, ForceNew) Enable multiple AZs for the cluster, only when using HA flavors. Changing this
  parameter will create a new cluster resource. This parameter and `masters` are alternative

* `masters` - (Optional, List, ForceNew) Advanced configuration of master nodes. Changing this creates a new cluster.
  This parameter and `multi_az` are alternative.

* `eip` - (Optional, String, ForceNew) EIP address of the cluster. Changing this parameter will create a new cluster
  resource.

* `kube_proxy_mode` - (Optional, String, ForceNew) Service forwarding mode. Two modes are available:

  + iptables: Traditional kube-proxy uses iptables rules to implement service load balancing. In this mode, too many
      iptables rules will be generated when many services are deployed. In addition, non-incremental updates will cause
      a latency and even obvious performance issues in the case of heavy service traffic.
  + ipvs: Optimized kube-proxy mode with higher throughput and faster speed. This mode supports incremental updates
      and can keep connections uninterrupted during service updates. It is suitable for large-sized clusters.

* `extend_param` - (Optional, Map, ForceNew) Extended parameter. Changing this parameter will create a new cluster
  resource.

* `charging_mode` - (Optional, String, ForceNew) Specifies the charging mode of the CCE cluster. Valid values are *
  prePaid* and *postPaid*, defaults to *postPaid*. Changing this creates a new cluster.

* `period_unit` - (Optional, String, ForceNew) Specifies the charging period unit of the CCE cluster. Valid values are *
  month* and *year*. This parameter is mandatory if `charging_mode` is set to *prePaid*. Changing this creates a new
  cluster.

* `period` - (Optional, Int, ForceNew) Specifies the charging period of the CCE cluster. If `period_unit` is set to *
  month*, the value ranges from 1 to 9. If `period_unit` is set to *year*, the value ranges from 1 to 3. This parameter
  is mandatory if `charging_mode` is set to *prePaid*. Changing this creates a new cluster.

* `auto_renew` - (Optional, String, ForceNew) Specifies whether auto renew is enabled. Valid values are "true" and "
  false". Changing this creates a new cluster.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id of the cce cluster. Changing this
  creates a new cluster.

* `delete_evs` - (Optional, String) Specified whether to delete associated EVS disks when deleting the CCE cluster.
  valid values are "true", "try" and "false". Default is false.

* `delete_obs` - (Optional, String) Specified whether to delete associated OBS buckets when deleting the CCE cluster.
  valid values are "true", "try" and "false". Default is false.

* `delete_sfs` - (Optional, String) Specified whether to delete associated SFS file systems when deleting the CCE
  cluster. valid values are "true", "try" and "false". Default is false.

* `delete_efs` - (Optional, String) Specified whether to unbind associated SFS Turbo file systems when deleting the CCE
  cluster. valid values are "true", "try" and "false". Default is false.

* `delete_all` - (Optional, String) Specified whether to delete all associated storage resources when deleting the CCE
  cluster. valid values are "true", "try" and "false". Default is false.

* `hibernate` - (Optional, Bool) Specifies whether to hibernate the CCE cluster. Defaults to false. After a cluster is
  hibernated, resources such as workloads cannot be created or managed in the cluster, and the cluster cannot be
  deleted.

The `masters` block supports:

* `availability_zone` - (Optional, String, ForceNew) Specifies the availability zone of the master node. Changing this
  creates a new cluster.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Id of the cluster resource.

* `status` - Cluster status information.

* `certificate_clusters/name` - The cluster name.

* `certificate_clusters/server` - The server IP address.

* `certificate_clusters/certificate_authority_data` - The certificate data.

* `certificate_users/name` - The user name.

* `certificate_users/client_certificate_data` - The client certificate data.

* `certificate_users/client_key_data` - The client key data.

* `security_group_id` - Security group ID of the cluster.

* `kube_config_raw` - Raw Kubernetes config to be used by kubectl and other compatible tools.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 30 minute.
* `update` - Default is 30 minute.
* `delete` - Default is 30 minute.

## Import

Cluster can be imported using the cluster id, e.g.

 ```
 $ terraform import huaweicloud_cce_cluster.cluster_1 4779ab1c-7c1a-44b1-a02e-93dfc361b32d
```

Note that the imported state may not be identical to your resource definition, due to some attrubutes missing from the
API response, security or some other reason. The missing attributes include:
`delete_efs`, `delete_eni`, `delete_evs`, `delete_net`, `delete_obs`, `delete_sfs` and `delete_all`. It is generally
recommended running `terraform plan` after importing an cce cluster. You can then decide if changes should be applied to
the cluster, or the resource definition should be updated to align with the cluster. Also you can ignore changes as
below.

```
resource "huaweicloud_cce_cluster" "cluster_1" {
    ...

  lifecycle {
    ignore_changes = [
      delete_efs, delete_obs,
    ]
  }
}
```
