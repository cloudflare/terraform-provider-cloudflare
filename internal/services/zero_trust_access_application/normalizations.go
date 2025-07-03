package zero_trust_access_application

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"slices"
)

func normalizeEmptyAndNullString(data *basetypes.StringValue, stateData basetypes.StringValue) {
	if data.ValueString() != "" || stateData.ValueString() != "" {
		return
	}
	*data = stateData
}

func normalizeFalseAndNullBool(data *basetypes.BoolValue, stateData basetypes.BoolValue) {
	if data.ValueBool() || stateData.ValueBool() {
		return
	}
	*data = stateData
}

func normalizeTrueAndNullBool(data *basetypes.BoolValue, stateData basetypes.BoolValue) {
	if (!data.IsNull() && !data.ValueBool()) || (!stateData.IsNull() && !stateData.ValueBool()) {
		return
	}
	if stateData.IsUnknown() {
		return
	}
	*data = stateData
}

type ListField interface {
	Elements() []attr.Value
}

func normalizeEmptyAndNullList[T ListField](data *T, stateData T) {
	if len((*data).Elements()) != 0 || len(stateData.Elements()) != 0 {
		return
	}
	*data = stateData
}

func normalizeEmptyAndNullSlice[T any](data **[]T, stateData *[]T) {
	if (*data != nil && len(**data) != 0) || (stateData != nil && len(*stateData) != 0) {
		return
	}
	*data = stateData
}

type IsNull interface {
	IsNull() bool
}

func persistNullFromState[T IsNull](data *T, stateData T) {
	if stateData.IsNull() {
		*data = stateData
	}
}

func normalizeReadZeroTrustApplicationSamlAppData(data *ZeroTrustAccessApplicationSaaSAppModel, stateData ZeroTrustAccessApplicationSaaSAppModel) {
	normalizeEmptyAndNullString(&data.SPEntityID, stateData.SPEntityID)
	normalizeEmptyAndNullString(&data.ConsumerServiceURL, stateData.ConsumerServiceURL)
}

func normalizeReadZeroTrustApplicationOidcAppData(data *ZeroTrustAccessApplicationSaaSAppModel, stateData ZeroTrustAccessApplicationSaaSAppModel) {
	// Prevent diffs on the default access_token_lifetime
	if data.AccessTokenLifetime.ValueString() == "5m" && stateData.AccessTokenLifetime == types.StringNull() {
		data.AccessTokenLifetime = stateData.AccessTokenLifetime
	}

	// client_secret is only returned when the app is first created, assigning here from the state
	// to prevent a diff when the app is updated
	if !stateData.ClientSecret.IsUnknown() && !stateData.ClientSecret.IsNull() {
		data.ClientSecret = stateData.ClientSecret
	}

	normalizeFalseAndNullBool(&data.AllowPKCEWithoutClientSecret, stateData.AllowPKCEWithoutClientSecret)
}

