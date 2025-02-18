---
subcategory: "CodeArts Inspector"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_codearts_inspector_websites"
description: |-
  Use this data source to get the list of CodeArts inspector websites.
---

# huaweicloud_codearts_inspector_websites

Use this data source to get the list of CodeArts inspector websites.

## Example Usage

```hcl
data "huaweicloud_codearts_inspector_websites" "test" {}
```

## Argument Reference

The following arguments are supported:

* `domain_id` - (Optional, String) Specifies the domain ID.

* `auth_status` - (Optional, String) Specifies the auth status of website.
  Valid values are:
  + **unauth**: Unauthorized.
  + **auth**: Authorized.
  + **invalid**: Authentication file is invalid.
  + **manual**: Manual authentication.
  + **skip**: Authentication free.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `top_level_domain_num` - Indicates the number of top level domain

* `websites` - Indicates the websites list.
  The [websites](#attrblock--websites) structure is documented below.

<a name="attrblock--websites"></a>
The `websites` block supports:

* `id` - Indicates the domain ID.

* `website_name` - Indicates the website name.

* `website_address` - Indicates the website address.

* `auth_status` - Indicates the auth status of website.

* `created_at` - Indicates the time to create website.

* `high` - Indicates the number of high-risk vulnerabilities.

* `hint` - Indicates the number of hint-risk vulnerabilities.

* `low` - Indicates the number of low-severity vulnerabilities.

* `middle` - Indicates the number of medium-risk vulnerabilities.

* `top_level_domain_id` - Indicates the top level domain ID.
