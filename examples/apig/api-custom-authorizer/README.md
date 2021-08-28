# Register an API with Custom Authorizer and FunctionGraph

Configuration in this directory creates an API with custom authorizer and function back-end, The example includes a
function graph, an ECS for Nginx server, an API Gateway instance, some VPC and APIG configurations. It also creates
custom response, authorizer for the API and bind an EIP to ECS.

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

Wait a couple of minutes for the ECS to install nginx, and then you will see the nginx welcome page on your browser
through the EIP address.

For API, you need to configure your back-end database to bind the EIP address and port. After the connection is
successful, you can debug or publish this API on the console.

## Requirements

| Name | Version |
| ---- | ---- |
| terraform | > = 0.12.0 |
| huaweicloud | > = 1.27.1 |
