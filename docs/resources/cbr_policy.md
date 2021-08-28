---
subcategory: "Cloud Backup and Recovery (CBR)"
---

# huaweicloud_cbr_policy

Manages a CBR Policy resource within Huaweicloud.

## Example Usage

### create a backup policy

```hcl
variable "policy_name" {}

resource "huaweicloud_cbr_policy" "test" {
  name        = var.policy_name
  type        = "backup"
  time_period = 20

  backup_cycle {
    frequency       = "WEEKLY"
    days            = "MO,TH"
    execution_times = ["06:00"]
  }
}
```

### create a replication policy

```hcl
variable "policy_name" {}
variable "dest_region" {}
variable "dest_project_id" {}

resource "huaweicloud_cbr_policy" "test" {
  name                   = var.policy_name
  type                   = "replication"
  destination_region     = var.dest_region
  destination_project_id = var.dest_project_id
  backup_quantity        = 20

  backup_cycle {
    frequency       = "DAILY"
    interval        = 5
    execution_times = ["21:00"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the CBR policy. If omitted, the
  provider-level region will be used. Changing this will create a new policy.

* `name` - (Required, String) Specifies a unique name of the CBR policy. This parameter can contain a maximum of 64
  characters, which may consist of chinese charactors, letters, digits, underscores(_) and hyphens (-).

* `type` - (Required, String, ForceNew) Specifies the protection type of the CBR policy. Valid values are *backup* and *
  replication*. Changing this will create a new policy.

* `backup_cycle` - (Required, List) Specifies the scheduling rule for the CBR policy backup execution. The backup_cycle
  structure is documented below.

* `enabled` - (Optional, Bool) Specifies whether to enable the CBR policy. Default to true.

* `destination_region` - (Optional, String) Specifies the name of the replication destination region, which is mandatory
  for cross-region replication. Required if `protection_type` is *replication*.

* `destination_project_id` - (Optional, String) Specifies the ID of the replication destination project, which is
  mandatory for cross-region replication. Required if `protection_type` is *replication*.

* `backup_quantity` - (Optional, Int) Specifies the maximum number of retained backups. The value ranges from 2 to
  99999. This parameter and `time_period` are alternative.

* `time_period` - (Optional, Int) Specifies the duration (in days) for retained backups. The value ranges from 2 to
  99999.

-> **NOTE:** If this `backup_quantity` and `time_period` are both left blank, the backups will be retained permanently.

The `backup_cycle` block supports:

* `days` - (Optional, String) Specifies the weekly backup day of backup schedule. It supports seven days a week (MO, TU,
  WE, TH, FR, SA, SU) and this parameter is separated by a comma (,) without spaces, between date and date during the
  configuration.

* `interval` - (Optional, Int) Specifies the interval (in days) of backup schedule. The value ranges from 1 to 30. This
  parameter and `days` are alternative.

* `execution_times` - (Required, List) Specifies the backup time. Automated backups will be triggered at the backup
  time. The current time is in the UTC format (HH:MM). The minutes in the list must be set to *00* and the hours cannot
  be repeated. In the replication policy, you are advised to set one time point for one day.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - A resource ID in UUID format.

## Import

Policies can be imported by their `id`. For example,

```
terraform import huaweicloud_cbr_policy.test 4d2c2939-774f-42ef-ab15-e5b126b11ace
```
