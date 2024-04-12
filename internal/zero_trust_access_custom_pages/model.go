// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_custom_pages

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessCustomPagesResultEnvelope struct {
	Result ZeroTrustAccessCustomPagesModel `json:"result,computed"`
}

type ZeroTrustAccessCustomPagesModel struct {
	Identifier types.String `tfsdk:"identifier" path:"identifier"`
	UUID       types.String `tfsdk:"uuid" path:"uuid"`
	CustomHTML types.String `tfsdk:"custom_html" json:"custom_html"`
	Name       types.String `tfsdk:"name" json:"name"`
	Type       types.String `tfsdk:"type" json:"type"`
	AppCount   types.Int64  `tfsdk:"app_count" json:"app_count"`
	CreatedAt  types.String `tfsdk:"created_at" json:"created_at,computed"`
	UID        types.String `tfsdk:"uid" json:"uid,computed"`
	UpdatedAt  types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}
