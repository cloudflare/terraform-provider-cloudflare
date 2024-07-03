// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_dex_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &DeviceDEXTestsDataSource{}
var _ datasource.DataSourceWithValidateConfig = &DeviceDEXTestsDataSource{}

func (r DeviceDEXTestsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
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
							Description: "Device settings profiles targeted by this test",
							Computed:    true,
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
						"targeted": schema.BoolAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *DeviceDEXTestsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *DeviceDEXTestsDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
