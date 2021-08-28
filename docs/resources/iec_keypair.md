---
subcategory: "Intelligent EdgeCloud (IEC)"
---

# huaweicloud_iec_keypair

Manages a keypair resource within HuaweiCloud IEC.

## Example Usage

```hcl
resource "huaweicloud_iec_keypair" "test_keypair" {
  name = "iec-keypair-demo"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String, ForceNew) Specifies a unique name for the keypair. This parameter can contain a maximum of
  64 characters, which may consist of letters, digits, underscores (_), and hyphens (-). Changing this parameter creates
  a new keypair resource.

* `public_key` - (Optional, String, ForceNew) Specifies a pregenerated OpenSSH-formatted public key. Changing this
  parameter creates a new keypair resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The keypair use the unique name as the ID.

* `fingerprint` - The finger of iec keypair. The value contains a encoding type(SHA256) and a string of 43 characters.

## Import

Keypairs can be imported using the `name`, e.g.

```
$ terraform import huaweicloud_iec_keypair.test_keypair iec-keypair-demo
```
