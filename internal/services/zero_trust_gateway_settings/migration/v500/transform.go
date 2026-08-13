package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts a v4 SourceV4ZeroTrustGatewaySettingsModel to the v5
// TargetV5ZeroTrustGatewaySettingsModel.
//
// Key transformations:
//   - Flat booleans (activity_log_enabled, etc.) → settings.*
//   - TypeList MaxItems:1 blocks → pointer structs under settings.*
//   - antivirus.notification_settings[0].message → notification_settings.msg
//   - logging, proxy, ssh_session_log, payload_log → dropped
//   - New v5 computed fields (created_at, updated_at, id) → null
func Transform(ctx context.Context, source SourceV4ZeroTrustGatewaySettingsModel) (*TargetV5ZeroTrustGatewaySettingsModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5ZeroTrustGatewaySettingsModel{
		// id, created_at, updated_at are computed by v5 API; set null for migration
		ID:        types.StringNull(),
		AccountID: source.AccountID,
		CreatedAt: timetypes.NewRFC3339Null(),
		UpdatedAt: timetypes.NewRFC3339Null(),
	}

	settings, settingsDiags := transformSettings(ctx, source)
	diags.Append(settingsDiags...)
	if diags.HasError() {
		return nil, diags
	}
	target.Settings = settings

	return target, diags
}

// transformSettings converts all v4 settings fields into the v5 TargetV5SettingsModel.
func transformSettings(ctx context.Context, source SourceV4ZeroTrustGatewaySettingsModel) (*TargetV5SettingsModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	settings := &TargetV5SettingsModel{
		// New in v5; not in v4 → null
		HostSelector: nil,
		Inspection:   nil,
		Sandbox:      nil,
	}

	// Flat boolean: activity_log_enabled → settings.activity_log.enabled
	if !source.ActivityLogEnabled.IsNull() && !source.ActivityLogEnabled.IsUnknown() {
		settings.ActivityLog = &TargetV5ActivityLogModel{
			Enabled: source.ActivityLogEnabled,
		}
	}

	// Flat boolean: tls_decrypt_enabled → settings.tls_decrypt.enabled
	if !source.TLSDecryptEnabled.IsNull() && !source.TLSDecryptEnabled.IsUnknown() {
		settings.TLSDecrypt = &TargetV5TLSDecryptModel{
			Enabled: source.TLSDecryptEnabled,
		}
	}

	// Flat boolean: protocol_detection_enabled → settings.protocol_detection.enabled
	if !source.ProtocolDetectionEnabled.IsNull() && !source.ProtocolDetectionEnabled.IsUnknown() {
		settings.ProtocolDetection = &TargetV5ProtocolDetectionModel{
			Enabled: source.ProtocolDetectionEnabled,
		}
	}

	// Browser isolation: combine two flat fields → settings.browser_isolation
	// Both fields had Default: false in v4, so they're always present in v4 state.
	// We create the browser_isolation block when either field is non-null.
	urlEnabled := source.URLBrowserIsolationEnabled
	nonIdentityEnabled := source.NonIdentityBrowserIsolationEnabled
	if (!urlEnabled.IsNull() && !urlEnabled.IsUnknown()) ||
		(!nonIdentityEnabled.IsNull() && !nonIdentityEnabled.IsUnknown()) {
		settings.BrowserIsolation = &TargetV5BrowserIsolationModel{
			URLBrowserIsolationEnabled: urlEnabled,
			// Rename: non_identity_browser_isolation_enabled → non_identity_enabled
			NonIdentityEnabled: nonIdentityEnabled,
		}
	}

	// block_page: TypeList MaxItems:1 array → pointer struct
	if len(source.BlockPage) > 0 {
		bp := source.BlockPage[0]
		settings.BlockPage = &TargetV5BlockPageModel{
			BackgroundColor: bp.BackgroundColor,
			Enabled:         bp.Enabled,
			FooterText:      bp.FooterText,
			HeaderText:      bp.HeaderText,
			LogoPath:        bp.LogoPath,
			MailtoAddress:   bp.MailtoAddress,
			MailtoSubject:   bp.MailtoSubject,
			Name:            bp.Name,
			// New v5 fields not in v4 → null (will be populated on next refresh)
			IncludeContext: types.BoolNull(),
			Mode:           types.StringNull(),
			ReadOnly:       types.BoolNull(),
			SourceAccount:  types.StringNull(),
			SuppressFooter: types.BoolNull(),
			TargetURI:      types.StringNull(),
			Version:        types.Int64Null(),
		}
	}

	// body_scanning: TypeList MaxItems:1 array → pointer struct
	if len(source.BodyScanning) > 0 {
		bs := source.BodyScanning[0]
		settings.BodyScanning = &TargetV5BodyScanningModel{
			InspectionMode: bs.InspectionMode,
		}
	}

	// fips: TypeList MaxItems:1 array → pointer struct
	if len(source.Fips) > 0 {
		f := source.Fips[0]
		settings.Fips = &TargetV5FipsModel{
			TLS: f.TLS,
		}
	}

	// antivirus: TypeList MaxItems:1 array → pointer struct
	// notification_settings is a nested TypeList MaxItems:1 with field rename
	if len(source.Antivirus) > 0 {
		av := source.Antivirus[0]
		antivirusModel, avDiags := transformAntivirus(ctx, av)
		diags.Append(avDiags...)
		if diags.HasError() {
			return nil, diags
		}
		settings.Antivirus = antivirusModel
	}

	// extended_email_matching: TypeList MaxItems:1 array → pointer struct
	if len(source.ExtendedEmailMatching) > 0 {
		eem := source.ExtendedEmailMatching[0]
		settings.ExtendedEmailMatching = &TargetV5ExtendedEmailMatchingModel{
			Enabled: eem.Enabled,
			// New v5 computed fields not in v4 → null
			ReadOnly:      types.BoolNull(),
			SourceAccount: types.StringNull(),
			Version:       types.Int64Null(),
		}
	}

	// custom_certificate: TypeList MaxItems:1 array → pointer struct
	if len(source.CustomCertificate) > 0 {
		cc := source.CustomCertificate[0]
		settings.CustomCertificate = transformCustomCertificate(cc)
	}

	// certificate: TypeList MaxItems:1 array → pointer struct
	if len(source.Certificate) > 0 {
		cert := source.Certificate[0]
		settings.Certificate = &TargetV5CertificateModel{
			ID: cert.ID,
		}
	}

	// logging, proxy, ssh_session_log, payload_log: dropped
	// These are handled by tf-migrate (logging → cloudflare_zero_trust_gateway_logging,
	// proxy → cloudflare_zero_trust_device_settings; ssh_session_log and payload_log dropped)

	return settings, diags
}

