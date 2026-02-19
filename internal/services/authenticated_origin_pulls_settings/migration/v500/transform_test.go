package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/authenticated_origin_pulls_settings/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransform(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		source         v500.SourceCloudflareAuthenticatedOriginPullsModel
		expectedTarget v500.TargetAuthenticatedOriginPullsSettingsModel
		expectError    bool
	}{
		{
			name: "basic transformation - all fields present",
			source: v500.SourceCloudflareAuthenticatedOriginPullsModel{
				ID:                                   types.StringValue("zone-123"),
				ZoneID:                              types.StringValue("zone-123"),
				Hostname:                            types.StringValue("example.com"),
				AuthenticatedOriginPullsCertificate: types.StringValue("cert-123"),
				Enabled:                             types.BoolValue(true),
			},
			expectedTarget: v500.TargetAuthenticatedOriginPullsSettingsModel{
				ID:      types.StringValue("zone-123"),
				ZoneID:  types.StringValue("zone-123"),
				Enabled: types.BoolValue(true),
			},
			expectError: false,
		},
		{
			name: "disabled state",
			source: v500.SourceCloudflareAuthenticatedOriginPullsModel{
				ID:                                   types.StringValue("zone-456"),
				ZoneID:                              types.StringValue("zone-456"),
				Hostname:                            types.StringValue("test.example.com"),
				AuthenticatedOriginPullsCertificate: types.StringValue("cert-456"),
				Enabled:                             types.BoolValue(false),
			},
			expectedTarget: v500.TargetAuthenticatedOriginPullsSettingsModel{
				ID:      types.StringValue("zone-456"),
				ZoneID:  types.StringValue("zone-456"),
				Enabled: types.BoolValue(false),
			},
			expectError: false,
		},
		{
			name: "null optional fields",
			source: v500.SourceCloudflareAuthenticatedOriginPullsModel{
				ID:                                   types.StringValue("zone-789"),
				ZoneID:                              types.StringValue("zone-789"),
				Hostname:                            types.StringNull(),
				AuthenticatedOriginPullsCertificate: types.StringNull(),
				Enabled:                             types.BoolValue(true),
			},
			expectedTarget: v500.TargetAuthenticatedOriginPullsSettingsModel{
				ID:      types.StringValue("zone-789"),
				ZoneID:  types.StringValue("zone-789"),
				Enabled: types.BoolValue(true),
			},
			expectError: false,
		},
		{
			name: "unknown values",
			source: v500.SourceCloudflareAuthenticatedOriginPullsModel{
				ID:                                   types.StringUnknown(),
				ZoneID:                              types.StringValue("zone-999"),
				Hostname:                            types.StringUnknown(),
				AuthenticatedOriginPullsCertificate: types.StringUnknown(),
				Enabled:                             types.BoolUnknown(),
			},
			expectedTarget: v500.TargetAuthenticatedOriginPullsSettingsModel{
				ID:      types.StringUnknown(),
				ZoneID:  types.StringValue("zone-999"),
				Enabled: types.BoolUnknown(),
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			target, diags := v500.Transform(ctx, tt.source)

			if tt.expectError {
				require.True(t, diags.HasError(), "Expected error but got none")
			} else {
				require.False(t, diags.HasError(), "Unexpected error: %v", diags)

				// Verify all fields are correctly mapped
				assert.Equal(t, tt.expectedTarget.ID, target.ID, "ID mismatch")
				assert.Equal(t, tt.expectedTarget.ZoneID, target.ZoneID, "ZoneID mismatch")
				assert.Equal(t, tt.expectedTarget.Enabled, target.Enabled, "Enabled mismatch")
			}
		})
	}
}

func TestTransform_FieldsDropped(t *testing.T) {
	ctx := context.Background()

	source := v500.SourceCloudflareAuthenticatedOriginPullsModel{
		ID:                                   types.StringValue("zone-123"),
		ZoneID:                              types.StringValue("zone-123"),
		Hostname:                            types.StringValue("should-be-dropped.example.com"),
		AuthenticatedOriginPullsCertificate: types.StringValue("cert-should-be-dropped"),
		Enabled:                             types.BoolValue(true),
	}

	target, diags := v500.Transform(ctx, source)
	require.False(t, diags.HasError())

	// Verify that hostname and authenticated_origin_pulls_certificate are not in target
	// (They should not exist in the target model at all - this is validated by the type system)
	assert.Equal(t, types.StringValue("zone-123"), target.ID)
	assert.Equal(t, types.StringValue("zone-123"), target.ZoneID)
	assert.Equal(t, types.BoolValue(true), target.Enabled)
}
