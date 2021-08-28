---
subcategory: "Dedicated Load Balance (Dedicated ELB)"
---

# huaweicloud_elb_certificate

Use this data source to get the certificate in HuaweiCloud Dedicated Load Balance (Dedicated ELB).

## Example Usage

```hcl
variable "certificate_name" {}

data "huaweicloud_elb_certificate" "test" {
  name = var.certificate_name
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) The region in which to obtain the Dedicated ELB certificate. If omitted, the
  provider-level region will be used.

* `name` - (Required, String) The name of certificate. The value is case sensitive and does not supports fuzzy matching.

  -> **NOTE:** The certificate name is not unique. Only returns the last created one when matched multiple certificates.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The certificate ID in UUID format.

* `domain` - The domain of the Certificate. This parameter is valid only when `type` is "server".

* `type` - Specifies the certificate type. The value can be one of the following:
  + `server`: indicates the server certificate.
  + `client`: indicates the CA certificate.

* `description` - Human-readable description for the Certificate.

* `expiration` - Indicates the time when the certificate expires.
