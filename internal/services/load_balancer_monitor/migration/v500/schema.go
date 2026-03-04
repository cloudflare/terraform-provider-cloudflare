// Package v500 implements state migration from legacy provider (v4) to current provider (v5)
// for the cloudflare_load_balancer_monitor resource.
package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceLoadBalancerMonitorSchema returns the source schema for the legacy load_balancer_monitor resource.
// Schema version: 0 (implicit in SDKv2, no explicit version field)
// Resource type: cloudflare_load_balancer_monitor
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed, ElementType).
// Validators, PlanModifiers, Defaults, and Descriptions are intentionally omitted.
func SourceLoadBalancerMonitorSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // SDKv2 implicit version (no explicit version in v4 schema)
		Attributes: map[string]schema.Attribute{
			// ===============================================
			// Required Fields
			// ===============================================
			"account_id": schema.StringAttribute{
				Required: true,
			},

			// ===============================================
			// Computed Fields
			// ===============================================
			"id": schema.StringAttribute{
				Computed: true,
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},

			// ===============================================
			// Optional Numeric Fields (TypeInt in v4)
			// ===============================================
			"interval": schema.Int64Attribute{
				Optional: true,
			},
			"port": schema.Int64Attribute{
				Optional: true,
			},
			"retries": schema.Int64Attribute{
				Optional: true,
			},
			"timeout": schema.Int64Attribute{
				Optional: true,
			},
			"consecutive_down": schema.Int64Attribute{
				Optional: true,
			},
			"consecutive_up": schema.Int64Attribute{
				Optional: true,
			},

			// ===============================================
			// Optional String Fields
			// ===============================================
			"description": schema.StringAttribute{
				Optional: true,
			},
			"expected_body": schema.StringAttribute{
				Optional: true,
			},
			"expected_codes": schema.StringAttribute{
				Optional: true,
			},
			"probe_zone": schema.StringAttribute{
				Optional: true,
			},
			"type": schema.StringAttribute{
				Optional: true,
			},

			// ===============================================
			// Optional + Computed String Fields
			// ===============================================
			"method": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"path": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},

			// ===============================================
			// Optional Boolean Fields
			// ===============================================
			"allow_insecure": schema.BoolAttribute{
				Optional: true,
			},
			"follow_redirects": schema.BoolAttribute{
				Optional: true,
			},

			// ===============================================
			// Complex Nested Field: Header (TypeSet in v4)
			// ===============================================
			// In v4, this is a TypeSet with nested Resource schema.
			// Schema structure:
			//   header {
			//     header = "Host"
			//     values = ["example.com"]
			//   }
			// State storage: [{"header": "Host", "values": ["example.com"]}]
			//
			// In v5, this becomes a MapAttribute with list values.
			"header": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"header": schema.StringAttribute{
							Required: true,
						},
						"values": schema.SetAttribute{
							ElementType: types.StringType,
							Required:    true,
						},
					},
				},
			},
		},
	}
}
