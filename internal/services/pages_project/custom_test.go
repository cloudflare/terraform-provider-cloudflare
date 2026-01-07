package pages_project

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TestPreserveSecretEnvVars_ReturnsSourceWhenNonSecretValuesMatch tests the fix for #6526.
func TestPreserveSecretEnvVars_ReturnsSourceWhenNonSecretValuesMatch(t *testing.T) {
	ctx := context.Background()

	// Create source (plan) with secret env var value
	sourcePreview := &PagesProjectDeploymentConfigsPreviewModel{
		CompatibilityDate: types.StringValue("2024-01-17"),
		FailOpen:          types.BoolValue(true),
		EnvVars: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
			"HUGO_VERSION": {
				Type:  types.StringValue("plain_text"),
				Value: types.StringValue("0.145.0"),
			},
			"SECRET_KEY": {
				Type:  types.StringValue("secret_text"),
				Value: types.StringValue("my-secret-value"),
			},
		},
	}
	sourceProduction := &PagesProjectDeploymentConfigsProductionModel{
		CompatibilityDate: types.StringValue("2024-01-17"),
		FailOpen:          types.BoolValue(true),
		EnvVars: &map[string]PagesProjectDeploymentConfigsProductionEnvVarsModel{
			"HUGO_VERSION": {
				Type:  types.StringValue("plain_text"),
				Value: types.StringValue("0.145.0"),
			},
			"SECRET_KEY": {
				Type:  types.StringValue("secret_text"),
				Value: types.StringValue("my-secret-value"),
			},
		},
	}

	sourcePreviewObj, _ := customfield.NewObject(ctx, sourcePreview)
	sourceProductionObj, _ := customfield.NewObject(ctx, sourceProduction)
	sourceConfigs := &PagesProjectDeploymentConfigsModel{
		Preview:    sourcePreviewObj,
		Production: sourceProductionObj,
	}
	source, _ := customfield.NewObject(ctx, sourceConfigs)

	// Create destination (API response) with empty secret value (as API returns)
	destPreview := &PagesProjectDeploymentConfigsPreviewModel{
		CompatibilityDate: types.StringValue("2024-01-17"),
		FailOpen:          types.BoolValue(true),
		EnvVars: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
			"HUGO_VERSION": {
				Type:  types.StringValue("plain_text"),
				Value: types.StringValue("0.145.0"),
			},
			"SECRET_KEY": {
				Type:  types.StringValue("secret_text"),
				Value: types.StringValue(""),
			},
		},
	}
	destProduction := &PagesProjectDeploymentConfigsProductionModel{
		CompatibilityDate: types.StringValue("2024-01-17"),
		FailOpen:          types.BoolValue(true),
		EnvVars: &map[string]PagesProjectDeploymentConfigsProductionEnvVarsModel{
			"HUGO_VERSION": {
				Type:  types.StringValue("plain_text"),
				Value: types.StringValue("0.145.0"),
			},
			"SECRET_KEY": {
				Type:  types.StringValue("secret_text"),
				Value: types.StringValue(""),
			},
		},
	}

	destPreviewObj, _ := customfield.NewObject(ctx, destPreview)
	destProductionObj, _ := customfield.NewObject(ctx, destProduction)
	destConfigs := &PagesProjectDeploymentConfigsModel{
		Preview:    destPreviewObj,
		Production: destProductionObj,
	}
	dest, _ := customfield.NewObject(ctx, destConfigs)

	// Call PreserveSecretEnvVars
	result, diags := PreserveSecretEnvVars(ctx, source, dest)

	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	// Verify values are equal - fix returns source directly when non-secret values match
	if !result.Equal(source) {
		t.Errorf("expected result to equal source")
	}

	// Verify the secret value is preserved
	resultValue, _ := result.Value(ctx)
	resultPreview, _ := resultValue.Preview.Value(ctx)
	if resultPreview.EnvVars == nil {
		t.Fatal("expected env_vars to be present")
	}
	secretVar, exists := (*resultPreview.EnvVars)["SECRET_KEY"]
	if !exists {
		t.Fatal("expected SECRET_KEY to exist")
	}
	if secretVar.Value.ValueString() != "my-secret-value" {
		t.Errorf("expected secret value to be preserved, got %q", secretVar.Value.ValueString())
	}
}

