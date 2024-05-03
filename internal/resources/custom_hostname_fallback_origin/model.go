// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname_fallback_origin

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnameFallbackOriginResultEnvelope struct {
	Result CustomHostnameFallbackOriginModel `json:"result,computed"`
}

type CustomHostnameFallbackOriginModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
	Origin types.String `tfsdk:"origin" json:"origin"`
}
