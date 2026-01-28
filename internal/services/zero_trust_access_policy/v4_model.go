package zero_trust_access_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// V4AccessPolicyModel represents the v4 cloudflare_access_policy state structure.
// This is used by MoveState to parse the source state from v4 provider.
type V4AccessPolicyModel struct {
	ID                           types.String `tfsdk:"id"`
	AccountID                    types.String `tfsdk:"account_id"`
	ZoneID                       types.String `tfsdk:"zone_id"`
	ApplicationID                types.String `tfsdk:"application_id"`
	Name                         types.String `tfsdk:"name"`
	Precedence                   types.Int64  `tfsdk:"precedence"`
	Decision                     types.String `tfsdk:"decision"`
	SessionDuration              types.String `tfsdk:"session_duration"`
	IsolationRequired            types.Bool   `tfsdk:"isolation_required"`
	PurposeJustificationRequired types.Bool   `tfsdk:"purpose_justification_required"`
	PurposeJustificationPrompt   types.String `tfsdk:"purpose_justification_prompt"`
	ApprovalRequired             types.Bool   `tfsdk:"approval_required"`

	// v4 stores these as list blocks
	Include         []V4ConditionGroupModel `tfsdk:"include"`
	Exclude         []V4ConditionGroupModel `tfsdk:"exclude"`
	Require         []V4ConditionGroupModel `tfsdk:"require"`
	ApprovalGroup   []V4ApprovalGroupModel  `tfsdk:"approval_group"`
	ConnectionRules []V4ConnectionRulesModel `tfsdk:"connection_rules"`
}

// V4ConditionGroupModel represents a condition group in v4 schema.
// In v4, include/exclude/require are lists of condition groups where each group
// can have multiple condition types.
type V4ConditionGroupModel struct {
	Everyone             types.Bool   `tfsdk:"everyone"`
	Certificate          types.Bool   `tfsdk:"certificate"`
	AnyValidServiceToken types.Bool   `tfsdk:"any_valid_service_token"`
	Email                types.List   `tfsdk:"email"`
	EmailDomain          types.List   `tfsdk:"email_domain"`
	IP                   types.List   `tfsdk:"ip"`
	Group                types.List   `tfsdk:"group"`
	Geo                  types.List   `tfsdk:"geo"`
	LoginMethod          types.List   `tfsdk:"login_method"`
	CommonName           types.String `tfsdk:"common_name"`
	AuthMethod           types.String `tfsdk:"auth_method"`
	// v4 uses simple string lists for these
	DevicePosture types.List `tfsdk:"device_posture"`
	EmailList     types.List `tfsdk:"email_list"`
	IPList        types.List `tfsdk:"ip_list"`
	ServiceToken  types.List `tfsdk:"service_token"`
	// v4 nested blocks
	SAML               []V4SAMLModel               `tfsdk:"saml"`
	OIDC               []V4OIDCModel               `tfsdk:"oidc"`
	AzureAD            []V4AzureADModel            `tfsdk:"azure"`
	Okta               []V4OktaModel               `tfsdk:"okta"`
	GSuite             []V4GSuiteModel             `tfsdk:"gsuite"`
	GitHub             []V4GitHubModel             `tfsdk:"github"`
	ExternalEvaluation []V4ExternalEvaluationModel `tfsdk:"external_evaluation"`
	AuthContext        []V4AuthContextModel        `tfsdk:"auth_context"`
}

// V4ApprovalGroupModel represents the v4 approval_group block.
type V4ApprovalGroupModel struct {
	ApprovalsNeeded types.Int64  `tfsdk:"approvals_needed"`
	EmailAddresses  types.List   `tfsdk:"email_addresses"`
	EmailListUUID   types.String `tfsdk:"email_list_uuid"`
}

// V4ConnectionRulesModel represents the v4 connection_rules block.
type V4ConnectionRulesModel struct {
	SSH []V4SSHModel `tfsdk:"ssh"`
}

// V4SSHModel represents the v4 ssh block within connection_rules.
type V4SSHModel struct {
	Usernames       types.List `tfsdk:"usernames"`
	AllowEmailAlias types.Bool `tfsdk:"allow_email_alias"`
}

// V4SAMLModel represents the v4 saml block.
type V4SAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name"`
	AttributeValue     types.String `tfsdk:"attribute_value"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

// V4OIDCModel represents the v4 oidc block.
type V4OIDCModel struct {
	ClaimName          types.String `tfsdk:"claim_name"`
	ClaimValue         types.String `tfsdk:"claim_value"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

// V4AzureADModel represents the v4 azure block.
type V4AzureADModel struct {
	ID                 types.List   `tfsdk:"id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

// V4OktaModel represents the v4 okta block.
type V4OktaModel struct {
	Name               types.List   `tfsdk:"name"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

// V4GSuiteModel represents the v4 gsuite block.
type V4GSuiteModel struct {
	Email              types.List   `tfsdk:"email"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

// V4GitHubModel represents the v4 github block.
type V4GitHubModel struct {
	Name               types.String `tfsdk:"name"`
	Teams              types.List   `tfsdk:"teams"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

// V4ExternalEvaluationModel represents the v4 external_evaluation block.
type V4ExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url"`
}

// V4AuthContextModel represents the v4 auth_context block.
type V4AuthContextModel struct {
	ID                 types.String `tfsdk:"id"`
	AcID               types.String `tfsdk:"ac_id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

