// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/cloudforce_one"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CloudforceOneRequestsResultDataSourceModel] `json:"result,computed"`
}

type CloudforceOneRequestsDataSourceModel struct {
	AccountID types.String                                                             `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[CloudforceOneRequestsResultDataSourceModel] `tfsdk:"result"`
}

func (m *CloudforceOneRequestsDataSourceModel) toListParams(_ context.Context) (params cloudforce_one.RequestListParams, diags diag.Diagnostics) {
	params = cloudforce_one.RequestListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type CloudforceOneRequestsResultDataSourceModel struct {
	ID            types.String      `tfsdk:"id" json:"id,computed"`
	Created       timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Priority      types.String      `tfsdk:"priority" json:"priority,computed"`
	Request       types.String      `tfsdk:"request" json:"request,computed"`
	Summary       types.String      `tfsdk:"summary" json:"summary,computed"`
	TLP           types.String      `tfsdk:"tlp" json:"tlp,computed"`
	Updated       timetypes.RFC3339 `tfsdk:"updated" json:"updated,computed" format:"date-time"`
	Completed     timetypes.RFC3339 `tfsdk:"completed" json:"completed,computed" format:"date-time"`
	MessageTokens types.Int64       `tfsdk:"message_tokens" json:"message_tokens,computed"`
	ReadableID    types.String      `tfsdk:"readable_id" json:"readable_id,computed"`
	Status        types.String      `tfsdk:"status" json:"status,computed"`
	Tokens        types.Int64       `tfsdk:"tokens" json:"tokens,computed"`
}
