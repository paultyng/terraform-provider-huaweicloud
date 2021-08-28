---
subcategory: "Elastic Cloud Server (ECS)"
---

# huaweicloud_compute_eip_associate

Associate an EIP to a compute instance.

## Example Usage

### Automatically detect the correct network

```hcl
resource "huaweicloud_compute_instance" "myinstance" {
  name              = "instance"
  image_id          = "ad091b52-742f-469e-8f3c-fd81cadf0743"
  flavor_id         = "s6.small.1"
  key_pair          = "my_key_pair_name"
  security_groups   = [
    "default"]
  availability_zone = "cn-north-4a"

  network {
    uuid = "55534eaa-533a-419d-9b40-ec427ea7195a"
  }
}

resource "huaweicloud_vpc_eip" "myeip" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name        = "test"
    size        = 8
    share_type  = "PER"
    charge_mode = "traffic"
  }
}

resource "huaweicloud_compute_eip_associate" "associated" {
  public_ip   = huaweicloud_vpc_eip.myeip.address
  instance_id = huaweicloud_compute_instance.myinstance.id
}
```

### Explicitly set the network to attach to

```hcl
resource "huaweicloud_compute_instance" "myinstance" {
  name              = "instance"
  image_id          = "ad091b52-742f-469e-8f3c-fd81cadf0743"
  flavor_id         = "s6.small.1"
  key_pair          = "my_key_pair_name"
  security_groups   = [
    "default"]
  availability_zone = "cn-north-4a"

  network {
    uuid = "55534eaa-533a-419d-9b40-ec427ea7195a"
  }

  network {
    uuid = "3c4a0d74-24b9-46cf-9d7f-8b7a4dc2f65c"
  }
}

resource "huaweicloud_vpc_eip" "myeip" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name        = "test"
    size        = 8
    share_type  = "PER"
    charge_mode = "traffic"
  }
}

resource "huaweicloud_compute_eip_associate" "associated" {
  public_ip   = huaweicloud_vpc_eip.myeip.address
  instance_id = huaweicloud_compute_instance.myinstance.id
  fixed_ip    = huaweicloud_compute_instance.myinstance.network.1.fixed_ip_v4
}
```

## Argument Reference

The following arguments are supported:

* `public_ip` - (Required, String, ForceNew) Specifies the EIP address to associate.

* `instance_id` - (Required, String, ForceNew) Specifies the instance to associte the EIP with.

* `fixed_ip` - (Optional, String, ForceNew) Specifies the IP address to direct traffic to.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

## Import

This resource can be imported by specifying all three arguments, separated by a forward slash:

```
$ terraform import huaweicloud_compute_eip_associate.eip_1 <eip>/<instance_id>/<fixed_ip>
```
