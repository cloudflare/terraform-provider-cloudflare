// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldResultEnvelope struct {
	Result APIShieldModel `json:"result"`
}

type APIShieldModel struct {
	ID                    types.String                            `tfsdk:"id" json:"-,computed"`
	ZoneID                types.String                            `tfsdk:"zone_id" path:"zone_id,required"`
	AuthIDCharacteristics *[]*APIShieldAuthIDCharacteristicsModel `tfsdk:"auth_id_characteristics" json:"auth_id_characteristics,required"`
	Success               types.Bool                              `tfsdk:"success" json:"success,computed"`
	Errors                customfield.List[jsontypes.Normalized]  `tfsdk:"errors" json:"errors,computed"`
	Messages              customfield.List[jsontypes.Normalized]  `tfsdk:"messages" json:"messages,computed"`
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
