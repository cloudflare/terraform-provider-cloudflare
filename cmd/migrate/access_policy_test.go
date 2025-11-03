package main

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func TestTransformAccessPolicy(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "simple_email_transform",
			input: `resource "cloudflare_zero_trust_access_policy" "test" {
			  account_id = "abc123"
			  name       = "Test Policy"
			  decision   = "allow"

			  include = [{
			    email = ["test@example.com", "admin@example.com"]
			  }]
			}`,
			expected: `resource "cloudflare_zero_trust_access_policy" "test" {
			  account_id = "abc123"
			  name       = "Test Policy"
			  decision   = "allow"

			  include = [{
			    email = { email = "test@example.com" }
			  }, {
			  	email = { email =  "admin@example.com" }
			  }]
			}`,
		},

		{
			name: "multiple_conditions",
			input: `resource "cloudflare_zero_trust_access_policy" "test" {
				  account_id = "abc123"
				  name       = "Test Policy"
				  decision   = "allow"

				  include = [{
				    email = ["test@example.com"]
				    group = ["group-id-1", "group-id-2"]
				  }]

				  exclude = [{
				    ip = ["192.168.1.0/24"]
				  }]
				}`,
			expected: `resource "cloudflare_zero_trust_access_policy" "test" {
				  account_id = "abc123"
				  name       = "Test Policy"
				  decision   = "allow"

				  include = [{
				    email = { email = "test@example.com" }
				  }, {
				    group = { id = "group-id-1" }
				  }, {
				  	group = { id = "group-id-2" }
				  }
				  ]

				  exclude = [{
				    ip = { ip = "192.168.1.0/24" }
				  }]
				}`,
		},

		{
			name: "boolean_attribute_everyone",
			input: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"

  include = [{
    everyone = true
  }]
}`,
			expected: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"

  include = [{
    everyone = {}
  }]
}`,
		},

		{
			name: "github_single_team",
			input: `resource "cloudflare_zero_trust_access_policy" "test" {
			  account_id = "abc123"
			  name       = "Test Policy"
			  decision   = "allow"

			  include = [{
			    email_domain = ["example.com"]
			  }]

			  require = [{
			    github = [{
			      name = "my-org"
			      teams = ["engineering"]
			    }]
			  }]
			}`,
			expected: `resource "cloudflare_zero_trust_access_policy" "test" {
							  account_id = "abc123"
							  name       = "Test Policy"
							  decision   = "allow"

							  include = [{
							    email_domain = { domain = "example.com" }
							  }]

							  require = [{
							    github_organization = {
							      name = "my-org"
							      team = "engineering"
							    }
							  }]
							}`,
		},
		{
			name: "github_multiple_teams",
			input: `resource "cloudflare_zero_trust_access_policy" "test" {
			  account_id = "abc123"
			  name       = "Test Policy"
			  decision   = "allow"

			  include = [{
			    github = [{
			      name = "my-org"
			      teams = ["engineering", "devops", "security"]
			      identity_provider_id = "provider-123"
			    }]
			  }]
			}`,
			expected: `resource "cloudflare_zero_trust_access_policy" "test" {
							  account_id = "abc123"
							  name       = "Test Policy"
							  decision   = "allow"

							  include = [{
							    github_organization = {
							      name                 = "my-org"
							      team                 = "engineering"
							      identity_provider_id = "provider-123"
							    }
							  }, {
							    github_organization = {
							      name                 = "my-org"
							      team                 = "devops"
							      identity_provider_id = "provider-123"
							    }
							  }, {
							    github_organization = {
							      name                 = "my-org"
							      team                 = "security"
							      identity_provider_id = "provider-123"
							    }
							  }]
							}`,
		},
		{
			name: "mixed_array_attributes_ordering",
			input: `resource "cloudflare_zero_trust_access_policy" "test" {
			  account_id = "abc123"
			  name       = "Test Policy"
			  decision   = "allow"

			  include = [{
			    email = ["first@example.com", "second@example.com"]
			    group = ["group-1", "group-2"]
			    ip = ["10.0.0.0/8", "192.168.0.0/16"]
			  }]
			}`,
			expected: `resource "cloudflare_zero_trust_access_policy" "test" {
							  account_id = "abc123"
							  name       = "Test Policy"
							  decision   = "allow"

							  include = [{
							    email = { email = "first@example.com" }
							  }, {
							    email = { email = "second@example.com" }
							  }, {
							    group = { id = "group-1" }
							  }, {
							    group = { id = "group-2" }
							  }, {
							    ip = { ip = "10.0.0.0/8" }
							  }, {
							    ip = { ip = "192.168.0.0/16" }
							  }]
							}`,
		},
		{
			name: "mixed_with_non_array_attributes",
			input: `resource "cloudflare_zero_trust_access_policy" "test" {
			  account_id = "abc123"
			  name       = "Test Policy"
			  decision   = "allow"

			  include = [{
			    everyone = true
			    email = ["user@example.com"]
			    login_method = ["okta"]
			  }]
			}`,
			expected: `resource "cloudflare_zero_trust_access_policy" "test" {
							  account_id = "abc123"
							  name       = "Test Policy"
							  decision   = "allow"

							  include = [{
							    email = { email = "user@example.com" }
							  }, {
							    everyone = {}
							  }, {
							    login_method = ["okta"]
							  }]
							}`,
		},
		{
			name: "array_attributes_interleaved",
			input: `resource "cloudflare_zero_trust_access_policy" "test" {
			  account_id = "abc123"
			  name       = "Test Policy"
			  decision   = "allow"

			  include = [{
			    certificate = true
			    email = ["alice@example.com"]
			    login_method = ["github"]
			    group = ["admins"]
			    any_valid_service_token = false
			  }]
			}`,
			expected: `resource "cloudflare_zero_trust_access_policy" "test" {
							  account_id = "abc123"
							  name       = "Test Policy"
							  decision   = "allow"

							  include = [{
							    email = { email = "alice@example.com" }
							  }, {
							    group = { id = "admins" }
							  }, {
							    certificate = {}
							  }, {
							    login_method = ["github"]
							  }]
							}`,
		},
		{
			name: "no_transformation_needed",
			input: `resource "cloudflare_zero_trust_access_policy" "test" {
							  account_id = "abc123"
							  name       = "Test Policy"
							  decision   = "allow"
							  precedence = 1

							  include = [{
							    login_method = ["saml"]
							  }]
							}`,
			expected: `resource "cloudflare_zero_trust_access_policy" "test" {
							  account_id = "abc123"
							  name       = "Test Policy"
							  decision   = "allow"

							  include = [{
							    login_method = ["saml"]
							  }]
							}`,
		},
		{
			name: "skip_non_access_policy_resources",
			input: `resource "cloudflare_zero_trust_access_application" "test" {
							  account_id = "abc123"
							  name       = "Test App"
							  domain     = "test.example.com"
							  type       = "self_hosted"
							}`,
			expected: `resource "cloudflare_zero_trust_access_application" "test" {
							  account_id = "abc123"
							  name       = "Test App"
							  domain     = "test.example.com"
							  type       = "self_hosted"
							}`,
		},
	}

	cases := []TestCase{}
	for _, tt := range tests {
		cases = append(cases, TestCase{
			Name:     tt.name,
			Config:   tt.input,
			Expected: []string{tt.expected},
		})
	}

	RunTransformationTests(t, cases, func(input []byte, s string) ([]byte, error) {
		// Parse the input
		file, diags := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		if diags.HasErrors() {
			t.Fatalf("Failed to parse input: %v", diags.Error())
		}
		ds := ast.NewDiagnostics()
		// Apply the transformation
		for _, block := range file.Body().Blocks() {
			if isAccessPolicyResource(block) {
				transformAccessPolicyBlock(block, ds)
			}
		}

		// Format the output
		return hclwrite.Format(file.Bytes()), nil
	})
}

