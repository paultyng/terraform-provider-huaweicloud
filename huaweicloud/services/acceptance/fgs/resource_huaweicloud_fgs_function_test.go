package fgs

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/chnsz/golangsdk/openstack/fgs/v2/function"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance/common"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/fgs"
)

func getResourceObj(cfg *config.Config, state *terraform.ResourceState) (interface{}, error) {
	client, err := cfg.NewServiceClient("fgs", acceptance.HW_REGION_NAME)
	if err != nil {
		return nil, fmt.Errorf("error creating FunctionGraph client: %s", err)
	}

	return fgs.GetFunctionMetadata(client, state.Primary.ID)
}

func TestAccFgsV2Function_basic(t *testing.T) {
	var (
		f              function.Function
		randName       = acceptance.RandomAccResourceName()
		obsOjectConfig = zipFileUploadResourcesConfig()
		resourceName   = "huaweicloud_fgs_function.test"
	)

	rc := acceptance.InitResourceCheck(
		resourceName,
		&f,
		getResourceObj,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccFgsV2Function_basic_step1(randName),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					// Default value is v2. Some regions support only v1, the default value is v1
					resource.TestMatchResourceAttr(resourceName, "functiongraph_version", regexp.MustCompile(`v1|v2`)),
					resource.TestCheckResourceAttr(resourceName, "description", "function test"),
					resource.TestCheckResourceAttr(resourceName, "tags.foo", "bar"),
					resource.TestCheckResourceAttr(resourceName, "tags.key", "value"),
					resource.TestCheckResourceAttrSet(resourceName, "urn"),
					resource.TestCheckResourceAttrSet(resourceName, "version"),
					resource.TestCheckResourceAttr(resourceName, "code_type", "inline"),
				),
			},
			{
				Config: testAccFgsV2Function_basic_step2(randName),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "description", "function test update"),
					resource.TestCheckResourceAttr(resourceName, "tags.foo", "baar"),
					resource.TestCheckResourceAttr(resourceName, "tags.newkey", "value"),
					resource.TestCheckResourceAttrSet(resourceName, "urn"),
					resource.TestCheckResourceAttrSet(resourceName, "version"),
					resource.TestCheckResourceAttrPair(resourceName, "vpc_id", "huaweicloud_vpc.test", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "network_id", "huaweicloud_vpc_subnet.test", "id"),
					resource.TestCheckResourceAttr(resourceName, "depend_list.#", "1"),
				),
			},
			{
				Config: testAccFgsV2Function_basic_step3(randName, obsOjectConfig),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "depend_list.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "code_type", "obs"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"func_code",
					"tags",
				},
			},
		},
	})
}

func TestAccFgsV2Function_withEpsId(t *testing.T) {
	var f function.Function
	randName := acceptance.RandomAccResourceName()
	resourceName := "huaweicloud_fgs_function.test"

	rc := acceptance.InitResourceCheck(
		resourceName,
		&f,
		getResourceObj,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckEpsID(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccFgsV2Function_withEpsId(randName),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "enterprise_project_id",
						acceptance.HW_ENTERPRISE_PROJECT_ID_TEST),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"func_code",
				},
			},
		},
	})
}

func TestAccFgsV2Function_text(t *testing.T) {
	var f function.Function
	randName := acceptance.RandomAccResourceName()
	resourceName := "huaweicloud_fgs_function.test"

	rc := acceptance.InitResourceCheck(
		resourceName,
		&f,
		getResourceObj,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccFgsV2Function_text(randName),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"func_code",
				},
			},
		},
	})
}

