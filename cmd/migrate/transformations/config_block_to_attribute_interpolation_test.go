package transformations

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// TestVariableInterpolationPreservation tests that variable interpolations are preserved
// during block-to-list transformations across various resource types
func TestVariableInterpolationPreservation(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "interpolation_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration file
	configPath := filepath.Join(tempDir, "test_config.yaml")
	configContent := `
version: "1.0"
description: "Test variable interpolation preservation"
transformations:
  cloudflare_load_balancer_pool:
    to_map:
      - load_shedding
      - origin_steering
    to_list:
      - origins
  cloudflare_load_balancer:
    to_list:
      - country_pools
      - pop_pools
      - region_pools
      - rules
  cloudflare_load_balancer_monitor:
    to_list:
      - header
  cloudflare_teams_location:
    to_list:
      - networks
  cloudflare_access_mutual_tls_hostname_settings:
    to_list:
      - settings
  cloudflare_teams_account:
    to_list:
      - antivirus
      - notification_settings
      - block_page
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Test cases
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "preserve simple variable references in origins",
			input: `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"

  origins {
    address = var.origin_address
    name    = var.origin_name
    enabled = true
  }
}`,
			expected: `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"

  origins = [
    {
      address = var.origin_address,
      enabled = true,
      name    = var.origin_name
    }
  ]
}`,
		},
		{
			name: "preserve string interpolations",
			input: `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"

  origins {
    address = "${var.base_url}:8080"
    name    = "${var.prefix}-origin-${count.index}"
    header  = "Host: ${var.hostname}"
  }
}`,
			expected: `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"

  origins = [
    {
      address = "${var.base_url}:8080",
      header  = "Host: ${var.hostname}",
      name    = "${var.prefix}-origin-${count.index}"
    }
  ]
}`,
		},
		{
			name: "preserve for expressions",
			input: `resource "cloudflare_load_balancer" "example" {
  zone_id = var.zone_id
  name    = var.lb_name

  region_pools {
    region   = "WNAM"
    pool_ids = [for p in var.pools : p.id]
  }

  rules {
    name     = "rule-1"
    condition = "http.request.uri.path eq '/api'"
    fixed_response = {for k, v in var.responses : k => v.content}
  }
}`,
			expected: `resource "cloudflare_load_balancer" "example" {
  zone_id = var.zone_id
  name    = var.lb_name

  region_pools = [
    {
      pool_ids = [for p in var.pools : p.id],
      region   = "WNAM"
    }
  ]
  rules = [
    {
      condition      = "http.request.uri.path eq '/api'",
      fixed_response = {for k, v in var.responses : k => v.content},
      name           = "rule-1"
    }
  ]
}`,
		},
		{
			name: "preserve conditional expressions",
			input: `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"

  origins {
    address = var.use_backup ? var.backup_address : var.primary_address
    name    = "origin-1"
    enabled = var.environment == "production" ? true : false
    weight  = var.weight != null ? var.weight : 100
  }
}`,
			expected: `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"

  origins = [
    {
      address = var.use_backup ? var.backup_address : var.primary_address,
      enabled = var.environment == "production" ? true : false,
      name    = "origin-1",
      weight  = var.weight != null ? var.weight : 100
    }
  ]
}`,
		},
		{
			name: "preserve function calls",
			input: `resource "cloudflare_load_balancer_monitor" "example" {
  type           = "https"
  expected_codes = "2xx"

  header {
    header = "Authorization"
    values = [base64encode(var.api_key)]
  }

  header {
    header = "X-Request-ID"
    values = [uuid()]
  }
}`,
			expected: `resource "cloudflare_load_balancer_monitor" "example" {
  type           = "https"
  expected_codes = "2xx"

  header = [
    {
      header = "Authorization",
      values = [base64encode(var.api_key)]
    },
    {
      header = "X-Request-ID",
      values = [uuid()]
    }
  ]
}`,
		},
		{
			name: "preserve complex nested expressions",
			input: `resource "cloudflare_load_balancer" "example" {
  zone_id = var.zone_id
  name    = var.lb_name

  country_pools {
    country  = "US"
    pool_ids = concat(var.primary_pools, var.backup_pools)
  }

  pop_pools {
    pop      = "LAX"
    pool_ids = [for p in var.pools : p.id if p.region == "west"]
  }

  region_pools {
    region   = "WNAM"
    pool_ids = distinct(flatten([
      for env in var.environments : [
        for pool in env.pools : pool.id
      ]
    ]))
  }
}`,
			expected: `resource "cloudflare_load_balancer" "example" {
  zone_id = var.zone_id
  name    = var.lb_name

  country_pools = [
    {
      country  = "US",
      pool_ids = concat(var.primary_pools, var.backup_pools)
    }
  ]
  pop_pools = [
    {
      pool_ids = [for p in var.pools : p.id if p.region == "west"],
      pop      = "LAX"
    }
  ]
  region_pools = [
    {
      pool_ids = distinct(flatten([
      for env in var.environments : [
        for pool in env.pools : pool.id
      ]
    ])),
      region   = "WNAM"
    }
  ]
}`,
		},
		{
			name: "preserve locals and data references",
			input: `resource "cloudflare_teams_location" "example" {
  name = "office"

  networks {
    id      = local.network_id
    network = data.cloudflare_ip_ranges.cloudflare.ipv4_cidr_blocks[0]
  }

  networks {
    id      = local.secondary_network_id
    network = join("/", [local.network_prefix, local.subnet_mask])
  }
}`,
			expected: `resource "cloudflare_teams_location" "example" {
  name = "office"

  networks = [
    {
      id      = local.network_id,
      network = data.cloudflare_ip_ranges.cloudflare.ipv4_cidr_blocks[0]
    },
    {
      id      = local.secondary_network_id,
      network = join("/", [local.network_prefix, local.subnet_mask])
    }
  ]
}`,
		},
		{
			name: "preserve multiple blocks with different variable types",
			input: `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"

  origins {
    address = var.origins[0].address
    name    = var.origins[0].name
    enabled = var.origins[0].enabled
  }

  origins {
    address = local.backup_origin
    name    = "backup-${var.environment}"
    weight  = lookup(var.weights, "backup", 50)
  }

  origins {
    address = "${data.aws_instance.app.private_ip}:8080"
    name    = data.aws_instance.app.tags["Name"]
    header  = jsonencode(var.custom_headers)
  }
}`,
			expected: `resource "cloudflare_load_balancer_pool" "example" {
  name = "example-pool"

  origins = [
    {
      address = var.origins[0].address,
      enabled = var.origins[0].enabled,
      name    = var.origins[0].name
    },
    {
      address = local.backup_origin,
      name    = "backup-${var.environment}",
      weight  = lookup(var.weights, "backup", 50)
    },
    {
      address = "${data.aws_instance.app.private_ip}:8080",
      header  = jsonencode(var.custom_headers),
      name    = data.aws_instance.app.tags["Name"]
    }
  ]
}`,
		},
		{
			name: "preserve zero trust resource with variables",
			input: `resource "cloudflare_access_mutual_tls_hostname_settings" "example" {
  zone_id = var.zone_id

  settings {
    hostname                      = var.hostname
    china_network                 = var.enable_china_network
    client_certificate_forwarding = var.cert_forwarding
  }

  settings {
    hostname                      = "${var.subdomain}.${var.domain}"
    china_network                 = false
    client_certificate_forwarding = true
  }
}`,
			expected: `resource "cloudflare_access_mutual_tls_hostname_settings" "example" {
  zone_id = var.zone_id

  settings = [
    {
      china_network                 = var.enable_china_network,
      client_certificate_forwarding = var.cert_forwarding,
      hostname                      = var.hostname
    },
    {
      china_network                 = false,
      client_certificate_forwarding = true,
      hostname                      = "${var.subdomain}.${var.domain}"
    }
  ]
}`,
		},
		{
			name: "preserve teams account with complex expressions",
			input: `resource "cloudflare_teams_account" "example" {
  account_id = var.account_id

  antivirus {
    enabled_download_phase = var.antivirus_config.download
    enabled_upload_phase   = var.antivirus_config.upload
    fail_closed            = var.environment == "production"
  }

  notification_settings {
    enabled     = true
    msg         = "${var.notification_prefix}: ${var.notification_message}"
    support_url = format("https://support.%s/help", var.domain)
  }

  block_page {
    background_color = var.theme.background
    enabled          = var.block_page_enabled
    logo_path        = var.logo_url != "" ? var.logo_url : "https://example.com/default.png"
  }
}`,
			expected: `resource "cloudflare_teams_account" "example" {
  account_id = var.account_id

  antivirus = [
    {
      enabled_download_phase = var.antivirus_config.download,
      enabled_upload_phase   = var.antivirus_config.upload,
      fail_closed            = var.environment == "production"
    }
  ]
  block_page = [
    {
      background_color = var.theme.background,
      enabled          = var.block_page_enabled,
      logo_path        = var.logo_url != "" ? var.logo_url : "https://example.com/default.png"
    }
  ]
  notification_settings = [
    {
      enabled     = true,
      msg         = "${var.notification_prefix}: ${var.notification_message}",
      support_url = format("https://support.%s/help", var.domain)
    }
  ]
}`,
		},
	}

	// Create transformer
	transformer, err := NewHCLTransformer(configPath)
	if err != nil {
		t.Fatalf("Failed to create transformer: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create input file
			inputPath := filepath.Join(tempDir, "input.tf")
			if err := os.WriteFile(inputPath, []byte(tt.input), 0644); err != nil {
				t.Fatalf("Failed to write input file: %v", err)
			}

			// Transform the file
			outputPath := filepath.Join(tempDir, "output.tf")
			if err := transformer.TransformFile(inputPath, outputPath); err != nil {
				t.Fatalf("Failed to transform file: %v", err)
			}

			// Read the output
			output, err := os.ReadFile(outputPath)
			if err != nil {
				t.Fatalf("Failed to read output file: %v", err)
			}

			// Use semantic comparison that ignores formatting differences and attribute order
			if !compareHCLBlocks(t, tt.expected, string(output)) {
				// For better debugging, show the actual output
				t.Errorf("Transformation mismatch\n\nExpected:\n%s\n\nGot:\n%s", tt.expected, string(output))
			}
		})
	}
}

func TestComplexCloudflareRulesetWithMultipleHeredocAndRatelimit(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "ruleset_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration file
	configPath := filepath.Join(tempDir, "test_config.yaml")
	configContent := `
