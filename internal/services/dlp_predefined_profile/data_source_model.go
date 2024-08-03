// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dlp_predefined_profile

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DLPPredefinedProfileResultDataSourceEnvelope struct {
	Result DLPPredefinedProfileDataSourceModel `json:"result,computed"`
}

type DLPPredefinedProfileDataSourceModel struct {
	AccountID         types.String                                         `tfsdk:"account_id" path:"account_id"`
	ProfileID         types.String                                         `tfsdk:"profile_id" path:"profile_id"`
	ID                types.String                                         `tfsdk:"id" json:"id"`
	Name              types.String                                         `tfsdk:"name" json:"name"`
	OCREnabled        types.Bool                                           `tfsdk:"ocr_enabled" json:"ocr_enabled"`
	Type              types.String                                         `tfsdk:"type" json:"type"`
	ContextAwareness  *DLPPredefinedProfileContextAwarenessDataSourceModel `tfsdk:"context_awareness" json:"context_awareness"`
	Entries           *[]*DLPPredefinedProfileEntriesDataSourceModel       `tfsdk:"entries" json:"entries"`
	AllowedMatchCount types.Float64                                        `tfsdk:"allowed_match_count" json:"allowed_match_count"`
}

type DLPPredefinedProfileContextAwarenessDataSourceModel struct {
	Enabled types.Bool                                                                        `tfsdk:"enabled" json:"enabled,computed"`
	Skip    customfield.NestedObject[DLPPredefinedProfileContextAwarenessSkipDataSourceModel] `tfsdk:"skip" json:"skip,computed"`
}

type DLPPredefinedProfileContextAwarenessSkipDataSourceModel struct {
	Files types.Bool `tfsdk:"files" json:"files,computed"`
}

type DLPPredefinedProfileEntriesDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled"`
	Name      types.String `tfsdk:"name" json:"name"`
	ProfileID types.String `tfsdk:"profile_id" json:"profile_id,computed"`
}
