// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_entry

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPEntryResultDataSourceEnvelope struct {
	Result ZeroTrustDLPEntryDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPEntryResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDLPEntryDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDLPEntryDataSourceModel struct {
	AccountID  types.String                                                         `tfsdk:"account_id" path:"account_id,optional"`
	EntryID    types.String                                                         `tfsdk:"entry_id" path:"entry_id,optional"`
	CreatedAt  timetypes.RFC3339                                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Enabled    types.Bool                                                           `tfsdk:"enabled" json:"enabled,computed"`
	ID         types.String                                                         `tfsdk:"id" json:"id,computed"`
	Name       types.String                                                         `tfsdk:"name" json:"name,computed"`
	ProfileID  types.String                                                         `tfsdk:"profile_id" json:"profile_id,computed"`
	Secret     types.Bool                                                           `tfsdk:"secret" json:"secret,computed"`
	Type       types.String                                                         `tfsdk:"type" json:"type,computed"`
	UpdatedAt  timetypes.RFC3339                                                    `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Confidence customfield.NestedObject[ZeroTrustDLPEntryConfidenceDataSourceModel] `tfsdk:"confidence" json:"confidence,computed"`
	Pattern    customfield.NestedObject[ZeroTrustDLPEntryPatternDataSourceModel]    `tfsdk:"pattern" json:"pattern,computed"`
	WordList   jsontypes.Normalized                                                 `tfsdk:"word_list" json:"word_list,computed"`
	Filter     *ZeroTrustDLPEntryFindOneByDataSourceModel                           `tfsdk:"filter"`
}

func (m *ZeroTrustDLPEntryDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DLPEntryGetParams, diags diag.Diagnostics) {
	params = zero_trust.DLPEntryGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustDLPEntryDataSourceModel) toListParams(_ context.Context) (params zero_trust.DLPEntryListParams, diags diag.Diagnostics) {
	params = zero_trust.DLPEntryListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPEntryConfidenceDataSourceModel struct {
	Available types.Bool `tfsdk:"available" json:"available,computed"`
}

type ZeroTrustDLPEntryPatternDataSourceModel struct {
	Regex      types.String `tfsdk:"regex" json:"regex,computed"`
	Validation types.String `tfsdk:"validation" json:"validation,computed"`
}

type ZeroTrustDLPEntryFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
