package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceAccessPolicyModel represents the v4 cloudflare_access_policy state structure.
// This is used by MoveState to parse the source state from v4 provider.
type SourceAccessPolicyModel struct {
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
	Include         []SourceConditionGroupModel  `tfsdk:"include"`
	Exclude         []SourceConditionGroupModel  `tfsdk:"exclude"`
	Require         []SourceConditionGroupModel  `tfsdk:"require"`
	ApprovalGroup   []SourceApprovalGroupModel   `tfsdk:"approval_group"`
	ConnectionRules []SourceConnectionRulesModel `tfsdk:"connection_rules"`
}

// SourceConditionGroupModel represents a condition group in v4 schema.
// In v4, include/exclude/require are lists of condition groups where each group
// can have multiple condition types.
type SourceConditionGroupModel struct {
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
	CommonNames          types.List   `tfsdk:"common_names"`
	AuthMethod           types.String `tfsdk:"auth_method"`
	// v4 uses simple string lists for these
	DevicePosture types.List `tfsdk:"device_posture"`
	EmailList     types.List `tfsdk:"email_list"`
	IPList        types.List `tfsdk:"ip_list"`
	ServiceToken  types.List `tfsdk:"service_token"`
	// v4 nested blocks
	SAML               []SourceSAMLModel               `tfsdk:"saml"`
	OIDC               []SourceOIDCModel               `tfsdk:"oidc"`
	AzureAD            []SourceAzureADModel            `tfsdk:"azure"`
	Okta               []SourceOktaModel               `tfsdk:"okta"`
	GSuite             []SourceGSuiteModel             `tfsdk:"gsuite"`
	GitHub             []SourceGitHubModel             `tfsdk:"github"`
	ExternalEvaluation []SourceExternalEvaluationModel `tfsdk:"external_evaluation"`
	AuthContext        []SourceAuthContextModel        `tfsdk:"auth_context"`
}

// SourceApprovalGroupModel represents the v4 approval_group block.
type SourceApprovalGroupModel struct {
	ApprovalsNeeded types.Int64  `tfsdk:"approvals_needed"`
	EmailAddresses  types.List   `tfsdk:"email_addresses"`
	EmailListUUID   types.String `tfsdk:"email_list_uuid"`
}

// SourceConnectionRulesModel represents the v4 connection_rules block.
type SourceConnectionRulesModel struct {
	SSH []SourceSSHModel `tfsdk:"ssh"`
}

// SourceSSHModel represents the v4 ssh block within connection_rules.
type SourceSSHModel struct {
	Usernames       types.List `tfsdk:"usernames"`
	AllowEmailAlias types.Bool `tfsdk:"allow_email_alias"`
}

// SourceSAMLModel represents the v4 saml block.
type SourceSAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name"`
	AttributeValue     types.String `tfsdk:"attribute_value"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

// SourceOIDCModel represents the v4 oidc block.
type SourceOIDCModel struct {
	ClaimName          types.String `tfsdk:"claim_name"`
	ClaimValue         types.String `tfsdk:"claim_value"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

