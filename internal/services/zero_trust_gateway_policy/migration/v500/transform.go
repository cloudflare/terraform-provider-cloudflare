package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4 cloudflare_teams_rule) state to target (current v5 cloudflare_zero_trust_gateway_policy) state.
// This function is shared by both UpgradeFromV4 and MoveState handlers.
func Transform(ctx context.Context, source *SourceCloudflareTeamsRuleModel) (*TargetZeroTrustGatewayPolicyModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for zero_trust_gateway_policy migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"name is required for zero_trust_gateway_policy migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.Action.IsNull() || source.Action.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"action is required for zero_trust_gateway_policy migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Initialize target with direct copies
	target := &TargetZeroTrustGatewayPolicyModel{
		ID:            source.ID,
		AccountID:     source.AccountID,
		Name:          source.Name,
		Action:        source.Action,
		Enabled:       source.Enabled,
		Traffic:       source.Traffic,
		Identity:      source.Identity,
		DevicePosture: source.DevicePosture,
		Precedence:    normalizeGatewayPolicyPrecedence(source.Precedence),
		Version:       source.Version, // Direct copy - both are types.Int64
	}

	// Step 3: Handle description (Required in v4, Optional in v5)
	// Note: v4 required this field, so it should always be present
	if !source.Description.IsNull() && !source.Description.IsUnknown() {
		target.Description = source.Description
	} else {
		// If somehow missing, set to empty string (v4 default)
		target.Description = types.StringValue("")
	}

	// Step 4: Handle filters list conversion (types.List → *[]types.String)
	if !source.Filters.IsNull() && !source.Filters.IsUnknown() {
		// Convert types.List to *[]types.String
		elements := source.Filters.Elements()
		filterSlice := make([]types.String, 0, len(elements))
		for _, elem := range elements {
			if str, ok := elem.(types.String); ok {
				filterSlice = append(filterSlice, str)
			}
		}
		target.Filters = &filterSlice
	} else {
		target.Filters = nil
	}

	// Step 5: Handle rule_settings transformation (array[0] → customfield.NestedObject)
	if len(source.RuleSettings) > 0 {
		rs, rsDiags := transformRuleSettings(ctx, &source.RuleSettings[0])
		diags.Append(rsDiags...)
		if !diags.HasError() {
			target.RuleSettings, _ = customfield.NewObject(ctx, rs)
		}
	} else {
		target.RuleSettings = customfield.NullObject[TargetRuleSettingsModel](ctx)
	}

	// Step 6: Set computed fields to null (will refresh from API)
	// These fields don't exist in v4 and will be populated by Read()
	target.CreatedAt = timetypes.NewRFC3339Null()
	target.UpdatedAt = timetypes.NewRFC3339Null()
	target.DeletedAt = timetypes.NewRFC3339Null()
	target.ReadOnly = types.BoolNull()
	target.Sharable = types.BoolNull()
	target.SourceAccount = types.StringNull()
	target.Version = types.Int64Null()
	target.WarningStatus = types.StringNull()

	return target, diags
}

func normalizeGatewayPolicyPrecedence(v types.Int64) types.Int64 {
	if v.IsNull() || v.IsUnknown() {
		return v
	}

	value := v.ValueInt64()
	// v4 API-backed state may persist precedence with a three-digit suffix
	// (e.g. 400412 for configured 400). During migration we normalize to
	// user-configured precedence when this suffix encoding is detected.
	if value >= 10000 {
		return types.Int64Value(value / 1000)
	}

	return v
}

