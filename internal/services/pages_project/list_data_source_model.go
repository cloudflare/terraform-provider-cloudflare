// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesProjectsResultListDataSourceEnvelope struct {
	Result *[]*PagesProjectsItemsDataSourceModel `json:"result,computed"`
}

type PagesProjectsDataSourceModel struct {
	AccountID types.String                          `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                           `tfsdk:"max_items"`
	Items     *[]*PagesProjectsItemsDataSourceModel `tfsdk:"items"`
}

type PagesProjectsItemsDataSourceModel struct {
	ID          types.String                                `tfsdk:"id" json:"id,computed"`
	Aliases     *[]types.String                             `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig types.String                                `tfsdk:"build_config" json:"build_config,computed"`
	CreatedOn   types.String                                `tfsdk:"created_on" json:"created_on,computed"`
	EnvVars     types.String                                `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment types.String                                `tfsdk:"environment" json:"environment,computed"`
	IsSkipped   types.Bool                                  `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage types.String                                `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn  types.String                                `tfsdk:"modified_on" json:"modified_on,computed"`
	ProjectID   types.String                                `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName types.String                                `tfsdk:"project_name" json:"project_name,computed"`
	ShortID     types.String                                `tfsdk:"short_id" json:"short_id,computed"`
	Source      types.String                                `tfsdk:"source" json:"source,computed"`
	Stages      *[]*PagesProjectsItemsStagesDataSourceModel `tfsdk:"stages" json:"stages,computed"`
	URL         types.String                                `tfsdk:"url" json:"url,computed"`
}

type PagesProjectsItemsDeploymentTriggerDataSourceModel struct {
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type PagesProjectsItemsDeploymentTriggerMetadataDataSourceModel struct {
	Branch        types.String `tfsdk:"branch" json:"branch,computed"`
	CommitHash    types.String `tfsdk:"commit_hash" json:"commit_hash,computed"`
	CommitMessage types.String `tfsdk:"commit_message" json:"commit_message,computed"`
}

type PagesProjectsItemsStagesDataSourceModel struct {
	EndedOn   types.String `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
	StartedOn types.String `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String `tfsdk:"status" json:"status,computed"`
}
