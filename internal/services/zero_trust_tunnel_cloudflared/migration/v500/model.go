package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceTunnelCloudflaredModel mirrors the v4 SDKv2 state for cloudflare_tunnel /
// cloudflare_zero_trust_tunnel.
type SourceTunnelCloudflaredModel struct {
	ID          types.String `tfsdk:"id"`
	AccountID   types.String `tfsdk:"account_id"`
	Name        types.String `tfsdk:"name"`
	Secret      types.String `tfsdk:"secret"`
	ConfigSrc   types.String `tfsdk:"config_src"`
	Cname       types.String `tfsdk:"cname"`
	TunnelToken types.String `tfsdk:"tunnel_token"`
}

// TargetTunnelCloudflaredModel mirrors the v5 Plugin Framework model for
// cloudflare_zero_trust_tunnel_cloudflared. Defined locally to avoid an import
// cycle with the parent service package.
type TargetTunnelCloudflaredModel struct {
	ID              types.String                                                          `tfsdk:"id"`
	AccountID       types.String                                                          `tfsdk:"account_id"`
	ConfigSrc       types.String                                                          `tfsdk:"config_src"`
	Name            types.String                                                          `tfsdk:"name"`
	TunnelSecret    types.String                                                          `tfsdk:"tunnel_secret"`
	AccountTag      types.String                                                          `tfsdk:"account_tag"`
	ConnsActiveAt   timetypes.RFC3339                                                     `tfsdk:"conns_active_at"`
	ConnsInactiveAt timetypes.RFC3339                                                     `tfsdk:"conns_inactive_at"`
	CreatedAt       timetypes.RFC3339                                                     `tfsdk:"created_at"`
	DeletedAt       timetypes.RFC3339                                                     `tfsdk:"deleted_at"`
	RemoteConfig    types.Bool                                                            `tfsdk:"remote_config"`
	Status          types.String                                                          `tfsdk:"status"`
	TunType         types.String                                                          `tfsdk:"tun_type"`
	Connections     customfield.NestedObjectList[TargetTunnelCloudflaredConnectionsModel] `tfsdk:"connections"`
	Metadata        jsontypes.Normalized                                                  `tfsdk:"metadata"`
}

// TargetTunnelCloudflaredConnectionsModel mirrors the v5 connections nested object.
type TargetTunnelCloudflaredConnectionsModel struct {
	ID                 types.String      `tfsdk:"id"`
	ClientID           types.String      `tfsdk:"client_id"`
	ClientVersion      types.String      `tfsdk:"client_version"`
	ColoName           types.String      `tfsdk:"colo_name"`
	IsPendingReconnect types.Bool        `tfsdk:"is_pending_reconnect"`
	OpenedAt           timetypes.RFC3339 `tfsdk:"opened_at"`
	OriginIP           types.String      `tfsdk:"origin_ip"`
	UUID               types.String      `tfsdk:"uuid"`
}
