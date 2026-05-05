// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_resource_library_category

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustResourceLibraryCategoriesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustResourceLibraryCategoriesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustResourceLibraryCategoriesDataSourceModel struct {
	AccountID types.String                                                                          `tfsdk:"account_id" path:"account_id,required"`
	Limit     types.Int64                                                                           `tfsdk:"limit" query:"limit,computed_optional"`
	Offset    types.Int64                                                                           `tfsdk:"offset" query:"offset,computed_optional"`
	MaxItems  types.Int64                                                                           `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustResourceLibraryCategoriesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustResourceLibraryCategoriesDataSourceModel) toListParams(_ context.Context) (params zero_trust.ResourceLibraryCategoryListParams, diags diag.Diagnostics) {
	params = zero_trust.ResourceLibraryCategoryListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Limit.IsNull() {
		params.Limit = cloudflare.F(m.Limit.ValueInt64())
	}
	if !m.Offset.IsNull() {
		params.Offset = cloudflare.F(m.Offset.ValueInt64())
	}

	return
}

type ZeroTrustResourceLibraryCategoriesResultDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	CreatedAt   types.String `tfsdk:"created_at" json:"created_at,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
}
