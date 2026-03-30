package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

// SourceAccessApplicationModel represents the v4 cloudflare_access_application state structure.
// This is used by MoveState and UpgradeFromV4 to parse the source state from v4 provider.
type SourceAccessApplicationModel struct {
	ID                       types.String `tfsdk:"id"`
	AccountID                types.String `tfsdk:"account_id"`
	ZoneID                   types.String `tfsdk:"zone_id"`
	Name                     types.String `tfsdk:"name"`
	Domain                   types.String `tfsdk:"domain"`
	Type                     types.String `tfsdk:"type"`
	SessionDuration          types.String `tfsdk:"session_duration"`
	AutoRedirectToIdentity   types.Bool   `tfsdk:"auto_redirect_to_identity"`
	EnableBindingCookie      types.Bool   `tfsdk:"enable_binding_cookie"`
	HTTPOnlyCookieAttribute  types.Bool   `tfsdk:"http_only_cookie_attribute"`
	SameSiteCookieAttribute  types.String `tfsdk:"same_site_cookie_attribute"`
	LogoURL                  types.String `tfsdk:"logo_url"`
	SkipInterstitial         types.Bool   `tfsdk:"skip_interstitial"`
	AppLauncherVisible       types.Bool   `tfsdk:"app_launcher_visible"`
	ServiceAuth401Redirect   types.Bool   `tfsdk:"service_auth_401_redirect"`
	CustomDenyMessage        types.String `tfsdk:"custom_deny_message"`
	CustomDenyURL            types.String `tfsdk:"custom_deny_url"`
	CustomNonIdentityDenyURL types.String `tfsdk:"custom_non_identity_deny_url"`
	AllowedIdPs              types.Set    `tfsdk:"allowed_idps"`
	Tags                     types.Set    `tfsdk:"tags"`
	SelfHostedDomains        types.Set    `tfsdk:"self_hosted_domains"`
	CustomPages              types.Set    `tfsdk:"custom_pages"`
	OptionsPreflightBypass   types.Bool   `tfsdk:"options_preflight_bypass"`
	PathCookieAttribute      types.Bool   `tfsdk:"path_cookie_attribute"`
	AUD                      types.String `tfsdk:"aud"`

	// Additional fields from v4 schema
	AppLauncherLogoURL       types.String `tfsdk:"app_launcher_logo_url"`
	HeaderBgColor            types.String `tfsdk:"header_bg_color"`
	BgColor                  types.String `tfsdk:"bg_color"`
	SkipAppLauncherLoginPage types.Bool   `tfsdk:"skip_app_launcher_login_page"`
	AllowAuthenticateViaWARP types.Bool   `tfsdk:"allow_authenticate_via_warp"`

	// v4 stores these as simple string arrays
	Policies []types.String `tfsdk:"policies"`

	// v4 stores these as list blocks with MaxItems:1 (arrays in state)
	CORSHeaders       []SourceCORSHeadersModel       `tfsdk:"cors_headers"`
	SaaSApp           []SourceSaaSAppModel           `tfsdk:"saas_app"`
	SCIMConfig        []SourceSCIMConfigModel        `tfsdk:"scim_config"`
	LandingPageDesign []SourceLandingPageDesignModel `tfsdk:"landing_page_design"`
	FooterLinks       []SourceFooterLinksModel       `tfsdk:"footer_links"`
	Destinations      []SourceDestinationsModel      `tfsdk:"destinations"`
	TargetCriteria    []SourceTargetCriteriaModel    `tfsdk:"target_criteria"`

	// Deprecated/removed fields
	DomainType types.String `tfsdk:"domain_type"`
}

// SourceCORSHeadersModel represents the v4 cors_headers block structure.
type SourceCORSHeadersModel struct {
	AllowAllHeaders  types.Bool  `tfsdk:"allow_all_headers"`
	AllowAllMethods  types.Bool  `tfsdk:"allow_all_methods"`
	AllowAllOrigins  types.Bool  `tfsdk:"allow_all_origins"`
	AllowCredentials types.Bool  `tfsdk:"allow_credentials"`
	AllowedHeaders   types.Set   `tfsdk:"allowed_headers"`
	AllowedMethods   types.Set   `tfsdk:"allowed_methods"`
	AllowedOrigins   types.Set   `tfsdk:"allowed_origins"`
	MaxAge           types.Int64 `tfsdk:"max_age"`
}

