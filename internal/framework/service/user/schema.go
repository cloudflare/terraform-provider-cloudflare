package user

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func (r *CloudflareUserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Use this data source to retrieve information about the currently authenticated user.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The user's unique identifier.",
				Computed:    true,
			},
			"email": schema.StringAttribute{
				Description: "The user's email address.",
				Computed:    true,
			},
			"username": schema.StringAttribute{
				Description: "The user's username.",
				Computed:    true,
			},
		},
	}
}
