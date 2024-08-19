// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname_fallback_origin

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnameFallbackOriginResultDataSourceEnvelope struct {
	Result CustomHostnameFallbackOriginDataSourceModel `json:"result,computed"`
}

type CustomHostnameFallbackOriginDataSourceModel struct {
	ZoneID    types.String      `tfsdk:"zone_id" path:"zone_id"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at"`
	Origin    types.String      `tfsdk:"origin" json:"origin"`
	Status    types.String      `tfsdk:"status" json:"status"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at"`
	Errors    *[]types.String   `tfsdk:"errors" json:"errors"`
}
