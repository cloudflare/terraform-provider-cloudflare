// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldResultEnvelope struct {
	Result APIShieldModel `json:"result"`
}

type APIShieldModel struct {
	ID                    types.String                                                      `tfsdk:"id" json:"-,computed"`
	ZoneID                types.String                                                      `tfsdk:"zone_id" path:"zone_id"`
	AuthIDCharacteristics customfield.NestedObjectList[APIShieldAuthIDCharacteristicsModel] `tfsdk:"auth_id_characteristics" json:"auth_id_characteristics,computed_optional"`
	Success               types.Bool                                                        `tfsdk:"success" json:"success,computed"`
	Errors                customfield.NestedObjectList[APIShieldErrorsModel]                `tfsdk:"errors" json:"errors,computed"`
	Messages              customfield.NestedObjectList[APIShieldMessagesModel]              `tfsdk:"messages" json:"messages,computed"`
}

type APIShieldAuthIDCharacteristicsModel struct {
	Name types.String `tfsdk:"name" json:"name,computed_optional"`
	Type types.String `tfsdk:"type" json:"type,computed_optional"`
}

type APIShieldErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type APIShieldMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}
