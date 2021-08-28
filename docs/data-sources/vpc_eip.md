---
subcategory: "Elastic IP (EIP)"
---

# huaweicloud_vpc_eip

Use this data source to get the details of an available EIP.

## Example Usage

```hcl
data "huaweicloud_vpc_eip" "by_address" {
  public_ip = "123.60.208.163"
}
```

## Argument Reference

* `region` - (Optional, String) The region in which to obtain the EIP. If omitted, the provider-level region will be
  used.

* `public_ip` - (Optional, String) The public ip address of the EIP.

* `port_id` - (Optional, String) The port id of the EIP.

* `enterprise_project_id` - (Optional, String) The enterprise project id of the EIP.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `status` - The status of the EIP.

* `type` - The type of the EIP.

* `private_ip` - The private ip of the EIP.

* `bandwidth_id` - The bandwidth id of the EIP.

* `bandwidth_size` - The bandwidth size of the EIP.

* `bandwidth_share_type` - The bandwidth share type of the EIP.
