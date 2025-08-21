package zero_trust_access_application

import (
	"context"
	"slices"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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

func normalizeZeroTrustApplicationPolicyConnectionRulesAPIData(_ context.Context, data, stateData *ZeroTrustAccessApplicationPoliciesConnectionRulesModel) {
	if data.SSH != nil && stateData.SSH != nil {
		normalizeFalseAndNullBool(&data.SSH.AllowEmailAlias, stateData.SSH.AllowEmailAlias)
	}
}

func normalizeZeroTrustApplicationPolicyAPIData(ctx context.Context, data, stateData *ZeroTrustAccessApplicationPoliciesModel) {
	// Preserve null values from the Terraform state, even if the API response returns actual values.
	// This is important because the API may populate these fields when it expands the attached reusable policy
	// from its given ID.
	//
	// However, we intentionally avoid storing the full expanded policy inside the application resource's
	// nested block, as its source of truth is the reusable policy resource itself.
	// Only the policy ID should be persisted in state for reusable policies.
	// For legacy policies, the ID should be ignored as they are not a standalone resource, but rather
	// live as a nested object owned by the application.
	persistNullFromState(&data.ID, stateData.ID)
	persistNullFromState(&data.Decision, stateData.Decision)
	persistNullFromState(&data.Name, stateData.Name)
	persistNullFromState(&data.Include, stateData.Include)
	persistNullFromState(&data.Require, stateData.Require)
	persistNullFromState(&data.Exclude, stateData.Exclude)

	if data.ConnectionRules != nil && stateData.ConnectionRules != nil {
		normalizeZeroTrustApplicationPolicyConnectionRulesAPIData(ctx, data.ConnectionRules, stateData.ConnectionRules)
	}
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

	normalizeFalseAndNullBool(&data.ServiceAuth401Redirect, stateData.ServiceAuth401Redirect)
	normalizeFalseAndNullBool(&data.EnableBindingCookie, stateData.EnableBindingCookie)
	normalizeFalseAndNullBool(&data.OptionsPreflightBypass, stateData.OptionsPreflightBypass)
	normalizeFalseAndNullBool(&data.AutoRedirectToIdentity, stateData.AutoRedirectToIdentity)
	if slices.Contains(selfHostedAppTypes, data.Type.ValueString()) {
		normalizeTrueAndNullBool(&data.HTTPOnlyCookieAttribute, stateData.HTTPOnlyCookieAttribute)
		normalizeFalseAndNullBool(&data.SkipInterstitial, stateData.SkipInterstitial)
		normalizeFalseAndNullBool(&data.AllowIframe, stateData.AllowIframe)
		normalizeFalseAndNullBool(&data.PathCookieAttribute, stateData.PathCookieAttribute)

		if data.CORSHeaders != nil && stateData.CORSHeaders != nil {
			// This is the only bool CORSHeaders field needed for normalization because
			// the other fields are not allowed to be false. e.g. AllowAllOrigins = false
			// requires AllowedOrigins list to be present, and these fields are mutually exclusive.
			normalizeFalseAndNullBool(&data.CORSHeaders.AllowCredentials, stateData.CORSHeaders.AllowCredentials)
		}
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
			if len(*stateData.Policies) <= i {
				break
			}
			normalizeZeroTrustApplicationPolicyAPIData(ctx, &(*data.Policies)[i], &(*stateData.Policies)[i])
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

// Normalizes the API request before sending it to the API
func normalizeWriteZeroTrustApplicationAPIData(ctx context.Context, data *ZeroTrustAccessApplicationModel, cfg *tfsdk.Config) diag.Diagnostics {
	var (
		diags   = make(diag.Diagnostics, 0)
		cfgData *ZeroTrustAccessApplicationModel
	)
	diags.Append(cfg.Get(ctx, &cfgData)...)

	if data.SCIMConfig != nil && cfgData.SCIMConfig != nil {
		// load config sensitive write values directly from the config.
		if data.SCIMConfig.Authentication != nil && cfgData.SCIMConfig.Authentication != nil {
			data.SCIMConfig.Authentication.Password = cfgData.SCIMConfig.Authentication.Password
			data.SCIMConfig.Authentication.Token = cfgData.SCIMConfig.Authentication.Token
			data.SCIMConfig.Authentication.ClientSecret = cfgData.SCIMConfig.Authentication.ClientSecret
		}
	}

	// If the API receives a null 'policies' array, it wont update the policies on the application, for historical reasons.
	// To avoid a diff, we need to ensure that the array is not nil
	if data.Policies == nil {
		data.Policies = &[]ZeroTrustAccessApplicationPoliciesModel{}
	}

	return diags
}

func normalizeImportZeroTrustAccessApplicationAPIData(ctx context.Context, data *ZeroTrustAccessApplicationModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)
	if data.AllowedIdPs != nil && len(*data.AllowedIdPs) == 0 {
		data.AllowedIdPs = nil
	}
	if data.Policies != nil && len(*data.Policies) == 0 {
		data.Policies = nil
	}

	if data.Policies != nil {
		for i := range *data.Policies {
			policy := &(*data.Policies)[i]
			if !policy.ID.IsNull() && !policy.ID.IsUnknown() {
				policy.Decision = types.StringNull()
				policy.Name = types.StringNull()
				policy.Include = customfield.NullObjectList[ZeroTrustAccessApplicationPoliciesIncludeModel](ctx)
				policy.Require = customfield.NullObjectList[ZeroTrustAccessApplicationPoliciesRequireModel](ctx)
				policy.Exclude = customfield.NullObjectList[ZeroTrustAccessApplicationPoliciesExcludeModel](ctx)
			} else {
				if !policy.Include.IsNull() && len(policy.Include.Elements()) == 0 {
					policy.Include = customfield.NullObjectList[ZeroTrustAccessApplicationPoliciesIncludeModel](ctx)
				}
				if !policy.Require.IsNull() && len(policy.Require.Elements()) == 0 {
					policy.Require = customfield.NullObjectList[ZeroTrustAccessApplicationPoliciesRequireModel](ctx)
				}
				if !policy.Exclude.IsNull() && len(policy.Exclude.Elements()) == 0 {
					policy.Exclude = customfield.NullObjectList[ZeroTrustAccessApplicationPoliciesExcludeModel](ctx)
				}
			}
		}
	}

	if !data.Tags.IsNull() && !data.Tags.IsUnknown() {
		if len(data.Tags.Elements()) == 0 {
			data.Tags = customfield.NullList[types.String](ctx)
		}
	}

	return diags
}
