package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceSnippetRulesSchema returns the source schema for the legacy cloudflare_snippet_rules resource.
// Schema version: 1 (v4 SDKv2 provider)
// Resource type: cloudflare_snippet_rules
//
// This minimal schema is used only for reading v4 state during migration.
// Validators, PlanModifiers, Defaults, and Descriptions are intentionally omitted.
//
// Note: Uses ListNestedAttribute (not ListNestedBlock) because SDKv2 stores block data
// in the attributes section of the state JSON. The Plugin Framework reads
// ListNestedAttribute from the attributes section, matching where SDKv2 placed it.
func SourceSnippetRulesSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"rules": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"expression": schema.StringAttribute{
							Optional: true,
						},
						"snippet_name": schema.StringAttribute{
							Optional: true,
						},
						"enabled": schema.BoolAttribute{
							Optional: true,
						},
						"description": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}