// Normalizing function to ensure consistency between the state and the meaning of the API response.
// Alters the API response before applying it to the state by laxing equalities between null & zero-value
// for some attributes, and nullifies fields that terraform should not be saving in the state.
func normalizeReadZeroTrustApplicationAPIData(ctx context.Context, data, stateData *ZeroTrustAccessApplicationModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)

	// Empty `allowed_idps` is the same as a null value. The API might return an empty array, so we need to normalize it
	// here  to avoid a diff
	normalizeEmptyAndNullSlice(&data.AllowedIdPs, stateData.AllowedIdPs)
	// `policies` might not be in the configuration, so we need to normalize it here to avoid a diff
	normalizeEmptyAndNullSlice(&data.Policies, stateData.Policies)
	// `tags` might not be in the configuration, so we need to normalize it here to avoid a diff
	normalizeEmptyAndNullList(&data.Tags, stateData.Tags)

	normalizeFalseAndNullBool(&data.EnableBindingCookie, stateData.EnableBindingCookie)
	normalizeFalseAndNullBool(&data.OptionsPreflightBypass, stateData.OptionsPreflightBypass)
	normalizeFalseAndNullBool(&data.AutoRedirectToIdentity, stateData.AutoRedirectToIdentity)
	if slices.Contains(selfHostedAppTypes, data.Type.String()) {
		normalizeTrueAndNullBool(&data.HTTPOnlyCookieAttribute, stateData.HTTPOnlyCookieAttribute)
	}

	if !data.SaaSApp.IsNull() && !stateData.SaaSApp.IsNull() {
		var dataSaasApp, stateDataSaasApp ZeroTrustAccessApplicationSaaSAppModel
		diags.Append(data.SaaSApp.As(ctx, &dataSaasApp, basetypes.ObjectAsOptions{})...)
		diags.Append(stateData.SaaSApp.As(ctx, &stateDataSaasApp, basetypes.ObjectAsOptions{})...)
		if diags.HasError() {
			return diags
		}

		switch dataSaasApp.AuthType.ValueString() {
		case "saml":
			normalizeReadZeroTrustApplicationSamlAppData(&dataSaasApp, stateDataSaasApp)
		case "oidc":
			normalizeReadZeroTrustApplicationOidcAppData(&dataSaasApp, stateDataSaasApp)
		}

		var saasDiags diag.Diagnostics
		data.SaaSApp, saasDiags = customfield.NewObject[ZeroTrustAccessApplicationSaaSAppModel](ctx, &dataSaasApp)
		diags.Append(saasDiags...)
		if diags.HasError() {
			return diags
		}
	}

	if data.Policies != nil && stateData.Policies != nil {
		for i := range *data.Policies {
			// Preserve null values from the Terraform state, even if the API response returns actual values.
			// This is important because the API may populate these fields when it expands the attached reusable policy
			// from its given ID.
			//
			// However, we intentionally avoid storing the full expanded policy inside the application resource's
			// nested block, as its source of truth is the reusable policy resource itself.
			// Only the policy ID should be persisted in state for reusable policies.
			// For legacy policies, the ID should be ignored as they are not a standalone resource, but rather
			// live as a nested object owned by the application.
			persistNullFromState(&(*data.Policies)[i].ID, (*stateData.Policies)[i].ID)
			persistNullFromState(&(*data.Policies)[i].Decision, (*stateData.Policies)[i].Decision)
			persistNullFromState(&(*data.Policies)[i].Name, (*stateData.Policies)[i].Name)
			persistNullFromState(&(*data.Policies)[i].Include, (*stateData.Policies)[i].Include)
			persistNullFromState(&(*data.Policies)[i].Require, (*stateData.Policies)[i].Require)
			persistNullFromState(&(*data.Policies)[i].Exclude, (*stateData.Policies)[i].Exclude)
		}
	}

	if data.SCIMConfig != nil && stateData.SCIMConfig != nil {
		if data.SCIMConfig.Authentication != nil && stateData.SCIMConfig.Authentication != nil {
			data.SCIMConfig.Authentication.Password = stateData.SCIMConfig.Authentication.Password
			data.SCIMConfig.Authentication.Token = stateData.SCIMConfig.Authentication.Token
			data.SCIMConfig.Authentication.ClientSecret = stateData.SCIMConfig.Authentication.ClientSecret
		}
	}

	return diags
}

// Some fields are write-only sensitive and should not be stored in the state.
// Usually these secrets are injected in the config from a secret store.
func loadConfigSensitiveValuesForWriting(ctx context.Context, data *ZeroTrustAccessApplicationModel, cfg *tfsdk.Config) diag.Diagnostics {
	var (
		diags   = make(diag.Diagnostics, 0)
		cfgData *ZeroTrustAccessApplicationModel
	)
	diags.Append(cfg.Get(ctx, &cfgData)...)

	if data.SCIMConfig != nil && cfgData.SCIMConfig != nil {
		if data.SCIMConfig.Authentication != nil && cfgData.SCIMConfig.Authentication != nil {
			data.SCIMConfig.Authentication.Password = cfgData.SCIMConfig.Authentication.Password
			data.SCIMConfig.Authentication.Token = cfgData.SCIMConfig.Authentication.Token
			data.SCIMConfig.Authentication.ClientSecret = cfgData.SCIMConfig.Authentication.ClientSecret
		}
	}
	return diags
}