func TestAccFgsV2Function_createByImage(t *testing.T) {
	var f function.Function
	randName := acceptance.RandomAccResourceName()
	rName1 := "huaweicloud_fgs_function.create_with_vpc_access"
	rName2 := "huaweicloud_fgs_function.create_without_vpc_access"

	rc1 := acceptance.InitResourceCheck(
		rName1,
		&f,
		getResourceObj,
	)

	rc2 := acceptance.InitResourceCheck(
		rName2,
		&f,
		getResourceObj,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			acceptance.TestAccPreCheckComponentDeployment(t)
			acceptance.TestAccPreCheckImageUrlUpdated(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc1.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccFgsV2Function_createByImage_step_1(randName),
				Check: resource.ComposeTestCheckFunc(
					rc1.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName1, "name", randName+"_1"),
					resource.TestCheckResourceAttr(rName1, "agency", "functiongraph_swr_trust"),
					resource.TestCheckResourceAttr(rName1, "runtime", "Custom Image"),
					resource.TestCheckResourceAttr(rName1, "handler", "-"),
					resource.TestCheckResourceAttr(rName1, "custom_image.0.url", acceptance.HW_BUILD_IMAGE_URL),
					resource.TestCheckResourceAttrPair(rName1, "vpc_id", "huaweicloud_vpc.test", "id"),
					resource.TestCheckResourceAttrPair(rName1, "network_id", "huaweicloud_vpc_subnet.test", "id"),
					resource.TestCheckResourceAttr(rName1, "custom_image.0.command", "/bin/sh"),
					resource.TestCheckResourceAttr(rName1, "custom_image.0.args", "-args,value"),
					resource.TestCheckResourceAttr(rName1, "custom_image.0.working_dir", "/"),
					resource.TestCheckResourceAttr(rName1, "concurrency_num", "1"),
					rc2.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName2, "name", randName+"_2"),
					resource.TestCheckResourceAttr(rName2, "agency", "functiongraph_swr_trust"),
					resource.TestCheckResourceAttr(rName2, "runtime", "Custom Image"),
					resource.TestCheckResourceAttr(rName2, "handler", "-"),
					resource.TestCheckResourceAttr(rName2, "custom_image.0.url", acceptance.HW_BUILD_IMAGE_URL),
					resource.TestCheckResourceAttr(rName2, "vpc_id", ""),
					resource.TestCheckResourceAttr(rName2, "network_id", ""),
				),
			},
			{
				Config: testAccFgsV2Function_createByImage_step_2(randName),
				Check: resource.ComposeTestCheckFunc(
					rc1.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName1, "handler", "-"),
					resource.TestCheckResourceAttr(rName1, "vpc_id", ""),
					resource.TestCheckResourceAttr(rName1, "network_id", ""),
					resource.TestCheckResourceAttr(rName1, "custom_image.0.url", acceptance.HW_BUILD_IMAGE_URL_UPDATED),
					resource.TestCheckResourceAttr(rName1, "custom_image.0.command", ""),
					resource.TestCheckResourceAttr(rName1, "custom_image.0.args", ""),
					resource.TestCheckResourceAttr(rName1, "custom_image.0.working_dir", "/"),
					resource.TestCheckResourceAttr(rName1, "concurrency_num", "1000"),
					rc2.CheckResourceExists(),
					resource.TestCheckResourceAttr(rName2, "handler", "-"),
					resource.TestCheckResourceAttrPair(rName2, "vpc_id", "huaweicloud_vpc.test", "id"),
					resource.TestCheckResourceAttrPair(rName2, "network_id", "huaweicloud_vpc_subnet.test", "id"),
					resource.TestCheckResourceAttr(rName2, "custom_image.0.url", acceptance.HW_BUILD_IMAGE_URL_UPDATED),
				),
			},
			{
				ResourceName:      rName1,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      rName2,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccFgsV2Function_logConfig(t *testing.T) {
	var f function.Function
	randName := acceptance.RandomAccResourceName()
	resourceName := "huaweicloud_fgs_function.test"

	rc := acceptance.InitResourceCheck(
		resourceName,
		&f,
		getResourceObj,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccFgsV2Function_logConfig(randName),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "functiongraph_version", "v1"),
					resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "log_stream_id"),
				),
			},
			{
				Config: testAccFgsV2Function_logConfigUpdate(randName),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "functiongraph_version", "v1"),
					resource.TestCheckResourceAttrSet(resourceName, "log_group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "log_stream_id"),
				),
			},
		},
	})
}

func zipFileUploadResourcesConfig() string {
	randName := acceptance.RandomAccResourceNameWithDash()

	return fmt.Sprintf(`
variable "script_content" {
  type    = string
  default = <<EOT
def main():  
    print("Hello, World!")  

if __name__ == "__main__":  
    main()
EOT
}

resource "huaweicloud_obs_bucket" "test" {
  bucket = "%[1]s"
  acl    = "private"

  provisioner "local-exec" {
    command = "echo '${var.script_content}' >> test.py\nzip -r test.zip test.py"
  }
  provisioner "local-exec" {
    command = "rm test.zip test.py"
    when    = destroy
  }
}

resource "huaweicloud_obs_bucket_object" "test" {
  bucket = huaweicloud_obs_bucket.test.bucket
  key    = "test.zip"
  source = abspath("./test.zip")
}`, randName)
}

