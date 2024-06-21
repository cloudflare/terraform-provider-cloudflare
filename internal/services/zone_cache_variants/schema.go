// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_variants

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r ZoneCacheVariantsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "ID of the zone setting.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("variants"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"value": schema.SingleNestedAttribute{
				Description: "Value of the zone setting.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"avif": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for avif.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"bmp": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for bmp.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"gif": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for gif.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"jp2": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for jp2.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"jpeg": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for jpeg.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"jpg": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for jpg.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"jpg2": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for jpg2.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"png": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for png.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"tif": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for tif.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"tiff": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for tiff.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"webp": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for webp.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}
