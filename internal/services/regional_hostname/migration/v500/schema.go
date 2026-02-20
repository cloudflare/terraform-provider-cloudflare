package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// SourceRegionalHostnameSchema returns the source schema for legacy cloudflare_regional_hostname resource.
// Schema version: 0 (v4 provider default)
// Resource type: cloudflare_regional_hostname
//
// This minimal schema is used only for reading v4 state during migration.
// Fields are identical between v4 and v5.
func SourceRegionalHostnameSchema() schema.Schema {
	return schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"hostname": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"routing": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"region_key": schema.StringAttribute{
				Required: true,
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
