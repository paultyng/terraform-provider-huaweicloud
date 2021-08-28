---
subcategory: "Distributed Cache Service"
---

# huaweicloud_dcs_product

Use this data source to get the ID of an available Flexibleengine dcs product.

## Example Usage

```hcl
data "huaweicloud_dcs_product" "product1" {
  spec_code = "redis.cluster.xu1.large.r2.4"
}
```

## Argument Reference

* `region` - (Optional, String) Specifies the region in which to obtain the dcs products. If omitted, the provider-level
  region will be used.

* `spec_code` - (Optional, String) Specifies the DCS instance specification code. You can query the code as follows:
  + Query the specifications
      in [DCS Instance Specifications](https://support.huaweicloud.com/intl/en-us/productdesc-dcs/dcs-pd-200713003.html)
  + Log in to the DCS console, click *Buy DCS Instance*, and find the corresponding instance specification.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - A data source ID in UUID format.
* `engine` - The cache engine. The value is *redis* or *memcached*.
* `engine_version` - The supported versions of a cache engine.
* `cache_mode` - The mode of a cache engine. The value is one of *single*, *ha*, *cluster*,
  *proxy* and *ha_rw_split*.
