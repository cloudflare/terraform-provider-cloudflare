// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_application

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessApplicationResultEnvelope struct {
	Result AccessApplicationModel `json:"result,computed"`
}

type AccessApplicationModel struct {
	ID                       types.String                       `tfsdk:"id" json:"id,computed"`
	AccountID                types.String                       `tfsdk:"account_id" path:"account_id"`
	ZoneID                   types.String                       `tfsdk:"zone_id" path:"zone_id"`
	Domain                   types.String                       `tfsdk:"domain" json:"domain"`
	Type                     types.String                       `tfsdk:"type" json:"type"`
	AllowAuthenticateViaWARP types.Bool                         `tfsdk:"allow_authenticate_via_warp" json:"allow_authenticate_via_warp"`
	AllowedIdPs              types.String                       `tfsdk:"allowed_idps" json:"allowed_idps"`
	AppLauncherVisible       types.Bool                         `tfsdk:"app_launcher_visible" json:"app_launcher_visible"`
	AutoRedirectToIdentity   types.Bool                         `tfsdk:"auto_redirect_to_identity" json:"auto_redirect_to_identity"`
	CORSHeaders              *AccessApplicationCORSHeadersModel `tfsdk:"cors_headers" json:"cors_headers"`
	CustomDenyMessage        types.String                       `tfsdk:"custom_deny_message" json:"custom_deny_message"`
	CustomDenyURL            types.String                       `tfsdk:"custom_deny_url" json:"custom_deny_url"`
	CustomNonIdentityDenyURL types.String                       `tfsdk:"custom_non_identity_deny_url" json:"custom_non_identity_deny_url"`
	CustomPages              types.String                       `tfsdk:"custom_pages" json:"custom_pages"`
	EnableBindingCookie      types.Bool                         `tfsdk:"enable_binding_cookie" json:"enable_binding_cookie"`
	HTTPOnlyCookieAttribute  types.Bool                         `tfsdk:"http_only_cookie_attribute" json:"http_only_cookie_attribute"`
	LogoURL                  types.String                       `tfsdk:"logo_url" json:"logo_url"`
	Name                     types.String                       `tfsdk:"name" json:"name"`
	OptionsPreflightBypass   types.Bool                         `tfsdk:"options_preflight_bypass" json:"options_preflight_bypass"`
	PathCookieAttribute      types.Bool                         `tfsdk:"path_cookie_attribute" json:"path_cookie_attribute"`
	SameSiteCookieAttribute  types.String                       `tfsdk:"same_site_cookie_attribute" json:"same_site_cookie_attribute"`
	SelfHostedDomains        types.String                       `tfsdk:"self_hosted_domains" json:"self_hosted_domains"`
	ServiceAuth401Redirect   types.Bool                         `tfsdk:"service_auth_401_redirect" json:"service_auth_401_redirect"`
	SessionDuration          types.String                       `tfsdk:"session_duration" json:"session_duration"`
	SkipInterstitial         types.Bool                         `tfsdk:"skip_interstitial" json:"skip_interstitial"`
	Tags                     types.String                       `tfsdk:"tags" json:"tags"`
	SaaSApp                  *AccessApplicationSaaSAppModel     `tfsdk:"saas_app" json:"saas_app"`
}

type AccessApplicationCORSHeadersModel struct {
	AllowAllHeaders  types.Bool    `tfsdk:"allow_all_headers" json:"allow_all_headers"`
	AllowAllMethods  types.Bool    `tfsdk:"allow_all_methods" json:"allow_all_methods"`
	AllowAllOrigins  types.Bool    `tfsdk:"allow_all_origins" json:"allow_all_origins"`
	AllowCredentials types.Bool    `tfsdk:"allow_credentials" json:"allow_credentials"`
	AllowedHeaders   types.String  `tfsdk:"allowed_headers" json:"allowed_headers"`
	AllowedMethods   types.String  `tfsdk:"allowed_methods" json:"allowed_methods"`
	AllowedOrigins   types.String  `tfsdk:"allowed_origins" json:"allowed_origins"`
	MaxAge           types.Float64 `tfsdk:"max_age" json:"max_age"`
}

