package main

import (
	"testing"
)

func TestZoneSettingsTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "simple attributes",
			Config: `
resource "cloudflare_zone_settings_override" "zone_settings" {
  zone_id = var.zone_id

  settings {
    automatic_https_rewrites = var.automatic_https_rewrites
    ssl                      = var.ssl
  }
}`,
			Expected: []string{`
resource "cloudflare_zone_setting" "zone_settings_automatic_https_rewrites" {
  zone_id    = var.zone_id
  setting_id = "automatic_https_rewrites"
  value      = var.automatic_https_rewrites
}`, `
resource "cloudflare_zone_setting" "zone_settings_ssl" {
  zone_id    = var.zone_id
  setting_id = "ssl"
  value      = var.ssl
}`, `
import {
  to = cloudflare_zone_setting.zone_settings_automatic_https_rewrites
  id = "${var.zone_id}/automatic_https_rewrites"
}`, `
import {
  to = cloudflare_zone_setting.zone_settings_ssl
  id = "${var.zone_id}/ssl"
}`,
			},
		},
		{
			Name: "with security header block",
			Config: `
resource "cloudflare_zone_settings_override" "zone_settings" {
  zone_id = var.zone_id

  settings {
    ssl = var.ssl

    security_header {
      enabled = var.security_header_enabled
      max_age = var.security_header_max_age
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_zone_setting" "zone_settings_ssl" {
  zone_id    = var.zone_id
  setting_id = "ssl"
  value      = var.ssl
}`, `
resource "cloudflare_zone_setting" "zone_settings_security_header" {
  zone_id    = var.zone_id
  setting_id = "security_header"
  value = {
    enabled  =  var.security_header_enabled
    max_age  =  var.security_header_max_age
  }
}`, `
import {
  to = cloudflare_zone_setting.zone_settings_security_header
  id = "${var.zone_id}/security_header"
}`,
			},
		},
		{
			Name: "with nel block",
			Config: `
resource "cloudflare_zone_settings_override" "zone_settings" {
  zone_id = var.zone_id

  settings {
    nel {
      enabled = var.enable_network_error_logging
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_zone_setting" "zone_settings_nel" {
  zone_id    = var.zone_id
  setting_id = "nel"
  value = {
    enabled  =  var.enable_network_error_logging
  }
}`, `
import {
  to = cloudflare_zone_setting.zone_settings_nel
  id = "${var.zone_id}/nel"
}`,
			},
		},
		{
			Name: "zero_rtt setting name mapping",
			Config: `
resource "cloudflare_zone_settings_override" "test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  settings {
    zero_rtt = "on"
  }
}`,
			Expected: []string{`
resource "cloudflare_zone_setting" "test_zero_rtt" {
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
  setting_id = "0rtt"
  value      = "on"
}`, `
import {
  to = cloudflare_zone_setting.test_zero_rtt
  id = "${"0da42c8d2132a9ddaf714f9e7c920711"}/0rtt"
}`,
			},
		},
		{
			Name: "empty settings block",
			Config: `
resource "cloudflare_zone_settings_override" "zone_settings" {
  zone_id = var.zone_id

  settings {
  }
}`,
			Expected: []string{},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}
