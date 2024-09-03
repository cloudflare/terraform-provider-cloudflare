// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_catch_all

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/email_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingCatchAllResultDataSourceEnvelope struct {
	Result EmailRoutingCatchAllDataSourceModel `json:"result,computed"`
}

type EmailRoutingCatchAllDataSourceModel struct {
	ZoneID   types.String                                    `tfsdk:"zone_id" path:"zone_id"`
	ID       types.String                                    `tfsdk:"id" json:"id"`
	Name     types.String                                    `tfsdk:"name" json:"name"`
	Tag      types.String                                    `tfsdk:"tag" json:"tag"`
	Actions  *[]*EmailRoutingCatchAllActionsDataSourceModel  `tfsdk:"actions" json:"actions"`
	Matchers *[]*EmailRoutingCatchAllMatchersDataSourceModel `tfsdk:"matchers" json:"matchers"`
	Enabled  types.Bool                                      `tfsdk:"enabled" json:"enabled,computed_optional"`
}

func (m *EmailRoutingCatchAllDataSourceModel) toReadParams(_ context.Context) (params email_routing.RuleCatchAllGetParams, diags diag.Diagnostics) {
	params = email_routing.RuleCatchAllGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type EmailRoutingCatchAllActionsDataSourceModel struct {
	Type  types.String                   `tfsdk:"type" json:"type,computed"`
	Value customfield.List[types.String] `tfsdk:"value" json:"value,computed"`
}

type EmailRoutingCatchAllMatchersDataSourceModel struct {
	Type types.String `tfsdk:"type" json:"type,computed"`
}
