// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*APIShieldDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account API Gateway",
				"Account API Gateway Read",
				"Domain API Gateway",
				"Domain API Gateway Read",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier.",
				Computed:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Optional:    true,
			},
			"normalize": schema.BoolAttribute{
				Description: "Ensures that the configuration is written or retrieved in normalized fashion",
				Optional:    true,
			},
			"auth_id_characteristics": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[APIShieldAuthIDCharacteristicsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "For `type: header` or `type: cookie`: The name of the characteristic field, i.e., the header or cookie name.\n\nFor `type: jwt`: Claim location expressed as `$(token_config_id):$(json_path)`, where `token_config_id` \nis the ID of the token configuration used in validating the JWT, and `json_path` is a RFC 9535 \nJSONPath (https://goessner.net/articles/JsonPath/, https://www.rfc-editor.org/rfc/rfc9535.html).\nThe JSONPath expression may be in dot or bracket notation, may only specify literal keys\nor array indexes, and must return a singleton value, which will be interpreted as a string.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "The type of characteristic.\nAvailable values: \"header\", \"cookie\", \"jwt\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"header",
									"cookie",
									"jwt",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (d *APIShieldDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *APIShieldDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
