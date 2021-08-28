---
subcategory: "SoftWare Repository for Container (SWR)"
---

# huaweicloud_swr_organization

Manages a SWR organization resource within HuaweiCloud.

## Example Usage

```hcl
resource "huaweicloud_swr_organization" "test" {
  name = "terraform-test"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the resource. If omitted, the
  provider-level region will be used. Changing this creates a new resource.

* `name` - (Required, String, ForceNew) Specifies the name of the organization. The organization name must be globally
  unique.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the organization.

* `creator` - The creator user name of the organization.

* `permission` - The permission of the organization, the value can be Manage, Write, and Read.

* `login_server` - The URL that can be used to log into the container registry.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

Organizations can be imported using the `name`, e.g.

```
$ terraform import huaweicloud_swr_organization.test org-name
```