// SourceAzureADModel represents the v4 azure block.
type SourceAzureADModel struct {
	ID                 types.List   `tfsdk:"id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

// SourceOktaModel represents the v4 okta block.
type SourceOktaModel struct {
	Name               types.List   `tfsdk:"name"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

// SourceGSuiteModel represents the v4 gsuite block.
type SourceGSuiteModel struct {
	Email              types.List   `tfsdk:"email"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

// SourceGitHubModel represents the v4 github block.
type SourceGitHubModel struct {
	Name               types.String `tfsdk:"name"`
	Teams              types.List   `tfsdk:"teams"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

// SourceExternalEvaluationModel represents the v4 external_evaluation block.
type SourceExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url"`
}

// SourceAuthContextModel represents the v4 auth_context block.
type SourceAuthContextModel struct {
	ID                 types.String `tfsdk:"id"`
	AcID               types.String `tfsdk:"ac_id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

// TargetAccessPolicyModel represents the v5 cloudflare_zero_trust_access_policy state structure.
// This is a copy of the main model to avoid import cycles.
type TargetAccessPolicyModel struct {
	ID                           types.String                  `tfsdk:"id"`
	AccountID                    types.String                  `tfsdk:"account_id"`
	Decision                     types.String                  `tfsdk:"decision"`
	Name                         types.String                  `tfsdk:"name"`
	ApprovalRequired             types.Bool                    `tfsdk:"approval_required"`
	IsolationRequired            types.Bool                    `tfsdk:"isolation_required"`
	PurposeJustificationPrompt   types.String                  `tfsdk:"purpose_justification_prompt"`
	PurposeJustificationRequired types.Bool                    `tfsdk:"purpose_justification_required"`
	ApprovalGroups               *[]*TargetApprovalGroupsModel `tfsdk:"approval_groups"`
	ConnectionRules              *TargetConnectionRulesModel   `tfsdk:"connection_rules"`
	MfaConfig                    *TargetMfaConfigModel         `tfsdk:"mfa_config"`
	SessionDuration              types.String                  `tfsdk:"session_duration"`
	Exclude                      []TargetConditionModel        `tfsdk:"exclude"`
	Include                      []TargetConditionModel        `tfsdk:"include"`
	Require                      []TargetConditionModel        `tfsdk:"require"`
}

type TargetApprovalGroupsModel struct {
	ApprovalsNeeded types.Float64   `tfsdk:"approvals_needed"`
	EmailAddresses  *[]types.String `tfsdk:"email_addresses"`
	EmailListUUID   types.String    `tfsdk:"email_list_uuid"`
}

type TargetConnectionRulesModel struct {
	RDP *TargetConnectionRulesRDPModel `tfsdk:"rdp"`
}

type TargetConnectionRulesRDPModel struct {
	AllowedClipboardLocalToRemoteFormats *[]types.String `tfsdk:"allowed_clipboard_local_to_remote_formats"`
	AllowedClipboardRemoteToLocalFormats *[]types.String `tfsdk:"allowed_clipboard_remote_to_local_formats"`
}

type TargetMfaConfigModel struct {
	AllowedAuthenticators *[]types.String `tfsdk:"allowed_authenticators"`
	MfaDisabled           types.Bool      `tfsdk:"mfa_disabled"`
	SessionDuration       types.String    `tfsdk:"session_duration"`
}

// TargetConditionModel represents a single condition in v5 format.
// In v5, each condition is a separate object with only one field populated.
type TargetConditionModel struct {
	Group                *TargetGroupModel                `tfsdk:"group"`
	AnyValidServiceToken *TargetAnyValidServiceTokenModel `tfsdk:"any_valid_service_token"`
	AuthContext          *TargetAuthContextModel          `tfsdk:"auth_context"`
	AuthMethod           *TargetAuthMethodModel           `tfsdk:"auth_method"`
	AzureAD              *TargetAzureADModel              `tfsdk:"azure_ad"`
	Certificate          *TargetCertificateModel          `tfsdk:"certificate"`
	CommonName           *TargetCommonNameModel           `tfsdk:"common_name"`
	Geo                  *TargetGeoModel                  `tfsdk:"geo"`
	DevicePosture        *TargetDevicePostureModel        `tfsdk:"device_posture"`
	EmailDomain          *TargetEmailDomainModel          `tfsdk:"email_domain"`
	EmailList            *TargetEmailListModel            `tfsdk:"email_list"`
	Email                *TargetEmailModel                `tfsdk:"email"`
	Everyone             *TargetEveryoneModel             `tfsdk:"everyone"`
	ExternalEvaluation   *TargetExternalEvaluationModel   `tfsdk:"external_evaluation"`
	GitHubOrganization   *TargetGitHubOrganizationModel   `tfsdk:"github_organization"`
	GSuite               *TargetGSuiteModel               `tfsdk:"gsuite"`
	LoginMethod          *TargetLoginMethodModel          `tfsdk:"login_method"`
	IPList               *TargetIPListModel               `tfsdk:"ip_list"`
	IP                   *TargetIPModel                   `tfsdk:"ip"`
	Okta                 *TargetOktaModel                 `tfsdk:"okta"`
	SAML                 *TargetSAMLModel                 `tfsdk:"saml"`
	OIDC                 *TargetOIDCModel                 `tfsdk:"oidc"`
	ServiceToken         *TargetServiceTokenModel         `tfsdk:"service_token"`
	LinkedAppToken       *TargetLinkedAppTokenModel       `tfsdk:"linked_app_token"`
	UserRiskScore        *TargetUserRiskScoreModel        `tfsdk:"user_risk_score"`
}

type TargetUserRiskScoreModel struct {
	UserRiskScore *[]types.String `tfsdk:"user_risk_score"`
}

type TargetGroupModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetAnyValidServiceTokenModel struct{}

type TargetAuthContextModel struct {
	ID                 types.String `tfsdk:"id"`
	AcID               types.String `tfsdk:"ac_id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method"`
}

type TargetAzureADModel struct {
	ID                 types.String `tfsdk:"id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetCertificateModel struct{}

type TargetCommonNameModel struct {
	CommonName types.String `tfsdk:"common_name"`
}

type TargetGeoModel struct {
	CountryCode types.String `tfsdk:"country_code"`
}

type TargetDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid"`
}

type TargetEmailDomainModel struct {
	Domain types.String `tfsdk:"domain"`
}

type TargetEmailListModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetEmailModel struct {
	Email types.String `tfsdk:"email"`
}

type TargetEveryoneModel struct{}

type TargetExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url"`
}

type TargetGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
	Name               types.String `tfsdk:"name"`
	Team               types.String `tfsdk:"team"`
}

type TargetGSuiteModel struct {
	Email              types.String `tfsdk:"email"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetLoginMethodModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetIPListModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetIPModel struct {
	IP types.String `tfsdk:"ip"`
}

type TargetOktaModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
	Name               types.String `tfsdk:"name"`
}

type TargetSAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name"`
	AttributeValue     types.String `tfsdk:"attribute_value"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetOIDCModel struct {
	ClaimName          types.String `tfsdk:"claim_name"`
	ClaimValue         types.String `tfsdk:"claim_value"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id"`
}

type TargetLinkedAppTokenModel struct {
	AppUID types.String `tfsdk:"app_uid"`
}
