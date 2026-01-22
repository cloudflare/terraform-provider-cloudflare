// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/pages"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*PagesProjectResource)(nil)
var _ resource.ResourceWithModifyPlan = (*PagesProjectResource)(nil)
var _ resource.ResourceWithImportState = (*PagesProjectResource)(nil)

func NewResource() resource.Resource {
	return &PagesProjectResource{}
}

// PagesProjectResource defines the resource implementation.
type PagesProjectResource struct {
	client *cloudflare.Client
}

func (r *PagesProjectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pages_project"
}

func (r *PagesProjectResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *PagesProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PagesProjectModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save plan deployment_configs to preserve secret env vars later
	planDeploymentConfigs := data.DeploymentConfigs

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := PagesProjectResultEnvelope{*data}
	_, err = r.client.Pages.Projects.New(
		ctx,
		pages.ProjectNewParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.Name

	// Normalize empty build_config and empty pointer slices to null before preserving env vars
	var normDiags diag.Diagnostics
	data, normDiags = NormalizeDeploymentConfigs(ctx, data)
	if normDiags.HasError() {
		resp.Diagnostics.Append(normDiags...)
		return
	}

	updatedDeploymentConfigs, diags := PreserveSecretEnvVars(ctx, planDeploymentConfigs, data.DeploymentConfigs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	data.DeploymentConfigs = updatedDeploymentConfigs

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PagesProjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PagesProjectModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *PagesProjectModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	planDeploymentConfigs := data.DeploymentConfigs

	// Convert nil binding maps to empty maps so the API deletes them.
	data, diags := PrepareForUpdate(ctx, data, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := PagesProjectResultEnvelope{*data}
	_, err = r.client.Pages.Projects.Edit(
		ctx,
		data.Name.ValueString(),
		pages.ProjectEditParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.Name

	// Normalize empty pointer slices and structs to null before preserving env vars
	var normDiags diag.Diagnostics
	data, normDiags = NormalizeDeploymentConfigs(ctx, data)
	if normDiags.HasError() {
		resp.Diagnostics.Append(normDiags...)
		return
	}

	updatedDeploymentConfigs, diags := PreserveSecretEnvVars(ctx, planDeploymentConfigs, data.DeploymentConfigs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	data.DeploymentConfigs = updatedDeploymentConfigs

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PagesProjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PagesProjectModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve state values for comparison
	stateData := *data
	stateDeploymentConfigs := data.DeploymentConfigs

	res := new(http.Response)
	env := PagesProjectResultEnvelope{*data}
	_, err := r.client.Pages.Projects.Get(
		ctx,
		data.Name.ValueString(),
		pages.ProjectGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.Name

	// Preserve build_config from state only if API didn't return it
	// This handles the case where the API inconsistently returns build_config
	if data.BuildConfig == nil && stateData.BuildConfig != nil {
		data.BuildConfig = stateData.BuildConfig
	}

	// Normalize empty pointer slices to null before preserving env vars
	var normDiags diag.Diagnostics
	data, normDiags = NormalizeDeploymentConfigs(ctx, data)
	if normDiags.HasError() {
		resp.Diagnostics.Append(normDiags...)
		return
	}

	updatedDeploymentConfigs, diags := PreserveSecretEnvVars(ctx, stateDeploymentConfigs, data.DeploymentConfigs)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	data.DeploymentConfigs = updatedDeploymentConfigs

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PagesProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PagesProjectModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Pages.Projects.Delete(
		ctx,
		data.Name.ValueString(),
		pages.ProjectDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.Name

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PagesProjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(PagesProjectModel)

	path_account_id := ""
	path_project_name := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<project_name>",
		&path_account_id,
		&path_project_name,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.Name = types.StringValue(path_project_name)

	res := new(http.Response)
	env := PagesProjectResultEnvelope{*data}
	_, err := r.client.Pages.Projects.Get(
		ctx,
		path_project_name,
		pages.ProjectGetParams{
			AccountID: cloudflare.F(path_account_id),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.Name

	// Normalize deployment_configs to match API behavior - empty pointer slices should be null
	var normDiags diag.Diagnostics
	data, normDiags = NormalizeDeploymentConfigs(ctx, data)
	if normDiags.HasError() {
		resp.Diagnostics.Append(normDiags...)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PagesProjectResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// If we're deleting the resource, no need to modify the plan
	if req.Plan.Raw.IsNull() {
		return
	}

	// If the state is null (resource doesn't exist yet), no need to modify the plan
	if req.State.Raw.IsNull() {
		return
	}

	var plan, state *PagesProjectModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve computed-only and computed_optional fields from state during refresh
	// These fields are marked as Computed or ComputedOptional in the schema and should
	// not change during a refresh unless the user explicitly updates the configuration

	// Preserve canonical_deployment if it's unknown in the plan but present in state
	if plan.CanonicalDeployment.IsUnknown() && !state.CanonicalDeployment.IsNull() {
		plan.CanonicalDeployment = state.CanonicalDeployment
	}

	// Preserve latest_deployment if it's unknown in the plan but present in state
	if plan.LatestDeployment.IsUnknown() && !state.LatestDeployment.IsNull() {
		plan.LatestDeployment = state.LatestDeployment
	}

	// Preserve deployment_configs if it's unknown or null in the plan but present in state.
	// This prevents drift when user omits deployment_configs (issue #5928).
	if (plan.DeploymentConfigs.IsUnknown() || plan.DeploymentConfigs.IsNull()) && !state.DeploymentConfigs.IsNull() {
		plan.DeploymentConfigs = state.DeploymentConfigs
	}

	// Preserve build_config if it's null in the plan but present in state.
	// This prevents drift when user omits build_config (issue #5928).
	if plan.BuildConfig == nil && state.BuildConfig != nil {
		plan.BuildConfig = state.BuildConfig
	}

	// Preserve other computed fields
	if plan.CreatedOn.IsUnknown() && !state.CreatedOn.IsNull() {
		plan.CreatedOn = state.CreatedOn
	}

	if plan.Framework.IsUnknown() && !state.Framework.IsNull() {
		plan.Framework = state.Framework
	}

	if plan.FrameworkVersion.IsUnknown() && !state.FrameworkVersion.IsNull() {
		plan.FrameworkVersion = state.FrameworkVersion
	}

	if plan.PreviewScriptName.IsUnknown() && !state.PreviewScriptName.IsNull() {
		plan.PreviewScriptName = state.PreviewScriptName
	}

	if plan.ProductionScriptName.IsUnknown() && !state.ProductionScriptName.IsNull() {
		plan.ProductionScriptName = state.ProductionScriptName
	}

	if plan.Subdomain.IsUnknown() && !state.Subdomain.IsNull() {
		plan.Subdomain = state.Subdomain
	}

	if plan.UsesFunctions.IsUnknown() && !state.UsesFunctions.IsNull() {
		plan.UsesFunctions = state.UsesFunctions
	}

	if plan.Domains.IsUnknown() && !state.Domains.IsNull() {
		plan.Domains = state.Domains
	}

	resp.Diagnostics.Append(resp.Plan.Set(ctx, plan)...)
}

// mergeBuildConfigFromState merges unknown or null fields from state into plan.
// This handles the case where a user specifies build_config with only some fields,
// and we need to preserve the computed values from state for the unspecified fields.
func mergeBuildConfigFromState(plan, state *PagesProjectBuildConfigModel) {
	if plan == nil || state == nil {
		return
	}

	// For each field, if the plan value is unknown or null, use the state value
	if plan.BuildCaching.IsUnknown() || plan.BuildCaching.IsNull() {
		plan.BuildCaching = state.BuildCaching
	}
	if plan.BuildCommand.IsUnknown() || plan.BuildCommand.IsNull() {
		plan.BuildCommand = state.BuildCommand
	}
	if plan.DestinationDir.IsUnknown() || plan.DestinationDir.IsNull() {
		plan.DestinationDir = state.DestinationDir
	}
	if plan.RootDir.IsUnknown() || plan.RootDir.IsNull() {
		plan.RootDir = state.RootDir
	}
	if plan.WebAnalyticsTag.IsUnknown() || plan.WebAnalyticsTag.IsNull() {
		plan.WebAnalyticsTag = state.WebAnalyticsTag
	}
	if plan.WebAnalyticsToken.IsUnknown() || plan.WebAnalyticsToken.IsNull() {
		plan.WebAnalyticsToken = state.WebAnalyticsToken
	}
}

// NormalizeDeploymentConfigs normalizes empty pointer slices to null
// to match API behavior and prevent drift during import/refresh
func NormalizeDeploymentConfigs(ctx context.Context, data *PagesProjectModel) (*PagesProjectModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if data == nil {
		return data, diags
	}

	// Normalize build_config to null if all fields are empty/null
	// This handles the case where the API returns empty strings for build_config fields,
	// which would otherwise cause "planned value for a non-computed attribute" errors
	// when the user doesn't specify build_config in their configuration.
	if data.BuildConfig != nil {
		bc := data.BuildConfig
		allFieldsEmpty := true
		// For bool fields, check if not null and not unknown
		if !bc.BuildCaching.IsNull() && !bc.BuildCaching.IsUnknown() {
			allFieldsEmpty = false
		}
		// For string fields, check if not null, not unknown, AND not empty string
		// The API often returns empty strings "" which are different from null
		if !bc.BuildCommand.IsNull() && !bc.BuildCommand.IsUnknown() && bc.BuildCommand.ValueString() != "" {
			allFieldsEmpty = false
		}
		if !bc.DestinationDir.IsNull() && !bc.DestinationDir.IsUnknown() && bc.DestinationDir.ValueString() != "" {
			allFieldsEmpty = false
		}
		if !bc.RootDir.IsNull() && !bc.RootDir.IsUnknown() && bc.RootDir.ValueString() != "" {
			allFieldsEmpty = false
		}
		if !bc.WebAnalyticsTag.IsNull() && !bc.WebAnalyticsTag.IsUnknown() && bc.WebAnalyticsTag.ValueString() != "" {
			allFieldsEmpty = false
		}
		if !bc.WebAnalyticsToken.IsNull() && !bc.WebAnalyticsToken.IsUnknown() && bc.WebAnalyticsToken.ValueString() != "" {
			allFieldsEmpty = false
		}
		if allFieldsEmpty {
			data.BuildConfig = nil
		}
	}

	// Normalize source.config empty lists to null
	if data.Source != nil && data.Source.Config != nil {
		config := data.Source.Config
		if !config.PathExcludes.IsNull() && !config.PathExcludes.IsUnknown() && len(config.PathExcludes.Elements()) == 0 {
			config.PathExcludes = customfield.NullList[types.String](ctx)
		}
		if !config.PathIncludes.IsNull() && !config.PathIncludes.IsUnknown() && len(config.PathIncludes.Elements()) == 0 {
			config.PathIncludes = customfield.NullList[types.String](ctx)
		}
		if !config.PreviewBranchExcludes.IsNull() && !config.PreviewBranchExcludes.IsUnknown() && len(config.PreviewBranchExcludes.Elements()) == 0 {
			config.PreviewBranchExcludes = customfield.NullList[types.String](ctx)
		}
		if !config.PreviewBranchIncludes.IsNull() && !config.PreviewBranchIncludes.IsUnknown() && len(config.PreviewBranchIncludes.Elements()) == 0 {
			config.PreviewBranchIncludes = customfield.NullList[types.String](ctx)
		}
	}

	// Normalize deployment_configs empty pointer slices
	if !data.DeploymentConfigs.IsNull() && !data.DeploymentConfigs.IsUnknown() {
		configsValue, d := data.DeploymentConfigs.Value(ctx)
		diags.Append(d...)
		if diags.HasError() {
			return data, diags
		}

		previewModified := false
		productionModified := false

		// Normalize Preview config
		if !configsValue.Preview.IsNull() && !configsValue.Preview.IsUnknown() {
			previewValue, d := configsValue.Preview.Value(ctx)
			diags.Append(d...)
			if diags.HasError() {
				return data, diags
			}

			if previewValue.CompatibilityFlags != nil && len(*previewValue.CompatibilityFlags) == 0 {
				previewValue.CompatibilityFlags = nil
				previewModified = true
			}

			// Normalize empty placement to null
			if previewValue.Placement != nil && (previewValue.Placement.Mode.IsNull() || previewValue.Placement.Mode.ValueString() == "") {
				previewValue.Placement = nil
				previewModified = true
			}

			// Normalize empty binding maps to nil
			if normalizeEmptyMapsPreview(previewValue) {
				previewModified = true
			}

			if previewModified {
				configsValue.Preview, d = customfield.NewObject(ctx, previewValue)
				diags.Append(d...)
				if diags.HasError() {
					return data, diags
				}
			}
		}

		// Normalize Production config
		if !configsValue.Production.IsNull() && !configsValue.Production.IsUnknown() {
			productionValue, d := configsValue.Production.Value(ctx)
			diags.Append(d...)
			if diags.HasError() {
				return data, diags
			}

			if productionValue.CompatibilityFlags != nil && len(*productionValue.CompatibilityFlags) == 0 {
				productionValue.CompatibilityFlags = nil
				productionModified = true
			}

			// Normalize empty placement to null
			if productionValue.Placement != nil && (productionValue.Placement.Mode.IsNull() || productionValue.Placement.Mode.ValueString() == "") {
				productionValue.Placement = nil
				productionModified = true
			}

			// Normalize empty binding maps to nil
			if normalizeEmptyMapsProduction(productionValue) {
				productionModified = true
			}

			if productionModified {
				configsValue.Production, d = customfield.NewObject(ctx, productionValue)
				diags.Append(d...)
				if diags.HasError() {
					return data, diags
				}
			}
		}

		if previewModified || productionModified {
			data.DeploymentConfigs, d = customfield.NewObject(ctx, configsValue)
			diags.Append(d...)
			if diags.HasError() {
				return data, diags
			}
		}
	}

	return data, diags
}
