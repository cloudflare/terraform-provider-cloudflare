// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_token

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/ai_search"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AISearchTokenResultDataSourceEnvelope struct {
	Result AISearchTokenDataSourceModel `json:"result,computed"`
}

type AISearchTokenDataSourceModel struct {
	ID         types.String                           `tfsdk:"id" path:"id,computed_optional"`
	AccountID  types.String                           `tfsdk:"account_id" path:"account_id,required"`
	CfAPIID    types.String                           `tfsdk:"cf_api_id" json:"cf_api_id,computed"`
	CreatedAt  timetypes.RFC3339                      `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CreatedBy  types.String                           `tfsdk:"created_by" json:"created_by,computed"`
	Enabled    types.Bool                             `tfsdk:"enabled" json:"enabled,computed"`
	Legacy     types.Bool                             `tfsdk:"legacy" json:"legacy,computed"`
	ModifiedAt timetypes.RFC3339                      `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	ModifiedBy types.String                           `tfsdk:"modified_by" json:"modified_by,computed"`
	Name       types.String                           `tfsdk:"name" json:"name,computed"`
	Filter     *AISearchTokenFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *AISearchTokenDataSourceModel) toReadParams(_ context.Context) (params ai_search.TokenReadParams, diags diag.Diagnostics) {
	params = ai_search.TokenReadParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *AISearchTokenDataSourceModel) toListParams(_ context.Context) (params ai_search.TokenListParams, diags diag.Diagnostics) {
	params = ai_search.TokenListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.OrderBy.IsNull() {
		params.OrderBy = cloudflare.F(ai_search.TokenListParamsOrderBy(m.Filter.OrderBy.ValueString()))
	}
	if !m.Filter.OrderByDirection.IsNull() {
		params.OrderByDirection = cloudflare.F(ai_search.TokenListParamsOrderByDirection(m.Filter.OrderByDirection.ValueString()))
	}

	return
}

type AISearchTokenFindOneByDataSourceModel struct {
	OrderBy          types.String `tfsdk:"order_by" query:"order_by,computed_optional"`
	OrderByDirection types.String `tfsdk:"order_by_direction" query:"order_by_direction,computed_optional"`
}
