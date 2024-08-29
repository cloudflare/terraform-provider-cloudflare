// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_settings

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingSettingsResultEnvelope struct {
	Result EmailRoutingSettingsModel `json:"result"`
}

type EmailRoutingSettingsModel struct {
	ID         types.String      `tfsdk:"id" json:"-,computed"`
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id"`
	Created    timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Enabled    types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	Modified   timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name       types.String      `tfsdk:"name" json:"name,computed"`
	SkipWizard types.Bool        `tfsdk:"skip_wizard" json:"skip_wizard,computed"`
	Status     types.String      `tfsdk:"status" json:"status,computed"`
	Tag        types.String      `tfsdk:"tag" json:"tag,computed"`
}
