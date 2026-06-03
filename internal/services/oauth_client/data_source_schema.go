// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oauth_client

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*OAuthClientDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"OAuth Client Read",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Required:    true,
			},
			"oauth_client_id": schema.StringAttribute{
				Description: "The unique identifier for an OAuth client.",
				Required:    true,
			},
			"client_id": schema.StringAttribute{
				Description: "The unique identifier for an OAuth client.",
				Computed:    true,
			},
			"client_name": schema.StringAttribute{
				Description: "Human-readable name of the OAuth client.",
				Computed:    true,
			},
			"client_uri": schema.StringAttribute{
				Description: "URL of the home page of the client.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Timestamp when the OAuth client was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"has_rotated_secret": schema.BoolAttribute{
				Description: "Indicates whether the client has a rotated secret that has not yet been deleted.",
				Computed:    true,
			},
			"logo_uri": schema.StringAttribute{
				Description: "URL of the client's logo.",
				Computed:    true,
			},
			"policy_uri": schema.StringAttribute{
				Description: "URL that points to a privacy policy document.",
				Computed:    true,
			},
			"promoted_at": schema.StringAttribute{
				Description: "Timestamp when the OAuth client was promoted to public visibility.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"token_endpoint_auth_method": schema.StringAttribute{
				Description: "The authentication method the client uses at the token endpoint.\nAvailable values: \"none\", \"client_secret_basic\", \"client_secret_post\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"none",
						"client_secret_basic",
						"client_secret_post",
					),
				},
			},
			"tos_uri": schema.StringAttribute{
				Description: "URL that points to a terms of service document.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "Timestamp when the OAuth client was last updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"visibility": schema.StringAttribute{
				Description: "Visibility of the OAuth client.\nAvailable values: \"public\", \"private\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("public", "private"),
				},
			},
			"allowed_cors_origins": schema.ListAttribute{
				Description: "Array of allowed CORS origins.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"grant_types": schema.ListAttribute{
				Description: "Array of OAuth grant types the client is allowed to use. `authorization_code` is required; `refresh_token` may be included optionally.",
				Computed:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive("authorization_code", "refresh_token"),
					),
				},
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"post_logout_redirect_uris": schema.ListAttribute{
				Description: "Array of allowed post-logout redirect URIs.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"redirect_uris": schema.ListAttribute{
				Description: "Array of allowed redirect URIs for the client.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"response_types": schema.ListAttribute{
				Description: "Array of OAuth response types the client is allowed to use.",
				Computed:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"token",
							"id_token",
							"code",
						),
					),
				},
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"scopes": schema.ListAttribute{
				Description: "Array of OAuth scopes the client is allowed to request. Colon-delimited scopes are not accepted. Dot-delimited scopes are validated against available OAuth API scopes; simple identity scopes are allowed. Protocol scopes `offline_access` and `openid` are added or removed automatically based on `grant_types` and `response_types`.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"client_uri_verification": schema.SingleNestedAttribute{
				Description: "Client URI domain control verification state.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[OAuthClientClientURIVerificationDataSourceModel](ctx),
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

func (d *OAuthClientDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *OAuthClientDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
