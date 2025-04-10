// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_dns

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*EmailRoutingDNSDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"subdomain": schema.StringAttribute{
				Description: "Domain of your zone.",
				Optional:    true,
			},
			"success": schema.BoolAttribute{
				Description: "Whether the API call was successful",
				Computed:    true,
			},
			"errors": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[EmailRoutingDNSErrorsDataSourceModel](ctx),
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
						"documentation_url": schema.StringAttribute{
							Computed: true,
						},
						"source": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[EmailRoutingDNSErrorsSourceDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"pointer": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
			"messages": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[EmailRoutingDNSMessagesDataSourceModel](ctx),
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
						"documentation_url": schema.StringAttribute{
							Computed: true,
						},
						"source": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[EmailRoutingDNSMessagesSourceDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"pointer": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
			"result": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[EmailRoutingDNSResultDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"errors": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[EmailRoutingDNSResultErrorsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"code": schema.StringAttribute{
									Computed: true,
								},
								"missing": schema.SingleNestedAttribute{
									Description: "List of records needed to enable an Email Routing zone.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[EmailRoutingDNSResultErrorsMissingDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"content": schema.StringAttribute{
											Description: "DNS record content.",
											Computed:    true,
										},
										"name": schema.StringAttribute{
											Description: "DNS record name (or @ for the zone apex).",
											Computed:    true,
										},
										"priority": schema.Float64Attribute{
											Description: "Required for MX, SRV and URI records. Unused by other record types. Records with lower priorities are preferred.",
											Computed:    true,
											Validators: []validator.Float64{
												float64validator.Between(0, 65535),
											},
										},
										"ttl": schema.Float64Attribute{
											Description: "Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1 for 'automatic'.",
											Computed:    true,
										},
										"type": schema.StringAttribute{
											Description: "DNS record type.\nAvailable values: \"A\", \"AAAA\", \"CNAME\", \"HTTPS\", \"TXT\", \"SRV\", \"LOC\", \"MX\", \"NS\", \"CERT\", \"DNSKEY\", \"DS\", \"NAPTR\", \"SMIMEA\", \"SSHFP\", \"SVCB\", \"TLSA\", \"URI\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"A",
													"AAAA",
													"CNAME",
													"HTTPS",
													"TXT",
													"SRV",
													"LOC",
													"MX",
													"NS",
													"CERT",
													"DNSKEY",
													"DS",
													"NAPTR",
													"SMIMEA",
													"SSHFP",
													"SVCB",
													"TLSA",
													"URI",
												),
											},
										},
									},
								},
							},
						},
					},
					"record": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[EmailRoutingDNSResultRecordDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"content": schema.StringAttribute{
									Description: "DNS record content.",
									Computed:    true,
								},
								"name": schema.StringAttribute{
									Description: "DNS record name (or @ for the zone apex).",
									Computed:    true,
								},
								"priority": schema.Float64Attribute{
									Description: "Required for MX, SRV and URI records. Unused by other record types. Records with lower priorities are preferred.",
									Computed:    true,
									Validators: []validator.Float64{
										float64validator.Between(0, 65535),
									},
								},
								"ttl": schema.Float64Attribute{
									Description: "Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1 for 'automatic'.",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "DNS record type.\nAvailable values: \"A\", \"AAAA\", \"CNAME\", \"HTTPS\", \"TXT\", \"SRV\", \"LOC\", \"MX\", \"NS\", \"CERT\", \"DNSKEY\", \"DS\", \"NAPTR\", \"SMIMEA\", \"SSHFP\", \"SVCB\", \"TLSA\", \"URI\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"A",
											"AAAA",
											"CNAME",
											"HTTPS",
											"TXT",
											"SRV",
											"LOC",
											"MX",
											"NS",
											"CERT",
											"DNSKEY",
											"DS",
											"NAPTR",
											"SMIMEA",
											"SSHFP",
											"SVCB",
											"TLSA",
											"URI",
										),
									},
								},
							},
						},
					},
					"content": schema.StringAttribute{
						Description: "DNS record content.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "DNS record name (or @ for the zone apex).",
						Computed:    true,
					},
					"priority": schema.Float64Attribute{
						Description: "Required for MX, SRV and URI records. Unused by other record types. Records with lower priorities are preferred.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 65535),
						},
					},
					"ttl": schema.Float64Attribute{
						Description: "Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1 for 'automatic'.",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "DNS record type.\nAvailable values: \"A\", \"AAAA\", \"CNAME\", \"HTTPS\", \"TXT\", \"SRV\", \"LOC\", \"MX\", \"NS\", \"CERT\", \"DNSKEY\", \"DS\", \"NAPTR\", \"SMIMEA\", \"SSHFP\", \"SVCB\", \"TLSA\", \"URI\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"A",
								"AAAA",
								"CNAME",
								"HTTPS",
								"TXT",
								"SRV",
								"LOC",
								"MX",
								"NS",
								"CERT",
								"DNSKEY",
								"DS",
								"NAPTR",
								"SMIMEA",
								"SSHFP",
								"SVCB",
								"TLSA",
								"URI",
							),
						},
					},
				},
			},
			"result_info": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[EmailRoutingDNSResultInfoDataSourceModel](ctx),
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
	}
}

func (d *EmailRoutingDNSDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *EmailRoutingDNSDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
