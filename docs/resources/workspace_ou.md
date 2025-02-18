---
subcategory: "Workspace"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_workspace_ou"
description: |-
  Manages an OU resource of Workspace within HuaweiCloud.
---
# huaweicloud_workspace_ou

Manages an OU resource of Workspace within HuaweiCloud.

-> Before creating an OU, you need to create OUs on the AD server and register the Workspace service
   to connect to the AD domain.

## Example Usage

```hcl
variable "ou_name" {}
variable "ad_domain_name" {}

resource "huaweicloud_workspace_ou" "test" {
  name   = var.ou_name
  domain = var.ad_domain_name
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the resource.
  If omitted, the provider-level region will be used.
  Changing this creates a new resource.

* `name` - (Required, String) Specifies the name of the OU.  
  The name only allows Chinese and English characters, digits, spaces and special characters (-_$!@*?.).
  Multiple levels of OUs are separated by slashes (/) and no space is allowed before or after the slashes (/),
  a maximum of five layers of OUs are supported. e.g. `ab/cd/ef`.

* `domain` - (Required, String, ForceNew) Specifies the AD domain name to which the OU belongs.

* `description` - (Optional, String) Specifies the description of the OU.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID, also OU ID.

## Import

The OU resource can be imported using `name`, e.g.

```bash
$ terraform import huaweicloud_workspace_ou.test <name>
```
