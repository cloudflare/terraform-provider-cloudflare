// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldResultEnvelope struct {
	Result APIShieldModel `json:"result"`
}

type APIShieldModel struct {
	ID                    types.String                                         `tfsdk:"id" json:"-,computed"`
	ZoneID                types.String                                         `tfsdk:"zone_id" path:"zone_id,required"`
	AuthIDCharacteristics *[]*APIShieldAuthIDCharacteristicsModel              `tfsdk:"auth_id_characteristics" json:"auth_id_characteristics,required"`
	Success               types.Bool                                           `tfsdk:"success" json:"success,computed,no_refresh"`
	Errors                customfield.NestedObjectList[APIShieldErrorsModel]   `tfsdk:"errors" json:"errors,computed,no_refresh"`
	Messages              customfield.NestedObjectList[APIShieldMessagesModel] `tfsdk:"messages" json:"messages,computed,no_refresh"`
}

func (m APIShieldModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m APIShieldModel) MarshalJSONForUpdate(state APIShieldModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type APIShieldAuthIDCharacteristicsModel struct {
	Name types.String `tfsdk:"name" json:"name,required"`
	Type types.String `tfsdk:"type" json:"type,required"`
}

type APIShieldErrorsModel struct {
	Code             types.Int64                                          `tfsdk:"code" json:"code,computed"`
	Message          types.String                                         `tfsdk:"message" json:"message,computed"`
	DocumentationURL types.String                                         `tfsdk:"documentation_url" json:"documentation_url,computed"`
	Source           customfield.NestedObject[APIShieldErrorsSourceModel] `tfsdk:"source" json:"source,computed"`
}

type APIShieldErrorsSourceModel struct {
	Pointer types.String `tfsdk:"pointer" json:"pointer,computed"`
}

type APIShieldMessagesModel struct {
	Code             types.Int64                                            `tfsdk:"code" json:"code,computed"`
	Message          types.String                                           `tfsdk:"message" json:"message,computed"`
	DocumentationURL types.String                                           `tfsdk:"documentation_url" json:"documentation_url,computed"`
	Source           customfield.NestedObject[APIShieldMessagesSourceModel] `tfsdk:"source" json:"source,computed"`
}

type APIShieldMessagesSourceModel struct {
	Pointer types.String `tfsdk:"pointer" json:"pointer,computed"`
}
