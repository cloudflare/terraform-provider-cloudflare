// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname_fallback_origin

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnameFallbackOriginResultEnvelope struct {
	Result CustomHostnameFallbackOriginModel `json:"result"`
}

type CustomHostnameFallbackOriginModel struct {
	ID        types.String      `tfsdk:"id" json:"-,computed"`
	ZoneID    types.String      `tfsdk:"zone_id" path:"zone_id"`
	Origin    types.String      `tfsdk:"origin" json:"origin"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Errors    types.List        `tfsdk:"errors" json:"errors,computed"`
}