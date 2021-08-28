---
subcategory: "Deprecated"
---

# huaweicloud\_cts\_tracker

!> **WARNING:** It has been deprecated.

Allows you to collect, store, and query cloud resource operation records.

## Example Usage

 ```hcl
variable "bucket_name" {}
variable "topic_id" {}

resource "huaweicloud_cts_tracker" "tracker_v1" {
  bucket_name               = var.bucket_name
  file_prefix_name          = "yO8Q"
  is_support_smn            = true
  topic_id                  = var.topic_id
  is_send_all_key_operation = false
  operations                = ["login"]
  need_notify_user_list     = ["user1"]
}

 ```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the CTS tracker resource. If omitted, the
  provider-level region will be used. Changing this creates a new CTS tracker resource.

* `bucket_name` - (Required, String) The OBS bucket name for a tracker.

* `file_prefix_name` - (Optional, String) The prefix of a log that needs to be stored in an OBS bucket.

* `is_support_smn` - (Required, Bool) Specifies whether SMN is supported. When the value is false, topic_id and
  operations can be left empty.

* `topic_id` - (Optional, String) Required if the value of `is_support_smn` is true. The theme of the SMN service, Is
  obtained from SMN and in the format of **urn:smn:([a-z]|[A-Z]|[0-9]|\-){1,32}:([a-z]|[A-Z]|[0-9]){32}:([a-z]|[A-Z]
  |[0-9]|\-|\_){1,256}**.

* `operations` - (Required, String) Trigger conditions for sending a notification.

* `is_send_all_key_operation` - (Required, Bool) When the value is **false**, operations cannot be left empty.

* `need_notify_user_list` - (Optional, String) The users using the login function. When these users log in,
  notifications will be sent.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

* `status` - The status of a tracker. The value is **enabled**.

* `tracker_name` - The tracker name. Currently, only tracker **system** is available.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

CTS tracker can be imported using  `tracker_name`, e.g.

```
$ terraform import huaweicloud_cts_tracker_v1.tracker system
```

