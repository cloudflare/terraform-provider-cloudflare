package infrastructure_access_target

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/flatteners"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	tftypes "github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &InfrastructureAccessTargetResource{}

func NewResource() resource.Resource {
	return &InfrastructureAccessTargetResource{}
}

// InfrastructureAccessTargetResource defines the resource implementation.
type InfrastructureAccessTargetResource struct {
	client *muxclient.Client
}

func (r *InfrastructureAccessTargetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infrastructure_access_target"
}

func (r *InfrastructureAccessTargetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*muxclient.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *muxclient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *InfrastructureAccessTargetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *InfrastructureAccessTargetModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountId := data.AccountID.ValueString()
	if accountId == "" {
		resp.Diagnostics.AddError("failed to create infrastructure access target", "account id cannot be an empty string")
		return
	}
	createTargetParams := cloudflare.CreateInfrastructureAccessTargetParams{
		InfrastructureAccessTargetParams: cloudflare.InfrastructureAccessTargetParams{
			Hostname: data.Hostname.String(),
			IP:       data.IP,
		},
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Infrastructure Target from struct %+v", createTargetParams))
	target, err := r.client.V1.CreateInfrastructureAccessTarget(ctx, cloudflare.AccountIdentifier(accountId), createTargetParams)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error creating Infrastructure Target for account %q", accountId), err.Error())
		return
	}

	data = buildTargetModelFromResponse(data.AccountID, target)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InfrastructureAccessTargetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *InfrastructureAccessTargetModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Retrieving Cloudflare Infrastructure Target with ID %s", data.ID))
	target, err := r.client.V1.GetInfrastructureAccessTarget(ctx, cloudflare.AccountIdentifier(data.AccountID.String()), data.ID.String())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			resp.Diagnostics.AddError(fmt.Sprintf("Infrastructure Target with ID %s does not exist", data.ID), err.Error())
			return
		}
		resp.Diagnostics.AddError(fmt.Sprintf("error finding Infrastructure Target with ID %s", data.ID), err.Error())
		return
	}

	data = buildTargetModelFromResponse(data.AccountID, target)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InfrastructureAccessTargetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *InfrastructureAccessTargetModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountId := data.AccountID.ValueString()
	if accountId == "" {
		resp.Diagnostics.AddError("failed to update infrastructure access target", "account id cannot be an empty string")
		return
	}
	updatedTargetParams := cloudflare.UpdateInfrastructureAccessTargetParams{
		ID: data.ID.String(),
		ModifyParams: cloudflare.InfrastructureAccessTargetParams{
			Hostname: data.Hostname.String(),
			IP:       data.IP,
		},
	}

	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Infrastructure Target from struct: %+v", updatedTargetParams))
	updatedTarget, err := r.client.V1.UpdateInfrastructureAccessTarget(ctx, cloudflare.AccountIdentifier(accountId), updatedTargetParams)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error updating Infrastructure Target with ID %s for account %q", data.ID, accountId), err.Error())
		return
	}

	data = buildTargetModelFromResponse(data.AccountID, updatedTarget)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InfrastructureAccessTargetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *InfrastructureAccessTargetModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Infrastructure Target with ID: %s", data.ID))
	err := r.client.V1.DeleteInfrastructureAccessTarget(ctx, cloudflare.AccountIdentifier(data.AccountID.String()), data.ID.String())
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error deleting Infrastructure Target with ID %s for account %q", data.ID, data.AccountID.String()), err.Error())
		return
	}
}

func buildTargetModelFromResponse(accountID tftypes.String, target cloudflare.Target) *InfrastructureAccessTargetModel {
	built := InfrastructureAccessTargetModel{
		AccountID:  accountID,
		Hostname:   flatteners.String(target.Hostname),
		ID:         flatteners.String(target.ID),
		IP:         target.IP,
		CreatedAt:  flatteners.String(target.CreatedAt),
		ModifiedAt: flatteners.String(target.ModifiedAt),
	}
	return &built
}
