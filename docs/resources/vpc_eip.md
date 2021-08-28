---
subcategory: "Elastic IP (EIP)"
---

# huaweicloud_vpc_eip

Manages an EIP resource within HuaweiCloud. This is an alternative to `huaweicloud_vpc_eip_v1`

## Example Usage

### EIP with Dedicated Bandwidth

```hcl
resource "huaweicloud_vpc_eip" "eip_1" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    share_type  = "PER"
    name        = "test"
    size        = 10
    charge_mode = "traffic"
  }
}
```

### EIP with Shared Bandwidth

```hcl
resource "huaweicloud_vpc_bandwidth" "bandwidth_1" {
  name = "bandwidth_1"
  size = 5
}

resource "huaweicloud_vpc_eip" "eip_1" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    share_type = "WHOLE"
    id         = huaweicloud_vpc_bandwidth.bandwidth_1.id
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the eip resource. If omitted, the provider-level
  region will be used. Changing this creates a new eip resource.

* `publicip` - (Required, List) The elastic IP address object.

* `bandwidth` - (Required, List) The bandwidth object.

* `tags` - (Optional, Map) Specifies the key/value pairs to associate with the elastic IP.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id of the elastic IP. Changing this
  creates a new eip.

* `charging_mode` - (Optional, String, ForceNew) Specifies the charging mode of the elastic IP. Valid values are *
  prePaid* and *postPaid*, defaults to *postPaid*. Changing this creates a new eip.

* `period_unit` - (Optional, String, ForceNew) Specifies the charging period unit of the elastic IP. Valid values are *
  month* and *year*. This parameter is mandatory if `charging_mode` is set to *prePaid*. Changing this creates a new
  eip.

* `period` - (Optional, Int, ForceNew) Specifies the charging period of the elastic IP. If `period_unit` is set to *
  month*, the value ranges from 1 to 9. If `period_unit` is set to *year*, the value ranges from 1 to 3. This parameter
  is mandatory if `charging_mode` is set to *prePaid*. Changing this creates a new resource.

* `auto_renew` - (Optional, String, ForceNew) Specifies whether auto renew is enabled. Valid values are "true" and "
  false". Changing this creates a new resource.

The `publicip` block supports:

* `type` - (Required, String, ForceNew) The type of the eip. Changing this creates a new eip.

* `ip_address` - (Optional, String, ForceNew) The value must be a valid IP address in the available IP address segment.
  Changing this creates a new eip.

* `port_id` - (Optional, String) The port id which this eip will associate with. If the value is "" or this not
  specified, the eip will be in unbind state.

The `bandwidth` block supports:

* `share_type` - (Required, String, ForceNew) Whether the bandwidth is dedicated or shared. Changing this creates a new
  eip. Possible values are as follows:
  + *PER*: Dedicated bandwidth
  + *WHOLE*: Shared bandwidth

* `name` - (Optional, String) The bandwidth name, which is a string of 1 to 64 characters that contain letters, digits,
  underscores (_), and hyphens (-). This parameter is mandatory when `share_type` is set to *PER*.

* `size` - (Optional, Int) The bandwidth size. The value ranges from 1 to 300 Mbit/s. This parameter is mandatory
  when `share_type` is set to *PER*.

* `id` - (Optional, String, ForceNew) The shared bandwidth id. This parameter is mandatory when
  `share_type` is set to *WHOLE*. Changing this creates a new eip.

* `charge_mode` - (Optional, String, ForceNew) Specifies whether the bandwidth is billed by traffic or by bandwidth
  size. The value can be *traffic* or *bandwidth*. Changing this creates a new eip.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID in UUID format.
* `address` - The IP address of the eip.
* `status` - The status of eip.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

EIPs can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_vpc_eip.eip_1 2c7f39f3-702b-48d1-940c-b50384177ee1
```
