// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_deployment

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersDeploymentResultDataSourceEnvelope struct {
	Result WorkersDeploymentDataSourceModel `json:"result,computed"`
}

type WorkersDeploymentDataSourceModel struct {
	ID           types.String                                                           `tfsdk:"id" path:"deployment_id,computed"`
	DeploymentID types.String                                                           `tfsdk:"deployment_id" path:"deployment_id,required"`
	AccountID    types.String                                                           `tfsdk:"account_id" path:"account_id,required"`
	ScriptName   types.String                                                           `tfsdk:"script_name" path:"script_name,required"`
	AuthorEmail  types.String                                                           `tfsdk:"author_email" json:"author_email,computed"`
	CreatedOn    timetypes.RFC3339                                                      `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Source       types.String                                                           `tfsdk:"source" json:"source,computed"`
	Strategy     types.String                                                           `tfsdk:"strategy" json:"strategy,computed"`
	Annotations  customfield.NestedObject[WorkersDeploymentAnnotationsDataSourceModel]  `tfsdk:"annotations" json:"annotations,computed"`
	Versions     customfield.NestedObjectList[WorkersDeploymentVersionsDataSourceModel] `tfsdk:"versions" json:"versions,computed"`
}

func (m *WorkersDeploymentDataSourceModel) toReadParams(_ context.Context) (params workers.ScriptDeploymentGetParams, diags diag.Diagnostics) {
	params = workers.ScriptDeploymentGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type WorkersDeploymentAnnotationsDataSourceModel struct {
	WorkersMessage     types.String `tfsdk:"workers_message" json:"workers/message,computed"`
	WorkersTriggeredBy types.String `tfsdk:"workers_triggered_by" json:"workers/triggered_by,computed"`
}

type WorkersDeploymentVersionsDataSourceModel struct {
	Percentage types.Float64 `tfsdk:"percentage" json:"percentage,computed"`
	VersionID  types.String  `tfsdk:"version_id" json:"version_id,computed"`
}
