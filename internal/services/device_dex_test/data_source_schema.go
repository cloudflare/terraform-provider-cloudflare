// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_dex_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &DeviceDEXTestDataSource{}
var _ datasource.DataSourceWithValidateConfig = &DeviceDEXTestDataSource{}

func (r DeviceDEXTestDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"dex_test_id": schema.StringAttribute{
				Description: "API UUID.",
				Optional:    true,
			},
			"data": schema.SingleNestedAttribute{
				Description: "The configuration object which contains the details for the WARP client to conduct the test.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"host": schema.StringAttribute{
						Description: "The desired endpoint to test.",
						Optional:    true,
					},
					"kind": schema.StringAttribute{
						Description: "The type of test.",
						Optional:    true,
					},
					"method": schema.StringAttribute{
						Description: "The HTTP request method type.",
						Optional:    true,
					},
				},
			},
			"enabled": schema.BoolAttribute{
				Description: "Determines whether or not the test is active.",
				Optional:    true,
			},
			"interval": schema.StringAttribute{
				Description: "How often the test will run.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the DEX test. Must be unique.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "Additional details about the test.",
				Optional:    true,
			},
			"target_policies": schema.ListNestedAttribute{
				Description: "Device settings profiles targeted by this test",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The id of the device settings profile",
							Optional:    true,
						},
						"default": schema.BoolAttribute{
							Description: "Whether the profile is the account default",
							Optional:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the device settings profile",
							Optional:    true,
						},
					},
				},
			},
			"targeted": schema.BoolAttribute{
				Optional: true,
			},
			"find_one_by": schema.SingleNestedAttribute{
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

func (r *DeviceDEXTestDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *DeviceDEXTestDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
