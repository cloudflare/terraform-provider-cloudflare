package container_application

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigure = (*ContainerApplicationResource)(nil)
var _ resource.ResourceWithImportState = (*ContainerApplicationResource)(nil)

func NewResource() resource.Resource {
	return &ContainerApplicationResource{}
}

type ContainerApplicationResource struct {
	client *cloudflare.Client
}

func (r *ContainerApplicationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_container_application"
}

func (r *ContainerApplicationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ContainerApplicationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ContainerApplicationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountID := data.AccountID.ValueString()
	scriptName := data.ScriptName.ValueString()
	className := data.ClassName.ValueString()

	// Compute default name if not provided
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		data.Name = types.StringValue(fmt.Sprintf("%s-%s", scriptName, className))
	}

	// Resolve Durable Object namespace
	namespaceID, err := r.resolveDONamespace(ctx, accountID, scriptName, className)
	if err != nil {
		resp.Diagnostics.AddError("failed to resolve Durable Object namespace", err.Error())
		return
	}

	// Build create request
	createReq, err := data.toCreateRequest(ctx, namespaceID)
	if err != nil {
		resp.Diagnostics.AddError("failed to build create request", err.Error())
		return
	}

	bodyBytes, err := json.Marshal(createReq)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize create request", err.Error())
		return
	}

	// POST /accounts/{account_id}/containers/applications
	res := new(http.Response)
	err = r.client.Post(
		ctx,
		fmt.Sprintf("accounts/%s/containers/applications", accountID),
		nil,
		&res,
		option.WithRequestBody("application/json", bodyBytes),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create container application", err.Error())
		return
	}
	defer res.Body.Close()

	respBytes, _ := io.ReadAll(res.Body)

	var envelope apiApplicationEnvelope
	err = json.Unmarshal(respBytes, &envelope)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize create response", err.Error())
		return
	}

	data.fromAPIApplication(envelope.Result)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContainerApplicationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ContainerApplicationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountID := data.AccountID.ValueString()
	appID := data.ApplicationID.ValueString()

	// Preserve state-only fields before API read (not returned by the API)
	scriptName := data.ScriptName
	className := data.ClassName
	rolloutStepPercentage := data.RolloutStepPercentage
	rolloutKind := data.RolloutKind
	constraints := data.Constraints
	vcpu := data.Vcpu
	memoryMib := data.MemoryMib
	diskSizeMb := data.DiskSizeMb
	observability := data.Observability
	wranglerSSH := data.WranglerSSH
	authorizedKeys := data.AuthorizedKeys
	trustedUserCAKeys := data.TrustedUserCAKeys

	app, err := r.getApplication(ctx, accountID, appID)
	if err != nil {
		resp.Diagnostics.AddError("failed to read container application", err.Error())
		return
	}
	if app == nil {
		resp.Diagnostics.AddWarning("Resource not found", "The container application was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}

	data.fromAPIApplication(*app)

	// Restore state-only fields
	data.ScriptName = scriptName
	data.ClassName = className
	data.RolloutStepPercentage = rolloutStepPercentage
	data.RolloutKind = rolloutKind

	// Preserve these from state if the API didn't return them
	if data.Constraints == nil {
		data.Constraints = constraints
	}
	if data.Vcpu.IsNull() || data.Vcpu.IsUnknown() {
		data.Vcpu = vcpu
	}
	if data.MemoryMib.IsNull() || data.MemoryMib.IsUnknown() {
		data.MemoryMib = memoryMib
	}
	if data.DiskSizeMb.IsNull() || data.DiskSizeMb.IsUnknown() {
		data.DiskSizeMb = diskSizeMb
	}
	if data.Observability == nil {
		data.Observability = observability
	}
	if data.WranglerSSH == nil {
		data.WranglerSSH = wranglerSSH
	}
	if data.AuthorizedKeys == nil {
		data.AuthorizedKeys = authorizedKeys
	}
	if data.TrustedUserCAKeys == nil {
		data.TrustedUserCAKeys = trustedUserCAKeys
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContainerApplicationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ContainerApplicationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountID := data.AccountID.ValueString()
	appID := data.ApplicationID.ValueString()

	// Build modify request
	modifyReq, err := data.toModifyRequest(ctx)
	if err != nil {
		resp.Diagnostics.AddError("failed to build modify request", err.Error())
		return
	}

	bodyBytes, err := json.Marshal(modifyReq)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize modify request", err.Error())
		return
	}

	// PATCH /accounts/{account_id}/containers/applications/{app_id}
	res := new(http.Response)
	err = r.client.Patch(
		ctx,
		fmt.Sprintf("accounts/%s/containers/applications/%s", accountID, appID),
		nil,
		&res,
		option.WithRequestBody("application/json", bodyBytes),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to modify container application", err.Error())
		return
	}
	defer res.Body.Close()

	respBytes, _ := io.ReadAll(res.Body)
	var envelope apiApplicationEnvelope
	err = json.Unmarshal(respBytes, &envelope)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize modify response", err.Error())
		return
	}

	// Create rollout if rollout_kind is not "none"
	rolloutKind := data.RolloutKind.ValueString()
	if rolloutKind != "none" {
		rolloutReq := data.toRolloutRequest()
		rolloutBytes, err := json.Marshal(rolloutReq)
		if err != nil {
			resp.Diagnostics.AddError("failed to serialize rollout request", err.Error())
			return
		}

		rolloutRes := new(http.Response)
		err = r.client.Post(
			ctx,
			fmt.Sprintf("accounts/%s/containers/applications/%s/rollouts", accountID, appID),
			nil,
			&rolloutRes,
			option.WithRequestBody("application/json", rolloutBytes),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			resp.Diagnostics.AddError("failed to create container rollout", err.Error())
			return
		}
		defer rolloutRes.Body.Close()
	}

	data.fromAPIApplication(envelope.Result)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContainerApplicationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ContainerApplicationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountID := data.AccountID.ValueString()
	appID := data.ApplicationID.ValueString()

	// DELETE /accounts/{account_id}/containers/applications/{app_id}
	res := new(http.Response)
	err := r.client.Delete(
		ctx,
		fmt.Sprintf("accounts/%s/containers/applications/%s", accountID, appID),
		nil,
		&res,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to delete container application", err.Error())
		return
	}
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
}

