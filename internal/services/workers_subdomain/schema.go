// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_subdomain

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*WorkersSubdomainResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Workers Admin",
				"Workers Editor",
				"Workers Scripts Write",
			},
		}.String(),
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"subdomain": schema.StringAttribute{
				Description: "The account's workers.dev subdomain, as in `<subdomain>.workers.dev`.",
				Required:    true,
			},
		},
	}
}

func (r *WorkersSubdomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WorkersSubdomainResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
