---
subcategory: "Cloud Application Engine (CAE)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_cae_component_action"
description: |-
  Use this resource to operate CAE component within HuaweiCloud.
---


# huaweicloud_cae_component_action

Use this resource to operate CAE component within HuaweiCloud.

-> This resource is only a one-time action resource for operating component. Deleting this resource will not clear
   the corresponding request record, but will only remove the resource information from the tfstate file.

## Example Usage

```hcl
variable "environment_id" {}
variable "application_id" {}
variable "component_id" {}
variable "upgraded_version" {}
variable "image_url" {}

resource "huaweicloud_cae_component_action" "test" {
  environment_id = var.environment_id
  application_id = var.application_id
  component_id   = var.component_id
  metadata {
    name = "upgrade"
    
    annotations = {
      version = var.upgraded_version
    }
  }

  spec = jsonencode({
    "source" : {
      "type" : "image",
      "url" : var.image_url
    },
    "resource_limit" : {
      "cpu_limit" : "500m",
      "memory_limit" : "2Gi"
    }
  })
}

```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the resource.
  If omitted, the provider-level region will be used.
  Changing this creates a new resource.

* `environment_id` - (Required, String, ForceNew) Specifies the ID of the environment where the application is located.
  Changing this creates a new resource.

* `application_id` - (Required, String, ForceNew) Specifies the ID of the application where the component is located.
  Changing this creates a new resource.

* `component_id` - (Required, String, ForceNew) Specifies the ID of the component being operated.
  Changing this creates a new resource.

* `metadata` - (Required, List) Specifies the metadata of this action request.
  The [metadata](#component_action_metadata) structure is documented below.

* `spec` - (Optional, String) Specifies the specification detail of the action, in JSON format.  
  Please following [reference documentation](https://support.huaweicloud.com/api-cae/ExecuteAction.html#ExecuteAction__request_ActionOnComponentSpec).

  -> If the `spec` parameter specified in this resource is inconsistent with the `huaweicloud_cae_component` resource,
     you can handle the changes in the `huaweicloud_cae_component` resource by `lifecycle.ignore_changes` or manual synchronization.

<a name="component_action_metadata"></a>
The `metadata` block supports:

* `name` - (Required, String) Specifies the action name.  
  The valid values are as follows:
  + **deploy**
  + **configure**
  + **upgrade**
  + **rollback**
  + **start**
  + **restart**
  + **stop**

* `annotations` - (Optional, Map) Specifies the key/value pairs parameters related to the component being operated.
  Currently, only `version` is supported.

  -> If the `annotations` parameter specified in this resource is inconsistent with the `huaweicloud_cae_component` resource,
     you can handle the changes in the `huaweicloud_cae_component` resource by `lifecycle.ignore_changes` or manual synchronization.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minutes.
* `update` - Default is 10 minutes.
