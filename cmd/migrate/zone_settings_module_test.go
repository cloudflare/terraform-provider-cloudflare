package main

import (
	"strings"
	"testing"
)

func TestExpandZoneSettingsModules(t *testing.T) {
	tests := []TestCase{
		{
			Name: "expand simple zone_settings module",
			Config: `module "zone_settings" {
  source         = "../zone_settings"
  zone_id        = cloudflare_zone.example_com.id
  security_level = "high"
  ssl            = "origin_pull"
}`,
			Expected: []string{
				`resource "cloudflare_zone_setting" "zone_settings_security_level" {
  zone_id    = cloudflare_zone.example_com.id
  setting_id = "security_level"
  value      = "high"
}`,
				`import {
  to = cloudflare_zone_setting.zone_settings_security_level
  id = "${cloudflare_zone.example_com.id}/security_level"
}`,
				`resource "cloudflare_zone_setting" "zone_settings_ssl" {
  zone_id    = cloudflare_zone.example_com.id
  setting_id = "ssl"
  value      = "origin_pull"
}`,
				`import {
  to = cloudflare_zone_setting.zone_settings_ssl
  id = "${cloudflare_zone.example_com.id}/ssl"
}`,
			},
		},
		{
			Name: "expand zone_settings module with zero_rtt mapping",
			Config: `module "zone_settings" {
  source   = "../zone_settings"
  zone_id  = var.zone_id
  zero_rtt = "on"
}`,
			Expected: []string{
				`resource "cloudflare_zone_setting" "zone_settings_zero_rtt" {
  zone_id    = var.zone_id
  setting_id = "0rtt"
  value      = "on"
}`,
				`import {
  to = cloudflare_zone_setting.zone_settings_zero_rtt
  id = "${var.zone_id}/0rtt"
}`,
			},
		},
		{
			Name: "expand zone_settings module with NEL setting",
			Config: `module "zone_settings" {
  source                        = "../zone_settings"
  zone_id                       = var.zone_id
  enable_network_error_logging  = true
}`,
			Expected: []string{
				`resource "cloudflare_zone_setting" "zone_settings_nel" {
  zone_id    = var.zone_id
  setting_id = "nel"
  value      = { enabled = true }
}`,
				`import {
  to = cloudflare_zone_setting.zone_settings_nel
  id = "${var.zone_id}/nel"
}`,
			},
		},
		{
			Name: "expand zone_settings module with security headers",
			Config: `module "zone_settings" {
  source                              = "../zone_settings"
  zone_id                             = var.zone_id
  security_header_enabled             = true
  security_header_include_subdomains  = true
  security_header_max_age            = 31536000
}`,
			Expected: []string{
				`resource "cloudflare_zone_setting" "zone_settings_security_header" {
  zone_id    = var.zone_id
  setting_id = "security_header"
  value = {
    strict_transport_security = {
      enabled = true
      include_subdomains = true
      max_age = 31536000
    }
  }
}`,
				`import {
  to = cloudflare_zone_setting.zone_settings_security_header
  id = "${var.zone_id}/security_header"
}`,
			},
		},
	}

	// Test the expand function directly since it's string-based
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := expandZoneSettingsModules(test.Config, false) // skipImports = false

			// Check that each expected string is present in the result
			for _, expected := range test.Expected {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain '%s', but it didn't.\nActual output:\n%s", expected, result)
				}
			}
		})
	}
}

func TestExpandZoneSettingsModulesSkipImports(t *testing.T) {
	tests := []TestCase{
		{
			Name: "skip imports - expand simple zone_settings module",
			Config: `module "zone_settings" {
  source         = "../zone_settings"
  zone_id        = cloudflare_zone.example_com.id
  security_level = "high"
  ssl            = "origin_pull"
}`,
			Expected: []string{
				`resource "cloudflare_zone_setting" "zone_settings_security_level" {`,
				`zone_id    = cloudflare_zone.example_com.id`,
				`setting_id = "security_level"`,
				`value      = "high"`,
				`resource "cloudflare_zone_setting" "zone_settings_ssl" {`,
				`zone_id    = cloudflare_zone.example_com.id`,
				`setting_id = "ssl"`,
				`value      = "origin_pull"`,
			},
		},
	}

	// Test the expand function with skipImports = true
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := expandZoneSettingsModules(test.Config, true) // skipImports = true

			// Check that each expected string is present in the result
			for _, expected := range test.Expected {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain '%s', but it didn't.\nActual output:\n%s", expected, result)
				}
			}

			// Verify that no import blocks are present
			if strings.Contains(result, "import {") {
				t.Errorf("Expected no import blocks when skipImports=true, but found some.\nActual output:\n%s", result)
			}
		})
	}
}

