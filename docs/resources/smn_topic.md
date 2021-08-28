---
subcategory: "Simple Message Notification (SMN)"
---

# huaweicloud_smn_topic

Manages a SMN Topic resource within HuaweiCloud. This is an alternative to `huaweicloud_smn_topic_v2`

## Example Usage

```hcl
resource "huaweicloud_smn_topic" "topic_1" {
  name         = "topic_1"
  display_name = "The display name of topic_1"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the SMN topic resource. If omitted, the
  provider-level region will be used. Changing this creates a new SMN Topic resource.

* `name` - (Required, String, ForceNew) The name of the topic to be created.

* `display_name` - (Optional, String) Topic display name, which is presented as the name of the email sender in an email
  message.

* `enterprise_project_id` - (Optional, String, ForceNew) Specifies the enterprise project id of the SMN Topic, Value 0
  indicates the default enterprise project. Changing this parameter will create a new resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

* `topic_urn` - Resource identifier of a topic, which is unique.

* `push_policy` - Message pushing policy. 0 indicates that the message sending fails and the message is cached in the
  queue. 1 indicates that the failed message is discarded.

* `create_time` - Time when the topic was created.

* `update_time` - Time when the topic was updated.

