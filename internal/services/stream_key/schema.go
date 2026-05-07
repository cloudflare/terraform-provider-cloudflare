// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_key

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*StreamKeyResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Stream Read",
				"Stream Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"created": schema.StringAttribute{
				Description: "The date and time a signing key was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"jwk": schema.StringAttribute{
				Description: "The signing key in JWK format.",
				Computed:    true,
				Sensitive:   true,
			},
			"key_id": schema.StringAttribute{
				Description: "The unique identifier for the signing key.",
				Computed:    true,
			},
			"pem": schema.StringAttribute{
				Description: "The signing key in PEM format.",
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func (r *StreamKeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *StreamKeyResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
