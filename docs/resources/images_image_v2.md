---
subcategory: "Deprecated"
---

# huaweicloud_images_image_v2

Manages a Image resource within HuaweiCloud IMS.

## Example Usage

```hcl
resource "huaweicloud_images_image_v2" "rancheros" {
  name             = "RancherOS"
  image_source_url = "https://releases.rancher.com/os/latest/rancheros-openstack.img"
  container_format = "bare"
  disk_format      = "qcow2"
  tags             = ["foo.bar", "tag.value"]
}
```

## Argument Reference

The following arguments are supported:

* `container_format` - (Required) The container format. Must be "bare".

* `disk_format` - (Required) The disk format. Must be one of "qcow2", "vhd".

* `local_file_path` - (Optional) This is the filepath of the raw image file
   that will be uploaded to Glance. Conflicts with `image_source_url`.

* `image_cache_path` - (Optional) This is the directory where the images will
   be downloaded. Images will be stored with a filename corresponding to
   the url's md5 hash. Defaults to "$HOME/.terraform/image_cache"

* `image_source_url` - (Optional) This is the url of the raw image that will
   be downloaded in the `image_cache_path` before being uploaded to Glance.
   Glance is able to download image from internet but the `golangsdk` library
   does not yet provide a way to do so.
   Conflicts with `local_file_path`.

* `min_disk_gb` - (Optional) Amount of disk space (in GB) required to boot image.
   Defaults to 0.

* `min_ram_mb` - (Optional) Amount of ram (in MB) required to boot image.
   Defauts to 0.

* `name` - (Required) The name of the image.

* `protected` - (Optional) If true, image will not be deletable.
   Defaults to false.

* `region` - (Optional) The region in which to create the V2 Glance client.
    A Glance client is needed to create an Image that can be used with
    a compute instance. If omitted, the `region` argument of the provider
    is used. Changing this creates a new Image.

* `tags` - (Optional) The tags of the image. It must be a list of strings.
    At this time, it is not possible to delete all tags of an image.

* `visibility` - (Optional) The visibility of the image. Must be "private".
   The ability to set the visibility depends upon the configuration of
   the HuaweiCloud cloud.

Note: The `properties` attribute handling in the golangsdk library is currently buggy
and needs to be fixed before being implemented in this resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `checksum` - The checksum of the data associated with the image.
* `created_at` - The date the image was created.
* `file` - the trailing path after the glance
   endpoint that represent the location of the image
   or the path to retrieve it.
* `id` - A unique ID assigned by Glance.
* `metadata` - The metadata associated with the image.
   Image metadata allow for meaningfully define the image properties
   and tags. See [metadata reference](http://docs.openstack.org/developer/glance/metadefs-concepts.html).
* `owner` - The id of the huaweicloud user who owns the image.
* `schema` - The path to the JSON-schema that represent
   the image or image
* `size_bytes` - The size in bytes of the data associated with the image.
* `status` - The status of the image. It can be "queued", "active"
   or "saving".
* `update_at` - The date the image was last updated.


## Timeouts
This resource provides the following timeouts configuration options:
* `create` - Default is 30 minute.

## Import

Images can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_images_image_v2.rancheros 89c60255-9bd6-460c-822a-e2b959ede9d2
```
