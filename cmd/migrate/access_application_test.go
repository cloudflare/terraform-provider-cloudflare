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
  
  policies = [ { id = cloudflare_zero_trust_access_policy.old_style.id } ]
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

func TestAccessApplicationDomainTypeRemoval(t *testing.T) {
	tests := []TestCase{
		{
			Name: "remove domain_type attribute",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id  = "abc123"
  name        = "Test App"
  domain      = "test.example.com"
  domain_type = "public"
  type        = "self_hosted"
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  type       = "self_hosted"
}`},
		},
		{
			Name: "remove domain_type with other attributes preserved",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id    = "abc123"
  name          = "Test App"
  domain        = "test.example.com"
  domain_type   = "public"
  type          = "self_hosted"
  session_duration = "24h"
  
  cors_headers {
    allow_all_origins = true
  }
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id       = "abc123"
  name             = "Test App"
  domain           = "test.example.com"
  type             = "self_hosted"
  session_duration = "24h"
  
  cors_headers {
    allow_all_origins = true
  }
}`},
		},
		{
			Name: "no domain_type to remove",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  type       = "self_hosted"
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  type       = "self_hosted"
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}

func TestAccessApplicationDestinationsBlocksToAttribute(t *testing.T) {
	tests := []TestCase{
		{
			Name: "convert single destinations block to list attribute",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations {
    uri = "https://example.com"
  }
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations = [
    {
      uri = "https://example.com"
    }
  ]
}`},
		},
		{
			Name: "convert multiple destinations blocks to list attribute",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations {
    uri = "https://example.com"
  }
  
  destinations {
    uri = "tcp://db.example.com:5432"
  }
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations = [
    {
      uri = "https://example.com"
    },
    {
      uri = "tcp://db.example.com:5432"
    }
  ]
}`},
		},
		{
			Name: "destinations block with multiple attributes",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations {
    uri         = "https://app.example.com"
    description = "Main application"
  }
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations = [
    {
      description = "Main application"
      uri         = "https://app.example.com"
    }
  ]
}`},
		},
		{
			Name: "no destinations blocks - no change",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "self_hosted"
  domain     = "test.example.com"
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "self_hosted"
  domain     = "test.example.com"
}`},
		},
		{
			Name: "destinations blocks with variable references",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations {
    uri = var.app_uri
  }
  
  destinations {
    uri = local.db_connection
  }
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations = [
    {
      uri = var.app_uri
    },
    {
      uri = local.db_connection
    }
  ]
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}

func TestAccessApplicationCombinedMigrations(t *testing.T) {
	tests := []TestCase{
		{
			Name: "combined domain_type removal and destinations conversion",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id  = "abc123"
  name        = "Test App"
  type        = "warp"
  domain_type = "public"
  
  destinations {
    uri = "https://example.com"
  }
  
  policies = ["policy-id-1", "policy-id-2"]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  policies = [{ id = "policy-id-1" }, { id = "policy-id-2" }]
  destinations = [
    {
      uri = "https://example.com"
    }
  ]
}`},
		},
		{
			Name: "all transformations together with allowed_idps",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id    = "abc123"
  name          = "Test App"
  type          = "warp"
  domain_type   = "public"
  allowed_idps  = toset(["idp-1", "idp-2"])
  
  destinations {
    uri = "https://example.com"
  }
  
  destinations {
    uri = "tcp://db.example.com:5432"
  }
  
  policies = [
    cloudflare_zero_trust_access_policy.allow.id,
    "literal-policy-id"
  ]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id   = "abc123"
  name         = "Test App"
  type         = "warp"
  allowed_idps = ["idp-1", "idp-2"]
  
  policies = [{ id = cloudflare_zero_trust_access_policy.allow.id }, { id = "literal-policy-id" }]
  destinations = [
    {
      uri = "https://example.com"
    },
    {
      uri = "tcp://db.example.com:5432"
    }
  ]
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}
