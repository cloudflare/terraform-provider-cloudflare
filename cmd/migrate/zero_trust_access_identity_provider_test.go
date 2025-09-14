package main

import (
	"strings"
	"testing"
)

func TestTransformZeroTrustAccessIdentityProvider(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		contains []string
		notContains []string
	}{
		{
			name: "Resource name transformation",
			input: `resource "cloudflare_access_identity_provider" "test" {
  account_id = "test"
  name = "test"
  type = "onetimepin"
}`,
			contains: []string{
				`resource "cloudflare_zero_trust_access_identity_provider" "test"`,
				"config     = {}",
			},
			notContains: []string{
				`resource "cloudflare_access_identity_provider"`,
			},
		},
		{
			name: "Config block to object transformation",
			input: `resource "cloudflare_access_identity_provider" "test" {
  account_id = "test"
  name = "test"
  type = "github"
  config {
    client_id = "test"
    client_secret = "secret"
  }
}`,
			contains: []string{
				"config = {",
				`client_id     = "test"`,
				`client_secret = "secret"`,
			},
			notContains: []string{
				"config {",
			},
		},
		{
			name: "Certificate field transformation",
			input: `resource "cloudflare_access_identity_provider" "test" {
  account_id = "test"
  name = "test"
  type = "saml"
  config {
    idp_public_cert = "CERT123"
  }
}`,
			contains: []string{
				`idp_public_certs = ["CERT123"]`,
			},
			notContains: []string{
				`idp_public_cert = "CERT123"`,
			},
		},
		{
			name: "SCIM config block to object transformation",
			input: `resource "cloudflare_access_identity_provider" "test" {
  account_id = "test"
  name = "test"
  type = "azureAD"
  scim_config {
    enabled = true
  }
}`,
			contains: []string{
				"scim_config = {",
				"enabled = true",
			},
			notContains: []string{
				"scim_config {",
			},
		},
		{
			name: "Deprecated field removal",
			input: `resource "cloudflare_access_identity_provider" "test" {
  account_id = "test"
  name = "test"
  type = "azureAD"
  config {
    api_token = "deprecated"
    client_id = "test"
  }
  scim_config {
    enabled = true
    group_member_deprovision = true
  }
}`,
			contains: []string{
				`client_id = "test"`,
				"enabled = true",
			},
			notContains: []string{
				`api_token = "deprecated"`,
				"group_member_deprovision = true",
			},
		},
		{
			name: "Type validation - remove sign_request for azureAD",
			input: `resource "cloudflare_access_identity_provider" "test" {
  account_id = "test"
  name = "test"
  type = "azureAD"
  config {
    client_id = "test"
    sign_request = false
    conditional_access_enabled = true
    directory_id = "abc123"
    support_groups = false
  }
}`,
			contains: []string{
				`client_id`,
				`"test"`,
				"conditional_access_enabled",
				"true",
				`directory_id`,
				`"abc123"`,
				"support_groups",
				"false",
			},
			notContains: []string{
				"sign_request = false",
			},
		},
		{
			name: "Type validation - remove azureAD-specific fields for saml",
			input: `resource "cloudflare_access_identity_provider" "test" {
  account_id = "test"
  name = "test"
  type = "saml"
  config {
    client_id = "test"
    sign_request = true
    conditional_access_enabled = true
    directory_id = "abc123"
    support_groups = false
  }
}`,
			contains: []string{
				`client_id`,
				`"test"`,
				"sign_request",
				"true",
			},
			notContains: []string{
				"conditional_access_enabled",
				`directory_id`,
				"support_groups",
			},
		},
		{
			name: "Type validation - remove both for onetimepin",
			input: `resource "cloudflare_access_identity_provider" "test" {
  account_id = "test"
  name = "test"
  type = "onetimepin"
  config {
    sign_request = false
    conditional_access_enabled = true
    directory_id = "abc123"
    support_groups = false
  }
}`,
			contains: []string{
				"config",
				"{}",
			},
			notContains: []string{
				"sign_request",
				"conditional_access_enabled", 
				`directory_id`,
				"support_groups",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := transformFileWithoutImports([]byte(test.input), "test.tf")
			if err != nil {
				t.Fatalf("transformFile failed: %v", err)
			}
			result := string(output)
			
			// Check that all expected strings are present
			for _, expected := range test.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected result to contain %q, but it didn't.\nResult:\n%s", expected, result)
				}
			}
			
			// Check that all unwanted strings are absent
			for _, unwanted := range test.notContains {
				if strings.Contains(result, unwanted) {
					t.Errorf("Expected result to NOT contain %q, but it did.\nResult:\n%s", unwanted, result)
				}
			}
		})
	}
}