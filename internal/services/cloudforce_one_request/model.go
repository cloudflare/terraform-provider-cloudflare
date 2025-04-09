// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestResultEnvelope struct {
Result CloudforceOneRequestModel `json:"result"`
}

type CloudforceOneRequestModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
AccountIdentifier types.String `tfsdk:"account_identifier" path:"account_identifier,required"`
Content types.String `tfsdk:"content" json:"content,optional"`
Priority types.String `tfsdk:"priority" json:"priority,optional"`
RequestType types.String `tfsdk:"request_type" json:"request_type,optional"`
Summary types.String `tfsdk:"summary" json:"summary,optional"`
TLP types.String `tfsdk:"tlp" json:"tlp,optional"`
Completed timetypes.RFC3339 `tfsdk:"completed" json:"completed,computed" format:"date-time"`
Created timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
MessageTokens types.Int64 `tfsdk:"message_tokens" json:"message_tokens,computed"`
ReadableID types.String `tfsdk:"readable_id" json:"readable_id,computed"`
Request types.String `tfsdk:"request" json:"request,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
Tokens types.Int64 `tfsdk:"tokens" json:"tokens,computed"`
Updated timetypes.RFC3339 `tfsdk:"updated" json:"updated,computed" format:"date-time"`
}

func (m CloudforceOneRequestModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m CloudforceOneRequestModel) MarshalJSONForUpdate(state CloudforceOneRequestModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
