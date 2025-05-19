// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_message

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/cloudforce_one"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestMessageResultDataSourceEnvelope struct {
	Result CloudforceOneRequestMessageDataSourceModel `json:"result,computed"`
}

type CloudforceOneRequestMessageDataSourceModel struct {
	AccountID         types.String      `tfsdk:"account_id" path:"account_id,required"`
	RequestID         types.String      `tfsdk:"request_id" path:"request_id,required"`
	Page              types.Int64       `tfsdk:"page" json:"page,required"`
	PerPage           types.Int64       `tfsdk:"per_page" json:"per_page,required"`
	After             timetypes.RFC3339 `tfsdk:"after" json:"after,optional" format:"date-time"`
	Before            timetypes.RFC3339 `tfsdk:"before" json:"before,optional" format:"date-time"`
	SortBy            types.String      `tfsdk:"sort_by" json:"sort_by,optional"`
	SortOrder         types.String      `tfsdk:"sort_order" json:"sort_order,optional"`
	Author            types.String      `tfsdk:"author" json:"author,computed"`
	Content           types.String      `tfsdk:"content" json:"content,computed"`
	Created           timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	ID                types.Int64       `tfsdk:"id" json:"id,computed"`
	IsFollowOnRequest types.Bool        `tfsdk:"is_follow_on_request" json:"is_follow_on_request,computed"`
	Updated           timetypes.RFC3339 `tfsdk:"updated" json:"updated,computed" format:"date-time"`
}

func (m *CloudforceOneRequestMessageDataSourceModel) toReadParams(_ context.Context) (params cloudforce_one.RequestMessageGetParams, diags diag.Diagnostics) {
	params = cloudforce_one.RequestMessageGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
