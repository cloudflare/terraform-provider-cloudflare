// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname_ipfs_universal_path_content_list_entry

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/web3"
	"github.com/cloudflare/cloudflare-terraform/internal/apijson"
	"github.com/cloudflare/cloudflare-terraform/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Web3HostnameIPFSUniversalPathContentListEntryResource{}

func NewResource() resource.Resource {
	return &Web3HostnameIPFSUniversalPathContentListEntryResource{}
}

// Web3HostnameIPFSUniversalPathContentListEntryResource defines the resource implementation.
type Web3HostnameIPFSUniversalPathContentListEntryResource struct {
	client *cloudflare.Client
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web3_hostname_ipfs_universal_path_content_list_entry"
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *Web3HostnameIPFSUniversalPathContentListEntryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *Web3HostnameIPFSUniversalPathContentListEntryModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := Web3HostnameIPFSUniversalPathContentListEntryResultEnvelope{*data}
	_, err = r.client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Entries.New(
		ctx,
		data.ZoneIdentifier.ValueString(),
		data.Identifier.ValueString(),
		web3.HostnameIPFSUniversalPathContentListEntryNewParams{},
		option.WithRequestBody("application/json", dataBytes),
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
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *Web3HostnameIPFSUniversalPathContentListEntryModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := Web3HostnameIPFSUniversalPathContentListEntryResultEnvelope{*data}
	_, err := r.client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Entries.Get(
		ctx,
		data.ZoneIdentifier.ValueString(),
		data.Identifier.ValueString(),
		data.ID.ValueString(),
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
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *Web3HostnameIPFSUniversalPathContentListEntryModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := Web3HostnameIPFSUniversalPathContentListEntryResultEnvelope{*data}
	_, err = r.client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Entries.Update(
		ctx,
		data.ZoneIdentifier.ValueString(),
		data.Identifier.ValueString(),
		data.ID.ValueString(),
		web3.HostnameIPFSUniversalPathContentListEntryUpdateParams{},
		option.WithRequestBody("application/json", dataBytes),
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
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *Web3HostnameIPFSUniversalPathContentListEntryModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Entries.Delete(
		ctx,
		data.ZoneIdentifier.ValueString(),
		data.Identifier.ValueString(),
		data.ID.ValueString(),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
