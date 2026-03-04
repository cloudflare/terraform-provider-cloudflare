package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Transform converts source (v4) state to target (v5) state.
func Transform(ctx context.Context, source SourceMTLSHostnameSettingsModel) (*TargetMTLSHostnameSettingsModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	tflog.Debug(ctx, "Transforming access_mtls_hostname_settings state from v4 to v5")

	target := &TargetMTLSHostnameSettingsModel{
		AccountID: source.AccountID,
		ZoneID:    source.ZoneID,
		// Root computed fields - set to null, will refresh from API
		ChinaNetwork:                types.BoolNull(),
		ClientCertificateForwarding: types.BoolNull(),
		Hostname:                    types.StringNull(),
	}

	// Transform settings: []SourceSettingsModel → *[]*TargetSettingsModel
	// Add false defaults for Optional→Required fields if missing
	if len(source.Settings) > 0 {
		settings := make([]*TargetSettingsModel, len(source.Settings))
		for i, s := range source.Settings {
			settings[i] = &TargetSettingsModel{
				Hostname: s.Hostname,
				// china_network: default false if null (was Optional in v4, Required in v5)
				ChinaNetwork: defaultFalseBool(s.ChinaNetwork),
				// client_certificate_forwarding: default false if null
				ClientCertificateForwarding: defaultFalseBool(s.ClientCertificateForwarding),
			}
		}
		target.Settings = &settings
	}

	return target, diags
}

// defaultFalseBool returns the value if set, or false if null/unknown.
func defaultFalseBool(val types.Bool) types.Bool {
	if val.IsNull() || val.IsUnknown() {
		return types.BoolValue(false)
	}
	return val
}
