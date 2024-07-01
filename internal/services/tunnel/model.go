// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelResultEnvelope struct {
	Result TunnelModel `json:"result,computed"`
}

type TunnelResultDataSourceEnvelope struct {
	Result TunnelDataSourceModel `json:"result,computed"`
}

type TunnelsResultDataSourceEnvelope struct {
	Result TunnelsDataSourceModel `json:"result,computed"`
}

type TunnelModel struct {
	ID           types.String               `tfsdk:"id" json:"id,computed"`
	AccountID    types.String               `tfsdk:"account_id" path:"account_id"`
	Name         types.String               `tfsdk:"name" json:"name"`
	TunnelSecret types.String               `tfsdk:"tunnel_secret" json:"tunnel_secret"`
	Connections  *[]*TunnelConnectionsModel `tfsdk:"connections" json:"connections,computed"`
	CreatedAt    types.String               `tfsdk:"created_at" json:"created_at,computed"`
	DeletedAt    types.String               `tfsdk:"deleted_at" json:"deleted_at,computed"`
}

type TunnelConnectionsModel struct {
	ColoName           types.String `tfsdk:"colo_name" json:"colo_name"`
	IsPendingReconnect types.Bool   `tfsdk:"is_pending_reconnect" json:"is_pending_reconnect"`
	UUID               types.String `tfsdk:"uuid" json:"uuid,computed"`
}

type TunnelDataSourceModel struct {
}

type TunnelsDataSourceModel struct {
}
