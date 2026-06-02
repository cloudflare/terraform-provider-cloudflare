package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// PriorSchema describes the v0 (pre-Version-tag) layout used to read raw state.
func PriorSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id":         schema.Int64Attribute{Computed: true},
			"account_id": schema.StringAttribute{Optional: true},
			"body": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"is_recent":     schema.BoolAttribute{Required: true},
						"is_regex":      schema.BoolAttribute{Required: true},
						"is_similarity": schema.BoolAttribute{Required: true},
						"pattern":       schema.StringAttribute{Required: true},
						"comments":      schema.StringAttribute{Optional: true},
					},
				},
			},
			"comments":      schema.StringAttribute{Optional: true},
			"is_recent":     schema.BoolAttribute{Optional: true},
			"is_regex":      schema.BoolAttribute{Optional: true},
			"is_similarity": schema.BoolAttribute{Optional: true},
			"pattern":       schema.StringAttribute{Optional: true},
			"created_at":    schema.StringAttribute{Computed: true, CustomType: timetypes.RFC3339Type{}},
			"last_modified": schema.StringAttribute{Computed: true, CustomType: timetypes.RFC3339Type{}},
		},
	}
}
