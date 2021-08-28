---
subcategory: "Deprecated"
---

# huaweicloud\_cts\_tracker

!> **WARNING:** It has been deprecated.

CTS Tracker data source allows access of Cloud Tracker.

## Example Usage

```hcl
variable "bucket_name" {}

data "huaweicloud_cts_tracker" "tracker_v1" {
  bucket_name = var.bucket_name
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) The region in which to obtain the Cloud Trackers. If omitted, the provider-level region
  will be used.

* `tracker_name` - (Optional, String) The tracker name.

* `bucket_name` - (Optional, String) The OBS bucket name for a tracker.

* `file_prefix_name` - (Optional, String) The prefix of a log that needs to be stored in an OBS bucket.

* `status` - (Optional, String) Status of a tracker.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a data source ID in UUID format.

* `is_support_smn` -Specifies SMN support.

* `topic_id` - The theme of the SMN service.

* `operations` -The trigger conditions for sending a notification

* `is_send_all_key_operation` - Specifies Typical or All operations for Trigger Condition.

* `need_notify_user_list` - The users using the login function.

