---
subcategory: "Cloud Stream"
---

# huaweicloud_cs_peering_connect

Cloud Stream Service cluster peering connect management
This is an alternative to `huaweicloud_cs_peering_connect_v1`

!> **WARNING:** It has been deprecated, use `huaweicloud_dli_queue` instead.

## Example Usage

### create a cluster peering connect

```hcl
resource "huaweicloud_cs_cluster" "cluster" {
  name = "terraform_cs_cluster_v1_test"
}

resource "huaweicloud_vpc" "vpc" {
  name = "terraform_vpc_v1_test"
  cidr = "192.168.0.0/16"
}

resource "huaweicloud_vpc_subnet" "subnet" {
  name       = "terraform_vpc_subnet_test"
  cidr       = "192.168.0.0/16"
  gateway_ip = "192.168.0.1"
  vpc_id     = huaweicloud_vpc.vpc.id
}

resource "huaweicloud_cs_peering_connect" "peering" {
  name = "terraform_cs_peering_connect_test"
  target_vpc_info {
    vpc_id = huaweicloud_vpc.vpc.id
  }
  cluster_id = huaweicloud_cs_cluster.cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the cs peering connection resource. If omitted, the provider-level region will be used. Changing this creates a new cs peering connection resource.

* `cluster_id` - (Required, String, ForceNew) The id of cloud stream cluster. Changing this parameter will create a new resource.

* `name` - (Required, String, ForceNew) The name of peering connection. Changing this parameter will create a new resource.

* `target_vpc_info` - (Optional, List, ForceNew) The information of target vpc. Structure is documented below. Changing this parameter will create a new resource.

The `target_vpc_info` block supports:

* `project_id` - (Optional, String, ForceNew) The project ID to which target vpc belongs. Changing this parameter will create a new resource.

* `vpc_id` - (Required, String, ForceNew) The VPC ID. Changing this parameter will create a new resource.

## Timeouts
This resource provides the following timeouts configuration options:
- `create` - Default is 30 minute.
- `delete` - Default is 30 minute.

