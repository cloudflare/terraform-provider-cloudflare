// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jinzhu/copier"
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

	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("migrations"), &data.Migrations)...)

	if resp.Diagnostics.HasError() {
		return
	}

	planMigrations := data.Migrations

	var assets *WorkersScriptMetadataAssetsModel
	if data.Assets != nil {
		assets = &WorkersScriptMetadataAssetsModel{
			Config:              data.Assets.Config,
			JWT:                 data.Assets.JWT,
			Directory:           data.Assets.Directory,
			AssetManifestSHA256: data.Assets.AssetManifestSHA256,
		}
	}
	err := handleAssets(ctx, r.client, data)
	if err != nil {
		resp.Diagnostics.AddError("failed to upload assets", err.Error())
		return
	}

	contentSHA256 := data.ContentSHA256
	contentType := data.ContentType

	if !data.ContentFile.IsNull() {
		content, err := readFile((data.ContentFile.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("failed to read file", err.Error())
			return
		}
		data.Content = types.StringValue(content)
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
	data.ContentSHA256 = contentSHA256
	data.ContentType = contentType
	data.Assets = assets
	data.Migrations = planMigrations

	// avoid storing `content` in state if `content_file` is configured
	if !data.ContentFile.IsNull() {
		data.Content = types.StringNull()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersScriptResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *WorkersScriptModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("migrations"), &data.Migrations)...)

	if resp.Diagnostics.HasError() {
		return
	}

	planMigrations := data.Migrations

	var assets *WorkersScriptMetadataAssetsModel
	if data.Assets != nil {
		assets = &WorkersScriptMetadataAssetsModel{
			Config:              data.Assets.Config,
			JWT:                 data.Assets.JWT,
			Directory:           data.Assets.Directory,
			AssetManifestSHA256: data.Assets.AssetManifestSHA256,
		}
	}
	err := handleAssets(ctx, r.client, data)
	if err != nil {
		resp.Diagnostics.AddError("failed to upload assets", err.Error())
		return
	}

	var state *WorkersScriptModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	contentSHA256 := data.ContentSHA256
	contentType := data.ContentType

	if !data.ContentFile.IsNull() {
		content, err := readFile((data.ContentFile.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("failed to read file", err.Error())
			return
		}
		data.Content = types.StringValue(content)
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
	data.ContentSHA256 = contentSHA256
	data.ContentType = contentType
	data.Assets = assets
	data.Migrations = planMigrations

	// avoid storing `content` in state if `content_file` is configured
	if !data.ContentFile.IsNull() {
		data.Content = types.StringNull()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersScriptResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *WorkersScriptModel
	var state *WorkersScriptModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	accountId := data.AccountID.ValueString()
	scriptName := data.ScriptName.ValueString()

	// fetch the script resource
	res := new(http.Response)
	path := fmt.Sprintf("accounts/%s/workers/services/%s", accountId, scriptName)
	err := r.client.Get(
		ctx,
		path,
		nil,
		&res,
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
	var service WorkersServiceResultEnvelope
	err = apijson.Unmarshal(bytes, &service)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	copier.CopyWithOption(&data, &service.Result.DefaultEnvironment.Script, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// fetch the script metadata and version settings
	res = new(http.Response)
	path = fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", accountId, scriptName)
	err = r.client.Get(
		ctx,
		path,
		nil,
		&res,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	bytes, _ = io.ReadAll(res.Body)
	var metadata WorkersScriptMetadataResultEnvelope
	err = apijson.Unmarshal(bytes, &metadata)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	copier.CopyWithOption(&data.WorkersScriptMetadataModel, &metadata.Result, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	// restore any secret_text `text` values from state since they aren't returned by the API
	var diags diag.Diagnostics
	data.Bindings, diags = UpdateSecretTextsFromState(
		ctx,
		data.Bindings,
		state.Bindings,
	)
	resp.Diagnostics.Append(diags...)

	if !state.Migrations.IsNull() {
		data.Migrations = state.Migrations
	}

	// fetch the script content
	scriptContentRes, err := r.client.Workers.Scripts.Content.Get(
		ctx,
		data.ScriptName.ValueString(),
		workers.ScriptContentGetParams{
			AccountID: cloudflare.F(accountId),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	switch scriptContentRes.StatusCode {
	case http.StatusOK:
		var content string
		mediaType, mediaTypeParams, err := mime.ParseMediaType(scriptContentRes.Header.Get("Content-Type"))
		if err != nil {
			resp.Diagnostics.AddError("failed parsing content-type", err.Error())
			return
		}
		if strings.HasPrefix(mediaType, "multipart/") {
			mr := multipart.NewReader(scriptContentRes.Body, mediaTypeParams["boundary"])
			p, err := mr.NextPart()
			if err != nil {
				resp.Diagnostics.AddError("failed to read response body", err.Error())
			}
			c, _ := io.ReadAll(p)
			content = string(c)
		} else {
			bytes, err = io.ReadAll(scriptContentRes.Body)
			if err != nil {
				resp.Diagnostics.AddError("failed to read response body", err.Error())
				return
			}
			content = string(bytes)
		}

		// only update `content` if `content_file` isn't being used instead
		if data.ContentFile.IsNull() {
			data.Content = types.StringValue(content)
		}

		// refresh the content hash in case the remote state has drifted
		if !data.ContentSHA256.IsNull() {
			hash, _ := calculateStringHash(content)
			data.ContentSHA256 = types.StringValue(hash)
		}
	case http.StatusNoContent:
		data.Content = types.StringNull()
	default:
		resp.Diagnostics.AddError("failed to fetch script content", fmt.Sprintf("%v %s", scriptContentRes.StatusCode, scriptContentRes.Status))
		return
	}

	// If the API returned an empty object for `placement`, treat it as null
	if data.Placement.Attributes()["mode"].IsNull() {
		data.Placement = data.Placement.NullValue(ctx).(customfield.NestedObject[WorkersScriptMetadataPlacementModel])
	}

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
	var data = new(WorkersScriptModel)

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
