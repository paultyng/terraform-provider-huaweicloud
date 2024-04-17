---
subcategory: "Content Delivery Network (CDN)"
---

# huaweicloud_cdn_domain

Manages a CDN domain resource within HuaweiCloud.

## Example Usage

### Create a CDN domain

```hcl
variable "domain_name" {}
variable "origin_server" {}

resource "huaweicloud_cdn_domain" "domain_1" {
  name         = var.domain_name
  type         = "web"
  service_area = "mainland_china"

  sources {
    origin      = var.origin_server
    origin_type = "ipaddr"
    active      = 1
  }

  tags = {
    key = "val"
    foo = "bar"
  }
}
```

### Create a CDN domain with cache rules

```hcl
variable "domain_name" {}
variable "origin_server" {}

resource "huaweicloud_cdn_domain" "domain_1" {
  name         = var.domain_name
  type         = "web"
  service_area = "mainland_china"

  sources {
    origin      = var.origin_server
    origin_type = "ipaddr"
    active      = 1
  }

  cache_settings {
    rules {
      rule_type = "all"
      ttl       = 180
      ttl_type  = "d"
      priority  = 2
    }
  }
}
```

### Create a CDN domain with configs

```hcl
variable "domain_name" {}
variable "origin_server" {}
variable "ip_or_domain" {}

resource "huaweicloud_cdn_domain" "domain_1" {
  name         = var.domain_name
  type         = "web"
  service_area = "mainland_china"

  sources {
    origin      = var.origin_server
    origin_type = "ipaddr"
    active      = 1
  }

  configs {
    origin_protocol               = "http"
    ipv6_enable                   = true
    range_based_retrieval_enabled = true
    description                   = "test description"

    https_settings {
      certificate_name     = "terraform-test"
      certificate_body     = file("your_directory/chain.cer")
      http2_enabled        = true
      https_enabled        = true
      private_key          = file("your_directory/server_private.key")
      ocsp_stapling_status = "on"
    }

    cache_url_parameter_filter {
      type = "ignore_url_params"
    }

    retrieval_request_header {
      name   = "test-name"
      value  = "test-val"
      action = "set"
    }

    http_response_header {
      name   = "test-name"
      value  = "test-val"
      action = "set"
    }

    url_signing {
      enabled     = true
      type        = "type_a"
      sign_method = "md5"
      match_type  = "all"
      sign_arg    = "Psd_123"
      key         = "A27jtfSTy13q7A0UnTA9vpxYXEb"
      backup_key  = "S36klgTFa60q3V8DmSK2hwfBOYp"
      time_format = "dec"
      expire_time = 30

      inherit_config {
        enabled           = true
        inherit_type      = "m3u8"
        inherit_time_type = "sys_time"
      }
    }

    flexible_origin {
      match_type = "all"
      priority   = 1

      back_sources {
        http_port    = 80
        https_port   = 443
        ip_or_domain = var.ip_or_domain
        sources_type = "ipaddr"
      }
    }

    remote_auth {
      enabled = true

      remote_auth_rules {
        auth_failed_status      = "503"
        auth_server             = "https://testdomain-update.com"
        auth_success_status     = "302"
        file_type_setting       = "all"
        request_method          = "POST"
        reserve_args_setting    = "reserve_all_args"
        reserve_headers_setting = "reserve_all_headers"
        response_status         = "206"
        timeout                 = 3000
        timeout_action          = "forbid"

        add_custom_args_rules {
          key   = "http_user_agent"
          type  = "nginx_preset_var"
          value = "$server_protocol"
        }
      }
    }

    compress {
      enabled = false
    }

    force_redirect {
      enabled = true
      type    = "http"
    }

    referer {
      type          = "white"
      value         = "*.common.com,192.187.2.43,www.test.top:4990"
      include_empty = false
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Specifies acceleration domain name. The domain name consists of one or more parts,
  representing domains at different levels. Domain names at all levels can only be composed of letters, digits,
  and hyphens (-), and the letters are equivalent in upper and lower case. Domain names at all levels are connected
  with (.). The domain name can contain up to `75` characters.

* `type` - (Required, String) Specifies the service type of the domain name. The valid values are as follows:
  + **web**: Static acceleration. For websites with many images and small files, such as portals and e-commerce websites.
  + **download**: Download acceleration. For large files, such as apps in app stores and game clients.
  + **video**: Streaming media acceleration. For video on demand (VOD) websites and online education websites.
  + **wholeSite**: Whole site acceleration. For websites with both dynamic and static content, such as online exam
    platforms, forums, and blogs.

  -> Currently, **wholeSite** acceleration cannot be changed to other service types.

* `sources` - (Required, List) Specifies an array of one or more objects specifying origin server settings.
  A maximum of `50` origin site configurations can be configured.
  The [sources](#sources_cdn_domain) structure is documented below.

* `service_area` - (Required, String) Specifies the area covered by the acceleration service.
  Valid values are as follows:
  + **mainland_china**: Indicates that the service scope is mainland China.
  + **outside_mainland_china**: Indicates that the service scope is outside mainland China.
  + **global**: Indicates that the service scope is global.

  -> The service area cannot be changed between Chinese mainland and outside Chinese mainland.

* `configs` - (Optional, List) Specifies the domain configuration items. The [configs](#configs_object) structure is
  documented below.

* `cache_settings` - (Optional, List) Specifies the cache configuration. The [cache_settings](#cache_settings_object) structure
  is documented below.

* `enterprise_project_id` - (Optional, String) Specifies the enterprise project ID.

* `tags` - (Optional, Map) Specifies the key/value pairs to associate with the domain.

<a name="sources_cdn_domain"></a>
The `sources` block supports:

* `origin` - (Required, String) Specifies the unique domain name or IP address of the origin server.
  + If `origin_type` is set to **ipaddr**, this field can only be set to IPv4 address.
  + If `origin_type` is set to **domain**, this field can only be set to domain name.
  + If `origin_type` is set to **obs_bucket**, this field can only be set to OBS bucket domain name. The OBS bucket
    domain name must end with `.myhuaweicloud.com` or `.myhuaweicloud.cn`.

* `origin_type` - (Required, String) Specifies the origin server type. The valid values are as follows:
  + **ipaddr**: Origin server IP address.
  + **domain**: Origin server domain name.
  + **obs_bucket**: OBS bucket domain name.

* `active` - (Optional, Int) Specifies whether the origin server is primary or standby. Valid values are as follows:
  + **1**: Primary.
  + **0**: Standby.

  Defaults to **1**.

* `obs_web_hosting_enabled` - (Optional, Bool) Specifies whether to enable static website hosting for the OBS bucket.
  This parameter is valid only when the `origin_type` is set to **obs_bucket**. Defaults to **false**.

* `http_port` - (Optional, Int) Specifies the HTTP port. The port number ranges from `1` to `65535`.
  Defaults to **80**.

* `https_port` - (Optional, Int) Specifies the HTTPS port. The port number ranges from `1` to `65535`.
  Default value: **443**.

-> Fields `http_port` and `https_port` are valid only when `origin_type` is set to **ipaddr** or **domain**.

* `retrieval_host` - (Optional, String) Specifies the retrieval host. Things to note when using this field are as follows:
  + If `origin_type` is set to **ipaddr** or **domain**, the acceleration domain name will be used by default.
  + If `origin_type` is set to **obs_bucket**, the bucket's domain name will be used by default.

<a name="configs_object"></a>
The `configs` block support:

* `origin_protocol` - (Optional, String) Specifies the content retrieval protocol. Possible values:
  + **follow**: Same as user requests.
  + **http**: HTTP, which is the default value.
  + **https**: HTTPS.

* `ipv6_enable` - (Optional, Bool) Specifies whether to enable IPv6.

* `range_based_retrieval_enabled` - (Optional, Bool) Specifies whether to enable range-based retrieval.

  -> The prerequisite for enabling range-based retrieval is that your origin site supports Range requests, that is, the
  HTTP request header contains the Range field. Otherwise, the back-to-origin may fail.

* `description` - (Optional, String) Specifies the description of the domain. The value contains up to `200` characters.

* `slice_etag_status` - (Optional, String) Specifies whether ETag is verified during origin pull.
  Valid values are as follows:
  + **on**: Enable.
  + **off**: Disable.

  Defaults to **on**.

* `origin_receive_timeout` - (Optional, Int) Specifies the origin response timeout.
  The value ranges from **5** to **60**, in seconds. Defaults to **30**.

* `https_settings` - (Optional, List) Specifies the certificate configuration. The [https_settings](#https_settings_object)
  structure is documented below.

* `retrieval_request_header` - (Optional, List) Specifies the retrieval request header settings.
  The [retrieval_request_header](#retrieval_request_header_object) structure is documented below.

* `http_response_header` - (Optional, List) Specifies the HTTP response header settings.
  The [http_response_header](#http_response_header_object) structure is documented below.

* `url_signing` - (Optional, List) Specifies the URL signing.
  The [url_signing](#url_signing_object) structure is documented below.

* `force_redirect` - (Optional, List) Specifies the force redirect.
  The [force_redirect](#force_redirect_object) structure is documented below.

* `compress` - (Optional, List) Specifies the smart compression. The [compress](#compress_object) structure
  is documented below.

* `cache_url_parameter_filter` - (Optional, List) Specifies the settings for caching URL parameters.
  The [cache_url_parameter_filter](#cache_url_parameter_filter_object) structure is documented below.

* `ip_frequency_limit` - (Optional, List) Specifies the IP access frequency limit.
  The [ip_frequency_limit](#ip_frequency_limit_object) structure is documented below.

  -> Restricting the IP access frequency can effectively defend against CC attacks, but it may affect normal access.
  Please set access thresholds carefully. After creating the domain name, please wait a few minutes before configuring
  this field, otherwise the configuration may fail.

* `websocket` - (Optional, List) Specifies the websocket settings. This field can only be configured if `type` is
  set to **wholeSite**. The [websocket](#websocket_object) structure is documented below.

  -> Websocket and HTTP/2 are incompatible and cannot be both enabled. Websocket will not take effect when
  origin cache control is enabled in the cache configuration.

* `flexible_origin` - (Optional, List) Specifies the advanced origin rules.
  The [flexible_origin](#flexible_origin_object) structure is documented below.

  -> Up to 20 advanced origin rules can be configured. When `type` is configured as **wholeSite**, configuring this
  field is not supported.

* `remote_auth` - (Optional, List) Specifies the remote authentication settings.
  The [remote_auth](#remote_auth_object) structure is documented below.

  -> Configure remote authentication to allow CDN to forward user requests to an authentication server and process the
  requests based on results returned by the authentication server.

* `quic` - (Optional, List) Specifies the QUIC protocol. The [quic](#quic_object) structure is documented below.

  -> This field can only be used when the HTTPS certificate is enabled. Disabling the HTTPS certificate will disable QUIC.

* `referer` - (Optional, List) Specifies the referer validation. The [referer](#referer_object) structure is documented below.

  -> You can define referer whitelists and blacklists to control who can access specific domain names.

<a name="https_settings_object"></a>
The `https_settings` block support:

* `https_enabled` - (Optional, Bool) Specifies whether to enable HTTPS. Defaults to **false**.

* `certificate_name` - (Optional, String) Specifies the certificate name. The value contains `3` to `32` characters.
  This parameter is mandatory when a certificate is configured.

* `certificate_body` - (Optional, String) Specifies the content of the certificate used by the HTTPS protocol.
  This parameter is mandatory when a certificate is configured. The value is in PEM format.

* `private_key` - (Optional, String) Specifies the private key used by the HTTPS protocol. This parameter is mandatory
  when a certificate is configured. The value is in PEM format.

* `certificate_source` - (Optional, Int) Specifies the certificate source. Currently, only **0** is supported, which means
  your own certificate. Defaults to **0**.

* `certificate_type` - (Optional, String) Specifies the certificate type. Currently, only **server** is supported, which
  means international certificate. Defaults to **server**.

* `http2_enabled` - (Optional, Bool) Specifies whether HTTP/2 is used. Defaults to **false**.
  When `https_enabled` is set to **false**, this parameter does not take effect.

  -> Currently, this field does not support closing after it is enabled.

* `tls_version` - (Optional, String) Specifies the transport Layer Security (TLS). Currently, **TLSv1.0**,
  **TLSv1.1**, **TLSv1.2**, and **TLSv1.3** are supported. By default, **TLSv1.1**, **TLSv1.2**, and **TLSv1.3** are
  enabled. You can enable a single version or consecutive versions. To enable multiple versions, use commas (,) to
  separate versions, for example, **TLSv1.1,TLSv1.2**.

* `ocsp_stapling_status` - (Optional, String) Specifies whether online certificate status protocol (OCSP) stapling is enabled.
  Valid values are as follows:
  + **on**: Enable.
  + **off**: Disable.

  Defaults to **off**.

<a name="retrieval_request_header_object"></a>
The `retrieval_request_header` block support:

* `name` - (Required, String) Specifies the name of a retrieval request header. The value contains `1` to `64` characters,
  including digits, letters, and hyphens (-). The value must start with a letter.

* `action` - (Required, String) Specifies the operation type of the retrieval request header. Valid values are **delete**
  and **set**. If the header does not exist in the original retrieval request, add the header before setting its value.

* `value` - (Optional, String) Specifies the value of the retrieval request header. The value contains `1` to `1000`
  characters, including letters, digits, and special characters `.-_*#!&+|^~'"/:;,=@?<>`. Variables, for example,
  `$client_ip` and `$remote_port`, are not supported.

