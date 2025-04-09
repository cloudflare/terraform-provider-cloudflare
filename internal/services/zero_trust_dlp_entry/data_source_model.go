// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_entry

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

type ZeroTrustDLPEntryResultDataSourceEnvelope struct {
Result ZeroTrustDLPEntryDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPEntryDataSourceModel struct {
ID types.String `tfsdk:"id" path:"entry_id,computed"`
EntryID types.String `tfsdk:"entry_id" path:"entry_id,optional"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
ProfileID types.String `tfsdk:"profile_id" json:"profile_id,computed"`
Secret types.Bool `tfsdk:"secret" json:"secret,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
Confidence customfield.NestedObject[ZeroTrustDLPEntryConfidenceDataSourceModel] `tfsdk:"confidence" json:"confidence,computed"`
Pattern customfield.NestedObject[ZeroTrustDLPEntryPatternDataSourceModel] `tfsdk:"pattern" json:"pattern,computed"`
WordList jsontypes.Normalized `tfsdk:"word_list" json:"word_list,computed"`
}

func (m *ZeroTrustDLPEntryDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DLPEntryGetParams, diags diag.Diagnostics) {
  params = zero_trust.DLPEntryGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type ZeroTrustDLPEntryConfidenceDataSourceModel struct {
AIContextAvailable types.Bool `tfsdk:"ai_context_available" json:"ai_context_available,computed"`
Available types.Bool `tfsdk:"available" json:"available,computed"`
}

type ZeroTrustDLPEntryPatternDataSourceModel struct {
Regex types.String `tfsdk:"regex" json:"regex,computed"`
Validation types.String `tfsdk:"validation" json:"validation,computed"`
}
