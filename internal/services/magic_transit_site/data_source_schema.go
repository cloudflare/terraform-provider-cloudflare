// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicTransitSiteDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"site_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"connector_id": schema.StringAttribute{
				Description: "Magic Connector identifier tag.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"ha_mode": schema.BoolAttribute{
				Description: "Site high availability mode. If set to true, the site can have two connectors and runs in high availability mode.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the site.",
				Computed:    true,
			},
			"secondary_connector_id": schema.StringAttribute{
				Description: "Magic Connector identifier tag. Used when high availability mode is on.",
				Computed:    true,
			},
			"location": schema.SingleNestedAttribute{
				Description: "Location of site in latitude and longitude.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[MagicTransitSiteLocationDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"lat": schema.StringAttribute{
						Description: "Latitude",
						Computed:    true,
					},
					"lon": schema.StringAttribute{
						Description: "Longitude",
						Computed:    true,
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"connector_identifier": schema.StringAttribute{
						Description: "Identifier",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *MagicTransitSiteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MagicTransitSiteDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("site_id"), path.MatchRoot("filter")),
	}
}
