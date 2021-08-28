---
subcategory: "Cloud Stream"
---

# huaweicloud_cs_route

Cloud Stream Service cluster peering connect route management This is an alternative to `huaweicloud_cs_route_v1`

## Example Usage

### create a cluster peering connect route

```hcl
resource "huaweicloud_cs_cluster" "cluster" {
  name = "terraform_cs_cluster_test"
}

resource "huaweicloud_vpc" "vpc" {
  name = "terraform_vpc_test"
  cidr = "192.168.0.0/16"
}

resource "huaweicloud_vpc_subnet" "subnet" {
  name       = "terraform_vpc_subnet_test"
  cidr       = "192.168.0.0/16"
  gateway_ip = "192.168.0.1"
  vpc_id     = huaweicloud_vpc.vpc.id
}

resource "huaweicloud_cs_peering_connect" "peering" {
  name       = "terraform_cs_peering_connect_test"
  target_vpc_info {
    vpc_id = huaweicloud_vpc.vpc.id
  }
  cluster_id = huaweicloud_cs_cluster.cluster.id
}

resource "huaweicloud_cs_route" "route" {
  cluster_id  = huaweicloud_cs_cluster.cluster.id
  peering_id  = huaweicloud_cs_peering_connect.peering.id
  destination = huaweicloud_vpc_subnet.subnet.cidr
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the cs peering connect route resource. If
  omitted, the provider-level region will be used. Changing this creates a new cs peering connect route resource.

* `cluster_id` - (Required, String, ForceNew) The id of cloud stream cluster. Changing this parameter will create a new
  resource.

* `destination` - (Required, String, ForceNew) Routing destination CIDR. Changing this parameter will create a new
  resource.

* `peering_id` - (Required, String, ForceNew) The peering connection id of cloud stream cluster. Changing this parameter
  will create a new resource.
