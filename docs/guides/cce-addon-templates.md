---
page_title: "CCE Addon Templates"
---

# CCE Addon Templates

Addon support configuration input depending on addon type and version. This page contains description of addon
arguments. You can get up to date reference of addon arguments for your cluster using data source
[`huaweicloud_cce_addon_template`](https://registry.terraform.io/providers/huaweicloud/huaweicloud/latest/docs/data-sources/cce_addon_template)
.

Following addon templates exist in the addon template list:

- [`autoscaler`](#autoscaler)
- [`coredns`](#coredns)
- [`everest`](#everest)
- [`metrics-server`](#metrics-server)
- [`gpu-beta`](#gpu-beta)

All addons accept `basic` and some can accept `custom` input values.

## Addon Inputs

### `autoscaler`

A component that automatically adjusts the size of a Kubernetes Cluster so that all pods have a place to run and there
are no unneeded nodes.
`template_version`: `1.19.1`

##### `basic`

```json
{
  "cceEndpoint": "https://cce.cn-north-4.myhuaweicloud.com",
  "ecsEndpoint": "https://ecs.cn-north-4.myhuaweicloud.com",
  "image_version": "1.19.6",
  "platform": "linux-amd64",
  "region": "cn-north-4",
  "swr_addr": "swr.cn-north-4.myhuaweicloud.com",
  "swr_user": "hwofficial"
}
```

##### `custom`

```json
{
  "cluster_id": "",
  "coresTotal": 32000,
  "expander": "priority",
  "logLevel": 4,
  "maxEmptyBulkDeleteFlag": 10,
  "maxNodeProvisionTime": 15,
  "maxNodesTotal": 1000,
  "memoryTotal": 128000,
  "scaleDownDelayAfterAdd": 10,
  "scaleDownDelayAfterDelete": 10,
  "scaleDownDelayAfterFailure": 3,
  "scaleDownEnabled": false,
  "scaleDownUnneededTime": 10,
  "scaleDownUtilizationThreshold": 0.5,
  "scaleUpCpuUtilizationThreshold": 1,
  "scaleUpMemUtilizationThreshold": 1,
  "scaleUpUnscheduledPodEnabled": true,
  "scaleUpUtilizationEnabled": true,
  "tenant_id": "",
  "unremovableNodeRecheckTimeout": 5
}
```

### `coredns`

CoreDNS is a DNS server that chains plugins and provides Kubernetes DNS Services.
`template_version`: `1.17.7`

##### `basic`

```json
{
  "cluster_ip": "10.247.3.10",
  "image_version": "1.17.7",
  "platform": "linux-amd64",
  "swr_addr": "swr.cn-north-4.myhuaweicloud.com",
  "swr_user": "hwofficial"
}
```

##### `custom`

```json
{
  "stub_domains": "",
  "upstream_nameservers": ""
}
```

### `everest`

Everest is a cloud native container storage system based on CSI, used to support cloud storages services for Kubernetes.
`template_version`: `1.2.9`

##### `basic`

```json
{
  "bms_url": "bms.cn-north-4.myhuaweicloud.com",
  "controller_image_version": "1.2.9",
  "driver_image_version": "1.2.9",
  "ecsEndpoint": "https://ecs.cn-north-4.myhuaweicloud.com",
  "evs_url": "evs.cn-north-4.myhuaweicloud.com",
  "iam_url": "iam.cn-north-4.myhuaweicloud.com",
  "ims_url": "ims.cn-north-4.myhuaweicloud.com",
  "obs_url": "obs.cn-north-4.myhuaweicloud.com",
  "platform": "linux-amd64",
  "sfs_turbo_url": "sfs-turbo.cn-north-4.myhuaweicloud.com",
  "sfs_url": "sfs.cn-north-4.myhuaweicloud.com",
  "supportHcs": false,
  "swr_addr": "swr.cn-north-4.myhuaweicloud.com",
  "swr_user": "hwofficial"
}
```

##### `custom`

```json
{
  "cluster_id": "",
  "default_vpc_id": "",
  "disable_auto_mount_secret": false,
  "project_id": ""
}
```

### `metrics-server`

Metrics Server is a cluster-level resource usage data aggregator.
`template_version`: `1.1.2`

##### `basic`

```json
{
  "image_version": "v0.4.4",
  "swr_addr": "swr.cn-north-4.myhuaweicloud.com",
  "swr_user": "hwofficial"
}
```

##### `custom`

_Not supported_

### `gpu-beta`

A device plugin for nvidia.com/gpu resource on nvidia driver.
`template_version`: `1.2.2`

##### `basic`

```json
{
  "device_version": "1.2.2",
  "driver_version": "1.2.2",
  "obs_url": "obs.cn-north-4.myhuaweicloud.com",
  "region": "cn-north-4",
  "swr_addr": "swr.cn-north-4.myhuaweicloud.com",
  "swr_user": "hwofficial"
}
```

##### `custom`

```json
{
  "is_driver_from_nvidia": true,
  "nvidia_driver_download_url": ""
}
```
