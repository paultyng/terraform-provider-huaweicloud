# Create a WAF dedicated mode domain.

This example creates a dedicated mode domain name based on the example
[examples/waf/waf-dedicated-instance](https://github.com/huaweicloud/terraform-provider-huaweicloud/tree/master/examples/waf/waf-dedicated-instance)
. You can replace the VPC, Security Group, WAF Policy, and WAF dedicated instance with the resources already created in
HuaweiCloud.

To run, configure your HuaweiCloud provider as described in the
[document](https://registry.terraform.io/providers/huaweicloud/huaweicloud/latest/docs).

This example assumes that you have created a random password. If you want to use key-pair and do not have one, please
visit the
[document](https://registry.terraform.io/providers/huaweicloud/huaweicloud/latest/docs/resources/compute_keypair)
to create a key-pair.

## Usage

```shell
terraform init
terraform plan
terraform apply
terraform destroy
```

The creation of the WAF dedicated instance takes about 5 minutes. After the creation is successful, the WAF policy and
dedicated mode domain start to be created.

## Requirements

| Name | Version |
| ---- | ---- |
| terraform | > = 0.12.0 |
| huaweicloud | > = 1.28.0 |
