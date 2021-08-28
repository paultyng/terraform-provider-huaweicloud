---
subcategory: "Blockchain Service (BCS)"
---

# huaweicloud_bcs_instance

## Example Usage

### Basic Instance

```hcl
variable "instance_name" {}

variable "instance_password" {}

variable "enterprise_project_id" {}

data "huaweicloud_availability_zones" "test" {}

data "huaweicloud_cce_cluster" "test" {
  ...
}

resource "huaweicloud_bcs_instance" "test" {
  name                  = var.instance_name
  cce_cluster_id        = data.huaweicloud_cce_cluster.test.id
  consensus             = "etcdraft"
  edition               = 1
  enterprise_project_id = var.enterprise_project_id
  fabric_version        = "2.0"
  password              = var.instance_password
  volume_type           = "nfs"
  org_disk_size         = 100
  security_mechanism    = "ECDSA"
  orderer_node_num      = 1
  delete_storage        = true

  peer_orgs {
    org_name = "organization01"
    count    = 2
  }
  channels {
    name      = "channel01"
    org_names = [
      "organization01",
    ]
  }
}
```

### Instance With kafka consensus strategy

```hcl
variable "instance_name" {}

variable "instance_password" {}

variable "enterprise_project_id" {}

variable "database_user_name" {}

variable "database_password" {}

data "huaweicloud_availability_zones" "test" {}

data "huaweicloud_cce_cluster" "test" {
  ...
}

resource "huaweicloud_bcs_instance" "test" {
  name                  = var.instance_name
  blockchain_type       = "private"
  cce_cluster_id        = data.huaweicloud_cce_cluster.test.id
  consensus             = "kafka"
  edition               = 4
  fabric_version        = "1.4"
  enterprise_project_id = var.enterprise_project_id
  password              = var.instance_password
  volume_type           = "efs"
  org_disk_size         = 500
  database_type         = "couchdb"
  orderer_node_num      = 2
  bandwidth_size        = 5
  delete_storage        = true
  delete_obs            = true

  couchdb {
    user_name = var.database_user_name
    password  = var.database_password
  }
  peer_orgs {
    org_name = "organization01"
    count    = 2
  }
  peer_orgs {
    org_name = "organization02"
    count    = 2
  }
  channels {
    name      = "channel01"
    org_names = [
      "organization02",
    ]
  }
  channels {
    name      = "channel02"
    org_names = [
      "organization01",
      "organization02",
    ]
  }
  sfs_turbo {
    share_type        = "STANDARD"
    type              = "efs-ha"
    flavor            = "sfs.turbo.standard"
    availability_zone = data.huaweicloud_availability_zones.test.names[0]
  }
  kafka {
    flavor            = "c3.mini"
    storage_size      = 600
    availability_zone = [
      data.huaweicloud_availability_zones.test.names[0],
      data.huaweicloud_availability_zones.test.names[1],
      data.huaweicloud_availability_zones.test.names[2],
    ]
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the instance. If omitted, the
  provider-level region will be used. Changing this will create a new instance.

* `name` - (Required, String, ForceNew) Specifies a unique name of the BCS instance. The name consists of 4 to 24
  characters, including letters, digits, chinese charactors and hyphens (-), and the name cannot start with a hyphen.
  Changing this will create a new instance.

* `edition` - (Required, Int, ForceNew) Specifies Service edition of the BCS instance. Valid values are `1`, `2` and `4`
  . Changing this will create a new instance.

* `fabric_version` - (Required, String, ForceNew) Specifies version of fabric for the BCS instance. Valid values
  are `1.4` and `2.0`
  Changing this will create a new instance.

* `consensus` - (Required, String, ForceNew) Specifies the consensus algorithm used by the BCS instance. The valid
  values of fabric 1.4 are `solo`, `kafka` and `SFLIC`, and the valid values of fabric 2.0 are `SFLIC`
  and `etcdraft`. Changing this will create a new instance.

* `orderer_node_num` - (Required, Int, ForceNew) Specifies the number of peers in the orderer organizaion. Changing this
  will create a new instance.

* `cce_cluster_id` - (Required, String, ForceNew) Specifies the ID of the CCE cluster to attach to the BCS instance. The
  BCS service needs to exclusively occupy the CCE cluster. Please make sure that the CCE cluster is not occupied before
  deploying the BCS service. Changing this will create a new instance.

* `enterprise_project_id` - (Required, String, ForceNew) Specifies the ID of the enterprise project that the BCS
  instance belong to. Changing this will create a new instance.

* `password` - (Required, String, ForceNew) Specifies the Resource access and blockchain management password. The
  password consists of 8 to 12 characters and must consist at least three of following: uppercase letters, lowercase
  letters, digits, chinese charactors, special charactors(!@$%^-_=+[{}]:,./?). Changing this will create a new instance.

* `volume_type` - (Required, String, ForceNew) Specifies the storage volume type to attach to each organization of the
  BCS instance. Valid values are `nfs` (SFS) and `efs` (SFS Turbo). Changing this will create a new instance.

* `org_disk_size` - (Required, Int, ForceNew) Specifies the storage capacity of peer organization. Changing this will
  create a new instance.
  + The minimum storage capacity of `efs` volume type is 500GB.

  The specifications are as follows when `volume_type` is `nfs`:
  + The minimum storage capacity of basic edition is 40 GB.
  + The minimum storage capacity of enterprise and professional edition is 100 GB.

* `block_info` - (Optional, List, ForceNew) Specifies the configuration of block generation. The block_info object
  structure is documented below.

* `blockchain_type` - (Optional, String, ForceNew) Specifies the blockchain type of the BCS instance. Valid values
  are `private` and  `union`. Default is `private`. Changing this will create a new instance.

* `channels` - (Optional, List, ForceNew) Specifies an array of one or more channels to attach to the BCS instance. If
  omitted, the bcs instance will create a `channels` named `channel` by default. Changing this will create a new
  instance. The channels object structure is documented below.

* `couchdb` - (Optional, List, ForceNew) Specifies the NoSQL database used by BCS instance. If omitted, the bcs instance
  will create a `goleveldb`(File Database) database by default. Changing this will create a new instance. The couchdb
  object structure is documented below.

* `delete_storage` - (Optional, Bool) Specified whether to delete the associated SFS resources when deleting BCS
  instance. Default is false.

* `delete_obs` - (Optional, Bool) Specified whether to delete the associated OBS bucket when deleting BCS instance.
  `delete_obs` is used to delete the OBS created by the BCS instance of the Kafka consensus strategy. Default is false.

* `eip_enable` - (Optional, Bool, ForceNew) Specifies whether to use the EIP of the CCE to bind the BCS instance.
  Changing this will create a new instance. Defalut is true.
  + `true` means an EIP bound to the cluster will be used as the blockchain network access address. If the cluster is
      not bound with any EIP, bind an EIP to the cluster first. Please make sure that the cluster is bound to EIP.
  + `false` means a private address of the cluster will be used ad the blockchain network access address to ensure
      that the application can communicate with the internal network of the cluster.

* `kafka` - (Optional, List, ForceNew) Specifies the kafka configuration for the BCS instance. Changing this will create
  a new instance. The kafka object structure is documented below.

* `peer_orgs` - (Optional, List, ForceNew) Specifies an array of one or more Peer organizations to attach to the BCS
  instance. Changing this will create a new instance. If omitted, the bcs instance will create a `peer_orgs`
  named `organization` by default and the node count is 2. The peer_orgs object structure is documented below.

* `restful_api_support` - (Optional, Bool, ForceNew) Specified whether to add RESTful API support. Changing this will
  create a new instance.

* `sfs_turbo` - (Optional, List, ForceNew) Specifies the information about the SFS Turbo file system. Changing this will
  create a new instance. The sfs_turbo object structure is documented below.

* `security_mechanism` - (Optional, String, ForceNew) Specifies the secutity mechanism used by the BCS instance. Valid
  values are `ECDSA` and `SM2`(Chinese cryptographic algorithms, The basic and professional don't support this
  algorithm). Default is `ECDSA`. Changing this will create a new instance.

* `tc3_need` - (Optional, Bool, ForceNew) Specified whether to add Trusted computing platform. Changing this will create
  a new instance.

The `peer_orgs` block supports:

* `org_name` - (Required, String, ForceNew) Specifies the name of the peer organization. Changing this creates a new
  instance.

* `count` - (Required, Int, ForceNew) Specifies the number of peers in organization. Changing this creates a new
  instance.

The `channels` block supports:

* `name` - (Required, String, ForceNew) Specifies the name of the channel. Changing this creates a new instance.

* `org_name` - (Optional, List, ForceNew) Specifies the name of the peer organization. Changing this creates a new
  instance.

The `couchdb` block supports:

* `user_name` - (Required, String, ForceNew) Specifies the user name of the couch datebase. Changing this creates a new
  instance.

* `password` - (Required, String, ForceNew) Specifies the password of the couch datebase. The password consists of 8 to
  26 characters and must consist at least three of following: uppercase letters, lowercase letters, digits, special
  charactors(!@$%^-_=+[{}]:,./?). Changing this creates a new instance.

The `sfs_turbo` block supports:

* `availability_zone` - (Optional, String, ForceNew) Specifies the availability zone in which to create the SFS turbo.
  Please following [reference](https://developer.huaweicloud.com/en-us/endpoint/?all) for the values. Changing this
  creates a new instance.

* `flavor` - (Optional, String, ForceNew) Specifies the flavor of SFS turbo. Changing this creates a new instance.

* `share_type` - (Optional, String, ForceNew) Specifies the share type of the SFS turbo. Changing this creates a new
  instance.

* `type` - (Optional, String, ForceNew) Specifies the type of SFS turbo. Changing this creates a new instance.

The `block_info` block supports:

* `transaction_quantity` - (Optional, Int, ForceNew) Specifies the number of transactions included in the block. The
  defalt value is 500. Changing this creates a new instance.

* `block_size` - (Optional, Int, ForceNew) Specifies the volume of the block, the unit is MB. The default value is 2.
  Changing this creates a new instance.

* `generation_interval` - (Optional, Int, ForceNew) Specifies the block generation time, the unit is second. The defalt
  value is 2. Changing this creates a new instance.

The `kafka` block supports:

* `availability_zone` - (Optional, List, ForceNew)  Specifies the availability zone in which to create the kafka. The
  list must contain one or more than three availability zone. Please
  following [reference](https://developer.huaweicloud.com/en-us/endpoint/?all) for the values. Changing this creates a
  new instance.

* `flavor` - (Optional, String, ForceNew) Specifies the kafka flavor type. Changing this creates a new instance.
  + `c3.mini` : Mini type, the reference bandwidth is 100MB/s.
  + `c3.small.2` : Small type, the reference bandwidth is 300MB/s.
  + `c3.middle.2` : Middle type, the reference bandwidth is 600MB/s.
  + `c3.high.2` : High type, the reference bandwidth is 1200MB/s.

* `storage_size` - (Optional, Int, ForceNew) Specifies the kafka storage capacity. The storage capacity must be an
  integral multiple of 100 and the maximum is 90000GB. Changing this creates a new instance.
  + The minimum storage capacity of mini type is 600GB.
  + The minimum storage capacity of small type is 1200GB.
  + The minimum storage capacity of middle type is 2400GB.
  + The minimum storage capacity of high type is 4800GB.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - A resource ID in UUID format.
* `status` - The status of the BCS instance.
* `version` - The service verison of the BCS instance.
* `purchase_type` - The deployment type of the BCS instance.
* `cross_region_support` - Whether the BCS instance is deployed across regions.
* `rollback_support` - Whether rollback is supported when the BCS service fails to br upgraded.
* `old_service_version` - The version of an old BCS service.
* `agent_portal_address` - The agent addresses and port numbers on the user data plane of the BCS service.
* `peer_orgs/pvc_name` - The name of the PersistenetVolumeClaim (PVC) used by the peer.
* `peer_orgs/status` - The peer status. The value contains `IsCreating`, `IsUpgrading`, `Adding/IsScaling`,
  `Isdeleting`, `Normal`, `AbNormal` and `Unknown`.
* `peer_orgs/status_detail` - The peer status in the format like `1/1`. The denominator is the total number of peers in
  the organization, and the numerator is the number of normal peers.
* `peer_orgs/address/domain_port` - The domain name address.
* `peer_orgs/address/ip_port` - The IP address.
* `kafka/name` - The Kafka instance name.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 90 minute.
* `delete` - Default is 30 minute.
