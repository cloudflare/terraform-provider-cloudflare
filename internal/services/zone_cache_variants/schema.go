// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_variants

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZoneCacheVariantsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"value": schema.SingleNestedAttribute{
				Description: "Value of the zone setting.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"avif": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for avif.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"bmp": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for bmp.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"gif": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for gif.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"jp2": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for jp2.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"jpeg": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for jpeg.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"jpg": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for jpg.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"jpg2": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for jpg2.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"png": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for png.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"tif": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for tif.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"tiff": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for tiff.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"webp": schema.ListAttribute{
						Description: "List of strings with the MIME types of all the variants that should be served for webp.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
				},
			},
			"modified_on": schema.StringAttribute{
				Description: "last time this setting was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ZoneCacheVariantsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZoneCacheVariantsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