func (r *ContainerApplicationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(ContainerApplicationModel)

	pathAccountID := ""
	pathAppID := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<application_id>",
		&pathAccountID,
		&pathAppID,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(pathAccountID)

	app, err := r.getApplication(ctx, pathAccountID, pathAppID)
	if err != nil {
		resp.Diagnostics.AddError("failed to read container application for import", err.Error())
		return
	}
	if app == nil {
		resp.Diagnostics.AddError("container application not found", fmt.Sprintf("Application %s not found in account %s", pathAppID, pathAccountID))
		return
	}

	data.fromAPIApplication(*app)

	// These fields aren't available from the API; user will need to set them after import
	data.ScriptName = types.StringValue("")
	data.ClassName = types.StringValue("")
	data.RolloutStepPercentage = types.Int64Value(100)
	data.RolloutKind = types.StringValue("full_auto")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// resolveDONamespace looks up the Durable Object namespace ID for the given
// script_name and class_name combination. It retries a few times because the
// namespace may not be immediately available after script deployment.
func (r *ContainerApplicationResource) resolveDONamespace(ctx context.Context, accountID, scriptName, className string) (string, error) {
	const maxRetries = 5
	const retryDelay = 3 * time.Second

	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			tflog.Info(ctx, fmt.Sprintf("Durable Object namespace not found yet, retrying in %s (attempt %d/%d)", retryDelay, attempt+1, maxRetries))
			time.Sleep(retryDelay)
		}

		nsID, err := r.fetchDONamespace(ctx, accountID, scriptName, className)
		if err != nil {
			return "", err
		}
		if nsID != "" {
			return nsID, nil
		}
	}

	return "", fmt.Errorf("Durable Object namespace not found for class %q in script %q after %d attempts. Ensure the Worker script is deployed with the Durable Object binding before creating the container application", className, scriptName, maxRetries)
}

func (r *ContainerApplicationResource) fetchDONamespace(ctx context.Context, accountID, scriptName, className string) (string, error) {
	res := new(http.Response)
	err := r.client.Get(
		ctx,
		fmt.Sprintf("accounts/%s/workers/durable_objects/namespaces?per_page=1000", accountID),
		nil,
		&res,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		return "", fmt.Errorf("failed to list Durable Object namespaces: %w", err)
	}
	defer res.Body.Close()

	respBytes, _ := io.ReadAll(res.Body)
	var envelope apiDONamespaceListEnvelope
	err = json.Unmarshal(respBytes, &envelope)
	if err != nil {
		return "", fmt.Errorf("failed to parse Durable Object namespaces: %w", err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Found %d Durable Object namespaces, looking for class=%q script=%q", len(envelope.Result), className, scriptName))

	for _, ns := range envelope.Result {
		if ns.Class == className && ns.Script == scriptName {
			tflog.Info(ctx, fmt.Sprintf("Resolved Durable Object namespace: %s", ns.ID))
			return ns.ID, nil
		}
	}

	return "", nil
}

// getApplication fetches a single container application by listing all applications
// and filtering by ID.
func (r *ContainerApplicationResource) getApplication(ctx context.Context, accountID, appID string) (*apiApplication, error) {
	res := new(http.Response)
	err := r.client.Get(
		ctx,
		fmt.Sprintf("accounts/%s/containers/applications", accountID),
		nil,
		&res,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list container applications: %w", err)
	}
	defer res.Body.Close()

	respBytes, _ := io.ReadAll(res.Body)
	var envelope apiApplicationListEnvelope
	err = json.Unmarshal(respBytes, &envelope)
	if err != nil {
		return nil, fmt.Errorf("failed to parse container applications: %w", err)
	}

	for _, app := range envelope.Result {
		if app.ID == appID {
			return &app, nil
		}
	}

	return nil, nil
}
