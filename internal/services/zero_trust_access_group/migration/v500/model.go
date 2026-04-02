package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source V4 Models (Legacy SDKv2 Provider)
// ============================================================================

// SourceV4ZeroTrustAccessGroupModel represents the zero_trust_access_group state from v4.x provider (SDKv2).
// Schema version: 0
//
// In SDKv2, TypeList with MaxItems:1 is stored as an array with 1 element.
// Simple TypeList fields (like email = ["a", "b"]) are stored as arrays of primitives.
type SourceV4ZeroTrustAccessGroupModel struct {
	ID        types.String `tfsdk:"id"`
	AccountID types.String `tfsdk:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id"`
	Name      types.String `tfsdk:"name"`

	Include []SourceV4AccessGroupOptionModel `tfsdk:"include"` // TypeList = array
	Exclude []SourceV4AccessGroupOptionModel `tfsdk:"exclude"` // TypeList = array
	Require []SourceV4AccessGroupOptionModel `tfsdk:"require"` // TypeList = array
}

// SourceV4AccessGroupOptionModel represents a selector block (include/exclude/require) from v4.x
// Each selector can have multiple fields, and most fields are arrays in v4.
//
// IMPORTANT: In v4 SDKv2, simple list fields (email, ip, etc.) are stored as []string directly.
// Complex nested objects (gsuite, github, etc.) are stored as []Object because they're TypeList MaxItems:1.
type SourceV4AccessGroupOptionModel struct {
	// Simple string array selectors (stored as []interface{} in state, parsed as types.List)
	Email         types.List `tfsdk:"email"`          // List[String]
	EmailDomain   types.List `tfsdk:"email_domain"`   // List[String]
	EmailList     types.List `tfsdk:"email_list"`     // List[String]
	IP            types.List `tfsdk:"ip"`             // List[String]
	IPList        types.List `tfsdk:"ip_list"`        // List[String]
	ServiceToken  types.List `tfsdk:"service_token"`  // List[String]
	Group         types.List `tfsdk:"group"`          // List[String]
	DevicePosture types.List `tfsdk:"device_posture"` // List[String]
	LoginMethod   types.List `tfsdk:"login_method"`   // List[String]
	Geo           types.List `tfsdk:"geo"`            // List[String]

	// String scalar selectors (single values, not arrays)
	CommonName types.String `tfsdk:"common_name"` // String
	AuthMethod types.String `tfsdk:"auth_method"` // String

	// Boolean selectors (converted to empty objects in v5)
	Everyone             types.Bool `tfsdk:"everyone"`
	Certificate          types.Bool `tfsdk:"certificate"`
	AnyValidServiceToken types.Bool `tfsdk:"any_valid_service_token"`

	// Special case: common_names overflow array (removed in v5)
	CommonNames types.List `tfsdk:"common_names"` // List[String] - overflow field

	// Complex nested object selectors (TypeList MaxItems:1 = array with 1 element)
	GitHub             []SourceV4GitHubModel             `tfsdk:"github"` // Renamed to github_organization in v5
	GSuite             []SourceV4GSuiteModel             `tfsdk:"gsuite"`
	Azure              []SourceV4AzureModel              `tfsdk:"azure"` // Renamed to azure_ad in v5
	Okta               []SourceV4OktaModel               `tfsdk:"okta"`
	SAML               []SourceV4SAMLModel               `tfsdk:"saml"`
	ExternalEvaluation []SourceV4ExternalEvaluationModel `tfsdk:"external_evaluation"`
	AuthContext        []SourceV4AuthContextModel        `tfsdk:"auth_context"`
}

// SourceV4GitHubModel represents the github selector from v4.x
// Renamed to github_organization in v5, and teams becomes team (singular)
type SourceV4GitHubModel struct {
	Name               types.String `tfsdk:"name"`                 // String
	Teams              types.List   `tfsdk:"teams"`                // List[String] - becomes singular 'team' in v5
	IdentityProviderID types.String `tfsdk:"identity_provider_id"` // String (optional in v4)
}

// SourceV4GSuiteModel represents the gsuite selector from v4.x
type SourceV4GSuiteModel struct {
	Email              types.List   `tfsdk:"email"`                // List[String] - extract first element in v5
	IdentityProviderID types.String `tfsdk:"identity_provider_id"` // String (required)
}

// SourceV4AzureModel represents the azure selector from v4.x
// Renamed to azure_ad in v5
type SourceV4AzureModel struct {
	ID                 types.List   `tfsdk:"id"`                   // List[String] - extract first element in v5
	IdentityProviderID types.String `tfsdk:"identity_provider_id"` // String (optional in v4)
}

// SourceV4OktaModel represents the okta selector from v4.x
type SourceV4OktaModel struct {
	Name               types.List   `tfsdk:"name"`                 // List[String] - extract first element in v5
	IdentityProviderID types.String `tfsdk:"identity_provider_id"` // String (optional in v4)
}

// SourceV4SAMLModel represents the saml selector from v4.x
type SourceV4SAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name"`       // String (optional in v4)
	AttributeValue     types.String `tfsdk:"attribute_value"`      // String (optional in v4)
	IdentityProviderID types.String `tfsdk:"identity_provider_id"` // String (optional in v4)
}

// SourceV4ExternalEvaluationModel represents the external_evaluation selector from v4.x
type SourceV4ExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url"` // String (optional in v4)
	KeysURL     types.String `tfsdk:"keys_url"`     // String (optional in v4)
}

// SourceV4AuthContextModel represents the auth_context selector from v4.x
type SourceV4AuthContextModel struct {
	ID                 types.String `tfsdk:"id"`                   // String (required in v4)
	IdentityProviderID types.String `tfsdk:"identity_provider_id"` // String (required in v4)
	AcID               types.String `tfsdk:"ac_id"`                // String (required in v4)
}

