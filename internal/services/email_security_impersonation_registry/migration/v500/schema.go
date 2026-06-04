package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// PriorSchema describes the schema_version=0 layout (id was Int64).
func PriorSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id":                         schema.Int64Attribute{Computed: true},
			"account_id":                 schema.StringAttribute{Optional: true},
			"email":                      schema.StringAttribute{Required: true},
			"is_email_regex":             schema.BoolAttribute{Required: true},
			"name":                       schema.StringAttribute{Required: true},
			"comments":                   schema.StringAttribute{Computed: true},
			"created_at":                 schema.StringAttribute{Computed: true, CustomType: timetypes.RFC3339Type{}},
			"directory_id":               schema.Int64Attribute{Computed: true},
			"directory_node_id":          schema.Int64Attribute{Computed: true},
			"external_directory_node_id": schema.StringAttribute{Computed: true},
			"last_modified":              schema.StringAttribute{Computed: true, CustomType: timetypes.RFC3339Type{}},
			"provenance":                 schema.StringAttribute{Computed: true},
		},
	}
}