// transformRuleSettings converts v4 rule_settings (nested block) to v5 rule_settings (nested object).
func transformRuleSettings(ctx context.Context, source *SourceRuleSettingsModel) (*TargetRuleSettingsModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetRuleSettingsModel{
		// Direct scalar field copies
		BlockPageEnabled:                source.BlockPageEnabled,
		OverrideHost:                    source.OverrideHost,
		IPCategories:                    source.IPCategories,
		IgnoreCNAMECategoryMatches:      source.IgnoreCNAMECategoryMatches,
		AllowChildBypass:                source.AllowChildBypass,
		BypassParentRule:                normalizeGatewayPolicyFalseBoolToNull(source.BypassParentRule),
		InsecureDisableDNSSECValidation: source.InsecureDisableDNSSECValidation,
		ResolveDNSThroughCloudflare:     source.ResolveDNSThroughCloudflare,
	}

	// Field rename: block_page_reason → block_reason
	target.BlockReason = source.BlockPageReason

	// Handle override_ips conversion (types.List → customfield.List)
	if !source.OverrideIPs.IsNull() && !source.OverrideIPs.IsUnknown() {
		target.OverrideIPs = customfield.NewListMust[types.String](ctx, source.OverrideIPs.Elements())
	} else {
		target.OverrideIPs = customfield.NullList[types.String](ctx)
	}

	// Handle add_headers type conversion: map[string]types.String → map[string]*[]types.String
	if source.AddHeaders != nil && len(*source.AddHeaders) > 0 {
		v5Headers := make(map[string]*[]types.String)
		for key, value := range *source.AddHeaders {
			if !value.IsNull() && !value.IsUnknown() {
				// Wrap single string in array (v4: string, v5: []string)
				arr := []types.String{value}
				v5Headers[key] = &arr
			}
		}
		target.AddHeaders = &v5Headers
	} else {
		target.AddHeaders = nil
	}

	// Transform nested structures (array[0] → pointer)

	// audit_ssh
	if len(source.AuditSSH) > 0 {
		target.AuditSSH = transformAuditSSH(&source.AuditSSH[0])
	}

	// l4override
	if len(source.L4override) > 0 {
		target.L4override = transformL4override(&source.L4override[0])
	}

	// biso_admin_controls - special handling for deprecated v1 fields
	if len(source.BISOAdminControls) > 0 {
		bac, bacDiags := transformBISOAdminControls(&source.BISOAdminControls[0])
		diags.Append(bacDiags...)
		target.BISOAdminControls = bac
	}

	// check_session
	if len(source.CheckSession) > 0 {
		target.CheckSession = transformCheckSession(&source.CheckSession[0])
	}

	// egress
	if len(source.Egress) > 0 {
		target.Egress = transformEgress(&source.Egress[0])
	}

	// untrusted_cert
	if len(source.UntrustedCERT) > 0 {
		target.UntrustedCERT = transformUntrustedCERT(&source.UntrustedCERT[0])
	}

	// payload_log
	if len(source.PayloadLog) > 0 {
		target.PayloadLog = transformPayloadLog(&source.PayloadLog[0])
	}

	// notification_settings - has field rename
	if len(source.NotificationSettings) > 0 {
		target.NotificationSettings = transformNotificationSettings(&source.NotificationSettings[0])
	}

	// dns_resolvers - special: outer structure array[0]→pointer, but ipv4/ipv6 REMAIN as arrays
	if len(source.DNSResolvers) > 0 {
		dr, drDiags := transformDNSResolvers(&source.DNSResolvers[0])
		diags.Append(drDiags...)
		target.DNSResolvers = dr
	}

	// resolve_dns_internally
	if len(source.ResolveDNSInternally) > 0 {
		target.ResolveDNSInternally = transformResolveDNSInternally(&source.ResolveDNSInternally[0])
	}

	return target, diags
}

func normalizeGatewayPolicyFalseBoolToNull(v types.Bool) types.Bool {
	if v.IsNull() || v.IsUnknown() {
		return v
	}
	if !v.ValueBool() {
		return types.BoolNull()
	}
	return v
}

// Helper transformation functions for nested structures

func transformAuditSSH(source *SourceAuditSSHModel) *TargetAuditSSHModel {
	if source == nil {
		return nil
	}
	return &TargetAuditSSHModel{
		CommandLogging: source.CommandLogging,
	}
}

func transformL4override(source *SourceL4overrideModel) *TargetL4overrideModel {
	if source == nil {
		return nil
	}
	return &TargetL4overrideModel{
		IP:   source.IP,
		Port: source.Port,
	}
}

