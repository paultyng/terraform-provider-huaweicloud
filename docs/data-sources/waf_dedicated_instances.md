---
subcategory: "Web Application Firewall (WAF)"
---

# huaweicloud_waf_dedicated_instances

Use this data source to get a list of WAF dedicated instances.

## Example Usage

```hcl
variable instance_name {}

data "huaweicloud_waf_dedicated_instances" "instances" {
  name = var.instance_name
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) The region in which to query the WAF dedicated instance. If omitted, the provider-level
  region will be used.

* `id` - (Optional, String) The id of WAF dedicated instance.

* `name` - (Optional, String) The name of WAF dedicated instance.

## Attributes Reference

* `id` - The data source ID in UUID format.

The following attributes are exported:

The `instances` block supports:

* `id` - The id of WAF dedicated instance.

* `name` - The name of WAF dedicated instance.

* `available_zone` - The available zone names for the WAF dedicated instances.

* `specification_code` - The specification code of instance. Different specifications have different throughput. Values
  are:
    + `waf.instance.professional` - The professional edition, throughput: 100 Mbit/s; QPS: 2,000 (Reference only).
    + `waf.instance.enterprise` - The enterprise edition, throughput: 500 Mbit/s; QPS: 10,000 (Reference only).

* `cpu_architecture` - The ECS cpu architecture of WAF dedicated instance.

* `ecs_flavor` - The flavor of the ECS used by the WAF instance.

* `vpc_id` - The VPC id of WAF dedicated instance.

* `subnet_id` - The subnet id of WAF dedicated instance VPC.

* `security_group` - The security group of the instance. This is an array of security group ids.

* `server_id` - The service of the instance.

* `service_ip` - The service ip of the instance.

* `run_status` - The running status of the instance. Values are:
    + `0` - Instance is creating.
    + `1` - Instance has created.
    + `2` - Instance is deleting.
    + `3` - Instance has deleted.
    + `4` - Instance create failed.

* `access_status` - The access status of the instance. `0`: inaccessible, `1`: accessible.

* `upgradable` - The instance is to support upgrades. `0`: Cannot be upgraded, `1`: Can be upgraded.