func testAccFgsV2Function_basic_step1(rName string) string {
	//nolint:revive
	return fmt.Sprintf(`
resource "huaweicloud_fgs_function" "test" {
  name        = "%s"
  app         = "default"
  description = "function test"
  handler     = "index.handler"
  memory_size = 128
  timeout     = 3
  runtime     = "Python2.7"
  code_type   = "inline"
  func_code   = "aW1wb3J0IGpzb24KZGVmIGhhbmRsZXIgKGV2ZW50LCBjb250ZXh0KToKICAgIG91dHB1dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganNvbi5kdW1wcyhldmVudCkKICAgIHJldHVybiBvdXRwdXQ="

  tags = {
    foo = "bar"
    key = "value"
  }
}
`, rName)
}

func testAccFgsV2Function_basic_step2(rName string) string {
	//nolint:revive
	return fmt.Sprintf(`
%[1]s

data "huaweicloud_fgs_dependencies" "test" {
  type    = "public"
  runtime = "Python2.7"
}

resource "huaweicloud_fgs_function" "test" {
  name        = "%[2]s"
  app         = "default"
  description = "function test update"
  handler     = "index.handler"
  memory_size = 128
  timeout     = 3
  runtime     = "Python2.7"
  code_type   = "inline"
  func_code   = "aW1wb3J0IGpzb24KZGVmIGhhbmRsZXIgKGV2ZW50LCBjb250ZXh0KToKICAgIG91dHB1dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganNvbi5kdW1wcyhldmVudCkKICAgIHJldHVybiBvdXRwdXQ="
  agency      = "function_vpc_trust"
  vpc_id      = huaweicloud_vpc.test.id
  network_id  = huaweicloud_vpc_subnet.test.id
  depend_list = try(slice(data.huaweicloud_fgs_dependencies.test.packages[*].id, 0, 1), [])

  tags = {
    foo    = "baar"
    newkey = "value"
  }
}
`, common.TestVpc(rName), rName)
}

func testAccFgsV2Function_basic_step3(rName, obsConfig string) string {
	//nolint:revive
	return fmt.Sprintf(`
%[1]s

%[2]s

data "huaweicloud_fgs_dependencies" "test" {
  type    = "public"
  runtime = "Python2.7"
}

resource "huaweicloud_fgs_function" "test" {
  name        = "%[3]s"
  app         = "default"
  description = "fuction test update"
  handler     = "index.handler"
  memory_size = 128
  timeout     = 3
  runtime     = "Python2.7"
  code_type   = "obs"
  code_url    = format("https://%%s/%%s", huaweicloud_obs_bucket.test.bucket_domain_name, huaweicloud_obs_bucket_object.test.key)
  agency      = "function_vpc_trust"
  vpc_id      = huaweicloud_vpc.test.id
  network_id  = huaweicloud_vpc_subnet.test.id
  depend_list = try(slice(data.huaweicloud_fgs_dependencies.test.packages[*].id, 0, 2), [])

  tags = {
    foo    = "baar"
    newkey = "value"
  }
}
`, common.TestVpc(rName), obsConfig, rName)
}

func testAccFgsV2Function_text(rName string) string {
	return fmt.Sprintf(`
resource "huaweicloud_fgs_function" "test" {
  name        = "%s"
  app         = "default"
  description = "fuction test"
  handler     = "index.handler"
  memory_size = 128
  timeout     = 3
  runtime     = "Python2.7"
  code_type   = "inline"

  func_code = <<EOF
# -*- coding:utf-8 -*-
import json
def handler (event, context):
    return {
        "statusCode": 200,
        "isBase64Encoded": False,
        "body": json.dumps(event),
        "headers": {
            "Content-Type": "application/json"
        }
    }
EOF
}
`, rName)
}

func testAccFgsV2Function_withEpsId(rName string) string {
	//nolint:revive
	return fmt.Sprintf(`
resource "huaweicloud_fgs_function" "test" {
  name                  = "%s"
  app                   = "default"
  description           = "fuction test"
  handler               = "index.handler"
  memory_size           = 128
  timeout               = 3
  runtime               = "Python2.7"
  code_type             = "inline"
  enterprise_project_id = "%s"
  func_code             = "aW1wb3J0IGpzb24KZGVmIGhhbmRsZXIgKGV2ZW50LCBjb250ZXh0KToKICAgIG91dHB1dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganNvbi5kdW1wcyhldmVudCkKICAgIHJldHVybiBvdXRwdXQ="
}
`, rName, acceptance.HW_ENTERPRISE_PROJECT_ID_TEST)
}

