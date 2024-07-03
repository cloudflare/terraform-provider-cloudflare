// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_posture_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &DevicePostureRulesDataSource{}
var _ datasource.DataSourceWithValidateConfig = &DevicePostureRulesDataSource{}

func (r DevicePostureRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
						"description": schema.StringAttribute{
							Description: "The description of the device posture rule.",
							Computed:    true,
						},
						"expiration": schema.StringAttribute{
							Description: "Sets the expiration time for a posture check result. If empty, the result remains valid until it is overwritten by new data from the WARP client.",
							Computed:    true,
						},
						"match": schema.ListNestedAttribute{
							Description: "The conditions that the client must match to run the rule.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"platform": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("windows", "mac", "linux", "android", "ios"),
										},
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the device posture rule.",
							Computed:    true,
						},
						"schedule": schema.StringAttribute{
							Description: "Polling frequency for the WARP client posture check. Default: `5m` (poll every five minutes). Minimum: `1m`.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "The type of device posture rule.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("file", "application", "tanium", "gateway", "warp", "disk_encryption", "sentinelone", "carbonblack", "firewall", "os_version", "domain_joined", "client_certificate", "unique_client_id", "kolide", "tanium_s2s", "crowdstrike_s2s", "intune", "workspace_one", "sentinelone_s2s"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *DevicePostureRulesDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *DevicePostureRulesDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
