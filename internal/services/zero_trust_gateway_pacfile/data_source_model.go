// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_pacfile

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayPacfileResultDataSourceEnvelope struct {
	Result ZeroTrustGatewayPacfileDataSourceModel `json:"result,computed"`
}

type ZeroTrustGatewayPacfileDataSourceModel struct {
	ID          types.String      `tfsdk:"id" path:"pacfile_id,computed"`
	PacfileID   types.String      `tfsdk:"pacfile_id" path:"pacfile_id,required"`
	AccountID   types.String      `tfsdk:"account_id" path:"account_id,required"`
	Contents    types.String      `tfsdk:"contents" json:"contents,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	Slug        types.String      `tfsdk:"slug" json:"slug,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	URL         types.String      `tfsdk:"url" json:"url,computed"`
}

func (m *ZeroTrustGatewayPacfileDataSourceModel) toReadParams(_ context.Context) (params zero_trust.GatewayPacfileGetParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayPacfileGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
