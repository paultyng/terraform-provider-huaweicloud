---
page_title: "HuaweiCloud Provider Mirror Guide"
---

# Setting a Network Mirror

Sometimes, running the `terraform init` command to download and install providers can be too slow,
or it may even fail as follows:

```text
Initializing provider plugins...
- Finding huaweicloud/huaweicloud versions matching "1.65.0"...
╷
│ Error: Failed to query available provider packages
│ 
│ Could not retrieve the list of available versions for provider huaweicloud/huaweicloud: could not connect to registry.terraform.io: failed to request discovery document: Get
│ "https://registry.terraform.io/.well-known/terraform.json": dial tcp 3.165.82.91:443: connect: no route to host
╵
```

or

```text
Initializing provider plugins...
- Finding huaweicloud/huaweicloud versions matching "1.66.0"...
╷
│ Error: Failed to query available provider packages
│
│ Could not retrieve the list of available versions for provider huaweicloud/huaweicloud: could not query provider registry for registry.terraform.io/huaweicloud/huaweicloud: the
│ request failed after 3 attempts, please try again later: Get "https://registry.terraform.io/v1/providers/huaweicloud/huaweicloud/versions": context deadline exceeded
│ (Client.Timeout exceeded while awaiting headers)
╵
```

Since Terraform CLI v0.13.2, it provides the
[network_mirror](https://developer.hashicorp.com/terraform/cli/config/config-file#network_mirror) feature.
To fix the issues of downloading HuaweiCloud providers, HuaweiCloud provides a mirror service.
You can set the following configuration in the [CLI Configuration File](https://developer.hashicorp.com/terraform/cli/config/config-file):

```terraform
provider_installation {
  network_mirror {
    url = "https://tf-mirror.obs.cn-north-4.myhuaweicloud.com/"
    # Set HuaweiCloud providers to download from the mirror service.
    include = ["registry.terraform.io/huaweicloud/huaweicloud"]
  }
  direct {
    # For other providers, download from the Terraform Registry.
    exclude = ["registry.terraform.io/huaweicloud/huaweicloud"]
  }
}
```

## Example on Linux

- Create CLI configuration file `.terraformrc` and placed directly in the home directory of the relevant user.

```shell
tee ~/.terraformrc <<-'EOF'
provider_installation {
  network_mirror {
    url = "https://tf-mirror.obs.cn-north-4.myhuaweicloud.com/"
    # Set HuaweiCloud providers to download from the mirror service.
    include = ["registry.terraform.io/huaweicloud/huaweicloud"]
  }
  direct {
    # For other providers, download from the Terraform Registry.
    exclude = ["registry.terraform.io/huaweicloud/huaweicloud"]
  }
}
EOF
```

- When you run `terraform init` you will find the following output in the log file:

```text
2024-07-29T11:02:36.511+0800 [INFO]  Go runtime version: go1.22.1
2024-07-29T11:02:36.511+0800 [INFO]  CLI args: []string{"terraform", "init"}
2024-07-29T11:02:36.511+0800 [DEBUG] Attempting to open CLI config file: /home/huawei/.terraformrc
2024-07-29T11:02:36.511+0800 [INFO]  Loading CLI configuration from /home/huawei/.terraformrc
2024-07-29T11:02:36.511+0800 [DEBUG] checking for credentials in "/home/huawei/.terraform.d/plugins"
2024-07-29T11:02:36.511+0800 [DEBUG] Explicit provider installation configuration is set
2024-07-29T11:02:36.512+0800 [INFO]  CLI command args: []string{"init"}
2024-07-29T11:02:36.514+0800 [DEBUG] New state was assigned lineage "f234445e-7ddf-cf48-ade2-bb3ba4a9b5aa"
2024-07-29T11:02:36.514+0800 [DEBUG] checking for provisioner in "."
2024-07-29T11:02:36.523+0800 [DEBUG] checking for provisioner in "/usr/bin"
2024-07-29T11:02:36.523+0800 [DEBUG] checking for provisioner in "/home/huawei/.terraform.d/plugins"
2024-07-29T11:02:36.525+0800 [DEBUG] Querying available versions of provider registry.terraform.io/huaweicloud/huaweicloud at network mirror https://tf-mirror.obs.cn-north-4.myhuaweicloud.com/
2024-07-29T11:02:36.525+0800 [DEBUG] GET https://tf-mirror.obs.cn-north-4.myhuaweicloud.com/registry.terraform.io/huaweicloud/huaweicloud/index.json
2024-07-29T11:02:36.919+0800 [DEBUG] Finding package URL for registry.terraform.io/huaweicloud/huaweicloud v1.66.3 on linux_amd64 via network mirror https://tf-mirror.obs.cn-north-4.myhuaweicloud.com/
2024-07-29T11:02:36.919+0800 [DEBUG] GET https://tf-mirror.obs.cn-north-4.myhuaweicloud.com/registry.terraform.io/huaweicloud/huaweicloud/1.66.3.json
```