// TestPreserveSecretEnvVars_MergesWhenNonSecretValuesDiffer tests merging when values differ.
func TestPreserveSecretEnvVars_MergesWhenNonSecretValuesDiffer(t *testing.T) {
	ctx := context.Background()

	// Create source (plan) with certain values
	sourcePreview := &PagesProjectDeploymentConfigsPreviewModel{
		CompatibilityDate: types.StringValue("2024-01-17"),
		FailOpen:          types.BoolValue(true),
		EnvVars: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
			"SECRET_KEY": {
				Type:  types.StringValue("secret_text"),
				Value: types.StringValue("my-secret-value"),
			},
		},
	}
	sourceProduction := &PagesProjectDeploymentConfigsProductionModel{
		CompatibilityDate: types.StringValue("2024-01-17"),
		FailOpen:          types.BoolValue(true),
	}

	sourcePreviewObj, _ := customfield.NewObject(ctx, sourcePreview)
	sourceProductionObj, _ := customfield.NewObject(ctx, sourceProduction)
	sourceConfigs := &PagesProjectDeploymentConfigsModel{
		Preview:    sourcePreviewObj,
		Production: sourceProductionObj,
	}
	source, _ := customfield.NewObject(ctx, sourceConfigs)

	// Create destination (API response) with DIFFERENT non-secret values
	destPreview := &PagesProjectDeploymentConfigsPreviewModel{
		CompatibilityDate: types.StringValue("2024-01-18"), // Different date!
		FailOpen:          types.BoolValue(true),
		EnvVars: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
			"SECRET_KEY": {
				Type:  types.StringValue("secret_text"),
				Value: types.StringValue(""),
			},
		},
	}
	destProduction := &PagesProjectDeploymentConfigsProductionModel{
		CompatibilityDate: types.StringValue("2024-01-18"), // Different date!
		FailOpen:          types.BoolValue(true),
	}

	destPreviewObj, _ := customfield.NewObject(ctx, destPreview)
	destProductionObj, _ := customfield.NewObject(ctx, destProduction)
	destConfigs := &PagesProjectDeploymentConfigsModel{
		Preview:    destPreviewObj,
		Production: destProductionObj,
	}
	dest, _ := customfield.NewObject(ctx, destConfigs)

	// Call PreserveSecretEnvVars
	result, diags := PreserveSecretEnvVars(ctx, source, dest)

	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	// Result should NOT equal source since non-secret values differ
	if result.Equal(source) {
		t.Errorf("expected result to differ from source when non-secret values differ")
	}

	// But secret values should still be preserved
	resultValue, _ := result.Value(ctx)
	resultPreview, _ := resultValue.Preview.Value(ctx)
	if resultPreview.EnvVars == nil {
		t.Fatal("expected env_vars to be present")
	}
	secretVar, exists := (*resultPreview.EnvVars)["SECRET_KEY"]
	if !exists {
		t.Fatal("expected SECRET_KEY to exist")
	}
	if secretVar.Value.ValueString() != "my-secret-value" {
		t.Errorf("expected secret value to be preserved, got %q", secretVar.Value.ValueString())
	}

	// And non-secret values should come from destination
	if resultPreview.CompatibilityDate.ValueString() != "2024-01-18" {
		t.Errorf("expected compatibility_date from dest, got %q", resultPreview.CompatibilityDate.ValueString())
	}
}

// TestPreserveSecretEnvVars_HandlesNullSource tests null source handling.
func TestPreserveSecretEnvVars_HandlesNullSource(t *testing.T) {
	ctx := context.Background()

	source := customfield.NullObject[PagesProjectDeploymentConfigsModel](ctx)

	destPreview := &PagesProjectDeploymentConfigsPreviewModel{
		CompatibilityDate: types.StringValue("2024-01-17"),
	}
	destPreviewObj, _ := customfield.NewObject(ctx, destPreview)
	destConfigs := &PagesProjectDeploymentConfigsModel{
		Preview: destPreviewObj,
	}
	dest, _ := customfield.NewObject(ctx, destConfigs)

	result, diags := PreserveSecretEnvVars(ctx, source, dest)

	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if !result.Equal(dest) {
		t.Error("expected result to equal dest when source is null")
	}
}

