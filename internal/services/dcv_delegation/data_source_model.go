// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dcv_delegation

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/dcv_delegation"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DCVDelegationResultDataSourceEnvelope struct {
	Result DCVDelegationDataSourceModel `json:"result,computed"`
}

type DCVDelegationDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
	UUID   types.String `tfsdk:"uuid" json:"uuid,optional"`
}

func (m *DCVDelegationDataSourceModel) toReadParams(_ context.Context) (params dcv_delegation.DCVDelegationGetParams, diags diag.Diagnostics) {
	params = dcv_delegation.DCVDelegationGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
