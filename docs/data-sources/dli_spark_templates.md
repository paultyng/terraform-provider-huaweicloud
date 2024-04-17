---
subcategory: "Data Lake Insight (DLI)"
---

# huaweicloud_dli_spark_templates

Use this data source to get the list of the DLI spark templates.

## Example Usage

```hcl
variable "template_id" {}

data "huaweicloud_dli_spark_templates" "test" {
  template_id = var.template_id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the resource.
  If omitted, the provider-level region will be used.

* `template_id` - (Optional, String) Specifies the ID of the spark template to be queried.

* `name` - (Optional, String) Specifies the name of the spark template to be queried.

* `group` - (Optional, String) Specifies the group name to which the spark templates belong.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `templates` - All templates that match the filter parameters.

  The [templates](#templates_struct) structure is documented below.

<a name="templates_struct"></a>
The `templates` block supports:

* `description` - The description of template.

* `id` - The ID of template.

* `name` - The name of template.

* `body` - The body of template.

  The [body](#templates_body_struct) structure is documented below.

* `group` - The group of template.

<a name="templates_body_struct"></a>
The `body` block supports:

* `app_parameters` - The input parameters of the main class, that is application parameters.

* `name` - The Spark job name.

* `obs_bucket` - The OBS bucket for storing the Spark jobs.

* `auto_recovery` - The whether to enable the retry function.

* `max_retry_times` - The maximum number of retries.

* `jars` - The name of the resource package of type jar upload to the DLI resource management system.

* `dependent_packages` - The list of package resource objects.

  The [dependent_packages](#body_dependent_packages_struct) structure is documented below.

* `configurations` - The configuration items of the DLI Spark.

* `executor_memory` - The executor memory of the Spark application.

* `queue_name` - The DLI queue name.

* `main_class` - The Spark main class of the template.

* `python_files` - The name of the resource package of type pyFile upload to the DLI resource management system.

* `files` - The name of the resource package of type file upload to the DLI resource management system.

* `modules` - The name of the dependent system resource module.
  DLI provides dependencies for executing datasource jobs.
  The dependent modules and corresponding services are as follows.
  + **sys.datasource.hbase**: CloudTable/MRS HBase
  + **sys.datasource.opentsdb**: CloudTable/MRS OpenTSDB
  + **sys.datasource.rds**: RDS MySQL
  + **sys.datasource.css**: CSS

* `driver_memory` - The driver memory of the Spark application.

* `executor_cores` - The number of CPU cores of each executor in the Spark application.

* `app_name` - The name of the uploaded JAR or pyFile type package.

* `specification` - The compute resource type. Currently, resource types A, B, and C are available.  
  The available types and related specifications are as follows, default to minimum configuration (type **A**).
  | type | resource | driver cores | executor cores | driver memory | executor memory | num executor |
  | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
  | A | 8 vCPUs, 32-GB memory | 2 | 1 | 7G | 4G | 6 |
  | B | 16 vCPUs, 64-GB memory | 2 | 2 | 7G | 8G | 7 |
  | C | 32 vCPUs, 128-GB memory | 4 | 2 | 12G | 8G | 14 |

* `resources` - The list of resource objects.

  The [resources](#body_resources_struct) structure is documented below.

* `driver_cores` - The number of CPU cores of the Spark application driver.

* `num_executors` - The number of Executors in a Spark application.

<a name="body_dependent_packages_struct"></a>
The `dependent_packages` block supports:

* `name` - The name of a user group.

* `resources` - The resources of a user group.

  The [resources](#dependent_packages_resources_struct) structure is documented below.

<a name="dependent_packages_resources_struct"></a>
The `resources` block supports:

* `name` - The name of resource.

* `type` - The type of resource.

<a name="body_resources_struct"></a>
The `resources` block supports:

* `name` - The name of resource.

* `type` - The type of resource.
