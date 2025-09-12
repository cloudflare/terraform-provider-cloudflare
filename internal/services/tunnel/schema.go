// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID of the tunnel.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Cloudflare account ID",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"config_src": schema.StringAttribute{
				Description: "Indicates if this is a locally or remotely configured tunnel. If `local`, manage the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the tunnel on the Zero Trust dashboard.\nAvailable values: \"local\", \"cloudflare\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("local", "cloudflare"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
				Default:       stringdefault.StaticString("local"),
			},
			"name": schema.StringAttribute{
				Description: "A user-friendly name for a tunnel.",
				Required:    true,
			},
			"tunnel_secret": schema.StringAttribute{
				Description: "Sets the password required to run a locally-managed tunnel. Must be at least 32 bytes and encoded as a base64 string.",
				Optional:    true,
				Sensitive:   true,
			},
			"account_tag": schema.StringAttribute{
				Description: "Cloudflare account ID",
				Computed:    true,
			},
			"conns_active_at": schema.StringAttribute{
				Description: "Timestamp of when the tunnel established at least one connection to Cloudflare's edge. If `null`, the tunnel is inactive.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"conns_inactive_at": schema.StringAttribute{
				Description: "Timestamp of when the tunnel became inactive (no connections to Cloudflare's edge). If `null`, the tunnel is active.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"created_at": schema.StringAttribute{
				Description: "Timestamp of when the resource was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"deleted_at": schema.StringAttribute{
				Description: "Timestamp of when the resource was deleted. If `null`, the resource has not been deleted.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"remote_config": schema.BoolAttribute{
				Description: "If `true`, the tunnel can be configured remotely from the Zero Trust dashboard. If `false`, the tunnel must be configured locally on the origin machine.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "The status of the tunnel. Valid values are `inactive` (tunnel has never been run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy state), `healthy` (tunnel is active and able to serve traffic), or `down` (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).\nAvailable values: \"inactive\", \"degraded\", \"healthy\", \"down\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"inactive",
						"degraded",
						"healthy",
						"down",
					),
				},
			},
			"tun_type": schema.StringAttribute{
				Description: "The type of tunnel.\nAvailable values: \"cfd_tunnel\", \"warp_connector\", \"warp\", \"magic\", \"ip_sec\", \"gre\", \"cni\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"cfd_tunnel",
						"warp_connector",
						"warp",
						"magic",
						"ip_sec",
						"gre",
						"cni",
					),
				},
			},
			"cname": schema.StringAttribute{
				Description: "The DNS name that can be used to route traffic to the tunnel.",
				Computed:    true,
			},
			"tunnel_token": schema.StringAttribute{
				Description: "The token used by the tunnel to authenticate with Cloudflare.",
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}
