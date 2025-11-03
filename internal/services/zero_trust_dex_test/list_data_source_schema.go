// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_test

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDEXTestsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDEXTestsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"data": schema.SingleNestedAttribute{
							Description: "The configuration object which contains the details for the WARP client to conduct the test.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustDEXTestsDataDataSourceModel](ctx),
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
							Description: "Determines whether or not the test is active.",
							Computed:    true,
						},
						"interval": schema.StringAttribute{
							Description: "How often the test will run.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the DEX test. Must be unique.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "Additional details about the test.",
							Computed:    true,
						},
						"target_policies": schema.ListNestedAttribute{
							Description: "DEX rules targeted by this test",
							Optional:    true,
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustDEXTestsTargetPoliciesDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "API Resource UUID tag.",
										Computed:    true,
									},
									"default": schema.BoolAttribute{
										Description: "Whether the DEX rule is the account default",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "The name of the DEX rule",
										Computed:    true,
									},
								},
							},
						},
						"targeted": schema.BoolAttribute{
							Computed: true,
						},
						"test_id": schema.StringAttribute{
							Description: "The unique identifier for the test.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDEXTestsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustDEXTestsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
