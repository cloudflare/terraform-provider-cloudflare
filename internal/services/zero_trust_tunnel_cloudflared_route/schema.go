// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_route

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustTunnelCloudflaredRouteResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "UUID of the route.",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "account_id": schema.StringAttribute{
        Description: "Cloudflare account ID",
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "network": schema.StringAttribute{
        Description: "The private IPv4 or IPv6 range connected by the route, in CIDR notation.",
        Required: true,
      },
      "tunnel_id": schema.StringAttribute{
        Description: "UUID of the tunnel.",
        Required: true,
      },
      "comment": schema.StringAttribute{
        Description: "Optional remark describing the route.",
        Optional: true,
      },
      "virtual_network_id": schema.StringAttribute{
        Description: "UUID of the virtual network.",
        Optional: true,
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
    },
  }
}

func (r *ZeroTrustTunnelCloudflaredRouteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustTunnelCloudflaredRouteResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
