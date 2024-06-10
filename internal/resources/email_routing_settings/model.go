// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_settings

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingSettingsResultEnvelope struct {
	Result EmailRoutingSettingsModel `json:"result,computed"`
}

type EmailRoutingSettingsModel struct {
	ID             types.String `tfsdk:"id" json:"id"`
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
}
