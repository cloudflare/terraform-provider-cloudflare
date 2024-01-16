package r2_bucket

import (
	"context"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r *R2BucketResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
			The [R2 Bucket](https://developers.cloudflare.com/r2/) resource allows you to manage Cloudflare R2 buckets.
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
				MarkdownDescription: "The name of the R2 bucket.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"location": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The location hint of the R2 bucket.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}
