// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*TokenValidationConfigResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"token_type": schema.StringAttribute{
				Description: `Available values: "JWT".`,
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("JWT"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"credentials": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"keys": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"alg": schema.StringAttribute{
									Description: "Algorithm\nAvailable values: \"ES256\", \"ES384\", \"RS256\", \"RS384\", \"RS512\", \"PS256\", \"PS384\", \"PS512\".",
									Required:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"ES256",
											"ES384",
											"RS256",
											"RS384",
											"RS512",
											"PS256",
											"PS384",
											"PS512",
										),
									},
								},
								"kid": schema.StringAttribute{
									Description: "Key ID",
									Required:    true,
								},
								"kty": schema.StringAttribute{
									Description: "Key Type\nAvailable values: \"EC\", \"RSA\".",
									Required:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("EC", "RSA"),
									},
								},
								"x": schema.StringAttribute{
									Description: "X EC coordinate",
									Optional:    true,
								},
								"y": schema.StringAttribute{
									Description: "Y EC coordinate",
									Optional:    true,
								},
								"e": schema.StringAttribute{
									Description: "RSA exponent",
									Optional:    true,
								},
								"n": schema.StringAttribute{
									Description: "RSA modulus",
									Optional:    true,
								},
							},
						},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Required: true,
			},
			"title": schema.StringAttribute{
				Required: true,
			},
			"token_sources": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"last_updated": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *TokenValidationConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *TokenValidationConfigResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
