---
page_title: "Configure Remote State Backend for HuaweiCloud"
---

# Configure Remote State Backend for HuaweiCloud

## [Terraform Remote State](https://www.terraform.io/docs/language/state/remote.html)

By default, Terraform stores state locally in a file named `terraform.tfstate`. When working with Terraform in a team,
use of a local file makes Terraform usage complicated because each user must make sure they always have the latest state
data before running Terraform and make sure that nobody else runs Terraform at the same time.

With *remote* state, Terraform writes the state data to a remote data store, which can then be shared between all
members of a team. Terraform supports storing state in Terraform Cloud, HashiCorp Consul, Amazon S3, Azure Blob Storage,
Google Cloud Storage, etcd, and more.

Remote state is implemented by a [backend](https://www.terraform.io/docs/language/settings/backends/index.html).
Backends are configured with a nested `backend` block within the top-level `terraform` block:

```
terraform {
  backend "s3" {
    ...
  }
}
```

There are some important limitations on backend configuration:
* A configuration can only provide one backend block.
* A backend block cannot refer to **named values** (like input variables, locals, or data source attributes).

## Configuration Backend for HuaweiCloud

As HuaweiCloud OBS (Object Storage Service) can be compatible with the AWS S3 interface, and
[Amazon S3](https://www.terraform.io/docs/language/settings/backends/s3.html) backend supports custom endpoints,
we can use S3 backend to store state files in OBS.

Although the terraform block does not accept variables or locals and all backend configuration values must be hardcoded,
you can provide the credentials via the **AWS_ACCESS_KEY_ID** and **AWS_SECRET_ACCESS_KEY** environment variables to access
OBS, respectively.

```bash
export AWS_ACCESS_KEY_ID="your accesskey"
export AWS_SECRET_ACCESS_KEY="your secretkey"
```

The backend configuration as follows:

```hcl
terraform {
  backend "s3" {
    bucket   = "terraformbucket"
    key      = "terraform.tfstate"
    region   = "cn-north-1"
    endpoint = "https://obs.cn-north-1.myhuaweicloud.com"

    skip_region_validation      = true
    skip_credentials_validation = true
    skip_metadata_api_check     = true
  }
}
```

### Argument Reference

The following arguments are supported:

* `access_key` - (Optional) Specifies the access key of the HuaweiCloud cloud to use. This can also be sourced
  from the *AWS_ACCESS_KEY_ID* environment variable, AWS shared credentials file (e.g. ~/.aws/credentials), or
  AWS shared configuration file (e.g. ~/.aws/config).

* `secret_key` - (Optional) Specifies the secret key of the HuaweiCloud cloud to use. This can also be sourced
  from the *AWS_SECRET_ACCESS_KEY* environment variable, AWS shared credentials file (e.g. ~/.aws/credentials), or
  AWS shared configuration file (e.g. ~/.aws/config).

* `bucket` - (Required) Specifies the bucket name where to store the state. Make sure to create it before.

* `key` - (Required) Specifies the path to the state file inside the bucket.

* `region` - (Required) Specifies the region where the bucket is located. This can also be sourced from the
  *AWS_DEFAULT_REGION* and *AWS_REGION* environment variables.

* `endpoint` - (Required) Specifies the endpoint for HuaweiCloud OBS. The value is `https://obs.{{region}}.myhuaweicloud.com`.
  This can also be sourced from the *AWS_S3_ENDPOINT* environment variable.

* `skip_credentials_validation` - (Required) Skip credentials validation via the STS API. It's mandatory for HuaweiCloud.

* `skip_region_validation` - (Required) Skip validation of provided region name. It's mandatory for HuaweiCloud.

* `skip_metadata_api_check` - (Required) Skip usage of EC2 Metadata API. It's mandatory for HuaweiCloud.

* `workspace_key_prefix` - (Optional) Specifies the prefix applied to the state path inside the bucket. This parameter
  is only valid when using a non-default [workspace](https://www.terraform.io/docs/language/state/workspaces.html).
  When using a non-default workspace, the state path will be `/workspace_key_prefix/workspace_name/key_name`.

## For More Information

* [Terraform Remote State](https://www.terraform.io/docs/language/state/remote.html)
* [Terraform Backends](https://www.terraform.io/docs/language/settings/backends/index.html)
* [Amazon S3 Backend](https://www.terraform.io/docs/language/settings/backends/s3.html)
* [Workspaces](https://www.terraform.io/docs/language/state/workspaces.html)
* [The terraform_remote_state Data Source](https://www.terraform.io/docs/language/state/remote-state-data.html)
