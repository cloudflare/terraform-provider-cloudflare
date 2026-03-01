package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// SourceManagedHeadersSchema returns the source schema for legacy cloudflare_managed_headers resource.
// Schema version: 0 (v4 provider default - managed_headers didn't set explicit version)
// Resource type: cloudflare_managed_headers
//
// This minimal schema is used only for reading v4 state during migration.
// Both managed_request_headers and managed_response_headers were optional in v4
// (could be null), but are required in v5 (empty sets allowed).
func SourceManagedHeadersSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // v4 provider default
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"managed_request_headers": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
						"enabled": schema.BoolAttribute{
							Required: true,
						},
					},
				},
			},
			"managed_response_headers": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
						"enabled": schema.BoolAttribute{
							Required: true,
						},
					},
				},
			},
		},
	}
}