// transformAntivirus handles the antivirus block transformation including the
// nested notification_settings with field rename (message → msg).
func transformAntivirus(ctx context.Context, source SourceV4AntivirusModel) (*TargetV5AntivirusModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5AntivirusModel{
		EnabledDownloadPhase: source.EnabledDownloadPhase,
		EnabledUploadPhase:   source.EnabledUploadPhase,
		FailClosed:           source.FailClosed,
	}

	if len(source.NotificationSettings) > 0 {
		ns := source.NotificationSettings[0]
		notifModel := &TargetV5NotificationSettingsModel{
			Enabled:    ns.Enabled,
			Msg:        ns.Message, // Rename: message → msg
			SupportURL: ns.SupportURL,
			// include_context is new in v5, not in v4 → null
			IncludeContext: types.BoolNull(),
		}
		notifObj, objDiags := customfield.NewObject(ctx, notifModel)
		diags.Append(objDiags...)
		if diags.HasError() {
			return nil, diags
		}
		target.NotificationSettings = notifObj
	} else {
		target.NotificationSettings = customfield.NullObject[TargetV5NotificationSettingsModel](ctx)
	}

	return target, diags
}

// transformCustomCertificate handles the custom_certificate block transformation.
// updated_at is a plain string in v4; v5 uses timetypes.RFC3339.
func transformCustomCertificate(source SourceV4CustomCertificateModel) *TargetV5CustomCertificateModel {
	target := &TargetV5CustomCertificateModel{
		Enabled: source.Enabled,
		ID:      source.ID,
		// binding_status is new computed in v5 → null
		BindingStatus: types.StringNull(),
	}

	// Convert updated_at: plain string → RFC3339
	// If the v4 string is empty or null, set to null RFC3339
	if source.UpdatedAt.IsNull() || source.UpdatedAt.IsUnknown() || source.UpdatedAt.ValueString() == "" {
		target.UpdatedAt = timetypes.NewRFC3339Null()
	} else {
		target.UpdatedAt = timetypes.NewRFC3339ValueMust(source.UpdatedAt.ValueString())
	}

	return target
}
