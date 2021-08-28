---
subcategory: "Data Warehouse Service (DWS)"
---

# huaweicloud_dws_cluster

Manages Cluster in the Data Warehouse Service.

## Example Usage

### Dws Cluster Example

```hcl
resource "huaweicloud_networking_secgroup" "secgroup" {
  name        = "security_group_2"
  description = "terraform security group"
}

resource "huaweicloud_dws_cluster" "cluster" {
  node_type         = "dws.m3.xlarge"
  number_of_node    = 3
  network_id        = "{{ network_id }}"
  vpc_id            = "{{ vpc_id }}"
  security_group_id = huaweicloud_networking_secgroup.secgroup.id
  availability_zone = "{{ availability_zone }}"
  name              = "terraform_dws_cluster_test"
  user_name         = "test_cluster_admin"
  user_pwd          = "cluster123@!"

  timeouts {
    create = "30m"
    delete = "30m"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the cluster resource. If omitted, the
  provider-level region will be used. Changing this creates a new cluster resource.

* `name` - (Required, String, ForceNew) Cluster name, which must be unique and contains 4 to 64 characters, which
  consist of letters, digits, hyphens(-), or underscores(_) only and must start with a letter.

* `network_id` - (Required, String, ForceNew) Network ID, which is used for configuring cluster network

* `node_type` - (Required, String, ForceNew) Node type

* `number_of_node` - (Required, Int, ForceNew) Number of nodes in a cluster. The value ranges from 3 to 32

* `security_group_id` - (Required, String, ForceNew) ID of a security group. The ID is used for configuring cluster
  network

* `user_name` - (Required, String, ForceNew) Administrator username for logging in to a data warehouse cluster The
  administrator username must:  Consist of lowercase letters, digits, or underscores. Start with a lowercase letter or
  an underscore. Contain 1 to 63 characters. Cannot be a keyword of the DWS database.

* `vpc_id` - (Required, String, ForceNew) VPC ID, which is used for configuring cluster network

* `user_pwd` - (Required, String, ForceNew) Administrator password for logging in to a data warehouse cluster A password
  must conform to the following rules:  Contains 8 to 32 characters. Cannot be the same as the username or the username
  written in reverse order. Contains three types of the following:
  Lowercase letters Uppercase letters Digits Special characters
  ~!@#%^&*()-_=+|[{}];:,<.>/?

* `enterprise_project_id` - (Optional, String, ForceNew) Specifies the enterprise project id of the dws cluster, Value 0
  indicates the default enterprise project. Changing this parameter will create a new resource.

* `availability_zone` - (Optional, String, ForceNew) AZ in a cluster

* `port` - (Optional, Int) Service port of a cluster (8000 to 10000). The default value is 8000.

* `public_ip` - (Optional, List, ForceNew) A nested object resource Structure is documented below.

The `public_ip` block supports:

* `eip_id` - (Optional, String, ForceNew) EIP ID

* `public_bind_type` - (Optional, String, ForceNew) Binding type of an EIP. The value can be either of the following:
  auto_assign not_use bind_existing The default value is not_use.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `created` - Cluster creation time. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ

* `endpoints` - View the private network connection information about the cluster. Structure is documented below.

* `id` - Cluster ID

* `public_endpoints` - Public network connection information about the cluster. If the value is not specified, the
  public network connection information is not used by default Structure is documented below.

* `recent_event` - The recent event number

* `status` - Cluster status, which can be one of the following:  CREATING AVAILABLE UNAVAILABLE CREATION FAILED

* `sub_status` - Sub-status of clusters in the AVAILABLE state. The value can be one of the following:  NORMAL READONLY
  REDISTRIBUTING REDISTRIBUTION-FAILURE UNBALANCED UNBALANCED | READONLY DEGRADED DEGRADED | READONLY DEGRADED |
  UNBALANCED UNBALANCED | REDISTRIBUTING UNBALANCED | REDISTRIBUTION-FAILURE READONLY | REDISTRIBUTION-FAILURE
  UNBALANCED | READONLY | REDISTRIBUTION-FAILURE DEGRADED | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED |
  REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | READONLY | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | READONLY

* `task_status` - Cluster management task. The value can be one of the following:
  RESTORING SNAPSHOTTING GROWING REBOOTING SETTING_CONFIGURATION CONFIGURING_EXT_DATASOURCE DELETING_EXT_DATASOURCE
  REBOOT_FAILURE RESIZE_FAILURE

* `updated` - Last modification time of a cluster. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ

* `version` - Data warehouse version

The `endpoints` block contains:

* `connect_info` - (Optional, String) Private network connection information

* `jdbc_url` - (Optional, String)
  JDBC URL. The following is the default format:
  jdbc:postgresql://< connect_info>/<YOUR_DATABASE_NAME>

The `public_endpoints` block contains:

* `jdbc_url` - (Optional, String)
  JDBC URL. The following is the default format:
  jdbc:postgresql://< public_connect_info>/<YOUR_DATABASE_NAME>

* `public_connect_info` - (Optional, String)
  Public network connection information

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

Cluster can be imported using the following format:

```
$ terraform import huaweicloud_dws_cluster.default {{ resource id}}
```