// SourceSaaSAppModel represents the v4 saas_app block structure.
type SourceSaaSAppModel struct {
	AuthType                 types.String                 `tfsdk:"auth_type"`
	ConsumerServiceURL       types.String                 `tfsdk:"consumer_service_url"`
	SPEntityID               types.String                 `tfsdk:"sp_entity_id"`
	IdPEntityID              types.String                 `tfsdk:"idp_entity_id"`
	PublicKey                types.String                 `tfsdk:"public_key"`
	NameIDFormat             types.String                 `tfsdk:"name_id_format"`
	NameIDTransformJsonata   types.String                 `tfsdk:"name_id_transform_jsonata"`
	SAMLAttrTransformJsonata types.String                 `tfsdk:"saml_attribute_transform_jsonata"`
	DefaultRelayState        types.String                 `tfsdk:"default_relay_state"`
	SSOEndpoint              types.String                 `tfsdk:"sso_endpoint"`
	AppLauncherURL           types.String                 `tfsdk:"app_launcher_url"`
	ClientID                 types.String                 `tfsdk:"client_id"`
	ClientSecret             types.String                 `tfsdk:"client_secret"`
	AccessTokenLifetime      types.String                 `tfsdk:"access_token_lifetime"`
	AllowPKCEWithoutSecret   types.Bool                   `tfsdk:"allow_pkce_without_client_secret"`
	GroupFilterRegex         types.String                 `tfsdk:"group_filter_regex"`
	GrantTypes               types.Set                    `tfsdk:"grant_types"`
	RedirectURIs             types.Set                    `tfsdk:"redirect_uris"`
	Scopes                   types.Set                    `tfsdk:"scopes"`
	CustomAttributes         []SourceCustomAttributeModel `tfsdk:"custom_attribute"`
	CustomClaims             []SourceCustomClaimModel     `tfsdk:"custom_claim"`
	HybridAndImplicitOptions []SourceHybridOptionsModel   `tfsdk:"hybrid_and_implicit_options"`
	RefreshTokenOptions      []SourceRefreshTokenModel    `tfsdk:"refresh_token_options"`
}

// SourceCustomAttributeModel represents the v4 custom_attribute block structure.
type SourceCustomAttributeModel struct {
	Name         types.String                       `tfsdk:"name"`
	FriendlyName types.String                       `tfsdk:"friendly_name"`
	NameFormat   types.String                       `tfsdk:"name_format"`
	Required     types.Bool                         `tfsdk:"required"`
	Source       []SourceCustomAttributeSourceModel `tfsdk:"source"`
}

// SourceCustomAttributeSourceModel represents the v4 custom_attribute source block.
type SourceCustomAttributeSourceModel struct {
	Name      types.String      `tfsdk:"name"`
	NameByIdP map[string]string `tfsdk:"name_by_idp"`
}

// SourceCustomClaimModel represents the v4 custom_claim block structure.
type SourceCustomClaimModel struct {
	Name     types.String                   `tfsdk:"name"`
	Required types.Bool                     `tfsdk:"required"`
	Scope    types.String                   `tfsdk:"scope"`
	Source   []SourceCustomClaimSourceModel `tfsdk:"source"`
}

// SourceCustomClaimSourceModel represents the v4 custom_claim source block.
type SourceCustomClaimSourceModel struct {
	Name      types.String            `tfsdk:"name"`
	NameByIdP map[string]types.String `tfsdk:"name_by_idp"`
}

// SourceHybridOptionsModel represents the v4 hybrid_and_implicit_options block.
type SourceHybridOptionsModel struct {
	ReturnAccessTokenFromAuthEndpoint types.Bool `tfsdk:"return_access_token_from_authorization_endpoint"`
	ReturnIDTokenFromAuthEndpoint     types.Bool `tfsdk:"return_id_token_from_authorization_endpoint"`
}

// SourceRefreshTokenModel represents the v4 refresh_token_options block.
type SourceRefreshTokenModel struct {
	Lifetime types.String `tfsdk:"lifetime"`
}

// SourceSCIMConfigModel represents the v4 scim_config block structure.
type SourceSCIMConfigModel struct {
	IdPUID             types.String              `tfsdk:"idp_uid"`
	RemoteURI          types.String              `tfsdk:"remote_uri"`
	Enabled            types.Bool                `tfsdk:"enabled"`
	DeactivateOnDelete types.Bool                `tfsdk:"deactivate_on_delete"`
	Authentication     []SourceSCIMAuthModel     `tfsdk:"authentication"`
	Mappings           []SourceSCIMMappingsModel `tfsdk:"mappings"`
}

// SourceSCIMAuthModel represents the v4 scim_config authentication block.
type SourceSCIMAuthModel struct {
	Scheme           types.String `tfsdk:"scheme"`
	User             types.String `tfsdk:"user"`
	Password         types.String `tfsdk:"password"`
	Token            types.String `tfsdk:"token"`
	AuthorizationURL types.String `tfsdk:"authorization_url"`
	ClientID         types.String `tfsdk:"client_id"`
	ClientSecret     types.String `tfsdk:"client_secret"`
	TokenURL         types.String `tfsdk:"token_url"`
	Scopes           types.Set    `tfsdk:"scopes"`
}

// SourceSCIMMappingsModel represents the v4 scim_config mappings block.
type SourceSCIMMappingsModel struct {
	Schema           types.String                `tfsdk:"schema"`
	Enabled          types.Bool                  `tfsdk:"enabled"`
	Filter           types.String                `tfsdk:"filter"`
	TransformJsonata types.String                `tfsdk:"transform_jsonata"`
	Strictness       types.String                `tfsdk:"strictness"`
	Operations       []SourceSCIMOperationsModel `tfsdk:"operations"`
}

