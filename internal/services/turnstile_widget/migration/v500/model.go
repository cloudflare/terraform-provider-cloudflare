package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudfareTurnstileWidgetModel represents the source cloudflare_turnstile_widget state structure.
// This corresponds to the legacy (plugin-framework v4) cloudflare provider.
// Schema version: 0 (implicit in v4)
// Resource type: cloudflare_turnstile_widget
//
// Used by UpgradeFromV4 to parse v4 state and migrate to v5 (schema version 500).
type SourceCloudfareTurnstileWidgetModel struct {
	ID           types.String `tfsdk:"id"`
	AccountID    types.String `tfsdk:"account_id"`
	Secret       types.String `tfsdk:"secret"`
	Name         types.String `tfsdk:"name"`
	Domains      types.Set    `tfsdk:"domains"` // v4: SetAttribute - will convert to List in v5
	Mode         types.String `tfsdk:"mode"`
	Region       types.String `tfsdk:"region"`
	BotFightMode types.Bool   `tfsdk:"bot_fight_mode"`
	OffLabel     types.Bool   `tfsdk:"offlabel"` // Note: field name case - v4 model uses OffLabel
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetTurnstileWidgetModel represents the target cloudflare_turnstile_widget state structure (v500).
// This corresponds to the current provider implementation.
// Schema version: 500 (target)
// Resource type: cloudflare_turnstile_widget (same name, no rename)
//
// Key differences from source (v4):
// - Domains: types.Set → *[]types.String (Set to ordered List with pointer)
// - Sitekey: NEW field in v5 (in v4, sitekey was stored as the ID field)
// - New computed fields: CreatedOn, ModifiedOn, ClearanceLevel, EphemeralID
// - OffLabel → Offlabel (field name only, tfsdk tag is same)
//
// CRITICAL: During migration, Sitekey MUST be copied from source.ID because in v4,
// the sitekey value was stored as the resource ID. Setting it to null will cause
// "missing required sitekey parameter" errors from the API.
//
// Note: This must match the model in turnstile_widget/model.go exactly.
// We duplicate it here to keep the migration package self-contained.
type TargetTurnstileWidgetModel struct {
	ID             types.String      `tfsdk:"id"`
	Sitekey        types.String      `tfsdk:"sitekey"`         // NEW in v5 (computed)
	AccountID      types.String      `tfsdk:"account_id"`
	Mode           types.String      `tfsdk:"mode"`
	Name           types.String      `tfsdk:"name"`
	Domains        *[]types.String   `tfsdk:"domains"`         // CHANGED: Set → *[]types.String
	BotFightMode   types.Bool        `tfsdk:"bot_fight_mode"`
	ClearanceLevel types.String      `tfsdk:"clearance_level"` // NEW in v5 (optional computed)
	EphemeralID    types.Bool        `tfsdk:"ephemeral_id"`    // NEW in v5 (optional computed)
	Offlabel       types.Bool        `tfsdk:"offlabel"`
	Region         types.String      `tfsdk:"region"`
	CreatedOn      timetypes.RFC3339 `tfsdk:"created_on"`  // NEW in v5 (computed timestamp)
	ModifiedOn     timetypes.RFC3339 `tfsdk:"modified_on"` // NEW in v5 (computed timestamp)
	Secret         types.String      `tfsdk:"secret"`
}
