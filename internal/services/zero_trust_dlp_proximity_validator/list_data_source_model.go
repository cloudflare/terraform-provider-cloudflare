// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_proximity_validator

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPProximityValidatorsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDLPProximityValidatorsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDLPProximityValidatorsDataSourceModel struct {
	AccountID types.String                                                                       `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                        `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustDLPProximityValidatorsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDLPProximityValidatorsDataSourceModel) toListParams(_ context.Context) (params zero_trust.DLPValidatorProximityWordListParams, diags diag.Diagnostics) {
	params = zero_trust.DLPValidatorProximityWordListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPProximityValidatorsResultDataSourceModel struct {
	ID           types.String                                                                         `tfsdk:"id" json:"id,computed"`
	CreatedAt    timetypes.RFC3339                                                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name         types.String                                                                         `tfsdk:"name" json:"name,computed"`
	ReferencedBy customfield.NestedObject[ZeroTrustDLPProximityValidatorsReferencedByDataSourceModel] `tfsdk:"referenced_by" json:"referenced_by,computed"`
	UpdatedAt    timetypes.RFC3339                                                                    `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Words        customfield.NestedObjectSet[ZeroTrustDLPProximityValidatorsWordsDataSourceModel]     `tfsdk:"words" json:"words,computed"`
}

type ZeroTrustDLPProximityValidatorsReferencedByDataSourceModel struct {
	Entries    customfield.NestedObjectList[ZeroTrustDLPProximityValidatorsReferencedByEntriesDataSourceModel] `tfsdk:"entries" json:"entries,computed"`
	EntryCount types.Int64                                                                                     `tfsdk:"entry_count" json:"entry_count,computed"`
}

type ZeroTrustDLPProximityValidatorsReferencedByEntriesDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type ZeroTrustDLPProximityValidatorsWordsDataSourceModel struct {
	Weight types.String `tfsdk:"weight" json:"weight,computed"`
	Word   types.String `tfsdk:"word" json:"word,computed"`
}
