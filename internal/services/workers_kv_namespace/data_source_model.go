// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv_namespace

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersKVNamespaceResultDataSourceEnvelope struct {
	Result WorkersKVNamespaceDataSourceModel `json:"result,computed"`
}

type WorkersKVNamespaceResultListDataSourceEnvelope struct {
	Result *[]*WorkersKVNamespaceDataSourceModel `json:"result,computed"`
}

type WorkersKVNamespaceDataSourceModel struct {
	AccountID           types.String                                `tfsdk:"account_id" path:"account_id"`
	NamespaceID         types.String                                `tfsdk:"namespace_id" path:"namespace_id"`
	ID                  types.String                                `tfsdk:"id" json:"id,computed"`
	SupportsURLEncoding types.Bool                                  `tfsdk:"supports_url_encoding" json:"supports_url_encoding,computed"`
	Title               types.String                                `tfsdk:"title" json:"title,computed"`
	Filter              *WorkersKVNamespaceFindOneByDataSourceModel `tfsdk:"filter"`
}

type WorkersKVNamespaceFindOneByDataSourceModel struct {
	AccountID types.String  `tfsdk:"account_id" path:"account_id"`
	Direction types.String  `tfsdk:"direction" query:"direction"`
	Order     types.String  `tfsdk:"order" query:"order"`
	Page      types.Float64 `tfsdk:"page" query:"page"`
	PerPage   types.Float64 `tfsdk:"per_page" query:"per_page"`
}
