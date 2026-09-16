// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_custom_page

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustAccessCustomPageResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Access: Custom Pages Read",
				"Access: Custom Pages Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"uid": schema.StringAttribute{
				Description:   "UUID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"custom_html": schema.StringAttribute{
				Description: "Custom page HTML.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Custom page name.",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "Custom page type.\nAvailable values: \"identity_denied\", \"forbidden\", \"login\", \"interstitial\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"identity_denied",
						"forbidden",
						"login",
						"interstitial",
					),
				},
			},
			"contract_version": schema.Int64Attribute{
				Description: "Contract version of the page's Liquid template. Present (>= 1) marks a sanitized template; absent or 0 marks a legacy page served verbatim.",
				Optional:    true,
				Computed:    true,
			},
			"warnings": schema.ListNestedAttribute{
				Description: "Advisory validation findings returned when creating or updating a template. Omitted when empty.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessCustomPageWarningsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"message": schema.StringAttribute{
							Description: "Human-readable description of the finding.",
							Computed:    true,
						},
						"tier": schema.StringAttribute{
							Description: "The validation tier that produced the finding (e.g. html, liquid).",
							Computed:    true,
						},
						"ref": schema.StringAttribute{
							Description: "Optional pointer to the part of the template the finding refers to.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *ZeroTrustAccessCustomPageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustAccessCustomPageResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
