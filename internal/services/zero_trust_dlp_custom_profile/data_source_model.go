// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_custom_profile

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPCustomProfileResultDataSourceEnvelope struct {
	Result ZeroTrustDLPCustomProfileDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPCustomProfileDataSourceModel struct {
	ID                  types.String                                                                            `tfsdk:"id" path:"profile_id,computed"`
	ProfileID           types.String                                                                            `tfsdk:"profile_id" path:"profile_id,required"`
	AccountID           types.String                                                                            `tfsdk:"account_id" path:"account_id,required"`
	AIContextEnabled    types.Bool                                                                              `tfsdk:"ai_context_enabled" json:"ai_context_enabled,computed"`
	AllowedMatchCount   types.Int64                                                                             `tfsdk:"allowed_match_count" json:"allowed_match_count,computed"`
	ConfidenceThreshold types.String                                                                            `tfsdk:"confidence_threshold" json:"confidence_threshold,computed"`
	CreatedAt           timetypes.RFC3339                                                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description         types.String                                                                            `tfsdk:"description" json:"description,computed"`
	Name                types.String                                                                            `tfsdk:"name" json:"name,computed"`
	OCREnabled          types.Bool                                                                              `tfsdk:"ocr_enabled" json:"ocr_enabled,computed"`
	OpenAccess          types.Bool                                                                              `tfsdk:"open_access" json:"open_access,computed"`
	Type                types.String                                                                            `tfsdk:"type" json:"type,computed"`
	UpdatedAt           timetypes.RFC3339                                                                       `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	DataClasses         customfield.List[types.String]                                                          `tfsdk:"data_classes" json:"data_classes,computed"`
	DataTags            customfield.List[types.String]                                                          `tfsdk:"data_tags" json:"data_tags,computed"`
	ContextAwareness    customfield.NestedObject[ZeroTrustDLPCustomProfileContextAwarenessDataSourceModel]      `tfsdk:"context_awareness" json:"context_awareness,computed"`
	Entries             customfield.NestedObjectList[ZeroTrustDLPCustomProfileEntriesDataSourceModel]           `tfsdk:"entries" json:"entries,computed"`
	SensitivityLevels   customfield.NestedObjectList[ZeroTrustDLPCustomProfileSensitivityLevelsDataSourceModel] `tfsdk:"sensitivity_levels" json:"sensitivity_levels,computed"`
	SharedEntries       customfield.NestedObjectList[ZeroTrustDLPCustomProfileSharedEntriesDataSourceModel]     `tfsdk:"shared_entries" json:"shared_entries,computed"`
}

func (m *ZeroTrustDLPCustomProfileDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DLPProfileCustomGetParams, diags diag.Diagnostics) {
	params = zero_trust.DLPProfileCustomGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPCustomProfileContextAwarenessDataSourceModel struct {
	Enabled types.Bool                                                                             `tfsdk:"enabled" json:"enabled,computed"`
	Skip    customfield.NestedObject[ZeroTrustDLPCustomProfileContextAwarenessSkipDataSourceModel] `tfsdk:"skip" json:"skip,computed"`
}

type ZeroTrustDLPCustomProfileContextAwarenessSkipDataSourceModel struct {
	Files types.Bool `tfsdk:"files" json:"files,computed"`
}

type ZeroTrustDLPCustomProfileEntriesDataSourceModel struct {
	ID            types.String                                                                        `tfsdk:"id" json:"id,computed"`
	CreatedAt     timetypes.RFC3339                                                                   `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Enabled       types.Bool                                                                          `tfsdk:"enabled" json:"enabled,computed"`
	Name          types.String                                                                        `tfsdk:"name" json:"name,computed"`
	Pattern       customfield.NestedObject[ZeroTrustDLPCustomProfileEntriesPatternDataSourceModel]    `tfsdk:"pattern" json:"pattern,computed"`
	Type          types.String                                                                        `tfsdk:"type" json:"type,computed"`
	UpdatedAt     timetypes.RFC3339                                                                   `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Description   types.String                                                                        `tfsdk:"description" json:"description,computed"`
	ProfileID     types.String                                                                        `tfsdk:"profile_id" json:"profile_id,computed"`
	Confidence    customfield.NestedObject[ZeroTrustDLPCustomProfileEntriesConfidenceDataSourceModel] `tfsdk:"confidence" json:"confidence,computed"`
	Variant       customfield.NestedObject[ZeroTrustDLPCustomProfileEntriesVariantDataSourceModel]    `tfsdk:"variant" json:"variant,computed"`
	CaseSensitive types.Bool                                                                          `tfsdk:"case_sensitive" json:"case_sensitive,computed"`
	Secret        types.Bool                                                                          `tfsdk:"secret" json:"secret,computed"`
	WordList      jsontypes.Normalized                                                                `tfsdk:"word_list" json:"word_list,computed"`
}

type ZeroTrustDLPCustomProfileEntriesPatternDataSourceModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex,computed"`
	Validation types.String `tfsdk:"validation" json:"validation,computed"`
}

type ZeroTrustDLPCustomProfileEntriesConfidenceDataSourceModel struct {
	AIContextAvailable types.Bool `tfsdk:"ai_context_available" json:"ai_context_available,computed"`
	Available          types.Bool `tfsdk:"available" json:"available,computed"`
}

type ZeroTrustDLPCustomProfileEntriesVariantDataSourceModel struct {
	TopicType   types.String `tfsdk:"topic_type" json:"topic_type,computed"`
	Type        types.String `tfsdk:"type" json:"type,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
}

type ZeroTrustDLPCustomProfileSensitivityLevelsDataSourceModel struct {
	GroupID types.String `tfsdk:"group_id" json:"group_id,computed"`
	LevelID types.String `tfsdk:"level_id" json:"level_id,computed"`
}

type ZeroTrustDLPCustomProfileSharedEntriesDataSourceModel struct {
	ID            types.String                                                                              `tfsdk:"id" json:"id,computed"`
	CreatedAt     timetypes.RFC3339                                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Enabled       types.Bool                                                                                `tfsdk:"enabled" json:"enabled,computed"`
	Name          types.String                                                                              `tfsdk:"name" json:"name,computed"`
	Pattern       customfield.NestedObject[ZeroTrustDLPCustomProfileSharedEntriesPatternDataSourceModel]    `tfsdk:"pattern" json:"pattern,computed"`
	Type          types.String                                                                              `tfsdk:"type" json:"type,computed"`
	UpdatedAt     timetypes.RFC3339                                                                         `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Description   types.String                                                                              `tfsdk:"description" json:"description,computed"`
	ProfileID     types.String                                                                              `tfsdk:"profile_id" json:"profile_id,computed"`
	Confidence    customfield.NestedObject[ZeroTrustDLPCustomProfileSharedEntriesConfidenceDataSourceModel] `tfsdk:"confidence" json:"confidence,computed"`
	Variant       customfield.NestedObject[ZeroTrustDLPCustomProfileSharedEntriesVariantDataSourceModel]    `tfsdk:"variant" json:"variant,computed"`
	CaseSensitive types.Bool                                                                                `tfsdk:"case_sensitive" json:"case_sensitive,computed"`
	Secret        types.Bool                                                                                `tfsdk:"secret" json:"secret,computed"`
	WordList      jsontypes.Normalized                                                                      `tfsdk:"word_list" json:"word_list,computed"`
}

type ZeroTrustDLPCustomProfileSharedEntriesPatternDataSourceModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex,computed"`
	Validation types.String `tfsdk:"validation" json:"validation,computed"`
}

type ZeroTrustDLPCustomProfileSharedEntriesConfidenceDataSourceModel struct {
	AIContextAvailable types.Bool `tfsdk:"ai_context_available" json:"ai_context_available,computed"`
	Available          types.Bool `tfsdk:"available" json:"available,computed"`
}

type ZeroTrustDLPCustomProfileSharedEntriesVariantDataSourceModel struct {
	TopicType   types.String `tfsdk:"topic_type" json:"topic_type,computed"`
	Type        types.String `tfsdk:"type" json:"type,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
}
