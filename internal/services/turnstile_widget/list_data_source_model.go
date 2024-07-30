// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TurnstileWidgetsResultListDataSourceEnvelope struct {
	Result *[]*TurnstileWidgetsResultDataSourceModel `json:"result,computed"`
}

type TurnstileWidgetsDataSourceModel struct {
	AccountID types.String                              `tfsdk:"account_id" path:"account_id"`
	Direction types.String                              `tfsdk:"direction" query:"direction"`
	Order     types.String                              `tfsdk:"order" query:"order"`
	Page      types.Float64                             `tfsdk:"page" query:"page"`
	PerPage   types.Float64                             `tfsdk:"per_page" query:"per_page"`
	MaxItems  types.Int64                               `tfsdk:"max_items"`
	Result    *[]*TurnstileWidgetsResultDataSourceModel `tfsdk:"result"`
}

type TurnstileWidgetsResultDataSourceModel struct {
	BotFightMode   types.Bool        `tfsdk:"bot_fight_mode" json:"bot_fight_mode,computed"`
	ClearanceLevel types.String      `tfsdk:"clearance_level" json:"clearance_level,computed"`
	CreatedOn      timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed"`
	Domains        *[]types.String   `tfsdk:"domains" json:"domains,computed"`
	Mode           types.String      `tfsdk:"mode" json:"mode,computed"`
	ModifiedOn     timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed"`
	Name           types.String      `tfsdk:"name" json:"name,computed"`
	Offlabel       types.Bool        `tfsdk:"offlabel" json:"offlabel,computed"`
	Region         types.String      `tfsdk:"region" json:"region,computed"`
	Sitekey        types.String      `tfsdk:"sitekey" json:"sitekey,computed"`
}
