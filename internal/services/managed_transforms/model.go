// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ManagedTransformsResultEnvelope struct {
	Result ManagedTransformsModel `json:"result"`
}

type ManagedTransformsModel struct {
	ID                     types.String                                     `tfsdk:"id" json:"-,computed"`
	ZoneID                 types.String                                     `tfsdk:"zone_id" path:"zone_id,required"`
	ManagedRequestHeaders  *[]*ManagedTransformsManagedRequestHeadersModel  `tfsdk:"managed_request_headers" json:"managed_request_headers,required"`
	ManagedResponseHeaders *[]*ManagedTransformsManagedResponseHeadersModel `tfsdk:"managed_response_headers" json:"managed_response_headers,required"`
}

func (m ManagedTransformsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ManagedTransformsModel) MarshalJSONForUpdate(state ManagedTransformsModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type ManagedTransformsManagedRequestHeadersModel struct {
	ID            types.String                   `tfsdk:"id" json:"id,required"`
	Enabled       types.Bool                     `tfsdk:"enabled" json:"enabled,required"`
	HAsConflict   types.Bool                     `tfsdk:"has_conflict" json:"has_conflict,computed"`
	ConflictsWith customfield.List[types.String] `tfsdk:"conflicts_with" json:"conflicts_with,computed"`
}

type ManagedTransformsManagedResponseHeadersModel struct {
	ID            types.String                   `tfsdk:"id" json:"id,required"`
	Enabled       types.Bool                     `tfsdk:"enabled" json:"enabled,required"`
	HAsConflict   types.Bool                     `tfsdk:"has_conflict" json:"has_conflict,computed"`
	ConflictsWith customfield.List[types.String] `tfsdk:"conflicts_with" json:"conflicts_with,computed"`
}
