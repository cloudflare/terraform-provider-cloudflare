// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_test

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDEXTestDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"dex_test_id": schema.StringAttribute{
				Description: "API UUID.",
				Optional:    true,
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
			"data": schema.SingleNestedAttribute{
				Description: "The configuration object which contains the details for the WARP client to conduct the test.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustDEXTestDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"host": schema.StringAttribute{
						Description: "The desired endpoint to test.",
						Computed:    true,
						Optional:    true,
					},
					"kind": schema.StringAttribute{
						Description: "The type of test.",
						Computed:    true,
						Optional:    true,
					},
					"method": schema.StringAttribute{
						Description: "The HTTP request method type.",
						Computed:    true,
						Optional:    true,
					},
				},
			},
			"description": schema.StringAttribute{
				Description: "Additional details about the test.",
				Computed:    true,
				Optional:    true,
			},
			"targeted": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"target_policies": schema.ListNestedAttribute{
				Description: "Device settings profiles targeted by this test",
				Computed:    true,
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The id of the device settings profile",
							Computed:    true,
							Optional:    true,
						},
						"default": schema.BoolAttribute{
							Description: "Whether the profile is the account default",
							Computed:    true,
							Optional:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the device settings profile",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Required: true,
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
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("dex_test_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("dex_test_id")),
	}
}
