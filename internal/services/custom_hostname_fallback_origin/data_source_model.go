// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname_fallback_origin

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/custom_hostnames"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnameFallbackOriginResultDataSourceEnvelope struct {
	Result CustomHostnameFallbackOriginDataSourceModel `json:"result,computed"`
}

type CustomHostnameFallbackOriginDataSourceModel struct {
	ZoneID    types.String                   `tfsdk:"zone_id" path:"zone_id,required"`
	CreatedAt timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Origin    types.String                   `tfsdk:"origin" json:"origin,computed"`
	Status    types.String                   `tfsdk:"status" json:"status,computed"`
	UpdatedAt timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Errors    customfield.List[types.String] `tfsdk:"errors" json:"errors,computed"`
}

func (m *CustomHostnameFallbackOriginDataSourceModel) toReadParams(_ context.Context) (params custom_hostnames.FallbackOriginGetParams, diags diag.Diagnostics) {
	params = custom_hostnames.FallbackOriginGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
