// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_block_sender

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/email_security"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*EmailSecurityBlockSenderResource)(nil)
var _ resource.ResourceWithModifyPlan = (*EmailSecurityBlockSenderResource)(nil)
var _ resource.ResourceWithImportState = (*EmailSecurityBlockSenderResource)(nil)

func NewResource() resource.Resource {
	return &EmailSecurityBlockSenderResource{}
}

// EmailSecurityBlockSenderResource defines the resource implementation.
type EmailSecurityBlockSenderResource struct {
	client *cloudflare.Client
}

func (r *EmailSecurityBlockSenderResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_email_security_block_sender"
}

func (r *EmailSecurityBlockSenderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *EmailSecurityBlockSenderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *EmailSecurityBlockSenderModel

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
	env := EmailSecurityBlockSenderResultEnvelope{data.Body}
	_, err = r.client.EmailSecurity.Settings.BlockSenders.New(
		ctx,
		email_security.SettingBlockSenderNewParams{
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
	data.Body = env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EmailSecurityBlockSenderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *EmailSecurityBlockSenderModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *EmailSecurityBlockSenderModel

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
	env := EmailSecurityBlockSenderResultEnvelope{data.Body}
	_, err = r.client.EmailSecurity.Settings.BlockSenders.Edit(
		ctx,
		data.ID.ValueInt64(),
		email_security.SettingBlockSenderEditParams{
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
	data.Body = env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EmailSecurityBlockSenderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *EmailSecurityBlockSenderModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := EmailSecurityBlockSenderResultEnvelope{data.Body}
	_, err := r.client.EmailSecurity.Settings.BlockSenders.Get(
		ctx,
		data.ID.ValueInt64(),
		email_security.SettingBlockSenderGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
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
	data.Body = env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EmailSecurityBlockSenderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *EmailSecurityBlockSenderModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.EmailSecurity.Settings.BlockSenders.Delete(
		ctx,
		data.ID.ValueInt64(),
		email_security.SettingBlockSenderDeleteParams{
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

func (r *EmailSecurityBlockSenderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *EmailSecurityBlockSenderModel = new(EmailSecurityBlockSenderModel)

	path_account_id := ""
	path_pattern_id := int64(0)
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<pattern_id>",
		&path_account_id,
		&path_pattern_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ID = types.Int64Value(path_pattern_id)

	res := new(http.Response)
	env := EmailSecurityBlockSenderResultEnvelope{data.Body}
	_, err := r.client.EmailSecurity.Settings.BlockSenders.Get(
		ctx,
		path_pattern_id,
		email_security.SettingBlockSenderGetParams{
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
	data.Body = env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EmailSecurityBlockSenderResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