// SourceSCIMOperationsModel represents the v4 scim_config mappings operations block.
type SourceSCIMOperationsModel struct {
	Create types.Bool `tfsdk:"create"`
	Update types.Bool `tfsdk:"update"`
	Delete types.Bool `tfsdk:"delete"`
}

// SourceLandingPageDesignModel represents the v4 landing_page_design block structure.
type SourceLandingPageDesignModel struct {
	ButtonColor     types.String `tfsdk:"button_color"`
	ButtonTextColor types.String `tfsdk:"button_text_color"`
	ImageURL        types.String `tfsdk:"image_url"`
	Message         types.String `tfsdk:"message"`
	Title           types.String `tfsdk:"title"`
}

// SourceFooterLinksModel represents the v4 footer_links block structure.
type SourceFooterLinksModel struct {
	Name types.String `tfsdk:"name"`
	URL  types.String `tfsdk:"url"`
}

// SourceDestinationsModel represents the v4 destinations block structure.
type SourceDestinationsModel struct {
	Type       types.String `tfsdk:"type"`
	URI        types.String `tfsdk:"uri"`
	Hostname   types.String `tfsdk:"hostname"`
	CIDR       types.String `tfsdk:"cidr"`
	PortRange  types.String `tfsdk:"port_range"`
	VnetID     types.String `tfsdk:"vnet_id"`
	L4Protocol types.String `tfsdk:"l4_protocol"`
}

// SourceTargetCriteriaModel represents the v4 target_criteria block structure.
type SourceTargetCriteriaModel struct {
	Port             types.Int64                   `tfsdk:"port"`
	Protocol         types.String                  `tfsdk:"protocol"`
	TargetAttributes []SourceTargetAttributesModel `tfsdk:"target_attributes"`
}

// SourceTargetAttributesModel represents the v4 target_attributes block.
type SourceTargetAttributesModel struct {
	Name   types.String `tfsdk:"name"`
	Values types.List   `tfsdk:"values"`
}

// TargetAccessApplicationModel represents the v5 cloudflare_zero_trust_access_application state structure.
// This is a copy of the main model to avoid import cycles.
type TargetAccessApplicationModel struct {
	ID                          types.String                                           `tfsdk:"id"`
	AccountID                   types.String                                           `tfsdk:"account_id"`
	ZoneID                      types.String                                           `tfsdk:"zone_id"`
	AllowAuthenticateViaWARP    types.Bool                                             `tfsdk:"allow_authenticate_via_warp"`
	AllowIframe                 types.Bool                                             `tfsdk:"allow_iframe"`
	AppLauncherLogoURL          types.String                                           `tfsdk:"app_launcher_logo_url"`
	BgColor                     types.String                                           `tfsdk:"bg_color"`
	CustomDenyMessage           types.String                                           `tfsdk:"custom_deny_message"`
	CustomDenyURL               types.String                                           `tfsdk:"custom_deny_url"`
	CustomNonIdentityDenyURL    types.String                                           `tfsdk:"custom_non_identity_deny_url"`
	Domain                      types.String                                           `tfsdk:"domain"`
	HeaderBgColor               types.String                                           `tfsdk:"header_bg_color"`
	LogoURL                     types.String                                           `tfsdk:"logo_url"`
	Name                        types.String                                           `tfsdk:"name"`
	OptionsPreflightBypass      types.Bool                                             `tfsdk:"options_preflight_bypass"`
	ReadServiceTokensFromHeader types.String                                           `tfsdk:"read_service_tokens_from_header"`
	SameSiteCookieAttribute     types.String                                           `tfsdk:"same_site_cookie_attribute"`
	ServiceAuth401Redirect      types.Bool                                             `tfsdk:"service_auth_401_redirect"`
	SkipInterstitial            types.Bool                                             `tfsdk:"skip_interstitial"`
	Type                        types.String                                           `tfsdk:"type"`
	AllowedIdPs                 *[]types.String                                        `tfsdk:"allowed_idps"`
	CustomPages                 *[]types.String                                        `tfsdk:"custom_pages"`
	CORSHeaders                 *TargetCORSHeadersModel                                `tfsdk:"cors_headers"`
	FooterLinks                 *[]*TargetFooterLinksModel                             `tfsdk:"footer_links"`
	OAuthConfiguration          *TargetOAuthConfigurationModel                         `tfsdk:"oauth_configuration"`
	SCIMConfig                  *TargetSCIMConfigModel                                 `tfsdk:"scim_config"`
	TargetCriteria              *[]*TargetTargetCriteriaModel                          `tfsdk:"target_criteria"`
	AppLauncherVisible          types.Bool                                             `tfsdk:"app_launcher_visible"`
	AutoRedirectToIdentity      types.Bool                                             `tfsdk:"auto_redirect_to_identity"`
	EnableBindingCookie         types.Bool                                             `tfsdk:"enable_binding_cookie"`
	HTTPOnlyCookieAttribute     types.Bool                                             `tfsdk:"http_only_cookie_attribute"`
	PathCookieAttribute         types.Bool                                             `tfsdk:"path_cookie_attribute"`
	SessionDuration             types.String                                           `tfsdk:"session_duration"`
	SkipAppLauncherLoginPage    types.Bool                                             `tfsdk:"skip_app_launcher_login_page"`
	SelfHostedDomains           customfield.Set[types.String]                          `tfsdk:"self_hosted_domains"`
	Tags                        customfield.Set[types.String]                          `tfsdk:"tags"`
	Destinations                customfield.NestedObjectList[TargetDestinationsModel]  `tfsdk:"destinations"`
	LandingPageDesign           customfield.NestedObject[TargetLandingPageDesignModel] `tfsdk:"landing_page_design"`
	Policies                    *[]TargetPoliciesModel                                 `tfsdk:"policies"`
	AUD                         types.String                                           `tfsdk:"aud"`
	SaaSApp                     *TargetSaaSAppModel                                    `tfsdk:"saas_app"`
}

