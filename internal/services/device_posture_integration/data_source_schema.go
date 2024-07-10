// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_posture_integration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &DevicePostureIntegrationDataSource{}
var _ datasource.DataSourceWithValidateConfig = &DevicePostureIntegrationDataSource{}

func (r DevicePostureIntegrationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"integration_id": schema.StringAttribute{
				Description: "API UUID.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "API UUID.",
				Computed:    true,
				Optional:    true,
			},
			"config": schema.SingleNestedAttribute{
				Description: "The configuration object containing third-party integration information.",
				Computed:    true,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"api_url": schema.StringAttribute{
						Description: "The Workspace One API URL provided in the Workspace One Admin Dashboard.",
						Computed:    true,
					},
					"auth_url": schema.StringAttribute{
						Description: "The Workspace One Authorization URL depending on your region.",
						Computed:    true,
					},
					"client_id": schema.StringAttribute{
						Description: "The Workspace One client ID provided in the Workspace One Admin Dashboard.",
						Computed:    true,
					},
				},
			},
			"interval": schema.StringAttribute{
				Description: "The interval between each posture check with the third-party API. Use `m` for minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).",
				Computed:    true,
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the device posture integration.",
				Computed:    true,
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of device posture integration.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("workspace_one", "crowdstrike_s2s", "uptycs", "intune", "kolide", "tanium", "sentinelone_s2s"),
				},
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

func (r *DevicePostureIntegrationDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *DevicePostureIntegrationDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
