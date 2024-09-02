package cloud_connector_rules

import "github.com/hashicorp/terraform-plugin-framework/types"

type CloudConnectorRules struct {
	ZoneID types.String         `tfsdk:"zone_id"`
	Rules  []CloudConnectorRule `tfsdk:"rules"`
}

type CloudConnectorRule struct {
	Enabled     types.Bool                   `tfsdk:"enabled"`
	Expression  types.String                 `tfsdk:"expression"`
	Provider    types.String                 `tfsdk:"provider"`
	Description types.String                 `tfsdk:"description"`
	Parameters  CloudConnectorRuleParameters `tfsdk:"parameters"`
}

type CloudConnectorRuleParameters struct {
	Host types.String `tfsdk:"host"`
}
