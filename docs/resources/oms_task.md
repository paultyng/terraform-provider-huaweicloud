---
subcategory: "Object Storage Migration Service"
---

# huaweicloud\_oms\_task

Manages Object Storage Migration task within HuaweiCloud.

## Example Usage:  Creating a OMS task

```hcl
resource "huaweicloud_oms_task" "task_1" {
  description = "migration task"
  enable_kms  = true
  thread_num  = 5

  src_node {
    region     = "cn-north-1"
    ak         = var.src_AK
    sk         = var.src_SK
    cloud_type = "HuaweiCloud"
    bucket     = "oms-bucket"
    object_key = "123.txt"
  }
  dst_node {
    region     = "cn-east-3"
    ak         = var.dst_AK
    sk         = var.dst_SK
    bucket     = "test-oms"
    object_key = "oms"
  }
}
```

## Argument Reference

The following arguments are supported:

* `src_node` - (Required, List, ForceNew) Specifies the source node information.

* `dst_node` - (Required, List, ForceNew) Specifies the destination node information.

* `enable_kms` - (Required, Bool, ForceNew) Specifies whether to use KMS encryption.

* `thread_num` - (Required, Int, ForceNew) Specifies the number of threads used by the migration task. The value cannot
  exceed 50.

* `description` - (Optional, String, ForceNew) Specifies tasks description, which cannot exceed 255 characters. The
  following special characters are not allowed: <>()"&

* `smn_info` - (Optional, List, ForceNew) Specifies the field used for sending messages using the Simple Message
  Notification (SMN) service.

The `src_node` block supports:

* `region` - (Required, String, ForceNew) Specifies the region where the source bucket locates.
* `ak` - (Required, String, ForceNew) Specifies the source bucket Access Key.
* `sk` - (Required, String, ForceNew) Specifies the source bucket Secret Key.
* `bucket` - (Required, String, ForceNew) Specifies the name of the source bucket.
* `object_key` - (Required, String, ForceNew) Specifies the name of the object to be selected in the source bucket.
* `cloud_type` - (Optional, String, ForceNew) Specifies the source cloud service provider. The value can be AWS, Aliyun,
  Tencent, HuaweiCloud, QingCloud, KingsoftCloud, Baidu, or Qiniu. The default value is Aliyun.

The `dst_node` block supports:

* `region` - (Required, String, ForceNew) Specifies the region where the destination bucket locates.
* `ak` - (Required, String, ForceNew) Specifies the destination bucket Access Key.
* `sk` - (Required, String, ForceNew) Specifies the destination bucket Secret Key.
* `bucket` - (Required, String, ForceNew) Specifies the name of the destination bucket.
* `object_key` - (Required, String, ForceNew) Specifies the name of the object to be selected in the destination bucket.

The `smn_info` block supports:

* `topic_urn` - (Required, String, ForceNew) Specifies the SMN message topic URN bound to a migration task.
* `language` - (Optional, String, ForceNew) Specifies the management console language used by the current users. Users
  can select en-us.
* `trigger_conditions` - (Required, List, ForceNew) Specifies the trigger conditions of sending messages using SMN. The
  value depending on the state of a migration task. The migration task status can be SUCCESS or FAIL.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.
* `name` - Indicates the name for a task.
* `status` - Indicates the task status as follows: 0: Not started, 1: Waiting to migrate, 2: Migrating, 3: Migration
  paused, 4: Migration failed, 5: Migration succeeded.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.

