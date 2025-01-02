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

type ZeroTrustGatewayCategoriesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustGatewayCategoriesDataSourceModel] `json:"result,computed"`
}

type ZeroTrustGatewayCategoriesDataSourceModel struct {
	Beta          types.Bool                                                                           `tfsdk:"beta" json:"beta,computed"`
	Class         types.String                                                                         `tfsdk:"class" json:"class,computed"`
	Description   types.String                                                                         `tfsdk:"description" json:"description,computed"`
	ID            types.Int64                                                                          `tfsdk:"id" json:"id,computed"`
	Name          types.String                                                                         `tfsdk:"name" json:"name,computed"`
	Subcategories customfield.NestedObjectList[ZeroTrustGatewayCategoriesSubcategoriesDataSourceModel] `tfsdk:"subcategories" json:"subcategories,computed"`
	Filter        *ZeroTrustGatewayCategoriesFindOneByDataSourceModel                                  `tfsdk:"filter"`
}

func (m *ZeroTrustGatewayCategoriesDataSourceModel) toListParams(_ context.Context) (params zero_trust.GatewayCategoryListParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayCategoryListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ZeroTrustGatewayCategoriesSubcategoriesDataSourceModel struct {
	ID          types.Int64  `tfsdk:"id" json:"id,computed"`
	Beta        types.Bool   `tfsdk:"beta" json:"beta,computed"`
	Class       types.String `tfsdk:"class" json:"class,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustGatewayCategoriesFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
