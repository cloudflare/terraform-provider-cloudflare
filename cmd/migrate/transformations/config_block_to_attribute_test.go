package transformations

import (
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func TestLoadBalancerPoolOriginsWithVariableInterpolation(t *testing.T) {
	// Create test config
	configContent := `
version: "1.0"
transformations:
  cloudflare_load_balancer_pool:
    to_list:
      - origins
`
	configPath := "test_config.yaml"
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(configPath)

	// Load config
	config, err := LoadConfig(configPath)
	if err != nil {
		t.Fatal(err)
	}

	// Test input with variable interpolation
	input := `resource "cloudflare_load_balancer_pool" "test" {
  name = "test-pool"
  origins {
    name    = "u0"
    address = "api.unified-0.${var.api_openai_com_domain}"
    enabled = true
    weight  = 1
  }
  origins {
    name    = "u1"
    address = "api.unified-1.${var.api_openai_com_domain}"
    enabled = true
    weight  = 1
  }
}`

	// Parse
	file, diags := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatal(diags)
	}

	// Transform
	for _, block := range file.Body().Blocks() {
		if block.Type() == "resource" && len(block.Labels()) > 0 && block.Labels()[0] == "cloudflare_load_balancer_pool" {
			err := TransformResourceBlock(config, block, "cloudflare_load_balancer_pool")
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	// Check result
	result := string(file.Bytes())

	// Should NOT contain double dollar signs (escaped variables)
	if strings.Contains(result, "$${") {
		t.Errorf("Result contains escaped variables ($$):\n%s", result)
	}

	// Should still contain the variable reference
	if !strings.Contains(result, "${var.api_openai_com_domain}") {
		t.Errorf("Variable reference was lost:\n%s", result)
	}

	// Should have converted origins blocks to list attribute
	if !strings.Contains(result, "origins = [") {
		t.Errorf("Origins blocks were not converted to list attribute:\n%s", result)
	}

	// Should have removed the blocks
	if strings.Contains(result, "origins {") {
		t.Errorf("Origins blocks were not removed:\n%s", result)
	}

	// Check that both origins are in the list
	if !strings.Contains(result, "api.unified-0.${var.api_openai_com_domain}") {
		t.Errorf("First origin address was not preserved correctly:\n%s", result)
	}

	if !strings.Contains(result, "api.unified-1.${var.api_openai_com_domain}") {
		t.Errorf("Second origin address was not preserved correctly:\n%s", result)
	}
}

func TestNonLoadBalancerPoolOriginsUsesDefaultTransform(t *testing.T) {
	// Create test config for a different resource type
	configContent := `
version: "1.0"
transformations:
  cloudflare_other_resource:
    to_list:
      - items
`
	configPath := "test_config.yaml"
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(configPath)

	// Load config
	config, err := LoadConfig(configPath)
	if err != nil {
		t.Fatal(err)
	}

	// Test input with simple values
	input := `resource "cloudflare_other_resource" "test" {
  name = "test"
  items {
    name  = "item1"
    value = "simple-value"
  }
  items {
    name  = "item2"
    value = "another-value"
  }
}`

	// Parse
	file, diags := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatal(diags)
	}

	// Transform
	for _, block := range file.Body().Blocks() {
		if block.Type() == "resource" && len(block.Labels()) > 0 && block.Labels()[0] == "cloudflare_other_resource" {
			err := TransformResourceBlock(config, block, "cloudflare_other_resource")
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	// Check result
	result := string(file.Bytes())

	// Should have converted items blocks to list attribute
	if !strings.Contains(result, "items = [") {
		t.Errorf("Items blocks were not converted to list attribute:\n%s", result)
	}

	// Should have removed the blocks
	if strings.Contains(result, "items {") {
		t.Errorf("Items blocks were not removed:\n%s", result)
	}

	// Values should be preserved
	if !strings.Contains(result, "simple-value") || !strings.Contains(result, "another-value") {
		t.Errorf("Values were not preserved correctly:\n%s", result)
	}
}

func TestEmptyOriginsList(t *testing.T) {
	// Create test config
	configContent := `
version: "1.0"
transformations:
  cloudflare_load_balancer_pool:
    to_list:
      - origins
`
	configPath := "test_config.yaml"
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(configPath)

	// Load config
	config, err := LoadConfig(configPath)
	if err != nil {
		t.Fatal(err)
	}

	// Test input with no origins blocks
	input := `resource "cloudflare_load_balancer_pool" "test" {
  name = "test-pool"
  description = "A pool with no origins"
}`

	// Parse
	file, diags := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatal(diags)
	}

	// Transform
	for _, block := range file.Body().Blocks() {
		if block.Type() == "resource" && len(block.Labels()) > 0 && block.Labels()[0] == "cloudflare_load_balancer_pool" {
			err := TransformResourceBlock(config, block, "cloudflare_load_balancer_pool")
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	// Check result
	result := string(file.Bytes())

	// Should not have added origins attribute if there were no blocks
	if strings.Contains(result, "origins = ") || strings.Contains(result, "origins=") {
		t.Errorf("Origins attribute was added when no origins blocks existed:\n%s", result)
	}

	// Name and description should be preserved
	if !strings.Contains(result, "test-pool") {
		t.Errorf("Name attribute value was not preserved:\n%s", result)
	}

	if !strings.Contains(result, "A pool with no origins") {
		t.Errorf("Description attribute value was not preserved:\n%s", result)
	}
}

func TestMixedOriginsWithAndWithoutVariables(t *testing.T) {
	// Create test config
	configContent := `
version: "1.0"
transformations:
  cloudflare_load_balancer_pool:
    to_list:
      - origins
`
	configPath := "test_config.yaml"
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(configPath)

	// Load config
	config, err := LoadConfig(configPath)
	if err != nil {
		t.Fatal(err)
	}

	// Test input with mixed origins - some with variables, some without
	input := `resource "cloudflare_load_balancer_pool" "test" {
  name = "test-pool"
  origins {
    name    = "static"
    address = "192.168.1.1"
    enabled = true
  }
  origins {
    name    = "dynamic"
    address = "api.${var.domain_suffix}"
    enabled = false
  }
  origins {
    name    = "complex"
    address = "${var.prefix}.${var.domain}.${var.suffix}"
    weight  = 2
  }
}`

	// Parse
	file, diags := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatal(diags)
	}

	// Transform
	for _, block := range file.Body().Blocks() {
		if block.Type() == "resource" && len(block.Labels()) > 0 && block.Labels()[0] == "cloudflare_load_balancer_pool" {
			err := TransformResourceBlock(config, block, "cloudflare_load_balancer_pool")
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	// Check result
	result := string(file.Bytes())

	// Should NOT contain double dollar signs
	if strings.Contains(result, "$${") {
		t.Errorf("Result contains escaped variables ($$):\n%s", result)
	}

	// All addresses should be preserved correctly
	if !strings.Contains(result, "192.168.1.1") {
		t.Errorf("Static IP address was not preserved:\n%s", result)
	}

	if !strings.Contains(result, "api.${var.domain_suffix}") {
		t.Errorf("Simple variable interpolation was not preserved:\n%s", result)
	}

	if !strings.Contains(result, "${var.prefix}.${var.domain}.${var.suffix}") {
		t.Errorf("Complex variable interpolation was not preserved:\n%s", result)
	}

	// Should have converted to list
	if !strings.Contains(result, "origins = [") {
		t.Errorf("Origins blocks were not converted to list attribute:\n%s", result)
	}
}
