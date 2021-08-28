---
subcategory: "Web Application Firewall (WAF)"
---

# huaweicloud_waf_reference_table

Manages a WAF reference table resource within HuaweiCloud.

-> **NOTE:** All WAF resources depend on WAF instances, and the WAF instances need to be purchased before they can be
used. The reference table can be used in Cloud Mode (professional version), Dedicated Mode and ELB Mode.

## Example Usage

```hcl
resource "huaweicloud_waf_reference_table" "ref_table" {
  name = "tf_ref_table_demo"
  type = "url"

  conditions = [
    "/admin",
    "/manage"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the WAF reference table resource. If omitted,
  the provider-level region will be used. Changing this setting will push a new reference table.

* `name` - (Required, String) The name of the reference table. Only letters, digits, and underscores(_) are allowed. The
  maximum length is 64 characters.

* `type` - (Required, String, ForceNew) The type of the reference table, The options are `url`, `user-agent`, `ip`,
  `params`, `cookie`, `referer` and `header`. Changing this setting will push a new reference table.

* `conditions` - (Required, List) The conditions of the reference table. The maximum length is 30. The maximum length of
  condition is 2048 characters.

* `description` - (Optional, String) The description of the reference table. The maximum length is 128 characters.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id of the reference table.

* `creation_time` - The server time when reference table was created.

## Import

The reference table can be imported using the `id`, e.g.

```sh
terraform import huaweicloud_waf_reference_table.ref_table 96e46e5e702b4e2aa5609ad287de4788
```
