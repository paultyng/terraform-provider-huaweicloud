---
subcategory: "Identity and Access Management (IAM)"
---

# huaweicloud_identity_role_assignment

Manages a Role assignment within group on HuaweiCloud IAM Service. This is an alternative
to `huaweicloud_identity_role_assignment_v3`

Note: You _must_ have admin privileges in your HuaweiCloud cloud to use this resource.

## Example Usage: Assign Role On Project Level

```hcl
data "huaweicloud_identity_role" "role_1" {
  name = "rds_adm"
  #RDS Administrator
}

resource "huaweicloud_identity_group" "group_1" {
  name = "group_1"
}

resource "huaweicloud_identity_role_assignment" "role_assignment_1" {
  role_id    = data.huaweicloud_identity_role.role_1.id
  group_id   = huaweicloud_identity_group.group_1.id
  project_id = var.project_id
}
```

## Example Usage: Assign Role On Domain Level

```hcl
data "huaweicloud_identity_role" "role_1" {
  name = "secu_admin"
  #Security Administrator
}

resource "huaweicloud_identity_group" "group_1" {
  name = "group_1"
}

resource "huaweicloud_identity_role_assignment" "role_assignment_1" {
  role_id   = data.huaweicloud_identity_role.role_1.id
  group_id  = huaweicloud_identity_group.group_1.id
  domain_id = var.domain_id
}
```

## Argument Reference

The following arguments are supported:

* `role_id` - (Required, String, ForceNew) Specifies the role to assign.

* `group_id` - (Required, String, ForceNew) Specifies the group to assign the role to.

* `domain_id` - (Optional, String, ForceNew; Required if `project_id` is empty) Specifies the domain to assign the role
  in.

* `project_id` - (Optional, String, ForceNew; Required if `domain_id` is empty) Specifies the project to assign the role
  in.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID in UUID format.
