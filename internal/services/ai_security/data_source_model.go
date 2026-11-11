// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_security

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/ai_security"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AISecurityResultDataSourceEnvelope struct {
	Result AISecurityDataSourceModel `json:"result,computed"`
}

type AISecurityDataSourceModel struct {
	ID      types.String `tfsdk:"id" path:"zone_id,computed"`
	ZoneID  types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
}

func (m *AISecurityDataSourceModel) toReadParams(_ context.Context) (params ai_security.AISecurityGetParams, diags diag.Diagnostics) {
	params = ai_security.AISecurityGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
