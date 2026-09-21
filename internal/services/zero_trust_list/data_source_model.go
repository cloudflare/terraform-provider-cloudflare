// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_list

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustListResultDataSourceEnvelope struct {
	Result ZeroTrustListDataSourceModel `json:"result,computed"`
}

type ZeroTrustListDataSourceModel struct {
	ID          types.String                                                   `tfsdk:"id" path:"list_id,computed"`
	ListID      types.String                                                   `tfsdk:"list_id" path:"list_id,optional"`
	AccountID   types.String                                                   `tfsdk:"account_id" path:"account_id,optional"`
	CreatedAt   timetypes.RFC3339                                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description types.String                                                   `tfsdk:"description" json:"description,computed"`
	ListCount   types.Float64                                                  `tfsdk:"list_count" json:"count,computed"`
	Name        types.String                                                   `tfsdk:"name" json:"name,computed"`
	Type        types.String                                                   `tfsdk:"type" json:"type,computed"`
	UpdatedAt   timetypes.RFC3339                                              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Items       customfield.NestedObjectSet[ZeroTrustListItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
	Filter      *ZeroTrustListFindOneByDataSourceModel                         `tfsdk:"filter"`
}

func (m *ZeroTrustListDataSourceModel) toReadParams(_ context.Context) (params zero_trust.GatewayListGetParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayListGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustListDataSourceModel) toListParams(_ context.Context) (params zero_trust.GatewayListListParams, diags diag.Diagnostics) {
	mFilterFilter := []interface{}{}
	if m.Filter.Filter != nil {
		for _, item := range *m.Filter.Filter {
			mFilterFilter = append(mFilterFilter, item.ValueString())
		}
	}

	params = zero_trust.GatewayListListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
		Filter:    cloudflare.F(mFilterFilter),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.GatewayListListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(zero_trust.GatewayListListParamsOrderBy(m.Filter.OrderBy.ValueString()))
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}
	if !m.Filter.Type.IsNull() {
		params.Type = cloudflare.F(zero_trust.GatewayListListParamsType(m.Filter.Type.ValueString()))
	}

	return
}

type ZeroTrustListItemsDataSourceModel struct {
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Value       types.String      `tfsdk:"value" json:"value,computed"`
}

type ZeroTrustListFindOneByDataSourceModel struct {
	Direction types.String    `tfsdk:"direction" query:"direction,optional"`
	Filter    *[]types.String `tfsdk:"filter" query:"filter,optional"`
	OrderBy   types.String    `tfsdk:"order_by" query:"order_by,optional"`
	Search    types.String    `tfsdk:"search" query:"search,optional"`
	Type      types.String    `tfsdk:"type" query:"type,optional"`
}
