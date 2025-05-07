// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_connector_rules

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/cloud_connector"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudConnectorRulesResultDataSourceEnvelope struct {
	Result CloudConnectorRulesDataSourceModel `json:"result,computed"`
}

type CloudConnectorRulesDataSourceModel struct {
	ZoneID      types.String                                                           `tfsdk:"zone_id" path:"zone_id,required"`
	Description types.String                                                           `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool                                                             `tfsdk:"enabled" json:"enabled,computed"`
	Expression  types.String                                                           `tfsdk:"expression" json:"expression,computed"`
	ID          types.String                                                           `tfsdk:"id" json:"id,computed"`
	Provider    types.String                                                           `tfsdk:"provider" json:"provider,computed"`
	Parameters  customfield.NestedObject[CloudConnectorRulesParametersDataSourceModel] `tfsdk:"parameters" json:"parameters,computed"`
}

func (m *CloudConnectorRulesDataSourceModel) toReadParams(_ context.Context) (params cloud_connector.RuleListParams, diags diag.Diagnostics) {
	params = cloud_connector.RuleListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type CloudConnectorRulesParametersDataSourceModel struct {
	Host types.String `tfsdk:"host" json:"host,computed"`
}
