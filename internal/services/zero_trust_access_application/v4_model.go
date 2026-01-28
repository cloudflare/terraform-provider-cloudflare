package zero_trust_access_application

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// V4AccessApplicationModel represents the v4 cloudflare_access_application state structure.
// This is used by MoveState to parse the source state from v4 provider.
type V4AccessApplicationModel struct {
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
	CORSHeaders       []V4CORSHeadersModel       `tfsdk:"cors_headers"`
	SaaSApp           []V4SaaSAppModel           `tfsdk:"saas_app"`
	SCIMConfig        []V4SCIMConfigModel        `tfsdk:"scim_config"`
	LandingPageDesign []V4LandingPageDesignModel `tfsdk:"landing_page_design"`
	FooterLinks       []V4FooterLinksModel       `tfsdk:"footer_links"`
	Destinations      []V4DestinationsModel      `tfsdk:"destinations"`
	TargetCriteria    []V4TargetCriteriaModel    `tfsdk:"target_criteria"`

	// Deprecated/removed fields
	DomainType types.String `tfsdk:"domain_type"`
}

// V4CORSHeadersModel represents the v4 cors_headers block structure.
type V4CORSHeadersModel struct {
	AllowAllHeaders  types.Bool  `tfsdk:"allow_all_headers"`
	AllowAllMethods  types.Bool  `tfsdk:"allow_all_methods"`
	AllowAllOrigins  types.Bool  `tfsdk:"allow_all_origins"`
	AllowCredentials types.Bool  `tfsdk:"allow_credentials"`
	AllowedHeaders   types.Set   `tfsdk:"allowed_headers"`
	AllowedMethods   types.Set   `tfsdk:"allowed_methods"`
	AllowedOrigins   types.Set   `tfsdk:"allowed_origins"`
	MaxAge           types.Int64 `tfsdk:"max_age"`
}

// V4SaaSAppModel represents the v4 saas_app block structure.
type V4SaaSAppModel struct {
	AuthType                 types.String             `tfsdk:"auth_type"`
	ConsumerServiceURL       types.String             `tfsdk:"consumer_service_url"`
	SPEntityID               types.String             `tfsdk:"sp_entity_id"`
	IdPEntityID              types.String             `tfsdk:"idp_entity_id"`
	PublicKey                types.String             `tfsdk:"public_key"`
	NameIDFormat             types.String             `tfsdk:"name_id_format"`
	NameIDTransformJsonata   types.String             `tfsdk:"name_id_transform_jsonata"`
	SAMLAttrTransformJsonata types.String             `tfsdk:"saml_attribute_transform_jsonata"`
	DefaultRelayState        types.String             `tfsdk:"default_relay_state"`
	SSOEndpoint              types.String             `tfsdk:"sso_endpoint"`
	AppLauncherURL           types.String             `tfsdk:"app_launcher_url"`
	ClientID                 types.String             `tfsdk:"client_id"`
	ClientSecret             types.String             `tfsdk:"client_secret"`
	AccessTokenLifetime      types.String             `tfsdk:"access_token_lifetime"`
	AllowPKCEWithoutSecret   types.Bool               `tfsdk:"allow_pkce_without_client_secret"`
	GroupFilterRegex         types.String             `tfsdk:"group_filter_regex"`
	GrantTypes               types.Set                `tfsdk:"grant_types"`
	RedirectURIs             types.Set                `tfsdk:"redirect_uris"`
	Scopes                   types.Set                `tfsdk:"scopes"`
	CustomAttributes         []V4CustomAttributeModel `tfsdk:"custom_attribute"`
	CustomClaims             []V4CustomClaimModel     `tfsdk:"custom_claim"`
	HybridAndImplicitOptions []V4HybridOptionsModel   `tfsdk:"hybrid_and_implicit_options"`
	RefreshTokenOptions      []V4RefreshTokenModel    `tfsdk:"refresh_token_options"`
}

// V4CustomAttributeModel represents the v4 custom_attribute block structure.
type V4CustomAttributeModel struct {
	Name         types.String                    `tfsdk:"name"`
	FriendlyName types.String                    `tfsdk:"friendly_name"`
	NameFormat   types.String                    `tfsdk:"name_format"`
	Required     types.Bool                      `tfsdk:"required"`
	Source       []V4CustomAttributeSourceModel  `tfsdk:"source"`
}

