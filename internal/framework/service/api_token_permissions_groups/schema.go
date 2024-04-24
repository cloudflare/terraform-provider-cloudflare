package api_token_permissions_groups

import (
	"context"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *APITokenPermissionsGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		Description: heredoc.Docf(`
			Use this data source to look up [API Token Permission Groups](https://developers.cloudflare.com/api/tokens/create/permissions).
			Commonly used as references within [%s](/docs/providers/cloudflare/r/api_token.html) resources.
		`, "`cloudflare_token`"),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Checksum of permissions.",
			},
			"permissions": schema.MapAttribute{
				Computed:           true,
				DeprecationMessage: "Use specific account, zone or user attributes instead.",
				Description:        "Map of all permissions available. Should not be used as some permissions will overlap resource scope. Instead, use resource level specific attributes.",
				ElementType:        types.StringType,
			},
			"zone": schema.MapAttribute{
				Computed:    true,
				Description: "Map of permissions for zone level resources.",
				ElementType: types.StringType,
			},
			"account": schema.MapAttribute{
				Computed:    true,
				Description: "Map of permissions for account level resources.",
				ElementType: types.StringType,
			},
			"user": schema.MapAttribute{
				Computed:    true,
				Description: "Map of permissions for user level resources.",
				ElementType: types.StringType,
			},
			"r2": schema.MapAttribute{
				Computed:    true,
				Description: "Map of permissions for r2 level resources.",
				ElementType: types.StringType,
			},
		},
	}
}
