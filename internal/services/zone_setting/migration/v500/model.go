package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceZoneSettingModel is the model for cloudflare_zone_setting at schema_version=1.
//
// This is used during the no-op version bump (1 → 500). The schema is identical to
// the current v5 model — no field changes occurred between schema_version=1 and 500.
//
// Note: The v4 cloudflare_zone_settings_override → cloudflare_zone_setting migration
// is a one-to-many transformation handled by tf-migrate, not the provider state upgrader.
type SourceZoneSettingModel struct {
	ID            types.String                       `tfsdk:"id"`
	SettingID     types.String                       `tfsdk:"setting_id"`
	ZoneID        types.String                       `tfsdk:"zone_id"`
	Value         customfield.NormalizedDynamicValue `tfsdk:"value"`
	Enabled       types.Bool                         `tfsdk:"enabled"`
	Editable      types.Bool                         `tfsdk:"editable"`
	ModifiedOn    timetypes.RFC3339                  `tfsdk:"modified_on"`
	TimeRemaining types.Float64                      `tfsdk:"time_remaining"`
}
