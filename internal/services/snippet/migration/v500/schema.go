package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareSnippetSchema returns the source schema for the legacy cloudflare_snippet resource.
// Schema version: 1 (v4 SDKv2 provider had explicit schema version 1)
// Resource type: cloudflare_snippet
//
// This minimal schema is used only for reading v4 state during migration.
// It includes only the properties needed for state parsing (Required, Optional, Computed).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
//
// In v4, the snippet resource had:
// - name (Required): The snippet name
// - zone_id (Required): The zone ID
// - main_module (Required): The main module filename
// - files: List of objects with name/content (stored as array in SDKv2 state attributes)
// - created_on, modified_on (Computed): Timestamps
//
// Note: files is defined as ListNestedAttribute (not ListNestedBlock) because SDKv2
// stores block data in the attributes section of the state JSON. The Plugin Framework
// reads ListNestedAttribute from the attributes section, matching where SDKv2 placed it.
func SourceCloudflareSnippetSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"main_module": schema.StringAttribute{
				Required: true,
			},
			"files": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional: true,
						},
						"content": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
