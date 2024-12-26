package content_scanning

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
	_ resource.Resource = &ContentScanningResource{}
)

func NewResource() resource.Resource {
	return &ContentScanningResource{}
}

type ContentScanningResource struct {
	client *muxclient.Client
}

func (r *ContentScanningResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_content_scanning"
}

func (r *ContentScanningResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ContentScanningResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ContentScanningModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	zoneID := cloudflare.ZoneIdentifier(plan.ZoneID.ValueString())
	if plan.Enabled.ValueBool() {
		params := cloudflare.ContentScanningEnableParams{}
		_, err := r.client.V1.ContentScanningEnable(ctx, zoneID, params)
		if err != nil {
			resp.Diagnostics.AddError("Error enabling (Create) Content Scanning", err.Error())
			return
		}
	} else {
		params := cloudflare.ContentScanningDisableParams{}
		_, err := r.client.V1.ContentScanningDisable(ctx, zoneID, params)
		if err != nil {
			resp.Diagnostics.AddError("Error disabling (Create) Content Scanning", err.Error())
			return
		}
	}
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ContentScanningResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ContentScanningModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zoneID := state.ZoneID.ValueString()
	status, err := r.client.V1.ContentScanningStatus(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.ContentScanningStatusParams{})
	if err != nil {
		resp.Diagnostics.AddError("Error reading Content Scanning status", err.Error())
		return
	}
	switch status.Result.Value {
	case "enabled":
		state.Enabled = types.BoolValue(true)
	case "disabled":
		state.Enabled = types.BoolValue(false)
	default:
		resp.Diagnostics.AddError("Unrecognized state", fmt.Sprintf("Unrecognized state = %s for Content Scanning", status.Result.Value))
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *ContentScanningResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ContentScanningModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zoneID := cloudflare.ZoneIdentifier(plan.ZoneID.ValueString())
	if plan.Enabled.ValueBool() {
		params := cloudflare.ContentScanningEnableParams{}
		_, err := r.client.V1.ContentScanningEnable(ctx, zoneID, params)
		if err != nil {
			resp.Diagnostics.AddError("Error enabling (Update) Content Scanning", err.Error())
			return
		}
	} else {
		params := cloudflare.ContentScanningDisableParams{}
		_, err := r.client.V1.ContentScanningDisable(ctx, zoneID, params)
		if err != nil {
			resp.Diagnostics.AddError("Error disabling (Update) Content Scanning", err.Error())
			return
		}
	}
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Delete disables the Content Scanning feature.
func (r *ContentScanningResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ContentScanningModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zoneID := cloudflare.ZoneIdentifier(state.ZoneID.ValueString())
	params := cloudflare.ContentScanningDisableParams{}
	_, err := r.client.V1.ContentScanningDisable(ctx, zoneID, params)
	if err != nil {
		resp.Diagnostics.AddError("Error disabling (Update) Content Scanning", err.Error())
		return
	}
}

func (r *ContentScanningResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// req.ID is the zoneID for which you want to import the state of the Content Scanning feature
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("zone_id"), req.ID)...)
}
