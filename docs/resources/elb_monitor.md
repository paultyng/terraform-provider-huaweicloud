---
subcategory: "Dedicated Load Balance (Dedicated ELB)"
---

# huaweicloud_elb_monitor

Manages an ELB monitor resource within HuaweiCloud.

## Example Usage

```hcl
resource "huaweicloud_elb_monitor" "monitor_1" {
  protocol    = "HTTP"
  interval    = 30
  timeout     = 15
  max_retries = 10
  url_path    = "/api"
  port        = 8888
  pool_id     = huaweicloud_elb_pool.test.id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the ELB monitor resource. If omitted, the
  provider-level region will be used. Changing this creates a new monitor.

* `pool_id` - (Required, String, ForceNew) The id of the pool that this monitor will be assigned to.

* `protocol` - (Required, String, ForceNew) The type of probe, which is TCP, HTTP, or HTTPS, that is sent by the load
  balancer to verify the member state. Changing this creates a new monitor.

* `domain_name` - (Optional, String) The Domain Name of the Monitor.

* `port` - (Optional, Int) Specifies the health check port. The value ranges from 1 to 65535.

* `interval` - (Required, Int) The time, in seconds, between sending probes to members.

* `timeout` - (Required, Int) Maximum number of seconds for a monitor to wait for a ping reply before it times out. The
  value must be less than the delay value.

* `max_retries` - (Required, Int) Number of permissible ping failures before changing the member's status to INACTIVE.
  Must be a number between 1 and 10.

* `url_path` - (Optional, String) Required for HTTP(S) types. URI path that will be accessed if monitor type is HTTP or
  HTTPS.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique ID for the monitor.
