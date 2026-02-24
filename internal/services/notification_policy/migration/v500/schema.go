// File generated for v4 to v5 state migration

package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareNotificationPolicySchema returns the source schema for legacy resource.
// Schema version: 0 (SDK v2 implicit version)
// Resource type: cloudflare_notification_policy
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed, ElementType).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceCloudflareNotificationPolicySchema() schema.Schema {
	return schema.Schema{
		Version: 0, // SDK v2 implicit version
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"alert_type": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"enabled": schema.BoolAttribute{
				Required: true,
			},
			"created": schema.StringAttribute{
				Computed: true,
			},
			"modified": schema.StringAttribute{
				Computed: true,
			},
			// filters: TypeList MaxItems:1 in v4 (stored as array in SDK v2 state)
			"filters": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// All filter fields are TypeSet in v4
						"actions": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"airport_code": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"affected_components": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"status": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"health_check_id": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"zones": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"services": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"product": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"limit": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"enabled": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"pool_id": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"slo": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"where": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"group_by": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"alert_trigger_preferences": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"requests_per_second": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"target_zone_name": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"target_hostname": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"target_ip": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"packets_per_second": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"protocol": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"project_id": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"environment": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"event": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"event_source": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"new_health": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"input_id": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"event_type": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"megabits_per_second": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"incident_impact": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"new_status": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"selectors": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"tunnel_id": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"tunnel_name": schema.SetAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			// Integration fields: TypeSet in v4
			"email_integration": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
						"name": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"webhooks_integration": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
						"name": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"pagerduty_integration": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
						"name": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}
