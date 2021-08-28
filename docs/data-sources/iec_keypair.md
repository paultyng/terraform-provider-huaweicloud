---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_keypair

Use this data source to get the details of a specific IEC keypair.

## Example Usage

```hcl
data "huaweicloud_iec_keypair" "kp_1" {
  name = "iec-keypair-demo"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Specifies a unique name for the keypair. This parameter can contain a maximum of 64
  characters, which may consist of letters, digits, underscores (_), and hyphens (-).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source use `name` as the ID.

* `public_key` - The pregenerated OpenSSH-formatted public key.

* `fingerprint` - The finger of iec keypair. The value contains a encoding type(SHA256) and a string of 43 characters.
