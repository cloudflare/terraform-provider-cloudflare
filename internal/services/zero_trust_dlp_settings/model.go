// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPSettingsResultEnvelope struct {
	Result ZeroTrustDLPSettingsModel `json:"result"`
}

type ZeroTrustDLPSettingsModel struct {
	ID                types.String                                                      `tfsdk:"id" json:"-,computed"`
	AccountID         types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	AIContextAnalysis types.Bool                                                        `tfsdk:"ai_context_analysis" json:"ai_context_analysis,computed_optional"`
	OCR               types.Bool                                                        `tfsdk:"ocr" json:"ocr,computed_optional"`
	PayloadLogging    customfield.NestedObject[ZeroTrustDLPSettingsPayloadLoggingModel] `tfsdk:"payload_logging" json:"payload_logging,computed_optional"`
}

func (m ZeroTrustDLPSettingsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDLPSettingsModel) MarshalJSONForUpdate(state ZeroTrustDLPSettingsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustDLPSettingsPayloadLoggingModel struct {
	MaskingLevel types.String `tfsdk:"masking_level" json:"masking_level,computed_optional"`
	PublicKey    types.String `tfsdk:"public_key" json:"public_key,optional"`
}