func transformBISOAdminControls(source *SourceBISOAdminControlsModel) (*TargetBISOAdminControlsModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if source == nil {
		return nil, diags
	}

	// Transform biso_admin_controls fields:
	// 1. Rename v1 bool fields from disable_* to shortened versions (dp, dd, dcp, dk, du)
	// 2. Keep v2 string fields unchanged (copy, download, keyboard, paste, printing, upload)
	// 3. Drop disable_clipboard_redirection (no v5 equivalent)
	target := &TargetBISOAdminControlsModel{
		Version: source.Version,

		// v1 fields: Rename from disable_* to shortened versions
		DP:  source.DisablePrinting,  // disable_printing → dp
		DCP: source.DisableCopyPaste, // disable_copy_paste → dcp
		DD:  source.DisableDownload,  // disable_download → dd
		DK:  source.DisableKeyboard,  // disable_keyboard → dk
		DU:  source.DisableUpload,    // disable_upload → du
		// disable_clipboard_redirection intentionally dropped (no v5 equivalent)

		// v2 fields: Copy unchanged
		Copy:     source.Copy,
		Download: source.Download,
		Keyboard: source.Keyboard,
		Paste:    source.Paste,
		Printing: source.Printing,
		Upload:   source.Upload,
	}

	return target, diags
}

func transformCheckSession(source *SourceCheckSessionModel) *TargetCheckSessionModel {
	if source == nil {
		return nil
	}
	return &TargetCheckSessionModel{
		Enforce:  source.Enforce,
		Duration: source.Duration,
	}
}

func transformEgress(source *SourceEgressModel) *TargetEgressModel {
	if source == nil {
		return nil
	}
	return &TargetEgressModel{
		IPV6:         source.IPV6,
		IPV4:         source.IPV4,
		IPV4Fallback: source.IPV4Fallback,
	}
}

func transformUntrustedCERT(source *SourceUntrustedCERTModel) *TargetUntrustedCERTModel {
	if source == nil {
		return nil
	}
	return &TargetUntrustedCERTModel{
		Action: source.Action,
	}
}

func transformPayloadLog(source *SourcePayloadLogModel) *TargetPayloadLogModel {
	if source == nil {
		return nil
	}
	return &TargetPayloadLogModel{
		Enabled: source.Enabled,
	}
}

func transformNotificationSettings(source *SourceNotificationSettingsModel) *TargetNotificationSettingsModel {
	if source == nil {
		return nil
	}

	// Field rename: message → msg
	return &TargetNotificationSettingsModel{
		Enabled:    source.Enabled,
		Msg:        source.Message, // RENAMED: message → msg
		SupportURL: source.SupportURL,
	}
}

func transformDNSResolvers(source *SourceDNSResolversModel) (*TargetDNSResolversModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if source == nil {
		return nil, diags
	}

	target := &TargetDNSResolversModel{}

	// CRITICAL: ipv4 and ipv6 remain as ARRAYS (not converted to pointers)
	// The dns_resolvers structure itself goes from array[0]→pointer,
	// but the nested ipv4/ipv6 fields stay as arrays

	if len(source.IPV4) > 0 {
		ipv4List := make([]*TargetDNSResolversIPV4Model, 0, len(source.IPV4))
		for _, v4 := range source.IPV4 {
			ipv4List = append(ipv4List, &TargetDNSResolversIPV4Model{
				IP:                         v4.IP,
				Port:                       v4.Port,
				VnetID:                     v4.VnetID,
				RouteThroughPrivateNetwork: v4.RouteThroughPrivateNetwork,
			})
		}
		target.IPV4 = &ipv4List
	}

	if len(source.IPV6) > 0 {
		ipv6List := make([]*TargetDNSResolversIPV6Model, 0, len(source.IPV6))
		for _, v6 := range source.IPV6 {
			ipv6List = append(ipv6List, &TargetDNSResolversIPV6Model{
				IP:                         v6.IP,
				Port:                       v6.Port,
				VnetID:                     v6.VnetID,
				RouteThroughPrivateNetwork: v6.RouteThroughPrivateNetwork,
			})
		}
		target.IPV6 = &ipv6List
	}

	return target, diags
}

func transformResolveDNSInternally(source *SourceResolveDNSInternallyModel) *TargetResolveDNSInternallyModel {
	if source == nil {
		return nil
	}
	return &TargetResolveDNSInternallyModel{
		ViewID:   source.ViewID,
		Fallback: source.Fallback,
	}
}
