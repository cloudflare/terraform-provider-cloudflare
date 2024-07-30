// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_script

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkerScriptResultListDataSourceEnvelope struct {
	Result *[]*WorkerScriptDataSourceModel `json:"result,computed"`
}

type WorkerScriptDataSourceModel struct {
	AccountID     types.String                                 `tfsdk:"account_id" path:"account_id"`
	ScriptName    types.String                                 `tfsdk:"script_name" path:"script_name"`
	ID            types.String                                 `tfsdk:"id" json:"id"`
	CreatedOn     timetypes.RFC3339                            `tfsdk:"created_on" json:"created_on"`
	Etag          types.String                                 `tfsdk:"etag" json:"etag"`
	Logpush       types.Bool                                   `tfsdk:"logpush" json:"logpush"`
	ModifiedOn    timetypes.RFC3339                            `tfsdk:"modified_on" json:"modified_on"`
	PlacementMode types.String                                 `tfsdk:"placement_mode" json:"placement_mode"`
	TailConsumers *[]*WorkerScriptTailConsumersDataSourceModel `tfsdk:"tail_consumers" json:"tail_consumers"`
	UsageModel    types.String                                 `tfsdk:"usage_model" json:"usage_model"`
	Filter        *WorkerScriptFindOneByDataSourceModel        `tfsdk:"filter"`
}

type WorkerScriptTailConsumersDataSourceModel struct {
	Service     types.String `tfsdk:"service" json:"service,computed"`
	Environment types.String `tfsdk:"environment" json:"environment"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace"`
}

type WorkerScriptFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
