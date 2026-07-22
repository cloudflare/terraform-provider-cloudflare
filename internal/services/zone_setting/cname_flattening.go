package zone_setting

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const CNAMEFlatteningSettingID = "cname_flattening"

const (
	cnameFlatteningDeprecationSummary = "The cname_flattening zone setting is no longer supported"
	cnameFlatteningDeprecationDetail  = "Use cloudflare_zone_dns_settings.flatten_all_cnames instead. " +
		"Map on or flatten_all to true, and off, apex, or flatten_at_root to false. " +
		"See the CNAME flattening migration guide for state migration instructions."
)

func isCNAMEFlatteningSetting(settingID types.String) bool {
	return !settingID.IsNull() && !settingID.IsUnknown() && strings.EqualFold(settingID.ValueString(), CNAMEFlatteningSettingID)
}

func addCNAMEFlatteningDeprecationError(diags *diag.Diagnostics) {
	diags.AddError(cnameFlatteningDeprecationSummary, cnameFlatteningDeprecationDetail)
}
