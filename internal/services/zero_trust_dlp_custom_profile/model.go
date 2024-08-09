// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_custom_profile

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPCustomProfileResultEnvelope struct {
	Result ZeroTrustDLPCustomProfileModel `json:"result,computed"`
}

type ZeroTrustDLPCustomProfileModel struct {
	AccountID         types.String                                    `tfsdk:"account_id" path:"account_id"`
	ProfileID         types.String                                    `tfsdk:"profile_id" path:"profile_id"`
	Profiles          *[]*ZeroTrustDLPCustomProfileProfilesModel      `tfsdk:"profiles" json:"profiles"`
	Description       types.String                                    `tfsdk:"description" json:"description"`
	Name              types.String                                    `tfsdk:"name" json:"name"`
	OCREnabled        types.Bool                                      `tfsdk:"ocr_enabled" json:"ocr_enabled"`
	ContextAwareness  *ZeroTrustDLPCustomProfileContextAwarenessModel `tfsdk:"context_awareness" json:"context_awareness"`
	Entries           *[]*ZeroTrustDLPCustomProfileEntriesModel       `tfsdk:"entries" json:"entries"`
	SharedEntries     *[]*ZeroTrustDLPCustomProfileSharedEntriesModel `tfsdk:"shared_entries" json:"shared_entries"`
	AllowedMatchCount types.Float64                                   `tfsdk:"allowed_match_count" json:"allowed_match_count"`
	CreatedAt         timetypes.RFC3339                               `tfsdk:"created_at" json:"created_at,computed"`
	ID                types.String                                    `tfsdk:"id" json:"id,computed"`
	Type              types.String                                    `tfsdk:"type" json:"type,computed"`
	UpdatedAt         timetypes.RFC3339                               `tfsdk:"updated_at" json:"updated_at,computed"`
}

type ZeroTrustDLPCustomProfileProfilesModel struct {
	AllowedMatchCount types.Float64                                           `tfsdk:"allowed_match_count" json:"allowed_match_count"`
	ContextAwareness  *ZeroTrustDLPCustomProfileProfilesContextAwarenessModel `tfsdk:"context_awareness" json:"context_awareness"`
	Description       types.String                                            `tfsdk:"description" json:"description"`
	Entries           *[]*ZeroTrustDLPCustomProfileProfilesEntriesModel       `tfsdk:"entries" json:"entries"`
	Name              types.String                                            `tfsdk:"name" json:"name"`
	OCREnabled        types.Bool                                              `tfsdk:"ocr_enabled" json:"ocr_enabled"`
}

type ZeroTrustDLPCustomProfileProfilesContextAwarenessModel struct {
	Enabled types.Bool                                                  `tfsdk:"enabled" json:"enabled"`
	Skip    *ZeroTrustDLPCustomProfileProfilesContextAwarenessSkipModel `tfsdk:"skip" json:"skip"`
}

type ZeroTrustDLPCustomProfileProfilesContextAwarenessSkipModel struct {
	Files types.Bool `tfsdk:"files" json:"files"`
}

type ZeroTrustDLPCustomProfileProfilesEntriesModel struct {
	Enabled types.Bool                                            `tfsdk:"enabled" json:"enabled"`
	Name    types.String                                          `tfsdk:"name" json:"name"`
	Pattern *ZeroTrustDLPCustomProfileProfilesEntriesPatternModel `tfsdk:"pattern" json:"pattern"`
}

type ZeroTrustDLPCustomProfileProfilesEntriesPatternModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex"`
	Validation types.String `tfsdk:"validation" json:"validation"`
}

type ZeroTrustDLPCustomProfileContextAwarenessModel struct {
	Enabled types.Bool                                          `tfsdk:"enabled" json:"enabled"`
	Skip    *ZeroTrustDLPCustomProfileContextAwarenessSkipModel `tfsdk:"skip" json:"skip"`
}

type ZeroTrustDLPCustomProfileContextAwarenessSkipModel struct {
	Files types.Bool `tfsdk:"files" json:"files"`
}

type ZeroTrustDLPCustomProfileEntriesModel struct {
	ID        types.String                                  `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339                             `tfsdk:"created_at" json:"created_at,computed"`
	Enabled   types.Bool                                    `tfsdk:"enabled" json:"enabled"`
	Name      types.String                                  `tfsdk:"name" json:"name"`
	Pattern   *ZeroTrustDLPCustomProfileEntriesPatternModel `tfsdk:"pattern" json:"pattern"`
	ProfileID types.String                                  `tfsdk:"profile_id" json:"profile_id,computed"`
	UpdatedAt timetypes.RFC3339                             `tfsdk:"updated_at" json:"updated_at,computed"`
}

type ZeroTrustDLPCustomProfileEntriesPatternModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex"`
	Validation types.String `tfsdk:"validation" json:"validation"`
}

type ZeroTrustDLPCustomProfileSharedEntriesModel struct {
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
	EntryID types.String `tfsdk:"entry_id" json:"entry_id,computed"`
}
