// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_connector_rules

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudConnectorRulesResultEnvelope struct {
	Result *[]*CloudConnectorRulesBodyModel `json:"result"`
}

type CloudConnectorRulesModel struct {
	ZoneID types.String                     `tfsdk:"zone_id" path:"zone_id,required"`
	Body   *[]*CloudConnectorRulesBodyModel `tfsdk:"body" json:"body,required"`
}

func (m CloudConnectorRulesModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Body)
}

func (m CloudConnectorRulesModel) MarshalJSONForUpdate(state CloudConnectorRulesModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m.Body, state.Body)
}

type CloudConnectorRulesBodyModel struct {
	ID          types.String                            `tfsdk:"id" json:"id,optional"`
	Description types.String                            `tfsdk:"description" json:"description,optional"`
	Enabled     types.Bool                              `tfsdk:"enabled" json:"enabled,optional"`
	Expression  types.String                            `tfsdk:"expression" json:"expression,optional"`
	Parameters  *CloudConnectorRulesBodyParametersModel `tfsdk:"parameters" json:"parameters,optional"`
	Provider    types.String                            `tfsdk:"provider" json:"provider,optional"`
}

type CloudConnectorRulesBodyParametersModel struct {
	Host types.String `tfsdk:"host" json:"host,optional"`
}
