package acceptance

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud"
)

var (
	HW_REGION_NAME        = os.Getenv("HW_REGION_NAME")
	HW_CUSTOM_REGION_NAME = os.Getenv("HW_CUSTOM_REGION_NAME")
	HW_AVAILABILITY_ZONE  = os.Getenv("HW_AVAILABILITY_ZONE")
	HW_ACCESS_KEY         = os.Getenv("HW_ACCESS_KEY")
	HW_SECRET_KEY         = os.Getenv("HW_SECRET_KEY")
	HW_PROJECT_ID         = os.Getenv("HW_PROJECT_ID")
	HW_DOMAIN_ID          = os.Getenv("HW_DOMAIN_ID")
	HW_DOMAIN_NAME        = os.Getenv("HW_DOMAIN_NAME")

	HW_FLAVOR_ID             = os.Getenv("HW_FLAVOR_ID")
	HW_FLAVOR_NAME           = os.Getenv("HW_FLAVOR_NAME")
	HW_IMAGE_ID              = os.Getenv("HW_IMAGE_ID")
	HW_IMAGE_NAME            = os.Getenv("HW_IMAGE_NAME")
	HW_VPC_ID                = os.Getenv("HW_VPC_ID")
	HW_NETWORK_ID            = os.Getenv("HW_NETWORK_ID")
	HW_SUBNET_ID             = os.Getenv("HW_SUBNET_ID")
	HW_ENTERPRISE_PROJECT_ID = os.Getenv("HW_ENTERPRISE_PROJECT_ID")
	HW_MAPREDUCE_CUSTOM      = os.Getenv("HW_MAPREDUCE_CUSTOM")

	HW_DEPRECATED_ENVIRONMENT = os.Getenv("HW_DEPRECATED_ENVIRONMENT")
)

var TestAccProviders map[string]*schema.Provider
var TestAccProvider *schema.Provider

func init() {
	TestAccProvider = huaweicloud.Provider()
	TestAccProviders = map[string]*schema.Provider{
		"huaweicloud": TestAccProvider,
	}
}

func preCheckRequiredEnvVars(t *testing.T) {
	if HW_REGION_NAME == "" {
		t.Fatal("HW_REGION_NAME must be set for acceptance tests")
	}
}

//lintignore:AT003
func TestAccPreCheck(t *testing.T) {
	// Do not run the test if this is a deprecated testing environment.
	if HW_DEPRECATED_ENVIRONMENT != "" {
		t.Skip("This environment only runs deprecated tests")
	}

	preCheckRequiredEnvVars(t)
}

//lintignore:AT003
func TestAccPrecheckCustomRegion(t *testing.T) {
	if HW_CUSTOM_REGION_NAME == "" {
		t.Skip("HW_CUSTOM_REGION_NAME must be set for acceptance tests")
	}
}

//lintignore:AT003
func TestAccPreCheckDeprecated(t *testing.T) {
	if HW_DEPRECATED_ENVIRONMENT == "" {
		t.Skip("This environment does not support deprecated tests")
	}

	preCheckRequiredEnvVars(t)
}

//lintignore:AT003
func TestAccPreCheckEpsID(t *testing.T) {
	if HW_ENTERPRISE_PROJECT_ID == "" {
		t.Skip("This environment does not support Enterprise Project ID tests")
	}
}

//lintignore:AT003
func TestAccPreCheckMrsCustom(t *testing.T) {
	if HW_MAPREDUCE_CUSTOM == "" {
		t.Skip("HW_MAPREDUCE_CUSTOM must be set for acceptance tests:custom type cluster of map reduce")
	}
}

func RandomAccResourceName() string {
	return fmt.Sprintf("tf_acc_test_%s", acctest.RandString(5))
}
