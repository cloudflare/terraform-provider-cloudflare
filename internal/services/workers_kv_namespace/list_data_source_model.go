// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv_namespace

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersKVNamespacesResultListDataSourceEnvelope struct {
	Result *[]*WorkersKVNamespacesResultDataSourceModel `json:"result,computed"`
}

type WorkersKVNamespacesDataSourceModel struct {
	AccountID types.String                                 `tfsdk:"account_id" path:"account_id"`
	Direction types.String                                 `tfsdk:"direction" query:"direction"`
	Order     types.String                                 `tfsdk:"order" query:"order"`
	MaxItems  types.Int64                                  `tfsdk:"max_items"`
	Result    *[]*WorkersKVNamespacesResultDataSourceModel `tfsdk:"result"`
}

type WorkersKVNamespacesResultDataSourceModel struct {
	ID                  types.String `tfsdk:"id" json:"id,computed"`
	Title               types.String `tfsdk:"title" json:"title,computed"`
	SupportsURLEncoding types.Bool   `tfsdk:"supports_url_encoding" json:"supports_url_encoding,computed"`
}