version: "1.0"
description: "Test ruleset transformation"
transformations:
  cloudflare_ruleset:
    to_list:
      - rules
    to_map:
      - ratelimit
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

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

	// Create transformer
	transformer, err := NewHCLTransformer(configPath)
	if err != nil {
		t.Fatalf("Failed to create transformer: %v", err)
	}

	// Create input file
	inputPath := filepath.Join(tempDir, "input.tf")
	if err := os.WriteFile(inputPath, []byte(input), 0644); err != nil {
		t.Fatalf("Failed to write input file: %v", err)
	}

	// Transform the file
	outputPath := filepath.Join(tempDir, "output.tf")
	if err := transformer.TransformFile(inputPath, outputPath); err != nil {
		t.Fatalf("Failed to transform file: %v", err)
	}

	// Read the output
	output, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	// Verify the output is valid HCL
	_, diags := hclwrite.ParseConfig(output, "output.tf", hcl.InitialPos)
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

	// Check that ratelimit blocks are converted to maps within the list
	// Use regex to handle potential spacing variations
	if !regexp.MustCompile(`ratelimit\s*=\s*\{`).MatchString(string(output)) {
		t.Errorf("Expected ratelimit blocks to be converted to map, but output doesn't contain 'ratelimit = {'")
	} else {
		t.Logf("✅ Ratelimit blocks correctly converted to map format within list items")
	}

	// Check for balanced EOF markers
	eofCount := strings.Count(string(output), "EOF")
	if eofCount%2 != 0 {
		t.Errorf("Expected balanced EOF markers (pairs), but found %d total EOF occurrences", eofCount)
	} else {
		t.Logf("✅ EOF markers are balanced (%d pairs)", eofCount/2)
	}

	// Verify the local.account_id reference is preserved
	if !strings.Contains(string(output), "local.account_id") {
		t.Errorf("Expected local.account_id to be preserved")
	} else {
		t.Logf("✅ Local references preserved correctly")
	}

	t.Logf("Successfully transformed complex heredoc with ratelimit rules")
}

