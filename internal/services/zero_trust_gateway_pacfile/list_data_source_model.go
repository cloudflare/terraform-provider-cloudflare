// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_pacfile

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayPacfilesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustGatewayPacfilesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustGatewayPacfilesDataSourceModel struct {
	AccountID types.String                                                                `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                 `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustGatewayPacfilesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustGatewayPacfilesDataSourceModel) toListParams(_ context.Context) (params zero_trust.GatewayPacfileListParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayPacfileListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustGatewayPacfilesResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	Slug        types.String      `tfsdk:"slug" json:"slug,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	URL         types.String      `tfsdk:"url" json:"url,computed"`
}
