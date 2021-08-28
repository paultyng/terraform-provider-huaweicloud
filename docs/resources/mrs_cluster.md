---
subcategory: "MapReduce Service (MRS)"
---

# huaweicloud\_mrs\_cluster

Manages resource cluster within HuaweiCloud MRS. It is recommend to use `huaweicloud_mapreduce_cluster`, which makes a
great improvement of managing MRS clusters.

## Example Usage: Creating an MRS cluster

```hcl
resource "huaweicloud_mrs_cluster" "cluster1" {
  cluster_name          = "mrs-cluster"
  cluster_version       = "MRS 1.8.10"
  cluster_type          = 0
  region                = "cn-north-1"
  available_zone_id     = "ae04cf9d61544df3806a3feeb401b204"
  billing_type          = 12
  master_node_num       = 2
  core_node_num         = 3
  master_node_size      = "c3.4xlarge.2.linux.bigdata"
  core_node_size        = "c3.xlarge.4.linux.bigdata"
  volume_type           = "SATA"
  volume_size           = 100
  safe_mode             = 0
  cluster_admin_secret  = var.admin_secret
  node_public_cert_name = var.keypair
  vpc_id                = var.vpc_id
  subnet_id             = var.subnet_id

  component_list {
    component_name = "Hadoop"
  }
  component_list {
    component_name = "Spark"
  }
  component_list {
    component_name = "Hive"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the msr cluster resource. If omitted, the
  provider-level region will be used. Changing this creates a new msr cluster resource.

* `billing_type` - (Required, Int, ForceNew) The value is 12, indicating on-demand payment.

* `cluster_name` - (Required, String, ForceNew) Cluster name, which is globally unique and contains only 1 to 64
  letters, digits, hyphens (-), and underscores (_).

* `cluster_version` - (Optional, String, ForceNew) Version of the clusters. Currently, MRS 1.8.10, MRS 1.9.2 and MRS
  2.1.0 are supported.

* `cluster_type` - (Optional, Int, ForceNew) Type of clusters. 0: analysis cluster 1: streaming cluster. The default
  value is 0.

* `master_node_num` - (Required, Int, ForceNew) Number of Master nodes. Set this parameter to 2 to enable cluster HA,
  set this parameter to 1 to enable cluster HA.

* `master_node_size` - (Required, String, ForceNew) Best match based on several years of commissioning experience. MRS
  supports specifications of hosts, and host specifications are determined by CPUs, memory, and disks space. MRS
  supports instance specifications detailed
  in [MRS specifications](https://support.huaweicloud.com/en-us/api-mrs/mrs_01_9006.html)

* `core_node_num` - (Required, Int, ForceNew) Number of Core nodes. Value range: 1 to 500.

* `core_node_size` - (Required, String, ForceNew) Instance specification of a Core node Configuration method of this
  parameter is identical to that of master_node_size.

* `available_zone_id` - (Required, String, ForceNew) ID of an available zone. The value as follows:

  CN North-Beijing1 AZ1 (cn-north-1a): ae04cf9d61544df3806a3feeb401b204<br>
  CN North-Beijing1 AZ2 (cn-north-1b): d573142f24894ef3bd3664de068b44b0<br>
  CN North-Beijing4 AZ1 (cn-north-4a): effdcbc7d4d64a02aa1fa26b42f56533<br>
  CN North-Beijing4 AZ2 (cn-north-4b): a0865121f83b41cbafce65930a22a6e8<br>
  CN North-Beijing4 AZ3 (cn-north-4c): 2dcb154ac2724a6d92e9bcc859657c1e<br>
  CN East-Shanghai1 AZ1 (cn-east-3a): e7afd64502d64fe3bfb60c2c82ec0ec6<br>
  CN East-Shanghai1 AZ2 (cn-east-3b): d90ff6d692954373bf53be49cf3900cb<br>
  CN East-Shanghai1 AZ3 (cn-east-3c): 2dafb4c708da4d509d0ad24864ae1c6d<br>
  CN East-Shanghai2 AZ1 (cn-east-2a): 72d50cedc49846b9b42c21495f38d81c<br>
  CN East-Shanghai2 AZ2 (cn-east-2b): 38b0f7a602344246bcb0da47b5d548e7<br>
  CN East-Shanghai2 AZ3 (cn-east-2c): 5547fd6bf8f84bb5a7f9db062ad3d015<br>
  CN South-Guangzhou AZ1 (cn-south-1a): 34f5ff4865cf4ed6b270f15382ebdec5<br>
  CN South-Guangzhou AZ2 (cn-south-2b): 043c7e39ecb347a08dc8fcb6c35a274e<br>
  CN South-Guangzhou AZ3 (cn-south-1c): af1687643e8c4ec1b34b688e4e3b8901<br>

* `vpc_id` - (Required, String, ForceNew) Specifies the VPC ID. Changing this parameter will create a new resource.

* `subnet_id` - (Required, String, ForceNew) Specifies the network id of a subnet. Changing this parameter will create a
  new resource.

* `volume_type` - (Required, String, ForceNew) Type of disks SATA and SSD are supported. SATA: common I/O; SAS
  high-speed I/O; SSD: super high-speed I/O.

* `volume_size` - (Required, Int, ForceNew) Data disk storage space of a Core node. Value range: 100 GB to 32000 GB

* `safe_mode` - (Required, Int, ForceNew) running mode of an MRS cluster.
  + 0: indicates that the Kerberos authentication is disabled. Users can use all functions provided by the cluster.
  + 1: indicates that the Kerberos authentication is enabled. Common users cannot use the file management or job
      management functions of an MRS cluster and cannot view cluster resource usage or the job records of Hadoop and
      Spark. To use these functions, the users must obtain the relevant permissions from the MRS Manager administrator.

* `cluster_admin_secret` - (Required, String, ForceNew) Indicates the password of the MRS Manager administrator. This
  parameter must meet the following requirements:
  + Must contain 8 to 26 characters.
  + Must contain at least three of the following: uppercase letters, lowercase letters, digits, and special
      characters: `~!@#$%^&*()-_=+\|[{}];:'",<.>/? and space.
  + Cannot be the username or the username spelled backwards.

