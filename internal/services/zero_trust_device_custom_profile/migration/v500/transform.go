// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"context"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TransformToCustomProfile transforms a v4 device profile state to v5 custom profile structure.
//
// Key transformations:
// - Remove fields: default, enabled (not in custom profile or computed)
// - Remove: fallback_domains (migrated to separate resource)
// - Extract policy_id: From composite ID "account_id/policy_id" → policy_id attribute
// - Type conversions: Int64 → Float64 (auto_connect, captive_portal, precedence, service_mode_v2.port)
// - Structure change: Flatten service_mode_v2 → Nested object
// - Keep fields: name, description, match, precedence (required/optional in custom profile)
func TransformToCustomProfile(ctx context.Context, source SourceDeviceProfileModel) (*TargetCustomProfileModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	target := &TargetCustomProfileModel{}

	// Direct field copies (no transformation)
	target.ID = source.ID
	target.AccountID = source.AccountID
	target.Name = source.Name
	target.Description = source.Description
	target.Match = source.Match
	target.LANAllowMinutes = source.LANAllowMinutes
	target.LANAllowSubnetSize = source.LANAllowSubnetSize
	target.AllowModeSwitch = source.AllowModeSwitch
	target.AllowUpdates = source.AllowUpdates
	target.AllowedToLeave = source.AllowedToLeave
	target.DisableAutoFallback = source.DisableAutoFallback
	target.ExcludeOfficeIPs = source.ExcludeOfficeIPs
	target.SupportURL = source.SupportURL
	target.SwitchLocked = source.SwitchLocked
	target.TunnelProtocol = source.TunnelProtocol

	// Extract policy_id from composite ID
	// v4 ID format: "account_id/policy_id"
	// v5: policy_id as separate attribute
	if !source.ID.IsNull() {
		idStr := source.ID.ValueString()
		if slashIdx := strings.Index(idStr, "/"); slashIdx != -1 && slashIdx < len(idStr)-1 {
			policyID := idStr[slashIdx+1:]
			target.PolicyID = types.StringValue(policyID)
		} else {
			// ID doesn't have expected format, use as-is
			target.PolicyID = source.ID
		}
	} else {
		target.PolicyID = types.StringNull()
	}

	// Type conversions: Int64 → Float64
	if !source.AutoConnect.IsNull() {
		target.AutoConnect = types.Float64Value(float64(source.AutoConnect.ValueInt64()))
	} else {
		target.AutoConnect = types.Float64Null()
	}

	if !source.CaptivePortal.IsNull() {
		target.CaptivePortal = types.Float64Value(float64(source.CaptivePortal.ValueInt64()))
	} else {
		target.CaptivePortal = types.Float64Null()
	}

	// Precedence: Int64 → Float64
	if !source.Precedence.IsNull() {
		target.Precedence = types.Float64Value(float64(source.Precedence.ValueInt64()))
	} else {
		target.Precedence = types.Float64Null()
	}

	// Transform service_mode_v2: Flatten → Nested object
	// v4: service_mode_v2_mode (string) + service_mode_v2_port (int64)
	// v5: service_mode_v2 { mode (string), port (float64) }
	// Edge case: v4 default is mode="warp" with no port - don't create nested object for this
	hasMode := !source.ServiceModeV2Mode.IsNull() && source.ServiceModeV2Mode.ValueString() != ""
	hasPort := !source.ServiceModeV2Port.IsNull() && source.ServiceModeV2Port.ValueInt64() != 0
	isV4Default := hasMode && source.ServiceModeV2Mode.ValueString() == "warp" && !hasPort

	if isV4Default {
		// Don't migrate v4 default value - let v5 use its own defaults
		target.ServiceModeV2 = customfield.NullObject[TargetCustomProfileServiceModeV2Model](ctx)
	} else if hasMode || hasPort {
		// Create nested object with transformed values
		serviceModeV2 := &TargetCustomProfileServiceModeV2Model{}
		if hasMode {
			serviceModeV2.Mode = source.ServiceModeV2Mode
		} else {
			serviceModeV2.Mode = types.StringNull()
		}
		if hasPort {
			serviceModeV2.Port = types.Float64Value(float64(source.ServiceModeV2Port.ValueInt64()))
		} else {
			serviceModeV2.Port = types.Float64Null()
		}
		nestedObj, objDiags := customfield.NewObject(ctx, serviceModeV2)
		diags.Append(objDiags...)
		target.ServiceModeV2 = nestedObj
	} else {
		// No service_mode_v2 in v4 - use null
		target.ServiceModeV2 = customfield.NullObject[TargetCustomProfileServiceModeV2Model](ctx)
	}

	// Enabled field: Optional+Computed in v5 custom profile (not removed like in default)
	// Copy from v4 if present
	target.Enabled = source.Enabled

	// Computed fields - set to null, API will populate
	target.Default = types.BoolNull()
	target.GatewayUniqueID = types.StringNull()

	// Optional+Computed fields that didn't exist in v4 - set to null, API will populate
	target.Exclude = customfield.NullObjectList[TargetCustomProfileExcludeModel](ctx)
	target.Include = customfield.NullObjectList[TargetCustomProfileIncludeModel](ctx)
	target.FallbackDomains = customfield.NullObjectList[TargetCustomProfileFallbackDomainsModel](ctx)
	target.TargetTests = customfield.NullObjectList[TargetCustomProfileTargetTestsModel](ctx)
	target.RegisterInterfaceIPWithDNS = types.BoolNull()
	target.SccmVpnBoundarySupport = types.BoolNull()

	// Fields explicitly NOT copied (removed in v5 custom profile):
	// - default (computed in v5, always false for custom)
	// - fallback_domains (migrated to separate resource)

	return target, diags
}
