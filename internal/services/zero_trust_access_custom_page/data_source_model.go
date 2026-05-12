// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_custom_page

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessCustomPageResultDataSourceEnvelope struct {
	Result ZeroTrustAccessCustomPageDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessCustomPageDataSourceModel struct {
	ID           types.String `tfsdk:"id" path:"custom_page_id,computed"`
	CustomPageID types.String `tfsdk:"custom_page_id" path:"custom_page_id,required"`
	AccountID    types.String `tfsdk:"account_id" path:"account_id,required"`
	CustomHTML   types.String `tfsdk:"custom_html" json:"custom_html,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
	Type         types.String `tfsdk:"type" json:"type,computed"`
	UID          types.String `tfsdk:"uid" json:"uid,computed"`
}

func (m *ZeroTrustAccessCustomPageDataSourceModel) toReadParams(_ context.Context) (params zero_trust.AccessCustomPageGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessCustomPageGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
