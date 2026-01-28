package zero_trust_access_application

import (
	"context"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.ResourceWithMoveState = (*ZeroTrustAccessApplicationResource)(nil)
var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessApplicationResource)(nil)

// MoveState handles moves from cloudflare_access_application (v4) to cloudflare_zero_trust_access_application (v5).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_application.example
//	    to   = cloudflare_zero_trust_access_application.example
//	}
func (r *ZeroTrustAccessApplicationResource) MoveState(ctx context.Context) []resource.StateMover {
	v4Schema := V4AccessApplicationSchema()
	return []resource.StateMover{
		{
			SourceSchema: &v4Schema,
			StateMover:   r.moveFromAccessApplication,
		},
	}
}

func (r *ZeroTrustAccessApplicationResource) moveFromAccessApplication(
	ctx context.Context,
	req resource.MoveStateRequest,
	resp *resource.MoveStateResponse,
) {
	// Verify source is cloudflare_access_application from cloudflare provider
	if req.SourceTypeName != "cloudflare_access_application" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Starting state move from cloudflare_access_application to cloudflare_zero_trust_access_application",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	// Parse the v4 state using the v4 schema
	var v4State V4AccessApplicationModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 state to v5 state
	v5State, diags := transformV4ToV5(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the v5 state
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, v5State)...)

	tflog.Info(ctx, "State move from cloudflare_access_application to cloudflare_zero_trust_access_application completed successfully")
}

// UpgradeState handles schema version upgrades.
func (r *ZeroTrustAccessApplicationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v4Schema := V4AccessApplicationSchema()
	v5Schema := ResourceSchema(ctx)
	return map[int64]resource.StateUpgrader{
		// Handle upgrades from earlier v5 versions (no schema changes, just version bump)
		0: {
			PriorSchema:   &v5Schema,
			StateUpgrader: r.upgradeFromV5,
		},
		// Handle state moved from cloudflare_access_application (v4 provider)
		// When users run `terraform state mv cloudflare_access_application.x cloudflare_zero_trust_access_application.x`,
		// the old schema_version is preserved, triggering this upgrader.
		2: {
			PriorSchema:   &v4Schema,
			StateUpgrader: r.upgradeFromV4,
		},
	}
}

func (r *ZeroTrustAccessApplicationResource) upgradeFromV5(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	// No-op upgrade: schema is compatible, just copy state through
	var state ZeroTrustAccessApplicationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ZeroTrustAccessApplicationResource) upgradeFromV4(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading access application state from v4 cloudflare_access_application format")

	// Parse the v4 state
	var v4State V4AccessApplicationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to v5
	v5State, diags := transformV4ToV5(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)

	tflog.Info(ctx, "State upgrade from v4 cloudflare_access_application completed successfully")
}

