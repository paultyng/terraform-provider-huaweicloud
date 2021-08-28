---
subcategory: "Deprecated"
---

# huaweicloud\_csbs\_backup\_policy

!> **WARNING:** It has been deprecated.

The HuaweiCloud CSBS Backup Policy data source allows access of backup Policy resources. This is an alternative
to `huaweicloud_csbs_backup_policy_v1`

## Example Usage

```hcl
variable "policy_id" {}

data "huaweicloud_csbs_backup_policy" "csbs_policy" {
  id = var.policy_id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) The region in which to obtain the CSBS backup policies. If omitted, the provider-level
  region will be used.

* `id` - (Optional, String) Specifies the ID of backup policy.

* `name` - (Optional, String) Specifies the backup policy name.

* `status` - (Optional, String) Specifies the backup policy status.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - Specifies the backup policy description.

* `provider_id` - Provides the Backup provider ID.

* `parameters` - Specifies the parameters of a backup policy.

* `scheduled_operation` block supports the following arguments:
  + `name` - Specifies Scheduling period name.

  + `description` - Specifies Scheduling period description.

  + `enabled` - Specifies whether the scheduling period is enabled.

  + `max_backups` - Specifies maximum number of backups that can be automatically created for a backup object.

  + `retention_duration_days` - Specifies duration of retaining a backup, in days.

  + `permanent` - Specifies whether backups are permanently retained.

  + `trigger_pattern` - Specifies Scheduling policy of the scheduler.

  + `operation_type` - Specifies Operation type, which can be backup.

  + `id` - Specifies Scheduling period ID.

  + `trigger_id` - Specifies Scheduler ID.

  + `trigger_name` - Specifies Scheduler name.

  + `trigger_type` - Specifies Scheduler type.

* `resource` block supports the following arguments:

  + `id` - Specifies the ID of the object to be backed up.

  + `type` - Entity object type of the backup object.

  + `name` - Specifies backup object name.
