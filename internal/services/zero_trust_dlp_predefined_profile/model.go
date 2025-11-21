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
	ID                  types.String                                                            `tfsdk:"id" json:"-,computed"`
	ProfileID           types.String                                                            `tfsdk:"profile_id" path:"profile_id,required"`
	AccountID           types.String                                                            `tfsdk:"account_id" path:"account_id,required"`
	EnabledEntries      *[]types.String                                                         `tfsdk:"enabled_entries" json:"enabled_entries,optional"`
	AIContextEnabled    types.Bool                                                              `tfsdk:"ai_context_enabled" json:"ai_context_enabled,computed_optional"`
	AllowedMatchCount   types.Int64                                                             `tfsdk:"allowed_match_count" json:"allowed_match_count,computed_optional"`
	ConfidenceThreshold types.String                                                            `tfsdk:"confidence_threshold" json:"confidence_threshold,computed_optional"`
	OCREnabled          types.Bool                                                              `tfsdk:"ocr_enabled" json:"ocr_enabled,computed_optional"`
	Entries             customfield.NestedObjectList[ZeroTrustDLPPredefinedProfileEntriesModel] `tfsdk:"entries" json:"entries,computed_optional"`
	Name                types.String                                                            `tfsdk:"name" json:"name,computed"`
	OpenAccess          types.Bool                                                              `tfsdk:"open_access" json:"open_access,computed"`
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
