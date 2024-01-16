package email_routing_address

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &EmailRoutingAddressResource{}
var _ resource.ResourceWithImportState = &EmailRoutingAddressResource{}

func NewResource() resource.Resource {
	return &EmailRoutingAddressResource{}
}

// EmailRoutingAddressResource defines the resource implementation.
type EmailRoutingAddressResource struct {
	client *cloudflare.API
}

func (r *EmailRoutingAddressResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_email_routing_address"
}

func (r *EmailRoutingAddressResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *EmailRoutingAddressResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *EmailRoutingAddressModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	emailAddress, err := r.client.CreateEmailRoutingDestinationAddress(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()),
		cloudflare.CreateEmailRoutingAddressParameters{
			Email: data.Email.ValueString(),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create email routing destination address", err.Error())
		return
	}
	data = buildEmailRoutingAddressModel(data.AccountID, emailAddress)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EmailRoutingAddressResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *EmailRoutingAddressModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	emailAddress, err := r.client.GetEmailRoutingDestinationAddress(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()), data.Tag.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed reading email routing destination address", err.Error())
		return
	}
	data = buildEmailRoutingAddressModel(data.AccountID, emailAddress)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EmailRoutingAddressResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *EmailRoutingAddressModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddError("failed to update email routing destination address", "Not implemented")
}

func (r *EmailRoutingAddressResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *EmailRoutingAddressModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.DeleteEmailRoutingDestinationAddress(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()), data.Tag.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("failed to delete email routing destination address", err.Error())
		return
	}
}

func buildEmailRoutingAddressModel(accountID types.String, address cloudflare.EmailRoutingDestinationAddress) *EmailRoutingAddressModel {
	var verifiedTime string
	verified := address.Verified
	if verified == nil {
		verifiedTime = "null"
	} else {
		verifiedTime = verified.String()
	}
	return &EmailRoutingAddressModel{
		AccountID: accountID,
		ID:        types.StringValue(address.Tag),
		Tag:       types.StringValue(address.Tag),
		Email:     types.StringValue(address.Email),
		Verified:  types.StringValue(verifiedTime),
		Created:   types.StringValue(address.Created.String()),
		Modified:  types.StringValue(address.Modified.String()),
	}
}

func (r *EmailRoutingAddressResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idparts := strings.Split(req.ID, "/")
	if len(idparts) != 2 {
		resp.Diagnostics.AddError("error importing email routing destination address", `invalid ID specified. Please specify the ID as "<account_id>/<email_routing_address_id>"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("account_id"), idparts[0],
	)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("tag"), idparts[1],
	)...)
}
