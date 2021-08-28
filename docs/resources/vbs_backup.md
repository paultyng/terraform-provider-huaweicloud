---
subcategory: "Deprecated"
---

# huaweicloud\_vbs\_backup

!> **WARNING:** It has been deprecated.

Provides an VBS Backup resource. This is an alternative to `huaweicloud_vbs_backup_v2`

# Example Usage

```hcl
resource "huaweicloud_evs_volume" "volume" {
  name              = "volume"
  description       = "my volume"
  volume_type       = "SATA"
  size              = 20
  availability_zone = "cn-north-4a"
}

resource "huaweicloud_evs_snapshot" "snapshot_1" {
  name        = "snapshot-001"
  description = "for vbs backup"
  volume_id   = huaweicloud_evs_volume.volume.id
}

resource "huaweicloud_vbs_backup" "backup_1" {
  volume_id   = huaweicloud_evs_volume.volume.id
  snapshot_id = huaweicloud_evs_snapshot.snapshot_1.id
  name        = "vbs-backup"
  description = "Backup_Demo"
  tags {
    key   = "bar"
    value = "value"
  }
}
```

# Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the VBS backup resource. If omitted, the
  provider-level region will be used. Changing this creates a new VBS Backup resource.

* `name` - (Required, String, ForceNew) The name of the vbs backup. Changing the parameter creates a new backup.

* `volume_id` - (Required, String, ForceNew) The id of the disk to be backed up. Changing the parameter creates a new
  backup.

* `snapshot_id` - (Optional, String, ForceNew) The snapshot id of the disk to be backed up. Changing the parameter
  creates a new backup.

* `description` - (Optional, String, ForceNew) The description of the vbs backup. Changing the parameter creates a new
  backup.

* `tags` - (Optional, List, ForceNew) List of tags to be configured for the backup resources. Changing the parameter
  creates a new backup.

  + `key` - (Required, String, ForceNew) Specifies the tag key. Changing the parameter creates a new backup.

  + `value` - (Required, String, ForceNew) Specifies the tag value. Changing the parameter creates a new backup.

# Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id of the vbs backup.

* `container` - The container of the backup.

* `created_at` - Backup creation time.

* `status` - The status of the VBS backup.

* `availability_zone` - The AZ where the backup resides.

* `size` - The size of the vbs backup.

* `object_count` - Number of objects on Object Storage Service (OBS) for the disk data.

* `service_metadata` - The metadata of the vbs backup.

# Import

VBS Backup can be imported using the `backup id`, e.g.

```
 $ terraform import huaweicloud_vbs_backup.backup_1 4779ab1c-7c1a-44b1-a02e-93dfc361b32d
```

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 3 minute.