type TargetCORSHeadersModel struct {
	AllowAllHeaders  types.Bool      `tfsdk:"allow_all_headers"`
	AllowAllMethods  types.Bool      `tfsdk:"allow_all_methods"`
	AllowAllOrigins  types.Bool      `tfsdk:"allow_all_origins"`
	AllowCredentials types.Bool      `tfsdk:"allow_credentials"`
	AllowedHeaders   *[]types.String `tfsdk:"allowed_headers"`
	AllowedMethods   *[]types.String `tfsdk:"allowed_methods"`
	AllowedOrigins   *[]types.String `tfsdk:"allowed_origins"`
	MaxAge           types.Float64   `tfsdk:"max_age"`
}

type TargetFooterLinksModel struct {
	Name types.String `tfsdk:"name"`
	URL  types.String `tfsdk:"url"`
}

type TargetOAuthConfigurationModel struct {
	Enabled                   types.Bool                                     `tfsdk:"enabled"`
	DynamicClientRegistration *TargetOAuthConfigurationDynamicClientRegModel `tfsdk:"dynamic_client_registration"`
	Grant                     *TargetOAuthConfigurationGrantModel            `tfsdk:"grant"`
}

type TargetOAuthConfigurationDynamicClientRegModel struct {
	Enabled             types.Bool      `tfsdk:"enabled"`
	AllowAnyOnLocalhost types.Bool      `tfsdk:"allow_any_on_localhost"`
	AllowAnyOnLoopback  types.Bool      `tfsdk:"allow_any_on_loopback"`
	AllowedURIs         *[]types.String `tfsdk:"allowed_uris"`
}

type TargetOAuthConfigurationGrantModel struct {
	AccessTokenLifetime types.String `tfsdk:"access_token_lifetime"`
	SessionDuration     types.String `tfsdk:"session_duration"`
}

type TargetSaaSAppModel struct {
	AuthType                      types.String                         `tfsdk:"auth_type"`
	ConsumerServiceURL            types.String                         `tfsdk:"consumer_service_url"`
	CustomAttributes              *[]*TargetCustomAttributesModel      `tfsdk:"custom_attributes"`
	DefaultRelayState             types.String                         `tfsdk:"default_relay_state"`
	IdPEntityID                   types.String                         `tfsdk:"idp_entity_id"`
	NameIDFormat                  types.String                         `tfsdk:"name_id_format"`
	NameIDTransformJsonata        types.String                         `tfsdk:"name_id_transform_jsonata"`
	PublicKey                     types.String                         `tfsdk:"public_key"`
	SAMLAttributeTransformJsonata types.String                         `tfsdk:"saml_attribute_transform_jsonata"`
	SPEntityID                    types.String                         `tfsdk:"sp_entity_id"`
	SSOEndpoint                   types.String                         `tfsdk:"sso_endpoint"`
	AccessTokenLifetime           types.String                         `tfsdk:"access_token_lifetime"`
	AllowPKCEWithoutClientSecret  types.Bool                           `tfsdk:"allow_pkce_without_client_secret"`
	AppLauncherURL                types.String                         `tfsdk:"app_launcher_url"`
	ClientID                      types.String                         `tfsdk:"client_id"`
	ClientSecret                  types.String                         `tfsdk:"client_secret"`
	CustomClaims                  *[]*TargetCustomClaimsModel          `tfsdk:"custom_claims"`
	GrantTypes                    *[]types.String                      `tfsdk:"grant_types"`
	GroupFilterRegex              types.String                         `tfsdk:"group_filter_regex"`
	HybridAndImplicitOptions      *TargetHybridAndImplicitOptionsModel `tfsdk:"hybrid_and_implicit_options"`
	RedirectURIs                  *[]types.String                      `tfsdk:"redirect_uris"`
	RefreshTokenOptions           *TargetRefreshTokenOptionsModel      `tfsdk:"refresh_token_options"`
	Scopes                        *[]types.String                      `tfsdk:"scopes"`
}