<a name="http_response_header_object"></a>
The `http_response_header` block support:

* `name` - (Required, String) Specifies the HTTP response header. Valid values are **Content-Disposition**, **Content-Language**,
  **Access-Control-Allow-Origin**, **Access-Control-Allow-Methods**, **Access-Control-Max-Age**, **Access-Control-Expose-Headers**,
  **Access-Control-Allow-Headers** or custom headers. A header contains `1` to `100` characters, including letters, digits,
  and hyphens (-), and starts with a letter.

* `action` - (Required, String) Specifies the operation type of the HTTP response header. The value can be **set** or **delete**.

* `value` - (Optional, String) Specifies the value of the HTTP response header. The value contains `1` to `128` characters,
  including letters, digits, and special characters `.-_*#!&+|^~'"/:;,=@?<>`.

<a name="url_signing_object"></a>
The `url_signing` block support:

* `enabled` - (Required, Bool) Specifies whether to enable of A/B/C URL signing method.

* `type` - (Optional, String) Specifies the signing Method, possible values are:
  **type_a**: Method A.
  **type_b**: Method B.
  **type_c1**: Method C1.
  **type_c2**: Method C2.

* `sign_method` - (Optional, String) Specifies the encryption algorithm type for URL authentication.
  The valid values are as following:
  + **md5**
  + **sha256**

