// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsResultEnvelope struct {
	Result AuthenticatedOriginPullsModel `json:"result,computed"`
}

type AuthenticatedOriginPullsModel struct {
	ZoneID   types.String                            `tfsdk:"zone_id" path:"zone_id"`
	Hostname types.String                            `tfsdk:"hostname" path:"hostname"`
	Config   *[]*AuthenticatedOriginPullsConfigModel `tfsdk:"config" json:"config"`
}

type AuthenticatedOriginPullsConfigModel struct {
	CERTID   types.String `tfsdk:"cert_id" json:"cert_id"`
	Enabled  types.Bool   `tfsdk:"enabled" json:"enabled"`
	Hostname types.String `tfsdk:"hostname" json:"hostname"`
}
