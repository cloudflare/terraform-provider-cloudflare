package zero_trust_access_identity_provider

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func normalizeFalseAndNullBool(data *basetypes.BoolValue, stateData basetypes.BoolValue) {
	if data.ValueBool() || stateData.ValueBool() {
		return
	}
	*data = stateData
}

func normalizeReadZeroTrustIDPScimConfigData(ctx context.Context, dataValue, stateValue *customfield.NestedObject[ZeroTrustAccessIdentityProviderSCIMConfigModel]) diag.Diagnostics {
	var (
		diags                           = make(diag.Diagnostics, 0)
		dataScimConfig, stateScimConfig ZeroTrustAccessIdentityProviderSCIMConfigModel
	)

	diags.Append(dataValue.As(ctx, &dataScimConfig, basetypes.ObjectAsOptions{})...)
	diags.Append(stateValue.As(ctx, &stateScimConfig, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return diags
	}

	if !stateScimConfig.Secret.IsUnknown() && !stateScimConfig.Secret.IsNull() {
		// Scim secret is only generated and returned in the create request, and null on reads.
		// so we need to load it from the state
		dataScimConfig.Secret = stateScimConfig.Secret
	}

	*dataValue, diags = customfield.NewObject[ZeroTrustAccessIdentityProviderSCIMConfigModel](ctx, &dataScimConfig)
	return diags
}

func normalizeReadZeroTrustIDPConfigData(ctx context.Context, dataValue *ZeroTrustAccessIdentityProviderModel, stateValue ZeroTrustAccessIdentityProviderModel) diag.Diagnostics {
	diag := make(diag.Diagnostics, 0)
	if dataValue.Config == nil || stateValue.Config == nil {
		return diag
	}

	normalizeFalseAndNullBool(&dataValue.Config.SignRequest, stateValue.Config.SignRequest)
	normalizeFalseAndNullBool(&dataValue.Config.ConditionalAccessEnabled, stateValue.Config.ConditionalAccessEnabled)
	normalizeFalseAndNullBool(&dataValue.Config.SupportGroups, stateValue.Config.SupportGroups)

	return diag
}

func normalizeReadZeroTrustIDPData(ctx context.Context, dataValue *ZeroTrustAccessIdentityProviderModel, state *tfsdk.State) diag.Diagnostics {
	var (
		diags      = make(diag.Diagnostics, 0)
		stateValue ZeroTrustAccessIdentityProviderModel
	)
	d := state.Get(ctx, &stateValue)

	diags.Append(normalizeReadZeroTrustIDPConfigData(ctx, dataValue, stateValue)...)
	if diags.HasError() {
		return diags
	}

	// scim_config.secret is only returned when the app is first created, assigning here from the state
	// to prevent a diff when the app is updated
	if !dataValue.SCIMConfig.IsNull() && (!stateValue.SCIMConfig.IsNull() && !stateValue.SCIMConfig.IsUnknown()) {
		diags.Append(normalizeReadZeroTrustIDPScimConfigData(ctx, &dataValue.SCIMConfig, &stateValue.SCIMConfig)...)
		if diags.HasError() {
			return diags
		}
	}

	return d
}

// Some fields are write-only sensitive and should not be stored in the state.
// Usually these secrets are injected in the config from a secret store.
func loadConfigSensitiveValuesForWriting(ctx context.Context, data *ZeroTrustAccessIdentityProviderModel, cfg *tfsdk.Config) diag.Diagnostics {
	var (
		diags   = make(diag.Diagnostics, 0)
		cfgData *ZeroTrustAccessIdentityProviderModel
	)
	diags.Append(cfg.Get(ctx, &cfgData)...)

	if data.Config != nil && cfgData.Config != nil {
		data.Config.ClientSecret = cfgData.Config.ClientSecret
	}

	return diags
}
