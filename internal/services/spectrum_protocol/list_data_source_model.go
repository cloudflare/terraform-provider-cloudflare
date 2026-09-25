// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_protocol

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/spectrum"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SpectrumProtocolsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SpectrumProtocolsResultDataSourceModel] `json:"result,computed"`
}

type SpectrumProtocolsDataSourceModel struct {
	ZoneID   types.String                                                         `tfsdk:"zone_id" path:"zone_id,required"`
	MaxItems types.Int64                                                          `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[SpectrumProtocolsResultDataSourceModel] `tfsdk:"result"`
}

func (m *SpectrumProtocolsDataSourceModel) toListParams(_ context.Context) (params spectrum.ProtocolListParams, diags diag.Diagnostics) {
	params = spectrum.ProtocolListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type SpectrumProtocolsResultDataSourceModel struct {
	Description types.String                  `tfsdk:"description" json:"description,computed"`
	Name        types.String                  `tfsdk:"name" json:"name,computed"`
	Ports       customfield.List[types.Int64] `tfsdk:"ports" json:"ports,computed"`
	Transport   types.String                  `tfsdk:"transport" json:"transport,computed"`
}
