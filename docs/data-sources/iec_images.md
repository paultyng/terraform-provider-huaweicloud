---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_images

Use this data source to get the available of HuaweiCloud IEC images.

## Example Usage

```hcl
data "huaweicloud_iec_images" "iec_image" {
  os_type = "Linux"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) The region in which to obtain the images. If omitted, the provider-level region will be
  used.

* `name` - (Optional, String) Specifies the image Name, which can be queried with a regular expression.

* `os_type` - (Optional, String) Specifies the os type of the iec image.
  "Linux", "Windows" and "Other" are supported.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `images` - An array of one or more image. The images object structure is documented below.

The `images` block supports:

* `id` - The id of the iec images.
* `name` - The name of the iec images.
* `status` - The status of the iec images.
* `os_type` - The os_type of the iec images.
