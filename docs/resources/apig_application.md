---
subcategory: "API Gateway (Dedicated APIG)"
---

# huaweicloud_apig_application

Manages an APIG application resource within HuaweiCloud.

## Example Usage

```hcl
variable "instance_id" {}
variable "app_name" {}
variable "app_code" {}

resource "huaweicloud_apig_application" "test" {
  instance_id = var.instance_id
  name        = var.app_name
  description = "Created by script"

  app_codes = [var.app_code]
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the APIG application resource. If
  omitted, the provider-level region will be used. Changing this will create a new APIG application resource.

* `instance_id` - (Required, String, ForceNew) Specifies an ID of the APIG dedicated instance to which the APIG
  application belongs to. Changing this will create a new APIG application resource.

* `name` - (Required, String) Specifies the name of the API application. The API group name consists of 3 to 64
  characters, starting with a letter. Only letters, digits and underscores (_) are allowed. Chinese characters must be
  in UTF-8 or Unicode format.

* `description` - (Optional, String) Specifies the description about the APIG application. The description contain a
  maximum of 255 characters and the angle brackets (< and >) are not allowed. Chinese characters must be in UTF-8 or
  Unicode format.

* `app_codes` - (Required, List) Specifies an array of one or more application codes which the APIG application belongs
  to. Up to five application codes can be created. The code consists of 64 to 180 characters, starting with a letter,
  digit, plus sign (+) or slash (/). Only letters, digits and following special special characters are allowed: !@#$%+-_
  /=

* `secret_action` - (Optional, String) Specifies the secret action to be done for the APIG application. The valid action
  is *RESET*.

  -> **NOTE:** The `secret_action` is a one-time action.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the APIG application.
* `registraion_time` - Registration time, in RFC-3339 format.
* `update_time` - Time when the API group was last modified, in RFC-3339 format.
* `app_key` - App key.
* `app_secret` - App secret.

## Import

APIG Applications can be imported using their `id` and ID of the APIG dedicated instance to which the application
belongs, separated by a slash, e.g.

```
$ terraform import huaweicloud_apig_application.test <instance id>/<id>
```
