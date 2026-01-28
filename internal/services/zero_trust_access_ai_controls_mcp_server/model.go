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
	ID                 types.String                                            `tfsdk:"id" json:"id,required"`
	AccountID          types.String                                            `tfsdk:"account_id" path:"account_id,required"`
	AuthType           types.String                                            `tfsdk:"auth_type" json:"auth_type,required"`
	Hostname           types.String                                            `tfsdk:"hostname" json:"hostname,required"`
	Name               types.String                                            `tfsdk:"name" json:"name,required"`
	AuthCredentials    types.String                                            `tfsdk:"auth_credentials" json:"auth_credentials,optional,no_refresh"`
	Description        types.String                                            `tfsdk:"description" json:"description,optional"`
	CreatedAt          timetypes.RFC3339                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy          types.String                                            `tfsdk:"created_by" json:"created_by,computed"`
	Error              types.String                                            `tfsdk:"error" json:"error,computed"`
	LastSuccessfulSync timetypes.RFC3339                                       `tfsdk:"last_successful_sync" json:"last_successful_sync,computed" format:"date-time"`
	LastSynced         timetypes.RFC3339                                       `tfsdk:"last_synced" json:"last_synced,computed" format:"date-time"`
	ModifiedAt         timetypes.RFC3339                                       `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	ModifiedBy         types.String                                            `tfsdk:"modified_by" json:"modified_by,computed"`
	Status             types.String                                            `tfsdk:"status" json:"status,computed"`
	Prompts            customfield.List[customfield.Map[jsontypes.Normalized]] `tfsdk:"prompts" json:"prompts,computed"`
	Tools              customfield.List[customfield.Map[jsontypes.Normalized]] `tfsdk:"tools" json:"tools,computed"`
}

func (m ZeroTrustAccessAIControlsMcpServerModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessAIControlsMcpServerModel) MarshalJSONForUpdate(state ZeroTrustAccessAIControlsMcpServerModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
