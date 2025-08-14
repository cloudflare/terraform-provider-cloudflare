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
							    everyone = {}
							    login_method = ["okta"]
							  }, {
							    email = { email = "user@example.com" }
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
							    certificate = {}
							    login_method = ["github"]
							  }, {
							    email = { email = "alice@example.com" }
							  }, {
							    group = { id = "admins" }
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
							  precedence = 1

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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
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
			output := string(hclwrite.Format(file.Bytes()))
			
			// Print diagnostics for debugging
			if ds.HclDiagnostics.HasErrors() || len(ds.HclDiagnostics) > 0 {
				for _, d := range ds.HclDiagnostics {
					t.Logf("Diagnostic: %s", d.Summary)
				}
			}

			// Normalize whitespace for comparison
			expected := normalizeWhitespace(tt.expected)
			actual := normalizeWhitespace(output)

			if expected != actual {
				t.Errorf("Transformation mismatch.\nExpected:\n%s\n\nGot:\n%s\n\nRaw output:\n%s", expected, actual, output)
			}
		})
	}
}
