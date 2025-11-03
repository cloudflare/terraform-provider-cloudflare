// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation_config

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*TokenValidationConfigDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "UUID.",
				Computed:    true,
			},
			"config_id": schema.StringAttribute{
				Description: "UUID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"title": schema.StringAttribute{
				Computed: true,
			},
			"token_type": schema.StringAttribute{
				Description: `Available values: "JWT".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("JWT"),
				},
			},
			"token_sources": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"credentials": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[TokenValidationConfigCredentialsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"keys": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[TokenValidationConfigCredentialsKeysDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"alg": schema.StringAttribute{
									Description: "Algorithm\nAvailable values: \"RS256\", \"RS384\", \"RS512\", \"PS256\", \"PS384\", \"PS512\", \"ES256\", \"ES384\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"RS256",
											"RS384",
											"RS512",
											"PS256",
											"PS384",
											"PS512",
											"ES256",
											"ES384",
										),
									},
								},
								"e": schema.StringAttribute{
									Description: "RSA exponent",
									Computed:    true,
								},
								"kid": schema.StringAttribute{
									Description: "Key ID",
									Computed:    true,
								},
								"kty": schema.StringAttribute{
									Description: "Key Type\nAvailable values: \"RSA\", \"EC\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("RSA", "EC"),
									},
								},
								"n": schema.StringAttribute{
									Description: "RSA modulus",
									Computed:    true,
								},
								"crv": schema.StringAttribute{
									Description: "Curve\nAvailable values: \"P-256\", \"P-384\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("P-256", "P-384"),
									},
								},
								"x": schema.StringAttribute{
									Description: "X EC coordinate",
									Computed:    true,
								},
								"y": schema.StringAttribute{
									Description: "Y EC coordinate",
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *TokenValidationConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *TokenValidationConfigDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
