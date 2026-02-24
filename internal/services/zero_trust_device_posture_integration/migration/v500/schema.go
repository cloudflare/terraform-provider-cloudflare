package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareDevicePostureIntegrationSchema returns the source schema for legacy resource.
// Schema version: 0 (SDKv2 implicit)
// Resource type: cloudflare_device_posture_integration
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceCloudflareDevicePostureIntegrationSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // SDKv2 implicit version
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
			"type": schema.StringAttribute{
				Required: true,
			},
			"identifier": schema.StringAttribute{
				Optional: true,
				// REMOVED in v5 - will be dropped during migration
			},
			"interval": schema.StringAttribute{
				Optional: true,
				// Optional in v4, Required in v5
			},
			// config: TypeList MaxItems:1 in v4
			// Stored as array in state: [{...}]
			"config": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"api_url": schema.StringAttribute{
							Optional: true,
						},
						"auth_url": schema.StringAttribute{
							Optional: true,
						},
						"client_id": schema.StringAttribute{
							Optional: true,
						},
						"client_secret": schema.StringAttribute{
							Optional: true,
						},
						"customer_id": schema.StringAttribute{
							Optional: true,
						},
						"client_key": schema.StringAttribute{
							Optional: true,
						},
						"access_client_id": schema.StringAttribute{
							Optional: true,
						},
						"access_client_secret": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}
