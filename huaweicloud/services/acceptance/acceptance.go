package acceptance

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils/fmtp"
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

func RandomAccResourceName() string {
	return fmt.Sprintf("tf_acc_test_%s", acctest.RandString(5))
}

const (
	// UNSET check if the resource is not set, corresponding to: TestCheckNoResourceAttrSet.
	UNSET = "#NOSET"
	// CHECKSET check if the resource is set, corresponding to: TestCheckResourceAttrSet.
	CHECKSET = "#CHECKSET"
	// REGEXMATCH use regular expression matching, corresponding to: TestMatchResourceAttr.
	//
	// e.g.: REGEXMATCH + "[a-zA-Z]{4,64}"`
	REGEXMATCH = "#REGEXMATCH:"
)

// ServiceFunc the HuaweiCloud resource query functions.
type ServiceFunc func(*config.Config, *terraform.ResourceState) (interface{}, error)

// resourceCheck resource check object, only used in the package.
type resourceCheck struct {
	resourceName    string
	checkMap        map[string]string
	resourceObject  interface{}
	getResourceFunc ServiceFunc
	resourceType    string
}

const (
	resourceTypeCode   = "resource"
	dataSourceTypeCode = "dataSource"
)

/*
InitDataSourceCheck build a 'resourceCheck' object. Only used to check datasource attributes.
  Parameters:
    resourceName:    The resource name is used to check in the terraform.State.e.g. : huaweicloud_waf_domain.domain_1.
  Return:
    *resourceCheck: resourceCheck object
*/
func InitDataSourceCheck(sourceName string) *resourceCheck {
	return &resourceCheck{
		resourceName: sourceName,
		resourceType: dataSourceTypeCode,
	}
}

/*
InitResourceCheck build a 'resourceCheck' object. The common test methods are provided in 'resourceCheck'.
  Parameters:
    resourceName:    The resource name is used to check in the terraform.State.e.g. : huaweicloud_waf_domain.domain_1.
    resourceObject:  Resource object, used to check whether the resource exists in HuaweiCloud.
    getResourceFunc: The function used to get the resource object.
  Return:
    *resourceCheck: resourceCheck object
*/
func InitResourceCheck(resourceName string, resourceObject interface{}, getResourceFunc ServiceFunc) *resourceCheck {
	return &resourceCheck{
		resourceName:    resourceName,
		resourceObject:  resourceObject,
		getResourceFunc: getResourceFunc,
		resourceType:    resourceTypeCode,
	}
}

// CheckResourceMapAttr check the value in the checkMap based on the resourceName.
func (rc *resourceCheck) CheckResourceMapAttr(checkMap map[string]string) resource.TestCheckFunc {
	attrCheck := func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rc.resourceName]
		if !ok {
			return fmtp.Errorf("can't find resource by id: %s", rc.resourceName)
		}
		if rs.Primary.ID == "" {
			return fmtp.Errorf("resource ID is not set")
		}

		var errors []string
		for key, value := range checkMap {
			var err error
			if strings.HasPrefix(value, REGEXMATCH) {
				var regex *regexp.Regexp
				regex, err = regexp.Compile(value[len(REGEXMATCH):])
				if err == nil {
					err = resource.TestMatchResourceAttr(rc.resourceName, key, regex)(s)
				} else {
					err = nil
				}
			} else if value == UNSET {
				err = resource.TestCheckNoResourceAttr(rc.resourceName, key)(s)
			} else if value == CHECKSET {
				err = resource.TestCheckResourceAttrSet(rc.resourceName, key)(s)
			} else {
				err = resource.TestCheckResourceAttr(rc.resourceName, key, value)(s)
			}
			if err != nil {
				errors = append(errors, err.Error())
			}
		}

		if len(errors) == 0 {
			return nil
		}
		return fmtp.Errorf("\n" + strings.Join(errors, "\n"))
	}
	// If provider is a data source, only check the attributes.
	if strings.Compare(rc.resourceType, dataSourceTypeCode) == 0 {
		return attrCheck
	}
	// If provider is a resource, check the resource whether exists first.
	return func(s *terraform.State) error {
		err := rc.CheckResourceExists()(s)
		if err == nil {
			return attrCheck(s)
		}
		return err
	}
}

// CheckResourceDestroy check whether resources destroied in HuaweiCloud.
func (rc *resourceCheck) CheckResourceDestroy() resource.TestCheckFunc {
	if strings.Compare(rc.resourceType, dataSourceTypeCode) == 0 {
		fmtp.Errorf("Error, you built a resourceCheck with 'InitDataSourceCheck', " +
			"it cannot run CheckResourceDestroy().")
		return nil
	}
	return func(s *terraform.State) error {
		strs := strings.Split(rc.resourceName, ".")
		var resourceType string
		for _, str := range strs {
			if strings.Contains(str, "huaweicloud_") {
				resourceType = strings.Trim(str, " ")
				break
			}
		}

		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceType {
				continue
			}

			conf := TestAccProvider.Meta().(*config.Config)
			if _, err := rc.getResourceFunc(conf, rs); err == nil {
				return fmtp.Errorf("failed to destroy resource. The resource of %s : %s still exists。", resourceType, rs.Primary.ID)
			}
		}
		return nil
	}
}

// CheckResourceExists check whether resources exist in HuaweiCloud.
func (rc *resourceCheck) CheckResourceExists() resource.TestCheckFunc {
	if strings.Compare(rc.resourceType, dataSourceTypeCode) == 0 {
		fmtp.Errorf("Error, you built a resourceCheck with 'InitDataSourceCheck', " +
			"it cannot run CheckResourceExists().")
		return nil
	}
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rc.resourceName]
		if !ok {
			return fmtp.Errorf("not found: %s", rc.resourceName)

		}
		if rs.Primary.ID == "" {
			return fmtp.Errorf("no id set for the resource: %s", rc.resourceName)
		}

		if rc.getResourceFunc != nil {
			conf := TestAccProvider.Meta().(*config.Config)
			r, err := rc.getResourceFunc(conf, rs)
			if err != nil {
				return fmtp.Errorf("checking resource %s %s exists error: %s ", rc.resourceName, rs.Primary.ID, err)
			}
			rc.resourceObject = r
		} else {
			return fmtp.Errorf("The 'getResourceFunc' is nil, please set it.")
		}

		return nil
	}
}