func testAccFgsV2Function_createByImage_step_1(rName string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_fgs_function" "create_with_vpc_access" {
  name                  = "%[2]s_1"
  app                   = "default"
  handler               = "-"
  functiongraph_version = "v2"
  memory_size           = 128
  runtime               = "Custom Image"
  timeout               = 3
  agency                = "functiongraph_swr_trust"

  custom_image {
    url         = "%[3]s"
    command     = "/bin/sh"
    args        = "-args,value"
    working_dir = "/"
  }

  vpc_id     = huaweicloud_vpc.test.id
  network_id = huaweicloud_vpc_subnet.test.id
}

resource "huaweicloud_fgs_function" "create_without_vpc_access" {
  name        = "%[2]s_2"
  app         = "default"
  handler     = "-"
  memory_size = 128
  runtime     = "Custom Image"
  timeout     = 3
  agency      = "functiongraph_swr_trust"

  custom_image {
    url = "%[3]s"
  }
}
`, common.TestVpc(rName), rName, acceptance.HW_BUILD_IMAGE_URL)
}

func testAccFgsV2Function_createByImage_step_2(rName string) string {
	return fmt.Sprintf(`
%[1]s

# Closs the VPC access
resource "huaweicloud_fgs_function" "create_with_vpc_access" {
  name                  = "%[2]s_1"
  app                   = "default"
  handler               = "-"
  functiongraph_version = "v2"
  memory_size           = 128
  runtime               = "Custom Image"
  timeout               = 3
  agency                = "functiongraph_swr_trust"

  custom_image {
    url = "%[3]s"
  }

  concurrency_num = 1000
}

# Open the VPC access
resource "huaweicloud_fgs_function" "create_without_vpc_access" {
  name        = "%[2]s_2"
  app         = "default"
  handler     = "-"
  memory_size = 128
  runtime     = "Custom Image"
  timeout     = 3
  agency      = "functiongraph_swr_trust"

  custom_image {
    url = "%[3]s"
  }

  vpc_id     = huaweicloud_vpc.test.id
  network_id = huaweicloud_vpc_subnet.test.id
}
`, common.TestVpc(rName), rName, acceptance.HW_BUILD_IMAGE_URL_UPDATED)
}

func TestAccFgsV2Function_strategy(t *testing.T) {
	var (
		f function.Function

		name         = acceptance.RandomAccResourceName()
		resourceName = "huaweicloud_fgs_function.test"
	)

	rc := acceptance.InitResourceCheck(
		resourceName,
		&f,
		getResourceObj,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccFunction_strategy_default(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "max_instance_num", "400"),
				),
			},
			{
				Config: testAccFunction_strategy_defined(name, 1000),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "max_instance_num", "1000"),
				),
			},
			{
				Config: testAccFunction_strategy_defined(name, 0),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "max_instance_num", "0"),
				),
			},
			{
				Config: testAccFunction_strategy_defined(name, -1),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "max_instance_num", "-1"),
				),
			},
			{
				Config: testAccFunction_strategy_default(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "max_instance_num", "-1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"func_code",
				},
			},
		},
	})
}

func testAccFunction_strategy_default(name string) string {
	return fmt.Sprintf(`
resource "huaweicloud_fgs_function" "test" {
  functiongraph_version = "v2"
  name                  = "%[1]s"
  app                   = "default"
  handler               = "index.handler"
  memory_size           = 128
  timeout               = 3
  runtime               = "Python2.7"
  code_type             = "inline"
  func_code             = "dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganN="
}
`, name)
}

func testAccFunction_strategy_defined(name string, maxInstanceNum int) string {
	return fmt.Sprintf(`
