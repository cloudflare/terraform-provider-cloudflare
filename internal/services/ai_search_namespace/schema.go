// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_namespace

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*AISearchNamespaceResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description: "Optional description for the namespace. Max 256 characters.",
				Optional:    true,
			},
			"public_endpoint_params": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[AISearchNamespacePublicEndpointParamsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"authorized_hosts": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"chat_completions_endpoint": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[AISearchNamespacePublicEndpointParamsChatCompletionsEndpointModel](ctx),
						Attributes: map[string]schema.Attribute{
							"disabled": schema.BoolAttribute{
								Description: "Disable chat completions endpoint for this public endpoint",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
						},
					},
					"custom_domains": schema.ListAttribute{
						Description: "Custom domain hostnames that alias this public endpoint. GET and create responses return the current set; on update (PUT) this field is only echoed back when supplied in the request body, otherwise it is null (omit it to leave domains unchanged).",
						Optional:    true,
						ElementType: types.StringType,
					},
					"default_domain_enabled": schema.BoolAttribute{
						Description: "When false, the instance is reachable only via a registered custom domain and the default <public_endpoint_id>.search.ai.cloudflare.com host returns 404. Requires at least one custom domain. Defaults to true. public_endpoint_params is replaced wholesale on update, so resend default_domain_enabled on every update to keep the default host off — omitting it resets to true.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(true),
					},
					"enabled": schema.BoolAttribute{
						Computed: true,
						Optional: true,
						Default:  booldefault.StaticBool(false),
					},
					"instances_allowed": schema.ListAttribute{
						Description: "Instance IDs exposed through the namespace public endpoint. Empty means nothing is searchable. Every ID must be an existing instance in this namespace, and the list cannot exceed the account's multi-instance search limit.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"mcp": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[AISearchNamespacePublicEndpointParamsMcpModel](ctx),
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								Computed: true,
								Optional: true,
								Default:  stringdefault.StaticString("Finds exactly what you're looking for"),
							},
							"disabled": schema.BoolAttribute{
								Description: "Disable MCP endpoint for this public endpoint",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
						},
					},
					"rate_limit": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"period_ms": schema.Int64Attribute{
								Optional: true,
								Validators: []validator.Int64{
									int64validator.Between(60000, 3600000),
								},
							},
							"requests": schema.Int64Attribute{
								Optional: true,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},
							"technique": schema.StringAttribute{
								Description: `Available values: "fixed", "sliding".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("fixed", "sliding"),
								},
							},
						},
					},
					"search_endpoint": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[AISearchNamespacePublicEndpointParamsSearchEndpointModel](ctx),
						Attributes: map[string]schema.Attribute{
							"disabled": schema.BoolAttribute{
								Description: "Disable search endpoint for this public endpoint",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"public_endpoint_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *AISearchNamespaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AISearchNamespaceResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
