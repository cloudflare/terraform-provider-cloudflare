package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// SourceZoneSettingSchema returns the schema used by the v5 cloudflare_zone_setting resource
// at schema_version=1 (the version written by early v5 releases before the version bump to 500).
//
// This schema is used as the PriorSchema for the version=1 → version=500 no-op upgrade.
// It must match exactly what was written to state by the v5 Plugin Framework provider.
//
// Note: The v4 cloudflare_zone_settings_override resource is a one-to-many transformation
// handled entirely by tf-migrate (config + state). The provider-side upgrader only needs
// to handle the v5 version bump (1 → 500).
func SourceZoneSettingSchema() schema.Schema {
	return schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"setting_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"value": schema.DynamicAttribute{
				Required:   true,
				CustomType: customfield.NormalizedDynamicType{},
				PlanModifiers: []planmodifier.Dynamic{
					customfield.NormalizeDynamicPlanModifier(),
				},
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"editable": schema.BoolAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"time_remaining": schema.Float64Attribute{
				Computed: true,
			},
		},
	}
}
