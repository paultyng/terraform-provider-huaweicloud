---
subcategory: "Log Tank Service (LTS)"
---

# huaweicloud_lts_stream

Manage a log stream resource within HuaweiCloud.

## Example Usage

### create a log stream

```hcl
resource "huaweicloud_lts_group" "test_group" {
  group_name  = "test_group"
  ttl_in_days = 1
}
resource "huaweicloud_lts_stream" "test_stream" {
  group_id    = huaweicloud_lts_group.test_group.id
  stream_name = "testacc_stream"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the log stream resource. If omitted, the
  provider-level region will be used. Changing this creates a new log stream resource.

* `group_id` - (Required, String, ForceNew) Specifies the ID of a created log group. Changing this parameter will create
  a new resource.

* `stream_name` - (Required, String, ForceNew) Specifies the log stream name. Changing this parameter will create a new
  resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The log stream ID.

* `filter_count` - Number of log stream filters.

## Import

Log stream can be imported using the lts group ID and stream ID separated by a slash, e.g.

```
$ terraform import huaweicloud_lts_stream.stream_1 393f2bfd-2244-11ea-adb7-286ed488c87f/72855918-20b1-11ea-80e0-286ed488c880
```
