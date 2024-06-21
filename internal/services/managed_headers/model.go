// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_headers

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ManagedHeadersModel struct {
	ZoneID                 types.String                                  `tfsdk:"zone_id" path:"zone_id"`
	ManagedRequestHeaders  *[]*ManagedHeadersManagedRequestHeadersModel  `tfsdk:"managed_request_headers" json:"managed_request_headers"`
	ManagedResponseHeaders *[]*ManagedHeadersManagedResponseHeadersModel `tfsdk:"managed_response_headers" json:"managed_response_headers"`
}

type ManagedHeadersManagedRequestHeadersModel struct {
	ID      types.String `tfsdk:"id" json:"id"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
}

type ManagedHeadersManagedResponseHeadersModel struct {
	ID      types.String `tfsdk:"id" json:"id"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
}
