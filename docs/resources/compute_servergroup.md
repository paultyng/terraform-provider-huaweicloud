---
subcategory: "Elastic Cloud Server (ECS)"
---

# huaweicloud_compute_servergroup

Manages Server Group resource within HuaweiCloud. This is an alternative to `huaweicloud_compute_servergroup_v2`

## Example Usage

```hcl
data "huaweicloud_compute_instance" "instance_demo" {
  name = "ecs-servergroup-demo"
}

resource "huaweicloud_compute_servergroup" "test-sg" {
  name     = "my-sg"
  policies = [
    "anti-affinity"]
  members  = [
    data.huaweicloud_compute_instance.instance_demo.id,
  ]
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the server group resource. If omitted,
  the provider-level region will be used. Changing this creates a new server group.

* `name` - (Required, String, ForceNew) Specifies a unique name for the server group. This parameter can contain a
  maximum of 255 characters, which may consist of letters, digits, underscores (_), and hyphens (-). Changing this
  creates a new server group.

* `policies` - (Required, List, ForceNew) Specifies the set of policies for the server group. Only *anti-affinity*
  policies are supported.

  + `anti-affinity`: All ECS in this group must be deployed on different hosts. Changing this creates a new server
      group.

* `members` - (Optional, Set) Specifies an array of one or more instance ID to attach server group.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - A resource ID in UUID format.

## Import

Server Groups can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_compute_servergroup.test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
```
