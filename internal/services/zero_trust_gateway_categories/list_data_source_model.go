// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_categories

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayCategoriesListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustGatewayCategoriesListResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustGatewayCategoriesListDataSourceModel struct {
	AccountID types.String                                                                      `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                       `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustGatewayCategoriesListResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustGatewayCategoriesListDataSourceModel) toListParams(_ context.Context) (params zero_trust.GatewayCategoryListParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayCategoryListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustGatewayCategoriesListResultDataSourceModel struct {
	ID            types.Int64                                                                              `tfsdk:"id" json:"id,computed"`
	Beta          types.Bool                                                                               `tfsdk:"beta" json:"beta,computed"`
	Class         types.String                                                                             `tfsdk:"class" json:"class,computed"`
	Description   types.String                                                                             `tfsdk:"description" json:"description,computed"`
	Name          types.String                                                                             `tfsdk:"name" json:"name,computed"`
	Subcategories customfield.NestedObjectList[ZeroTrustGatewayCategoriesListSubcategoriesDataSourceModel] `tfsdk:"subcategories" json:"subcategories,computed"`
}

type ZeroTrustGatewayCategoriesListSubcategoriesDataSourceModel struct {
	ID          types.Int64  `tfsdk:"id" json:"id,computed"`
	Beta        types.Bool   `tfsdk:"beta" json:"beta,computed"`
	Class       types.String `tfsdk:"class" json:"class,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
}
