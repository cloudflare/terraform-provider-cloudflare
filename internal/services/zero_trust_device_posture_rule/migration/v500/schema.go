package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceDevicePostureRuleSchema returns the minimal schema for the legacy cloudflare_device_posture_rule resource.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_device_posture_rule
//
// This schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed, ElementType).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceDevicePostureRuleSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // SDKv2 default schema version
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Optional: true, // Was Optional in v4, Required in v5
				Computed: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"expiration": schema.StringAttribute{
				Optional: true,
			},
			"schedule": schema.StringAttribute{
				Optional: true,
			},
			// In v4 SDKv2, input was a block (MaxItems:1) stored as a list in state JSON.
			"input": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"operating_system": schema.StringAttribute{
							Optional: true,
						},
						"path": schema.StringAttribute{
							Optional: true,
						},
						"exists": schema.BoolAttribute{
							Optional: true,
						},
						"sha256": schema.StringAttribute{
							Optional: true,
						},
						"thumbprint": schema.StringAttribute{
							Optional: true,
						},
						"id": schema.StringAttribute{
							Optional: true,
						},
						"domain": schema.StringAttribute{
							Optional: true,
						},
						"operator": schema.StringAttribute{
							Optional: true,
						},
						"version": schema.StringAttribute{
							Optional: true,
						},
						"os_distro_name": schema.StringAttribute{
							Optional: true,
						},
						"os_distro_revision": schema.StringAttribute{
							Optional: true,
						},
						"os_version_extra": schema.StringAttribute{
							Optional: true,
						},
						"enabled": schema.BoolAttribute{
							Optional: true,
						},
						// check_disks was a Set in v4
						"check_disks": schema.SetAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"require_all": schema.BoolAttribute{
							Optional: true,
						},
						"certificate_id": schema.StringAttribute{
							Optional: true,
						},
						"cn": schema.StringAttribute{
							Optional: true,
						},
						"check_private_key": schema.BoolAttribute{
							Optional: true,
						},
						"extended_key_usage": schema.ListAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						// locations was a nested block in v4
						"locations": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"paths": schema.ListAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
									"trust_stores": schema.ListAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
								},
							},
						},
						"subject_alternative_names": schema.ListAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"update_window_days": schema.Float64Attribute{
							Optional: true,
						},
						"compliance_status": schema.StringAttribute{
							Optional: true,
						},
						"connection_id": schema.StringAttribute{
							Optional: true,
						},
						"last_seen": schema.StringAttribute{
							Optional: true,
						},
						"os": schema.StringAttribute{
							Optional: true,
						},
						"overall": schema.StringAttribute{
							Optional: true,
						},
						"sensor_config": schema.StringAttribute{
							Optional: true,
						},
						"state": schema.StringAttribute{
							Optional: true,
						},
						"version_operator": schema.StringAttribute{
							Optional: true,
						},
						"count_operator": schema.StringAttribute{
							Optional: true,
						},
						"issue_count": schema.StringAttribute{
							Optional: true,
						},
						"eid_last_seen": schema.StringAttribute{
							Optional: true,
						},
						"risk_level": schema.StringAttribute{
							Optional: true,
						},
						"score_operator": schema.StringAttribute{
							Optional: true,
						},
						"total_score": schema.Float64Attribute{
							Optional: true,
						},
						"active_threats": schema.Float64Attribute{
							Optional: true,
						},
						"infected": schema.BoolAttribute{
							Optional: true,
						},
						"is_active": schema.BoolAttribute{
							Optional: true,
						},
						"network_status": schema.StringAttribute{
							Optional: true,
						},
						"operational_state": schema.StringAttribute{
							Optional: true,
						},
						"score": schema.Float64Attribute{
							Optional: true,
						},
						// running existed in v4 but was removed in v5
						"running": schema.BoolAttribute{
							Optional: true,
						},
					},
				},
			},
			// In v4 SDKv2, match was multiple blocks stored as a list in state JSON.
			"match": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"platform": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}
