// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_ai_controls_mcp_portal

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessAIControlsMcpPortalsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessAIControlsMcpPortalsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessAIControlsMcpPortalsDataSourceModel struct {
	AccountID types.String                                                                           `tfsdk:"account_id" path:"account_id,required"`
	Search    types.String                                                                           `tfsdk:"search" query:"search,optional"`
	MaxItems  types.Int64                                                                            `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustAccessAIControlsMcpPortalsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessAIControlsMcpPortalsDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessAIControlMcpPortalListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessAIControlMcpPortalListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	return
}

type ZeroTrustAccessAIControlsMcpPortalsResultDataSourceModel struct {
	ID               types.String                                                                            `tfsdk:"id" json:"id,computed"`
	Hostname         types.String                                                                            `tfsdk:"hostname" json:"hostname,computed"`
	Name             types.String                                                                            `tfsdk:"name" json:"name,computed"`
	Servers          customfield.NestedObjectList[ZeroTrustAccessAIControlsMcpPortalsServersDataSourceModel] `tfsdk:"servers" json:"servers,computed"`
	AllowCodeMode    types.Bool                                                                              `tfsdk:"allow_code_mode" json:"allow_code_mode,computed"`
	CreatedAt        timetypes.RFC3339                                                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy        types.String                                                                            `tfsdk:"created_by" json:"created_by,computed"`
	Description      types.String                                                                            `tfsdk:"description" json:"description,computed"`
	ModifiedAt       timetypes.RFC3339                                                                       `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	ModifiedBy       types.String                                                                            `tfsdk:"modified_by" json:"modified_by,computed"`
	SecureWebGateway types.Bool                                                                              `tfsdk:"secure_web_gateway" json:"secure_web_gateway,computed"`
}

type ZeroTrustAccessAIControlsMcpPortalsServersDataSourceModel struct {
	ID                 types.String                                            `tfsdk:"id" json:"id,computed"`
	AuthType           types.String                                            `tfsdk:"auth_type" json:"auth_type,computed"`
	Hostname           types.String                                            `tfsdk:"hostname" json:"hostname,computed"`
	Name               types.String                                            `tfsdk:"name" json:"name,computed"`
	Prompts            customfield.List[customfield.Map[jsontypes.Normalized]] `tfsdk:"prompts" json:"prompts,computed"`
	Tools              customfield.List[customfield.Map[jsontypes.Normalized]] `tfsdk:"tools" json:"tools,computed"`
	UpdatedPrompts     customfield.NormalizedDynamicValue                      `tfsdk:"updated_prompts" json:"updated_prompts,computed"`
	UpdatedTools       customfield.NormalizedDynamicValue                      `tfsdk:"updated_tools" json:"updated_tools,computed"`
	CreatedAt          timetypes.RFC3339                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy          types.String                                            `tfsdk:"created_by" json:"created_by,computed"`
	DefaultDisabled    types.Bool                                              `tfsdk:"default_disabled" json:"default_disabled,computed"`
	Description        types.String                                            `tfsdk:"description" json:"description,computed"`
	Error              types.String                                            `tfsdk:"error" json:"error,computed"`
	LastSuccessfulSync timetypes.RFC3339                                       `tfsdk:"last_successful_sync" json:"last_successful_sync,computed" format:"date-time"`
	LastSynced         timetypes.RFC3339                                       `tfsdk:"last_synced" json:"last_synced,computed" format:"date-time"`
	ModifiedAt         timetypes.RFC3339                                       `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	ModifiedBy         types.String                                            `tfsdk:"modified_by" json:"modified_by,computed"`
	OnBehalf           types.Bool                                              `tfsdk:"on_behalf" json:"on_behalf,computed"`
	Status             types.String                                            `tfsdk:"status" json:"status,computed"`
}
