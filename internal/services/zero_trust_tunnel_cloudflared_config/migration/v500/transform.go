package v500

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (v4 SDKv2) state to target (v5 Plugin Framework) state.
// This function handles all field transformations, type conversions, and nested structure migrations.
func Transform(ctx context.Context, source SourceV4TunnelConfigModel) (*TargetV5TunnelConfigModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Initialize target with direct copies
	target := &TargetV5TunnelConfigModel{
		ID:        source.ID,
		TunnelID:  source.TunnelID,
		AccountID: source.AccountID,
	}

	// Set v5-only computed fields to null (will be populated by API on refresh)
	target.CreatedAt = timetypes.NewRFC3339Null()
	target.Source = types.StringNull()
	target.Version = types.Int64Null()

	// Transform config: array[0] → pointer
	if len(source.Config) > 0 {
		configTarget, configDiags := transformConfig(ctx, source.Config[0])
		diags.Append(configDiags...)
		if diags.HasError() {
			return nil, diags
		}
		target.Config = configTarget
	} else {
		target.Config = nil
	}

	return target, diags
}

// transformConfig converts SourceV4ConfigModel (TypeList MaxItems:1 element) to TargetV5ConfigModel (pointer).
func transformConfig(ctx context.Context, source SourceV4ConfigModel) (*TargetV5ConfigModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5ConfigModel{}

	// origin_request: TypeList MaxItems:1 → SingleNestedAttribute (pointer)
	// warp_routing is dropped (removed in v5)
	// Guard against v4 storing an unset TypeList MaxItems:1 as [{}] (one all-null element).
	if len(source.OriginRequest) > 0 && !isOriginRequestEmpty(source.OriginRequest[0]) {
		originReq, orDiags := transformConfigOriginRequest(ctx, source.OriginRequest[0])
		diags.Append(orDiags...)
		if diags.HasError() {
			return nil, diags
		}
		target.OriginRequest = originReq
	} else {
		target.OriginRequest = nil
	}

	// ingress_rule → ingress (renamed; same list structure)
	ingress, ingressDiags := transformIngressRules(ctx, source.IngressRule)
	diags.Append(ingressDiags...)
	if diags.HasError() {
		return nil, diags
	}
	target.Ingress = ingress

	return target, diags
}

// transformIngressRules converts []SourceV4IngressRuleModel → *[]*TargetV5IngressModel.
func transformIngressRules(ctx context.Context, source []SourceV4IngressRuleModel) (*[]*TargetV5IngressModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(source) == 0 {
		return nil, diags
	}

	result := make([]*TargetV5IngressModel, 0, len(source))
	for _, srcIngress := range source {
		targetIngress := &TargetV5IngressModel{
			Hostname: srcIngress.Hostname,
			Service:  srcIngress.Service,
			Path:     srcIngress.Path,
		}

		// origin_request per ingress: TypeList MaxItems:1 → pointer
		// Guard against v4 storing an unset TypeList MaxItems:1 as [{}] (one all-null element).
		if len(srcIngress.OriginRequest) > 0 && !isOriginRequestEmpty(srcIngress.OriginRequest[0]) {
			ingressOR, orDiags := transformIngressOriginRequest(ctx, srcIngress.OriginRequest[0])
			diags.Append(orDiags...)
			if diags.HasError() {
				return nil, diags
			}
			targetIngress.OriginRequest = ingressOR
		} else {
			targetIngress.OriginRequest = nil
		}

		result = append(result, targetIngress)
	}

	return &result, diags
}

