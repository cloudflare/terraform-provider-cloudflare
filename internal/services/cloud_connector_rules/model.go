// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_connector_rules

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudConnectorRulesResultEnvelope struct {
	Result *[]*CloudConnectorRulesRulesModel `json:"result"`
}

type CloudConnectorRulesModel struct {
	ID     types.String                      `tfsdk:"id" json:"-,computed"`
	ZoneID types.String                      `tfsdk:"zone_id" path:"zone_id,required"`
	Rules  *[]*CloudConnectorRulesRulesModel `tfsdk:"rules" json:"rules,optional"`
}

func (m CloudConnectorRulesModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Rules)
}

func (m CloudConnectorRulesModel) MarshalJSONForUpdate(state CloudConnectorRulesModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m.Rules, state.Rules)
}

type CloudConnectorRulesRulesModel struct {
	ID          types.String                             `tfsdk:"id" json:"id,computed"`
	Description types.String                             `tfsdk:"description" json:"description,optional"`
	Enabled     types.Bool                               `tfsdk:"enabled" json:"enabled,optional"`
	Expression  types.String                             `tfsdk:"expression" json:"expression,optional"`
	Parameters  *CloudConnectorRulesRulesParametersModel `tfsdk:"parameters" json:"parameters,optional"`
	Provider    types.String                             `tfsdk:"provider" json:"provider,optional"`
}

type CloudConnectorRulesRulesParametersModel struct {
	Host types.String `tfsdk:"host" json:"host,optional"`
}