// transformV4ToV5 converts a v4 cloudflare_access_application state to v5 cloudflare_zero_trust_access_application state.
func transformV4ToV5(ctx context.Context, v4 V4AccessApplicationModel) (*ZeroTrustAccessApplicationModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	v5 := &ZeroTrustAccessApplicationModel{
		ID:                       v4.ID,
		AccountID:                v4.AccountID,
		ZoneID:                   v4.ZoneID,
		Name:                     v4.Name,
		Domain:                   v4.Domain,
		SessionDuration:          v4.SessionDuration,
		AutoRedirectToIdentity:   v4.AutoRedirectToIdentity,
		EnableBindingCookie:      v4.EnableBindingCookie,
		HTTPOnlyCookieAttribute:  v4.HTTPOnlyCookieAttribute,
		SameSiteCookieAttribute:  v4.SameSiteCookieAttribute,
		LogoURL:                  v4.LogoURL,
		SkipInterstitial:         v4.SkipInterstitial,
		AppLauncherVisible:       v4.AppLauncherVisible,
		ServiceAuth401Redirect:   v4.ServiceAuth401Redirect,
		CustomDenyMessage:        v4.CustomDenyMessage,
		CustomDenyURL:            v4.CustomDenyURL,
		CustomNonIdentityDenyURL: v4.CustomNonIdentityDenyURL,
		OptionsPreflightBypass:   v4.OptionsPreflightBypass,
		PathCookieAttribute:      v4.PathCookieAttribute,
		AUD:                      v4.AUD,
		AppLauncherLogoURL:       v4.AppLauncherLogoURL,
		HeaderBgColor:            v4.HeaderBgColor,
		BgColor:                  v4.BgColor,
		SkipAppLauncherLoginPage: v4.SkipAppLauncherLoginPage,
		AllowAuthenticateViaWARP: v4.AllowAuthenticateViaWARP,
	}

	// Type: v4 may not have this, v5 defaults to "self_hosted"
	if !v4.Type.IsNull() && !v4.Type.IsUnknown() && v4.Type.ValueString() != "" {
		v5.Type = v4.Type
	} else {
		v5.Type = types.StringValue("self_hosted")
	}

	// AllowedIdPs: v4 Set -> v5 *[]types.String
	if !v4.AllowedIdPs.IsNull() && !v4.AllowedIdPs.IsUnknown() {
		allowedIdPs, d := setToStringSlice(ctx, v4.AllowedIdPs)
		diags.Append(d...)
		if !diags.HasError() {
			v5.AllowedIdPs = allowedIdPs
		}
	}

	// Tags: v4 Set -> v5 customfield.Set
	if !v4.Tags.IsNull() && !v4.Tags.IsUnknown() {
		tags, d := setToCustomfieldSet(ctx, v4.Tags)
		diags.Append(d...)
		if !diags.HasError() {
			v5.Tags = tags
		}
	}

	// SelfHostedDomains: v4 Set -> v5 customfield.Set
	if !v4.SelfHostedDomains.IsNull() && !v4.SelfHostedDomains.IsUnknown() {
		domains, d := setToCustomfieldSet(ctx, v4.SelfHostedDomains)
		diags.Append(d...)
		if !diags.HasError() {
			v5.SelfHostedDomains = domains
		}
	}

	// CustomPages: v4 Set -> v5 *[]types.String
	if !v4.CustomPages.IsNull() && !v4.CustomPages.IsUnknown() {
		customPages, d := setToStringSlice(ctx, v4.CustomPages)
		diags.Append(d...)
		if !diags.HasError() {
			v5.CustomPages = customPages
		}
	}

	// CORSHeaders: v4 []V4CORSHeadersModel (list block) -> v5 *ZeroTrustAccessApplicationCORSHeadersModel (object)
	if len(v4.CORSHeaders) > 0 {
		corsHeaders, d := transformV4CORSHeadersToV5(ctx, v4.CORSHeaders[0])
		diags.Append(d...)
		if !diags.HasError() {
			v5.CORSHeaders = corsHeaders
		}
	}

	// SaaSApp: v4 []V4SaaSAppModel (list block) -> v5 *ZeroTrustAccessApplicationSaaSAppModel (object)
	if len(v4.SaaSApp) > 0 {
		saasApp, d := transformV4SaaSAppToV5(ctx, v4.SaaSApp[0])
		diags.Append(d...)
		if !diags.HasError() {
			v5.SaaSApp = saasApp
		}
	}

	// SCIMConfig: v4 []V4SCIMConfigModel (list block) -> v5 *ZeroTrustAccessApplicationSCIMConfigModel (object)
	if len(v4.SCIMConfig) > 0 {
		scimConfig, d := transformV4SCIMConfigToV5(ctx, v4.SCIMConfig[0])
		diags.Append(d...)
		if !diags.HasError() {
			v5.SCIMConfig = scimConfig
		}
	}

	// LandingPageDesign: v4 []V4LandingPageDesignModel (list block) -> v5 customfield.NestedObject (object)
	if len(v4.LandingPageDesign) > 0 {
		landingPage, d := transformV4LandingPageToV5(ctx, v4.LandingPageDesign[0])
		diags.Append(d...)
		if !diags.HasError() {
			v5.LandingPageDesign = landingPage
		}
	}

	// FooterLinks: v4 []V4FooterLinksModel -> v5 *[]*ZeroTrustAccessApplicationFooterLinksModel
	if len(v4.FooterLinks) > 0 {
		footerLinks := make([]*ZeroTrustAccessApplicationFooterLinksModel, 0, len(v4.FooterLinks))
		for _, link := range v4.FooterLinks {
			footerLinks = append(footerLinks, &ZeroTrustAccessApplicationFooterLinksModel{
				Name: link.Name,
				URL:  link.URL,
			})
		}
		v5.FooterLinks = &footerLinks
	}

	// Destinations: v4 []V4DestinationsModel -> v5 customfield.NestedObjectList
	if len(v4.Destinations) > 0 {
		destinations, d := transformV4DestinationsToV5(ctx, v4.Destinations)
		diags.Append(d...)
		if !diags.HasError() {
			v5.Destinations = destinations
		}
	}

	// TargetCriteria: v4 []V4TargetCriteriaModel -> v5 *[]*ZeroTrustAccessApplicationTargetCriteriaModel
	if len(v4.TargetCriteria) > 0 {
		targetCriteria, d := transformV4TargetCriteriaToV5(ctx, v4.TargetCriteria)
		diags.Append(d...)
		if !diags.HasError() {
			v5.TargetCriteria = targetCriteria
		}
	}

	// Policies: v4 []types.String (string array) -> v5 *[]ZeroTrustAccessApplicationPoliciesModel (object array with id/precedence)
	if len(v4.Policies) > 0 {
		policies := make([]ZeroTrustAccessApplicationPoliciesModel, 0, len(v4.Policies))
		for i, policyID := range v4.Policies {
			if !policyID.IsNull() && !policyID.IsUnknown() {
				policies = append(policies, ZeroTrustAccessApplicationPoliciesModel{
					ID:         policyID,
					Precedence: types.Int64Value(int64(i + 1)),
				})
			}
		}
		if len(policies) > 0 {
			v5.Policies = &policies
		}
	}

	// Note: domain_type is removed in v5 (deprecated)

	return v5, diags
}

