// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_webhook

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCasbWebhookResultEnvelope struct {
	Result ZeroTrustCasbWebhookModel `json:"result"`
}

type ZeroTrustCasbWebhookModel struct {
	ID                 types.String                         `tfsdk:"id" json:"id,computed"`
	AccountID          types.String                         `tfsdk:"account_id" path:"account_id,required"`
	AuthenticationType types.String                         `tfsdk:"authentication_type" json:"authentication_type,required"`
	DestinationURL     types.String                         `tfsdk:"destination_url" json:"destination_url,required"`
	Label              types.String                         `tfsdk:"label" json:"label,required"`
	SigningSecret      types.String                         `tfsdk:"signing_secret" json:"signing_secret,optional,no_refresh"`
	Headers            *[]*ZeroTrustCasbWebhookHeadersModel `tfsdk:"headers" json:"headers,optional"`
	Status             types.String                         `tfsdk:"status" json:"status,computed_optional"`
	CreatedAt          timetypes.RFC3339                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt          timetypes.RFC3339                    `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Version            types.Int64                          `tfsdk:"version" json:"version,computed"`
}

func (m ZeroTrustCasbWebhookModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustCasbWebhookModel) MarshalJSONForUpdate(state ZeroTrustCasbWebhookModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustCasbWebhookHeadersModel struct {
	Key   types.String `tfsdk:"key" json:"key,required"`
	Value types.String `tfsdk:"value" json:"value,optional"`
}
