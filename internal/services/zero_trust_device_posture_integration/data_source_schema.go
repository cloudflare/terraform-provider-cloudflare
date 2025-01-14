// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_posture_integration

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDevicePostureIntegrationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
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
					stringvalidator.OneOfCaseInsensitive(
						"workspace_one",
						"crowdstrike_s2s",
						"uptycs",
						"intune",
						"kolide",
						"tanium",
						"sentinelone_s2s",
						"custom_s2s",
					),
				},
			},
			"config": schema.SingleNestedAttribute{
				Description: "The Workspace One Config Response.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustDevicePostureIntegrationConfigDataSourceModel](ctx),
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

func (d *ZeroTrustDevicePostureIntegrationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDevicePostureIntegrationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("integration_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("integration_id")),
	}
}
