// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesProjectsResultListDataSourceEnvelope struct {
	Result *[]*PagesProjectsResultDataSourceModel `json:"result,computed"`
}

type PagesProjectsDataSourceModel struct {
	AccountID types.String                           `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                            `tfsdk:"max_items"`
	Result    *[]*PagesProjectsResultDataSourceModel `tfsdk:"result"`
}

type PagesProjectsResultDataSourceModel struct {
	ID                types.String                                                            `tfsdk:"id" json:"id,computed"`
	Aliases           *[]jsontypes.Normalized                                                 `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig       jsontypes.Normalized                                                    `tfsdk:"build_config" json:"build_config,computed"`
	CreatedOn         timetypes.RFC3339                                                       `tfsdk:"created_on" json:"created_on,computed"`
	DeploymentTrigger customfield.NestedObject[PagesProjectsDeploymentTriggerDataSourceModel] `tfsdk:"deployment_trigger" json:"deployment_trigger,computed"`
	EnvVars           jsontypes.Normalized                                                    `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment       types.String                                                            `tfsdk:"environment" json:"environment,computed"`
	IsSkipped         types.Bool                                                              `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage       jsontypes.Normalized                                                    `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn        timetypes.RFC3339                                                       `tfsdk:"modified_on" json:"modified_on,computed"`
	ProjectID         types.String                                                            `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName       types.String                                                            `tfsdk:"project_name" json:"project_name,computed"`
	ShortID           types.String                                                            `tfsdk:"short_id" json:"short_id,computed"`
	Source            jsontypes.Normalized                                                    `tfsdk:"source" json:"source,computed"`
	Stages            *[]*PagesProjectsStagesDataSourceModel                                  `tfsdk:"stages" json:"stages,computed"`
	URL               types.String                                                            `tfsdk:"url" json:"url,computed"`
}

type PagesProjectsDeploymentTriggerDataSourceModel struct {
	Metadata *PagesProjectsDeploymentTriggerMetadataDataSourceModel `tfsdk:"metadata" json:"metadata"`
	Type     types.String                                           `tfsdk:"type" json:"type,computed"`
}

type PagesProjectsDeploymentTriggerMetadataDataSourceModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectsStagesDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}