* `node_password` - (Optional, String, ForceNew) Password of user **root** for logging in to a cluster node. This
  parameter and `node_public_cert_name` are alternative. A password must meet the following requirements:
  + Must be 8 to 26 characters.
  + Must contain at least three of the following: uppercase letters, lowercase letters, digits, and special
      characters (!@$%^-_=+[{}]:,./?), but must not contain spaces.
  + Cannot be the username or the username spelled backwards.

* `node_public_cert_name` - (Optional, String, ForceNew) Name of a key pair. You can use a key to log in to the Master
  node in the cluster. This parameter and `node_password` are alternative.

* `log_collection` - (Optional, Int, ForceNew) Indicates whether logs are collected when cluster installation fails. 0:
  not collected; 1: collected. The default value is 1, indicating that OBS buckets will be created and only used to
  collect logs that record MRS cluster creation failures.

* `component_list` - (Required, List, ForceNew) List of service components to be installed. Structure is documented
  below.

* `add_jobs` - (Optional, List, ForceNew) Jobs can be submitted when a cluster is created. Currently, only one job can
  be created. Structure is documented below.

* `tags` - (Optional) The key/value pairs to associate with the cluster.

The `component_list` block supports:

* `component_name` - (Required, String, ForceNew) Component name.
  + MRS 2.1.0 supports: Presto, Hadoop, Spark, HBase, Hive, Tez, Hue, Loader, Flink, Impala, Kudu, Flume, Kafka, and
      Storm;
  + MRS 1.9.2 supports: Presto, Hadoop, Spark, HBase, OpenTSDB, Hive, Hue, Loader, Tez, Flink, Alluxio, Ranger, Flume,
      Kafka, KafkaManager, and Storm;
  + MRS 1.8.10 supports: Presto, Hadoop, Spark, HBase, OpenTSDB, Hive, Hue, Loader, Flink, Flume, Kafka, KafkaManager,
      and Storm;

The `add_jobs` block supports:

* `job_type` - (Required, Int, ForceNew) Job type code. 1: MapReduce; 2: Spark; 3: Hive Script; 4: HiveQL
  (not supported currently); 5: DistCp, importing and exporting data (not supported currently); 6: Spark Script; 7:
  Spark SQL, submitting Spark SQL statements
  (not supported currently). NOTE: Spark and Hive jobs can be added to only clusters including Spark and Hive
  components.

