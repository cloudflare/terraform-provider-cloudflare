// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_ai_controls_mcp_server

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessAIControlsMcpServerResultEnvelope struct {
	Result ZeroTrustAccessAIControlsMcpServerModel `json:"result"`
}

type ZeroTrustAccessAIControlsMcpServerModel struct {
	ID                           types.String                                                                       `tfsdk:"id" json:"id,required"`
	AccountID                    types.String                                                                       `tfsdk:"account_id" path:"account_id,required"`
	AuthType                     types.String                                                                       `tfsdk:"auth_type" json:"auth_type,required"`
	Hostname                     types.String                                                                       `tfsdk:"hostname" json:"hostname,required"`
	Name                         types.String                                                                       `tfsdk:"name" json:"name,required"`
	AuthCredentials              types.String                                                                       `tfsdk:"auth_credentials" json:"auth_credentials,optional,no_refresh"`
	ClientSecret                 types.String                                                                       `tfsdk:"client_secret" json:"client_secret,optional,no_refresh"`
	Description                  types.String                                                                       `tfsdk:"description" json:"description,optional"`
	UpdatedPrompts               *[]*ZeroTrustAccessAIControlsMcpServerUpdatedPromptsModel                          `tfsdk:"updated_prompts" json:"updated_prompts,optional"`
	UpdatedTools                 *[]*ZeroTrustAccessAIControlsMcpServerUpdatedToolsModel                            `tfsdk:"updated_tools" json:"updated_tools,optional"`
	IsSharedOAuthCallbackEnabled types.Bool                                                                         `tfsdk:"is_shared_oauth_callback_enabled" json:"is_shared_oauth_callback_enabled,computed_optional"`
	SecureWebGateway             types.Bool                                                                         `tfsdk:"secure_web_gateway" json:"secure_web_gateway,computed_optional"`
	AuthenticationStatus         types.String                                                                       `tfsdk:"authentication_status" json:"authentication_status,computed"`
	CreatedAt                    timetypes.RFC3339                                                                  `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy                    types.String                                                                       `tfsdk:"created_by" json:"created_by,computed"`
	Error                        types.String                                                                       `tfsdk:"error" json:"error,computed"`
	LastSuccessfulSync           timetypes.RFC3339                                                                  `tfsdk:"last_successful_sync" json:"last_successful_sync,computed" format:"date-time"`
	LastSynced                   timetypes.RFC3339                                                                  `tfsdk:"last_synced" json:"last_synced,computed" format:"date-time"`
	ModifiedAt                   timetypes.RFC3339                                                                  `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	ModifiedBy                   types.String                                                                       `tfsdk:"modified_by" json:"modified_by,computed"`
	Status                       types.String                                                                       `tfsdk:"status" json:"status,computed"`
	Prompts                      customfield.List[customfield.Map[jsontypes.Normalized]]                            `tfsdk:"prompts" json:"prompts,computed"`
	Tools                        customfield.List[customfield.Map[jsontypes.Normalized]]                            `tfsdk:"tools" json:"tools,computed"`
	AuthConfigSummary            customfield.NestedObject[ZeroTrustAccessAIControlsMcpServerAuthConfigSummaryModel] `tfsdk:"auth_config_summary" json:"auth_config_summary,computed"`
	ErrorDetails                 customfield.NestedObject[ZeroTrustAccessAIControlsMcpServerErrorDetailsModel]      `tfsdk:"error_details" json:"error_details,computed"`
}

func (m ZeroTrustAccessAIControlsMcpServerModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessAIControlsMcpServerModel) MarshalJSONForUpdate(state ZeroTrustAccessAIControlsMcpServerModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustAccessAIControlsMcpServerUpdatedPromptsModel struct {
	Name        types.String `tfsdk:"name" json:"name,required"`
	Alias       types.String `tfsdk:"alias" json:"alias,optional"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
}

type ZeroTrustAccessAIControlsMcpServerUpdatedToolsModel struct {
	Name        types.String `tfsdk:"name" json:"name,required"`
	Alias       types.String `tfsdk:"alias" json:"alias,optional"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
}

type ZeroTrustAccessAIControlsMcpServerAuthConfigSummaryModel struct {
	AuthMode            types.String                                                                                       `tfsdk:"auth_mode" json:"auth_mode,computed"`
	ClientSecretVersion types.Float64                                                                                      `tfsdk:"client_secret_version" json:"client_secret_version,computed"`
	Config              customfield.NestedObject[ZeroTrustAccessAIControlsMcpServerAuthConfigSummaryConfigModel]           `tfsdk:"config" json:"config,computed"`
	HasClientSecret     types.Bool                                                                                         `tfsdk:"has_client_secret" json:"has_client_secret,computed"`
	RegistrationInfo    customfield.NestedObject[ZeroTrustAccessAIControlsMcpServerAuthConfigSummaryRegistrationInfoModel] `tfsdk:"registration_info" json:"registration_info,computed"`
}

type ZeroTrustAccessAIControlsMcpServerAuthConfigSummaryConfigModel struct {
	AuthorizationEndpoint types.String `tfsdk:"authorization_endpoint" json:"authorization_endpoint,computed"`
	Issuer                types.String `tfsdk:"issuer" json:"issuer,computed"`
	Resource              types.String `tfsdk:"resource" json:"resource,computed"`
	RevocationEndpoint    types.String `tfsdk:"revocation_endpoint" json:"revocation_endpoint,computed"`
	TokenEndpoint         types.String `tfsdk:"token_endpoint" json:"token_endpoint,computed"`
}

type ZeroTrustAccessAIControlsMcpServerAuthConfigSummaryRegistrationInfoModel struct {
	ClientID                types.String                   `tfsdk:"client_id" json:"client_id,computed"`
	RedirectURIs            customfield.List[types.String] `tfsdk:"redirect_uris" json:"redirect_uris,computed"`
	Scope                   types.String                   `tfsdk:"scope" json:"scope,computed"`
	TokenEndpointAuthMethod types.String                   `tfsdk:"token_endpoint_auth_method" json:"token_endpoint_auth_method,computed"`
}

type ZeroTrustAccessAIControlsMcpServerErrorDetailsModel struct {
	Cause      types.String  `tfsdk:"cause" json:"cause,computed"`
	IsUpstream types.Bool    `tfsdk:"is_upstream" json:"is_upstream,computed"`
	McpCode    types.Float64 `tfsdk:"mcp_code" json:"mcp_code,computed"`
	Retryable  types.Bool    `tfsdk:"retryable" json:"retryable,computed"`
	StatusCode types.Float64 `tfsdk:"status_code" json:"status_code,computed"`
}
