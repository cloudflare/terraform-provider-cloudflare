// Package v500 handles state migration from cloudflare_api_shield v4 (schema_version=0)
// to cloudflare_api_shield v5 (version=500).
package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceAPIShieldSchema returns the legacy cloudflare_api_shield schema (schema_version=0).
// This is used by UpgradeFromV4 to parse state from the legacy SDKv2 provider.
//
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_api_shield.go
//
// This minimal schema includes only the properties needed for state parsing:
//   - Required, Optional, Computed field markers
//   - Block structure definitions
//
// Intentionally excluded (not needed for reading state):
//   - Validators (state is already valid from v4)
//   - PlanModifiers (not planning, just reading)
//   - Descriptions (not shown to user during migration)
func SourceAPIShieldSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
		},
		Blocks: map[string]schema.Block{
			// In v4 SDKv2: TypeList with Elem: &schema.Resource (block syntax)
			// In v5 Framework: ListNestedAttribute (attribute syntax)
			//
			// SDKv2 TypeList with nested Resource is represented as ListNestedBlock in Framework
			// for source schema parsing.
			"auth_id_characteristics": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional: true,
						},
						"type": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}