// transformConfigOriginRequest converts a config-level SourceV4OriginRequestModel to TargetV5OriginRequestModel.
func transformConfigOriginRequest(ctx context.Context, source SourceV4OriginRequestModel) (*TargetV5OriginRequestModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5OriginRequestModel{}

	// Duration fields: string ("30s") → Int64 (seconds)
	target.ConnectTimeout = parseDurationToInt64(source.ConnectTimeout)
	target.TLSTimeout = parseDurationToInt64(source.TLSTimeout)
	target.TCPKeepAlive = parseDurationToInt64(source.TCPKeepAlive)
	target.KeepAliveTimeout = parseDurationToInt64(source.KeepAliveTimeout)

	// Direct copies
	target.KeepAliveConnections = source.KeepAliveConnections
	target.NoHappyEyeballs = source.NoHappyEyeballs
	target.NoTLSVerify = source.NoTLSVerify
	target.DisableChunkedEncoding = source.DisableChunkedEncoding
	target.HTTP2Origin = source.HTTP2Origin
	target.HTTPHostHeader = source.HTTPHostHeader
	target.OriginServerName = source.OriginServerName
	target.CAPool = source.CAPool
	target.ProxyType = source.ProxyType

	// v5-only field: set to null
	target.MatchSnItoHost = types.BoolNull()

	// Dropped: bastion_mode, proxy_address, proxy_port, ip_rules (not copied)

	// access: TypeList MaxItems:1 → SingleNestedAttribute (pointer)
	// Drop access if missing aud_tag or team_name
	accessTarget, accessDiags := transformAccess(ctx, source.Access)
	diags.Append(accessDiags...)
	if diags.HasError() {
		return nil, diags
	}
	target.Access = accessTarget

	return target, diags
}

// transformIngressOriginRequest converts an ingress-level SourceV4OriginRequestModel to TargetV5IngressOriginRequestModel.
func transformIngressOriginRequest(ctx context.Context, source SourceV4OriginRequestModel) (*TargetV5IngressOriginRequestModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5IngressOriginRequestModel{}

	// Duration fields: string ("30s") → Int64 (seconds)
	target.ConnectTimeout = parseDurationToInt64(source.ConnectTimeout)
	target.TLSTimeout = parseDurationToInt64(source.TLSTimeout)
	target.TCPKeepAlive = parseDurationToInt64(source.TCPKeepAlive)
	target.KeepAliveTimeout = parseDurationToInt64(source.KeepAliveTimeout)

	// Direct copies
	target.KeepAliveConnections = source.KeepAliveConnections
	target.NoHappyEyeballs = source.NoHappyEyeballs
	target.NoTLSVerify = source.NoTLSVerify
	target.DisableChunkedEncoding = source.DisableChunkedEncoding
	target.HTTP2Origin = source.HTTP2Origin
	target.HTTPHostHeader = source.HTTPHostHeader
	target.OriginServerName = source.OriginServerName
	target.CAPool = source.CAPool
	target.ProxyType = source.ProxyType

	// v5-only field: set to null
	target.MatchSnItoHost = types.BoolNull()

	// Dropped: bastion_mode, proxy_address, proxy_port, ip_rules (not copied)

	// access: TypeList MaxItems:1 → SingleNestedAttribute (pointer)
	accessTarget, accessDiags := transformAccess(ctx, source.Access)
	diags.Append(accessDiags...)
	if diags.HasError() {
		return nil, diags
	}
	target.Access = accessTarget

	return target, diags
}

// transformAccess converts a []SourceV4AccessModel (TypeList MaxItems:1) to *TargetV5AccessModel.
// Returns nil if access is empty or if aud_tag / team_name are missing (as per tf-migrate logic).
func transformAccess(ctx context.Context, source []SourceV4AccessModel) (*TargetV5AccessModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(source) == 0 {
		return nil, diags
	}

	src := source[0]

	// Drop access if team_name is missing/empty
	if src.TeamName.IsNull() || src.TeamName.IsUnknown() || src.TeamName.ValueString() == "" {
		return nil, diags
	}

	// Drop access if aud_tag is missing/null/empty
	if src.AUDTag.IsNull() || src.AUDTag.IsUnknown() {
		return nil, diags
	}

	target := &TargetV5AccessModel{
		Required: src.Required,
		TeamName: src.TeamName,
	}

	// aud_tag: TypeSet → *[]types.String
	audTags, audDiags := convertSetToStringSlice(ctx, src.AUDTag)
	diags.Append(audDiags...)
	if diags.HasError() {
		return nil, diags
	}

	// Drop access if aud_tag slice is empty after conversion
	if len(audTags) == 0 {
		return nil, diags
	}

	target.AUDTag = &audTags

	return target, diags
}

// ============================================================================
// Helper Functions
// ============================================================================

