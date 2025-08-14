package main

import (
	"testing"
)

func TestTransformAccessPolicyBasic(t *testing.T) {
	tests := []TestCase{
		{
			Name: "single_email_transform",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [{
    email = ["test@example.com"]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [
    { email = { email = "test@example.com" } }
  ]
}`},
		},
		{
			Name: "multiple_emails_transform",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [{
    email = ["user1@example.com", "user2@example.com"]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [
    { email = { email = "user1@example.com" } },
    { email = { email = "user2@example.com" } }
  ]
}`},
		},
		{
			Name: "mixed_condition_types_transform",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [{
    email = ["user@example.com"]
    email_domain = ["example.com"]
    everyone = true
    any_valid_service_token = true
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [
    { everyone = {} },
    { any_valid_service_token = {} },
    { email = { email = "user@example.com" } },
    { email_domain = { domain = "example.com" } }
  ]
}`},
		},
		{
			Name: "ip_range_transform",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [{
    ip = ["192.168.1.0/24", "10.0.0.0/8"]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [
    { ip = { ip = "192.168.1.0/24" } },
    { ip = { ip = "10.0.0.0/8" } }
  ]
}`},
		},
		{
			Name: "exclude_conditions_transform",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  exclude = [{
    email = ["blocked@example.com"]
    geo = ["CN", "RU"]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  exclude = [
    { email = { email = "blocked@example.com" } },
    { geo = { country_code = "CN" } },
    { geo = { country_code = "RU" } }
  ]
}`},
		},
		{
			Name: "require_certificate_transform",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  require = [{
    certificate = true
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  require = [
    { certificate = {} }
  ]
}`},
		},
		{
			Name: "boolean_false_removal",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [{
    email = ["user@example.com"]
    everyone = false
    certificate = false
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [
    { email = { email = "user@example.com" } }
  ]
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}