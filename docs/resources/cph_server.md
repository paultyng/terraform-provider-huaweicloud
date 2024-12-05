---
subcategory: "Cloud Phone (CPH)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_cph_server"
description: ""
---

# huaweicloud_cph_server

Manages a CPH server resource within HuaweiCloud.  

## Example Usage

```HCL
  variable "name" {}
  variable "server_flavor" {}
  variable "phone_flavor" {}
  variable "image_id" {}
  variable "keypair" {}
  variable "vpc_id" {}
  variable "subnet_id" {}

  resource "huaweicloud_cph_server" "test" {
    name          = var.name
    server_flavor = var.server_flavor
    phone_flavor  = var.phone_flavor
    image_id      = var.image_id
    keypair_name  = var.keypair

    vpc_id    = var.vpc_id
    subnet_id = var.subnet_id
    eip_type  = "5_bgp"

    bandwidth {
      share_type  = "0"
      charge_mode = "1"
      size        = 300
    }

    period_unit = "month"
    period      = 1
    auto_renew  = "true"

  }
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the resource.
  If omitted, the provider-level region will be used. Changing this parameter will create a new resource.

* `name` - (Required, String) Server name.  
  The name can contain `1` to `60` characters, only English letters, Chinese characters, digits, underscore (_) and
  hyphens (-) are allowed.

* `server_flavor` - (Required, String, ForceNew) The CPH server flavor.

  Changing this parameter will create a new resource.

* `phone_flavor` - (Required, String) The cloud phone flavor.
  
* `image_id` - (Required, String) The cloud phone image ID.

* `vpc_id` - (Required, String) The ID of VPC which the cloud server belongs to.

* `subnet_id` - (Required, String) The ID of subnet which the cloud server belongs to.

* `availability_zone` - (Optional, String, ForceNew) The name of the AZ where the cloud server is located.

  Changing this parameter will create a new resource.

* `eip_id` - (Optional, String, ForceNew) The ID of an **existing** EIP assigned to the cloud server.
  This parameter and `eip_type`, `bandwidth` are alternative.

* `eip_type` - (Optional, String, ForceNew) The type of an EIP that will be automatically assigned to the cloud server.
  The options are as follows:
    + **5_telcom**: China Telecom.
    + **5_union**: China Unicom.
    + **5_bgp**: Dynamic BGP.
    + **5_sbgp**: Static BGP.

  Changing this parameter will create a new resource.

* `bandwidth` - (Optional, List, ForceNew) The bandwidth of an EIP that will be automatically assigned to
  the cloud server.
  The [BandWidth](#cphServer_BandWidth) structure is documented below.

* `period_unit` - (Required, String, ForceNew) The charging period unit.  
  Valid values are **month** and **year**.

  Changing this parameter will create a new resource.

* `period` - (Required, Int, ForceNew) The charging period.  
  If `period_unit` is set to **month**, the value ranges from `1` to `9`.
  If `period_unit` is set to **year**, the value ranges from `1` to `3`.

  Changing this parameter will create a new resource.

* `auto_renew` - (Required, String, ForceNew) Whether auto renew is enabled. Valid values are **true** and **false**.
  Defaults to false.  

  Changing this parameter will create a new resource.

* `keypair_name` - (Optional, String) Specifies the key pair name, which is used for logging in to
  the cloud phone through ADB.

* `enterprise_project_id` - (Optional, String) The enterprise project ID.

* `ports` - (Optional, List) The application port enabled by the cloud phone.
  The [ApplicationPort](#cphServer_ApplicationPort) structure is documented below.

* `phone_data_volume` - (Optional, List) The phone data volume.
  The [phone_data_volume](#phone_data_volume) structure is documented below.

* `server_share_data_volume` - (Optional, List) The server share data volume.
  The [server_share_data_volume](#server_share_data_volume) structure is documented below.

<a name="cphServer_BandWidth"></a>
The `BandWidth` block supports:

* `share_type` - (Required, String) The bandwidth type.  
  The options are as follows:
    + **0**: Dedicated bandwidth.
    + **1**: Shared bandwidth.

* `id` - (Optional, String) The bandwidth ID.  
 You can specify an existing shared bandwidth when assigning an EIP for a shared bandwidth.
 This parameter is mandatory when you create a shared bandwidth.

* `size` - (Optional, Int) The bandwidth (Mbit/s).  
  The valid value is range from `1` to `2,000`.  
  This parameter is mandatory for a dedicated bandwidth.

* `charge_mode` - (Optional, String) Which the bandwidth used by the CPH server is billed.  
 This parameter is mandatory for a dedicated bandwidth.
 The options are as follows:
   + **0**: Billed by bandwidth.
   + **1**: Billed by traffic.

<a name="cphServer_ApplicationPort"></a>
The `ApplicationPort` block supports:

* `name` - (Required, String) The application port name, which can contain a maximum of 16 bytes.  
 The key service name cannot be **adb** or **vnc**.

* `listen_port` - (Required, Int) The port number, which ranges from `10,000` to `50,000`.

* `internet_accessible` - (Required, String) Whether public network access is mapped.
  The options are as follows:
    + **true**: public network access is mapped.
    + **false**: no mapping is performed.

<a name="phone_data_volume"></a>
The `phone_data_volume` block supports:

* `volume_size` - (Optional, Int, ForceNew) Specifies the volume size, the unit is GB.
  Changing this parameter will create a new resource.

* `volume_type` - (Optional, String, ForceNew) Specifies the volume type.
  Changing this parameter will create a new resource.

<a name="server_share_data_volume"></a>
The `server_share_data_volume` block supports:

* `volume_type` - (Optional, String, ForceNew) Specifies the share volume type.
  Changing this parameter will create a new resource.

* `size` - (Optional, Int, ForceNew) Specifies the share volume size, the unit is GB.
  Changing this parameter will create a new resource.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.

* `order_id` - The order ID.

* `addresses` - The IP addresses of the CPH server.
  The [Address](#cphServer_Address) structure is documented below.

* `security_groups` - The list of the security groups bound to the extension NIC of the CPH server.

* `status` - The CPH server status.  
  The options are as follows:
    + **0**, **1**, **3**, and **4**: Creating.
    + **2**: Abnormal.
    + **5**: Normal.
    + **8**: Frozen.
    + **10**: Stopped.
    + **11**: Being stopped.
    + **12**: Stopping failed.
    + **13**: Starting.

<a name="cphServer_Address"></a>
The `Address` block supports:

* `server_ip` - The internal IP address of the CPH server.  

* `public_ip` - The public IP address of the CPH server.  

* `phone_data_volume` - The phone data volume.
  The [phone_data_volume](#attr_phone_data_volume) structure is documented below.

* `server_share_data_volume` - The server share data volume.
  The [server_share_data_volume](#attr_server_share_data_volume) structure is documented below.

<a name="attr_phone_data_volume"></a>
The `phone_data_volume` block supports:

* `volume_id` - The volume ID.

* `volume_name` - The volume name.

* `created_at` - The creation time.

* `updated_at` - The update time.

<a name="attr_server_share_data_volume"></a>
The `server_share_data_volume` block supports:

* `version` - The share volume type.

## Import

The CPH server can be imported using the `id`, e.g.

```bash
$ terraform import huaweicloud_cph_server.test 0ce123456a00f2591fabc00385ff1234
```

Note that the imported state may not be identical to your resource definition, due to some attributes missing from the
API response, security or some other reason. The missing attributes include: `image_id`, `eip_id`, `eip_type`, `auto_renew`,
`period`, and `period_unit`. It is generally recommended running `terraform plan` after importing a resource.
You can then decide if changes should be applied to the resource, or the resource definition should be updated to
align with the resource. Also you can ignore changes as below.

```hcl
resource "huaweicloud_cph_server" "test" {
    ...

    lifecycle {
      ignore_changes = [
        image_id, eip_id, eip_type, auto_renew, period, period_unit,
      ]
    }
}
