// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_profile

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPPredefinedProfileResultEnvelope struct {
	Result ZeroTrustDLPPredefinedProfileModel `json:"result"`
}

type ZeroTrustDLPPredefinedProfileModel struct {
	ID                  types.String                                        `tfsdk:"id" json:"-,computed"`
	ProfileID           types.String                                        `tfsdk:"profile_id" path:"profile_id,required"`
	AccountID           types.String                                        `tfsdk:"account_id" path:"account_id,required"`
	Entries             *[]*ZeroTrustDLPPredefinedProfileEntriesModel       `tfsdk:"entries" json:"entries,required"`
	AIContextEnabled    types.Bool                                          `tfsdk:"ai_context_enabled" json:"ai_context_enabled,optional"`
	AllowedMatchCount   types.Int64                                         `tfsdk:"allowed_match_count" json:"allowed_match_count,optional"`
	ConfidenceThreshold types.String                                        `tfsdk:"confidence_threshold" json:"confidence_threshold,optional"`
	OCREnabled          types.Bool                                          `tfsdk:"ocr_enabled" json:"ocr_enabled,optional"`
	ContextAwareness    *ZeroTrustDLPPredefinedProfileContextAwarenessModel `tfsdk:"context_awareness" json:"context_awareness,optional"`
	CreatedAt           timetypes.RFC3339                                   `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description         types.String                                        `tfsdk:"description" json:"description,computed"`
	Name                types.String                                        `tfsdk:"name" json:"name,computed"`
	OpenAccess          types.Bool                                          `tfsdk:"open_access" json:"open_access,computed"`
	Type                types.String                                        `tfsdk:"type" json:"type,computed"`
	UpdatedAt           timetypes.RFC3339                                   `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
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
