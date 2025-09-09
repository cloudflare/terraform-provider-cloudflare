package zero_trust_connectivity_settings

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func (r *ConnectivitySettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages Zero Trust Connectivity Settings for an account.",
		Attributes: map[string]schema.Attribute{
			consts.AccountIDSchemaKey: schema.StringAttribute{
				Description: consts.AccountIDSchemaDescription,
				Required:    true,
			},
			"icmp_proxy_enabled": schema.BoolAttribute{
				Description: "A flag to enable the ICMP proxy for the account network.",
				Required:    true,
			},
			"offramp_warp_enabled": schema.BoolAttribute{
				Description: "A flag to enable WARP to WARP traffic.",
				Required:    true,
			},
		},
	}
}
