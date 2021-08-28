---
subcategory: "Auto Scaling"
---

# huaweicloud_as_policy

Manages a AS Policy resource within HuaweiCloud. This is an alternative to `huaweicloud_as_policy_v1`

## Example Usage

### AS Recurrence Policy

```hcl
resource "huaweicloud_as_policy" "my_aspolicy" {
  scaling_policy_name = "my_aspolicy"
  scaling_policy_type = "RECURRENCE"
  scaling_group_id    = "4579f2f5-cbe8-425a-8f32-53dcb9d9053a"
  cool_down_time      = 900

  scaling_policy_action {
    operation       = "ADD"
    instance_number = 1
  }
  scheduled_policy {
    launch_time     = "07:00"
    recurrence_type = "Daily"
    start_time      = "2020-11-30T12:00Z"
    end_time        = "2020-12-30T12:00Z"
  }
}
```

### AS Scheduled Policy

```hcl
resource "huaweicloud_as_policy" "my_aspolicy_1" {
  scaling_policy_name = "my_aspolicy_1"
  scaling_policy_type = "SCHEDULED"
  scaling_group_id    = "4579f2f5-cbe8-425a-8f32-53dcb9d9053a"
  cool_down_time      = 900

  scaling_policy_action {
    operation       = "REMOVE"
    instance_number = 1
  }
  scheduled_policy {
    launch_time = "2020-12-22T12:00Z"
  }
}
```

Please note that the `launch_time` of the `SCHEDULED` policy cannot be earlier than the current time.

### AS Alarm Policy

```hcl
resource "huaweicloud_ces_alarmrule" "alarm_rule" {
  alarm_name = "as_alarm_rule"

  metric {
    namespace   = "SYS.AS"
    metric_name = "cpu_util"
    dimensions {
      name  = "AutoScalingGroup"
      value = "4579f2f5-cbe8-425a-8f32-53dcb9d9053a"
    }
  }
  condition {
    period              = 300
    filter              = "average"
    comparison_operator = ">="
    value               = 60
    unit                = "%"
    count               = 1
  }
  alarm_actions {
    type              = "autoscaling"
    notification_list = []
  }
}

resource "huaweicloud_as_policy" "my_aspolicy_2" {
  scaling_policy_name = "my_aspolicy_2"
  scaling_policy_type = "ALARM"
  scaling_group_id    = "4579f2f5-cbe8-425a-8f32-53dcb9d9053a"
  alarm_id            = huaweicloud_ces_alarmrule.alarm_rule.id
  cool_down_time      = 900

  scaling_policy_action {
    operation       = "ADD"
    instance_number = 1
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the AS policy. If omitted, the `region` argument
  of the provider is used. Changing this creates a new AS policy.

* `scaling_policy_name` - (Required, String) The name of the AS policy. The name can contain letters, digits,
  underscores(_), and hyphens(-), and cannot exceed 64 characters.

* `scaling_policy_type` - (Required, String) The AS policy type. The values can be `ALARM`, `SCHEDULED`,
  and `RECURRENCE`.

* `scaling_group_id` - (Required, String, ForceNew) The AS group ID. Changing this creates a new AS policy.

* `alarm_id` - (Optional, String) The alarm rule ID. This argument is mandatory when `scaling_policy_type`
  is set to `ALARM`. You can create an alarm rule with `huaweicloud_ces_alarmrule`.

* `scheduled_policy` - (Optional, List) The periodic or scheduled AS policy. This argument is mandatory
  when `scaling_policy_type` is set to `SCHEDULED` or `RECURRENCE`. The scheduled_policy structure is documented below.

* `scaling_policy_action` - (Optional, List) The action of the AS policy. The scaling_policy_action structure is
  documented below.

* `cool_down_time` - (Optional, Int) The cooling duration (in seconds), and is 900 by default.

The `scheduled_policy` block supports:

* `launch_time` - (Required, String) The time when the scaling action is triggered.
  + If `scaling_policy_type` is set to `SCHEDULED`, the time format is YYYY-MM-DDThh:mmZ.
  + If `scaling_policy_type` is set to `RECURRENCE`, the time format is hh:mm.

* `recurrence_type` - (Optional, String) The periodic triggering type. This argument is mandatory when
  `scaling_policy_type` is set to `RECURRENCE`. The options include `Daily`, `Weekly`, and `Monthly`.

* `recurrence_value` - (Optional, String) The frequency at which scaling actions are triggered.

* `start_time` - (Optional, String) The start time of the scaling action triggered periodically. The time format
  complies with UTC. The current time is used by default. The time format is YYYY-MM-DDThh:mmZ.

* `end_time` - (Optional, String) The end time of the scaling action triggered periodically. The time format complies
  with UTC. This argument is mandatory when `scaling_policy_type`
  is set to `RECURRENCE`. The time format is YYYY-MM-DDThh:mmZ.

The `scaling_policy_action` block supports:

* `operation` - (Optional, String) The operation to be performed. The options include `ADD` (default), `REMOVE`,
  and `SET`.

* `instance_number` - (Optional, Int) The number of instances to be operated. The default number is 1.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

