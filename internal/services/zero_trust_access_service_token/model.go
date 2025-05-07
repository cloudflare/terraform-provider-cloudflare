// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_service_token

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessServiceTokenResultEnvelope struct {
	Result ZeroTrustAccessServiceTokenModel `json:"result"`
}

type ZeroTrustAccessServiceTokenModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	AccountID    types.String      `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID       types.String      `tfsdk:"zone_id" path:"zone_id,optional"`
	Name         types.String      `tfsdk:"name" json:"name,required"`
	Duration     types.String      `tfsdk:"duration" json:"duration,computed_optional"`
	ClientID     types.String      `tfsdk:"client_id" json:"client_id,computed"`
	ClientSecret types.String      `tfsdk:"client_secret" json:"client_secret,computed,no_refresh"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ExpiresAt    timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	LastSeenAt   timetypes.RFC3339 `tfsdk:"last_seen_at" json:"last_seen_at,computed" format:"date-time"`
	UpdatedAt    timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustAccessServiceTokenModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessServiceTokenModel) MarshalJSONForUpdate(state ZeroTrustAccessServiceTokenModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
