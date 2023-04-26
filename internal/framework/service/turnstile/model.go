package turnstile

import "github.com/hashicorp/terraform-plugin-framework/types"

type TurnstileWidgetModel struct {
	AccountID    types.String `tfsdk:"account_id"`
	ID           types.String `tfsdk:"id"`
	Domains      types.Set    `tfsdk:"domains"`
	Name         types.String `tfsdk:"name"`
	Secret       types.String `tfsdk:"secret"`
	Region       types.String `tfsdk:"region"`
	Mode         types.String `tfsdk:"mode"`
	BotFightMode types.Bool   `tfsdk:"bot_fight_mode"`
	OffLabel     types.Bool   `tfsdk:"offlabel"`
}