// V4CustomAttributeSourceModel represents the v4 custom_attribute source block.
type V4CustomAttributeSourceModel struct {
	Name      types.String       `tfsdk:"name"`
	NameByIdP map[string]string  `tfsdk:"name_by_idp"`
}

// V4CustomClaimModel represents the v4 custom_claim block structure.
type V4CustomClaimModel struct {
	Name     types.String              `tfsdk:"name"`
	Required types.Bool                `tfsdk:"required"`
	Scope    types.String              `tfsdk:"scope"`
	Source   []V4CustomClaimSourceModel `tfsdk:"source"`
}

// V4CustomClaimSourceModel represents the v4 custom_claim source block.
type V4CustomClaimSourceModel struct {
	Name      types.String           `tfsdk:"name"`
	NameByIdP map[string]types.String `tfsdk:"name_by_idp"`
}

// V4HybridOptionsModel represents the v4 hybrid_and_implicit_options block.
type V4HybridOptionsModel struct {
	ReturnAccessTokenFromAuthEndpoint types.Bool `tfsdk:"return_access_token_from_authorization_endpoint"`
	ReturnIDTokenFromAuthEndpoint     types.Bool `tfsdk:"return_id_token_from_authorization_endpoint"`
}

// V4RefreshTokenModel represents the v4 refresh_token_options block.
type V4RefreshTokenModel struct {
	Lifetime types.String `tfsdk:"lifetime"`
}

// V4SCIMConfigModel represents the v4 scim_config block structure.
type V4SCIMConfigModel struct {
	IdPUID             types.String             `tfsdk:"idp_uid"`
	RemoteURI          types.String             `tfsdk:"remote_uri"`
	Enabled            types.Bool               `tfsdk:"enabled"`
	DeactivateOnDelete types.Bool               `tfsdk:"deactivate_on_delete"`
	Authentication     []V4SCIMAuthModel        `tfsdk:"authentication"`
	Mappings           []V4SCIMMappingsModel    `tfsdk:"mappings"`
}

// V4SCIMAuthModel represents the v4 scim_config authentication block.
type V4SCIMAuthModel struct {
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

// V4SCIMMappingsModel represents the v4 scim_config mappings block.
type V4SCIMMappingsModel struct {
	Schema           types.String               `tfsdk:"schema"`
	Enabled          types.Bool                 `tfsdk:"enabled"`
	Filter           types.String               `tfsdk:"filter"`
	TransformJsonata types.String               `tfsdk:"transform_jsonata"`
	Strictness       types.String               `tfsdk:"strictness"`
	Operations       []V4SCIMOperationsModel    `tfsdk:"operations"`
}

// V4SCIMOperationsModel represents the v4 scim_config mappings operations block.
type V4SCIMOperationsModel struct {
	Create types.Bool `tfsdk:"create"`
	Update types.Bool `tfsdk:"update"`
	Delete types.Bool `tfsdk:"delete"`
}

// V4LandingPageDesignModel represents the v4 landing_page_design block structure.
type V4LandingPageDesignModel struct {
	ButtonColor     types.String `tfsdk:"button_color"`
	ButtonTextColor types.String `tfsdk:"button_text_color"`
	ImageURL        types.String `tfsdk:"image_url"`
	Message         types.String `tfsdk:"message"`
	Title           types.String `tfsdk:"title"`
}

// V4FooterLinksModel represents the v4 footer_links block structure.
type V4FooterLinksModel struct {
	Name types.String `tfsdk:"name"`
	URL  types.String `tfsdk:"url"`
}

// V4DestinationsModel represents the v4 destinations block structure.
type V4DestinationsModel struct {
	Type       types.String `tfsdk:"type"`
	URI        types.String `tfsdk:"uri"`
	Hostname   types.String `tfsdk:"hostname"`
	CIDR       types.String `tfsdk:"cidr"`
	PortRange  types.String `tfsdk:"port_range"`
	VnetID     types.String `tfsdk:"vnet_id"`
	L4Protocol types.String `tfsdk:"l4_protocol"`
}

// V4TargetCriteriaModel represents the v4 target_criteria block structure.
type V4TargetCriteriaModel struct {
	Port             types.Int64                  `tfsdk:"port"`
	Protocol         types.String                 `tfsdk:"protocol"`
	TargetAttributes []V4TargetAttributesModel    `tfsdk:"target_attributes"`
}

// V4TargetAttributesModel represents the v4 target_attributes block.
type V4TargetAttributesModel struct {
	Name   types.String `tfsdk:"name"`
	Values types.List   `tfsdk:"values"`
}
