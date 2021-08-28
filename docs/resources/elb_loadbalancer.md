---
subcategory: "Dedicated Load Balance (Dedicated ELB)"
---

# huaweicloud_elb_loadbalancer

Manages a Dedicated Load Balancer resource within HuaweiCloud.

## Example Usage

### Basic Loadbalancer

```hcl
resource "huaweicloud_elb_loadbalancer" "basic" {
  name              = "basic"
  description       = "basic example"
  cross_vpc_backend = true

  vpc_id         = "{{ vpc_id }}"
  ipv4_subnet_id = "{{ subnet_id }}"

  l4_flavor_id = "{{ l4_flavor_id }}"
  l7_flavor_id = "{{ l7_flavor_id }}"

  availability_zone     = [
    "cn-north-4a",
    "cn-north-4b",
  ]
  enterprise_project_id = "{{ eps_id }}"
}
```

### Loadbalancer With Existing EIP

```hcl
resource "huaweicloud_elb_loadbalancer" "basic" {
  name              = "basic"
  description       = "basic example"
  cross_vpc_backend = true

  vpc_id            = "{{ vpc_id }}"
  ipv6_network_id   = "{{ ipv6_network_id }}"
  ipv6_bandwidth_id = "{{ ipv6_bandwidth_id }}"
  ipv4_subnet_id    = "{{ subnet_id }}"

  l4_flavor_id = "{{ l4_flavor_id }}"
  l7_flavor_id = "{{ l7_flavor_id }}"

  availability_zone     = [
    "cn-north-4a",
    "cn-north-4b",
  ]
  enterprise_project_id = "{{ eps_id }}"

  ipv4_eip_id = "{{ eip_id }}"
}
```

### Loadbalancer With EIP

```hcl
resource "huaweicloud_elb_loadbalancer" "basic" {
  name              = "basic"
  description       = "basic example"
  cross_vpc_backend = true

  vpc_id            = "{{ vpc_id }}"
  ipv6_network_id   = "{{ ipv6_network_id }}"
  ipv6_bandwidth_id = "{{ ipv6_bandwidth_id }}"
  ipv4_subnet_id    = "{{ subnet_id }}"

  l4_flavor_id = "{{ l4_flavor_id }}"
  l7_flavor_id = "{{ l7_flavor_id }}"

  availability_zone     = [
    "cn-north-4a",
    "cn-north-4b",
  ]
  enterprise_project_id = "{{ eps_id }}"

  iptype                = "5_bgp"
  bandwidth_charge_mode = "traffic"
  sharetype             = "PER"
  bandwidth_size        = 10
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the loadbalancer resource. If omitted, the
  provider-level region will be used. Changing this creates a new loadbalancer.

* `availability_zone` - (Required, List, ForceNew) Specifies the list of AZ names. Changing this parameter will create a
  new resource.

* `name` - (Required, String) Human-readable name for the loadbalancer.

* `description` - (Optional, String) Human-readable description for the loadbalancer.

* `cross_vpc_backend` - (Optional, Bool) Enable this if you want to associate the IP addresses of backend servers with
  your load balancer. Can only be true when updating.

* `vpc_id` - (Optional, String, ForceNew) The vpc on which to create the loadbalancer. Changing this creates a new
  loadbalancer.

* `ipv4_subnet_id` - (Optional, String) The subnet on which to allocate the loadbalancer's ipv4 address.

* `ipv6_network_id` - (Optional, String) The network on which to allocate the loadbalancer's ipv6 address.

* `ipv6_bandwidth_id` - (Optional, String) The ipv6 bandwidth id. Only support shared bandwidth.

* `ipv4_address` - (Optional, String) The ipv4 address of the load balancer.

* `ipv4_eip_id` - (Optional, String, ForceNew) The ID of the EIP. Changing this parameter will create a new resource.

-> **NOTE:** If the ipv4_eip_id parameter is configured, you do not need to configure the bandwidth parameters:
`iptype`, `bandwidth_charge_mode`, `bandwidth_size` and `share_type`.

* `iptype` - (Optional, String, ForceNew) Elastic IP type. Changing this parameter will create a new resource.

* `bandwidth_charge_mode` - (Optional, String, ForceNew) Bandwidth billing type. Changing this parameter will create a
  new resource.

* `sharetype` - (Optional, String, ForceNew) Bandwidth sharing type. Changing this parameter will create a new resource.

* `bandwidth_size` - (Optional, Int, ForceNew) Bandwidth size. Changing this parameter will create a new resource.

* `l4_flavor_id` - (Optional, String) The L4 flavor id of the load balancer.

* `l7_flavor_id` - (Optional, String) The L7 flavor id of the load balancer.

* `tags` - (Optional, Map) The key/value pairs to associate with the loadbalancer.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id of the loadbalancer. Changing this
  creates a new loadbalancer.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `ipv4_eip` - The ipv4 eip address of the Load Balancer.
* `ipv6_eip` - The ipv6 eip address of the Load Balancer.
* `ipv6_eip_id` - The ipv6 eip id of the Load Balancer.
* `ipv6_address` - The ipv6 address of the Load Balancer.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `update` - Default is 10 minute.
* `delete` - Default is 5 minute.