type TargetCustomAttributesModel struct {
	FriendlyName types.String                       `tfsdk:"friendly_name"`
	Name         types.String                       `tfsdk:"name"`
	NameFormat   types.String                       `tfsdk:"name_format"`
	Required     types.Bool                         `tfsdk:"required"`
	Source       *TargetCustomAttributesSourceModel `tfsdk:"source"`
}

type TargetCustomAttributesSourceModel struct {
	Name      types.String             `tfsdk:"name"`
	NameByIdP *[]*TargetNameByIdPModel `tfsdk:"name_by_idp"`
}

type TargetNameByIdPModel struct {
	IdPID      types.String `tfsdk:"idp_id"`
	SourceName types.String `tfsdk:"source_name"`
}

type TargetCustomClaimsModel struct {
	Name     types.String                   `tfsdk:"name"`
	Required types.Bool                     `tfsdk:"required"`
	Scope    types.String                   `tfsdk:"scope"`
	Source   *TargetCustomClaimsSourceModel `tfsdk:"source"`
}

type TargetCustomClaimsSourceModel struct {
	Name      types.String             `tfsdk:"name"`
	NameByIdP *map[string]types.String `tfsdk:"name_by_idp"`
}

type TargetHybridAndImplicitOptionsModel struct {
	ReturnAccessTokenFromAuthorizationEndpoint types.Bool `tfsdk:"return_access_token_from_authorization_endpoint"`
	ReturnIDTokenFromAuthorizationEndpoint     types.Bool `tfsdk:"return_id_token_from_authorization_endpoint"`
}

type TargetRefreshTokenOptionsModel struct {
	Lifetime types.String `tfsdk:"lifetime"`
}

type TargetSCIMConfigModel struct {
	IdPUID             types.String                   `tfsdk:"idp_uid"`
	RemoteURI          types.String                   `tfsdk:"remote_uri"`
	Authentication     *TargetSCIMAuthenticationModel `tfsdk:"authentication"`
	DeactivateOnDelete types.Bool                     `tfsdk:"deactivate_on_delete"`
	Enabled            types.Bool                     `tfsdk:"enabled"`
	Mappings           *[]*TargetSCIMMappingsModel    `tfsdk:"mappings"`
}

type TargetSCIMAuthenticationModel struct {
	Password         types.String    `tfsdk:"password"`
	Scheme           types.String    `tfsdk:"scheme"`
	User             types.String    `tfsdk:"user"`
	Token            types.String    `tfsdk:"token"`
	AuthorizationURL types.String    `tfsdk:"authorization_url"`
	ClientID         types.String    `tfsdk:"client_id"`
	ClientSecret     types.String    `tfsdk:"client_secret"`
	TokenURL         types.String    `tfsdk:"token_url"`
	Scopes           *[]types.String `tfsdk:"scopes"`
}

type TargetSCIMMappingsModel struct {
	Schema           types.String               `tfsdk:"schema"`
	Enabled          types.Bool                 `tfsdk:"enabled"`
	Filter           types.String               `tfsdk:"filter"`
	Operations       *TargetSCIMOperationsModel `tfsdk:"operations"`
	Strictness       types.String               `tfsdk:"strictness"`
	TransformJsonata types.String               `tfsdk:"transform_jsonata"`
}

type TargetSCIMOperationsModel struct {
	Create types.Bool `tfsdk:"create"`
	Delete types.Bool `tfsdk:"delete"`
	Update types.Bool `tfsdk:"update"`
}

type TargetTargetCriteriaModel struct {
	Port             types.Int64                 `tfsdk:"port"`
	Protocol         types.String                `tfsdk:"protocol"`
	TargetAttributes *map[string]*[]types.String `tfsdk:"target_attributes"`
}

type TargetDestinationsModel struct {
	Type        types.String `tfsdk:"type"`
	URI         types.String `tfsdk:"uri"`
	CIDR        types.String `tfsdk:"cidr"`
	Hostname    types.String `tfsdk:"hostname"`
	L4Protocol  types.String `tfsdk:"l4_protocol"`
	PortRange   types.String `tfsdk:"port_range"`
	VnetID      types.String `tfsdk:"vnet_id"`
	McpServerID types.String `tfsdk:"mcp_server_id"`
}

type TargetLandingPageDesignModel struct {
	ButtonColor     types.String `tfsdk:"button_color"`
	ButtonTextColor types.String `tfsdk:"button_text_color"`
	ImageURL        types.String `tfsdk:"image_url"`
	Message         types.String `tfsdk:"message"`
	Title           types.String `tfsdk:"title"`
}

