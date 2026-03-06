// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TransformToDefaultProfile transforms a v4 device profile state to v5 default profile structure.
//
// Key transformations:
// - Remove fields: default, enabled, name, description, match, precedence (not in default profile)
// - Remove: fallback_domains (migrated to separate resource)
// - Type conversions: Int64 → Float64 (auto_connect, captive_portal, service_mode_v2.port)
// - Structure change: Flatten service_mode_v2 → Nested object
// - Add fields: register_interface_ip_with_dns, sccm_vpn_boundary_support (with defaults)
func TransformToDefaultProfile(ctx context.Context, source SourceDeviceProfileModel) (*TargetDefaultProfileModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	target := &TargetDefaultProfileModel{}

	// Direct field copies (no transformation)
	target.ID = source.ID
	target.AccountID = source.AccountID
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

	// Transform service_mode_v2: Flatten → Nested object
	// v4: service_mode_v2_mode (string) + service_mode_v2_port (int64)
	// v5: service_mode_v2 { mode (string), port (float64) }
	// Edge case: v4 default is mode="warp" with no port - don't create nested object for this
	hasMode := !source.ServiceModeV2Mode.IsNull() && source.ServiceModeV2Mode.ValueString() != ""
	hasPort := !source.ServiceModeV2Port.IsNull() && source.ServiceModeV2Port.ValueInt64() != 0
	isV4Default := hasMode && source.ServiceModeV2Mode.ValueString() == "warp" && !hasPort

	if isV4Default {
		// Don't migrate v4 default value - let v5 use its own defaults
		target.ServiceModeV2 = customfield.NullObject[TargetDefaultProfileServiceModeV2Model](ctx)
	} else if hasMode || hasPort {
		// Create nested object with transformed values
		serviceModeV2 := &TargetDefaultProfileServiceModeV2Model{}
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
		target.ServiceModeV2 = customfield.NullObject[TargetDefaultProfileServiceModeV2Model](ctx)
	}

	// Add v5 default profile specific fields with default values
	// These fields don't exist in v4 but are required in v5
	target.RegisterInterfaceIPWithDNS = types.BoolValue(true)  // v5 default
	target.SccmVpnBoundarySupport = types.BoolValue(false)     // v5 default

	// Computed fields - set to null, API will populate
	target.Default = types.BoolNull()
	target.Enabled = types.BoolNull()
	target.GatewayUniqueID = types.StringNull()

	// Optional+Computed fields that didn't exist in v4 - set to null, API will populate
	target.Exclude = customfield.NullObjectList[TargetDefaultProfileExcludeModel](ctx)
	target.Include = customfield.NullObjectList[TargetDefaultProfileIncludeModel](ctx)
	target.FallbackDomains = customfield.NullObjectList[TargetDefaultProfileFallbackDomainsModel](ctx)

	// Fields explicitly NOT copied (removed in v5 default profile):
	// - default (computed in v5)
	// - enabled (computed in v5)
	// - name (not in default profile)
	// - description (not in default profile)
	// - match (not in default profile)
	// - precedence (not in default profile)
	// - fallback_domains (migrated to separate resource)

	return target, diags
}
