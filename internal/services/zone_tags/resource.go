// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_tags

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zones"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ZoneTagsResource)(nil)
var _ resource.ResourceWithImportState = (*ZoneTagsResource)(nil)

func NewResource() resource.Resource {
	return &ZoneTagsResource{}
}

// ZoneTagsResource defines the resource implementation.
type ZoneTagsResource struct {
	client *cloudflare.Client
}

func (r *ZoneTagsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zone_tags"
}

func (r *ZoneTagsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZoneTagsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZoneTagsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tagMap := make(map[string]string)
	data.Tags.ElementsAs(ctx, &tagMap, false)

	// Call imaginary cloudflare-go function
	_, err := r.client.Zones.UpdateTags(ctx, zones.ZoneUpdateTagsParams{
		ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		Tags:   tagMap,
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to create zone tags", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneTagsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZoneTagsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Call imaginary cloudflare-go function
	tagsResp, err := r.client.Zones.GetTags(ctx, zones.ZoneGetTagsParams{
		ZoneID: cloudflare.F(data.ZoneID.ValueString()),
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to read zone tags", err.Error())
		return
	}

	if len(tagsResp.Tags) > 0 {
		data.Tags, _ = types.MapValueFrom(ctx, types.StringType, tagsResp.Tags)
	} else {
		data.Tags = types.MapNull(types.StringType)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneTagsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZoneTagsModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tagMap := make(map[string]string)
	data.Tags.ElementsAs(ctx, &tagMap, false)

	// Call imaginary cloudflare-go function
	_, err := r.client.Zones.UpdateTags(ctx, zones.ZoneUpdateTagsParams{
		ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		Tags:   tagMap,
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to update zone tags", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZoneTagsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ZoneTagsModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete all tags by setting empty map
	_, err := r.client.Zones.UpdateTags(ctx, zones.ZoneUpdateTagsParams{
		ZoneID: cloudflare.F(data.ZoneID.ValueString()),
		Tags:   make(map[string]string),
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to delete zone tags", err.Error())
		return
	}
}

func (r *ZoneTagsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *ZoneTagsModel = new(ZoneTagsModel)

	path := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<zone_id>",
		&path,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ZoneID = types.StringValue(path)

	// Read the current tags to populate state
	tagsResp, err := r.client.Zones.GetTags(ctx, zones.ZoneGetTagsParams{
		ZoneID: cloudflare.F(path),
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to read zone tags during import", err.Error())
		return
	}

	if len(tagsResp.Tags) > 0 {
		data.Tags, _ = types.MapValueFrom(ctx, types.StringType, tagsResp.Tags)
	} else {
		data.Tags = types.MapNull(types.StringType)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}