// TargetPoliciesModel represents the v5 policies structure.
// This is a copy of ZeroTrustAccessApplicationPoliciesModel to maintain upgrade path stability.
type TargetPoliciesModel struct {
	ID              types.String                                            `tfsdk:"id"`
	Precedence      types.Int64                                             `tfsdk:"precedence"`
	Decision        types.String                                            `tfsdk:"decision"`
	Include         customfield.NestedObjectSet[TargetPoliciesIncludeModel] `tfsdk:"include"`
	Name            types.String                                            `tfsdk:"name"`
	ConnectionRules *TargetPoliciesConnectionRulesModel                     `tfsdk:"connection_rules"`
	Exclude         customfield.NestedObjectSet[TargetPoliciesExcludeModel] `tfsdk:"exclude"`
	Require         customfield.NestedObjectSet[TargetPoliciesRequireModel] `tfsdk:"require"`
}

type TargetPoliciesIncludeModel struct {
	Group                *TargetPoliciesIncludeGroupModel                `tfsdk:"group"`
	AnyValidServiceToken *TargetPoliciesIncludeAnyValidServiceTokenModel `tfsdk:"any_valid_service_token"`
	AuthContext          *TargetPoliciesIncludeAuthContextModel          `tfsdk:"auth_context"`
	AuthMethod           *TargetPoliciesIncludeAuthMethodModel           `tfsdk:"auth_method"`
	AzureAD              *TargetPoliciesIncludeAzureADModel              `tfsdk:"azure_ad"`
	Certificate          *TargetPoliciesIncludeCertificateModel          `tfsdk:"certificate"`
	CommonName           *TargetPoliciesIncludeCommonNameModel           `tfsdk:"common_name"`
	Geo                  *TargetPoliciesIncludeGeoModel                  `tfsdk:"geo"`
	DevicePosture        *TargetPoliciesIncludeDevicePostureModel        `tfsdk:"device_posture"`
	EmailDomain          *TargetPoliciesIncludeEmailDomainModel          `tfsdk:"email_domain"`
	EmailList            *TargetPoliciesIncludeEmailListModel            `tfsdk:"email_list"`
	Email                *TargetPoliciesIncludeEmailModel                `tfsdk:"email"`
	Everyone             *TargetPoliciesIncludeEveryoneModel             `tfsdk:"everyone"`
	ExternalEvaluation   *TargetPoliciesIncludeExternalEvaluationModel   `tfsdk:"external_evaluation"`
	GitHubOrganization   *TargetPoliciesIncludeGitHubOrganizationModel   `tfsdk:"github_organization"`
	GSuite               *TargetPoliciesIncludeGSuiteModel               `tfsdk:"gsuite"`
	LoginMethod          *TargetPoliciesIncludeLoginMethodModel          `tfsdk:"login_method"`
	IPList               *TargetPoliciesIncludeIPListModel               `tfsdk:"ip_list"`
	IP                   *TargetPoliciesIncludeIPModel                   `tfsdk:"ip"`
	Okta                 *TargetPoliciesIncludeOktaModel                 `tfsdk:"okta"`
	SAML                 *TargetPoliciesIncludeSAMLModel                 `tfsdk:"saml"`
	OIDC                 *TargetPoliciesIncludeOIDCModel                 `tfsdk:"oidc"`
	ServiceToken         *TargetPoliciesIncludeServiceTokenModel         `tfsdk:"service_token"`
	LinkedAppToken       *TargetPoliciesIncludeLinkedAppTokenModel       `tfsdk:"linked_app_token"`
}

type TargetPoliciesIncludeGroupModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPoliciesIncludeAnyValidServiceTokenModel struct{}

type TargetPoliciesIncludeAuthContextModel struct {
	ID                 types.String `tfsdk:"id"`
	AcID               types.String `tfsdk:"ac_id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesIncludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method"`
}

type TargetPoliciesIncludeAzureADModel struct {
	ID                 types.String `tfsdk:"id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesIncludeCertificateModel struct{}

type TargetPoliciesIncludeCommonNameModel struct {
	CommonName types.String `tfsdk:"common_name"`
}

type TargetPoliciesIncludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code"`
}

type TargetPoliciesIncludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid"`
}

type TargetPoliciesIncludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain"`
}

type TargetPoliciesIncludeEmailListModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPoliciesIncludeEmailModel struct {
	Email types.String `tfsdk:"email"`
}

type TargetPoliciesIncludeEveryoneModel struct{}

type TargetPoliciesIncludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url"`
}

type TargetPoliciesIncludeGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
	Name               types.String `tfsdk:"name"`
	Team               types.String `tfsdk:"team"`
}

type TargetPoliciesIncludeGSuiteModel struct {
	Email              types.String `tfsdk:"email"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesIncludeLoginMethodModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPoliciesIncludeIPListModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPoliciesIncludeIPModel struct {
	IP types.String `tfsdk:"ip"`
}

type TargetPoliciesIncludeOktaModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
	Name               types.String `tfsdk:"name"`
}

type TargetPoliciesIncludeSAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name"`
	AttributeValue     types.String `tfsdk:"attribute_value"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesIncludeOIDCModel struct {
	ClaimName          types.String `tfsdk:"claim_name"`
	ClaimValue         types.String `tfsdk:"claim_value"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesIncludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id"`
}

type TargetPoliciesIncludeLinkedAppTokenModel struct {
	AppUID types.String `tfsdk:"app_uid"`
}