* `match_type` - (Optional, String) Specifies the authentication scope.
  Currently, only spuuort value is **all**, indicates that all files are involved in authentication.

* `sign_arg` - (Optional, String) Specifies the authentication parameters.
  The valid length is limited from `1` to `100` characters, only letters, digits, and underscores (_) are allowed.
  The value cang not start with a digit.

* `inherit_config` - (Optional, List) Specifies the details of the authentication inheritance.
  The [inherit_config](#inherit_config_object) structure is documented below.

* `key` - (Optional, String) Specifies the authentication key contains `16` to `32` characters, including letters and digits.

* `backup_key` - (Optional, String) Specifies the standby authentication key contains `16` to `32` characters,
  including letters and digits.

* `time_format` - (Optional, String) Specifies the time format. Possible values are:
  **dec**: Decimal, can be used in Method A, Method B and Method C2.
  **hex**: Hexadecimal, can be used in Method C1 and Method C2.

* `expire_time` - (Optional, Int) Specifies the expiration time. The value ranges from `0` to `31536000`, in seconds.

<a name="inherit_config_object"></a>
The `inherit_config` blocks support:

* `enabled` - (Required, Bool) Specifies whether to enable authentication inheritance.

* `inherit_type` - (Optional, String) Specifies the authentication inheritance configuration.
  The valid values are as follows:
  + **m3u8**
  + **mpd**

-> Separate multiple values with commas (,). e.g. **m3u8,mpd**.
-> This parameter is mandatory when authentication inheritance is enabled.

* `inherit_time_type` - (Optional, String) Specifies the start time of authentication inheritance.
  The valid values are as follows:
  + **sys_time**: The current system time.
  + **parent_url_time**: The time consistent with the M3U8 and MPD access links.

  -> This parameter is mandatory when authentication inheritance is enabled.

<a name="force_redirect_object"></a>
The `force_redirect` blocks support:

* `enabled` - (Required, Bool) Specifies whether to enable force redirect.

* `type` - (Optional, String) Specifies the force redirect type.
  Possible values are: **http** (force redirect to HTTP) and **https** (force redirect to HTTPS).

  -> Force redirect **https** type can be set only if https is enabled.

* `redirect_code` - (Optional, Int) Specifies the force redirect status code. Valid values are: **301** and **302**.
  Defaults to **302**.

<a name="compress_object"></a>
The `compress` blocks support:

* `enabled` - (Required, Bool) Specifies whether to enable smart compression.

* `type` - (Optional, String) Specifies the smart compression type.
  Possible values are: **gzip** (gzip) and **br** (Brotli).

<a name="cache_url_parameter_filter_object"></a>
The `cache_url_parameter_filter` block support:

* `type` - (Optional, String) Specifies the operation type for caching URL parameters. Valid values are:
  **full_url**: Cache all parameters
  **ignore_url_params**: Ignore all parameters
  **del_params**: Ignore specific URL parameters
  **reserve_params**: Reserve specified URL parameters

* `value` - (Optional, String) Specifies the parameter values. Multiple values are separated by semicolons (;).

<a name="ip_frequency_limit_object"></a>
The `ip_frequency_limit` block support:

* `enabled` - (Required, Bool) Specifies whether to enable IP access frequency.

* `qps` - (Optional, Int) Specifies the access threshold, in times/second. The value ranges from **1** to **100,000**.
  This field is required when enable IP access frequency.

<a name="websocket_object"></a>
The `websocket` block support:

* `enabled` - (Required, Bool) Specifies whether to enable websocket settings.

* `timeout` - (Optional, Int) Specifies the duration for keeping a connection open, in seconds. The value ranges
  from **1** to **300**. This field is required when enable websocket settings.

<a name="flexible_origin_object"></a>
The `flexible_origin` block support:

* `match_type` - (Required, String) Specifies the URI match mode. Valid values are as follows:
  + **all**: All files.
  + **file_extension**: File name extension.
  + **file_path**: Directory.

* `priority` - (Required, Int) Specifies the priority. The value of this field must be unique. Value ranges from **1**
  to **100**. A greater number indicates a higher priority.

* `back_sources` - (Required, List) Specifies the back source information. The length of this array field cannot exceed `1`.
  The [back_sources](#flexible_origin_back_sources_object) structure is documented below.

* `match_pattern` - (Optional, String) Specifies the URI match rule. The usage rules are as follows:
  + When `match_type` is set to **all**, set this field to empty.
  + When `match_type` is set to **file_extension**, the value of this field should start with a period (.).
    Enter up to 20 file name extensions and use semicolons (;) to separate them. Example: **.jpg;.zip;.exe**.
  + When `match_type` is set to **file_path**, the value of this field should start with a slash (/).
    Enter up to 20 paths and use semicolons (;) to separate them. Example: **/test/folder01;/test/folder02**.

<a name="flexible_origin_back_sources_object"></a>
The `back_sources` block support:

* `sources_type` - (Required, String) Specifies the origin server type. Valid values are as follows:
  + **ipaddr**: IP address.
  + **domain**: Domain name.
  + **obs_bucket**: OBS bucket.

* `ip_or_domain` - (Required, String) Specifies the IP address or domain name of the origin server.
  + When `sources_type` is set to **ipaddr**, the value of this field can only be set to a valid IPv4 or Ipv6 address.
  + When `sources_type` is set to **domain**, the value of this field can only be set to a domain name.
  + When `sources_type` is set to **obs_bucket**, the value of this field can only be set to an OBS bucket access
    domain name.

* `obs_bucket_type` - (Optional, String) Specifies the OBS bucket type. Valid values are **private** and **public**.
  This field is required when `sources_type` is set to **obs_bucket**.

* `http_port` - (Optional, Int) Specifies the HTTP port, ranging from **1** to **65535**. Defaults to **80**.

* `https_port` - (Optional, Int) Specifies the HTTPS port, ranging from **1** to **65535**. Defaults to **443**.

-> Fields `http_port` and `https_port` do not support editing when `sources_type` is set to **obs_bucket**.

<a name="remote_auth_object"></a>
The `remote_auth` block support:

* `enabled` - (Required, Bool) Specifies whether to enable remote authentication.

* `remote_auth_rules` - (Optional, List) Specifies the remote authentication settings. The length of this array field
  cannot exceed `1`. The [remote_auth_rules](#remote_auth_rules_object) structure is documented below.

<a name="remote_auth_rules_object"></a>
The `remote_auth_rules` block support:

* `auth_server` - (Required, String) Specifies the address of a reachable server. The address must include **http://** or
  **https://**. The address cannot be a local address such as **localhost** or **127.0.0.1**. The address cannot be an
  acceleration domain name added on CDN.

* `request_method` - (Required, String) Specifies the request method supported by the authentication server. Valid values
  are **GET**, **POST**, and **HEAD**.

* `file_type_setting` - (Required, String) Specifies the authentication file type settings. Valid values are:
  + **all**: Requests for all files are authenticated.
  + **specific_file**: Requests for files of specific types are authenticated.

* `reserve_args_setting` - (Required, String) Specifies the parameters that need to be authenticated in user requests.
  Valid values are as follows:
  + **reserve_all_args**: Retain all URL parameters.
  + **reserve_specific_args**: Retain specified URL parameters.
  + **ignore_all_args**: Ignore all URL parameters.

* `reserve_headers_setting` - (Required, String) Specifies the headers to be authenticated in user requests.
  Valid values are as follows:
  + **reserve_all_headers**: Retain all request headers.
  + **reserve_specific_headers**: Retain specified request headers.
  + **ignore_all_headers**: Ignore all request headers.

* `auth_success_status` - (Required, String) Specifies the status code returned by the remote authentication server
  to CDN nodes when authentication is successful. Value range: **2xx** and **3xx**.

* `auth_failed_status` - (Required, String) Specifies the status code returned by the remote authentication server
  to CDN nodes when authentication is failed. Value range: **4xx** and **5xx**.

* `response_status` - (Required, String) Specifies the status code returned by CDN nodes to users when authentication
  is failed. Value range: **2xx**, **3xx**, **4xx**, and **5xx**.

* `timeout` - (Required, Int) Specifies the duration from the time when a CDN node forwards an authentication request
  to the time when the CDN node receives the result returned by the remote authentication server. Enter **0** or a value
  ranging from **50** to **3000**. The unit is millisecond.

* `timeout_action` - (Required, String) Specifies the action of the CDN nodes to process user requests after the
  authentication timeout. Valid values are as follows:
  + **pass**: The user request is allowed and the corresponding resource is returned after the authentication times out.
  + **forbid**: The user request is rejected after the authentication times out and the configured status code is
    returned to the user.

* `specified_file_type` - (Optional, String) Specifies the specific file types. The value contains letters and digits.
  The value contains up to `512` characters. File types are not case-sensitive, and multiple file types are separated
  by vertical bars (|). For example: **jpg|MP4**. This parameter is mandatory when `file_type_setting` is set to
  **specific_file**. In other cases, this parameter is left blank.

* `reserve_args` - (Optional, String) Specifies the reserve args. Multiple args are separated by vertical bars (|).
  For example: **key1|key2**. This parameter is mandatory when `reserve_args_setting` is set to **reserve_specific_args**.
  In other cases, this parameter is left blank.

* `reserve_headers` - (Optional, String) Specifies the reserve headers. Multiple headers are separated by vertical bars (|).
  For example: **key1|key2**. This parameter is mandatory when `reserve_headers_setting` is set to **reserve_specific_headers**.
  In other cases, this parameter is left blank.

* `add_custom_args_rules` - (Optional, List) Specifies the URL validation parameters.
  The [add_custom_args_rules](#add_custom_rules_object) structure is documented below.

* `add_custom_headers_rules` - (Optional, List) Specifies the request header authentication parameters.
  The [add_custom_headers_rules](#add_custom_rules_object) structure is documented below.

<a name="add_custom_rules_object"></a>
The `add_custom_args_rules` and `add_custom_headers_rules` block support:

* `type` - (Required, String) Specifies the parameter type. Valid values are:
  + **custom_var**: Custom parameter type.
  + **nginx_preset_var**: Preset variable type.

* `key` - (Required, String) Specifies the parameter key. The value contains up to `256` characters. The value can be
  composed of digits, uppercase letters, lowercase letters, and special characters (._-*#%|+^@?=).

* `value` - (Required, String) Specifies the parameter value. The usage restrictions of this field are as follows:
  + When `type` is set to **custom_var**, the value contains up to `256` characters. The value can be composed of digits,
    uppercase letters, lowercase letters, and special characters (._-*#%|+^@?=).
  + When `type` is set to **nginx_preset_var**, the value can only be **$http_host**, **$http_user_agent**,
    **$http_referer**, **$http_x_forwarded_for**, **$http_content_type**, **$remote_addr**, **$scheme**,
    **$server_protocol**, **$request_uri**, **$uri**, **$args**, or **$request_method**.

<a name="quic_object"></a>
The `quic` block support:

* `enabled` - (Required, Bool) Specifies whether to enable QUIC.

<a name="referer_object"></a>
The `referer` block support:

* `type` - (Required, String) Specifies the referer validation type. Valid values are as follows:
  + **off**: Disable referer validation.
  + **black**: Referer blacklist.
  + **white**: Referer whitelist.

* `value` - (Optional, String) Specifies the domain names or IP addresses, which are separated by commas (,).
  Wildcard domain names and domain names with port numbers are supported. Enter up to `400` domain names and IP addresses.
  The port number ranges from `1` to `65535`. This field is required when `type` is set to **black** or **white**.

* `include_empty` - (Optional, Bool) Specifies whether blank `referers` are included.
  A referer blacklist including blank `referers` indicates that requests without any `referers` are not allowed to access.
  A referer whitelist including blank `referers` indicates that requests without any `referers` are allowed to access.

  Defaults to **false**.

<a name="cache_settings_object"></a>
The `cache_settings` block support:

* `follow_origin` - (Optional, Bool) Specifies whether to enable origin cache control. Defaults to **false**.

* `rules` - (Optional, List) Specifies the cache rules, which overwrite the previous rule configurations.
  Blank rules are reset to default rules. The [rules](#rules_object) structure is documented below.

<a name="rules_object"></a>
The `rules` block support:

* `rule_type` - (Required, String) Specifies the rule type. Possible value are:
  + **all**: All types of files are matched. It is the default value. The cloud will create a cache rule with **all**
    rule type by default.
  + **file_extension**: Files are matched based on their suffixes.
  + **catalog**: Files are matched based on their directories.
  + **full_path**: Files are matched based on their full paths.
  + **home_page**: Files are matched based on their homepage.

* `ttl` - (Required, Int) Specifies the cache age. The maximum cache age is 365 days.

* `ttl_type` - (Required, String) Specifies the unit of the cache age. Possible values:
  + **s**: Second
  + **m**: Minute
  + **h**: Hour
  + **d**: Day

* `priority` - (Required, Int) Specifies the priority weight of this rule. The default value is 1.
  A larger value indicates a higher priority. The value ranges from 1 to 100. The weight values must be unique.

* `content` - (Optional, String) Specifies the content that matches `rule_type`.
  + If `rule_type` is set to **all** or **home_page**, keep this parameter empty.
  + If `rule_type` is set to **file_extension**, the value of this parameter is a list of file name
    extensions. A file name extension starts with a period (.). File name extensions are separated by semicolons (;),
    for example, `.jpg;.zip;.exe`. Up to 20 file types are supported.
  + If `rule_type` is set to **catalog**, the value of this parameter is a list of directories. A directory starts with
    a slash (/). Directories are separated by semicolons (;), for example, `/test/folder01;/test/folder02`.
    Up to 20 directories are supported.
  + If `rule_type` is set to **full_path**, the value must start with a slash (/) and cannot end with an asterisk.
    Example: `/test/index.html` or `/test/*.jpg`

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The acceleration domain name ID.

* `cname` - The CNAME of the acceleration domain name.

* `domain_status` - The status of the acceleration domain name. The available values are
  'online', 'offline', 'configuring', 'configure_failed', 'checking', 'check_failed' and 'deleting.'

* `configs/https_settings/https_status` - The status of the https. The available values are 'on' and 'off'.

* `configs/https_settings/http2_status` - The status of the http 2.0. The available values are 'on' and 'off'.

* `configs/url_signing/status` - The status of the url_signing. The available values are 'on' and 'off'.

* `configs/url_signing/inherit_config/status` - The status of the authentication inheritance.
  The valid values are **on** and **off**.

* `configs/force_redirect/status` - The status of the force redirect. The available values are 'on' and 'off'.

* `configs/compress/status` - The status of the compress. The available values are 'on' and 'off'.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 30 minutes.
* `update` - Default is 30 minutes.
* `delete` - Default is 20 minutes.

## Import

The CDN domain resource can be imported using the domain `name`, e.g.

```bash
$ terraform import huaweicloud_cdn_domain.test <name>
```

Note that the imported state may not be identical to your resource definition, due to some attributes missing from the
API response, security or some other reason. The missing attributes include: `enterprise_project_id`,
`configs.0.url_signing.0.key`, `configs.0.https_settings.0.certificate_body`, `configs.0.https_settings.0.private_key`,
and `cache_settings`.
It is generally recommended running `terraform plan` after importing a resource.
You can then decide if changes should be applied to the resource, or the resource definition should be updated to align
with the resource. Also, you can ignore changes as below.

```hcl
resource "huaweicloud_cdn_domain" "test" {
  ...
  
  lifecycle {
    ignore_changes = [
      enterprise_project_id, configs.0.url_signing.0.key, configs.0.https_settings.0.certificate_body,
      configs.0.https_settings.0.private_key, cache_settings,
    ]
  }
}
```
