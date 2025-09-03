// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_posture_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDevicePostureRuleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "API UUID.",
				Computed:    true,
			},
			"rule_id": schema.StringAttribute{
				Description: "API UUID.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the device posture rule.",
				Computed:    true,
			},
			"expiration": schema.StringAttribute{
				Description: "Sets the expiration time for a posture check result. If empty, the result remains valid until it is overwritten by new data from the WARP client.",
				Computed:    true,
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
				Description: "The type of device posture rule.\nAvailable values: \"file\", \"application\", \"tanium\", \"gateway\", \"warp\", \"disk_encryption\", \"serial_number\", \"sentinelone\", \"carbonblack\", \"firewall\", \"os_version\", \"domain_joined\", \"client_certificate\", \"client_certificate_v2\", \"unique_client_id\", \"kolide\", \"tanium_s2s\", \"crowdstrike_s2s\", \"intune\", \"workspace_one\", \"sentinelone_s2s\", \"custom_s2s\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"file",
						"application",
						"tanium",
						"gateway",
						"warp",
						"disk_encryption",
						"serial_number",
						"sentinelone",
						"carbonblack",
						"firewall",
						"os_version",
						"domain_joined",
						"client_certificate",
						"client_certificate_v2",
						"unique_client_id",
						"kolide",
						"tanium_s2s",
						"crowdstrike_s2s",
						"intune",
						"workspace_one",
						"sentinelone_s2s",
						"custom_s2s",
					),
				},
			},
			"input": schema.SingleNestedAttribute{
				Description: "The value to be checked against.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustDevicePostureRuleInputDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"operating_system": schema.StringAttribute{
						Description: "Operating system.\nAvailable values: \"windows\", \"linux\", \"mac\", \"android\", \"ios\", \"chromeos\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"windows",
								"linux",
								"mac",
								"android",
								"ios",
								"chromeos",
							),
						},
					},
					"path": schema.StringAttribute{
						Description: "File path.",
						Computed:    true,
					},
					"exists": schema.BoolAttribute{
						Description: "Whether or not file exists.",
						Computed:    true,
					},
					"sha256": schema.StringAttribute{
						Description: "SHA-256.",
						Computed:    true,
					},
					"thumbprint": schema.StringAttribute{
						Description: "Signing certificate thumbprint.",
						Computed:    true,
					},
					"id": schema.StringAttribute{
						Description: "List ID.",
						Computed:    true,
					},
					"domain": schema.StringAttribute{
						Description: "Domain.",
						Computed:    true,
					},
					"operator": schema.StringAttribute{
						Description: "Operator.\nAvailable values: \"<\", \"<=\", \">\", \">=\", \"==\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"<",
								"<=",
								">",
								">=",
								"==",
							),
						},
					},
					"version": schema.StringAttribute{
						Description: "Version of OS.",
						Computed:    true,
					},
					"os_distro_name": schema.StringAttribute{
						Description: "Operating System Distribution Name (linux only).",
						Computed:    true,
					},
					"os_distro_revision": schema.StringAttribute{
						Description: "Version of OS Distribution (linux only).",
						Computed:    true,
					},
					"os_version_extra": schema.StringAttribute{
						Description: "Additional version data. For Mac or iOS, the Product Version Extra. For Linux, the kernel release version. (Mac, iOS, and Linux only).",
						Computed:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Enabled.",
						Computed:    true,
					},
					"check_disks": schema.ListAttribute{
						Description: "List of volume names to be checked for encryption.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"require_all": schema.BoolAttribute{
						Description: "Whether to check all disks for encryption.",
						Computed:    true,
					},
					"certificate_id": schema.StringAttribute{
						Description: "UUID of Cloudflare managed certificate.",
						Computed:    true,
					},
					"cn": schema.StringAttribute{
						Description: "Common Name that is protected by the certificate.",
						Computed:    true,
					},
					"check_private_key": schema.BoolAttribute{
						Description: "Confirm the certificate was not imported from another device. We recommend keeping this enabled unless the certificate was deployed without a private key.",
						Computed:    true,
					},
					"extended_key_usage": schema.ListAttribute{
						Description: "List of values indicating purposes for which the certificate public key can be used.",
						Computed:    true,
						Validators: []validator.List{
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive("clientAuth", "emailProtection"),
							),
						},
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"locations": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustDevicePostureRuleInputLocationsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"paths": schema.ListAttribute{
								Description: "List of paths to check for client certificate on linux.",
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"trust_stores": schema.ListAttribute{
								Description: "List of trust stores to check for client certificate.",
								Computed:    true,
								Validators: []validator.List{
									listvalidator.ValueStringsAre(
										stringvalidator.OneOfCaseInsensitive("system", "user"),
									),
								},
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
						},
					},
					"subject_alternative_names": schema.ListAttribute{
						Description: "List of certificate Subject Alternative Names.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"compliance_status": schema.StringAttribute{
						Description: "Compliance Status.\nAvailable values: \"compliant\", \"noncompliant\", \"unknown\", \"notapplicable\", \"ingraceperiod\", \"error\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"compliant",
								"noncompliant",
								"unknown",
								"notapplicable",
								"ingraceperiod",
								"error",
							),
						},
					},
					"connection_id": schema.StringAttribute{
						Description: "Posture Integration ID.",
						Computed:    true,
					},
					"last_seen": schema.StringAttribute{
						Description: "For more details on last seen, please refer to the Crowdstrike documentation.",
						Computed:    true,
					},
					"os": schema.StringAttribute{
						Description: "Os Version.",
						Computed:    true,
					},
					"overall": schema.StringAttribute{
						Description: "Overall.",
						Computed:    true,
					},
					"sensor_config": schema.StringAttribute{
						Description: "SensorConfig.",
						Computed:    true,
					},
					"state": schema.StringAttribute{
						Description: "For more details on state, please refer to the Crowdstrike documentation.\nAvailable values: \"online\", \"offline\", \"unknown\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"online",
								"offline",
								"unknown",
							),
						},
					},
					"version_operator": schema.StringAttribute{
						Description: "Version Operator.\nAvailable values: \"<\", \"<=\", \">\", \">=\", \"==\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"<",
								"<=",
								">",
								">=",
								"==",
							),
						},
					},
					"count_operator": schema.StringAttribute{
						Description: "Count Operator.\nAvailable values: \"<\", \"<=\", \">\", \">=\", \"==\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"<",
								"<=",
								">",
								">=",
								"==",
							),
						},
					},
					"issue_count": schema.StringAttribute{
						Description: "The Number of Issues.",
						Computed:    true,
					},
					"eid_last_seen": schema.StringAttribute{
						Description: "For more details on eid last seen, refer to the Tanium documentation.",
						Computed:    true,
					},
					"risk_level": schema.StringAttribute{
						Description: "For more details on risk level, refer to the Tanium documentation.\nAvailable values: \"low\", \"medium\", \"high\", \"critical\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"low",
								"medium",
								"high",
								"critical",
							),
						},
					},
					"score_operator": schema.StringAttribute{
						Description: "Score Operator.\nAvailable values: \"<\", \"<=\", \">\", \">=\", \"==\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"<",
								"<=",
								">",
								">=",
								"==",
							),
						},
					},
					"total_score": schema.Float64Attribute{
						Description: "For more details on total score, refer to the Tanium documentation.",
						Computed:    true,
					},
					"active_threats": schema.Float64Attribute{
						Description: "The Number of active threats.",
						Computed:    true,
					},
					"infected": schema.BoolAttribute{
						Description: "Whether device is infected.",
						Computed:    true,
					},
					"is_active": schema.BoolAttribute{
						Description: "Whether device is active.",
						Computed:    true,
					},
					"network_status": schema.StringAttribute{
						Description: "Network status of device.\nAvailable values: \"connected\", \"disconnected\", \"disconnecting\", \"connecting\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"connected",
								"disconnected",
								"disconnecting",
								"connecting",
							),
						},
					},
					"operational_state": schema.StringAttribute{
						Description: "Agent operational state.\nAvailable values: \"na\", \"partially_disabled\", \"auto_fully_disabled\", \"fully_disabled\", \"auto_partially_disabled\", \"disabled_error\", \"db_corruption\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"na",
								"partially_disabled",
								"auto_fully_disabled",
								"fully_disabled",
								"auto_partially_disabled",
								"disabled_error",
								"db_corruption",
							),
						},
					},
					"score": schema.Float64Attribute{
						Description: "A value between 0-100 assigned to devices set by the 3rd party posture provider.",
						Computed:    true,
					},
				},
			},
			"match": schema.ListNestedAttribute{
				Description: "The conditions that the client must match to run the rule.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDevicePostureRuleMatchDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"platform": schema.StringAttribute{
							Description: `Available values: "windows", "mac", "linux", "android", "ios", "chromeos".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"windows",
									"mac",
									"linux",
									"android",
									"ios",
									"chromeos",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDevicePostureRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDevicePostureRuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
