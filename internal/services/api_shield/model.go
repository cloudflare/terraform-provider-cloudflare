// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldResultEnvelope struct {
	Result APIShieldModel `json:"result,computed"`
}

type APIShieldModel struct {
	ID                    types.String                            `tfsdk:"id" json:"-,computed"`
	ZoneID                types.String                            `tfsdk:"zone_id" path:"zone_id"`
	AuthIDCharacteristics *[]*APIShieldAuthIDCharacteristicsModel `tfsdk:"auth_id_characteristics" json:"auth_id_characteristics"`
	Errors                *[]*APIShieldErrorsModel                `tfsdk:"errors" json:"errors,computed"`
	Messages              *[]*APIShieldMessagesModel              `tfsdk:"messages" json:"messages,computed"`
	Success               types.Bool                              `tfsdk:"success" json:"success,computed"`
}

type APIShieldAuthIDCharacteristicsModel struct {
	Name types.String `tfsdk:"name" json:"name"`
	Type types.String `tfsdk:"type" json:"type"`
}

type APIShieldErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

type APIShieldMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}
