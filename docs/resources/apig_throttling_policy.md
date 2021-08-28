---
subcategory: "API Gateway (Dedicated APIG)"
---

# huaweicloud_apig_throttling_policy

Manages an APIG (API) throttling policy resource within HuaweiCloud.

## Example Usage

### Create a basic throttling policy

```hcl
variable "instance_id" {}
variable "policy_name" {}
variable "description" {}

resource "huaweicloud_apig_throttling_policy" "test" {
  instance_id       = var.instance_id
  name              = var.policy_name
  description       = var.description
  type              = "API-based"
  period            = 10
  period_unit       = "MINUTE"
  max_api_requests  = 70
  max_user_requests = 45
  max_app_requests  = 45
  max_ip_requests   = 45
}
```

### Create a throttling policy with a special throttle

```hcl
variable "instance_id" {}
variable "policy_name" {}
variable "description" {}
variable "application_id" {}

resource "huaweicloud_apig_throttling_policy" "test" {
  instance_id       = var.instance_id
  name              = var.policy_name
  description       = var.description
  type              = "API-based"
  period            = 10
  period_unit       = "MINUTE"
  max_api_requests  = 70
  max_user_requests = 45
  max_app_requests  = 45
  max_ip_requests   = 45

  app_throttles {
    max_api_requests     = 40
    throttling_object_id = var.application_id
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the API throttling policy resource. If
  omitted, the provider-level region will be used. Changing this will create a new API throttling policy resource.

* `instance_id` - (Required, String, ForceNew) Specifies an ID of the APIG dedicated instance to which the API
  throttling policy belongs to. Changing this will create a new API throttling policy resource.

* `name` - (Required, String) Specifies the name of the API throttling policy. The policy name consists of 3 to 64
  characters, starting with a letter. Only letters, digits and underscores (_) are allowed.

* `period` - (Required, Int) Specifies the period of time for limiting the number of API calls. This parameter applies
  with each of the API call limits: `max_api_requests`, `max_app_requests`, `max_ip_requests`
  and `max_user_requests`.

* `max_api_requests` - (Required, Int) Specifies the maximum number of times an API can be accessed within a specified
  period. The value of this parameter cannot exceed the default limit 200 TPS.

* `max_app_requests` - (Optional, Int) Specifies the maximum number of times the API can be accessed by an app within
  the same period. The value of this parameter must be less than or equal to the value of `max_user_requests`.

* `max_ip_requests` - (Optional, Int) Specifies the maximum number of times the API can be accessed by an IP address
  within the same period. The value of this parameter must be less than or equal to the value of `max_api_requests`.

* `max_user_requests` - (Optional, Int) Specifies the maximum number of times the API can be accessed by a user within
  the same period. The value of this parameter must be less than or equal to the value of `max_api_requests`.

* `type` - (Optional, String) Specifies the type of the request throttling policy. The valid values are as follows:
  + API-based: limiting the maximum number of times a single API bound to the policy can be called within the
      specified period.
  + API-shared: limiting the maximum number of times all APIs bound to the policy can be called within the specified
      period.

* `description` - (Optional, String) Specifies the description about the API throttling policy. The description contain
  a maximum of 255 characters and the angle brackets (< and >) are not allowed. Chinese characters must be in UTF-8 or
  Unicode format.

* `period_unit` - (Optional, String) Specifies the time unit for limiting the number of API calls. The valid values
  are *SECOND*, *MINUTE*, *HOUR* and *DAY*, default to *MINUTE*.

* `user_throttles` - (Optional, List) Specifies an array of one or more special throttling policies for IAM user limit.
  The `throttle` object of the `user_throttles` structure is documented below.

* `app_throttles` - (Optional, List) Specifies an array of one or more special throttling policies for APP limit.
  The `throttle` object of the `user_throttles` structure is documented below.

The `throttle` block supports:

* `max_api_requests` - (Required, Int) Specifies the maximum number of times an API can be accessed within a specified
  period.

* `throttling_object_id` - (Required, String) Specifies the object ID which the special throttling policy belongs.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the API throttling policy.

* `user_throttles` - An array of one or more special throttling policies for IAM user limit.
  + `throttling_object_name` - The object name which the special user throttling policy belongs.
  + `id` - ID of the special user throttling policy.

* `app_throttles` - An array of one or more special throttling policies for APP limit.
  + `throttling_object_name` - The object name which the special application throttling policy belongs.
  + `id` - ID of the special application throttling policy.

* `create_time` - Time when the API throttling policy was created.

## Import

API Throttling Policies of APIG can be imported using their `name` and the ID of the APIG instances to which the
environment belongs, separated by a slash, e.g.

```
$ terraform import huaweicloud_apig_throttling_policy.test <instance ID>/<name>
```
