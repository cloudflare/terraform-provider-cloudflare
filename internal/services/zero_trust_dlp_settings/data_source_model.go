// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPSettingsResultDataSourceEnvelope struct {
	Result ZeroTrustDLPSettingsDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPSettingsDataSourceModel struct {
	ID                types.String                                                                `tfsdk:"id" path:"account_id,computed"`
	AccountID         types.String                                                                `tfsdk:"account_id" path:"account_id,required"`
	AIContextAnalysis types.Bool                                                                  `tfsdk:"ai_context_analysis" json:"ai_context_analysis,computed"`
	OCR               types.Bool                                                                  `tfsdk:"ocr" json:"ocr,computed"`
	PayloadLogging    customfield.NestedObject[ZeroTrustDLPSettingsPayloadLoggingDataSourceModel] `tfsdk:"payload_logging" json:"payload_logging,computed"`
}

func (m *ZeroTrustDLPSettingsDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DLPSettingGetParams, diags diag.Diagnostics) {
	params = zero_trust.DLPSettingGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPSettingsPayloadLoggingDataSourceModel struct {
	UpdatedAt    timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	MaskingLevel types.String      `tfsdk:"masking_level" json:"masking_level,computed"`
	PublicKey    types.String      `tfsdk:"public_key" json:"public_key,computed"`
}
