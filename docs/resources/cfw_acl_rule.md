---
subcategory: "Cloud Firewall (CFW)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_cfw_acl_rule"
description: |-
  Manages a CFW ACL rule resource within HuaweiCloud.
---

# huaweicloud_cfw_acl_rule

Manages a CFW ACL rule resource within HuaweiCloud.

## Example Usage

### Create a basic rule

```hcl
variable "name" {}
variable "description" {}
variable "object_id" {}

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = var.name
  object_id           = var.object_id
  description         = var.description
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1

  source_addresses      = ["1.1.1.1"] 
  destination_addresses = ["1.1.1.2"]

  custom_services {
    protocol    = 6
    source_port = 81
    dest_port   = 82
  }

  sequence {
    top = 1
  }

  tags = {
    key = "value"
  }
}
```

### Create a rule with the source address using the region list

```hcl
variable "name" {}
variable "description" {}
variable "object_id" {}

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = var.name
  object_id           = var.object_id
  description         = var.description
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1

  source_region_list {
    description_cn = "中国"
    description_en = "Chinese Mainland"
    region_id      = "CN"
    region_type    = 0
  }

  destination_addresses = ["1.1.1.2"]

  custom_services {
    protocol    = 6
    source_port = 81
    dest_port   = 82
  }

  sequence {
    top = 1
  }

  tags = {
    key = "value"
  }
}
```

### Create a rule with the custom service groups

```hcl
variable "name" {}
variable "description" {}
variable "object_id" {}
variable "service_group_id" {}
variable "protocol" {}

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = var.name
  object_id           = var.object_id
  description         = var.description
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1

  source_addresses      = ["1.1.1.1"] 
  destination_addresses = ["1.1.1.2"]

  custom_service_groups {
    protocols = [var.protocol]
    group_ids = [
      var.service_group_id
    ]
  }

  sequence {
    top = 1
  }

  tags = {
    key = "value"
  }
}
```

### Create a rule with any service

```hcl
variable "name" {}
variable "description" {}
variable "object_id" {}
variable "service_group_id" {}
variable "protocol" {}

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = var.name
  object_id           = var.object_id
  description         = var.description
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1

  source_addresses      = ["1.1.1.1"] 
  destination_addresses = ["1.1.1.2"]

  sequence {
    top = 1
  }

  tags = {
    key = "value"
  }
}
```

### Create a rule with any source address

