// Package v500 handles state migration from cloudflare_api_shield v4 (schema_version=0)
// to cloudflare_api_shield v5 (version=500).
package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x SDKv2)
// ============================================================================

// SourceAPIShieldModel represents the legacy cloudflare_api_shield resource state from v4.x provider.
// Schema version: 0 (implicit in SDKv2)
// Resource type: cloudflare_api_shield
//
// Key differences from v5:
//   - auth_id_characteristics is Optional (v5: Required)
//   - Nested fields (type, name) are Optional (v5: Required)
//   - Stores auth_id_characteristics as array in state (same as v5)
type SourceAPIShieldModel struct {
	ID                    types.String                             `tfsdk:"id"`
	ZoneID                types.String                             `tfsdk:"zone_id"`
	AuthIDCharacteristics *[]*SourceAuthIDCharacteristicsModel `tfsdk:"auth_id_characteristics"`
}

// SourceAuthIDCharacteristicsModel represents nested auth_id_characteristics structure from v4.x provider.
//
// Key differences from v5:
//   - type is Optional (v5: Required)
//   - name is Optional (v5: Required)
//   - v4 only supported ["header", "cookie"], v5 adds "jwt" (backward compatible)
type SourceAuthIDCharacteristicsModel struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+ Framework)
// ============================================================================

// TargetAPIShieldModel represents the current cloudflare_api_shield resource state from v5.x+ provider.
// Schema version: 500 (when TF_MIG_TEST=1, otherwise version 1)
// Resource type: cloudflare_api_shield
//
// Note: This matches the APIShieldModel in the parent package's model.go file.
// We duplicate it here to keep the migration package self-contained.
//
// Key differences from v4:
//   - auth_id_characteristics is Required (v4: Optional)
//   - Nested fields (type, name) are Required (v4: Optional)
//   - Supports additional type value "jwt" (backward compatible)
type TargetAPIShieldModel struct {
	ID                    types.String                             `tfsdk:"id"`
	ZoneID                types.String                             `tfsdk:"zone_id"`
	AuthIDCharacteristics *[]*TargetAuthIDCharacteristicsModel `tfsdk:"auth_id_characteristics"`
}

// TargetAuthIDCharacteristicsModel represents nested auth_id_characteristics structure from v5.x+ provider.
//
// Key differences from v4:
//   - type is Required (v4: Optional)
//   - name is Required (v4: Optional)
//   - Supports values: "header", "cookie", "jwt" (v4 only had header/cookie)
type TargetAuthIDCharacteristicsModel struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}
