package workers_for_platforms

import (
	"context"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r *WorkersForPlatformsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
			The [Workers for Platforms](https://developers.cloudflare.com/cloudflare-for-platforms/workers-for-platforms/) resource allows you 
			to manage Cloudflare Workers for Platforms namespaces.
	`),

		Attributes: map[string]schema.Attribute{
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			consts.IDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.IDSchemaDescription,
				Computed:            true,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the Workers for Platforms namespace.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}
