// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*HyperdriveConfigDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"hyperdrive_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "When the Hyperdrive configuration was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "When the Hyperdrive configuration was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"caching": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[HyperdriveConfigCachingDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"disabled": schema.BoolAttribute{
						Description: "When set to true, disables the caching of SQL responses. (Default: false)",
						Computed:    true,
					},
					"max_age": schema.Int64Attribute{
						Description: "When present, specifies max duration for which items should persist in the cache. Not returned if set to default. (Default: 60)",
						Computed:    true,
					},
					"stale_while_revalidate": schema.Int64Attribute{
						Description: "When present, indicates the number of seconds cache may serve the response after it becomes stale. Not returned if set to default. (Default: 15)",
						Computed:    true,
					},
				},
			},
			"origin": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[HyperdriveConfigOriginDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						Description: "The name of your origin database.",
						Computed:    true,
					},
					"host": schema.StringAttribute{
						Description: "The host (hostname or IP) of your origin database.",
						Computed:    true,
					},
					"password": schema.StringAttribute{
						Description: "The password required to access your origin database. This value is write-only and never returned by the API.",
						Computed:    true,
						Sensitive:   true,
					},
					"port": schema.Int64Attribute{
						Description: "The port (default: 5432 for Postgres) of your origin database.",
						Computed:    true,
					},
					"scheme": schema.StringAttribute{
						Description: "Specifies the URL scheme used to connect to your origin database.\nAvailable values: \"postgres\", \"postgresql\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("postgres", "postgresql"),
						},
					},
					"user": schema.StringAttribute{
						Description: "The user of your origin database.",
						Computed:    true,
					},
					"access_client_id": schema.StringAttribute{
						Description: "The Client ID of the Access token to use when connecting to the origin database.",
						Computed:    true,
					},
					"access_client_secret": schema.StringAttribute{
						Description: "The Client Secret of the Access token to use when connecting to the origin database. This value is write-only and never returned by the API.",
						Computed:    true,
						Sensitive:   true,
					},
				},
			},
		},
	}
}

func (d *HyperdriveConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *HyperdriveConfigDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
