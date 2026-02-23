// File generated for StateUpgrader migration from v4 to v5

package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// SourceWorkersForPlatformsNamespaceSchema returns the v4 schema definition.
// Both cloudflare_workers_for_platforms_namespace (deprecated) and
// cloudflare_workers_for_platforms_dispatch_namespace (current v4) had identical schemas.
func SourceWorkersForPlatformsNamespaceSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // v4 schema version (Plugin Framework resources default to version 0)
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}
