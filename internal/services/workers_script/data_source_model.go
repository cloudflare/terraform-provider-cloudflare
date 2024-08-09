// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersScriptResultListDataSourceEnvelope struct {
	Result *[]*WorkersScriptDataSourceModel `json:"result,computed"`
}

type WorkersScriptDataSourceModel struct {
	AccountID     types.String                                  `tfsdk:"account_id" path:"account_id"`
	ScriptName    types.String                                  `tfsdk:"script_name" path:"script_name"`
	CreatedOn     timetypes.RFC3339                             `tfsdk:"created_on" json:"created_on"`
	Etag          types.String                                  `tfsdk:"etag" json:"etag"`
	ID            types.String                                  `tfsdk:"id" json:"id"`
	Logpush       types.Bool                                    `tfsdk:"logpush" json:"logpush"`
	ModifiedOn    timetypes.RFC3339                             `tfsdk:"modified_on" json:"modified_on"`
	PlacementMode types.String                                  `tfsdk:"placement_mode" json:"placement_mode"`
	UsageModel    types.String                                  `tfsdk:"usage_model" json:"usage_model"`
	TailConsumers *[]*WorkersScriptTailConsumersDataSourceModel `tfsdk:"tail_consumers" json:"tail_consumers"`
	Filter        *WorkersScriptFindOneByDataSourceModel        `tfsdk:"filter"`
}

type WorkersScriptTailConsumersDataSourceModel struct {
	Service     types.String `tfsdk:"service" json:"service,computed"`
	Environment types.String `tfsdk:"environment" json:"environment"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace"`
}

type WorkersScriptFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
