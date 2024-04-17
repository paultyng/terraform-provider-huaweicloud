package cdn

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/chnsz/golangsdk/openstack/cdn/v1/domains"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func getResourceExtensionOpts(epsId string) *domains.ExtensionOpts {
	if epsId != "" {
		return &domains.ExtensionOpts{
			EnterpriseProjectId: epsId,
		}
	}
	return nil
}

func getCdnDomainFunc(cfg *config.Config, state *terraform.ResourceState) (interface{}, error) {
	client, err := cfg.CdnV1Client(acceptance.HW_REGION_NAME)
	if err != nil {
		return nil, fmt.Errorf("error creating CDN v1 client: %s", err)
	}

	opts := getResourceExtensionOpts(state.Primary.Attributes["enterprise_project_id"])
	return domains.GetByName(client, state.Primary.Attributes["name"], opts).Extract()
}

func TestAccCdnDomain_basic(t *testing.T) {
	var (
		obj          interface{}
		resourceName = "huaweicloud_cdn_domain.test"
	)

	rc := acceptance.InitResourceCheck(
		resourceName,
		&obj,
		getCdnDomainFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheckCDN(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccCdnDomain_basic,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "type", "web"),
					resource.TestCheckResourceAttr(resourceName, "service_area", "outside_mainland_china"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.origin_protocol", "http"),
					resource.TestCheckResourceAttr(resourceName, "sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.active", "1"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.origin", "100.254.53.75"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.origin_type", "ipaddr"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.http_port", "80"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.https_port", "443"),
					resource.TestCheckResourceAttr(resourceName, "tags.key", "val"),
					resource.TestCheckResourceAttr(resourceName, "tags.foo", "bar"),
				),
			},
			{
				Config: testAccCdnDomain_update1,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "type", "download"),
					resource.TestCheckResourceAttr(resourceName, "service_area", "global"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "0"),
				),
			},
			{
				Config: testAccCdnDomain_update2,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "type", "web"),
					resource.TestCheckResourceAttr(resourceName, "service_area", "mainland_china"),
					resource.TestCheckResourceAttr(resourceName, "sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.retrieval_host", "customize.test.huaweicloud.com"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.http_port", "8001"),
					resource.TestCheckResourceAttr(resourceName, "sources.0.https_port", "8002"),
				),
			},
			{
				Config: testAccCdnDomain_update3,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "sources.#", "2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testCDNDomainImportState(resourceName),
				ImportStateVerifyIgnore: []string{
					"enterprise_project_id", "configs.0.url_signing.0.key", "configs.0.https_settings.0.certificate_body",
					"configs.0.https_settings.0.private_key", "cache_settings",
				},
			},
		},
	})
}

var testAccCdnDomain_basic = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "web"
  service_area          = "outside_mainland_china"

  configs {
    origin_protocol = "http"
  }

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
    http_port   = 80
    https_port  = 443
  }

  cache_settings {
    rules {
      rule_type = "all"
      ttl       = 365
      ttl_type  = "d"
      priority  = 2
    }
  }

  tags = {
    key = "val"
    foo = "bar"
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME)

var testAccCdnDomain_update1 = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "download"
  service_area          = "global"

  configs {
    origin_protocol = "http"
  }

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
    http_port   = 80
    https_port  = 443
  }

  cache_settings {
    follow_origin = true
    rules {
      rule_type = "file_extension"
      content   = ".jpg"
      ttl       = 0
      ttl_type  = "d"
      priority  = 3
    }
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME)