// Helper functions

func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}

func setToStringSlice(ctx context.Context, set types.Set) (*[]types.String, diag.Diagnostics) {
	var diags diag.Diagnostics
	var values []string
	diags.Append(set.ElementsAs(ctx, &values, false)...)
	if diags.HasError() {
		return nil, diags
	}
	result := make([]types.String, len(values))
	for i, v := range values {
		result[i] = types.StringValue(v)
	}
	return &result, diags
}

func setToCustomfieldSet(ctx context.Context, set types.Set) (customfield.Set[types.String], diag.Diagnostics) {
	var diags diag.Diagnostics
	var values []string
	diags.Append(set.ElementsAs(ctx, &values, false)...)
	if diags.HasError() {
		return customfield.Set[types.String]{}, diags
	}
	attrs := make([]attr.Value, len(values))
	for i, v := range values {
		attrs[i] = types.StringValue(v)
	}
	return customfield.NewSetMust[types.String](ctx, attrs), diags
}

func setToCustomfieldList(ctx context.Context, set types.Set) (customfield.List[types.String], diag.Diagnostics) {
	var diags diag.Diagnostics
	var values []string
	diags.Append(set.ElementsAs(ctx, &values, false)...)
	if diags.HasError() {
		return customfield.List[types.String]{}, diags
	}
	attrs := make([]attr.Value, len(values))
	for i, v := range values {
		attrs[i] = types.StringValue(v)
	}
	return customfield.NewListMust[types.String](ctx, attrs), diags
}

// transformV4CORSHeadersToV5 converts v4 cors_headers to v5 format.
func transformV4CORSHeadersToV5(ctx context.Context, v4 V4CORSHeadersModel) (*ZeroTrustAccessApplicationCORSHeadersModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	v5 := &ZeroTrustAccessApplicationCORSHeadersModel{
		AllowAllHeaders:  v4.AllowAllHeaders,
		AllowAllMethods:  v4.AllowAllMethods,
		AllowAllOrigins:  v4.AllowAllOrigins,
		AllowCredentials: v4.AllowCredentials,
	}

	// MaxAge: v4 Int64 -> v5 Float64
	if !v4.MaxAge.IsNull() && !v4.MaxAge.IsUnknown() {
		v5.MaxAge = types.Float64Value(float64(v4.MaxAge.ValueInt64()))
	}

	// AllowedHeaders: v4 Set -> v5 *[]types.String
	if !v4.AllowedHeaders.IsNull() && !v4.AllowedHeaders.IsUnknown() {
		headers, d := setToStringSlice(ctx, v4.AllowedHeaders)
		diags.Append(d...)
		if !diags.HasError() {
			v5.AllowedHeaders = headers
		}
	}

	// AllowedMethods: v4 Set -> v5 *[]types.String
	if !v4.AllowedMethods.IsNull() && !v4.AllowedMethods.IsUnknown() {
		methods, d := setToStringSlice(ctx, v4.AllowedMethods)
		diags.Append(d...)
		if !diags.HasError() {
			v5.AllowedMethods = methods
		}
	}

	// AllowedOrigins: v4 Set -> v5 *[]types.String
	if !v4.AllowedOrigins.IsNull() && !v4.AllowedOrigins.IsUnknown() {
		origins, d := setToStringSlice(ctx, v4.AllowedOrigins)
		diags.Append(d...)
		if !diags.HasError() {
			v5.AllowedOrigins = origins
		}
	}

	return v5, diags
}

