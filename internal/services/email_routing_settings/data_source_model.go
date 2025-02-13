// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/email_routing"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingSettingsResultDataSourceEnvelope struct {
	Result EmailRoutingSettingsDataSourceModel `json:"result,computed"`
}

type EmailRoutingSettingsDataSourceModel struct {
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Created    timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Enabled    types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	ID         types.String      `tfsdk:"id" json:"id,computed"`
	Modified   timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name       types.String      `tfsdk:"name" json:"name,computed"`
	SkipWizard types.Bool        `tfsdk:"skip_wizard" json:"skip_wizard,computed"`
	Status     types.String      `tfsdk:"status" json:"status,computed"`
	Tag        types.String      `tfsdk:"tag" json:"tag,computed"`
}

func (m *EmailRoutingSettingsDataSourceModel) toReadParams(_ context.Context) (params email_routing.EmailRoutingGetParams, diags diag.Diagnostics) {
	params = email_routing.EmailRoutingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
