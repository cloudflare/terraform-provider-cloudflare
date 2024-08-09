// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_profile

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPPredefinedProfileResultDataSourceEnvelope struct {
	Result ZeroTrustDLPPredefinedProfileDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPPredefinedProfileDataSourceModel struct {
	AccountID         types.String                                                  `tfsdk:"account_id" path:"account_id"`
	ProfileID         types.String                                                  `tfsdk:"profile_id" path:"profile_id"`
	ID                types.String                                                  `tfsdk:"id" json:"id"`
	Name              types.String                                                  `tfsdk:"name" json:"name"`
	OCREnabled        types.Bool                                                    `tfsdk:"ocr_enabled" json:"ocr_enabled"`
	Type              types.String                                                  `tfsdk:"type" json:"type"`
	ContextAwareness  *ZeroTrustDLPPredefinedProfileContextAwarenessDataSourceModel `tfsdk:"context_awareness" json:"context_awareness"`
	Entries           *[]*ZeroTrustDLPPredefinedProfileEntriesDataSourceModel       `tfsdk:"entries" json:"entries"`
	AllowedMatchCount types.Float64                                                 `tfsdk:"allowed_match_count" json:"allowed_match_count"`
}

type ZeroTrustDLPPredefinedProfileContextAwarenessDataSourceModel struct {
	Enabled types.Bool                                                                                 `tfsdk:"enabled" json:"enabled,computed"`
	Skip    customfield.NestedObject[ZeroTrustDLPPredefinedProfileContextAwarenessSkipDataSourceModel] `tfsdk:"skip" json:"skip,computed"`
}

type ZeroTrustDLPPredefinedProfileContextAwarenessSkipDataSourceModel struct {
	Files types.Bool `tfsdk:"files" json:"files,computed"`
}

type ZeroTrustDLPPredefinedProfileEntriesDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled"`
	Name      types.String `tfsdk:"name" json:"name"`
	ProfileID types.String `tfsdk:"profile_id" json:"profile_id,computed"`
}
