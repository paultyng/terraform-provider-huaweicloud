---
subcategory: "FunctionGraph"
---

# huaweicloud_fgs_dependencies

Use this data source to filter dependent packages of FGS from HuaweiCloud.

## Example Usage

### Obtain all public dependent packages

```hcl
data "huaweicloud_fgs_dependencies" "test" {}
```

### Obtain specific public dependent package by name

```hcl
data "huaweicloud_fgs_dependencies" "test" {
  type = "public"
  name = "obssdk-3.0.2"
}
```

### Obtain all public Python2.7 dependent packages

```hcl
data "huaweicloud_fgs_dependencies" "test" {
  type    = "public"
  runtime = "Python2.7"
}
```

## Argument Reference

* `region` - (Optional, String) Specifies the region in which to obtain the dependent packages. If omitted, the
  provider-level region will be used.

* `type` - (Optional, String) Specifies the dependent package type to match. Valid values: **public** and **private**.

* `runtime` - (Optional, String) Specifies the dependent package runtime to match. Valid values: **Java8**,
  **Node.js6.10**, **Node.js8.10**, **Node.js10.16**, **Node.js12.13**, **Python2.7**, **Python3.6**, **Go1.8**,
  **Go1.x**, **C#(.NET Core 2.0)**, **C#(.NET Core 2.1)**, **C#(.NET Core 3.1)** and **PHP7.3**.

* `name` - (Optional, String) Specifies the dependent package runtime to match.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - A data source ID.

* `packages` - All dependent packages that match.

  + `id` - Dependent package ID.

  + `name` - Dependent package name.

  + `owner` - Dependent package owner.

  + `link` - URL of the dependent package in the OBS console.

  + `etag` - Unique ID of the dependent package.

  + `size` - Dependent package size.

  + `file_name` - File name of the Dependent package.

  + `runtime` - Dependent package runtime.
