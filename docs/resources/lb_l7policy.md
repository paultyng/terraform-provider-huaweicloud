---
subcategory: "Elastic Load Balance (ELB)"
---

# huaweicloud_lb_l7policy

Manages an ELB L7 Policy resource within HuaweiCloud. This is an alternative to `huaweicloud_lb_l7policy_v2`

## Example Usage

```hcl
resource "huaweicloud_lb_loadbalancer" "loadbalancer_1" {
  name          = "loadbalancer_1"
  vip_subnet_id = var.subnet_id
}

resource "huaweicloud_lb_listener" "listener_1" {
  name            = "listener_1"
  protocol        = "HTTP"
  protocol_port   = 8080
  loadbalancer_id = huaweicloud_lb_loadbalancer.loadbalancer_1.id
}

resource "huaweicloud_lb_pool" "pool_1" {
  name            = "pool_1"
  protocol        = "HTTP"
  lb_method       = "ROUND_ROBIN"
  loadbalancer_id = huaweicloud_lb_loadbalancer.loadbalancer_1.id
}

resource "huaweicloud_lb_l7policy" "l7policy_1" {
  name             = "test"
  action           = "REDIRECT_TO_POOL"
  description      = "test l7 policy"
  position         = 1
  listener_id      = huaweicloud_lb_listener.listener_1.id
  redirect_pool_id = huaweicloud_lb_pool.pool_1.id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the L7 Policy resource. If omitted, the
  provider-level region will be used. Changing this creates a new L7 Policy.

* `name` - (Optional, String) Human-readable name for the L7 Policy. Does not have to be unique.

* `description` - (Optional, String) Human-readable description for the L7 Policy.

* `action` - (Required, String, ForceNew) The L7 Policy action - can either be REDIRECT\_TO\_POOL, or
  REDIRECT\_TO\_LISTENER. Changing this creates a new L7 Policy.

* `listener_id` - (Required, String, ForceNew) The Listener on which the L7 Policy will be associated with. Changing
  this creates a new L7 Policy.

* `position` - (Optional, Int, ForceNew) The position of this policy on the listener. Positions start at 1. Changing
  this creates a new L7 Policy.

* `redirect_pool_id` - (Optional, String) Requests matching this policy will be redirected to the pool with this ID.
  Only valid if action is REDIRECT\_TO\_POOL.

* `redirect_listener_id` - (Optional, String) Requests matching this policy will be redirected to the listener with this
  ID. Only valid if action is REDIRECT\_TO\_LISTENER.

* `admin_state_up` - (Optional, Bool) The administrative state of the L7 Policy. This value can only be true (UP).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique ID for the L7 {olicy.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `update` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.:

```
$ terraform import huaweicloud_lb_l7policy.l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74
```
