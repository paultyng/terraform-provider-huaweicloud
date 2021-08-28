---
subcategory: "Elastic Volume Service (EVS)"
---

# huaweicloud_evs_snapshot

Provides an EVS snapshot resource.

# Example Usage

```hcl
resource "huaweicloud_evs_volume" "myvolume" {
  name        = "volume"
  description = "my volume"
  volume_type = "SATA"
  size        = 20

  availability_zone = "cn-north-4a"

  tags = {
    foo = "bar"
    key = "value"
  }
}

resource "huaweicloud_evs_snapshot" "snapshot_1" {
  name        = "snapshot-001"
  description = "Daily backup"
  volume_id   = huaweicloud_evs_volume.myvolume.id
}
```

# Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the evs snapshot resource. If omitted, the
  provider-level region will be used. Changing this creates a new EVS snapshot resource.

* `volume_id` - (Required, String, ForceNew) The id of the snapshot's source disk. Changing the parameter creates a new
  snapshot.

* `name` - (Required, String) The name of the snapshot. The value can contain a maximum of 255 bytes.

* `description` - (Optional, String) The description of the snapshot. The value can contain a maximum of 255 bytes.

* `force` - (Optional, Bool) Specifies the flag for forcibly creating a snapshot. Default to false.

# Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id of the snapshot.

* `status` - The status of the snapshot.

* `size` - The size of the snapshot in GB.

# Import

EVS snapshot can be imported using the `snapshot id`, e.g.

```
 $ terraform import huaweicloud_evs_snapshot.snapshot_1 3a11b255-3bb6-46f3-91e4-3338baa92dd6
```

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 3 minute.