// TestPreserveSecretEnvVars_HandlesNullDest tests null destination handling.
func TestPreserveSecretEnvVars_HandlesNullDest(t *testing.T) {
	ctx := context.Background()

	sourcePreview := &PagesProjectDeploymentConfigsPreviewModel{
		CompatibilityDate: types.StringValue("2024-01-17"),
	}
	sourcePreviewObj, _ := customfield.NewObject(ctx, sourcePreview)
	sourceConfigs := &PagesProjectDeploymentConfigsModel{
		Preview: sourcePreviewObj,
	}
	source, _ := customfield.NewObject(ctx, sourceConfigs)

	dest := customfield.NullObject[PagesProjectDeploymentConfigsModel](ctx)

	result, diags := PreserveSecretEnvVars(ctx, source, dest)

	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if !result.IsNull() {
		t.Error("expected result to be null when dest is null")
	}
}

// TestEnvVarsMatchIgnoringSecrets tests env var comparison.
func TestEnvVarsMatchIgnoringSecrets(t *testing.T) {
	tests := []struct {
		name     string
		source   *map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel
		dest     *map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel
		expected bool
	}{
		{
			name:     "both nil",
			source:   nil,
			dest:     nil,
			expected: true,
		},
		{
			name:   "source nil, dest empty",
			source: nil,
			dest:   &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{},
			expected: true,
		},
		{
			name: "plain_text values match",
			source: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
				"KEY": {Type: types.StringValue("plain_text"), Value: types.StringValue("value")},
			},
			dest: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
				"KEY": {Type: types.StringValue("plain_text"), Value: types.StringValue("value")},
			},
			expected: true,
		},
		{
			name: "plain_text values differ",
			source: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
				"KEY": {Type: types.StringValue("plain_text"), Value: types.StringValue("value1")},
			},
			dest: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
				"KEY": {Type: types.StringValue("plain_text"), Value: types.StringValue("value2")},
			},
			expected: false,
		},
		{
			name: "secret_text values differ but should be ignored",
			source: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
				"SECRET": {Type: types.StringValue("secret_text"), Value: types.StringValue("secret-value")},
			},
			dest: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
				"SECRET": {Type: types.StringValue("secret_text"), Value: types.StringValue("")}, // API returns empty
			},
			expected: true,
		},
		{
			name: "mixed plain and secret, plain values match",
			source: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
				"PLAIN":  {Type: types.StringValue("plain_text"), Value: types.StringValue("plain-value")},
				"SECRET": {Type: types.StringValue("secret_text"), Value: types.StringValue("secret-value")},
			},
			dest: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
				"PLAIN":  {Type: types.StringValue("plain_text"), Value: types.StringValue("plain-value")},
				"SECRET": {Type: types.StringValue("secret_text"), Value: types.StringValue("")},
			},
			expected: true,
		},
		{
			name: "different keys",
			source: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
				"KEY1": {Type: types.StringValue("plain_text"), Value: types.StringValue("value")},
			},
			dest: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
				"KEY2": {Type: types.StringValue("plain_text"), Value: types.StringValue("value")},
			},
			expected: false,
		},
		{
			name: "type mismatch",
			source: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
				"KEY": {Type: types.StringValue("plain_text"), Value: types.StringValue("value")},
			},
			dest: &map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel{
				"KEY": {Type: types.StringValue("secret_text"), Value: types.StringValue("value")},
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := envVarsMatchIgnoringSecrets(tc.source, tc.dest)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestCompatibilityFlagsEqual tests flag comparison.
func TestCompatibilityFlagsEqual(t *testing.T) {
	tests := []struct {
		name     string
		source   *[]types.String
		dest     *[]types.String
		expected bool
	}{
		{
			name:     "both nil",
			source:   nil,
			dest:     nil,
			expected: true,
		},
		{
			name:     "source nil, dest empty",
			source:   nil,
			dest:     &[]types.String{},
			expected: true,
		},
		{
			name:     "source empty, dest nil",
			source:   &[]types.String{},
			dest:     nil,
			expected: true,
		},
		{
			name:     "same values",
			source:   &[]types.String{types.StringValue("flag1"), types.StringValue("flag2")},
			dest:     &[]types.String{types.StringValue("flag1"), types.StringValue("flag2")},
			expected: true,
		},
		{
			name:     "different values",
			source:   &[]types.String{types.StringValue("flag1")},
			dest:     &[]types.String{types.StringValue("flag2")},
			expected: false,
		},
		{
			name:     "different lengths",
			source:   &[]types.String{types.StringValue("flag1")},
			dest:     &[]types.String{types.StringValue("flag1"), types.StringValue("flag2")},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := compatibilityFlagsEqual(tc.source, tc.dest)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