resource "huaweicloud_fgs_function" "test" {
  functiongraph_version = "v2"
  name                  = "%[1]s"
  app                   = "default"
  handler               = "index.handler"
  memory_size           = 128
  timeout               = 3
  runtime               = "Python2.7"
  code_type             = "inline"
  func_code             = "dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganN="
  max_instance_num      = %[2]d
}
`, name, maxInstanceNum)
}

func TestAccFgsV2Function_versions(t *testing.T) {
	var (
		f function.Function

		name         = acceptance.RandomAccResourceName()
		resourceName = "huaweicloud_fgs_function.test"
	)

	rc := acceptance.InitResourceCheck(
		resourceName,
		&f,
		getResourceObj,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccFunction_versions_step1(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "versions.#", "0"),
				),
			},
			{
				Config: testAccFunction_versions_step2(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "versions.#", "1"),
				),
			},
			{
				Config: testAccFunction_versions_step3(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "versions.#", "2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"func_code",
				},
			},
		},
	})
}

func testAccFunction_versions_step1(name string) string {
	return fmt.Sprintf(`
resource "huaweicloud_fgs_function" "test" {
  functiongraph_version = "v2"
  name                  = "%[1]s"
  app                   = "default"
  handler               = "index.handler"
  memory_size           = 128
  timeout               = 3
  runtime               = "Python2.7"
  code_type             = "inline"
  func_code             = "dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganN="
}
`, name)
}

func testAccFunction_versions_step2(name string) string {
	return fmt.Sprintf(`
resource "huaweicloud_fgs_function" "test" {
  functiongraph_version = "v2"
  name                  = "%[1]s"
  app                   = "default"
  handler               = "index.handler"
  memory_size           = 128
  timeout               = 3
  runtime               = "Python2.7"
  code_type             = "inline"
  func_code             = "dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganN="

  versions {
    name = "%[1]s"

    aliases {
      name        = "custom_alias"
      description = "This is a description of the custom alias"
    }
  }
}
`, name)
}

func testAccFunction_versions_step3(name string) string {
	return fmt.Sprintf(`
resource "huaweicloud_fgs_function" "test" {
  functiongraph_version = "v2"
  name                  = "%[1]s"
  app                   = "default"
  handler               = "index.handler"
  memory_size           = 128
  timeout               = 3
  runtime               = "Python2.7"
  code_type             = "inline"
  func_code             = "dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganN="

  versions {
    name = "latest"

    aliases {
      name = "demo"
    }
  }
  versions {
    name = "%[1]s"

    aliases {
      name = "custom_alias_update"
    }
  }
}
`, name)
}

func TestAccFgsV2Function_domain(t *testing.T) {
	var (
		f function.Function

		name         = acceptance.RandomAccResourceName()
		resourceName = "huaweicloud_fgs_function.test"
	)

	rc := acceptance.InitResourceCheck(
		resourceName,
		&f,
		getResourceObj,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccFunction_domain_step1(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
				),
			},
			{
				Config: testAccFunction_domain_step2(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"func_code",
				},
			},
		},
	})
}

func testAccFunction_domain_base(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_dns_zone" "test" {
  count = 3

  zone_type = "private"
  name      = format("functiondebug.example%%d.com.", count.index)

  router {
    router_id = huaweicloud_vpc.test.id
  }
}
`, common.TestVpc(name))
}

func testAccFunction_domain_step1(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_fgs_function" "test" {
  name        = "%[2]s"
  app         = "default"
  handler     = "index.handler"
  code_type   = "inline"
  memory_size = 128
  runtime     = "Python3.10"
  timeout     = 3
  func_code   = "dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganN="

  # VPC access and DNS configuration
  agency     = "function_all_trust" # Allow VPC and DNS permissions for FunctionGraph service
  vpc_id     = huaweicloud_vpc.test.id
  network_id = huaweicloud_vpc_subnet.test.id
  dns_list   = jsonencode(
    [for v in slice(huaweicloud_dns_zone.test[*], 0, 2) : tomap({id=v.id, domain_name=v.name})]
  )
}
`, testAccFunction_domain_base(name), name)
}

func testAccFunction_domain_step2(name string) string {
	return fmt.Sprintf(`
%[1]s

resource "huaweicloud_fgs_function" "test" {
  name        = "%[2]s"
  app         = "default"
  handler     = "index.handler"
  code_type   = "inline"
  memory_size = 128
  runtime     = "Python3.10"
  timeout     = 3
  func_code   = "dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganN="

  # VPC access and DNS configuration
  agency     = "function_all_trust" # Allow VPC and DNS permissions for FunctionGraph service
  vpc_id     = huaweicloud_vpc.test.id
  network_id = huaweicloud_vpc_subnet.test.id
  dns_list   = jsonencode(
    [for v in slice(huaweicloud_dns_zone.test[*], 1, 3) : tomap({id=v.id, domain_name=v.name})]
  )
}
`, testAccFunction_domain_base(name), name)
}

func testAccFgsV2Function_logConfig(rName string) string {
	return fmt.Sprintf(`
