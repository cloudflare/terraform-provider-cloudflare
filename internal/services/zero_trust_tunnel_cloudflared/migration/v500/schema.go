package v500

import "github.com/hashicorp/terraform-plugin-framework/resource/schema"

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
