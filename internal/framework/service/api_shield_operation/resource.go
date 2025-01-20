package apishieldoperation

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/flatteners"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &APIShieldOperationResource{}
var _ resource.ResourceWithImportState = &APIShieldOperationResource{}

func NewResource() resource.Resource {
	return &APIShieldOperationResource{}
}

type APIShieldOperationResource struct {
	client *muxclient.Client
}

func (r *APIShieldOperationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_shield_operation"
}

func (r *APIShieldOperationResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*muxclient.Client)
}

func (r *APIShieldOperationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model APIShieldOperationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ops, err := r.client.V1.CreateAPIShieldOperations(
		ctx,
		cloudflare.ZoneIdentifier(model.ZoneID.ValueString()),
		cloudflare.CreateAPIShieldOperationsParams{
			Operations: []cloudflare.APIShieldBasicOperation{
				{
					Method:   model.Method.ValueString(),
					Host:     model.Host.ValueString(),
					Endpoint: model.Endpoint.ValueString(),
				},
			},
		},
	)

	if err != nil {
		resp.Diagnostics.AddError("Error creating API Shield Operation", err.Error())
		return
	}

	if len(ops) != 1 {
		resp.Diagnostics.AddError("Error creating API Shield Operation", fmt.Sprintf("expected 1 operation in response but got %d", len(ops)))
		return
	}

	op := ops[0]
	// The API normalizes the response, so we must not override it on create.
	// See https://github.com/hashicorp/terraform/blob/main/docs/resource-instance-change-lifecycle.md
	// Normalization: the remote API has returned some data in a different form than was recorded in the Previous Run State, but the meaning is unchanged.
	// In this case, the provider should return the exact value from the Previous Run State,
	// thereby preserving the value as it was written by the user in the configuration and thus avoiding unwanted cascading changes to elsewhere in the configuration.
	model.ID = flatteners.String(op.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, model)...)
}

func (r *APIShieldOperationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var model APIShieldOperationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	op, err := r.client.V1.GetAPIShieldOperation(
		ctx,
		cloudflare.ZoneIdentifier(model.ZoneID.ValueString()),
		cloudflare.GetAPIShieldOperationParams{
			OperationID: model.ID.ValueString(),
		})

	if err != nil {
		resp.Diagnostics.AddError("Error reading API Shield Operation", err.Error())
		return
	}

	model.ID = flatteners.String(op.ID)
	model.Method = flatteners.String(op.Method)
	model.Host = flatteners.String(op.Host)
	model.Endpoint = NewEndpointValue(op.Endpoint)

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *APIShieldOperationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Update API shield Operation is not supported",
		"Update should never have been called because the resource is configured to be replaced instead",
	)
}

func (r *APIShieldOperationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var model APIShieldOperationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.V1.DeleteAPIShieldOperation(
		ctx,
		cloudflare.ZoneIdentifier(model.ZoneID.ValueString()),
		cloudflare.DeleteAPIShieldOperationParams{
			OperationID: model.ID.ValueString(),
		},
	)

	if err != nil {
		resp.Diagnostics.AddError("Error deleting API Shield Operation", err.Error())
		return
	}
}

func (r *APIShieldOperationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idparts := strings.Split(req.ID, "/")
	if len(idparts) != 2 {
		resp.Diagnostics.AddError("error importing api_shield_operation", `invalid ID specified. Please specify the ID as "<zone_id>/<operation_id>"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("zone_id"), idparts[0],
	)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("id"), idparts[1],
	)...)
}
