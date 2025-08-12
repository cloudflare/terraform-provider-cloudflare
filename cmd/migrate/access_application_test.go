package main

import (
	"testing"
)

func TestAccessApplicationPoliciesTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "transform policies from list of strings to list of objects",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = [
    cloudflare_zero_trust_access_policy.allow.id,
    cloudflare_zero_trust_access_policy.deny.id
  ]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = [
    { id = cloudflare_zero_trust_access_policy.allow.id },
    { id = cloudflare_zero_trust_access_policy.deny.id }
  ]
}`},
		},
		{
			Name: "transform policies with literal IDs",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = ["policy-id-1", "policy-id-2"]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = [{ id = "policy-id-1" }, { id = "policy-id-2" }]
}`},
		},
		{
			Name: "mixed references and literals",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = [
    cloudflare_zero_trust_access_policy.allow.id,
    "literal-policy-id",
    cloudflare_zero_trust_access_policy.deny.id
  ]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = [
    { id = cloudflare_zero_trust_access_policy.allow.id },
    { id = "literal-policy-id" },
    { id = cloudflare_zero_trust_access_policy.deny.id }
  ]
}`},
		},
		{
			Name: "handle old resource name references",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = [
    cloudflare_access_policy.old_style.id
  ]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  
  policies = [
    { id = cloudflare_zero_trust_access_policy.old_style.id }
  ]
}`},
		},
		{
			Name: "no policies attribute",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}