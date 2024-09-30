// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_headers

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ManagedHeadersModel struct {
	ID                     types.String                                  `tfsdk:"id" json:"-,computed"`
	ZoneID                 types.String                                  `tfsdk:"zone_id" path:"zone_id,required"`
	ManagedRequestHeaders  *[]*ManagedHeadersManagedRequestHeadersModel  `tfsdk:"managed_request_headers" json:"managed_request_headers,required"`
	ManagedResponseHeaders *[]*ManagedHeadersManagedResponseHeadersModel `tfsdk:"managed_response_headers" json:"managed_response_headers,required"`
}

type ManagedHeadersManagedRequestHeadersModel struct {
	ID      types.String `tfsdk:"id" json:"id,optional"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
}

type ManagedHeadersManagedResponseHeadersModel struct {
	ID      types.String `tfsdk:"id" json:"id,optional"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
}