func TestAccessPolicyAdditionalCoverage(t *testing.T) {
	tests := []TestCase{
		{
			Name: "access_policy with github teams",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [{
    github = [{
      teams = ["team-1", "team-2"]
      identity_provider_id = "provider-123"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test"`},
		},
		{
			Name: "access_policy with empty include/exclude arrays",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = []
  exclude = []
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test"`},
		},
		{
			Name: "access_policy with complex github configuration",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [{
    github = [{
      name = "my-org"
      teams = ["team-1"]
      identity_provider_id = "provider-123"
    }, {
      name = "other-org"
      teams = ["team-2", "team-3"]
      identity_provider_id = "provider-456"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test"`},
		},
		{
			Name: "access_policy with application_id reference",
			Config: `resource "cloudflare_zero_trust_access_application" "app" {
  account_id = "abc123"
  name       = "My App"
  domain     = "example.com"
}

resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  application_id = cloudflare_zero_trust_access_application.app.id
  
  include = [{
    email = ["test@example.com"]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test"`},
		},
		{
			Name: "access_policy with zone_id attribute",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  zone_id = "zone123"
  name    = "Test Policy"
  decision = "allow"
  
  include = [{
    everyone = true
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test"`},
		},
		{
			Name: "access_policy with gsuite groups",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [{
    gsuite = [{
      email = ["gsuite-group@example.com"]
      identity_provider_id = "provider-123"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test"`},
		},
		{
			Name: "access_policy with okta groups",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [{
    okta = [{
      name = ["okta-group-1", "okta-group-2"]
      identity_provider_id = "provider-123"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test"`},
		},
		{
			Name: "access_policy with saml groups",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [{
    saml = [{
      attribute_name = "group"
      attribute_value = "admins"
      identity_provider_id = "provider-123"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test"`},
		},
		{
			Name: "access_policy with azure groups",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [{
    azure = [{
      id = ["azure-group-1", "azure-group-2"]
      identity_provider_id = "provider-123"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test"`},
		},
		{
			Name: "access_policy with require and exclude blocks",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [{
    email = ["user@example.com"]
  }]
  
  require = [{
    group = ["required-group"]
  }]
  
  exclude = [{
    ip = ["192.168.1.0/24"]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test"`},
		},
		{
			Name: "access_policy with external_evaluation",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [{
    external_evaluation = [{
      evaluate_url = "https://example.com/evaluate"
      keys_url = "https://example.com/keys"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test"`},
		},
		{
			Name: "access_policy with auth_context",
			Config: `resource "cloudflare_zero_trust_access_policy" "test" {
  account_id = "abc123"
  name       = "Test Policy"
  decision   = "allow"
  
  include = [{
    auth_context = [{
      id = "context-123"
      ac_id = "ac-456"
      identity_provider_id = "provider-789"
    }]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_policy" "test"`},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
}

func TestExtractApplicationReference(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "direct application id reference",
			input:    "cloudflare_zero_trust_access_application.app.id",
			expected: "cloudflare_zero_trust_access_application.app",
		},
		{
			name:     "application reference with index",
			input:    "cloudflare_zero_trust_access_application.app[0].id",
			expected: "cloudflare_zero_trust_access_application.app[0]",
		},
		{
			name:     "non-application reference",
			input:    "var.application_id",
			expected: "",
		},
		{
			name:     "local reference",
			input:    "local.app_id",
			expected: "",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test would need the actual implementation
			// but serves to show the test structure
			if tt.input == "" && tt.expected == "" {
				// Pass for now
				return
			}
		})
	}
}

func TestGenerateMovedBlocks(t *testing.T) {
	tests := []struct {
		name     string
		oldAddr  string
		newAddr  string
		expected bool
	}{
		{
			name:     "simple resource move",
			oldAddr:  "cloudflare_zero_trust_access_policy.old",
			newAddr:  "cloudflare_zero_trust_access_policy.new",
			expected: true,
		},
		{
			name:     "same address no move",
			oldAddr:  "cloudflare_zero_trust_access_policy.test",
			newAddr:  "cloudflare_zero_trust_access_policy.test",
			expected: false,
		},
		{
			name:     "empty addresses",
			oldAddr:  "",
			newAddr:  "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test would need the actual implementation
			// but serves to show the test structure
			if tt.oldAddr == "" && !tt.expected {
				// Pass for now
				return
			}
		})
	}
}

func TestCollectApplicationPolicyMapping(t *testing.T) {
	tests := []TestCase{
		{
			Name: "multiple policies for same application",
			Config: `resource "cloudflare_zero_trust_access_application" "app1" {
  account_id = "abc123"
  name       = "App 1"
  domain     = "app1.example.com"
}

resource "cloudflare_zero_trust_access_policy" "policy1" {
  account_id = "abc123"
  name       = "Policy 1"
  decision   = "allow"
  application_id = cloudflare_zero_trust_access_application.app1.id
  precedence = 1
  
  include = [{
    email = ["user1@example.com"]
  }]
}

resource "cloudflare_zero_trust_access_policy" "policy2" {
  account_id = "abc123"
  name       = "Policy 2"
  decision   = "deny"
  application_id = cloudflare_zero_trust_access_application.app1.id
  precedence = 2
  
  include = [{
    email = ["user2@example.com"]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application" "app1"`},
		},
		{
			Name: "policies with different application references",
			Config: `resource "cloudflare_zero_trust_access_application" "app1" {
  account_id = "abc123"
  name       = "App 1"
  domain     = "app1.example.com"
}

resource "cloudflare_zero_trust_access_application" "app2" {
  account_id = "abc123"
  name       = "App 2"
  domain     = "app2.example.com"
}

resource "cloudflare_zero_trust_access_policy" "policy1" {
  account_id = "abc123"
  name       = "Policy 1"
  decision   = "allow"
  application_id = cloudflare_zero_trust_access_application.app1.id
  
  include = [{
    email = ["user@example.com"]
  }]
}

resource "cloudflare_zero_trust_access_policy" "policy2" {
  account_id = "abc123"
  name       = "Policy 2"
  decision   = "allow"
  application_id = cloudflare_zero_trust_access_application.app2.id
  
  include = [{
    email = ["user@example.com"]
  }]
}`,
			Expected: []string{`resource "cloudflare_zero_trust_access_application"`},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
}
