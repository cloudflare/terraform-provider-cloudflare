// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TurnstileWidgetResultEnvelope struct {
	Result TurnstileWidgetModel `json:"result,computed"`
}

type TurnstileWidgetModel struct {
	ID             types.String      `tfsdk:"id" json:"-,computed"`
	Sitekey        types.String      `tfsdk:"sitekey" json:"sitekey,computed"`
	AccountID      types.String      `tfsdk:"account_id" path:"account_id"`
	Region         types.String      `tfsdk:"region" json:"region"`
	Mode           types.String      `tfsdk:"mode" json:"mode"`
	Name           types.String      `tfsdk:"name" json:"name"`
	Domains        *[]types.String   `tfsdk:"domains" json:"domains"`
	BotFightMode   types.Bool        `tfsdk:"bot_fight_mode" json:"bot_fight_mode"`
	ClearanceLevel types.String      `tfsdk:"clearance_level" json:"clearance_level"`
	Offlabel       types.Bool        `tfsdk:"offlabel" json:"offlabel"`
	CreatedOn      timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn     timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed"`
	Secret         types.String      `tfsdk:"secret" json:"secret,computed"`
}
