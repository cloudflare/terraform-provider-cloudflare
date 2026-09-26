// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dns_location

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
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

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDNSLocationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Cloudflare Zero Trust Secure DNS Locations Write",
				"Zero Trust Read",
				"Zero Trust Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"location_id": schema.StringAttribute{
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Optional:    true,
			},
			"client_default": schema.BoolAttribute{
				Description: "Indicate whether this location is the default location.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"dns_destination_ips_id": schema.StringAttribute{
				Description: "Indicate the identifier of the pair of IPv4 addresses assigned to this location.",
				Computed:    true,
			},
			"dns_destination_ipv6_block_id": schema.StringAttribute{
				Description: "Specify the UUID of the IPv6 block brought to the gateway so that this location's IPv6 address is allocated from the Bring Your Own IPv6 (BYOIPv6) block rather than the standard Cloudflare IPv6 block.",
				Computed:    true,
			},
			"doh_subdomain": schema.StringAttribute{
				Description: "Specify the DNS over HTTPS domain that receives DNS requests. Gateway automatically generates this value.",
				Computed:    true,
			},
			"ecs_support": schema.BoolAttribute{
				Description: "Indicate whether the location must resolve EDNS queries.",
				Computed:    true,
			},
			"ip": schema.StringAttribute{
				Description: "Defines the automatically generated IPv6 destination IP assigned to this location. Gateway counts all DNS requests sent to this IP as requests under this location.",
				Computed:    true,
			},
			"ipv4_destination": schema.StringAttribute{
				Description: "Show the primary destination IPv4 address from the pair identified dns_destination_ips_id. This field read-only.",
				Computed:    true,
			},
			"ipv4_destination_backup": schema.StringAttribute{
				Description: "Show the backup destination IPv4 address from the pair identified dns_destination_ips_id. This field read-only.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Specify the location name.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"endpoints": schema.SingleNestedAttribute{
				Description: "Configure the destination endpoints for this location.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustDNSLocationEndpointsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"doh": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustDNSLocationEndpointsDOHDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Indicate whether the DOH endpoint is enabled for this location.",
								Computed:    true,
							},
							"networks": schema.ListNestedAttribute{
								Description: "Specify the list of allowed source IP network ranges for this endpoint. When the list is empty, the endpoint allows all source IPs. The list takes effect only if the endpoint is enabled for this location.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectListType[ZeroTrustDNSLocationEndpointsDOHNetworksDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"network": schema.StringAttribute{
											Description: "Specify the IP address or IP CIDR.",
											Computed:    true,
										},
									},
								},
							},
							"require_token": schema.BoolAttribute{
								Description: "Specify whether the DOH endpoint requires user identity authentication.",
								Computed:    true,
							},
						},
					},
					"dot": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustDNSLocationEndpointsDOTDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Indicate whether the DOT endpoint is enabled for this location.",
								Computed:    true,
							},
							"networks": schema.ListNestedAttribute{
								Description: "Specify the list of allowed source IP network ranges for this endpoint. When the list is empty, the endpoint allows all source IPs. The list takes effect only if the endpoint is enabled for this location.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectListType[ZeroTrustDNSLocationEndpointsDOTNetworksDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"network": schema.StringAttribute{
											Description: "Specify the IP address or IP CIDR.",
											Computed:    true,
										},
									},
								},
							},
						},
					},
					"ipv4": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustDNSLocationEndpointsIPV4DataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Indicate whether the IPv4 endpoint is enabled for this location.",
								Computed:    true,
							},
						},
					},
					"ipv6": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustDNSLocationEndpointsIPV6DataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Indicate whether the IPV6 endpoint is enabled for this location.",
								Computed:    true,
							},
							"networks": schema.ListNestedAttribute{
								Description: "Specify the list of allowed source IPv6 network ranges for this endpoint. When the list is empty, the endpoint allows all source IPs. The list takes effect only if the endpoint is enabled for this location.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectListType[ZeroTrustDNSLocationEndpointsIPV6NetworksDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"network": schema.StringAttribute{
											Description: "Specify the IPv6 address or IPv6 CIDR.",
											Computed:    true,
										},
									},
								},
							},
						},
					},
				},
			},
			"max_ttl": schema.SingleNestedAttribute{
				Description: "Controls how DNS response TTLs are capped for this location relative to the account `max_ttl_secs` setting. Omitting `max_ttl` on update resets it to `inherit`.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustDNSLocationMaxTTLDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						Description: "`inherit` uses the account `max_ttl_secs`. `override` uses this location's `ttl_secs`. `disabled` leaves returned TTLs unchanged.\nAvailable values: \"inherit\", \"override\", \"disabled\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"inherit",
								"override",
								"disabled",
							),
						},
					},
					"ttl_secs": schema.Int64Attribute{
						Description: "Location-specific cap on DNS response TTLs, in seconds. Required when `mode` is `override`. Must be omitted when `mode` is `inherit` or `disabled`.",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.Between(60, 36000),
						},
					},
				},
			},
			"networks": schema.ListNestedAttribute{
				Description: "Specify the list of network ranges from which requests at this location originate. The list takes effect only if it is non-empty and the IPv4 endpoint is enabled for this location.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDNSLocationNetworksDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network": schema.StringAttribute{
							Description: "Specify the IPv4 address or IPv4 CIDR. Limit IPv4 CIDRs to a maximum of /24.",
							Computed:    true,
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"direction": schema.StringAttribute{
						Description: "Sort direction. Only takes effect when `order_by` is also provided; it\nis ignored otherwise. When `direction` is omitted the effective\ndirection is field-specific: `created_at` and `updated_at` default to\ndescending (newest first); `name` defaults to ascending.\n  * `asc` — ascending.\n  * `desc` — descending.\nAvailable values: \"asc\", \"desc\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"filter": schema.ListAttribute{
						Description: "Filter the returned locations by one or more `field:value` pairs.\nRepeat the parameter to apply multiple filters; they are combined with\nlogical AND (a location must satisfy every filter to be returned).\n\nSupported fields and their matching behaviour:\n  * `name` — case-insensitive substring match on the location name.\n  * `id` — substring match on the location ID (UUID), with or without dashes.\n  * `is_default` — whether it is the default for the account.\n\nEach entry must match one of the per-field patterns below:\n  * the field must be one of `name`, `id`, or `is_default`;\n  * `name`/`id` accept any value;\n  * `is_default` only accepts `true` or `false`; any other value returns `400`",
						Optional:    true,
						ElementType: types.StringType,
					},
					"order_by": schema.StringAttribute{
						Description: "Field to sort the returned locations by. When omitted, the order of\nresults is unspecified. Supported values:\n  * `name` — sort alphabetically by location name.\n  * `created_at` — sort by creation time; defaults to descending unless `direction` is set.\n  * `updated_at` — sort by last-modified time; defaults to descending unless `direction` is set.\nAvailable values: \"name\", \"created_at\", \"updated_at\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"name",
								"created_at",
								"updated_at",
							),
						},
					},
					"search": schema.StringAttribute{
						Description: "Case-insensitive substring match on the location name. When combined\nwith `filter`, both must match (logical AND).",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDNSLocationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDNSLocationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("location_id"), path.MatchRoot("filter")),
	}
}
