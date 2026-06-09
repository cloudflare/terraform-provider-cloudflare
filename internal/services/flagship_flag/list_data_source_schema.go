// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship_flag

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*FlagshipFlagsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account Firewall Access Rules Read",
				"Account Firewall Access Rules Write",
				"Account Settings Read",
				"Account Settings Write",
				"Billing Read",
				"Billing Write",
				"DDoS Botnet Feed Read",
				"DDoS Botnet Feed Write",
				"DDoS Protection Read",
				"DDoS Protection Write",
				"DNS Firewall Read",
				"DNS Firewall Write",
				"DNS View Read",
				"DNS View Write",
				"Load Balancers Account Read",
				"Load Balancers Account Write",
				"Load Balancing: Monitors and Pools Read",
				"Load Balancing: Monitors and Pools Write",
				"SCIM Provisioning",
				"Trust and Safety Read",
				"Trust and Safety Write",
				"Workers KV Storage Read",
				"Workers KV Storage Write",
				"Workers R2 Storage Read",
				"Workers R2 Storage Write",
				"Workers Scripts Read",
				"Workers Scripts Write",
				"Workers Tail Read",
				"Zero Trust: PII Read",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account ID.",
				Required:    true,
			},
			"app_id": schema.StringAttribute{
				Description: "App identifier.",
				Required:    true,
			},
			"limit": schema.StringAttribute{
				Description: "Max items to return (1–200).",
				Optional:    true,
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
				CustomType:  customfield.NewNestedObjectListType[FlagshipFlagsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"default_variation": schema.StringAttribute{
							Description: "Variation served when no rule matches or the flag is disabled. Must be a key in `variations`.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "When false, the flag bypasses all rules and always serves `default_variation`.",
							Computed:    true,
						},
						"key": schema.StringAttribute{
							Description: "Unique identifier for the flag within an app. Used in all evaluation and SDK calls.",
							Computed:    true,
						},
						"rules": schema.ListNestedAttribute{
							Description: "Targeting rules evaluated in ascending `priority`; the first matching rule wins. An empty array means the flag always serves `default_variation`.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[FlagshipFlagsRulesDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"conditions": schema.ListNestedAttribute{
										Description: "Conditions the context must satisfy for this rule to match. An empty array matches all contexts.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectListType[FlagshipFlagsRulesConditionsDataSourceModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"attribute": schema.StringAttribute{
													Computed: true,
												},
												"operator": schema.StringAttribute{
													Description: `Available values: "equals", "not_equals", "greater_than", "less_than", "greater_than_or_equals", "less_than_or_equals", "contains", "starts_with", "ends_with", "in", "not_in".`,
													Computed:    true,
													Validators: []validator.String{
														stringvalidator.OneOfCaseInsensitive(
															"equals",
															"not_equals",
															"greater_than",
															"less_than",
															"greater_than_or_equals",
															"less_than_or_equals",
															"contains",
															"starts_with",
															"ends_with",
															"in",
															"not_in",
														),
													},
												},
												"value": schema.StringAttribute{
													Description: "Value to compare against the context attribute. Must be an array for `in` and `not_in`; numeric and ISO-8601 datetime strings are accepted by the ordering operators.",
													Computed:    true,
													CustomType:  jsontypes.NormalizedType{},
												},
												"clauses": schema.ListNestedAttribute{
													Computed:   true,
													CustomType: customfield.NewNestedObjectListType[FlagshipFlagsRulesConditionsClausesDataSourceModel](ctx),
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"attribute": schema.StringAttribute{
																Computed: true,
															},
															"operator": schema.StringAttribute{
																Description: `Available values: "equals", "not_equals", "greater_than", "less_than", "greater_than_or_equals", "less_than_or_equals", "contains", "starts_with", "ends_with", "in", "not_in".`,
																Computed:    true,
																Validators: []validator.String{
																	stringvalidator.OneOfCaseInsensitive(
																		"equals",
																		"not_equals",
																		"greater_than",
																		"less_than",
																		"greater_than_or_equals",
																		"less_than_or_equals",
																		"contains",
																		"starts_with",
																		"ends_with",
																		"in",
																		"not_in",
																	),
																},
															},
															"value": schema.StringAttribute{
																Description: "Value to compare against the context attribute. Must be an array for `in` and `not_in`; numeric and ISO-8601 datetime strings are accepted by the ordering operators.",
																Computed:    true,
																CustomType:  jsontypes.NormalizedType{},
															},
															"clauses": schema.ListNestedAttribute{
																Computed:   true,
																CustomType: customfield.NewNestedObjectListType[FlagshipFlagsRulesConditionsClausesClausesDataSourceModel](ctx),
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"attribute": schema.StringAttribute{
																			Computed: true,
																		},
																		"operator": schema.StringAttribute{
																			Description: `Available values: "equals", "not_equals", "greater_than", "less_than", "greater_than_or_equals", "less_than_or_equals", "contains", "starts_with", "ends_with", "in", "not_in".`,
																			Computed:    true,
																			Validators: []validator.String{
																				stringvalidator.OneOfCaseInsensitive(
																					"equals",
																					"not_equals",
																					"greater_than",
																					"less_than",
																					"greater_than_or_equals",
																					"less_than_or_equals",
																					"contains",
																					"starts_with",
																					"ends_with",
																					"in",
																					"not_in",
																				),
																			},
																		},
																		"value": schema.StringAttribute{
																			Description: "Value to compare against the context attribute. Must be an array for `in` and `not_in`; numeric and ISO-8601 datetime strings are accepted by the ordering operators.",
																			Computed:    true,
																			CustomType:  jsontypes.NormalizedType{},
																		},
																		"clauses": schema.ListNestedAttribute{
																			Computed:   true,
																			CustomType: customfield.NewNestedObjectListType[FlagshipFlagsRulesConditionsClausesClausesClausesDataSourceModel](ctx),
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"attribute": schema.StringAttribute{
																						Computed: true,
																					},
																					"operator": schema.StringAttribute{
																						Description: `Available values: "equals", "not_equals", "greater_than", "less_than", "greater_than_or_equals", "less_than_or_equals", "contains", "starts_with", "ends_with", "in", "not_in".`,
																						Computed:    true,
																						Validators: []validator.String{
																							stringvalidator.OneOfCaseInsensitive(
																								"equals",
																								"not_equals",
																								"greater_than",
																								"less_than",
																								"greater_than_or_equals",
																								"less_than_or_equals",
																								"contains",
																								"starts_with",
																								"ends_with",
																								"in",
																								"not_in",
																							),
																						},
																					},
																					"value": schema.StringAttribute{
																						Description: "Value to compare against the context attribute. Must be an array for `in` and `not_in`; numeric and ISO-8601 datetime strings are accepted by the ordering operators.",
																						Computed:    true,
																						CustomType:  jsontypes.NormalizedType{},
																					},
																					"clauses": schema.ListNestedAttribute{
																						Computed:   true,
																						CustomType: customfield.NewNestedObjectListType[FlagshipFlagsRulesConditionsClausesClausesClausesClausesDataSourceModel](ctx),
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"attribute": schema.StringAttribute{
																									Computed: true,
																								},
																								"operator": schema.StringAttribute{
																									Description: `Available values: "equals", "not_equals", "greater_than", "less_than", "greater_than_or_equals", "less_than_or_equals", "contains", "starts_with", "ends_with", "in", "not_in".`,
																									Computed:    true,
																									Validators: []validator.String{
																										stringvalidator.OneOfCaseInsensitive(
																											"equals",
																											"not_equals",
																											"greater_than",
																											"less_than",
																											"greater_than_or_equals",
																											"less_than_or_equals",
																											"contains",
																											"starts_with",
																											"ends_with",
																											"in",
																											"not_in",
																										),
																									},
																								},
																								"value": schema.StringAttribute{
																									Description: "Value to compare against the context attribute. Must be an array for `in` and `not_in`; numeric and ISO-8601 datetime strings are accepted by the ordering operators.",
																									Computed:    true,
																									CustomType:  jsontypes.NormalizedType{},
																								},
																								"clauses": schema.ListNestedAttribute{
																									Computed:   true,
																									CustomType: customfield.NewNestedObjectListType[FlagshipFlagsRulesConditionsClausesClausesClausesClausesClausesDataSourceModel](ctx),
																									NestedObject: schema.NestedAttributeObject{
																										Attributes: map[string]schema.Attribute{
																											"attribute": schema.StringAttribute{
																												Computed: true,
																											},
																											"operator": schema.StringAttribute{
																												Description: `Available values: "equals", "not_equals", "greater_than", "less_than", "greater_than_or_equals", "less_than_or_equals", "contains", "starts_with", "ends_with", "in", "not_in".`,
																												Computed:    true,
																												Validators: []validator.String{
																													stringvalidator.OneOfCaseInsensitive(
																														"equals",
																														"not_equals",
																														"greater_than",
																														"less_than",
																														"greater_than_or_equals",
																														"less_than_or_equals",
																														"contains",
																														"starts_with",
																														"ends_with",
																														"in",
																														"not_in",
																													),
																												},
																											},
																											"value": schema.StringAttribute{
																												Description: "Value to compare against the context attribute. Must be an array for `in` and `not_in`; numeric and ISO-8601 datetime strings are accepted by the ordering operators.",
																												Computed:    true,
																												CustomType:  jsontypes.NormalizedType{},
																											},
																											"clauses": schema.ListAttribute{
																												Computed:    true,
																												CustomType:  customfield.NewListType[types.String](ctx),
																												ElementType: types.StringType,
																											},
																											"logical_operator": schema.StringAttribute{
																												Description: `Available values: "AND", "OR".`,
																												Computed:    true,
																												Validators: []validator.String{
																													stringvalidator.OneOfCaseInsensitive("AND", "OR"),
																												},
																											},
																										},
																									},
																								},
																								"logical_operator": schema.StringAttribute{
																									Description: `Available values: "AND", "OR".`,
																									Computed:    true,
																									Validators: []validator.String{
																										stringvalidator.OneOfCaseInsensitive("AND", "OR"),
																									},
																								},
																							},
																						},
																					},
																					"logical_operator": schema.StringAttribute{
																						Description: `Available values: "AND", "OR".`,
																						Computed:    true,
																						Validators: []validator.String{
																							stringvalidator.OneOfCaseInsensitive("AND", "OR"),
																						},
																					},
																				},
																			},
																		},
																		"logical_operator": schema.StringAttribute{
																			Description: `Available values: "AND", "OR".`,
																			Computed:    true,
																			Validators: []validator.String{
																				stringvalidator.OneOfCaseInsensitive("AND", "OR"),
																			},
																		},
																	},
																},
															},
															"logical_operator": schema.StringAttribute{
																Description: `Available values: "AND", "OR".`,
																Computed:    true,
																Validators: []validator.String{
																	stringvalidator.OneOfCaseInsensitive("AND", "OR"),
																},
															},
														},
													},
												},
												"logical_operator": schema.StringAttribute{
													Description: `Available values: "AND", "OR".`,
													Computed:    true,
													Validators: []validator.String{
														stringvalidator.OneOfCaseInsensitive("AND", "OR"),
													},
												},
											},
										},
									},
									"priority": schema.Int64Attribute{
										Description: "Evaluation order; lower numbers are evaluated first. Must be unique across the flag's rules.",
										Computed:    true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
									"serve_variation": schema.StringAttribute{
										Description: "Variation served when this rule matches. Must be a key in `variations`.",
										Computed:    true,
									},
									"rollout": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[FlagshipFlagsRulesRolloutDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"percentage": schema.Float64Attribute{
												Description: "Percentage of matching traffic (0–100) served this variation. For multi-way splits, use cumulative upper bounds across rules (e.g. 30, 70, 100).",
												Computed:    true,
												Validators: []validator.Float64{
													float64validator.Between(0, 100),
												},
											},
											"attribute": schema.StringAttribute{
												Description: "Context attribute used for sticky bucketing. Defaults to `targetingKey`. If absent at evaluation time, bucketing is random per request.",
												Computed:    true,
											},
										},
									},
								},
							},
						},
						"variations": schema.MapAttribute{
							Description: "Map of variation name to value. All values must be the same type (boolean, string, number, or JSON object/array). Each serialized value must be 10KB or smaller.",
							Computed:    true,
							CustomType:  customfield.NewMapType[types.String](ctx),
							ElementType: types.StringType,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Description: "Value type of the flag's variations. Inferred from the variation values on write, so it may be omitted in requests.\nAvailable values: \"boolean\", \"string\", \"number\", \"json\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"boolean",
									"string",
									"number",
									"json",
								),
							},
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
						},
						"updated_by": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *FlagshipFlagsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *FlagshipFlagsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
