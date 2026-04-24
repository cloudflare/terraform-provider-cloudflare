// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*ManagedTransformsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account Rulesets Read",
				"Account Rulesets Write",
				"Account WAF Read",
				"Account WAF Write",
				"Bot Management Read",
				"Bot Management Write",
				"Cache Settings Read",
				"Cache Settings Write",
				"Config Settings Read",
				"Config Settings Write",
				"Custom Errors Read",
				"Custom Errors Write",
				"Dynamic URL Redirects Read",
				"Dynamic URL Redirects Write",
				"HTTP DDoS Managed Ruleset Read",
				"HTTP DDoS Managed Ruleset Write",
				"L4 DDoS Managed Ruleset Read",
				"L4 DDoS Managed Ruleset Write",
				"Logs Read",
				"Logs Write",
				"Magic Firewall Read",
				"Magic Firewall Write",
				"Managed headers Read",
				"Managed headers Write",
				"Mass URL Redirects Read",
				"Mass URL Redirects Write",
				"Origin Read",
				"Origin Write",
				"Response Compression Read",
				"Response Compression Write",
				"Sanitize Read",
				"Sanitize Write",
				"Select Configuration Read",
				"Select Configuration Write",
				"Transform Rules Read",
				"Transform Rules Write",
				"Zone Transform Rules Read",
				"Zone Transform Rules Write",
				"Zone WAF Read",
				"Zone WAF Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The unique ID of the zone.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "The unique ID of the zone.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"managed_request_headers": schema.SetNestedAttribute{
				Description: "The list of Managed Request Transforms.",
				Required:    true,
				CustomType:  customfield.NewNestedObjectSetType[ManagedTransformsManagedRequestHeadersModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The human-readable identifier of the Managed Transform.",
							Required:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the Managed Transform is enabled.",
							Required:    true,
						},
					},
				},
			},
			"managed_response_headers": schema.SetNestedAttribute{
				Description: "The list of Managed Response Transforms.",
				Required:    true,
				CustomType:  customfield.NewNestedObjectSetType[ManagedTransformsManagedResponseHeadersModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The human-readable identifier of the Managed Transform.",
							Required:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the Managed Transform is enabled.",
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func (r *ManagedTransformsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ManagedTransformsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
