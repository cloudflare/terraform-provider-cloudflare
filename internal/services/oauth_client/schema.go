// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oauth_client

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*OAuthClientResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"OAuth Client Read",
				"OAuth Client Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Account identifier tag.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"oauth_client_id": schema.StringAttribute{
				Description:   "The unique identifier for an OAuth client.",
				Optional:      true,
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"client_name": schema.StringAttribute{
				Description: "Human-readable name of the OAuth client.",
				Required:    true,
			},
			"token_endpoint_auth_method": schema.StringAttribute{
				Description: "The authentication method the client uses at the token endpoint.\nAvailable values: \"none\", \"client_secret_basic\", \"client_secret_post\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"none",
						"client_secret_basic",
						"client_secret_post",
					),
				},
			},
			"grant_types": schema.ListAttribute{
				Description: "Array of OAuth grant types the client is allowed to use. `authorization_code` is required; `refresh_token` may be included optionally.",
				Required:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive("authorization_code", "refresh_token"),
					),
				},
				ElementType: types.StringType,
			},
			"redirect_uris": schema.ListAttribute{
				Description: "Array of allowed redirect URIs for the client.",
				Required:    true,
				ElementType: types.StringType,
			},
			"response_types": schema.ListAttribute{
				Description: "Array of OAuth response types the client is allowed to use.",
				Required:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"token",
							"id_token",
							"code",
						),
					),
				},
				ElementType: types.StringType,
			},
			"scopes": schema.ListAttribute{
				Description: "Array of OAuth scopes the client is allowed to request. Colon-delimited scopes are not accepted. Dot-delimited scopes are validated against available OAuth API scopes; simple identity scopes are allowed. Protocol scopes `offline_access` and `openid` are added or removed automatically based on `grant_types` and `response_types`.",
				Required:    true,
				ElementType: types.StringType,
			},
			"client_uri": schema.StringAttribute{
				Description: "URL of the home page of the client.",
				Optional:    true,
			},
			"logo_uri": schema.StringAttribute{
				Description: "URL of the client's logo.",
				Optional:    true,
			},
			"policy_uri": schema.StringAttribute{
				Description: "URL that points to a privacy policy document.",
				Optional:    true,
			},
			"tos_uri": schema.StringAttribute{
				Description: "URL that points to a terms of service document.",
				Optional:    true,
			},
			"visibility": schema.StringAttribute{
				Description: "Promote the OAuth client from private to public visibility. Only `public` is accepted; demotion to `private` is not supported. Promotion requires a non-empty client name, logo URI, verified client URI host, and at least one non-identity scope.\nAvailable values: \"public\".",
				Optional:    true,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("public"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"allowed_cors_origins": schema.ListAttribute{
				Description: "Array of allowed CORS origins.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"post_logout_redirect_uris": schema.ListAttribute{
				Description: "Array of allowed post-logout redirect URIs.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"client_id": schema.StringAttribute{
				Description:   "The unique identifier for an OAuth client.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"client_secret": schema.StringAttribute{
				Description:   "The client secret. This is the only time the secret is returned in a response.",
				Computed:      true,
				Sensitive:     true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created_at": schema.StringAttribute{
				Description:   "Timestamp when the OAuth client was created.",
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"has_rotated_secret": schema.BoolAttribute{
				Description:   "Indicates whether the client has a rotated secret that has not yet been deleted.",
				Computed:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"promoted_at": schema.StringAttribute{
				Description:   "Timestamp when the OAuth client was promoted to public visibility.",
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"updated_at": schema.StringAttribute{
				Description: "Timestamp when the OAuth client was last updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"client_uri_verification": schema.SingleNestedAttribute{
				Description: "Client URI domain control verification state.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[OAuthClientClientURIVerificationModel](ctx),
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Description: "Current verification status for the client URI host.\nAvailable values: \"pending\", \"in_progress\", \"verified\", \"failed\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"pending",
								"in_progress",
								"verified",
								"failed",
							),
						},
					},
					"text": schema.StringAttribute{
						Description: "Exact TXT record value that must be added to DNS to prove ownership of the client URI host.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *OAuthClientResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *OAuthClientResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
