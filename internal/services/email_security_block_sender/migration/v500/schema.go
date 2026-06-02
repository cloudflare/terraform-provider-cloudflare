package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// PriorSchema describes the schema_version=0 layout (id was Int64).
func PriorSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id":            schema.Int64Attribute{Computed: true},
			"account_id":    schema.StringAttribute{Optional: true},
			"is_regex":      schema.BoolAttribute{Required: true},
			"pattern":       schema.StringAttribute{Required: true},
			"pattern_type":  schema.StringAttribute{Required: true},
			"comments":      schema.StringAttribute{Optional: true},
			"created_at":    schema.StringAttribute{Computed: true, CustomType: timetypes.RFC3339Type{}},
			"last_modified": schema.StringAttribute{Computed: true, CustomType: timetypes.RFC3339Type{}},
		},
	}
}