```hcl
variable "name" {}
variable "description" {}
variable "object_id" {}
variable "service_group_id" {}
variable "protocol" {}

resource "huaweicloud_cfw_acl_rule" "test" {
  name                = var.name
  object_id           = var.object_id
  description         = var.description
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1

  destination_addresses = ["1.1.1.2"]

  custom_services {
    protocol    = 6
    source_port = 81
    dest_port   = 82
  }
  
  sequence {
    top = 1
  }

  tags = {
    key = "value"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the resource.
  If omitted, the provider-level region will be used.
  Changing this creates a new resource.

* `object_id` - (Required, String, NonUpdatable) The protected object ID.

* `name` - (Required, String) The rule name.

* `type` - (Required, Int) The rule type.
  The value can be `0` (Internet rule), `1` (VPC rule), or `2` (NAT rule).

* `action_type` - (Required, Int) The action type.
  The value can be `0` (allow), `1` (deny).

* `address_type` - (Required, Int) The address type.
  The value can be `0` (IPv4), `1` (IPv6).

* `applications` - (Optional, List) The application list.
  The valid value can be **HTTP**, **HTTPS**, **TLS1**, **DNS**, **SSH**, **MYSQL**, **SMTP**, **RDP**, **RDPS**,
  **VNC**, **POP3**, **IMAP4**, **SMTPS**, **POP3S**, **FTPS**, **ANY**, **BGP** and so on.

* `sequence` - (Required, List) The sequence configuration.
  The [sequence](#Sequence) structure is documented below.

* `status` - (Required, Int) The rule status. The options are as follows:
  + **0**: disabled;
  + **1**: enabled;

* `long_connect_enable` - (Required, Int) Whether to support persistent connections.

* `custom_services` - (Optional, List) The custom service configuration.
  The [custom_services](#CustomServices) structure is documented below.

* `custom_service_groups` - (Optional, List) The custom service group list.
  The [custom_service_groups](#CustomServiceGroups) structure is documented below.

* `predefined_service_groups` - (Optional, List) The predefined service group list.
  The [predefined_service_groups](#PredefinedServiceGroups) structure is documented below.

* `source_addresses` - (Optional, List) The source IP address list.

* `source_region_list` - (Optional, List) The source region list.
  The [source_region_list](#SourceRegionList) structure is documented below.

* `source_address_groups` - (Optional, List) The source address group list.

* `source_predefined_groups` - (Optional, List) The source predefined address group list.

* `source_address_type` - (Optional, Int) The source address type.
  The value can be `0` (IPv4), `1` (IPv6).

* `destination_addresses` - (Optional, List) The destination IP address list.

* `destination_region_list` - (Optional, List) The destination region list.
  The [destination_region_list](#DestinationRegionList) structure is documented below.

* `destination_domain_address_name` - (Optional, String) The destination domain address name.

* `destination_domain_group_id` - (Optional, String) The destination domain group ID.

* `destination_domain_group_name` - (Optional, String) The destination domain group name.

* `destination_domain_group_type` - (Optional, Int) The destination domain group type.
  The options are as follows:
  + **4**: application domain name group;
  + **6**: network domain name group;

* `destination_address_groups` - (Optional, List) The destination address group list.

* `destination_address_type` - (Optional, Int) The destination address type.
  The value can be `0` (IPv4), `1` (IPv6).

* `long_connect_time_hour` - (Optional, Int) The persistent connection duration (hour).

* `long_connect_time_minute` - (Optional, Int) The persistent connection duration (minute).

* `long_connect_time_second` - (Optional, Int) The persistent Connection Duration (second).

* `description` - (Optional, String) The rule description.

* `direction` - (Optional, Int) The rule direction. The options are as follows:
  + **0**: inbound;
  + **1**: outbound;

* `rule_hit_count` - (Optional, String) The number of times the ACL rule is hit.
  Setting the value to **0** will clear the hit count. Value options: **0**.

* `tags` - (Optional, Map) The key/value pairs to associate with the ACL rule.

<a name="Sequence"></a>
The `sequence` block supports:

* `bottom` - (Optional, Int) Whether to pin on bottom.
  The options are as follows:
  + **0**: no;
  + **1**: yes;

* `dest_rule_id` - (Optional, String) The ID of the rule that the added rule will follow.

* `top` - (Optional, Int) Whether to pin on top.
  The options are as follows:
  + **0**: no;
  + **1**: yes;

<a name="CustomServices"></a>
The `custom_services` block supports:

* `dest_port` - (Required, String) The destination port.

* `protocol` - (Required, Int) The protocol type.

* `source_port` - (Required, String) The source port.

<a name="CustomServiceGroups"></a>
The `custom_service_groups` block supports:

* `group_ids` - (Required, List) The IDs of the custom service groups.

* `protocols` - (Required, List) The protocols used in the custom service groups.

<a name="PredefinedServiceGroups"></a>
The `predefined_service_groups` block supports:

* `group_ids` - (Required, List) The IDs of the predefined service groups.

* `protocols` - (Required, List) The protocols used in the predefined service groups.

<a name="SourceRegionList"></a>
The `source_region_list` block supports:

* `region_id` - (Required, String) The region ID.

* `region_type` - (Required, Int) The region type.

* `description_cn` - (Optional, String) The Chinese description of the region.

* `description_en` - (Optional, String) The English description of the region.

<a name="DestinationRegionList"></a>
The `destination_region_list` block supports:

* `region_id` - (Required, String) The region ID.

* `region_type` - (Required, Int) The region type.

* `description_cn` - (Optional, String) The Chinese description of the region.

* `description_en` - (Optional, String) The English description of the region.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.

## Import

The ACL rule can be imported using `object_id`, `id`, separated by a slash, e.g.

```bash
$ terraform import huaweicloud_cfw_acl_rule.test <object_id>/<id>
```

Note that the imported state may not be identical to your resource definition, due to some attributes missing from the
API response, security or some other reason.
The missing attributes include: `sequence`, `type`, `predefined_service_groups` and `source_predefined_groups`.
It is generally recommended running `terraform plan` after importing the resource.
You can then decide if changes should be applied to the instance, or the resource definition should be updated to
align with the instance. Also you can ignore changes as below.

```hcl
resource "huaweicloud_cfw_acl_rule" "test" {
    ...

  lifecycle {
    ignore_changes = [
      sequence, type, predefined_service_groups, source_predefined_groups,
    ]
  }
}
```
