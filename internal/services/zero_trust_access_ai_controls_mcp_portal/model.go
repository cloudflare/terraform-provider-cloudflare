// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_ai_controls_mcp_portal

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessAIControlsMcpPortalResultEnvelope struct {
	Result ZeroTrustAccessAIControlsMcpPortalModel `json:"result"`
}

type ZeroTrustAccessAIControlsMcpPortalModel struct {
	ID               types.String                                                                 `tfsdk:"id" json:"id,required"`
	AccountID        types.String                                                                 `tfsdk:"account_id" path:"account_id,required"`
	Hostname         types.String                                                                 `tfsdk:"hostname" json:"hostname,required"`
	Name             types.String                                                                 `tfsdk:"name" json:"name,required"`
	Description      types.String                                                                 `tfsdk:"description" json:"description,optional"`
	SecureWebGateway types.Bool                                                                   `tfsdk:"secure_web_gateway" json:"secure_web_gateway,optional"`
	Servers          customfield.NestedObjectList[ZeroTrustAccessAIControlsMcpPortalServersModel] `tfsdk:"servers" json:"servers,computed_optional"`
	CreatedAt        timetypes.RFC3339                                                            `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy        types.String                                                                 `tfsdk:"created_by" json:"created_by,computed"`
	ModifiedAt       timetypes.RFC3339                                                            `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	ModifiedBy       types.String                                                                 `tfsdk:"modified_by" json:"modified_by,computed"`
}

func (m ZeroTrustAccessAIControlsMcpPortalModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessAIControlsMcpPortalModel) MarshalJSONForUpdate(state ZeroTrustAccessAIControlsMcpPortalModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustAccessAIControlsMcpPortalServersModel struct {
	ServerID        types.String                                                     `tfsdk:"server_id" json:"server_id,required,no_refresh"`
	DefaultDisabled types.Bool                                                       `tfsdk:"default_disabled" json:"default_disabled,computed_optional"`
	OnBehalf        types.Bool                                                       `tfsdk:"on_behalf" json:"on_behalf,computed_optional"`
	UpdatedPrompts  *[]*ZeroTrustAccessAIControlsMcpPortalServersUpdatedPromptsModel `tfsdk:"updated_prompts" json:"updated_prompts,optional,no_refresh"`
	UpdatedTools    *[]*ZeroTrustAccessAIControlsMcpPortalServersUpdatedToolsModel   `tfsdk:"updated_tools" json:"updated_tools,optional,no_refresh"`
}

type ZeroTrustAccessAIControlsMcpPortalServersUpdatedPromptsModel struct {
	Name        types.String `tfsdk:"name" json:"name,required"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
}

type ZeroTrustAccessAIControlsMcpPortalServersUpdatedToolsModel struct {
	Name        types.String `tfsdk:"name" json:"name,required"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
}
