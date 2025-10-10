package cloudflare_zone_settings_override

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/registry"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
)

func TestZoneSettingsTransformation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name: "simple attributes",
			input: `
resource "cloudflare_zone_settings_override" "zone_settings" {
  zone_id = var.zone_id

  settings {
    automatic_https_rewrites = var.automatic_https_rewrites
    ssl                      = var.ssl
  }
}`,
			expected: []string{
				`resource "cloudflare_zone_setting" "zone_settings_automatic_https_rewrites"`,
				`setting_id = "automatic_https_rewrites"`,
				`value      = var.automatic_https_rewrites`,
				`resource "cloudflare_zone_setting" "zone_settings_ssl"`,
				`setting_id = "ssl"`,
				`value      = var.ssl`,
				`import {`,
				`to = cloudflare_zone_setting.zone_settings_automatic_https_rewrites`,
				`id = "${var.zone_id}/automatic_https_rewrites"`,
				`to = cloudflare_zone_setting.zone_settings_ssl`,
				`id = "${var.zone_id}/ssl"`,
			},
		},
		{
			name: "with security header block",
			input: `
resource "cloudflare_zone_settings_override" "zone_settings" {
  zone_id = var.zone_id

  settings {
    ssl = var.ssl

    security_header {
      enabled = var.security_header_enabled
      max_age = var.security_header_max_age
    }
  }
}`,
			expected: []string{
				`resource "cloudflare_zone_setting" "zone_settings_ssl"`,
				`setting_id = "ssl"`,
				`value      = var.ssl`,
				`resource "cloudflare_zone_setting" "zone_settings_security_header"`,
				`setting_id = "security_header"`,
				`strict_transport_security`,
				`enabled = var.security_header_enabled`,
				`max_age = var.security_header_max_age`,
				`import {`,
				`to = cloudflare_zone_setting.zone_settings_security_header`,
				`id = "${var.zone_id}/security_header"`,
			},
		},
		{
			name: "with nel block",
			input: `
resource "cloudflare_zone_settings_override" "zone_settings" {
  zone_id = var.zone_id

  settings {
    nel {
      enabled = var.enable_network_error_logging
    }
  }
}`,
			expected: []string{
				`resource "cloudflare_zone_setting" "zone_settings_nel"`,
				`setting_id = "nel"`,
				`value = {`,
				`enabled = var.enable_network_error_logging`,
				`import {`,
				`to = cloudflare_zone_setting.zone_settings_nel`,
				`id = "${var.zone_id}/nel"`,
			},
		},
		{
			name: "zero_rtt setting name mapping",
			input: `
resource "cloudflare_zone_settings_override" "test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  settings {
    zero_rtt = "on"
  }
}`,
			expected: []string{
				`resource "cloudflare_zone_setting" "test_zero_rtt"`,
				`setting_id = "0rtt"`, // zero_rtt maps to 0rtt
				`value      = "on"`,
				`import {`,
				`to = cloudflare_zone_setting.test_zero_rtt`,
				`id = "${"0da42c8d2132a9ddaf714f9e7c920711"}/0rtt"`,
			},
		},
		{
			name: "empty settings block",
			input: `
resource "cloudflare_zone_settings_override" "zone_settings" {
  zone_id = var.zone_id

  settings {
  }
}`,
			expected: []string{}, // No resources should be created
		},
		{
			name: "excludes deprecated universal_ssl setting",
			input: `
resource "cloudflare_zone_settings_override" "zone_settings" {
  zone_id = var.zone_id

  settings {
    ssl = "strict"
    universal_ssl = ""
    automatic_https_rewrites = "on"
  }
}`,
			expected: []string{
				`resource "cloudflare_zone_setting" "zone_settings_ssl"`,
				`setting_id = "ssl"`,
				`value      = "strict"`,
				`resource "cloudflare_zone_setting" "zone_settings_automatic_https_rewrites"`,
				`setting_id = "automatic_https_rewrites"`,
				`value      = "on"`,
				// universal_ssl should NOT appear in output
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create registry and register zone settings transformer
			reg := registry.NewStrategyRegistry()
			transformer := NewZoneSettingsOverride()
			reg.Register(transformer)

			// Parse the input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			assert.False(t, diags.HasErrors())

			// Find the zone settings override block
			var zsoBlock *hclwrite.Block
			for _, block := range file.Body().Blocks() {
				if block.Type() == "resource" && len(block.Labels()) >= 2 && block.Labels()[0] == "cloudflare_zone_settings_override" {
					zsoBlock = block
					break
				}
			}

			if zsoBlock == nil && len(tt.expected) > 0 {
				t.Fatal("No zone settings override block found")
			}

			if zsoBlock != nil {
				// TransformConfig using the wrapper function directly
				newBlocks := transformZoneSettingsBlock(zsoBlock, false)

				// Build output from all blocks
				outputFile := hclwrite.NewEmptyFile()
				for _, block := range newBlocks {
					outputFile.Body().AppendBlock(block)
				}

				result := string(outputFile.Bytes())

				// Check that all expected strings are present
				for _, expected := range tt.expected {
					assert.Contains(t, result, expected, "Missing expected content: %s", expected)
				}

				// Special check: ensure universal_ssl doesn't appear in output
				if tt.name == "excludes deprecated universal_ssl setting" {
					assert.NotContains(t, result, "universal_ssl", "Deprecated universal_ssl setting should not appear in output")
				}
			}
		})
	}
}

