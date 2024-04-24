package workers_for_platforms

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
var _ resource.Resource = &WorkersForPlatformsResource{}
var _ resource.ResourceWithImportState = &WorkersForPlatformsResource{}

func NewResource() resource.Resource {
	return &WorkersForPlatformsResource{}
}

// WorkersForPlatformsResource defines the resource implementation.
type WorkersForPlatformsResource struct {
	client *muxclient.Client
}

func (r *WorkersForPlatformsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workers_for_platforms_namespace"
}

func (r *WorkersForPlatformsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *WorkersForPlatformsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *WorkersForPlatformsNamespaceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	namespace, err := r.client.V1.CreateWorkersForPlatformsDispatchNamespace(ctx, cfv1.AccountIdentifier(data.AccountID.ValueString()),
		cfv1.CreateWorkersForPlatformsDispatchNamespaceParams{
			Name: data.Name.ValueString(),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create Workers for Platforms namespace", err.Error())
		return
	}
	data.ID = types.StringValue(namespace.Result.NamespaceName)
	data.Name = types.StringValue(namespace.Result.NamespaceName)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersForPlatformsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *WorkersForPlatformsNamespaceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	namespace, err := r.client.V1.GetWorkersForPlatformsDispatchNamespace(ctx, cfv1.AccountIdentifier(data.AccountID.ValueString()), data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed reading Workers for Platforms namespace", err.Error())
		return
	}
	data.ID = types.StringValue(namespace.Result.NamespaceName)
	data.Name = types.StringValue(namespace.Result.NamespaceName)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersForPlatformsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *WorkersForPlatformsNamespaceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddError("failed to update Workers for Platforms namespace", "Not implemented")
}

func (r *WorkersForPlatformsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *WorkersForPlatformsNamespaceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.V1.DeleteWorkersForPlatformsDispatchNamespace(ctx, cfv1.AccountIdentifier(data.AccountID.ValueString()), data.Name.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("failed to delete Workers for Platforms namespace", err.Error())
		return
	}
}

func (r *WorkersForPlatformsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idparts := strings.Split(req.ID, "/")
	if len(idparts) != 2 {
		resp.Diagnostics.AddError("error importing Workers for Platforms namespace", "invalid ID specified. Please specify the ID as \"account_id/name\"")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("account_id"), idparts[0],
	)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("id"), idparts[1],
	)...)
}
