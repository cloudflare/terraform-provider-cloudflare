// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/accounts"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*AccountMemberResource)(nil)
var _ resource.ResourceWithModifyPlan = (*AccountMemberResource)(nil)
var _ resource.ResourceWithImportState = (*AccountMemberResource)(nil)

func NewResource() resource.Resource {
	return &AccountMemberResource{}
}

// AccountMemberResource defines the resource implementation.
type AccountMemberResource struct {
	client *cloudflare.Client
}

func (r *AccountMemberResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account_member"
}

func (r *AccountMemberResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AccountMemberResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AccountMemberModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.marshalCustom()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := AccountMemberResultEnvelope{*data}
	_, err = r.client.Accounts.Members.New(
		ctx,
		accounts.MemberNewParams{
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

	// Due to a bug in the POST endpoint, roles are not returned on the POST
	// response. But, they are returned later on the GET response (and PUT). So,
	// to get around this we are doing a GET after every POST to retrieve the
	// full actual state of the member.
	// This can be removed once either the POST endpoint is fixed or roles are
	// removed.
	res = new(http.Response)
	_, err = r.client.Accounts.Members.Get(
		ctx,
		data.ID.ValueString(),
		accounts.MemberGetParams{
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
	bytes, _ = io.ReadAll(res.Body)
	data, err = unmarshalCustom(bytes, data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccountMemberResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AccountMemberModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *AccountMemberModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// check if the user has configured roles or policies
	var config *AccountMemberModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	permissionType := checkConfiguredPermissionType(config)

	dataBytes, err := data.marshalCustomForUpdate(*state, permissionType)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	_, err = r.client.Accounts.Members.Update(
		ctx,
		data.ID.ValueString(),
		accounts.MemberUpdateParams{
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
	data, err = unmarshalCustom(bytes, data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccountMemberResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AccountMemberModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	_, err := r.client.Accounts.Members.Get(
		ctx,
		data.ID.ValueString(),
		accounts.MemberGetParams{
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
	data, err = unmarshalComputedCustom(bytes, data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccountMemberResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AccountMemberModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Accounts.Members.Delete(
		ctx,
		data.ID.ValueString(),
		accounts.MemberDeleteParams{
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

func (r *AccountMemberResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(AccountMemberModel)

	path_account_id := ""
	path_member_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<member_id>",
		&path_account_id,
		&path_member_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ID = types.StringValue(path_member_id)

	res := new(http.Response)
	_, err := r.client.Accounts.Members.Get(
		ctx,
		path_member_id,
		accounts.MemberGetParams{
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
	data, err = unmarshalCustom(bytes, data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccountMemberResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var config, state, plan *AccountMemberModel

	// Get config, state, and plan
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if config == nil && plan == nil {
		// no plan changes for destroy
		return
	}

	if state != nil {
		// Always suppress user field diffs since it's computed and can change independently
		// The user field contains computed values that may be updated by the API
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("user"), state.User)...)
	}

	if config != nil {
		configurePermissionType := checkConfiguredPermissionType(config)
		switch configurePermissionType {
		case Roles:
			if state == nil || !reflect.DeepEqual(plan.Roles, state.Roles) {
				// if roles are changing, set policies to unknown
				resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("policies"), plan.Policies.UnknownValue(ctx))...)
			} else {
				// else preserve state policies
				resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("policies"), state.Policies)...)
			}
		case Policies:
			if state == nil || !plan.Policies.Equal(state.Policies) {
				// if policies are changing, set roles to unknown
				resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("roles"), types.SetUnknown(types.StringType))...)
			} else {
				// else preserve state roles
				resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("roles"), state.Roles)...)
			}
		}
	}
}
