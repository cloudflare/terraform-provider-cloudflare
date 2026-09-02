// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ct_alerting

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zones"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CTAlertingResultDataSourceEnvelope struct {
	Result CTAlertingDataSourceModel `json:"result,computed"`
}

type CTAlertingDataSourceModel struct {
	ID      types.String                   `tfsdk:"id" path:"zone_id,computed"`
	ZoneID  types.String                   `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled types.Bool                     `tfsdk:"enabled" json:"enabled,computed"`
	Emails  customfield.List[types.String] `tfsdk:"emails" json:"emails,computed"`
}

func (m *CTAlertingDataSourceModel) toReadParams(_ context.Context) (params zones.CTAlertingGetParams, diags diag.Diagnostics) {
	params = zones.CTAlertingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
