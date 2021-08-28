---
page_title: "CCE Cluster Enhanced Authentication"
---

# CCE Cluster Enhanced Authentication

This page contains an example of how to use the cce cluster enhanced authentication with terraform.
When you use the cce cluster enhanced authentication, you should set the value of parameter `authentication_mode`
to authenticating_proxy, and there will be another 3 parameters(`authenticating_proxy_ca`, `authenticating_proxy_cert`
and `authenticating_proxy_private_key`) needed.

## Genetate CA Root Certificate, Client Certificate and Client Certificate Private Key with openssl

1. Generate CA private key.

    ```bash
    openssl genrsa -out ca.key 2048
    ```

2. Generate CSR.

    ```bash
    openssl req -new -key ca.key -out ca.csr
    ```

3. Generate CA root certificate.

    ```bash
    openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.crt
    ```

4. Generate client private key.

    ```bash
    openssl genrsa -out client.key 2048
    ```

5. Generate CSR.

    ```bash
    openssl req -new -key client.key -out client.csr
    ```

6. Generate Client certificate.

    ```bash
    openssl ca -in client.csr -out client.crt -cert ca.crt -keyfile ca.key -days 3650
    ```

7. Copy `ca.crt`, `client.crt` and `client.key` to your directory.

## Create CCE cluster with Enhance Authentication

```hcl
resource "huaweicloud_vpc" "myvpc" {
  name = "vpc"
  cidr = "192.168.0.0/16"
}

resource "huaweicloud_vpc_subnet" "mysubnet" {
  name          = "subnet"
  cidr          = "192.168.0.0/16"
  gateway_ip    = "192.168.0.1"
  //dns is required for cce node installing
  primary_dns   = "100.125.1.250"
  secondary_dns = "100.125.21.250"
  vpc_id        = huaweicloud_vpc.myvpc.id
}

resource "huaweicloud_cce_cluster" "cluster" {
  name                   = "cluster"
  flavor_id              = "cce.s1.small"
  vpc_id                 = huaweicloud_vpc.myvpc.id
  subnet_id              = huaweicloud_vpc_subnet.mysubnet.id
  container_network_type = "overlay_l2"

  authentication_mode              = "authenticating_proxy"
  authenticating_proxy_ca          = filebase64("your_directory/ca.crt")
  authenticating_proxy_cert        = filebase64("your_directory/client.crt")
  authenticating_proxy_private_key = filebase64("your_directory/client.key")
}
```
