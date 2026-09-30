package zero_trust_casb_integration

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ZeroTrustCasbIntegrationModel is manually maintained.
//
// Vendors follow a `vendor_name = { auth_method = { ... } }` shape: each vendor
// block nests one attribute per supported authentication method, named after
// the auth method slug the API expects. Exactly one vendor and exactly one auth
// method must be configured; this is enforced by ConfigValidators in schema.go.
//
// Only vendor/auth method combinations that are live in production, are not
// alpha-gated, and are createable through the public API are exposed. OAuth
// based auth methods are excluded because they require an interactive
// browser redirect that Terraform cannot perform.
//
// Within each auth method, confidential arguments are write-only and never
// reach state; only SecretsDigest is persisted for them. Non-confidential
// arguments are stored in plain text so ordinary Terraform diffing applies.
type ZeroTrustCasbIntegrationModel struct {
	ID        types.String `tfsdk:"id"`
	AccountID types.String `tfsdk:"account_id"`
	Name      types.String `tfsdk:"name"`
	Paused    types.Bool   `tfsdk:"paused"`

	DLPProfiles customfield.Set[types.String] `tfsdk:"dlp_profiles"`
	Permissions customfield.Set[types.String] `tfsdk:"permissions"`
	UseCases    customfield.Set[types.String] `tfsdk:"use_cases"`

	// SecretsDigest is a randomly salted SHA3-256 digest of the confidential arguments only.
	// It is how a secret rotation becomes visible to Terraform, since the
	// secrets themselves are write-only.
	SecretsDigest types.String `tfsdk:"secrets_digest"`

	Anthropic           *ZeroTrustCasbIntegrationAnthropicModel           `tfsdk:"anthropic"`
	AWS                 *ZeroTrustCasbIntegrationAWSModel                 `tfsdk:"aws"`
	Box                 *ZeroTrustCasbIntegrationBoxModel                 `tfsdk:"box"`
	GoogleCloudPlatform *ZeroTrustCasbIntegrationGoogleCloudPlatformModel `tfsdk:"google_cloud_platform"`
	GoogleWorkspace     *ZeroTrustCasbIntegrationGoogleWorkspaceModel     `tfsdk:"google_workspace"`
	OpenAI              *ZeroTrustCasbIntegrationOpenAIModel              `tfsdk:"openai"`
}

// ── Anthropic ───────────────────────────────────────────────────────────────

type ZeroTrustCasbIntegrationAnthropicModel struct {
	AdminAPIKey      *ZeroTrustCasbIntegrationAnthropicAdminAPIKeyModel      `tfsdk:"anthropic_admin_api_key"`
	WorkspaceAPIKey  *ZeroTrustCasbIntegrationAnthropicWorkspaceAPIKeyModel  `tfsdk:"anthropic_workspace_api_key"`
	ComplianceAPIKey *ZeroTrustCasbIntegrationAnthropicComplianceAPIKeyModel `tfsdk:"anthropic_compliance_api_key"`
}

type ZeroTrustCasbIntegrationAnthropicAdminAPIKeyModel struct {
	APIKey   types.String `tfsdk:"api_key"`
	TenantID types.String `tfsdk:"tenant_id"`
}

type ZeroTrustCasbIntegrationAnthropicWorkspaceAPIKeyModel struct {
	APIKey   types.String `tfsdk:"api_key"`
	TenantID types.String `tfsdk:"tenant_id"`
}

type ZeroTrustCasbIntegrationAnthropicComplianceAPIKeyModel struct {
	ComplianceAPIKey types.String `tfsdk:"compliance_api_key"`
	TenantID         types.String `tfsdk:"tenant_id"`
}

// ── AWS ─────────────────────────────────────────────────────────────────────

type ZeroTrustCasbIntegrationAWSModel struct {
	IAMRole *ZeroTrustCasbIntegrationAWSIAMRoleModel `tfsdk:"aws_iam_role"`
}

// AWS IAM role delegation has no confidential arguments: the external ID is a
// customer-chosen value that is useless without Cloudflare's AWS principal.
type ZeroTrustCasbIntegrationAWSIAMRoleModel struct {
	RoleARN    types.String `tfsdk:"role_arn"`
	ExternalID types.String `tfsdk:"external_id"`
}

// ── Box ─────────────────────────────────────────────────────────────────────

type ZeroTrustCasbIntegrationBoxModel struct {
	ServerAuthentication *ZeroTrustCasbIntegrationBoxServerAuthenticationModel `tfsdk:"box_server_authentication"`
}

// Box server authentication uses a Cloudflare-side application, so the caller
// supplies only the enterprise ID.
type ZeroTrustCasbIntegrationBoxServerAuthenticationModel struct {
	EnterpriseID types.String `tfsdk:"enterprise_id"`
}

// ── Google Cloud Platform ───────────────────────────────────────────────────

type ZeroTrustCasbIntegrationGoogleCloudPlatformModel struct {
	ServiceAccount *ZeroTrustCasbIntegrationGoogleCloudPlatformServiceAccountModel `tfsdk:"google_cloud_platform_service_account"`
}

type ZeroTrustCasbIntegrationGoogleCloudPlatformServiceAccountModel struct {
	ServiceAccountKeyJSON jsontypes.Normalized `tfsdk:"service_account_key_json"`
}

// ── Google Workspace ────────────────────────────────────────────────────────

// Workload Identity Federation is intentionally absent: it is flagged
// is_in_alpha and is therefore not offered for new integrations in production.
type ZeroTrustCasbIntegrationGoogleWorkspaceModel struct {
	DomainWideDelegationServiceAccount *ZeroTrustCasbIntegrationGoogleWorkspaceDomainWideDelegationServiceAccountModel `tfsdk:"google_domain_wide_delegation_service_account"`
}

type ZeroTrustCasbIntegrationGoogleWorkspaceDomainWideDelegationServiceAccountModel struct {
	ServiceAccountKeyJSON jsontypes.Normalized `tfsdk:"service_account_key_json"`
	AdministratorEmail    types.String         `tfsdk:"administrator_email"`
}

// ── OpenAI ──────────────────────────────────────────────────────────────────

type ZeroTrustCasbIntegrationOpenAIModel struct {
	StandardAPIKey   *ZeroTrustCasbIntegrationOpenAIStandardAPIKeyModel   `tfsdk:"chatgpt_standard_api_key"`
	ComplianceAPIKey *ZeroTrustCasbIntegrationOpenAIComplianceAPIKeyModel `tfsdk:"chatgpt_compliance_api_key"`
}

type ZeroTrustCasbIntegrationOpenAIStandardAPIKeyModel struct {
	AdminAPIKey    types.String `tfsdk:"admin_api_key"`
	OrganizationID types.String `tfsdk:"organization_id"`
	ProjectAPIKey  types.String `tfsdk:"project_api_key"`
	ProjectID      types.String `tfsdk:"project_id"`
}

type ZeroTrustCasbIntegrationOpenAIComplianceAPIKeyModel struct {
	AdminAPIKey      types.String `tfsdk:"admin_api_key"`
	OrganizationID   types.String `tfsdk:"organization_id"`
	ComplianceAPIKey types.String `tfsdk:"compliance_api_key"`
	WorkspaceID      types.String `tfsdk:"workspace_id"`
	ProjectAPIKey    types.String `tfsdk:"project_api_key"`
	ProjectID        types.String `tfsdk:"project_id"`
}
