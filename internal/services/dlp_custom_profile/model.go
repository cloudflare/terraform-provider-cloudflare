// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dlp_custom_profile

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DLPCustomProfileResultEnvelope struct {
	Result DLPCustomProfileModel `json:"result,computed"`
}

type DLPCustomProfileModel struct {
	AccountID         types.String                           `tfsdk:"account_id" path:"account_id"`
	ProfileID         types.String                           `tfsdk:"profile_id" path:"profile_id"`
	Profiles          *[]*DLPCustomProfileProfilesModel      `tfsdk:"profiles" json:"profiles"`
	AllowedMatchCount types.Float64                          `tfsdk:"allowed_match_count" json:"allowed_match_count"`
	ContextAwareness  *DLPCustomProfileContextAwarenessModel `tfsdk:"context_awareness" json:"context_awareness"`
	Description       types.String                           `tfsdk:"description" json:"description"`
	Entries           *[]*DLPCustomProfileEntriesModel       `tfsdk:"entries" json:"entries"`
	Name              types.String                           `tfsdk:"name" json:"name"`
	OCREnabled        types.Bool                             `tfsdk:"ocr_enabled" json:"ocr_enabled"`
	SharedEntries     *[]*DLPCustomProfileSharedEntriesModel `tfsdk:"shared_entries" json:"shared_entries"`
	ID                types.String                           `tfsdk:"id" json:"id,computed"`
	CreatedAt         types.String                           `tfsdk:"created_at" json:"created_at,computed"`
	Type              types.String                           `tfsdk:"type" json:"type,computed"`
	UpdatedAt         types.String                           `tfsdk:"updated_at" json:"updated_at,computed"`
}

type DLPCustomProfileProfilesModel struct {
	AllowedMatchCount types.Float64                                  `tfsdk:"allowed_match_count" json:"allowed_match_count"`
	ContextAwareness  *DLPCustomProfileProfilesContextAwarenessModel `tfsdk:"context_awareness" json:"context_awareness"`
	Description       types.String                                   `tfsdk:"description" json:"description"`
	Entries           *[]*DLPCustomProfileProfilesEntriesModel       `tfsdk:"entries" json:"entries"`
	Name              types.String                                   `tfsdk:"name" json:"name"`
	OCREnabled        types.Bool                                     `tfsdk:"ocr_enabled" json:"ocr_enabled"`
}

type DLPCustomProfileProfilesContextAwarenessModel struct {
	Enabled types.Bool                                         `tfsdk:"enabled" json:"enabled"`
	Skip    *DLPCustomProfileProfilesContextAwarenessSkipModel `tfsdk:"skip" json:"skip"`
}

type DLPCustomProfileProfilesContextAwarenessSkipModel struct {
	Files types.Bool `tfsdk:"files" json:"files"`
}

type DLPCustomProfileProfilesEntriesModel struct {
	Enabled types.Bool                                   `tfsdk:"enabled" json:"enabled"`
	Name    types.String                                 `tfsdk:"name" json:"name"`
	Pattern *DLPCustomProfileProfilesEntriesPatternModel `tfsdk:"pattern" json:"pattern"`
}

type DLPCustomProfileProfilesEntriesPatternModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex"`
	Validation types.String `tfsdk:"validation" json:"validation"`
}

type DLPCustomProfileContextAwarenessModel struct {
	Enabled types.Bool                                 `tfsdk:"enabled" json:"enabled"`
	Skip    *DLPCustomProfileContextAwarenessSkipModel `tfsdk:"skip" json:"skip"`
}

type DLPCustomProfileContextAwarenessSkipModel struct {
	Files types.Bool `tfsdk:"files" json:"files"`
}

type DLPCustomProfileEntriesModel struct {
	ID        types.String                         `tfsdk:"id" json:"id,computed"`
	CreatedAt types.String                         `tfsdk:"created_at" json:"created_at,computed"`
	Enabled   types.Bool                           `tfsdk:"enabled" json:"enabled"`
	Name      types.String                         `tfsdk:"name" json:"name"`
	Pattern   *DLPCustomProfileEntriesPatternModel `tfsdk:"pattern" json:"pattern"`
	ProfileID types.String                         `tfsdk:"profile_id" json:"profile_id"`
	UpdatedAt types.String                         `tfsdk:"updated_at" json:"updated_at,computed"`
}

type DLPCustomProfileEntriesPatternModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex"`
	Validation types.String `tfsdk:"validation" json:"validation"`
}

type DLPCustomProfileSharedEntriesModel struct {
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
	EntryID types.String `tfsdk:"entry_id" json:"entry_id,computed"`
}
