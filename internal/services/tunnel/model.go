// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelResultEnvelope struct {
	Result TunnelModel `json:"result,computed"`
}

type TunnelModel struct {
	ID           types.String               `tfsdk:"id" json:"id,computed"`
	AccountID    types.String               `tfsdk:"account_id" path:"account_id"`
	Name         types.String               `tfsdk:"name" json:"name"`
	TunnelSecret types.String               `tfsdk:"tunnel_secret" json:"tunnel_secret"`
	CreatedAt    timetypes.RFC3339          `tfsdk:"created_at" json:"created_at,computed"`
	DeletedAt    timetypes.RFC3339          `tfsdk:"deleted_at" json:"deleted_at,computed"`
	Connections  *[]*TunnelConnectionsModel `tfsdk:"connections" json:"connections,computed"`
}

type TunnelConnectionsModel struct {
	ColoName           types.String `tfsdk:"colo_name" json:"colo_name"`
	IsPendingReconnect types.Bool   `tfsdk:"is_pending_reconnect" json:"is_pending_reconnect"`
	UUID               types.String `tfsdk:"uuid" json:"uuid,computed"`
}
