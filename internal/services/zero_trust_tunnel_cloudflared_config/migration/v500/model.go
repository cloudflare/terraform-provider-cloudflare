package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source V4 Models (Legacy SDKv2 Provider)
// ============================================================================

// SourceV4TunnelConfigModel represents the zero_trust_tunnel_cloudflared_config state
// from the v4.x provider (SDKv2). Both cloudflare_tunnel_config (deprecated) and
// cloudflare_zero_trust_tunnel_cloudflared_config use the same schema.
// Schema version: 0 (default SDKv2, no explicit SchemaVersion defined)
type SourceV4TunnelConfigModel struct {
	ID        types.String            `tfsdk:"id"`
	TunnelID  types.String            `tfsdk:"tunnel_id"`
	AccountID types.String            `tfsdk:"account_id"`
	Config    []SourceV4ConfigModel   `tfsdk:"config"` // TypeList MaxItems:1 = array
}

// SourceV4ConfigModel represents the config block from v4.x (TypeList MaxItems:1).
type SourceV4ConfigModel struct {
	WarpRouting   []SourceV4WarpRoutingModel   `tfsdk:"warp_routing"`   // TypeList MaxItems:1 = array; removed in v5
	OriginRequest []SourceV4OriginRequestModel `tfsdk:"origin_request"` // TypeList MaxItems:1 = array
	IngressRule   []SourceV4IngressRuleModel   `tfsdk:"ingress_rule"`   // TypeList; renamed to ingress in v5
}

// SourceV4WarpRoutingModel represents the warp_routing block from v4.x.
// This block is removed entirely in v5.
type SourceV4WarpRoutingModel struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// SourceV4OriginRequestModel represents the origin_request block from v4.x.
// Used at both config-level and ingress-level.
type SourceV4OriginRequestModel struct {
	// Duration fields: stored as strings ("30s", "10s") in v4, converted to Int64 seconds in v5
	ConnectTimeout   types.String `tfsdk:"connect_timeout"`
	TLSTimeout       types.String `tfsdk:"tls_timeout"`
	TCPKeepAlive     types.String `tfsdk:"tcp_keep_alive"`
	KeepAliveTimeout types.String `tfsdk:"keep_alive_timeout"`

	// Int fields
	KeepAliveConnections types.Int64 `tfsdk:"keep_alive_connections"`

	// Bool fields
	NoHappyEyeballs        types.Bool `tfsdk:"no_happy_eyeballs"`
	NoTLSVerify            types.Bool `tfsdk:"no_tls_verify"`
	DisableChunkedEncoding types.Bool `tfsdk:"disable_chunked_encoding"`
	HTTP2Origin            types.Bool `tfsdk:"http2_origin"`

	// String fields
	HTTPHostHeader   types.String `tfsdk:"http_host_header"`
	OriginServerName types.String `tfsdk:"origin_server_name"`
	CAPool           types.String `tfsdk:"ca_pool"`
	ProxyType        types.String `tfsdk:"proxy_type"`

	// Removed in v5
	BastionMode  types.Bool   `tfsdk:"bastion_mode"`
	ProxyAddress types.String `tfsdk:"proxy_address"`
	ProxyPort    types.Int64  `tfsdk:"proxy_port"`

	// TypeSet of objects - removed in v5 (must parse but will be dropped)
	IPRules []SourceV4IPRuleModel `tfsdk:"ip_rules"`

	// TypeList MaxItems:1 = array
	Access []SourceV4AccessModel `tfsdk:"access"`
}

// SourceV4IPRuleModel represents ip_rules set entries from v4.x.
// This block is removed in v5.
type SourceV4IPRuleModel struct {
	Prefix types.String `tfsdk:"prefix"`
	Ports  types.List   `tfsdk:"ports"` // TypeList of TypeInt
	Allow  types.Bool   `tfsdk:"allow"`
}

// SourceV4AccessModel represents the access block from v4.x (TypeList MaxItems:1).
type SourceV4AccessModel struct {
	Required types.Bool   `tfsdk:"required"`
	TeamName types.String `tfsdk:"team_name"`
	AUDTag   types.Set    `tfsdk:"aud_tag"` // TypeSet of TypeString
}

// SourceV4IngressRuleModel represents ingress_rule entries from v4.x.
// Renamed to ingress in v5.
type SourceV4IngressRuleModel struct {
	Hostname      types.String                 `tfsdk:"hostname"`
	Path          types.String                 `tfsdk:"path"`
	Service       types.String                 `tfsdk:"service"`
	OriginRequest []SourceV4OriginRequestModel `tfsdk:"origin_request"` // TypeList MaxItems:1 = array
}

// ============================================================================
// Target V5 Models (Current Plugin Framework Provider)
// ============================================================================

