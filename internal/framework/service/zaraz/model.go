package zaraz

import "github.com/hashicorp/terraform-plugin-framework/types"

type ZarazConfigModel struct {
	AccountId types.String `tfsdk:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id"`
	Config    ZarazConfig  `tfsdk:"config"`
}

type ZarazConfig struct {
	DebugKey      types.String    `tfsdk:"debug_key"`
	DefaultFields types.Map       `tfsdk:"default_fields"`
	Tools         map[string]Tool `tfsdk:"tools"`
}

type Tool struct {
	Worker  types.Map    `tfsdk:"worker"`
	Type    types.String `tfsdk:"type"`
	Actions types.Map    `tfsdk:"actions"`
	Mode    ToolMode     `tfsdk:"mode,omitempty"`
}

type ModeSegment struct {
	Start types.Number `tfsdk:"start"`
	End   types.Number `tfsdk:"end"`
}

type ToolMode struct {
	Light      types.Bool  `tfsdk:"light"`
	Cloud      types.Bool  `tfsdk:"cloud"`
	Sample     types.Bool  `tfsdk:"sample"`
	Segment    ModeSegment `tfsdk:"segment"`
	Trigger    []string    `tfsdk:"trigger"`
	IignoreSPA types.Bool  `tfsdk:"ignore_spa,omitempty"`
}
