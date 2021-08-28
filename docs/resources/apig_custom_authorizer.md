---
subcategory: "API Gateway (Dedicated APIG)"
---

# huaweicloud_apig_custom_authorizer

Manages an APIG custom authorizer resource within HuaweiCloud.

## Example Usage

```hcl
variable "instance_id" {}
variable "authorizer_name" {}
variable "function_urn" {}

resource "huaweicloud_apig_custom_authorizer" "test" {
  instance_id  = var.instance_id
  name         = var.authorizer_name
  function_urn = var.function_urn
  type         = "FRONTEND"
  cache_age    = 60

  identity {
    name     = "user_name"
    location = "QUERY"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the custom authorizer resource.
  If omitted, the provider-level region will be used.
  Changing this will create a new custom authorizer resource.

* `instance_id` - (Required, String, ForceNew) Specifies an ID of the APIG dedicated instance to which the
  custom authorizer belongs to.
  Changing this will create a new custom authorizer resource.

* `name` - (Required, String, ForceNew) Specifies the name of the custom authorizer.
  The custom authorizer name consists of 3 to 64 characters, starting with a letter.
  Only letters, digits and underscores (_) are allowed.
  Changing this will create a new custom authorizer resource.

* `type` - (Required, String, ForceNew) Specifies the custom authoriz type.
  The valid values are *FRONTEND* and *BACKEND*.
  Changing this will create a new custom authorizer resource.

* `function_urn` - (Required, String, ForceNew) Specifies the uniform function URN of the function graph resource.
  Changing this will create a new custom authorizer resource.

* `is_body_send` - (Optional, Bool, ForceNew) Specifies whether to send the body.
  Changing this will create a new custom authorizer resource.

* `cache_age` - (Optional, String) Specifies the maximum cache age.
  Changing this will create a new custom authorizer resource.

* `user_data` - (Optional, String, ForceNew) Specifies the user data, which can contain a maximum of 2,048 characters.
  The user data is used by APIG to invoke the specified authentication function when accessing the backend service.
  Changing this will create a new custom authorizer resource.

  -> **NOTE:** The user data will be displayed in plain text on the console.

* `identity` - (Optional, List) Specifies an array of one or more parameter identities of the custom authorizer.
  The object structure is documented below.

The `identity` block supports:

* `name` - (Required, String, ForceNew) Specifies the name of the parameter to be verified.
  The parameter includes front-end and back-end parameters.
  Changing this will create a new custom authorizer resource.

* `location` - (Required, String, ForceNew) Specifies the parameter location, which support 'HEADER' and 'QUERY'.
  Changing this will create a new custom authorizer resource.

* `validation` - (Required, String, ForceNew) Specifies the parameter verification expression.
  If omitted, the custom authorizer will not perform verification.
  The valid value is range form 1 to 2,048.
  Changing this will create a new custom authorizer resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the custom authorizer.
* `create_time` - Time when the APIG custom authorizer was created.

## Import

Custom Authorizers of the APIG can be imported using their `name` and the ID of the APIG instance to which the group belongs,
separated by a slash, e.g.

```
$ terraform import huaweicloud_apig_custom_authorizer.test <instance id>/<name>
```
