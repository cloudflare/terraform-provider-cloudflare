// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_smart_routing

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/argo"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ArgoSmartRoutingResultDataSourceEnvelope struct {
	Result ArgoSmartRoutingDataSourceModel `json:"result,computed"`
}

type ArgoSmartRoutingDataSourceModel struct {
	ID         types.String      `tfsdk:"id" path:"zone_id,computed"`
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Editable   types.Bool        `tfsdk:"editable" json:"editable,computed"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Value      types.String      `tfsdk:"value" json:"value,computed"`
}

func (m *ArgoSmartRoutingDataSourceModel) toReadParams(_ context.Context) (params argo.SmartRoutingGetParams, diags diag.Diagnostics) {
	params = argo.SmartRoutingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
