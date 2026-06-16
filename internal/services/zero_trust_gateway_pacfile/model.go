// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_pacfile

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayPacfileResultEnvelope struct {
	Result ZeroTrustGatewayPacfileModel `json:"result"`
}

type ZeroTrustGatewayPacfileModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	AccountID   types.String      `tfsdk:"account_id" path:"account_id,required"`
	Slug        types.String      `tfsdk:"slug" json:"slug,optional"`
	Contents    types.String      `tfsdk:"contents" json:"contents,required"`
	Name        types.String      `tfsdk:"name" json:"name,required"`
	Description types.String      `tfsdk:"description" json:"description,optional"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	URL         types.String      `tfsdk:"url" json:"url,computed"`
}

func (m ZeroTrustGatewayPacfileModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustGatewayPacfileModel) MarshalJSONForUpdate(state ZeroTrustGatewayPacfileModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
