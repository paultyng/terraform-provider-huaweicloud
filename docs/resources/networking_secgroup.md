---
subcategory: "Virtual Private Cloud (VPC)"
---

# huaweicloud_networking_secgroup

Manages a Security Group resource within HuaweiCloud. This is an alternative to `huaweicloud_networking_secgroup_v2`

## Example Usage

```hcl
resource "huaweicloud_networking_secgroup" "secgroup" {
  name        = "secgroup_1"
  description = "My security group"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the security group resource. If omitted, the
  provider-level region will be used. Changing this creates a new security group resource.

* `name` - (Required, String) Specifies a unique name for the security group.

* `description` - (Optional, String) Specifies the description for the security group.

* `enterprise_project_id` - (Optional, String, ForceNew) Specifies the enterprise project id of the security group.
  Changing this creates a new security group.

* `delete_default_rules` - (Optional, Bool, ForceNew) Specifies whether or not to delete the default security rules.
  This is `false` by default.

-> **NOTE:** The default security rules are described
in [HuaweiCloud](https://support.huaweicloud.com/intl/en-us/usermanual-vpc/SecurityGroup_0003.html). See the below
section for more information.

## Default Security Group Rules

In most cases, HuaweiCloud will create some security group rules for each new security group. These security group rules
will not be managed by Terraform, so if you prefer to have *all*
aspects of your infrastructure managed by Terraform, set `delete_default_rules` to `true`
and then create separate security group rules such as the following:

```hcl
resource "huaweicloud_networking_secgroup_rule" "secgroup_rule_v4" {
  security_group_id = huaweicloud_networking_secgroup.secgroup.id
  direction         = "egress"
  ethertype         = "IPv4"
}

resource "huaweicloud_networking_secgroup_rule" "secgroup_rule_v6" {
  security_group_id = huaweicloud_networking_secgroup.secgroup.id
  direction         = "egress"
  ethertype         = "IPv6"
}

resource "huaweicloud_networking_secgroup_rule" "allow_ssh" {
  security_group_id = huaweicloud_networking_secgroup.secgroup.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 22
  port_range_max    = 22
  remote_ip_prefix  = "0.0.0.0/0"
}
```

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID in UUID format.

## Timeouts

This resource provides the following timeouts configuration options:

* `delete` - Default is 10 minute.

## Import

Security Groups can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_networking_secgroup.secgroup_1 38809219-5e8a-4852-9139-6f461c90e8bc
```
