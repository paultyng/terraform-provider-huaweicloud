---
subcategory: "Domain Name Service (DNS)"
---

# huaweicloud_dns_nameservers

Use this data source to get the list of DNS name servers.

## Example Usage

```hcl
data "huaweicloud_dns_nameservers" "test" {
  type = "public"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the resource.
  If omitted, the provider-level region will be used.

* `type` - (Optional, String) Specifies the type of the name server.
  The valid values are as follows:
   + **public**
   + **private**

* `server_region` - (Optional, String) Specifies the region to which the name server belongs.
  This parameter cannot be set when `type` is `public`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `nameservers` - All name servers that match the filter parameters.

  The [nameservers](#nameservers_struct) structure is documented below.

<a name="nameservers_struct"></a>
The `nameservers` block supports:

* `type` - The type of the name server.

* `region` - The region where the name server is located.

* `ns_records` - The list of name servers.

  The [ns_records](#nameservers_ns_records_struct) structure is documented below.

<a name="nameservers_ns_records_struct"></a>
The `ns_records` block supports:

* `hostname` - The host name of the public name server.
  If `type` is set to `private`, the value is an empty string.

* `address` - The  address of the private name server.
  If `type` is set to `public`, the value is an empty string.

* `priority` - The priority of  the name server.
  The smaller value means the higher priority.