func TestZoneSettingsTransformationSkipImports(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []string
		notExpected []string
	}{
		{
			name: "skip imports - simple attributes",
			input: `
resource "cloudflare_zone_settings_override" "zone_settings" {
  zone_id = var.zone_id

  settings {
    automatic_https_rewrites = var.automatic_https_rewrites
    ssl                      = var.ssl
  }
}`,
			expected: []string{
				`resource "cloudflare_zone_setting" "zone_settings_automatic_https_rewrites"`,
				`setting_id = "automatic_https_rewrites"`,
				`value      = var.automatic_https_rewrites`,
				`resource "cloudflare_zone_setting" "zone_settings_ssl"`,
				`setting_id = "ssl"`,
				`value      = var.ssl`,
			},
			notExpected: []string{
				`import {`, // No import blocks should be generated
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			assert.False(t, diags.HasErrors())

			// Find the zone settings override block
			var zsoBlock *hclwrite.Block
			for _, block := range file.Body().Blocks() {
				if block.Type() == "resource" && len(block.Labels()) >= 2 && block.Labels()[0] == "cloudflare_zone_settings_override" {
					zsoBlock = block
					break
				}
			}

			if zsoBlock != nil {
				// TransformConfig with skipImports=true
				newBlocks := transformZoneSettingsBlock(zsoBlock, true)

				// Build output from all blocks
				outputFile := hclwrite.NewEmptyFile()
				for _, block := range newBlocks {
					outputFile.Body().AppendBlock(block)
				}

				result := string(outputFile.Bytes())

				// Check that all expected strings are present
				for _, expected := range tt.expected {
					assert.Contains(t, result, expected, "Missing expected content: %s", expected)
				}

				// Check that non-expected strings are not present
				for _, notExpected := range tt.notExpected {
					assert.NotContains(t, result, notExpected, "Found unexpected content: %s", notExpected)
				}
			}
		})
	}
}

// TestZoneSettingsStateTransformation tests the state file transformation
func TestZoneSettingsStateTransformation(t *testing.T) {
	tests := []struct {
		name         string
		inputState   string
		resourcePath string
		expected     string
	}{
		{
			name: "deletes zone settings override from state",
			inputState: `{
				"version": 4,
				"terraform_version": "1.0.0",
				"serial": 1,
				"lineage": "test",
				"outputs": {},
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_zone_settings_override",
						"name": "test",
						"provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
						"instances": []
					},
					{
						"mode": "managed",
						"type": "cloudflare_zone",
						"name": "test",
						"provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
						"instances": []
					}
				]
			}`,
			resourcePath: "resources.0",
			expected:     `"cloudflare_zone"`, // Only the zone resource should remain
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transformZoneSettingsStateJSON(tt.inputState, tt.resourcePath)

			// Check the expected content is present
			assert.Contains(t, result, tt.expected)

			// Check that zone_settings_override was removed
			assert.NotContains(t, result, "cloudflare_zone_settings_override")
		})
	}
}

// TestZoneSettingsIntegration tests the full pipeline integration
func TestZoneSettingsIntegration(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name: "complete zone settings transformation through pipeline",
			input: `
resource "cloudflare_zone_settings_override" "test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"

  settings {
    always_online = "on"
    zero_rtt      = "on"
    ssl           = "flexible"

    security_header {
      enabled            = true
      preload            = true
      include_subdomains = true
      max_age            = 86400
    }
  }
}`,
			expected: []string{
				`cloudflare_zone_setting" "test_always_online"`,
				`cloudflare_zone_setting" "test_zero_rtt"`,
				`cloudflare_zone_setting" "test_ssl"`,
				`cloudflare_zone_setting" "test_security_header"`,
				`setting_id = "0rtt"`, // zero_rtt should map to 0rtt
				`strict_transport_security`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create registry and register zone settings transformer
			reg := registry.NewStrategyRegistry()
			reg.Register(NewZoneSettingsOverride())

			// Build and run pipeline
			pipeline := poc.BuildPipeline(reg)
			result, err := pipeline.Transform([]byte(tt.input), "test.tf")
			assert.NoError(t, err)

			resultStr := string(result)

			// Check that all expected strings are present
			for _, expected := range tt.expected {
				assert.Contains(t, resultStr, expected, "Missing expected content: %s", expected)
			}
		})
	}
}
