---
subcategory: "API Gateway (Dedicated APIG)"
---

# huaweicloud_apig_api

Manages an APIG API resource within HuaweiCloud.

## Example Usage

```hcl
variable "instance_id" {}
variable "group_id" {}
variable "api_name" {}
variable "custom_response_id" {}
variable "custom_auth_id" {}
variable "vpc_channel_id" {}

resource "huaweicloud_apig_api" "test" {
  instance_id             = var.instance_id
  group_id                = var.group_id
  type                    = "Public"
  name                    = var.api_name
  request_protocol        = "HTTP"
  request_method          = "POST"
  request_path            = "/terraform/users"
  security_authentication = "AUTHORIZER"
  matching                = "Exact"
  success_response        = "Successful"
  response_id             = var.custom_response_id
  authorizer_id           = var.custom_auth_id

  backend_params {
    type     = "SYSTEM"
    name     = "X-User-Auth"
    location = "HEADER"
    value    = "user_name"
  }

  web {
    path             = "/backend/users"
    vpc_channel_id   = var.vpc_channel_id
    request_method   = "POST"
    request_protocol = "HTTP"
    timeout          = 5000
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the API resource. If omitted, the
  provider-level region will be used. Changing this will create a new API resource.

* `instance_id` - (Required, String, ForceNew) Specifies an ID of the APIG dedicated instance to which the API belongs
  to. Changing this will create a new API resource.

* `group_id` - (Required, String) Specifies an ID of the APIG group to which the API belongs to.

* `type` - (Required, String) Specifies the API type. The valid values are __Public__ and __Private__.

* `name` - (Required, String) Specifies the API name, which can consists of 3 to 64 characters, starting with a letter.
  Only letters, digits and underscores (_) are allowed. Chinese characters must be in UTF-8 or Unicode format.

* `request_method` - (Required, String) Specifies the request method of the API. The valid values are __GET__, __POST__
  , __PUT__, __DELETE__, __HEAD__, __PATCH__, __OPTIONS__ and __ANY__.

* `request_path` - (Required, String) Specifies the request address, which can contain a maximum of 512 characters
  request parameters enclosed with brackets ({}).
  + The address can contain special characters, such as asterisks (), percent signs (%), hyphens (-), and
      underscores (_) and must comply with URI specifications.
  + The address can contain environment variables, each starting with a letter and consisting of 3 to 32 characters.
      Only letters, digits, hyphens (-), and underscores (_) are allowed in environment variables.

* `request_protocol` - (Required, String) Specifies the request protocol of the API. The valid value are
  __HTTP__, __HTTPS__ and __BOTH__.

* `request_params` - (Optional, List) Specifies an array of one or more request parameters of the front-end. The maximum
  of request parameters is 50. The [object](#request_params) structure is documented below.

* `backend_params` - (Optional, List) Specifies an array of one or more backend parameters.
  The [object](#backend_params) structure is documented below. The maximum of request parameters is 50.

* `security_authentication` - (Optional, String) Specifies the security authentication mode. The valid values are
  __NONE__, __APP__ and __IAM__, default to __NONE__.

* `simple_authentication` - (Optional, Bool) Specifies whether AppCode authentication is enabled. The applicaiton code
  must located in the header when `simple_authentication` is true.

* `authorizer_id` - (Optional, String) Specifies ID of the front-end custom authorizer.

* `body_description` - (Optional, String) Specifies the description of the API request body, which can be an example
  request body, media type or parameters. The request body does not exceed 20,480 characters. Chinese characters must be
  in UTF-8 or Unicode format.

* `cors` - (Optional, Bool) Specifies whether CORS is supported, default to false.

* `description` - (Optional, String) Specifies the API description, which can contain a maximum of 255 characters. The
  Chinese characters must be in UTF-8 or Unicode format.

* `matching` - (Optional, String) Specifies the route matching mode. The valid value are __Exact__ and __Prefix__,
  default to __Exact__.

* `response_id` - (Optional, String) Specifies the APIG group response ID.

* `success_response` - (Optional, String) Specifies the example response for a successful request. Ensure that the
  response does not exceed 20,480 characters. Chinese characters must be in UTF-8 or Unicode format.

* `failure_response` - (Optional, String) Specifies the example response for a successful request. Ensure that the
  response does not exceed 20,480 characters. Chinese characters must be in UTF-8 or Unicode format.

* `mock` - (Optional, List, ForceNew) Specifies the mock backend details. The [object](#mock) structure is documented
  below. Changing this will create a new API resource.

* `func_graph` - (Optional, List, ForceNew) Specifies the function graph backend details. The [object](#func_graph)
  structure is documented below. Changing this will create a new API resource.

* `web` - (Optional, List, ForceNew) Specifies the web backend details. The [object](#web) structure is documented
  below. Changing this will create a new API resource.

* `mock_policy` - (Optional, List) Specifies the Mock policy backends. The maximum of the policy is 5.
  The [object](#mock_policy) structure is documented below.

* `func_graph_policy` - (Optional, List) Specifies the Mock policy backends. The maximum of the policy is 5.
  The [object](#func_graph_policy) structure is documented below.

* `web_policy` - (Optional, List) Specifies the example response for a failed request. The maximum of the policy is 5.
  The [object](#web_policy) structure is documented below.

The <span id="request_params">`request_params`</span> block supports:

* `name` - (Required, String) Specifies the request parameter name, which contain of 1 to 32 characters and start with a
  letter. Only letters, digits, hyphens (-), underscores (_) and periods (.) are allowed. If Location is specified as
  __HEADER__ and `security_authentication` is specified as __APP__, the parameter name is not 'Authorization' (
  case-insensitive) and cannot contain underscores.

* `required` - (Required, Bool) Specifies whether the request parameter is required.

* `location` - (Optional, String) Specifies the location of the request parameter. The valid values are __PATH__,
  __QUERY__ and __HEADER__, default to __PATH__.

* `type` - (Optional, String) Specifies the request parameter type. The valid values are __STRING__ and __NUMBER__,
  default to __STRING__.

* `maximum` - (Optional, Int) Specifies the maximum value or size of the request parameter.

* `minimum` - (Optional, Int) Specifies the minimum value or size of the request parameter. For string type,
  The `maximum` and `minimum` means size. For number type, they means value.

* `example` - (Optional, String) Specifies the example value of the request parameter, which contain a maximum of 255
  characters, and the angle brackets (< and >) are not allowed.

* `default` - (Optional, String) Specifies the default value of the request parameter, which contain a maximum of 255
  characters, and the angle brackets (< and >) are not allowed.

* `description` - (Optional, String) Specifies the description of the request parameter, which contain a maximum of 255
  characters, and the angle brackets (< and >) are not allowed.

The <span id="backend_params">`backend_params`</span> block supports:

* `type` - (Required, String) Specifies the backend parameter type. The valid values are __REQUEST__, __CONSTANT__
  and __SYSTEM__.

* `name` - (Required, String) Specifies the backend parameter name, which contain of 1 to 32 characters and start with a
  letter. Only letters, digits, hyphens (-), underscores (_) and periods (.) are allowed. The parameter name is not
  case-sensitive. It cannot start with 'x-apig-' or 'x-sdk-' and cannot be 'x-stage'. If the location is specified as
  __HEADER__, the name cannot contain underscores.

* `location` - (Required, String) Specifies the location of the backend parameter. The valid values are __PATH__,
  __QUERY__ and __HEADER__.

* `value` - (Required, String) Specifies the request parameter name corresponding to the request parameter name of the
  back-end parameter.

* `description` - (Optional, String) Specifies the description of the constant or system parameter, which contain a
  maximum of 255 characters, and the angle brackets (< and >) are not allowed.

The <span id="mock">`mock`</span> block supports:

* `response` - (Required, String) Specifies the response of the backend policy, which contain a maximum of 2,048
  characters, and the angle brackets (< and >) are not allowed.

  -> **NOTE:**  Mock enables APIG to return a response without sending the request to the backend. This is useful for
  testing APIs when the backend is not available.

* `authorizer_id` - (Optional, String) Specifies the ID of the backend custom authorization.

The <span id="func_graph">`func_graph`</span> block supports:

* `function_urn` - (Required, String) Specifies the function graph URN.

* `version` - (Required, String) Specifies the version of the function graph.

* `timeout` - (Optional, Int) Specifies the location of the backend parameter. The valid value is range form 1 to
  600,000, default to 5,000.

* `invocation_type` - (Optional, String) Specifies the invocation mode. The valid values are __async__ and __sync__,
  default to __sync__.

* `authorizer_id` - (Optional, String) Specifies the ID of the backend custom authorization.

The <span id="web">`web`</span> block supports:

* `path` - (Required, String) Specifies the backend request address, which can contain a maximum of 512 characters and
  must comply with URI specifications.
  + The request address can contain request parameters enclosed with brackets ({}).
  + The request address can contain special characters, such as asterisks (*), percent signs (%), hyphens (-) and
      underscores (_) and must comply with URI specifications.
  + The address can contain environment variables, each starting with a letter and consisting of 3 to 32 characters.
      Only letters, digits, hyphens (-), and underscores (_) are allowed in environment variables.

* `host_header` - (Optional, String) Specifies the proxy host header. The host header can be customized for requests to
  be forwarded to cloud servers through the VPC channel. By default, the original host header of the request is used.

* `vpc_channel_id` - (Optional, String) Specifies the VPC channel ID. This parameter and `backend_address` are
  alternative.

* `backend_address` - (Optional, String) Specifies the backend service address, which consists of a domain name or IP
  address, and a port number, with not more than 255 characters. The backend service address must be in the format "Host
  name:Port number", for example, apig.example.com:7443. If the port number is not specified, the default HTTPS port
  443, or the default HTTP port 80 is used. The backend service address can contain environment variables, each starting
  with a letter and consisting of 3 to 32 characters. Only letters, digits, hyphens (-), and underscores (_) are
  allowed.

* `request_method` - (Optional, String) Specifies the backend request method of the API. The valid types are __GET__,
  __POST__, __PUT__, __DELETE__, __HEAD__, __PATCH__, __OPTIONS__ and __ANY__.

* `request_protocol` - (Optional, String) Specifies the backend request protocol. The valid values are __HTTP__ and
  __HTTPS__, default to __HTTPS__.

* `timeout` - (Optional, Int) Specifies the timeout, in ms, which allowed for APIG to request the backend service. The
  valid value is range from 1 to 600,000, default to 5,000.

* `ssl_enable` - (Optional, Bool) Specifies the indicates whether to enable two-way authentication, default to false.

* `authorizer_id` - (Optional, String) Specifies the ID of the backend custom authorization.

The <span id="mock_policy">`mock_policy`</span> block supports:

* `name` - (Required, String) Specifies the backend policy name, which can contains of 3 to 64 characters and start with
  a letter. Only letters, digits, and underscores (_) are allowed.

* `conditions` - (Required, List) Specifies an array of one or more policy conditions. Up to five conditions can be set.
  The [object](#conditions) structure is documented below.

* `response` - (Optional, String) Specifies the response of the backend policy, which contain a maximum of 2,048
  characters, and the angle brackets (< and >) are not allowed.

* `effective_mode` - (Optional, String) Specifies the effective mode of the backend policy. The valid values are __ALL__
  and __ANY__, default to __ANY__.

* `backend_params` - (Optional, List) Specifies an array of one or more backend parameters. The maximum of request
  parameters is 50. The [object](#backend_params) structure is documented above.

* `authorizer_id` - (Optional, String) Specifies the ID of the backend custom authorization.

The <span id="func_graph_policy">`func_graph_policy`</span> block supports:

* `name` - (Required, String) Specifies the backend policy name, which can contains of 3 to 64 characters and start with
  a letter. Only letters, digits, and underscores (_) are allowed.

* `function_urn` - (Required, String) Specifies the URN of the function graph.

* `conditions` - (Required, List) Specifies an array of one or more policy conditions. Up to five conditions can be set.
  The [object](#conditions) structure is documented below.

* `invocation_mode` - (Optional, String) Specifies the invocation mode of the function graph. The valid values are
  __async__ and __sync__, default to __sync__.

* `effective_mode` - (Optional, String) Specifies the effective mode of the backend policy. The valid values are __ALL__
  and __ANY__, default to __ANY__.

* `timeout` - (Optional, Int) Specifies the timeout, in ms, which allowed for APIG to request the backend service. The
  valid value is range from 1 to 600,000, default to 5,000.

* `version` - (Optional, String) Specifies the version of the function graph.

* `backend_params` - (Optional, List) Specifies an array of one or more backend parameters. The maximum of request
  parameters is 50. The [object](#backend_params) structure is documented above.

* `authorizer_id` - (Optional, String) Specifies the ID of the backend custom authorization.

The <span id="web_policy">`web_policy`</span> block supports:

* `name` - (Required, String) Specifies the backend policy name, which can contains of 3 to 64 characters and start with
  a letter. Only letters, digits, and underscores (_) are allowed.

* `path` - (Required, String) Specifies the backend request address, which can contain a maximum of 512 characters and
  must comply with URI specifications.
  + The request address can contain request parameters enclosed with brackets ({}).
  + The request address can contain special characters, such as asterisks (*), percent signs (%), hyphens (-) and
      underscores (_) and must comply with URI specifications.
  + The address can contain environment variables, each starting with a letter and consisting of 3 to 32 characters.
      Only letters, digits, hyphens (-), and underscores (_) are allowed in environment variables.

* `request_method` - (Required, String) Specifies the backend request method of the API. The valid types are __GET__,
  __POST__, __PUT__, __DELETE__, __HEAD__, __PATCH__, __OPTIONS__ and __ANY__.

* `conditions` - (Required, List) Specifies an array of one or more policy conditions. Up to five conditions can be set.
  The [object](#conditions) structure is documented below.

* `host_header` - (Optional, String) Specifies the proxy host header. The host header can be customized for requests to
  be forwarded to cloud servers through the VPC channel. By default, the original host header of the request is used.

* `vpc_channel_id` - (Optional, String) Specifies the VPC channel ID. This parameter and `backend_address` are
  alternative.

* `backend_address` - (Optional, String) Specifies the backend service address, which consists of a domain name or IP
  address, and a port number, with not more than 255 characters. The backend service address must be in the format "Host
  name:Port number", for example, apig.example.com:7443. If the port number is not specified, the default HTTPS port 443
  or the default HTTP port 80 is used. The backend service address can contain environment variables, each starting with
  a letter and consisting of 3 to 32 characters. Only letters, digits, hyphens (-), and underscores (_) are allowed.

* `request_protocol` - (Optional, String) Specifies the backend request protocol. The valid values are __HTTP__ and
  __HTTPS__, default to __HTTPS__.

* `effective_mode` - (Optional, String) Specifies the effective mode of the backend policy. The valid values are __ALL__
  and __ANY__, default to __ANY__.

* `timeout` - (Optional, Int) Specifies the timeout, in ms, which allowed for APIG to request the backend service. The
  valid value is range from 1 to 600,000, default to 5,000.

* `backend_params` - (Optional, List) Specifies an array of one or more backend parameters. The maximum of request
  parameters is 50. The [object](#backend_params) structure is documented above.

* `authorizer_id` - (Optional, String) Specifies the ID of the backend custom authorization.

The <span id="conditions">`conditions`</span> block supports:

* `value` - (Required, String) Specifies the condition type. For a condition with the input parameter source:
  + If the condition type is __Enumerated__, separate condition values with commas.
  + If the condition type is __Matching__, enter a regular expression compatible with PERL.

  For a condition with the Source IP address source, enter IPv4 addresses and separate them with commas. The CIDR
  address format is supported.

* `param_name` - (Optional, String) Specifies the request parameter name. This parameter is required if the policy type
  is param.

* `source` - (Optional, String) Specifies the policy type. The valid values are __param__ and __source__, default to
  __source__.

* `type` - (Optional, String) Specifies the condition type of the backend policy. The valid values are __Equal__,
  __Enumerated__ and __Matching__, default to __Equal__.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the APIG API.
* `register_time` - Time when the API is registered, in UTC format.
* `update_time` - Time when the API was last modified, in UTC format.

## Import

APIs can be imported using their `name` and ID of the APIG dedicated instance to which the API belongs, separated by a
slash, e.g.

```
$ terraform import huaweicloud_apig_api.test <instance_id>/<name>
```
