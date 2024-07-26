// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelResultDataSourceEnvelope struct {
	Result TunnelDataSourceModel `json:"result,computed"`
}

type TunnelResultListDataSourceEnvelope struct {
	Result *[]*TunnelDataSourceModel `json:"result,computed"`
}

type TunnelDataSourceModel struct {
	AccountID   types.String                         `tfsdk:"account_id" path:"account_id"`
	TunnelID    types.String                         `tfsdk:"tunnel_id" path:"tunnel_id"`
	ID          types.String                         `tfsdk:"id" json:"id"`
	Connections *[]*TunnelConnectionsDataSourceModel `tfsdk:"connections" json:"connections"`
	CreatedAt   types.String                         `tfsdk:"created_at" json:"created_at"`
	Name        types.String                         `tfsdk:"name" json:"name"`
	DeletedAt   types.String                         `tfsdk:"deleted_at" json:"deleted_at"`
	FindOneBy   *TunnelFindOneByDataSourceModel      `tfsdk:"find_one_by"`
}

type TunnelConnectionsDataSourceModel struct {
	ColoName           types.String `tfsdk:"colo_name" json:"colo_name"`
	IsPendingReconnect types.Bool   `tfsdk:"is_pending_reconnect" json:"is_pending_reconnect"`
	UUID               types.String `tfsdk:"uuid" json:"uuid,computed"`
}

type TunnelFindOneByDataSourceModel struct {
	AccountID     types.String  `tfsdk:"account_id" path:"account_id"`
	ExcludePrefix types.String  `tfsdk:"exclude_prefix" query:"exclude_prefix"`
	ExistedAt     types.String  `tfsdk:"existed_at" query:"existed_at"`
	IncludePrefix types.String  `tfsdk:"include_prefix" query:"include_prefix"`
	IsDeleted     types.Bool    `tfsdk:"is_deleted" query:"is_deleted"`
	Name          types.String  `tfsdk:"name" query:"name"`
	Page          types.Float64 `tfsdk:"page" query:"page"`
	PerPage       types.Float64 `tfsdk:"per_page" query:"per_page"`
	Status        types.String  `tfsdk:"status" query:"status"`
	TunTypes      types.String  `tfsdk:"tun_types" query:"tun_types"`
	UUID          types.String  `tfsdk:"uuid" query:"uuid"`
	WasActiveAt   types.String  `tfsdk:"was_active_at" query:"was_active_at"`
	WasInactiveAt types.String  `tfsdk:"was_inactive_at" query:"was_inactive_at"`
}
