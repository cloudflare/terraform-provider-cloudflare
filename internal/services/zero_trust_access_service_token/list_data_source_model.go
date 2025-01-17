// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_service_token

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessServiceTokensResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessServiceTokensResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessServiceTokensDataSourceModel struct {
	AccountID types.String                                                                    `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                                    `tfsdk:"zone_id" path:"zone_id,optional"`
	Name      types.String                                                                    `tfsdk:"name" query:"name,optional"`
	Search    types.String                                                                    `tfsdk:"search" query:"search,optional"`
	MaxItems  types.Int64                                                                     `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustAccessServiceTokensResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessServiceTokensDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessServiceTokenListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessServiceTokenListParams{}

	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessServiceTokensResultDataSourceModel struct {
	ID         types.String      `tfsdk:"id" json:"id,computed"`
	ClientID   types.String      `tfsdk:"client_id" json:"client_id,computed"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Duration   types.String      `tfsdk:"duration" json:"duration,computed"`
	ExpiresAt  timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	LastSeenAt timetypes.RFC3339 `tfsdk:"last_seen_at" json:"last_seen_at,computed" format:"date-time"`
	Name       types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt  timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}
