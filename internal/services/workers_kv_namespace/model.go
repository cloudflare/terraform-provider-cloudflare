// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv_namespace

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersKVNamespaceResultEnvelope struct {
	Result WorkersKVNamespaceModel `json:"result,computed"`
}

type WorkersKVNamespaceResultDataSourceEnvelope struct {
	Result WorkersKVNamespaceDataSourceModel `json:"result,computed"`
}

type WorkersKVNamespacesResultDataSourceEnvelope struct {
	Result WorkersKVNamespacesDataSourceModel `json:"result,computed"`
}

type WorkersKVNamespaceModel struct {
	ID                  types.String `tfsdk:"id" json:"id,computed"`
	AccountID           types.String `tfsdk:"account_id" path:"account_id"`
	Title               types.String `tfsdk:"title" json:"title"`
	SupportsURLEncoding types.Bool   `tfsdk:"supports_url_encoding" json:"supports_url_encoding,computed"`
}

type WorkersKVNamespaceDataSourceModel struct {
}

type WorkersKVNamespacesDataSourceModel struct {
}
