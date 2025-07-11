// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_entry

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPPredefinedEntryResultEnvelope struct {
	Result ZeroTrustDLPPredefinedEntryModel `json:"result"`
}

type ZeroTrustDLPPredefinedEntryModel struct {
	ID         types.String                                                         `tfsdk:"id" json:"id,computed"`
	AccountID  types.String                                                         `tfsdk:"account_id" path:"account_id,required"`
	EntryID    types.String                                                         `tfsdk:"entry_id" json:"entry_id,required"`
	ProfileID  types.String                                                         `tfsdk:"profile_id" json:"profile_id,optional"`
	Enabled    types.Bool                                                           `tfsdk:"enabled" json:"enabled,required"`
	Name       types.String                                                         `tfsdk:"name" json:"name,computed"`
	Confidence customfield.NestedObject[ZeroTrustDLPPredefinedEntryConfidenceModel] `tfsdk:"confidence" json:"confidence,computed"`
}

func (m ZeroTrustDLPPredefinedEntryModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDLPPredefinedEntryModel) MarshalJSONForUpdate(state ZeroTrustDLPPredefinedEntryModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDLPPredefinedEntryConfidenceModel struct {
	AIContextAvailable types.Bool `tfsdk:"ai_context_available" json:"ai_context_available,computed"`
	Available          types.Bool `tfsdk:"available" json:"available,computed"`
}
