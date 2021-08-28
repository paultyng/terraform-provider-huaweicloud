package waf

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/huaweicloud/golangsdk/openstack/waf/v1/certificates"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccWafCertificateV1_basic(t *testing.T) {
	var certificate certificates.Certificate
	resourceName := "huaweicloud_waf_certificate.certificate_1"
	name := acceptance.RandomAccResourceName()
	updateName := name + "_update"

	rc := acceptance.InitResourceCheck(
		resourceName,
		&certificate,
		func(c *config.Config, state *terraform.ResourceState) (interface{}, error) {
			wafClient, err := c.WafV1Client(acceptance.HW_REGION_NAME)
			if err != nil {
				return nil, fmt.Errorf("error creating HuaweiCloud WAF client: %s", err)
			}
			return certificates.Get(wafClient, state.Primary.ID).Extract()
		},
	)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAccPreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccWafCertificateV1_conf(name),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceMapAttr(map[string]string{
						"name":       name,
						"expiration": acceptance.CHECKSET,
					}),
				),
			},
			{
				Config: testAccWafCertificateV1_conf(updateName),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceMapAttr(map[string]string{
						"name": updateName,
					}),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"certificate", "private_key"},
			},
		},
	})
}

func testAccWafCertificateV1_conf(name string) string {
	return fmt.Sprintf(`
resource "huaweicloud_waf_certificate" "certificate_1" {
  name = "%s"

  certificate = <<EOT
-----BEGIN CERTIFICATE-----
MIIDazCCAlOgAwIBAgIUehx07qc7un7IB7/X9lHCLkt/jPowDQYJKoZIhvcNAQEL
BQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yMTA1MzEwOTI1NTJaFw0yMjA1
MzEwOTI1NTJaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEB
AQUAA4IBDwAwggEKAoIBAQCvmuH5ViGtGOlevJ8vOoN3Ak4pp3SescdAfQa/r4cO
z/bmBqBcZJTX9HODhiQzdemyLLs9aOkQXYIc8OrcaIsjns92XITVDpFW0ThGyjhT
ZdELj9LsbIcVzNPPclTcebZBlzAyX0oLqpHK73OUYQY2E6l44U9G8Id763Bnws9N
Rn3cg0qufrlUgdim/pYZ8ubjvlDJ9eEIhcsu9zu8c8i2+8qLjEsonx5PrwzNlYP3
JqAmZ2dcbQeSPfv5U6ZceKEZfegK+Cxv4rFd5F4Rdxl+SAIY+6mr7qu1dAlcVMLS
QcLlJLRWQ5NmqL9xju7Fbj2VZt+L6nb512iKaedPo2GfAgMBAAGjUzBRMB0GA1Ud
DgQWBBR5yzB/GujpSlLrn0l2p+BslakGzjAfBgNVHSMEGDAWgBR5yzB/GujpSlLr
n0l2p+BslakGzjAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQCj
TqvcIk0Au/yOxOfIGUZzVkTiORxbwATAfRN6n/+mrgWnIbHG4XFqqjmFr7gGvHeH
+BuyU06VXJgKYaPUqbYl7eBd4Spm5v3Wq7C7i96dOHmG8fVcjQnTWleyEmUsEarv
A6/lhTqXV1+AuNUaH+9EbBUBsrCHGLkECBMKl0+cJN8lo5XncAtp7z1+O/Mn0Zi6
XyNOyvqcmmn8HUkSIS4RlJ2ohuZN6oFC3sYX9g9Vo++IkjGl3dRbf/7JutqBGHNE
RVKoPyaivymDDIIL/qSy/Pi2s0hzUhwc1M8td0K/AMxyeigwNG7mTH0RzX32bUkf
ZoURg5WiRskhtHEvBsLF
-----END CERTIFICATE-----
EOT

  private_key = <<EOT
-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCvmuH5ViGtGOle
vJ8vOoN3Ak4pp3SescdAfQa/r4cOz/bmBqBcZJTX9HODhiQzdemyLLs9aOkQXYIc
8OrcaIsjns92XITVDpFW0ThGyjhTZdELj9LsbIcVzNPPclTcebZBlzAyX0oLqpHK
73OUYQY2E6l44U9G8Id763Bnws9NRn3cg0qufrlUgdim/pYZ8ubjvlDJ9eEIhcsu
9zu8c8i2+8qLjEsonx5PrwzNlYP3JqAmZ2dcbQeSPfv5U6ZceKEZfegK+Cxv4rFd
5F4Rdxl+SAIY+6mr7qu1dAlcVMLSQcLlJLRWQ5NmqL9xju7Fbj2VZt+L6nb512iK
aedPo2GfAgMBAAECggEAeMAvDS3uAEI2Dx/y8h3xUn9yUfBFH+6tTanrXxoK6+OT
Kj96O64qL4l3eQRflkdJiGx74FFomglCtDXxudflfXvxurkJ2hunUySQ5xScwLQt
mB6w8kP6a8IqD+bVdbn32ohk6u5dU0JZ+ErJlklVZRAGJAoCYox5DXwrEh6CP+bJ
pItgjv71tEEnX5sScQwV7FMRbjsPzXoJp8vCQjlUdetM1fk9rs3R2WSeFbPgLLtC
xY0+8Hexy0q6BLmyPZvFCaVIAzAHCYeCyzPK3xcm4odbrBmRL/amOg24CCny065N
MU9RFhEjQsY1RaK7dgkvjsntUZvU+aDcL8o6djOTuQKBgQDlDN/j2ntpGCtbTWH0
cVTW13Ze7U7iE3BfDO3m4VYP3Xi/v5FI8nHlmLrcl30H1dPKvMTec0dCBOqD1wzF
KiqHy8ELowO2CbXMYJpjuPzXH40/AE3eOJVTJM8mOeuFdeFgYCd/9cB7o5jfTA5Y
4zj8EmcRzsH1rNSnvo7/O9q6+wKBgQDERDSvP8RScEbzDKuN6uhzj1K2CAEnY6//
rDA1so18UhAie9NcAvlKa46jQTOcYD77g5h0WSlNt9ZbK9Plq9CY9psI0KNqN3Fl
YVKOKdD5m6Rifmg+lt8KLc/WocQ10DXpPTXzzuRlN/TaMDdN2pedEre/0AAMs8Ia
MIUnu4oyrQKBgQC6b6BNdqi9Ak9IIdR5g0XrGbXfzolGu0vcEkoSg5fpkfuXF/bJ
yY2rtIVkyGmc1w9tFfmol2yI8Ddy2LgsRAYaQl7/edCre3vev0LrqMck0ynE/hpj
purkojF6i+qI10p7h8ie/wmNmbv1BZMoBst7Yf9DH2gA8IynfRQn7DA9wQKBgGaU
M2kJDgX8UsjDbYKuLTIAzb0AMAIzUxBxIX1fRh2dEnvDdjOYBk1EK/fdoyjvENwJ
6ouc8j6BgBKEtKpMg6j+8wbHbTGdqrHPDQPqjSN4mpEz+i4EUqySRxep0tBBc3vl
FybHko3okhvbqXwSbL2Ww90HzI7XAPMJOv8KQO+9AoGBAJxxftNWvypBXGkPCdH2
f3ikvT2Vef9QZjqkvtipCecAkjM6ReLshVsdqFSv/ZmsVUeNKoTHvX2GnhweJM44
x7N2mFK4skBzVtMVbjAHVjG78UitVu+FrzqGreaJXHaduhgUH2iFWfw09joOotAM
X7ioLbTeWGBqFM+C80PkdBNp
-----END PRIVATE KEY-----
EOT
}
`, name)
}
