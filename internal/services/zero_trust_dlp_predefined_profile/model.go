// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_profile

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPPredefinedProfileResultEnvelope struct {
	Result ZeroTrustDLPPredefinedProfileModel `json:"result"`
}

type ZeroTrustDLPPredefinedProfileModel struct {
	ID                types.String                                                                 `tfsdk:"id" json:"-,computed"`
	ProfileID         types.String                                                                 `tfsdk:"profile_id" path:"profile_id,required"`
	AccountID         types.String                                                                 `tfsdk:"account_id" path:"account_id,required"`
	Entries           *[]*ZeroTrustDLPPredefinedProfileEntriesModel                                `tfsdk:"entries" json:"entries,required"`
	AllowedMatchCount types.Int64                                                                  `tfsdk:"allowed_match_count" json:"allowed_match_count,optional"`
	OCREnabled        types.Bool                                                                   `tfsdk:"ocr_enabled" json:"ocr_enabled,optional"`
	ContextAwareness  customfield.NestedObject[ZeroTrustDLPPredefinedProfileContextAwarenessModel] `tfsdk:"context_awareness" json:"context_awareness,computed_optional"`
}

func (m ZeroTrustDLPPredefinedProfileModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDLPPredefinedProfileModel) MarshalJSONForUpdate(state ZeroTrustDLPPredefinedProfileModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDLPPredefinedProfileEntriesModel struct {
	ID      types.String `tfsdk:"id" json:"id,required"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,required"`
}

type ZeroTrustDLPPredefinedProfileContextAwarenessModel struct {
	Enabled types.Bool                                              `tfsdk:"enabled" json:"enabled,required"`
	Skip    *ZeroTrustDLPPredefinedProfileContextAwarenessSkipModel `tfsdk:"skip" json:"skip,required"`
}

type ZeroTrustDLPPredefinedProfileContextAwarenessSkipModel struct {
	Files types.Bool `tfsdk:"files" json:"files,required"`
}
