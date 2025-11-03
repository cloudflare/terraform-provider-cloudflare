// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation_config

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*TokenValidationConfigsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[TokenValidationConfigsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "UUID.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"credentials": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[TokenValidationConfigsCredentialsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"keys": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[TokenValidationConfigsCredentialsKeysDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"alg": schema.StringAttribute{
												Description: "Algorithm\nAvailable values: \"ES256\", \"ES384\", \"RS256\", \"RS384\", \"RS512\", \"PS256\", \"PS384\", \"PS512\".",
												Computed:    true,
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
												Computed:    true,
											},
											"kty": schema.StringAttribute{
												Description: "Key Type\nAvailable values: \"EC\", \"RSA\".",
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("EC", "RSA"),
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
											"e": schema.StringAttribute{
												Description: "RSA exponent",
												Computed:    true,
											},
											"n": schema.StringAttribute{
												Description: "RSA modulus",
												Computed:    true,
											},
										},
									},
								},
							},
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
						"token_sources": schema.ListAttribute{
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"token_type": schema.StringAttribute{
							Description: `Available values: "JWT".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("JWT"),
							},
						},
					},
				},
			},
		},
	}
}

func (d *TokenValidationConfigsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *TokenValidationConfigsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