// ============================================================================
// Target V5 Model (Plugin Framework Provider)
// ============================================================================

// TargetV5ZeroTrustAccessGroupModel represents the zero_trust_access_group state for v5 (Plugin Framework).
// This mirrors the structure in the parent package's ZeroTrustAccessGroupModel.
// We duplicate it here to avoid import cycles (migrations.go imports v500, which can't import parent).
type TargetV5ZeroTrustAccessGroupModel struct {
	ID        types.String                                 `tfsdk:"id"`
	AccountID types.String                                 `tfsdk:"account_id"`
	ZoneID    types.String                                 `tfsdk:"zone_id"`
	Name      types.String                                 `tfsdk:"name"`
	Include   *[]*TargetV5ZeroTrustAccessGroupIncludeModel `tfsdk:"include"`
	IsDefault types.Bool                                   `tfsdk:"is_default"`
	Exclude   *[]*TargetV5ZeroTrustAccessGroupIncludeModel `tfsdk:"exclude"` // Same structure as Include
	Require   *[]*TargetV5ZeroTrustAccessGroupIncludeModel `tfsdk:"require"` // Same structure as Include
}

// TargetV5ZeroTrustAccessGroupIncludeModel represents a selector block in v5.
// In v5, each selector is a SingleNestedAttribute, so only one field is set per object.
// Note: Exclude and Require use the same model structure.
type TargetV5ZeroTrustAccessGroupIncludeModel struct {
	Group                *TargetV5GroupModel                `tfsdk:"group"`
	AnyValidServiceToken *TargetV5AnyValidServiceTokenModel `tfsdk:"any_valid_service_token"`
	AuthContext          *TargetV5AuthContextModel          `tfsdk:"auth_context"`
	AuthMethod           *TargetV5AuthMethodModel           `tfsdk:"auth_method"`
	AzureAD              *TargetV5AzureADModel              `tfsdk:"azure_ad"`
	Certificate          *TargetV5CertificateModel          `tfsdk:"certificate"`
	CommonName           *TargetV5CommonNameModel           `tfsdk:"common_name"`
	Geo                  *TargetV5GeoModel                  `tfsdk:"geo"`
	DevicePosture        *TargetV5DevicePostureModel        `tfsdk:"device_posture"`
	EmailDomain          *TargetV5EmailDomainModel          `tfsdk:"email_domain"`
	EmailList            *TargetV5EmailListModel            `tfsdk:"email_list"`
	Email                *TargetV5EmailModel                `tfsdk:"email"`
	Everyone             *TargetV5EveryoneModel             `tfsdk:"everyone"`
	ExternalEvaluation   *TargetV5ExternalEvaluationModel   `tfsdk:"external_evaluation"`
	GitHubOrganization   *TargetV5GitHubOrganizationModel   `tfsdk:"github_organization"`
	GSuite               *TargetV5GSuiteModel               `tfsdk:"gsuite"`
	LoginMethod          *TargetV5LoginMethodModel          `tfsdk:"login_method"`
	IPList               *TargetV5IPListModel               `tfsdk:"ip_list"`
	IP                   *TargetV5IPModel                   `tfsdk:"ip"`
	Okta                 *TargetV5OktaModel                 `tfsdk:"okta"`
	SAML                 *TargetV5SAMLModel                 `tfsdk:"saml"`
	OIDC                 *TargetV5OIDCModel                 `tfsdk:"oidc"`
	ServiceToken         *TargetV5ServiceTokenModel         `tfsdk:"service_token"`
	LinkedAppToken       *TargetV5LinkedAppTokenModel       `tfsdk:"linked_app_token"`
	UserRiskScore        *TargetV5UserRiskScoreModel        `tfsdk:"user_risk_score"`
}

type TargetV5UserRiskScoreModel struct {
	UserRiskScore *[]types.String `tfsdk:"user_risk_score"`
}

// Target V5 nested selector models
type TargetV5GroupModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetV5AnyValidServiceTokenModel struct{}

type TargetV5AuthContextModel struct {
	ID                 types.String `tfsdk:"id"`
	AcID               types.String `tfsdk:"ac_id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetV5AuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method"`
}

type TargetV5AzureADModel struct {
	ID                 types.String `tfsdk:"id"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetV5CertificateModel struct{}

type TargetV5CommonNameModel struct {
	CommonName types.String `tfsdk:"common_name"`
}

type TargetV5GeoModel struct {
	CountryCode types.String `tfsdk:"country_code"`
}

type TargetV5DevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid"`
}

type TargetV5EmailDomainModel struct {
	Domain types.String `tfsdk:"domain"`
}

type TargetV5EmailListModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetV5EmailModel struct {
	Email types.String `tfsdk:"email"`
}

type TargetV5EveryoneModel struct{}

type TargetV5ExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url"`
}

type TargetV5GitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
	Name               types.String `tfsdk:"name"`
	Team               types.String `tfsdk:"team"`
}

type TargetV5GSuiteModel struct {
	Email              types.String `tfsdk:"email"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetV5LoginMethodModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetV5IPListModel struct {
	ID types.String `tfsdk:"id"`
}

type TargetV5IPModel struct {
	IP types.String `tfsdk:"ip"`
}

type TargetV5OktaModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
	Name               types.String `tfsdk:"name"`
}

type TargetV5SAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name"`
	AttributeValue     types.String `tfsdk:"attribute_value"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetV5OIDCModel struct {
	ClaimName          types.String `tfsdk:"claim_name"`
	ClaimValue         types.String `tfsdk:"claim_value"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id"`
}

type TargetV5ServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id"`
}

type TargetV5LinkedAppTokenModel struct {
	AppUID types.String `tfsdk:"app_uid"`
}
