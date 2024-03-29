// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_tags

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessTagsResultEnvelope struct {
	Result ZeroTrustAccessTagsModel `json:"result,computed"`
}

type ZeroTrustAccessTagsModel struct {
	Identifier types.String `tfsdk:"identifier" path:"identifier"`
	TagName    types.String `tfsdk:"tag_name" path:"tag_name"`
	Name       types.String `tfsdk:"name" json:"name"`
	AppCount   types.Int64  `tfsdk:"app_count" json:"app_count"`
	CreatedAt  types.String `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt  types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}
