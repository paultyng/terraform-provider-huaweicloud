---
subcategory: "API Gateway (Dedicated APIG)"
---

# huaweicloud_apig_response

Manages an APIG (API) custom response resource within HuaweiCloud.

## Example Usage

```hcl
variable "instance_id" {}
variable "group_id" {}
variable "response_name" {}

resource "huaweicloud_apig_response" "test" {
  name        = var.response_name
  instance_id = var.instance_id
  group_id    = var.group_id

  rule {
    error_type  = "AUTHORIZER_FAILURE"
    body        = "{\"code\":\"$context.authorizer.frontend.code\",\"message\":\"$context.authorizer.frontend.message\"}"
    status_code = 401
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the API custom response resource. If
  omitted, the provider-level region will be used. Changing this will create a new API custom response resource.

* `group_id` - (Required, String, ForceNew) Specifies the ID of the API group to which the API response belongs to.
  Changing this will create a new API custom response resource.

* `instance_id` - (Required, String, ForceNew) Specifies the ID of the APIG dedicated instance to which the API group
  where the API custom response belongs. Changing this will create a new API custom response resource.

* `name` - (Required, String) Specifies the name of the API custom response. The name consists of 1 to 64 characters,
  and only letters, digits, hyphens(-), and underscores (_) are allowed.

* `rule` - (Optional, List) Specifies the API custom response rules definition. The object structure is documented
  below.

The `rule` block supports:

* `error_type` - (Required, String) Specifies the type of the API custom response rule.
  + **AUTH_FAILURE**: Authentication failed.
  + **AUTH_HEADER_MISSING**: The identity source is missing.
  + **AUTHORIZER_FAILURE**: Custom authentication failed.
  + **AUTHORIZER_CONF_FAILURE**: There has been a custom authorizer error.
  + **AUTHORIZER_IDENTITIES_FAILURE**: The identity source of the custom authorizer is invalid.
  + **BACKEND_UNAVAILABLE**: The backend service is unavailable.
  + **BACKEND_TIMEOUT**: Communication with the backend service timed out.
  + **THROTTLED**: The request was rejected due to request throttling.
  + **UNAUTHORIZED**: The app you are using has not been authorized to call the API.
  + **ACCESS_DENIED**: Access denied.
  + **NOT_FOUND**: No API is found.
  + **REQUEST_PARAMETERS_FAILURE**: The request parameters are incorrect.
  + **DEFAULT_4XX**: Another 4XX error occurred.
  + **DEFAULT_5XX**: Another 5XX error occurred.

* `body` - (Required, String) Specifies the body template of the API response rule, e.g.
  `{\"code\":\"$context.authorizer.frontend.code\",\"message\":\"$context.authorizer.frontend.message\"}`

* `status_code` - (Optional, Int) Specifies the HTTP status code of the API response rule.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the API custom response.
* `create_time` - Time when the API custom response is created.
* `update_time` - Time when the API custom response was last modified.

## Import

API Responses can be imported using their `name` and IDs of the APIG dedicated instances and API groups to which the API
response belongs, separated by a slash, e.g.

```
$ terraform import huaweicloud_apig_response.test <instance id>/<group id>/<name>
```
