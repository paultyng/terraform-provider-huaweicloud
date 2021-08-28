---
subcategory: "Virtual Private Cloud (VPC)"
---

# huaweicloud_networking_secgroup

Use this data source to get the ID of an available HuaweiCloud security group.

## Example Usage

```hcl
data "huaweicloud_networking_secgroup" "secgroup" {
  name = "tf_test_secgroup"
}
```

## Argument Reference

* `region` - (Optional, String) Specifies the region in which to obtain the security group. If omitted, the
  provider-level region will be used.

* `secgroup_id` - (Optional, String) Specifiest he ID of the security group.

* `name` - (Optional, String) Specifies the name of the security group.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID in UUID format.

* `description`- The description of the security group.
