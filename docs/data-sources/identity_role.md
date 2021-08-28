---
subcategory: "Identity and Access Management (IAM)"
---

# huaweicloud_identity_role

Use this data source to get the ID of an IAM **system-defined** role or policy.

The Role in Terraform is the same as Policy. We can get all **System-Defined Policies** form
[HuaweiCloud](https://support.huaweicloud.com/intl/en-us/usermanual-permissions/iam_01_0001.html). Please refer to the
following table to configuration:

Display Name | Role/Policy Name | Description
---- | --- | ---
Server Administrator | server_adm | Server Administrator
ECS FullAccess | system_all_3 | All permissions of ECS service
ECS CommonOperations | system_all_1 | Permissions for basic ECS operations,such as start,stop restart a ECS,query ECS details,automatic recovery of ECSs and so on
ECS ReadOnlyAccess | system_all_8 | The read-only permissions to all ECS resources, which can be used for statistics and survey
IMS Administrator | ims_adm | IMS Administrator
IMS FullAccess | system_all_16 | All permissions of Image Management Service
IMS ReadOnlyAccess | system_all_22 | The read-only permissions to all IMS resources, which can be used for statistics and survey
AutoScaling Administrator | as_adm | Auto Scaling administrator with full permissions
AutoScaling FullAccess | system_all_23 | Full permissions for Auto Scaling
AutoScaling ReadOnlyAccess | system_all_13 | Read-only permissions for Auto Scaling
EVS FullAccess | system_all_6 | Full permissions for Elastic Volume Service, including creating, expanding, attaching, detaching, querying, and deleting EVS disks
EVS ReadOnlyAccess | system_all_2 | Read-only permissions for Elastic Volume Service
SFS Administrator | sfs_adm | Scalable File Service Administrator
SFS FullAccess | system_all_58 | All permissions of SFS service
SFS ReadOnlyAccess | system_all_57 | The read-only permissions to all SFS resources
SFS Turbo FullAccess | system_all_99 | All permissions of SFS Turbo resources
SFS Turbo ReadOnlyAccess | system_all_98 | The read-only permissions to all SFS Turbo resources
OBS Administrator | system_all_159 | Object Storage Service Administrator
OBS OperateAccess | system_all_72 | Basic operation permissions to view the bucket list, obtain bucket metadata, list objects in a bucket, query bucket location, upload objects, download objects, delete objects, and obtain object ACLs
OBS ReadOnlyAccess | system_all_64 | Permissions to view the bucket list, obtain bucket metadata, list objects in a bucket, and query bucket location
OBS Buckets Viewer | obs_b_list | Permissions to view the bucket list, obtain bucket metadata, and query bucket location
CSBS Administrator | csbs_adm | Cloud Server Backup Service Administrator
SDRS Administrator | sdrs_adm | Storage Disaster Recovery Service Administrator
VPC Administrator | vpc_netadm | VPC Administrator
VPC FullAccess | system_all_7 | All permissions of VPC service
VPC ReadOnlyAccess | system_all_5 | The read-only permissions to all VPC resources, which can be used for statistics and survey
ELB Administrator | elb_adm | Elastic Load Balance administrator with full permissions for this service
ELB FullAccess | system_all_56 | All permissions of ELB service
ELB ReadOnlyAccess | system_all_55 | Read-only permissions for Elastic Load Balance
DNS Administrator | dns_adm | DNS Administrator
DNS FullAccess | system_all_102 | Allow users to perform all operations, including creating, deleting, querying, and modifying DNS resources
DNS ReadOnlyAccess | system_all_103 | Read-only permissions, which only allow users to query DNS resources
NAT Administrator | nat_adm | NAT Gateway administrator with full permissions for this service
NAT FullAccess | system_all_75 | All permissions of NAT Gateway service
NAT ReadOnlyAccess | system_all_76 | The read-only permissions to all NAT Gateway resources
VPCEndpoint Administrator | vpcep_adm | VPCEndpoint service enables you to privately connect your VPC to supported services
RDS Administrator | rds_adm | RDS Administrator
RDS FullAccess | system_all_14 | Full permissions for Relational Database Service
RDS ReadOnlyAccess | system_all_12 | Read-only permissions for Relational Database Service
DDS Administrator | dds_adm | Document Database Service Administrator
CCE Administrator | cce_adm | CCE Administrator
CCE FullAccess | system_all_32 | Common operation permissions on CCE cluster resources
CCE ReadOnlyAccess | system_all_31 | Permissions to view CCE cluster resources, excluding the namespace-level permissions of the clusters (with Kubernetes RBAC enabled)
CSS FullAccess | system_all_153 | All permissions for Cloud Search Service
CSS ReadOnlyAccess | system_all_154 | Read-only permissions for viewing Cloud Search Service
ServiceStage Administrator | svcstg_adm | ServiceStage administrator, who has full permissions for this service
ServiceStage Developer | svcstg_dev | ServiceStage developer, who has full permissions for this service but does not have the permission for creating infrastructure
ServiceStage Operator | svcstg_opr | ServiceStage operator, who has the read-only permission for this service
Anti-DDoS Administrator | ddos_adm | Anti-DDoS Administrator
APM Administrator | apm_adm | Application Performance Monitor Admin
BCS Administrator | bcs_adm | BlockChain Service Administrator
CES Administrator | ces_adm | CES Administrator
CS Tenant Admin | cs_adm | Cloud Stream Service Tenant Administrator, can manage multiple CS users
CS Tenant User | cs_user | Cloud Stream Service User
CTS Administrator | cts_adm | CTS Administrator
DCS Administrator | dcs_admin | Distributed Cache Service Administrator
DIS Administrator | dis_adm | DIS Administrator
KMS Administrator | kms_adm | KMS Administrator
MRS Administrator | mrs_adm | MRS Administrator
SWR Admin | swr_adm | Software Repository Admin
SMN Administrator | smn_adm | SMN Administrator
TMS Administrator | tms_adm | Tag Management Service Administrator
Security Administrator | secu_admin | Full permissions for Identity and Access Management
Tenant Administrator | te_admin | Tenant Administrator (Exclude IAM)
Tenant Guest | readonly | Tenant Guest (Exclude IAM)
EPS FullAccess | system_all_10 | All operations on the Enterprise Project Management service
FullAccess | system_all_1001 | Full permissions for all services that support policy-based authorization

## Example Usage

```hcl
data "huaweicloud_identity_role" "kms_adm" {
  display_name = "KMS Administrator"
}
```

## Argument Reference

* `display_name` - (Optional, String; Required if `name` is empty) Specifies the display name of the role displayed on
  the console.

* `name` - (Optional, String; Required if `display_name` is empty) Specifies the name of the role for internal use.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID in UUID format.
* `description` - The description of the policy.
* `catalog` - The service catalog of the policy.
* `type` - The display mode of the policy.
* `policy` - The content of the policy.
