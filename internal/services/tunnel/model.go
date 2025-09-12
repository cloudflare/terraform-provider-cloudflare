// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TunnelResultEnvelope struct {
	Result TunnelModel `json:"result"`
}

type TunnelModel struct {
	ID              types.String      `tfsdk:"id" json:"id,computed"`
	AccountID       types.String      `tfsdk:"account_id" path:"account_id,required"`
	ConfigSrc       types.String      `tfsdk:"config_src" json:"config_src,computed"`
	Name            types.String      `tfsdk:"name" json:"name,required"`
	TunnelSecret    types.String      `tfsdk:"tunnel_secret" json:"tunnel_secret,computed"`
	AccountTag      types.String      `tfsdk:"account_tag" json:"account_tag,computed"`
	ConnsActiveAt   timetypes.RFC3339 `tfsdk:"conns_active_at" json:"conns_active_at,computed"`
	ConnsInactiveAt timetypes.RFC3339 `tfsdk:"conns_inactive_at" json:"conns_inactive_at,computed"`
	CreatedAt       timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	DeletedAt       timetypes.RFC3339 `tfsdk:"deleted_at" json:"deleted_at,computed"`
	RemoteConfig    types.Bool        `tfsdk:"remote_config" json:"remote_config,computed"`
	Status          types.String      `tfsdk:"status" json:"status,computed"`
	TunType         types.String      `tfsdk:"tun_type" json:"tun_type,computed"`
	Cname           types.String      `tfsdk:"cname" json:"cname,computed"`
	TunnelToken     types.String      `tfsdk:"tunnel_token" json:"tunnel_token,computed"`
}
