package transformations

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func TestComplexCloudflareRulesetWithMultipleHeredocAndRatelimit(t *testing.T) {
	input := `resource "cloudflare_ruleset" "global-rate-limit" {
  account_id  = local.account_id
  kind        = "custom"
  name        = "Global rate limit"
  phase       = "http_ratelimit"
  description = "*These rules are managed by Terraform*"
  rules {
    action      = "log"
    description = "Log Global rate limit 1000/10s non authenticated"
    enabled     = true
    expression  = <<-EOF
      not (
        any(http.request.headers["authorization"][*] contains "sess-")
        or any(http.request.headers["authorization"][*] contains "sk-")
        or any(http.request.headers["authorization"][*] contains "eyJhbGc")
        or http.cookie contains "__Secure-next-auth."
        or http.cookie contains "__Host-next-auth.csrf-token="
        or http.cookie contains "oai-did="
      )
      and not http.user_agent eq "Stripe/1.0 (+https://stripe.com/docs/webhooks)"
    EOF
    ratelimit {
      characteristics     = ["cf.unique_visitor_id", "cf.colo.id"]
      mitigation_timeout  = 3600
      period              = 10
      requests_per_period = 1000
    }
  }
  rules {
    action      = "log"
    description = "Log Global rate limit 2500/1m non authenticated"
    enabled     = true
    expression  = <<-EOF
      not (
        any(http.request.headers["authorization"][*] contains "sess-")
        or any(http.request.headers["authorization"][*] contains "sk-")
        or any(http.request.headers["authorization"][*] contains "eyJhbGc")
        or http.cookie contains "__Secure-next-auth."
        or http.cookie contains "__Host-next-auth.csrf-token="
        or http.cookie contains "oai-did="
      )
      and not http.user_agent eq "Stripe/1.0 (+https://stripe.com/docs/webhooks)"
    EOF
    ratelimit {
      characteristics     = ["cf.unique_visitor_id", "cf.colo.id"]
      mitigation_timeout  = 3600
      period              = 60
      requests_per_period = 2500
    }
  }
}`

	// Test the transformation
	config := &TransformationConfig{
		Transformations: map[string]ResourceTransform{
			"cloudflare_ruleset": {
				ToList: []string{"rules"},
			},
		},
	}

	// Parse the input
	file, diags := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("Failed to parse input: %v", diags)
	}

	// Find the cloudflare_ruleset resource block and apply transformations
	blocks := file.Body().Blocks()
	var transformed *hclwrite.File
	for _, block := range blocks {
		if block.Type() == "resource" && len(block.Labels()) >= 2 && block.Labels()[0] == "cloudflare_ruleset" {
			err := TransformResourceBlock(config, block, "cloudflare_ruleset")
			if err != nil {
				t.Fatalf("Failed to transform resource block: %v", err)
			}
			transformed = file
			break
		}
	}
	if transformed == nil {
		t.Fatalf("Could not find cloudflare_ruleset resource block to transform")
	}

	// Format and apply post-processing
	output := hclwrite.Format(transformed.Bytes())
	output = fixCloudflareRulesetDoubleDollarSigns(output)

	// Verify the output is valid HCL
	_, diags = hclwrite.ParseConfig(output, "output.tf", hcl.InitialPos)
	if diags.HasErrors() {
		t.Errorf("Transformed output is not valid HCL: %v\nOutput:\n%s", diags, string(output))
	} else {
		t.Logf("✅ Transformation succeeded - output is valid HCL")
	}

	// Check that transformation actually happened (should contain "rules = [")
	if !strings.Contains(string(output), "rules = [") {
		t.Errorf("Expected transformation to convert rules blocks to list, but output doesn't contain 'rules = ['")
	} else {
		t.Logf("✅ Rules blocks correctly converted to list format")
	}

	// Check that heredoc expressions are preserved
	if !strings.Contains(string(output), "<<-EOF") {
		t.Errorf("Expected heredoc expressions to be preserved, but output doesn't contain '<<-EOF'")
	} else {
		t.Logf("✅ Heredoc expressions preserved correctly")
	}

	// Check for balanced EOF markers
	eofCount := strings.Count(string(output), "EOF")
	if eofCount%2 != 0 {
		t.Errorf("Expected balanced EOF markers (pairs), but found %d total EOF occurrences", eofCount)
	} else {
		t.Logf("✅ EOF markers are balanced (%d pairs)", eofCount/2)
	}

	t.Logf("Successfully transformed complex heredoc with ratelimit rules")
}

// TestExactMigrationTestCase tests the exact same configuration that fails in the migration tests
func TestExactMigrationTestCase(t *testing.T) {
	input := `resource "cloudflare_ruleset" "test" {
  account_id  = "test123"
  kind        = "zone"
  name        = "Complex Heredoc Test"
  phase       = "http_ratelimit"
  description = "Test complex heredoc with ratelimit"
  
  rules {
    action      = "log"
    description = "Log rate limit with complex expression"
    enabled     = true
    expression  = <<-EOF
      not (
        any(http.request.headers["authorization"][*] contains "sess-")
        or any(http.request.headers["authorization"][*] contains "sk-")
        or http.cookie contains "__Secure-next-auth."
      )
      and not http.user_agent eq "Stripe/1.0 (+https://stripe.com/docs/webhooks)"
    EOF
    ratelimit {
      characteristics     = ["cf.unique_visitor_id", "cf.colo.id"]
      mitigation_timeout  = 3600
      period              = 10
      requests_per_period = 1000
    }
  }
}`

	// Test the transformation
	config := &TransformationConfig{
		Transformations: map[string]ResourceTransform{
			"cloudflare_ruleset": {
				ToList: []string{"rules"},
			},
		},
	}

	// Parse the input
	file, diags := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("Failed to parse input: %v", diags)
	}

	// Find the cloudflare_ruleset resource block and apply transformations
	blocks := file.Body().Blocks()
	var transformed *hclwrite.File
	for _, block := range blocks {
		if block.Type() == "resource" && len(block.Labels()) >= 2 && block.Labels()[0] == "cloudflare_ruleset" {
			err := TransformResourceBlock(config, block, "cloudflare_ruleset")
			if err != nil {
				t.Fatalf("Failed to transform resource block: %v", err)
			}
			transformed = file
			break
		}
	}
	if transformed == nil {
		t.Fatalf("Could not find cloudflare_ruleset resource block to transform")
	}

	// Format and apply post-processing
	output := hclwrite.Format(transformed.Bytes())
	output = fixCloudflareRulesetDoubleDollarSigns(output)

	// Verify the output is valid HCL
	_, diags = hclwrite.ParseConfig(output, "output.tf", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("Transformed output is not valid HCL: %v\nOutput:\n%s", diags, string(output))
	}

	// Check that it has the expected structure
	if !strings.Contains(string(output), "rules = [") {
		t.Errorf("Expected rules to be converted to list format")
	}

	if !strings.Contains(string(output), "ratelimit = {") {
		t.Errorf("Expected ratelimit to be converted to object format")
	}

	// Debug: log the output before and after post-processing
	outputBeforeProcessing := hclwrite.Format(transformed.Bytes())
	t.Logf("Before post-processing (raw):\n%q", string(outputBeforeProcessing))
	t.Logf("Before post-processing (formatted):\n%s", string(outputBeforeProcessing))
	
	// Log the output for inspection
	t.Logf("After post-processing (raw):\n%q", string(output))
	t.Logf("After post-processing (formatted):\n%s", string(output))
}