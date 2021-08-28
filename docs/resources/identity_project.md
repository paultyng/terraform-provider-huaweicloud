---
subcategory: "Identity and Access Management (IAM)"
---

# huaweicloud_identity_project

Manages a Project resource within HuaweiCloud Identity And Access Management service. This is an alternative
to `huaweicloud_identity_project_v3`

Note: You _must_ have security admin privileges in your HuaweiCloud cloud to use this resource.

## Example Usage

```hcl
resource "huaweicloud_identity_project" "project_1" {
  name        = "cn-north-1_project1"
  description = "This is a test project"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Specifies the name of the project. it must start with an existing _region_ and be less
  than or equal to 64 characters. Example: cn-north-1_project1.

* `description` - (Optional, String) Specifies the description of the project.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - A resource ID in UUID format.

* `parent_id` - The parent of this project.

* `enabled` - Enabling status of this project.

## Import

Projects can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_identity_project.project_1 89c60255-9bd6-460c-822a-e2b959ede9d2
```