// TestExactMigrationTestCase tests the exact same configuration that fails in the migration tests
func TestExactMigrationTestCase(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "exact_migration_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test configuration file
	configPath := filepath.Join(tempDir, "test_config.yaml")
	configContent := `
version: "1.0"
description: "Test exact migration case"
transformations:
  cloudflare_ruleset:
    to_list:
      - rules
    to_map:
      - ratelimit
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	input := `resource "cloudflare_ruleset" "test" {
  account_id  = var.account_id
  kind        = "zone"
  name        = "Complex Heredoc Test"
  phase       = "http_ratelimit"
  description = "Test complex heredoc with ratelimit for ${var.environment}"

  rules {
    action      = "log"
    description = "Log rate limit with complex expression"
    enabled     = var.enable_rate_limiting
    expression  = <<-EOF
      not (
        any(http.request.headers["authorization"][*] contains "sess-")
        or any(http.request.headers["authorization"][*] contains "sk-")
        or http.cookie contains "__Secure-next-auth."
      )
      and not http.user_agent eq "Stripe/1.0 (+https://stripe.com/docs/webhooks)"
    EOF
    ratelimit {
      characteristics     = var.rate_characteristics
      mitigation_timeout  = var.mitigation_timeout
      period              = 10
      requests_per_period = 1000
    }
  }
}`

	// Create transformer
	transformer, err := NewHCLTransformer(configPath)
	if err != nil {
		t.Fatalf("Failed to create transformer: %v", err)
	}

	// Create input file
	inputPath := filepath.Join(tempDir, "input.tf")
	if err := os.WriteFile(inputPath, []byte(input), 0644); err != nil {
		t.Fatalf("Failed to write input file: %v", err)
	}

	// Transform the file
	outputPath := filepath.Join(tempDir, "output.tf")
	if err := transformer.TransformFile(inputPath, outputPath); err != nil {
		t.Fatalf("Failed to transform file: %v", err)
	}

	// Read the output
	output, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	
	// Log the output for debugging
	t.Logf("Generated output:\n%s", string(output))

	// Verify the output is valid HCL
	_, diags := hclwrite.ParseConfig(output, "output.tf", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("Transformed output is not valid HCL: %v", diags)
	}
	// Check that it has the expected structure
	if !strings.Contains(string(output), "rules = [") {
		t.Errorf("Expected rules to be converted to list format")
	} else {
		t.Logf("✅ Rules blocks correctly converted to list format")
	}

	// Check that ratelimit blocks are converted to maps within the list
	// Use regex to handle potential spacing variations
	if !regexp.MustCompile(`ratelimit\s*=\s*\{`).MatchString(string(output)) {
		t.Errorf("Expected ratelimit blocks to be converted to map, but output doesn't contain 'ratelimit = {'")
	} else {
		t.Logf("✅ Ratelimit blocks correctly converted to map format within list items")
	}

	// Check that variable interpolations are preserved
	if !strings.Contains(string(output), "var.account_id") {
		t.Errorf("Expected var.account_id to be preserved")
	}
	if !strings.Contains(string(output), "${var.environment}") {
		t.Errorf("Expected ${var.environment} interpolation to be preserved")
	}
	if !strings.Contains(string(output), "var.enable_rate_limiting") {
		t.Errorf("Expected var.enable_rate_limiting to be preserved")
	}
	if !strings.Contains(string(output), "var.rate_characteristics") {
		t.Errorf("Expected var.rate_characteristics to be preserved in nested ratelimit block")
	}
	if !strings.Contains(string(output), "var.mitigation_timeout") {
		t.Errorf("Expected var.mitigation_timeout to be preserved in nested ratelimit block")
	} else {
		t.Logf("✅ Variable interpolations preserved correctly including in nested blocks")
	}

	// Check that heredoc expressions are preserved
	if !strings.Contains(string(output), "<<-EOF") {
		t.Errorf("Expected heredoc expressions to be preserved")
	} else {
		t.Logf("✅ Heredoc expressions preserved correctly")
	}

	t.Logf("Successfully transformed configuration with heredocs and variable interpolations")
}
