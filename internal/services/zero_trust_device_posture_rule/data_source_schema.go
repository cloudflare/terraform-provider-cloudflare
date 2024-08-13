// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_posture_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &ZeroTrustDevicePostureRuleDataSource{}

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"rule_id": schema.StringAttribute{
				Description: "API UUID.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the device posture rule.",
				Computed:    true,
				Optional:    true,
			},
			"expiration": schema.StringAttribute{
				Description: "Sets the expiration time for a posture check result. If empty, the result remains valid until it is overwritten by new data from the WARP client.",
				Computed:    true,
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "API UUID.",
				Computed:    true,
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the device posture rule.",
				Computed:    true,
				Optional:    true,
			},
			"schedule": schema.StringAttribute{
				Description: "Polling frequency for the WARP client posture check. Default: `5m` (poll every five minutes). Minimum: `1m`.",
				Computed:    true,
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of device posture rule.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"file",
						"application",
						"tanium",
						"gateway",
						"warp",
						"disk_encryption",
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
					),
				},
			},
			"input": schema.SingleNestedAttribute{
				Description: "The value to be checked against.",
				Computed:    true,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"operating_system": schema.StringAttribute{
						Description: "Operating system",
						Computed:    true,
						Optional:    true,
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
						Optional:    true,
					},
					"exists": schema.BoolAttribute{
						Description: "Whether or not file exists",
						Computed:    true,
						Optional:    true,
					},
					"sha256": schema.StringAttribute{
						Description: "SHA-256.",
						Computed:    true,
						Optional:    true,
					},
					"thumbprint": schema.StringAttribute{
						Description: "Signing certificate thumbprint.",
						Computed:    true,
						Optional:    true,
					},
					"id": schema.StringAttribute{
						Description: "List ID.",
						Computed:    true,
						Optional:    true,
					},
					"domain": schema.StringAttribute{
						Description: "Domain",
						Computed:    true,
						Optional:    true,
					},
					"operator": schema.StringAttribute{
						Description: "operator",
						Computed:    true,
						Optional:    true,
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
						Description: "Version of OS",
						Computed:    true,
						Optional:    true,
					},
					"os_distro_name": schema.StringAttribute{
						Description: "Operating System Distribution Name (linux only)",
						Computed:    true,
						Optional:    true,
					},
					"os_distro_revision": schema.StringAttribute{
						Description: "Version of OS Distribution (linux only)",
						Computed:    true,
						Optional:    true,
					},
					"os_version_extra": schema.StringAttribute{
						Description: "Additional version data. For Mac or iOS, the Product Verison Extra. For Linux, the kernel release version. (Mac, iOS, and Linux only)",
						Computed:    true,
						Optional:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Enabled",
						Computed:    true,
						Optional:    true,
					},
					"check_disks": schema.ListAttribute{
						Description: "List of volume names to be checked for encryption.",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"require_all": schema.BoolAttribute{
						Description: "Whether to check all disks for encryption.",
						Computed:    true,
						Optional:    true,
					},
					"certificate_id": schema.StringAttribute{
						Description: "UUID of Cloudflare managed certificate.",
						Computed:    true,
						Optional:    true,
					},
					"cn": schema.StringAttribute{
						Description: "Common Name that is protected by the certificate",
						Computed:    true,
						Optional:    true,
					},
					"check_private_key": schema.BoolAttribute{
						Description: "Confirm the certificate was not imported from another device. We recommend keeping this enabled unless the certificate was deployed without a private key.",
						Computed:    true,
						Optional:    true,
					},
					"extended_key_usage": schema.ListAttribute{
						Description: "List of values indicating purposes for which the certificate public key can be used",
						Computed:    true,
						Optional:    true,
						Validators: []validator.List{
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive("clientAuth", "emailProtection"),
							),
						},
						ElementType: types.StringType,
					},
					"locations": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"paths": schema.ListAttribute{
								Description: "List of paths to check for client certificate on linux.",
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
							"trust_stores": schema.ListAttribute{
								Description: "List of trust stores to check for client certificate.",
								Computed:    true,
								Optional:    true,
								Validators: []validator.List{
									listvalidator.ValueStringsAre(
										stringvalidator.OneOfCaseInsensitive("system", "user"),
									),
								},
								ElementType: types.StringType,
							},
						},
					},
					"compliance_status": schema.StringAttribute{
						Description: "Compliance Status",
						Computed:    true,
						Optional:    true,
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
						Optional:    true,
					},
					"last_seen": schema.StringAttribute{
						Description: "For more details on last seen, please refer to the Crowdstrike documentation.",
						Computed:    true,
						Optional:    true,
					},
					"os": schema.StringAttribute{
						Description: "Os Version",
						Computed:    true,
						Optional:    true,
					},
					"overall": schema.StringAttribute{
						Description: "overall",
						Computed:    true,
						Optional:    true,
					},
					"sensor_config": schema.StringAttribute{
						Description: "SensorConfig",
						Computed:    true,
						Optional:    true,
					},
					"state": schema.StringAttribute{
						Description: "For more details on state, please refer to the Crowdstrike documentation.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"online",
								"offline",
								"unknown",
							),
						},
					},
					"version_operator": schema.StringAttribute{
						Description: "Version Operator",
						Computed:    true,
						Optional:    true,
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
						Description: "Count Operator",
						Computed:    true,
						Optional:    true,
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
						Optional:    true,
					},
					"eid_last_seen": schema.StringAttribute{
						Description: "For more details on eid last seen, refer to the Tanium documentation.",
						Computed:    true,
						Optional:    true,
					},
					"risk_level": schema.StringAttribute{
						Description: "For more details on risk level, refer to the Tanium documentation.",
						Computed:    true,
						Optional:    true,
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
						Description: "Score Operator",
						Computed:    true,
						Optional:    true,
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
						Optional:    true,
					},
					"active_threats": schema.Float64Attribute{
						Description: "The Number of active threats.",
						Computed:    true,
						Optional:    true,
					},
					"infected": schema.BoolAttribute{
						Description: "Whether device is infected.",
						Computed:    true,
						Optional:    true,
					},
					"is_active": schema.BoolAttribute{
						Description: "Whether device is active.",
						Computed:    true,
						Optional:    true,
					},
					"network_status": schema.StringAttribute{
						Description: "Network status of device.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"connected",
								"disconnected",
								"disconnecting",
								"connecting",
							),
						},
					},
				},
			},
			"match": schema.ListNestedAttribute{
				Description: "The conditions that the client must match to run the rule.",
				Computed:    true,
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"platform": schema.StringAttribute{
							Computed: true,
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"windows",
									"mac",
									"linux",
									"android",
									"ios",
								),
							},
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

func (d *ZeroTrustDevicePostureRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDevicePostureRuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
