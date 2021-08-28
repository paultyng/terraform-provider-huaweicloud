package waf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceWafCertificateV1_basic(t *testing.T) {
	name := acceptance.RandomAccResourceName()
	resourceName := "data.huaweicloud_waf_certificate.cert_1"

	rc := acceptance.InitDataSourceCheck(
		resourceName,
	)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acceptance.TestAccPreCheck(t) },
		Providers: acceptance.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccWafCertificateListV1_conf(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceMapAttr(map[string]string{
						"name":          name,
						"id":            acceptance.CHECKSET,
						"expire_status": acceptance.CHECKSET,
						"expiration":    acceptance.CHECKSET,
					}),
				),
			},
		},
	})
}

func testAccWafCertificateListV1_conf(name string) string {
	return fmt.Sprintf(`
%s

data "huaweicloud_waf_certificate" "cert_1" {
  name       = huaweicloud_waf_certificate.certificate_1.name
  depends_on = [huaweicloud_waf_certificate.certificate_1]
}
`, testAccWafCertificateV1_conf(name))
}
