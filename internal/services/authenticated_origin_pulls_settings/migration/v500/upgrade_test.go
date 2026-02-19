package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/authenticated_origin_pulls_settings/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpgradeFromLegacyV0(t *testing.T) {
	ctx := context.Background()
	sourceSchema := v500.SourceCloudflareAuthenticatedOriginPullsSchema()

	tests := []struct {
		name           string
		sourceState    map[string]interface{}
		expectedTarget v500.TargetAuthenticatedOriginPullsSettingsModel
		expectError    bool
	}{
		{
			name: "upgrade with all fields",
			sourceState: map[string]interface{}{
				"id":                                     "zone-123",
				"zone_id":                                "zone-123",
				"hostname":                               "example.com",
				"authenticated_origin_pulls_certificate": "cert-123",
				"enabled":                                true,
			},
			expectedTarget: v500.TargetAuthenticatedOriginPullsSettingsModel{
				ID:      types.StringValue("zone-123"),
				ZoneID:  types.StringValue("zone-123"),
				Enabled: types.BoolValue(true),
			},
			expectError: false,
		},
		{
			name: "upgrade with disabled state",
			sourceState: map[string]interface{}{
				"id":                                     "zone-456",
				"zone_id":                                "zone-456",
				"hostname":                               "test.example.com",
				"authenticated_origin_pulls_certificate": "cert-456",
				"enabled":                                false,
			},
			expectedTarget: v500.TargetAuthenticatedOriginPullsSettingsModel{
				ID:      types.StringValue("zone-456"),
				ZoneID:  types.StringValue("zone-456"),
				Enabled: types.BoolValue(false),
			},
			expectError: false,
		},
		{
			name: "upgrade with null optional fields",
			sourceState: map[string]interface{}{
				"id":                                     "zone-789",
				"zone_id":                                "zone-789",
				"hostname":                               nil,
				"authenticated_origin_pulls_certificate": nil,
				"enabled":                                true,
			},
			expectedTarget: v500.TargetAuthenticatedOriginPullsSettingsModel{
				ID:      types.StringValue("zone-789"),
				ZoneID:  types.StringValue("zone-789"),
				Enabled: types.BoolValue(true),
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create source state
			sourceStateValue := createTerraformValue(t, sourceSchema, tt.sourceState)
			sourceState := tfsdk.State{
				Raw:    sourceStateValue,
				Schema: sourceSchema,
			}

			// Create request and response
			req := resource.UpgradeStateRequest{
				State: &sourceState,
			}
			resp := &resource.UpgradeStateResponse{
				State: tfsdk.State{
					Schema: schema.Schema{
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Computed: true,
							},
							"zone_id": schema.StringAttribute{
								Required: true,
							},
							"enabled": schema.BoolAttribute{
								Required: true,
							},
						},
					},
				},
			}

			// Run upgrade
			v500.UpgradeFromLegacyV0(ctx, req, resp)

			if tt.expectError {
				require.True(t, resp.Diagnostics.HasError(), "Expected error but got none")
			} else {
				require.False(t, resp.Diagnostics.HasError(), "Unexpected error: %v", resp.Diagnostics)

				// Extract upgraded state
				var upgradedState v500.TargetAuthenticatedOriginPullsSettingsModel
				diags := resp.State.Get(ctx, &upgradedState)
				require.False(t, diags.HasError(), "Failed to extract upgraded state: %v", diags)

				// Verify fields
				assert.Equal(t, tt.expectedTarget.ID, upgradedState.ID, "ID mismatch")
				assert.Equal(t, tt.expectedTarget.ZoneID, upgradedState.ZoneID, "ZoneID mismatch")
				assert.Equal(t, tt.expectedTarget.Enabled, upgradedState.Enabled, "Enabled mismatch")
			}
		})
	}
}

func TestUpgradeFromV1(t *testing.T) {
	ctx := context.Background()

	// V1 and V500 have identical schemas, so this is a no-op
	targetSchema := schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"enabled": schema.BoolAttribute{
				Required: true,
			},
		},
	}

	tests := []struct {
		name        string
		sourceState map[string]interface{}
		expectError bool
	}{
		{
			name: "no-op upgrade preserves state",
			sourceState: map[string]interface{}{
				"id":      "zone-123",
				"zone_id": "zone-123",
				"enabled": true,
			},
			expectError: false,
		},
		{
			name: "no-op upgrade with disabled state",
			sourceState: map[string]interface{}{
				"id":      "zone-456",
				"zone_id": "zone-456",
				"enabled": false,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create source state
			sourceStateValue := createTerraformValue(t, targetSchema, tt.sourceState)
			sourceState := tfsdk.State{
				Raw:    sourceStateValue,
				Schema: targetSchema,
			}

			// Create request and response
			req := resource.UpgradeStateRequest{
				State: &sourceState,
			}
			resp := &resource.UpgradeStateResponse{
				State: tfsdk.State{
					Schema: targetSchema,
				},
			}

			// Run upgrade
			v500.UpgradeFromV1(ctx, req, resp)

			if tt.expectError {
				require.True(t, resp.Diagnostics.HasError(), "Expected error but got none")
			} else {
				require.False(t, resp.Diagnostics.HasError(), "Unexpected error: %v", resp.Diagnostics)

				// Extract upgraded state
				var upgradedState v500.TargetAuthenticatedOriginPullsSettingsModel
				diags := resp.State.Get(ctx, &upgradedState)
				require.False(t, diags.HasError(), "Failed to extract upgraded state: %v", diags)

				// Verify state unchanged (no-op)
				assert.Equal(t, types.StringValue(tt.sourceState["id"].(string)), upgradedState.ID)
				assert.Equal(t, types.StringValue(tt.sourceState["zone_id"].(string)), upgradedState.ZoneID)
				assert.Equal(t, types.BoolValue(tt.sourceState["enabled"].(bool)), upgradedState.Enabled)
			}
		})
	}
}

// createTerraformValue creates a tftypes.Value from a schema and map of values
func createTerraformValue(t *testing.T, s schema.Schema, values map[string]interface{}) tftypes.Value {
	t.Helper()

	// Build the tftypes.Type from schema
	attrTypes := make(map[string]tftypes.Type)
	for name := range s.Attributes {
		switch s.Attributes[name].(type) {
		case schema.StringAttribute:
			attrTypes[name] = tftypes.String
		case schema.BoolAttribute:
			attrTypes[name] = tftypes.Bool
		default:
			t.Fatalf("Unsupported attribute type for %s", name)
		}
	}

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	// Build the value map
	valueMap := make(map[string]tftypes.Value)
	for name, val := range values {
		if val == nil {
			valueMap[name] = tftypes.NewValue(attrTypes[name], nil)
		} else {
			valueMap[name] = tftypes.NewValue(attrTypes[name], val)
		}
	}

	return tftypes.NewValue(objectType, valueMap)
}
