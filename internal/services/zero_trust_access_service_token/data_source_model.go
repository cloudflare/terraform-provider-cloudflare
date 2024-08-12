// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_service_token

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessServiceTokenResultDataSourceEnvelope struct {
	Result ZeroTrustAccessServiceTokenDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessServiceTokenResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustAccessServiceTokenDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessServiceTokenDataSourceModel struct {
	AccountID      types.String      `tfsdk:"account_id" path:"account_id"`
	ServiceTokenID types.String      `tfsdk:"service_token_id" path:"service_token_id"`
	ZoneID         types.String      `tfsdk:"zone_id" path:"zone_id"`
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	Duration       types.String      `tfsdk:"duration" json:"duration,computed"`
	UpdatedAt      timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
	ClientID       types.String      `tfsdk:"client_id" json:"client_id"`
	ExpiresAt      timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at"`
	ID             types.String      `tfsdk:"id" json:"id"`
	Name           types.String      `tfsdk:"name" json:"name"`
}
