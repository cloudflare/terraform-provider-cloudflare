// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_entry

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

type ZeroTrustDLPPredefinedEntriesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDLPPredefinedEntriesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDLPPredefinedEntriesDataSourceModel struct {
	AccountID types.String                                                                     `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                      `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustDLPPredefinedEntriesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDLPPredefinedEntriesDataSourceModel) toListParams(_ context.Context) (params zero_trust.DLPEntryPredefinedListParams, diags diag.Diagnostics) {
	params = zero_trust.DLPEntryPredefinedListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPPredefinedEntriesResultDataSourceModel struct {
	ID            types.String                                                                     `tfsdk:"id" json:"id,computed"`
	CreatedAt     timetypes.RFC3339                                                                `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Enabled       types.Bool                                                                       `tfsdk:"enabled" json:"enabled,computed"`
	Name          types.String                                                                     `tfsdk:"name" json:"name,computed"`
	Pattern       customfield.NestedObject[ZeroTrustDLPPredefinedEntriesPatternDataSourceModel]    `tfsdk:"pattern" json:"pattern,computed"`
	Type          types.String                                                                     `tfsdk:"type" json:"type,computed"`
	UpdatedAt     timetypes.RFC3339                                                                `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	ProfileID     types.String                                                                     `tfsdk:"profile_id" json:"profile_id,computed"`
	UploadStatus  types.String                                                                     `tfsdk:"upload_status" json:"upload_status,computed"`
	Confidence    customfield.NestedObject[ZeroTrustDLPPredefinedEntriesConfidenceDataSourceModel] `tfsdk:"confidence" json:"confidence,computed"`
	Variant       customfield.NestedObject[ZeroTrustDLPPredefinedEntriesVariantDataSourceModel]    `tfsdk:"variant" json:"variant,computed"`
	CaseSensitive types.Bool                                                                       `tfsdk:"case_sensitive" json:"case_sensitive,computed"`
	Secret        types.Bool                                                                       `tfsdk:"secret" json:"secret,computed"`
	WordList      jsontypes.Normalized                                                             `tfsdk:"word_list" json:"word_list,computed"`
}

type ZeroTrustDLPPredefinedEntriesPatternDataSourceModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex,computed"`
	Validation types.String `tfsdk:"validation" json:"validation,computed"`
}

type ZeroTrustDLPPredefinedEntriesConfidenceDataSourceModel struct {
	AIContextAvailable types.Bool `tfsdk:"ai_context_available" json:"ai_context_available,computed"`
	Available          types.Bool `tfsdk:"available" json:"available,computed"`
}

type ZeroTrustDLPPredefinedEntriesVariantDataSourceModel struct {
	TopicType   types.String `tfsdk:"topic_type" json:"topic_type,computed"`
	Type        types.String `tfsdk:"type" json:"type,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
}
