---
subcategory: "Cloud Container Instance (CCI)"
---

# huaweicloud_cci_pvc

Manages a CCI Persistent Volume Claim resource within HuaweiCloud.

## Example Usage

### Import an EVS volume

```hcl
variable "volume_id" {}

variable "namespace" {}

variable "pvc_name" {}

resource "huaweicloud_cci_pvc" "test" {
  namespace   = var.namespace
  name        = var.pvc_name
  volume_type = "ssd"
  volume_id   = var.volume_id
}
```

### Import an OBS bucket

```hcl
variable "obs_bucket_name" {}

variable "namespace" {}

variable "pvc_name" {}

resource "huaweicloud_cci_pvc" "test" {
  namespace   = var.namespace
  name        = var.pvc_name
  volume_type = "obs"
  volume_id   = var.obs_bucket_name
}
```

### Import an SFS

```hcl
variable "sfs_id" {}

variable "namespace" {}

variable "pvc_name" {}

variable "export_location" {}

resource "huaweicloud_cci_pvc" "test" {
  namespace         = var.namespace
  name              = var.pvc_name
  volume_type       = "nfs-rw"
  volume_id         = var.sfs_id
  device_mount_path = var.export_location
}
```

### Import an SFS Turbo

```hcl
variable "sfs_turbo_id" {}

variable "namespace" {}

variable "pvc_name" {}

variable "export_location" {}

resource "huaweicloud_cci_pvc" "test" {
  namespace         = var.namespace
  name              = var.pvc_name
  volume_type       = "efs-standard"
  volume_id         = var.sfs_turbo_id
  device_mount_path = var.export_location
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the PVC resource. If omitted, the
  provider-level region will be used. Changing this will create a new PVC resource.

* `namespace` - (Required, String, ForceNew) Specifies the namespace to logically divide your cloud container instances
  into different group. Changing this will create a new PVC resource.

* `name` - (Required, String, ForceNew) Specifies the unique name of the PVC resource. This parameter can contain a
  maximum of 63 characters, which may consist of lowercase letters, digits and hyphens, and must start and end with
  lowercase letters and digits. Changing this will create a new PVC resource.

* `volume_id` - (Required, String, ForceNew) Specifies the ID of the storage bound to the CCI Namespace. Changing this
  will create a new PVC resource.

* `volume_type` - (Optional, String, ForceNew) Specifies the type of the storage bound to the CCI Namespace. The valid
  values are *sas*, *ssd*, *sata*, *obs*, *nfs-rw*, *efs-standard* and *efs-performance*, Defalut to *sas*. Changing
  this will create a new PVC resource.

* `device_mount_path` - (Optional, String, ForceNew) Specifies the share path of the SFS storage bound to the CCI
  Namespace. Required if `volume_type` is *nfs-rw*, *efs-standard* or *efs-performance*. Changing this will create a new
  PVC resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The PVC ID in UUID format.

* `access_modes` - The access mode the volume should have.

* `status` - The current phase of the PVC.

* `creation_timestamp` - The server time when PVC was created.

* `enable` - Whether the PVC is available.

## Import

PVCs can be imported using the `namespace`, `volume_type` and `id`, e.g.

```
$ terraform import huaweicloud_cci_pvc.test <namespace>/<volume_type>/<id>
```

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 5 minute.
* `delete` - Default is 3 minute.
