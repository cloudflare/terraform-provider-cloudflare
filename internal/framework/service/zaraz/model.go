package zaraz

import "github.com/hashicorp/terraform-plugin-framework/types"

type ZarazConfigModel struct {
	AccountId types.String `tfsdk:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id"`
	Config    *ZarazConfig `tfsdk:"config"`
}

type ZarazLoadRuleOp string

type ZarazRuleType string

type ZarazTriggerSystem string

type ZarazSelectorType string

type ZarazToolType string

type ZarazRuleSettings struct {
	Type        *ZarazSelectorType `json:"type"`
	Selector    types.String       `json:"selector"`
	WaitForTags types.Int64        `json:"wait_for_tags"`
	Interval    types.Int64        `json:"interval"`
	Limit       types.Int64        `json:"limit"`
	Validate    *types.Bool        `json:"validate"`
	Variable    types.String       `json:"variable"`
	Match       types.String       `json:"match"`
	Positions   types.String       `json:"positions"`
	Op          *ZarazLoadRuleOp   `json:"op"`
	Value       types.String       `json:"value"`
}

type ZarazTriggerRule struct {
	Id       types.String      `tfsdk:"id"`
	Match    types.String       `tfsdk:"match"`
	Op       ZarazLoadRuleOp    `tfsdk:"op"`
	Value    types.String       `tfsdk:"value"`
	Action   ZarazRuleType      `tfsdk:"action"`
	Settings *ZarazRuleSettings `tfsdk:"settings"`
}

type ZarazTrigger struct {
	Name         types.String       `tfsdk:"name"`
	Description  types.String       `tfsdk:"description"`
	LoadRules    []ZarazTriggerRule `tfsdk:"load_rules"`
	ExcludeRules []ZarazTriggerRule `tfsdk:"exclude_rules"`
	System       ZarazTriggerSystem `tfsdk:"system"`
}

type ZarazConfig struct {
	DebugKey     types.String            `tfsdk:"debug_key"`
	ZarazVersion types.Int64             `tfsdk:"zaraz_version"`
	Tools        map[string]ZarazTool    `tfsdk:"tools"`
	Triggers     map[string]ZarazTrigger `tfsdk:"triggers"`
}

type ZarazWorker struct {
	EscapedWorkerName string       `tfsdk:"escaped_worker_name"`
	WorkerTag         string       `tfsdk:"worker_tag"`
	MutableId         types.String `tfsdk:"mutable_id"`
}

type ZarazTool struct {
	DefaultFields    map[string]any         `tfsdk:"default_fields"`
	Worker           *ZarazWorker           `tfsdk:"worker"`
	Type             types.String           `tfsdk:"type"`
	Actions          map[string]ZarazAction `tfsdk:"actions"`
	NeoEvents        []ZarazAction          `tfsdk:"neo_events"`
	Mode             *ToolMode              `tfsdk:"mode"`
	Name             types.String           `tfsdk:"name"`
	BlockingTriggers []string               `tfsdk:"blocking_triggers"`
	Enabled          types.Bool             `tfsdk:"enabled"`
	DefaultPurpose   types.String           `tfsdk:"default_purpose"`
	Library          types.String           `tfsdk:"library"`
	Component        types.String           `tfsdk:"component"`
	Permissions      []string               `tfsdk:"permissions"`
	Settings         map[string]any         `tfsdk:"settings"`
}

type ModeSegment struct {
	Start types.Number `tfsdk:"start"`
	End   types.Number `tfsdk:"end"`
}

type ToolMode struct {
	Light      types.Bool     `tfsdk:"light"`
	Cloud      types.Bool     `tfsdk:"cloud"`
	Sample     types.Bool     `tfsdk:"sample"`
	Segment    *ModeSegment   `tfsdk:"segment"`
	Trigger    []types.String `tfsdk:"trigger"`
	IignoreSPA types.Bool     `tfsdk:"ignore_spa"`
}

type ZarazAction struct {
	BlockingTriggers []types.String `tfsdk:"blocking_triggers"`
	FiringTriggers   []types.String `tfsdk:"firing_triggers"`
	Data             map[string]any `tfsdk:"data"`
	ActionType       types.String   `tfsdk:"action_type"`
}
