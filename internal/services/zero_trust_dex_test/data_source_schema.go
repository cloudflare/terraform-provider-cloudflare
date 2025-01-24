// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_test

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDEXTestDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The unique identifier for the test.",
				Computed:    true,
			},
			"dex_test_id": schema.StringAttribute{
				Description: "The unique identifier for the test.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Description: "Additional details about the test.",
				Computed:    true,
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
			"targeted": schema.BoolAttribute{
				Computed: true,
			},
			"test_id": schema.StringAttribute{
				Description: "The unique identifier for the test.",
				Computed:    true,
			},
			"data": schema.SingleNestedAttribute{
				Description: "The configuration object which contains the details for the WARP client to conduct the test.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustDEXTestDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"host": schema.StringAttribute{
						Description: "The desired endpoint to test.",
						Computed:    true,
					},
					"kind": schema.StringAttribute{
						Description: "The type of test.",
						Computed:    true,
					},
					"method": schema.StringAttribute{
						Description: "The HTTP request method type.",
						Computed:    true,
					},
				},
			},
			"target_policies": schema.ListNestedAttribute{
				Description: "Device settings profiles targeted by this test",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDEXTestTargetPoliciesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The id of the device settings profile",
							Computed:    true,
						},
						"default": schema.BoolAttribute{
							Description: "Whether the profile is the account default",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the device settings profile",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDEXTestDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDEXTestDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
