package basic

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReferenceUpdater(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		mappings       map[string]string
		expectedOutput string
		checkFunc      func(t *testing.T, output string)
	}{
		{
			name: "update simple resource references",
			input: `resource "cloudflare_new_resource" "example" {
  zone_id = "abc123"
  policy_id = cloudflare_access_policy.main.id
  group_id = cloudflare_access_group.admins.id
}`,
			mappings: map[string]string{
				"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
				"cloudflare_access_group":  "cloudflare_zero_trust_access_group",
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "cloudflare_zero_trust_access_policy.main.id")
				assert.Contains(t, output, "cloudflare_zero_trust_access_group.admins.id")
				assert.NotContains(t, output, "cloudflare_access_policy")
				assert.NotContains(t, output, "cloudflare_access_group")
			},
		},
		{
			name: "update references in interpolations",
			input: `resource "cloudflare_new_resource" "example" {
  zone_id = "abc123"
  description = "Policy: ${cloudflare_access_policy.main.id}"
  combined = "${cloudflare_access_group.admins.id}-${cloudflare_access_policy.main.name}"
}`,
			mappings: map[string]string{
				"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
				"cloudflare_access_group":  "cloudflare_zero_trust_access_group",
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "${cloudflare_zero_trust_access_policy.main.id}")
				assert.Contains(t, output, "${cloudflare_zero_trust_access_group.admins.id}")
				assert.Contains(t, output, "${cloudflare_zero_trust_access_policy.main.name}")
			},
		},
		{
			name: "update data source references",
			input: `resource "cloudflare_new_resource" "example" {
  zone_id = data.cloudflare_zone.example.id
  account_id = data.cloudflare_account.main.id
  origin_ca = data.cloudflare_origin_ca.cert.id
}`,
			mappings: map[string]string{
				"cloudflare_account":   "cloudflare_accounts",
				"cloudflare_origin_ca": "cloudflare_origin_ca_certificate",
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "data.cloudflare_accounts.main.id")
				assert.Contains(t, output, "data.cloudflare_origin_ca_certificate.cert.id")
				assert.Contains(t, output, "data.cloudflare_zone.example.id") // unchanged
			},
		},
		{
			name: "update nested block references",
			input: `resource "cloudflare_new_resource" "example" {
  zone_id = "abc123"
  
  rule {
    policy = cloudflare_access_policy.main.id
    group = cloudflare_access_group.users.id
  }
  
  condition {
    reference = cloudflare_access_application.app.id
  }
}`,
			mappings: map[string]string{
				"cloudflare_access_policy":      "cloudflare_zero_trust_access_policy",
				"cloudflare_access_group":       "cloudflare_zero_trust_access_group",
				"cloudflare_access_application": "cloudflare_zero_trust_access_application",
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "cloudflare_zero_trust_access_policy.main.id")
				assert.Contains(t, output, "cloudflare_zero_trust_access_group.users.id")
				assert.Contains(t, output, "cloudflare_zero_trust_access_application.app.id")
			},
		},
		{
			name: "handle multiple resource types in single attribute",
			input: `resource "cloudflare_new_resource" "example" {
  zone_id = "abc123"
  refs = [
    cloudflare_access_policy.one.id,
    cloudflare_access_policy.two.id,
    cloudflare_access_group.admin.id
  ]
}`,
			mappings: map[string]string{
				"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
				"cloudflare_access_group":  "cloudflare_zero_trust_access_group",
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "cloudflare_zero_trust_access_policy.one.id")
				assert.Contains(t, output, "cloudflare_zero_trust_access_policy.two.id")
				assert.Contains(t, output, "cloudflare_zero_trust_access_group.admin.id")
			},
		},
		{
			name: "preserve non-matching references",
			input: `resource "cloudflare_new_resource" "example" {
  zone_id = var.zone_id
  account = local.account_id
  module_ref = module.networking.subnet_id
  other_ref = aws_instance.example.id
}`,
			mappings: map[string]string{
				"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "var.zone_id")
				assert.Contains(t, output, "local.account_id")
				assert.Contains(t, output, "module.networking.subnet_id")
				assert.Contains(t, output, "aws_instance.example.id")
			},
		},
		{
			name: "update for_each references",
			input: `resource "cloudflare_new_resource" "example" {
  for_each = cloudflare_access_policy.all
  
  zone_id = "abc123"
  policy_id = each.value.id
  parent = cloudflare_access_policy.all[each.key].id
}`,
			mappings: map[string]string{
				"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "cloudflare_zero_trust_access_policy.all")
				assert.Contains(t, output, "cloudflare_zero_trust_access_policy.all[each.key].id")
			},
		},
		{
			name: "avoid partial replacements",
			input: `resource "cloudflare_new_resource" "example" {
  zone_id = "abc123"
  short = cloudflare_zone.example.id
  longer = cloudflare_zone_settings.example.id
}`,
			mappings: map[string]string{
				"cloudflare_zone_settings": "cloudflare_zone_settings_override",
				"cloudflare_zone":          "cloudflare_zones", // Should not affect cloudflare_zone_settings
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "cloudflare_zones.example.id")
				assert.Contains(t, output, "cloudflare_zone_settings_override.example.id")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			// Get the first block
			blocks := file.Body().Blocks()
			require.Len(t, blocks, 1)
			block := blocks[0]

			// Create context and apply transformation
			ctx := &TransformContext{}
			updater := ReferenceUpdater(tt.mappings)
			
			err := updater(block, ctx)
			require.NoError(t, err)

			// Get output
			output := string(hclwrite.Format(file.Bytes()))

			// Run custom checks
			if tt.checkFunc != nil {
				tt.checkFunc(t, output)
			}
		})
	}
}

