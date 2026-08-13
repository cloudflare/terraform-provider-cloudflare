package main

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
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

  policies = [{ id = cloudflare_zero_trust_access_policy.allow.id }, { id = cloudflare_zero_trust_access_policy.deny.id }]
  type     = "self_hosted"
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
  type     = "self_hosted"
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

  policies = [{ id = cloudflare_zero_trust_access_policy.allow.id }, { id = "literal-policy-id" }, { id = cloudflare_zero_trust_access_policy.deny.id }]
  type     = "self_hosted"
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

  policies = [{ id = cloudflare_zero_trust_access_policy.old_style.id }]
  type     = "self_hosted"
}`},
		},
		{
			Name: "no policies attribute but add default type",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  type       = "self_hosted"
}`},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
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

	RunTransformationTests(t, tests, transformFileDefault)
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

	RunTransformationTests(t, tests, transformFileDefault)
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

	RunTransformationTests(t, tests, transformFileDefault)
}


func TestAccessApplicationSkipAppLauncherLoginPageRemoval(t *testing.T) {
	tests := []TestCase{
		{
			Name: "remove skip_app_launcher_login_page when type is not app_launcher",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id                  = "abc123"
  name                        = "Test App"
  domain                      = "test.example.com"
  type                        = "self_hosted"
  skip_app_launcher_login_page = false
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  type       = "self_hosted"
}`},
		},
		{
			Name: "preserve skip_app_launcher_login_page when type is app_launcher",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id                  = "abc123"
  name                        = "Test App"
  type                        = "app_launcher"
  skip_app_launcher_login_page = true
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id                   = "abc123"
  name                         = "Test App"
  type                         = "app_launcher"
  skip_app_launcher_login_page = true
}`},
		},
		{
			Name: "remove skip_app_launcher_login_page when type is warp",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id                  = "abc123"
  name                        = "Test App"
  type                        = "warp"
  skip_app_launcher_login_page = false
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
}`},
		},
		{
			Name: "remove skip_app_launcher_login_page when no type attribute and add default type",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id                  = "abc123"
  name                        = "Test App"
  domain                      = "test.example.com"
  skip_app_launcher_login_page = false
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  type       = "self_hosted"
}`},
		},
		{
			Name: "no skip_app_launcher_login_page to remove",
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
	}

	RunTransformationTests(t, tests, transformFileDefault)
}

func TestAccessApplicationSetToListTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "transform toset to list for allowed_idps",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id   = "abc123"
  name         = "Test App"
  domain       = "test.example.com"
  allowed_idps = toset(["idp-1", "idp-2", "idp-3"])
  type         = "self_hosted"
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id   = "abc123"
  name         = "Test App"
  domain       = "test.example.com"
  allowed_idps = ["idp-1", "idp-2", "idp-3"]
  type         = "self_hosted"
}`},
		},
		{
			Name: "handle already list format for allowed_idps",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id   = "abc123"
  name         = "Test App"
  domain       = "test.example.com"
  allowed_idps = ["idp-1", "idp-2"]
  type         = "self_hosted"
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id   = "abc123"
  name         = "Test App"
  domain       = "test.example.com"
  allowed_idps = ["idp-1", "idp-2"]
  type         = "self_hosted"
}`},
		},
		{
			Name: "transform toset for custom_pages",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id    = "abc123"
  name          = "Test App"
  domain        = "test.example.com"
  custom_pages  = toset(["page1", "page2"])
  type          = "self_hosted"
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id   = "abc123"
  name         = "Test App"
  domain       = "test.example.com"
  custom_pages = ["page1", "page2"]
  type         = "self_hosted"
}`},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
}

func TestAccessApplicationPoliciesEdgeCases(t *testing.T) {
	tests := []TestCase{
		{
			Name: "empty policies array",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  policies   = []
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  policies   = []
  type       = "self_hosted"
}`},
		},
		{
			Name: "complex policy references with expressions",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  policies   = concat(
    [cloudflare_zero_trust_access_policy.main.id],
    var.additional_policies
  )
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  policies   = concat([cloudflare_zero_trust_access_policy.main.id], var.additional_policies)
  type       = "self_hosted"
}`},
		},
		{
			Name: "policies with for expression",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  policies   = [for p in var.policy_ids : p]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  domain     = "test.example.com"
  policies = [
    for p in
    var.policy_ids
    : p
  ]
  type = "self_hosted"
}`},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
}

func TestAccessApplicationDestinationsEdgeCases(t *testing.T) {
	tests := []TestCase{
		{
			Name: "destinations with expressions",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations {
    uri = format("https://%s.example.com", var.subdomain)
  }
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations = [
    {
      uri = format("https://%s.example.com", var.subdomain)
    }
  ]
}`},
		},
		{
			Name: "destinations with conditional expression",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations {
    uri = var.use_ssl ? "https://app.example.com" : "http://app.example.com"
  }
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations = [
    {
      uri = var.use_ssl ? "https://app.example.com" : "http://app.example.com"
    }
  ]
}`},
		},
		{
			Name: "destinations block without uri",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations {
    description = "Test destination"
  }
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations = [
    {
      description = "Test destination"
    }
  ]
}`},
		},
		{
			Name: "empty destinations block",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations {
  }
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations = [
    {}
  ]
}`},
		},
		{
			Name: "multiple destinations with mixed content",
			Config: `resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations {
    uri = "https://app1.example.com"
    description = "Primary app"
  }
  
  destinations {
  }
  
  destinations {
    uri = "tcp://db.example.com:3306"
  }
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "test" {
  account_id = "abc123"
  name       = "Test App"
  type       = "warp"
  
  destinations = [
    {
      description = "Primary app"
      uri         = "https://app1.example.com"
    },
    {},
    {
      uri = "tcp://db.example.com:3306"
    }
  ]
}`},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
}

func TestCreatePoliciesAttribute(t *testing.T) {
	tests := []struct {
		name     string
		policies []PolicyReference
		expected string
	}{
		{
			name:     "no policies",
			policies: []PolicyReference{},
			expected: "",
		},
		{
			name: "single policy",
			policies: []PolicyReference{
				{ResourceName: "cloudflare_zero_trust_access_policy.test1", Precedence: 1},
			},
			expected: `policies = [
  {
    id         = cloudflare_zero_trust_access_policy.test1.id
    precedence = 1
  }
]`,
		},
		{
			name: "multiple policies",
			policies: []PolicyReference{
				{ResourceName: "cloudflare_zero_trust_access_policy.test1", Precedence: 1},
				{ResourceName: "cloudflare_zero_trust_access_policy.test2", Precedence: 2},
				{ResourceName: "cloudflare_zero_trust_access_policy.test3", Precedence: 3},
			},
			expected: `policies = [
  {
    id         = cloudflare_zero_trust_access_policy.test1.id
    precedence = 1
  },
  {
    id         = cloudflare_zero_trust_access_policy.test2.id
    precedence = 2
  },
  {
    id         = cloudflare_zero_trust_access_policy.test3.id
    precedence = 3
  }
]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := hclwrite.NewEmptyFile()
			body := file.Body()
			
			createPoliciesAttribute(body, tt.policies)
			
			result := string(file.Bytes())
			if tt.expected == "" {
				assert.Equal(t, "", strings.TrimSpace(result))
			} else {
				// Check that the expected content is in the result
				assert.Contains(t, result, tt.expected)
				if len(tt.policies) > 0 {
					assert.Contains(t, result, "# Policies auto-migrated from v4 access_policy resources")
				}
			}
		})
	}
}
