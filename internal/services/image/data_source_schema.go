// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ImageDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Optional:    true,
			},
			"image_id": schema.StringAttribute{
				Description: "Image unique identifier.",
				Optional:    true,
			},
			"filename": schema.StringAttribute{
				Description: "Image file name.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "Image unique identifier.",
				Computed:    true,
			},
			"require_signed_urls": schema.BoolAttribute{
				Description: "Indicates whether the image can be a accessed only using it's UID. If set to true, a signed token needs to be generated with a signing key to view the image.",
				Computed:    true,
			},
			"uploaded": schema.StringAttribute{
				Description: "When the media item was uploaded.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"variants": schema.ListAttribute{
				Description: "Object specifying available variants for an image.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"images": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ImageImagesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Image unique identifier.",
							Computed:    true,
						},
						"filename": schema.StringAttribute{
							Description: "Image file name.",
							Computed:    true,
						},
						"meta": schema.StringAttribute{
							Description: "User modifiable key-value store. Can be used for keeping references to another system of record for managing images. Metadata must not exceed 1024 bytes.",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"require_signed_urls": schema.BoolAttribute{
							Description: "Indicates whether the image can be a accessed only using it's UID. If set to true, a signed token needs to be generated with a signing key to view the image.",
							Computed:    true,
						},
						"uploaded": schema.StringAttribute{
							Description: "When the media item was uploaded.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"variants": schema.ListAttribute{
							Description: "Object specifying available variants for an image.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
					},
				},
			},
			"meta": schema.StringAttribute{
				Description: "User modifiable key-value store. Can be used for keeping references to another system of record for managing images. Metadata must not exceed 1024 bytes.",
				Computed:    true,
				CustomType:  jsontypes.NormalizedType{},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Account identifier tag.",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *ImageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ImageDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("image_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("image_id")),
	}
}
