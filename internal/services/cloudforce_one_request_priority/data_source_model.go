// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_priority

import (
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestPriorityResultDataSourceEnvelope struct {
Result CloudforceOneRequestPriorityDataSourceModel `json:"result,computed"`
}

type CloudforceOneRequestPriorityDataSourceModel struct {
AccountIdentifier types.String `tfsdk:"account_identifier" path:"account_identifier,required"`
PriorityIdentifer types.String `tfsdk:"priority_identifer" path:"priority_identifer,required"`
Completed timetypes.RFC3339 `tfsdk:"completed" json:"completed,computed" format:"date-time"`
Content types.String `tfsdk:"content" json:"content,computed"`
Created timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
ID types.String `tfsdk:"id" json:"id,computed"`
MessageTokens types.Int64 `tfsdk:"message_tokens" json:"message_tokens,computed"`
Priority timetypes.RFC3339 `tfsdk:"priority" json:"priority,computed" format:"date-time"`
ReadableID types.String `tfsdk:"readable_id" json:"readable_id,computed"`
Request types.String `tfsdk:"request" json:"request,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
Summary types.String `tfsdk:"summary" json:"summary,computed"`
TLP types.String `tfsdk:"tlp" json:"tlp,computed"`
Tokens types.Int64 `tfsdk:"tokens" json:"tokens,computed"`
Updated timetypes.RFC3339 `tfsdk:"updated" json:"updated,computed" format:"date-time"`
}
