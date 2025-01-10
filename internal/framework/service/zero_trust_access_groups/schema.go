package zero_trust_access_groups

import (
	"context"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func (r *ZeroTrustAccessGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		Description: heredoc.Docf(`
			Use this data source to look up [Zero Trust Access Groups](https://developers.cloudflare.com/cloudflare-one/identity/users/groups/).
			Commonly used as references within [%s](/docs/providers/cloudflare/r/zero_trust_access_policy.html) resources.
		`, "`cloudflare_zero_trust_access_policy`"),
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required:    true,
				Description: "Cloudflare Account ID",
			},
			"groups": schema.ListNestedAttribute{
				Computed:    true,
				Description: "A list of Zero Trust Access Groups.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "The identifier for this group.",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "The name of the group.",
						},
					},
				},
			},
		},
	}
}
