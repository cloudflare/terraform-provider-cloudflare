package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x SDKv2)
// ============================================================================

// SourceCloudflareCustomPagesModel represents the legacy custom_pages resource
// state from v4.x provider (SDKv2).
//
// Schema version: 0 (v4 SDKv2 resources don't have explicit schema versions)
// Resource type: cloudflare_custom_pages
//
// Key differences from v5:
// - Uses "type" field instead of "identifier"
// - state field is Optional (not Required)
// - url field is Required (not Optional+Computed)
// - Missing computed fields: created_on, modified_on, description, preview_target, required_tokens
type SourceCloudflareCustomPagesModel struct {
	ID        types.String `tfsdk:"id"`         // Computed in v4, also Computed in v5
	Type      types.String `tfsdk:"type"`       // Renamed to Identifier in v5
	State     types.String `tfsdk:"state"`      // Optional in v4, Required in v5
	URL       types.String `tfsdk:"url"`        // Required in v4, Optional+Computed in v5
	ZoneID    types.String `tfsdk:"zone_id"`    // Optional, mutually exclusive with account_id
	AccountID types.String `tfsdk:"account_id"` // Optional, mutually exclusive with zone_id
}

// ============================================================================
// Target Models (Current Provider - v5.x+ Plugin Framework)
// ============================================================================

// TargetCustomPagesModel represents the current custom_pages resource state
// from v5.x+ provider (Plugin Framework).
//
// Schema version: 500 (or 1 with migrations.schema version)
// Resource type: cloudflare_custom_pages
//
// This matches the structure in the parent package's model.go.
// We duplicate it here to keep the migration package self-contained.
type TargetCustomPagesModel struct {
	ID             types.String                   `tfsdk:"id"`              // NEW: Computed field (API-assigned)
	Identifier     types.String                   `tfsdk:"identifier"`      // Renamed from Type
	AccountID      types.String                   `tfsdk:"account_id"`      // Same as v4
	ZoneID         types.String                   `tfsdk:"zone_id"`         // Same as v4
	State          types.String                   `tfsdk:"state"`           // Now Required (was Optional)
	URL            types.String                   `tfsdk:"url"`             // Now Optional+Computed (was Required)
	CreatedOn      timetypes.RFC3339              `tfsdk:"created_on"`      // NEW: Computed timestamp
	Description    types.String                   `tfsdk:"description"`     // NEW: Computed field
	ModifiedOn     timetypes.RFC3339              `tfsdk:"modified_on"`     // NEW: Computed timestamp
	PreviewTarget  types.String                   `tfsdk:"preview_target"`  // NEW: Computed field
	RequiredTokens customfield.List[types.String] `tfsdk:"required_tokens"` // NEW: Computed list
}
