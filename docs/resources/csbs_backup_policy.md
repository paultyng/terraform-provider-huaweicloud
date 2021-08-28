---
subcategory: "Deprecated"
---

# huaweicloud_csbs_backup_policy

!> **WARNING:** It has been deprecated.

Provides an HuaweiCloud Backup Policy of Resources. This is an alternative to `huaweicloud_csbs_backup_policy_v1`

## Example Usage

 ```hcl
variable "name" {}
variable "id" {}
variable "resource_name" {}

resource "huaweicloud_csbs_backup_policy" "backup_policy" {
  name = var.name
  resource {
    id   = var.id
    type = "OS::Nova::Server"
    name = var.resource_name
  }
  scheduled_operation {
    enabled         = true
    operation_type  = "backup"
    trigger_pattern = "BEGIN:VCALENDAR\r\nBEGIN:VEVENT\r\nRRULE:FREQ=WEEKLY;BYDAY=TH;BYHOUR=12;BYMINUTE=27\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n"
  }
}

 ```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the Backup Policy resource. If omitted, the
  provider-level region will be used. Changing this creates a new Backup Policy resource.

* `name` - (Required, String, ForceNew) Specifies the name of backup policy. The value consists of 1 to 255 characters
  and can contain only letters, digits, underscores (_), and hyphens (-).

* `description` - (Optional, String, ForceNew) Backup policy description. The value consists of 0 to 255 characters and
  must not contain a greater-than sign (>) or less-than sign (<).

* `provider_id` - (Required, String) Specifies backup provider ID. Default value is **
  fc4d5750-22e7-4798-8a46-f48f62c4c1da**

* `common` - (Optional, Map) General backup policy parameters, which are blank by default.

* `scheduled_operation` block supports the following arguments:

  + `name` - (Optional, String) Specifies Scheduling period name.The value consists of 1 to 255 characters and can
      contain only letters, digits, underscores (_), and hyphens (-).

  + `description` - (Optional, String) Specifies Scheduling period description.The value consists of 0 to 255
      characters and must not contain a greater-than sign (>) or less-than sign (<).

  + `enabled` - (Optional, Bool) Specifies whether the scheduling period is enabled. Default value is **true**

  + `max_backups` - (Optional, Int) Specifies maximum number of backups that can be automatically created for a backup
      object.

  + `retention_duration_days` - (Optional, Int) Specifies duration of retaining a backup, in days.

  + `permanent` - (Optional, Bool) Specifies whether backups are permanently retained.

  + `trigger_pattern` - (Required, String) Specifies Scheduling policy of the scheduler.

  + `operation_type` - (Required, String) Specifies Operation type, which can be backup.

* `resource` block supports the following arguments:

  + `id` - (Required, String) Specifies the ID of the object to be backed up.

  + `type` - (Required, String) Entity object type of the backup object. If the type is VMs, the value is **OS::Nova::
      Server**.

  + `name` - (Required, String) Specifies backup object name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `status` - Status of Backup Policy.

* `id` - Backup Policy ID.

* `created_at` - Creation time.

* `scheduled_operation` - Backup plan information

  + `id` - Specifies Scheduling period ID.

  + `trigger_id` - Specifies Scheduler ID.

  + `trigger_name` - Specifies Scheduler name.

  + `trigger_type` - Specifies Scheduler type.

## Import

Backup Policy can be imported using  `id`, e.g.

```
$ terraform import huaweicloud_csbs_backup_policy.backup_policy 7056d636-ac60-4663-8a6c-82d3c32c1c64
```
