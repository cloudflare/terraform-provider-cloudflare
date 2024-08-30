// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpull_retention

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/logs"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LogpullRetentionResultDataSourceEnvelope struct {
	Result LogpullRetentionDataSourceModel `json:"result,computed"`
}

type LogpullRetentionDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
	Flag   types.Bool   `tfsdk:"flag" json:"flag"`
}

func (m *LogpullRetentionDataSourceModel) toReadParams(_ context.Context) (params logs.ControlRetentionGetParams, diags diag.Diagnostics) {
	params = logs.ControlRetentionGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