// transformV4SaaSAppToV5 converts v4 saas_app to v5 format.
func transformV4SaaSAppToV5(ctx context.Context, v4 V4SaaSAppModel) (*ZeroTrustAccessApplicationSaaSAppModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	v5 := &ZeroTrustAccessApplicationSaaSAppModel{
		AuthType:                      v4.AuthType,
		ConsumerServiceURL:            v4.ConsumerServiceURL,
		SPEntityID:                    v4.SPEntityID,
		IdPEntityID:                   v4.IdPEntityID,
		PublicKey:                     v4.PublicKey,
		NameIDFormat:                  v4.NameIDFormat,
		NameIDTransformJsonata:        v4.NameIDTransformJsonata,
		SAMLAttributeTransformJsonata: v4.SAMLAttrTransformJsonata,
		DefaultRelayState:             v4.DefaultRelayState,
		SSOEndpoint:                   v4.SSOEndpoint,
		AppLauncherURL:                v4.AppLauncherURL,
		ClientID:                      v4.ClientID,
		ClientSecret:                  v4.ClientSecret,
		AccessTokenLifetime:           v4.AccessTokenLifetime,
		AllowPKCEWithoutClientSecret:  v4.AllowPKCEWithoutSecret,
		GroupFilterRegex:              v4.GroupFilterRegex,
	}

	// GrantTypes: v4 Set -> v5 *[]types.String
	if !v4.GrantTypes.IsNull() && !v4.GrantTypes.IsUnknown() {
		grantTypes, d := setToStringSlice(ctx, v4.GrantTypes)
		diags.Append(d...)
		if !diags.HasError() {
			v5.GrantTypes = grantTypes
		}
	}

	// RedirectURIs: v4 Set -> v5 *[]types.String
	if !v4.RedirectURIs.IsNull() && !v4.RedirectURIs.IsUnknown() {
		redirectURIs, d := setToStringSlice(ctx, v4.RedirectURIs)
		diags.Append(d...)
		if !diags.HasError() {
			v5.RedirectURIs = redirectURIs
		}
	}

	// Scopes: v4 Set -> v5 *[]types.String
	if !v4.Scopes.IsNull() && !v4.Scopes.IsUnknown() {
		scopes, d := setToStringSlice(ctx, v4.Scopes)
		diags.Append(d...)
		if !diags.HasError() {
			v5.Scopes = scopes
		}
	}

	// CustomAttributes: v4 custom_attribute -> v5 custom_attributes
	if len(v4.CustomAttributes) > 0 {
		customAttrs := make([]*ZeroTrustAccessApplicationSaaSAppCustomAttributesModel, 0, len(v4.CustomAttributes))
		for _, attr := range v4.CustomAttributes {
			v5Attr := &ZeroTrustAccessApplicationSaaSAppCustomAttributesModel{
				Name:         attr.Name,
				FriendlyName: attr.FriendlyName,
				NameFormat:   attr.NameFormat,
				Required:     attr.Required,
			}
			if len(attr.Source) > 0 {
				v5Attr.Source = transformV4CustomAttrSourceToV5(ctx, attr.Source[0])
			}
			customAttrs = append(customAttrs, v5Attr)
		}
		v5.CustomAttributes = &customAttrs
	}

	// CustomClaims: v4 custom_claim -> v5 custom_claims
	if len(v4.CustomClaims) > 0 {
		customClaims := make([]*ZeroTrustAccessApplicationSaaSAppCustomClaimsModel, 0, len(v4.CustomClaims))
		for _, claim := range v4.CustomClaims {
			v5Claim := &ZeroTrustAccessApplicationSaaSAppCustomClaimsModel{
				Name:     claim.Name,
				Required: claim.Required,
				Scope:    claim.Scope,
			}
			if len(claim.Source) > 0 {
				v5Claim.Source = transformV4CustomClaimSourceToV5(ctx, claim.Source[0])
			}
			customClaims = append(customClaims, v5Claim)
		}
		v5.CustomClaims = &customClaims
	}

	// HybridAndImplicitOptions: v4 []V4HybridOptionsModel -> v5 *Model
	if len(v4.HybridAndImplicitOptions) > 0 {
		v5.HybridAndImplicitOptions = &ZeroTrustAccessApplicationSaaSAppHybridAndImplicitOptionsModel{
			ReturnAccessTokenFromAuthorizationEndpoint: v4.HybridAndImplicitOptions[0].ReturnAccessTokenFromAuthEndpoint,
			ReturnIDTokenFromAuthorizationEndpoint:     v4.HybridAndImplicitOptions[0].ReturnIDTokenFromAuthEndpoint,
		}
	}

	// RefreshTokenOptions: v4 []V4RefreshTokenModel -> v5 *Model
	if len(v4.RefreshTokenOptions) > 0 {
		v5.RefreshTokenOptions = &ZeroTrustAccessApplicationSaaSAppRefreshTokenOptionsModel{
			Lifetime: v4.RefreshTokenOptions[0].Lifetime,
		}
	}

	return v5, diags
}

