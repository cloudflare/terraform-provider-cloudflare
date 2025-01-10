// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv_namespace

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/kv"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersKVNamespaceResultDataSourceEnvelope struct {
	Result WorkersKVNamespaceDataSourceModel `json:"result,computed"`
}

type WorkersKVNamespaceResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkersKVNamespaceDataSourceModel] `json:"result,computed"`
}

type WorkersKVNamespaceDataSourceModel struct {
	AccountID           types.String                                `tfsdk:"account_id" path:"account_id,optional"`
	NamespaceID         types.String                                `tfsdk:"namespace_id" path:"namespace_id,optional"`
	ID                  types.String                                `tfsdk:"id" json:"id,computed"`
	SupportsURLEncoding types.Bool                                  `tfsdk:"supports_url_encoding" json:"supports_url_encoding,computed"`
	Title               types.String                                `tfsdk:"title" json:"title,computed"`
	Filter              *WorkersKVNamespaceFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *WorkersKVNamespaceDataSourceModel) toReadParams(_ context.Context) (params kv.NamespaceGetParams, diags diag.Diagnostics) {
	params = kv.NamespaceGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *WorkersKVNamespaceDataSourceModel) toListParams(_ context.Context) (params kv.NamespaceListParams, diags diag.Diagnostics) {
	params = kv.NamespaceListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(kv.NamespaceListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(kv.NamespaceListParamsOrder(m.Filter.Order.ValueString()))
	}

	return
}

type WorkersKVNamespaceFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	Direction types.String `tfsdk:"direction" query:"direction,optional"`
	Order     types.String `tfsdk:"order" query:"order,optional"`
}
