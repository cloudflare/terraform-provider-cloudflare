package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceTeamsListSchema returns the v4 cloudflare_teams_list schema (schema_version=0).
// Used by both MoveState and UpgradeState[0] for reading v4 state.
func SourceTeamsListSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"type": schema.StringAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			// V4 items stored as string list in SDKv2 state
			"items": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			// V4 items_with_description stored as list of objects in SDKv2 state
			"items_with_description": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"value":       schema.StringAttribute{Optional: true},
						"description": schema.StringAttribute{Optional: true},
					},
				},
			},
		},
	}
}
