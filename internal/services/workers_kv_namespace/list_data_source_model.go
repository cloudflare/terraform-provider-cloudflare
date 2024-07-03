// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv_namespace

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersKVNamespacesResultListDataSourceEnvelope struct {
	Result *[]*WorkersKVNamespacesItemsDataSourceModel `json:"result,computed"`
}

type WorkersKVNamespacesDataSourceModel struct {
	AccountID types.String                                `tfsdk:"account_id" path:"account_id"`
	Direction types.String                                `tfsdk:"direction" query:"direction"`
	Order     types.String                                `tfsdk:"order" query:"order"`
	Page      types.Float64                               `tfsdk:"page" query:"page"`
	PerPage   types.Float64                               `tfsdk:"per_page" query:"per_page"`
	MaxItems  types.Int64                                 `tfsdk:"max_items"`
	Items     *[]*WorkersKVNamespacesItemsDataSourceModel `tfsdk:"items"`
}

type WorkersKVNamespacesItemsDataSourceModel struct {
	ID                  types.String `tfsdk:"id" json:"id,computed"`
	Title               types.String `tfsdk:"title" json:"title,computed"`
	SupportsURLEncoding types.Bool   `tfsdk:"supports_url_encoding" json:"supports_url_encoding,computed"`
}
