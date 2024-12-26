package snippets

import (
	"context"
	"fmt"
	"strings"

	cfv1 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SnippetResource{}
var _ resource.ResourceWithImportState = &SnippetResource{}

// SnippetResource defines the resource implementation.
type SnippetResource struct {
	client *muxclient.Client
}

func NewResource() resource.Resource {
	return &SnippetResource{}
}

func (r *SnippetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snippet"
}

func (r *SnippetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func resourceFromAPIResponse(data *Snippet, snippet cfv1.Snippet) {
	data.SnippetName = types.StringValue(snippet.SnippetName)
}

func requestFromResource(data *Snippet) cfv1.SnippetRequest {
	files := make([]cfv1.SnippetFile, 0, len(data.SnippetFile))
	for _, file := range data.SnippetFile {
		files = append(files, cfv1.SnippetFile{
			FileName: file.FileName.ValueString(),
			Content:  file.Content.ValueString(),
		})
	}
	return cfv1.SnippetRequest{
		SnippetName: data.SnippetName.ValueString(),
		MainFile:    data.MainModule.ValueString(),
		Files:       files,
	}
}

func (r *SnippetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *Snippet

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	snippet, err := r.client.V1.UpdateZoneSnippet(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()),
		requestFromResource(data),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create snippet", err.Error())
		return
	}

	resourceFromAPIResponse(data, *snippet)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnippetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *Snippet

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	snippet, err := r.client.V1.GetZoneSnippet(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()), data.SnippetName.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed reading snippet", err.Error())
		return
	}
	resourceFromAPIResponse(data, *snippet)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnippetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *Snippet

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	snippet, err := r.client.V1.UpdateZoneSnippet(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()),
		requestFromResource(data),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create snippet", err.Error())
		return
	}

	resourceFromAPIResponse(data, *snippet)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnippetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *Snippet

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.V1.DeleteZoneSnippet(ctx, cfv1.ZoneIdentifier(data.ZoneID.ValueString()), data.SnippetName.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("failed to delete snippet", err.Error())
		return
	}
}

func (r *SnippetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idparts := strings.Split(req.ID, "/")
	if len(idparts) != 2 {
		resp.Diagnostics.AddError("error importing snippet", `invalid ID specified. Please specify the ID as "<zone_id>/<snippet_id>"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("zone_id"), idparts[0],
	)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("id"), idparts[1],
	)...)
}