var testAccCdnDomain_update2 = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "web"
  service_area          = "mainland_china"

  configs {
    origin_protocol = "http"
  }

  sources {
    active         = 1
    origin         = "100.254.53.75"
    origin_type    = "ipaddr"
    retrieval_host = "customize.test.huaweicloud.com"
    http_port      = 8001
    https_port     = 8002
  }

  cache_settings {}
}
`, acceptance.HW_CDN_DOMAIN_NAME)

var testAccCdnDomain_update3 = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "web"
  service_area          = "mainland_china"

  sources {
    active      = 1
    origin      = "14.215.177.39"
    origin_type = "ipaddr"
  }
  sources {
    active      = 0
    origin      = "220.181.28.52"
    origin_type = "ipaddr"
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME)

// Prepare the HTTPS certificate before running this test case
// All configuration item modifications may trigger `CDN.0163`. This is a problem that we have no way to solve.
// When a `CDN.0163` error occurs, you can avoid this error by adjusting the test case configuration items.
func TestAccCdnDomain_configHttpSettings(t *testing.T) {
	var (
		obj          interface{}
		resourceName = "huaweicloud_cdn_domain.test"
	)

	rc := acceptance.InitResourceCheck(
		resourceName,
		&obj,
		getCdnDomainFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheckCDN(t)
			acceptance.TestAccPreCheckCERT(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccCdnDomain_configHttpSettings,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "configs.0.origin_protocol", "http"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.ipv6_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.range_based_retrieval_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.certificate_name", "terraform-test"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.https_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.http2_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.certificate_source", "0"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.certificate_type", "server"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.ocsp_stapling_status", "on"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.https_status", "on"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.http2_status", "on"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.quic.0.enabled", "true"),
					testAccCheckTLSVersion(resourceName, "TLSv1.1,TLSv1.2"),

					resource.TestCheckResourceAttrSet(resourceName, "configs.0.https_settings.0.certificate_body"),
					resource.TestCheckResourceAttrSet(resourceName, "configs.0.https_settings.0.private_key"),
				),
			},
			{
				Config: testAccCdnDomain_configHttpSettings_update1,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.certificate_name", "terraform-update"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.https_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.http2_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.certificate_source", "0"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.certificate_type", "server"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.ocsp_stapling_status", "off"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.https_status", "on"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.http2_status", "on"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.quic.0.enabled", "false"),
					testAccCheckTLSVersion(resourceName, "TLSv1.1,TLSv1.2,TLSv1.3"),
				),
			},
			{
				Config: testAccCdnDomain_configHttpSettings_update2,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.https_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.http2_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.https_status", "off"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.https_settings.0.http2_status", "off"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testCDNDomainImportState(resourceName),
				ImportStateVerifyIgnore: []string{
					"enterprise_project_id", "configs.0.url_signing.0.key", "configs.0.https_settings.0.certificate_body",
					"configs.0.https_settings.0.private_key", "cache_settings",
				},
			},
		},
	})
}

// The response value order of field `tls_version` will be modified.
// For example `TLSv1.1,TLSv1.2` will be modified to `TLSv1.2,TLSv1.1`.
func testAccCheckTLSVersion(n string, tlsVersion string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource (%s) not found: %s", n, rs)
		}

		tlsVersionAttr := rs.Primary.Attributes["configs.0.https_settings.0.tls_version"]
		if tlsVersionAttr == "" {
			return fmt.Errorf("attribute `configs.0.https_settings.0.tls_version` is not found from state")
		}

		tlsVersionArray := strings.Split(tlsVersion, ",")
		tlsVersionAttrArray := strings.Split(tlsVersionAttr, ",")
		sort.Strings(tlsVersionArray)
		sort.Strings(tlsVersionAttrArray)
		if reflect.DeepEqual(tlsVersionArray, tlsVersionAttrArray) {
			return nil
		}
		return fmt.Errorf("attribute 'configs.0.https_settings.0.tls_version' expected (%s), got (%s)",
			tlsVersion, tlsVersionAttr)
	}
}

var testAccCdnDomain_configHttpSettings = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "web"
  service_area          = "outside_mainland_china"
  enterprise_project_id = 0

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
  }

  configs {
    origin_protocol               = "http"
    ipv6_enable                   = true
    range_based_retrieval_enabled = "true"

    https_settings {
      certificate_name     = "terraform-test"
      certificate_body     = file("%s")
      http2_enabled        = true
      https_enabled        = true
      private_key          = file("%s")
      tls_version          = "TLSv1.1,TLSv1.2"
      certificate_source   = 0
      certificate_type     = "server"
      ocsp_stapling_status = "on"
    }

    quic {
      enabled = true
    }
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME, acceptance.HW_CDN_CERT_PATH, acceptance.HW_CDN_PRIVATE_KEY_PATH)

var testAccCdnDomain_configHttpSettings_update1 = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "web"
  service_area          = "outside_mainland_china"
  enterprise_project_id = 0

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
  }

  configs {
    origin_protocol               = "http"
    ipv6_enable                   = true
    range_based_retrieval_enabled = "true"

    https_settings {
      certificate_name     = "terraform-update"
      certificate_body     = file("%s")
      http2_enabled        = true
      https_enabled        = true
      private_key          = file("%s")
      tls_version          = "TLSv1.1,TLSv1.2,TLSv1.3"
      certificate_source   = 0
      ocsp_stapling_status = "off"
    }

    quic {
      enabled = false
    }
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME, acceptance.HW_CDN_CERT_PATH, acceptance.HW_CDN_PRIVATE_KEY_PATH)

var testAccCdnDomain_configHttpSettings_update2 = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "web"
  service_area          = "outside_mainland_china"
  enterprise_project_id = 0

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
  }

  configs {
    origin_protocol               = "http"
    ipv6_enable                   = true
    range_based_retrieval_enabled = "true"

    https_settings {
      https_enabled = false
    }
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME)

// All configuration item modifications may trigger `CDN.0163`. This is a problem that we have no way to solve.
// When a `CDN.0163` error occurs, you can avoid this error by adjusting the test case configuration items.
func TestAccCdnDomain_configs(t *testing.T) {
	var (
		obj          interface{}
		resourceName = "huaweicloud_cdn_domain.test"
	)

	rc := acceptance.InitResourceCheck(
		resourceName,
		&obj,
		getCdnDomainFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheckCDN(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccCdnDomain_configs,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "configs.0.origin_protocol", "http"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.ipv6_enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.range_based_retrieval_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.description", "test description"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.slice_etag_status", "on"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.origin_receive_timeout", "60"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.cache_url_parameter_filter.0.type", "ignore_url_params"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.retrieval_request_header.0.name", "test-name"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.type", "type_a"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.sign_method", "md5"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.match_type", "all"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.sign_arg", "Psd_123"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.key", "A27jtfSTy13q7A0UnTA9vpxYXEb"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.backup_key", "S36klgTFa60q3V8DmSK2hwfBOYp"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.time_format", "dec"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.expire_time", "0"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.status", "on"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.inherit_config.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.inherit_config.0.inherit_type", "m3u8"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.inherit_config.0.inherit_time_type", "sys_time"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.inherit_config.0.status", "on"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.compress.0.status", "off"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.force_redirect.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.force_redirect.0.type", "http"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.force_redirect.0.redirect_code", "301"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.force_redirect.0.status", "on"),

					resource.TestCheckResourceAttr(resourceName, "configs.0.referer.0.type", "white"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.referer.0.include_empty", "false"),

					resource.TestCheckResourceAttr(resourceName, "configs.0.flexible_origin.#", "2"),

					resource.TestCheckResourceAttr(resourceName, "configs.0.remote_auth.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.auth_failed_status", "403"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.auth_server", "https://testdomain.com"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.auth_success_status", "200"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.file_type_setting", "specific_file"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.request_method", "GET"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.reserve_args", "k1|k2|key33"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.reserve_args_setting", "reserve_specific_args"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.reserve_headers_setting", "reserve_specific_headers"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.reserve_headers", "key1|key2"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.response_status", "403"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.specified_file_type", "jpg|mp4"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.timeout", "50"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.timeout_action", "pass"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.add_custom_args_rules.#", "2"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.add_custom_headers_rules.#", "2"),
				),
			},
			{
				Config: testAccCdnDomain_configsUpdate1,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "configs.0.origin_protocol", "follow"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.ipv6_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.range_based_retrieval_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.description", "update description"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.slice_etag_status", "off"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.origin_receive_timeout", "30"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.cache_url_parameter_filter.0.type", "del_params"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.cache_url_parameter_filter.0.value", "test_value"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.retrieval_request_header.0.name", "test-name-update"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.retrieval_request_header.0.value", "test-val-update"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.retrieval_request_header.0.action", "set"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.type", "type_c2"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.sign_method", "sha256"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.match_type", "all"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.sign_arg", "Dma_001"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.key", "3P7k9s4r0aey9CB1mvvDHG2"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.backup_key", "5F8a6c3r1xgp7DL0jkeBYZ4"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.time_format", "hex"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.expire_time", "31536000"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.status", "on"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.inherit_config.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.inherit_config.0.inherit_type", "mpd"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.inherit_config.0.inherit_time_type", "parent_url_time"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.inherit_config.0.status", "on"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.compress.0.status", "off"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.force_redirect.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.force_redirect.0.type", "http"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.force_redirect.0.redirect_code", "302"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.force_redirect.0.status", "on"),

					resource.TestCheckResourceAttr(resourceName, "configs.0.referer.0.type", "black"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.referer.0.include_empty", "true"),

					resource.TestCheckResourceAttr(resourceName, "configs.0.flexible_origin.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.flexible_origin.0.match_type", "file_path"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.flexible_origin.0.match_pattern", "/test/folder01;/test/folder02"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.flexible_origin.0.priority", "3"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.flexible_origin.0.back_sources.0.http_port", "83"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.flexible_origin.0.back_sources.0.https_port", "470"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.flexible_origin.0.back_sources.0.ip_or_domain", "www.hshs.cdd"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.flexible_origin.0.back_sources.0.sources_type", "domain"),

					resource.TestCheckResourceAttr(resourceName, "configs.0.remote_auth.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.auth_failed_status", "503"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.auth_server", "https://testdomain-update.com"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.auth_success_status", "302"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.file_type_setting", "all"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.request_method", "POST"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.reserve_args_setting", "reserve_all_args"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.reserve_headers_setting", "reserve_all_headers"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.response_status", "206"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.timeout", "3000"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.timeout_action", "forbid"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.add_custom_args_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.add_custom_headers_rules.#", "0"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.add_custom_args_rules.0.key", "http_user_agent"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.add_custom_args_rules.0.type", "nginx_preset_var"),
					resource.TestCheckResourceAttr(resourceName,
						"configs.0.remote_auth.0.remote_auth_rules.0.add_custom_args_rules.0.value", "$server_protocol"),
				),
			},
			{
				Config: testAccCdnDomain_configsUpdate2,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "configs.0.description", ""),
					resource.TestCheckResourceAttr(resourceName, "configs.0.slice_etag_status", "on"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.origin_receive_timeout", "5"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.flexible_origin.#", "0"),

					resource.TestCheckResourceAttr(resourceName, "configs.0.remote_auth.0.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.remote_auth.0.remote_auth_rules.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.status", "on"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.inherit_config.0.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.inherit_config.0.status", "off"),
				),
			},
			{
				Config: testAccCdnDomain_configsUpdate3,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "configs.0.flexible_origin.#", "0"),

					resource.TestCheckResourceAttr(resourceName, "configs.0.remote_auth.0.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.remote_auth.0.remote_auth_rules.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.url_signing.0.status", "off"),

					resource.TestCheckResourceAttr(resourceName, "configs.0.force_redirect.0.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.force_redirect.0.status", "off"),

					resource.TestCheckResourceAttr(resourceName, "configs.0.referer.0.type", "off"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testCDNDomainImportState(resourceName),
				ImportStateVerifyIgnore: []string{
					"enterprise_project_id", "configs.0.url_signing.0.key", "configs.0.url_signing.0.backup_key",
					"configs.0.https_settings.0.certificate_body", "configs.0.https_settings.0.private_key", "cache_settings",
				},
			},
		},
	})
}

var testAccCdnDomain_configs = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "web"
  service_area          = "outside_mainland_china"
  enterprise_project_id = 0

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
  }

  configs {
    origin_protocol               = "http"
    ipv6_enable                   = true
    range_based_retrieval_enabled = true
    description                   = "test description"
    origin_receive_timeout        = "60"

    cache_url_parameter_filter {
      type = "ignore_url_params"
    }

    retrieval_request_header {
      name   = "test-name"
      value  = "test-val"
      action = "set"
    }

    http_response_header {
      name   = "test-name"
      value  = "test-val"
      action = "set"
    }

    url_signing {
      enabled     = true
      type        = "type_a"
      sign_method = "md5"
      match_type  = "all"
      sign_arg    = "Psd_123"
      key         = "A27jtfSTy13q7A0UnTA9vpxYXEb"
      backup_key  = "S36klgTFa60q3V8DmSK2hwfBOYp"
      time_format = "dec"
      expire_time = 0

      inherit_config {
        enabled           = true
        inherit_type      = "m3u8"
        inherit_time_type = "sys_time"
      }
    }

    compress {
      enabled = false
    }

    force_redirect {
      enabled       = true
      type          = "http"
      redirect_code = 301
    }

    referer {
      type          = "white"
      value         = "*.common.com,192.187.2.43,www.test.top:4990"
      include_empty = false
    }

    flexible_origin {
      match_type = "all"
      priority   = 1

      back_sources {
        http_port    = 1
        https_port   = 65535
        ip_or_domain = "165.132.12.2"
        sources_type = "ipaddr"
      }
    }

    flexible_origin {
      match_type    = "file_extension"
      match_pattern = ".jpg;.zip;.exe"
      priority      = 2

      back_sources {
        http_port    = 65535
        https_port   = 1
        ip_or_domain = "165.5.1.4"
        sources_type = "ipaddr"
      }
    }

    remote_auth {
      enabled = true

      remote_auth_rules {
        auth_failed_status      = "403"
        auth_server             = "https://testdomain.com"
        auth_success_status     = "200"
        file_type_setting       = "specific_file"
        request_method          = "GET"
        reserve_args            = "k1|k2|key33"
        reserve_args_setting    = "reserve_specific_args"
        reserve_headers_setting = "reserve_specific_headers"
        reserve_headers         = "key1|key2"
        response_status         = "403"
        specified_file_type     = "jpg|mp4"
        timeout                 = 50
        timeout_action          = "pass"

        add_custom_args_rules {
          key   = "http_user_agent"
          type  = "nginx_preset_var"
          value = "$http_host"
        }

        add_custom_args_rules {
          key   = "args_custom_key"
          type  = "custom_var"
          value = "args_custom_value"
        }

        add_custom_headers_rules {
          key   = "http_user_agent"
          type  = "nginx_preset_var"
          value = "$remote_addr"
        }

        add_custom_headers_rules {
          key   = "headers_custom_key"
          type  = "custom_var"
          value = "headers_custom_value"
        }
      }
    }
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME)

var testAccCdnDomain_configsUpdate1 = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "web"
  service_area          = "outside_mainland_china"
  enterprise_project_id = 0

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
  }

  configs {
    origin_protocol               = "follow"
    ipv6_enable                   = false
    range_based_retrieval_enabled = false
    description                   = "update description"
    slice_etag_status             = "off"
    origin_receive_timeout        = "30"

    cache_url_parameter_filter {
      type  = "del_params"
      value = "test_value"
    }

    retrieval_request_header {
      name   = "test-name-update"
      value  = "test-val-update"
      action = "set"
    }

    http_response_header {
      name   = "Content-Disposition"
      value  = "test-val-update"
      action = "set"
    }

    url_signing {
      enabled     = true
      type        = "type_c2"
      sign_method = "sha256"
      match_type  = "all"
      sign_arg    = "Dma_001"
      key         = "3P7k9s4r0aey9CB1mvvDHG2"
      backup_key  = "5F8a6c3r1xgp7DL0jkeBYZ4"
      time_format = "hex"
      expire_time = 31536000

      inherit_config {
        enabled           = true
        inherit_type      = "mpd"
        inherit_time_type = "parent_url_time"
      }
    }

    compress {
      enabled = false
    }

    force_redirect {
      enabled       = true
      type          = "http"
      redirect_code = 302
    }

    referer {
      type          = "black"
      value         = "*.common.com,192.187.2.43"
      include_empty = true
    }

    flexible_origin {
      match_type    = "file_path"
      match_pattern = "/test/folder01;/test/folder02"
      priority      = 3

      back_sources {
        http_port    = 83
        https_port   = 470
        ip_or_domain = "www.hshs.cdd"
        sources_type = "domain"
      }
    }

    remote_auth {
      enabled = true

      remote_auth_rules {
        auth_failed_status      = "503"
        auth_server             = "https://testdomain-update.com"
        auth_success_status     = "302"
        file_type_setting       = "all"
        request_method          = "POST"
        reserve_args_setting    = "reserve_all_args"
        reserve_headers_setting = "reserve_all_headers"
        response_status         = "206"
        timeout                 = 3000
        timeout_action          = "forbid"

        add_custom_args_rules {
          key   = "http_user_agent"
          type  = "nginx_preset_var"
          value = "$server_protocol"
        }
      }
    }
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME)

var testAccCdnDomain_configsUpdate2 = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "web"
  service_area          = "outside_mainland_china"
  enterprise_project_id = 0

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
  }

  configs {
    origin_protocol               = "follow"
    ipv6_enable                   = false
    range_based_retrieval_enabled = false
    slice_etag_status             = "on"
    origin_receive_timeout        = "5"

    remote_auth {
      enabled = false
    }

    url_signing {
      enabled     = true
      type        = "type_c2"
      sign_method = "sha256"
      match_type  = "all"
      sign_arg    = "web506"
      key         = "3P7k9s4r0aey9CB1mvvDHG2"
      backup_key  = "5F8a6c3r1xgp7DL0jkeBYZ4"
      time_format = "hex"
      expire_time = 31536000

      inherit_config {
        enabled = false
      }
    }
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME)

var testAccCdnDomain_configsUpdate3 = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "web"
  service_area          = "outside_mainland_china"
  enterprise_project_id = 0

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
  }

  configs {
    origin_protocol               = "follow"
    ipv6_enable                   = false
    range_based_retrieval_enabled = false

    remote_auth {
      enabled = false
    }

    url_signing {
      enabled = false
    }

    force_redirect {
      enabled = false
    }

    referer {
      type = "off"
    }
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME)

// All configuration item modifications may trigger `CDN.0163`. This is a problem that we have no way to solve.
// When a `CDN.0163` error occurs, you can avoid this error by adjusting the test case configuration items.
func TestAccCdnDomain_configTypeWholeSite(t *testing.T) {
	var (
		obj          interface{}
		resourceName = "huaweicloud_cdn_domain.test"
	)

	rc := acceptance.InitResourceCheck(
		resourceName,
		&obj,
		getCdnDomainFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheckCDN(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccCdnDomain_wholeSite,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "type", "wholeSite"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.websocket.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.websocket.0.timeout", "1"),
				),
			},
			{
				Config: testAccCdnDomain_wholeSiteUpdate1,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "type", "wholeSite"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.websocket.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.websocket.0.timeout", "300"),
				),
			},
			{
				Config: testAccCdnDomain_wholeSiteUpdate2,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "type", "wholeSite"),
					resource.TestCheckResourceAttr(resourceName, "configs.0.websocket.0.enabled", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testCDNDomainImportState(resourceName),
				ImportStateVerifyIgnore: []string{
					"enterprise_project_id", "configs.0.url_signing.0.key", "configs.0.https_settings.0.certificate_body",
					"configs.0.https_settings.0.private_key", "cache_settings",
				},
			},
		},
	})
}

var testAccCdnDomain_wholeSite = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "wholeSite"
  service_area          = "outside_mainland_china"
  enterprise_project_id = 0

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
  }

  configs {
    origin_protocol = "http"
    ipv6_enable     = true

    websocket {
      enabled = true
      timeout = 1
    }
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME)

var testAccCdnDomain_wholeSiteUpdate1 = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "wholeSite"
  service_area          = "outside_mainland_china"
  enterprise_project_id = 0

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
  }

  configs {
    origin_protocol = "http"
    ipv6_enable     = true

    websocket {
      enabled = true
      timeout = 300
    }
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME)

var testAccCdnDomain_wholeSiteUpdate2 = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "wholeSite"
  service_area          = "outside_mainland_china"
  enterprise_project_id = 0

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
  }

  configs {
    origin_protocol = "http"
    ipv6_enable     = true

    websocket {
      enabled = false
    }
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME)

func TestAccCdnDomain_epsID_migrate(t *testing.T) {
	var (
		obj          interface{}
		resourceName = "huaweicloud_cdn_domain.test"
	)

	rc := acceptance.InitResourceCheck(
		resourceName,
		&obj,
		getCdnDomainFunc,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acceptance.TestAccPreCheckCDN(t)
			acceptance.TestAccPreCheckEpsID(t)
		},
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: TestAccCdnDomain_epsID_basic,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "enterprise_project_id", "0"),
				),
			},
			{
				Config: TestAccCdnDomain_epsID_update1,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "enterprise_project_id", acceptance.HW_ENTERPRISE_PROJECT_ID_TEST),
				),
			},
			{
				Config: TestAccCdnDomain_epsID_update2,
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", acceptance.HW_CDN_DOMAIN_NAME),
					resource.TestCheckResourceAttr(resourceName, "enterprise_project_id", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testCDNDomainImportState(resourceName),
				ImportStateVerifyIgnore: []string{
					"enterprise_project_id", "configs.0.url_signing.0.key", "configs.0.https_settings.0.certificate_body",
					"configs.0.https_settings.0.private_key", "cache_settings",
				},
			},
		},
	})
}

var TestAccCdnDomain_epsID_basic = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "wholeSite"
  service_area          = "outside_mainland_china"
  enterprise_project_id = "0"

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME)

var TestAccCdnDomain_epsID_update1 = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "wholeSite"
  service_area          = "outside_mainland_china"
  enterprise_project_id = "%s"

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME, acceptance.HW_ENTERPRISE_PROJECT_ID_TEST)

var TestAccCdnDomain_epsID_update2 = fmt.Sprintf(`
resource "huaweicloud_cdn_domain" "test" {
  name                  = "%s"
  type                  = "wholeSite"
  service_area          = "outside_mainland_china"
  enterprise_project_id = "0"

  sources {
    active      = 1
    origin      = "100.254.53.75"
    origin_type = "ipaddr"
  }
}
`, acceptance.HW_CDN_DOMAIN_NAME)

// testCDNDomainImportState use to return an ID using `name`
func testCDNDomainImportState(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("resource (%s) not found", name)
		}

		return rs.Primary.Attributes["name"], nil
	}
}
