---
subcategory: "SSL Certificate Manager (SCM)"
---

# huaweicloud_scm_certificate

SSL Certificate Manager (SCM) allows you to purchase Secure Sockets Layer (SSL) certificates from the world's leading
digital certificate authorities (CAs), upload existing SSL certificates, and centrally manage all your SSL certificates
in one place.

## Example Usage

### Load the certificate contents from the local files.

```hcl
resource "huaweicloud_scm_certificate" "certificate_1" {
  name              = "certificate_1"
  certificate       = file("/usr/local/data/certificate/cert_xxx/xxx_ca.crt")
  certificate_chain = file("/usr/local/data/certificate/cert_xxx/xxx_ca_chain.crt")
  private_key       = file("/usr/local/data/certificate/cert_xxx/xxx_server.key")
}
```

### Write the contents of the certificate into the Terrafrom script.

```hcl
resource "huaweicloud_scm_certificate" "certificate_2" {
  name              = "certificate_2"
  certificate       = <<EOT
-----BEGIN CERTIFICATE-----
MIIC9DCCAl2gAwIBAgIUUcJZn3ep4l8iHu6lL/jE2UV+G8gwDQYJKoZIhvcNAQEL
ZWlqaW5nMQswC...
(This is an example, please replace it with a encrypted key of valid SSL certificate.)
-----END CERTIFICATE----------
EOT
  certificate_chain = <<EOT
-----BEGIN CERTIFICATE-----
MIIC9DCCAl2gAwIBAgIUUcJZn3ep4l8iHu6lL/jE2UV+G8gwDQYJKoZIhvcNAQEL
BQAwgYsxCzAJB...
(This is an example, please replace it with a encrypted key of valid SSL certificate.)
-----END CERTIFICATE----------
EOT
  private_key       = <<EOT
-----BEGIN PRIVATE KEY-----
QWH3GbHx5bGQyexHj2hre4yEahn4dAKKdjSAMUuSfLWygp2pEdNFOegYTdqk/snv
mhNmxp74oUcVfi1Msw6KY2...
(This is an example, please replace it with a encrypted key of valid SSL certificate.)
-----END PRIVATE KEY-----
EOT
}
```

### Push the SSL certificate to another HUAWEI CLOUD service.

```hcl
# Load the certificate contents from the local files.
resource "huaweicloud_scm_certificate" "certificate_3" {
  name              = "certificate_3"
  certificate       = file("/usr/local/data/certificate/cert_xxx/xxx_ca.crt")
  certificate_chain = file("/usr/local/data/certificate/cert_xxx/xxx_ca_chain.crt")
  private_key       = file("/usr/local/data/certificate/cert_xxx/xxx_server.key")

  target {
    project = [
      "la-south-2"]
    service = "Enhance_ELB"
  }

  target {
    service = "CDN"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the SCM certificate resource. If omitted, the
  provider-level region will be used. Changing this setting will push a new certificate.
* `name` - (Optional, String, ForceNew) Human-readable name for the Certificate. Does not have to be unique. The value
  contains a maximum of 63 characters.
* `certificate` - (Required, String, ForceNew) The public encrypted key of the Certificate, PEM format.
* `certificate_chain` - (Required, String, ForceNew) The chain of the certificate. It can be extracted from the _
  server.crt_ file in the Nginx directory, usually after the second paragraph is the certificate chain.
* `private_key` - (Required, String, ForceNew) The private encrypted key of the Certificate, PEM format.
* `target` - (Optional, List) The service to which the certificate needs to be pushed.

The `target` block supports:

* `service` - (Required, String) Service to which the certificate is pushed. The options include `CDN`,`WAF`
  and `Enhance_ELB`.
* `project` - (Optional, String) The project where the service you want to push a certificate to. The same certificate
  can be pushed repeatedly to the same WAF or ELB service in the same `project`, but the CDN service can only be pushed
  once.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.
* `domain` - Domain name bound to a certificate.
* `domain_count` - Number of domain names can be bound to a certificate.
* `push_support` - Whether a certificate can be pushed.
* `not_before` - Time when the certificate takes effect. If no valid value is obtained, this parameter is left blank.
* `not_after` - Time when the certificate becomes invalid. If no valid value is obtained, this parameter is left blank.
* `status` - Certificate status. The value can be:
  + `PAID` - The certificate has been paid, and needs to be applied for from the CA.
  + `ISSUED` - The certificate has been issued.
  + `CHECKING` - The certificate application is under review.
  + `CANCELCHECKING` - The application of certificate cancellation is under review.
  + `UNPASSED` - The certificate application fails.
  + `EXPIRED` - The certificate has expired.
  + `REVOKING` - The application of certificate revocation is under review.
  + `REVOKED` - The certificate has been revoked.
  + `UPLOAD` - The certificate is being hosted.
  + `SUPPLEMENTCHECKING` - The application for the new domain name of the multi-domain certificate is under review.
  + `CANCELSUPPLEMENTING` - The cancellation on additional domain names to be added is being reviewed.
* `authentifications` - (List) Domain ownership verification information. This is a list, each item of data is as
  follows:
  + `record_name` - Name of a domain ownership verification value.
  + `record_type` - Type of the domain name verification value.
  + `record_value` - Domain verification value.
  + `domain` - Domain name mapping to the verification value

## Import

Certificates can be imported using the `id`, e.g.

```sh
terraform import huaweicloud_scm_certificate.certificate_1 scs1627959834994
```

