// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_dns

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*EmailRoutingDNSResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Domain of your zone.",
				Required:    true,
			},
			"created": schema.StringAttribute{
				Description: "The date and time the settings have been created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"enabled": schema.BoolAttribute{
				Description: "State of the zone settings for Email Routing.",
				Computed:    true,
			},
			"modified": schema.StringAttribute{
				Description: "The date and time the settings have been modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"skip_wizard": schema.BoolAttribute{
				Description: "Flag to check if the user skipped the configuration wizard.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Show the state of your account, and the type or configuration error.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ready",
						"unconfigured",
						"misconfigured",
						"misconfigured/locked",
						"unlocked",
					),
				},
			},
			"success": schema.BoolAttribute{
				Description: "Whether the API call was successful",
				Computed:    true,
			},
			"tag": schema.StringAttribute{
				Description: "Email Routing settings tag. (Deprecated, replaced by Email Routing settings identifier)",
				Computed:    true,
			},
			"errors": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[EmailRoutingDNSErrorsModel](ctx),
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
				CustomType: customfield.NewNestedObjectListType[EmailRoutingDNSMessagesModel](ctx),
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
			"result": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[EmailRoutingDNSResultModel](ctx),
				Attributes: map[string]schema.Attribute{
					"errors": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[EmailRoutingDNSResultErrorsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"code": schema.StringAttribute{
									Computed: true,
								},
								"missing": schema.SingleNestedAttribute{
									Description: "List of records needed to enable an Email Routing zone.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[EmailRoutingDNSResultErrorsMissingModel](ctx),
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
											Validators: []validator.Float64{
												float64validator.Between(1, 86400),
											},
										},
										"type": schema.StringAttribute{
											Description: "DNS record type.",
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
						CustomType: customfield.NewNestedObjectListType[EmailRoutingDNSResultRecordModel](ctx),
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
									Validators: []validator.Float64{
										float64validator.Between(1, 86400),
									},
								},
								"type": schema.StringAttribute{
									Description: "DNS record type.",
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
						Validators: []validator.Float64{
							float64validator.Between(1, 86400),
						},
					},
					"type": schema.StringAttribute{
						Description: "DNS record type.",
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
				CustomType: customfield.NewNestedObjectType[EmailRoutingDNSResultInfoModel](ctx),
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

func (r *EmailRoutingDNSResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *EmailRoutingDNSResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
