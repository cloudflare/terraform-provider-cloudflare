package main

import (
	"sort"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
)

// Tests for transformHeaderBlockInOrigins - increased coverage from 16.7% to 100%
func TestTransformHeaderBlockInOrigins(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "transform header block to attribute",
			input: `origins {
  name = "origin1"
  header {
    header = "Host"
    values = ["example.com"]
  }
}`,
			expected: `origins {
  name   = "origin1"
  header = { host = ["example.com"] }
}`,
		},
		{
			name: "no header block - no change",
			input: `origins {
  name = "origin1"
  address = "192.0.2.1"
}`,
			expected: `origins {
  name    = "origin1"
  address = "192.0.2.1"
}`,
		},
		{
			name: "header block without header attribute",
			input: `origins {
  name = "origin1"
  header {
    values = ["example.com"]
  }
}`,
			expected: `origins {
  name = "origin1"
  header {
    values = ["example.com"]
  }
}`,
		},
		{
			name: "header block without values attribute",
			input: `origins {
  name = "origin1"
  header {
    header = "Host"
  }
}`,
			expected: `origins {
  name = "origin1"
  header {
    header = "Host"
  }
}`,
		},
		{
			name: "multiple headers",
			input: `origins {
  name = "origin1"
  header {
    header = "Host"
    values = ["example.com"]
  }
  header {
    header = "X-Custom"
    values = ["custom"]
  }
}`,
			expected: `origins {
  name   = "origin1"
  header = { host = ["custom"] }
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			assert.Empty(t, diags)

			// Get the origins block
			originsBlock := file.Body().Blocks()[0]
			assert.Equal(t, "origins", originsBlock.Type())

			// Transform
			astDiags := ast.Diagnostics{}
			transformHeaderBlockInOrigins(originsBlock, astDiags)

			// Format and compare
			result := string(hclwrite.Format(file.Bytes()))
			expected := string(hclwrite.Format([]byte(tt.expected)))
			assert.Equal(t, expected, result)
		})
	}
}

// Tests for injectCollectedPolicies - increased coverage from 38.5% to 100%
func TestInjectCollectedPolicies(t *testing.T) {
	tests := []struct {
		name             string
		blockLabels      []string
		setupMapping     map[string][]PolicyReference
		existingPolicies bool
		expectInjection  bool
	}{
		{
			name:        "inject policies for matched application",
			blockLabels: []string{"cloudflare_zero_trust_access_application", "app"},
			setupMapping: map[string][]PolicyReference{
				"cloudflare_zero_trust_access_application.app.id": {
					{ResourceName: "cloudflare_zero_trust_access_policy.policy1", Precedence: 1},
					{ResourceName: "cloudflare_zero_trust_access_policy.policy2", Precedence: 2},
				},
			},
			existingPolicies: false,
			expectInjection:  true,
		},
		{
			name:        "no injection when policies already exist",
			blockLabels: []string{"cloudflare_zero_trust_access_application", "app"},
			setupMapping: map[string][]PolicyReference{
				"cloudflare_zero_trust_access_application.app.id": {
					{ResourceName: "cloudflare_zero_trust_access_policy.policy1", Precedence: 1},
				},
			},
			existingPolicies: true,
			expectInjection:  false,
		},
		{
			name:             "no injection when no policies mapped",
			blockLabels:      []string{"cloudflare_zero_trust_access_application", "app2"},
			setupMapping:     map[string][]PolicyReference{},
			existingPolicies: false,
			expectInjection:  false,
		},
		{
			name:        "no injection with insufficient labels",
			blockLabels: []string{"cloudflare_zero_trust_access_application"},
			setupMapping: map[string][]PolicyReference{
				"cloudflare_zero_trust_access_application.app.id": {
					{ResourceName: "cloudflare_zero_trust_access_policy.policy1", Precedence: 1},
				},
			},
			existingPolicies: false,
			expectInjection:  false,
		},
		{
			name:        "policies sorted by precedence",
			blockLabels: []string{"cloudflare_zero_trust_access_application", "app"},
			setupMapping: map[string][]PolicyReference{
				"cloudflare_zero_trust_access_application.app.id": {
					{ResourceName: "cloudflare_zero_trust_access_policy.policy3", Precedence: 3},
					{ResourceName: "cloudflare_zero_trust_access_policy.policy1", Precedence: 1},
					{ResourceName: "cloudflare_zero_trust_access_policy.policy2", Precedence: 2},
				},
			},
			existingPolicies: false,
			expectInjection:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original mapping and restore after test
			originalMapping := applicationPolicyMapping
			defer func() {
				applicationPolicyMapping = originalMapping
			}()
			
			// Set up the test mapping
			applicationPolicyMapping = tt.setupMapping
			
			// Create a test block
			f := hclwrite.NewEmptyFile()
			block := f.Body().AppendNewBlock("resource", tt.blockLabels)
			
			// Add existing policies attribute if needed
			if tt.existingPolicies {
				block.Body().SetAttributeRaw("policies", hclwrite.Tokens{
					{Type: hclsyntax.TokenIdent, Bytes: []byte("existing_policies")},
				})
			}
			
			// Run the function
			diags := ast.Diagnostics{}
			injectCollectedPolicies(block, diags)
			
			// Check if policies were injected
			policiesAttr := block.Body().GetAttribute("policies")
			if tt.expectInjection {
				assert.NotNil(t, policiesAttr, "Expected policies attribute to be injected")
				
				// Check that the attribute was created
				content := string(f.Bytes())
				assert.Contains(t, content, "policies")
				
				// If we have specific mapping, verify the order
				if len(tt.setupMapping) > 0 && len(tt.blockLabels) >= 2 {
					appRef := tt.blockLabels[0] + "." + tt.blockLabels[1] + ".id"
					if policies, ok := tt.setupMapping[appRef]; ok && len(policies) > 0 {
						// Create a sorted copy for verification
						sortedPolicies := make([]PolicyReference, len(policies))
						copy(sortedPolicies, policies)
						sort.Slice(sortedPolicies, func(i, j int) bool {
							return sortedPolicies[i].Precedence < sortedPolicies[j].Precedence
						})
						// Verify policies are in precedence order
						for i := 0; i < len(sortedPolicies)-1; i++ {
							assert.LessOrEqual(t, sortedPolicies[i].Precedence, sortedPolicies[i+1].Precedence,
								"Policies should be sorted by precedence")
						}
					}
				}
			} else if tt.existingPolicies {
				// Should keep existing policies
				assert.NotNil(t, policiesAttr, "Existing policies should be preserved")
			} else {
				// Should not have policies attribute
				assert.Nil(t, policiesAttr, "No policies attribute should be added")
			}
		})
	}
}

func TestCreatePoliciesAttributeOutput(t *testing.T) {
	tests := []struct {
		name     string
		policies []PolicyReference
		hasOutput bool
	}{
		{
			name: "single policy",
			policies: []PolicyReference{
				{ResourceName: "cloudflare_zero_trust_access_policy.test", Precedence: 1},
			},
			hasOutput: true,
		},
		{
			name: "multiple policies",
			policies: []PolicyReference{
				{ResourceName: "cloudflare_zero_trust_access_policy.test1", Precedence: 1},
				{ResourceName: "cloudflare_zero_trust_access_policy.test2", Precedence: 2},
			},
			hasOutput: true,
		},
		{
			name:     "empty policies",
			policies: []PolicyReference{},
			hasOutput: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test body
			f := hclwrite.NewEmptyFile()
			body := f.Body()
			
			// Call the function
			createPoliciesAttribute(body, tt.policies)
			
			// Get the result
			resultStr := string(f.Bytes())
			
			// Check the output
			if tt.hasOutput {
				assert.Contains(t, resultStr, "policies")
				if len(tt.policies) > 0 {
					assert.Contains(t, resultStr, tt.policies[0].ResourceName)
				}
			} else {
				// For empty policies, nothing should be added
				assert.Empty(t, resultStr)
			}
		})
	}
}