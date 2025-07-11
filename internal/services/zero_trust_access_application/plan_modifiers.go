package zero_trust_access_application

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"regexp"
	"slices"
)

var (
	selfHostedAppTypes         = []string{"self_hosted", "ssh", "vnc", "rdp"}
	saasAppTypes               = []string{"saas", "dash_sso"}
	appLauncherVisibleAppTypes = []string{"self_hosted", "ssh", "vnc", "rdp", "saas", "bookmark", "infrastructure"}
	targetCompatibleAppTypes   = []string{"rdp", "infrastructure"}
	durationRegex              = regexp.MustCompile(`^(?:0|[-+]?(\d+(?:\.\d*)?|\.\d+)(?:ns|us|µs|ms|s|m|h)(?:(\d+(?:\.\d*)?|\.\d+)(?:ns|us|µs|ms|s|m|h))*)$`)
)

// Sets a specific default value for a computed attribute specific to a set of app types, in case the attribute is unknown.
// If the app type is not in the list, it sets the second default value.
func setDefaultAccordingToAppTypes[T attr.Value](wantAppTypes []string, gotAppType string, planAttribute *T, default1, default2 T) {
	if planAttribute == nil || !(*planAttribute).IsUnknown() {
		return
	}
	if slices.Contains(wantAppTypes, gotAppType) {
		*planAttribute = default1
	} else {
		*planAttribute = default2
	}
}

// Sets a specific default value for a computed attribute specific to a given app type, in case the attribute is unknown.
// If the app type does not match, it sets the second default value.
func setDefaultAccordingToAppType[T attr.Value](wantAppType string, gotAppType string, planAttribute *T, default1, default2 T) {
	setDefaultAccordingToAppTypes([]string{wantAppType}, gotAppType, planAttribute, default1, default2)
}

func modifyPlanForDomains(ctx context.Context, planApp, stateApp *ZeroTrustAccessApplicationModel) {
	appType := planApp.Type.ValueString()

	setDefaultAccordingToAppTypes(selfHostedAppTypes, appType, &planApp.SelfHostedDomains, customfield.UnknownList[types.String](ctx), customfield.NullList[types.String](ctx))
	setDefaultAccordingToAppTypes(selfHostedAppTypes, appType, &planApp.Destinations, customfield.UnknownObjectList[ZeroTrustAccessApplicationDestinationsModel](ctx), customfield.NullObjectList[ZeroTrustAccessApplicationDestinationsModel](ctx))
	setDefaultAccordingToAppTypes(selfHostedAppTypes, appType, &planApp.HTTPOnlyCookieAttribute, types.BoolUnknown(), types.BoolNull())

	// A self_hosted_app's 'domain', 'self_hosted_domains', and 'destinations' are all tied together in the API.
	// changing one, causes the others to change. So we need to tell TF to set the other two to unknown if any of them
	// changes from the previous state.
	if stateApp == nil ||
		(!planApp.Domain.IsUnknown() && !planApp.Domain.Equal(stateApp.Domain)) ||
		(!planApp.SelfHostedDomains.IsUnknown() && !planApp.SelfHostedDomains.Equal(stateApp.SelfHostedDomains)) ||
		(!planApp.Destinations.IsUnknown() && !planApp.Destinations.Equal(stateApp.Destinations)) {

		if planApp.Domain.IsNull() {
			planApp.Domain = types.StringUnknown()
		}
		if planApp.SelfHostedDomains.IsNull() {
			planApp.SelfHostedDomains = customfield.UnknownList[types.String](ctx)
		}
		if planApp.Destinations.IsNull() {
			planApp.Destinations = customfield.UnknownObjectList[ZeroTrustAccessApplicationDestinationsModel](ctx)
		}
	} else {
		// If the domain, self_hosted_domains, and destinations have not changed, we can copy all of them from the state.
		planApp.Domain = stateApp.Domain
		planApp.SelfHostedDomains = stateApp.SelfHostedDomains
		planApp.Destinations = stateApp.Destinations
	}
}

