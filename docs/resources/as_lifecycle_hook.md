---
subcategory: "Auto Scaling"
---

# huaweicloud_as_lifecycle_hook

Manages an AS Lifecycle Hook resource within HuaweiCloud.

## Example Usage

### Basic Lifecycle Hook

```hcl
variable "hook_name" {}

variable "as_group_id" {}

variable "smn_topic_urn" {}

resource "huaweicloud_as_lifecycle_hook" "test" {
  name                   = var.hook_name
  type                   = "ADD"
  scaling_group_id       = var.as_group_id
  default_result         = "ABANDON"
  notification_topic_urn = var.smn_topic_urn
  notification_message   = "This is a test message"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the AS lifecycle hook. If omitted, the `region`
  argument of the provider is used. Changing this creates a new AS lifecycle hook.

* `name` - (Required, String) Specifies the lifecycle hook name. This parameter can contain a maximum of 32 characters,
  which may consist of letters, digits, underscores (_) and hyphens (-).

* `type` - (Required, String) Specifies the lifecycle hook type. The valid values are following strings:
  + `ADD`: The hook suspends the instance when the instance is started.
  + `REMOVE`: The hook suspends the instance when the instance is terminated.

* `notification_topic_urn` - (Required, String) Specifies a unique topic in SMN.

* `scaling_group_id` - (Required, String, ForceNew) Specifies the ID of the AS group in UUID format. Changing this
  creates a new AS lifecycle hook.

* `default_result` - (Optional, String) Specifies the default lifecycle hook callback operation. This operation is
  performed when the timeout duration expires. The valid values are *ABANDON* and *CONTINUE*, default to *ABANDON*.

* `timeout` - (Optional, Int) Specifies the lifecycle hook timeout duration, which ranges from 300 to 86400 in the unit
  of second, default to 3600.

* `notification_message` - (Optional, String) Specifies a customized notification. This parameter can contains a maximum
  of 256 characters, which cannot contain the following characters: <>&'().

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.

* `notification_topic_name` - The topic name in SMN.

* `create_time` - The server time in UTC format when the lifecycle hook is created.

## Import

Lifecycle hooks can be imported using the AS group ID and hook ID separated by a slash, e.g.

```
$ terraform import huaweicloud_as_lifecycle_hook.test <AS group ID>/<Lifecycle hook ID>
```

