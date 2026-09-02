package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceTunnelCloudflaredSchema returns the minimal v4 schema used to parse
// cloudflare_tunnel / cloudflare_zero_trust_tunnel state during MoveState and
// UpgradeState. Version is not set (defaults to 0, matching SDKv2 behaviour).
func SourceTunnelCloudflaredSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id":           schema.StringAttribute{Computed: true},
			"account_id":   schema.StringAttribute{Required: true},
			"name":         schema.StringAttribute{Required: true},
			"secret":       schema.StringAttribute{Required: true, Sensitive: true},
			"config_src":   schema.StringAttribute{Optional: true},
			"cname":        schema.StringAttribute{Computed: true},
			"tunnel_token": schema.StringAttribute{Computed: true, Sensitive: true},
		},
	}
}

// OldTargetTunnelCloudflaredSchema returns the v5.0–v5.23 schema for
// cloudflare_zero_trust_tunnel_cloudflared. It is used as PriorSchema for the
// version-1 → version-500 UpgradeState handler so that the framework can
// correctly decode state written by those older provider releases.
//
// The critical difference from the current schema is that the connections list
// nested object includes is_pending_reconnect, which was removed in v5.24.0.
func OldTargetTunnelCloudflaredSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id":               schema.StringAttribute{Computed: true},
			"account_id":       schema.StringAttribute{Required: true},
			"config_src":       schema.StringAttribute{Optional: true, Computed: true},
			"name":             schema.StringAttribute{Required: true},
			"tunnel_secret":    schema.StringAttribute{Optional: true, Sensitive: true},
			"account_tag":      schema.StringAttribute{Computed: true},
			"conns_active_at":  schema.StringAttribute{Computed: true, CustomType: timetypes.RFC3339Type{}},
			"conns_inactive_at": schema.StringAttribute{Computed: true, CustomType: timetypes.RFC3339Type{}},
			"created_at":       schema.StringAttribute{Computed: true, CustomType: timetypes.RFC3339Type{}},
			"deleted_at":       schema.StringAttribute{Computed: true, CustomType: timetypes.RFC3339Type{}},
			"remote_config":    schema.BoolAttribute{Computed: true},
			"status":           schema.StringAttribute{Computed: true},
			"tun_type":         schema.StringAttribute{Computed: true},
			"metadata":         schema.StringAttribute{Computed: true, CustomType: jsontypes.NormalizedType{}},
			"connections": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[TargetTunnelCloudflaredConnectionsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id":                   schema.StringAttribute{Computed: true},
						"client_id":            schema.StringAttribute{Computed: true},
						"client_version":       schema.StringAttribute{Computed: true},
						"colo_name":            schema.StringAttribute{Computed: true},
						"is_pending_reconnect": schema.BoolAttribute{Computed: true},
						"opened_at":            schema.StringAttribute{Computed: true, CustomType: timetypes.RFC3339Type{}},
						"origin_ip":            schema.StringAttribute{Computed: true},
						"uuid":                 schema.StringAttribute{Computed: true},
					},
				},
			},
		},
	}
}
