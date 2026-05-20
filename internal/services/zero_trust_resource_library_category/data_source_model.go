// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_resource_library_category

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustResourceLibraryCategoryResultDataSourceEnvelope struct {
	Result ZeroTrustResourceLibraryCategoryDataSourceModel `json:"result,computed"`
}

type ZeroTrustResourceLibraryCategoryDataSourceModel struct {
	AccountID   types.String `tfsdk:"account_id" path:"account_id,required"`
	ID          types.String `tfsdk:"id" path:"id,required"`
	CreatedAt   types.String `tfsdk:"created_at" json:"created_at,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
}

func (m *ZeroTrustResourceLibraryCategoryDataSourceModel) toReadParams(_ context.Context) (params zero_trust.ResourceLibraryCategoryGetParams, diags diag.Diagnostics) {
	params = zero_trust.ResourceLibraryCategoryGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
