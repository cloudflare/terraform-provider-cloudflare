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

type ZeroTrustDLPProximityValidatorResultDataSourceEnvelope struct {
	Result ZeroTrustDLPProximityValidatorDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPProximityValidatorDataSourceModel struct {
	ID           types.String                                                                        `tfsdk:"id" path:"validator_id,computed"`
	ValidatorID  types.String                                                                        `tfsdk:"validator_id" path:"validator_id,required"`
	AccountID    types.String                                                                        `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt    timetypes.RFC3339                                                                   `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name         types.String                                                                        `tfsdk:"name" json:"name,computed"`
	UpdatedAt    timetypes.RFC3339                                                                   `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	ReferencedBy customfield.NestedObject[ZeroTrustDLPProximityValidatorReferencedByDataSourceModel] `tfsdk:"referenced_by" json:"referenced_by,computed"`
	Words        customfield.NestedObjectSet[ZeroTrustDLPProximityValidatorWordsDataSourceModel]     `tfsdk:"words" json:"words,computed"`
}

func (m *ZeroTrustDLPProximityValidatorDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DLPValidatorProximityWordGetParams, diags diag.Diagnostics) {
	params = zero_trust.DLPValidatorProximityWordGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPProximityValidatorReferencedByDataSourceModel struct {
	Entries    customfield.NestedObjectList[ZeroTrustDLPProximityValidatorReferencedByEntriesDataSourceModel] `tfsdk:"entries" json:"entries,computed"`
	EntryCount types.Int64                                                                                    `tfsdk:"entry_count" json:"entry_count,computed"`
}

type ZeroTrustDLPProximityValidatorReferencedByEntriesDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type ZeroTrustDLPProximityValidatorWordsDataSourceModel struct {
	Weight types.String `tfsdk:"weight" json:"weight,computed"`
	Word   types.String `tfsdk:"word" json:"word,computed"`
}
