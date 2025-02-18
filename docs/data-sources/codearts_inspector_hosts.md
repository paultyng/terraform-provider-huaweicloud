---
subcategory: "CodeArts Inspector"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_codearts_inspector_hosts"
description: |-
  Use this data source to get the list of CodeArts inspector hosts.
---

# huaweicloud_codearts_inspector_hosts

Use this data source to get the list of CodeArts inspector hosts.

## Example Usage

```hcl
data "huaweicloud_codearts_inspector_hosts" "test" {}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Optional, String) Specifies the host group ID.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `hosts` - Indicates the host list.

  The [hosts](#hosts_struct) structure is documented below.

<a name="hosts_struct"></a>
The `hosts` block supports:

* `id` - Indicates the host ID.

* `name` - Indicates the host name.

* `os_type` - Indicates the host os type.

* `ip` - Indicates the host IP.

* `ssh_credential_id` - Indicates the host ssh credential ID

* `jumper_server_id` - Indicates the jumper server ID.

* `smb_credential_id` - Indicates the smb credential ID.

* `group_id` - Indicates the host group ID.

* `auth_status` - Indicates the auth status.
  Value can be as follows:
  + **-1**: unknown
  + **0**: connected
  + **1**: unreachable
  + **2**: login failed

* `last_scan_id` - Indicates the last scan ID.

* `last_scan_info` - Indicates the last scan info.

  The [last_scan_info](#hosts_last_scan_info_struct) structure is documented below.

<a name="hosts_last_scan_info_struct"></a>
The `last_scan_info` block supports:

* `status` - Indicates the scan status.
  Value can be as follows:
  + **0**: running
  + **1**: completed
  + **2**: cancel
  + **3**: waiting
  + **4**: failed
  + **5**: scheduled

* `enable_weak_passwd` - Indicates whether weak password check enabled.

* `create_time` - Indicates the scan task create time.

* `end_time` - Indicates the scan task end time.

* `progress` - Indicates the task progress.

* `reason` - Indicates the task description.
