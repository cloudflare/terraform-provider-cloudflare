// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_headers

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ManagedHeadersDataSourceModel struct {
	ZoneID                 types.String                                            `tfsdk:"zone_id" path:"zone_id"`
	ManagedRequestHeaders  *[]*ManagedHeadersManagedRequestHeadersDataSourceModel  `tfsdk:"managed_request_headers" json:"managed_request_headers"`
	ManagedResponseHeaders *[]*ManagedHeadersManagedResponseHeadersDataSourceModel `tfsdk:"managed_response_headers" json:"managed_response_headers"`
}

type ManagedHeadersManagedRequestHeadersDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
}

type ManagedHeadersManagedResponseHeadersDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
}
