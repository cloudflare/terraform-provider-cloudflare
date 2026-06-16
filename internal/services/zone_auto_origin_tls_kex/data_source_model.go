// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_auto_origin_tls_kex

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/ssl"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneAutoOriginTLSKexResultDataSourceEnvelope struct {
	Result ZoneAutoOriginTLSKexDataSourceModel `json:"result,computed"`
}

type ZoneAutoOriginTLSKexDataSourceModel struct {
	ID         types.String      `tfsdk:"id" path:"zone_id,computed"`
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled    types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (m *ZoneAutoOriginTLSKexDataSourceModel) toReadParams(_ context.Context) (params ssl.AutoOriginTLSKexGetParams, diags diag.Diagnostics) {
	params = ssl.AutoOriginTLSKexGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