type TargetPoliciesConnectionRulesModel struct {
	SSH *TargetPoliciesConnectionRulesSSHModel `tfsdk:"ssh"`
	RDP *TargetPoliciesConnectionRulesRDPModel `tfsdk:"rdp"`
}

type TargetPoliciesConnectionRulesSSHModel struct {
	Usernames       *[]types.String `tfsdk:"usernames"`
	AllowEmailAlias types.Bool      `tfsdk:"allow_email_alias"`
}

type TargetPoliciesConnectionRulesRDPModel struct {
	AllowedClipboardLocalToRemoteFormats *[]types.String `tfsdk:"allowed_clipboard_local_to_remote_formats"`
	AllowedClipboardRemoteToLocalFormats *[]types.String `tfsdk:"allowed_clipboard_remote_to_local_formats"`
}

type TargetPoliciesExcludeModel struct {
	Group                *TargetPoliciesExcludeGroupModel                `tfsdk:"group"`
	AnyValidServiceToken *TargetPoliciesExcludeAnyValidServiceTokenModel `tfsdk:"any_valid_service_token"`
	AuthContext          *TargetPoliciesExcludeAuthContextModel          `tfsdk:"auth_context"`
	AuthMethod           *TargetPoliciesExcludeAuthMethodModel           `tfsdk:"auth_method"`
	AzureAD              *TargetPoliciesExcludeAzureADModel              `tfsdk:"azure_ad"`
	Certificate          *TargetPoliciesExcludeCertificateModel          `tfsdk:"certificate"`
	CommonName           *TargetPoliciesExcludeCommonNameModel           `tfsdk:"common_name"`
	Geo                  *TargetPoliciesExcludeGeoModel                  `tfsdk:"geo"`
	DevicePosture        *TargetPoliciesExcludeDevicePostureModel        `tfsdk:"device_posture"`
	EmailDomain          *TargetPoliciesExcludeEmailDomainModel          `tfsdk:"email_domain"`
	EmailList            *TargetPoliciesExcludeEmailListModel            `tfsdk:"email_list"`
	Email                *TargetPoliciesExcludeEmailModel                `tfsdk:"email"`
	Everyone             *TargetPoliciesExcludeEveryoneModel             `tfsdk:"everyone"`
	ExternalEvaluation   *TargetPoliciesExcludeExternalEvaluationModel   `tfsdk:"external_evaluation"`
	GitHubOrganization   *TargetPoliciesExcludeGitHubOrganizationModel   `tfsdk:"github_organization"`
	GSuite               *TargetPoliciesExcludeGSuiteModel               `tfsdk:"gsuite"`
	LoginMethod          *TargetPoliciesExcludeLoginMethodModel          `tfsdk:"login_method"`
	IPList               *TargetPoliciesExcludeIPListModel               `tfsdk:"ip_list"`
	IP                   *TargetPoliciesExcludeIPModel                   `tfsdk:"ip"`
	Okta                 *TargetPoliciesExcludeOktaModel                 `tfsdk:"okta"`
	SAML                 *TargetPoliciesExcludeSAMLModel                 `tfsdk:"saml"`
	OIDC                 *TargetPoliciesExcludeOIDCModel                 `tfsdk:"oidc"`
	ServiceToken         *TargetPoliciesExcludeServiceTokenModel         `tfsdk:"service_token"`
	LinkedAppToken       *TargetPoliciesExcludeLinkedAppTokenModel       `tfsdk:"linked_app_token"`
}

type TargetPoliciesExcludeGroupModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPoliciesExcludeAnyValidServiceTokenModel struct{}

type TargetPoliciesExcludeAuthContextModel struct {
	ID                 types.String `tfsdk:"id"`
	AcID               types.String `tfsdk:"ac_id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesExcludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method"`
}

type TargetPoliciesExcludeAzureADModel struct {
	ID                 types.String `tfsdk:"id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesExcludeCertificateModel struct{}

type TargetPoliciesExcludeCommonNameModel struct {
	CommonName types.String `tfsdk:"common_name"`
}

type TargetPoliciesExcludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code"`
}

type TargetPoliciesExcludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid"`
}

type TargetPoliciesExcludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain"`
}

type TargetPoliciesExcludeEmailListModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPoliciesExcludeEmailModel struct {
	Email types.String `tfsdk:"email"`
}

type TargetPoliciesExcludeEveryoneModel struct{}

type TargetPoliciesExcludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url"`
}

type TargetPoliciesExcludeGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
	Name               types.String `tfsdk:"name"`
	Team               types.String `tfsdk:"team"`
}

type TargetPoliciesExcludeGSuiteModel struct {
	Email              types.String `tfsdk:"email"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesExcludeLoginMethodModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPoliciesExcludeIPListModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPoliciesExcludeIPModel struct {
	IP types.String `tfsdk:"ip"`
}

type TargetPoliciesExcludeOktaModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
	Name               types.String `tfsdk:"name"`
}

type TargetPoliciesExcludeSAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name"`
	AttributeValue     types.String `tfsdk:"attribute_value"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesExcludeOIDCModel struct {
	ClaimName          types.String `tfsdk:"claim_name"`
	ClaimValue         types.String `tfsdk:"claim_value"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesExcludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id"`
}

