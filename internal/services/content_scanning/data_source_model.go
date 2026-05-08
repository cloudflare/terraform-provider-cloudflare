// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package content_scanning

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/content_scanning"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ContentScanningResultDataSourceEnvelope struct {
	Result ContentScanningDataSourceModel `json:"result,computed"`
}

type ContentScanningDataSourceModel struct {
	ZoneID   types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Modified types.String `tfsdk:"modified" json:"modified,computed"`
	Value    types.String `tfsdk:"value" json:"value,computed"`
}

func (m *ContentScanningDataSourceModel) toReadParams(_ context.Context) (params content_scanning.ContentScanningGetParams, diags diag.Diagnostics) {
	params = content_scanning.ContentScanningGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
