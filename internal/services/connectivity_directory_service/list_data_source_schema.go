// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connectivity_directory_service

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ConnectivityDirectoryServicesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: `Available values: "tcp", "http".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("tcp", "http"),
				},
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ConnectivityDirectoryServicesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ConnectivityDirectoryServicesHostDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ipv4": schema.StringAttribute{
									Computed: true,
								},
								"network": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[ConnectivityDirectoryServicesHostNetworkDataSourceModel](ctx),
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
									CustomType: customfield.NewNestedObjectType[ConnectivityDirectoryServicesHostResolverNetworkDataSourceModel](ctx),
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
						"name": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Description: `Available values: "tcp", "http".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("tcp", "http"),
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
						"service_id": schema.StringAttribute{
							Computed: true,
						},
						"tls_settings": schema.SingleNestedAttribute{
							Description: "TLS settings for a connectivity service.\n\nIf omitted, the default mode (`verify_full`) is used.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ConnectivityDirectoryServicesTLSSettingsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"cert_verification_mode": schema.StringAttribute{
									Description: "TLS certificate verification mode for the connection to the origin.\n\n- `\"verify_full\"` — verify certificate chain and hostname (default)\n- `\"verify_ca\"` — verify certificate chain only, skip hostname check\n- `\"disabled\"` — do not verify the server certificate at all",
									Computed:    true,
								},
							},
						},
						"updated_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"app_protocol": schema.StringAttribute{
							Description: `Available values: "postgresql", "mysql".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("postgresql", "mysql"),
							},
						},
						"tcp_port": schema.Int64Attribute{
							Computed: true,
							Validators: []validator.Int64{
								int64validator.AtLeast(1),
							},
						},
					},
				},
			},
		},
	}
}

func (d *ConnectivityDirectoryServicesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ConnectivityDirectoryServicesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
