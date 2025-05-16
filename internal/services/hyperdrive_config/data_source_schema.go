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
				Description: "Define configurations using a unique string identifier.",
				Computed:    true,
			},
			"hyperdrive_id": schema.StringAttribute{
				Description: "Define configurations using a unique string identifier.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Define configurations using a unique string identifier.",
				Required:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "Defines the creation time of the Hyperdrive configuration.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "Defines the last modified time of the Hyperdrive configuration.",
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
						Description: "Set to true to disable caching of SQL responses. Default is false.",
						Computed:    true,
					},
					"max_age": schema.Int64Attribute{
						Description: "Specify the maximum duration items should persist in the cache. Not returned if set to the default (60).",
						Computed:    true,
					},
					"stale_while_revalidate": schema.Int64Attribute{
						Description: "Specify the number of seconds the cache may serve a stale response. Omitted if set to the default (15).",
						Computed:    true,
					},
				},
			},
			"mtls": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[HyperdriveConfigMTLSDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"ca_certificate_id": schema.StringAttribute{
						Description: "Define CA certificate ID obtained after uploading CA cert.",
						Computed:    true,
					},
					"mtls_certificate_id": schema.StringAttribute{
						Description: "Define mTLS certificate ID obtained after uploading client cert.",
						Computed:    true,
					},
					"sslmode": schema.StringAttribute{
						Description: "Set SSL mode to 'require', 'verify-ca', or 'verify-full' to verify the CA.",
						Computed:    true,
					},
				},
			},
			"origin": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[HyperdriveConfigOriginDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						Description: "Set the name of your origin database.",
						Computed:    true,
					},
					"host": schema.StringAttribute{
						Description: "Defines the host (hostname or IP) of your origin database.",
						Computed:    true,
					},
					"password": schema.StringAttribute{
						Description: "Set the password needed to access your origin database. The API never returns this write-only value.",
						Computed:    true,
						Sensitive:   true,
					},
					"port": schema.Int64Attribute{
						Description: "Defines the port (default: 5432 for Postgres) of your origin database.",
						Computed:    true,
					},
					"scheme": schema.StringAttribute{
						Description: "Specifies the URL scheme used to connect to your origin database.\nAvailable values: \"postgres\", \"postgresql\", \"mysql\".",
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
						Description: "Set the user of your origin database.",
						Computed:    true,
					},
					"access_client_id": schema.StringAttribute{
						Description: "Defines the Client ID of the Access token to use when connecting to the origin database.",
						Computed:    true,
					},
					"access_client_secret": schema.StringAttribute{
						Description: "Defines the Client Secret of the Access Token to use when connecting to the origin database. The API never returns this write-only value.",
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
