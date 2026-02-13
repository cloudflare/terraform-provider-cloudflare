package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareAPITokenSchema returns the v4 SDKv2 api_token schema for state parsing.
//
// In v4, "policy" and "condition" were TypeList/TypeSet blocks (stored as JSON arrays).
// We use ListNestedBlock so the Plugin Framework can correctly deserialize v4 state arrays.
//
// Key differences from v5:
//   - "policy" (block, array) instead of "policies" (set attribute)
//   - permission_groups: Set of strings instead of Set of objects
//   - resources: Map of strings instead of JSON string
//   - condition/request_ip: List blocks (arrays) instead of SingleNestedAttribute (objects)
//   - policy.id: computed field that was removed in v5
func SourceCloudflareAPITokenSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"status": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"value": schema.StringAttribute{
				Computed:  true,
				Sensitive: true,
			},
			"issued_on": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},
			"expires_on": schema.StringAttribute{
				Optional: true,
			},
			"not_before": schema.StringAttribute{
				Optional: true,
			},
		},
		Blocks: map[string]schema.Block{
			// v4: "policy" was a TypeSet block (stored as array in state)
			"policy": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"effect": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						// v4: permission_groups was TypeSet of strings
						"permission_groups": schema.SetAttribute{
							ElementType: types.StringType,
							Required:    true,
						},
						// v4: resources was TypeMap of strings
						"resources": schema.MapAttribute{
							ElementType: types.StringType,
							Required:    true,
						},
					},
				},
			},
			// v4: "condition" was a TypeList block with MaxItems:1
			"condition": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Blocks: map[string]schema.Block{
						// v4: "request_ip" was a TypeList block with MaxItems:1
						"request_ip": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"in": schema.ListAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
									"not_in": schema.ListAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
