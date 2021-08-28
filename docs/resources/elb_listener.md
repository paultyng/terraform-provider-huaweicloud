---
subcategory: "Dedicated Load Balance (Dedicated ELB)"
---

# huaweicloud_elb_listener

Manages an ELB listener resource within HuaweiCloud.

## Example Usage

```hcl
resource "huaweicloud_elb_loadbalancer" "test" {
  ...
}

resource "huaweicloud_elb_listener" "basic" {
  name            = "basic"
  description     = "basic description"
  protocol        = "HTTP"
  protocol_port   = 8080
  loadbalancer_id = huaweicloud_elb_loadbalancer.test.id

  idle_timeout     = 60
  request_timeout  = 60
  response_timeout = 60

  tags = {
    key = "value"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the listener resource. If omitted, the
  provider-level region will be used. Changing this creates a new listener.

* `protocol` - (Required, String, ForceNew) The protocol can either be TCP, UDP, HTTP or HTTPS. Changing this creates a
  new listener.

* `protocol_port` - (Required, Int, ForceNew) The port on which to listen for client traffic. Changing this creates a
  new listener.

* `loadbalancer_id` - (Required, String, ForceNew) The load balancer on which to provision this listener. Changing this
  creates a new listener.

* `name` - (Optional, String) Human-readable name for the listener.

* `default_pool_id` - (Optional, String) The ID of the default pool with which the listener is associated. Changing this
  creates a new listener.

* `description` - (Optional, String) Human-readable description for the listener.

* `http2_enable` - (Optional, Bool) Specifies whether to use HTTP/2. The default value is false. This parameter is valid
  only when the protocol is set to *HTTPS*.

* `forward_eip` - (Optional, Bool) Specifies whether transfer the load balancer EIP in the X-Forward-EIP header to
  backend servers. The default value is false. This parameter is valid only when the protocol is set to *HTTP* or *
  HTTPS*.

* `access_policy` - (Optional, String) Specifies the access policy for the listener. Valid options are *white* and *
  black*.

* `ip_group` - (Optional, String) Specifies the ip group id for the listener.

* `server_certificate` - (Optional, String) Specifies the ID of the server certificate used by the listener. This
  parameter is mandatory when protocol is set to *HTTPS*.

* `sni_certificate` - (Optional, List) Lists the IDs of SNI certificates (server certificates with a domain name) used
  by the listener. This parameter is valid when protocol is set to *HTTPS*.

* `ca_certificate` - (Optional, String) Specifies the ID of the CA certificate used by the listener. This parameter is
  valid when protocol is set to *HTTPS*.

* `tls_ciphers_policy` - (Optional, String) Specifies the TLS cipher policy for the listener. Valid options are:
  tls-1-0-inherit, tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict, tls-1-2-fs, tls-1-0-with-1-3, and tls-1-2-fs-with-1-3.
  This parameter is valid when protocol is set to *HTTPS*.

* `idle_timeout` - (Optional, Int) Specifies the idle timeout for the listener. Value range: 0 to 4000.

* `request_timeout` - (Optional, Int) Specifies the request timeout for the listener. Value range: 1 to 300. This
  parameter is valid when protocol is set to *HTTP* or *HTTPS*.

* `response_timeout` - (Optional, Int) Specifies the response timeout for the listener. Value range: 1 to 300. This
  parameter is valid when protocol is set to *HTTP* or *HTTPS*.

* `tags` - (Optional, Map) The key/value pairs to associate with the listener.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique ID for the listener.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `update` - Default is 10 minute.
* `delete` - Default is 10 minute.
