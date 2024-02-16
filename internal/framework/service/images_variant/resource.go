package images_variant

import (
	"context"
	"fmt"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/pkg/errors"
)

var _ resource.Resource = &ImagesVariant{}
var _ resource.ResourceWithImportState = &ImagesVariant{}

func NewResource() resource.Resource {
	return &ImagesVariant{}
}

type ImagesVariant struct {
	client *cloudflare.API
}

func (r *ImagesVariant) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_images_variant"
}

func (r *ImagesVariant) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ImagesVariant) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ImagesVariantModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	variant, err := r.client.CreateImagesVariant(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()), cloudflare.CreateImagesVariantParams{
		ID:                     data.ID.ValueString(),
		NeverRequireSignedURLs: data.NeverRequireSignedUrls.ValueBoolPointer(),
		Options: cloudflare.ImagesVariantsOptions{
			Fit:      data.Options.Fit.ValueString(),
			Metadata: data.Options.Metadata.ValueString(),
			Height:   int(data.Options.Height.ValueInt64()),
			Width:    int(data.Options.Width.ValueInt64()),
		},
	})

	if err != nil {
		resp.Diagnostics.AddError("failed to create images variant rule", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, toImagesVariantModel(data.AccountID.ValueString(), data.ID.ValueString(), variant))...)
}

func (r *ImagesVariant) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ImagesVariantModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	imageVariant, err := getImagesVariant(ctx, r.client, data)
	if imageVariant.ID.ValueString() == "" {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed reading Images Variant with ID: %q", data.ID), err.Error())
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, imageVariant)...)
}

func (r *ImagesVariant) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ImagesVariantModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.UpdateImagesVariant(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()), cloudflare.UpdateImagesVariantParams{
		ID:                     data.ID.ValueString(),
		NeverRequireSignedURLs: data.NeverRequireSignedUrls.ValueBoolPointer(),
		Options: cloudflare.ImagesVariantsOptions{
			Fit:      data.Options.Fit.ValueString(),
			Metadata: data.Options.Metadata.ValueString(),
			Width:    int(data.Options.Width.ValueInt64()),
			Height:   int(data.Options.Height.ValueInt64()),
		},
	})

	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to update Images Variant: %s", data.ID), err.Error())
		return
	}

	imageVariant, err := getImagesVariant(ctx, r.client, data)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed reading Images Variant with ID: %q", data.ID), err.Error())
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, imageVariant)...)
}

func (r *ImagesVariant) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ImagesVariantModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteImagesVariant(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()), data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to delete images variant: %s", data.ID.ValueString()), err.Error())
		return
	}
}

func (r *ImagesVariant) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idparts := strings.Split(req.ID, "/")
	if len(idparts) != 2 {
		resp.Diagnostics.AddError("error importing images variant", `invalid ID specified. Please specify the ID as "<accountID>/<imagesVariantID>"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("account_id"), idparts[0],
	)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("id"), idparts[1],
	)...)
}

func getImagesVariant(ctx context.Context, client *cloudflare.API, data *ImagesVariantModel) (*ImagesVariantModel, error) {
	variantID := data.ID.ValueString()
	accountIdentifier := cloudflare.AccountIdentifier(data.AccountID.ValueString())
	imageVariant, err := client.GetImagesVariant(ctx, accountIdentifier, variantID)

	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Images Variant %s no longer exists", data.ID.ValueString()))
			return &ImagesVariantModel{}, nil
		}
		return &ImagesVariantModel{}, errors.Wrap(err, fmt.Sprintf("error reading Images Variant with ID %q", data.ID.ValueString()))
	}

	return toImagesVariantModel(data.AccountID.ValueString(), data.ID.ValueString(), imageVariant), nil
}

func toImagesVariantModel(accountID string, variantID string, item cloudflare.ImagesVariant) *ImagesVariantModel {
	model := ImagesVariantModel{
		AccountID:              types.StringValue(accountID),
		ID:                     types.StringValue(variantID),
		NeverRequireSignedUrls: types.BoolValue(*item.NeverRequireSignedURLs),
	}

	model.Options = &ImagesVariantOptionsModel{
		Fit:      types.StringValue(item.Options.Fit),
		Metadata: types.StringValue(item.Options.Metadata),
		Width:    types.Int64Value(int64(item.Options.Width)),
		Height:   types.Int64Value(int64(item.Options.Height)),
	}

	return &model
}