func TestSpecificAttributeReferenceUpdater(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		attributes []string
		mappings   map[string]string
		checkFunc  func(t *testing.T, output string)
	}{
		{
			name: "update only specified attributes",
			input: `resource "cloudflare_new_resource" "example" {
  zone_id = "abc123"
  policy_id = cloudflare_access_policy.main.id
  group_id = cloudflare_access_group.admins.id
  app_id = cloudflare_access_application.app.id
}`,
			attributes: []string{"policy_id", "app_id"},
			mappings: map[string]string{
				"cloudflare_access_policy":      "cloudflare_zero_trust_access_policy",
				"cloudflare_access_group":       "cloudflare_zero_trust_access_group",
				"cloudflare_access_application": "cloudflare_zero_trust_access_application",
			},
			checkFunc: func(t *testing.T, output string) {
				// Only policy_id and app_id should be updated
				assert.Contains(t, output, "cloudflare_zero_trust_access_policy.main.id")
				assert.Contains(t, output, "cloudflare_zero_trust_access_application.app.id")
				// group_id should not be updated
				assert.Contains(t, output, "cloudflare_access_group.admins.id")
			},
		},
		{
			name: "ignore non-specified attributes",
			input: `resource "cloudflare_new_resource" "example" {
  zone_id = cloudflare_zone.example.id
  account = cloudflare_account.main.id
  target = cloudflare_record.target.id
}`,
			attributes: []string{"target"},
			mappings: map[string]string{
				"cloudflare_zone":    "cloudflare_zones",
				"cloudflare_account": "cloudflare_accounts",
				"cloudflare_record":  "cloudflare_dns_record",
			},
			checkFunc: func(t *testing.T, output string) {
				// Only target should be updated
				assert.Contains(t, output, "cloudflare_dns_record.target.id")
				// Others should remain unchanged
				assert.Contains(t, output, "cloudflare_zone.example.id")
				assert.Contains(t, output, "cloudflare_account.main.id")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &TransformContext{}
			
			updater := SpecificAttributeReferenceUpdater(tt.attributes, tt.mappings)
			err := updater(block, ctx)
			require.NoError(t, err)

			output := string(hclwrite.Format(file.Bytes()))
			if tt.checkFunc != nil {
				tt.checkFunc(t, output)
			}
		})
	}
}

func TestPatternBasedReferenceUpdater(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		patterns  []ReferencePattern
		checkFunc func(t *testing.T, output string)
	}{
		{
			name: "simple pattern replacement",
			input: `resource "cloudflare_new_resource" "example" {
  zone_id = "abc123"
  old_ref = cloudflare_old_resource.example.id
  another = cloudflare_old_thing.test.arn
}`,
			patterns: []ReferencePattern{
				{
					Pattern:     `cloudflare_old_resource`,
					Replacement: `cloudflare_new_resource`,
				},
				{
					Pattern:     `cloudflare_old_thing`,
					Replacement: `cloudflare_new_thing`,
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "cloudflare_new_resource.example.id")
				assert.Contains(t, output, "cloudflare_new_thing.test.arn")
			},
		},
		{
			name: "complex regex patterns",
			input: `resource "cloudflare_new_resource" "example" {
  zone_id = "abc123"
  ref1 = cloudflare_access_policy.main.id
  ref2 = cloudflare_access_group.users.id
  ref3 = cloudflare_access_application.app.id
}`,
			patterns: []ReferencePattern{
				{
					Pattern:     `cloudflare_access_(\w+)`,
					Replacement: `cloudflare_zero_trust_access_$1`,
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "cloudflare_zero_trust_access_policy.main.id")
				assert.Contains(t, output, "cloudflare_zero_trust_access_group.users.id")
				assert.Contains(t, output, "cloudflare_zero_trust_access_application.app.id")
			},
		},
		{
			name: "attribute suffix patterns",
			input: `resource "cloudflare_new_resource" "example" {
  zone_id = "abc123"
  ref_id = cloudflare_resource.example.resource_id
  ref_arn = cloudflare_resource.example.resource_arn
}`,
			patterns: []ReferencePattern{
				{
					Pattern:     `\.resource_id\b`,
					Replacement: `.id`,
				},
				{
					Pattern:     `\.resource_arn\b`,
					Replacement: `.arn`,
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "cloudflare_resource.example.id")
				assert.Contains(t, output, "cloudflare_resource.example.arn")
				assert.NotContains(t, output, "resource_id")
				assert.NotContains(t, output, "resource_arn")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &TransformContext{}
			
			updater := PatternBasedReferenceUpdater(tt.patterns)
			err := updater(block, ctx)
			require.NoError(t, err)

			output := string(hclwrite.Format(file.Bytes()))
			if tt.checkFunc != nil {
				tt.checkFunc(t, output)
			}
		})
	}
}

