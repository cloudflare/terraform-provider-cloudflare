package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
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

// TargetZeroTrustDLPCustomProfileModel represents the current resource state.
// Schema version: 500
// Resource type: cloudflare_zero_trust_dlp_custom_profile
type TargetZeroTrustDLPCustomProfileModel struct {
	ID                  types.String                     `tfsdk:"id"`
	AccountID           types.String                     `tfsdk:"account_id"`
	Name                types.String                     `tfsdk:"name"`
	Description         types.String                     `tfsdk:"description"`
	DataClasses         *[]types.String                  `tfsdk:"data_classes"`
	DataTags            *[]types.String                  `tfsdk:"data_tags"`
	ContextAwareness    *TargetContextAwarenessModel     `tfsdk:"context_awareness"`
	Entries             *[]*TargetEntriesModel           `tfsdk:"entries"`
	SensitivityLevels   *[]*TargetSensitivityLevelsModel `tfsdk:"sensitivity_levels"`
	SharedEntries       *[]*TargetSharedEntriesModel     `tfsdk:"shared_entries"`
	AIContextEnabled    types.Bool                       `tfsdk:"ai_context_enabled"`
	AllowedMatchCount   types.Int64                      `tfsdk:"allowed_match_count"`
	ConfidenceThreshold types.String                     `tfsdk:"confidence_threshold"`
	OCREnabled          types.Bool                       `tfsdk:"ocr_enabled"`
	CreatedAt           timetypes.RFC3339                `tfsdk:"created_at"`
	OpenAccess          types.Bool                       `tfsdk:"open_access"`
	Type                types.String                     `tfsdk:"type"`
	UpdatedAt           timetypes.RFC3339                `tfsdk:"updated_at"`
}

// TargetContextAwarenessModel represents context_awareness in v5.
type TargetContextAwarenessModel struct {
	Enabled types.Bool                       `tfsdk:"enabled"`
	Skip    *TargetContextAwarenessSkipModel `tfsdk:"skip"`
}

// TargetContextAwarenessSkipModel represents context_awareness.skip in v5.
type TargetContextAwarenessSkipModel struct {
	Files types.Bool `tfsdk:"files"`
}

// TargetEntriesModel represents a single entry in v5.
type TargetEntriesModel struct {
	Description types.String        `tfsdk:"description"`
	Enabled     types.Bool          `tfsdk:"enabled"`
	EntryID     types.String        `tfsdk:"entry_id"`
	Name        types.String        `tfsdk:"name"`
	Pattern     *TargetPatternModel `tfsdk:"pattern"`
}

// TargetPatternModel represents a pattern in a v5 entry.
type TargetPatternModel struct {
	Regex      types.String `tfsdk:"regex"`
	Validation types.String `tfsdk:"validation"`
}

// TargetSharedEntriesModel represents a shared entry in v5.
type TargetSharedEntriesModel struct {
	Enabled   types.Bool   `tfsdk:"enabled"`
	EntryID   types.String `tfsdk:"entry_id"`
	EntryType types.String `tfsdk:"entry_type"`
}

type TargetSensitivityLevelsModel struct {
	GroupID types.String `tfsdk:"group_id"`
	LevelID types.String `tfsdk:"level_id"`
}
