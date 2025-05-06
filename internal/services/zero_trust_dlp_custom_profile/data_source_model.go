// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_custom_profile

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
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
	AccountID           types.String                                                                       `tfsdk:"account_id" path:"account_id,required"`
	ProfileID           types.String                                                                       `tfsdk:"profile_id" path:"profile_id,required"`
	AIContextEnabled    types.Bool                                                                         `tfsdk:"ai_context_enabled" json:"ai_context_enabled,computed"`
	AllowedMatchCount   types.Int64                                                                        `tfsdk:"allowed_match_count" json:"allowed_match_count,computed"`
	ConfidenceThreshold types.String                                                                       `tfsdk:"confidence_threshold" json:"confidence_threshold,computed"`
	CreatedAt           timetypes.RFC3339                                                                  `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description         types.String                                                                       `tfsdk:"description" json:"description,computed"`
	ID                  types.String                                                                       `tfsdk:"id" json:"id,computed"`
	Name                types.String                                                                       `tfsdk:"name" json:"name,computed"`
	OCREnabled          types.Bool                                                                         `tfsdk:"ocr_enabled" json:"ocr_enabled,computed"`
	OpenAccess          types.Bool                                                                         `tfsdk:"open_access" json:"open_access,computed"`
	Type                types.String                                                                       `tfsdk:"type" json:"type,computed"`
	UpdatedAt           timetypes.RFC3339                                                                  `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	ContextAwareness    customfield.NestedObject[ZeroTrustDLPCustomProfileContextAwarenessDataSourceModel] `tfsdk:"context_awareness" json:"context_awareness,computed"`
	Entries             customfield.NestedObjectList[ZeroTrustDLPCustomProfileEntriesDataSourceModel]      `tfsdk:"entries" json:"entries,computed"`
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
	ProfileID     types.String                                                                        `tfsdk:"profile_id" json:"profile_id,computed"`
	Confidence    customfield.NestedObject[ZeroTrustDLPCustomProfileEntriesConfidenceDataSourceModel] `tfsdk:"confidence" json:"confidence,computed"`
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
