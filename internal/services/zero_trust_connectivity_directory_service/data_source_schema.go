// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_connectivity_directory_service

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

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustConnectivityDirectoryServiceDataSource)(nil)

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
			"type": schema.StringAttribute{
				Description: `Available values: "http".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("http"),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"host": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustConnectivityDirectoryServiceHostDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"ipv4": schema.StringAttribute{
						Computed: true,
					},
					"network": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustConnectivityDirectoryServiceHostNetworkDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[ZeroTrustConnectivityDirectoryServiceHostResolverNetworkDataSourceModel](ctx),
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description: `Available values: "http".`,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("http"),
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustConnectivityDirectoryServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustConnectivityDirectoryServiceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("service_id"), path.MatchRoot("filter")),
	}
}
