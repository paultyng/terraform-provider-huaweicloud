---
subcategory: "Web Application Firewall (WAF)"
---

# huaweicloud_waf_dedicated_instance

Manages a WAF dedicated instance resource within HuaweiCloud.

## Example Usage

```hcl
variable az_name {}
variable ecs_flavor_id {}
variable vpc_id {}
variable subnet_id {}
variable security_group_id {}

resource "huaweicloud_waf_dedicated_instance" "instance_1" {
  name               = "instance_1"
  available_zone     = var.az_name
  specification_code = "waf.instance.professional"
  ecs_flavor         = var.ecs_flavor_id
  vpc_id             = var.vpc_id
  subnet_id          = var.subnet_id

  security_group = [
    var.security_group_id
  ]
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the WAF dedicated instance. If omitted, the
  provider-level region will be used. Changing this setting will create a new instance.

* `name` - (Required, String) The name of WAF dedicated instance. Duplicate names are allowed, we suggest to keeping the
  name unique.

* `available_zone` - (Required, String, ForceNew) The available zone names for the dedicated instances. It can be
  obtained through this data source `huaweicloud_availability_zones`. Changing this will create a new instance.

* `specification_code` - (Required, String, ForceNew) The specification code of instance. Different specifications have
  different throughput. Changing this will create a new instance. Values are:
  + `waf.instance.professional` - The professional edition, throughput: 100 Mbit/s; QPS: 2,000 (Reference only).
  + `waf.instance.enterprise` - The enterprise edition, throughput: 500 Mbit/s; QPS: 10,000 (Reference only).

* `ecs_flavor` - (Required, String, ForceNew) The flavor of the ECS used by the WAF instance. Flavors can be obtained
  through this data source `huaweicloud_compute_flavors`. Changing this will create a new instance.

  -> **NOTE:** If the instance specification is the professional edition, the ECS specification should be 2U4G. If the
  instance specification is the enterprise edition, the ECS specification should be 8U16G.

* `vpc_id` - (Required, String, ForceNew) The VPC id of WAF dedicated instance. Changing this will create a new
  instance.

* `subnet_id` - (Required, String, ForceNew) The subnet id of WAF dedicated instance VPC. Changing this will create a
  new instance.

* `security_group` - (Required, List, ForceNew) The security group of the instance. This is an array of security group
  ids. Changing this will create a new instance.

* `cpu_architecture` - (Optional, String, ForceNew) The ECS cpu architecture of instance, Default value is `x86`.
  Changing this will create a new instance.

## Attributes Reference

The following attributes are exported:

* `id` - The id of the instance.

* `server_id` - The id of the instance server.

* `service_ip` - The ip of the instance service.

* `run_status` - The running status of the instance. Values are:
  + `0` - Instance is creating.
  + `1` - Instance has created.
  + `2` - Instance is deleting.
  + `3` - Instance has deleted.
  + `4` - Instance create failed.

* `access_status` - The access status of the instance. `0`: inaccessible, `1`: accessible.

* `upgradable` - The instance is to support upgrades. `0`: Cannot be upgraded, `1`: Can be upgraded.

## Import

WAF dedicated instance can be imported using the `id`, e.g.

```sh
terraform import huaweicloud_waf_dedicated_instance.instance_1 2f87641090206b821f07e0f6bd6
```
