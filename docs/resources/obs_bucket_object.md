---
subcategory: "Object Storage Service (OBS)"
---

# huaweicloud_obs_bucket_object

Provides an OBS bucket object resource.

## Example Usage

### Uploading to a bucket

```hcl
resource "huaweicloud_obs_bucket_object" "object" {
  bucket       = "your_bucket_name"
  key          = "new_key_from_content"
  content      = "some object content"
  content_type = "application/xml"
}
```

### Uploading a file to a bucket

```hcl
resource "huaweicloud_obs_bucket" "examplebucket" {
  bucket = "examplebuckettftest"
  acl    = "private"
}

resource "huaweicloud_obs_bucket_object" "object" {
  bucket = huaweicloud_obs_bucket.examplebucket.bucket
  key    = "new_key_from_file"
  source = "index.html"
}
```

### Server Side Encryption with OBS Default Master Key

```hcl
resource "huaweicloud_obs_bucket_object" "examplebucket_object" {
  bucket     = "your_bucket_name"
  key        = "someobject"
  source     = "index.html"
  encryption = true
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the OBS bucket object resource. If omitted, the
  provider-level region will be used. Changing this creates a new OBS bucket object resource.

* `bucket` - (Required, String, ForceNew) The name of the bucket to put the file in.

* `key` - (Required, String, ForceNew) The name of the object once it is in the bucket.

* `source` - (Optional, String) The path to the source file being uploaded to the bucket.

* `content` - (Optional, String) The literal content being uploaded to the bucket.

* `acl` - (Optional, String) The ACL policy to apply. Defaults to `private`.

* `storage_class` - (Optioanl, String) Specifies the storage class of the object. Defaults to `STANDARD`.

* `content_type` - (Optional, String) A standard MIME type describing the format of the object data, e.g.
  application/octet-stream. All Valid MIME Types are valid for this input.

* `encryption` - (Optional, Bool) Whether enable server-side encryption of the object in SSE-KMS mode.

* `sse_kms_key_id` - (Optional, String) The ID of the kms key. If omitted, the default master key will be used.

* `etag` - (Optional, String) Specifies the unique identifier of the object content. It can be used to trigger updates.
  The only meaningful value is `md5(file("path_to_file"))`.

Either `source` or `content` must be provided to specify the bucket content. These two arguments are mutually-exclusive.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - the `key` of the resource supplied above.
* `etag` - the ETag generated for the object (an MD5 sum of the object content). When the object is encrypted on the
  server side, the ETag value is not the MD5 value of the object, but the unique identifier calculated through the
  server-side encryption.
* `size` - the size of the object in bytes.
* `version_id` - A unique version ID value for the object, if bucket versioning is enabled.

