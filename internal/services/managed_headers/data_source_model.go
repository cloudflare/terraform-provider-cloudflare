// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_headers

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/managed_headers"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ManagedHeadersDataSourceModel struct {
	ZoneID                 types.String                                            `tfsdk:"zone_id" path:"zone_id"`
	ManagedRequestHeaders  *[]*ManagedHeadersManagedRequestHeadersDataSourceModel  `tfsdk:"managed_request_headers" json:"managed_request_headers"`
	ManagedResponseHeaders *[]*ManagedHeadersManagedResponseHeadersDataSourceModel `tfsdk:"managed_response_headers" json:"managed_response_headers"`
}

func (m *ManagedHeadersDataSourceModel) toReadParams() (params managed_headers.ManagedHeaderListParams, diags diag.Diagnostics) {
	params = managed_headers.ManagedHeaderListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type ManagedHeadersManagedRequestHeadersDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
}

type ManagedHeadersManagedResponseHeadersDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
}
