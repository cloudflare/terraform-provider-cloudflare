// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/managed_transforms"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ManagedTransformsDataSourceModel struct {
	ZoneID                 types.String                                               `tfsdk:"zone_id" path:"zone_id,required"`
	ManagedRequestHeaders  *[]*ManagedTransformsManagedRequestHeadersDataSourceModel  `tfsdk:"managed_request_headers" json:"managed_request_headers,optional"`
	ManagedResponseHeaders *[]*ManagedTransformsManagedResponseHeadersDataSourceModel `tfsdk:"managed_response_headers" json:"managed_response_headers,optional"`
}

func (m *ManagedTransformsDataSourceModel) toReadParams(_ context.Context) (params managed_transforms.ManagedTransformListParams, diags diag.Diagnostics) {
	params = managed_transforms.ManagedTransformListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type ManagedTransformsManagedRequestHeadersDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
}

type ManagedTransformsManagedResponseHeadersDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
}
