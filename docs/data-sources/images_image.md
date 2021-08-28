---
subcategory: "Image Management Service (IMS)"
---

# huaweicloud_images_image

Use this data source to get the ID of an available HuaweiCloud image. This is an alternative
to `huaweicloud_images_image_v2`

## Example Usage

```hcl
data "huaweicloud_images_image" "ubuntu" {
  name        = "Ubuntu 18.04 server 64bit"
  visibility  = "public"
  most_recent = true
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the images. If omitted, the provider-level region will be
  used.

* `most_recent` - (Optional, Bool) If more than one result is returned, use the most recent image.

* `name` - (Optional, String) The name of the image.

* `owner` - (Optional, String) The owner (UUID) of the image.

* `size_min` - (Optional, Int) The minimum size (in bytes) of the image to return.

* `size_max` - (Optional, Int) The maximum size (in bytes) of the image to return.

* `sort_direction` - (Optional, String) Order the results in either `asc` or `desc`.

* `sort_key` - (Optional, String) Sort images based on a certain key. Must be one of
  "name", "container_format", "disk_format", "status", "id" or "size". Defaults to `name`.

* `tag` - (Optional, String) Search for images with a specific tag.

* `visibility` - (Optional, String) The visibility of the image. Must be one of
  "public", "private", "community", or "shared". Defaults to `private`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.
* `checksum` - The checksum of the data associated with the image.
* `container_format` - The format of the image's container.
* `disk_format` - The format of the image's disk.
* `file` - the trailing path after the glance endpoint that represent the location of the image or the path to retrieve
  it.
* `metadata` - The metadata associated with the image. Image metadata allow for meaningfully define the image properties
  and tags.
* `min_disk_gb` - The minimum amount of disk space required to use the image.
* `min_ram_mb` - The minimum amount of ram required to use the image.
* `properties` - Freeform information about the image.
* `protected` - Whether or not the image is protected.
* `schema` - The path to the JSON-schema that represent the image or image.
* `size_bytes` - The size of the image (in bytes).
* `status` - The status of the image.
* `created_at` - The date when the image was created.
* `update_at` - The date when the image was last updated.