resource "huaweicloud_lts_group" "test" {
  group_name  = "%[1]s"
  ttl_in_days = 30
}

resource "huaweicloud_lts_stream" "test" {
  group_id    = huaweicloud_lts_group.test.id
  stream_name = "%[1]s"
}

resource "huaweicloud_fgs_function" "test" {
  name                  = "%[1]s"
  app                   = "default"
  description           = "fuction test"
  handler               = "index.handler"
  functiongraph_version = "v1"
  memory_size           = 128
  timeout               = 3
  runtime               = "Python2.7"
  code_type             = "inline"
  func_code             = "dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganN="

  log_group_id    = huaweicloud_lts_group.test.id
  log_stream_id   = huaweicloud_lts_stream.test.id
  log_group_name  = huaweicloud_lts_group.test.group_name
  log_stream_name = huaweicloud_lts_stream.test.stream_name
}
`, rName)
}

func testAccFgsV2Function_logConfigUpdate(rName string) string {
	return fmt.Sprintf(`
resource "huaweicloud_lts_group" "test" {
  group_name  = "%[1]s"
  ttl_in_days = 30
}

resource "huaweicloud_lts_stream" "test" {
  group_id    = huaweicloud_lts_group.test.id
  stream_name = "%[1]s"
}

resource "huaweicloud_lts_group" "test1" {
  group_name  = "%[1]s-new"
  ttl_in_days = 30
}

resource "huaweicloud_lts_stream" "test1" {
  group_id    = huaweicloud_lts_group.test1.id
  stream_name = "%[1]s-new"
}

resource "huaweicloud_fgs_function" "test" {
  name                  = "%[1]s"
  app                   = "default"
  description           = "fuction test"
  handler               = "index.handler"
  functiongraph_version = "v1"
  memory_size           = 128
  timeout               = 3
  runtime               = "Python2.7"
  code_type             = "inline"
  func_code             = "dCA9ICdIZWxsbyBtZXNzYWdlOiAnICsganN="

  log_group_id    = huaweicloud_lts_group.test1.id
  log_stream_id   = huaweicloud_lts_stream.test1.id
  log_group_name  = huaweicloud_lts_group.test1.group_name
  log_stream_name = huaweicloud_lts_stream.test1.stream_name
}
`, rName)
}

func TestAccFgsV2Function_reservedInstance_version(t *testing.T) {
	var (
		f            function.Function
		name         = acceptance.RandomAccResourceName()
		resourceName = "huaweicloud_fgs_function.test"
	)

	rc := acceptance.InitResourceCheck(
		resourceName,
		&f,
		getResourceObj,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccFgsV2Function_reservedInstance_step1(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.qualifier_name", "latest"),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.qualifier_type", "version"),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.count", "1"),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.idle_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.tactics_config.0.cron_configs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.tactics_config.0.cron_configs.0.count", "2"),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.tactics_config.0.cron_configs.0.cron", "0 */10 * * * ?"),
					resource.TestCheckResourceAttrSet(resourceName, "reserved_instances.0.tactics_config.0.cron_configs.0.start_time"),
					resource.TestCheckResourceAttrSet(resourceName, "reserved_instances.0.tactics_config.0.cron_configs.0.expired_time"),
					resource.TestCheckResourceAttrSet(resourceName, "reserved_instances.0.tactics_config.0.cron_configs.0.name"),
				),
			},
			{
				Config: testAccFgsV2Function_reservedInstance_step2(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.count", "2"),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.idle_mode", "false"),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.tactics_config.#", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"func_code",
					"tags",
				},
			},
		},
	})
}

func TestAccFgsV2Function_reservedInstance_alias(t *testing.T) {
	var (
		f               function.Function
		name            = acceptance.RandomAccResourceName()
		updateAliasName = acceptance.RandomAccResourceName()
		resourceName    = "huaweicloud_fgs_function.test"
	)

	rc := acceptance.InitResourceCheck(
		resourceName,
		&f,
		getResourceObj,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccFgsV2Function_reservedInstance_alias(name, name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.qualifier_name", name),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.qualifier_type", "alias"),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.count", "1"),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.idle_mode", "false"),
				),
			},
			{
				Config: testAccFgsV2Function_reservedInstance_alias(name, updateAliasName),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.qualifier_name", updateAliasName),
					resource.TestCheckResourceAttr(resourceName, "reserved_instances.0.qualifier_type", "alias"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"func_code",
					"tags",
				},
			},
		},
	})
}

func testAccFgsV2Function_reservedInstance_step1(rName string) string {
	return fmt.Sprintf(`
