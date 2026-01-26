// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_ai_controls_mcp_server

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

type ZeroTrustAccessAIControlsMcpServersResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessAIControlsMcpServersResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessAIControlsMcpServersDataSourceModel struct {
	AccountID types.String                                                                           `tfsdk:"account_id" path:"account_id,required"`
	Search    types.String                                                                           `tfsdk:"search" query:"search,optional"`
	MaxItems  types.Int64                                                                            `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustAccessAIControlsMcpServersResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessAIControlsMcpServersDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessAIControlMcpServerListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessAIControlMcpServerListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	return
}

type ZeroTrustAccessAIControlsMcpServersResultDataSourceModel struct {
	ID                 types.String                                            `tfsdk:"id" json:"id,computed"`
	AuthType           types.String                                            `tfsdk:"auth_type" json:"auth_type,computed"`
	Hostname           types.String                                            `tfsdk:"hostname" json:"hostname,computed"`
	Name               types.String                                            `tfsdk:"name" json:"name,computed"`
	Prompts            customfield.List[customfield.Map[jsontypes.Normalized]] `tfsdk:"prompts" json:"prompts,computed"`
	Tools              customfield.List[customfield.Map[jsontypes.Normalized]] `tfsdk:"tools" json:"tools,computed"`
	CreatedAt          timetypes.RFC3339                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy          types.String                                            `tfsdk:"created_by" json:"created_by,computed"`
	Description        types.String                                            `tfsdk:"description" json:"description,computed"`
	Error              types.String                                            `tfsdk:"error" json:"error,computed"`
	LastSuccessfulSync timetypes.RFC3339                                       `tfsdk:"last_successful_sync" json:"last_successful_sync,computed" format:"date-time"`
	LastSynced         timetypes.RFC3339                                       `tfsdk:"last_synced" json:"last_synced,computed" format:"date-time"`
	ModifiedAt         timetypes.RFC3339                                       `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	ModifiedBy         types.String                                            `tfsdk:"modified_by" json:"modified_by,computed"`
	Status             types.String                                            `tfsdk:"status" json:"status,computed"`
}
