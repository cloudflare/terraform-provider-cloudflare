// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_posture_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r DevicePostureRuleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "API UUID.",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"account_id": schema.StringAttribute{
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"name": schema.StringAttribute{
						Description: "The name of the device posture rule.",
						Required:    true,
					},
					"type": schema.StringAttribute{
						Description: "The type of device posture rule.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("file", "application", "tanium", "gateway", "warp", "disk_encryption", "sentinelone", "carbonblack", "firewall", "os_version", "domain_joined", "client_certificate", "unique_client_id", "kolide", "tanium_s2s", "crowdstrike_s2s", "intune", "workspace_one", "sentinelone_s2s"),
						},
					},
					"description": schema.StringAttribute{
						Description: "The description of the device posture rule.",
						Optional:    true,
					},
					"expiration": schema.StringAttribute{
						Description: "Sets the expiration time for a posture check result. If empty, the result remains valid until it is overwritten by new data from the WARP client.",
						Optional:    true,
					},
					"input": schema.SingleNestedAttribute{
						Description: "The value to be checked against.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"operating_system": schema.StringAttribute{
								Description: "Operating system",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("windows", "linux", "mac", "android", "ios", "chromeos"),
								},
							},
							"path": schema.StringAttribute{
								Description: "File path.",
								Optional:    true,
							},
							"exists": schema.BoolAttribute{
								Description: "Whether or not file exists",
								Optional:    true,
							},
							"sha256": schema.StringAttribute{
								Description: "SHA-256.",
								Optional:    true,
							},
							"thumbprint": schema.StringAttribute{
								Description: "Signing certificate thumbprint.",
								Optional:    true,
							},
							"id": schema.StringAttribute{
								Description: "List ID.",
								Optional:    true,
							},
							"domain": schema.StringAttribute{
								Description: "Domain",
								Optional:    true,
							},
							"operator": schema.StringAttribute{
								Description: "operator",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("<", "<=", ">", ">=", "=="),
								},
							},
							"version": schema.StringAttribute{
								Description: "Version of OS",
								Optional:    true,
							},
							"os_distro_name": schema.StringAttribute{
								Description: "Operating System Distribution Name (linux only)",
								Optional:    true,
							},
							"os_distro_revision": schema.StringAttribute{
								Description: "Version of OS Distribution (linux only)",
								Optional:    true,
							},
							"os_version_extra": schema.StringAttribute{
								Description: "Additional version data. For Mac or iOS, the Product Verison Extra. For Linux, the kernel release version. (Mac, iOS, and Linux only)",
								Optional:    true,
							},
							"enabled": schema.BoolAttribute{
								Description: "Enabled",
								Optional:    true,
							},
							"check_disks": schema.ListAttribute{
								Description: "List of volume names to be checked for encryption.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"require_all": schema.BoolAttribute{
								Description: "Whether to check all disks for encryption.",
								Optional:    true,
							},
							"certificate_id": schema.StringAttribute{
								Description: "UUID of Cloudflare managed certificate.",
								Optional:    true,
							},
							"cn": schema.StringAttribute{
								Description: "Common Name that is protected by the certificate",
								Optional:    true,
							},
							"compliance_status": schema.StringAttribute{
								Description: "Compliance Status",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("compliant", "noncompliant", "unknown", "notapplicable", "ingraceperiod", "error"),
								},
							},
							"connection_id": schema.StringAttribute{
								Description: "Posture Integration ID.",
								Optional:    true,
							},
							"last_seen": schema.StringAttribute{
								Description: "For more details on last seen, please refer to the Crowdstrike documentation.",
								Optional:    true,
							},
							"os": schema.StringAttribute{
								Description: "Os Version",
								Optional:    true,
							},
							"overall": schema.StringAttribute{
								Description: "overall",
								Optional:    true,
							},
							"sensor_config": schema.StringAttribute{
								Description: "SensorConfig",
								Optional:    true,
							},
							"state": schema.StringAttribute{
								Description: "For more details on state, please refer to the Crowdstrike documentation.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("online", "offline", "unknown"),
								},
							},
							"version_operator": schema.StringAttribute{
								Description: "Version Operator",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("<", "<=", ">", ">=", "=="),
								},
							},
							"count_operator": schema.StringAttribute{
								Description: "Count Operator",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("<", "<=", ">", ">=", "=="),
								},
							},
							"issue_count": schema.StringAttribute{
								Description: "The Number of Issues.",
								Optional:    true,
							},
							"eid_last_seen": schema.StringAttribute{
								Description: "For more details on eid last seen, refer to the Tanium documentation.",
								Optional:    true,
							},
							"risk_level": schema.StringAttribute{
								Description: "For more details on risk level, refer to the Tanium documentation.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("low", "medium", "high", "critical"),
								},
							},
							"score_operator": schema.StringAttribute{
								Description: "Score Operator",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("<", "<=", ">", ">=", "=="),
								},
							},
							"total_score": schema.Float64Attribute{
								Description: "For more details on total score, refer to the Tanium documentation.",
								Optional:    true,
							},
							"active_threats": schema.Float64Attribute{
								Description: "The Number of active threats.",
								Optional:    true,
							},
							"infected": schema.BoolAttribute{
								Description: "Whether device is infected.",
								Optional:    true,
							},
							"is_active": schema.BoolAttribute{
								Description: "Whether device is active.",
								Optional:    true,
							},
							"network_status": schema.StringAttribute{
								Description: "Network status of device.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("connected", "disconnected", "disconnecting", "connecting"),
								},
							},
						},
					},
					"match": schema.ListNestedAttribute{
						Description: "The conditions that the client must match to run the rule.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"platform": schema.StringAttribute{
									Optional: true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("windows", "mac", "linux", "android", "ios"),
									},
								},
							},
						},
					},
					"schedule": schema.StringAttribute{
						Description: "Polling frequency for the WARP client posture check. Default: `5m` (poll every five minutes). Minimum: `1m`.",
						Optional:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state DevicePostureRuleModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
