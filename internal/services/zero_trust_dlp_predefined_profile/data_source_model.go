// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_profile

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

type ZeroTrustDLPPredefinedProfileResultDataSourceEnvelope struct {
Result ZeroTrustDLPPredefinedProfileDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPPredefinedProfileDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
ProfileID types.String `tfsdk:"profile_id" path:"profile_id,required"`
AIContextEnabled types.Bool `tfsdk:"ai_context_enabled" json:"ai_context_enabled,computed"`
AllowedMatchCount types.Int64 `tfsdk:"allowed_match_count" json:"allowed_match_count,computed"`
ConfidenceThreshold types.String `tfsdk:"confidence_threshold" json:"confidence_threshold,computed"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
Description types.String `tfsdk:"description" json:"description,computed"`
ID types.String `tfsdk:"id" json:"id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
OCREnabled types.Bool `tfsdk:"ocr_enabled" json:"ocr_enabled,computed"`
OpenAccess types.Bool `tfsdk:"open_access" json:"open_access,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
ContextAwareness customfield.NestedObject[ZeroTrustDLPPredefinedProfileContextAwarenessDataSourceModel] `tfsdk:"context_awareness" json:"context_awareness,computed"`
Entries customfield.NestedObjectList[ZeroTrustDLPPredefinedProfileEntriesDataSourceModel] `tfsdk:"entries" json:"entries,computed"`
}

func (m *ZeroTrustDLPPredefinedProfileDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DLPProfilePredefinedGetParams, diags diag.Diagnostics) {
  params = zero_trust.DLPProfilePredefinedGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type ZeroTrustDLPPredefinedProfileContextAwarenessDataSourceModel struct {
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
Skip customfield.NestedObject[ZeroTrustDLPPredefinedProfileContextAwarenessSkipDataSourceModel] `tfsdk:"skip" json:"skip,computed"`
}

type ZeroTrustDLPPredefinedProfileContextAwarenessSkipDataSourceModel struct {
Files types.Bool `tfsdk:"files" json:"files,computed"`
}

type ZeroTrustDLPPredefinedProfileEntriesDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Pattern customfield.NestedObject[ZeroTrustDLPPredefinedProfileEntriesPatternDataSourceModel] `tfsdk:"pattern" json:"pattern,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
ProfileID types.String `tfsdk:"profile_id" json:"profile_id,computed"`
Confidence customfield.NestedObject[ZeroTrustDLPPredefinedProfileEntriesConfidenceDataSourceModel] `tfsdk:"confidence" json:"confidence,computed"`
Secret types.Bool `tfsdk:"secret" json:"secret,computed"`
WordList jsontypes.Normalized `tfsdk:"word_list" json:"word_list,computed"`
}

type ZeroTrustDLPPredefinedProfileEntriesPatternDataSourceModel struct {
Regex types.String `tfsdk:"regex" json:"regex,computed"`
Validation types.String `tfsdk:"validation" json:"validation,computed"`
}

type ZeroTrustDLPPredefinedProfileEntriesConfidenceDataSourceModel struct {
AIContextAvailable types.Bool `tfsdk:"ai_context_available" json:"ai_context_available,computed"`
Available types.Bool `tfsdk:"available" json:"available,computed"`
}
