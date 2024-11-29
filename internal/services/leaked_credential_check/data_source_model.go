// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package leaked_credential_check

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LeakedCredentialCheckResultDataSourceEnvelope struct {
	Result LeakedCredentialCheckDataSourceModel `json:"result,computed"`
}

type LeakedCredentialCheckDataSourceModel struct {
	ZoneID  types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
}

func (m *LeakedCredentialCheckDataSourceModel) toReadParams(_ context.Context) (params cloudflare.LeakedCredentialCheckGetParams, diags diag.Diagnostics) {
	params = cloudflare.LeakedCredentialCheckGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
