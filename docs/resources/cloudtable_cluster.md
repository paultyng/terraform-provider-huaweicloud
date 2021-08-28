---
subcategory: "Cloud Table"
---

# huaweicloud_cloudtable_cluster

Cloud table cluster management This is an alternative to `huaweicloud_cloudtable_cluster_v2`

## Example Usage

### create a CloudTable cluster

```hcl
resource "huaweicloud_networking_secgroup" "secgroup" {
  name        = "terraform_test_security_group"
  description = "terraform security group acceptance test"
}

resource "huaweicloud_cloudtable_cluster" "cluster" {
  availability_zone = "{{ availability_zone }}"
  name              = "terraform-test-cluster"
  rs_num            = 2
  security_group_id = huaweicloud_networking_secgroup.secgroup.id
  subnet_id         = "{{ network_id }}"
  vpc_id            = "{{ vpc_id }}"
  storage_type      = "COMMON"
}
```

## Argument Reference

The following arguments are supported:

* `availability_zone` - (Required, String, ForceNew) Availability zone (AZ). Changing this parameter will create a new
  resource.

* `name` - (Required, String, ForceNew) Cluster name. The value must be between 4 and 64 characters long and start with
  a letter. Only letters, digits, and hyphens (-) are allowed. It is case insensitive. Changing this parameter will
  create a new resource.

* `rs_num` - (Required, Int, ForceNew) Number of computing units. Value range: 2 to 10. Changing this parameter will
  create a new resource.

* `security_group_id` - (Required, String, ForceNew) Security group ID. Changing this parameter will create a new
  resource.

* `storage_type` - (Required, String, ForceNew) Storage I/O type. The value are ULTRAHIGH and COMMON. Changing this
  parameter will create a new resource.

* `subnet_id` - (Required, String, ForceNew) Subnet ID. Changing this parameter will create a new resource.

* `vpc_id` - (Required, String, ForceNew) VPC of the cluster. Changing this parameter will create a new resource.

* `enable_iam_auth` - (Optional, Bool, ForceNew) Whether to enable IAM authentication for OpenTSDB. Changing this
  parameter will create a new resource.

* `lemon_num` - (Optional, Int, ForceNew) Number of Lemon nodes Value range: 2 to 10. Changing this parameter will
  create a new resource.

* `opentsdb_num` - (Optional, Int, ForceNew) Number of OpenTSDB nodes Value range: 2 to 10. Changing this parameter will
  create a new resource.

* `tags` - (Optional, List, ForceNew) Enterprise project information. Structure is documented below. Changing this
  parameter will create a new resource.

The `tags` block supports:

* `key` - (Optional, String, ForceNew) Tag value. Changing this parameter will create a new resource.

* `value` - (Optional, String, ForceNew) Tag key. Changing this parameter will create a new resource.

* `region` - (Optional, String, ForceNew) The region in which to create the cloud table cluster resource. If omitted,
  the provider-level region will be used. Changing this creates a new cloud table cluster resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

* `created` - Time when the cluster was created.

* `hbase_public_endpoint` - HBase link of the public network.

* `lemon_link` - Lemon link of the intranet.

* `open_tsdb_link` - OpenTSDB link of the intranet.

* `opentsdb_public_endpoint` - OpenTSDB link of the public network.

* `storage_quota` - Storage quota.

* `used_storage_size` - Used storage space.

* `zookeeper_link` - ZooKeeper link of the intranet.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 30 minute.

