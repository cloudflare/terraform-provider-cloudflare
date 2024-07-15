// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_settings

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingSettingsResultEnvelope struct {
	Result EmailRoutingSettingsModel `json:"result,computed"`
}

type EmailRoutingSettingsModel struct {
	ID             types.String      `tfsdk:"id" json:"id"`
	ZoneIdentifier types.String      `tfsdk:"zone_identifier" path:"zone_identifier"`
	Created        timetypes.RFC3339 `tfsdk:"created" json:"created,computed"`
	Enabled        types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	Modified       timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed"`
	Name           types.String      `tfsdk:"name" json:"name,computed"`
	SkipWizard     types.Bool        `tfsdk:"skip_wizard" json:"skip_wizard,computed"`
	Status         types.String      `tfsdk:"status" json:"status,computed"`
	Tag            types.String      `tfsdk:"tag" json:"tag,computed"`
}
