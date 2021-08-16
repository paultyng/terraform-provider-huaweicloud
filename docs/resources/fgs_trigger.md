---
subcategory: "FunctionGraph"
---

# huaweicloud_fgs_trigger

Manages a trigger resource within HuaweiCloud FunctionGraph.

## Example Usage

### Timing Trigger with rate schedule type

```hcl
variable "function_urn" {}
variable "timer_name" {}

resource "huaweicloud_fgs_trigger" "test" {
  function_urn = var.function_urn
  type         = "TIMER"

  timer {
    name          = var.timer_name
    schedule_type = "Rate"
    schedule      = "1d"
  }
}
```

### Timing Trigger with cron schedule type

```hcl
variable "function_urn" {}
variable "timer_name" {}

resource "huaweicloud_fgs_trigger" "test" {
  function_urn = huaweicloud_fgs_function.test.urn
  type         = "TIMER"

  timer {
    name          = var.timer_name
    schedule_type = "Cron"
    schedule      = "@every 1h30m"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the trigger resource.
  If omitted, the provider-level region will be used.
  Changing this will create a new trigger resource.

* `function_urn` - (Required, String, ForceNew) Specifies the Uniform Resource Name (URN) of the function.
  Changing this will create a new trigger resource.

* `type` - (Required, String, ForceNew) Specifies the type of the function.
  The valid values currently only support __TIMER__.
  Changing this will create a new trigger resource.

* `status` - (Optional, String) Specifies whether trigger is enabled, default to true.

* `timer` - (Optional, List, ForceNew) Specifies the configuration of the timing trigger.
  Changing this will create a new trigger resource.
  The `timer` object structure is documented below.

The `timer` block supports:

* `name` - (Required, String, ForceNew) Specifies the trigger name, which can contains of 1 to 64 characters.
  The name must start with a letter, only letters, digits, hyphens (-) and underscores (_) are allowed.
  Changing this will create a new trigger resource.

* `schedule_type` - (Required, String, ForceNew) Specifies the type of the time schedule.
  The valid values are __Rate__ and __Corn__.
  Changing this will create a new trigger resource.

* `schedule` - (Required, String, ForceNew) Specifies the time schedule.
  For the rate type, schedule is composed of time and time unit. The time unit supports minutes, hours and days.
  For the corn expression, please refer to the
  HuaweiCloud [document](https://support.huaweicloud.com/en-us/usermanual-functiongraph/functiongraph_01_0908.html).
  Changing this will create a new trigger resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - resource ID in UUID format.

## Import

Triggers can be imported using their `id` and URN of the FunctionGraph to which the trigger belongs, separated by a slash, e.g.

```
$ terraform import huaweicloud_fgs_trigger.test 80e4640f-9e97-40ae-a787-8cb9b2f68528
```
