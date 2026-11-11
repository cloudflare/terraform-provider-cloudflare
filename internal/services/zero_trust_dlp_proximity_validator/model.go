// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_proximity_validator

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPProximityValidatorResultEnvelope struct {
	Result ZeroTrustDLPProximityValidatorModel `json:"result"`
}

type ZeroTrustDLPProximityValidatorModel struct {
	ID           types.String                                                              `tfsdk:"id" json:"id,computed"`
	AccountID    types.String                                                              `tfsdk:"account_id" path:"account_id,required"`
	Name         types.String                                                              `tfsdk:"name" json:"name,required"`
	Words        *[]*ZeroTrustDLPProximityValidatorWordsModel                              `tfsdk:"words" json:"words,required"`
	CreatedAt    timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt    timetypes.RFC3339                                                         `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	ReferencedBy customfield.NestedObject[ZeroTrustDLPProximityValidatorReferencedByModel] `tfsdk:"referenced_by" json:"referenced_by,computed"`
}

func (m ZeroTrustDLPProximityValidatorModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDLPProximityValidatorModel) MarshalJSONForUpdate(state ZeroTrustDLPProximityValidatorModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDLPProximityValidatorWordsModel struct {
	Weight types.String `tfsdk:"weight" json:"weight,required"`
	Word   types.String `tfsdk:"word" json:"word,required"`
}

type ZeroTrustDLPProximityValidatorReferencedByModel struct {
	Entries    customfield.NestedObjectList[ZeroTrustDLPProximityValidatorReferencedByEntriesModel] `tfsdk:"entries" json:"entries,computed"`
	EntryCount types.Int64                                                                          `tfsdk:"entry_count" json:"entry_count,computed"`
}

type ZeroTrustDLPProximityValidatorReferencedByEntriesModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}
