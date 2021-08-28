# Create a professional edition WAF dedicated instance.

This example first creates a WAF dedicated instance, and then creates a policy. The WAF dedicated instance requires a
VPC, Subnet, and security group. In this example, we all create them. You can replace them with resources already
created in Huawei Cloud.

To run, configure your Huaweicloud provider as described in the
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

The creation of the WAF dedicated instance takes about 5 minutes. After the creation is complete, the WAF policy starts
to be created.

## Requirements

| Name | Version |
| ---- | ---- |
| terraform | > = 0.12.0 |
| huaweicloud | > = 1.28.0 |
