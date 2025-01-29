// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_entry

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPEntriesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDLPEntriesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDLPEntriesDataSourceModel struct {
	AccountID types.String                                                           `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                            `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustDLPEntriesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDLPEntriesDataSourceModel) toListParams(_ context.Context) (params zero_trust.DLPEntryListParams, diags diag.Diagnostics) {
	params = zero_trust.DLPEntryListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPEntriesResultDataSourceModel struct {
	ID         types.String                                                           `tfsdk:"id" json:"id,computed"`
	CreatedAt  timetypes.RFC3339                                                      `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Enabled    types.Bool                                                             `tfsdk:"enabled" json:"enabled,computed"`
	Name       types.String                                                           `tfsdk:"name" json:"name,computed"`
	Pattern    customfield.NestedObject[ZeroTrustDLPEntriesPatternDataSourceModel]    `tfsdk:"pattern" json:"pattern,computed"`
	Type       types.String                                                           `tfsdk:"type" json:"type,computed"`
	UpdatedAt  timetypes.RFC3339                                                      `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	ProfileID  types.String                                                           `tfsdk:"profile_id" json:"profile_id,computed"`
	Confidence customfield.NestedObject[ZeroTrustDLPEntriesConfidenceDataSourceModel] `tfsdk:"confidence" json:"confidence,computed"`
	Secret     types.Bool                                                             `tfsdk:"secret" json:"secret,computed"`
	WordList   jsontypes.Normalized                                                   `tfsdk:"word_list" json:"word_list,computed"`
}

type ZeroTrustDLPEntriesPatternDataSourceModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex,computed"`
	Validation types.String `tfsdk:"validation" json:"validation,computed"`
}

type ZeroTrustDLPEntriesConfidenceDataSourceModel struct {
	Available types.Bool `tfsdk:"available" json:"available,computed"`
}
