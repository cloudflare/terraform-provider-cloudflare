// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_priority

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestPriorityResultEnvelope struct {
	Result CloudforceOneRequestPriorityModel `json:"result"`
}

type CloudforceOneRequestPriorityModel struct {
	ID                types.String      `tfsdk:"id" json:"id,computed"`
	AccountIdentifier types.String      `tfsdk:"account_identifier" path:"account_identifier,required"`
	Priority          types.Int64       `tfsdk:"priority" json:"priority,required"`
	Requirement       types.String      `tfsdk:"requirement" json:"requirement,required"`
	Tlp               types.String      `tfsdk:"tlp" json:"tlp,required"`
	Labels            *[]types.String   `tfsdk:"labels" json:"labels,required"`
	Completed         timetypes.RFC3339 `tfsdk:"completed" json:"completed,computed" format:"date-time"`
	Content           types.String      `tfsdk:"content" json:"content,computed"`
	Created           timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	MessageTokens     types.Int64       `tfsdk:"message_tokens" json:"message_tokens,computed"`
	ReadableID        types.String      `tfsdk:"readable_id" json:"readable_id,computed"`
	Request           types.String      `tfsdk:"request" json:"request,computed"`
	Status            types.String      `tfsdk:"status" json:"status,computed"`
	Summary           types.String      `tfsdk:"summary" json:"summary,computed"`
	Tokens            types.Int64       `tfsdk:"tokens" json:"tokens,computed"`
	Updated           timetypes.RFC3339 `tfsdk:"updated" json:"updated,computed" format:"date-time"`
}

func (m CloudforceOneRequestPriorityModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CloudforceOneRequestPriorityModel) MarshalJSONForUpdate(state CloudforceOneRequestPriorityModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
