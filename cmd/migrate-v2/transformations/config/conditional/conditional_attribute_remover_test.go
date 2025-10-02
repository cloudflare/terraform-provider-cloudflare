package conditional

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConditionalAttributeRemover(t *testing.T) {
	tests := []struct {
		name           string
		hclContent     string
		targetAttr     string
		conditionAttr  string
		allowedValues  []string
		expectRemoved  bool
	}{
		{
			name: "RemoveWhenConditionNotInAllowed",
			hclContent: `resource "test" "example" {
  type = "self_hosted"
  skip_app_launcher_login_page = true
}`,
			targetAttr:    "skip_app_launcher_login_page",
			conditionAttr: "type",
			allowedValues: []string{"app_launcher"},
			expectRemoved: true,
		},
		{
			name: "KeepWhenConditionInAllowed",
			hclContent: `resource "test" "example" {
  type = "app_launcher"
  skip_app_launcher_login_page = true
}`,
			targetAttr:    "skip_app_launcher_login_page",
			conditionAttr: "type",
			allowedValues: []string{"app_launcher"},
			expectRemoved: false,
		},
		{
			name: "RemoveWhenConditionMissing",
			hclContent: `resource "test" "example" {
  skip_app_launcher_login_page = true
}`,
			targetAttr:    "skip_app_launcher_login_page",
			conditionAttr: "type",
			allowedValues: []string{"app_launcher"},
			expectRemoved: true,
		},
		{
			name: "NoOpWhenTargetMissing",
			hclContent: `resource "test" "example" {
  type = "app_launcher"
}`,
			targetAttr:    "skip_app_launcher_login_page",
			conditionAttr: "type",
			allowedValues: []string{"app_launcher"},
			expectRemoved: false,
		},
		{
			name: "MultipleAllowedValues",
			hclContent: `resource "test" "example" {
  type = "saml"
  auth_attribute = "value"
}`,
			targetAttr:    "auth_attribute",
			conditionAttr: "type",
			allowedValues: []string{"saml", "oauth"},
			expectRemoved: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hclContent), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &basic.TransformContext{}

			transformer := ConditionalAttributeRemover(
				tt.targetAttr,
				tt.conditionAttr,
				tt.allowedValues,
				true, // add diagnostic
			)
			err := transformer(block, ctx)
			require.NoError(t, err)

			body := block.Body()
			attr := body.GetAttribute(tt.targetAttr)

			if tt.expectRemoved {
				assert.Nil(t, attr, "Attribute should have been removed")
				if attr == nil {
					assert.NotEmpty(t, ctx.Diagnostics, "Should have diagnostic when removing")
				}
			} else {
				if body.GetAttribute(tt.targetAttr) != nil {
					assert.NotNil(t, attr, "Attribute should have been kept")
				}
			}
		})
	}
}

func TestConditionalAttributeRemoverWithDefault(t *testing.T) {
	tests := []struct {
		name          string
		hclContent    string
		targetAttr    string
		conditionAttr string
		allowedValues []string
		defaultValue  string
		expectRemoved bool
	}{
		{
			name: "UseDefaultWhenConditionMissing",
			hclContent: `resource "test" "example" {
  skip_app_launcher_login_page = true
}`,
			targetAttr:    "skip_app_launcher_login_page",
			conditionAttr: "type",
			allowedValues: []string{"app_launcher"},
			defaultValue:  "self_hosted",
			expectRemoved: true,
		},
		{
			name: "UseActualValueWhenPresent",
			hclContent: `resource "test" "example" {
  type = "app_launcher"
  skip_app_launcher_login_page = true
}`,
			targetAttr:    "skip_app_launcher_login_page",
			conditionAttr: "type",
			allowedValues: []string{"app_launcher"},
			defaultValue:  "self_hosted",
			expectRemoved: false,
		},
		{
			name: "DefaultValueInAllowedList",
			hclContent: `resource "test" "example" {
  feature_flag = true
}`,
			targetAttr:    "feature_flag",
			conditionAttr: "mode",
			allowedValues: []string{"advanced", "expert"},
			defaultValue:  "advanced",
			expectRemoved: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hclContent), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &basic.TransformContext{}

			transformer := ConditionalAttributeRemoverWithDefault(
				tt.targetAttr,
				tt.conditionAttr,
				tt.allowedValues,
				tt.defaultValue,
				true, // add diagnostic
			)
			err := transformer(block, ctx)
			require.NoError(t, err)

			body := block.Body()
			attr := body.GetAttribute(tt.targetAttr)

			if tt.expectRemoved {
				assert.Nil(t, attr, "Attribute should have been removed")
				if attr == nil {
					assert.NotEmpty(t, ctx.Diagnostics, "Should have diagnostic when removing")
				}
			} else {
				if body.GetAttribute(tt.targetAttr) != nil {
					assert.NotNil(t, attr, "Attribute should have been kept")
				}
			}
		})
	}
}

func TestConditionalAttributeRemover_Diagnostics(t *testing.T) {
	hclContent := `resource "test" "example" {
  type = "other"
  skip_app_launcher_login_page = true
}`

	file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	block := file.Body().Blocks()[0]
	ctx := &basic.TransformContext{}

	transformer := ConditionalAttributeRemover(
		"skip_app_launcher_login_page",
		"type",
		[]string{"app_launcher"},
		true, // add diagnostic
	)
	err := transformer(block, ctx)
	require.NoError(t, err)

	// Check diagnostic was added
	assert.Len(t, ctx.Diagnostics, 1)
	assert.Contains(t, ctx.Diagnostics[0], "Removed skip_app_launcher_login_page")
	assert.Contains(t, ctx.Diagnostics[0], "type is one of: app_launcher")
}