// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship_app

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*FlagshipAppResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account Firewall Access Rules Read",
				"Account Firewall Access Rules Write",
				"Account Settings Read",
				"Account Settings Write",
				"Billing Read",
				"Billing Write",
				"DDoS Botnet Feed Read",
				"DDoS Botnet Feed Write",
				"DDoS Protection Read",
				"DDoS Protection Write",
				"DNS Firewall Read",
				"DNS Firewall Write",
				"DNS View Read",
				"DNS View Write",
				"Load Balancers Account Read",
				"Load Balancers Account Write",
				"Load Balancing: Monitors and Pools Read",
				"Load Balancing: Monitors and Pools Write",
				"SCIM Provisioning",
				"Trust and Safety Read",
				"Trust and Safety Write",
				"Workers KV Storage Read",
				"Workers KV Storage Write",
				"Workers R2 Storage Read",
				"Workers R2 Storage Write",
				"Workers Scripts Read",
				"Workers Scripts Write",
				"Workers Tail Read",
				"Zero Trust: PII Read",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Cloudflare account ID.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"updated_by": schema.StringAttribute{
				Description: "Email of the actor who last modified the app, or `edge-gateway` for gateway-authenticated changes.",
				Computed:    true,
			},
		},
	}
}

func (r *FlagshipAppResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *FlagshipAppResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
