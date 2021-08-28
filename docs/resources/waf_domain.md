---
subcategory: "Web Application Firewall (WAF)"
---

# huaweicloud_waf_domain

Manages a WAF domain resource within HuaweiCloud.

## Example Usage

```hcl
resource "huaweicloud_waf_certificate" "certificate_1" {
  name        = "cert_1"
  certificate = <<EOT
-----BEGIN CERTIFICATE-----
MIIFmQl5dh2QUAeo39TIKtadgAgh4zHx09kSgayS9Wph9LEqq7MA+2042L3J9aOa
DAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUR+SosWwALt6PkP0J9iOIxA6RW8gVsLwq
...
+HhDvD/VeOHytX3RAs2GeTOtxyAV5XpKY5r+PkyUqPJj04t3d0Fopi0gNtLpMF=
-----END CERTIFICATE-----
EOT
  private_key = <<EOT
-----BEGIN PRIVATE KEY-----
MIIJwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAM
ATAwMC4GCCsGAQUFBwIBFiJodHRwOi8vY3BzLnJvb3QteDEubGV0c2VuY3J5cHQu
...
he8Y4IWS6wY7bCkjCWDcRQJMEhg76fsO3txE+FiYruq9RUWhiF1myv4Q6W+CyBFC
1qoJFlcDyqSMo5iHq3HLjs
-----END PRIVATE KEY-----
EOT
}

resource "huaweicloud_waf_domain" "domain_1" {
  domain           = "www.example.com"
  certificate_id   = huaweicloud_waf_certificate.certificate_1.id
  certificate_name = huaweicloud_waf_certificate.certificate_1.name
  proxy            = true

  server {
    client_protocol = "HTTPS"
    server_protocol = "HTTP"
    address         = "119.8.0.13"
    port            = "8080"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the WAF domain resource. If omitted, the
  provider-level region will be used. Changing this setting will push a new certificate.

* `domain` - (Required, String, ForceNew) Specifies the domain name to be protected. For example, www.example.com or *
  .example.com. Changing this creates a new domain.

* `server` - (Required, List) Specifies an array of origin web servers. The object structure is documented below.

* `certificate_id` - (Required, String) Specifies the certificate ID. This parameter is mandatory when `client_protocol`
  is set to HTTPS.

* `certificate_name` - (Required, String) Specifies the certificate name. This parameter is mandatory
  when `client_protocol` is set to HTTPS.

* `policy_id` - (Optional, String, ForceNew) Specifies the policy ID associated with the domain. If not specified, a new
  policy will be created automatically. Changing this create a new domain.

* `keep_proxy` - (Optional, Bool) Specifies whether to retain the policy when deleting a domain name. Defaults to true.

* `proxy` - (Optional, Bool) Specifies whether a proxy is configured.

The `server` block supports:

* `client_protocol` - (Required, String) Protocol type of the client. The options include `HTTP` and `HTTPS`.

* `server_protocol` - (Required, String) Protocol used by WAF to forward client requests to the server. The options
  include `HTTP` and `HTTPS`.

* `address` - (Required, String) IP address or domain name of the web server that the client accesses. For example,
  192.168.1.1 or www.a.com.

* `port` - (Required, Int) Port number used by the web server. The value ranges from 0 to 65535, for example, 8080.

## Attributes Reference

The following attributes are exported:

* `id` - ID of the domain.

* `protect_status` - The WAF mode. -1: bypassed, 0: disabled, 1: enabled.

* `access_status` - Whether a domain name is connected to WAF. 0: The domain name is not connected to WAF, 1: The domain
  name is connected to WAF.

* `protocol` - The protocol type of the client. The options are HTTP, HTTPS, and HTTP&HTTPS.

## Import

Domains can be imported using the `id`, e.g.

```sh
terraform import huaweicloud_waf_domain.domain_2 7902bd9e01104cb794dcb668f235e0c5
```
