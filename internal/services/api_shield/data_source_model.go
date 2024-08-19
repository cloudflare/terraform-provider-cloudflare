// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldResultDataSourceEnvelope struct {
	Result APIShieldDataSourceModel `json:"result,computed"`
}

type APIShieldDataSourceModel struct {
	ZoneID                types.String                                      `tfsdk:"zone_id" path:"zone_id"`
	Properties            *[]types.String                                   `tfsdk:"properties" query:"properties"`
	AuthIDCharacteristics *[]*APIShieldAuthIDCharacteristicsDataSourceModel `tfsdk:"auth_id_characteristics" json:"auth_id_characteristics"`
}

type APIShieldAuthIDCharacteristicsDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}
