// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*WorkerResource)(nil)
var _ resource.ResourceWithModifyPlan = (*WorkerResource)(nil)
var _ resource.ResourceWithImportState = (*WorkerResource)(nil)

// WorkerBuildsTriggerResponse represents the API response for a builds trigger.
type WorkerBuildsTriggerResponse struct {
	Result WorkerBuildsTriggerResult `json:"result"`
}

// WorkerBuildsTriggerResult represents the trigger configuration from the API.
type WorkerBuildsTriggerResult struct {
	TriggerUUID                     string   `json:"trigger_uuid"`
	ScriptName                      string   `json:"script_name"`
	BuildCommand                    string   `json:"build_command"`
	DeployCommand                   string   `json:"deploy_command"`
	Branch                          string   `json:"branch"`
	BranchIncludes                  []string `json:"branch_includes"`
	BranchExcludes                  []string `json:"branch_excludes"`
	NonProductionDeploymentsEnabled bool     `json:"non_production_deployments_enabled"`
	ProviderType                    string   `json:"provider_type"`
	ProviderAccountName             string   `json:"provider_account_name"`
	ProviderAccountID               string   `json:"provider_account_id"`
	RepoID                          string   `json:"repo_id"`
	RepoName                        string   `json:"repo_name"`
	CreatedOn                       string   `json:"created_on"`
	ModifiedOn                      string   `json:"modified_on"`
}

// WorkerBuildsRepoConnectionRequest represents the request to connect a repository.
type WorkerBuildsRepoConnectionRequest struct {
	Builds []WorkerBuildsRepoConnectionBuild `json:"builds"`
}

// WorkerBuildsRepoConnectionBuild represents a build configuration for repo connection.
type WorkerBuildsRepoConnectionBuild struct {
	ScriptName    string `json:"script_name"`
	RepoID        string `json:"repo_id"`
	ProviderType  string `json:"provider_type"`
	Branch        string `json:"branch"`
	BuildCommand  string `json:"build_command,omitempty"`
	DeployCommand string `json:"deploy_command,omitempty"`
	RootDir       string `json:"root_dir,omitempty"`
}

// WorkerBuildsTriggerRequest represents the request to create/update a trigger.
type WorkerBuildsTriggerRequest struct {
	ScriptName                      string   `json:"script_name"`
	BuildCommand                    string   `json:"build_command,omitempty"`
	DeployCommand                   string   `json:"deploy_command,omitempty"`
	Branch                          string   `json:"branch,omitempty"`
	BranchIncludes                  []string `json:"branch_includes,omitempty"`
	BranchExcludes                  []string `json:"branch_excludes,omitempty"`
	NonProductionDeploymentsEnabled bool     `json:"non_production_deployments_enabled"`
	ProviderType                    string   `json:"provider_type,omitempty"`
	ProviderAccountName             string   `json:"provider_account_name,omitempty"`
	ProviderAccountID               string   `json:"provider_account_id,omitempty"`
	RepoID                          string   `json:"repo_id,omitempty"`
	RepoName                        string   `json:"repo_name,omitempty"`
	Enabled                         bool     `json:"enabled"`
}

// getBuildsTrigger retrieves the builds trigger for a worker.
func (r *WorkerResource) getBuildsTrigger(ctx context.Context, accountID, scriptName string) (*WorkerBuildsTriggerResult, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/builds/workers/%s/triggers", accountID, scriptName), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := r.doAuthenticatedRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil, nil
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get builds trigger: %s", res.Status)
	}

	var triggerResp WorkerBuildsTriggerResponse
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &triggerResp); err != nil {
		return nil, err
	}

	return &triggerResp.Result, nil
}

// createBuildsTrigger creates a builds trigger for a worker.
func (r *WorkerResource) createBuildsTrigger(ctx context.Context, accountID string, trigger WorkerBuildsTriggerRequest) (*WorkerBuildsTriggerResult, error) {
	triggerBytes, err := json.Marshal(trigger)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/builds/triggers", accountID), bytes.NewReader(triggerBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := r.doAuthenticatedRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("failed to create builds trigger: %s - %s", res.Status, string(bodyBytes))
	}

	var triggerResp WorkerBuildsTriggerResponse
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &triggerResp); err != nil {
		return nil, err
	}

	return &triggerResp.Result, nil
}

