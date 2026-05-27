// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dls_prefix_binding

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/dls"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DLSPrefixBindingResultDataSourceEnvelope struct {
	Result DLSPrefixBindingDataSourceModel `json:"result,computed"`
}

type DLSPrefixBindingDataSourceModel struct {
	ID        types.String `tfsdk:"id" path:"binding_id,computed"`
	BindingID types.String `tfsdk:"binding_id" path:"binding_id,required"`
	AccountID types.Int64  `tfsdk:"account_id" path:"account_id,required"`
	CIDR      types.String `tfsdk:"cidr" json:"cidr,computed"`
	PrefixID  types.String `tfsdk:"prefix_id" json:"prefix_id,computed"`
	RegionKey types.String `tfsdk:"region_key" json:"region_key,computed"`
}

func (m *DLSPrefixBindingDataSourceModel) toReadParams(_ context.Context) (params dls.RegionalServicePrefixBindingGetParams, diags diag.Diagnostics) {
	params = dls.RegionalServicePrefixBindingGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueInt64()),
	}

	return
}