// TargetV5TunnelConfigModel represents the zero_trust_tunnel_cloudflared_config
// state in the v5.x provider (Plugin Framework). Schema version: 500.
//
// Note: Config uses *T (pointer) even though the schema uses customfield.NewNestedObjectType.
// This follows the page_rule migration pattern and works because the framework
// converts the pointer struct to the appropriate object type during state serialization.
type TargetV5TunnelConfigModel struct {
	ID        types.String           `tfsdk:"id"`
	TunnelID  types.String           `tfsdk:"tunnel_id"`
	AccountID types.String           `tfsdk:"account_id"`
	Config    *TargetV5ConfigModel   `tfsdk:"config"`
	CreatedAt timetypes.RFC3339      `tfsdk:"created_at"` // v5-only, computed; set to null
	Source    types.String           `tfsdk:"source"`     // v5-only, computed/optional; set to null
	Version   types.Int64            `tfsdk:"version"`    // v5-only, computed; set to null
}

// TargetV5ConfigModel represents the config object in v5.x (SingleNestedAttribute).
type TargetV5ConfigModel struct {
	Ingress       *[]*TargetV5IngressModel       `tfsdk:"ingress"`        // ListNestedAttribute (from ingress_rule)
	OriginRequest *TargetV5OriginRequestModel     `tfsdk:"origin_request"` // SingleNestedAttribute
}

// TargetV5IngressModel represents a single ingress item in v5.x.
type TargetV5IngressModel struct {
	Hostname      types.String                        `tfsdk:"hostname"`
	Service       types.String                        `tfsdk:"service"`
	OriginRequest *TargetV5IngressOriginRequestModel  `tfsdk:"origin_request"`
	Path          types.String                        `tfsdk:"path"`
}

// TargetV5IngressOriginRequestModel represents the origin_request within an ingress item.
// Matches ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestModel.
type TargetV5IngressOriginRequestModel struct {
	Access                 *TargetV5AccessModel `tfsdk:"access"`
	CAPool                 types.String         `tfsdk:"ca_pool"`
	ConnectTimeout         types.Int64          `tfsdk:"connect_timeout"`
	DisableChunkedEncoding types.Bool           `tfsdk:"disable_chunked_encoding"`
	HTTP2Origin            types.Bool           `tfsdk:"http2_origin"`
	HTTPHostHeader         types.String         `tfsdk:"http_host_header"`
	KeepAliveConnections   types.Int64          `tfsdk:"keep_alive_connections"`
	KeepAliveTimeout       types.Int64          `tfsdk:"keep_alive_timeout"`
	MatchSnItoHost         types.Bool           `tfsdk:"match_sn_ito_host"` // v5-only
	NoHappyEyeballs        types.Bool           `tfsdk:"no_happy_eyeballs"`
	NoTLSVerify            types.Bool           `tfsdk:"no_tls_verify"`
	OriginServerName       types.String         `tfsdk:"origin_server_name"`
	ProxyType              types.String         `tfsdk:"proxy_type"`
	TCPKeepAlive           types.Int64          `tfsdk:"tcp_keep_alive"`
	TLSTimeout             types.Int64          `tfsdk:"tls_timeout"`
}

// TargetV5OriginRequestModel represents the config-level origin_request in v5.x.
// Matches ZeroTrustTunnelCloudflaredConfigConfigOriginRequestModel.
type TargetV5OriginRequestModel struct {
	Access                 *TargetV5AccessModel `tfsdk:"access"`
	CAPool                 types.String         `tfsdk:"ca_pool"`
	ConnectTimeout         types.Int64          `tfsdk:"connect_timeout"`
	DisableChunkedEncoding types.Bool           `tfsdk:"disable_chunked_encoding"`
	HTTP2Origin            types.Bool           `tfsdk:"http2_origin"`
	HTTPHostHeader         types.String         `tfsdk:"http_host_header"`
	KeepAliveConnections   types.Int64          `tfsdk:"keep_alive_connections"`
	KeepAliveTimeout       types.Int64          `tfsdk:"keep_alive_timeout"`
	MatchSnItoHost         types.Bool           `tfsdk:"match_sn_ito_host"` // v5-only
	NoHappyEyeballs        types.Bool           `tfsdk:"no_happy_eyeballs"`
	NoTLSVerify            types.Bool           `tfsdk:"no_tls_verify"`
	OriginServerName       types.String         `tfsdk:"origin_server_name"`
	ProxyType              types.String         `tfsdk:"proxy_type"`
	TCPKeepAlive           types.Int64          `tfsdk:"tcp_keep_alive"`
	TLSTimeout             types.Int64          `tfsdk:"tls_timeout"`
}

// TargetV5AccessModel represents the access block in v5.x (SingleNestedAttribute).
// Matches ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestAccessModel
// and ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessModel.
type TargetV5AccessModel struct {
	AUDTag   *[]types.String `tfsdk:"aud_tag"`   // v4 TypeSet → v5 ListAttribute
	TeamName types.String    `tfsdk:"team_name"`
	Required types.Bool      `tfsdk:"required"`
}
