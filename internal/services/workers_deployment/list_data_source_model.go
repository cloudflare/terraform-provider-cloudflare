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
	Items customfield.NestedObjectList[WorkersDeploymentsResultDataSourceModel] `json:"items,computed"`
}

type WorkersDeploymentsDataSourceModel struct {
	AccountID  types.String                                                          `tfsdk:"account_id" path:"account_id,required"`
	ScriptName types.String                                                          `tfsdk:"script_name" path:"script_name,required"`
	Since      timetypes.RFC3339                                                     `tfsdk:"since" query:"since,optional" format:"date-time"`
	Until      timetypes.RFC3339                                                     `tfsdk:"until" query:"until,optional" format:"date-time"`
	MaxItems   types.Int64                                                           `tfsdk:"max_items"`
	Result     customfield.NestedObjectList[WorkersDeploymentsResultDataSourceModel] `tfsdk:"result"`
}

func (m *WorkersDeploymentsDataSourceModel) toListParams(_ context.Context) (params workers.ScriptDeploymentListParams, diags diag.Diagnostics) {
	mSince, errs := m.Since.ValueRFC3339Time()
	diags.Append(errs...)
	mUntil, errs := m.Until.ValueRFC3339Time()
	diags.Append(errs...)

	params = workers.ScriptDeploymentListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Since.IsNull() {
		params.Since = cloudflare.F(mSince)
	}
	if !m.Until.IsNull() {
		params.Until = cloudflare.F(mUntil)
	}

	return
}

type WorkersDeploymentsResultDataSourceModel struct {
	Deployments customfield.NestedObjectList[WorkersDeploymentsDeploymentsDataSourceModel] `tfsdk:"deployments" json:"deployments,computed"`
}

type WorkersDeploymentsDeploymentsDataSourceModel struct {
	ID          types.String                                                                       `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339                                                                  `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Source      types.String                                                                       `tfsdk:"source" json:"source,computed"`
	Strategy    types.String                                                                       `tfsdk:"strategy" json:"strategy,computed"`
	Versions    customfield.NestedObjectList[WorkersDeploymentsDeploymentsVersionsDataSourceModel] `tfsdk:"versions" json:"versions,computed"`
	Annotations customfield.NestedObject[WorkersDeploymentsDeploymentsAnnotationsDataSourceModel]  `tfsdk:"annotations" json:"annotations,computed"`
	AuthorEmail types.String                                                                       `tfsdk:"author_email" json:"author_email,computed"`
}

type WorkersDeploymentsDeploymentsVersionsDataSourceModel struct {
	Percentage types.Float64 `tfsdk:"percentage" json:"percentage,computed"`
	VersionID  types.String  `tfsdk:"version_id" json:"version_id,computed"`
}

type WorkersDeploymentsDeploymentsAnnotationsDataSourceModel struct {
	WorkersMessage     types.String `tfsdk:"workers_message" json:"workers/message,computed"`
	WorkersTriggeredBy types.String `tfsdk:"workers_triggered_by" json:"workers/triggered_by,computed"`
}
