// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_posture_integration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &DevicePostureIntegrationsDataSource{}
var _ datasource.DataSourceWithValidateConfig = &DevicePostureIntegrationsDataSource{}

func (r DevicePostureIntegrationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
						"id": schema.StringAttribute{
							Description: "API UUID.",
							Computed:    true,
						},
						"interval": schema.StringAttribute{
							Description: "The interval between each posture check with the third-party API. Use `m` for minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the device posture integration.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "The type of device posture integration.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("workspace_one", "crowdstrike_s2s", "uptycs", "intune", "kolide", "tanium", "sentinelone_s2s"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *DevicePostureIntegrationsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *DevicePostureIntegrationsDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
