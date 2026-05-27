// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dls_prefix_binding

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/dls"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DLSPrefixBindingsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[DLSPrefixBindingsResultDataSourceModel] `json:"result,computed"`
}

type DLSPrefixBindingsDataSourceModel struct {
	AccountID types.Int64                                                          `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                          `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[DLSPrefixBindingsResultDataSourceModel] `tfsdk:"result"`
}

func (m *DLSPrefixBindingsDataSourceModel) toListParams(_ context.Context) (params dls.RegionalServicePrefixBindingListParams, diags diag.Diagnostics) {
	params = dls.RegionalServicePrefixBindingListParams{
		AccountID: cloudflare.F(m.AccountID.ValueInt64()),
	}

	return
}

type DLSPrefixBindingsResultDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	CIDR      types.String `tfsdk:"cidr" json:"cidr,computed"`
	PrefixID  types.String `tfsdk:"prefix_id" json:"prefix_id,computed"`
	RegionKey types.String `tfsdk:"region_key" json:"region_key,computed"`
}
