// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_script_secret

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersForPlatformsScriptSecretsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"dispatch_namespace": schema.StringAttribute{
				Description: "Name of the Workers for Platforms dispatch namespace.",
				Required:    true,
			},
			"script_name": schema.StringAttribute{
				Description: "Name of the script, used in URLs and route configuration.",
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
				CustomType:  customfield.NewNestedObjectListType[WorkersForPlatformsScriptSecretsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "A JavaScript variable name for the binding.",
							Computed:    true,
						},
						"text": schema.StringAttribute{
							Description: "The secret value to use.",
							Computed:    true,
							Sensitive:   true,
						},
						"type": schema.StringAttribute{
							Description: "The kind of resource that the binding provides.\nAvailable values: \"secret_text\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("secret_text", "secret_key"),
							},
						},
						"algorithm": schema.StringAttribute{
							Description: "Algorithm-specific key parameters ([learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm)).",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"format": schema.StringAttribute{
							Description: "Data format of the key ([learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format)).\nAvailable values: \"raw\", \"pkcs8\", \"spki\", \"jwk\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"raw",
									"pkcs8",
									"spki",
									"jwk",
								),
							},
						},
						"usages": schema.ListAttribute{
							Description: "Allowed operations with the key ([learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages)).",
							Computed:    true,
							Validators: []validator.List{
								listvalidator.ValueStringsAre(
									stringvalidator.OneOfCaseInsensitive(
										"encrypt",
										"decrypt",
										"sign",
										"verify",
										"deriveKey",
										"deriveBits",
										"wrapKey",
										"unwrapKey",
									),
								),
							},
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"key_base64": schema.StringAttribute{
							Description: "Base64-encoded key data. Required if `format` is \"raw\", \"pkcs8\", or \"spki\".",
							Computed:    true,
							Sensitive:   true,
						},
						"key_jwk": schema.StringAttribute{
							Description: "Key data in [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key) format. Required if `format` is \"jwk\".",
							Computed:    true,
							Sensitive:   true,
							CustomType:  jsontypes.NormalizedType{},
						},
					},
				},
			},
		},
	}
}

func (d *WorkersForPlatformsScriptSecretsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *WorkersForPlatformsScriptSecretsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
