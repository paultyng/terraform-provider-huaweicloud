---
subcategory: "Distributed Message Service (DMS)"
---

# huaweicloud_dms_queue

Manages a DMS queue in the huaweicloud DMS Service. This is an alternative to `huaweicloud_dms_queue_v1`

## Example Usage

### Automatically detect the correct network

```hcl
resource "huaweicloud_dms_queue" "queue_1" {
  name              = "queue_1"
  description       = "test create dms queue"
  queue_mode        = "FIFO"
  redrive_policy    = "enable"
  max_consume_count = 80
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the DMS queue resource. If omitted, the
  provider-level region will be used. Changing this creates a new DMS queue resource.

* `name` - (Required, String, ForceNew) Indicates the unique name of a queue. A string of 1 to 64 characters that
  contain a-z, A-Z, 0-9, hyphens (-), and underscores (_). The name cannot be modified once specified.

* `queue_mode` - (Optional, String, ForceNew) Indicates the queue type. It only support 'NORMAL' and 'FIFO'. NORMAL:
  Standard queue. Best-effort ordering. Messages might be retrieved in an order different from which they were sent.
  Select standard queues when throughput is important. FIFO: First-ln-First-out (FIFO) queue. FIFO delivery. Messages
  are retrieved in the order they were sent. Select FIFO queues when the order of messages is important. Default value:
  NORMAL.

* `description` - (Optional, String, ForceNew) Indicates the basic information about a queue. The queue description must
  be 0 to 160 characters in length, and does not contain angle brackets (<) and (>).

* `redrive_policy` - (Optional, String, ForceNew) Indicates whether to enable dead letter messages. Dead letter messages
  indicate messages that cannot be normally consumed. The redrive_policy should be set to 'enable' or 'disable'. The
  default value is 'disable'.

* `max_consume_count` - (Optional, Int, ForceNew) This parameter is mandatory only when redrive_policy is set to enable.
  This parameter indicates the maximum number of allowed message consumption failures. When a message fails to be
  consumed after the number of consumption attempts of this message reaches this value, DMS stores this message into the
  dead letter queue. The max_consume_count value range is 1–100.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.
* `created` - Indicates the time when a queue is created.
* `reservation` - Indicates the retention period (unit: min) of a message in a queue.
* `max_msg_size_byte` - Indicates the maximum message size (unit: byte) that is allowed in queue.
* `produced_messages` - Indicates the total number of messages (not including the messages that have expired and been
  deleted) in a queue.
* `group_count` - Indicates the total number of consumer groups in a queue.
