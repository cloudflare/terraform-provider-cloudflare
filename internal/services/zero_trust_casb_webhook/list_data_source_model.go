// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_webhook

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCasbWebhooksResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustCasbWebhooksResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustCasbWebhooksDataSourceModel struct {
	AccountID types.String                                                             `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustCasbWebhooksResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustCasbWebhooksDataSourceModel) toListParams(_ context.Context) (params zero_trust.CasbPostureWebhookListParams, diags diag.Diagnostics) {
	params = zero_trust.CasbPostureWebhookListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustCasbWebhooksResultDataSourceModel struct {
	ID                 types.String                                                              `tfsdk:"id" json:"id,computed"`
	AuthenticationType types.String                                                              `tfsdk:"authentication_type" json:"authentication_type,computed"`
	CreatedAt          timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DestinationURL     types.String                                                              `tfsdk:"destination_url" json:"destination_url,computed"`
	Label              types.String                                                              `tfsdk:"label" json:"label,computed"`
	Status             types.String                                                              `tfsdk:"status" json:"status,computed"`
	UpdatedAt          timetypes.RFC3339                                                         `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Version            types.Int64                                                               `tfsdk:"version" json:"version,computed"`
	Headers            customfield.NestedObjectList[ZeroTrustCasbWebhooksHeadersDataSourceModel] `tfsdk:"headers" json:"headers,computed"`
}

type ZeroTrustCasbWebhooksHeadersDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}
