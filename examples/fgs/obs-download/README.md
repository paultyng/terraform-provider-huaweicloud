# Cretae an Timing Trigger to download files periodically

Configuration in this directory creates an Timing Trigger and download files from the OBS bucket to the
FGS cache periodically. The example includes a function of FGS, an OBS bucket and an IAM agency.
The change of the download address requires the user to modify the corresponding python code.
In this use case, the local file will be uploaded to the OBS bucket for function download.
The object name, object address, etc. need to be configured by the user.

To run, configure your Huaweicloud provider as described in the
[document](https://registry.terraform.io/providers/huaweicloud/huaweicloud/latest/docs).

If you want to use cron expression, please visit the
[document](https://support.huaweicloud.com/en-us/usermanual-functiongraph/functiongraph_01_0908.html)
and according to the example of the provider document.

## Usage

```
terraform init
terraform plan
terraform apply
terraform destroy
```

## Requirements

| Name | Version |
| ---- | ---- |
| terraform | >= 0.12.0 |
| huaweicloud | >= 1.27.2 |
