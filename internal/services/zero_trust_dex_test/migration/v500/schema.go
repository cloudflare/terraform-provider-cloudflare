package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareDeviceDexTestSchema returns the legacy cloudflare_device_dex_test schema (schema_version=0).
// This is used by MoveState and UpgradeFromLegacyV3 to parse state from the legacy SDKv2 provider.
//
// Schema version 0 is the default for SDKv2 resources (no explicit version set).
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_device_dex_tests.go
//
// This minimal schema includes only the properties needed for state parsing (Required, Optional, Computed).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceCloudflareDeviceDexTestSchema() schema.Schema {
	return schema.Schema{
		// Version 0 is implicit for SDKv2 resources without explicit SchemaVersion
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
			"description": schema.StringAttribute{
				Required: true,
			},
			"interval": schema.StringAttribute{
				Required: true,
			},
			"enabled": schema.BoolAttribute{
				Required: true,
			},
			// Computed timestamp fields (removed in v5)
			"updated": schema.StringAttribute{
				Computed: true,
			},
			"created": schema.StringAttribute{
				Computed: true,
			},
		},
		Blocks: map[string]schema.Block{
			// In SDKv2, TypeList with Elem: &schema.Resource is represented as a ListNestedBlock
			// MaxItems: 1 is stored as an array with one element in state
			"data": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"kind": schema.StringAttribute{
							Required: true,
						},
						"host": schema.StringAttribute{
							Required: true,
						},
						"method": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}