func modifySaasAppNestedObjectPlan(ctx context.Context, planApp *ZeroTrustAccessApplicationModel) diag.Diagnostics {
	diags := diag.Diagnostics{}
	if planApp.SaaSApp.IsNull() {
		return diags
	}
	var planSaasApp ZeroTrustAccessApplicationSaaSAppModel
	diags.Append(planApp.SaaSApp.As(ctx, &planSaasApp, basetypes.ObjectAsOptions{})...)

	oidcType, samlType, currentType := "oidc", "saml", planSaasApp.AuthType.ValueString()

	// These fields are non-existent for non-oidc saas_apps. So we can set them to null and avoid the recurring
	// diffs due to unknown computed values.
	setDefaultAccordingToAppType(oidcType, currentType, &planSaasApp.ClientID, types.StringUnknown(), types.StringNull())
	setDefaultAccordingToAppType(oidcType, currentType, &planSaasApp.ClientSecret, types.StringUnknown(), types.StringNull())
	setDefaultAccordingToAppType(oidcType, currentType, &planSaasApp.AllowPKCEWithoutClientSecret, types.BoolValue(false), types.BoolNull())
	setDefaultAccordingToAppType(oidcType, currentType, &planSaasApp.AccessTokenLifetime, types.StringValue("5m"), types.StringNull())

	// These fields are non-existent for non-saml saas_apps. So we can set them to null and avoid the recurring
	// diffs due to unknown computed values.
	setDefaultAccordingToAppType(samlType, currentType, &planSaasApp.IdPEntityID, types.StringUnknown(), types.StringNull())
	setDefaultAccordingToAppType(samlType, currentType, &planSaasApp.NameIDFormat, types.StringUnknown(), types.StringNull())
	setDefaultAccordingToAppType(samlType, currentType, &planSaasApp.SSOEndpoint, types.StringUnknown(), types.StringNull())

	planApp.SaaSApp, _ = customfield.NewObject[ZeroTrustAccessApplicationSaaSAppModel](ctx, &planSaasApp)
	return diags
}

func modifyNestedPoliciesPlan(_ context.Context, planApp *ZeroTrustAccessApplicationModel) {
	lastKnownPrecedence := 0
	for i := range *planApp.Policies {
		if (*planApp.Policies)[i].Precedence.IsUnknown() {
			(*planApp.Policies)[i].Precedence = types.Int64Value(int64(lastKnownPrecedence + 1))
		}
		lastKnownPrecedence = int((*planApp.Policies)[i].Precedence.ValueInt64())
	}
}

func modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, res *resource.ModifyPlanResponse) {
	var planApp, stateApp *ZeroTrustAccessApplicationModel
	res.Diagnostics.Append(req.Plan.Get(ctx, &planApp)...)
	res.Diagnostics.Append(req.State.Get(ctx, &stateApp)...)
	if res.Diagnostics.HasError() || planApp == nil {
		return
	}

	modifyPlanForDomains(ctx, planApp, stateApp)

	appType := planApp.Type.ValueString()

	// Add default values for some app type specific attributes
	setDefaultAccordingToAppTypes(selfHostedAppTypes, appType, &planApp.HTTPOnlyCookieAttribute, types.BoolValue(true), types.BoolNull())
	setDefaultAccordingToAppTypes(appLauncherVisibleAppTypes, appType, &planApp.AppLauncherVisible, types.BoolValue(true), types.BoolNull())
	setDefaultAccordingToAppType("app_launcher", appType, &planApp.SkipAppLauncherLoginPage, types.BoolValue(false), types.BoolNull())

	if appType == "saas" {
		res.Diagnostics.Append(modifySaasAppNestedObjectPlan(ctx, planApp)...)
	}

	if planApp.Policies != nil {
		modifyNestedPoliciesPlan(ctx, planApp)
	}

	res.Plan.Set(ctx, &planApp)
}