type TargetPoliciesExcludeLinkedAppTokenModel struct {
	AppUID types.String `tfsdk:"app_uid"`
}

type TargetPoliciesRequireModel struct {
	Group                *TargetPoliciesRequireGroupModel                `tfsdk:"group"`
	AnyValidServiceToken *TargetPoliciesRequireAnyValidServiceTokenModel `tfsdk:"any_valid_service_token"`
	AuthContext          *TargetPoliciesRequireAuthContextModel          `tfsdk:"auth_context"`
	AuthMethod           *TargetPoliciesRequireAuthMethodModel           `tfsdk:"auth_method"`
	AzureAD              *TargetPoliciesRequireAzureADModel              `tfsdk:"azure_ad"`
	Certificate          *TargetPoliciesRequireCertificateModel          `tfsdk:"certificate"`
	CommonName           *TargetPoliciesRequireCommonNameModel           `tfsdk:"common_name"`
	Geo                  *TargetPoliciesRequireGeoModel                  `tfsdk:"geo"`
	DevicePosture        *TargetPoliciesRequireDevicePostureModel        `tfsdk:"device_posture"`
	EmailDomain          *TargetPoliciesRequireEmailDomainModel          `tfsdk:"email_domain"`
	EmailList            *TargetPoliciesRequireEmailListModel            `tfsdk:"email_list"`
	Email                *TargetPoliciesRequireEmailModel                `tfsdk:"email"`
	Everyone             *TargetPoliciesRequireEveryoneModel             `tfsdk:"everyone"`
	ExternalEvaluation   *TargetPoliciesRequireExternalEvaluationModel   `tfsdk:"external_evaluation"`
	GitHubOrganization   *TargetPoliciesRequireGitHubOrganizationModel   `tfsdk:"github_organization"`
	GSuite               *TargetPoliciesRequireGSuiteModel               `tfsdk:"gsuite"`
	LoginMethod          *TargetPoliciesRequireLoginMethodModel          `tfsdk:"login_method"`
	IPList               *TargetPoliciesRequireIPListModel               `tfsdk:"ip_list"`
	IP                   *TargetPoliciesRequireIPModel                   `tfsdk:"ip"`
	Okta                 *TargetPoliciesRequireOktaModel                 `tfsdk:"okta"`
	SAML                 *TargetPoliciesRequireSAMLModel                 `tfsdk:"saml"`
	OIDC                 *TargetPoliciesRequireOIDCModel                 `tfsdk:"oidc"`
	ServiceToken         *TargetPoliciesRequireServiceTokenModel         `tfsdk:"service_token"`
	LinkedAppToken       *TargetPoliciesRequireLinkedAppTokenModel       `tfsdk:"linked_app_token"`
}

type TargetPoliciesRequireGroupModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPoliciesRequireAnyValidServiceTokenModel struct{}

type TargetPoliciesRequireAuthContextModel struct {
	ID                 types.String `tfsdk:"id"`
	AcID               types.String `tfsdk:"ac_id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesRequireAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method"`
}

type TargetPoliciesRequireAzureADModel struct {
	ID                 types.String `tfsdk:"id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesRequireCertificateModel struct{}

type TargetPoliciesRequireCommonNameModel struct {
	CommonName types.String `tfsdk:"common_name"`
}

type TargetPoliciesRequireGeoModel struct {
	CountryCode types.String `tfsdk:"country_code"`
}

type TargetPoliciesRequireDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid"`
}

type TargetPoliciesRequireEmailDomainModel struct {
	Domain types.String `tfsdk:"domain"`
}

type TargetPoliciesRequireEmailListModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPoliciesRequireEmailModel struct {
	Email types.String `tfsdk:"email"`
}

type TargetPoliciesRequireEveryoneModel struct{}

type TargetPoliciesRequireExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url"`
}

type TargetPoliciesRequireGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
	Name               types.String `tfsdk:"name"`
	Team               types.String `tfsdk:"team"`
}

type TargetPoliciesRequireGSuiteModel struct {
	Email              types.String `tfsdk:"email"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesRequireLoginMethodModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPoliciesRequireIPListModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetPoliciesRequireIPModel struct {
	IP types.String `tfsdk:"ip"`
}

type TargetPoliciesRequireOktaModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
	Name               types.String `tfsdk:"name"`
}

type TargetPoliciesRequireSAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name"`
	AttributeValue     types.String `tfsdk:"attribute_value"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesRequireOIDCModel struct {
	ClaimName          types.String `tfsdk:"claim_name"`
	ClaimValue         types.String `tfsdk:"claim_value"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetPoliciesRequireServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id"`
}

type TargetPoliciesRequireLinkedAppTokenModel struct {
	AppUID types.String `tfsdk:"app_uid"`
}
