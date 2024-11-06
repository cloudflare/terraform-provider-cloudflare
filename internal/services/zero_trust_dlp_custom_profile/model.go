// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_custom_profile

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPCustomProfileResultEnvelope struct {
	Result ZeroTrustDLPCustomProfileModel `json:"result"`
}

type ZeroTrustDLPCustomProfileModel struct {
	ID                  types.String                                                              `tfsdk:"id" json:"id,computed"`
	AccountID           types.String                                                              `tfsdk:"account_id" path:"account_id,required"`
	Profiles            customfield.NestedObjectList[ZeroTrustDLPCustomProfileProfilesModel]      `tfsdk:"profiles" json:"profiles,computed_optional"`
	ConfidenceThreshold types.String                                                              `tfsdk:"confidence_threshold" json:"confidence_threshold,optional"`
	Description         types.String                                                              `tfsdk:"description" json:"description,optional"`
	Name                types.String                                                              `tfsdk:"name" json:"name,optional"`
	OCREnabled          types.Bool                                                                `tfsdk:"ocr_enabled" json:"ocr_enabled,optional"`
	AllowedMatchCount   types.Int64                                                               `tfsdk:"allowed_match_count" json:"allowed_match_count,computed_optional"`
	ContextAwareness    customfield.NestedObject[ZeroTrustDLPCustomProfileContextAwarenessModel]  `tfsdk:"context_awareness" json:"context_awareness,computed_optional"`
	Entries             customfield.NestedObjectList[ZeroTrustDLPCustomProfileEntriesModel]       `tfsdk:"entries" json:"entries,computed_optional"`
	SharedEntries       customfield.NestedObjectList[ZeroTrustDLPCustomProfileSharedEntriesModel] `tfsdk:"shared_entries" json:"shared_entries,computed_optional"`
}

func (m ZeroTrustDLPCustomProfileModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDLPCustomProfileModel) MarshalJSONForUpdate(state ZeroTrustDLPCustomProfileModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDLPCustomProfileProfilesModel struct {
	Entries             *[]*ZeroTrustDLPCustomProfileProfilesEntriesModel                                 `tfsdk:"entries" json:"entries,required"`
	Name                types.String                                                                      `tfsdk:"name" json:"name,required"`
	AllowedMatchCount   types.Int64                                                                       `tfsdk:"allowed_match_count" json:"allowed_match_count,computed_optional"`
	ConfidenceThreshold types.String                                                                      `tfsdk:"confidence_threshold" json:"confidence_threshold,optional"`
	ContextAwareness    customfield.NestedObject[ZeroTrustDLPCustomProfileProfilesContextAwarenessModel]  `tfsdk:"context_awareness" json:"context_awareness,computed_optional"`
	Description         types.String                                                                      `tfsdk:"description" json:"description,optional"`
	OCREnabled          types.Bool                                                                        `tfsdk:"ocr_enabled" json:"ocr_enabled,optional"`
	SharedEntries       customfield.NestedObjectList[ZeroTrustDLPCustomProfileProfilesSharedEntriesModel] `tfsdk:"shared_entries" json:"shared_entries,computed_optional"`
}

type ZeroTrustDLPCustomProfileProfilesEntriesModel struct {
	Enabled types.Bool                                            `tfsdk:"enabled" json:"enabled,required"`
	Name    types.String                                          `tfsdk:"name" json:"name,required"`
	Pattern *ZeroTrustDLPCustomProfileProfilesEntriesPatternModel `tfsdk:"pattern" json:"pattern,optional"`
	Words   *[]types.String                                       `tfsdk:"words" json:"words,optional"`
}

type ZeroTrustDLPCustomProfileProfilesEntriesPatternModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex,required"`
	Validation types.String `tfsdk:"validation" json:"validation,optional"`
}

type ZeroTrustDLPCustomProfileProfilesContextAwarenessModel struct {
	Enabled types.Bool                                                  `tfsdk:"enabled" json:"enabled,required"`
	Skip    *ZeroTrustDLPCustomProfileProfilesContextAwarenessSkipModel `tfsdk:"skip" json:"skip,required"`
}

type ZeroTrustDLPCustomProfileProfilesContextAwarenessSkipModel struct {
	Files types.Bool `tfsdk:"files" json:"files,required"`
}

type ZeroTrustDLPCustomProfileProfilesSharedEntriesModel struct {
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled,required"`
	EntryID   types.String `tfsdk:"entry_id" json:"entry_id,required"`
	EntryType types.String `tfsdk:"entry_type" json:"entry_type,required"`
}

type ZeroTrustDLPCustomProfileContextAwarenessModel struct {
	Enabled types.Bool                                          `tfsdk:"enabled" json:"enabled,required"`
	Skip    *ZeroTrustDLPCustomProfileContextAwarenessSkipModel `tfsdk:"skip" json:"skip,required"`
}

type ZeroTrustDLPCustomProfileContextAwarenessSkipModel struct {
	Files types.Bool `tfsdk:"files" json:"files,required"`
}

type ZeroTrustDLPCustomProfileEntriesModel struct {
	Enabled types.Bool                                                             `tfsdk:"enabled" json:"enabled,required"`
	Name    types.String                                                           `tfsdk:"name" json:"name,required"`
	Pattern customfield.NestedObject[ZeroTrustDLPCustomProfileEntriesPatternModel] `tfsdk:"pattern" json:"pattern,computed_optional"`
	Words   *[]types.String                                                        `tfsdk:"words" json:"words,optional"`
}

type ZeroTrustDLPCustomProfileEntriesPatternModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex,required"`
	Validation types.String `tfsdk:"validation" json:"validation,optional"`
}

type ZeroTrustDLPCustomProfileSharedEntriesModel struct {
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled,required"`
	EntryID   types.String `tfsdk:"entry_id" json:"entry_id,required"`
	EntryType types.String `tfsdk:"entry_type" json:"entry_type,required"`
}
