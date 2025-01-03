// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersScriptsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkersScriptsResultDataSourceModel] `json:"result,computed"`
}

type WorkersScriptsDataSourceModel struct {
	AccountID types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                       `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[WorkersScriptsResultDataSourceModel] `tfsdk:"result"`
}

func (m *WorkersScriptsDataSourceModel) toListParams(_ context.Context) (params workers.ScriptListParams, diags diag.Diagnostics) {
	params = workers.ScriptListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type WorkersScriptsResultDataSourceModel struct {
	ID              types.String                                                             `tfsdk:"id" json:"id,computed"`
	CreatedOn       timetypes.RFC3339                                                        `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Etag            types.String                                                             `tfsdk:"etag" json:"etag,computed"`
	HasAssets       types.Bool                                                               `tfsdk:"has_assets" json:"has_assets,computed"`
	HasModules      types.Bool                                                               `tfsdk:"has_modules" json:"has_modules,computed"`
	Logpush         types.Bool                                                               `tfsdk:"logpush" json:"logpush,computed"`
	ModifiedOn      timetypes.RFC3339                                                        `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Placement       customfield.NestedObject[WorkersScriptsPlacementDataSourceModel]         `tfsdk:"placement" json:"placement,computed"`
	PlacementMode   types.String                                                             `tfsdk:"placement_mode" json:"placement_mode,computed"`
	PlacementStatus types.String                                                             `tfsdk:"placement_status" json:"placement_status,computed"`
	TailConsumers   customfield.NestedObjectList[WorkersScriptsTailConsumersDataSourceModel] `tfsdk:"tail_consumers" json:"tail_consumers,computed"`
	UsageModel      types.String                                                             `tfsdk:"usage_model" json:"usage_model,computed"`
}

type WorkersScriptsPlacementDataSourceModel struct {
	Mode   types.String `tfsdk:"mode" json:"mode,computed"`
	Status types.String `tfsdk:"status" json:"status,computed"`
}

type WorkersScriptsTailConsumersDataSourceModel struct {
	Service     types.String `tfsdk:"service" json:"service,computed"`
	Environment types.String `tfsdk:"environment" json:"environment,computed"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace,computed"`
}
