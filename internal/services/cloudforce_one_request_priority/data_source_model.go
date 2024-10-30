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
	AccountIdentifier types.String      `tfsdk:"account_identifier" path:"account_identifier,required"`
	PriorityIdentifer types.String      `tfsdk:"priority_identifer" path:"priority_identifer,required"`
	Completed         timetypes.RFC3339 `tfsdk:"completed" json:"completed,optional" format:"date-time"`
	Content           types.String      `tfsdk:"content" json:"content,optional"`
	Created           timetypes.RFC3339 `tfsdk:"created" json:"created,optional" format:"date-time"`
	ID                types.String      `tfsdk:"id" json:"id,optional"`
	MessageTokens     types.Int64       `tfsdk:"message_tokens" json:"message_tokens,optional"`
	Priority          timetypes.RFC3339 `tfsdk:"priority" json:"priority,optional" format:"date-time"`
	ReadableID        types.String      `tfsdk:"readable_id" json:"readable_id,optional"`
	Request           types.String      `tfsdk:"request" json:"request,optional"`
	Status            types.String      `tfsdk:"status" json:"status,optional"`
	Summary           types.String      `tfsdk:"summary" json:"summary,optional"`
	TLP               types.String      `tfsdk:"tlp" json:"tlp,optional"`
	Tokens            types.Int64       `tfsdk:"tokens" json:"tokens,optional"`
	Updated           timetypes.RFC3339 `tfsdk:"updated" json:"updated,optional" format:"date-time"`
}
