package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareDLPProfileModel represents the legacy DLP profile state from v4.x provider.
// Schema version: 0 (SDKv2 default)
// Resource types: cloudflare_dlp_profile, cloudflare_zero_trust_dlp_profile
type SourceCloudflareDLPProfileModel struct {
	ID                types.String                  `tfsdk:"id"`
	AccountID         types.String                  `tfsdk:"account_id"`
	Name              types.String                  `tfsdk:"name"`
	Description       types.String                  `tfsdk:"description"`
	Type              types.String                  `tfsdk:"type"`
	Entry             []SourceEntryModel            `tfsdk:"entry"`
	AllowedMatchCount types.Int64                   `tfsdk:"allowed_match_count"`
	ContextAwareness  []SourceContextAwarenessModel `tfsdk:"context_awareness"`
	OCREnabled        types.Bool                    `tfsdk:"ocr_enabled"`
}

// SourceEntryModel represents a single entry in the v4 DLP profile.
type SourceEntryModel struct {
	ID      types.String         `tfsdk:"id"`
	Name    types.String         `tfsdk:"name"`
	Enabled types.Bool           `tfsdk:"enabled"`
	Pattern []SourcePatternModel `tfsdk:"pattern"`
}

// SourcePatternModel represents a pattern in a v4 DLP profile entry.
type SourcePatternModel struct {
	Regex      types.String `tfsdk:"regex"`
	Validation types.String `tfsdk:"validation"`
}

// SourceContextAwarenessModel represents the context_awareness block in v4.
type SourceContextAwarenessModel struct {
	Enabled types.Bool                        `tfsdk:"enabled"`
	Skip    []SourceContextAwarenessSkipModel `tfsdk:"skip"`
}

// SourceContextAwarenessSkipModel represents context_awareness.skip in v4.
type SourceContextAwarenessSkipModel struct {
	Files types.Bool `tfsdk:"files"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetZeroTrustDLPPredefinedProfileModel represents the current resource state.
// Schema version: 500
// Resource type: cloudflare_zero_trust_dlp_predefined_profile
type TargetZeroTrustDLPPredefinedProfileModel struct {
	ID                  types.String                                                            `tfsdk:"id"`
	ProfileID           types.String                                                            `tfsdk:"profile_id"`
	AccountID           types.String                                                            `tfsdk:"account_id"`
	EnabledEntries      *[]types.String                                                         `tfsdk:"enabled_entries"`
	AIContextEnabled    types.Bool                                                              `tfsdk:"ai_context_enabled"`
	AllowedMatchCount   types.Int64                                                             `tfsdk:"allowed_match_count"`
	ConfidenceThreshold types.String                                                            `tfsdk:"confidence_threshold"`
	OCREnabled          types.Bool                                                              `tfsdk:"ocr_enabled"`
	Entries             customfield.NestedObjectList[TargetPredefinedProfileEntriesModel]        `tfsdk:"entries"`
	Name                types.String                                                            `tfsdk:"name"`
	OpenAccess          types.Bool                                                              `tfsdk:"open_access"`
}

// TargetPredefinedProfileEntriesModel represents an entry in v5 predefined profile.
type TargetPredefinedProfileEntriesModel struct {
	ID      types.String `tfsdk:"id"`
	Enabled types.Bool   `tfsdk:"enabled"`
}
