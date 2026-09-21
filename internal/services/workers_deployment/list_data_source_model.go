// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_deployment

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersDeploymentsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[WorkersDeploymentsResultDataSourceModel] `json:"deployments,computed"`
}

type WorkersDeploymentsDataSourceModel struct {
	AccountID  types.String                                                          `tfsdk:"account_id" path:"account_id,required"`
	ScriptName types.String                                                          `tfsdk:"script_name" path:"script_name,required"`
	MaxItems   types.Int64                                                           `tfsdk:"max_items"`
	Result     customfield.NestedObjectList[WorkersDeploymentsResultDataSourceModel] `tfsdk:"result"`
}

func (m *WorkersDeploymentsDataSourceModel) toListParams(_ context.Context) (params workers.ScriptDeploymentListParams, diags diag.Diagnostics) {
	params = workers.ScriptDeploymentListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type WorkersDeploymentsResultDataSourceModel struct {
	ID          types.String                                                            `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339                                                       `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Source      types.String                                                            `tfsdk:"source" json:"source,computed"`
	Strategy    types.String                                                            `tfsdk:"strategy" json:"strategy,computed"`
	Versions    customfield.NestedObjectList[WorkersDeploymentsVersionsDataSourceModel] `tfsdk:"versions" json:"versions,computed"`
	Annotations customfield.NestedObject[WorkersDeploymentsAnnotationsDataSourceModel]  `tfsdk:"annotations" json:"annotations,computed"`
	AuthorEmail types.String                                                            `tfsdk:"author_email" json:"author_email,computed"`
}

type WorkersDeploymentsVersionsDataSourceModel struct {
	Percentage types.Float64 `tfsdk:"percentage" json:"percentage,computed"`
	VersionID  types.String  `tfsdk:"version_id" json:"version_id,computed"`
}

type WorkersDeploymentsAnnotationsDataSourceModel struct {
	WorkersMessage     types.String `tfsdk:"workers_message" json:"workers/message,computed"`
	WorkersTriggeredBy types.String `tfsdk:"workers_triggered_by" json:"workers/triggered_by,computed"`
}
