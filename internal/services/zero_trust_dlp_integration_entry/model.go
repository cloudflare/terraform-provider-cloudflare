// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_integration_entry

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPIntegrationEntryResultEnvelope struct {
	Result ZeroTrustDLPIntegrationEntryModel `json:"result"`
}

type ZeroTrustDLPIntegrationEntryModel struct {
	ID            types.String                                                          `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                                                          `tfsdk:"account_id" path:"account_id,required"`
	EntryID       types.String                                                          `tfsdk:"entry_id" json:"entry_id,required,no_refresh"`
	ProfileID     types.String                                                          `tfsdk:"profile_id" json:"profile_id,optional"`
	Enabled       types.Bool                                                            `tfsdk:"enabled" json:"enabled,required"`
	CaseSensitive types.Bool                                                            `tfsdk:"case_sensitive" json:"case_sensitive,computed"`
	CreatedAt     timetypes.RFC3339                                                     `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name          types.String                                                          `tfsdk:"name" json:"name,computed"`
	Secret        types.Bool                                                            `tfsdk:"secret" json:"secret,computed"`
	Type          types.String                                                          `tfsdk:"type" json:"type,computed"`
	UpdatedAt     timetypes.RFC3339                                                     `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Confidence    customfield.NestedObject[ZeroTrustDLPIntegrationEntryConfidenceModel] `tfsdk:"confidence" json:"confidence,computed"`
	Pattern       customfield.NestedObject[ZeroTrustDLPIntegrationEntryPatternModel]    `tfsdk:"pattern" json:"pattern,computed"`
	Variant       customfield.NestedObject[ZeroTrustDLPIntegrationEntryVariantModel]    `tfsdk:"variant" json:"variant,computed"`
	WordList      jsontypes.Normalized                                                  `tfsdk:"word_list" json:"word_list,computed"`
}

func (m ZeroTrustDLPIntegrationEntryModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDLPIntegrationEntryModel) MarshalJSONForUpdate(state ZeroTrustDLPIntegrationEntryModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDLPIntegrationEntryConfidenceModel struct {
	AIContextAvailable types.Bool `tfsdk:"ai_context_available" json:"ai_context_available,computed"`
	Available          types.Bool `tfsdk:"available" json:"available,computed"`
}

type ZeroTrustDLPIntegrationEntryPatternModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex,computed"`
	Validation types.String `tfsdk:"validation" json:"validation,computed"`
}

type ZeroTrustDLPIntegrationEntryVariantModel struct {
	TopicType   types.String `tfsdk:"topic_type" json:"topic_type,computed"`
	Type        types.String `tfsdk:"type" json:"type,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
}
