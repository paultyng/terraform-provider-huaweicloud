---
subcategory: "Data Lake Insight (DLI)"
---

# huaweicloud_dli_queue

Manages DLI Queue resource within HuaweiCloud

## Example Usage

### create a queue

```hcl
resource "huaweicloud_dli_queue" "queue" {
  name     = "terraform_dli_queue_test"
  cu_count = 16

  tags = {
    foo = "bar"
    key = "value"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the dli queue resource. If omitted,
  the provider-level region will be used. Changing this will create a new VPC channel resource.

* `name` - (Required, String, ForceNew) Name of a queue. Name of a newly created resource queue. The name can contain
  only digits, letters, and underscores (_), but cannot contain only digits or start with an underscore (_). Length
  range: 1 to 128 characters. Changing this parameter will create a new resource.

* `queue_type` - (Optional, String, ForceNew) Indicates the queue type. Changing this parameter will create a new
  resource. The options are as follows:
  + sql
  + general
      > NOTE: If the type is not specified, the default value sql is used.

* `description` - (Optional, String, ForceNew) Description of a queue. Changing this parameter will create a new
  resource.

* `cu_count` - (Required, Int) Minimum number of CUs that are bound to a queue. Initial value can be `16`,
  `64`, or `256`. When scale_out or scale_in, the number must be a multiple of 16

* `enterprise_project_id` - (Optional, String, ForceNew) Enterprise project ID. The value 0 indicates the default
  enterprise project. Changing this parameter will create a new resource.

* `platform` - (Optional, String, ForceNew) CPU architecture of queue compute resources. Changing this parameter will
  create a new resource. The options are as follows:
  + x86_64 : default value
  + aarch64

* `resource_mode` - (Optional, String, ForceNew) Queue resource mode. Changing this parameter will create a new
  resource. The options are as follows:
  + 0: indicates the shared resource mode.
  + 1: indicates the exclusive resource mode.

* `feature` - (Optional, String, ForceNew)Indicates the queue feature. Changing this parameter will create a new
  resource. The options are as follows:
  + basic: basic type (default value)
  + ai: AI-enhanced (Only the SQL x86_64 dedicated queue supports this option.)

* `tags` - (Optional, Map, ForceNew) Label of a queue. Changing this parameter will create a new resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

* `create_time` - Time when a queue is created.

## Timeouts

This resource provides the following timeouts configuration options:

* `update` - Default is 45 minute.

## Import

DLI queue can be imported by  `id`. For example,

```
terraform import huaweicloud_dli_queue.example  abc123
```
