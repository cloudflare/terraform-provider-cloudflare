// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_deployment

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/workers"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersDeploymentResultDataSourceEnvelope struct {
Result WorkersDeploymentDataSourceModel `json:"result,computed"`
}

type WorkersDeploymentDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
ScriptName types.String `tfsdk:"script_name" path:"script_name,required"`
Deployments customfield.NestedObjectList[WorkersDeploymentDeploymentsDataSourceModel] `tfsdk:"deployments" json:"deployments,computed"`
}

func (m *WorkersDeploymentDataSourceModel) toReadParams(_ context.Context) (params workers.ScriptDeploymentGetParams, diags diag.Diagnostics) {
  params = workers.ScriptDeploymentGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type WorkersDeploymentDeploymentsDataSourceModel struct {
Strategy types.String `tfsdk:"strategy" json:"strategy,computed"`
Versions customfield.NestedObjectList[WorkersDeploymentDeploymentsVersionsDataSourceModel] `tfsdk:"versions" json:"versions,computed"`
ID types.String `tfsdk:"id" json:"id,computed"`
Annotations customfield.NestedObject[WorkersDeploymentDeploymentsAnnotationsDataSourceModel] `tfsdk:"annotations" json:"annotations,computed"`
AuthorEmail types.String `tfsdk:"author_email" json:"author_email,computed"`
CreatedOn types.String `tfsdk:"created_on" json:"created_on,computed"`
Source types.String `tfsdk:"source" json:"source,computed"`
}

type WorkersDeploymentDeploymentsVersionsDataSourceModel struct {
Percentage types.Float64 `tfsdk:"percentage" json:"percentage,computed"`
VersionID types.String `tfsdk:"version_id" json:"version_id,computed"`
}

type WorkersDeploymentDeploymentsAnnotationsDataSourceModel struct {
WorkersMessage types.String `tfsdk:"workers_message" json:"workers/message,computed"`
}
