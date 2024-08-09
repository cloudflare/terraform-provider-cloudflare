// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredResultEnvelope struct {
	Result ZeroTrustTunnelCloudflaredModel `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredModel struct {
	ID           types.String                                   `tfsdk:"id" json:"id,computed"`
	AccountID    types.String                                   `tfsdk:"account_id" path:"account_id"`
	Name         types.String                                   `tfsdk:"name" json:"name"`
	TunnelSecret types.String                                   `tfsdk:"tunnel_secret" json:"tunnel_secret"`
	CreatedAt    timetypes.RFC3339                              `tfsdk:"created_at" json:"created_at,computed"`
	DeletedAt    timetypes.RFC3339                              `tfsdk:"deleted_at" json:"deleted_at,computed"`
	Connections  *[]*ZeroTrustTunnelCloudflaredConnectionsModel `tfsdk:"connections" json:"connections,computed"`
}

type ZeroTrustTunnelCloudflaredConnectionsModel struct {
	ColoName           types.String `tfsdk:"colo_name" json:"colo_name"`
	IsPendingReconnect types.Bool   `tfsdk:"is_pending_reconnect" json:"is_pending_reconnect"`
	UUID               types.String `tfsdk:"uuid" json:"uuid,computed"`
}
