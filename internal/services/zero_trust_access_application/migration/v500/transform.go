package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/migrations"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// normalizeBoolFalseToNull converts false boolean values to null.
// The v5 provider's API treats false and null as equivalent for these optional boolean fields.
// By normalizing false to null during migration, we prevent drift after the v5 provider refreshes state.
func normalizeBoolFalseToNull(b types.Bool) types.Bool {
	if b.IsNull() || b.IsUnknown() {
		return b
	}
	if !b.ValueBool() {
		// false -> null (they are semantically equivalent)
		return types.BoolNull()
	}
	return b
}

// Transform converts a v4 cloudflare_access_application state to v5 cloudflare_zero_trust_access_application state.
func Transform(ctx context.Context, v4 SourceAccessApplicationModel) (*TargetAccessApplicationModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	v5 := &TargetAccessApplicationModel{
		ID:                       v4.ID,
		AccountID:                v4.AccountID,
		ZoneID:                   v4.ZoneID,
		Name:                     v4.Name,
		Domain:                   v4.Domain,
		SessionDuration:          v4.SessionDuration,
		AutoRedirectToIdentity:   normalizeBoolFalseToNull(v4.AutoRedirectToIdentity),
		EnableBindingCookie:      normalizeBoolFalseToNull(v4.EnableBindingCookie),
		HTTPOnlyCookieAttribute:  normalizeBoolFalseToNull(v4.HTTPOnlyCookieAttribute),
		SameSiteCookieAttribute:  migrations.FalseyStringToNull(v4.SameSiteCookieAttribute),
		LogoURL:                  migrations.FalseyStringToNull(v4.LogoURL),
		SkipInterstitial:         normalizeBoolFalseToNull(v4.SkipInterstitial),
		AppLauncherVisible:       v4.AppLauncherVisible,
		ServiceAuth401Redirect:   normalizeBoolFalseToNull(v4.ServiceAuth401Redirect),
		CustomDenyMessage:        migrations.FalseyStringToNull(v4.CustomDenyMessage),
		CustomDenyURL:            migrations.FalseyStringToNull(v4.CustomDenyURL),
		CustomNonIdentityDenyURL: migrations.FalseyStringToNull(v4.CustomNonIdentityDenyURL),
		OptionsPreflightBypass:   normalizeBoolFalseToNull(v4.OptionsPreflightBypass),
		PathCookieAttribute:      normalizeBoolFalseToNull(v4.PathCookieAttribute),
		AUD:                      v4.AUD,
		AppLauncherLogoURL:       migrations.FalseyStringToNull(v4.AppLauncherLogoURL),
		HeaderBgColor:            migrations.FalseyStringToNull(v4.HeaderBgColor),
		BgColor:                  migrations.FalseyStringToNull(v4.BgColor),
		SkipAppLauncherLoginPage: normalizeBoolFalseToNull(v4.SkipAppLauncherLoginPage),
		AllowAuthenticateViaWARP: migrations.FalseyBoolToNull(v4.AllowAuthenticateViaWARP),
	}

	// Type: v4 may not have this, v5 defaults to "self_hosted"
	if !v4.Type.IsNull() && !v4.Type.IsUnknown() && v4.Type.ValueString() != "" {
		v5.Type = v4.Type
	} else {
		v5.Type = types.StringValue("self_hosted")
	}

	// AllowedIdPs: v4 Set -> v5 *[]types.String
	// Only set if non-empty; empty arrays should remain null to prevent drift
	if !v4.AllowedIdPs.IsNull() && !v4.AllowedIdPs.IsUnknown() && len(v4.AllowedIdPs.Elements()) > 0 {
		allowedIdPs, d := setToStringSlice(ctx, v4.AllowedIdPs)
		diags.Append(d...)
		if !diags.HasError() {
			v5.AllowedIdPs = allowedIdPs
		}
	}

	// Tags: v4 Set -> v5 customfield.Set
	// Only set if non-empty; empty sets should remain null to prevent drift
	if !v4.Tags.IsNull() && !v4.Tags.IsUnknown() && len(v4.Tags.Elements()) > 0 {
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
		if !diags.HasError() && customPages != nil {
			v5.CustomPages = customPages
		}
	}

	// CORSHeaders: v4 []SourceCORSHeadersModel (list block) -> v5 *TargetCORSHeadersModel (object)
	if len(v4.CORSHeaders) > 0 {
		corsHeaders, d := transformCORSHeaders(ctx, v4.CORSHeaders[0])
		diags.Append(d...)
		if !diags.HasError() {
			v5.CORSHeaders = corsHeaders
		}
	}

	// SaaSApp: v4 []SourceSaaSAppModel (list block) -> v5 *TargetSaaSAppModel (object)
	if len(v4.SaaSApp) > 0 {
		saasApp, d := transformSaaSApp(ctx, v4.SaaSApp[0])
		diags.Append(d...)
		if !diags.HasError() {
			v5.SaaSApp = saasApp
		}
	}

	// SCIMConfig: v4 []SourceSCIMConfigModel (list block) -> v5 *TargetSCIMConfigModel (object)
	if len(v4.SCIMConfig) > 0 {
		scimConfig, d := transformSCIMConfig(ctx, v4.SCIMConfig[0])
		diags.Append(d...)
		if !diags.HasError() {
			v5.SCIMConfig = scimConfig
		}
	}

	// LandingPageDesign: v4 []SourceLandingPageDesignModel (list block) -> v5 customfield.NestedObject (object)
	if len(v4.LandingPageDesign) > 0 {
		landingPage, d := transformLandingPage(ctx, v4.LandingPageDesign[0])
		diags.Append(d...)
		if !diags.HasError() {
			v5.LandingPageDesign = landingPage
		}
	}

	// FooterLinks: v4 []SourceFooterLinksModel -> v5 *[]*TargetFooterLinksModel
	if len(v4.FooterLinks) > 0 {
		footerLinks := make([]*TargetFooterLinksModel, 0, len(v4.FooterLinks))
		for _, link := range v4.FooterLinks {
			footerLinks = append(footerLinks, &TargetFooterLinksModel{
				Name: link.Name,
				URL:  link.URL,
			})
		}
		v5.FooterLinks = &footerLinks
	}

	// Destinations: v4 []SourceDestinationsModel -> v5 customfield.NestedObjectList
	if len(v4.Destinations) > 0 {
		destinations, d := transformDestinations(ctx, v4.Destinations)
		diags.Append(d...)
		if !diags.HasError() {
			v5.Destinations = destinations
		}
	}

	// TargetCriteria: v4 []SourceTargetCriteriaModel -> v5 *[]*TargetTargetCriteriaModel
	if len(v4.TargetCriteria) > 0 {
		targetCriteria, d := transformTargetCriteria(ctx, v4.TargetCriteria)
		diags.Append(d...)
		if !diags.HasError() {
			v5.TargetCriteria = targetCriteria
		}
	}

	// Policies: v4 []types.String (string array) -> v5 *[]TargetPoliciesModel (object array with id/precedence)
	if len(v4.Policies) > 0 {
		policies := make([]TargetPoliciesModel, 0, len(v4.Policies))
		for i, policyID := range v4.Policies {
			if !policyID.IsNull() && !policyID.IsUnknown() {
				policies = append(policies, TargetPoliciesModel{
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

func setToStringSlice(ctx context.Context, set types.Set) (*[]types.String, diag.Diagnostics) {
	var diags diag.Diagnostics
	var values []string
	diags.Append(set.ElementsAs(ctx, &values, false)...)
	if diags.HasError() {
		return nil, diags
	}
	if len(values) == 0 {
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

// transformCORSHeaders converts v4 cors_headers to v5 format.
func transformCORSHeaders(ctx context.Context, v4 SourceCORSHeadersModel) (*TargetCORSHeadersModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	v5 := &TargetCORSHeadersModel{
		AllowAllHeaders:  normalizeBoolFalseToNull(v4.AllowAllHeaders),
		AllowAllMethods:  normalizeBoolFalseToNull(v4.AllowAllMethods),
		AllowAllOrigins:  normalizeBoolFalseToNull(v4.AllowAllOrigins),
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
		if !diags.HasError() && headers != nil {
			v5.AllowedHeaders = headers
		}
	}

	// AllowedMethods: v4 Set -> v5 *[]types.String
	if !v4.AllowedMethods.IsNull() && !v4.AllowedMethods.IsUnknown() {
		methods, d := setToStringSlice(ctx, v4.AllowedMethods)
		diags.Append(d...)
		if !diags.HasError() && methods != nil {
			v5.AllowedMethods = methods
		}
	}

	// AllowedOrigins: v4 Set -> v5 *[]types.String
	if !v4.AllowedOrigins.IsNull() && !v4.AllowedOrigins.IsUnknown() {
		origins, d := setToStringSlice(ctx, v4.AllowedOrigins)
		diags.Append(d...)
		if !diags.HasError() && origins != nil {
			v5.AllowedOrigins = origins
		}
	}

	return v5, diags
}

// transformSaaSApp converts v4 saas_app to v5 format.
func transformSaaSApp(ctx context.Context, v4 SourceSaaSAppModel) (*TargetSaaSAppModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	v5 := &TargetSaaSAppModel{
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
		AllowPKCEWithoutClientSecret:  normalizeBoolFalseToNull(v4.AllowPKCEWithoutSecret),
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
		customAttrs := make([]*TargetCustomAttributesModel, 0, len(v4.CustomAttributes))
		for _, attr := range v4.CustomAttributes {
			v5Attr := &TargetCustomAttributesModel{
				Name:         attr.Name,
				FriendlyName: attr.FriendlyName,
				NameFormat:   attr.NameFormat,
				Required:     attr.Required,
			}
			if len(attr.Source) > 0 {
				v5Attr.Source = transformCustomAttrSource(ctx, attr.Source[0])
			}
			customAttrs = append(customAttrs, v5Attr)
		}
		v5.CustomAttributes = &customAttrs
	}

	// CustomClaims: v4 custom_claim -> v5 custom_claims
	if len(v4.CustomClaims) > 0 {
		customClaims := make([]*TargetCustomClaimsModel, 0, len(v4.CustomClaims))
		for _, claim := range v4.CustomClaims {
			v5Claim := &TargetCustomClaimsModel{
				Name:     claim.Name,
				Required: claim.Required,
				Scope:    claim.Scope,
			}
			if len(claim.Source) > 0 {
				v5Claim.Source = transformCustomClaimSource(ctx, claim.Source[0])
			}
			customClaims = append(customClaims, v5Claim)
		}
		v5.CustomClaims = &customClaims
	}

	// HybridAndImplicitOptions: v4 []SourceHybridOptionsModel -> v5 *Model
	if len(v4.HybridAndImplicitOptions) > 0 {
		v5.HybridAndImplicitOptions = &TargetHybridAndImplicitOptionsModel{
			ReturnAccessTokenFromAuthorizationEndpoint: v4.HybridAndImplicitOptions[0].ReturnAccessTokenFromAuthEndpoint,
			ReturnIDTokenFromAuthorizationEndpoint:     v4.HybridAndImplicitOptions[0].ReturnIDTokenFromAuthEndpoint,
		}
	}

	// RefreshTokenOptions: v4 []SourceRefreshTokenModel -> v5 *Model
	if len(v4.RefreshTokenOptions) > 0 {
		v5.RefreshTokenOptions = &TargetRefreshTokenOptionsModel{
			Lifetime: v4.RefreshTokenOptions[0].Lifetime,
		}
	}

	return v5, diags
}

func transformCustomAttrSource(ctx context.Context, v4 SourceCustomAttributeSourceModel) *TargetCustomAttributesSourceModel {
	v5 := &TargetCustomAttributesSourceModel{
		Name: v4.Name,
	}
	// name_by_idp: v4 map[string]string -> v5 *[]*TargetNameByIdPModel
	if len(v4.NameByIdP) > 0 {
		nameByIdP := make([]*TargetNameByIdPModel, 0, len(v4.NameByIdP))
		for idpID, sourceName := range v4.NameByIdP {
			nameByIdP = append(nameByIdP, &TargetNameByIdPModel{
				IdPID:      types.StringValue(idpID),
				SourceName: types.StringValue(sourceName),
			})
		}
		v5.NameByIdP = &nameByIdP
	}
	return v5
}

func transformCustomClaimSource(ctx context.Context, v4 SourceCustomClaimSourceModel) *TargetCustomClaimsSourceModel {
	v5 := &TargetCustomClaimsSourceModel{
		Name: v4.Name,
	}
	// name_by_idp: v4 map[string]types.String -> v5 *map[string]types.String
	if len(v4.NameByIdP) > 0 {
		v5.NameByIdP = &v4.NameByIdP
	}
	return v5
}

// transformSCIMConfig converts v4 scim_config to v5 format.
func transformSCIMConfig(ctx context.Context, v4 SourceSCIMConfigModel) (*TargetSCIMConfigModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	v5 := &TargetSCIMConfigModel{
		IdPUID:             v4.IdPUID,
		RemoteURI:          v4.RemoteURI,
		Enabled:            v4.Enabled,
		DeactivateOnDelete: v4.DeactivateOnDelete,
	}

	// Authentication: v4 []SourceSCIMAuthModel -> v5 *Model
	if len(v4.Authentication) > 0 {
		auth := v4.Authentication[0]
		v5Auth := &TargetSCIMAuthenticationModel{
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

	// Mappings: v4 []SourceSCIMMappingsModel -> v5 *[]*Model
	if len(v4.Mappings) > 0 {
		mappings := make([]*TargetSCIMMappingsModel, 0, len(v4.Mappings))
		for _, m := range v4.Mappings {
			v5Mapping := &TargetSCIMMappingsModel{
				Schema:           m.Schema,
				Enabled:          m.Enabled,
				Filter:           m.Filter,
				TransformJsonata: m.TransformJsonata,
				Strictness:       m.Strictness,
			}
			if len(m.Operations) > 0 {
				v5Mapping.Operations = &TargetSCIMOperationsModel{
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

// transformLandingPage converts v4 landing_page_design to v5 format.
func transformLandingPage(ctx context.Context, v4 SourceLandingPageDesignModel) (customfield.NestedObject[TargetLandingPageDesignModel], diag.Diagnostics) {
	var diags diag.Diagnostics

	v5Model := &TargetLandingPageDesignModel{
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

// transformDestinations converts v4 destinations to v5 format.
func transformDestinations(ctx context.Context, v4 []SourceDestinationsModel) (customfield.NestedObjectList[TargetDestinationsModel], diag.Diagnostics) {
	var diags diag.Diagnostics

	destinations := make([]TargetDestinationsModel, 0, len(v4))
	for _, d := range v4 {
		dest := TargetDestinationsModel{
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

// transformTargetCriteria converts v4 target_criteria to v5 format.
func transformTargetCriteria(ctx context.Context, v4 []SourceTargetCriteriaModel) (*[]*TargetTargetCriteriaModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	criteria := make([]*TargetTargetCriteriaModel, 0, len(v4))
	for _, c := range v4 {
		v5Criteria := &TargetTargetCriteriaModel{
			Port:     c.Port,
			Protocol: c.Protocol,
		}

		// TargetAttributes: v4 []SourceTargetAttributesModel -> v5 *map[string]*[]types.String
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
