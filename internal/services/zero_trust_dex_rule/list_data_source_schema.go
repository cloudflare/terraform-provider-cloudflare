// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDEXRulesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Description: "Filter results by rule name",
				Optional:    true,
			},
			"sort_by": schema.StringAttribute{
				Description: "Which property to sort results by\nAvailable values: \"name\", \"created_at\", \"updated_at\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"name",
						"created_at",
						"updated_at",
					),
				},
			},
			"sort_order": schema.StringAttribute{
				Description: "Sort direction for sort_by property\nAvailable values: \"ASC\", \"DESC\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("ASC", "DESC"),
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDEXRulesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"rules": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[ZeroTrustDEXRulesRulesDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "API Resource UUID tag.",
										Computed:    true,
									},
									"created_at": schema.StringAttribute{
										Computed: true,
									},
									"match": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"description": schema.StringAttribute{
										Computed: true,
									},
									"targeted_tests": schema.ListNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectListType[ZeroTrustDEXRulesRulesTargetedTestsDataSourceModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"data": schema.SingleNestedAttribute{
													Description: "The configuration object which contains the details for the WARP client to conduct the test.",
													Computed:    true,
													CustomType:  customfield.NewNestedObjectType[ZeroTrustDEXRulesRulesTargetedTestsDataDataSourceModel](ctx),
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description: "The desired endpoint to test.",
															Computed:    true,
														},
														"kind": schema.StringAttribute{
															Description: "The type of test.\nAvailable values: \"http\", \"traceroute\".",
															Computed:    true,
															Validators: []validator.String{
																stringvalidator.OneOfCaseInsensitive("http", "traceroute"),
															},
														},
														"method": schema.StringAttribute{
															Description: "The HTTP request method type.\nAvailable values: \"GET\".",
															Computed:    true,
															Validators: []validator.String{
																stringvalidator.OneOfCaseInsensitive("GET"),
															},
														},
													},
												},
												"enabled": schema.BoolAttribute{
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
												"test_id": schema.StringAttribute{
													Computed: true,
												},
											},
										},
									},
									"updated_at": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDEXRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustDEXRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
