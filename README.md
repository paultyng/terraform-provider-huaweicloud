Terraform HuaweiCloud Provider
==============================

<!-- markdownlint-disable-next-line MD034 -->
* Website: https://www.terraform.io
* [![Documentation](https://img.shields.io/badge/documentation-blue)](https://registry.terraform.io/providers/huaweicloud/huaweicloud/latest/docs)
  
* Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<a href="https://www.huaweicloud.com/">
  <img src="https://console-static.huaweicloud.com/static/authui/20210202115135/public/custom/images/logo-en.svg"
    alt="HUAWEI CLOUD" width="450px" height="102px">
</a>

Requirements
------------

* [Terraform](https://www.terraform.io/downloads.html) 0.12.x
* [Go](https://golang.org/doc/install) 1.18 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/huaweicloud/terraform-provider-huaweicloud`

```sh
$ mkdir -p $GOPATH/src/github.com/huaweicloud; cd $GOPATH/src/github.com/huaweicloud
$ git clone https://github.com/huaweicloud/terraform-provider-huaweicloud
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/huaweicloud/terraform-provider-huaweicloud
$ make build
```

Using the provider
------------------

Please see the documentation at [provider usage](docs/index.md).

Or you can browse the documentation within this repo [here](https://github.com/huaweicloud/terraform-provider-huaweicloud/tree/master/docs).

Developing the Provider
-----------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed
on your machine (version 1.14+ is *required*).
You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH),
as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`.
This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-huaweicloud
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

License
-------

Terraform-Provider-Huaweicloud is under the Mozilla Public License 2.0. See the [LICENSE](LICENSE) file for details.
