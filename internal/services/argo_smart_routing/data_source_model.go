// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_smart_routing

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/argo"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ArgoSmartRoutingResultDataSourceEnvelope struct {
	Result ArgoSmartRoutingDataSourceModel `json:"result,computed"`
}

type ArgoSmartRoutingDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
}

func (m *ArgoSmartRoutingDataSourceModel) toReadParams(_ context.Context) (params argo.SmartRoutingGetParams, diags diag.Diagnostics) {
	params = argo.SmartRoutingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
