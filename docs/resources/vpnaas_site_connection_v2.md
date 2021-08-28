---
subcategory: "Deprecated"
---

# huaweicloud_vpnaas_site_connection_v2

Manages a V2 IPSec site connection resource within HuaweiCloud.

## Example Usage

```hcl
resource "huaweicloud_vpnaas_site_connection_v2" "conn_1" {
  name              = "connection_1"
  ikepolicy_id      = huaweicloud_vpnaas_ike_policy_v2.policy_2.id
  ipsecpolicy_id    = huaweicloud_vpnaas_ipsec_policy_v2.policy_1.id
  vpnservice_id     = huaweicloud_vpnaas_service_v2.service_1.id
  psk               = "secret"
  peer_address      = "192.168.10.1"
  local_ep_group_id = huaweicloud_vpnaas_endpoint_group_v2.group_2.id
  peer_ep_group_id  = huaweicloud_vpnaas_endpoint_group_v2.group_1.id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional) The region in which to create the V2 Networking client. A Networking client is needed to create
  an IPSec site connection. If omitted, the
  `region` argument of the provider is used. Changing this creates a new site connection.

* `name` - (Optional) The name of the connection. Changing this updates the name of the existing connection.

* `description` - (Optional) The human-readable description for the connection. Changing this updates the description of
  the existing connection.

* `admin_state_up` - (Optional) The administrative state of the resource. Can either be up(true) or down(false).
  Defaults to `true`. Changing this updates the administrative state of the existing connection.

* `ikepolicy_id` - (Required) The ID of the IKE policy. Changing this creates a new connection.

* `vpnservice_id` - (Required) The ID of the VPN service. Changing this creates a new connection.

* `local_ep_group_id` - (Optional) The ID for the endpoint group that contains private subnets for the local side of the
  connection. You must specify this parameter with the peer_ep_group_id parameter unless in backward- compatible mode
  where peer_cidrs is provided with a subnet_id for the VPN service. Changing this updates the existing connection.

* `ipsecpolicy_id` - (Required) The ID of the IPsec policy. Changing this creates a new connection.

* `peer_id` - (Required) The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address,
  e-mail address, key ID, or FQDN. Typically, this value matches the peer_address value. Changing this updates the
  existing policy.

* `peer_ep_group_id` - (Optional) The ID for the endpoint group that contains private CIDRs in the form < net_address >
  / < prefix > for the peer side of the connection. You must specify this parameter with the local_ep_group_id parameter
  unless in backward-compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.

* `local_id` - (Optional) An ID to be used instead of the external IP address for a virtual router used in traffic
  between instances on different networks in east-west traffic. Most often, local ID would be domain name, email
  address, etc. If this is not configured then the external IP address will be used as the ID.

* `peer_address` - (Required) The peer gateway public IPv4 or IPv6 address or FQDN.

* `psk` - (Required) The pre-shared key. A valid value is any string.

* `initiator` - (Optional) A valid value is response-only or bi-directional. Default is bi-directional.

* `peer_cidrs` - (Optional) Unique list of valid peer private CIDRs in the form < net_address > / < prefix > .

* `dpd` - (Optional) A dictionary with dead peer detection (DPD) protocol controls.
  + `action` - (Optional) The dead peer detection (DPD) action. A valid value is clear, hold, restart, disabled, or
      restart-by-peer. Default value is hold.

  + `timeout` - (Optional) The dead peer detection (DPD) timeout in seconds. A valid value is a positive integer that
      is greater than the DPD interval value. Default is 120.

  + `interval` - (Optional) The dead peer detection (DPD) interval, in seconds. A valid value is a positive integer.
      Default is 30.

* `mtu` - (Optional) The maximum transmission unit (MTU) value to address fragmentation. Minimum value is 68 for IPv4,
  and 1280 for IPv6.

* `value_specs` - (Optional) Map of additional options.

* `tags` - (Optional) The key/value pairs to associate with the connection.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `update` - Default is 10 minute.
* `delete` - Default is 10 minute.

## Import

Site Connections can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_vpnaas_site_connection_v2.conn_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
```
