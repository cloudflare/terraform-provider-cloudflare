package leaked_credential_check

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource = &LeakedCredentialCheckResource{}
)

func NewResource() resource.Resource {
	return &LeakedCredentialCheckResource{}
}

type LeakedCredentialCheckResource struct {
	client *muxclient.Client
}

func (r *LeakedCredentialCheckResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_leaked_credential_check"
}

func (r *LeakedCredentialCheckResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *LeakedCredentialCheckResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan LeakedCredentialCheckModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zoneID := plan.ZoneID.ValueString()
	createParams := cloudflare.LeakCredentialCheckSetStatusParams{
		Enabled: plan.Enabled.ValueBoolPointer(),
	}

	_, err := r.client.V1.LeakedCredentialCheckSetStatus(ctx, cloudflare.ZoneIdentifier(zoneID), createParams)
	if err != nil {
		resp.Diagnostics.AddError("Error enabling (Create) Leaked Credential Check", err.Error())
		return
	}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *LeakedCredentialCheckResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state LeakedCredentialCheckModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zoneID := state.ZoneID.ValueString()
	status, err := r.client.V1.LeakedCredentialCheckGetStatus(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.LeakedCredentialCheckGetStatusParams{})
	if err != nil {
		resp.Diagnostics.AddError("Error reading Leaked Credential Check status", err.Error())
		return
	}
	state.Enabled = types.BoolPointerValue(status.Enabled)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *LeakedCredentialCheckResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan LeakedCredentialCheckModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zoneID := plan.ZoneID.ValueString()
	updateParams := cloudflare.LeakCredentialCheckSetStatusParams{
		Enabled: plan.Enabled.ValueBoolPointer(),
	}

	_, err := r.client.V1.LeakedCredentialCheckSetStatus(ctx, cloudflare.ZoneIdentifier(zoneID), updateParams)
	if err != nil {
		resp.Diagnostics.AddError("Error updating status of Leaked Credential Check", err.Error())
		return
	}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Delete disables (enabled: false) the Leaked Credential Check functionality
func (r *LeakedCredentialCheckResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state LeakedCredentialCheckModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zoneID := state.ZoneID.ValueString()
	params := cloudflare.LeakCredentialCheckSetStatusParams{
		Enabled: cloudflare.BoolPtr(false),
	}

	_, err := r.client.V1.LeakedCredentialCheckSetStatus(ctx, cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating status of Leaked Credential Check", err.Error())
		return
	}
}

func (r *LeakedCredentialCheckResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// req.ID is the zoneID for which you want to import the state of the
	// Leaked Credential Check feature
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("zone_id"), req.ID)...)
}
