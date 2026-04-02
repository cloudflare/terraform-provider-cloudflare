// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connectivity_directory_service

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ConnectivityDirectoryServiceDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"service_id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"app_protocol": schema.StringAttribute{
				Description: `Available values: "postgresql", "mysql".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("postgresql", "mysql"),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"http_port": schema.Int64Attribute{
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"https_port": schema.Int64Attribute{
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"tcp_port": schema.Int64Attribute{
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"type": schema.StringAttribute{
				Description: `Available values: "tcp", "http".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("tcp", "http"),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"host": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ConnectivityDirectoryServiceHostDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"ipv4": schema.StringAttribute{
						Computed: true,
					},
					"network": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[ConnectivityDirectoryServiceHostNetworkDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"tunnel_id": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"ipv6": schema.StringAttribute{
						Computed: true,
					},
					"hostname": schema.StringAttribute{
						Computed: true,
					},
					"resolver_network": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[ConnectivityDirectoryServiceHostResolverNetworkDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"tunnel_id": schema.StringAttribute{
								Computed: true,
							},
							"resolver_ips": schema.ListAttribute{
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"tls_settings": schema.SingleNestedAttribute{
				Description: "TLS settings for a connectivity service.\n\nIf omitted, the default mode (`verify_full`) is used.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ConnectivityDirectoryServiceTLSSettingsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cert_verification_mode": schema.StringAttribute{
						Description: "TLS certificate verification mode for the connection to the origin.\n\n- `\"verify_full\"` — verify certificate chain and hostname (default)\n- `\"verify_ca\"` — verify certificate chain only, skip hostname check\n- `\"disabled\"` — do not verify the server certificate at all",
						Computed:    true,
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description: `Available values: "tcp", "http".`,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("tcp", "http"),
						},
					},
				},
			},
		},
	}
}

func (d *ConnectivityDirectoryServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ConnectivityDirectoryServiceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("service_id"), path.MatchRoot("filter")),
	}
}
