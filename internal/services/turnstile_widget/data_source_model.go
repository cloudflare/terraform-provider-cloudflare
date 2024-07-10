// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TurnstileWidgetResultDataSourceEnvelope struct {
	Result TurnstileWidgetDataSourceModel `json:"result,computed"`
}

type TurnstileWidgetResultListDataSourceEnvelope struct {
	Result *[]*TurnstileWidgetDataSourceModel `json:"result,computed"`
}

type TurnstileWidgetDataSourceModel struct {
	AccountID      types.String                             `tfsdk:"account_id" path:"account_id"`
	Sitekey        types.String                             `tfsdk:"sitekey" json:"sitekey"`
	BotFightMode   types.Bool                               `tfsdk:"bot_fight_mode" json:"bot_fight_mode,computed"`
	ClearanceLevel types.String                             `tfsdk:"clearance_level" json:"clearance_level,computed"`
	CreatedOn      types.String                             `tfsdk:"created_on" json:"created_on,computed"`
	Domains        types.String                             `tfsdk:"domains" json:"domains,computed"`
	Mode           types.String                             `tfsdk:"mode" json:"mode,computed"`
	ModifiedOn     types.String                             `tfsdk:"modified_on" json:"modified_on,computed"`
	Name           types.String                             `tfsdk:"name" json:"name,computed"`
	Offlabel       types.Bool                               `tfsdk:"offlabel" json:"offlabel,computed"`
	Region         types.String                             `tfsdk:"region" json:"region,computed"`
	Secret         types.String                             `tfsdk:"secret" json:"secret"`
	FindOneBy      *TurnstileWidgetFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type TurnstileWidgetFindOneByDataSourceModel struct {
	AccountID types.String  `tfsdk:"account_id" path:"account_id"`
	Direction types.String  `tfsdk:"direction" query:"direction"`
	Order     types.String  `tfsdk:"order" query:"order"`
	Page      types.Float64 `tfsdk:"page" query:"page"`
	PerPage   types.Float64 `tfsdk:"per_page" query:"per_page"`
}
