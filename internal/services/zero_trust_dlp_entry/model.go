// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_entry

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPEntryResultEnvelope struct {
	Result ZeroTrustDLPEntryModel `json:"result"`
}

type ZeroTrustDLPEntryModel struct {
	ID         types.String                                               `tfsdk:"id" json:"id,computed"`
	AccountID  types.String                                               `tfsdk:"account_id" path:"account_id,required"`
	ProfileID  types.String                                               `tfsdk:"profile_id" json:"profile_id,required"`
	Enabled    types.Bool                                                 `tfsdk:"enabled" json:"enabled,required"`
	Name       types.String                                               `tfsdk:"name" json:"name,required"`
	Pattern    *ZeroTrustDLPEntryPatternModel                             `tfsdk:"pattern" json:"pattern,required"`
	Type       types.String                                               `tfsdk:"type" json:"type,optional"`
	CreatedAt  timetypes.RFC3339                                          `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Secret     types.Bool                                                 `tfsdk:"secret" json:"secret,computed"`
	UpdatedAt  timetypes.RFC3339                                          `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Confidence customfield.NestedObject[ZeroTrustDLPEntryConfidenceModel] `tfsdk:"confidence" json:"confidence,computed"`
	WordList   jsontypes.Normalized                                       `tfsdk:"word_list" json:"word_list,computed"`
}

func (m ZeroTrustDLPEntryModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDLPEntryModel) MarshalJSONForUpdate(state ZeroTrustDLPEntryModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDLPEntryPatternModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex,required"`
	Validation types.String `tfsdk:"validation" json:"validation,optional"`
}

type ZeroTrustDLPEntryConfidenceModel struct {
	Available types.Bool `tfsdk:"available" json:"available,computed"`
}
