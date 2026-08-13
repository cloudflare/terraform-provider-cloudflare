package main

import (
	"testing"
)

// Configuration transformation tests - test the full migration pipeline
// including both YAML transformations (block-to-list) and resource renames
func TestZeroTrustAccessPolicyConfigTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "transforms_include_block_to_list",
			Config: `
resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "test-account"
  name       = "test-policy"
  
  include {
    email = ["user@example.com"]
  }
}`,
			Expected: []string{`include = [{ email = { email = "user@example.com" } }]`},
		},
		{
			Name: "transforms_multiple_include_blocks",
			Config: `
resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "test-account"
  name       = "test-policy"
  
  include {
    email = ["user1@example.com"]
  }
  
  include {
    email = ["user2@example.com"]
  }
}`,
			Expected: []string{`include = [{ email = { email = "user1@example.com" } }, { email = { email = "user2@example.com" } }]`},
		},
		{
			Name: "transforms_exclude_and_require_blocks",
			Config: `
resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "test-account"
  name       = "test-policy"
  
  include {
    everyone = true
  }
  
  exclude {
    email = ["blocked@example.com"]
  }
  
  require {
    email_domain = ["example.com"]
  }
}`,
			Expected: []string{
				`include = [{ everyone = {} }]`,
				`exclude = [{ email = { email = "blocked@example.com" } }]`,
				`require = [{ email_domain = { domain = "example.com" } }]`,
			},
		},
		{
			Name: "transforms_nested_identity_provider_blocks",
			Config: `
resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "test-account"
  name       = "test-policy"
  
  include {
    azure {
      id = "azure-id"
      identity_provider_id = "provider-id"
    }
  }
}`,
			Expected: []string{`include = [{
    azure = {
      id                   = "azure-id"
      identity_provider_id = "provider-id"
    }
  }]`},
		},
		{
			Name: "renames_access_policy_to_zero_trust",
			Config: `
resource "cloudflare_access_policy" "test" {
  account_id = "test-account"
  name       = "test-policy"
  
  include {
    email = ["user@example.com"]
  }
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test"`, `include    = [{ email = { email = "user@example.com" } }]`},
		},
	}

	// Use RunFullTransformationTests to apply both YAML and Go transformations
	RunFullTransformationTests(t, tests)
}