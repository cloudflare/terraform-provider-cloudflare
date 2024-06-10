// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldResultEnvelope struct {
	Result APIShieldModel `json:"result,computed"`
}

type APIShieldModel struct {
	ZoneID                types.String                            `tfsdk:"zone_id" path:"zone_id"`
	AuthIDCharacteristics *[]*APIShieldAuthIDCharacteristicsModel `tfsdk:"auth_id_characteristics" json:"auth_id_characteristics"`
}

type APIShieldAuthIDCharacteristicsModel struct {
	Name types.String `tfsdk:"name" json:"name"`
	Type types.String `tfsdk:"type" json:"type"`
}
