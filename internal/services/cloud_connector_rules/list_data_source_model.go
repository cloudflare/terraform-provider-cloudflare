// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_connector_rules

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/cloud_connector"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudConnectorRulesListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CloudConnectorRulesListResultDataSourceModel] `json:"result,computed"`
}

type CloudConnectorRulesListDataSourceModel struct {
	ZoneID   types.String                                                               `tfsdk:"zone_id" path:"zone_id,required"`
	MaxItems types.Int64                                                                `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[CloudConnectorRulesListResultDataSourceModel] `tfsdk:"result"`
}

func (m *CloudConnectorRulesListDataSourceModel) toListParams(_ context.Context) (params cloud_connector.RuleListParams, diags diag.Diagnostics) {
	params = cloud_connector.RuleListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type CloudConnectorRulesListResultDataSourceModel struct {
	ID            types.String                                                               `tfsdk:"id" json:"id,computed"`
	Description   types.String                                                               `tfsdk:"description" json:"description,computed"`
	Enabled       types.Bool                                                                 `tfsdk:"enabled" json:"enabled,computed"`
	Expression    types.String                                                               `tfsdk:"expression" json:"expression,computed"`
	Parameters    customfield.NestedObject[CloudConnectorRulesListParametersDataSourceModel] `tfsdk:"parameters" json:"parameters,computed"`
	CloudProvider types.String                                                               `tfsdk:"cloud_provider" json:"provider,computed"`
}

type CloudConnectorRulesListParametersDataSourceModel struct {
	Host types.String `tfsdk:"host" json:"host,computed"`
}
