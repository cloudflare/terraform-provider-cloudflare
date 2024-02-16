package images_variant

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r *ImagesVariant) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Specify variants that allow you to resize images for different use cases.",
		Version:             1,
		Attributes: map[string]schema.Attribute{
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			consts.IDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.IDSchemaDescription,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"never_require_signed_urls": schema.BoolAttribute{
				MarkdownDescription: "Indicates whether the variant can access an image without a signature, regardless of image access control.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"options": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"metadata": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "What EXIF data should be preserved in the output image.",
						Validators:          []validator.String{stringvalidator.OneOf("keep", "copyright", "none")},
					},
					"fit": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "The fit property describes how the width and height dimensions should be interpreted.",
						Validators:          []validator.String{stringvalidator.OneOf("scale-down", "contain", "cover", "crop", "pad")},
					},
					"height": schema.Int64Attribute{
						Required:            true,
						MarkdownDescription: "Maximum height in image pixels.",
						Validators:          []validator.Int64{int64validator.AtLeast(1)},
					},
					"width": schema.Int64Attribute{
						Required:            true,
						MarkdownDescription: "Maximum width in image pixels.",
						Validators:          []validator.Int64{int64validator.AtLeast(1)},
					},
				},
			},
		},
	}
}
