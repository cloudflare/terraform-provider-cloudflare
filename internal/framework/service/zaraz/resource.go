package zaraz

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ZarazConfigResource{}
var _ resource.ResourceWithImportState = &ZarazConfigResource{}

func NewResource() resource.Resource {
	return &ZarazConfigResource{}
}

// ZarazConfigResource defines the resource implementation.
type ZarazConfigResource struct {
	client *cloudflare.API
}

func (r *ZarazConfigResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zaraz_config"
}

func (r *ZarazConfigResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.API)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.API, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ZarazConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	log.Printf("In create")
	var data *ZarazConfigModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	rc := cloudflare.ZoneIdentifier(data.ZoneID.ValueString())
	response, err := r.client.UpdateZarazConfig(ctx, rc, cloudflare.UpdateZarazConfigParams{
		DebugKey: data.Config.DebugKey.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to create Zaraz config", err.Error())
		return
	}
	data.ZoneID = types.StringValue(data.ZoneID.String())
	data.Config.DebugKey = types.StringValue(response.Result.DebugKey)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZarazConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("In read")

	var data *ZarazConfigModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	rc := cloudflare.ZoneIdentifier(data.ZoneID.ValueString())

	response, err := r.client.GetZarazConfig(ctx, rc)
	if err != nil {
		resp.Diagnostics.AddError("failed reading D1 database", err.Error())
		return
	}
	data.ZoneID = types.StringValue(data.ZoneID.String())
	data.Config.DebugKey = types.StringValue(response.Result.DebugKey)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZarazConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZarazConfigModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddError("failed to update Zaraz Config", "Not implemented")
}

// Todo: Can we implment Zaraz TF without delete?
// If no, then what does delete for zaraz config mean in the terraform context?
// Does deleting a config mean, resetting the config?
func (r *ZarazConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var data *ZarazConfigModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	rc := cloudflare.ZoneIdentifier(data.ZoneID.ValueString())
	response, err := r.client.UpdateZarazConfig(ctx, rc, cloudflare.UpdateZarazConfigParams{
		DebugKey: "123",
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to delete Zaraz config", err.Error())
		return
	}
	data.ZoneID = types.StringValue(data.ZoneID.String())
	data.Config.DebugKey = types.StringValue(response.Result.DebugKey)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZarazConfigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, "/")

	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"invalid import identifier",
			fmt.Sprintf("expected import identifier to be resourceLevel/resourceIdentifier/rulesetID. got: %q", req.ID),
		)
		return
	}
	resourceLevel, resourceIdentifier, rulesetID := idParts[0], idParts[1], idParts[2]

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), rulesetID)...)
	if resourceLevel == "zone" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("zone_id"), resourceIdentifier)...)
	}
}
