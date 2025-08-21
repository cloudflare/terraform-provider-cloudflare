package zero_trust_device_custom_profile

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMarshalJSONForUpdate_ProblemScenario reproduces the exact scenario from GitHub issue #5514
func TestMarshalJSONForUpdate_ProblemScenario(t *testing.T) {
	// Plan: user only specifies required fields, computed_optional fields are null
	plan := ZeroTrustDeviceCustomProfileModel{
		AccountID:   types.StringValue("test-account-id"),
		Match:       types.StringValue(`identity.email == "test@cloudflare.com"`),
		Name:        types.StringValue("Allow Developers"),
		Precedence:  types.Float64Value(100),
		Description: types.StringNull(), // This was causing the problem in the logs!
		// All other computed_optional fields are null (user didn't specify them)
		AllowModeSwitch:            types.BoolNull(),
		AllowUpdates:               types.BoolNull(),
		AllowedToLeave:             types.BoolNull(),
		AutoConnect:                types.Float64Null(),
		CaptivePortal:              types.Float64Null(),
		DisableAutoFallback:        types.BoolNull(),
		ExcludeOfficeIPs:           types.BoolNull(),
		RegisterInterfaceIPWithDNS: types.BoolNull(),
		SccmVpnBoundarySupport:     types.BoolNull(),
		SupportURL:                 types.StringNull(),
		SwitchLocked:               types.BoolNull(),
		TunnelProtocol:             types.StringNull(),
		Enabled:                    types.BoolNull(),
		FallbackDomains:            customfield.NullObjectList[ZeroTrustDeviceCustomProfileFallbackDomainsModel](context.TODO()),
		TargetTests:                customfield.NullObjectList[ZeroTrustDeviceCustomProfileTargetTestsModel](context.TODO()),
		Default:                    types.BoolNull(),
		GatewayUniqueID:            types.StringNull(),
		PolicyID:                   types.StringNull(),
		ID:                         types.StringNull(),
		Exclude:                    nil,
	}

	// State: API returned default values
	state := ZeroTrustDeviceCustomProfileModel{
		AccountID:                  types.StringValue("test-account-id"),
		Match:                      types.StringValue(`identity.email == "test@cloudflare.com"`),
		Name:                       types.StringValue("Allow Developers"),
		Precedence:                 types.Float64Value(100),
		Description:                types.StringValue(""), // API returns empty string, not null
		Enabled:                    types.BoolValue(true),
		AllowModeSwitch:            types.BoolValue(false),
		AllowUpdates:               types.BoolValue(false),
		AllowedToLeave:             types.BoolValue(true),
		AutoConnect:                types.Float64Value(0),
		CaptivePortal:              types.Float64Value(180),
		DisableAutoFallback:        types.BoolValue(false),
		ExcludeOfficeIPs:           types.BoolValue(false),
		RegisterInterfaceIPWithDNS: types.BoolValue(true),
		SccmVpnBoundarySupport:     types.BoolValue(false),
		SupportURL:                 types.StringValue(""),
		SwitchLocked:               types.BoolValue(false),
		TunnelProtocol:             types.StringValue(""),
		FallbackDomains:            customfield.NullObjectList[ZeroTrustDeviceCustomProfileFallbackDomainsModel](context.TODO()),
		TargetTests:                customfield.NullObjectList[ZeroTrustDeviceCustomProfileTargetTestsModel](context.TODO()),
		Default:                    types.BoolValue(false),
		GatewayUniqueID:            types.StringValue("test-gateway-id"),
		PolicyID:                   types.StringValue("test-policy-id"),
		ID:                         types.StringValue("test-policy-id"),
		Exclude:                    nil,
	}

	// This should NOT produce JSON with null values
	jsonBytes, err := plan.MarshalJSONForUpdate(state)
	require.NoError(t, err)

	jsonStr := string(jsonBytes)
	t.Logf("Generated JSON: '%s'", jsonStr)

	// The key test: we should NOT see null values for ANY field
	assert.NotContains(t, jsonStr, `null`, "JSON should not contain any null values")
	assert.NotContains(t, jsonStr, `"description":null`, "description should not be explicitly nullified")

	// An empty JSON (no changes) is perfectly acceptable and prevents API errors
	if jsonStr == "" {
		t.Log("Empty JSON indicates no changes needed - this prevents API errors")
	}
}