type AccessApplicationSaaSAppModel struct {
	AuthType                      types.String                                   `tfsdk:"auth_type" json:"auth_type"`
	ConsumerServiceURL            types.String                                   `tfsdk:"consumer_service_url" json:"consumer_service_url"`
	CreatedAt                     types.String                                   `tfsdk:"created_at" json:"created_at,computed"`
	CustomAttributes              *AccessApplicationSaaSAppCustomAttributesModel `tfsdk:"custom_attributes" json:"custom_attributes"`
	DefaultRelayState             types.String                                   `tfsdk:"default_relay_state" json:"default_relay_state"`
	IdPEntityID                   types.String                                   `tfsdk:"idp_entity_id" json:"idp_entity_id"`
	NameIDFormat                  types.String                                   `tfsdk:"name_id_format" json:"name_id_format"`
	NameIDTransformJsonata        types.String                                   `tfsdk:"name_id_transform_jsonata" json:"name_id_transform_jsonata"`
	PublicKey                     types.String                                   `tfsdk:"public_key" json:"public_key"`
	SAMLAttributeTransformJsonata types.String                                   `tfsdk:"saml_attribute_transform_jsonata" json:"saml_attribute_transform_jsonata"`
	SPEntityID                    types.String                                   `tfsdk:"sp_entity_id" json:"sp_entity_id"`
	SSOEndpoint                   types.String                                   `tfsdk:"sso_endpoint" json:"sso_endpoint"`
	UpdatedAt                     types.String                                   `tfsdk:"updated_at" json:"updated_at,computed"`
	AppLauncherURL                types.String                                   `tfsdk:"app_launcher_url" json:"app_launcher_url"`
	ClientID                      types.String                                   `tfsdk:"client_id" json:"client_id"`
	ClientSecret                  types.String                                   `tfsdk:"client_secret" json:"client_secret"`
	CustomClaims                  *AccessApplicationSaaSAppCustomClaimsModel     `tfsdk:"custom_claims" json:"custom_claims"`
	GrantTypes                    types.String                                   `tfsdk:"grant_types" json:"grant_types"`
	GroupFilterRegex              types.String                                   `tfsdk:"group_filter_regex" json:"group_filter_regex"`
	RedirectURIs                  types.String                                   `tfsdk:"redirect_uris" json:"redirect_uris"`
	Scopes                        types.String                                   `tfsdk:"scopes" json:"scopes"`
}

type AccessApplicationSaaSAppCustomAttributesModel struct {
	FriendlyName types.String                                         `tfsdk:"friendly_name" json:"friendly_name"`
	Name         types.String                                         `tfsdk:"name" json:"name"`
	NameFormat   types.String                                         `tfsdk:"name_format" json:"name_format"`
	Required     types.Bool                                           `tfsdk:"required" json:"required"`
	Source       *AccessApplicationSaaSAppCustomAttributesSourceModel `tfsdk:"source" json:"source"`
}

type AccessApplicationSaaSAppCustomAttributesSourceModel struct {
	Name      types.String            `tfsdk:"name" json:"name"`
	NameByIdP map[string]types.String `tfsdk:"name_by_idp" json:"name_by_idp"`
}

type AccessApplicationSaaSAppCustomClaimsModel struct {
	Name      types.String                                     `tfsdk:"name" json:"name"`
	NameByIdP map[string]types.String                          `tfsdk:"name_by_idp" json:"name_by_idp"`
	Required  types.Bool                                       `tfsdk:"required" json:"required"`
	Scope     types.String                                     `tfsdk:"scope" json:"scope"`
	Source    *AccessApplicationSaaSAppCustomClaimsSourceModel `tfsdk:"source" json:"source"`
}

type AccessApplicationSaaSAppCustomClaimsSourceModel struct {
	Name types.String `tfsdk:"name" json:"name"`
}
