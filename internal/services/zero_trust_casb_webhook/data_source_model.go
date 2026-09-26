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

type ZeroTrustCasbWebhookResultDataSourceEnvelope struct {
	Result ZeroTrustCasbWebhookDataSourceModel `json:"result,computed"`
}

type ZeroTrustCasbWebhookDataSourceModel struct {
	ID                 types.String                                                             `tfsdk:"id" path:"webhook_id,computed"`
	WebhookID          types.String                                                             `tfsdk:"webhook_id" path:"webhook_id,required"`
	AccountID          types.String                                                             `tfsdk:"account_id" path:"account_id,required"`
	AuthenticationType types.String                                                             `tfsdk:"authentication_type" json:"authentication_type,computed"`
	CreatedAt          timetypes.RFC3339                                                        `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DestinationURL     types.String                                                             `tfsdk:"destination_url" json:"destination_url,computed"`
	Label              types.String                                                             `tfsdk:"label" json:"label,computed"`
	Status             types.String                                                             `tfsdk:"status" json:"status,computed"`
	UpdatedAt          timetypes.RFC3339                                                        `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Version            types.Int64                                                              `tfsdk:"version" json:"version,computed"`
	Headers            customfield.NestedObjectList[ZeroTrustCasbWebhookHeadersDataSourceModel] `tfsdk:"headers" json:"headers,computed"`
}

func (m *ZeroTrustCasbWebhookDataSourceModel) toReadParams(_ context.Context) (params zero_trust.CasbPostureWebhookGetParams, diags diag.Diagnostics) {
	params = zero_trust.CasbPostureWebhookGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustCasbWebhookHeadersDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}