func TestReferenceUpdaterForState(t *testing.T) {
	tests := []struct {
		name          string
		state         map[string]interface{}
		mappings      map[string]string
		expectedState map[string]interface{}
	}{
		{
			name: "update references in state attributes",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id":   "abc123",
					"policy_id": "cloudflare_access_policy.main.id",
					"group_ref": "${cloudflare_access_group.admins.id}",
				},
			},
			mappings: map[string]string{
				"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
				"cloudflare_access_group":  "cloudflare_zero_trust_access_group",
			},
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id":   "abc123",
					"policy_id": "cloudflare_zero_trust_access_policy.main.id",
					"group_ref": "${cloudflare_zero_trust_access_group.admins.id}",
				},
			},
		},
		{
			name: "update nested map references",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"config": map[string]interface{}{
						"policy": "cloudflare_access_policy.main.id",
						"group":  "cloudflare_access_group.users.id",
					},
				},
			},
			mappings: map[string]string{
				"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
				"cloudflare_access_group":  "cloudflare_zero_trust_access_group",
			},
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"config": map[string]interface{}{
						"policy": "cloudflare_zero_trust_access_policy.main.id",
						"group":  "cloudflare_zero_trust_access_group.users.id",
					},
				},
			},
		},
		{
			name: "update array references",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"refs": []interface{}{
						"cloudflare_access_policy.one.id",
						"cloudflare_access_policy.two.id",
						"cloudflare_access_group.admin.id",
					},
				},
			},
			mappings: map[string]string{
				"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
				"cloudflare_access_group":  "cloudflare_zero_trust_access_group",
			},
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"refs": []interface{}{
						"cloudflare_zero_trust_access_policy.one.id",
						"cloudflare_zero_trust_access_policy.two.id",
						"cloudflare_zero_trust_access_group.admin.id",
					},
				},
			},
		},
		{
			name: "update dependencies",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
				},
				"dependencies": []interface{}{
					"cloudflare_access_policy.main",
					"cloudflare_access_group.users",
				},
			},
			mappings: map[string]string{
				"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
				"cloudflare_access_group":  "cloudflare_zero_trust_access_group",
			},
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
				},
				"dependencies": []interface{}{
					"cloudflare_zero_trust_access_policy.main",
					"cloudflare_zero_trust_access_group.users",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updater := ReferenceUpdaterForState(tt.mappings)
			err := updater(tt.state)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedState, tt.state)
		})
	}
}

func TestGlobalReferenceUpdater(t *testing.T) {
	input := `
locals {
  policy = cloudflare_access_policy.main.id
}

resource "cloudflare_new_resource" "example" {
  zone_id = "abc123"
  policy_id = cloudflare_access_policy.other.id
}

data "cloudflare_zone" "example" {
  name = "example.com"
  account_id = cloudflare_account.main.id
}`

	mappings := map[string]string{
		"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
		"cloudflare_account":       "cloudflare_accounts",
	}

	file, diags := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	updater := GlobalReferenceUpdater(mappings)
	err := updater(file)
	require.NoError(t, err)

	output := string(hclwrite.Format(file.Bytes()))

	// Check all references were updated
	assert.Contains(t, output, "cloudflare_zero_trust_access_policy.main.id")
	assert.Contains(t, output, "cloudflare_zero_trust_access_policy.other.id")
	assert.Contains(t, output, "cloudflare_accounts.main.id")
	assert.NotContains(t, output, "cloudflare_access_policy")
	assert.NotContains(t, output, "cloudflare_account.main")
}