resource "huaweicloud_fgs_function" "test" {
  name        = "%[1]s"
  app         = "default"
  handler     = "index.handler"
  memory_size = 128
  timeout     = 3
  runtime     = "Node.js16.17"
  code_type   = "inline"

  reserved_instances {
    qualifier_type = "version"
    qualifier_name = "latest"
    count          = 1
    idle_mode      = true

    tactics_config {
      cron_configs {
        name         = "scheme-waekcy"
        cron         = "0 */10 * * * ?"
        start_time   = "1708342889"
        expired_time = "1739878889"
        count        = 2
      }
    }
  }
}
`, rName)
}

func testAccFgsV2Function_reservedInstance_step2(rName string) string {
	return fmt.Sprintf(`
resource "huaweicloud_fgs_function" "test" {
  name        = "%s"
  app         = "default"
  handler     = "index.handler"
  memory_size = 128
  timeout     = 3
  runtime     = "Node.js16.17"
  code_type   = "inline"
	  
  reserved_instances {
    qualifier_type = "version"
    qualifier_name = "latest"
    count          = 2
    idle_mode      = false
  }
}
`, rName)
}

func testAccFgsV2Function_reservedInstance_alias(rName string, aliasName string) string {
	return fmt.Sprintf(`
resource "huaweicloud_fgs_function" "test" {
  name        = "%[1]s"
  app         = "default"
  handler     = "index.handler"
  memory_size = 128
  timeout     = 3
  runtime     = "Node.js16.17"
  code_type   = "inline"

  versions {
    name = "latest"
	
    aliases {
      name = "%[2]s"
    }
  }

  reserved_instances {
    qualifier_type = "alias"
    qualifier_name = "%[2]s"
    count          = 1
    idle_mode      = false
  }
}
`, rName, aliasName)
}

func TestAccFgsV2Function_gpuMemory(t *testing.T) {
	var (
		f function.Function

		name         = acceptance.RandomAccResourceName()
		resourceName = "huaweicloud_fgs_function.test"
	)

	rc := acceptance.InitResourceCheck(
		resourceName,
		&f,
		getResourceObj,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheck(t)
			// Only the cn-east-3 region supports this function, and you need to submit a service ticket to enable the function.
			acceptance.TestAccPreCheckFgsGpuType(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccFunction_gpuMemory(name, 1024),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "gpu_memory", "1024"),
					resource.TestCheckResourceAttr(resourceName, "gpu_type", acceptance.HW_FGS_GPU_TYPE),
				),
			},
			{
				Config: testAccFunction_gpuMemory(name, 2048),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "gpu_memory", "2048"),
					resource.TestCheckResourceAttr(resourceName, "gpu_type", acceptance.HW_FGS_GPU_TYPE),
				),
			},
			{
				Config: testAccFunction_default(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "gpu_memory", "0"),
					resource.TestCheckResourceAttr(resourceName, "gpu_type", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"func_code",
				},
			},
		},
	})
}

func testAccFunction_gpuMemory(name string, memory int) string {
	return fmt.Sprintf(`
resource "huaweicloud_fgs_function" "test" {
  name        = "%[1]s"
  app         = "default"
  handler     = "bootstrap"
  memory_size = 128
  timeout     = 3
  runtime     = "Custom"
  code_type   = "inline"
  gpu_memory  = %[2]d
  gpu_type    = "%[3]s"
}
`, name, memory, acceptance.HW_FGS_GPU_TYPE)
}

func testAccFunction_default(name string) string {
	return fmt.Sprintf(`
resource "huaweicloud_fgs_function" "test" {
  name        = "%[1]s"
  app         = "default"
  handler     = "bootstrap"
  memory_size = 128
  timeout     = 3
  runtime     = "Custom"
  code_type   = "inline"
}
`, name)
}