func transformV4CustomAttrSourceToV5(ctx context.Context, v4 V4CustomAttributeSourceModel) *ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceModel {
	v5 := &ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceModel{
		Name: v4.Name,
	}
	// name_by_idp: v4 map[string]string -> v5 *[]*NameByIdPModel
	if len(v4.NameByIdP) > 0 {
		nameByIdP := make([]*ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceNameByIdPModel, 0, len(v4.NameByIdP))
		for idpID, sourceName := range v4.NameByIdP {
			nameByIdP = append(nameByIdP, &ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceNameByIdPModel{
				IdPID:      types.StringValue(idpID),
				SourceName: types.StringValue(sourceName),
			})
		}
		v5.NameByIdP = &nameByIdP
	}
	return v5
}

func transformV4CustomClaimSourceToV5(ctx context.Context, v4 V4CustomClaimSourceModel) *ZeroTrustAccessApplicationSaaSAppCustomClaimsSourceModel {
	v5 := &ZeroTrustAccessApplicationSaaSAppCustomClaimsSourceModel{
		Name: v4.Name,
	}
	// name_by_idp: v4 map[string]types.String -> v5 *map[string]types.String
	if len(v4.NameByIdP) > 0 {
		v5.NameByIdP = &v4.NameByIdP
	}
	return v5
}

// transformV4SCIMConfigToV5 converts v4 scim_config to v5 format.
func transformV4SCIMConfigToV5(ctx context.Context, v4 V4SCIMConfigModel) (*ZeroTrustAccessApplicationSCIMConfigModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	v5 := &ZeroTrustAccessApplicationSCIMConfigModel{
		IdPUID:             v4.IdPUID,
		RemoteURI:          v4.RemoteURI,
		Enabled:            v4.Enabled,
		DeactivateOnDelete: v4.DeactivateOnDelete,
	}

	// Authentication: v4 []V4SCIMAuthModel -> v5 *Model
	if len(v4.Authentication) > 0 {
		auth := v4.Authentication[0]
		v5Auth := &ZeroTrustAccessApplicationSCIMConfigAuthenticationModel{
			Scheme:           auth.Scheme,
			User:             auth.User,
			Password:         auth.Password,
			Token:            auth.Token,
			AuthorizationURL: auth.AuthorizationURL,
			ClientID:         auth.ClientID,
			ClientSecret:     auth.ClientSecret,
			TokenURL:         auth.TokenURL,
		}
		// Scopes: v4 Set -> v5 *[]types.String
		if !auth.Scopes.IsNull() && !auth.Scopes.IsUnknown() {
			scopes, d := setToStringSlice(ctx, auth.Scopes)
			diags.Append(d...)
			if !diags.HasError() {
				v5Auth.Scopes = scopes
			}
		}
		v5.Authentication = v5Auth
	}

	// Mappings: v4 []V4SCIMMappingsModel -> v5 *[]*Model
	if len(v4.Mappings) > 0 {
		mappings := make([]*ZeroTrustAccessApplicationSCIMConfigMappingsModel, 0, len(v4.Mappings))
		for _, m := range v4.Mappings {
			v5Mapping := &ZeroTrustAccessApplicationSCIMConfigMappingsModel{
				Schema:           m.Schema,
				Enabled:          m.Enabled,
				Filter:           m.Filter,
				TransformJsonata: m.TransformJsonata,
				Strictness:       m.Strictness,
			}
			if len(m.Operations) > 0 {
				v5Mapping.Operations = &ZeroTrustAccessApplicationSCIMConfigMappingsOperationsModel{
					Create: m.Operations[0].Create,
					Update: m.Operations[0].Update,
					Delete: m.Operations[0].Delete,
				}
			}
			mappings = append(mappings, v5Mapping)
		}
		v5.Mappings = &mappings
	}

	return v5, diags
}

