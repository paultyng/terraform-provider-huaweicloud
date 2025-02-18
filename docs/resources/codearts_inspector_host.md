---
subcategory: "CodeArts Inspector"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_codearts_inspector_host"
description: |-
  Manages a CodeArts inspector host resource within HuaweiCloud.
---
# huaweicloud_codearts_inspector_host

Manages a CodeArts inspector host resource within HuaweiCloud.

## Example Usage

```hcl
variable "name" {}
variable "ip" {}
variable "ssh_credential_id" {}

resource "huaweicloud_codearts_inspector_host" "test" {
  name              = var.name
  ip                = var.ip
  os_type           = "linux"
  ssh_credential_id = var.ssh_credential_id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String, ForceNew) Specifies the host name.
  Changing this creates a new resource.

* `ip` - (Required, String, ForceNew) Specifies the host IP.
  Changing this creates a new resource.

* `os_type` - (Required, String, ForceNew) Specifies the host os type. Valid values are **windows** and **linux**.
  Changing this creates a new resource.

* `group_id` - (Optional, String, ForceNew) Specifies the host group ID.
  Changing this creates a new resource.

* `jumper_server_id` - (Optional, String, ForceNew) Specifies the jumper server ID. Only available for **linux** host.
  Changing this creates a new resource.

* `smb_credential_id` - (Optional, String, ForceNew) Specifies the smb credential ID for **windows** host.
  Changing this creates a new resource.

* `ssh_credential_id` - (Optional, String, ForceNew) Specifies the host ssh credential ID for **linux** host.
  Changing this creates a new resource.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.

* `auth_status` - Indicates the auth status.
  Value can be as follows:
  + **-1**: unknown
  + **0**: connected
  + **1**: unreachable
  + **2**: login failed

* `last_scan_id` - Indicates the last scan ID.

* `last_scan_info` - Indicates the last scan informations.
  The [last_scan_info](#attrblock--last_scan_info) structure is documented below.

<a name="attrblock--last_scan_info"></a>
The `last_scan_info` block supports:

* `enable_weak_passwd` - Indicates whether weak password check enabled.

* `progress` - Indicates the task progress.

* `reason` - Indicates the task description.

* `create_time` - Indicates the scan task create time.

* `end_time` - Indicates the scan task end time.

* `status` - Indicates the task status.
  Value can be as follows:
  + **0**: running
  + **1**: completed
  + **2**: cancel
  + **3**: waiting
  + **4**: failed
  + **5**: scheduled

## Import

The host can be imported using `id`, e.g.

```bash
$ terraform import huaweicloud_codearts_inspector_host.test <id>
```
