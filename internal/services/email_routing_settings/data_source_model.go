// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_settings

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingSettingsResultDataSourceEnvelope struct {
	Result EmailRoutingSettingsDataSourceModel `json:"result,computed"`
}

type EmailRoutingSettingsDataSourceModel struct {
	ZoneIdentifier types.String      `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String      `tfsdk:"id" json:"id"`
	Created        timetypes.RFC3339 `tfsdk:"created" json:"created"`
	Enabled        types.Bool        `tfsdk:"enabled" json:"enabled"`
	Modified       timetypes.RFC3339 `tfsdk:"modified" json:"modified"`
	Name           types.String      `tfsdk:"name" json:"name"`
	SkipWizard     types.Bool        `tfsdk:"skip_wizard" json:"skip_wizard"`
	Status         types.String      `tfsdk:"status" json:"status"`
	Tag            types.String      `tfsdk:"tag" json:"tag"`
}
