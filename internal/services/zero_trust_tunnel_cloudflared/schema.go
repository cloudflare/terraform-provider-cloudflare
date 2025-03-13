// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustTunnelCloudflaredResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "UUID of the tunnel.",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "account_id": schema.StringAttribute{
        Description: "Cloudflare account ID",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "config_src": schema.StringAttribute{
        Description: "Indicates if this is a locally or remotely configured tunnel. If `local`, manage the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the tunnel on the Zero Trust dashboard.\nAvailable values: \"local\", \"cloudflare\".",
        Computed: true,
        Optional: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive("local", "cloudflare"),
        },
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
        Default: stringdefault.  StaticString("local"),
      },
      "name": schema.StringAttribute{
        Description: "A user-friendly name for a tunnel.",
        Required: true,
      },
      "tunnel_secret": schema.StringAttribute{
        Description: "Sets the password required to run a locally-managed tunnel. Must be at least 32 bytes and encoded as a base64 string.",
        Optional: true,
        Sensitive: true,
      },
      "account_tag": schema.StringAttribute{
        Description: "Cloudflare account ID",
        Computed: true,
      },
      "conns_active_at": schema.StringAttribute{
        Description: "Timestamp of when the tunnel established at least one connection to Cloudflare's edge. If `null`, the tunnel is inactive.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "conns_inactive_at": schema.StringAttribute{
        Description: "Timestamp of when the tunnel became inactive (no connections to Cloudflare's edge). If `null`, the tunnel is active.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "created_at": schema.StringAttribute{
        Description: "Timestamp of when the resource was created.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "deleted_at": schema.StringAttribute{
        Description: "Timestamp of when the resource was deleted. If `null`, the resource has not been deleted.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "remote_config": schema.BoolAttribute{
        Description: "If `true`, the tunnel can be configured remotely from the Zero Trust dashboard. If `false`, the tunnel must be configured locally on the origin machine.",
        Computed: true,
      },
      "status": schema.StringAttribute{
        Description: "The status of the tunnel. Valid values are `inactive` (tunnel has never been run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy state), `healthy` (tunnel is active and able to serve traffic), or `down` (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).\nAvailable values: \"inactive\", \"degraded\", \"healthy\", \"down\".",
        Computed: true,
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
        Computed: true,
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
      "connections": schema.ListNestedAttribute{
        Description: "The Cloudflare Tunnel connections between your origin and Cloudflare's edge.",
        Computed: true,
        CustomType: customfield.NewNestedObjectListType[ZeroTrustTunnelCloudflaredConnectionsModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Description: "UUID of the Cloudflare Tunnel connection.",
              Computed: true,
            },
            "client_id": schema.StringAttribute{
              Description: "UUID of the Cloudflare Tunnel connector.",
              Computed: true,
            },
            "client_version": schema.StringAttribute{
              Description: "The cloudflared version used to establish this connection.",
              Computed: true,
            },
            "colo_name": schema.StringAttribute{
              Description: "The Cloudflare data center used for this connection.",
              Computed: true,
            },
            "is_pending_reconnect": schema.BoolAttribute{
              Description: "Cloudflare continues to track connections for several minutes after they disconnect. This is an optimization to improve latency and reliability of reconnecting.  If `true`, the connection has disconnected but is still being tracked. If `false`, the connection is actively serving traffic.",
              Computed: true,
            },
            "opened_at": schema.StringAttribute{
              Description: "Timestamp of when the connection was established.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "origin_ip": schema.StringAttribute{
              Description: "The public IP address of the host running cloudflared.",
              Computed: true,
            },
            "uuid": schema.StringAttribute{
              Description: "UUID of the Cloudflare Tunnel connection.",
              Computed: true,
            },
          },
        },
      },
      "metadata": schema.StringAttribute{
        Description: "Metadata associated with the tunnel.",
        Computed: true,
        CustomType: jsontypes.NormalizedType{

        },
      },
    },
  }
}

func (r *ZeroTrustTunnelCloudflaredResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustTunnelCloudflaredResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
