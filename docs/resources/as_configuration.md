---
subcategory: "Auto Scaling"
---

# huaweicloud_as_configuration

Manages a AS Configuration resource within HuaweiCloud. This is an alternative to `huaweicloud_as_configuration_v1`

## Example Usage

### Basic AS Configuration

```hcl
resource "huaweicloud_as_configuration" "my_as_config" {
  scaling_configuration_name = "my_as_config"
  instance_config {
    flavor    = var.flavor
    image     = var.image_id
    disk {
      size        = 40
      volume_type = "SSD"
      disk_type   = "SYS"
    }
    key_name  = var.keyname
    user_data = file("userdata.txt")
  }
}
```

### AS Configuration With Encrypted Data Disk

```hcl
resource "huaweicloud_as_configuration" "my_as_config" {
  scaling_configuration_name = "my_as_config"
  instance_config {
    flavor    = var.flavor
    image     = var.image_id
    disk {
      size        = 40
      volume_type = "SSD"
      disk_type   = "SYS"
    }
    disk {
      size        = 100
      volume_type = "SSD"
      disk_type   = "DATA"
      kms_id      = var.kms_id
    }
    key_name  = var.keyname
    user_data = file("userdata.txt")
  }
}
```

### AS Configuration With User Data and Metadata

```hcl
resource "huaweicloud_as_configuration" "my_as_config" {
  scaling_configuration_name = "my_as_config"
  instance_config {
    flavor    = var.flavor
    image     = var.image_id
    disk {
      size        = 40
      volume_type = "SSD"
      disk_type   = "SYS"
    }
    key_name  = var.keyname
    user_data = file("userdata.txt")
    metadata  = {
      some_key = "some_value"
    }
  }
}
```

### AS Configuration uses the existing instance specifications as the template

```hcl
resource "huaweicloud_as_configuration" "my_as_config" {
  scaling_configuration_name = "my_as_config"
  instance_config {
    instance_id = "4579f2f5-cbe8-425a-8f32-53dcb9d9053a"
    key_name    = var.keyname
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the AS configuration. If omitted, the `region`
  argument of the provider is used. Changing this creates a new AS configuration.

* `scaling_configuration_name` - (Required, String, ForceNew) The name of the AS configuration. The name can contain
  letters, digits, underscores(_), and hyphens(-), and cannot exceed 64 characters.

* `instance_config` - (Required, List, ForceNew) The information about instance configurations. The instance_config
  dictionary data structure is documented below.

The `instance_config` block supports:

* `instance_id` - (Optional, String) When using the existing instance specifications as the template to create AS
  configurations, specify this argument. In this case, flavor, image, and disk arguments do not take effect. If the
  instance_id argument is not specified, flavor, image, and disk arguments are mandatory.

* `flavor` - (Required, String) The flavor name.

* `image` - (Optional, String) The image ID.

* `disk` - (Optional, List) The disk group information. System disks are mandatory and data disks are optional. The dick
  structure is described below.

* `key_name` - (Required, String) The name of the SSH key pair used to log in to the instance.

* `user_data` - (Optional, String, ForceNew) The user data to provide when launching the instance. The file content must
  be encoded with Base64.

* `personality` - (Optional, List) Customize the personality of an instance by defining one or more files and their
  contents. The personality structure is described below.

* `public_ip` - (Optional, List) The elastic IP address of the instance. The public_ip structure is described below.

* `metadata` - (Optional, Map) Metadata key/value pairs to make available from within the instance.

The `disk` block supports:

* `size` - (Required, Int) The disk size. The unit is GB. The system disk size ranges from 40 to 32768, and the data
  disk size ranges from 10 to 32768.

* `volume_type` - (Required, String) The disk type, which must be the same as the disk type available in the system.
  Available options are:
  + `SAS`: high I/O disk type.
  + `SSD`: ultra-high I/O disk type.
  + `GPSSD`: general purpose SSD disk type.

* `disk_type` - (Required, String) Whether the disk is a system disk or a data disk. Option `DATA` indicates a data
  disk, option `SYS` indicates a system disk.

* `kms_id` - (Optional, String, ForceNew) The Encryption KMS ID of the data disk.

The `personality` block supports:

* `path` - (Required, String) The absolute path of the destination file.

* `contents` - (Required, String) The content of the injected file, which must be encoded with base64.

The `public_ip` block supports:

* `eip` - (Required, List) The configuration parameter for creating an elastic IP address that will be automatically
  assigned to the instance. The eip structure is described below.

The `eip` block supports:

* `ip_type` - (Required, String) The IP address type. The system only supports `5_bgp` (indicates dynamic BGP).

* `bandwidth` - (Required, List) The bandwidth information. The structure is described below.

The `bandwidth` block supports:

* `size` - (Required, Int) The bandwidth (Mbit/s). The value range is 1 to 300.

* `share_type` - (Required, String) The bandwidth sharing type. The system only supports `PER` (indicates exclusive
  bandwidth).

* `charging_mode` - (Required, String) The bandwidth charging mode. The system only supports `traffic`.
