package hyperdrive_config

import (
	"context"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r *HyperdriveConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
		The [Hyperdrive Config](https://developers.cloudflare.com/hyperdrive/) resource allows you to manage Cloudflare Hyperdrive Configs.
`),

		Attributes: map[string]schema.Attribute{
			consts.IDSchemaKey: schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: consts.IDSchemaDescription + " This is the hyperdrive config value.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Hyperdrive configuration.",
				Required:            true,
			},
			"origin": schema.SingleNestedAttribute{
				MarkdownDescription: "The origin details for the Hyperdrive configuration.",
				Required:            true,
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						MarkdownDescription: "The name of your origin database.",
						Required:            true,
					},
					"password": schema.StringAttribute{
						MarkdownDescription: "The password of the Hyperdrive configuration.",
						Required:            true,
						Sensitive:           true,
					},
					"host": schema.StringAttribute{
						MarkdownDescription: "The host (hostname or IP) of your origin database.",
						Required:            true,
					},
					"port": schema.Int64Attribute{
						MarkdownDescription: "The port (default: 5432 for Postgres) of your origin database.",
						Optional:            true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(65535),
						},
					},
					"scheme": schema.StringAttribute{
						MarkdownDescription: "Specifies the URL scheme used to connect to your origin database.",
						Required:            true,
					},
					"user": schema.StringAttribute{
						MarkdownDescription: "The user of your origin database.",
						Required:            true,
					},
					"access_client_id": schema.StringAttribute{
						MarkdownDescription: "Client ID associated with the Cloudflare Access Service Token used to connect via Access.",
						Optional:            true,
						Validators: []validator.String{
							stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("access_client_secret")),
							stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("port")),
						},
					},
					"access_client_secret": schema.StringAttribute{
						MarkdownDescription: "Client Secret associated with the Cloudflare Access Service Token used to connect via Access.",
						Optional:            true,
						Validators: []validator.String{
							stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("access_client_id")),
							stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("port")),
						},
					},
				},
			},
			"caching": schema.SingleNestedAttribute{
				MarkdownDescription: "The caching details for the Hyperdrive configuration.",
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"disabled": schema.BoolAttribute{
						MarkdownDescription: "Disable caching for this Hyperdrive configuration.",
						Optional:            true,
						Computed:            true,
					},
					"max_age": schema.Int64Attribute{
						MarkdownDescription: "Configure the `max_age` value of this Hyperdrive configuration.",
						Optional:            true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(3600),
						},
					},
					"stale_while_revalidate": schema.Int64Attribute{
						MarkdownDescription: "Disable caching for this Hyperdrive configuration.",
						Optional:            true,
						Validators: []validator.Int64{
							int64validator.AlsoRequires(path.MatchRelative().AtParent().AtName("max_age")),
							int64validator.AtLeast(0),
							int64validator.AtMost(3600), // Max value is the value of max_age, API will reject if not
						},
					},
				},
			},
		},
	}
}
