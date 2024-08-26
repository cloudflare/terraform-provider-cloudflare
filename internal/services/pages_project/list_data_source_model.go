// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/pages"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagesProjectsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PagesProjectsResultDataSourceModel] `json:"result,computed"`
}

type PagesProjectsDataSourceModel struct {
	AccountID types.String                                                     `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                                      `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[PagesProjectsResultDataSourceModel] `tfsdk:"result"`
}

func (m *PagesProjectsDataSourceModel) toListParams() (params pages.ProjectListParams, diags diag.Diagnostics) {
	params = pages.ProjectListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type PagesProjectsResultDataSourceModel struct {
	ID                types.String                                                            `tfsdk:"id" json:"id,computed"`
	Aliases           types.List                                                              `tfsdk:"aliases" json:"aliases,computed"`
	BuildConfig       *PagesProjectsBuildConfigDataSourceModel                                `tfsdk:"build_config" json:"build_config"`
	CreatedOn         timetypes.RFC3339                                                       `tfsdk:"created_on" json:"created_on,computed"`
	DeploymentTrigger customfield.NestedObject[PagesProjectsDeploymentTriggerDataSourceModel] `tfsdk:"deployment_trigger" json:"deployment_trigger,computed"`
	EnvVars           map[string]jsontypes.Normalized                                         `tfsdk:"env_vars" json:"env_vars,computed"`
	Environment       types.String                                                            `tfsdk:"environment" json:"environment,computed"`
	IsSkipped         types.Bool                                                              `tfsdk:"is_skipped" json:"is_skipped,computed"`
	LatestStage       customfield.NestedObject[PagesProjectsLatestStageDataSourceModel]       `tfsdk:"latest_stage" json:"latest_stage,computed"`
	ModifiedOn        timetypes.RFC3339                                                       `tfsdk:"modified_on" json:"modified_on,computed"`
	ProjectID         types.String                                                            `tfsdk:"project_id" json:"project_id,computed"`
	ProjectName       types.String                                                            `tfsdk:"project_name" json:"project_name,computed"`
	ShortID           types.String                                                            `tfsdk:"short_id" json:"short_id,computed"`
	Source            customfield.NestedObject[PagesProjectsSourceDataSourceModel]            `tfsdk:"source" json:"source,computed"`
	Stages            customfield.NestedObjectList[PagesProjectsStagesDataSourceModel]        `tfsdk:"stages" json:"stages,computed"`
	URL               types.String                                                            `tfsdk:"url" json:"url,computed"`
}

type PagesProjectsBuildConfigDataSourceModel struct {
	BuildCaching      types.Bool   `tfsdk:"build_caching" json:"build_caching"`
	BuildCommand      types.String `tfsdk:"build_command" json:"build_command"`
	DestinationDir    types.String `tfsdk:"destination_dir" json:"destination_dir"`
	RootDir           types.String `tfsdk:"root_dir" json:"root_dir"`
	WebAnalyticsTag   types.String `tfsdk:"web_analytics_tag" json:"web_analytics_tag"`
	WebAnalyticsToken types.String `tfsdk:"web_analytics_token" json:"web_analytics_token"`
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

type PagesProjectsLatestStageDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}

type PagesProjectsSourceDataSourceModel struct {
	Config *PagesProjectsSourceConfigDataSourceModel `tfsdk:"config" json:"config"`
	Type   types.String                              `tfsdk:"type" json:"type"`
}

type PagesProjectsSourceConfigDataSourceModel struct {
	DeploymentsEnabled           types.Bool      `tfsdk:"deployments_enabled" json:"deployments_enabled"`
	Owner                        types.String    `tfsdk:"owner" json:"owner"`
	PathExcludes                 *[]types.String `tfsdk:"path_excludes" json:"path_excludes"`
	PathIncludes                 *[]types.String `tfsdk:"path_includes" json:"path_includes"`
	PrCommentsEnabled            types.Bool      `tfsdk:"pr_comments_enabled" json:"pr_comments_enabled"`
	PreviewBranchExcludes        *[]types.String `tfsdk:"preview_branch_excludes" json:"preview_branch_excludes"`
	PreviewBranchIncludes        *[]types.String `tfsdk:"preview_branch_includes" json:"preview_branch_includes"`
	PreviewDeploymentSetting     types.String    `tfsdk:"preview_deployment_setting" json:"preview_deployment_setting"`
	ProductionBranch             types.String    `tfsdk:"production_branch" json:"production_branch"`
	ProductionDeploymentsEnabled types.Bool      `tfsdk:"production_deployments_enabled" json:"production_deployments_enabled"`
	RepoName                     types.String    `tfsdk:"repo_name" json:"repo_name"`
}

type PagesProjectsStagesDataSourceModel struct {
	EndedOn   timetypes.RFC3339 `tfsdk:"ended_on" json:"ended_on,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	StartedOn timetypes.RFC3339 `tfsdk:"started_on" json:"started_on,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
}
