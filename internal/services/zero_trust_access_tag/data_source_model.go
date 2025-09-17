// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_tag

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessTagResultDataSourceEnvelope struct {
	Result ZeroTrustAccessTagDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessTagDataSourceModel struct {
	ID        types.String `tfsdk:"id" path:"tag_name,computed"`
	TagName   types.String `tfsdk:"tag_name" path:"tag_name,optional"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
}

func (m *ZeroTrustAccessTagDataSourceModel) toReadParams(_ context.Context) (params zero_trust.AccessTagGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessTagGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
