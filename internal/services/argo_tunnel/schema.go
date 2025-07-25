// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_tunnel

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ArgoTunnelResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID of the Argo Tunnel.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			consts.AccountIDSchemaKey: schema.StringAttribute{
				Description:   "The account identifier to target for the resource.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "A user-friendly name for a tunnel.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"secret": schema.StringAttribute{
				Description:   "Sets the password required to run a locally-managed tunnel. Must be at least 32 bytes and encoded as a base64 string.",
				Required:      true,
				Sensitive:     true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"cname": schema.StringAttribute{
				Description: "The CNAME record that points to your tunnel.",
				Computed:    true,
			},
			"tunnel_token": schema.StringAttribute{
				Description: "The tunnel token.",
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}
