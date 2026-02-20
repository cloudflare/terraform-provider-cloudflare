package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Source models (v4 cloudflare_access_identity_provider)

type SourceAccessIdentityProviderModel struct {
	ID         types.String            `tfsdk:"id"`
	AccountID  types.String            `tfsdk:"account_id"`
	ZoneID     types.String            `tfsdk:"zone_id"`
	Name       types.String            `tfsdk:"name"`
	Type       types.String            `tfsdk:"type"`
	Config     []SourceConfigModel     `tfsdk:"config"`
	ScimConfig []SourceScimConfigModel `tfsdk:"scim_config"`
}

type SourceConfigModel struct {
	Claims                   types.List   `tfsdk:"claims"`
	ClientID                 types.String `tfsdk:"client_id"`
	ClientSecret             types.String `tfsdk:"client_secret"`
	ConditionalAccessEnabled types.Bool   `tfsdk:"conditional_access_enabled"`
	DirectoryID              types.String `tfsdk:"directory_id"`
	EmailClaimName           types.String `tfsdk:"email_claim_name"`
	Prompt                   types.String `tfsdk:"prompt"`
	SupportGroups            types.Bool   `tfsdk:"support_groups"`
	CentrifyAccount          types.String `tfsdk:"centrify_account"`
	CentrifyAppID            types.String `tfsdk:"centrify_app_id"`
	AppsDomain               types.String `tfsdk:"apps_domain"`
	AuthURL                  types.String `tfsdk:"auth_url"`
	CERTsURL                 types.String `tfsdk:"certs_url"`
	PKCEEnabled              types.Bool   `tfsdk:"pkce_enabled"`
	Scopes                   types.List   `tfsdk:"scopes"`
	TokenURL                 types.String `tfsdk:"token_url"`
	AuthorizationServerID    types.String `tfsdk:"authorization_server_id"`
	OktaAccount              types.String `tfsdk:"okta_account"`
	OneloginAccount          types.String `tfsdk:"onelogin_account"`
	PingEnvID                types.String `tfsdk:"ping_env_id"`
	Attributes               types.List   `tfsdk:"attributes"`
	EmailAttributeName       types.String `tfsdk:"email_attribute_name"`
	HeaderAttributes         types.List   `tfsdk:"header_attributes"`
	IssuerURL                types.String `tfsdk:"issuer_url"`
	SignRequest              types.Bool   `tfsdk:"sign_request"`
	SSOTargetURL             types.String `tfsdk:"sso_target_url"`
	RedirectURL              types.String `tfsdk:"redirect_url"`
	ApiToken                 types.String `tfsdk:"api_token"`       // deprecated
	IdpPublicCert            types.String `tfsdk:"idp_public_cert"` // renamed to idp_public_certs in v5
}

type SourceScimConfigModel struct {
	Enabled                types.Bool   `tfsdk:"enabled"`
	IdentityUpdateBehavior types.String `tfsdk:"identity_update_behavior"`
	SCIMBaseURL            types.String `tfsdk:"scim_base_url"`
	SeatDeprovision        types.Bool   `tfsdk:"seat_deprovision"`
	Secret                 types.String `tfsdk:"secret"`
	UserDeprovision        types.Bool   `tfsdk:"user_deprovision"`
	GroupMemberDeprovision types.Bool   `tfsdk:"group_member_deprovision"` // deprecated
}

// Target models (v5 cloudflare_zero_trust_access_identity_provider)

type TargetAccessIdentityProviderModel struct {
	ID         types.String                                      `tfsdk:"id"`
	AccountID  types.String                                      `tfsdk:"account_id"`
	ZoneID     types.String                                      `tfsdk:"zone_id"`
	Name       types.String                                      `tfsdk:"name"`
	Type       types.String                                      `tfsdk:"type"`
	Config     *TargetConfigModel                                `tfsdk:"config"`
	SCIMConfig customfield.NestedObject[TargetScimConfigModel]   `tfsdk:"scim_config"`
}

type TargetConfigModel struct {
	Claims                   *[]types.String                 `tfsdk:"claims"`
	ClientID                 types.String                    `tfsdk:"client_id"`
	ClientSecret             types.String                    `tfsdk:"client_secret"`
	ConditionalAccessEnabled types.Bool                      `tfsdk:"conditional_access_enabled"`
	DirectoryID              types.String                    `tfsdk:"directory_id"`
	EmailClaimName           types.String                    `tfsdk:"email_claim_name"`
	Prompt                   types.String                    `tfsdk:"prompt"`
	SupportGroups            types.Bool                      `tfsdk:"support_groups"`
	CentrifyAccount          types.String                    `tfsdk:"centrify_account"`
	CentrifyAppID            types.String                    `tfsdk:"centrify_app_id"`
	AppsDomain               types.String                    `tfsdk:"apps_domain"`
	AuthURL                  types.String                    `tfsdk:"auth_url"`
	CERTsURL                 types.String                    `tfsdk:"certs_url"`
	PKCEEnabled              types.Bool                      `tfsdk:"pkce_enabled"`
	Scopes                   *[]types.String                 `tfsdk:"scopes"`
	TokenURL                 types.String                    `tfsdk:"token_url"`
	AuthorizationServerID    types.String                    `tfsdk:"authorization_server_id"`
	OktaAccount              types.String                    `tfsdk:"okta_account"`
	OneloginAccount          types.String                    `tfsdk:"onelogin_account"`
	PingEnvID                types.String                    `tfsdk:"ping_env_id"`
	Attributes               *[]types.String                 `tfsdk:"attributes"`
	EmailAttributeName       types.String                    `tfsdk:"email_attribute_name"`
	HeaderAttributes         *[]*TargetHeaderAttributesModel `tfsdk:"header_attributes"`
	IdPPublicCERTs           *[]types.String                 `tfsdk:"idp_public_certs"`
	IssuerURL                types.String                    `tfsdk:"issuer_url"`
	SignRequest              types.Bool                      `tfsdk:"sign_request"`
	SSOTargetURL             types.String                    `tfsdk:"sso_target_url"`
	RedirectURL              types.String                    `tfsdk:"redirect_url"`
}

type TargetHeaderAttributesModel struct {
	AttributeName types.String `tfsdk:"attribute_name"`
	HeaderName    types.String `tfsdk:"header_name"`
}

type TargetScimConfigModel struct {
	Enabled                types.Bool   `tfsdk:"enabled"`
	IdentityUpdateBehavior types.String `tfsdk:"identity_update_behavior"`
	SCIMBaseURL            types.String `tfsdk:"scim_base_url"`
	SeatDeprovision        types.Bool   `tfsdk:"seat_deprovision"`
	Secret                 types.String `tfsdk:"secret"`
	UserDeprovision        types.Bool   `tfsdk:"user_deprovision"`
}