// transformV4LandingPageToV5 converts v4 landing_page_design to v5 format.
func transformV4LandingPageToV5(ctx context.Context, v4 V4LandingPageDesignModel) (customfield.NestedObject[ZeroTrustAccessApplicationLandingPageDesignModel], diag.Diagnostics) {
	var diags diag.Diagnostics

	v5Model := &ZeroTrustAccessApplicationLandingPageDesignModel{
		ButtonColor:     v4.ButtonColor,
		ButtonTextColor: v4.ButtonTextColor,
		ImageURL:        v4.ImageURL,
		Message:         v4.Message,
		Title:           v4.Title,
	}

	result, d := customfield.NewObject(ctx, v5Model)
	diags.Append(d...)
	return result, diags
}

// transformV4DestinationsToV5 converts v4 destinations to v5 format.
func transformV4DestinationsToV5(ctx context.Context, v4 []V4DestinationsModel) (customfield.NestedObjectList[ZeroTrustAccessApplicationDestinationsModel], diag.Diagnostics) {
	var diags diag.Diagnostics

	destinations := make([]ZeroTrustAccessApplicationDestinationsModel, 0, len(v4))
	for _, d := range v4 {
		dest := ZeroTrustAccessApplicationDestinationsModel{
			Type:       d.Type,
			URI:        d.URI,
			Hostname:   d.Hostname,
			CIDR:       d.CIDR,
			PortRange:  d.PortRange,
			VnetID:     d.VnetID,
			L4Protocol: d.L4Protocol,
		}
		destinations = append(destinations, dest)
	}

	result, d := customfield.NewObjectList(ctx, destinations)
	diags.Append(d...)
	return result, diags
}

// transformV4TargetCriteriaToV5 converts v4 target_criteria to v5 format.
func transformV4TargetCriteriaToV5(ctx context.Context, v4 []V4TargetCriteriaModel) (*[]*ZeroTrustAccessApplicationTargetCriteriaModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	criteria := make([]*ZeroTrustAccessApplicationTargetCriteriaModel, 0, len(v4))
	for _, c := range v4 {
		v5Criteria := &ZeroTrustAccessApplicationTargetCriteriaModel{
			Port:     c.Port,
			Protocol: c.Protocol,
		}

		// TargetAttributes: v4 []V4TargetAttributesModel -> v5 *map[string]*[]types.String
		if len(c.TargetAttributes) > 0 {
			targetAttrs := make(map[string]*[]types.String)
			for _, attr := range c.TargetAttributes {
				if !attr.Name.IsNull() && !attr.Name.IsUnknown() {
					var values []string
					diags.Append(attr.Values.ElementsAs(ctx, &values, false)...)
					if diags.HasError() {
						return nil, diags
					}
					strValues := make([]types.String, len(values))
					for i, v := range values {
						strValues[i] = types.StringValue(v)
					}
					targetAttrs[attr.Name.ValueString()] = &strValues
				}
			}
			v5Criteria.TargetAttributes = &targetAttrs
		}

		criteria = append(criteria, v5Criteria)
	}

	return &criteria, diags
}
