package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceHealthcheckSchema returns the source schema for legacy healthcheck resource.
// Schema version: 0 (implicit in v4 - no Version field specified)
// Resource type: cloudflare_healthcheck
//
// IMPORTANT: This schema represents the v4 FLAT structure.
// All HTTP/TCP-specific fields (method, port, path, etc.) are at the ROOT level in v4.
// The v5 schema nests these into http_config or tcp_config objects.
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed, ElementType).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceHealthcheckSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // v4 schema has no explicit Version field (defaults to 0)
		Attributes: map[string]schema.Attribute{
			// Core identity fields
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},

			// Required fields
			"name": schema.StringAttribute{
				Required: true,
			},
			"address": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},

			// Optional configuration fields
			"description": schema.StringAttribute{
				Optional: true,
			},
			"check_regions": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
			"suspended": schema.BoolAttribute{
				Optional: true,
			},
			"consecutive_fails": schema.Int64Attribute{
				Optional: true,
			},
			"consecutive_successes": schema.Int64Attribute{
				Optional: true,
			},
			"interval": schema.Int64Attribute{
				Optional: true,
			},
			"retries": schema.Int64Attribute{
				Optional: true,
			},
			"timeout": schema.Int64Attribute{
				Optional: true,
			},

			// FLAT HTTP/TCP fields at ROOT level (v4 structure)
			// These will move into http_config or tcp_config in v5
			"method": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"port": schema.Int64Attribute{
				Optional: true,
			},
			"path": schema.StringAttribute{
				Optional: true,
			},
			"expected_codes": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"expected_body": schema.StringAttribute{
				Optional: true,
			},
			"follow_redirects": schema.BoolAttribute{
				Optional: true,
			},
			"allow_insecure": schema.BoolAttribute{
				Optional: true,
			},

			// Header as Set of nested objects (v4 format)
			// v4: Set of {header: string, values: []string}
			// v5: Map[string][]string
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

			// Computed/output fields
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
