// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_connector_rules

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudConnectorRulesResultEnvelope struct {
	Result customfield.NestedObjectList[CloudConnectorRulesRulesModel] `json:"result"`
}

type CloudConnectorRulesModel struct {
	ID          types.String                                                 `tfsdk:"id" json:"id,computed"`
	ZoneID      types.String                                                 `tfsdk:"zone_id" path:"zone_id,required"`
	Rules       customfield.NestedObjectList[CloudConnectorRulesRulesModel]  `tfsdk:"rules" json:"rules,computed_optional"`
	Description types.String                                                 `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool                                                   `tfsdk:"enabled" json:"enabled,computed"`
	Expression  types.String                                                 `tfsdk:"expression" json:"expression,computed"`
	Provider    types.String                                                 `tfsdk:"provider" json:"provider,computed"`
	Parameters  customfield.NestedObject[CloudConnectorRulesParametersModel] `tfsdk:"parameters" json:"parameters,computed"`
}

func (m CloudConnectorRulesModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Rules)
}

func (m CloudConnectorRulesModel) MarshalJSONForUpdate(state CloudConnectorRulesModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m.Rules, state.Rules)
}

type CloudConnectorRulesRulesModel struct {
	ID          types.String                                                      `tfsdk:"id" json:"id,optional"`
	Description types.String                                                      `tfsdk:"description" json:"description,optional"`
	Enabled     types.Bool                                                        `tfsdk:"enabled" json:"enabled,optional"`
	Expression  types.String                                                      `tfsdk:"expression" json:"expression,optional"`
	Parameters  customfield.NestedObject[CloudConnectorRulesRulesParametersModel] `tfsdk:"parameters" json:"parameters,computed_optional"`
	Provider    types.String                                                      `tfsdk:"provider" json:"provider,optional"`
}

type CloudConnectorRulesRulesParametersModel struct {
	Host types.String `tfsdk:"host" json:"host,optional"`
}

type CloudConnectorRulesParametersModel struct {
	Host types.String `tfsdk:"host" json:"host,computed"`
}
