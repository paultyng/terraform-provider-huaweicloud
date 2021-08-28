---
subcategory: "MapReduce Service (MRS)"
---

# huaweicloud_mapreduce_job

Manage a job resource within HuaweiCloud MRS.

## Example Usage

```hcl
variable "cluster_id" {}
variable "job_name" {}
variable "program_path" {}
variable "access_key" {}
variable "secret_key" {}

resource "huaweicloud_mapreduce_job" "test" {
  cluster_id   = var.cluster_id
  type         = "SparkSubmit"
  name         = var.job_name
  program_path = var.program_path
  parameters   = "${var.access_key} ${var.secret_key} 1 obs://obs-demo-analysis/input obs://obs-demo-analysis/output"

  program_parameters = {
    "--class" = "com.huawei.bigdata.spark.examples.DriverBehavior"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the MapReduce job resource. If
  omitted, the provider-level region will be used. Changing this will create a new MapReduce job resource.

* `cluster_id` - (Required, String, ForceNew) Specifies an ID of the MapReduce cluster to which the job belongs to.
  Changing this will create a new MapReduce job resource.

* `name` - (Required, String, ForceNew) Specifies the name of the MapReduce job. The name can contain 1 to 64
  characters, which may consist of letters, digits, underscores (_) and hyphens (-). Changing this will create a new
  MapReduce job resource.

* `type` - (Required, String, ForceNew) Specifies the job type. The valid values are as <span id="jump">follows</span>:
  + [Flink](https://support.huaweicloud.com/intl/en-us/usermanual-mrs/mrs_01_0527.html)
  + [HiveSql](https://support.huaweicloud.com/intl/en-us/usermanual-mrs/mrs_01_0525.html)
  + [HiveScript](https://support.huaweicloud.com/intl/en-us/usermanual-mrs/mrs_01_0525.html)
  + [MapReduce](https://support.huaweicloud.com/intl/en-us/usermanual-mrs/mrs_01_0052.html)
  + [SparkSubmit](https://support.huaweicloud.com/intl/en-us/usermanual-mrs/mrs_01_0524.html)
  + [SparkSql](https://support.huaweicloud.com/intl/en-us/usermanual-mrs/mrs_01_0526.html)
  + [SparkScript](https://support.huaweicloud.com/intl/en-us/usermanual-mrs/mrs_01_0526.html)

  Changing this will create a new MapReduce job resource.

  -> **NOTE:** Spark and Hive jobs can be added to only clusters including Spark and Hive components.

* `program_path` - (Optional, String, ForceNew) Specifies the .jar package path or .py file path for program execution.
  The parameter must meet the following requirements:
  + Contains a maximum of 1023 characters, excluding special characters such as `;|&><'$`.
  + The address cannot be empty or full of spaces.
  + The program support OBS or DHFS to storage program file or package. For OBS, starts with (OBS:) **obs://** and end
      with **.jar** or **.py**. For DHFS, starts with (DHFS:) **/user**.

  Required if `type` is __MapReduce__ or __SparkSubmit__. Changing this will create a new MapReduce job resource.

* `parameters` - (Optional, String, ForceNew) Specifies the parameters for the MapReduce job. Add an at sign (@) before
  each parameter can prevent the parameters being saved in plaintext format. Each parameters are separated with spaces.
  This parameter can be set when `type` is __Flink__, __MapReduce__ or __SparkSubmit__. Changing this will create a new
  MapReduce job resource.

* `program_parameters` - (Optional, Map, ForceNew) Specifies the the key/value pairs of the program parameters, such as
  thread, memory, and vCPUs, are used to optimize resource usage and improve job execution performance. This parameter
  can be set when `type` is __Flink__, __SparkSubmit__, __SparkSql__, __SparkScript__, __HiveSql__ or
  __HiveScript__. Refer to the documents for each [type](#jump) of support key-values. Changing this will create a new
  MapReduce job resource.

* `service_parameters` - (Optional, Map, ForceNew) Specifies the key/value pairs used to modify service configuration.
  Parameter configurations of services are available on the Service Configuration tab page of MapReduce Manager.
  Changing this will create a new MapReduce job resource.

* `sql` - (Optional, String, ForceNew) Specifies the SQL command or file path. Only required if `type` is __HiveSql__
  or __SparkSql__. Changing this will create a new MapReduce job resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the MapReduce job in UUID format.
* `status` - Status of the MapReduce job.
* `start_time` - The creation time of the MapReduce job, in 13-bit timestamp format.
* `submit_time` - The submission time of the MapReduce job, in 13-bit timestamp format.
* `finish_time` - The completion time of the MapReduce job, in 13-bit timestamp format.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 20 minute.

## Import

MapReduce jobs can be imported using their `id` and the IDs of the MapReduce cluster to which the job belongs, separated
by a slash, e.g.

```
$ terraform import huaweicloud_mapreduce_job.test <cluster_id>/<id>
```
