---
subcategory: "Elastic Load Balance (ELB)"
---

# huaweicloud_lb_certificate

Use this data source to get the certificates in HuaweiCloud Elastic Load Balance (ELB).

## Example Usage

```hcl
variable "certificate_name" {}

data "huaweicloud_lb_certificate" "test" {
  name = var.certificate_name
  type = "server"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) The region in which to obtain the ELB certificate. If omitted, the provider-level region
  will be used.

* `type` - (Optional, String) Specifies the certificate type. The default value is `server`. The value can be one of the
  following:
  + `server`: indicates the server certificate.
  + `client`: indicates the CA certificate.

* `name` - (Required, String) The name of certificate. The value is case sensitive and does not supports fuzzy matching.

  -> **NOTE:** The certificate name is not unique. Only returns the last created one when matched multiple certificates.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The certificate ID in UUID format.

* `domain` - The domain of the Certificate. This parameter is valid only when `type` is "server".

* `description` - Human-readable description for the Certificate.

* `expiration` - Indicates the time when the certificate expires.
