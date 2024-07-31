// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dlp_predefined_profile

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DLPPredefinedProfileResultEnvelope struct {
	Result DLPPredefinedProfileModel `json:"result,computed"`
}

type DLPPredefinedProfileModel struct {
	ID                types.String                               `tfsdk:"id" json:"-,computed"`
	AccountID         types.String                               `tfsdk:"account_id" path:"account_id"`
	ProfileID         types.String                               `tfsdk:"profile_id" path:"profile_id"`
	AllowedMatchCount types.Float64                              `tfsdk:"allowed_match_count" json:"allowed_match_count"`
	ContextAwareness  *DLPPredefinedProfileContextAwarenessModel `tfsdk:"context_awareness" json:"context_awareness"`
	Entries           *[]*DLPPredefinedProfileEntriesModel       `tfsdk:"entries" json:"entries"`
	OCREnabled        types.Bool                                 `tfsdk:"ocr_enabled" json:"ocr_enabled"`
	Name              types.String                               `tfsdk:"name" json:"name,computed"`
	Type              types.String                               `tfsdk:"type" json:"type,computed"`
}

type DLPPredefinedProfileContextAwarenessModel struct {
	Enabled types.Bool                                     `tfsdk:"enabled" json:"enabled"`
	Skip    *DLPPredefinedProfileContextAwarenessSkipModel `tfsdk:"skip" json:"skip"`
}

type DLPPredefinedProfileContextAwarenessSkipModel struct {
	Files types.Bool `tfsdk:"files" json:"files"`
}

type DLPPredefinedProfileEntriesModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
}
