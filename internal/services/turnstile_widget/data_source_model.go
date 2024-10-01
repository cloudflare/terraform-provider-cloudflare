// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/turnstile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TurnstileWidgetResultDataSourceEnvelope struct {
	Result TurnstileWidgetDataSourceModel `json:"result,computed"`
}

type TurnstileWidgetResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[TurnstileWidgetDataSourceModel] `json:"result,computed"`
}

type TurnstileWidgetDataSourceModel struct {
	AccountID      types.String                             `tfsdk:"account_id" path:"account_id,optional"`
	Sitekey        types.String                             `tfsdk:"sitekey" path:"sitekey,computed_optional"`
	Secret         types.String                             `tfsdk:"secret" json:"secret,optional"`
	BotFightMode   types.Bool                               `tfsdk:"bot_fight_mode" json:"bot_fight_mode,computed"`
	ClearanceLevel types.String                             `tfsdk:"clearance_level" json:"clearance_level,computed"`
	CreatedOn      timetypes.RFC3339                        `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	EphemeralID    types.Bool                               `tfsdk:"ephemeral_id" json:"ephemeral_id,computed"`
	Mode           types.String                             `tfsdk:"mode" json:"mode,computed"`
	ModifiedOn     timetypes.RFC3339                        `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name           types.String                             `tfsdk:"name" json:"name,computed"`
	Offlabel       types.Bool                               `tfsdk:"offlabel" json:"offlabel,computed"`
	Region         types.String                             `tfsdk:"region" json:"region,computed"`
	Domains        customfield.List[types.String]           `tfsdk:"domains" json:"domains,computed"`
	Filter         *TurnstileWidgetFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *TurnstileWidgetDataSourceModel) toReadParams(_ context.Context) (params turnstile.WidgetGetParams, diags diag.Diagnostics) {
	params = turnstile.WidgetGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *TurnstileWidgetDataSourceModel) toListParams(_ context.Context) (params turnstile.WidgetListParams, diags diag.Diagnostics) {
	params = turnstile.WidgetListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(turnstile.WidgetListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(turnstile.WidgetListParamsOrder(m.Filter.Order.ValueString()))
	}

	return
}

type TurnstileWidgetFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	Direction types.String `tfsdk:"direction" query:"direction,optional"`
	Order     types.String `tfsdk:"order" query:"order,optional"`
}
