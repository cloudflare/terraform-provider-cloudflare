// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*WorkersScriptResource)(nil)
var _ resource.ResourceWithModifyPlan = (*WorkersScriptResource)(nil)
var _ resource.ResourceWithImportState = (*WorkersScriptResource)(nil)

func NewResource() resource.Resource {
	return &WorkersScriptResource{}
}

// WorkersScriptResource defines the resource implementation.
type WorkersScriptResource struct {
	client *cloudflare.Client
}

func (r *WorkersScriptResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workers_script"
}

func (r *WorkersScriptResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *WorkersScriptResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *WorkersScriptModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Custom asset and content handling is not currently implemented
	// The required fields (Content, ContentFile, Migrations, etc.) are not exposed in the model
	// This needs to be reimplemented once the schema is updated

	dataBytes, formDataContentType, err := data.MarshalMultipart()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize multipart http request", err.Error())
		return
	}
	res := new(http.Response)
	env := WorkersScriptResultEnvelope{*data}
	_, err = r.client.Workers.Scripts.Update(
		ctx,
		data.ScriptName.ValueString(),
		workers.ScriptUpdateParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithRequestBody(formDataContentType, dataBytes),
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
	data.ID = data.ScriptName

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersScriptResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *WorkersScriptModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Custom asset and content handling is not currently implemented
	// The required fields (Content, ContentFile, Migrations, etc.) are not exposed in the model
	// This needs to be reimplemented once the schema is updated

	var state *WorkersScriptModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, formDataContentType, err := data.MarshalMultipart()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize multipart http request", err.Error())
		return
	}
	res := new(http.Response)
	env := WorkersScriptResultEnvelope{*data}
	_, err = r.client.Workers.Scripts.Update(
		ctx,
		data.ScriptName.ValueString(),
		workers.ScriptUpdateParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithRequestBody(formDataContentType, dataBytes),
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
	data.ID = data.ScriptName

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersScriptResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *WorkersScriptModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Custom Read logic is not currently implemented
	// The original implementation referenced fields and types that don't exist in the model
	// This needs to be reimplemented once the schema is updated

	// TODO: Workers script Read is not fully implemented
	// The custom logic was removed due to incompatible model changes
	// For now, we'll keep the data from state as-is
	// This needs to be reimplemented with proper API calls once the schema is updated

	data.ID = data.ScriptName

	// Note: The following custom features are not implemented:
	// - Content/ContentFile handling
	// - Assets handling
	// - Secret text restoration from state
	// - Custom metadata fetching

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersScriptResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *WorkersScriptModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Workers.Scripts.Delete(
		ctx,
		data.ScriptName.ValueString(),
		workers.ScriptDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.ScriptName

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersScriptResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *WorkersScriptModel = new(WorkersScriptModel)

	path_account_id := ""
	path_script_name := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<script_name>",
		&path_account_id,
		&path_script_name,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ScriptName = types.StringValue(path_script_name)

	res := new(http.Response)
	_, err := r.client.Workers.Scripts.Get(
		ctx,
		path_script_name,
		workers.ScriptGetParams{
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
	err = apijson.Unmarshal(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data.ID = data.ScriptName

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersScriptResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}

	// After running all the provider plan modification, if there are no differences/updates, return.
	if req.Plan.Raw.Equal(req.State.Raw) {
		return
	}

	// Terraform Framework checks if there are planned changes and if so, marks computed attribute values as unknown.
	// This occurs before any plan modifiers are run, so if a change doesn't get planned until running a plan modifier (such as recomputing `asset_manifest_sha256`),
	// any computed attribute values from the previous state are carried over without being marked as unknown.
	// Since these are now considered "known values", they MUST match after apply, or else Terraform will throw
	// "Error: Provider produced inconsistent result after apply".
	// To prevent this, we must explicitly mark any computed attributes we know can change as unknown.
	resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("modified_on"), timetypes.NewRFC3339Unknown())...)
	resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("has_assets"), types.BoolUnknown())...)
}