func TestChainedReferenceUpdater(t *testing.T) {
	input := `resource "cloudflare_new_resource" "example" {
  zone_id = "abc123"
  policy_id = cloudflare_access_policy.main.id
  group_id = cloudflare_access_group.admins.id
  app_id = cloudflare_access_application.app.id
}`

	updates := []ReferenceUpdate{
		{
			// First update - always apply
			Mappings: map[string]string{
				"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
			},
		},
		{
			// Second update - conditional
			Condition: func(block *hclwrite.Block) bool {
				// Only apply if block has group_id attribute
				return block.Body().GetAttribute("group_id") != nil
			},
			Mappings: map[string]string{
				"cloudflare_access_group": "cloudflare_zero_trust_access_group",
			},
		},
		{
			// Third update - always apply
			Mappings: map[string]string{
				"cloudflare_access_application": "cloudflare_zero_trust_access_application",
			},
		},
	}

	file, diags := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	block := file.Body().Blocks()[0]
	ctx := &TransformContext{}
	
	updater := ChainedReferenceUpdater(updates)
	err := updater(block, ctx)
	require.NoError(t, err)

	output := string(hclwrite.Format(file.Bytes()))

	// All updates should be applied
	assert.Contains(t, output, "cloudflare_zero_trust_access_policy.main.id")
	assert.Contains(t, output, "cloudflare_zero_trust_access_group.admins.id")
	assert.Contains(t, output, "cloudflare_zero_trust_access_application.app.id")
}

func TestReferenceUpdaterEdgeCases(t *testing.T) {
	t.Run("empty mappings", func(t *testing.T) {
		updater := ReferenceUpdater(map[string]string{})
		
		input := `resource "test" "example" { ref = cloudflare_policy.main.id }`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		ctx := &TransformContext{}
		err := updater(block, ctx)
		assert.NoError(t, err)
		
		// Should remain unchanged
		output := string(file.Bytes())
		assert.Contains(t, output, "cloudflare_policy.main.id")
	})
	
	t.Run("nil mappings", func(t *testing.T) {
		updater := ReferenceUpdater(nil)
		
		input := `resource "test" "example" { ref = cloudflare_policy.main.id }`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		ctx := &TransformContext{}
		err := updater(block, ctx)
		assert.NoError(t, err)
	})
	
	t.Run("no references in attributes", func(t *testing.T) {
		mappings := map[string]string{
			"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
		}
		
		input := `resource "test" "example" { 
  zone_id = "abc123"
  name = "example"
  count = 5
}`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		ctx := &TransformContext{}
		updater := ReferenceUpdater(mappings)
		err := updater(block, ctx)
		assert.NoError(t, err)
		
		// Should remain unchanged
		output := string(hclwrite.Format(file.Bytes()))
		assert.Contains(t, output, `zone_id = "abc123"`)
		assert.Contains(t, output, `name    = "example"`)
	})
	
	t.Run("deeply nested blocks", func(t *testing.T) {
		mappings := map[string]string{
			"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
		}
		
		input := `resource "test" "example" {
  zone_id = "abc123"
  
  level1 {
    ref = cloudflare_access_policy.one.id
    
    level2 {
      ref = cloudflare_access_policy.two.id
      
      level3 {
        ref = cloudflare_access_policy.three.id
      }
    }
  }
}`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		ctx := &TransformContext{}
		updater := ReferenceUpdater(mappings)
		err := updater(block, ctx)
		assert.NoError(t, err)
		
		output := string(file.Bytes())
		assert.Contains(t, output, "cloudflare_zero_trust_access_policy.one.id")
		assert.Contains(t, output, "cloudflare_zero_trust_access_policy.two.id")
		assert.Contains(t, output, "cloudflare_zero_trust_access_policy.three.id")
	})
	
	t.Run("complex expressions", func(t *testing.T) {
		mappings := map[string]string{
			"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
		}
		
		input := `resource "test" "example" {
  zone_id = "abc123"
  complex = join(",", [cloudflare_access_policy.one.id, cloudflare_access_policy.two.id])
  ternary = var.condition ? cloudflare_access_policy.main.id : ""
}`
		file, _ := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		
		ctx := &TransformContext{}
		updater := ReferenceUpdater(mappings)
		err := updater(block, ctx)
		assert.NoError(t, err)
		
		output := string(file.Bytes())
		assert.Contains(t, output, "cloudflare_zero_trust_access_policy.one.id")
		assert.Contains(t, output, "cloudflare_zero_trust_access_policy.two.id")
		assert.Contains(t, output, "cloudflare_zero_trust_access_policy.main.id")
	})
}