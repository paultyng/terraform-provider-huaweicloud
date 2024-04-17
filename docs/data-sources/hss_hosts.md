---
subcategory: "Host Security Service (HSS)"
---

# huaweicloud_hss_hosts

Use this data source to get the list of HSS hosts within HuaweiCloud.

## Example Usage

```hcl
variable host_id {}

data "hss_hosts" "test" {
  host_id = var.host_id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the HSS hosts.
  If omitted, the provider-level region will be used.

* `host_id` - (Optional, String) Specifies the ID of the host to be queried.

* `name` - (Optional, String) Specifies the name of the host to be queried.
  This field will undergo a fuzzy matching query, the query result is for all hosts whose names contain this value.

* `status` - (Optional, String) Specifies the status of the hosts to be queried.  
  The valid values are as follows:
  + **ACTIVE**
  + **SHUTOFF**
  + **ERROR**

* `os_type` - (Optional, String) Specifies the operating system type of the hosts to be queried.  
  The valid values are as follows:
  + **Linux**
  + **Windows**

* `agent_status` - (Optional, String) Specifies the agent status of the hosts to be queried.  
  The valid values are as follows:
  + **installed**
  + **not_installed**
  + **online**
  + **offline**
  + **install_failed**

* `protect_status` - (Optional, String) Specifies the protection status of the hosts to be queried.  
  The valid values are as follows:
  + **closed**
  + **opened**

* `protect_version` - (Optional, String) Specifies the protection version enabled by the hosts to be queried.  
  The valid values are as follows:
  + **hss.version.basic**
  + **hss.version.advanced**
  + **hss.version.enterprise**
  + **hss.version.premium**
  + **hss.version.wtp**
  + **hss.version.container.enterprise**

* `protect_charging_mode` - (Optional, String) Specifies the charging mode for the hosts protection quota to be queried.
  
  The valid values are as follows:
  + **prePaid**
  + **postPaid**

* `detect_result` - (Optional, String) Specifies the security detection result of the hosts to be queried.  
  The valid values are as follows:
  + **undetected**
  + **clean**
  + **risk**

* `group_id` - (Optional, String) Specifies the host group ID of the hosts to be queried.

* `policy_group_id` - (Optional, String) Specifies the policy group ID of the hosts to be queried.

* `asset_value` - (Optional, String) Specifies the asset importance of the hosts to be queried.  
  The valid values are as follows:
  + **important**
  + **common**
  + **test**

* `enterprise_project_id` - (Optional, String) Specifies the ID of the enterprise project to which the hosts belong.
  If omitted, will query the hosts under all enterprise projects.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID in UUID format.

* `hosts` - All hosts that match the filter parameters.  
  The [hosts](#hss_hosts) structure is documented below.

<a name="hss_hosts"></a>
The `hosts` block supports:

* `id` - The ID of the host.

* `name` - The name of the host.

* `status` - The status of the host.

* `os_type` - The operating system type of the host.

* `agent_id` - The agent ID installed on the host.

* `agent_status` - The agent status of the host.

* `protect_status` - The protection status of the host.

* `protect_version` - The protection version enabled by the host.

* `protect_charging_mode` - The charging mode for the host protection quota.

* `quota_id` - The protection quota ID of the host.

* `detect_result` - The security detection result of the host.

* `group_id` - The host group ID to which the host belongs.

* `policy_group_id` - The policy group ID to which the host belongs.

* `asset_value` - The asset importance of the host.

* `open_time` - The time to enable host protection.

* `private_ip` - The private IP address of the host.

* `public_ip` - The elastic public IP address of the host.

* `asset_risk_num` - The number of asset risks in the host

* `vulnerability_risk_num` - The number of vulnerability risks in the host.

* `baseline_risk_num` - The number of baseline risks in the host.

* `intrusion_risk_num` - The number of intrusion risks in the host.

* `enterprise_project_id` - The ID of the enterprise project to which the host belongs.