* `job_name` - (Required, String, ForceNew) Job name. It contains 1 to 64 characters. Only letters, digits, hyphens (-),
  and underscores (_) are allowed. NOTE: Identical job names are allowed but not recommended.

* `jar_path` - (Required, String, ForceNew) Path of the .jar file or .sql file for program execution The parameter must
  meet the following requirements:
  + Contains a maximum of 1,023 characters, excluding special characters such as ;|&><'$. The parameter value cannot
      be empty or full of spaces.
  + Files can be stored in HDFS or OBS. The path varies depending on the file system. OBS: The path must start with
      s3a://. Files or programs encrypted by KMS are not supported. HDFS: The path starts with a slash (/).
  + Spark Script must end with .sql while MapReduce and Spark Jar must end with .jar. sql and jar are
      case-insensitive.

* `arguments` - (Optional, String, ForceNew) Key parameter for program execution. The parameter is specified by the
  function of the user's program. MRS is only responsible for loading the parameter. The parameter contains a maximum of
  2047 characters, excluding special characters such as ;|&>'<$, and can be empty.

* `input` - (Optional, String, ForceNew) Path for inputting data, which must start with / or s3a://. A correct OBS path
  is required. The parameter contains a maximum of 1023 characters, excluding special characters such as ;|&>'<$, and
  can be empty.

* `output` - (Optional, String, ForceNew) Path for outputting data, which must start with / or s3a://. A correct OBS
  path is required. If the path does not exist, the system automatically creates it. The parameter contains a maximum of
  1023 characters, excluding special characters such as ;|&>'<$, and can be empty.

* `job_log` - (Optional, String, ForceNew) Path for storing job logs that record job running status. This path must
  start with / or s3a://. A correct OBS path is required. The parameter contains a maximum of 1023 characters, excluding
  special characters such as ;|&>'<$, and can be empty.

* `shutdown_cluster` - (Optional, Bool, ForceNew) Whether to delete the cluster after the jobs are complete.

* `file_action` - (Optional, String, ForceNew) Data import and export. Valid values include: import, export.

* `submit_job_once_cluster_run` - (Required, Bool, ForceNew) Whether the job is submitted during the cluster creation or
  after the cluster is created.

* `hql` - (Optional, String, ForceNew) HiveQL statement.

* `hive_script_path` - (Optional, String, ForceNew) SQL program path This parameter is needed by Spark Script and Hive
  Script jobs only and must meet the following requirements:
  Contains a maximum of 1023 characters, excluding special characters such as ;|&><'$. The address cannot be empty or
  full of spaces. Starts with / or s3a://. Ends with .sql. sql is case-insensitive.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Indicates the MRS cluster ID.
* `available_zone_name` - Indicates the name of an availability zone.
* `component_list` - See Argument Reference below.
* `order_id` - Order ID for creating clusters.
* `instance_id` - Instance ID.
* `hadoop_version` - Hadoop version.
* `master_node_ip` - IP address of a Master node.
* `externalIp` - Internal IP address.
* `private_ip_first` - Primary private IP address.
* `external_ip` - External IP address.
* `slave_security_groups_id` - Standby security group ID.
* `security_groups_id` - Security group ID.
* `external_alternate_ip` - Backup external IP address.
* `master_node_spec_id` - Specification ID of a Master node.
* `core_node_spec_id` - Specification ID of a Core node.
* `master_node_product_id` - Product ID of a Master node.
* `core_node_product_id` - Product ID of a Core node.
* `duration` - Cluster subscription duration.
* `vnc` - URI address for remote login of the elastic cloud server.
* `fee` - Cluster creation fee, which is automatically calculated.
* `deployment_id` - Deployment ID of a cluster.
* `cluster_state` - Cluster status. Valid values include: starting, running, terminating, terminated, failed, abnormal,
  frozen, scaling-out, scaling-in.
* `create_at` - Cluster creation time.
* `update_at` - Cluster update time.
* `error_info` - Error information.
* `charging_start_time` - Time when charging starts.
* `remark` - Remarks of a cluster.

The `component_list` attributes supports:

* `component_id` - Indicates the component ID.
* `component_name` - Indicates the component name.
* `component_version` - Indicates the component version.
* `component_desc` - Indicates the component description.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 60 minute.
* `delete` - Default is 10 minute.

