// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/api_gateway"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldResultDataSourceEnvelope struct {
	Result APIShieldDataSourceModel `json:"result,computed"`
}

type APIShieldDataSourceModel struct {
	ZoneID                types.String                                                                `tfsdk:"zone_id" path:"zone_id,required"`
	Properties            *[]types.String                                                             `tfsdk:"properties" query:"properties,optional"`
	AuthIDCharacteristics customfield.NestedObjectList[APIShieldAuthIDCharacteristicsDataSourceModel] `tfsdk:"auth_id_characteristics" json:"auth_id_characteristics,computed"`
}

func (m *APIShieldDataSourceModel) toReadParams(_ context.Context) (params api_gateway.ConfigurationGetParams, diags diag.Diagnostics) {
	params = api_gateway.ConfigurationGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type APIShieldAuthIDCharacteristicsDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}
