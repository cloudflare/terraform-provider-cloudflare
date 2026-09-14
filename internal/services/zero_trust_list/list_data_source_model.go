// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_list

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustListsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustListsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustListsDataSourceModel struct {
	AccountID types.String                                                      `tfsdk:"account_id" path:"account_id,optional"`
	Direction types.String                                                      `tfsdk:"direction" query:"direction,optional"`
	OrderBy   types.String                                                      `tfsdk:"order_by" query:"order_by,optional"`
	Search    types.String                                                      `tfsdk:"search" query:"search,optional"`
	Type      types.String                                                      `tfsdk:"type" query:"type,optional"`
	Filter    *[]jsontypes.Normalized                                           `tfsdk:"filter" query:"filter,optional"`
	MaxItems  types.Int64                                                       `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustListsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustListsDataSourceModel) toListParams(_ context.Context) (params zero_trust.GatewayListListParams, diags diag.Diagnostics) {
	mFilter := []interface{}{}
	if m.Filter != nil {
		for _, item := range *m.Filter {
			mFilter = append(mFilter, item.ValueString())
		}
	}

	params = zero_trust.GatewayListListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
		Filter:    cloudflare.F(mFilter),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.GatewayListListParamsDirection(m.Direction.ValueString()))
	}
	if !m.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(zero_trust.GatewayListListParamsOrderBy(m.OrderBy.ValueString()))
	}
	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}
	if !m.Type.IsNull() {
		params.Type = cloudflare.F(zero_trust.GatewayListListParamsType(m.Type.ValueString()))
	}

	return
}

type ZeroTrustListsResultDataSourceModel struct {
	ID          types.String                                                    `tfsdk:"id" json:"id,computed"`
	ListCount   types.Float64                                                   `tfsdk:"list_count" json:"count,computed"`
	CreatedAt   timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description types.String                                                    `tfsdk:"description" json:"description,computed"`
	Items       customfield.NestedObjectSet[ZeroTrustListsItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
	Name        types.String                                                    `tfsdk:"name" json:"name,computed"`
	Type        types.String                                                    `tfsdk:"type" json:"type,computed"`
	UpdatedAt   timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustListsItemsDataSourceModel struct {
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Value       types.String      `tfsdk:"value" json:"value,computed"`
}
