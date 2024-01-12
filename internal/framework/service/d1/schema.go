package d1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r *DatabaseResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
			The [D1 Database](https://developers.cloudflare.com/d1/) resource allows you to manage Cloudflare D1 databases.
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
				MarkdownDescription: "The name of the D1 Database.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-z0-9][a-z0-9-_]*$`),
						"must contain only lowercase alphanumeric characters",
					),
				},
			},
			"version": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The backend version of the database.",
			},
		},
	}
}
