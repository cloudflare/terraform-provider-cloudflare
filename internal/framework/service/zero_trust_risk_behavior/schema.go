package zero_trust_risk_behavior

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r *ZeroTrustRiskBehaviorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
			The [Risk Behavior](https://developers.cloudflare.com/cloudflare-one/insights/risk-score/) resource allows you to configure Cloudflare Risk Behaviors for an account.
	`),

		Attributes: map[string]schema.Attribute{
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},

		Blocks: map[string]schema.Block{

			"behavior": schema.SetNestedBlock{
				MarkdownDescription: "Zero Trust risk behaviors configured on this account",
				Validators: []validator.Set{
					setvalidator.SizeAtLeast(1),
					setvalidator.IsRequired(),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Name of this risk behavior type",
						},
						"enabled": schema.BoolAttribute{
							Required:            true,
							MarkdownDescription: "Whether this risk behavior type is enabled.",
						},
						"risk_level": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: fmt.Sprintf("Risk level. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{"low", "medium", "high"})),
							Validators: []validator.String{
								stringvalidator.OneOf("low", "medium", "high"),
							},
						},
					},
				},
			},
		},
	}
}
