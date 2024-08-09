// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_custom_profile

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPCustomProfileResultDataSourceEnvelope struct {
	Result ZeroTrustDLPCustomProfileDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPCustomProfileDataSourceModel struct {
	AccountID         types.String                                              `tfsdk:"account_id" path:"account_id"`
	ProfileID         types.String                                              `tfsdk:"profile_id" path:"profile_id"`
	CreatedAt         timetypes.RFC3339                                         `tfsdk:"created_at" json:"created_at"`
	Description       types.String                                              `tfsdk:"description" json:"description"`
	ID                types.String                                              `tfsdk:"id" json:"id"`
	Name              types.String                                              `tfsdk:"name" json:"name"`
	OCREnabled        types.Bool                                                `tfsdk:"ocr_enabled" json:"ocr_enabled"`
	Type              types.String                                              `tfsdk:"type" json:"type"`
	UpdatedAt         timetypes.RFC3339                                         `tfsdk:"updated_at" json:"updated_at"`
	ContextAwareness  *ZeroTrustDLPCustomProfileContextAwarenessDataSourceModel `tfsdk:"context_awareness" json:"context_awareness"`
	Entries           *[]*ZeroTrustDLPCustomProfileEntriesDataSourceModel       `tfsdk:"entries" json:"entries"`
	AllowedMatchCount types.Float64                                             `tfsdk:"allowed_match_count" json:"allowed_match_count"`
}

type ZeroTrustDLPCustomProfileContextAwarenessDataSourceModel struct {
	Enabled types.Bool                                                                             `tfsdk:"enabled" json:"enabled,computed"`
	Skip    customfield.NestedObject[ZeroTrustDLPCustomProfileContextAwarenessSkipDataSourceModel] `tfsdk:"skip" json:"skip,computed"`
}

type ZeroTrustDLPCustomProfileContextAwarenessSkipDataSourceModel struct {
	Files types.Bool `tfsdk:"files" json:"files,computed"`
}

type ZeroTrustDLPCustomProfileEntriesDataSourceModel struct {
	ID        types.String                                            `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339                                       `tfsdk:"created_at" json:"created_at,computed"`
	Enabled   types.Bool                                              `tfsdk:"enabled" json:"enabled"`
	Name      types.String                                            `tfsdk:"name" json:"name"`
	Pattern   *ZeroTrustDLPCustomProfileEntriesPatternDataSourceModel `tfsdk:"pattern" json:"pattern"`
	ProfileID types.String                                            `tfsdk:"profile_id" json:"profile_id,computed"`
	UpdatedAt timetypes.RFC3339                                       `tfsdk:"updated_at" json:"updated_at,computed"`
}

type ZeroTrustDLPCustomProfileEntriesPatternDataSourceModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex,computed"`
	Validation types.String `tfsdk:"validation" json:"validation"`
}
