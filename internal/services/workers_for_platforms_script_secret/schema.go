// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_script_secret

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*WorkersForPlatformsScriptSecretResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "A JavaScript variable name for the binding.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "A JavaScript variable name for the binding.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"dispatch_namespace": schema.StringAttribute{
				Description:   "Name of the Workers for Platforms dispatch namespace.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"script_name": schema.StringAttribute{
				Description:   "Name of the script, used in URLs and route configuration.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Description: "The kind of resource that the binding provides.\nAvailable values: \"secret_text\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("secret_text", "secret_key"),
				},
			},
			"format": schema.StringAttribute{
				Description: "Data format of the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).\nAvailable values: \"raw\", \"pkcs8\", \"spki\", \"jwk\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"raw",
						"pkcs8",
						"spki",
						"jwk",
					),
				},
			},
			"key_base64": schema.StringAttribute{
				Description: "Base64-encoded key data. Required if `format` is \"raw\", \"pkcs8\", or \"spki\".",
				Optional:    true,
				Sensitive:   true,
			},
			"text": schema.StringAttribute{
				Description: "The secret value to use.",
				Optional:    true,
				Sensitive:   true,
			},
			"usages": schema.ListAttribute{
				Description: "Allowed operations with the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).",
				Optional:    true,
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
				ElementType: types.StringType,
			},
			"algorithm": schema.StringAttribute{
				Description: "Algorithm-specific key parameters. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).",
				Optional:    true,
				CustomType:  jsontypes.NormalizedType{},
			},
			"key_jwk": schema.StringAttribute{
				Description: "Key data in [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key) format. Required if `format` is \"jwk\".",
				Optional:    true,
				Sensitive:   true,
				CustomType:  jsontypes.NormalizedType{},
			},
		},
	}
}

func (r *WorkersForPlatformsScriptSecretResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WorkersForPlatformsScriptSecretResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
