---
subcategory: "Elastic Load Balance (ELB)"
---

# huaweicloud_lb_whitelist

Manages an ELB whitelist resource within HuaweiCloud. This is an alternative to `huaweicloud_lb_whitelist_v2`

## Example Usage

```hcl
resource "huaweicloud_lb_listener" "listener_1" {
  name            = "listener_1"
  protocol        = "HTTP"
  protocol_port   = 8080
  loadbalancer_id = var.loadbalancer_id
}

resource "huaweicloud_lb_whitelist" "whitelist_1" {
  enable_whitelist = true
  whitelist        = "192.168.11.1,192.168.0.1/24,192.168.201.18/8"
  listener_id      = huaweicloud_lb_listener.listener_1.id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the ELB whitelist resource. If omitted, the
  provider-level region will be used. Changing this creates a new whitelist.

* `listener_id` - (Required, String, ForceNew) The Listener ID that the whitelist will be associated with. Changing this
  creates a new whitelist.

* `enable_whitelist` - (Optional, Bool) Specify whether to enable access control.

* `whitelist` - (Optional, String) Specifies the IP addresses in the whitelist. Use commas(,) to separate the multiple
  IP addresses.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique ID for the whitelist.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `update` - Default is 10 minute.
* `delete` - Default is 10 minute.
