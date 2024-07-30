// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RulesetsResultListDataSourceEnvelope struct {
	Result *[]*RulesetsResultDataSourceModel `json:"result,computed"`
}

type RulesetsDataSourceModel struct {
	AccountID types.String                      `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                      `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                       `tfsdk:"max_items"`
	Result    *[]*RulesetsResultDataSourceModel `tfsdk:"result"`
}

type RulesetsResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	Kind        types.String      `tfsdk:"kind" json:"kind,computed"`
	LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	Phase       types.String      `tfsdk:"phase" json:"phase,computed"`
	Version     types.String      `tfsdk:"version" json:"version,computed"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
}
