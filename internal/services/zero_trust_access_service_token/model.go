// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_service_token

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessServiceTokenResultEnvelope struct {
	Result ZeroTrustAccessServiceTokenModel `json:"result,computed"`
}

type ZeroTrustAccessServiceTokenModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	AccountID    types.String      `tfsdk:"account_id" path:"account_id"`
	ZoneID       types.String      `tfsdk:"zone_id" path:"zone_id"`
	Name         types.String      `tfsdk:"name" json:"name"`
	Duration     types.String      `tfsdk:"duration" json:"duration"`
	ClientID     types.String      `tfsdk:"client_id" json:"client_id,computed"`
	ClientSecret types.String      `tfsdk:"client_secret" json:"client_secret,computed"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	ExpiresAt    timetypes.RFC3339 `tfsdk:"expires_at" json:"expires_at,computed"`
	UpdatedAt    timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
}
