---
subcategory: "Cloud Data Migration (CDM)"
---

# huaweicloud_cdm_cluster

CDM cluster management This is an alternative to `huaweicloud_cdm_cluster_v1`

## Example Usage

### create a cdm cluster

```hcl
resource "huaweicloud_networking_secgroup" "secgroup" {
  name        = "terraform_test_security_group"
  description = "terraform security group acceptance test"
}

resource "huaweicloud_cdm_cluster" "cluster" {
  availability_zone = "{{ availability_zone }}"
  flavor_id         = "{{ flavor_id }}"
  name              = "terraform_test_cdm_cluster"
  security_group_id = huaweicloud_networking_secgroup.secgroup.id
  subnet_id         = "{{ network_id }}"
  vpc_id            = "{{ vpc_id }}"
  version           = "{{ version }}"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the cluster resource. If omitted, the
  provider-level region will be used. Changing this creates a new CDM cluster resource.

* `availability_zone` - (Required, String, ForceNew) Available zone. Changing this parameter will create a new resource.

* `flavor_id` - (Required, String, ForceNew) Flavor id. Changing this parameter will create a new resource.

* `name` - (Required, String, ForceNew) Cluster name. Changing this parameter will create a new resource.

* `security_group_id` - (Required, String, ForceNew) Security group ID. Changing this parameter will create a new
  resource.

* `subnet_id` - (Required, String, ForceNew) Subnet ID. Changing this parameter will create a new resource.

* `version` - (Required, String, ForceNew) Cluster version. Changing this parameter will create a new resource.

* `vpc_id` - (Required, String, ForceNew) VPC ID. Changing this parameter will create a new resource.

* `email` - (Optional, List, ForceNew) Notification email addresses. The max number is 5. Changing this parameter will
  create a new resource.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id. Changing this parameter will create
  a new resource.

* `is_auto_off` - (Optional, Bool, ForceNew) Whether to automatically shut down. Changing this parameter will create a
  new resource.

* `phone_num` - (Optional, List, ForceNew) Notification phone numbers. The max number is 5. Changing this parameter will
  create a new resource.

* `schedule_boot_time` - (Optional, String, ForceNew) Timed boot time. Changing this parameter will create a new
  resource.

* `schedule_off_time` - (Optional, String, ForceNew) Timed shutdown time. Changing this parameter will create a new
  resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

* `created` - Create time.

* `instances` - Instance list. Structure is documented below.

* `publid_ip` - Public ip.

The `instances` block contains:

* `id` - Instance ID.

* `name` - Instance name.

* `public_ip` - Public IP.

* `role` - Role.

* `traffic_ip` - Traffic IP.

* `type` - Instance type.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 30 minute.

