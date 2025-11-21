// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ManagedTransformsDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The unique ID of the zone.",
				Computed:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The unique ID of the zone.",
				Required:    true,
			},
			"managed_request_headers": schema.ListNestedAttribute{
				Description: "The list of Managed Request Transforms.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ManagedTransformsManagedRequestHeadersDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The human-readable identifier of the Managed Transform.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the Managed Transform is enabled.",
							Computed:    true,
						},
					},
				},
			},
			"managed_response_headers": schema.ListNestedAttribute{
				Description: "The list of Managed Response Transforms.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ManagedTransformsManagedResponseHeadersDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The human-readable identifier of the Managed Transform.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the Managed Transform is enabled.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ManagedTransformsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ManagedTransformsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
