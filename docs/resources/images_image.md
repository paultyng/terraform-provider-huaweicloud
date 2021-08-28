---
subcategory: "Image Management Service (IMS)"
---

# huaweicloud_images_image

Manages an Image resource within HuaweiCloud IMS.

## Example Usage

### Creating an image from ECS

```hcl
data "huaweicloud_availability_zones" "test" {}

data "huaweicloud_vpc_subnet" "test" {
  name = "subnet-default"
}

resource "huaweicloud_compute_instance" "test" {
  name              = "%s"
  image_name        = "Ubuntu 18.04 server 64bit"
  security_groups   = [
    "default"]
  availability_zone = data.huaweicloud_availability_zones.test.names[0]

  network {
    uuid = data.huaweicloud_vpc_subnet.test.id
  }
}

resource "huaweicloud_images_image" "test" {
  name        = "%s"
  instance_id = huaweicloud_compute_instance.test.id
  description = "created by Terraform"

  tags = {
    foo = "bar"
    key = "value"
  }
}
```

### Creating an image from OBS bucket

```hcl
resource "huaweicloud_images_image" "ims_test_file" {
  name        = "ims_test_file"
  image_url   = "ims-image:centos70.qcow2"
  min_disk    = 40
  description = "Create an image from the OBS bucket."

  tags = {
    foo = "bar1"
    key = "value"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) The name of the image.

* `description` - (Optional, String, ForceNew) A description of the image.

* `min_ram` - (Optional, Int, ForceNew) The minimum memory of the image in the unit of MB. The default value is 0,
  indicating that the memory is not restricted.

* `max_ram` - (Optional, Int, ForceNew) The maximum memory of the image in the unit of MB.

* `tags` - (Optional, Map) The tags of the image.

* `instance_id` - (Optional, String, ForceNew) The ID of the ECS that needs to be converted into an image. This
  parameter is mandatory when you create a privete image from an ECS.

* `image_url` - (Optional, String, ForceNew) The URL of the external image file in the OBS bucket. This parameter is
  mandatory when you create a private image from an external file uploaded to an OBS bucket. The format is *OBS bucket
  name:Image file name*.

* `min_disk` - (Optional, Int, ForceNew) The minimum size of the system disk in the unit of GB. This parameter is
  mandatory when you create a private image from an external file uploaded to an OBS bucket. The value ranges from 1 GB
  to 1024 GB.

* `os_version` - (Optional, String, ForceNew) The OS version. This parameter is valid when you create a private image
  from an external file uploaded to an OBS bucket.

* `is_config` - (Optional, Bool, ForceNew) If automatic configuration is required, set the value to true. Otherwise, set
  the value to false.

* `cmk_id` - (Optional, String, ForceNew) The master key used for encrypting an image.

* `type` - (Optional, String, ForceNew) The image type. Must be one of `ECS`, `FusionCompute`, `BMS`, or `Ironic`.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id of the image. Changing this creates a
  new image.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - A unique ID assigned by IMS.

* `visibility` - Whether the image is visible to other tenants.

* `data_origin` - The image resource. The pattern can be 'instance,*instance_id*' or 'file,*image_url*'.

* `disk_format` - The image file format. The value can be `vhd`, `zvhd`, `raw`, `zvhd2`, or `qcow2`.

* `image_size` - The size(bytes) of the image file format.

* `checksum` - The checksum of the data associated with the image.

* `status` - The status of the image.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 3 minute.

## Import

Images can be imported using the `id`, e.g.

```sh
terraform import huaweicloud_images_image.my_image 7886e623-f1b3-473e-b882-67ba1c35887f
```
