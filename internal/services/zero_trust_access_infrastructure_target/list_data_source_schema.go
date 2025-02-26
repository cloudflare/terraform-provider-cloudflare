// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_infrastructure_target

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

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessInfrastructureTargetsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier",
				Required:    true,
			},
			"created_after": schema.StringAttribute{
				Description: "Date and time at which the target was created after (inclusive)",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"created_before": schema.StringAttribute{
				Description: "Date and time at which the target was created before (inclusive)",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"direction": schema.StringAttribute{
				Description: "The sorting direction.\navailable values: \"asc\", \"desc\"",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"hostname": schema.StringAttribute{
				Description: "Hostname of a target",
				Optional:    true,
			},
			"hostname_contains": schema.StringAttribute{
				Description: "Partial match to the hostname of a target",
				Optional:    true,
			},
			"ip_like": schema.StringAttribute{
				Description: "Filters for targets whose IP addresses look like the specified string.\nSupports `*` as a wildcard character",
				Optional:    true,
			},
			"ip_v4": schema.StringAttribute{
				Description: "IPv4 address of the target",
				Optional:    true,
			},
			"ip_v6": schema.StringAttribute{
				Description: "IPv6 address of the target",
				Optional:    true,
			},
			"ipv4_end": schema.StringAttribute{
				Description: "Defines an IPv4 filter range's ending value (inclusive). Requires\n`ipv4_start` to be specified as well.",
				Optional:    true,
			},
			"ipv4_start": schema.StringAttribute{
				Description: "Defines an IPv4 filter range's starting value (inclusive). Requires\n`ipv4_end` to be specified as well.",
				Optional:    true,
			},
			"ipv6_end": schema.StringAttribute{
				Description: "Defines an IPv6 filter range's ending value (inclusive). Requires\n`ipv6_start` to be specified as well.",
				Optional:    true,
			},
			"ipv6_start": schema.StringAttribute{
				Description: "Defines an IPv6 filter range's starting value (inclusive). Requires\n`ipv6_end` to be specified as well.",
				Optional:    true,
			},
			"modified_after": schema.StringAttribute{
				Description: "Date and time at which the target was modified after (inclusive)",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_before": schema.StringAttribute{
				Description: "Date and time at which the target was modified before (inclusive)",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"order": schema.StringAttribute{
				Description: "The field to sort by.\navailable values: \"hostname\", \"created_at\"",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("hostname", "created_at"),
				},
			},
			"virtual_network_id": schema.StringAttribute{
				Description: "Private virtual network identifier of the target",
				Optional:    true,
			},
			"ips": schema.ListAttribute{
				Description: "Filters for targets that have any of the following IP addresses. Specify\n`ips` multiple times in query parameter to build list of candidates.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"target_ids": schema.ListAttribute{
				Description: "Filters for targets that have any of the following UUIDs. Specify\n`target_ids` multiple times in query parameter to build list of\ncandidates.",
				Optional:    true,
				ElementType: types.StringType,
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessInfrastructureTargetsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Target identifier",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "Date and time at which the target was created",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"hostname": schema.StringAttribute{
							Description: "A non-unique field that refers to a target",
							Computed:    true,
						},
						"ip": schema.SingleNestedAttribute{
							Description: "The IPv4/IPv6 address that identifies where to reach a target",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessInfrastructureTargetsIPDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ipv4": schema.SingleNestedAttribute{
									Description: "The target's IPv4 address",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessInfrastructureTargetsIPIPV4DataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"ip_addr": schema.StringAttribute{
											Description: "IP address of the target",
											Computed:    true,
										},
										"virtual_network_id": schema.StringAttribute{
											Description: "(optional) Private virtual network identifier for the target. If omitted, the default virtual network ID will be used.",
											Computed:    true,
										},
									},
								},
								"ipv6": schema.SingleNestedAttribute{
									Description: "The target's IPv6 address",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessInfrastructureTargetsIPIPV6DataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"ip_addr": schema.StringAttribute{
											Description: "IP address of the target",
											Computed:    true,
										},
										"virtual_network_id": schema.StringAttribute{
											Description: "(optional) Private virtual network identifier for the target. If omitted, the default virtual network ID will be used.",
											Computed:    true,
										},
									},
								},
							},
						},
						"modified_at": schema.StringAttribute{
							Description: "Date and time at which the target was modified",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustAccessInfrastructureTargetsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustAccessInfrastructureTargetsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
