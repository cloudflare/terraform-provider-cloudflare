// File generated for StateUpgrader migration from v4 to v5

package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// SourceWorkersKVNamespaceSchema returns the v4 schema definition
// This matches the v4 SDKv2 provider schema structure
func SourceWorkersKVNamespaceSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // v4 schema version (SDKv2 had implicit version 0)
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"title": schema.StringAttribute{
				Required: true,
			},
		},
	}
}
