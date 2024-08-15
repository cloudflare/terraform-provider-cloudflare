// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &HyperdriveConfigsDataSource{}

func (d *HyperdriveConfigsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"caching": schema.SingleNestedAttribute{
							Computed: true,
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"disabled": schema.BoolAttribute{
									Description: "When set to true, disables the caching of SQL responses. (Default: false)",
									Computed:    true,
									Optional:    true,
								},
								"max_age": schema.Int64Attribute{
									Description: "When present, specifies max duration for which items should persist in the cache. (Default: 60)",
									Computed:    true,
									Optional:    true,
								},
								"stale_while_revalidate": schema.Int64Attribute{
									Description: "When present, indicates the number of seconds cache may serve the response after it becomes stale. (Default: 15)",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"name": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"origin": schema.SingleNestedAttribute{
							Computed: true,
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"database": schema.StringAttribute{
									Description: "The name of your origin database.",
									Computed:    true,
								},
								"host": schema.StringAttribute{
									Description: "The host (hostname or IP) of your origin database.",
									Computed:    true,
								},
								"scheme": schema.StringAttribute{
									Description: "Specifies the URL scheme used to connect to your origin database.",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"postgres",
											"postgresql",
											"mysql",
										),
									},
								},
								"user": schema.StringAttribute{
									Description: "The user of your origin database.",
									Computed:    true,
								},
								"access_client_id": schema.StringAttribute{
									Description: "The Client ID of the Access token to use when connecting to the origin database",
									Computed:    true,
									Optional:    true,
								},
								"port": schema.Int64Attribute{
									Description: "The port (default: 5432 for Postgres) of your origin database.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *HyperdriveConfigsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
