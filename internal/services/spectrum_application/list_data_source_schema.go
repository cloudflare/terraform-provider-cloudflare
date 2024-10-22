// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customvalidator"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ datasource.DataSourceWithConfigValidators = (*SpectrumApplicationsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Zone identifier.",
				Required:    true,
			},
			"direction": schema.StringAttribute{
				Description: "Sets the direction by which results are ordered.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"order": schema.StringAttribute{
				Description: "Application field by which results are ordered.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"protocol",
						"app_id",
						"created_on",
						"modified_on",
						"dns",
					),
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
				CustomType:  customfield.NewNestedObjectListType[SpectrumApplicationsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"errors": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[SpectrumApplicationsErrorsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"code": schema.Int64Attribute{
										Computed: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1000),
										},
									},
									"message": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"messages": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[SpectrumApplicationsMessagesDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"code": schema.Int64Attribute{
										Computed: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1000),
										},
									},
									"message": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"success": schema.BoolAttribute{
							Description: "Whether the API call was successful",
							Computed:    true,
						},
						"result": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[SpectrumApplicationsResultDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "App identifier.",
										Computed:    true,
									},
									"created_on": schema.StringAttribute{
										Description: "When the Application was created.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"dns": schema.SingleNestedAttribute{
										Description: "The name and type of DNS record for the Spectrum application.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[SpectrumApplicationsResultDNSDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the DNS record associated with the application.",
												Computed:    true,
											},
											"type": schema.StringAttribute{
												Description: "The type of DNS record associated with the application.",
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("CNAME", "ADDRESS"),
												},
											},
										},
									},
									"ip_firewall": schema.BoolAttribute{
										Description: "Enables IP Access Rules for this application.\nNotes: Only available for TCP applications.",
										Computed:    true,
									},
									"modified_on": schema.StringAttribute{
										Description: "When the Application was last modified.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"protocol": schema.StringAttribute{
										Description: "The port configuration at Cloudflare's edge. May specify a single port, for example `\"tcp/1000\"`, or a range of ports, for example `\"tcp/1000-2000\"`.",
										Computed:    true,
									},
									"proxy_protocol": schema.StringAttribute{
										Description: "Enables Proxy Protocol to the origin. Refer to [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/) for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple Proxy Protocol.",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"off",
												"v1",
												"v2",
												"simple",
											),
										},
									},
									"tls": schema.StringAttribute{
										Description: "The type of TLS termination associated with the application.",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"off",
												"flexible",
												"full",
												"strict",
											),
										},
									},
									"traffic_type": schema.StringAttribute{
										Description: "Determines how data travels from the edge to your origin. When set to \"direct\", Spectrum will send traffic directly to your origin, and the application's type is derived from the `protocol`. When set to \"http\" or \"https\", Spectrum will apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and the application type matches this property exactly.",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"direct",
												"http",
												"https",
											),
										},
									},
									"argo_smart_routing": schema.BoolAttribute{
										Description: "Enables Argo Smart Routing for this application.\nNotes: Only available for TCP applications with traffic_type set to \"direct\".",
										Computed:    true,
									},
									"edge_ips": schema.SingleNestedAttribute{
										Description: "The anycast edge IP configuration for the hostname of this application.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[SpectrumApplicationsResultEdgeIPsDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"connectivity": schema.StringAttribute{
												Description: "The IP versions supported for inbound connections on Spectrum anycast IPs.",
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive(
														"all",
														"ipv4",
														"ipv6",
													),
												},
											},
											"type": schema.StringAttribute{
												Description: "The type of edge IP configuration specified. Dynamically allocated edge IPs use Spectrum anycast IPs in accordance with the connectivity you specify. Only valid with CNAME DNS names.",
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("dynamic", "static"),
												},
											},
											"ips": schema.ListAttribute{
												Description: "The array of customer owned IPs we broadcast via anycast for this hostname and application.",
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
										},
									},
									"origin_direct": schema.ListAttribute{
										Description: "List of origin IP addresses. Array may contain multiple IP addresses for load balancing.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"origin_dns": schema.SingleNestedAttribute{
										Description: "The name and type of DNS record for the Spectrum application.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[SpectrumApplicationsResultOriginDNSDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the DNS record associated with the origin.",
												Computed:    true,
											},
											"ttl": schema.Int64Attribute{
												Description: "The TTL of our resolution of your DNS record in seconds.",
												Computed:    true,
												Validators: []validator.Int64{
													int64validator.AtLeast(600),
												},
											},
											"type": schema.StringAttribute{
												Description: "The type of DNS record associated with the origin. \"\" is used to specify a combination of A/AAAA records.",
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive(
														"",
														"A",
														"AAAA",
														"SRV",
													),
												},
											},
										},
									},
									"origin_port": schema.DynamicAttribute{
										Description: "The destination port at the origin. Only specified in conjunction with origin_dns. May use an integer to specify a single origin port, for example `1000`, or a string to specify a range of origin ports, for example `\"1000-2000\"`.\nNotes: If specifying a port range, the number of ports in the range must match the number of ports specified in the \"protocol\" field.",
										Computed:    true,
										Validators: []validator.Dynamic{
											customvalidator.AllowedSubtypes(basetypes.Int64Type{}, basetypes.StringType{}),
										},
									},
								},
							},
						},
						"result_info": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[SpectrumApplicationsResultInfoDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"count": schema.Float64Attribute{
									Description: "Total number of results for the requested service",
									Computed:    true,
								},
								"page": schema.Float64Attribute{
									Description: "Current page within paginated list of results",
									Computed:    true,
								},
								"per_page": schema.Float64Attribute{
									Description: "Number of results per page of results",
									Computed:    true,
								},
								"total_count": schema.Float64Attribute{
									Description: "Total results available without any search parameters",
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *SpectrumApplicationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *SpectrumApplicationsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