// isOriginRequestEmpty reports whether a v4 origin_request element has all-null/unknown/zero fields.
// In SDKv2, an unset TypeList MaxItems:1 block can be stored as [{}] — one element with every
// field at its zero value. Optional string fields may be stored as "" (empty string, not null),
// bool fields as null, and int fields as null. We treat such an element as "not configured"
// so that the migrated v5 state has origin_request = null, matching tf-migrate output for configs
// that never had an origin_request block.
//
// The nested access block (also TypeList MaxItems:1) can also be stored as [{}] inside an otherwise
// empty origin_request, so we check whether the access element itself is empty rather than just
// checking len(src.Access) == 0.
func isOriginRequestEmpty(src SourceV4OriginRequestModel) bool {
	isNullOrEmpty := func(s types.String) bool {
		return s.IsNull() || s.IsUnknown() || s.ValueString() == ""
	}
	accessEmpty := len(src.Access) == 0 ||
		(len(src.Access) == 1 && isAccessElementEmpty(src.Access[0]))
	return isNullOrEmpty(src.ConnectTimeout) &&
		isNullOrEmpty(src.TLSTimeout) &&
		isNullOrEmpty(src.TCPKeepAlive) &&
		isNullOrEmpty(src.KeepAliveTimeout) &&
		(src.KeepAliveConnections.IsNull() || src.KeepAliveConnections.IsUnknown()) &&
		(src.NoHappyEyeballs.IsNull() || src.NoHappyEyeballs.IsUnknown()) &&
		(src.NoTLSVerify.IsNull() || src.NoTLSVerify.IsUnknown()) &&
		(src.DisableChunkedEncoding.IsNull() || src.DisableChunkedEncoding.IsUnknown()) &&
		(src.HTTP2Origin.IsNull() || src.HTTP2Origin.IsUnknown()) &&
		isNullOrEmpty(src.HTTPHostHeader) &&
		isNullOrEmpty(src.OriginServerName) &&
		isNullOrEmpty(src.CAPool) &&
		isNullOrEmpty(src.ProxyType) &&
		accessEmpty
}

// isAccessElementEmpty reports whether a v4 access element has no meaningful configuration.
// An access element is meaningless without both team_name and aud_tag, which together identify
// the Access app to validate against. SDKv2 may store an unset access TypeList MaxItems:1 as
// [{}] with an all-zero element; this function detects that case.
func isAccessElementEmpty(src SourceV4AccessModel) bool {
	teamNameEmpty := src.TeamName.IsNull() || src.TeamName.IsUnknown() || src.TeamName.ValueString() == ""
	audTagEmpty := src.AUDTag.IsNull() || src.AUDTag.IsUnknown() || len(src.AUDTag.Elements()) == 0
	return teamNameEmpty && audTagEmpty
}

// parseDurationToInt64 converts a Go duration string (e.g., "30s", "1m30s") to Int64 seconds.
// Returns null Int64 if the value is null/unknown/empty or cannot be parsed.
func parseDurationToInt64(val types.String) types.Int64 {
	if val.IsNull() || val.IsUnknown() {
		return types.Int64Null()
	}

	s := strings.TrimSpace(val.ValueString())
	if s == "" {
		return types.Int64Null()
	}

	d, err := time.ParseDuration(s)
	if err != nil {
		// Cannot parse; try appending "s" for bare numbers
		d, err = time.ParseDuration(fmt.Sprintf("%ss", s))
		if err != nil {
			return types.Int64Null()
		}
	}

	return types.Int64Value(int64(d.Seconds()))
}

// convertSetToStringSlice converts types.Set to []types.String for Set[String] → List[String] conversions.
// Extracts to []string first, then converts to []types.String to avoid attr.Value issues.
func convertSetToStringSlice(ctx context.Context, set types.Set) ([]types.String, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Extract to []string first
	var rawStrings []string
	diags.Append(set.ElementsAs(ctx, &rawStrings, false)...)
	if diags.HasError() {
		return nil, diags
	}

	// Convert []string to []types.String
	result := make([]types.String, 0, len(rawStrings))
	for _, str := range rawStrings {
		result = append(result, types.StringValue(str))
	}
	return result, diags
}
