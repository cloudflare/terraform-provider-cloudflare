package zaraz

import "github.com/hashicorp/terraform-plugin-framework/types"

type ZarazConfigModel struct {
	AccountId     types.String `tfsdk:"account_id"`
	ZoneID        types.String `tfsdk:"zone_id"`
	DebugKey      types.String `tfsdk:"debugKey"`
	DefaultFields types.Map    `tfsdk:"default_fields"`
	Tools         []Tool       `tfsdk:"tools"`
}

type Tool struct {
	Worker  types.Map    `tfsdk:"worker"`
	Type    types.String `tfsdk:"type"`
	Actions types.Map    `tfsdk:"actions"`
}
