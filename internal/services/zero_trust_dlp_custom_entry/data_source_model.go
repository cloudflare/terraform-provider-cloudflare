// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_custom_entry

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

type ZeroTrustDLPCustomEntryResultDataSourceEnvelope struct {
	Result ZeroTrustDLPCustomEntryDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPCustomEntryDataSourceModel struct {
	ID            types.String                                                               `tfsdk:"id" path:"entry_id,computed"`
	EntryID       types.String                                                               `tfsdk:"entry_id" path:"entry_id,required"`
	AccountID     types.String                                                               `tfsdk:"account_id" path:"account_id,required"`
	CaseSensitive types.Bool                                                                 `tfsdk:"case_sensitive" json:"case_sensitive,computed"`
	CreatedAt     timetypes.RFC3339                                                          `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Enabled       types.Bool                                                                 `tfsdk:"enabled" json:"enabled,computed"`
	Name          types.String                                                               `tfsdk:"name" json:"name,computed"`
	ProfileID     types.String                                                               `tfsdk:"profile_id" json:"profile_id,computed"`
	Secret        types.Bool                                                                 `tfsdk:"secret" json:"secret,computed"`
	Type          types.String                                                               `tfsdk:"type" json:"type,computed"`
	UpdatedAt     timetypes.RFC3339                                                          `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Confidence    customfield.NestedObject[ZeroTrustDLPCustomEntryConfidenceDataSourceModel] `tfsdk:"confidence" json:"confidence,computed"`
	Pattern       customfield.NestedObject[ZeroTrustDLPCustomEntryPatternDataSourceModel]    `tfsdk:"pattern" json:"pattern,computed"`
	Variant       customfield.NestedObject[ZeroTrustDLPCustomEntryVariantDataSourceModel]    `tfsdk:"variant" json:"variant,computed"`
	WordList      jsontypes.Normalized                                                       `tfsdk:"word_list" json:"word_list,computed"`
}

func (m *ZeroTrustDLPCustomEntryDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DLPEntryCustomGetParams, diags diag.Diagnostics) {
	params = zero_trust.DLPEntryCustomGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPCustomEntryConfidenceDataSourceModel struct {
	AIContextAvailable types.Bool `tfsdk:"ai_context_available" json:"ai_context_available,computed"`
	Available          types.Bool `tfsdk:"available" json:"available,computed"`
}

type ZeroTrustDLPCustomEntryPatternDataSourceModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex,computed"`
	Validation types.String `tfsdk:"validation" json:"validation,computed"`
}

type ZeroTrustDLPCustomEntryVariantDataSourceModel struct {
	TopicType   types.String `tfsdk:"topic_type" json:"topic_type,computed"`
	Type        types.String `tfsdk:"type" json:"type,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
}
