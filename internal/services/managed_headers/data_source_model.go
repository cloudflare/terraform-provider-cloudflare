// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_headers

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/managed_headers"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ManagedHeadersDataSourceModel struct {
	ZoneID                 types.String                                            `tfsdk:"zone_id" path:"zone_id,required"`
	ManagedRequestHeaders  *[]*ManagedHeadersManagedRequestHeadersDataSourceModel  `tfsdk:"managed_request_headers" json:"managed_request_headers,optional"`
	ManagedResponseHeaders *[]*ManagedHeadersManagedResponseHeadersDataSourceModel `tfsdk:"managed_response_headers" json:"managed_response_headers,optional"`
}

func (m *ManagedHeadersDataSourceModel) toReadParams(_ context.Context) (params managed_headers.ManagedHeaderListParams, diags diag.Diagnostics) {
	params = managed_headers.ManagedHeaderListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type ManagedHeadersManagedRequestHeadersDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
}

type ManagedHeadersManagedResponseHeadersDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
}
