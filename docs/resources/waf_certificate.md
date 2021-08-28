---
subcategory: "Web Application Firewall (WAF)"
---

# huaweicloud_waf_certificate

Manages a WAF certificate resource within HuaweiCloud.

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
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the WAF certificate resource. If omitted, the
  provider-level region will be used. Changing this setting will push a new certificate.

* `name` - (Required, String) Specifies the certificate name. The maximum length is 256 characters. Only digits,
  letters, underscores(`_`), and hyphens(`-`) are allowed.

* `certificate` - (Required, String, ForceNew) Specifies the certificate content. Changing this creates a new
  certificate.

* `private_key` - (Required, String, ForceNew) Specifies the private key. Changing this creates a new certificate.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The certificate ID in UUID format.

* `expiration` - Indicates the time when the certificate expires.

## Import

Certificates can be imported using the `id`, e.g.

```sh
terraform import huaweicloud_waf_certificate.certificate_2 3ebd3201238d41f9bfc3623b61435954
```

Note that the imported state is not identical to your resource definition, due to security reason. The missing
attributes include `certificate`, and `private_key`. You can ignore changes as below.

```
resource "huaweicloud_waf_certificate" "certificate_2" {
    ...
  lifecycle {
    ignore_changes = [
      certificate, private_key,
    ]
  }
}
```
