// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_deployment

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersDeploymentResultEnvelope struct {
	Result WorkersDeploymentModel `json:"result"`
}

type WorkersDeploymentModel struct {
	ID          types.String                                                    `tfsdk:"id" json:"id,computed"`
	AccountID   types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	ScriptName  types.String                                                    `tfsdk:"script_name" path:"script_name,required"`
	Strategy    types.String                                                    `tfsdk:"strategy" json:"strategy,required"`
	Versions    *[]*WorkersDeploymentVersionsModel                              `tfsdk:"versions" json:"versions,required"`
	Annotations customfield.NestedObject[WorkersDeploymentAnnotationsModel]     `tfsdk:"annotations" json:"annotations,computed_optional"`
	AuthorEmail types.String                                                    `tfsdk:"author_email" json:"author_email,computed"`
	CreatedOn   types.String                                                    `tfsdk:"created_on" json:"created_on,computed"`
	Source      types.String                                                    `tfsdk:"source" json:"source,computed"`
	Deployments customfield.NestedObjectList[WorkersDeploymentDeploymentsModel] `tfsdk:"deployments" json:"deployments,computed"`
}

func (m WorkersDeploymentModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WorkersDeploymentModel) MarshalJSONForUpdate(state WorkersDeploymentModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type WorkersDeploymentVersionsModel struct {
	Percentage types.Float64 `tfsdk:"percentage" json:"percentage,required"`
	VersionID  types.String  `tfsdk:"version_id" json:"version_id,required"`
}

type WorkersDeploymentAnnotationsModel struct {
	WorkersMessage types.String `tfsdk:"workers_message" json:"workers/message,optional"`
}

type WorkersDeploymentDeploymentsModel struct {
	Strategy    types.String                                                            `tfsdk:"strategy" json:"strategy,computed"`
	Versions    customfield.NestedObjectList[WorkersDeploymentDeploymentsVersionsModel] `tfsdk:"versions" json:"versions,computed"`
	ID          types.String                                                            `tfsdk:"id" json:"id,computed"`
	Annotations customfield.NestedObject[WorkersDeploymentDeploymentsAnnotationsModel]  `tfsdk:"annotations" json:"annotations,computed"`
	AuthorEmail types.String                                                            `tfsdk:"author_email" json:"author_email,computed"`
	CreatedOn   types.String                                                            `tfsdk:"created_on" json:"created_on,computed"`
	Source      types.String                                                            `tfsdk:"source" json:"source,computed"`
}

type WorkersDeploymentDeploymentsVersionsModel struct {
	Percentage types.Float64 `tfsdk:"percentage" json:"percentage,computed"`
	VersionID  types.String  `tfsdk:"version_id" json:"version_id,computed"`
}

type WorkersDeploymentDeploymentsAnnotationsModel struct {
	WorkersMessage types.String `tfsdk:"workers_message" json:"workers/message,computed"`
}