// updateBuildsTrigger updates a builds trigger.
func (r *WorkerResource) updateBuildsTrigger(ctx context.Context, accountID, triggerUUID string, trigger WorkerBuildsTriggerRequest) (*WorkerBuildsTriggerResult, error) {
	triggerBytes, err := json.Marshal(trigger)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/builds/triggers/%s", accountID, triggerUUID), bytes.NewReader(triggerBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := r.doAuthenticatedRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("failed to update builds trigger: %s - %s", res.Status, string(bodyBytes))
	}

	var triggerResp WorkerBuildsTriggerResponse
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &triggerResp); err != nil {
		return nil, err
	}

	return &triggerResp.Result, nil
}

// deleteBuildsTrigger deletes a builds trigger.
func (r *WorkerResource) deleteBuildsTrigger(ctx context.Context, accountID, triggerUUID string) error {
	req, err := http.NewRequestWithContext(ctx, "DELETE", fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/builds/triggers/%s", accountID, triggerUUID), nil)
	if err != nil {
		return err
	}

	res, err := r.doAuthenticatedRequest(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 && res.StatusCode != 404 {
		return fmt.Errorf("failed to delete builds trigger: %s", res.Status)
	}

	return nil
}

// connectBuildsRepo connects a repository to a worker.
func (r *WorkerResource) connectBuildsRepo(ctx context.Context, accountID string, conn WorkerBuildsRepoConnectionRequest) error {
	connBytes, err := json.Marshal(conn)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/builds/repos/connections", accountID), bytes.NewReader(connBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := r.doAuthenticatedRequest(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(res.Body)
		return fmt.Errorf("failed to connect builds repo: %s - %s", res.Status, string(bodyBytes))
	}

	return nil
}

// doAuthenticatedRequest performs an HTTP request using the SDK's authentication.
func (r *WorkerResource) doAuthenticatedRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+os.Getenv("CLOUDFLARE_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	return httpClient.Do(req)
}

func NewResource() resource.Resource {
	return &WorkerResource{}
}

// WorkerResource defines the resource implementation.
type WorkerResource struct {
	client *cloudflare.Client
}

func (r *WorkerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_worker"
}

func (r *WorkerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *WorkerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *WorkerModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := WorkerResultEnvelope{*data}
	_, err = r.client.Workers.Beta.Workers.New(
		ctx,
		workers.BetaWorkerNewParams{
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

	if len(data.Builds) > 0 {
		builds := data.Builds[0]

		if !builds.Enabled.IsNull() && builds.Enabled.ValueBool() {
			accountID := data.AccountID.ValueString()
			scriptName := data.Name.ValueString()

			connReq := WorkerBuildsRepoConnectionRequest{
				Builds: []WorkerBuildsRepoConnectionBuild{
					{
						ScriptName:   scriptName,
						RepoID:       builds.RepoID.ValueString(),
						ProviderType: builds.ProviderType.ValueString(),
						Branch:       builds.Branch.ValueString(),
					},
				},
			}
			if !builds.BuildCommand.IsNull() {
				connReq.Builds[0].BuildCommand = builds.BuildCommand.ValueString()
			}
			if !builds.DeployCommand.IsNull() {
				connReq.Builds[0].DeployCommand = builds.DeployCommand.ValueString()
			}

			if err := r.connectBuildsRepo(ctx, accountID, connReq); err != nil {
				resp.Diagnostics.AddError("failed to connect builds repo", err.Error())
				return
			}

			triggerReq := WorkerBuildsTriggerRequest{
				ScriptName:                      scriptName,
				BuildCommand:                    builds.BuildCommand.ValueString(),
				DeployCommand:                   builds.DeployCommand.ValueString(),
				Branch:                          builds.Branch.ValueString(),
				NonProductionDeploymentsEnabled: builds.NonProductionDeploymentsEnabled.ValueBool(),
				ProviderType:                    builds.ProviderType.ValueString(),
				ProviderAccountName:             builds.ProviderAccountName.ValueString(),
				ProviderAccountID:               builds.ProviderAccountID.ValueString(),
				RepoID:                          builds.RepoID.ValueString(),
				RepoName:                        builds.RepoName.ValueString(),
				Enabled:                         builds.Enabled.ValueBool(),
			}

			if !builds.BranchIncludes.IsNull() {
				var branchIncludes []string
				builds.BranchIncludes.ElementsAs(ctx, &branchIncludes, false)
				triggerReq.BranchIncludes = branchIncludes
			}
			if !builds.BranchExcludes.IsNull() {
				var branchExcludes []string
				builds.BranchExcludes.ElementsAs(ctx, &branchExcludes, false)
				triggerReq.BranchExcludes = branchExcludes
			}

			if _, err := r.createBuildsTrigger(ctx, accountID, triggerReq); err != nil {
				resp.Diagnostics.AddError("failed to create builds trigger", err.Error())
				return
			}
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *WorkerModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *WorkerModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := WorkerResultEnvelope{*data}
	_, err = r.client.Workers.Beta.Workers.Update(
		ctx,
		data.ID.ValueString(),
		workers.BetaWorkerUpdateParams{
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

	if len(data.Builds) > 0 {
		planBuilds := data.Builds[0]

		if !planBuilds.Enabled.IsNull() && planBuilds.Enabled.ValueBool() {
			accountID := data.AccountID.ValueString()
			scriptName := data.Name.ValueString()

			existingTrigger, err := r.getBuildsTrigger(ctx, accountID, scriptName)
			if err != nil {
				resp.Diagnostics.AddError("failed to get builds trigger", err.Error())
				return
			}

			if existingTrigger != nil {
				triggerReq := WorkerBuildsTriggerRequest{
					ScriptName:                      scriptName,
					BuildCommand:                    planBuilds.BuildCommand.ValueString(),
					DeployCommand:                   planBuilds.DeployCommand.ValueString(),
					Branch:                          planBuilds.Branch.ValueString(),
					NonProductionDeploymentsEnabled: planBuilds.NonProductionDeploymentsEnabled.ValueBool(),
					ProviderType:                    planBuilds.ProviderType.ValueString(),
					ProviderAccountName:             planBuilds.ProviderAccountName.ValueString(),
					ProviderAccountID:               planBuilds.ProviderAccountID.ValueString(),
					RepoID:                          planBuilds.RepoID.ValueString(),
					RepoName:                        planBuilds.RepoName.ValueString(),
					Enabled:                         planBuilds.Enabled.ValueBool(),
				}

				if !planBuilds.BranchIncludes.IsNull() {
					var branchIncludes []string
					planBuilds.BranchIncludes.ElementsAs(ctx, &branchIncludes, false)
					triggerReq.BranchIncludes = branchIncludes
				}
				if !planBuilds.BranchExcludes.IsNull() {
					var branchExcludes []string
					planBuilds.BranchExcludes.ElementsAs(ctx, &branchExcludes, false)
					triggerReq.BranchExcludes = branchExcludes
				}

				if _, err := r.updateBuildsTrigger(ctx, accountID, existingTrigger.TriggerUUID, triggerReq); err != nil {
					resp.Diagnostics.AddError("failed to update builds trigger", err.Error())
					return
				}
			} else {
				connReq := WorkerBuildsRepoConnectionRequest{
					Builds: []WorkerBuildsRepoConnectionBuild{
						{
							ScriptName:   scriptName,
							RepoID:       planBuilds.RepoID.ValueString(),
							ProviderType: planBuilds.ProviderType.ValueString(),
							Branch:       planBuilds.Branch.ValueString(),
						},
					},
				}
				if !planBuilds.BuildCommand.IsNull() {
					connReq.Builds[0].BuildCommand = planBuilds.BuildCommand.ValueString()
				}
				if !planBuilds.DeployCommand.IsNull() {
					connReq.Builds[0].DeployCommand = planBuilds.DeployCommand.ValueString()
				}

				if err := r.connectBuildsRepo(ctx, accountID, connReq); err != nil {
					resp.Diagnostics.AddError("failed to connect builds repo", err.Error())
					return
				}

				triggerReq := WorkerBuildsTriggerRequest{
					ScriptName:                      scriptName,
					BuildCommand:                    planBuilds.BuildCommand.ValueString(),
					DeployCommand:                   planBuilds.DeployCommand.ValueString(),
					Branch:                          planBuilds.Branch.ValueString(),
					NonProductionDeploymentsEnabled: planBuilds.NonProductionDeploymentsEnabled.ValueBool(),
					ProviderType:                    planBuilds.ProviderType.ValueString(),
					ProviderAccountName:             planBuilds.ProviderAccountName.ValueString(),
					ProviderAccountID:               planBuilds.ProviderAccountID.ValueString(),
					RepoID:                          planBuilds.RepoID.ValueString(),
					RepoName:                        planBuilds.RepoName.ValueString(),
					Enabled:                         planBuilds.Enabled.ValueBool(),
				}

				if !planBuilds.BranchIncludes.IsNull() {
					var branchIncludes []string
					planBuilds.BranchIncludes.ElementsAs(ctx, &branchIncludes, false)
					triggerReq.BranchIncludes = branchIncludes
				}
				if !planBuilds.BranchExcludes.IsNull() {
					var branchExcludes []string
					planBuilds.BranchExcludes.ElementsAs(ctx, &branchExcludes, false)
					triggerReq.BranchExcludes = branchExcludes
				}

				if _, err := r.createBuildsTrigger(ctx, accountID, triggerReq); err != nil {
					resp.Diagnostics.AddError("failed to create builds trigger", err.Error())
					return
				}
			}
		}
	} else if len(state.Builds) > 0 && (len(data.Builds) == 0 || data.Builds[0].Enabled.IsNull() || !data.Builds[0].Enabled.ValueBool()) {
		accountID := data.AccountID.ValueString()
		scriptName := data.Name.ValueString()

		existingTrigger, err := r.getBuildsTrigger(ctx, accountID, scriptName)
		if err != nil {
			resp.Diagnostics.AddError("failed to get builds trigger", err.Error())
			return
		}

		if existingTrigger != nil {
			if err := r.deleteBuildsTrigger(ctx, accountID, existingTrigger.TriggerUUID); err != nil {
				resp.Diagnostics.AddError("failed to delete builds trigger", err.Error())
				return
			}
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *WorkerModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := WorkerResultEnvelope{*data}
	_, err := r.client.Workers.Beta.Workers.Get(
		ctx,
		data.ID.ValueString(),
		workers.BetaWorkerGetParams{
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *WorkerModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Workers.Beta.Workers.Delete(
		ctx,
		data.ID.ValueString(),
		workers.BetaWorkerDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(WorkerModel)

	path_account_id := ""
	path_worker_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<worker_id>",
		&path_account_id,
		&path_worker_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ID = types.StringValue(path_worker_id)

	res := new(http.Response)
	env := WorkerResultEnvelope{*data}
	_, err := r.client.Workers.Beta.Workers.Get(
		ctx,
		path_worker_id,
		workers.BetaWorkerGetParams{
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkerResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var plan, state *WorkerModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || plan == nil || state == nil {
		return
	}

	// If there are any meaningful changes to user-configurable attributes, do
	// nothing so that updated_on can legitimately change on update.
	if (!plan.Name.IsUnknown() && !plan.Name.Equal(state.Name)) ||
		(!plan.Logpush.IsUnknown() && !plan.Logpush.Equal(state.Logpush)) ||
		(!plan.Tags.IsUnknown() && !plan.Tags.Equal(state.Tags)) ||
		(!plan.Observability.IsUnknown() && !plan.Observability.Equal(state.Observability)) ||
		(!plan.Subdomain.IsUnknown() && !plan.Subdomain.Equal(state.Subdomain)) ||
		(!plan.TailConsumers.IsUnknown() && !plan.TailConsumers.Equal(state.TailConsumers)) {
		return
	}

	// No changes to user-configurable attributes, so copy updated_on timestamp
	// from state to avoid spurious changes.
	resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("updated_on"), state.UpdatedOn)...)
	resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("references"), state.References)...)
}
